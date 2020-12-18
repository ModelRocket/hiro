/*
 * This file is part of the Model Rocket Hiro Stack
 * Copyright (c) 2020 Model Rocket LLC.
 *
 * https://github.com/ModelRocket/hiro
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package hiro

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/ModelRocket/hiro/pkg/api"
	"github.com/ModelRocket/hiro/pkg/api/session"
	"github.com/ModelRocket/hiro/pkg/env"
	"github.com/ModelRocket/hiro/pkg/hiro/pb"
	"github.com/ModelRocket/hiro/pkg/oauth"
	"github.com/apex/log"
	"github.com/go-co-op/gocron"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type (
	// Daemon is the core hiro service object that all hiro base platforms use
	Daemon struct {
		name        string
		apiServer   *api.Server
		apiOptions  []api.Option
		rpcServer   *grpc.Server
		webRPCPath  string
		backOptions []BackendOption
		ctrl        Controller
		oauthPath   string
		oauthCtrl   oauth.Controller
		sessionCtrl session.Controller
		sessionMgr  *session.Manager
		serverAddr  string
		shutdown    chan int
		wg          sync.WaitGroup
		sched       *gocron.Scheduler
	}

	// Job is a job handler that the daemon will schedule
	Job struct {
		Function interface{}
		Params   []interface{}
		Interval time.Duration
		At       *time.Time
	}

	// DaemonOption is a daemon option
	DaemonOption func(d *Daemon)
)

// NewDaemon creates a new daemon object
func NewDaemon(opts ...DaemonOption) (*Daemon, error) {
	const (
		defaultServerAddr = "127.0.0.0:9000"
		defaultOAuthPath  = "/oauth"
		defaultWebRPCPath = "/rpc"
	)

	d := &Daemon{
		serverAddr:  defaultServerAddr,
		apiOptions:  []api.Option{api.WithLog(log.Log)},
		backOptions: make([]BackendOption, 0),
		oauthPath:   defaultOAuthPath,
		webRPCPath:  defaultWebRPCPath,
		shutdown:    make(chan int),
		sched:       gocron.NewScheduler(time.UTC),
	}

	for _, opt := range opts {
		opt(d)
	}

	if d.ctrl == nil {
		back, err := New(d.backOptions...)
		if err != nil {
			return nil, err
		}

		d.ctrl = back
	}

	if d.oauthCtrl == nil {
		if back, ok := d.ctrl.(*Backend); ok {
			d.oauthCtrl = back.OAuthController()
		} else {
			return nil, errors.New("OAuthController is required")
		}
	}

	// start the token cleanup job
	d.sched.Every(uint64(env.Duration("TOKEN_CLEANUP_INTERVAL", time.Minute*15).Minutes())).
		Minutes().
		StartImmediately().
		Do(d.oauthCtrl.TokenCleanup, context.Background())

	if d.sessionCtrl == nil {
		if back, ok := d.ctrl.(*Backend); ok {
			d.sessionCtrl = back.SessionController()
		} else {
			return nil, errors.New("SessionController is required")
		}
	}

	// start the session cleanup job
	d.sched.Every(uint64(env.Duration("SESSION_CLEANUP_INTERVAL", time.Minute*15).Minutes())).
		Minutes().
		StartImmediately().
		Do(d.sessionCtrl.SessionCleanup, context.Background())

	if d.apiServer == nil {
		d.apiServer = api.NewServer(d.apiOptions...)
	}

	d.sessionMgr = session.NewManager(d.sessionCtrl)

	// setup the oauth router
	d.apiServer.Router(d.oauthPath).
		WithSessionManager(d.sessionMgr).
		WithRoutes(oauth.Routes(d.oauthCtrl)...)

	if d.rpcServer == nil {
		d.rpcServer = grpc.NewServer()
	}

	// register the hiro rpc server
	pb.RegisterHiroServer(d.rpcServer, NewRPCServer(d.ctrl))

	// add grpc-web support
	ws := grpcweb.WrapServer(d.rpcServer)

	d.apiServer.Router(d.webRPCPath).
		Headers("Content-Type", "application/grpc-web").
		Handler(ws)

	d.sched.StartAsync()

	return d, nil
}

// WithName sets the daemon name
func WithName(name string) DaemonOption {
	return func(d *Daemon) {
		d.name = name
	}
}

// WithServerAddr sets the daemon listening address
func WithServerAddr(addr string) DaemonOption {
	return func(d *Daemon) {
		d.serverAddr = addr
	}
}

// WithBackendOptions sets backend options
func WithBackendOptions(o []BackendOption) DaemonOption {
	return func(d *Daemon) {
		d.backOptions = o
	}
}

// WithController sets the daemon controller
func WithController(c Controller) DaemonOption {
	return func(d *Daemon) {
		d.ctrl = c
	}
}

// WithOAuthController set the daemon oauth controller
func WithOAuthController(o oauth.Controller) DaemonOption {
	return func(d *Daemon) {
		d.oauthCtrl = o
	}
}

// WithSessionController set the daemon session controller
func WithSessionController(s session.Controller) DaemonOption {
	return func(d *Daemon) {
		d.sessionCtrl = s
	}
}

// WithAPIServer sets the daemon api server; mutally exclusive with WithAPIOptions
func WithAPIServer(s *api.Server) DaemonOption {
	return func(d *Daemon) {
		d.apiServer = s
	}
}

// WithAPIOptions sets api server options; mutally exclusive with WithAPIServer
func WithAPIOptions(o ...api.Option) DaemonOption {
	return func(d *Daemon) {
		d.apiOptions = o
	}
}

// WithRPCServer sets the daemon rpc server
func WithRPCServer(s *grpc.Server) DaemonOption {
	return func(d *Daemon) {
		d.rpcServer = s
	}
}

// Serve starts the dameon server
func (d *Daemon) Serve(ready func()) error {
	l, err := net.Listen("tcp", d.serverAddr)
	if err != nil {
		return err
	}

	// add a shutdown handler
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()

		for {
			select {
			case <-d.shutdown:
				l.Close()
				return
			}
		}
	}()

	// create the multiplexer
	m := cmux.New(l)

	// use an error group to collect errors
	errs := new(errgroup.Group)

	// start the http server
	errs.Go(func() error {
		l := m.Match(cmux.HTTP1Fast())
		s := &http.Server{Handler: d.apiServer}

		return s.Serve(l)
	})

	// start the rpc server
	errs.Go(func() error {
		l := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))

		return d.rpcServer.Serve(l)
	})

	// start the mux
	errs.Go(m.Serve)

	ready()

	return errs.Wait()
}

// Shutdown terminates the daemon services
func (d *Daemon) Shutdown(ctx context.Context) error {
	done := make(chan bool)

	d.sched.Stop()

	go func() {
		close(d.shutdown)
		d.wg.Wait()
		done <- true
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-done:
			return nil
		}
	}
}

// AddJob adds a job to the daemon scheduler
func (d *Daemon) AddJob(job Job) error {
	j := d.sched.Every(uint64(job.Interval.Seconds())).Seconds()

	if job.At != nil {
		j = j.At(job.At.UTC().Format("15:04:05"))
	}

	j.Do(job.Function, job.Params...)

	return nil
}
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

package oauth

import (
	"context"
	"time"

	"github.com/ModelRocket/hiro/pkg/api"
	"github.com/ModelRocket/hiro/pkg/safe"
	"github.com/ModelRocket/hiro/pkg/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	// AuthorizeParams contains all the bound params for the authorize operation
	AuthorizeParams struct {
		AppURI              URI                  `json:"app_uri"`
		Audience            string               `json:"audience"`
		ClientID            string               `json:"client_id"`
		CodeChallenge       string               `json:"code_challenge"`
		CodeChallengeMethod *CodeChallengeMethod `json:"code_challenge_method"`
		RedirectURI         URI                  `json:"redirect_uri"`
		ResponseType        string               `json:"response_type"`
		Scope               Scope                `json:"scope"`
		State               *string              `json:"state"`
	}
)

var (
	// DefaultCodeChallengeMethod is the only challenge method
	DefaultCodeChallengeMethod = "S256"
)

const (
	// RequestTokenParam is the name of the request token parameter passed on redirect from /authorize
	RequestTokenParam = "request_token"
)

// Validate validates the params
func (p AuthorizeParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.AppURI, validation.Required, is.RequestURI),
		validation.Field(&p.Audience, validation.Required),
		validation.Field(&p.ClientID, validation.Required),
		validation.Field(&p.CodeChallenge, validation.Required),
		validation.Field(&p.CodeChallengeMethod, validation.NilOrNotEmpty),
		validation.Field(&p.RedirectURI, validation.Required, is.RequestURI),
		validation.Field(&p.ResponseType, validation.Required),
		validation.Field(&p.Scope, validation.NilOrNotEmpty),
	)
}

func authorize(ctx context.Context, params *AuthorizeParams) api.Responder {
	ctrl := api.Context(ctx).(Controller)
	log := api.Log(ctx).WithField("operation", "authorize")

	// ensure this is a valid client
	client, err := ctrl.ClientGet(ctx, params.ClientID)
	if err != nil {
		log.Error(err.Error())

		return ErrAccessDenied.WithError(err)
	}

	// authorize this client for the grant, uris, and scope
	if err := client.ClientAuthorize(
		ctx,
		params.Audience,
		GrantTypeAuthCode,
		[]URI{params.AppURI, params.RedirectURI},
		params.Scope,
	); err != nil {
		log.Error(err.Error())

		return ErrAccessDenied.WithError(err)
	}

	// parse the api uri
	u, err := params.AppURI.Parse()
	if err != nil {
		log.Error(err.Error())

		return ErrAccessDenied.WithError(err)
	}

	// create a new request
	token, err := ctrl.RequestCreate(ctx, Request{
		AudienceID:          types.ID(params.Audience),
		ApplicationID:       types.ID(params.ClientID),
		ExpiresAt:           time.Now().Add(time.Minute * 10),
		Scope:               params.Scope,
		CodeChallenge:       params.CodeChallenge,
		CodeChallengeMethod: CodeChallengeMethod(safe.String(params.CodeChallengeMethod, CodeChallengeMethodS256)),
		AppURI:              params.AppURI,
		RedirectURI:         params.RedirectURI,
		State:               params.State,
	})
	if err != nil {
		log.Error(err.Error())

		return api.Redirect(u, map[string]string{
			"error":             "server_error",
			"error_description": err.Error(),
		})
	}
	log.Debugf("request token %s created", token)

	q := u.Query()
	q.Set(RequestTokenParam, token)

	u.RawQuery = q.Encode()

	return api.Redirect(u)
}

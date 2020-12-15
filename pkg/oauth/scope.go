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
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type (
	// Scope is an oauth scope
	Scope []string

	// ScopeSet represents a map between an audiece and a scope
	ScopeSet map[string]Scope
)

const (
	// ScopeOpenID is the openid scope
	ScopeOpenID = "openid"

	// ScopeProfile is the scope required to query for a users profile
	ScopeProfile = "profile"

	// ScopeOfflineAccess is the scope necessary to request a refresh_token
	ScopeOfflineAccess = "offline_access"

	// ScopeVerifyEmail is the scope required to verify a user's email address
	ScopeVerifyEmail = "verify:email"

	// ScopeVerifyPhone is the scope required to verify a user's phone number
	ScopeVerifyPhone = "verify:phone"
)

var (
	// Scopes is the list of all oauth scopes
	Scopes = Scope{
		ScopeOpenID,
		ScopeProfile,
		ScopeOfflineAccess,
		ScopeVerifyEmail,
		ScopeVerifyPhone,
	}
)

// MakeScope returns a Scope from the string scopes
func MakeScope(scopes ...string) Scope {
	return Scope(scopes)
}

// Contains return true if the scope contains the value
func (s Scope) Contains(value string) bool {
	for _, v := range s {
		if string(v) == value {
			return true
		}
	}

	return false
}

// Every returns true if every element is contained in the scope
func (s Scope) Every(elements ...string) bool {
	for _, elem := range elements {
		if !s.Contains(elem) {
			return false
		}
	}
	return true
}

// Some returns true if at least one of the elements is contained in the scope
func (s Scope) Some(elements ...string) bool {
	for _, elem := range elements {
		if s.Contains(elem) {
			return true
		}
	}
	return false
}

// Without returns the scope excluding the elements
func (s Scope) Without(elements ...string) Scope {
	r := make(Scope, 0)
	for _, v := range s {
		if !Scope(elements).Contains(v) {
			r = append(r, v)
		}
	}

	return r
}

// Unique returns a scope with only unique values
func (s Scope) Unique() Scope {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UnmarshalText handles text unmarshaling
func (s *Scope) UnmarshalText(v []byte) error {
	input := string(v)
	if strings.Contains(input, ",") {
		*s = strings.Split(input, ",")
	} else {
		*s = strings.Fields(input)
	}

	return nil
}

// MarshalJSON handles json marshaling of this type
func (s Scope) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(s.Unique()))
}

// Value returns Permissions as a value that can be stored as json in the database
func (s Scope) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Scan reads a json value from the database into a Permissions
func (s *Scope) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	val := make([]string, 0)
	if err := json.Unmarshal(b, &val); err != nil {
		return err
	}
	*s = val

	return nil
}

// Get returns the scope for the audience
func (p ScopeSet) Get(a string) Scope {
	if s, ok := p[a]; ok {
		return s
	}
	return Scope{}
}

// Set sets a value in scope set
func (p ScopeSet) Set(a string, s ...string) {
	p[a] = s
}

// Append appends to the scope set
func (p ScopeSet) Append(a string, s ...string) {
	if p[a] == nil {
		p[a] = make(Scope, 0)
	}
	p[a] = append(p[a], s...)
}

// Value returns PermissionSet as a value that can be stored as json in the database
func (p ScopeSet) Value() (driver.Value, error) {
	perms := make(ScopeSet)
	for k, v := range p {
		perms[k] = v.Unique()
	}

	return json.Marshal(perms)
}

// Scan reads a json value from the database into a PermissionSet
func (p ScopeSet) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	return nil
}
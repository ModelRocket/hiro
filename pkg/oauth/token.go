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
	"encoding/json"
	"time"

	"github.com/ModelRocket/hiro/pkg/types"
	"github.com/fatih/structs"
)

type (
	// Token represents a revokable set of claims
	Token struct {
		ID        *types.ID `json:"jti,omitempty"`
		Issuer    *URI      `json:"iss,omitempty"`
		Subject   *types.ID `json:"sub,omitempty"`
		Audience  types.ID  `json:"aud,omitempty"`
		ClientID  types.ID  `json:"azp,omitempty"`
		Use       TokenUse  `json:"use,omitempty"`
		AuthTime  *Time     `json:"auth_time,omitempty"`
		Scope     Scope     `json:"scope,omitempty"`
		IssuedAt  Time      `json:"iat,omitempty"`
		ExpiresAt *Time     `json:"exp,omitempty"`
		Revokable bool      `json:"-"`
		RevokedAt *Time     `json:"-"`
		Claims    Claims    `json:"-"`
	}

	// TokenUse defines token usage
	TokenUse string

	// Time is a time structure used for tokens
	Time time.Time
)

const (
	// TokenUseAccess is a token to be used for access
	TokenUseAccess TokenUse = "access"

	// TokenUseIdentity is a token to be used for identity
	TokenUseIdentity TokenUse = "identity"
)

// NewToken intializes a token of use type
func NewToken(use TokenUse) Token {
	return Token{
		Use:    use,
		Claims: make(Claims),
	}
}

// Sign generates an encoded and sign token using the secret
func (t Token) Sign(s TokenSecret) (string, error) {
	// create the full token claims
	enc := structs.New(t)
	enc.TagName = "json"

	c := Claims(enc.Map())

	for k, v := range t.Claims {
		c[k] = v
	}

	return c.Sign(s)
}

// MarshalJSON markshals the time to an epoch
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).UTC().Unix())
}

// Time casts the oauth time back to a time.Time
func (t Time) Time() time.Time {
	return time.Time(t)
}

// Ptr returns a pointer to this time
func (t Time) Ptr() *Time {
	return &t
}

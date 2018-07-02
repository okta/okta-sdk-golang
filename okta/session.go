/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"encoding/json"
	"time"
)

type CreateSessionRequest struct {
	SessionToken string `json:"sessionToken, omitempty"`
}

func (m *CreateSessionRequest) WithSessionToken(v string) *CreateSessionRequest {
	m.SessionToken = v
	return m
}

type SessionResource resource

type Session struct {
	Links map[string]interface{} `json:"_links, omitempty"`
	Amr map[string]interface{} `json:"amr, omitempty"`
	CreatedAt time.Time `json:"createdAt, omitempty"`
	ExpiresAt time.Time `json:"expiresAt, omitempty"`
	Id string `json:"id, omitempty"`
	Idp SessionIdentityProvider `json:"idp, omitempty"`
	LastFactorVerification time.Time `json:"lastFactorVerification, omitempty"`
	LastPasswordVerification time.Time `json:"lastPasswordVerification, omitempty"`
	Login string `json:"login, omitempty"`
	Status string `json:"status, omitempty"`
	UserId string `json:"userId, omitempty"`
}


func (m *SessionResource) Refresh(sessionId string) Session {
	resp, err := m.client.requestExecutor.Post("/api/v1/sessions/"+sessionId+"/lifecycle/refresh", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Session{}

	json.Unmarshal(resp, &r)

	return r
}


type SessionAuthenticationMethod struct {
}



type SessionIdentityProvider struct {
	Id string `json:"id, omitempty"`
	Type string `json:"type, omitempty"`
}



type SessionIdentityProviderType struct {
}



type SessionStatus struct {
}






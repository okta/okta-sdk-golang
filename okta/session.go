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
	"time"
	"github.com/okta/okta-sdk-golang/okta/query"
	"encoding/json"
)

type SessionResource resource

type Session struct {
	Links []string `json:"_links,omitempty"`
	Amr []string `json:"amr,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Id string `json:"id,omitempty"`
	Idp *SessionIdentityProvider `json:"idp,omitempty"`
	LastFactorVerification *time.Time `json:"lastFactorVerification,omitempty"`
	LastPasswordVerification *time.Time `json:"lastPasswordVerification,omitempty"`
	Login string `json:"login,omitempty"`
	Status string `json:"status,omitempty"`
	UserId string `json:"userId,omitempty"`
}

func (m *SessionResource) GetSession(sessionId string, qp *query.Params)  (*Session, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/sessions/"+sessionId+"", qp)
	if err != nil  {
		return nil, err
	}
	
	r := Session{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *SessionResource) EndSession(sessionId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/sessions/"+sessionId+"", qp)
	if err != nil  {
		return err
	}
	return nil
}


func (m *SessionResource) RefreshSession(sessionId string, qp *query.Params)  (*Session, error) {
	resp, err := m.client.requestExecutor.Post("/api/v1/sessions/"+sessionId+"/lifecycle/refresh", nil, qp)
	if err != nil  {
		return nil, err
	}
	
	r := Session{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

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
	"encoding/json"
	"bytes"
)

type AppUserResource resource

type AppUser struct {
	Embedded []string `json:"_embedded,omitempty"`
	Links []string `json:"_links,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Credentials *AppUserCredentials `json:"credentials,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
	Id string `json:"id,omitempty"`
	LastSync *time.Time `json:"lastSync,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	PasswordChanged *time.Time `json:"passwordChanged,omitempty"`
	Profile []string `json:"profile,omitempty"`
	Scope string `json:"scope,omitempty"`
	Status string `json:"status,omitempty"`
	StatusChanged *time.Time `json:"statusChanged,omitempty"`
	SyncState string `json:"syncState,omitempty"`
}

func (m *AppUser) WithCredentials(v *AppUserCredentials) *AppUser {
	m.Credentials = v
	return m
}

func (m *AppUser) WithId(v string) *AppUser {
	m.Id = v
	return m
}

func (m *AppUser) WithProfile(v []string) *AppUser {
	m.Profile = v
	return m
}

func (m *AppUser) WithScope(v string) *AppUser {
	m.Scope = v
	return m
}

func (m *AppUserResource) UpdateApplicationUser(appId string, userId string, body AppUser)  (*AppUser, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/users/"+userId+"", bytes.NewReader(iobytes))
	if err != nil  {
		return nil, err
	}
	
	r := AppUser{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *AppUserResource) DeleteApplicationUser(appId string, userId string) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/apps/"+appId+"/users/"+userId+"")
	if err != nil  {
		return err
	}
	return nil
}



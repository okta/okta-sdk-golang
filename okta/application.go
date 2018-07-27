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

type ApplicationResource resource

type Application struct {
	Embedded []string `json:"_embedded,omitempty"`
	Links []string `json:"_links,omitempty"`
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Credentials *ApplicationCredentials `json:"credentials,omitempty"`
	Features []string `json:"features,omitempty"`
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Licensing *ApplicationLicensing `json:"licensing,omitempty"`
	Name string `json:"name,omitempty"`
	Settings *ApplicationSettings `json:"settings,omitempty"`
	SignOnMode string `json:"signOnMode,omitempty"`
	Status string `json:"status,omitempty"`
	Visibility *ApplicationVisibility `json:"visibility,omitempty"`
}

func (m *Application) WithAccessibility(v *ApplicationAccessibility) *Application {
	m.Accessibility = v
	return m
}

func (m *Application) WithCredentials(v *ApplicationCredentials) *Application {
	m.Credentials = v
	return m
}

func (m *Application) WithFeatures(v []string) *Application {
	m.Features = v
	return m
}

func (m *Application) WithLabel(v string) *Application {
	m.Label = v
	return m
}

func (m *Application) WithLicensing(v *ApplicationLicensing) *Application {
	m.Licensing = v
	return m
}

func (m *Application) WithSettings(v *ApplicationSettings) *Application {
	m.Settings = v
	return m
}

func (m *Application) WithSignOnMode(v string) *Application {
	m.SignOnMode = v
	return m
}

func (m *Application) WithVisibility(v *ApplicationVisibility) *Application {
	m.Visibility = v
	return m
}

func (m *ApplicationResource) GetApplication(appId string)  (*Application, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"")
	if err != nil  {
		return nil, err
	}
	
	r := Application{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *ApplicationResource) UpdateApplication(appId string, body Application)  (*Application, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/apps/"+appId+"", bytes.NewReader(iobytes))
	if err != nil  {
		return nil, err
	}
	
	r := Application{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *ApplicationResource) DeleteApplication(appId string) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/apps/"+appId+"")
	if err != nil  {
		return err
	}
	return nil
}


func (m *ApplicationResource) ActivateApplication(appId string) error {
	_, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/lifecycle/activate", nil)
	if err != nil  {
		return err
	}
	return nil
}
func (m *ApplicationResource) DeactivateApplication(appId string) error {
	_, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/lifecycle/deactivate", nil)
	if err != nil  {
		return err
	}
	return nil
}
func (m *ApplicationResource) ListApplicationUsers(appId string)  (*AppUser, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/users")
	if err != nil  {
		return nil, err
	}
	
	r := AppUser{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) AssignUserToApplication(appId string, body AppUser)  (*AppUser, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/users", bytes.NewReader(iobytes))
	if err != nil  {
		return nil, err
	}
	
	r := AppUser{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) GetApplicationUser(appId string, userId string)  (*AppUser, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/users/"+userId+"")
	if err != nil  {
		return nil, err
	}
	
	r := AppUser{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) CreateApplicationGroupAssignment(appId string, groupId string, body ApplicationGroupAssignment)  (*ApplicationGroupAssignment, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/apps/"+appId+"/groups/"+groupId+"", bytes.NewReader(iobytes))
	if err != nil  {
		return nil, err
	}
	
	r := ApplicationGroupAssignment{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) GetApplicationGroupAssignment(appId string, groupId string)  (*ApplicationGroupAssignment, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/groups/"+groupId+"")
	if err != nil  {
		return nil, err
	}
	
	r := ApplicationGroupAssignment{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) CloneApplicationKey(appId string, keyId string)  (*JsonWebKey, error) {
	resp, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/credentials/keys/"+keyId+"/clone", nil)
	if err != nil  {
		return nil, err
	}
	
	r := JsonWebKey{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) GetApplicationKey(appId string, keyId string)  (*JsonWebKey, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/credentials/keys/"+keyId+"")
	if err != nil  {
		return nil, err
	}
	
	r := JsonWebKey{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) ListApplicationGroupAssignments(appId string)  (*ApplicationGroupAssignment, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/groups")
	if err != nil  {
		return nil, err
	}
	
	r := ApplicationGroupAssignment{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}
func (m *ApplicationResource) ListApplicationKeys(appId string)  (*JsonWebKey, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/credentials/keys")
	if err != nil  {
		return nil, err
	}
	
	r := JsonWebKey{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

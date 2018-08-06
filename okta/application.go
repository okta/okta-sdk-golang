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
	"fmt"
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

func (m *ApplicationResource) GetApplication(appId string, qp *query.Params)  (*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var application *Application
	resp, err := m.client.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (m *ApplicationResource) UpdateApplication(appId string, body Application, qp *query.Params)  (*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}


	var application *Application
	resp, err := m.client.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (m *ApplicationResource) DeleteApplication(appId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}


	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (m *ApplicationResource) ActivateApplication(appId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/activate", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}


	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *ApplicationResource) DeactivateApplication(appId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/deactivate", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}


	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *ApplicationResource) ListApplicationUsers(appId string, qp *query.Params)  ([]*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var appUser []*AppUser
	resp, err := m.client.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (m *ApplicationResource) AssignUserToApplication(appId string, body AppUser, qp *query.Params)  (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}


	var appUser *AppUser
	resp, err := m.client.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (m *ApplicationResource) GetApplicationUser(appId string, userId string, qp *query.Params)  (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var appUser *AppUser
	resp, err := m.client.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (m *ApplicationResource) CreateApplicationGroupAssignment(appId string, groupId string, body ApplicationGroupAssignment, qp *query.Params)  (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}


	var applicationGroupAssignment *ApplicationGroupAssignment
	resp, err := m.client.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (m *ApplicationResource) GetApplicationGroupAssignment(appId string, groupId string, qp *query.Params)  (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var applicationGroupAssignment *ApplicationGroupAssignment
	resp, err := m.client.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (m *ApplicationResource) CloneApplicationKey(appId string, keyId string, qp *query.Params)  (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v/clone", appId, keyId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var jsonWebKey *JsonWebKey
	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}
func (m *ApplicationResource) GetApplicationKey(appId string, keyId string, qp *query.Params)  (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v", appId, keyId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var jsonWebKey *JsonWebKey
	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}
func (m *ApplicationResource) ListApplicationGroupAssignments(appId string, qp *query.Params)  ([]*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var applicationGroupAssignment []*ApplicationGroupAssignment
	resp, err := m.client.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (m *ApplicationResource) ListApplicationKeys(appId string, qp *query.Params)  ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys", appId)
	if &qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var jsonWebKey []*JsonWebKey
	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}

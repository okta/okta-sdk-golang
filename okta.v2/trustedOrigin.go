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
	"github.com/okta/okta-sdk-golang/okta.v2/query"
	"time"
)

type TrustedOriginResource resource

type TrustedOrigin struct {
	Links         interface{} `json:"_links,omitempty"`
	Created       *time.Time  `json:"created,omitempty"`
	CreatedBy     string      `json:"createdBy,omitempty"`
	Id            string      `json:"id,omitempty"`
	LastUpdated   *time.Time  `json:"lastUpdated,omitempty"`
	LastUpdatedBy string      `json:"lastUpdatedBy,omitempty"`
	Name          string      `json:"name,omitempty"`
	Origin        string      `json:"origin,omitempty"`
	Scopes        []*Scope    `json:"scopes,omitempty"`
	Status        string      `json:"status,omitempty"`
}

func (m *TrustedOriginResource) CreateOrigin(body TrustedOrigin) (*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin *TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

func (m *TrustedOriginResource) GetOrigin(trustedOriginId string) (*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins/%v", trustedOriginId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin *TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

func (m *TrustedOriginResource) UpdateOrigin(trustedOriginId string, body TrustedOrigin) (*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins/%v", trustedOriginId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin *TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

func (m *TrustedOriginResource) DeleteOrigin(trustedOriginId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins/%v", trustedOriginId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *TrustedOriginResource) ListOrigins(qp *query.Params) ([]*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin []*TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

func (m *TrustedOriginResource) ActivateOrigin(trustedOriginId string) (*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins/%v/lifecycle/activate", trustedOriginId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin *TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

func (m *TrustedOriginResource) DeactivateOrigin(trustedOriginId string) (*TrustedOrigin, *Response, error) {
	url := fmt.Sprintf("/api/v1/trustedOrigins/%v/lifecycle/deactivate", trustedOriginId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var trustedOrigin *TrustedOrigin

	resp, err := m.client.requestExecutor.Do(req, &trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, nil
}

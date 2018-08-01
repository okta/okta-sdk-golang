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
	"github.com/okta/okta-sdk-golang/okta/query"
	"fmt"
)

type FactorResource resource

type Factor struct {
	Embedded []string `json:"_embedded,omitempty"`
	Links []string `json:"_links,omitempty"`
	Device string `json:"device,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	FactorType string `json:"factorType,omitempty"`
	Id string `json:"id,omitempty"`
	MfaStateTokenId string `json:"mfaStateTokenId,omitempty"`
	Profile *FactorProfile `json:"profile,omitempty"`
	Provider string `json:"provider,omitempty"`
	RechallengeExistingFactor bool `json:"rechallengeExistingFactor,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	Status string `json:"status,omitempty"`
	UserId string `json:"userId,omitempty"`
	Verify *VerifyFactorRequest `json:"verify,omitempty"`
}

func (m *Factor) WithDevice(v string) *Factor {
	m.Device = v
	return m
}

func (m *Factor) WithFactorType(v string) *Factor {
	m.FactorType = v
	return m
}

func (m *Factor) WithMfaStateTokenId(v string) *Factor {
	m.MfaStateTokenId = v
	return m
}

func (m *Factor) WithProfile(v *FactorProfile) *Factor {
	m.Profile = v
	return m
}

func (m *Factor) WithProvider(v string) *Factor {
	m.Provider = v
	return m
}

func (m *Factor) WithRechallengeExistingFactor(v bool) *Factor {
	m.RechallengeExistingFactor = v
	return m
}

func (m *Factor) WithSessionId(v string) *Factor {
	m.SessionId = v
	return m
}

func (m *Factor) WithUserId(v string) *Factor {
	m.UserId = v
	return m
}

func (m *Factor) WithVerify(v *VerifyFactorRequest) *Factor {
	m.Verify = v
	return m
}

func (m *FactorResource) DeleteFactor(userId string, factorId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
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

	func (m *FactorResource) ActivateFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params)  (*Factor, *Response, error) {
		url := fmt.Sprintf("/api/v1/users/%v/factors/%v/lifecycle/activate", userId, factorId)
		if &qp != nil {
			url = url + qp.String()
		}
		req, err := m.client.requestExecutor.NewRequest("POST", url, body)
		if err != nil {
			return nil, nil, err
		}
	
	
		var factor *Factor
		resp, err := m.client.requestExecutor.Do(req, &factor)
		if err != nil {
			return nil, resp, err
		}
		return factor, resp, nil
	}
	func (m *FactorResource) VerifyFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params)  (*VerifyFactorResponse, *Response, error) {
		url := fmt.Sprintf("/api/v1/users/%v/factors/%v/verify", userId, factorId)
		if &qp != nil {
			url = url + qp.String()
		}
		req, err := m.client.requestExecutor.NewRequest("POST", url, body)
		if err != nil {
			return nil, nil, err
		}
	
	
		var verifyFactorResponse *VerifyFactorResponse
		resp, err := m.client.requestExecutor.Do(req, &verifyFactorResponse)
		if err != nil {
			return nil, resp, err
		}
		return verifyFactorResponse, resp, nil
	}

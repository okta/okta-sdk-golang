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
	"encoding/json"
	"fmt"
	"github.com/okta/okta-sdk-golang/okta/query"
)

type UserFactor interface {
	IsUserFactorInstance() bool
}

type UserFactors []UserFactor

func (uf *UserFactors) UnmarshalJSON(data []byte) error {
	rawFactors := []json.RawMessage{}
	if err := json.Unmarshal(data, &rawFactors); err != nil {
		return err
	}

	factors := make([]map[string]interface{}, len(rawFactors))
	if err := json.Unmarshal(data, &factors); err != nil {
		return err
	}

	for i, f := range factors {
		var factor UserFactor

		switch ft := f["factorType"]; ft {
		case "call":
			factor = new(CallFactor)
		case "email":
			factor = new(EmailFactor)
		case "push":
			factor = new(PushFactor)
		case "question":
			factor = new(SecurityQuestionFactor)
		case "sms":
			factor = new(SmsFactor)
		case "token":
			factor = new(TokenFactor)
		case "token:hardware":
			factor = new(HardwareFactor)
		case "token:software:totp":
			factor = new(TotpFactor)
		case "u2f":
			factor = new(U2fFactor)
		case "web":
			factor = new(WebFactor)
		default:
			return fmt.Errorf("unknown factor type %q", ft)
		}

		if err := json.Unmarshal(rawFactors[i], factor); err != nil {
			return err
		}

		*uf = append(*uf, factor)
	}

	return nil
}

type FactorResource resource

type Factor struct {
	Embedded                  interface{}          `json:"_embedded,omitempty"`
	Links                     interface{}          `json:"_links,omitempty"`
	Device                    string               `json:"device,omitempty"`
	DeviceType                string               `json:"deviceType,omitempty"`
	FactorType                string               `json:"factorType,omitempty"`
	Id                        string               `json:"id,omitempty"`
	MfaStateTokenId           string               `json:"mfaStateTokenId,omitempty"`
	Profile                   *FactorProfile       `json:"profile,omitempty"`
	Provider                  string               `json:"provider,omitempty"`
	RechallengeExistingFactor *bool                `json:"rechallengeExistingFactor,omitempty"`
	SessionId                 string               `json:"sessionId,omitempty"`
	Status                    string               `json:"status,omitempty"`
	TokenLifetimeSeconds      int64                `json:"tokenLifetimeSeconds,omitempty"`
	UserId                    string               `json:"userId,omitempty"`
	Verify                    *VerifyFactorRequest `json:"verify,omitempty"`
}

func NewFactor() *Factor {
	return &Factor{}
}

func (a *Factor) IsUserFactorInstance() bool {
	return true
}

func (m *FactorResource) DeleteFactor(userId string, factorId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
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
func (m *FactorResource) ListFactors(userId string) ([]UserFactor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor UserFactors
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) AddFactor(userId string, body UserFactor, qp *query.Params) (interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	factor := body
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ListSupportedFactors(userId string) ([]interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/catalog", userId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []interface{}
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ListSupportedSecurityQuestions(userId string) ([]*SecurityQuestion, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/questions", userId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var securityQuestion []*SecurityQuestion
	resp, err := m.client.requestExecutor.Do(req, &securityQuestion)
	if err != nil {
		return nil, resp, err
	}
	return securityQuestion, resp, nil
}
func (m *FactorResource) GetFactor(userId string, factorId string, factorInstance UserFactor) (interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	factor := factorInstance
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ActivateFactor(userId string, factorId string, body VerifyFactorRequest) (interface{}, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/lifecycle/activate", userId, factorId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	factor := body
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) VerifyFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params) (*VerifyFactorResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/verify", userId, factorId)
	if qp != nil {
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

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

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"context"
	"fmt"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type PolicyResource resource

type Policy struct {
	Embedded    interface{}           `json:"_embedded,omitempty"`
	Links       interface{}           `json:"_links,omitempty"`
	Conditions  *PolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time            `json:"created,omitempty"`
	Description string                `json:"description,omitempty"`
	Id          string                `json:"id,omitempty"`
	LastUpdated *time.Time            `json:"lastUpdated,omitempty"`
	Name        string                `json:"name,omitempty"`
	Priority    int64                 `json:"priority,omitempty"`
	Status      string                `json:"status,omitempty"`
	System      *bool                 `json:"system,omitempty"`
	Type        string                `json:"type,omitempty"`
}

// Gets a policy.
func (m *PolicyResource) GetPolicy(ctx context.Context, policyId string, qp *query.Params) (*Policy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policy *Policy

	resp, err := m.client.requestExecutor.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Updates a policy.
func (m *PolicyResource) UpdatePolicy(ctx context.Context, policyId string, body Policy) (*Policy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policy *Policy

	resp, err := m.client.requestExecutor.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Removes a policy.
func (m *PolicyResource) DeletePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets all policies with the specified type.
func (m *PolicyResource) ListPolicies(ctx context.Context, qp *query.Params) ([]*AuthorizationServerPolicy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicy []*AuthorizationServerPolicy

	resp, err := m.client.requestExecutor.Do(ctx, req, &authorizationServerPolicy)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicy, resp, nil
}

// Creates a policy.
func (m *PolicyResource) CreatePolicy(ctx context.Context, body Policy, qp *query.Params) (*Policy, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policy *Policy

	resp, err := m.client.requestExecutor.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Activates a policy.
func (m *PolicyResource) ActivatePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/activate", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates a policy.
func (m *PolicyResource) DeactivatePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/deactivate", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates all policy rules.
func (m *PolicyResource) ListPolicyRules(ctx context.Context, policyId string) ([]*PolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule []*PolicyRule

	resp, err := m.client.requestExecutor.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Creates a policy rule.
func (m *PolicyResource) CreatePolicyRule(ctx context.Context, policyId string, body PolicyRule) (*PolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *PolicyRule

	resp, err := m.client.requestExecutor.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Removes a policy rule.
func (m *PolicyResource) DeletePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a policy rule.
func (m *PolicyResource) GetPolicyRule(ctx context.Context, policyId string, ruleId string) (*PolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *PolicyRule

	resp, err := m.client.requestExecutor.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Updates a policy rule.
func (m *PolicyResource) UpdatePolicyRule(ctx context.Context, policyId string, ruleId string, body PolicyRule) (*PolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *PolicyRule

	resp, err := m.client.requestExecutor.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Activates a policy rule.
func (m *PolicyResource) ActivatePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/activate", policyId, ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates a policy rule.
func (m *PolicyResource) DeactivatePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/deactivate", policyId, ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

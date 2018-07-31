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
	"bytes"
)

type GroupRuleResource resource

type GroupRule struct {
	Actions *GroupRuleAction `json:"actions,omitempty"`
	Conditions *GroupRuleConditions `json:"conditions,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Type string `json:"type,omitempty"`
}

func (m *GroupRule) WithActions(v *GroupRuleAction) *GroupRule {
	m.Actions = v
	return m
}

func (m *GroupRule) WithConditions(v *GroupRuleConditions) *GroupRule {
	m.Conditions = v
	return m
}

func (m *GroupRule) WithName(v string) *GroupRule {
	m.Name = v
	return m
}

func (m *GroupRule) WithType(v string) *GroupRule {
	m.Type = v
	return m
}

func (m *GroupRuleResource) UpdateRule(ruleId string, body GroupRule, qp *query.Params)  (*GroupRule, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/groups/rules/"+ruleId+"", bytes.NewReader(iobytes), qp)
	if err != nil  {
		return nil, err
	}
	
	r := GroupRule{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *GroupRuleResource) DeleteRule(ruleId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/groups/rules/"+ruleId+"", qp)
	if err != nil  {
		return err
	}
	return nil
}


func (m *GroupRuleResource) ActivateRule(ruleId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/groups/rules/"+ruleId+"/lifecycle/activate", nil, qp)
	if err != nil  {
		return err
	}
	return nil
}
func (m *GroupRuleResource) DeactivateRule(ruleId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/groups/rules/"+ruleId+"/lifecycle/deactivate", nil, qp)
	if err != nil  {
		return err
	}
	return nil
}

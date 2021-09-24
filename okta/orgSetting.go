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
)

type OrgSettingResource resource

type OrgSetting struct {
	Links                 interface{} `json:"_links,omitempty"`
	Address1              string      `json:"address1,omitempty"`
	Address2              string      `json:"address2,omitempty"`
	City                  string      `json:"city,omitempty"`
	CompanyName           string      `json:"companyName,omitempty"`
	Country               string      `json:"country,omitempty"`
	Created               *time.Time  `json:"created,omitempty"`
	EndUserSupportHelpURL string      `json:"endUserSupportHelpURL,omitempty"`
	ExpiresAt             *time.Time  `json:"expiresAt,omitempty"`
	Id                    string      `json:"id,omitempty"`
	LastUpdated           *time.Time  `json:"lastUpdated,omitempty"`
	PhoneNumber           string      `json:"phoneNumber,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	State                 string      `json:"state,omitempty"`
	Status                string      `json:"status,omitempty"`
	Subdomain             string      `json:"subdomain,omitempty"`
	SupportPhoneNumber    string      `json:"supportPhoneNumber,omitempty"`
	Website               string      `json:"website,omitempty"`
}

// Update settings of your organization.
func (m *OrgSettingResource) UpdateOrgSetting(ctx context.Context, body OrgSetting) (*OrgSetting, *Response, error) {
	url := fmt.Sprintf("/api/v1/org")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var orgSetting *OrgSetting

	resp, err := rq.Do(ctx, req, &orgSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgSetting, resp, nil
}

// Partial update settings of your organization.
func (m *OrgSettingResource) PartialUpdateOrgSetting(ctx context.Context, body OrgSetting) (*OrgSetting, *Response, error) {
	url := fmt.Sprintf("/api/v1/org")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var orgSetting *OrgSetting

	resp, err := rq.Do(ctx, req, &orgSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgSetting, resp, nil
}
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
)

type OrgOktaCommunicationSettingResource resource

type OrgOktaCommunicationSetting struct {
	Links            interface{} `json:"_links,omitempty"`
	OptOutEmailUsers *bool       `json:"optOutEmailUsers,omitempty"`
}

// Gets Okta Communication Settings of your organization.
func (m *OrgOktaCommunicationSettingResource) GetOktaCommunicationSettings(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := fmt.Sprintf("/api/v1/org/privacy/oktaCommunication")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

// Opts in all users of this org to Okta Communication emails.
func (m *OrgOktaCommunicationSettingResource) OptInUsersToOktaCommunicationEmails(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := fmt.Sprintf("/api/v1/org/privacy/oktaCommunication/optIn")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

// Opts out all users of this org from Okta Communication emails.
func (m *OrgOktaCommunicationSettingResource) OptOutUsersFromOktaCommunicationEmails(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := fmt.Sprintf("/api/v1/org/privacy/oktaCommunication/optOut")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

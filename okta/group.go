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

type GroupResource resource

type Group struct {
	Embedded []string `json:"_embedded,omitempty"`
	Links []string `json:"_links,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	LastMembershipUpdated *time.Time `json:"lastMembershipUpdated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	ObjectClass []string `json:"objectClass,omitempty"`
	Profile *GroupProfile `json:"profile,omitempty"`
	Type string `json:"type,omitempty"`
}

func (m *GroupResource) UpdateGroup(groupId string, body Group, qp *query.Params)  (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}


	var group *Group
	resp, err := m.client.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (m *GroupResource) DeleteGroup(groupId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)
	if qp != nil {
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

func (m *GroupResource) RemoveGroupUser(groupId string, userId string, qp *query.Params)  (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users/%v", groupId, userId)
	if qp != nil {
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
func (m *GroupResource) ListGroupUsers(groupId string, qp *query.Params)  ([]*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}


	var user []*User
	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}

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

func (m *Group) WithProfile(v *GroupProfile) *Group {
	m.Profile = v
	return m
}

func (m *GroupResource) UpdateGroup(groupId string, body Group)  (*Group, error) {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/groups/"+groupId+"", bytes.NewReader(iobytes))
	if err != nil  {
		return nil, err
	}
	
	r := Group{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

func (m *GroupResource) DeleteGroup(groupId string) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/groups/"+groupId+"")
	if err != nil  {
		return err
	}
	return nil
}


func (m *GroupResource) RemoveGroupUser(groupId string, userId string) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/groups/"+groupId+"/users/"+userId+"")
	if err != nil  {
		return err
	}
	return nil
}
func (m *GroupResource) ListGroupUsers(groupId string)  (*User, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/groups/"+groupId+"/users")
	if err != nil  {
		return nil, err
	}
	
	r := User{}
	
	json.Unmarshal(resp, &r)
	
	return &r, nil
}

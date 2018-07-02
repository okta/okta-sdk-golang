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
	"encoding/json"
	"time"
)
type GroupResource resource

type Group struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Id string `json:"id, omitempty"`
	LastMembershipUpdated time.Time `json:"lastMembershipUpdated, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	ObjectClass map[string]interface{} `json:"objectClass, omitempty"`
	Profile GroupProfile `json:"profile, omitempty"`
	Type string `json:"type, omitempty"`
}

func (m *Group) WithProfile(v GroupProfile) *Group {
	m.Profile = v
	return m
}

func (m *GroupResource) Removeuser(groupId string, userId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/groups/"+groupId+"/users/"+userId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *GroupResource) Listusers(groupId string) User {
	resp, err := m.client.requestExecutor.Get("/api/v1/groups/"+groupId+"/users")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := User{}

	json.Unmarshal(resp, &r)

	return r
}


type GroupProfile struct {
	Description string `json:"description, omitempty"`
	Name string `json:"name, omitempty"`
}

func (m *GroupProfile) WithDescription(v string) *GroupProfile {
	m.Description = v
	return m
}
func (m *GroupProfile) WithName(v string) *GroupProfile {
	m.Name = v
	return m
}





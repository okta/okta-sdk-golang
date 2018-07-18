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
)

type ApplicationGroupAssignment struct {
	Embedded []string `json:"_embedded, omitempty"`
	Links []string `json:"_links, omitempty"`
	Id string `json:"id, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	Priority int64 `json:"priority, omitempty"`
	Profile []string `json:"profile, omitempty"`
}

func (m *ApplicationGroupAssignment) WithPriority(v int64) *ApplicationGroupAssignment {
	m.Priority = v
	return m
}

func (m *ApplicationGroupAssignment) WithProfile(v []string) *ApplicationGroupAssignment {
	m.Profile = v
	return m
}

func (m *ApplicationGroupAssignmentResource) DeleteApplicationGroupAssignment(appId string, groupId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/apps/"+appId+"/groups/"+groupId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}



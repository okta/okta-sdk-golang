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
	"github.com/okta/okta-sdk-golang/okta.v2/query"
	"time"
)

type GroupResource resource

type Group struct {
	Embedded              interface{}   `json:"_embedded,omitempty"`
	Links                 interface{}   `json:"_links,omitempty"`
	Created               *time.Time    `json:"created,omitempty"`
	Id                    string        `json:"id,omitempty"`
	LastMembershipUpdated *time.Time    `json:"lastMembershipUpdated,omitempty"`
	LastUpdated           *time.Time    `json:"lastUpdated,omitempty"`
	ObjectClass           []string      `json:"objectClass,omitempty"`
	Profile               *GroupProfile `json:"profile,omitempty"`
	Type                  string        `json:"type,omitempty"`
}

// Updates the profile for a group with &#x60;OKTA_GROUP&#x60; type from your organization.
func (m *GroupResource) UpdateGroup(groupId string, body Group) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
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

// Removes a group with &#x60;OKTA_GROUP&#x60; type from your organization.
func (m *GroupResource) DeleteGroup(groupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates groups in your organization with pagination. A subset of groups can be returned that match a supported filter expression or query.
func (m *GroupResource) ListGroups(qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group

	resp, err := m.client.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// Adds a new group with &#x60;OKTA_GROUP&#x60; type to your organization.
func (m *GroupResource) CreateGroup(body Group) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
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

// Lists all group rules for your organization.
func (m *GroupResource) ListGroupRules(qp *query.Params) ([]*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var groupRule []*GroupRule

	resp, err := m.client.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}

	return groupRule, resp, nil
}

// Creates a group rule to dynamically add users to the specified group if they match the condition
func (m *GroupResource) CreateGroupRule(body GroupRule) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule

	resp, err := m.client.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}

	return groupRule, resp, nil
}

// Removes a specific group rule by id from your organization
func (m *GroupResource) DeleteGroupRule(ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches a specific group rule by id from your organization
func (m *GroupResource) GetGroupRule(ruleId string, qp *query.Params) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule

	resp, err := m.client.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}

	return groupRule, resp, nil
}

// Updates a group rule. Only &#x60;INACTIVE&#x60; rules can be updated.
func (m *GroupResource) UpdateGroupRule(ruleId string, body GroupRule) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule

	resp, err := m.client.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}

	return groupRule, resp, nil
}

// Activates a specific group rule by id from your organization
func (m *GroupResource) ActivateGroupRule(ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v/lifecycle/activate", ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates a specific group rule by id from your organization
func (m *GroupResource) DeactivateGroupRule(ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v/lifecycle/deactivate", ruleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all group rules for your organization.
func (m *GroupResource) GetGroup(groupId string) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
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

// Enumerates all applications that are assigned to a group.
func (m *GroupResource) ListAssignedApplicationsForGroup(groupId string, qp *query.Params) ([]App, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/apps", groupId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application []Application

	resp, err := m.client.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}

	apps := make([]App, len(application))
	for i := range application {
		apps[i] = &application[i]
	}
	return apps, resp, nil

}

func (m *GroupResource) ListGroupAssignedRoles(groupId string, qp *query.Params) ([]*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles", groupId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var role []*Role

	resp, err := m.client.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// Assigns a Role to a Group
func (m *GroupResource) AssignRoleToGroup(groupId string, body AssignRoleRequest, qp *query.Params) (*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles", groupId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var role *Role

	resp, err := m.client.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// Unassigns a Role from a Group
func (m *GroupResource) RemoveRoleFromGroup(groupId string, roleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v", groupId, roleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *GroupResource) GetRole(groupId string, roleId string) (*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v", groupId, roleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var role *Role

	resp, err := m.client.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// Lists all App targets for an &#x60;APP_ADMIN&#x60; Role assigned to a Group. This methods return list may include full Applications or Instances. The response for an instance will have an &#x60;ID&#x60; value, while Application will not have an ID.
func (m *GroupResource) ListApplicationTargetsForApplicationAdministratorRoleForGroup(groupId string, roleId string, qp *query.Params) ([]App, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/catalog/apps", groupId, roleId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application []Application

	resp, err := m.client.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}

	apps := make([]App, len(application))
	for i := range application {
		apps[i] = &application[i]
	}
	return apps, resp, nil

}

func (m *GroupResource) RemoveApplicationTargetFromApplicationAdministratorRoleGivenToGroup(groupId string, roleId string, appName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/catalog/apps/%v", groupId, roleId, appName)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *GroupResource) AddApplicationTargetToAdminRoleGivenToGroup(groupId string, roleId string, appName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/catalog/apps/%v", groupId, roleId, appName)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Remove App Instance Target to App Administrator Role given to a Group
func (m *GroupResource) RemoveApplicationTargetFromAdministratorRoleGivenToGroup(groupId string, roleId string, appName string, applicationId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/catalog/apps/%v/%v", groupId, roleId, appName, applicationId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Add App Instance Target to App Administrator Role given to a Group
func (m *GroupResource) AddApplicationInstanceTargetToAppAdminRoleGivenToGroup(groupId string, roleId string, appName string, applicationId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/catalog/apps/%v/%v", groupId, roleId, appName, applicationId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *GroupResource) ListGroupTargetsForGroupRole(groupId string, roleId string, qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/groups", groupId, roleId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group

	resp, err := m.client.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

//
func (m *GroupResource) RemoveGroupTargetFromGroupAdministratorRoleGivenToGroup(groupId string, roleId string, targetGroupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/groups/%v", groupId, roleId, targetGroupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

//
func (m *GroupResource) AddGroupTargetToGroupAdministratorRoleForGroup(groupId string, roleId string, targetGroupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/roles/%v/targets/groups/%v", groupId, roleId, targetGroupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates all users that are a member of a group.
func (m *GroupResource) ListGroupUsers(groupId string, qp *query.Params) ([]*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users", groupId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
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

// Removes a user from a group with &#x27;OKTA_GROUP&#x27; type.
func (m *GroupResource) RemoveUserFromGroup(groupId string, userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users/%v", groupId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Adds a user to a group with &#x27;OKTA_GROUP&#x27; type.
func (m *GroupResource) AddUserToGroup(groupId string, userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users/%v", groupId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

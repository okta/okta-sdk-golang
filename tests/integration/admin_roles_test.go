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

package integration

import (
	"context"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const ADMIN_GROUP_NAME = "Assign Admin Role To Group"

func Test_can_add_an_admin_role_to_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "add_admin_role"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}

	user, _, err := client.User.CreateUser(ctx, *u, nil)
	require.NoError(t, err, "Creating a new user should not error")
	role := okta.AssignRoleRequest{
		Type: "SUPER_ADMIN",
	}
	createdRole, response, err := client.User.AssignRoleToUser(ctx, user.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Role{}, createdRole, "did not return `*okta.Role` as first variable")
	assert.Equal(t, "POST", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, createdRole.Id, "id should not be empty")
	assert.NotEmpty(t, createdRole.Label, "label should not be empty")
	assert.NotEmpty(t, createdRole.Type, "type should not be empty")
	assert.NotEmpty(t, createdRole.Status, "status should not be empty")
	assert.NotEmpty(t, createdRole.Created, "created should not be empty")
	assert.NotEmpty(t, createdRole.LastUpdated, "lastUpdated should not be empty")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_add_an_admin_role_to_group(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	role := okta.AssignRoleRequest{
		Type: "ORG_ADMIN",
	}
	createdRole, response, err := client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Role{}, createdRole, "did not return `*okta.Role` as first variable")
	assert.Equal(t, "POST", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, createdRole.Id, "id should not be empty")
	assert.NotEmpty(t, createdRole.Label, "label should not be empty")
	assert.NotEmpty(t, createdRole.Type, "type should not be empty")
	assert.NotEmpty(t, createdRole.Status, "status should not be empty")
	assert.NotEmpty(t, createdRole.Created, "created should not be empty")
	assert.NotEmpty(t, createdRole.LastUpdated, "lastUpdated should not be empty")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_remove_an_admin_role_to_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "delete_admin_role"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}

	user, _, err := client.User.CreateUser(ctx, *u, nil)
	require.NoError(t, err, "Creating a new user should not error")
	role := okta.AssignRoleRequest{
		Type: "SUPER_ADMIN",
	}
	createdRole, response, err := client.User.AssignRoleToUser(ctx, user.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Role{}, createdRole, "did not return `*okta.Role` as first variable")
	assert.Equal(t, "POST", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	response, err = client.User.RemoveRoleFromUser(ctx, user.Id, createdRole.Id)
	require.NoError(t, err, "removing role from user must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, 204, response.StatusCode, "did not return a 204")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_remove_an_admin_role_to_group(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	role := okta.AssignRoleRequest{
		Type: "ORG_ADMIN",
	}
	createdRole, response, err := client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Role{}, createdRole, "did not return `*okta.Role` as first variable")
	assert.Equal(t, "POST", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	response, err = client.Group.RemoveRoleFromGroup(ctx, group.Id, createdRole.Id)
	require.NoError(t, err, "removing role from group must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, 204, response.StatusCode, "did not return a 204")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_list_roles_assigned_to_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "list_roles"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}

	user, _, err := client.User.CreateUser(ctx, *u, nil)
	require.NoError(t, err, "Creating a new user should not error")

	role := okta.AssignRoleRequest{
		Type: "SUPER_ADMIN",
	}

	_, _, err = client.User.AssignRoleToUser(ctx, user.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	roles, response, err := client.User.ListAssignedRolesForUser(ctx, user.Id, nil)

	require.NoError(t, err, "listing administrator roles must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Role{}, roles, "did not return `[]*okta.Role` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, roles, "listing roles result should not be empty")

	assert.NotEmpty(t, roles[0].Id, "id should not be empty")
	assert.NotEmpty(t, roles[0].Label, "label should not be empty")
	assert.NotEmpty(t, roles[0].Type, "type should not be empty")
	assert.NotEmpty(t, roles[0].Status, "status should not be empty")
	assert.NotEmpty(t, roles[0].Created, "created should not be empty")
	assert.NotEmpty(t, roles[0].LastUpdated, "lastUpdated should not be empty")
	assert.NotEmpty(t, roles[0].AssignmentType, "assignmentType should not be empty")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_list_roles_assigned_to_a_group(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	role := okta.AssignRoleRequest{
		Type: "ORG_ADMIN",
	}
	_, _, err = client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	roles, response, err := client.Group.ListGroupAssignedRoles(ctx, group.Id, nil)

	require.NoError(t, err, "listing administrator roles must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Role{}, roles, "did not return `[]*okta.Role` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, roles, "listing roles result should not be empty")

	assert.NotEmpty(t, roles[0].Id, "id should not be empty")
	assert.NotEmpty(t, roles[0].Label, "label should not be empty")
	assert.NotEmpty(t, roles[0].Type, "type should not be empty")
	assert.NotEmpty(t, roles[0].Status, "status should not be empty")
	assert.NotEmpty(t, roles[0].Created, "created should not be empty")
	assert.NotEmpty(t, roles[0].LastUpdated, "lastUpdated should not be empty")
	assert.NotEmpty(t, roles[0].AssignmentType, "assignmentType should not be empty")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_add_group_targets_for_the_group_administrator_role_given_to_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "add-group-targets"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}

	user, _, err := client.User.CreateUser(ctx, *u, nil)
	require.NoError(t, err, "Creating a new user should not error")

	role := okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}

	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")

	addedRole, _, err := client.User.AssignRoleToUser(ctx, user.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	response, err := client.User.AddGroupTargetToRole(ctx, user.Id, addedRole.Id, group.Id)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, "PUT", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles/"+addedRole.Id+"/targets/groups/"+group.Id, response.Response.Request.URL.Path, "path for request was incorrect")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_add_group_targets_for_the_group_administrator_role_given_to_a_group(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	role := okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}

	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")

	addedRole, _, err := client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	response, err := client.Group.AddGroupTargetToGroupAdministratorRoleForGroup(ctx, group.Id, addedRole.Id, group.Id)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, "PUT", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles/"+addedRole.Id+"/targets/groups/"+group.Id, response.Response.Request.URL.Path, "path for request was incorrect")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_list_group_targets_for_the_group_administrator_role_given_to_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "add-group-targets"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}

	user, _, err := client.User.CreateUser(ctx, *u, nil)
	require.NoError(t, err, "Creating a new user should not error")

	role := okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}

	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")

	addedRole, _, err := client.User.AssignRoleToUser(ctx, user.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	response, err := client.User.AddGroupTargetToRole(ctx, user.Id, addedRole.Id, group.Id)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, "PUT", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles/"+addedRole.Id+"/targets/groups/"+group.Id, response.Response.Request.URL.Path, "path for request was incorrect")

	listRoles, response, err := client.User.ListGroupTargetsForRole(ctx, user.Id, addedRole.Id, nil)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Group{}, listRoles, "did not return `[]*okta.Group` type as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/users/"+user.Id+"/roles/"+addedRole.Id+"/targets/groups", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, listRoles[0].Id, "id should not be empty")
	assert.NotEmpty(t, listRoles[0].ObjectClass, "objectClass should not be empty")
	assert.NotEmpty(t, listRoles[0].Profile, "profile should not be empty")
	assert.IsType(t, &okta.GroupProfile{}, listRoles[0].Profile, "profile should be instance of *okta.GroupProfile")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_list_group_targets_for_the_group_administrator_role_given_to_a_group(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	role := okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}

	gp := &okta.GroupProfile{
		Name: testName(ADMIN_GROUP_NAME),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")

	addedRole, _, err := client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
	require.NoError(t, err, "adding role to user must not error")

	response, err := client.Group.AddGroupTargetToGroupAdministratorRoleForGroup(ctx, group.Id, addedRole.Id, group.Id)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, "PUT", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles/"+addedRole.Id+"/targets/groups/"+group.Id, response.Response.Request.URL.Path, "path for request was incorrect")

	listRoles, response, err := client.Group.ListGroupTargetsForGroupRole(ctx, group.Id, addedRole.Id, nil)
	require.NoError(t, err, "list group assignments must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Group{}, listRoles, "did not return `[]*okta.Group` type as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/groups/"+group.Id+"/roles/"+addedRole.Id+"/targets/groups", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, listRoles[0].Id, "id should not be empty")
	assert.NotEmpty(t, listRoles[0].ObjectClass, "objectClass should not be empty")
	assert.NotEmpty(t, listRoles[0].Profile, "profile should not be empty")
	assert.IsType(t, &okta.GroupProfile{}, listRoles[0].Profile, "profile should be instance of *okta.GroupProfile")

	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

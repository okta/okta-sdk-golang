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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanGetAGroup(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST Get Test Group"),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Get the group by ID → GET /api/v1/groups/{{groupId}}
	foundGroup, _, err := client.Group.GetGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when finding a group")
	assert.Equal(t, group.Id, foundGroup.Id, "Group that was found was not correct")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")

	// Verify that the group is deleted by calling get on group (Exception thrown with 404 error message) → GET /api/v1/groups/{{groupId}}
	_, resp, err := client.Group.GetGroup(ctx, group.Id)
	assert.Error(t, err, "Finding a group by id should have reported an error")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode,
		"Should have resulted in a 404 when finding a deleted group")
}

func TestCanListGroups(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST List Test Group"),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// List all groups and find the group created → GET /api/v1/groups
	groupList, _, err := client.Group.ListGroups(ctx, nil)
	require.NoError(t, err, "Listing groups should not error")
	found := false
	for _, grp := range groupList {
		if grp.Id == group.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not find group from list")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func TestCanSearchForAGroup(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a new group → POST /api/v1/groups
	groupName := testName("SDK_TEST Search Test Group")
	gp := &okta.GroupProfile{
		Name: groupName,
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Search the group by name with query parameter → GET /api/v1/groups?q=Search
	groupList, _, err := client.Group.ListGroups(ctx, query.NewQueryParams(query.WithQ(groupName)))
	assert.Len(t, groupList, 1, "Did not find correct amount of groups")
	require.NoError(t, err, "Listing groups should not error")
	found := false
	for _, grp := range groupList {
		if grp.Id == group.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not find group from list")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func TestCanUpdateAGroup(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a new group → POST /api/v1/groups
	groupName := testName("SDK_TEST Update Test Group")
	gp := &okta.GroupProfile{
		Name: groupName,
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Update the group name and description → PUT /api/v1/groups/{{groupId}}
	newGroupName := testName("SDK_TEST Updated Name")
	ngp := &okta.GroupProfile{
		Name: newGroupName,
	}
	client.Group.UpdateGroup(ctx, group.Id, okta.Group{Profile: ngp})

	// Verify that group profile is updated by calling get on the group and verifying the profile → GET /api/v1/groups/{{groupId}}
	updatedGroup, _, err := client.Group.GetGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when getting updated group")
	assert.Equal(t, newGroupName, updatedGroup.Profile.Name, "The group was not updated")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func TestGroupUserOperations(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a user with credentials → POST /api/v1/users?activate=false
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "With-Group"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")
	assert.IsType(t, &okta.User{}, user)

	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST Group-Member API Test Group"),
	}
	g := &okta.Group{
		Profile: gp,
	}

	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Add user to the group  → POST /api/v1/groups/{{groupId}}/users/{{userId}}
	_, err = client.Group.AddUserToGroup(ctx, group.Id, user.Id)
	require.NoError(t, err, "Should not error when adding user to group")

	// Validate user present in group → GET /api/v1/groups/{{groupId}}/users
	users, _, err := client.Group.ListGroupUsers(ctx, group.Id, nil)
	require.NoError(t, err)
	found := false
	for _, tmpuser := range users {
		if tmpuser.Id == user.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not find user in group")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Delete the group →  DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting a group")
}

func TestGroupRuleOperations(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: testPassword(10),
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "With-Group-Rule"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")
	assert.IsType(t, &okta.User{}, user)

	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST Group-Member-Rule API Test Group"),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Create a group rule and verify rule executes → POST /api/v1/groups/rules
	// The rule below adds the user created in step 1 to the group created in step 2 upon rule execution/activation
	lastName := profile["lastName"].(string)
	grce := &okta.GroupRuleExpression{
		Type:  "urn:okta:expression:1.0",
		Value: "user.lastName==\"" + lastName + "\"",
	}
	grc := &okta.GroupRuleConditions{
		Expression: grce,
	}
	grga := &okta.GroupRuleGroupAssignment{
		GroupIds: []string{group.Id},
	}
	gra := &okta.GroupRuleAction{
		AssignUserToGroups: grga,
	}
	gr := &okta.GroupRule{
		Actions:    gra,
		Conditions: grc,
		Type:       "group_rule",
		Name:       testName("SDK_TEST group rule"),
	}
	groupRule, _, err := client.Group.CreateGroupRule(ctx, *gr)
	require.NoError(t, err, "Should not error when creating a group Rule")
	assert.IsType(t, &okta.GroupRule{}, groupRule)

	// Activate the above rule and verify that user is added to the group → POST /api/v1/groups/rules/{{ruleId}}/lifecycle/activate
	_, err = client.Group.ActivateGroupRule(ctx, groupRule.Id)
	require.NoError(t, err, "Should not error when activating rule")

	users := []*okta.User{}

	// Use a backoff to check the user is in the group as there can be eventual
	// consistency issues in adding users to groups.
	operation := func() error {
		users, _, err = client.Group.ListGroupUsers(ctx, group.Id, nil)
		if err != nil {
			return err
		}
		for _, tmpuser := range users {
			if tmpuser.Id == user.Id {
				return nil
			}
		}
		return fmt.Errorf("returning error so backoff continues to looking for user being added")
	}
	bOff := backoff.NewExponentialBackOff()
	bOff.MaxElapsedTime = 30 * time.Second
	err = backoff.Retry(operation, bOff)
	require.NoError(t, err, "Inspecting group for user addition had an issue.")

	// List the group rules and validate the above rule is present → POST /api/v1/groups/rules
	groupRules, _, err := client.Group.ListGroupRules(ctx, nil)
	require.NoError(t, err, "Error should not happen when listing rules")
	found := false
	for _, tmpRules := range groupRules {
		if tmpRules.Id == groupRule.Id {
			found = true
		}
	}
	assert.True(t, found, "Group rule execution did not happen")

	// Deactivate the rule  → POST /api/v1/groups/rules/{{ruleId}}/lifecycle/deactivate
	_, err = client.Group.DeactivateGroupRule(ctx, groupRule.Id)
	require.NoError(t, err, "Error should not happen when deactivating rule")

	// Update the rule (Rule can only be updated when it's deactivated) → POST /api/v1/groups/rules/{{ruleId}}
	grce = &okta.GroupRuleExpression{
		Type:  "urn:okta:expression:1.0",
		Value: "user.lastName==\"Incorrect\"",
	}
	grc = &okta.GroupRuleConditions{
		Expression: grce,
	}
	grga = &okta.GroupRuleGroupAssignment{
		GroupIds: []string{group.Id},
	}
	gra = &okta.GroupRuleAction{
		AssignUserToGroups: grga,
	}
	gr = &okta.GroupRule{
		Actions:    gra,
		Conditions: grc,
		Type:       "group_rule",
		Name:       testName("SDK_TEST group rule Updated"),
	}
	newGroupRule, _, err := client.Group.UpdateGroupRule(ctx, groupRule.Id, *gr)
	require.NoError(t, err, "Should not error when updating rule")

	// Activate the updated rule and verify that the user is removed from the group →  POST /api/v1/groups/rules/{{ruleId}}/lifecycle/activate
	_, err = client.Group.ActivateGroupRule(ctx, newGroupRule.Id)
	require.NoError(t, err, "Should not error when activating the group rule")

	bOff.Reset()
	// Use a backoff to check the user has been removed from the group as there
	// can be eventual consistency issues in removing users from groups.
	operation = func() error {
		users, _, err = client.Group.ListGroupUsers(ctx, group.Id, nil)
		if err != nil {
			return err
		}
		for _, tmpuser := range users {
			if tmpuser.Id == user.Id {
				return fmt.Errorf("returning error so backoff continues user still listed in group")
			}
		}
		return nil
	}
	err = backoff.Retry(operation, bOff)
	require.NoError(t, err, "Inspecting group for user removal had an issue.")

	// Deactivate the user, group and group rule → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.Group.DeactivateGroupRule(ctx, newGroupRule.Id)
	require.NoError(t, err, "should not error when deactivating rule")

	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "should not error when deactivating user")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting user")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.Group.DeleteGroup(ctx, group.Id)
	require.NoError(t, err, "Should not error when deleting Group")

	// Delete the group rule → DELETE /api/v1/groups/rules/{{ruleId}}
	_, err = client.Group.DeleteGroupRule(ctx, groupRule.Id, &query.Params{})
	require.NoError(t, err, "Should not error when deleting Rule")
}

func TestGroupProfileSerialization(t *testing.T) {
	gp := okta.GroupProfile{
		Name:        "test",
		Description: "tester",
		GroupProfileMap: okta.GroupProfileMap{
			"custom": "value",
		},
	}

	gpExpected := okta.GroupProfile{
		Name:        "test",
		Description: "tester",
		GroupProfileMap: okta.GroupProfileMap{
			"custom": "value",
		},
	}

	b, err := json.Marshal(&gp)
	require.NoError(t, err)

	var gpCopy okta.GroupProfile
	err = json.Unmarshal(b, &gpCopy)
	require.NoError(t, err)

	assert.Equal(t, gpExpected, gpCopy, "expected marshal to unmarshal to produce exact copy of group profile")
}

func TestListAssignedApplicationsForGroup(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)

	gp := &okta.GroupProfile{
		Name: testName("SDK_TEST Get Test Group"),
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	apps, _, err := client.Group.ListAssignedApplicationsForGroup(ctx, group.Id, nil)
	require.NoError(t, err, "Should not error when listing assigned apps for group")
	assert.Equal(t, 0, len(apps), "there shouldn't be any apps assigned to group")

	app := okta.NewBookmarkApplication()
	app.Settings = &okta.BookmarkApplicationSettings{
		App: &okta.BookmarkApplicationSettingsApplication{
			RequestIntegration: new(bool),
			Url:                "https://example.com/bookmark.htm",
		},
	}
	_, _, err = client.Application.CreateApplication(ctx, app, nil)
	require.NoError(t, err, "Creating an application should not error")

	_, _, err = client.Application.CreateApplicationGroupAssignment(ctx, app.Id, group.Id, okta.ApplicationGroupAssignment{})
	require.NoError(t, err, "Assigning application to group should not error")

	apps, _, err = client.Group.ListAssignedApplicationsForGroup(ctx, group.Id, nil)
	require.NoError(t, err, "Should not error when listing assigned apps for group")
	assert.Equal(t, 1, len(apps), "there should be one app assigned to group")

	// teardown
	client.Application.DeactivateApplication(ctx, app.Id)
	client.Application.DeleteApplication(ctx, app.Id)
	client.Group.DeleteGroup(ctx, group.Id)
}

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
	"net/http"
	"testing"

	"github.com/okta/okta-sdk-golang/okta/query"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_get_a_group(t *testing.T) {
	client := tests.NewClient()
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "Get Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Get the group by ID → GET /api/v1/groups/{{groupId}}
	foundGroup, _, err := client.GetGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when finding a group")
	assert.Equal(t, group.Id, foundGroup.Id, "Group that was found was not correct")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.DeleteGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when deleting a group")

	// Verify that the group is deleted by calling get on group (Exception thrown with 404 error message) → GET /api/v1/groups/{{groupId}}
	_, resp, err := client.GetGroup(group.Id, nil)
	assert.Error(t, err, "Finding a group by id should have reported an error")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode,
		"Should have resulted in a 404 when finding a deleted group")
}

func Test_can_list_groups(t *testing.T) {
	client := tests.NewClient()
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "List Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// List all groups and find the group created → GET /api/v1/groups
	groupList, _, err := client.ListGroups(nil)
	require.NoError(t, err, "Listing groups should not error")
	found := false
	for _, grp := range groupList {
		if grp.Id == group.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not find group from list")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.DeleteGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_search_for_a_group(t *testing.T) {
	client := tests.NewClient()
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "Search Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Search the group by name with query parameter → GET /api/v1/groups?q=Search
	groupList, _, err := client.ListGroups(query.NewQueryParams(query.WithQ("Search Test Group")))
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
	_, err = client.DeleteGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_can_update_a_group(t *testing.T) {
	client := tests.NewClient()
	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "Update Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Update the group name and description → PUT /api/v1/groups/{{groupId}}
	ngp := &okta.GroupProfile{
		Name: "Updated Name",
	}
	client.Group.UpdateGroup(group.Id, okta.Group{Profile: ngp}, nil)

	// Verify that group profile is updated by calling get on the group and verifying the profile → GET /api/v1/groups/{{groupId}}
	updatedGroup, _, err := client.GetGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when getting updated group")
	assert.Equal(t, "Updated Name", updatedGroup.Profile.Name, "The group was not updated")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	_, err = client.DeleteGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when deleting a group")
}

func Test_group_user_operations(t *testing.T) {
	client := tests.NewClient()
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
	profile["email"] = "john-with-group@example.com"
	profile["login"] = "john-with-group@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")
	assert.IsType(t, &okta.User{}, user)

	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "Group-Member API Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}

	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Should not error when creating a group")
	assert.IsType(t, &okta.Group{}, group)

	// Add user to the group  → POST /api/v1/groups/{{groupId}}/users/{{userId}}
	_, err = client.AddUserToGroup(group.Id, user.Id, nil)
	require.NoError(t, err, "Should not error when adding user to group")

	// Validate user present in group → GET /api/v1/groups/{{groupId}}/users
	users, _, err := client.ListGroupUsers(group.Id, nil)
	found := false
	for _, tmpuser := range users {
		if tmpuser.Id == user.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not find user in group")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Delete the group →  DELETE /api/v1/groups/{{groupId}}
	_, err = client.DeleteGroup(group.Id, nil)
	require.NoError(t, err, "Should not error when deleting a group")
}

//func Test_group_rule_operations(t *testing.T) {
//	client := tests.NewClient()
//	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
//	p := new(okta.PasswordCredential).WithValue("Abcd1234")
//	uc := new(okta.UserCredentials).WithPassword(p)
//	profile := okta.UserProfile{}
//	profile["firstName"] = "John"
//	profile["lastName"] = "With-Group-Rule"
//	profile["email"] = "john-with-group-rule@example.com"
//	profile["login"] = "john-with-group-rule@example.com"
//	u := new(okta.User).WithCredentials(uc).WithProfile(&profile)
//	qp := query.NewQueryParams(query.WithActivate(true))
//
//	user, _, err := client.CreateUser(*u, qp)
//	require.NoError(t, err, "Creating an user should not error")
//	assert.IsType(t, &okta.User{}, user)
//
//	// Create a new group → POST /api/v1/groups
//	gp := new(okta.GroupProfile).
//		WithName("Group-Member-Rule API Test Group")
//	g := new(okta.Group).
//		WithProfile(gp)
//	group, _, err := client.CreateGroup(*g, nil)
//	require.NoError(t, err, "Should not error when creating a group")
//	assert.IsType(t, &okta.Group{}, group)
//
//	// Create a group rule and verify rule executes → POST /api/v1/groups/rules
//	// The rule below adds the user created in step 1 to the group created in step 2 upon rule execution/activation
//	grce := new(okta.GroupRuleExpression).
//		WithType("urn:okta:expression:1.0").
//		WithValue("user.lastName==\"${createdUser.profile.lastName}\"")
//	grcp := new(okta.GroupRulePeopleCondition).
//		WithUsers(new(okta.GroupRuleUserCondition)).
//		WithGroups(new(okta.GroupRuleGroupCondition))
//	grc := new(okta.GroupRuleConditions).
//		WithExpression(grce).
//		WithPeople(grcp)
//
//	grga := new(okta.GroupRuleGroupAssignment).
//		WithGroupIds([]string{group.Id})
//	gra := new(okta.GroupRuleAction).
//		WithAssignUserToGroups(grga)
//	gr := new(okta.GroupRule).
//		WithActions(gra).
//		WithConditions(grc).
//		WithType("group_rule").
//		WithName("Test group rule")
//	groupRule, _, err := client.CreateRule(*gr, nil)
//	require.NoError(t, err, "Should not error when creating a group Rule")
//	assert.IsType(t, &okta.GroupRule{}, groupRule)
//
//	// Activate the above rule and verify that user is added to the group → POST /api/v1/groups/rules/{{ruleId}}/lifecycle/activate
//	_, err = client.ActivateRule(groupRule.Id, nil)
//	require.NoError(t, err, "Should not error when activating rule")
//
//	time.Sleep(2 * time.Second)
//	users, _, err := client.ListGroupUsers(group.Id, nil)
//	found := false
//	for _, tmpuser := range users {
//		if tmpuser.Id == user.Id {
//			found = true
//		}
//	}
//	assert.True(t, found, "Group rule execution did not happen")
//
//	// List the group rules and validate the above rule is present → POST /api/v1/groups/rules
//	groupRules, _, err := client.ListRules(nil)
//	require.NoError(t, err, "Error should not happen when listing rules")
//	found = false
//	for _, tmpRules := range groupRules {
//		if tmpRules.Id == groupRule.Id {
//			found = true
//		}
//	}
//	assert.True(t, found, "Group rule execution did not happen")
//
//	// Deactivate the rule  → POST /api/v1/groups/rules/{{ruleId}}/lifecycle/deactivate
//	_, err = client.DeactivateRule(groupRule.Id, nil)
//	require.NoError(t, err, "Error should not happen when deactivating rule")
//
//	// Update the rule (Rule can only be updated when it's deactivated) → POST /api/v1/groups/rules/{{ruleId}}
//	grce = new(okta.GroupRuleExpression).
//		WithValue("user.lastName==\"incorrect\"")
//	grc = new(okta.GroupRuleConditions).
//		WithExpression(grce)
//	gr = new(okta.GroupRule).
//		WithName("Test group rule updated")
//	newGroupRule, _, err := client.UpdateRule(groupRule.Id, *gr, nil)
//	require.NoError(t, err, "Should not error when updating rule")
//
//	// Activate the updated rule and verify that the user is removed from the group →  POST /api/v1/groups/rules/{{ruleId}}/lifecycle/activate
//	_, err = client.ActivateRule(newGroupRule.Id, nil)
//	require.NoError(t, err, "Should not error when activating the group rule")
//	time.Sleep(2 * time.Second)
//	users, _, err = client.ListGroupUsers(group.Id, nil)
//	found = false
//	for _, tmpuser := range users {
//		if tmpuser.Id == user.Id {
//			found = true
//		}
//	}
//	assert.False(t, found, "Group rule execution did not happen to remove user")
//
//	// Deactivate the user, group and group rule → POST /api/v1/users/{{userId}}/lifecycle/deactivate
//	_, err = client.DeactivateRule(newGroupRule.Id, nil)
//	require.NoError(t, err, "should not error when deactivating rule")
//
//	_, err = client.DeactivateUser(user.Id, nil)
//	require.NoError(t, err, "should not error when deactivating user")
//
//	// Delete the user → DELETE /api/v1/users/{{userId}}
//	_, err = client.DeactivateOrDeleteUser(user.Id, nil)
//	require.NoError(t, err, "Should not error when deleting user")
//
//	// Delete the group → DELETE /api/v1/groups/{{groupId}}
//	_, err = client.DeleteGroup(group.Id, nil)
//	require.NoError(t, err, "Should not error when deleting Group")
//
//	// Delete the group rule → DELETE /api/v1/groups/rules/{{ruleId}}
//	_, err = client.DeleteRule(groupRule.Id, nil)
//	require.NoError(t, err, "Should not error when deleting Rule")
//}

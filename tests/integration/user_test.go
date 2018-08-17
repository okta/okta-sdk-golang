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
	"time"

	"github.com/okta/okta-sdk-golang/okta/query"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_get_a_user(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=false
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Get-User"
	profile["email"] = "john-get-user@example.com"
	profile["login"] = "john-get-user@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by id should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Id")

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, _, err = client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by id should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Id")

	// Get the user by login name → GET /api/v1/users/{{login}}
	ubln, _, err := client.User.GetUser(profile["login"].(string), nil)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubln.Id, "Could not find user by Login")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_activate_a_user(t *testing.T) {
	client := tests.NewClient()
	//Create user with credentials → POST /api/v1/users?activate=false
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Activate"
	profile["email"] = "john-activate@example.com"
	profile["login"] = "john-activate@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Activate the user → POST /api/v1/users/{{userId}}/lifecycle/activate?sendEmail=false
	token, _, err := client.User.ActivateUser(user.Id, query.NewQueryParams(query.WithSendEmail(false)))
	require.NoError(t, err, "Could not activate the user")
	assert.NotEmpty(t, token, "Token was not provided")
	assert.IsType(t, &okta.UserActivationToken{}, token, "Activation did not return correct type")

	// Verify that the user is in the list of ACTIVE users with query parameter → GET /api/v1/users?filter=status eq "ACTIVE"
	filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
	users, _, err := client.ListUsers(filter)
	require.NoError(t, err, "Could not get active users")
	found := false
	for _, u := range users {
		if user.Id == u.Id {
			found = true
		}
	}
	assert.True(t, found, "The user was not found")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_update_user_profile(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=false
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Profile-Update"
	profile["email"] = "john-profile-update@example.com"
	profile["login"] = "john-profile-update@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Update the user's profile by adding a nickname → PUT /api/v1/users/{{userId}}
	newProfile := *user.Profile
	newProfile["nickName"] = "Batman"
	updatedUser := &okta.User{
		Profile: &newProfile,
	}
	_, _, err = client.User.UpdateUser(user.Id, *updatedUser, nil)
	require.NoError(t, err, "Could not update the user")

	// Verify that user profile is updated by calling get on the user → GET /api/v1/users/{{userId}}
	tmpUser, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "User was not available to get")
	tmpProfile := *tmpUser.Profile
	assert.Equal(t, "Batman", tmpProfile["nickName"])
	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_suspend_a_user(t *testing.T) {
	client := tests.NewClient()
	//Create user with credentials → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Suspend"
	profile["email"] = "john-suspend@example.com"
	profile["login"] = "john-suspend@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Suspend the user → POST /api/v1/users/{{userId}}/lifecycle/suspend
	_, err = client.User.SuspendUser(user.Id, nil)
	require.NoError(t, err, "Could not suspend the user")

	// Verify that user is in the list of suspended users → GET /api/v1/users?filter=status eq "SUSPENDED"
	filter := query.NewQueryParams(query.WithFilter("status eq \"SUSPENDED\""))
	users, _, err := client.ListUsers(filter)
	require.NoError(t, err, "Could not get suspended users")
	found := false
	for _, u := range users {
		if user.Id == u.Id {
			found = true
		}
	}
	assert.True(t, found, "The user was not found")

	// Unsuspend the user →  POST /api/v1/users/{{userId}}/lifecycle/unsuspend
	_, err = client.User.UnsuspendUser(user.Id, nil)
	require.NoError(t, err, "Could not unsuspend the user")

	// Verify that user is in the list of active users → GET /api/v1/users?filter=status eq "ACTIVE"
	filter = query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
	users, _, err = client.ListUsers(filter)
	require.NoError(t, err, "Could not get active users")
	found = false
	for _, u := range users {
		if user.Id == u.Id {
			found = true
		}
	}

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_change_users_password(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Change-Password"
	profile["email"] = "john-change-password@example.com"
	profile["login"] = "john-change-password@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	//Sleep 1 second to make sure time has passed for password chagned timestamps
	time.Sleep(1 * time.Second)

	// Change the password to '1234Abcd' → POST /api/v1/users/{{userId}}/credentials/change_password
	op := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	np := &okta.PasswordCredential{
		Value: "1234Abcd",
	}
	npr := &okta.ChangePasswordRequest{
		OldPassword: op,
		NewPassword: np,
	}
	_, _, err = client.User.ChangePassword(user.Id, *npr, nil)
	require.NoError(t, err, "Could not change password")

	// Get the user and verify that 'passwordChanged' field has increased → GET /api/v1/users/{{userId}}/
	ubid, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Login")
	assert.True(t, ubid.PasswordChanged.After(*user.PasswordChanged), "Appears that password change did not happen")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_get_reset_password_link_for_user(t *testing.T) {
	client := tests.NewClient()
	// Create user with credentials → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Get-Reset-Password-Url"
	profile["email"] = "john-get-reset-password-url@example.com"
	profile["login"] = "john-get-reset-password-url@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Reset the user password → POST /api/v1/users/{{userId}}/lifecycle/reset_password?sendEmail=false
	rpt, _, err := client.User.ResetPassword(user.Id, query.NewQueryParams(query.WithSendEmail(false)))
	require.NoError(t, err, "Could not reset password")

	// Verify that the response returned contains the reset password link
	assert.IsType(t, &okta.ResetPasswordToken{}, rpt)
	assert.NotEmpty(t, rpt.ResetPasswordUrl, "Reset Password is not set")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_expire_a_users_password_and_get_a_temp_one(t *testing.T) {
	client := tests.NewClient()
	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Expire-Password"
	profile["email"] = "john-expire-password@example.com"
	profile["login"] = "john-expire-password@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Expire the user password → POST /api/v1/users/{{userId}}/lifecycle/expire_password?tempPassword=true
	ep, _, err := client.User.ExpirePassword(user.Id, query.NewQueryParams(query.WithTempPassword(true)))
	require.NoError(t, err, "Could not reset password")

	// Verify that the returned response contains a temporary password
	assert.IsType(t, &okta.TempPassword{}, ep)
	assert.NotEmpty(t, ep.TempPassword, "Temp Password not provided")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_change_user_recovery_question(t *testing.T) {
	client := tests.NewClient()
	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Change-Recovery-Question"
	profile["email"] = "john-change-recovery-question@example.com"
	profile["login"] = "john-change-recovery-question@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Update the user's recovery question → POST /api/v1/users/{{userId}}/credentials/change_recovery_question
	nucp := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	nucrq := &okta.RecoveryQuestionCredential{
		Question: "How many roads must a man walk down?",
		Answer:   "forty two",
	}
	nuc := &okta.UserCredentials{
		Password:         nucp,
		RecoveryQuestion: nucrq,
	}
	tmpuc, _, err := client.User.ChangeRecoveryQuestion(user.Id, *nuc, nil)
	require.NoError(t, err, "Could not change recovery question")
	assert.IsType(t, &okta.UserCredentials{}, tmpuc)

	// Update the user's password using updated recovery question credentials passing the body below → POST /api/v1/users/{{userId}}/credentials/forgot_password
	np := &okta.PasswordCredential{
		Value: "1234Abcd",
	}
	rq := &okta.RecoveryQuestionCredential{
		Answer: "forty two",
	}
	ucfp := &okta.UserCredentials{
		Password:         np,
		RecoveryQuestion: rq,
	}
	_, _, err = client.User.ForgotPassword(user.Id, *ucfp, nil)
	require.NoError(t, err, "Could not change password with recovery question")

	// Get the user and verify that 'passwordChanged' field has increased → GET /api/v1/users/{{userId}}
	ubid, _, err := client.User.GetUser(user.Id, nil)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Login")
	assert.True(t, ubid.PasswordChanged.After(*user.PasswordChanged), "Appears that password change did not happen")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_assign_a_user_to_a_role(t *testing.T) {
	client := tests.NewClient()
	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Role"
	profile["email"] = "john-role@example.com"
	profile["login"] = "john-role@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Add 'USER_ADMIN' role to the user → POST /api/v1/users/{{userId}}/roles (Body → { type: 'USER_ADMIN'  })
	r := &okta.Role{
		Type: "USER_ADMIN",
	}
	_, _, err = client.User.AddRoleToUser(user.Id, *r, nil)
	require.NoError(t, err, "Should not have had an error when adding role to user")

	// List roles for the user and verify added role → GET /api/v1/users/{{userId}}/roles
	roles, _, err := client.User.ListAssignedRoles(user.Id, nil)
	found := false
	roleId := ""
	for _, role := range roles {
		if role.Type == "USER_ADMIN" {
			found = true
			roleId = role.Id
		}
	}
	assert.True(t, found, "Could not verify USER_ADMIN was added to the user")

	// Remove role for the user → DELETE /api/v1/users/{{userId}}//roles/{{roleId}}/
	_, err = client.User.RemoveRoleFromUser(user.Id, roleId, nil)
	require.NoError(t, err, "Should not have had an error when removing role to user")

	// List roles for user and verify role was removed → GET /api/v1/users/{{userId}}/roles
	roles, _, err = client.User.ListAssignedRoles(user.Id, nil)
	found = false
	roleId = ""
	for _, role := range roles {
		if role.Type == "USER_ADMIN" {
			found = true
			roleId = role.Id
		}
	}
	assert.False(t, found, "Could not verify USER_ADMIN was removed to the user")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, resp, err := client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should not have been able to find user")
}

func Test_user_group_target_role(t *testing.T) {
	client := tests.NewClient()
	// Create a user with credentials, activated by default → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Group-Target"
	profile["email"] = "john-group-target@example.com"
	profile["login"] = "john-group-target@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "Group-Target Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.CreateGroup(*g, nil)
	require.NoError(t, err, "Creating an group should not error")

	// Add 'USER_ADMIN' role to the user → POST /api/v1/users/{{userId}}/roles (Body → { type: 'USER_ADMIN'  })
	r := &okta.Role{
		Type: "USER_ADMIN",
	}
	r, _, err = client.User.AddRoleToUser(user.Id, *r, nil)
	require.NoError(t, err, "Should not have had an error when adding role to user")

	// Add Group Target to 'USER_ADMIN' role → PUT /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups/{{groupId}}
	resp, err := client.User.AddGroupTargetToRole(user.Id, r.Id, group.Id, nil)
	require.NoError(t, err, "Should not have had an error when adding group target to role")

	// List Group Targets for role → GET  /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups
	groups, _, err := client.User.ListGroupTargetsForRole(user.Id, r.Id, nil)
	found := false
	for _, tmpgroup := range groups {
		if tmpgroup.Id == group.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not verify group target")

	//Remove Group Target from Admin User Role and verify removed → DELETE /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups/{{groupId}}
	gp = &okta.GroupProfile{
		Name: "TMP - Group-Target Test Group",
	}
	g = &okta.Group{
		Profile: gp,
	}
	newgroup, _, err := client.CreateGroup(*g, nil)
	_, err = client.User.AddGroupTargetToRole(user.Id, r.Id, newgroup.Id, nil)
	_, err = client.User.RemoveGroupTargetFromRole(user.Id, r.Id, group.Id, nil)
	require.NoError(t, err, "Should not have had an error when removing group target to role")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, resp, err = client.User.GetUser(user.Id, nil)
	require.Error(t, err, "User should not exist, but does")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should not have been able to find user")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	client.DeleteGroup(group.Id, nil)
	client.DeleteGroup(newgroup.Id, nil)
}

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
	"net/http"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_get_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Get the user by ID → GET /api/v1/users/{{userId}}
	ubid, _, err := client.User.GetUser(ctx, user.Id)
	require.NoError(t, err, "Getting a user by id should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Id")

	// Get the user by login name → GET /api/v1/users/{{login}}
	ubln, _, err := client.User.GetUser(ctx, profile["login"].(string))
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubln.Id, "Could not find user by Login")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_activate_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	// Create user with credentials → POST /api/v1/users?activate=false
	uc := &okta.UserCredentials{
		Password: &okta.PasswordCredential{
			Value: "Abcd1234",
		},
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Activate"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Activate the user → POST /api/v1/users/{{userId}}/lifecycle/activate?sendEmail=false
	token, _, err := client.User.ActivateUser(ctx, user.Id, query.NewQueryParams(query.WithSendEmail(false)))
	require.NoError(t, err, "Could not activate the user")
	assert.NotEmpty(t, token, "Token was not provided")
	assert.IsType(t, &okta.UserActivationToken{}, token, "Activation did not return correct type")

	crUser, _, err := client.User.GetUser(ctx, user.Id)
	require.NoError(t, err, "Could not get user by ID")
	assert.NotNil(t, crUser.Activated, "users activation time is missing")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_update_user_profile(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Update the user's profile by adding a nickname → PUT /api/v1/users/{{userId}}
	newProfile := *user.Profile
	newProfile["nickName"] = "Batman"
	updatedUser := &okta.User{
		Profile: &newProfile,
	}
	_, _, err = client.User.UpdateUser(ctx, user.Id, *updatedUser, nil)
	require.NoError(t, err, "Could not update the user")

	// Verify that user profile is updated by calling get on the user → GET /api/v1/users/{{userId}}
	tmpUser, _, err := client.User.GetUser(ctx, user.Id)
	require.NoError(t, err, "User was not available to get")
	tmpProfile := *tmpUser.Profile
	assert.Equal(t, "Batman", tmpProfile["nickName"])
	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_suspend_a_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
	// Create user with credentials → POST /api/v1/users?activate=true
	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Suspend"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Suspend the user → POST /api/v1/users/{{userId}}/lifecycle/suspend
	_, err = client.User.SuspendUser(ctx, user.Id)
	require.NoError(t, err, "Could not suspend the user")

	// Verify that user is in the list of suspended users → GET /api/v1/users?filter=status eq "SUSPENDED"
	filter := query.NewQueryParams(query.WithFilter("status eq \"SUSPENDED\""))
	users, _, err := client.User.ListUsers(ctx, filter)
	require.NoError(t, err, "Could not get suspended users")
	found := false
	for _, u := range users {
		if user.Id == u.Id {
			found = true
		}
	}
	assert.True(t, found, "The user was not found")

	// Unsuspend the user →  POST /api/v1/users/{{userId}}/lifecycle/unsuspend
	_, err = client.User.UnsuspendUser(ctx, user.Id)
	require.NoError(t, err, "Could not unsuspend the user")

	// Verify that user is in the list of active users → GET /api/v1/users?filter=status eq "ACTIVE"
	filter = query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))
	users, _, err = client.User.ListUsers(ctx, filter)
	require.NoError(t, err, "Could not get active users")
	found = false
	for _, u := range users {
		if user.Id == u.Id {
			found = true
		}
	}
	assert.True(t, found, "The user was not found")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_change_users_password(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Sleep 1 second to make sure time has passed for password chagned timestamps
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
	_, _, err = client.User.ChangePassword(ctx, user.Id, *npr, nil)
	require.NoError(t, err, "Could not change password")

	// Get the user and verify that 'passwordChanged' field has increased → GET /api/v1/users/{{userId}}/
	ubid, _, err := client.User.GetUser(ctx, user.Id)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Login")
	assert.True(t, ubid.PasswordChanged.After(*user.PasswordChanged), "Appears that password change did not happen")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_get_reset_password_link_for_user(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Reset the user password → POST /api/v1/users/{{userId}}/lifecycle/reset_password?sendEmail=false
	rpt, _, err := client.User.ResetPassword(ctx, user.Id, query.NewQueryParams(query.WithSendEmail(false)))
	require.NoError(t, err, "Could not reset password")

	// Verify that the response returned contains the reset password link
	assert.IsType(t, &okta.ResetPasswordToken{}, rpt)
	assert.NotEmpty(t, rpt.ResetPasswordUrl, "Reset Password is not set")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_expire_a_users_password_and_get_a_temp_one(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Expire the user password → POST /api/v1/users/{{userId}}/lifecycle/expire_password?tempPassword=true
	ep, _, err := client.User.ExpirePasswordAndGetTemporaryPassword(ctx, user.Id)
	require.NoError(t, err, "Could not reset password")

	// Verify that the returned response contains a temporary password
	assert.IsType(t, &okta.TempPassword{}, ep)
	assert.NotEmpty(t, ep.TempPassword, "Temp Password not provided")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_change_user_recovery_question(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

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
	tmpuc, _, err := client.User.ChangeRecoveryQuestion(ctx, user.Id, *nuc)
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
	_, _, err = client.User.ForgotPasswordSetNewPassword(ctx, user.Id, *ucfp, nil)
	require.NoError(t, err, "Could not change password with recovery question")

	// Get the user and verify that 'passwordChanged' field has increased → GET /api/v1/users/{{userId}}
	ubid, _, err := client.User.GetUser(ctx, user.Id)
	require.NoError(t, err, "Getting a user by login should not error")
	assert.Equal(t, user.Id, ubid.Id, "Could not find user by Login")
	assert.True(t, ubid.PasswordChanged.After(*user.PasswordChanged), "Appears that password change did not happen")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, _, err = client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
}

func Test_can_assign_a_user_to_a_role(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Add 'USER_ADMIN' role to the user → POST /api/v1/users/{{userId}}/roles (Body → { type: 'USER_ADMIN'  })
	arr := &okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}
	_, _, err = client.User.AssignRoleToUser(ctx, user.Id, *arr, nil)
	require.NoError(t, err, "Should not have had an error when adding role to user")

	// List roles for the user and verify added role → GET /api/v1/users/{{userId}}/roles
	roles, _, err := client.User.ListAssignedRolesForUser(ctx, user.Id, nil)
	require.NoError(t, err)
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
	_, err = client.User.RemoveRoleFromUser(ctx, user.Id, roleId)
	require.NoError(t, err, "Should not have had an error when removing role to user")

	// List roles for user and verify role was removed → GET /api/v1/users/{{userId}}/roles
	roles, _, err = client.User.ListAssignedRolesForUser(ctx, user.Id, nil)
	require.NoError(t, err)
	found = false
	for _, role := range roles {
		if role.Type == "USER_ADMIN" {
			found = true
		}
	}
	assert.False(t, found, "Could not verify USER_ADMIN was removed to the user")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, resp, err := client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should not have been able to find user")
}

func Test_user_group_target_role(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)
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
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	// Create a new group → POST /api/v1/groups
	gp := &okta.GroupProfile{
		Name: "SDK_TEST Group-Target Test Group",
	}
	g := &okta.Group{
		Profile: gp,
	}
	group, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err, "Creating an group should not error")

	// Add 'USER_ADMIN' role to the user → POST /api/v1/users/{{userId}}/roles (Body → { type: 'USER_ADMIN'  })
	arr := &okta.AssignRoleRequest{
		Type: "USER_ADMIN",
	}
	r, _, err := client.User.AssignRoleToUser(ctx, user.Id, *arr, nil)
	require.NoError(t, err, "Should not have had an error when adding role to user")

	// Add Group Target to 'USER_ADMIN' role → PUT /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups/{{groupId}}
	_, err = client.User.AddGroupTargetToRole(ctx, user.Id, r.Id, group.Id)
	require.NoError(t, err, "Should not have had an error when adding group target to role")

	// List Group Targets for role → GET  /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups
	groups, _, err := client.User.ListGroupTargetsForRole(ctx, user.Id, r.Id, nil)
	require.NoError(t, err)
	found := false
	for _, tmpgroup := range groups {
		if tmpgroup.Id == group.Id {
			found = true
		}
	}
	assert.True(t, found, "Could not verify group target")

	// Remove Group Target from Admin User Role and verify removed → DELETE /api/v1/users/{{userId}}/roles/{{roleId}}/targets/groups/{{groupId}}
	gp = &okta.GroupProfile{
		Name: "SDK_TEST TMP - Group-Target Test Group",
	}
	g = &okta.Group{
		Profile: gp,
	}
	newgroup, _, err := client.Group.CreateGroup(ctx, *g)
	require.NoError(t, err)
	_, err = client.User.AddGroupTargetToRole(ctx, user.Id, r.Id, newgroup.Id)
	require.NoError(t, err)
	_, err = client.User.RemoveGroupTargetFromRole(ctx, user.Id, r.Id, group.Id)
	require.NoError(t, err, "Should not have had an error when removing group target to role")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Verify that the user is deleted by calling get on user (Exception thrown with 404 error message) → GET /api/v1/users/{{userId}}
	_, resp, err := client.User.GetUser(ctx, user.Id)
	require.Error(t, err, "User should not exist, but does")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Should not have been able to find user")

	// Delete the group → DELETE /api/v1/groups/{{groupId}}
	client.Group.DeleteGroup(ctx, group.Id)
	client.Group.DeleteGroup(ctx, newgroup.Id)
}

func Test_can_get_user_with_cache_enabled(t *testing.T) {
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
	profile["lastName"] = "Test-Cache"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	createdUser, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	for i := 0; i < 50; i++ {
		user, resp, err := client.User.GetUser(ctx, profile["email"].(string))
		assert.NoError(t, err, "Should not error when getting user")
		assert.NotNil(t, user, "user should not be nil")
		assert.NotNil(t, resp, "resp should not be nil")
	}

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, createdUser.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, createdUser.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

func Test_can_paginate_across_users(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err)

	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile1 := okta.UserProfile{}
	profile1["firstName"] = "John"
	profile1["lastName"] = "page-test"
	profile1["email"] = "john-page-1@example.com"
	profile1["login"] = "SDK_TESTjohn-page-1@example.com"
	u1 := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile1,
	}
	profile2 := okta.UserProfile{}
	profile2["firstName"] = "John"
	profile2["lastName"] = "page-test"
	profile2["email"] = "john-page-2@example.com"
	profile2["login"] = "SDK_TESTjohn-page-2@example.com"
	u2 := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile2,
	}
	qp := query.NewQueryParams(query.WithActivate(true))

	createdUser1, _, err := client.User.CreateUser(ctx, *u1, qp)
	require.NoError(t, err, "Creating a new user should not error")
	createdUser2, _, err := client.User.CreateUser(ctx, *u2, qp)
	require.NoError(t, err, "Creating a new user should not error")

	query := query.NewQueryParams(query.WithLimit(1))
	user1, resp, err := client.User.ListUsers(ctx, query)
	require.NoError(t, err)
	assert.Equal(t, 1, len(user1), "User1 did not reutrn 1 user")
	user1Profile := *user1[0].Profile

	hasNext := resp.HasNextPage()
	assert.True(t, hasNext, "Should return true for HasNextPage")

	var user2 []*okta.User
	_, err = resp.Next(ctx, &user2)
	require.NoError(t, err)

	assert.Equal(t, 1, len(user2), "User2 did not reutrn 1 user")
	user2Profile := *user2[0].Profile

	assert.NotEqual(t, user2Profile["email"], user1Profile["email"], "Emails should not be the same")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, createdUser1.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, createdUser1.Id, nil)
	require.NoError(t, err, "Should not error when deleting")

	// Deactivate the user → POST /api/v1/users/{{userId}}/lifecycle/deactivate
	_, err = client.User.DeactivateUser(ctx, createdUser2.Id, nil)
	require.NoError(t, err, "Should not error when deactivating")

	// Delete the user → DELETE /api/v1/users/{{userId}}
	_, err = client.User.DeactivateOrDeleteUser(ctx, createdUser2.Id, nil)
	require.NoError(t, err, "Should not error when deleting")
}

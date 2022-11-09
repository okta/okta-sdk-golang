package okta

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUser(activate bool) (*User, *APIResponse, CreateUserRequest, error) {
	req := apiClient.UserApi.CreateUser(apiClient.cfg.Context)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	req = req.Body(body)
	req = req.Activate(activate)
	user, res, err := req.Execute()
	return user, res, body, err
}

func cleanUpUser(userId string) error {
	_, err := apiClient.UserApi.DeactivateUser(apiClient.cfg.Context, userId).Execute()
	if err != nil {
		return err
	}
	_, err = apiClient.UserApi.DeleteUser(apiClient.cfg.Context, userId).Execute()
	if err != nil {
		return err
	}
	_, res, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, userId).Execute()
	if err != nil && res.StatusCode == http.StatusNotFound {
		err = nil
	}
	return err
}

func Test_Get_User(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("get user by id", func(t *testing.T) {
		ubid, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.Equal(t, user.GetId(), ubid.GetId())
	})
	t.Run("get user by login", func(t *testing.T) {
		ubid, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, *user.Profile.Login).Execute()
		require.NoError(t, err, "Could not get user by login")
		assert.Equal(t, user.GetId(), ubid.GetId())
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Activate_User(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(false)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("activate users", func(t *testing.T) {
		req := apiClient.UserApi.ActivateUser(apiClient.cfg.Context, user.GetId())
		req = req.SendEmail(false)
		token, _, err := req.Execute()
		require.NoError(t, err, "Could not activate the user")
		assert.NotEmpty(t, token, "Token was not provided")
		assert.IsType(t, &UserActivationToken{}, token, "Activation did not return correct type")
	})
	t.Run("get user by id", func(t *testing.T) {
		crUser, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.NotNil(t, crUser.Activated, "users activation time is missing")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Update_User_Profile(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	nickName := "Batman"
	t.Run("update user", func(t *testing.T) {
		newProfile := user.Profile
		newProfile.NickName = &nickName
		req := apiClient.UserApi.UpdateUser(apiClient.cfg.Context, user.GetId())
		body := UpdateUserRequest{Profile: newProfile}
		req = req.User(body)
		_, _, err := req.Execute()
		require.NoError(t, err, "Could not update user by ID")
	})
	t.Run("get user", func(t *testing.T) {
		updatedUser, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.Equal(t, nickName, *updatedUser.Profile.NickName)
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Suspend_User(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("suspend user", func(t *testing.T) {
		_, err := apiClient.UserApi.SuspendUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not suspend user")
		susReq := apiClient.UserApi.ListUsers(apiClient.cfg.Context)
		susReq = susReq.Filter("status eq \"SUSPENDED\"")
		listUsers, _, err := susReq.Execute()
		require.NoError(t, err, "Could not get suspended user")
		var found bool
		for _, u := range listUsers {
			if user.GetId() == u.GetId() {
				found = true
				break
			}
		}
		assert.True(t, found, "The user was not found")
	})
	t.Run("unsuspend user", func(t *testing.T) {
		_, err := apiClient.UserApi.UnsuspendUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not unsuspend user")
		unsusReq := apiClient.UserApi.ListUsers(apiClient.cfg.Context)
		unsusReq = unsusReq.Filter("status eq \"ACTIVE\"")
		listUsers, _, err := unsusReq.Execute()
		require.NoError(t, err, "Could not get active user")
		var found bool
		for _, u := range listUsers {
			if user.GetId() == u.GetId() {
				found = true
				break
			}
		}
		assert.True(t, found, "The user was not found")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Change_User_Password(t *testing.T) {
	t.Parallel()
	user, _, payload, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	time.Sleep(1 * time.Second)
	t.Run("change users password", func(t *testing.T) {
		req := apiClient.UserApi.ChangePassword(apiClient.cfg.Context, *user.Id)
		newPassword := NewPasswordCredential()
		newPassword.SetValue(testPassword(10))
		payload := ChangePasswordRequest{
			OldPassword: payload.GetCredentials().Password,
			NewPassword: newPassword,
		}
		req = req.ChangePasswordRequest(payload)
		_, _, err := req.Execute()
		require.NoError(t, err, "Could not change user password")
		ubid, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.Equal(t, user.GetId(), ubid.GetId())
		assert.True(t, ubid.HasPasswordChanged())
		assert.True(t, ubid.GetPasswordChanged().After(user.GetPasswordChanged()), "Password has not changed")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Get_Reset_Password_Link_User(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("reset password", func(t *testing.T) {
		req := apiClient.UserApi.GenerateResetPasswordToken(apiClient.cfg.Context, user.GetId())
		req = req.SendEmail(false)
		rpt, _, err := req.Execute()
		require.NoError(t, err, "Could not reset password")
		assert.IsType(t, &ResetPasswordToken{}, rpt)
		assert.NotEmpty(t, rpt.ResetPasswordUrl, "Reset Password is not set")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Expire_Password_User_Get_Temp(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("expire password", func(t *testing.T) {
		tp, _, err := apiClient.UserApi.ExpirePasswordAndGetTemporaryPassword(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not reset password")
		assert.IsType(t, &TempPassword{}, tp)
		assert.NotEmpty(t, tp.TempPassword, "Temp Password not provided")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Change_User_Recovery_Question(t *testing.T) {
	t.Parallel()
	user, _, body, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("change recovery question", func(t *testing.T) {
		req := apiClient.UserApi.ChangeRecoveryQuestion(apiClient.cfg.Context, user.GetId())
		payload := UserCredentials{
			Password:         body.GetCredentials().Password,
			RecoveryQuestion: testFactory.NewValidTestRecoveryQuestionCredential(),
		}
		req = req.UserCredentials(payload)
		tmpuc, _, err := req.Execute()
		require.NoError(t, err, "Could not change recovery question")
		assert.IsType(t, &UserCredentials{}, tmpuc)
	})
	t.Run("update password using recovery question", func(t *testing.T) {
		req := apiClient.UserApi.ForgotPasswordSetNewPassword(apiClient.cfg.Context, user.GetId())
		payload := UserCredentials{
			Password:         testFactory.NewValidTestPasswordCredential(),
			RecoveryQuestion: testFactory.NewValidTestRecoveryQuestionCredential(),
		}
		req = req.UserCredentials(payload)
		_, _, err := req.Execute()
		require.NoError(t, err, "Could not change password with recovery question")
		ubid, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.Equal(t, user.GetId(), ubid.GetId())
		assert.True(t, ubid.GetPasswordChanged().After(user.GetPasswordChanged()), "Password change did not happen")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Assign_User_To_A_Role(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	var roleId string
	role := ROLETYPE_USER_ADMIN
	t.Run("add role to user", func(t *testing.T) {
		req := apiClient.RoleAssignmentApi.AssignRoleToUser(apiClient.cfg.Context, user.GetId())
		payload := AssignRoleRequest{
			Type: &role,
		}
		req = req.AssignRoleRequest(payload)
		_, _, err = req.Execute()
		require.NoError(t, err, "Should not have had an error when adding role to user")
		listRoles, _, err := apiClient.RoleAssignmentApi.ListAssignedRolesForUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Should not have had an error when getting user's assigned role")
		var found bool
		for _, r := range listRoles {
			if r.GetType() == role {
				found = true
				roleId = r.GetId()
				break
			}
		}
		assert.True(t, found, "Could not verify USER_ADMIN was added to the user")
	})
	t.Run("remove role from user", func(t *testing.T) {
		_, err = apiClient.RoleAssignmentApi.UnassignRoleFromUser(apiClient.cfg.Context, user.GetId(), roleId).Execute()
		require.NoError(t, err, "Should not have had an error when removing role to user")
		listRoles, _, err := apiClient.RoleAssignmentApi.ListAssignedRolesForUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Should not have had an error when getting user's assigned role")
		var found bool
		for _, r := range listRoles {
			if r.GetType() == role {
				found = true
				break
			}
		}
		assert.False(t, found, "Could not verify USER_ADMIN was removed to the user")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_User_Group_Target_Role(t *testing.T) {
	t.Parallel()
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	var groupId string
	var roleId string
	var newGroupId string
	t.Run("add group target to role", func(t *testing.T) {
		greq := apiClient.GroupApi.CreateGroup(apiClient.cfg.Context)
		gp := NewGroupProfile()
		gp.SetName("SDK_TEST Group-Target Test Group")
		gpayload := Group{Profile: gp}
		greq = greq.Group(gpayload)
		group, _, err := greq.Execute()
		require.NoError(t, err, "Creating an group should not error")
		areq := apiClient.RoleAssignmentApi.AssignRoleToUser(apiClient.cfg.Context, user.GetId())
		payload := NewAssignRoleRequest()
		payload.SetType("USER_ADMIN")
		areq = areq.AssignRoleRequest(*payload)
		role, _, err := areq.Execute()
		require.NoError(t, err, "Should not have had an error when adding role to user")
		_, err = apiClient.RoleTargetApi.AssignGroupTargetToUserRole(apiClient.cfg.Context, user.GetId(), role.GetId(), group.GetId()).Execute()
		require.NoError(t, err, "Should not have had an error when adding group target to role")
		groups, _, err := apiClient.RoleTargetApi.ListGroupTargetsForRole(apiClient.cfg.Context, user.GetId(), role.GetId()).Execute()
		require.NoError(t, err)
		var found bool
		for _, tmpgroup := range groups {
			if tmpgroup.GetId() == group.GetId() {
				found = true
				groupId = group.GetId()
				roleId = role.GetId()
				break
			}
		}
		assert.True(t, found, "Could not verify group target")
	})
	t.Run("remove group target from role", func(t *testing.T) {
		greq := apiClient.GroupApi.CreateGroup(apiClient.cfg.Context)
		gp := NewGroupProfile()
		gp.SetName("SDK_TEST Group TMP-Target Test Group")
		gpayload := Group{Profile: gp}
		greq = greq.Group(gpayload)
		newGroup, _, err := greq.Execute()
		require.NoError(t, err, "Should not have had an error when adding role to user")
		newGroupId = newGroup.GetId()
		_, err = apiClient.RoleTargetApi.AssignGroupTargetToUserRole(apiClient.cfg.Context, user.GetId(), roleId, newGroup.GetId()).Execute()
		require.NoError(t, err)
		_, err = apiClient.RoleTargetApi.UnassignGroupTargetFromUserAdminRole(apiClient.cfg.Context, user.GetId(), roleId, groupId).Execute()
		require.NoError(t, err, "Should not have had an error when removing group target to role")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
	err = cleanUpGroup(groupId)
	require.NoError(t, err, "Clean up group should not error")
	err = cleanUpGroup(newGroupId)
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Get_User_With_Cache_Enabled(t *testing.T) {
	t.Parallel()
	configuration := NewConfiguration()
	configuration.Debug = true
	cachedApiClient := NewAPIClient(configuration)
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("get user with cache", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			u, resp, err := cachedApiClient.UserApi.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
			assert.NoError(t, err, "Should not error when getting user")
			assert.NotNil(t, u, "user should not be nil")
			assert.NotNil(t, resp, "resp should not be nil")
		}
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_List_User_Subscriptions(t *testing.T) {
	t.Parallel()
	user, _, err := apiClient.UserApi.GetUser(apiClient.cfg.Context, "me").Execute()
	require.NoError(t, err, "Getting the current user should not error")
	t.Run("get user subscription", func(t *testing.T) {
		subscriptions, _, err := apiClient.SubscriptionApi.ListUserSubscriptions(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Should not error listing user subscriptions")
		assert.True(t, len(subscriptions) > 0, "User should have subscriptions")
	})
	t.Run("get user subscription by notification type", func(t *testing.T) {
		expectedNotificationType := "OKTA_ANNOUNCEMENT"
		subscription, _, err := apiClient.SubscriptionApi.ListUserSubscriptionsByNotificationType(apiClient.cfg.Context, user.GetId(), expectedNotificationType).Execute()
		require.NoError(t, err, "Should not error getting user subscription by notification types")
		assert.Equal(t, subscription.GetNotificationType(), NOTIFICATIONTYPE_OKTA_ANNOUNCEMENT, "User should have subscription notification type %q, got %q", expectedNotificationType, subscription.NotificationType)
	})
}

func TestCanPaginateAcrossUsers(t *testing.T) {
	createdUser1, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	createdUser2, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	user1, resp, err := apiClient.UserApi.ListUsers(apiClient.cfg.Context).Limit(1).Execute()
	require.NoError(t, err)
	assert.Equal(t, 1, len(user1), "User1 did not return 1 user")
	user1Profile := user1[0].GetProfile()
	hasNext := resp.HasNextPage()
	assert.True(t, hasNext, "Should return true for HasNextPage")
	var user2 []User
	res, err := resp.Next(&user2)
	require.NoError(t, err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	require.NoError(t, err)
	assert.NotEmpty(t, string(body), "body is empty")
	assert.Equal(t, 1, len(user2), "User2 did not return 1 user")
	user2Profile := user2[0].GetProfile()
	assert.NotEqual(t, user2Profile.GetEmail(), user1Profile.GetEmail(), "Emails should not be the same")
	err = cleanUpUser(createdUser1.GetId())
	require.NoError(t, err, "Should not error when deactivating")
	err = cleanUpUser(createdUser2.GetId())
	require.NoError(t, err, "Should not error when deactivating")
}

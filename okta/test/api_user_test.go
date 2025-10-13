/*
Okta Admin Management

Testing UserAPIService

*/

package okta

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	okta "github.com/okta/okta-sdk-golang/v6/okta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	apiClient   *okta.APIClient
	testDataMgr *okta.TestDataManager
	testContext context.Context
)

// TestMain provides setup and teardown for all user tests
func TestMain(m *testing.M) {
	if err := okta.ValidateTestEnvironment(); err != nil {
		log.Printf("Test environment validation failed: %v", err)
		log.Printf("Please set OKTA_CLIENT_ORGURL and OKTA_CLIENT_TOKEN environment variables")
		os.Exit(1)
	}

	var err error
	configuration, err := okta.NewConfiguration()
	if err != nil {
		log.Fatalf("Failed to create configuration: %v", err)
	}

	apiClient = okta.NewAPIClient(configuration)
	testDataMgr = okta.GetTestDataManager()
	testContext = context.Background()

	exitCode := m.Run()

	log.Println("Cleaning up test resources...")
	testDataMgr.CleanupAllTestUsers()

	os.Exit(exitCode)
}

func Test_okta_UserAPIService(t *testing.T) {

	t.Run("Test UserAPIService CreateUser", func(t *testing.T) {
		testFactory := &okta.TestFactory{}
		profile := testFactory.NewValidTestUserProfile()
		credentials := testFactory.NewValidTestUserCredentialsWithPassword()

		createReq := okta.NewCreateUserRequest(profile)
		createReq.SetCredentials(*credentials)

		req := apiClient.UserAPI.CreateUser(testContext).Body(*createReq).Activate(false)
		resp, httpRes, err := req.Execute()

		require.NoError(t, err, "CreateUser should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")

		assert.NotNil(t, resp.Id, "User ID should be set")
		assert.NotNil(t, resp.Profile, "User profile should be set")
		if resp.Profile != nil {
			assert.Equal(t, profile.GetFirstName(), resp.Profile.GetFirstName(), "First name should match")
			assert.Equal(t, profile.GetLastName(), resp.Profile.GetLastName(), "Last name should match")
			assert.Equal(t, profile.GetEmail(), resp.Profile.GetEmail(), "Email should match")
		}

		if resp.Id != nil {
			testDataMgr.TrackUser(*resp.Id)
		}
	})

	t.Run("Test UserAPIService GetUser", func(t *testing.T) {
		createdUser, err := testDataMgr.CreateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id, "Created user should have an ID")

		resp, httpRes, err := apiClient.UserAPI.GetUser(testContext, *createdUser.Id).Execute()

		require.NoError(t, err, "GetUser should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")

		assert.Equal(t, createdUser.Id, resp.Id, "User IDs should match")
		if resp.Profile != nil && createdUser.Profile != nil {
			assert.Equal(t, createdUser.Profile.GetEmail(), resp.Profile.GetEmail(), "Emails should match")
		}
	})

	t.Run("Test UserAPIService ListUsers", func(t *testing.T) {
		resp, httpRes, err := apiClient.UserAPI.ListUsers(testContext).Execute()

		require.NoError(t, err, "ListUsers should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")

		assert.IsType(t, []okta.User{}, resp, "Response should be a slice of User")
	})

	t.Run("Test UserAPIService UpdateUser", func(t *testing.T) {
		createdUser, err := testDataMgr.CreateAndActivateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id, "Created user should have an ID")

		testFactory := &okta.TestFactory{}
		updatedProfile := testFactory.NewTestUserProfileUpdate()
		updateReq := okta.NewUpdateUserRequest()
		updateReq.SetProfile(updatedProfile)

		req := apiClient.UserAPI.UpdateUser(testContext, *createdUser.Id).User(*updateReq)
		resp, httpRes, err := req.Execute()

		require.NoError(t, err, "UpdateUser should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")

		if resp.Profile != nil {
			assert.Equal(t, updatedProfile.GetFirstName(), resp.Profile.GetFirstName(), "First name should be updated")
			assert.Equal(t, updatedProfile.GetLastName(), resp.Profile.GetLastName(), "Last name should be updated")
		}
	})

	t.Run("Test UserAPIService ReplaceUser", func(t *testing.T) {
		createdUser, err := testDataMgr.CreateAndActivateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id, "Created user should have an ID")

		testFactory := &okta.TestFactory{}
		replacementProfile := testFactory.NewTestUserProfileUpdate()
		replaceReq := okta.NewUpdateUserRequest()
		replaceReq.SetProfile(replacementProfile)

		req := apiClient.UserAPI.ReplaceUser(testContext, *createdUser.Id).User(*replaceReq)
		resp, httpRes, err := req.Execute()

		require.NoError(t, err, "ReplaceUser should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")

		if resp.Profile != nil {
			assert.Equal(t, replacementProfile.GetFirstName(), resp.Profile.GetFirstName(), "First name should be replaced")
			assert.Equal(t, replacementProfile.GetLastName(), resp.Profile.GetLastName(), "Last name should be replaced")
		}
	})

	t.Run("Test UserAPIService DeleteUser", func(t *testing.T) {
		createdUser, err := testDataMgr.CreateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id, "Created user should have an ID")

		_, deactivateErr := apiClient.UserLifecycleAPI.DeactivateUser(testContext, *createdUser.Id).Execute()
		require.NoError(t, deactivateErr, "DeactivateUser should not return an error")

		// Wait a moment for deactivation to complete
		time.Sleep(2 * time.Second)

		httpRes, err := apiClient.UserAPI.DeleteUser(testContext, *createdUser.Id).Execute()

		require.NoError(t, err, "DeleteUser should not return an error")
		assert.Equal(t, 204, httpRes.StatusCode, "HTTP status should be 200")

		testDataMgr.RemoveUserFromTracking(*createdUser.Id)

		_, _, getErr := apiClient.UserAPI.GetUser(testContext, *createdUser.Id).Execute()
		assert.Error(t, getErr, "Getting deleted user should return an error")
	})

	// FIXME: This feature needs permission in the org and skipping this test.
	t.Run("Test UserAPIService ListUserBlocks", func(t *testing.T) {
		t.Skip()
		createdUser, err := testDataMgr.CreateAndActivateTestUser()
		require.NoError(t, err, "Failed to create test user")
		require.NotNil(t, createdUser.Id, "Created user should have an ID")

		resp, httpRes, err := apiClient.UserAPI.ListUserBlocks(testContext, *createdUser.Id).Execute()
		fmt.Println(httpRes.Body)

		require.NoError(t, err, "ListUserBlocks should not return an error")
		require.NotNil(t, resp, "Response should not be nil")
		assert.Equal(t, 200, httpRes.StatusCode, "HTTP status should be 200")
	})
}

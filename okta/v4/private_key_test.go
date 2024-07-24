package okta

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Private_Key_Request_Can_Create_User(t *testing.T) {
	if os.Getenv("OKTA_CCI") != "yes" {
		t.Skip("Skipping testing not in CI environment")
	}
	configuration, err := NewConfiguration(WithAuthorizationMode("PrivateKey"), WithScopes([]string{"okta.users.manage"}))
	require.NoError(t, err, "Creating a new config should not error")
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserAPI.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

func Test_JWT_Request_Can_Create_User(t *testing.T) {
	if os.Getenv("OKTA_CCI") != "yes" {
		t.Skip("Skipping testing not in CI environment")
	}
	configuration, err := NewConfiguration(WithAuthorizationMode("JWT"), WithScopes([]string{"okta.users.manage"}))
	require.NoError(t, err, "Creating a new config should not error")
	privateKeySigner, err := createKeySigner(configuration.Okta.Client.PrivateKey, configuration.Okta.Client.PrivateKeyId)
	require.NoError(t, err)
	clientAssertion, err := createClientAssertion(configuration.Okta.Client.OrgUrl, configuration.Okta.Client.ClientId, privateKeySigner)
	require.NoError(t, err)
	configuration.Okta.Client.ClientAssertion = clientAssertion
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserAPI.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

func Test_JWK_Request_Can_Create_User(t *testing.T) {
	if os.Getenv("OKTA_CCI") != "yes" {
		t.Skip("Skipping testing not in CI environment")
	}
	configuration, err := NewConfiguration(WithAuthorizationMode("JWK"), WithScopes([]string{"okta.users.manage"}), WithJWK(""), WithEncryptionType("RSA"))
	require.NoError(t, err, "Creating a new config should not error")
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserAPI.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

func Test_Dpop_Get_User(t *testing.T) {
	configuration, err := NewConfiguration(WithAuthorizationMode("PrivateKey"), WithScopes([]string{"okta.users.manage", "okta.users.read"}))
	require.NoError(t, err, "Creating a new config should not error")
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserAPI.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
	t.Run("get user by id", func(t *testing.T) {
		ubid, _, err := client.UserAPI.GetUser(apiClient.cfg.Context, user.GetId()).Execute()
		require.NoError(t, err, "Could not get user by ID")
		assert.Equal(t, user.GetId(), ubid.GetId())
	})
	t.Run("get user by login", func(t *testing.T) {
		ubid, _, err := client.UserAPI.GetUser(apiClient.cfg.Context, *user.Profile.Login).Execute()
		require.NoError(t, err, "Could not get user by login")
		assert.Equal(t, user.GetId(), ubid.GetId())
	})
	t.Run("list user", func(t *testing.T) {
		ubid, _, err := client.UserAPI.ListUsers(apiClient.cfg.Context).Limit(1).Execute()
		require.NoError(t, err, "Could not get user by login")
		assert.Equal(t, 1, len(ubid))
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Dpop_Pagination(t *testing.T) {
	configuration, err := NewConfiguration(WithAuthorizationMode("PrivateKey"), WithScopes([]string{"okta.users.manage", "okta.users.read"}))
	require.NoError(t, err, "Creating a new config should not error")
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	createdUser1, _, err := client.UserAPI.CreateUser(client.cfg.Context).Body(body).Activate(true).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	uc = testFactory.NewValidTestUserCredentialsWithPassword()
	profile = testFactory.NewValidTestUserProfile()
	body = CreateUserRequest{Credentials: uc, Profile: profile}
	createdUser2, _, err := client.UserAPI.CreateUser(client.cfg.Context).Body(body).Activate(true).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	user1, resp, err := client.UserAPI.ListUsers(client.cfg.Context).Limit(1).Execute()
	require.NoError(t, err)
	assert.Equal(t, 1, len(user1), "User1 did not return 1 user")
	user1Profile := user1[0].GetProfile()
	hasNext := resp.HasNextPage()
	assert.True(t, hasNext, "Should return true for HasNextPage")
	var user2 []User
	res, err := resp.Next(&user2)
	require.NoError(t, err)
	respBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	require.NoError(t, err)
	assert.NotEmpty(t, string(respBody), "body is empty")
	assert.Equal(t, 1, len(user2), "User2 did not return 1 user")
	user2Profile := user2[0].GetProfile()
	assert.NotEqual(t, user2Profile.GetEmail(), user1Profile.GetEmail(), "Emails should not be the same")
	err = cleanUpUser(createdUser1.GetId())
	require.NoError(t, err, "Should not error when deactivating")
	err = cleanUpUser(createdUser2.GetId())
	require.NoError(t, err, "Should not error when deactivating")
}

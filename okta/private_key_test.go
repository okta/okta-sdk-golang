package okta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Private_Key_Request_Can_Create_User(t *testing.T) {
	if os.Getenv("OKTA_CCI") != "yes" {
		t.Skip("Skipping testing not in CI environment")
	}
	configuration := NewConfiguration(WithAuthorizationMode("PrivateKey"), WithScopes([]string{"okta.users.manage"}))
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserApi.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

func Test_JWT_Request_Can_Create_User(t *testing.T) {
	if os.Getenv("OKTA_CCI") != "yes" {
		t.Skip("Skipping testing not in CI environment")
	}
	configuration := NewConfiguration(WithAuthorizationMode("JWT"), WithScopes([]string{"okta.users.manage"}))
	privateKeySigner, err := createKeySigner(configuration.Okta.Client.PrivateKey, configuration.Okta.Client.PrivateKeyId)
	require.NoError(t, err)
	clientAssertion, err := createClientAssertion(configuration.Okta.Client.OrgUrl, configuration.Okta.Client.ClientId, privateKeySigner)
	require.NoError(t, err)
	configuration.Okta.Client.ClientAssertion = clientAssertion
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserApi.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

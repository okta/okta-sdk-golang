package okta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Private_Key_Request_Can_Create_User(t *testing.T) {
	configuration := NewConfiguration(WithAuthorizationMode("PrivateKey"), WithScopes([]string{"okta.users.manage"}))
	client := NewAPIClient(configuration)
	uc := testFactory.NewValidTestUserCredentialsWithPassword()
	profile := testFactory.NewValidTestUserProfile()
	body := CreateUserRequest{Credentials: uc, Profile: profile}
	user, _, err := client.UserApi.CreateUser(apiClient.cfg.Context).Body(body).Execute()
	require.NoError(t, err, "Creating a new user should not error")
	assert.NotNil(t, user, "User should not be nil")
}

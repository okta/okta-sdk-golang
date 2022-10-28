package okta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Applications(t *testing.T) {
	t.Run("can get all applicaitons", func(t *testing.T) {
		req := apiClient.ApplicationApi.ListApplications(apiClient.cfg.Context).Limit(2)
		applications, _, _ := req.Execute()
		assert.Equal(t, 2, len(applications))
	})
}

func Test_OIDC_Application(t *testing.T) {
	t.Run("create an oidc application", func(t *testing.T) {
		// ctx := context.WithValue(
		// 	context.Background(),
		// 	ContextAPIKeys,
		// 	map[string]APIKey{
		// 		"api_token": {
		// 			Key:    "***",
		// 			Prefix: "SSWS",
		// 		},
		// 	},
		// )
	})
}

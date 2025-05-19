package okta

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Get_User_Schema(t *testing.T) {
	t.Run("get default user schema", func(t *testing.T) {
		schema, _, err := apiClient.SchemaAPI.GetUserSchema(apiClient.cfg.Context, "default").Execute()
		require.NoError(t, err, "Could not get default user schema")
		assert.NotEmpty(t, schema, "User schema is empty")
		assert.Equal(t, "Username", schema.Definitions.Base.Properties.Login.GetTitle())
		assert.Equal(t, "READ_WRITE", schema.Definitions.Base.Properties.Login.GetMutability())
		assert.Equal(t, "NONE", schema.Definitions.Base.Properties.Login.GetScope())
		assert.Equal(t, int32(5), schema.Definitions.Base.Properties.Login.GetMinLength())
		assert.Equal(t, int32(100), schema.Definitions.Base.Properties.Login.GetMaxLength())
		assert.NotEmpty(t, schema.Definitions.Base.Properties.Login.GetPermissions())
		assert.Equal(t, "SELF", schema.Definitions.Base.Properties.Login.GetPermissions()[0].GetPrincipal())
	})
}

func Test_Update_Property_To_User_Schema(t *testing.T) {
	schema, _, err := apiClient.SchemaAPI.GetUserSchema(apiClient.cfg.Context, "default").Execute()
	require.NoError(t, err)

	t.Run("remove all properties with SDK_TEST prefix by setting to null", func(t *testing.T) {
		req := apiClient.SchemaAPI.UpdateUserProfile(apiClient.cfg.Context, "default")

		properties := schema.Definitions.Custom.GetProperties()
		updatedProperties := make(map[string]*UserSchemaAttribute)

		for key := range properties {
			if strings.HasPrefix(key, "SDK_TEST") {
				// Set to nil to remove property
				updatedProperties[key] = nil
			} else {
				// Keep existing property
				updatedProperties[key] = properties[key]
			}
		}

		schema.Definitions.Custom.SetProperties(updatedProperties)

		req = req.UserSchema(*schema)
		_, _, err = req.Execute()
		require.NoError(t, err)
	})
}

func convertStringToMap(jsonString *string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(*jsonString), &data) // Dereference the pointer
	if err != nil {
		return nil, err
	}
	return data, nil
}

func convertMapToString(data map[string]interface{}) (*string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	jsonString := string(jsonBytes)
	return &jsonString, nil
}

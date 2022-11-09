package okta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Get_User_Schema(t *testing.T) {
	t.Parallel()
	t.Run("get default user schema", func(t *testing.T) {
		schema, _, err := apiClient.SchemaApi.GetUserSchema(apiClient.cfg.Context, "default").Execute()
		require.NoError(t, err, "Could not get default user schema")
		assert.NotEmpty(t, schema, "User schema is empty")
		assert.Equal(t, "Username", schema.Definitions.Base.Properties.Login.GetTitle())
		assert.Equal(t, "READ_WRITE", schema.Definitions.Base.Properties.Login.GetMutability())
		assert.Equal(t, USERSCHEMAATTRIBUTESCOPE_NONE, schema.Definitions.Base.Properties.Login.GetScope())
		assert.Equal(t, int32(5), schema.Definitions.Base.Properties.Login.GetMinLength())
		assert.Equal(t, int32(100), schema.Definitions.Base.Properties.Login.GetMaxLength())
		assert.NotEmpty(t, schema.Definitions.Base.Properties.Login.GetPermissions())
		assert.Equal(t, "SELF", schema.Definitions.Base.Properties.Login.GetPermissions()[0].GetPrincipal())
	})
}

func Test_Update_Property_To_User_Schema(t *testing.T) {
	t.Parallel()
	schema, _, err := apiClient.SchemaApi.GetUserSchema(apiClient.cfg.Context, "default").Execute()
	require.NoError(t, err, "Could not get default user schema")
	assert.NotEmpty(t, schema, "User schema is empty")
	t.Run("get update user schema", func(t *testing.T) {
		req := apiClient.SchemaApi.UpdateUserProfile(apiClient.cfg.Context, "default")
		customAttributeName := testPrefix + randomTestString()
		customAttributeDetail := UserSchemaAttribute{}
		customAttributeDetail.SetTitle(customAttributeName)
		customAttributeDetail.SetType(USERSCHEMAATTRIBUTETYPE_STRING)
		customAttributeDetail.SetMinLength(1)
		customAttributeDetail.SetMaxLength(20)
		customAttribute := make(map[string]UserSchemaAttribute)
		customAttribute[customAttributeName] = customAttributeDetail
		payload := UserSchemaPublic{Properties: &customAttribute}
		schema.Definitions.SetCustom(payload)
		req = req.UserSchema(*schema)
		updateSchema, _, err := req.Execute()
		require.NoError(t, err, "Could not update default user schema")
		assert.NotEmpty(t, updateSchema, "User schema is empty")
		updateAttribute := schema.Definitions.Custom.GetProperties()[customAttributeName]
		assert.Equal(t, customAttributeName, updateAttribute.GetTitle())
		assert.Equal(t, int32(1), updateAttribute.GetMinLength())
		assert.Equal(t, int32(20), updateAttribute.GetMaxLength())
	})
}

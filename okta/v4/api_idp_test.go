package okta

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupIdp(name string) (*IdentityProvider, *APIResponse, error) {
	req := apiClient.IdentityProviderAPI.CreateIdentityProvider(apiClient.cfg.Context)
	payload := testFactory.NewValidTestIdentityProvider()
	payload.SetName(name)
	req = req.IdentityProvider(*payload)
	return req.Execute()
}

func cleanUpIdp(idpId string) (err error) {
	_, err = apiClient.IdentityProviderAPI.DeleteIdentityProvider(apiClient.cfg.Context, idpId).Execute()
	return
}

func Test_Get_Identity_Provider(t *testing.T) {
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	t.Run("get idp by id", func(t *testing.T) {
		ridp, _, err := apiClient.IdentityProviderAPI.GetIdentityProvider(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not get idp by ID")
		assert.Equal(t, createdIdp.GetId(), ridp.GetId())
		assert.Equal(t, createdIdp.GetProtocol(), ridp.GetProtocol())
		assert.Equal(t, createdIdp.GetPolicy(), ridp.GetPolicy())
	})
	t.Run("get list idp", func(t *testing.T) {
		lidp, _, err := apiClient.IdentityProviderAPI.ListIdentityProviders(apiClient.cfg.Context).Execute()
		require.NoError(t, err, "Could not get list idp")
		var createdIdpInList bool
		for _, idp := range lidp {
			if idp.GetId() == createdIdp.GetId() {
				createdIdpInList = true
			}
		}
		assert.True(t, createdIdpInList)
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Activate_Identity_Provider(t *testing.T) {
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	assert.Equal(t, "ACTIVE", createdIdp.GetStatus())
	t.Run("deactivate idp", func(t *testing.T) {
		didp, _, err := apiClient.IdentityProviderAPI.DeactivateIdentityProvider(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not deactivate idp")
		assert.Equal(t, "INACTIVE", didp.GetStatus())
	})
	t.Run("activate idp", func(t *testing.T) {
		aidp, _, err := apiClient.IdentityProviderAPI.ActivateIdentityProvider(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not activate idp")
		assert.Equal(t, "ACTIVE", aidp.GetStatus())
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Update_Identity_Provider(t *testing.T) {
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	t.Run("update idp", func(t *testing.T) {
		req := apiClient.IdentityProviderAPI.ReplaceIdentityProvider(apiClient.cfg.Context, createdIdp.GetId())
		createdIdp.SetName(fmt.Sprintf("%v%v", testPrefix, "Update"))
		req = req.IdentityProvider(*createdIdp)
		uidp, _, err := req.Execute()
		require.NoError(t, err, "Could not update idp")
		assert.Equal(t, fmt.Sprintf("%v%v", testPrefix, "Update"), uidp.GetName())
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Get_Key(t *testing.T) {
	req := apiClient.IdentityProviderAPI.CreateIdentityProviderKey(apiClient.cfg.Context)
	payload := testFactory.NewValidTestJsonWebKey()
	req = req.JsonWebKey(*payload)
	createdKey, _, err := req.Execute()
	require.NoError(t, err, "Creating a new idp key should not error")
	t.Run("get idp key by id", func(t *testing.T) {
		ridpk, _, err := apiClient.IdentityProviderAPI.GetIdentityProviderKey(apiClient.cfg.Context, createdKey.GetKid()).Execute()
		require.NoError(t, err, "Could not get idp key by ID")
		assert.Equal(t, createdKey.GetKid(), ridpk.GetKid())
	})
	t.Run("get list idp keys", func(t *testing.T) {
		lidpk, _, err := apiClient.IdentityProviderAPI.ListIdentityProviderKeys(apiClient.cfg.Context).Execute()
		require.NoError(t, err, "Could not get list idp")
		assert.Equal(t, len(lidpk), 3)
	})
	_, err = apiClient.IdentityProviderAPI.DeleteIdentityProviderKey(apiClient.cfg.Context, createdKey.GetKid()).Execute()
	require.NoError(t, err, "Clean up idp key should not error")
}

func Test_List_Signing_Keys(t *testing.T) {
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	var generatedKey *JsonWebKey
	t.Run("generate signing key", func(t *testing.T) {
		req := apiClient.IdentityProviderAPI.GenerateIdentityProviderSigningKey(apiClient.cfg.Context, createdIdp.GetId())
		req = req.ValidityYears(int32(2))
		generatedKey, _, err = req.Execute()
		require.NoError(t, err, "Generating a new signing key should not error")
		assert.NotNil(t, generatedKey)
		assert.NotEmpty(t, generatedKey.X5c)
	})
	t.Run("list signing keys", func(t *testing.T) {
		retrievedKeys, _, err := apiClient.IdentityProviderAPI.ListIdentityProviderSigningKeys(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Retrieveing signing keys should not error")
		var result bool
		for _, v := range retrievedKeys {
			if contain(generatedKey.X5c[0], v.X5c) {
				result = true
			}
		}
		assert.True(t, result)
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Clone_Signing_Key(t *testing.T) {
	createdIdp1, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	createdIdp2, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	t.Run("clone signing key", func(t *testing.T) {
		greq := apiClient.IdentityProviderAPI.GenerateIdentityProviderSigningKey(apiClient.cfg.Context, createdIdp1.GetId())
		greq = greq.ValidityYears(int32(2))
		generatedKey, _, err := greq.Execute()
		require.NoError(t, err, "Generating a new signing key should not error")
		creq := apiClient.IdentityProviderAPI.CloneIdentityProviderKey(apiClient.cfg.Context, createdIdp1.GetId(), generatedKey.GetKid())
		creq = creq.TargetIdpId(createdIdp2.GetId())
		clonedKey, _, err := creq.Execute()
		require.NoError(t, err, "Could not clone signing key")
		assert.NotNil(t, clonedKey)
		assert.Equal(t, generatedKey.GetKid(), clonedKey.GetKid())
	})
	err = cleanUpIdp(createdIdp1.GetId())
	require.NoError(t, err, "Clean up idp should not error")
	err = cleanUpIdp(createdIdp2.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Get_CSR(t *testing.T) {
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	var generatedCsr *Csr
	t.Run("generate CSR", func(t *testing.T) {
		req := apiClient.IdentityProviderAPI.GenerateCsrForIdentityProvider(apiClient.cfg.Context, createdIdp.GetId())
		req = req.Metadata(*testFactory.NewValidTestCSRMetadata())
		generatedCsr, _, err = req.Execute()
		require.NoError(t, err, "Generating a new csr should not error")
		assert.NotNil(t, generatedCsr)
		assert.Equal(t, "RSA", generatedCsr.GetKty())
		assert.NotNil(t, generatedCsr.Csr)
	})
	t.Run("get CSR by ID", func(t *testing.T) {
		rcsr, _, err := apiClient.IdentityProviderAPI.GetCsrForIdentityProvider(apiClient.cfg.Context, createdIdp.GetId(), generatedCsr.GetId()).Execute()
		require.NoError(t, err, "Could not get csr by ID")
		assert.NotNil(t, rcsr)
		assert.Equal(t, generatedCsr.GetKty(), rcsr.GetKty())
		assert.NotNil(t, generatedCsr.GetCsr(), rcsr.GetCsr())
	})
	t.Run("list CSR", func(t *testing.T) {
		listCSRs, _, err := apiClient.IdentityProviderAPI.ListCsrsForIdentityProvider(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not list csr by idp ID")
		assert.NotEmpty(t, listCSRs)
		var result bool
		for _, csr := range listCSRs {
			if csr.GetId() == generatedCsr.GetId() {
				result = true
			}
		}
		assert.True(t, result)
	})
	t.Run("revoke CSR", func(t *testing.T) {
		_, err := apiClient.IdentityProviderAPI.RevokeCsrForIdentityProvider(apiClient.cfg.Context, createdIdp.GetId(), generatedCsr.GetId()).Execute()
		require.NoError(t, err, "Unable to revoke csr")
	})
	t.Run("list CSR", func(t *testing.T) {
		listCSRs, _, err := apiClient.IdentityProviderAPI.ListCsrsForIdentityProvider(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not list csr by idp ID")
		assert.Empty(t, listCSRs)
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
}

func Test_Get_Linked_User(t *testing.T) {
	createdUser, _, _, err := setupUser(false)
	require.NoError(t, err, "Creating a new user should not error")
	createdIdp, _, err := setupIdp(randomTestString())
	require.NoError(t, err, "Creating a new idp should not error")
	t.Run("link idp to user", func(t *testing.T) {
		req := apiClient.IdentityProviderAPI.LinkUserToIdentityProvider(apiClient.cfg.Context, createdIdp.GetId(), createdUser.GetId())
		externalId := "externalId"
		req = req.UserIdentityProviderLinkRequest(UserIdentityProviderLinkRequest{ExternalId: &externalId})
		_, _, err := req.Execute()
		require.NoError(t, err, "Could not link user and idp")
	})
	t.Run("get linked user for idps", func(t *testing.T) {
		linkUser, _, err := apiClient.IdentityProviderAPI.GetIdentityProviderApplicationUser(apiClient.cfg.Context, createdIdp.GetId(), createdUser.GetId()).Execute()
		require.NoError(t, err, "Could not get user's idp")
		assert.Equal(t, createdUser.GetId(), linkUser.GetId())
		var idpInLink bool
		if idp, ok := linkUser.Links.AdditionalProperties["idp"]; ok {
			if hrefMap, ok := idp.(map[string]interface{}); ok {
				if href, ok := hrefMap["href"]; ok {
					if strings.Contains(fmt.Sprintf("%v", href), createdIdp.GetId()) {
						idpInLink = true
					}
				}
			}
		}
		assert.True(t, idpInLink)
	})
	t.Run("list linked idp user", func(t *testing.T) {
		listIdp, _, err := apiClient.IdentityProviderAPI.ListIdentityProviderApplicationUsers(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not list idp's user")
		assert.Equal(t, 1, len(listIdp))
	})
	t.Run("unlink idp from user", func(t *testing.T) {
		_, err := apiClient.IdentityProviderAPI.UnlinkUserFromIdentityProvider(apiClient.cfg.Context, createdIdp.GetId(), createdUser.GetId()).Execute()
		require.NoError(t, err, "Could unlink idp and user")
	})
	t.Run("list linked idp user", func(t *testing.T) {
		listIdp, _, err := apiClient.IdentityProviderAPI.ListIdentityProviderApplicationUsers(apiClient.cfg.Context, createdIdp.GetId()).Execute()
		require.NoError(t, err, "Could not list idp's user")
		assert.Empty(t, listIdp)
	})
	err = cleanUpIdp(createdIdp.GetId())
	require.NoError(t, err, "Clean up idp should not error")
	err = cleanUpUser(createdUser.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

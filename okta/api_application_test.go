package okta

import (
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupBasicAuthApplication(label string) (*ListApplications200ResponseInner, *APIResponse, error) {
	req := apiClient.ApplicationApi.CreateApplication(apiClient.cfg.Context)
	req = req.Application(ListApplications200ResponseInner{BasicAuthApplication: testFactory.NewValidBasicAuthApplication(label)})
	return req.Execute()
}

func setupOrg2OrgApplication(label string) (*ListApplications200ResponseInner, *APIResponse, error) {
	req := apiClient.ApplicationApi.CreateApplication(apiClient.cfg.Context)
	req = req.Application(ListApplications200ResponseInner{SamlApplication: testFactory.NewValidOrg2OrgApplication(label)})
	return req.Execute()
}

func setupBookmarkApplication(label string) (*ListApplications200ResponseInner, *APIResponse, error) {
	req := apiClient.ApplicationApi.CreateApplication(apiClient.cfg.Context)
	req = req.Application(ListApplications200ResponseInner{BookmarkApplication: testFactory.NewValidBookmarkApplication(label)})
	return req.Execute()
}

func setupOIDCApplication(label string) (*ListApplications200ResponseInner, *APIResponse, error) {
	req := apiClient.ApplicationApi.CreateApplication(apiClient.cfg.Context)
	req = req.Application(ListApplications200ResponseInner{OpenIdConnectApplication: testFactory.NewValidOIDCApplication(label)})
	return req.Execute()
}

func cleanUpApplication(appId string) error {
	_, err := apiClient.ApplicationApi.DeactivateApplication(apiClient.cfg.Context, appId).Execute()
	if err != nil {
		return err
	}
	_, err = apiClient.ApplicationApi.DeleteApplication(apiClient.cfg.Context, appId).Execute()
	if err != nil {
		return err
	}
	return err
}

func Test_Get_Applications(t *testing.T) {
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("get applications by id", func(t *testing.T) {
		app, _, err := apiClient.ApplicationApi.GetApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not get app by ID")
		assert.Equal(t, createdApp.BasicAuthApplication.GetId(), app.BasicAuthApplication.GetId())
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Get_List_Applications(t *testing.T) {
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("get all applications", func(t *testing.T) {
		apps, _, err := apiClient.ApplicationApi.ListApplications(apiClient.cfg.Context).Limit(100).Execute()
		require.NoError(t, err, "Could not get list apps")
		var createAppInList bool
		for _, a := range apps {
			if (a.BasicAuthApplication != nil) && (a.BasicAuthApplication.GetId() == createdApp.BasicAuthApplication.GetId()) {
				createAppInList = true
				break
			}
		}
		assert.True(t, createAppInList, "Could not find app from list")
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Update_App(t *testing.T) {
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("update applications", func(t *testing.T) {
		newName := randomTestString()
		payload := testFactory.NewValidBasicAuthApplication(newName)
		payload.Settings.App.SetAuthURL("https://example.org/auth.html")
		payload.Settings.App.SetUrl("https://example.org/auth.html")
		app, _, err := apiClient.ApplicationApi.ReplaceApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Application(ListApplications200ResponseInner{BasicAuthApplication: payload}).Execute()
		require.NoError(t, err, "Could not update apps")
		require.NotNil(t, app.BasicAuthApplication)
		assert.Equal(t, newName, app.BasicAuthApplication.GetLabel())
		assert.Equal(t, "https://example.org/auth.html", app.BasicAuthApplication.Settings.App.GetAuthURL())
		assert.Equal(t, "https://example.org/auth.html", app.BasicAuthApplication.Settings.App.GetUrl())
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Activate_Application(t *testing.T) {
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("deactivate applications", func(t *testing.T) {
		_, err = apiClient.ApplicationApi.DeactivateApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not deactivate the app")
		app, _, err := apiClient.ApplicationApi.GetApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not get app by ID")
		assert.Equal(t, createdApp.BasicAuthApplication.GetId(), app.BasicAuthApplication.GetId())
		assert.Equal(t, APPLICATIONLIFECYCLESTATUS_INACTIVE, app.BasicAuthApplication.GetStatus())
	})
	t.Run("activate applications", func(t *testing.T) {
		_, err = apiClient.ApplicationApi.ActivateApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not activate the app")
		newapp, _, err := apiClient.ApplicationApi.GetApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not get app by ID")
		assert.Equal(t, createdApp.BasicAuthApplication.GetId(), newapp.BasicAuthApplication.GetId())
		assert.Equal(t, APPLICATIONLIFECYCLESTATUS_ACTIVE, newapp.BasicAuthApplication.GetStatus())
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Application_Users_Operations(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(90))
	}
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	appUserList, _, err := apiClient.ApplicationApi.ListApplicationUsers(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
	require.NoError(t, err, "Could not get list app users")
	require.Empty(t, appUserList, "App User List should be empty")
	user, _, _, err := setupUser(false)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("assign user to application", func(t *testing.T) {
		pwcredentials := AppUserPasswordCredential{}
		pwcredentials.SetValue(randomTestString())
		credentials := AppUserCredentials{}
		credentials.SetPassword(pwcredentials)
		credentials.SetUserName(randomTestString())
		payload := AppUser{}
		payload.SetId(user.GetId())
		payload.SetCredentials(credentials)
		appUser, _, err := apiClient.ApplicationApi.AssignUserToApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).AppUser(payload).Execute()
		require.NoError(t, err, "Assigning user to application should not error")
		appUserList, _, err := apiClient.ApplicationApi.ListApplicationUsers(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not get list app users")
		require.NotEmpty(t, appUserList, "App User List should not be empty")
		var found bool
		for _, ae := range appUserList {
			if ae.GetId() == appUser.GetId() {
				found = true
				break
			}
		}
		assert.True(t, found)
	})
	t.Run("get application user", func(t *testing.T) {
		appUser, _, err := apiClient.ApplicationApi.GetApplicationUser(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), user.GetId()).Execute()
		require.NoError(t, err, "Could not get app user")
		assert.Equal(t, user.GetId(), appUser.GetId())
	})
	t.Run("update application user", func(t *testing.T) {
		appUser, _, err := apiClient.ApplicationApi.GetApplicationUser(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), user.GetId()).Execute()
		require.NoError(t, err, "Could not get app user")
		oldUserName := appUser.Credentials.GetUserName()
		newUserName := randomTestString()
		pwcredentials := AppUserPasswordCredential{}
		pwcredentials.SetValue(randomTestString())
		credentials := AppUserCredentials{}
		credentials.SetPassword(pwcredentials)
		credentials.SetUserName(newUserName)
		appUser.SetCredentials(credentials)
		updatedAppUser, _, err := apiClient.ApplicationApi.UpdateApplicationUser(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), user.GetId()).AppUser(*appUser).Execute()
		require.NoError(t, err, "Could not update app user")
		assert.NotEqual(t, oldUserName, updatedAppUser.Credentials.GetUserName())
		assert.Equal(t, newUserName, updatedAppUser.Credentials.GetUserName())
	})
	t.Run("remove application from user", func(t *testing.T) {
		_, err = apiClient.ApplicationApi.UnassignUserFromApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), user.GetId()).Execute()
		require.NoError(t, err, "Delete user to application should not error")
		appUserList, _, err := apiClient.ApplicationApi.ListApplicationUsers(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not get list apps")
		require.Empty(t, appUserList, "App User List should be empty")
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
}

func Test_Application_Groups_Operations(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(90))
	}
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	t.Run("assign group to application", func(t *testing.T) {
		payload := ApplicationGroupAssignment{}
		payload.SetPriority(5)
		appGroup, _, err := apiClient.ApplicationApi.AssignGroupToApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), group.GetId()).ApplicationGroupAssignment(payload).Execute()
		require.NoError(t, err, "Create app group assignment should not error")
		assert.NotNil(t, appGroup)
	})
	t.Run("get application group", func(t *testing.T) {
		appGroup, _, err := apiClient.ApplicationApi.GetApplicationGroupAssignment(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), group.GetId()).Execute()
		require.NoError(t, err, "Get app group assignment should not error")
		assert.Equal(t, int32(5), appGroup.GetPriority())
	})
	t.Run("list application group", func(t *testing.T) {
		appGroupList, _, err := apiClient.ApplicationApi.ListApplicationGroupAssignments(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Get list app group assignment should not error")
		assert.NotEmpty(t, appGroupList)
	})
	t.Run("remove application from group", func(t *testing.T) {
		_, err = apiClient.ApplicationApi.UnassignApplicationFromGroup(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), group.GetId()).Execute()
		require.NoError(t, err, "Delete app group assignment should not error")
		_, _, err := apiClient.ApplicationApi.GetApplicationGroupAssignment(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), group.GetId()).Execute()
		assert.Equal(t, "404 Not Found", err.Error())
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_CSR_For_Application(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(90))
	}
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	var generatedCsr *Csr
	t.Run("generate csr", func(t *testing.T) {
		generatedCsr, _, err = apiClient.ApplicationApi.GenerateCsrForApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Metadata(*testFactory.NewValidTestCSRMetadata()).Execute()
		require.NoError(t, err, "Generating a new csr should not error")
		assert.NotNil(t, generatedCsr)
		assert.Equal(t, "RSA", generatedCsr.GetKty())
		assert.NotNil(t, generatedCsr.Csr)
	})
	t.Run("get CSR by ID", func(t *testing.T) {
		rcsr, _, err := apiClient.ApplicationApi.GetCsrForApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), generatedCsr.GetId()).Execute()
		require.NoError(t, err, "Could not get csr by ID")
		assert.NotNil(t, rcsr)
		assert.Equal(t, generatedCsr.GetKty(), rcsr.GetKty())
		assert.NotNil(t, generatedCsr.GetCsr(), rcsr.GetCsr())
	})
	t.Run("list CSR", func(t *testing.T) {
		listCSRs, _, err := apiClient.ApplicationApi.ListCsrsForApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not list csr by app ID")
		assert.NotEmpty(t, listCSRs)
		var result bool
		for _, csr := range listCSRs {
			if csr.GetId() == generatedCsr.GetId() {
				result = true
				break
			}
		}
		assert.True(t, result)
	})
	t.Run("revoke csr", func(t *testing.T) {
		_, err := apiClient.ApplicationApi.RevokeCsrFromApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId(), generatedCsr.GetId()).Execute()
		require.NoError(t, err, "Unable to revoke csr")
		listCSRs, _, err := apiClient.ApplicationApi.ListCsrsForApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "Could not list csr by app ID")
		assert.Empty(t, listCSRs)
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func TestGetDefaultProvisioningConnectionForApplication(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(90))
	}
	createdApp, _, err := setupOrg2OrgApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("get provisioning", func(t *testing.T) {
		conn, _, err := apiClient.ApplicationApi.GetDefaultProvisioningConnectionForApplication(apiClient.cfg.Context, createdApp.SamlApplication.GetId()).Execute()
		require.NoError(t, err, "getting default provisioning connection for application should not error.")
		assert.NotEmpty(t, conn.GetAuthScheme())
		assert.NotEmpty(t, conn.GetStatus())
	})
	t.Run("set provisioning", func(t *testing.T) {
		profile := &ProvisioningConnectionProfile{}
		profile.SetAuthScheme("TOKEN")
		profile.SetToken("TEST")
		payload := ProvisioningConnectionRequest{Profile: profile}
		conn, _, err := apiClient.ApplicationApi.UpdateDefaultProvisioningConnectionForApplication(apiClient.cfg.Context, createdApp.SamlApplication.GetId()).ProvisioningConnectionRequest(payload).Activate(false).Execute()
		require.NoError(t, err, "setting default provisioning connection for application should not error.")
		assert.Equal(t, PROVISIONINGCONNECTIONAUTHSCHEME_TOKEN, conn.GetAuthScheme())
	})
	err = cleanUpApplication(createdApp.SamlApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

// support org2org apps only
// func Test_List_Features_For_Application(t *testing.T) {
// 	createdApp, _, err := setupOrg2OrgApplication(randomTestString())
// 	require.NoError(t, err, "Creating a new application should not error")
// 	t.Run("list feature for application", func(t *testing.T) {
// 		profile := &ProvisioningConnectionProfile{}
// 		profile.SetAuthScheme("TOKEN")
// 		profile.SetToken("FIXME_WITH_REAL_ORG_TOKEN")
// 		payload := ProvisioningConnectionRequest{Profile: profile}
// 		_, _, err := apiClient.ApplicationApi.SetDefaultProvisioningConnectionForApplication(apiClient.cfg.Context, createdApp.SamlApplication.GetId()).ProvisioningConnectionRequest(payload).Activate(true).Execute()
// 		require.NoError(t, err, "setting default provisioning connection for application should not error.")
// 		features, _, err := apiClient.ApplicationApi.ListFeaturesForApplication(apiClient.cfg.Context, createdApp.SamlApplication.GetId()).Execute()
// 		require.NoError(t, err, "listing features for application should not error.")
// 		var found bool
// 		for _, feature := range features {
// 			if feature.GetName() == "USER_PROVISIONING" {
// 				found = true
// 				break
// 			}
// 		}
// 		assert.True(t, found)
// 	})
// 	err = cleanUpApplication(createdApp.SamlApplication.GetId())
// 	require.NoError(t, err, "Clean up app should not error")
// }

func Test_Upload_Application_Logo(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(180))
	}
	createdApp, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	t.Run("upload application logo", func(t *testing.T) {
		fileDir, _ := os.Getwd()
		fileName := "asset/logo.png"
		filePath := path.Join(fileDir, fileName)
		file, err := os.Open(filePath)
		require.NoError(t, err, "opening application logo should not error.")
		_, err = apiClient.ApplicationApi.UploadApplicationLogo(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).File(file).Execute()
		require.NoError(t, err, "uploading application logo should not error.")
	})
	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Application_Key_Operation(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(180))
	}
	createdApp1, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	createdApp2, _, err := setupBasicAuthApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	var createdKey *JsonWebKey
	t.Run("generate application key", func(t *testing.T) {
		createdKey, _, err = apiClient.ApplicationApi.GenerateApplicationKey(apiClient.cfg.Context, createdApp1.BasicAuthApplication.GetId()).ValidityYears(2).Execute()
		assert.Nil(t, err, "generate new application key should not error")
	})
	t.Run("clone application key", func(t *testing.T) {
		ckey, _, err := apiClient.ApplicationApi.CloneApplicationKey(apiClient.cfg.Context, createdApp1.BasicAuthApplication.GetId(), createdKey.GetKid()).TargetAid(createdApp2.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "clone application key should not error")
		assert.Equal(t, createdKey.GetKid(), ckey.GetKid())
		assert.Equal(t, createdKey.GetExpiresAt(), ckey.GetExpiresAt())
		assert.Equal(t, createdKey.GetX5c(), ckey.GetX5c())
	})
	t.Run("get application key credentials", func(t *testing.T) {
		rkey, _, err := apiClient.ApplicationApi.GetApplicationKey(apiClient.cfg.Context, createdApp1.BasicAuthApplication.GetId(), createdKey.GetKid()).Execute()
		require.NoError(t, err, "get application key should not error")
		assert.Equal(t, createdKey.GetKid(), rkey.GetKid())
		assert.Equal(t, createdKey.GetCreated(), rkey.GetCreated())
		assert.Equal(t, createdKey.GetExpiresAt(), rkey.GetExpiresAt())
		assert.Equal(t, createdKey.GetX5c(), rkey.GetX5c())
	})
	t.Run("list application key credentials", func(t *testing.T) {
		rkeys, _, err := apiClient.ApplicationApi.ListApplicationKeys(apiClient.cfg.Context, createdApp1.BasicAuthApplication.GetId()).Execute()
		require.NoError(t, err, "get application key should not error")
		assert.Equal(t, 2, len(rkeys))
	})
	err = cleanUpApplication(createdApp1.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
	err = cleanUpApplication(createdApp2.BasicAuthApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

func Test_Scope_Consent_Grant_Operation_For_Application(t *testing.T) {
	if os.Getenv("OKTA_TRAVIS_CI") == "yes" {
		time.Sleep(time.Duration(180))
	}
	createdApp, _, err := setupOIDCApplication(randomTestString())
	require.NoError(t, err, "Creating a new application should not error")
	var grant *OAuth2ScopeConsentGrant
	t.Run("grant consent to scope", func(t *testing.T) {
		payload := OAuth2ScopeConsentGrant{}
		payload.SetIssuer(apiClient.cfg.Okta.Client.OrgUrl)
		payload.SetScopeId("okta.users.read")
		grant, _, err = apiClient.ApplicationApi.GrantConsentToScope(apiClient.cfg.Context, createdApp.OpenIdConnectApplication.GetId()).OAuth2ScopeConsentGrant(payload).Execute()
		assert.Nil(t, err, "grant consent to scope should not error")
	})
	t.Run("get scope consent grant", func(t *testing.T) {
		rgrant, _, err := apiClient.ApplicationApi.GetScopeConsentGrant(apiClient.cfg.Context, createdApp.OpenIdConnectApplication.GetId(), grant.GetId()).Execute()
		require.NoError(t, err, "get scope consent grant should not error")
		assert.Equal(t, grant.GetId(), rgrant.GetId())
		assert.Equal(t, grant.GetClientId(), rgrant.GetClientId())
	})
	t.Run("list scope consent grant", func(t *testing.T) {
		rgrants, _, err := apiClient.ApplicationApi.ListScopeConsentGrants(apiClient.cfg.Context, createdApp.OpenIdConnectApplication.GetId()).Execute()
		require.NoError(t, err, "list scope consent grant should not error")
		assert.NotEmpty(t, rgrants)
	})
	t.Run("revoke consent grant", func(t *testing.T) {
		_, err = apiClient.ApplicationApi.RevokeScopeConsentGrant(apiClient.cfg.Context, createdApp.OpenIdConnectApplication.GetId(), grant.GetId()).Execute()
		assert.Nil(t, err, "revoke consent to scope should not error")
	})
	err = cleanUpApplication(createdApp.OpenIdConnectApplication.GetId())
	require.NoError(t, err, "Clean up app should not error")
}

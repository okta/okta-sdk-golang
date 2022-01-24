/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta/query"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"

	"github.com/okta/okta-sdk-golang/v2/tests"
)

func TestCanGetApplicationByID(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	foundApplication, _, err := client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
	require.NoError(t, err, "Should not error when getting an applicaiton by id")

	assert.Equal(t, appId, foundApplication.(*okta.BasicAuthApplication).Id, "Application found was not correct")

	client.Application.DeactivateApplication(ctx, appId)
	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanUpdateApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	newBasicApplication := okta.NewBasicAuthApplication()
	newBasicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.org/auth.html",
			Url:     "https://example.org/auth.html",
		},
	}

	newApp, _, err := client.Application.UpdateApplication(ctx, appId, newBasicApplication)
	require.NoError(t, err, "Updating an application caused an error")

	assert.Equal(t, "https://example.org/auth.html", newApp.(*okta.BasicAuthApplication).Settings.App.Url, "The application did not update")

	_, err = client.Application.DeactivateApplication(ctx, appId)
	require.NoError(t, err, "Deactivating an application should not error")

	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanCreateABookmarkApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = &okta.BookmarkApplicationSettings{
		App: &okta.BookmarkApplicationSettingsApplication{
			RequestIntegration: new(bool),
			Url:                "https://example.com/bookmark.htm",
		},
	}

	assert.Empty(t, bookmarkApplication.Id)
	application, _, err := client.Application.CreateApplication(ctx, bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BookmarkApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BookmarkApplication).Id)

	client.Application.DeactivateApplication(ctx, application.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(ctx, application.(*okta.BookmarkApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanCreateABasicAuthenticationApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BasicAuthApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BasicAuthApplication).Id)
	assert.NotEmpty(t, basicApplication.Id)

	client.Application.DeactivateApplication(ctx, application.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(ctx, application.(*okta.BasicAuthApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestListApplicationAllowsCastingToCorrectType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	app1, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = &okta.BookmarkApplicationSettings{
		App: &okta.BookmarkApplicationSettingsApplication{
			RequestIntegration: new(bool),
			Url:                "https://example.com/bookmark.htm",
		},
	}

	app2, _, err := client.Application.CreateApplication(ctx, bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	applist, _, err := client.Application.ListApplications(ctx, nil)
	require.NoError(t, err, "List applications should not error")

	for _, a := range applist {
		if a.(*okta.Application).Name == "bookmark" {
			if a.(*okta.Application).Id == app2.(*okta.BookmarkApplication).Id {
				ba, _, _ := client.Application.GetApplication(ctx, a.(*okta.Application).Id, okta.NewBookmarkApplication(), nil)
				requestIntegration := ba.(*okta.BookmarkApplication).Settings.App.RequestIntegration
				assert.False(t, *requestIntegration)
			}
		}
	}

	client.Application.DeactivateApplication(ctx, app1.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(ctx, app1.(*okta.BasicAuthApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")

	client.Application.DeactivateApplication(ctx, app2.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(ctx, app2.(*okta.BookmarkApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanActivateApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, query.NewQueryParams(query.WithActivate(false)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.ActivateApplication(ctx, appId)
	require.NoError(t, err, "Activationg an application should not error")
	application, _, _ = client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	client.Application.DeactivateApplication(ctx, appId)
	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanDeactivateApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, query.NewQueryParams(query.WithActivate(true)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeactivateApplication(ctx, appId)
	require.NoError(t, err, "Deactivating an application should not error")
	application, _, err = client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanAddAndRemoveApplicationUsers(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	appUserList, _, _ := client.Application.ListApplicationUsers(ctx, appId, nil)
	assert.Empty(t, appUserList, "App Users should be empty")

	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Get-User"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	req := okta.AppUser{
		Credentials: &okta.AppUserCredentials{
			Password: &okta.AppUserPasswordCredential{
				Value: "abcd1234",
			},
			UserName: "appUser",
		},
		Id: user.Id,
	}

	appUser, _, err := client.Application.AssignUserToApplication(ctx, appId, req)
	require.NoError(t, err, "Assigning user to application show not error")

	assert.IsType(t, okta.AppUser{}, *appUser, "Type returned from assigning user to application was incorrect")

	appUserList, _, _ = client.Application.ListApplicationUsers(ctx, appId, nil)
	assert.NotEmpty(t, appUserList, "App Users should not be empty")

	client.Application.DeleteApplicationUser(ctx, appId, appUser.Id, nil)
	client.GetRequestExecutor().RefreshNext()
	appUserList, _, _ = client.Application.ListApplicationUsers(ctx, appId, nil)
	assert.Empty(t, appUserList, "App Users should be empty after deleting")

	client.Application.DeactivateApplication(ctx, appId)
	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")

	client.User.DeactivateUser(ctx, user.Id, nil)
	client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
}

func TestCanSetApplicationSettingsDuringCreation(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, okta.BasicApplicationSettingsApplication{}, *application.(*okta.BasicAuthApplication).Settings.App, "The returned type of application settings application was not correct type")
	assert.Equal(t, "https://example.com/auth.html", application.(*okta.BasicAuthApplication).Settings.App.Url)

	appId := application.(*okta.BasicAuthApplication).Id
	client.Application.DeactivateApplication(ctx, appId)
	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanSetApplicationSettingsDuringUpdate(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	myApp, _, err := client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
	app := myApp.(*okta.BasicAuthApplication)
	myAppId := app.Id
	myAppSettingsApp := app.Settings.App

	myAppSettingsApp.Url = "https://okta.com/auth"

	_, _, _ = client.Application.UpdateApplication(ctx, myAppId, app)

	updatedApp, _, err := client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "https://okta.com/auth", updatedApp.(*okta.BasicAuthApplication).Settings.App.Url, "The URL was not updated'")

	client.Application.DeactivateApplication(ctx, appId)
	_, err = client.Application.DeleteApplication(ctx, appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestCanCreateCSRForApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := createBasicAuthApplication(t)

	subject := okta.CsrMetadataSubject{
		CountryName:            "US",
		StateOrProvinceName:    "California",
		LocalityName:           "San Francisco",
		OrganizationName:       "Okta, Inc",
		OrganizationalUnitName: "Dev",
		CommonName:             "SP Issuer",
	}

	subjectAltNames := okta.CsrMetadataSubjectAltNames{
		DnsNames: []string{"dev.okta.com"},
	}

	csr := okta.CsrMetadata{
		Subject:         &subject,
		SubjectAltNames: &subjectAltNames,
	}

	csrs, _, err := client.Application.GenerateCsrForApplication(ctx, application.Id, csr)
	require.NoError(t, err, "Creating an application Csr should not error")

	assert.IsType(t, &okta.Csr{}, csrs, "did not return a `okta.Csr` object")

	client.Application.DeactivateApplication(ctx, application.Id)
	_, err = client.Application.DeleteApplication(ctx, application.Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func TestGetDefaultProvisioningConnectionForApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := createOrg2OrgApplication(t)

	conn, _, err := client.Application.GetDefaultProvisioningConnectionForApplication(ctx, application.Id)
	require.NoError(t, err, "getting default provisioning connection for application should not error.")
	assert.NotEmpty(t, conn.AuthScheme, "connection auth scheme shouldn't be empty")
	assert.NotEmpty(t, conn.Status, "connection status shouldn't be empty")

	client.Application.DeactivateApplication(ctx, application.Id)
	_, err = client.Application.DeleteApplication(ctx, application.Id)
	require.NoError(t, err, "Deleting an application should not error")
}

func TestSetDefaultProvisioningConnectionForApplication(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := createOrg2OrgApplication(t)

	provisionConnectionRequest := okta.ProvisioningConnectionRequest{
		Profile: &okta.ProvisioningConnectionProfile{
			AuthScheme: "TOKEN",
			Token:      "TEST",
		},
	}

	conn, _, err := client.Application.SetDefaultProvisioningConnectionForApplication(ctx, application.Id, provisionConnectionRequest, query.NewQueryParams(query.WithActivate(false)))
	require.NoError(t, err, "setting default provisioning connection for application should not error.")
	assert.Equal(t, "TOKEN", conn.AuthScheme, "expected auth scheme %q, go %q", "TOKEN", conn.AuthScheme)

	client.Application.DeactivateApplication(ctx, application.Id)
	_, err = client.Application.DeleteApplication(ctx, application.Id)
	require.NoError(t, err, "Deleting an application should not error")
}

func TestListFeaturesForApplication(t *testing.T) {
	t.Skip("listing application features is specific to an org2org")

	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := createOrg2OrgApplication(t)

	// FIXME this needs a second org to run against for this IT to work
	// need to activate a provisioning connection between orgs
	provisionConnectionRequest := okta.ProvisioningConnectionRequest{
		Profile: &okta.ProvisioningConnectionProfile{
			AuthScheme: "TOKEN",
			Token:      "FIXME_WITH_REAL_ORG_TOKEN",
		},
	}
	_, _, err = client.Application.SetDefaultProvisioningConnectionForApplication(ctx, application.Id, provisionConnectionRequest, query.NewQueryParams(query.WithActivate(false)))
	require.NoError(t, err, "setting default provisioning connection for application should not error.")

	features, _, err := client.Application.ListFeaturesForApplication(ctx, application.Id)
	require.NoError(t, err, "listing features for application should not error.")

	foundUserProvisiontingFeature := false
	for _, feature := range features {
		// NOTE: Provisioning must be enabled for the application. To activate
		// provisioning, see Provisioning Connections. The only application
		// Feature currently supported is USER_PROVISIONING.
		// https://developer.okta.com/docs/reference/api/apps/#list-features-for-application
		if feature.Name == "USER_PROVISIONING" {
			foundUserProvisiontingFeature = true
			break
		}
	}
	if !foundUserProvisiontingFeature {
		assert.FailNow(t, "the org2org application should at least have USER_PROVISIONING feature")
	}

	client.Application.DeactivateApplication(ctx, application.Id)
	_, err = client.Application.DeleteApplication(ctx, application.Id)
	require.NoError(t, err, "Deleting an application should not error")
}

func TestUploadApplicationLogo(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := createBasicAuthApplication(t)

	fileDir, _ := os.Getwd()
	fileName := "../fixtures/logo.png"
	filePath := path.Join(fileDir, fileName)
	_, err = client.Application.UploadApplicationLogo(ctx, application.Id, filePath)
	require.NoError(t, err, "uploading application logo should not error.")

	client.Application.DeactivateApplication(ctx, application.Id)
	_, err = client.Application.DeleteApplication(ctx, application.Id)
	require.NoError(t, err, "Deleting an application should not error")
}

func createBasicAuthApplication(t *testing.T) *okta.BasicAuthApplication {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = &okta.BasicApplicationSettings{
		App: &okta.BasicApplicationSettingsApplication{
			AuthURL: "https://example.com/auth.html",
			Url:     "https://example.com/auth.html",
		},
	}

	application, _, err := client.Application.CreateApplication(ctx, basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	return application.(*okta.BasicAuthApplication)
}

func createOrg2OrgApplication(t *testing.T) *okta.Org2OrgApplication {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	application := okta.NewOrg2OrgApplication()
	application.Label = "Sample Okta Org2Org App"
	application.Name = "okta_org2org"
	application.Settings = &okta.Org2OrgApplicationSettings{
		App: &okta.Org2OrgApplicationSettingsApp{
			AcsUrl:         "https://example.okta.com/sso/saml2/exampleid",
			AudRestriction: "https://www.okta.com/saml2/service-provider/exampleid",
			BaseUrl:        "https://example.okta.com",
		},
	}

	app, _, err := client.Application.CreateApplication(ctx, application, query.NewQueryParams(query.WithActivate(true)))
	require.NoError(t, err, "Creating an application should not error")

	return app.(*okta.Org2OrgApplication)
}

func deleteAllApps(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	applicationList, _, err := client.Application.ListApplications(ctx, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	for _, a := range applicationList {
		if strings.HasPrefix(a.(*okta.Application).Label, "Template Basic Auth App") ||
			strings.HasPrefix(a.(*okta.Application).Label, "Bookmark App") {
			client.Application.DeactivateApplication(ctx, a.(*okta.Application).Id)
			client.Application.DeleteApplication(ctx, a.(*okta.Application).Id)
		}
	}
}

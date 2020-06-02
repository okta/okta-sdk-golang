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
	"strings"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta/query"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"

	"github.com/okta/okta-sdk-golang/v2/tests"
)

func Test_can_get_applicaiton_by_id(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	foundApplication, _, err := client.Application.GetApplication(context.TODO(), appId, okta.NewBasicAuthApplication(), nil)
	require.NoError(t, err, "Should not error when getting an applicaiton by id")

	assert.Equal(t, appId, foundApplication.(*okta.BasicAuthApplication).Id, "Application found was not correct")

	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_update_application(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	newBasicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	newBasicApplicationSettingsApplication.AuthURL = "https://example.org/auth.html"
	newBasicApplicationSettingsApplication.Url = "https://example.org/auth.html"

	newBasicApplicationSettings := okta.NewBasicApplicationSettings()
	newBasicApplicationSettings.App = newBasicApplicationSettingsApplication

	newBasicApplication := okta.NewBasicAuthApplication()
	newBasicApplication.Settings = newBasicApplicationSettings

	newApp, _, err := client.Application.UpdateApplication(context.TODO(), appId, newBasicApplication)
	require.NoError(t, err, "Updating an application caused an error")

	assert.Equal(t, "https://example.org/auth.html", newApp.(*okta.BasicAuthApplication).Settings.App.Url, "The application did not update")

	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_create_a_bookmark_application(t *testing.T) {
	client, _ := tests.NewClient()

	bookmarkApplicationSettingsApplication := okta.NewBookmarkApplicationSettingsApplication()
	bookmarkApplicationSettingsApplication.RequestIntegration = new(bool)
	bookmarkApplicationSettingsApplication.Url = "https://example.com/bookmark.htm"

	bookmarkApplicationSettings := okta.NewBookmarkApplicationSettings()
	bookmarkApplicationSettings.App = bookmarkApplicationSettingsApplication

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = bookmarkApplicationSettings
	assert.Empty(t, bookmarkApplication.Id)
	application, _, err := client.Application.CreateApplication(context.TODO(), bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BookmarkApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BookmarkApplication).Id)

	client.Application.DeactivateApplication(context.TODO(), application.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(context.TODO(), application.(*okta.BookmarkApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_create_a_basic_authentication_application(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BasicAuthApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BasicAuthApplication).Id)
	assert.NotEmpty(t, basicApplication.Id)

	client.Application.DeactivateApplication(context.TODO(), application.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(context.TODO(), application.(*okta.BasicAuthApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_list_application_allows_casting_to_correct_type(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	app1, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	bookmarkApplicationSettingsApplication := okta.NewBookmarkApplicationSettingsApplication()
	bookmarkApplicationSettingsApplication.RequestIntegration = new(bool)
	bookmarkApplicationSettingsApplication.Url = "https://example.com/bookmark.htm"

	bookmarkApplicationSettings := okta.NewBookmarkApplicationSettings()
	bookmarkApplicationSettings.App = bookmarkApplicationSettingsApplication

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = bookmarkApplicationSettings

	app2, _, err := client.Application.CreateApplication(context.TODO(), bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	applist, _, err := client.Application.ListApplications(context.TODO(), nil)
	require.NoError(t, err, "List applicaitons should not error")

	for _, a := range applist {

		if a.(*okta.Application).Name == "bookmark" {
			if a.(*okta.Application).Id == app2.(*okta.BookmarkApplication).Id {
				ba, _, _ := client.Application.GetApplication(context.TODO(), a.(*okta.Application).Id, okta.NewBookmarkApplication(), nil)
				requestIntegration := ba.(*okta.BookmarkApplication).Settings.App.RequestIntegration
				assert.False(t, *requestIntegration)
			}
		}
	}

	client.Application.DeactivateApplication(context.TODO(), app1.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(context.TODO(), app1.(*okta.BasicAuthApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")

	client.Application.DeactivateApplication(context.TODO(), app2.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(context.TODO(), app2.(*okta.BookmarkApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_activate_application(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, query.NewQueryParams(query.WithActivate(false)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.ActivateApplication(context.TODO(), appId)
	require.NoError(t, err, "Activationg an application should not error")
	application, _, _ = client.Application.GetApplication(context.TODO(), appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_deactivate_application(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, query.NewQueryParams(query.WithActivate(true)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeactivateApplication(context.TODO(), appId)
	require.NoError(t, err, "Deactivating an application should not error")
	application, _, err = client.Application.GetApplication(context.TODO(), appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_add_and_remove_application_users(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	appUserList, _, _ := client.Application.ListApplicationUsers(context.TODO(), appId, nil)
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
	profile["email"] = "john-get-user@example.com"
	profile["login"] = "john-get-user@example.com"
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(context.TODO(), *u, qp)
	require.NoError(t, err, "Creating an user should not error")

	appUserPasswordCredentials := okta.NewAppUserPasswordCredential()
	appUserPasswordCredentials.Value = "abcd1234"

	appUserCredentials := okta.NewAppUserCredentials()
	appUserCredentials.UserName = "appUser"
	appUserCredentials.Password = appUserPasswordCredentials

	appUser := okta.NewAppUser()
	appUser.Credentials = appUserCredentials
	appUser.Id = user.Id

	appUser, _, err = client.Application.AssignUserToApplication(context.TODO(), appId, *appUser)
	require.NoError(t, err, "Assigning user to application show not error")

	assert.IsType(t, okta.AppUser{}, *appUser, "Type returned from assigning user to application was incorrect")

	appUserList, _, _ = client.Application.ListApplicationUsers(context.TODO(), appId, nil)
	assert.NotEmpty(t, appUserList, "App Users should not be empty")

	client.Application.DeleteApplicationUser(context.TODO(), appId, appUser.Id, nil)
	appUserList, _, _ = client.Application.ListApplicationUsers(context.TODO(), appId, nil)
	assert.Empty(t, appUserList, "App Users should be empty after deleting")

	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")

	client.User.DeactivateUser(context.TODO(), user.Id, nil)
	client.User.DeactivateOrDeleteUser(context.TODO(), user.Id, nil)
}

func Test_can_set_application_settings_during_creation(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, okta.BasicApplicationSettingsApplication{}, *application.(*okta.BasicAuthApplication).Settings.App, "The returned type of application settings application was not correct type")
	assert.Equal(t, "https://example.com/auth.html", application.(*okta.BasicAuthApplication).Settings.App.Url)

	appId := application.(*okta.BasicAuthApplication).Id
	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_set_application_settings_during_update(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	assert.Empty(t, basicApplication.Id)
	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	myApp, _, err := client.Application.GetApplication(context.TODO(), appId, okta.NewBasicAuthApplication(), nil)
	app := myApp.(*okta.BasicAuthApplication)
	myAppId := app.Id
	myAppSettingsApp := app.Settings.App

	myAppSettingsApp.Url = "https://okta.com/auth"

	_, _, _ = client.Application.UpdateApplication(context.TODO(), myAppId, app)

	updatedApp, _, err := client.Application.GetApplication(context.TODO(), appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "https://okta.com/auth", updatedApp.(*okta.BasicAuthApplication).Settings.App.Url, "The URL was not updated'")

	client.Application.DeactivateApplication(context.TODO(), appId)
	_, err = client.Application.DeleteApplication(context.TODO(), appId)

	require.NoError(t, err, "Deleting an application should not error")
}

func Test_can_create_csr_for_application(t *testing.T) {
	client, _ := tests.NewClient()

	application := create_application(t)

	subject := okta.CSRMetadataSubject{
		CountryName:            "US",
		StateOrProvinceName:    "California",
		LocalityName:           "San Francisco",
		OrganizationName:       "Okta, Inc",
		OrganizationalUnitName: "Dev",
		CommonName:             "SP Issuer",
	}

	subjectAltNames := okta.CSRMetadataSubjectAltNames{
		DnsNames: []string{"dev.okta.com"},
	}

	csr := okta.CSRMetadata{
		Subject:         &subject,
		SubjectAltNames: &subjectAltNames,
	}

	csrs, _, err := client.Application.GenerateCsrForApplication(context.TODO(), application.Id, csr)
	require.NoError(t, err, "Creating an application CSR should not error")

	assert.IsType(t, &okta.CSR{}, csrs, "did not return a `okta.CSR` object")

	client.Application.DeactivateApplication(context.TODO(), application.Id)
	_, err = client.Application.DeleteApplication(context.TODO(), application.Id)

	require.NoError(t, err, "Deleting an application should not error")
}

func create_application(t *testing.T) *okta.BasicAuthApplication {
	client, _ := tests.NewClient()
	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(context.TODO(), basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	return application.(*okta.BasicAuthApplication)
}

func delete_all_apps(t *testing.T) {
	client, _ := tests.NewClient()

	applicationList, _, err := client.Application.ListApplications(context.TODO(), nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	for _, a := range applicationList {
		if strings.HasPrefix(a.(*okta.Application).Label, "Template Basic Auth App") ||
			strings.HasPrefix(a.(*okta.Application).Label, "Bookmark App") {
			client.Application.DeactivateApplication(context.TODO(), a.(*okta.Application).Id)
			client.Application.DeleteApplication(context.TODO(), a.(*okta.Application).Id)
		}
	}
}

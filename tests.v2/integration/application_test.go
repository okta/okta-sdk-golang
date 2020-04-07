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
	"fmt"
	"strings"
	"testing"

	"github.com/okta/okta-sdk-golang/okta.v2/query"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/okta.v2"

	"github.com/okta/okta-sdk-golang/tests.v2"
)

const myCApem = `-----BEGIN CERTIFICATE-----
MIIDGjCCAgICCQCf0kZljlDWgTANBgkqhkiG9w0BAQsFADBPMQswCQYDVQQGEwJV
UzENMAsGA1UECAwET2hpbzETMBEGA1UEBwwKU3ByaW5nYm9ybzENMAsGA1UECgwE
U2VsZjENMAsGA1UECwwEU2VsZjAeFw0yMDAzMTkxNjI1MzZaFw0yNTAzMTgxNjI1
MzZaME8xCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMRMwEQYDVQQHDApTcHJp
bmdib3JvMQ0wCwYDVQQKDARTZWxmMQ0wCwYDVQQLDARTZWxmMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8DC7b9I93e4jZWYZ/601sAYPDg4g+dXZBPv9
qo4N0D19bMx1LsRSk7b8BKCtmJW/Oej70Mh8hTR2Vfkgc6ex/Pmt5pIMnwwpU6Uq
BpjmYOw7jGUxutLSXVsf8tGt+76HquqpIyGwo9aZA4ZdVsC/h4IZFcD/qeeJ7asG
9jyHMvXHK9eEc8RURBKiPU3kEak6/bmf+6XPp8GcNn+N0pLhjlVjLv0TxojrkjRU
ONTuDN5yMaIru7DYnAx15AiwqS52HEsSXYXqb4oYE1B5sePaytTUyVchdyTSpo6d
d84+kVxM+5FGwKq9z1WOm83KQ3pJnBSE/puHf2yzrkBSZzjeZwIDAQABMA0GCSqG
SIb3DQEBCwUAA4IBAQCmRcoZvRgJMeOneQXQBu+weqyyzov8mbHn4aw7DDdzgwh7
YgoOdMT5I4Q0W/ct9yskvm6OQVtRao8n+1HHrgzsVw6B1cIaEpCLrNXoeD28eglf
ZUofH0ZOxEavZt2p4qRADCG0MnPY3re10YY1i19fFOIvM5aSc0W4Ep0QG4gnjlHb
2hBWj88kPWK+ECvaI9fDiwURfOVmTp2KgVmGCtMZf2WFDDuFOuCE9MrqLSXoUyJX
lCQEpomZRF6bwBBz2Ca0iQVoSM9C80wXAyP9naeGLRdnrXVo99KL6HYSti3iD2yO
YqGxQq1rRjEgipWQ3qhjV+7l57Na0BYaFX1+HFnJ
-----END CERTIFICATE-----`

const myCAkey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA8DC7b9I93e4jZWYZ/601sAYPDg4g+dXZBPv9qo4N0D19bMx1
LsRSk7b8BKCtmJW/Oej70Mh8hTR2Vfkgc6ex/Pmt5pIMnwwpU6UqBpjmYOw7jGUx
utLSXVsf8tGt+76HquqpIyGwo9aZA4ZdVsC/h4IZFcD/qeeJ7asG9jyHMvXHK9eE
c8RURBKiPU3kEak6/bmf+6XPp8GcNn+N0pLhjlVjLv0TxojrkjRUONTuDN5yMaIr
u7DYnAx15AiwqS52HEsSXYXqb4oYE1B5sePaytTUyVchdyTSpo6dd84+kVxM+5FG
wKq9z1WOm83KQ3pJnBSE/puHf2yzrkBSZzjeZwIDAQABAoIBAQDt/Ns7qO10AIlB
5zDLjSwtBVPVcVprMeCed7CYVbiKJOMp8kwJ0qyfgCelzi8ziOy4zIj2DjCTK7A0
72ugLQDGz/3m/79RuBMatgQ2FTnvvyIhsgLcQhf+OFQnnGrvjZGPYIPGM5N6Qx/J
xlClrMYZ1mZUj67DApA/1b5ILSEo80LCE+2KwHzXcv1xySV5G1lsdr/GDaQYjZlw
mZGNevY9r94REhe2240zu+gq9YHA2TJgrjjViSGsupPY3EE8vYMYWv/N1h8sLFkD
ntehUX32nK75Wq3NaB0OjBBreGtC52/wQwG2EOV3KRcw8WxqPhQzHYSO73aWlqbh
v97WGqGBAoGBAPpdTUhO1OEHpH8aCEqP8Sm2Q7hKK0W3ua9++4jEWHGcnvRYvECJ
OYyiHxzdMadyy8MvZKI7dWDQiqxEcwpm7rbJCJXzOjy+hdbJvV5myeHlwvzfaWe7
FR3t4XtC/kBtidm/hkhRkqlC3U1UJ9nDiCkbfNkrakJlbRY5+coTZW3vAoGBAPWY
zZprNiFMc9um86IjsuKaP8SSXDColsmO6dmRQQOjOFuYN0QMwlBGbnyvQZTrRgwq
s8UYJY0323tEgBGUynSWOVbRV0PWoOM1G0cx1aFv/Mj43xhJlXvMmrgUi4bxIPBZ
lUiTy32CaiHT5glekWdde6qDqtyO7Gr9maKDOQ8JAoGBAJLGoikS9iBaz6goBdZY
nsSacwcWjFnaBQUKx8H9gfBRJqsPXoXjLRbycJUGZDbLyQNLxI6LlxvEBphJpLvj
bm1AXEU0i97SvzoVmWw/jHlfrrl67JuAhTe/nuIZe18gGKHMc5fwIrASYBUWkipL
RIb882uJ1UjJl3NhV7yNNHiHAoGBAITTIz9EhH31zyMYY+No0zJioeI6FcnrI8HW
nPqh6DuDZtOCu0D+dYjczpx4XEuiArxJy/foW0bI0tcT8P+RLP1o0ZH2ne9+gHzh
F+OlPBiXbGt0zZNhGItf2L19vwg4GMxkZqxd4kv64FNzOpIOpyz0DhHmK94lHg+v
IAwYVB+hAoGAbfEb0nrWg2S4euP2KrZ2jnEzOgR7RLm3QYuxV+0UQQ4lSoGE6thq
Nd9jmO2rvpK47rjR2Db7p3/E/xfp27x/RMi5KvGF5l6E1xsXXrMCR2M6ghWdqv/C
wjcsLCBxDJreGsTbPeET2Sae6pVvSkKdMJVIwyTMZxyKOukerA9X1g4=
-----END RSA PRIVATE KEY-----`

func Test_can_get_applicaiton_by_id(t *testing.T) {
	client, _ := tests.NewClient()

	basicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	basicApplicationSettingsApplication.AuthURL = "https://example.com/auth.html"
	basicApplicationSettingsApplication.Url = "https://example.com/auth.html"

	basicApplicationSettings := okta.NewBasicApplicationSettings()
	basicApplicationSettings.App = basicApplicationSettingsApplication

	basicApplication := okta.NewBasicAuthApplication()
	basicApplication.Settings = basicApplicationSettings

	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	foundApplication, _, err := client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)
	require.NoError(t, err, "Should not error when getting an applicaiton by id")

	assert.Equal(t, appId, foundApplication.(*okta.BasicAuthApplication).Id, "Application found was not correct")

	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

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

	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	newBasicApplicationSettingsApplication := okta.NewBasicApplicationSettingsApplication()
	newBasicApplicationSettingsApplication.AuthURL = "https://example.org/auth.html"
	newBasicApplicationSettingsApplication.Url = "https://example.org/auth.html"

	newBasicApplicationSettings := okta.NewBasicApplicationSettings()
	newBasicApplicationSettings.App = newBasicApplicationSettingsApplication

	newBasicApplication := okta.NewBasicAuthApplication()
	newBasicApplication.Settings = newBasicApplicationSettings

	newApp, _, err := client.Application.UpdateApplication(appId, newBasicApplication)
	require.NoError(t, err, "Updating an application caused an error")

	assert.Equal(t, "https://example.org/auth.html", newApp.(*okta.BasicAuthApplication).Settings.App.Url, "The application did not update")

	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

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
	application, _, err := client.Application.CreateApplication(bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BookmarkApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BookmarkApplication).Id)

	client.Application.DeactivateApplication(application.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(application.(*okta.BookmarkApplication).Id)

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
	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BasicAuthApplication{}, application, "Application type returned was incorrect")
	assert.NotEmpty(t, application.(*okta.BasicAuthApplication).Id)
	assert.NotEmpty(t, basicApplication.Id)

	client.Application.DeactivateApplication(application.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(application.(*okta.BasicAuthApplication).Id)

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

	app1, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	bookmarkApplicationSettingsApplication := okta.NewBookmarkApplicationSettingsApplication()
	bookmarkApplicationSettingsApplication.RequestIntegration = new(bool)
	bookmarkApplicationSettingsApplication.Url = "https://example.com/bookmark.htm"

	bookmarkApplicationSettings := okta.NewBookmarkApplicationSettings()
	bookmarkApplicationSettings.App = bookmarkApplicationSettingsApplication

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = bookmarkApplicationSettings

	app2, _, err := client.Application.CreateApplication(bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	applist, _, err := client.Application.ListApplications(nil)
	require.NoError(t, err, "List applicaitons should not error")

	for _, a := range applist {

		if a.(*okta.Application).Name == "bookmark" {
			if a.(*okta.Application).Id == app2.(*okta.BookmarkApplication).Id {
				ba, _, _ := client.Application.GetApplication(a.(*okta.Application).Id, okta.NewBookmarkApplication(), nil)
				requestIntegration := ba.(*okta.BookmarkApplication).Settings.App.RequestIntegration
				assert.False(t, *requestIntegration)
			}
		}
	}

	client.Application.DeactivateApplication(app1.(*okta.BasicAuthApplication).Id)
	_, err = client.Application.DeleteApplication(app1.(*okta.BasicAuthApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")

	client.Application.DeactivateApplication(app2.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(app2.(*okta.BookmarkApplication).Id)

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

	application, _, err := client.Application.CreateApplication(basicApplication, query.NewQueryParams(query.WithActivate(false)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.ActivateApplication(appId)
	require.NoError(t, err, "Activationg an application should not error")
	application, _, _ = client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

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

	application, _, err := client.Application.CreateApplication(basicApplication, query.NewQueryParams(query.WithActivate(true)))
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	assert.Equal(t, "ACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeactivateApplication(appId)
	require.NoError(t, err, "Deactivating an application should not error")
	application, _, err = client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "INACTIVE", application.(*okta.BasicAuthApplication).Status, "Application is not inactive")

	_, err = client.Application.DeleteApplication(appId)

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

	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	appUserList, _, _ := client.Application.ListApplicationUsers(appId, nil)
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

	user, _, err := client.User.CreateUser(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	appUserPasswordCredentials := okta.NewAppUserPasswordCredential()
	appUserPasswordCredentials.Value = "abcd1234"

	appUserCredentials := okta.NewAppUserCredentials()
	appUserCredentials.UserName = "appUser"
	appUserCredentials.Password = appUserPasswordCredentials

	appUser := okta.NewAppUser()
	appUser.Credentials = appUserCredentials
	appUser.Id = user.Id

	appUser, _, err = client.Application.AssignUserToApplication(appId, *appUser)
	require.NoError(t, err, "Assigning user to application show not error")

	assert.IsType(t, okta.AppUser{}, *appUser, "Type returned from assigning user to application was incorrect")

	appUserList, _, _ = client.Application.ListApplicationUsers(appId, nil)
	assert.NotEmpty(t, appUserList, "App Users should not be empty")

	client.Application.DeleteApplicationUser(appId, appUser.Id, nil)
	appUserList, _, _ = client.Application.ListApplicationUsers(appId, nil)
	assert.Empty(t, appUserList, "App Users should be empty after deleting")

	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

	require.NoError(t, err, "Deleting an application should not error")

	client.User.DeactivateUser(user.Id, nil)
	client.User.DeactivateOrDeleteUser(user.Id, nil)
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
	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, okta.BasicApplicationSettingsApplication{}, *application.(*okta.BasicAuthApplication).Settings.App, "The returned type of application settings application was not correct type")
	assert.Equal(t, "https://example.com/auth.html", application.(*okta.BasicAuthApplication).Settings.App.Url)

	appId := application.(*okta.BasicAuthApplication).Id
	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

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
	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	appId := application.(*okta.BasicAuthApplication).Id

	myApp, _, err := client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)
	app := myApp.(*okta.BasicAuthApplication)
	myAppId := app.Id
	myAppSettingsApp := app.Settings.App

	myAppSettingsApp.Url = "https://okta.com/auth"

	_, _, _ = client.Application.UpdateApplication(myAppId, app)

	updatedApp, _, err := client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)
	assert.Equal(t, "https://okta.com/auth", updatedApp.(*okta.BasicAuthApplication).Settings.App.Url, "The URL was not updated'")

	client.Application.DeactivateApplication(appId)
	_, err = client.Application.DeleteApplication(appId)

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

	csrs, _, err := client.Application.GenerateCsrForApplication(application.Id, csr)
	require.NoError(t, err, "Creating an application CSR should not error")

	assert.IsType(t, &okta.CSR{}, csrs, "did not return a `okta.CSR` object")

	client.Application.DeactivateApplication(application.Id)
	_, err = client.Application.DeleteApplication(application.Id)

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

	application, _, err := client.Application.CreateApplication(basicApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	return application.(*okta.BasicAuthApplication)
}

func delete_all_apps(t *testing.T) {
	client, _ := tests.NewClient()

	applicationList, _, err := client.Application.ListApplications(nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	for _, a := range applicationList {
		if strings.HasPrefix(a.(*okta.Application).Label, "Template Basic Auth App") ||
			strings.HasPrefix(a.(*okta.Application).Label, "Bookmark App") {
			client.Application.DeactivateApplication(a.(*okta.Application).Id)
			client.Application.DeleteApplication(a.(*okta.Application).Id)
		}
	}
}

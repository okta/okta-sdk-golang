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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/okta"

	"github.com/okta/okta-sdk-golang/tests"
)

func Test_can_create_a_bookmark_application(t *testing.T) {
	client := tests.NewClient()

	bookmarkApplicationSettingsApplication := okta.BookmarkApplicationSettingsApplication{
		RequestIntegration: new(bool),
		Url:                "https://example.com/bookmark.htm",
	}

	bookmarkApplicationSettings := okta.BookmarkApplicationSettings{
		App: &bookmarkApplicationSettingsApplication,
	}

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = &bookmarkApplicationSettings
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
	client := tests.NewClient()

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

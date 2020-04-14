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

package unit

import (
	"testing"

	"github.com/okta/okta-sdk-golang/okta"

	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
)

func Test_error_on_empty_url(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl(""))
	assert.Error(t, err, "Does not error when org url is missing")
}

func Test_error_when_url_contains_yourOktaDomain(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl("https://{yourOktaDomain}"))
	assert.Error(t, err, "Does not error when org url contains {yourOktaDomain}")
}

func Test_error_when_url_contains_admin_okta_com(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl("https://test-admin.okta.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.okta.com")
}

func Test_error_when_url_contains_admin_oktapreview_com(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl("https://test-admin.oktapreview.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.oktapreview.com")
}

func Test_error_when_url_contains_admin_okta_emea_com(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl("https://test-admin.okta-emea.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.okta-emea.com")
}

func Test_error_when_url_contains_com_com(t *testing.T) {
	_, err := tests.NewClient(okta.WithOrgUrl("https://test.okta.com.com"))
	assert.Error(t, err, "Does not error when org url contains .com.com")
}

func Test_error_when_url_does_not_begin_with_https(t *testing.T) {
	_, err := tests.NewClient(okta.WithTestingDisableHttpsCheck(false), okta.WithOrgUrl("http://test.okta.com"))
	assert.Error(t, err, "Does not error when url contains only http")
}

func Test_error_when_api_token_is_empty(t *testing.T) {
	_, err := tests.NewClient(okta.WithToken(""))
	assert.Error(t, err, "Does not error when api token is empty")
}

func Test_error_when_api_token_contains_placeholder(t *testing.T) {
	_, err := tests.NewClient(okta.WithToken("{apiToken}"))
	assert.Error(t, err, "Does not error when api token contains {apiToken}")
}

func Test_error_when_authorization_mode_is_not_valid(t *testing.T) {
	_, err := tests.NewClient(okta.WithAuthorizationMode("invalid"))
	assert.Error(t, err, "Does not error when authorization mode is invalid")
}

func Test_does_not_error_when_authorization_mode_is_valid(t *testing.T) {
	_, err := tests.NewClient(okta.WithAuthorizationMode("SSWS"))
	assert.NoError(t, err, "Should not error when authorization mode is SSWS")
}

func Test_will_error_if_private_key_authorization_type_with_missing_properties(t *testing.T) {
	_, err := tests.NewClient(okta.WithAuthorizationMode("PrivateKey"), okta.WithClientId(""))
	assert.Error(t, err, "Does not error if private key selected with no other required options")
}

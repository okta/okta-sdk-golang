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

func Test_panic_on_empty_url(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl(""))
	}, "Does not panic when org url is missing")
}

func Test_panic_when_url_contains_yourOktaDomain(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl("https://{yourOktaDomain}"))
	}, "Does not panic when org url contains {yourOktaDomain}")
}

func Test_panic_when_url_contains_admin_okta_com(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl("https://test-admin.okta.com"))
	}, "Does not panic when org url contains test-admin.okta.com")
}

func Test_panic_when_url_contains_admin_oktapreview_com(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl("https://test-admin.oktapreview.com"))
	}, "Does not panic when org url contains test-admin.oktapreview.com")
}

func Test_panic_when_url_contains_admin_okta_emea_com(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl("https://test-admin.okta-emea.com"))
	}, "Does not panic when org url contains test-admin.okta-emea.com")
}

func Test_panic_when_url_contains_com_com(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithOrgUrl("https://test.okta.com.com"))
	}, "Does not panic when org url contains .com.com")
}

func Test_panic_when_url_does_not_begin_with_https(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithTestingDisableHttpsCheck(false), okta.WithOrgUrl("http://test.okta.com"))
	}, "Does not panic when url contains only http")
}

func Test_panic_when_api_token_is_empty(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithToken(""))
	}, "Does not panic when api token is empty")
}

func Test_panic_when_api_token_contains_placeholder(t *testing.T) {
	assert.Panics(t, func() {
		_, _ = tests.NewClient(okta.WithToken("{apiToken}"))
	}, "Does not panic when api token contains {apiToken}")
}

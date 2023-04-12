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
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/okta/okta-sdk-golang/v2/okta"

	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
)

func Test_error_on_empty_url(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl(""))
	assert.Error(t, err, "Does not error when org url is missing")
}

func Test_error_when_url_contains_yourOktaDomain(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://{yourOktaDomain}"))
	assert.Error(t, err, "Does not error when org url contains {yourOktaDomain}")
}

func Test_error_when_url_contains_admin_okta_com(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://test-admin.okta.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.okta.com")
}

func Test_error_when_url_contains_admin_oktapreview_com(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://test-admin.oktapreview.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.oktapreview.com")
}

func Test_error_when_url_contains_admin_okta_emea_com(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://test-admin.okta-emea.com"))
	assert.Error(t, err, "Does not error when org url contains test-admin.okta-emea.com")
}

func Test_error_when_url_contains_com_com(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://test.okta.com.com"))
	assert.Error(t, err, "Does not error when org url contains .com.com")
}

func Test_error_when_url_does_not_begin_with_https(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithTestingDisableHttpsCheck(false), okta.WithOrgUrl("http://test.okta.com"))
	assert.Error(t, err, "Does not error when url contains only http")
}

func Test_error_when_api_token_is_empty(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithToken(""))
	assert.Error(t, err, "Does not error when api token is empty")
}

func Test_error_when_api_token_contains_placeholder(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithToken("{apiToken}"))
	assert.Error(t, err, "Does not error when api token contains {apiToken}")
}

func Test_error_when_authorization_mode_is_not_valid(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithAuthorizationMode("invalid"))
	assert.Error(t, err, "Does not error when authorization mode is invalid")
}

func Test_does_not_error_when_authorization_mode_is_valid(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithAuthorizationMode("SSWS"))
	assert.NoError(t, err, "Should not error when authorization mode is SSWS")
}

func Test_does_not_error_when_authorization_mode_is_brearer(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithAuthorizationMode("Bearer"))
	assert.NoError(t, err, "Should not error when authorization mode is Bearer")
}

func Test_will_error_if_private_key_authorization_type_with_missing_properties(t *testing.T) {
	_, _, err := tests.NewClient(context.TODO(), okta.WithAuthorizationMode("PrivateKey"), okta.WithClientId(""))
	assert.Error(t, err, "Does not error if private key selected with no other required options")
}

type InterceptingRoundTripperTest struct {
	Name                    string
	Blocking                bool
	Interceptor             func(*http.Request) error
	ExpectedTransportCalls  int
	ExpectInterceptorCalled bool
	ExpectSdkErrorThrown    bool
}

func Test_Intercepting_RoundTripper(t *testing.T) {
	interceptorCalled := false
	testsToRun := []InterceptingRoundTripperTest{
		{
			Name:     "Calls interceptor",
			Blocking: false,
			Interceptor: func(r *http.Request) error {
				interceptorCalled = true
				return nil
			},
			ExpectedTransportCalls:  1,
			ExpectInterceptorCalled: true,
			ExpectSdkErrorThrown:    false,
		},
		{
			Name:     "Does not call transport when interceptor panics when blocking",
			Blocking: true,
			Interceptor: func(r *http.Request) error {
				interceptorCalled = true
				panic("Some err")
			},
			ExpectedTransportCalls:  0,
			ExpectInterceptorCalled: true,
			ExpectSdkErrorThrown:    true,
		},
		{
			Name:     "Calls transport when interceptor panics when non blocking",
			Blocking: false,
			Interceptor: func(r *http.Request) error {
				interceptorCalled = true
				panic("Some err")
			},
			ExpectedTransportCalls:  1,
			ExpectInterceptorCalled: true,
			ExpectSdkErrorThrown:    false,
		},
		{
			Name:     "Does not call transport when interceptor throws err when blocking",
			Blocking: true,
			Interceptor: func(r *http.Request) error {
				interceptorCalled = true
				return fmt.Errorf("Some error")
			},
			ExpectedTransportCalls:  0,
			ExpectInterceptorCalled: true,
			ExpectSdkErrorThrown:    true,
		},
		{
			Name:     "Calls transport when interceptor throws err when not blocking",
			Blocking: false,
			Interceptor: func(r *http.Request) error {
				interceptorCalled = true
				return fmt.Errorf("Some error")
			},
			ExpectedTransportCalls:  1,
			ExpectInterceptorCalled: true,
			ExpectSdkErrorThrown:    false,
		},
	}

	for _, test := range testsToRun {
		t.Run(
			test.Name,
			func(t *testing.T) {
				mockHttpClient := http.DefaultClient
				mockTransport := httpmock.DefaultTransport
				mockTransport.RegisterNoResponder(func(r *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200}, nil
				})
				mockHttpClient.Transport = mockTransport

				_, oktaClient, err := tests.NewClient(
					context.TODO(),
					okta.WithHttpInterceptorAndHttpClientPtr(test.Interceptor, mockHttpClient, test.Blocking),
				)
				assert.NoError(t, err)

				_, _, err = oktaClient.IdentityProvider.ActivateIdentityProvider(context.TODO(), "Anything")

				if test.ExpectSdkErrorThrown {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}

				assert.Equal(t, test.ExpectInterceptorCalled, interceptorCalled)

				callCount := mockTransport.GetTotalCallCount()

				assert.Equal(t, test.ExpectedTransportCalls, callCount)

				interceptorCalled = false
				mockTransport.ZeroCallCounters()
			},
		)
	}
}

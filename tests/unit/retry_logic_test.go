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
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/okta"

	"github.com/jarcoal/httpmock"
	"github.com/okta/okta-sdk-golang/tests"
)

func Test_429_Will_Automatically_Retry(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := tests.NewClient(okta.WithCache(false))

	httpmock.RegisterResponder("GET", "/api/v1/users",
		tests.MockResponse(
			Mock429Response(),
			MockValidResponse(),
		),
	)

	_, resp, err := client.User.ListUsers(nil)
	require.Nil(t, err, "Error should have been nil")
	require.NotNil(t, resp, "Response was nil")

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	require.Equal(t, 2, info["GET /api/v1/users"], "did not make exactly 2 call to /api/v1/users")
}

func Test_Will_Stop_Retrying_Based_On_Max_Retry_Configuration(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := tests.NewClient(okta.WithRequestTimeout(0), okta.WithRateLimitMaxRetries(1))

	httpmock.RegisterResponder("GET", "/api/v1/users",
		tests.MockResponse(
			Mock429Response(),
			Mock429Response(),
		),
	)

	_, _, err := client.User.ListUsers(nil)
	require.NotNil(t, err, "error was nil, but should have told the user they reached their max retry limit")

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	require.False(t, info["GET /api/v1/users"] < 2, "should have done at least one retry call to /api/v1/users")
	require.Equal(t, 2, info["GET /api/v1/users"], "should have thrown error after first retry to /api/v1/users")
}

func Test_Will_Handle_Backoff_Strategy_For_429(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := tests.NewClient(okta.WithRequestTimeout(1), okta.WithRateLimitMaxRetries(3))

	httpmock.RegisterResponder("GET", "/api/v1/users",
		tests.MockResponse(
			Mock429Response(),
			Mock429Response(),
			Mock429Response(),
			MockValidResponse(),
		),
	)

	_, _, err := client.User.ListUsers(nil)
	require.NotNil(t, err, "error was nil, but should have told the user they reached their max retry limit")

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()

	require.Equal(t, 1, info["GET /api/v1/users"], "should have thrown error before first retry to /api/v1/users due to request timeout")
}

func Test_a_429_with_x_reset_header_throws_error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := tests.NewClient()

	httpmock.RegisterResponder("GET", "/api/v1/users",
		tests.MockResponse(
			Mock429ResponseNoResetHeader(),
		),
	)

	_, _, err := client.User.ListUsers(nil)

	require.NotNil(t, err, "error should not be nil. It should let user know the reset header is required")
}

func Test_a_429_with_no_date_header_throws_error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := tests.NewClient()

	httpmock.RegisterResponder("GET", "/api/v1/users",
		tests.MockResponse(
			Mock429ResponseNoDateHeader(),
		),
	)

	_, _, err := client.User.ListUsers(nil)

	require.NotNil(t, err, "error should not be nil. It should let user know the date header is required")
}

func Test_gets_the_correct_backoff_time(t *testing.T) {
	backoff := okta.Get429BackoffTime(Mock429Response())

	require.Equal(t, int64(2), backoff, "backoff time should have only been 1 second")
}

func Test_with_multiple_x_rate_limit_request_times_still_retries(t *testing.T) {
	backoff := okta.Get429BackoffTime(Mock429ResponseMultipleHeaders())

	require.Equal(t, int64(11), backoff, "Backoff time should handle the correct header")
}

func Mock429Response() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+1, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 Z"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString(""),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseNoResetHeader() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 Z"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString(""),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseNoDateHeader() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+1, 10))
	header.Add("X-Okta-Request-id", "a-request-id")

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString(""),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseMultipleHeaders() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+20, 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+10, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 Z"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString(""),
		Header:        header,
		ContentLength: -1,
	}
}

func MockValidResponse() *http.Response {
	header := http.Header{}
	header.Add("X-Okta-Request-id", "another-request-id")
	header.Add("Date", time.Now().Add(time.Duration(time.Second*10)).Format(time.RFC3339))

	return &http.Response{
		Status:        strconv.Itoa(200),
		StatusCode:    200,
		Body:          httpmock.NewRespBodyFromString("[]"),
		Header:        header,
		ContentLength: -1,
	}
}

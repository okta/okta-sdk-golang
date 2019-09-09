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
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	testUserId = "abc123"
	testUrl    = "https://dev.oktapreview.com/"
	testToken  = "xxx"
)

var (
	backoffCalls []*Call
)

type (
	Call struct {
		args []interface{}
	}
)

func MockBackoff(t time.Duration) {
	backoffCalls = append(backoffCalls, &Call{[]interface{}{t}})
}

func before() {
	gock.Off()
	gock.DisableNetworking()
	okta.Backoff = MockBackoff
	backoffCalls = []*Call{}
}

func testUser() *okta.User {
	return &okta.User{
		Id: "zebraman",
		Profile: &okta.UserProfile{
			"firstName": "Zebra",
			"lastName":  "Centaur",
			"email":     "test@g.com",
		},
	}
}

func mockError(times int) {
	gock.New(testUrl).
		Times(times).
		ReplyError(errors.New("stuff is jacked bro"))
}

func mockErrorResponse(times, statusCode int) {
	gock.New(testUrl).
		Times(times).
		Get("/api/v1/users").
		Reply(statusCode).
		JSON(okta.Error{})
}

func mockSuccess(times int) {
	gock.New(testUrl).
		Times(times).
		Get(fmt.Sprintf("/api/v1/users/%s", testUserId)).
		Reply(200).
		JSON(testUser())
}

func Test_backoff_with_retry(t *testing.T) {
	before()
	mockErrorResponse(2, http.StatusTooManyRequests)
	mockSuccess(1)
	client, err := tests.NewClient(
		okta.WithOrgUrl(testUrl),
		okta.WithToken(testToken),
		okta.WithBackoff(true),
		okta.WithRetries(3),
	)
	assert.Equal(t, err, nil, "Error should be nil")
	_, resp, err := client.User.GetUser(testUserId)
	assert.Equal(t, err, nil, "Error for GetUser call should be nil")
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status code should be 200")
	assert.Equal(t, gock.IsDone(), true, "All gock calls should have completed")
	assert.Equal(t, len(backoffCalls), 2, "Backoff should have been called twice")
}

func Test_backoff_with_retry_and_response_body(t *testing.T) {
	before()
	gock.New(testUrl).
		Times(2).
		Post("/api/v1/users").
		Reply(http.StatusTooManyRequests)
	gock.New(testUrl).
		Times(1).
		Post("/api/v1/users").
		Reply(http.StatusCreated).
		JSON(testUser())
	client, err := tests.NewClient(
		okta.WithOrgUrl(testUrl),
		okta.WithToken(testToken),
		okta.WithBackoff(true),
		okta.WithRetries(3),
	)
	assert.Equal(t, err, nil, "Error should be nil")
	_, resp, err := client.User.CreateUser(*testUser(), nil)
	assert.Equal(t, err, nil, "Error for GetUser call should be nil")
	assert.Equal(t, resp.StatusCode, http.StatusCreated, "Status code should be 200")
	assert.Equal(t, gock.IsDone(), true, "All gock calls should have completed")
	assert.Equal(t, len(backoffCalls), 2, "Backoff should have been called twice")
}

func Test_retries_without_backoff(t *testing.T) {
	before()
	mockError(2)
	mockSuccess(1)
	client, err := tests.NewClient(
		okta.WithOrgUrl(testUrl),
		okta.WithToken(testToken),
		okta.WithRetries(3),
	)
	assert.Equal(t, err, nil, "Error should be nil")
	_, resp, err := client.User.GetUser(testUserId)
	assert.Equal(t, err, nil, "Error for GetUser call should be nil")
	assert.Equal(t, resp.StatusCode, http.StatusOK, "Status code should be 200")
	assert.Equal(t, gock.IsDone(), true, "All gock calls should have completed")
}

func Test_default_no_backoff_or_retries(t *testing.T) {
	before()
	mockErrorResponse(1, http.StatusTooManyRequests)
	client, err := tests.NewClient(
		okta.WithOrgUrl(testUrl),
		okta.WithToken(testToken),
	)
	assert.Equal(t, err, nil, "Error should be nil")
	_, resp, err := client.User.GetUser(testUserId)
	assert.NotEqual(t, err, nil, "error should be returned for 429 response")
	assert.Equal(t, resp.StatusCode, http.StatusTooManyRequests, "Status code should be 429")
	assert.Equal(t, gock.IsDone(), true, "All gock calls should have completed")
}

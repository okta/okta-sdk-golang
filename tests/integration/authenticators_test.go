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
	"errors"
	"fmt"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAuthenticators(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	authenticators, resp, err := client.Authenticator.ListAuthenticators(ctx)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200, "List authenticators should have 200 status.")
	assert.True(t, len(authenticators) > 0, "One or more authenticators should be present.")

	_, err = phoneAuthenticator(authenticators)
	require.NoError(t, err, "There should at least be a phone authenticator.")
}

func TestGetAuthenticator(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	phoneAuthenticator, err := fetchPhoneAuthenticator(ctx, client)
	require.NoError(t, err)

	authenticator, resp, err := client.Authenticator.GetAuthenticator(ctx, phoneAuthenticator.Id)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200, "Get authenticator should have 200 status.")
	assert.Equal(t, authenticator.Id, phoneAuthenticator.Id, "Expected authenticator getter should equal phone authenticator.")
}

func TestActivateDeactivateAuthenticator(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	// Find an inactive authenticator first and activate it. Then deactivate it.
	// For integration test purposes, randomly deactivating an authenticator can
	// return an error if that authenticator is already used in a policy. So
	// start with a inactive authenticator as the test subject.

	authenticator, err := fetchInactiveAuthenticator(ctx, client)
	require.NoError(t, err)
	authenticatorId := authenticator.Id
	assert.Equal(t, "INACTIVE", authenticator.Status, "Expected authenticator status to be inactive.")

	authenticator, _, err = client.Authenticator.ActivateAuthenticator(ctx, authenticatorId)
	require.NoError(t, err)
	assert.Equal(t, "ACTIVE", authenticator.Status, "Expected authenticator status to be active.")

	authenticator, _, err = client.Authenticator.DeactivateAuthenticator(ctx, authenticatorId)
	require.NoError(t, err)
	assert.Equal(t, "INACTIVE", authenticator.Status, "Expected authenticator status to be inactive.")
}

// phoneAuthenticator helper to pick the phone authenticator from the
// authenticators list
func phoneAuthenticator(authenticators []*okta.Authenticator) (*okta.Authenticator, error) {
	for _, authenticator := range authenticators {
		if authenticator.Key == "phone_number" {
			return authenticator, nil
		}
	}

	return nil, errors.New("Phone number authenticator not found")
}

// fetchPhoneAuthenticator helper to get the phone authenticator from the API
func fetchPhoneAuthenticator(ctx context.Context, client *okta.Client) (*okta.Authenticator, error) {
	authenticators, _, err := client.Authenticator.ListAuthenticators(ctx)
	if err != nil {
		return nil, err
	}
	return phoneAuthenticator(authenticators)
}

// fetchActiveAuthenticator helper fetches the first active authenicator it
// finds.
func fetchActiveAuthenticator(ctx context.Context, client *okta.Client) (*okta.Authenticator, error) {
	return fetchAuthenticatorByStatus("ACTIVE", ctx, client)
}

// fetchInactiveAuthenticator helper fetches the first inactive authenicator it
// finds.
func fetchInactiveAuthenticator(ctx context.Context, client *okta.Client) (*okta.Authenticator, error) {
	return fetchAuthenticatorByStatus("INACTIVE", ctx, client)
}

// fetchAuthenticatorByStatus helper fetches the first authenicator it finds
// with the given status.
func fetchAuthenticatorByStatus(status string, ctx context.Context, client *okta.Client) (*okta.Authenticator, error) {
	authenticators, _, err := client.Authenticator.ListAuthenticators(ctx)
	if err != nil {
		return nil, err
	}

	for _, authenticator := range authenticators {
		if authenticator.Status == status {
			return authenticator, nil
		}
	}

	return nil, fmt.Errorf("Couldn't find an authenticator with satus %q.", status)
}

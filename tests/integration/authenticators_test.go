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

	// Find the phone authenticator. Activate it if inactive, then deactivate
	// it. Else, deactivate it, then activate it.

	authenticator, err := fetchPhoneAuthenticator(ctx, client)
	require.NoError(t, err)

	var ops []string

	if authenticator.Status == "INACTIVE" {
		ops = []string{"activate", "deactivate"}
	} else {
		ops = []string{"deactivate", "activate"}
	}

	for _, op := range ops {
		switch op {
		case "activate":
			authenticator, _, err = client.Authenticator.ActivateAuthenticator(ctx, authenticator.Id)
			require.NoError(t, err)
			assert.Equal(t, "ACTIVE", authenticator.Status, "Expected authenticator status to be inactive.")
		case "deactivate":
			authenticator, _, err = client.Authenticator.DeactivateAuthenticator(ctx, authenticator.Id)
			require.NoError(t, err)
			assert.Equal(t, "INACTIVE", authenticator.Status, "Expected authenticator status to be inactive.")
		}
	}
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

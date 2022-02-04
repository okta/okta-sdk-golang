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
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	PhoneNumberKey = "phone_number"
	EmailKey       = "okta_email"
)

func TestListAuthenticators(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	authenticators, resp, err := client.Authenticator.ListAuthenticators(ctx)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200, "List authenticators should have 200 status.")
	assert.True(t, len(authenticators) > 0, "One or more authenticators should be present.")

	_, err = pickAuthenticator(PhoneNumberKey, authenticators)
	require.NoError(t, err, "There should at least be a phone authenticator.")
}

func TestGetAuthenticator(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	phoneAuthenticator, err := fetchAuthenticator(PhoneNumberKey, ctx, client)
	require.NoError(t, err)

	authenticator, resp, err := client.Authenticator.GetAuthenticator(ctx, phoneAuthenticator.Id)
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200, "Get authenticator should have 200 status.")
	assert.Equal(t, authenticator.Id, phoneAuthenticator.Id, "Expected authenticator getter should equal phone authenticator.")
}

func TestUpdateAuthenticator(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	emailAuthenticator, err := fetchAuthenticator(EmailKey, ctx, client)
	require.NoError(t, err)

	tokenLifetimeInMinutes := emailAuthenticator.Settings.TokenLifetimeInMinutes + 1
	if tokenLifetimeInMinutes > 10 {
		tokenLifetimeInMinutes = int64(5)
	}

	updateAuthenticator := okta.Authenticator{
		Name: emailAuthenticator.Name,
		Settings: &okta.AuthenticatorSettings{
			TokenLifetimeInMinutes: tokenLifetimeInMinutes,
		},
	}
	authenticator, resp, err := client.Authenticator.UpdateAuthenticator(ctx, emailAuthenticator.Id, okta.Authenticator(updateAuthenticator))
	require.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, authenticator.Id, emailAuthenticator.Id)
	assert.Equal(t, authenticator.Settings.TokenLifetimeInMinutes, tokenLifetimeInMinutes, "Expected authenticator token life in minutes to be updated.")
}

func TestActivateDeactivateAuthenticator(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	err = setupDisablePhoneNumberOnMFAEnrollPolicy()
	require.NoError(t, err)

	// Find the phone authenticator. Activate it if inactive, then deactivate
	// it. Else, deactivate it, then activate it.
	authenticator, err := fetchAuthenticator(PhoneNumberKey, ctx, client)
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
			result, _, err := client.Authenticator.ActivateAuthenticator(ctx, authenticator.Id)
			require.NoError(t, err, "authenticator: "+authenticator.Id)
			assert.Equal(t, "ACTIVE", result.Status, "Expected authenticator status to be active ("+authenticator.Id+").")
		case "deactivate":
			result, _, err := client.Authenticator.DeactivateAuthenticator(ctx, authenticator.Id)
			require.NoError(t, err, "authenticator: "+authenticator.Id)
			assert.Equal(t, "INACTIVE", result.Status, "Expected authenticator status to be inactive ("+authenticator.Id+").")
		}
	}
}

// pickAuthenticator helper to pick the phone authenticator from the
// authenticators list
func pickAuthenticator(key string, authenticators []*okta.Authenticator) (*okta.Authenticator, error) {
	for _, authenticator := range authenticators {
		if authenticator.Key == key {
			return authenticator, nil
		}
	}

	return nil, fmt.Errorf("%q authenticator not found", key)
}

// fetchAuthenticator helper to get the phone authenticator from the API
func fetchAuthenticator(key string, ctx context.Context, client *okta.Client) (*okta.Authenticator, error) {
	authenticators, _, err := client.Authenticator.ListAuthenticators(ctx)
	if err != nil {
		return nil, err
	}
	return pickAuthenticator(key, authenticators)
}

func setupDisablePhoneNumberOnMFAEnrollPolicy() error {
	ctx, client, err := tests.NewClient(context.TODO())
	if err != nil {
		return err
	}

	rq := client.CloneRequestExecutor()
	url := "/api/v1/policies?type=MFA_ENROLL"
	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	var policies []interface{}
	_, err = client.GetRequestExecutor().Do(ctx, req, &policies)
	if err != nil {
		return err
	}

	// note: updating policy settings directly using generic golang structs
	policy := policies[0].(map[string]interface{})
	if settings, ok := policy["settings"].(map[string]interface{}); ok {
		if authenticators, ok := settings["authenticators"].([]interface{}); ok {
			for _, authenticator := range authenticators {
				key := authenticator.(map[string]interface{})["key"]
				if key != "phone_number" {
					continue
				}
				enroll := authenticator.(map[string]interface{})["enroll"].(map[string]interface{})
				enroll["self"] = "NOT_ALLOWED"
				break
			}
		}
	}

	url = fmt.Sprintf("/api/v1/policies/%s", policy["id"])
	req, err = rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, policy)
	if err != nil {
		return err
	}

	// note: have executor read an interface to prevent a panic
	var result interface{}
	_, err = client.GetRequestExecutor().Do(ctx, req, &result)

	return err
}

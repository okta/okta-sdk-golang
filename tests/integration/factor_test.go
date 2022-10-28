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
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_exercise_factor_lifecycle(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Factor-Lifecycle"
	profile["email"] = randomEmail()
	profile["login"] = profile["email"]
	u := &okta.CreateUserRequest{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.CreateUser(ctx, *u, qp)
	require.NoError(t, err, "Creating a new user should not error")

	allowedFactors, _, _ := client.UserFactor.ListSupportedFactors(ctx, user.Id)
	continueTesting := false
	for _, f := range allowedFactors {
		if f.(*okta.UserFactor).FactorType == "sms" {
			continueTesting = true
		}
	}

	if continueTesting {
		factors, _, listFactorsError := client.UserFactor.ListFactors(ctx, user.Id)
		require.NoError(t, listFactorsError, "Should not error when listing factors")

		assert.Empty(t, factors, "Factors list should be empty")

		factorProfile := okta.NewSmsUserFactorProfile()
		factorProfile.PhoneNumber = "16284001133"

		factor := okta.NewSmsUserFactor()
		factor.Profile = factorProfile

		addedFactor, resp, err := client.UserFactor.EnrollFactor(ctx, user.Id, factor, nil)
		require.NotEmpty(t, resp, "Response should not be empty")
		require.NoError(t, err, "Adding factor should not error")
		assert.IsType(t, okta.NewSmsUserFactor(), addedFactor)

		foundFactor, _, err := client.UserFactor.GetFactor(ctx, user.Id, addedFactor.(*okta.SmsUserFactor).Id, okta.NewSmsUserFactor())
		require.NoError(t, err, "Getting the factor should not error")

		client.UserFactor.DeleteFactor(ctx, user.Id, foundFactor.(*okta.SmsUserFactor).Id)
	} else {
		t.Skip("Skipping exercise factor lifecycle testing. SMS factor was not enabled")
	}

	client.User.DeactivateUser(ctx, user.Id, nil)
	client.User.DeactivateOrDeleteUser(ctx, user.Id, nil)
}

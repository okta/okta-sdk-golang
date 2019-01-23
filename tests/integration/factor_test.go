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

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/query"
	"github.com/okta/okta-sdk-golang/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_exercise_factor_lifecycle(t *testing.T) {
	client := tests.NewClient()

	p := &okta.PasswordCredential{
		Value: "Abcd1234",
	}
	uc := &okta.UserCredentials{
		Password: p,
	}
	profile := okta.UserProfile{}
	profile["firstName"] = "John"
	profile["lastName"] = "Factor-Lifecycle"
	profile["email"] = "john-factor-lifecycle@example.com"
	profile["login"] = "john-factor-lifecycle@example.com"
	u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
	}
	qp := query.NewQueryParams(query.WithActivate(false))

	user, _, err := client.User.Create(*u, qp)
	require.NoError(t, err, "Creating an user should not error")

	allowedFactors, _, err := client.Factor.ListSupported(user.Id)
	continueTesting := false
	for _, f := range allowedFactors {
		if f.(map[string]interface{})["factorType"] == "sms" {
			continueTesting = true
		}
	}

	if continueTesting {
		factors, _, err := client.Factor.List(user.Id)
		require.NoError(t, err, "Should not error when listing factors")

		assert.Empty(t, factors, "Factors list should be empty")

		factorProfile := okta.NewSmsFactorProfile()
		factorProfile.PhoneNumber = "15055550006"

		factor := okta.NewSmsFactor()
		factor.Profile = factorProfile

		addedFactor, _, err := client.Factor.Add(user.Id, factor, nil)
		require.NoError(t, err, "Adding factor should not error")
		assert.IsType(t, okta.NewSmsFactor(), addedFactor)

		foundFactor, _, err := client.Factor.Get(user.Id, addedFactor.(*okta.SmsFactor).Id, okta.NewSmsFactor())
		require.NoError(t, err, "Getting the factor should not error")

		client.Factor.Delete(user.Id, foundFactor.(*okta.SmsFactor).Id)
	} else {
		t.Skip("Skipping exercise factor lifecycle testing. SMS factor was not enabled")
	}

	client.User.Deactivate(user.Id)
	client.User.DeactivateOrDelete(user.Id)
}

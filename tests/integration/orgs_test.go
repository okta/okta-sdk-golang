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
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrgContactTypes(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	contactTypes, _, err := client.OrgSetting.GetOrgContactTypes(ctx)
	require.NoError(t, err)

	// there should be a billing and technical contact type
	var foundBilling, foundTechnical bool
	for _, concontactType := range contactTypes {
		if concontactType.ContactType == "BILLING" {
			foundBilling = true
		}
		if concontactType.ContactType == "TECHNICAL" {
			foundTechnical = true
		}
	}
	assert.True(t, foundBilling, "There should be billing contact type.")
	assert.True(t, foundTechnical, "There should be technical contact type.")
}

func TestGetAndUpdateOrgContactUser(t *testing.T) {
	// switch the billing and technical contacts

	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	technicalUser, _, err := client.OrgSetting.GetOrgContactUser(ctx, "TECHNICAL")
	require.NoError(t, err)

	billingUser, _, err := client.OrgSetting.GetOrgContactUser(ctx, "BILLING")
	require.NoError(t, err)

	_, _, err = client.OrgSetting.UpdateOrgContactUser(ctx, "TECHNICAL", okta.UserIdString{UserId: billingUser.UserId})
	require.NoError(t, err)

	_, _, err = client.OrgSetting.UpdateOrgContactUser(ctx, "BILLING", okta.UserIdString{UserId: technicalUser.UserId})
	require.NoError(t, err)
}

func TestOptInOptOutOktaCommunicationSettings(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	// get the communication settings
	// if opted in, opt out, then opt in
	// if opted opted out, opt in, then opt out

	settings, _, err := client.OrgSetting.GetOktaCommunicationSettings(ctx)
	require.NoError(t, err)

	var foundOptIn, foundOptOut bool
	links := settings.Links.(map[string]interface{})
	_, foundOptIn = links["optIn"]
	_, foundOptOut = links["optOut"]

	assert.True(t, (foundOptIn || foundOptOut), "There should be optIn or optOut setting")

	if foundOptIn {
		_, _, err := client.OrgSetting.OptInUsersToOktaCommunicationEmails(ctx)
		require.NoError(t, err)
		_, _, err = client.OrgSetting.OptOutUsersFromOktaCommunicationEmails(ctx)
		require.NoError(t, err)
	}

	if foundOptOut {
		_, _, err := client.OrgSetting.OptOutUsersFromOktaCommunicationEmails(ctx)
		require.NoError(t, err)
		_, _, err = client.OrgSetting.OptInUsersToOktaCommunicationEmails(ctx)
		require.NoError(t, err)
	}
}

func TestGetGrantExtentRevokeOktaCommunicationSettings(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	// get the support settings
	// if support not granted, grant it, extend it, revoke it
	// if support granted, revoke it, grant it, extend it

	settings, _, err := client.OrgSetting.GetOrgOktaSupportSettings(ctx)
	require.NoError(t, err)

	ops := []string{"revoke", "grant", "extend"}
	opIndex := 0
	if settings.Support != "ENABLED" {
		opIndex = 1
	}

	// will grant, extend, revoke
	// or will revoke, grant, extend
	countOps := 0
	for countOps < len(ops) {
		switch ops[opIndex] {
		case "revoke":
			_, _, err := client.OrgSetting.RevokeOktaSupport(ctx)
			require.NoError(t, err)
		case "grant":
			_, _, err := client.OrgSetting.GrantOktaSupport(ctx)
			require.NoError(t, err)
		case "extend":
			_, _, err := client.OrgSetting.ExtendOktaSupport(ctx)
			require.NoError(t, err)
		}

		opIndex++
		if opIndex >= len(ops) {
			opIndex = 0
		}
		countOps++
	}
}

func TestGetPerfsShowHideFooter(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	// get the org preferences
	// if hidden footer, show it, then hide it
	// if showing footer, hide it, then show it

	preferences, _, err := client.OrgSetting.GetOrgPreferences(ctx)
	require.NoError(t, err)

	var showEndUserFooter, hideEndUserFooter bool
	links := preferences.Links.(map[string]interface{})
	_, showEndUserFooter = links["showEndUserFooter"]
	_, hideEndUserFooter = links["hideEndUserFooter"]
	assert.True(t, (showEndUserFooter || hideEndUserFooter), "There should be show or hide end user preference")

	if showEndUserFooter {
		_, _, err = client.OrgSetting.HideOktaUIFooter(ctx)
		require.NoError(t, err)
		_, _, err := client.OrgSetting.ShowOktaUIFooter(ctx)
		require.NoError(t, err)
	}

	if hideEndUserFooter {
		_, _, err := client.OrgSetting.ShowOktaUIFooter(ctx)
		require.NoError(t, err)
		_, _, err = client.OrgSetting.HideOktaUIFooter(ctx)
		require.NoError(t, err)
	}
}

func TestGetAndPartialUpdateOrgSettings(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	// get org settings, partial update org with new website
	settings, _, err := client.OrgSetting.GetOrgSettings(ctx)
	require.NoError(t, err)

	website := fmt.Sprintf("https://developer.okta.com/golang-sdk-oie/testing/%d", time.Now().UnixNano())
	assert.True(t, settings.Website != website)

	partialSettings := okta.OrgSetting{Website: website}
	newSettings, _, err := client.OrgSetting.PartialUpdateOrgSetting(ctx, partialSettings)
	require.NoError(t, err)

	assert.True(t, newSettings.Website == website, "Expected org website %q, got %q", website, newSettings.Website)
}

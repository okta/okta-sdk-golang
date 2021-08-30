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
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanGetGroupProperties(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	gc, response, err := client.GroupSchema.GetGroupSchema(ctx)
	require.NoError(t, err, "getting group schema must not error")
	require.NotNil(t, gc, "group schema should not be nil")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.GroupSchema{}, gc, "did not return `*okta.GroupSchema{}` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/meta/schemas/group/default", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.Equal(t, "#custom", gc.Definitions.Custom.Id)
	assert.Equal(t, "#base", gc.Definitions.Base.Id)
	assert.Equal(t, "Name", gc.Definitions.Base.Properties["name"].Title)
	assert.Equal(t, "Description", gc.Definitions.Base.Properties["description"].Title)
}

func TestCanUpdateCustomGroupProperty(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	gc, _, err := client.GroupSchema.GetGroupSchema(ctx)
	require.NoError(t, err, "getting group schema must not error")
	require.NotNil(t, gc, "group schema should not be nil")

	testProperty1 := randomTestString()
	testProperty2 := randomTestString()

	gc.Definitions.Custom.Properties[testProperty1] = &okta.GroupSchemaAttribute{
		Description: "testing",
		Items: &okta.UserSchemaAttributeItems{
			Enum: []string{"test", "1", "2"},
			OneOf: []*okta.UserSchemaAttributeEnum{
				{
					Const: "test",
					Title: "test",
				},
				{
					Const: "1",
					Title: "1",
				},
				{
					Const: "2",
					Title: "2",
				},
			},
			Type: "string",
		},
		Master: &okta.UserSchemaAttributeMaster{
			Type: "OKTA",
		},
		Mutability: "READ_WRITE",
		Permissions: []*okta.UserSchemaAttributePermission{
			{
				Action:    "READ_ONLY",
				Principal: "SELF",
			},
		},
		Scope: "NONE",
		Title: "Property Title",
		Type:  "array",
	}
	gc.Definitions.Custom.Properties[testProperty2] = &okta.GroupSchemaAttribute{
		Description:  "User's username for twitter.com",
		ExternalName: "External Twitter username",
		Master: &okta.UserSchemaAttributeMaster{
			Type: "PROFILE_MASTER",
		},
		MaxLength:  20,
		MinLength:  1,
		Mutability: "READ_WRITE",
		Permissions: []*okta.UserSchemaAttributePermission{
			{
				Action:    "READ_WRITE",
				Principal: "SELF",
			},
		},
		Scope:  "NONE",
		Title:  "Twitter username",
		Type:   "string",
		Unique: "UNIQUE_VALIDATED",
	}
	updatedGC, _, err := client.GroupSchema.UpdateGroupSchema(ctx, *gc)
	require.NoError(t, err, "updating group schema must not error")
	require.NotNil(t, updatedGC, "updated group schema should not be nil")

	assert.Equal(t, "Property Title", updatedGC.Definitions.Custom.Properties[testProperty1].Title)
	assert.Equal(t, 3, len(updatedGC.Definitions.Custom.Properties[testProperty1].Items.Enum))
	assert.Equal(t, 3, len(updatedGC.Definitions.Custom.Properties[testProperty1].Items.OneOf))
	assert.Equal(t, "string", updatedGC.Definitions.Custom.Properties[testProperty1].Items.Type)
	// assert.Equal(t, "OKTA", updatedGC.Definitions.Custom.Properties[testProperty1].Master.Type)

	assert.Equal(t, "Twitter username", updatedGC.Definitions.Custom.Properties[testProperty2].Title)
	assert.Nil(t, updatedGC.Definitions.Custom.Properties[testProperty2].Items)
	// assert.Equal(t, "PROFILE_MASTER", updatedGC.Definitions.Custom.Properties[testProperty2].Master.Type)
	assert.Equal(t, int64(1), updatedGC.Definitions.Custom.Properties[testProperty2].MinLength)
	assert.Equal(t, int64(20), updatedGC.Definitions.Custom.Properties[testProperty2].MaxLength)
	assert.Equal(t, "UNIQUE_VALIDATED", updatedGC.Definitions.Custom.Properties[testProperty2].Unique)

	updatedGC.Definitions.Custom.Properties[testProperty1] = nil
	updatedGC.Definitions.Custom.Properties[testProperty2] = nil

	noCustomGC, _, err := client.GroupSchema.UpdateGroupSchema(ctx, *updatedGC)
	require.NoError(t, err, "updating group schema must not error")
	require.NotNil(t, noCustomGC, "updated group schema should not be nil")

	assert.Nil(t, noCustomGC.Definitions.Custom.Properties[testProperty1], "property should be removed")
	assert.Nil(t, noCustomGC.Definitions.Custom.Properties[testProperty2], "property should be removed")
}

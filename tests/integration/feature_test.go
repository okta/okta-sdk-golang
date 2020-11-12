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

func Test_can_list_all_features_for_organization(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	features, response, err := client.Feature.ListFeatures(ctx)

	require.NoError(t, err, "listing features must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Feature{}, features, "did not return `[]*okta.Feature` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/features", response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, features, "listing features result should not be empty")

	assert.NotEmpty(t, features[0].Description, "description should not be empty")
	assert.NotEmpty(t, features[0].Id, "id should not be empty")
	assert.NotEmpty(t, features[0].Name, "name should not be empty")
	assert.NotEmpty(t, features[0].Stage, "stage should not be empty")
	assert.NotEmpty(t, features[0].Status, "status should not be empty")
	assert.NotEmpty(t, features[0].Type, "type should not be empty")

}

func Test_can_get_a_feature(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	features, _, err := client.Feature.ListFeatures(ctx)
	require.NoError(t, err, "listing features must not error")

	firstFeatureId := features[0].Id
	feature, response, err := client.Feature.GetFeature(ctx, firstFeatureId)

	require.NoError(t, err, "getting a feature must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Feature{}, feature, "did not return `okta.Feature` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/features/"+firstFeatureId, response.Response.Request.URL.Path, "path for request was incorrect")

	assert.NotEmpty(t, feature, "listing features result should not be empty")

	assert.NotEmpty(t, feature.Description, "description should not be empty")
	assert.NotEmpty(t, feature.Id, "id should not be empty")
	assert.NotEmpty(t, feature.Name, "name should not be empty")
	assert.NotEmpty(t, feature.Stage, "stage should not be empty")
	assert.NotEmpty(t, feature.Status, "status should not be empty")
	assert.NotEmpty(t, feature.Type, "type should not be empty")

}

func Test_can_get_feature_dependencies(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	features, _, err := client.Feature.ListFeatures(ctx)
	require.NoError(t, err, "listing features must not error")

	firstFeatureId := features[0].Id
	featureDependencies, response, err := client.Feature.ListFeatureDependencies(ctx, firstFeatureId)

	require.NoError(t, err, "getting a features dependencies must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Feature{}, featureDependencies, "did not return `[]*okta.Feature` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/features/"+firstFeatureId+"/dependencies", response.Response.Request.URL.Path, "path for request was incorrect")

}

func Test_can_get_feature_dependants(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	features, _, err := client.Feature.ListFeatures(ctx)
	require.NoError(t, err, "listing features must not error")

	firstFeatureId := features[0].Id
	featureDependants, response, err := client.Feature.ListFeatureDependents(ctx, firstFeatureId)

	require.NoError(t, err, "getting a features dependants must not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, []*okta.Feature{}, featureDependants, "did not return `[]*okta.Feature` as first variable")
	assert.Equal(t, "GET", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/features/"+firstFeatureId+"/dependents", response.Response.Request.URL.Path, "path for request was incorrect")

}

func Test_can_update_a_feature_lifecycle(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	features, _, err := client.Feature.ListFeatures(ctx)
	require.NoError(t, err, "listing features must not error")

	firstFeatureId := features[0].Id
	origStatus := features[0].Status
	statusTo := "disable"
	if origStatus == "DISABLED" {
		statusTo = "enable"
	}
	statusChange, response, err := client.Feature.UpdateFeatureLifecycle(ctx, firstFeatureId, statusTo, nil)

	require.NoError(t, err, "updating feature status should not error")
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	require.IsType(t, &okta.Feature{}, statusChange, "did not return `[]*okta.Feature` as first variable")
	assert.Equal(t, "POST", response.Response.Request.Method, "did not make a get request")
	assert.Equal(t, "/api/v1/features/"+firstFeatureId+"/"+statusTo, response.Response.Request.URL.Path, "path for request was incorrect")
	assert.NotEqual(t, origStatus, statusChange.Status, "did not change the status")

	toOrigStatus := "enable"
	if origStatus == "DISABLED" {
		toOrigStatus = "disable"
	}
	updatedStatusChange, response, err := client.Feature.UpdateFeatureLifecycle(ctx, firstFeatureId, toOrigStatus, nil)
	require.NoError(t, err, "updating feature status should not error")
	assert.Equal(t, origStatus, updatedStatusChange.Status, "did not change the status back")
	assert.Equal(t, "/api/v1/features/"+firstFeatureId+"/"+toOrigStatus, response.Response.Request.URL.Path, "path for request was incorrect")
}

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
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_create_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("Create an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	assert_event_hook_model(t, eventHook)

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	_, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func Test_get_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("get an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	foundEventHook, response, err := client.EventHook.GetEventHook(ctx, eventHook.Id)

	require.NoError(t, err, "get an event hook should not error")
	tests.Assert_response(t, response, "GET", "/api/v1/eventHooks/"+eventHook.Id)

	assert_event_hook_model(t, foundEventHook)
	require.Equal(t, eventHook.Id, foundEventHook.Id, "did not find the same event hook from id")

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	_, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func Test_update_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("Create an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	eventHook.Name = "GO SDK: Updated Event Hook"
	updatedEventHook, response, err := client.EventHook.UpdateEventHook(ctx, eventHook.Id, *eventHook)

	assert_event_hook_model(t, updatedEventHook)
	require.Equal(t, eventHook.Name, updatedEventHook.Name, "update of event hook did not work")

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	_, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func Test_deactivate_and_delete_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("deactivate and delete an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	assert_event_hook_model(t, eventHook)
	require.Equal(t, "ACTIVE", eventHook.Status, "event hook was not active")

	_, response, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks/"+eventHook.Id+"/lifecycle/deactivate")

	foundEventHook, _, err := client.EventHook.GetEventHook(ctx, eventHook.Id)

	require.NoError(t, err, "get an event hook should not error")
	require.Equal(t, "INACTIVE", foundEventHook.Status, "event hook was not inactive after deactivate")

	response, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
	tests.Assert_response(t, response, "DELETE", "/api/v1/eventHooks/"+eventHook.Id)
}

func Test_activate_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("activate an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	assert_event_hook_model(t, eventHook)
	require.Equal(t, "ACTIVE", eventHook.Status, "event hook was not active")

	eventHook, response, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks/"+eventHook.Id+"/lifecycle/deactivate")
	require.Equal(t, "INACTIVE", eventHook.Status, "event hook was not active")

	eventHook, response, err = client.EventHook.ActivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "activating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks/"+eventHook.Id+"/lifecycle/activate")
	require.Equal(t, "ACTIVE", eventHook.Status, "event hook was not active")

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	response, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func Test_list_event_hooks(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("List event hooks")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	eventHookList, response, err := client.EventHook.ListEventHooks(ctx)
	assert.IsType(t, []*okta.EventHook{}, eventHookList, "did not return a list of eventHook objects")

	found := false
	for _, eh := range eventHookList {
		if eh.Id == eventHook.Id {
			found = true
		}
	}

	assert.True(t, found, "did not find the eventHook in the list")

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	_, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func Test_verify_an_event_hook(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	eventHookRequest := createEventHookRequestObject("Verify an event hook")

	eventHook, response, err := client.EventHook.CreateEventHook(ctx, eventHookRequest)

	require.NoError(t, err, "creating an event hook should not error")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks")

	_, response, err = client.EventHook.VerifyEventHook(ctx, eventHook.Id)

	// We expect this call to error.  Our test is just making sure we have the correct endpoint.
	// To fully test this, we have to have an acutal endpoint that will respond to event hooks.
	require.Error(t, err, "should have thrown an error because our event hook uri does not actually exist")
	require.Equal(t, 400, response.StatusCode, "Should have errored with a 400 status code")
	tests.Assert_response(t, response, "POST", "/api/v1/eventHooks/"+eventHook.Id+"/lifecycle/verify")

	_, _, err = client.EventHook.DeactivateEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deactivating an event hook should not error")
	_, err = client.EventHook.DeleteEventHook(ctx, eventHook.Id)
	require.NoError(t, err, "deleting an event hook should not error")
}

func createEventHookRequestObject(name string) okta.EventHook {
	eventHookEvent := okta.EventSubscriptions{
		Type: "EVENT_TYPE",
		Items: []string{
			"user.lifecycle.create",
			"user.lifecycle.activate",
		},
	}

	eventHookChannelConfigHeader := okta.EventHookChannelConfigHeader{
		Key:   "X-Other-Header",
		Value: "some-other-value",
	}

	eventHookChannelConfigAuthScheme := okta.EventHookChannelConfigAuthScheme{
		Type:  "HEADER",
		Key:   "Authorization",
		Value: "MyApiKey",
	}

	eventHookChannelConfig := okta.EventHookChannelConfig{
		Uri: "https://www.example.com/eventHooks",
		Headers: []*okta.EventHookChannelConfigHeader{
			&eventHookChannelConfigHeader,
		},
		AuthScheme: &eventHookChannelConfigAuthScheme,
	}

	eventHookChannel := okta.EventHookChannel{
		Type:    "HTTP",
		Version: "1.0.0",
		Config:  &eventHookChannelConfig,
	}

	eventHookRequest := okta.EventHook{
		Name:    "GO SDK: " + name,
		Events:  &eventHookEvent,
		Channel: &eventHookChannel,
	}

	return eventHookRequest
}

func assert_event_hook_model(t *testing.T, eventHook *okta.EventHook) {
	require.IsType(t, &okta.EventHook{}, eventHook, "did not return `*okta.EventHook` type as first variable")
	assert.NotEmpty(t, eventHook.Links, "links should not be empty")
	assert.NotEmpty(t, eventHook.Channel, "created should not be empty")
	assert.IsType(t, &okta.EventHookChannel{}, eventHook.Channel, "channel should be of type `*okta.EventHookChannel`")
	assert.NotEmpty(t, eventHook.Created, "created should not be empty")
	assert.IsType(t, &time.Time{}, eventHook.Created, "created should not be of type `*time.Time`")
	assert.NotEmpty(t, eventHook.CreatedBy, "createdBy should not be empty")
	assert.NotEmpty(t, eventHook.Events, "events should not be empty")
	assert.IsType(t, &okta.EventSubscriptions{}, eventHook.Events, "events should be of type `*okta.EventSubscriptions`")
	assert.NotEmpty(t, eventHook.Id, "id should not be empty")
	assert.NotEmpty(t, eventHook.LastUpdated, "lastUpdated should not be empty")
	assert.IsType(t, &time.Time{}, eventHook.LastUpdated, "lastUpdated should not be of type `*time.Time`")
	assert.NotEmpty(t, eventHook.Name, "name should not be empty")
	assert.NotEmpty(t, eventHook.Status, "status should not be empty")
	assert.NotEmpty(t, eventHook.VerificationStatus, "verificationStatus should not be empty")
}

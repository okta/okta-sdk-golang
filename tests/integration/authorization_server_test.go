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
	"net/http"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_can_create_an_authorizaiton_server(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")

	assert_authorization_server_model(t, authorizationServer)

	_, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
}

func Test_can_get_an_authorizaiton_server(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")

	authorizationServer, response, err = client.AuthorizationServer.GetAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err, "getting an authorizaiton server should not error")
	tests.AssertResponse(t, response, "GET", "/api/v1/authorizationServers/"+authorizationServer.Id)

	assert_authorization_server_model(t, authorizationServer)

	assert.Equal(t, as.Name, authorizationServer.Name, "did not return the same authorization server name")
	assert.Equal(t, as.Description, authorizationServer.Description, "did not return the same authorization server description")

	_, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
}

func Test_can_update_an_authorizaiton_server(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")

	authorizationServer, response, err = client.AuthorizationServer.GetAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err, "getting an authorizaiton server should not error")
	tests.AssertResponse(t, response, "GET", "/api/v1/authorizationServers/"+authorizationServer.Id)

	assert_authorization_server_model(t, authorizationServer)

	assert.Equal(t, as.Name, authorizationServer.Name, "did not return the same authorization server name")
	assert.Equal(t, as.Description, authorizationServer.Description, "did not return the same authorization server description")

	updatedName := "Updated Authorization Server - Golang"
	updatedDescription := "Updated Authorization Server Description"
	authorizationServer.Name = updatedName
	authorizationServer.Description = updatedDescription

	authorizationServer, response, err = client.AuthorizationServer.UpdateAuthorizationServer(ctx, authorizationServer.Id, *authorizationServer)
	require.NoError(t, err, "getting an authorizaiton server should not error")
	tests.AssertResponse(t, response, "PUT", "/api/v1/authorizationServers/"+authorizationServer.Id)

	assert_authorization_server_model(t, authorizationServer)

	assert.Equal(t, updatedName, authorizationServer.Name, "did not return the same authorization server name")
	assert.Equal(t, updatedDescription, authorizationServer.Description, "did not return the same authorization server description")

	_, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
}

func Test_can_delete_an_authorizaiton_server(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")

	authorizationServer, response, err = client.AuthorizationServer.GetAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err, "getting an authorizaiton server should not error")
	tests.AssertResponse(t, response, "GET", "/api/v1/authorizationServers/"+authorizationServer.Id)

	assert_authorization_server_model(t, authorizationServer)

	assert.Equal(t, as.Name, authorizationServer.Name, "did not return the same authorization server name")
	assert.Equal(t, as.Description, authorizationServer.Description, "did not return the same authorization server description")

	response, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode, "did not return a 204 status code during delete")

	_, response, err = client.AuthorizationServer.GetAuthorizationServer(ctx, authorizationServer.Id)
	assert.Error(t, err, "Finding an authorization server by id should have reported an error")
	assert.Equal(t, http.StatusNotFound, response.StatusCode, "Should have resulted in a 404 when finding a deleted authorization server")
}

func Test_can_list_authorizaiton_servers(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")

	authorizationServerList, response, err := client.AuthorizationServer.ListAuthorizationServers(ctx, nil)
	require.NoError(t, err, "list authorizaiton servers should not error")
	tests.AssertResponse(t, response, "GET", "/api/v1/authorizationServers")

	found := false
	for _, authServer := range authorizationServerList {
		if authServer.Id == authorizationServer.Id {
			assert_authorization_server_model(t, authServer)
			found = true
		}
	}
	assert.True(t, found, "Could not find authorization from list")

	_, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
}

func Test_can_activate_an_authorizaiton_server(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	as := okta.AuthorizationServer{
		Name:        testName("Sample Authorizaiton Server - Golang"),
		Description: "Sample Authorizaiton Server Description for Golang",
		Audiences:   []string{"api://default"},
	}

	authorizationServer, response, err := client.AuthorizationServer.CreateAuthorizationServer(ctx, as)
	require.NoError(t, err, "creating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers")
	assert.Equal(t, "ACTIVE", authorizationServer.Status, "should have active status after creating")

	response, err = client.AuthorizationServer.DeactivateAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err, "deactivating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers/"+authorizationServer.Id+"/lifecycle/deactivate")

	authorizationServer, _, _ = client.AuthorizationServer.GetAuthorizationServer(ctx, authorizationServer.Id)
	assert.Equal(t, "INACTIVE", authorizationServer.Status, "should have inactive status after deactivating")

	response, err = client.AuthorizationServer.ActivateAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err, "activating an authorizaiton server should not error")
	tests.AssertResponse(t, response, "POST", "/api/v1/authorizationServers/"+authorizationServer.Id+"/lifecycle/activate")

	// authorizationServer, response, _ = client.AuthorizationServer.GetAuthorizationServer(authorizationServer.Id)
	// assert.Equal(t, "ACTIVE", authorizationServer.Status, "should have active status after activating")

	_, err = client.AuthorizationServer.DeleteAuthorizationServer(ctx, authorizationServer.Id)
	require.NoError(t, err)
}

func assert_authorization_server_model(t *testing.T, authorizationServer *okta.AuthorizationServer) {
	require.IsType(t, &okta.AuthorizationServer{}, authorizationServer, "did not return `*okta.AuthorizationServer` type as first variable")
	assert.NotEmpty(t, authorizationServer.Id, "id should not be empty")
	assert.NotEmpty(t, authorizationServer.Audiences, "audiences should not be empty")
	assert.NotEmpty(t, authorizationServer.Created, "created should not be empty")
	assert.IsType(t, &time.Time{}, authorizationServer.Created, "created should not be of type `*time.Time`")
	assert.NotEmpty(t, authorizationServer.Credentials, "credentials should not be empty")
	assert.IsType(t, &okta.AuthorizationServerCredentials{}, authorizationServer.Credentials, "credentials should not be of type `*okta.AuthorizationServerCredentials`")
	assert.NotEmpty(t, authorizationServer.Description, "description should not be empty")
	assert.NotEmpty(t, authorizationServer.Issuer, "issuer should not be empty")
	assert.NotEmpty(t, authorizationServer.IssuerMode, "issuerMode should not be empty")
	assert.NotEmpty(t, authorizationServer.LastUpdated, "lastUpdated should not be empty")
	assert.IsType(t, &time.Time{}, authorizationServer.LastUpdated, "lastUpdated should not be of type `*time.Time`")
	assert.NotEmpty(t, authorizationServer.Name, "name should not be empty")
	assert.NotEmpty(t, authorizationServer.Status, "status should not be empty")
}

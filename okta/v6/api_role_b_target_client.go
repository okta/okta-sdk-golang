/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RoleBTargetClientAPI interface {

	/*
			AssignAppTargetInstanceRoleForClient Assign a client role app instance target

			Assigns an app instance target to an `APP_ADMIN` role assignment to a client. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
		The role no longer applies to all app targets, but applies only to the specified target.

		> **Note:** You can target a mixture of both OIN app and app instance targets, but you can't assign permissions to manage all instances of an OIN app and then assign a subset of permissions to the same app.
		For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage only specific configurations of the Salesforce app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clientId Client app ID
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiAssignAppTargetInstanceRoleForClientRequest
	*/
	AssignAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, appName string, appId string) ApiAssignAppTargetInstanceRoleForClientRequest

	// AssignAppTargetInstanceRoleForClientExecute executes the request
	AssignAppTargetInstanceRoleForClientExecute(r ApiAssignAppTargetInstanceRoleForClientRequest) (*APIResponse, error)

	/*
			AssignAppTargetRoleToClient Assign a client role app target

			Assigns an OIN app target for an `APP_ADMIN` role assignment to a client. When you assign an app target from the OIN catalog, you reduce the scope of the role assignment.
		The role assignment applies to only app instances that are included in the specified OIN app target.

		An assigned OIN app target overrides any existing app instance targets.
		For example, if a user is assigned to administer a specific Facebook instance, a successful request to add an OIN app target with `facebook` for `appName` makes that user the administrator for all Facebook instances.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clientId Client app ID
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@return ApiAssignAppTargetRoleToClientRequest
	*/
	AssignAppTargetRoleToClient(ctx context.Context, clientId string, roleAssignmentId string, appName string) ApiAssignAppTargetRoleToClientRequest

	// AssignAppTargetRoleToClientExecute executes the request
	AssignAppTargetRoleToClientExecute(r ApiAssignAppTargetRoleToClientRequest) (*APIResponse, error)

	/*
		AssignGroupTargetRoleForClient Assign a client role group target

		Assigns a group target to a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client app. When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets, but applies only to the specified target.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId Client app ID
		@param roleAssignmentId The `id` of the role assignment
		@param groupId The `id` of the group
		@return ApiAssignGroupTargetRoleForClientRequest
	*/
	AssignGroupTargetRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, groupId string) ApiAssignGroupTargetRoleForClientRequest

	// AssignGroupTargetRoleForClientExecute executes the request
	AssignGroupTargetRoleForClientExecute(r ApiAssignGroupTargetRoleForClientRequest) (*APIResponse, error)

	/*
		ListAppTargetRoleToClient List all client role app targets

		Lists all OIN app targets for an `APP_ADMIN` role that's assigned to a client (by `clientId`).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId Client app ID
		@param roleAssignmentId The `id` of the role assignment
		@return ApiListAppTargetRoleToClientRequest
	*/
	ListAppTargetRoleToClient(ctx context.Context, clientId string, roleAssignmentId string) ApiListAppTargetRoleToClientRequest

	// ListAppTargetRoleToClientExecute executes the request
	//  @return []CatalogApplication
	ListAppTargetRoleToClientExecute(r ApiListAppTargetRoleToClientRequest) ([]CatalogApplication, *APIResponse, error)

	/*
		ListGroupTargetRoleForClient List all client role group targets

		Lists all group targets for a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client. If the role isn't scoped to specific group targets, Okta returns an empty array `[]`.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clientId Client app ID
		@param roleAssignmentId The `id` of the role assignment
		@return ApiListGroupTargetRoleForClientRequest
	*/
	ListGroupTargetRoleForClient(ctx context.Context, clientId string, roleAssignmentId string) ApiListGroupTargetRoleForClientRequest

	// ListGroupTargetRoleForClientExecute executes the request
	//  @return []Group
	ListGroupTargetRoleForClientExecute(r ApiListGroupTargetRoleForClientRequest) ([]Group, *APIResponse, error)

	/*
			RemoveAppTargetInstanceRoleForClient Unassign a client role app instance target

			Unassigns an app instance target from a role assignment to a client app

		> **Note:** You can't remove the last app instance target from a role assignment.
		> If you need a role assignment that applies to all the apps, delete the role assignment with the instance target and create another one.  See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clientId Client app ID
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiRemoveAppTargetInstanceRoleForClientRequest
	*/
	RemoveAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, appName string, appId string) ApiRemoveAppTargetInstanceRoleForClientRequest

	// RemoveAppTargetInstanceRoleForClientExecute executes the request
	RemoveAppTargetInstanceRoleForClientExecute(r ApiRemoveAppTargetInstanceRoleForClientRequest) (*APIResponse, error)

	/*
			RemoveAppTargetRoleFromClient Unassign a client role app target

			Unassigns an OIN app target for a role assignment to a client app

		> **Note:** You can't remove the last OIN app target from a role assignment.
		> If you need a role assignment that applies to all apps, delete the role assignment with the target and create another one. See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clientId Client app ID
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@return ApiRemoveAppTargetRoleFromClientRequest
	*/
	RemoveAppTargetRoleFromClient(ctx context.Context, clientId string, roleAssignmentId string, appName string) ApiRemoveAppTargetRoleFromClientRequest

	// RemoveAppTargetRoleFromClientExecute executes the request
	RemoveAppTargetRoleFromClientExecute(r ApiRemoveAppTargetRoleFromClientRequest) (*APIResponse, error)

	/*
			RemoveGroupTargetRoleFromClient Unassign a client role group target

			Unassigns a Group target from a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client app.

		> **Note:** You can't remove the last group target from a role assignment. If you need a role assignment that applies to all groups, delete the role assignment with the target and create another one. See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param clientId Client app ID
			@param roleAssignmentId The `id` of the role assignment
			@param groupId The `id` of the group
			@return ApiRemoveGroupTargetRoleFromClientRequest
	*/
	RemoveGroupTargetRoleFromClient(ctx context.Context, clientId string, roleAssignmentId string, groupId string) ApiRemoveGroupTargetRoleFromClientRequest

	// RemoveGroupTargetRoleFromClientExecute executes the request
	RemoveGroupTargetRoleFromClientExecute(r ApiRemoveGroupTargetRoleFromClientRequest) (*APIResponse, error)
}

// RoleBTargetClientAPIService RoleBTargetClientAPI service
type RoleBTargetClientAPIService service

type ApiAssignAppTargetInstanceRoleForClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiAssignAppTargetInstanceRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetInstanceRoleForClientExecute(r)
}

/*
AssignAppTargetInstanceRoleForClient Assign a client role app instance target

Assigns an app instance target to an `APP_ADMIN` role assignment to a client. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
The role no longer applies to all app targets, but applies only to the specified target.

> **Note:** You can target a mixture of both OIN app and app instance targets, but you can't assign permissions to manage all instances of an OIN app and then assign a subset of permissions to the same app.
For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage only specific configurations of the Salesforce app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiAssignAppTargetInstanceRoleForClientRequest
*/
func (a *RoleBTargetClientAPIService) AssignAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, appName string, appId string) ApiAssignAppTargetInstanceRoleForClientRequest {
	return ApiAssignAppTargetInstanceRoleForClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) AssignAppTargetInstanceRoleForClientExecute(r ApiAssignAppTargetInstanceRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.AssignAppTargetInstanceRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiAssignAppTargetRoleToClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiAssignAppTargetRoleToClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetRoleToClientExecute(r)
}

/*
AssignAppTargetRoleToClient Assign a client role app target

Assigns an OIN app target for an `APP_ADMIN` role assignment to a client. When you assign an app target from the OIN catalog, you reduce the scope of the role assignment.
The role assignment applies to only app instances that are included in the specified OIN app target.

An assigned OIN app target overrides any existing app instance targets.
For example, if a user is assigned to administer a specific Facebook instance, a successful request to add an OIN app target with `facebook` for `appName` makes that user the administrator for all Facebook instances.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiAssignAppTargetRoleToClientRequest
*/
func (a *RoleBTargetClientAPIService) AssignAppTargetRoleToClient(ctx context.Context, clientId string, roleAssignmentId string, appName string) ApiAssignAppTargetRoleToClientRequest {
	return ApiAssignAppTargetRoleToClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) AssignAppTargetRoleToClientExecute(r ApiAssignAppTargetRoleToClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.AssignAppTargetRoleToClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiAssignGroupTargetRoleForClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	groupId          string
	retryCount       int32
}

func (r ApiAssignGroupTargetRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetRoleForClientExecute(r)
}

/*
AssignGroupTargetRoleForClient Assign a client role group target

Assigns a group target to a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client app. When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets, but applies only to the specified target.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param groupId The `id` of the group
	@return ApiAssignGroupTargetRoleForClientRequest
*/
func (a *RoleBTargetClientAPIService) AssignGroupTargetRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, groupId string) ApiAssignGroupTargetRoleForClientRequest {
	return ApiAssignGroupTargetRoleForClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		groupId:          groupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) AssignGroupTargetRoleForClientExecute(r ApiAssignGroupTargetRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.AssignGroupTargetRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiListAppTargetRoleToClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListAppTargetRoleToClientRequest) After(after string) ApiListAppTargetRoleToClientRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListAppTargetRoleToClientRequest) Limit(limit int32) ApiListAppTargetRoleToClientRequest {
	r.limit = &limit
	return r
}

func (r ApiListAppTargetRoleToClientRequest) Execute() ([]CatalogApplication, *APIResponse, error) {
	return r.ApiService.ListAppTargetRoleToClientExecute(r)
}

/*
ListAppTargetRoleToClient List all client role app targets

Lists all OIN app targets for an `APP_ADMIN` role that's assigned to a client (by `clientId`).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListAppTargetRoleToClientRequest
*/
func (a *RoleBTargetClientAPIService) ListAppTargetRoleToClient(ctx context.Context, clientId string, roleAssignmentId string) ApiListAppTargetRoleToClientRequest {
	return ApiListAppTargetRoleToClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []CatalogApplication
func (a *RoleBTargetClientAPIService) ListAppTargetRoleToClientExecute(r ApiListAppTargetRoleToClientRequest) ([]CatalogApplication, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CatalogApplication
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.ListAppTargetRoleToClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiListGroupTargetRoleForClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListGroupTargetRoleForClientRequest) After(after string) ApiListGroupTargetRoleForClientRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListGroupTargetRoleForClientRequest) Limit(limit int32) ApiListGroupTargetRoleForClientRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetRoleForClientRequest) Execute() ([]Group, *APIResponse, error) {
	return r.ApiService.ListGroupTargetRoleForClientExecute(r)
}

/*
ListGroupTargetRoleForClient List all client role group targets

Lists all group targets for a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client. If the role isn't scoped to specific group targets, Okta returns an empty array `[]`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListGroupTargetRoleForClientRequest
*/
func (a *RoleBTargetClientAPIService) ListGroupTargetRoleForClient(ctx context.Context, clientId string, roleAssignmentId string) ApiListGroupTargetRoleForClientRequest {
	return ApiListGroupTargetRoleForClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []Group
func (a *RoleBTargetClientAPIService) ListGroupTargetRoleForClientExecute(r ApiListGroupTargetRoleForClientRequest) ([]Group, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Group
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.ListGroupTargetRoleForClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiRemoveAppTargetInstanceRoleForClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiRemoveAppTargetInstanceRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveAppTargetInstanceRoleForClientExecute(r)
}

/*
RemoveAppTargetInstanceRoleForClient Unassign a client role app instance target

# Unassigns an app instance target from a role assignment to a client app

> **Note:** You can't remove the last app instance target from a role assignment.
> If you need a role assignment that applies to all the apps, delete the role assignment with the instance target and create another one.  See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiRemoveAppTargetInstanceRoleForClientRequest
*/
func (a *RoleBTargetClientAPIService) RemoveAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleAssignmentId string, appName string, appId string) ApiRemoveAppTargetInstanceRoleForClientRequest {
	return ApiRemoveAppTargetInstanceRoleForClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) RemoveAppTargetInstanceRoleForClientExecute(r ApiRemoveAppTargetInstanceRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.RemoveAppTargetInstanceRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiRemoveAppTargetRoleFromClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiRemoveAppTargetRoleFromClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveAppTargetRoleFromClientExecute(r)
}

/*
RemoveAppTargetRoleFromClient Unassign a client role app target

# Unassigns an OIN app target for a role assignment to a client app

> **Note:** You can't remove the last OIN app target from a role assignment.
> If you need a role assignment that applies to all apps, delete the role assignment with the target and create another one. See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiRemoveAppTargetRoleFromClientRequest
*/
func (a *RoleBTargetClientAPIService) RemoveAppTargetRoleFromClient(ctx context.Context, clientId string, roleAssignmentId string, appName string) ApiRemoveAppTargetRoleFromClientRequest {
	return ApiRemoveAppTargetRoleFromClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) RemoveAppTargetRoleFromClientExecute(r ApiRemoveAppTargetRoleFromClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.RemoveAppTargetRoleFromClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiRemoveGroupTargetRoleFromClientRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetClientAPI
	clientId         string
	roleAssignmentId string
	groupId          string
	retryCount       int32
}

func (r ApiRemoveGroupTargetRoleFromClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveGroupTargetRoleFromClientExecute(r)
}

/*
RemoveGroupTargetRoleFromClient Unassign a client role group target

Unassigns a Group target from a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a client app.

> **Note:** You can't remove the last group target from a role assignment. If you need a role assignment that applies to all groups, delete the role assignment with the target and create another one. See [Unassign a client role](/openapi/okta-management/management/tag/RoleAssignmentClient/#tag/RoleAssignmentClient/operation/deleteRoleFromClient).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId Client app ID
	@param roleAssignmentId The `id` of the role assignment
	@param groupId The `id` of the group
	@return ApiRemoveGroupTargetRoleFromClientRequest
*/
func (a *RoleBTargetClientAPIService) RemoveGroupTargetRoleFromClient(ctx context.Context, clientId string, roleAssignmentId string, groupId string) ApiRemoveGroupTargetRoleFromClientRequest {
	return ApiRemoveGroupTargetRoleFromClientRequest{
		ApiService:       a,
		ctx:              ctx,
		clientId:         clientId,
		roleAssignmentId: roleAssignmentId,
		groupId:          groupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetClientAPIService) RemoveGroupTargetRoleFromClientExecute(r ApiRemoveGroupTargetRoleFromClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetClientAPIService.RemoveGroupTargetRoleFromClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleAssignmentId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

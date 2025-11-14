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

API version: 2025.08.0
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

type RoleAssignmentAUserAPI interface {

	/*
			AssignRoleToUser Assign a user role

			Assigns a [standard role](/openapi/okta-management/guides/roles/#standard-roles) to a user.

		You can also assign a custom role to a user, but the preferred method to assign a custom role to a user is to create a binding between the custom role, the resource set, and the user. See [Create a role resource set binding](/openapi/okta-management/management/tag/RoleDResourceSetBinding/#tag/RoleDResourceSetBinding/operation/createResourceSetBinding).

		> **Notes:**
		> * The request payload is different for standard and custom role assignments.
		> * For IAM-based standard role assignments, use the request payload for standard roles. However, the response payload for IAM-based role assignments is similar to the custom role's assignment response.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@return ApiAssignRoleToUserRequest
	*/
	AssignRoleToUser(ctx context.Context, userId string) ApiAssignRoleToUserRequest

	// AssignRoleToUserExecute executes the request
	//  @return AssignRoleToUser201Response
	AssignRoleToUserExecute(r ApiAssignRoleToUserRequest) (*AssignRoleToUser201Response, *APIResponse, error)

	/*
		GetRoleAssignmentGovernanceGrant Retrieve a user role governance source

		Retrieves a governance source (identified by `grantId`) for a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@param grantId Grant ID
		@return ApiGetRoleAssignmentGovernanceGrantRequest
	*/
	GetRoleAssignmentGovernanceGrant(ctx context.Context, userId string, roleAssignmentId string, grantId string) ApiGetRoleAssignmentGovernanceGrantRequest

	// GetRoleAssignmentGovernanceGrantExecute executes the request
	//  @return RoleGovernanceSource
	GetRoleAssignmentGovernanceGrantExecute(r ApiGetRoleAssignmentGovernanceGrantRequest) (*RoleGovernanceSource, *APIResponse, error)

	/*
		GetRoleAssignmentGovernanceGrantResources Retrieve the user role governance source resources

		Retrieves the resources of a governance source (identified by `grantId`) for a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@param grantId Grant ID
		@return ApiGetRoleAssignmentGovernanceGrantResourcesRequest
	*/
	GetRoleAssignmentGovernanceGrantResources(ctx context.Context, userId string, roleAssignmentId string, grantId string) ApiGetRoleAssignmentGovernanceGrantResourcesRequest

	// GetRoleAssignmentGovernanceGrantResourcesExecute executes the request
	//  @return RoleGovernanceResources
	GetRoleAssignmentGovernanceGrantResourcesExecute(r ApiGetRoleAssignmentGovernanceGrantResourcesRequest) (*RoleGovernanceResources, *APIResponse, error)

	/*
		GetUserAssignedRole Retrieve a user role assignment

		Retrieves a role assigned to a user (identified by `userId`). The `roleAssignmentId` parameter is the unique identifier for either a standard role assignment object or a custom role resource set binding object.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@return ApiGetUserAssignedRoleRequest
	*/
	GetUserAssignedRole(ctx context.Context, userId string, roleAssignmentId string) ApiGetUserAssignedRoleRequest

	// GetUserAssignedRoleExecute executes the request
	//  @return AssignRoleToGroup200Response
	GetUserAssignedRoleExecute(r ApiGetUserAssignedRoleRequest) (*AssignRoleToGroup200Response, *APIResponse, error)

	/*
		GetUserAssignedRoleGovernance Retrieve all user role governance sources

		Retrieves the governance sources of a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@return ApiGetUserAssignedRoleGovernanceRequest
	*/
	GetUserAssignedRoleGovernance(ctx context.Context, userId string, roleAssignmentId string) ApiGetUserAssignedRoleGovernanceRequest

	// GetUserAssignedRoleGovernanceExecute executes the request
	//  @return RoleGovernance
	GetUserAssignedRoleGovernanceExecute(r ApiGetUserAssignedRoleGovernanceRequest) (*RoleGovernance, *APIResponse, error)

	/*
		ListAssignedRolesForUser List all user role assignments

		Lists all roles assigned to a user (identified by `userId`)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@return ApiListAssignedRolesForUserRequest
	*/
	ListAssignedRolesForUser(ctx context.Context, userId string) ApiListAssignedRolesForUserRequest

	// ListAssignedRolesForUserExecute executes the request
	//  @return []AssignRoleToGroup200Response
	ListAssignedRolesForUserExecute(r ApiListAssignedRolesForUserRequest) ([]AssignRoleToGroup200Response, *APIResponse, error)

	/*
		ListUsersWithRoleAssignments List all users with role assignments

		Lists all users with role assignments

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListUsersWithRoleAssignmentsRequest
	*/
	ListUsersWithRoleAssignments(ctx context.Context) ApiListUsersWithRoleAssignmentsRequest

	// ListUsersWithRoleAssignmentsExecute executes the request
	//  @return RoleAssignedUsers
	ListUsersWithRoleAssignmentsExecute(r ApiListUsersWithRoleAssignmentsRequest) (*RoleAssignedUsers, *APIResponse, error)

	/*
		UnassignRoleFromUser Unassign a user role

		Unassigns a role assignment (identified by `roleAssignmentId`) from a user (identified by `userId`)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@return ApiUnassignRoleFromUserRequest
	*/
	UnassignRoleFromUser(ctx context.Context, userId string, roleAssignmentId string) ApiUnassignRoleFromUserRequest

	// UnassignRoleFromUserExecute executes the request
	UnassignRoleFromUserExecute(r ApiUnassignRoleFromUserRequest) (*APIResponse, error)
}

// RoleAssignmentAUserAPIService RoleAssignmentAUserAPI service
type RoleAssignmentAUserAPIService service

type ApiAssignRoleToUserRequest struct {
	ctx                  context.Context
	ApiService           RoleAssignmentAUserAPI
	userId               string
	assignRoleRequest    *AssignRoleToUserRequest
	disableNotifications *bool
	retryCount           int32
}

func (r ApiAssignRoleToUserRequest) AssignRoleRequest(assignRoleRequest AssignRoleToUserRequest) ApiAssignRoleToUserRequest {
	r.assignRoleRequest = &assignRoleRequest
	return r
}

// Setting this to &#x60;true&#x60; grants the user third-party admin status
func (r ApiAssignRoleToUserRequest) DisableNotifications(disableNotifications bool) ApiAssignRoleToUserRequest {
	r.disableNotifications = &disableNotifications
	return r
}

func (r ApiAssignRoleToUserRequest) Execute() (*AssignRoleToUser201Response, *APIResponse, error) {
	return r.ApiService.AssignRoleToUserExecute(r)
}

/*
AssignRoleToUser Assign a user role

Assigns a [standard role](/openapi/okta-management/guides/roles/#standard-roles) to a user.

You can also assign a custom role to a user, but the preferred method to assign a custom role to a user is to create a binding between the custom role, the resource set, and the user. See [Create a role resource set binding](/openapi/okta-management/management/tag/RoleDResourceSetBinding/#tag/RoleDResourceSetBinding/operation/createResourceSetBinding).

> **Notes:**
> * The request payload is different for standard and custom role assignments.
> * For IAM-based standard role assignments, use the request payload for standard roles. However, the response payload for IAM-based role assignments is similar to the custom role's assignment response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@return ApiAssignRoleToUserRequest
*/
func (a *RoleAssignmentAUserAPIService) AssignRoleToUser(ctx context.Context, userId string) ApiAssignRoleToUserRequest {
	return ApiAssignRoleToUserRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return AssignRoleToUser201Response
func (a *RoleAssignmentAUserAPIService) AssignRoleToUserExecute(r ApiAssignRoleToUserRequest) (*AssignRoleToUser201Response, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssignRoleToUser201Response
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.AssignRoleToUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assignRoleRequest == nil {
		return localVarReturnValue, nil, reportError("assignRoleRequest is required and must be specified")
	}

	if r.disableNotifications != nil {
		localVarQueryParams.Add("disableNotifications", parameterToString(*r.disableNotifications, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.assignRoleRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiGetRoleAssignmentGovernanceGrantRequest struct {
	ctx              context.Context
	ApiService       RoleAssignmentAUserAPI
	userId           string
	roleAssignmentId string
	grantId          string
	retryCount       int32
}

func (r ApiGetRoleAssignmentGovernanceGrantRequest) Execute() (*RoleGovernanceSource, *APIResponse, error) {
	return r.ApiService.GetRoleAssignmentGovernanceGrantExecute(r)
}

/*
GetRoleAssignmentGovernanceGrant Retrieve a user role governance source

Retrieves a governance source (identified by `grantId`) for a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param grantId Grant ID
	@return ApiGetRoleAssignmentGovernanceGrantRequest
*/
func (a *RoleAssignmentAUserAPIService) GetRoleAssignmentGovernanceGrant(ctx context.Context, userId string, roleAssignmentId string, grantId string) ApiGetRoleAssignmentGovernanceGrantRequest {
	return ApiGetRoleAssignmentGovernanceGrantRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		grantId:          grantId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return RoleGovernanceSource
func (a *RoleAssignmentAUserAPIService) GetRoleAssignmentGovernanceGrantExecute(r ApiGetRoleAssignmentGovernanceGrantRequest) (*RoleGovernanceSource, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleGovernanceSource
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.GetRoleAssignmentGovernanceGrant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/governance/{grantId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"grantId"+"}", url.PathEscape(parameterToString(r.grantId, "")), -1)

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

type ApiGetRoleAssignmentGovernanceGrantResourcesRequest struct {
	ctx              context.Context
	ApiService       RoleAssignmentAUserAPI
	userId           string
	roleAssignmentId string
	grantId          string
	retryCount       int32
}

func (r ApiGetRoleAssignmentGovernanceGrantResourcesRequest) Execute() (*RoleGovernanceResources, *APIResponse, error) {
	return r.ApiService.GetRoleAssignmentGovernanceGrantResourcesExecute(r)
}

/*
GetRoleAssignmentGovernanceGrantResources Retrieve the user role governance source resources

Retrieves the resources of a governance source (identified by `grantId`) for a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param grantId Grant ID
	@return ApiGetRoleAssignmentGovernanceGrantResourcesRequest
*/
func (a *RoleAssignmentAUserAPIService) GetRoleAssignmentGovernanceGrantResources(ctx context.Context, userId string, roleAssignmentId string, grantId string) ApiGetRoleAssignmentGovernanceGrantResourcesRequest {
	return ApiGetRoleAssignmentGovernanceGrantResourcesRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		grantId:          grantId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return RoleGovernanceResources
func (a *RoleAssignmentAUserAPIService) GetRoleAssignmentGovernanceGrantResourcesExecute(r ApiGetRoleAssignmentGovernanceGrantResourcesRequest) (*RoleGovernanceResources, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleGovernanceResources
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.GetRoleAssignmentGovernanceGrantResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/governance/{grantId}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"grantId"+"}", url.PathEscape(parameterToString(r.grantId, "")), -1)

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

type ApiGetUserAssignedRoleRequest struct {
	ctx              context.Context
	ApiService       RoleAssignmentAUserAPI
	userId           string
	roleAssignmentId string
	retryCount       int32
}

func (r ApiGetUserAssignedRoleRequest) Execute() (*AssignRoleToGroup200Response, *APIResponse, error) {
	return r.ApiService.GetUserAssignedRoleExecute(r)
}

/*
GetUserAssignedRole Retrieve a user role assignment

Retrieves a role assigned to a user (identified by `userId`). The `roleAssignmentId` parameter is the unique identifier for either a standard role assignment object or a custom role resource set binding object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiGetUserAssignedRoleRequest
*/
func (a *RoleAssignmentAUserAPIService) GetUserAssignedRole(ctx context.Context, userId string, roleAssignmentId string) ApiGetUserAssignedRoleRequest {
	return ApiGetUserAssignedRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return AssignRoleToGroup200Response
func (a *RoleAssignmentAUserAPIService) GetUserAssignedRoleExecute(r ApiGetUserAssignedRoleRequest) (*AssignRoleToGroup200Response, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssignRoleToGroup200Response
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.GetUserAssignedRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)

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

type ApiGetUserAssignedRoleGovernanceRequest struct {
	ctx              context.Context
	ApiService       RoleAssignmentAUserAPI
	userId           string
	roleAssignmentId string
	retryCount       int32
}

func (r ApiGetUserAssignedRoleGovernanceRequest) Execute() (*RoleGovernance, *APIResponse, error) {
	return r.ApiService.GetUserAssignedRoleGovernanceExecute(r)
}

/*
GetUserAssignedRoleGovernance Retrieve all user role governance sources

Retrieves the governance sources of a role (identified by `roleAssignmentId`) that's assigned to a user (identified by `userId`)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiGetUserAssignedRoleGovernanceRequest
*/
func (a *RoleAssignmentAUserAPIService) GetUserAssignedRoleGovernance(ctx context.Context, userId string, roleAssignmentId string) ApiGetUserAssignedRoleGovernanceRequest {
	return ApiGetUserAssignedRoleGovernanceRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return RoleGovernance
func (a *RoleAssignmentAUserAPIService) GetUserAssignedRoleGovernanceExecute(r ApiGetUserAssignedRoleGovernanceRequest) (*RoleGovernance, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleGovernance
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.GetUserAssignedRoleGovernance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/governance"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)

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

type ApiListAssignedRolesForUserRequest struct {
	ctx        context.Context
	ApiService RoleAssignmentAUserAPI
	userId     string
	expand     *string
	retryCount int32
}

// An optional parameter used to return targets configured for the standard role assignment in the &#x60;embedded&#x60; property. Supported values: &#x60;targets/groups&#x60; or &#x60;targets/catalog/apps&#x60;
func (r ApiListAssignedRolesForUserRequest) Expand(expand string) ApiListAssignedRolesForUserRequest {
	r.expand = &expand
	return r
}

func (r ApiListAssignedRolesForUserRequest) Execute() ([]AssignRoleToGroup200Response, *APIResponse, error) {
	return r.ApiService.ListAssignedRolesForUserExecute(r)
}

/*
ListAssignedRolesForUser List all user role assignments

Lists all roles assigned to a user (identified by `userId`)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@return ApiListAssignedRolesForUserRequest
*/
func (a *RoleAssignmentAUserAPIService) ListAssignedRolesForUser(ctx context.Context, userId string) ApiListAssignedRolesForUserRequest {
	return ApiListAssignedRolesForUserRequest{
		ApiService: a,
		ctx:        ctx,
		userId:     userId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []AssignRoleToGroup200Response
func (a *RoleAssignmentAUserAPIService) ListAssignedRolesForUserExecute(r ApiListAssignedRolesForUserRequest) ([]AssignRoleToGroup200Response, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AssignRoleToGroup200Response
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.ListAssignedRolesForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
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

type ApiListUsersWithRoleAssignmentsRequest struct {
	ctx        context.Context
	ApiService RoleAssignmentAUserAPI
	after      *string
	limit      *int32
	retryCount int32
}

// Specifies the pagination cursor for the next page of targets
func (r ApiListUsersWithRoleAssignmentsRequest) After(after string) ApiListUsersWithRoleAssignmentsRequest {
	r.after = &after
	return r
}

// Specifies the number of results returned. Defaults to &#x60;100&#x60;.
func (r ApiListUsersWithRoleAssignmentsRequest) Limit(limit int32) ApiListUsersWithRoleAssignmentsRequest {
	r.limit = &limit
	return r
}

func (r ApiListUsersWithRoleAssignmentsRequest) Execute() (*RoleAssignedUsers, *APIResponse, error) {
	return r.ApiService.ListUsersWithRoleAssignmentsExecute(r)
}

/*
ListUsersWithRoleAssignments List all users with role assignments

Lists all users with role assignments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListUsersWithRoleAssignmentsRequest
*/
func (a *RoleAssignmentAUserAPIService) ListUsersWithRoleAssignments(ctx context.Context) ApiListUsersWithRoleAssignmentsRequest {
	return ApiListUsersWithRoleAssignmentsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return RoleAssignedUsers
func (a *RoleAssignmentAUserAPIService) ListUsersWithRoleAssignmentsExecute(r ApiListUsersWithRoleAssignmentsRequest) (*RoleAssignedUsers, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RoleAssignedUsers
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.ListUsersWithRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/assignees/users"

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

type ApiUnassignRoleFromUserRequest struct {
	ctx              context.Context
	ApiService       RoleAssignmentAUserAPI
	userId           string
	roleAssignmentId string
	retryCount       int32
}

func (r ApiUnassignRoleFromUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignRoleFromUserExecute(r)
}

/*
UnassignRoleFromUser Unassign a user role

Unassigns a role assignment (identified by `roleAssignmentId`) from a user (identified by `userId`)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiUnassignRoleFromUserRequest
*/
func (a *RoleAssignmentAUserAPIService) UnassignRoleFromUser(ctx context.Context, userId string, roleAssignmentId string) ApiUnassignRoleFromUserRequest {
	return ApiUnassignRoleFromUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleAssignmentAUserAPIService) UnassignRoleFromUserExecute(r ApiUnassignRoleFromUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleAssignmentAUserAPIService.UnassignRoleFromUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)

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

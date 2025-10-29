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

API version: 2025.10.0
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

type RoleBTargetAdminAPI interface {

	/*
		AssignAllAppsAsTargetToRoleForUser Assign all apps as target to admin role

		Assigns all apps as target to an `APP_ADMIN` role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@return ApiAssignAllAppsAsTargetToRoleForUserRequest
	*/
	AssignAllAppsAsTargetToRoleForUser(ctx context.Context, userId string, roleAssignmentId string) ApiAssignAllAppsAsTargetToRoleForUserRequest

	// AssignAllAppsAsTargetToRoleForUserExecute executes the request
	AssignAllAppsAsTargetToRoleForUserExecute(r ApiAssignAllAppsAsTargetToRoleForUserRequest) (*APIResponse, error)

	/*
			AssignAppInstanceTargetToAppAdminRoleForUser Assign an admin role app instance target

			Assigns an app instance target to an `APP_ADMIN` role assignment to an admin user. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
		The role no longer applies to all app targets, but applies only to the specified target.

		> **Note:** You can target a mixture of both OIN app and app instance targets, but can't assign permissions to manage all instances of an OIN app and then assign a subset of permission to the same OIN app.
		> For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage specific configurations of the Salesforce app.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest
	*/
	AssignAppInstanceTargetToAppAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest

	// AssignAppInstanceTargetToAppAdminRoleForUserExecute executes the request
	AssignAppInstanceTargetToAppAdminRoleForUserExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) (*APIResponse, error)

	/*
			AssignAppTargetToAdminRoleForUser Assign an admin role app target

			Assigns an OIN app target for an `APP_ADMIN` role assignment to an admin user. When you assign the first app target, you reduce the scope of the role assignment.
		The role no longer applies to all app targets, but applies only to the specified target.

		Assigning an OIN app target overrides any existing app instance targets of the OIN app.
		For example, if a user was assigned to administer a specific Facebook instance, a successful request to add an OIN app target with `facebook` for `appName` makes that user the admin for all Facebook instances.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@return ApiAssignAppTargetToAdminRoleForUserRequest
	*/
	AssignAppTargetToAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string) ApiAssignAppTargetToAdminRoleForUserRequest

	// AssignAppTargetToAdminRoleForUserExecute executes the request
	AssignAppTargetToAdminRoleForUserExecute(r ApiAssignAppTargetToAdminRoleForUserRequest) (*APIResponse, error)

	/*
			AssignGroupTargetToUserRole Assign an admin role group target

			Assigns a group target for a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.
		When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets but applies only to the specified target.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param groupId The `id` of the group
			@return ApiAssignGroupTargetToUserRoleRequest
	*/
	AssignGroupTargetToUserRole(ctx context.Context, userId string, roleAssignmentId string, groupId string) ApiAssignGroupTargetToUserRoleRequest

	// AssignGroupTargetToUserRoleExecute executes the request
	AssignGroupTargetToUserRoleExecute(r ApiAssignGroupTargetToUserRoleRequest) (*APIResponse, error)

	/*
			GetRoleTargetsByUserIdAndRoleId Retrieve a role target by assignment type

			Retrieves all role targets for an `APP_ADMIN`, `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user by user or group assignment type.
		If the role isn't scoped to specific group targets or any app targets, an empty array `[]` is returned.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleIdOrEncodedRoleId The `id` of the role or Base32 encoded `id` of the role name
			@return ApiGetRoleTargetsByUserIdAndRoleIdRequest
	*/
	GetRoleTargetsByUserIdAndRoleId(ctx context.Context, userId string, roleIdOrEncodedRoleId string) ApiGetRoleTargetsByUserIdAndRoleIdRequest

	// GetRoleTargetsByUserIdAndRoleIdExecute executes the request
	//  @return []RoleTarget
	GetRoleTargetsByUserIdAndRoleIdExecute(r ApiGetRoleTargetsByUserIdAndRoleIdRequest) ([]RoleTarget, *APIResponse, error)

	/*
		ListApplicationTargetsForApplicationAdministratorRoleForUser List all admin role app targets

		Lists all app targets for an `APP_ADMIN` role assigned to a user. The response is a list that includes OIN-cataloged apps or app instances. The response payload for an app instance contains the `id` property, but an OIN-cataloged app payload doesn't.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userId ID of an existing Okta user
		@param roleAssignmentId The `id` of the role assignment
		@return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest
	*/
	ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx context.Context, userId string, roleAssignmentId string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest

	// ListApplicationTargetsForApplicationAdministratorRoleForUserExecute executes the request
	//  @return []CatalogApplication
	ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) ([]CatalogApplication, *APIResponse, error)

	/*
			ListGroupTargetsForRole List all admin role group targets

			Lists all group targets for a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.
		If the role isn't scoped to specific group targets, an empty array `[]` is returned.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@return ApiListGroupTargetsForRoleRequest
	*/
	ListGroupTargetsForRole(ctx context.Context, userId string, roleAssignmentId string) ApiListGroupTargetsForRoleRequest

	// ListGroupTargetsForRoleExecute executes the request
	//  @return []Group
	ListGroupTargetsForRoleExecute(r ApiListGroupTargetsForRoleRequest) ([]Group, *APIResponse, error)

	/*
			UnassignAppInstanceTargetFromAdminRoleForUser Unassign an admin role app instance target

			Unassigns an app instance target from an `APP_ADMIN` role assignment to an admin user.

		> **Note:** You can't remove the last app instance target from a role assignment since this causes an exception.
		> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment and recreate a new one.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest
	*/
	UnassignAppInstanceTargetFromAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string, appId string) ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest

	// UnassignAppInstanceTargetFromAdminRoleForUserExecute executes the request
	UnassignAppInstanceTargetFromAdminRoleForUserExecute(r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) (*APIResponse, error)

	/*
			UnassignAppTargetFromAppAdminRoleForUser Unassign an admin role app target

			Unassigns an OIN app target from an `APP_ADMIN` role assignment to an admin user.

		> **Note:** You can't remove the last OIN app target from a role assignment since this causes an exception.
		> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment to the user and recreate a new one.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@return ApiUnassignAppTargetFromAppAdminRoleForUserRequest
	*/
	UnassignAppTargetFromAppAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string) ApiUnassignAppTargetFromAppAdminRoleForUserRequest

	// UnassignAppTargetFromAppAdminRoleForUserExecute executes the request
	UnassignAppTargetFromAppAdminRoleForUserExecute(r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) (*APIResponse, error)

	/*
			UnassignGroupTargetFromUserAdminRole Unassign an admin role group target

			Unassigns a group target from a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.

		> **Note:** You can't remove the last group target from a role assignment since this causes an exception.
		> If you need a role assignment that applies to all groups, delete the role assignment to the user and recreate a new one.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param userId ID of an existing Okta user
			@param roleAssignmentId The `id` of the role assignment
			@param groupId The `id` of the group
			@return ApiUnassignGroupTargetFromUserAdminRoleRequest
	*/
	UnassignGroupTargetFromUserAdminRole(ctx context.Context, userId string, roleAssignmentId string, groupId string) ApiUnassignGroupTargetFromUserAdminRoleRequest

	// UnassignGroupTargetFromUserAdminRoleExecute executes the request
	UnassignGroupTargetFromUserAdminRoleExecute(r ApiUnassignGroupTargetFromUserAdminRoleRequest) (*APIResponse, error)
}

// RoleBTargetAdminAPIService RoleBTargetAdminAPI service
type RoleBTargetAdminAPIService service

type ApiAssignAllAppsAsTargetToRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	retryCount       int32
}

func (r ApiAssignAllAppsAsTargetToRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAllAppsAsTargetToRoleForUserExecute(r)
}

/*
AssignAllAppsAsTargetToRoleForUser Assign all apps as target to admin role

Assigns all apps as target to an `APP_ADMIN` role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiAssignAllAppsAsTargetToRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) AssignAllAppsAsTargetToRoleForUser(ctx context.Context, userId string, roleAssignmentId string) ApiAssignAllAppsAsTargetToRoleForUserRequest {
	return ApiAssignAllAppsAsTargetToRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) AssignAllAppsAsTargetToRoleForUserExecute(r ApiAssignAllAppsAsTargetToRoleForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.AssignAllAppsAsTargetToRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps"
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

type ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppInstanceTargetToAppAdminRoleForUserExecute(r)
}

/*
AssignAppInstanceTargetToAppAdminRoleForUser Assign an admin role app instance target

Assigns an app instance target to an `APP_ADMIN` role assignment to an admin user. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
The role no longer applies to all app targets, but applies only to the specified target.

> **Note:** You can target a mixture of both OIN app and app instance targets, but can't assign permissions to manage all instances of an OIN app and then assign a subset of permission to the same OIN app.
> For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage specific configurations of the Salesforce app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) AssignAppInstanceTargetToAppAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest {
	return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) AssignAppInstanceTargetToAppAdminRoleForUserExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.AssignAppInstanceTargetToAppAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiAssignAppTargetToAdminRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiAssignAppTargetToAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetToAdminRoleForUserExecute(r)
}

/*
AssignAppTargetToAdminRoleForUser Assign an admin role app target

Assigns an OIN app target for an `APP_ADMIN` role assignment to an admin user. When you assign the first app target, you reduce the scope of the role assignment.
The role no longer applies to all app targets, but applies only to the specified target.

Assigning an OIN app target overrides any existing app instance targets of the OIN app.
For example, if a user was assigned to administer a specific Facebook instance, a successful request to add an OIN app target with `facebook` for `appName` makes that user the admin for all Facebook instances.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiAssignAppTargetToAdminRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) AssignAppTargetToAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string) ApiAssignAppTargetToAdminRoleForUserRequest {
	return ApiAssignAppTargetToAdminRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) AssignAppTargetToAdminRoleForUserExecute(r ApiAssignAppTargetToAdminRoleForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.AssignAppTargetToAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiAssignGroupTargetToUserRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	groupId          string
	retryCount       int32
}

func (r ApiAssignGroupTargetToUserRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetToUserRoleExecute(r)
}

/*
AssignGroupTargetToUserRole Assign an admin role group target

Assigns a group target for a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.
When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets but applies only to the specified target.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param groupId The `id` of the group
	@return ApiAssignGroupTargetToUserRoleRequest
*/
func (a *RoleBTargetAdminAPIService) AssignGroupTargetToUserRole(ctx context.Context, userId string, roleAssignmentId string, groupId string) ApiAssignGroupTargetToUserRoleRequest {
	return ApiAssignGroupTargetToUserRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		groupId:          groupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) AssignGroupTargetToUserRoleExecute(r ApiAssignGroupTargetToUserRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.AssignGroupTargetToUserRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiGetRoleTargetsByUserIdAndRoleIdRequest struct {
	ctx                   context.Context
	ApiService            RoleBTargetAdminAPI
	userId                string
	roleIdOrEncodedRoleId string
	assignmentType        *string
	after                 *string
	limit                 *int32
	retryCount            int32
}

// Specifies the assignment type of the user
func (r ApiGetRoleTargetsByUserIdAndRoleIdRequest) AssignmentType(assignmentType string) ApiGetRoleTargetsByUserIdAndRoleIdRequest {
	r.assignmentType = &assignmentType
	return r
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiGetRoleTargetsByUserIdAndRoleIdRequest) After(after string) ApiGetRoleTargetsByUserIdAndRoleIdRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiGetRoleTargetsByUserIdAndRoleIdRequest) Limit(limit int32) ApiGetRoleTargetsByUserIdAndRoleIdRequest {
	r.limit = &limit
	return r
}

func (r ApiGetRoleTargetsByUserIdAndRoleIdRequest) Execute() ([]RoleTarget, *APIResponse, error) {
	return r.ApiService.GetRoleTargetsByUserIdAndRoleIdExecute(r)
}

/*
GetRoleTargetsByUserIdAndRoleId Retrieve a role target by assignment type

Retrieves all role targets for an `APP_ADMIN`, `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user by user or group assignment type.
If the role isn't scoped to specific group targets or any app targets, an empty array `[]` is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleIdOrEncodedRoleId The `id` of the role or Base32 encoded `id` of the role name
	@return ApiGetRoleTargetsByUserIdAndRoleIdRequest
*/
func (a *RoleBTargetAdminAPIService) GetRoleTargetsByUserIdAndRoleId(ctx context.Context, userId string, roleIdOrEncodedRoleId string) ApiGetRoleTargetsByUserIdAndRoleIdRequest {
	return ApiGetRoleTargetsByUserIdAndRoleIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		userId:                userId,
		roleIdOrEncodedRoleId: roleIdOrEncodedRoleId,
		retryCount:            0,
	}
}

// Execute executes the request
//
//	@return []RoleTarget
func (a *RoleBTargetAdminAPIService) GetRoleTargetsByUserIdAndRoleIdExecute(r ApiGetRoleTargetsByUserIdAndRoleIdRequest) ([]RoleTarget, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RoleTarget
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.GetRoleTargetsByUserIdAndRoleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleIdOrEncodedRoleId}/targets"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrEncodedRoleId"+"}", url.PathEscape(parameterToString(r.roleIdOrEncodedRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.assignmentType != nil {
		localVarQueryParams.Add("assignmentType", parameterToString(*r.assignmentType, ""))
	}
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

type ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) After(after string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) Limit(limit int32) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	r.limit = &limit
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) Execute() ([]CatalogApplication, *APIResponse, error) {
	return r.ApiService.ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r)
}

/*
ListApplicationTargetsForApplicationAdministratorRoleForUser List all admin role app targets

Lists all app targets for an `APP_ADMIN` role assigned to a user. The response is a list that includes OIN-cataloged apps or app instances. The response payload for an app instance contains the `id` property, but an OIN-cataloged app payload doesn't.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx context.Context, userId string, roleAssignmentId string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []CatalogApplication
func (a *RoleBTargetAdminAPIService) ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) ([]CatalogApplication, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.ListApplicationTargetsForApplicationAdministratorRoleForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiListGroupTargetsForRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListGroupTargetsForRoleRequest) After(after string) ApiListGroupTargetsForRoleRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListGroupTargetsForRoleRequest) Limit(limit int32) ApiListGroupTargetsForRoleRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetsForRoleRequest) Execute() ([]Group, *APIResponse, error) {
	return r.ApiService.ListGroupTargetsForRoleExecute(r)
}

/*
ListGroupTargetsForRole List all admin role group targets

Lists all group targets for a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.
If the role isn't scoped to specific group targets, an empty array `[]` is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListGroupTargetsForRoleRequest
*/
func (a *RoleBTargetAdminAPIService) ListGroupTargetsForRole(ctx context.Context, userId string, roleAssignmentId string) ApiListGroupTargetsForRoleRequest {
	return ApiListGroupTargetsForRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []Group
func (a *RoleBTargetAdminAPIService) ListGroupTargetsForRoleExecute(r ApiListGroupTargetsForRoleRequest) ([]Group, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.ListGroupTargetsForRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppInstanceTargetFromAdminRoleForUserExecute(r)
}

/*
UnassignAppInstanceTargetFromAdminRoleForUser Unassign an admin role app instance target

Unassigns an app instance target from an `APP_ADMIN` role assignment to an admin user.

> **Note:** You can't remove the last app instance target from a role assignment since this causes an exception.
> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment and recreate a new one.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) UnassignAppInstanceTargetFromAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string, appId string) ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest {
	return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) UnassignAppInstanceTargetFromAdminRoleForUserExecute(r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.UnassignAppInstanceTargetFromAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUnassignAppTargetFromAppAdminRoleForUserRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppTargetFromAppAdminRoleForUserExecute(r)
}

/*
UnassignAppTargetFromAppAdminRoleForUser Unassign an admin role app target

Unassigns an OIN app target from an `APP_ADMIN` role assignment to an admin user.

> **Note:** You can't remove the last OIN app target from a role assignment since this causes an exception.
> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment to the user and recreate a new one.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiUnassignAppTargetFromAppAdminRoleForUserRequest
*/
func (a *RoleBTargetAdminAPIService) UnassignAppTargetFromAppAdminRoleForUser(ctx context.Context, userId string, roleAssignmentId string, appName string) ApiUnassignAppTargetFromAppAdminRoleForUserRequest {
	return ApiUnassignAppTargetFromAppAdminRoleForUserRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) UnassignAppTargetFromAppAdminRoleForUserExecute(r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.UnassignAppTargetFromAppAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUnassignGroupTargetFromUserAdminRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetAdminAPI
	userId           string
	roleAssignmentId string
	groupId          string
	retryCount       int32
}

func (r ApiUnassignGroupTargetFromUserAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignGroupTargetFromUserAdminRoleExecute(r)
}

/*
UnassignGroupTargetFromUserAdminRole Unassign an admin role group target

Unassigns a group target from a `USER_ADMIN`, `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to an admin user.

> **Note:** You can't remove the last group target from a role assignment since this causes an exception.
> If you need a role assignment that applies to all groups, delete the role assignment to the user and recreate a new one.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleAssignmentId The `id` of the role assignment
	@param groupId The `id` of the group
	@return ApiUnassignGroupTargetFromUserAdminRoleRequest
*/
func (a *RoleBTargetAdminAPIService) UnassignGroupTargetFromUserAdminRole(ctx context.Context, userId string, roleAssignmentId string, groupId string) ApiUnassignGroupTargetFromUserAdminRoleRequest {
	return ApiUnassignGroupTargetFromUserAdminRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		userId:           userId,
		roleAssignmentId: roleAssignmentId,
		groupId:          groupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetAdminAPIService) UnassignGroupTargetFromUserAdminRoleExecute(r ApiUnassignGroupTargetFromUserAdminRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetAdminAPIService.UnassignGroupTargetFromUserAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleAssignmentId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
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

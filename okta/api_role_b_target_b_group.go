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

type RoleBTargetBGroupAPI interface {

	/*
			AssignAppInstanceTargetToAppAdminRoleForGroup Assign a group role app instance target

			Assigns an app instance target to an `APP_ADMIN` role assignment to a group. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
		The role no longer applies to all app targets, but applies only to the specified target.

		> **Note:** You can target a mixture of both OIN app and app instance targets, but you can't assign permissions to manage all instances of an OIN app and then assign a subset of permissions to the same app.
		> For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage specific configurations of the Salesforce app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId The `id` of the group
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest
	*/
	AssignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest

	// AssignAppInstanceTargetToAppAdminRoleForGroupExecute executes the request
	AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error)

	/*
		AssignAppTargetToAdminRoleForGroup Assign a group role app target

		Assigns an OIN app target to an `APP_ADMIN` role assignment to a group. When you assign the first OIN app target, you reduce the scope of the role assignment. The role no longer applies to all app targets, but applies only to the specified target. An OIN app target that's assigned to the role overrides any existing instance targets of the OIN app. For example, if a user is assigned to administer a specific Facebook instance, a successful request to add an OIN app with `facebook` for `appName` makes that user the administrator for all Facebook instances.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId The `id` of the group
		@param roleAssignmentId The `id` of the role assignment
		@param appName Name of the app definition (the OIN catalog app key name)
		@return ApiAssignAppTargetToAdminRoleForGroupRequest
	*/
	AssignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string) ApiAssignAppTargetToAdminRoleForGroupRequest

	// AssignAppTargetToAdminRoleForGroupExecute executes the request
	AssignAppTargetToAdminRoleForGroupExecute(r ApiAssignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error)

	/*
			AssignGroupTargetToGroupAdminRole Assign a group role group target

			Assigns a group target to a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.
		When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets but applies only to the specified target.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId The `id` of the group
			@param roleAssignmentId The `id` of the role assignment
			@param targetGroupId
			@return ApiAssignGroupTargetToGroupAdminRoleRequest
	*/
	AssignGroupTargetToGroupAdminRole(ctx context.Context, groupId string, roleAssignmentId string, targetGroupId string) ApiAssignGroupTargetToGroupAdminRoleRequest

	// AssignGroupTargetToGroupAdminRoleExecute executes the request
	AssignGroupTargetToGroupAdminRoleExecute(r ApiAssignGroupTargetToGroupAdminRoleRequest) (*APIResponse, error)

	/*
		ListApplicationTargetsForApplicationAdministratorRoleForGroup List all group role app targets

		Lists all app targets for an `APP_ADMIN` role assignment to a group. The response includes a list of OIN-cataloged apps or app instances. The response payload for an app instance contains the `id` property, but an OIN-cataloged app doesn't.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId The `id` of the group
		@param roleAssignmentId The `id` of the role assignment
		@return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest
	*/
	ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest

	// ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute executes the request
	//  @return []CatalogApplication
	ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) ([]CatalogApplication, *APIResponse, error)

	/*
			ListGroupTargetsForGroupRole List all group role group targets

			Lists all group targets for a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.
		If the role isn't scoped to specific group targets, Okta returns an empty array `[]`.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId The `id` of the group
			@param roleAssignmentId The `id` of the role assignment
			@return ApiListGroupTargetsForGroupRoleRequest
	*/
	ListGroupTargetsForGroupRole(ctx context.Context, groupId string, roleAssignmentId string) ApiListGroupTargetsForGroupRoleRequest

	// ListGroupTargetsForGroupRoleExecute executes the request
	//  @return []Group
	ListGroupTargetsForGroupRoleExecute(r ApiListGroupTargetsForGroupRoleRequest) ([]Group, *APIResponse, error)

	/*
			UnassignAppInstanceTargetToAppAdminRoleForGroup Unassign a group role app instance target

			Unassigns an app instance target from an `APP_ADMIN` role assignment to a group

		> **Note:** You can't remove the last app instance target from a role assignment.
		> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment with the target and create another one. See [Unassign a group role](/openapi/okta-management/management/tag/RoleAssignmentBGroup/#tag/RoleAssignmentBGroup/operation/unassignRoleFromGroup).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId The `id` of the group
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@param appId Application ID
			@return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest
	*/
	UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string, appId string) ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest

	// UnassignAppInstanceTargetToAppAdminRoleForGroupExecute executes the request
	UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error)

	/*
			UnassignAppTargetToAdminRoleForGroup Unassign a group role app target

			Unassigns an OIN app target from an `APP_ADMIN` role assignment to a group

		> **Note:** You can't remove the last app target from a role assignment.
		> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment with the target and create another one. See [Unassign a group role](/openapi/okta-management/management/tag/RoleAssignmentBGroup/#tag/RoleAssignmentBGroup/operation/unassignRoleFromGroup).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId The `id` of the group
			@param roleAssignmentId The `id` of the role assignment
			@param appName Name of the app definition (the OIN catalog app key name)
			@return ApiUnassignAppTargetToAdminRoleForGroupRequest
	*/
	UnassignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string) ApiUnassignAppTargetToAdminRoleForGroupRequest

	// UnassignAppTargetToAdminRoleForGroupExecute executes the request
	UnassignAppTargetToAdminRoleForGroupExecute(r ApiUnassignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error)

	/*
		UnassignGroupTargetFromGroupAdminRole Unassign a group role group target

		Unassigns a group target from a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId The `id` of the group
		@param roleAssignmentId The `id` of the role assignment
		@param targetGroupId
		@return ApiUnassignGroupTargetFromGroupAdminRoleRequest
	*/
	UnassignGroupTargetFromGroupAdminRole(ctx context.Context, groupId string, roleAssignmentId string, targetGroupId string) ApiUnassignGroupTargetFromGroupAdminRoleRequest

	// UnassignGroupTargetFromGroupAdminRoleExecute executes the request
	UnassignGroupTargetFromGroupAdminRoleExecute(r ApiUnassignGroupTargetFromGroupAdminRoleRequest) (*APIResponse, error)
}

// RoleBTargetBGroupAPIService RoleBTargetBGroupAPI service
type RoleBTargetBGroupAPIService service

type ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r)
}

/*
AssignAppInstanceTargetToAppAdminRoleForGroup Assign a group role app instance target

Assigns an app instance target to an `APP_ADMIN` role assignment to a group. When you assign the first OIN app or app instance target, you reduce the scope of the role assignment.
The role no longer applies to all app targets, but applies only to the specified target.

> **Note:** You can target a mixture of both OIN app and app instance targets, but you can't assign permissions to manage all instances of an OIN app and then assign a subset of permissions to the same app.
> For example, you can't specify that an admin has access to manage all instances of the Salesforce app and then also manage specific configurations of the Salesforce app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest
*/
func (a *RoleBTargetBGroupAPIService) AssignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest {
	return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.AssignAppInstanceTargetToAppAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiAssignAppTargetToAdminRoleForGroupRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiAssignAppTargetToAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetToAdminRoleForGroupExecute(r)
}

/*
AssignAppTargetToAdminRoleForGroup Assign a group role app target

Assigns an OIN app target to an `APP_ADMIN` role assignment to a group. When you assign the first OIN app target, you reduce the scope of the role assignment. The role no longer applies to all app targets, but applies only to the specified target. An OIN app target that's assigned to the role overrides any existing instance targets of the OIN app. For example, if a user is assigned to administer a specific Facebook instance, a successful request to add an OIN app with `facebook` for `appName` makes that user the administrator for all Facebook instances.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiAssignAppTargetToAdminRoleForGroupRequest
*/
func (a *RoleBTargetBGroupAPIService) AssignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string) ApiAssignAppTargetToAdminRoleForGroupRequest {
	return ApiAssignAppTargetToAdminRoleForGroupRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) AssignAppTargetToAdminRoleForGroupExecute(r ApiAssignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.AssignAppTargetToAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiAssignGroupTargetToGroupAdminRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	targetGroupId    string
	retryCount       int32
}

func (r ApiAssignGroupTargetToGroupAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetToGroupAdminRoleExecute(r)
}

/*
AssignGroupTargetToGroupAdminRole Assign a group role group target

Assigns a group target to a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.
When you assign the first group target, you reduce the scope of the role assignment. The role no longer applies to all targets but applies only to the specified target.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param targetGroupId
	@return ApiAssignGroupTargetToGroupAdminRoleRequest
*/
func (a *RoleBTargetBGroupAPIService) AssignGroupTargetToGroupAdminRole(ctx context.Context, groupId string, roleAssignmentId string, targetGroupId string) ApiAssignGroupTargetToGroupAdminRoleRequest {
	return ApiAssignGroupTargetToGroupAdminRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		targetGroupId:    targetGroupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) AssignGroupTargetToGroupAdminRoleExecute(r ApiAssignGroupTargetToGroupAdminRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.AssignGroupTargetToGroupAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", url.PathEscape(parameterToString(r.targetGroupId, "")), -1)

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

type ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) After(after string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) Limit(limit int32) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	r.limit = &limit
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) Execute() ([]CatalogApplication, *APIResponse, error) {
	return r.ApiService.ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r)
}

/*
ListApplicationTargetsForApplicationAdministratorRoleForGroup List all group role app targets

Lists all app targets for an `APP_ADMIN` role assignment to a group. The response includes a list of OIN-cataloged apps or app instances. The response payload for an app instance contains the `id` property, but an OIN-cataloged app doesn't.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest
*/
func (a *RoleBTargetBGroupAPIService) ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []CatalogApplication
func (a *RoleBTargetBGroupAPIService) ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) ([]CatalogApplication, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.ListApplicationTargetsForApplicationAdministratorRoleForGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiListGroupTargetsForGroupRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	after            *string
	limit            *int32
	retryCount       int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListGroupTargetsForGroupRoleRequest) After(after string) ApiListGroupTargetsForGroupRoleRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListGroupTargetsForGroupRoleRequest) Limit(limit int32) ApiListGroupTargetsForGroupRoleRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetsForGroupRoleRequest) Execute() ([]Group, *APIResponse, error) {
	return r.ApiService.ListGroupTargetsForGroupRoleExecute(r)
}

/*
ListGroupTargetsForGroupRole List all group role group targets

Lists all group targets for a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.
If the role isn't scoped to specific group targets, Okta returns an empty array `[]`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@return ApiListGroupTargetsForGroupRoleRequest
*/
func (a *RoleBTargetBGroupAPIService) ListGroupTargetsForGroupRole(ctx context.Context, groupId string, roleAssignmentId string) ApiListGroupTargetsForGroupRoleRequest {
	return ApiListGroupTargetsForGroupRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []Group
func (a *RoleBTargetBGroupAPIService) ListGroupTargetsForGroupRoleExecute(r ApiListGroupTargetsForGroupRoleRequest) ([]Group, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.ListGroupTargetsForGroupRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	appName          string
	appId            string
	retryCount       int32
}

func (r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r)
}

/*
UnassignAppInstanceTargetToAppAdminRoleForGroup Unassign a group role app instance target

Unassigns an app instance target from an `APP_ADMIN` role assignment to a group

> **Note:** You can't remove the last app instance target from a role assignment.
> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment with the target and create another one. See [Unassign a group role](/openapi/okta-management/management/tag/RoleAssignmentBGroup/#tag/RoleAssignmentBGroup/operation/unassignRoleFromGroup).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@param appId Application ID
	@return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest
*/
func (a *RoleBTargetBGroupAPIService) UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string, appId string) ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest {
	return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		appId:            appId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.UnassignAppInstanceTargetToAppAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiUnassignAppTargetToAdminRoleForGroupRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	appName          string
	retryCount       int32
}

func (r ApiUnassignAppTargetToAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppTargetToAdminRoleForGroupExecute(r)
}

/*
UnassignAppTargetToAdminRoleForGroup Unassign a group role app target

Unassigns an OIN app target from an `APP_ADMIN` role assignment to a group

> **Note:** You can't remove the last app target from a role assignment.
> If you need a role assignment that applies to all apps, delete the `APP_ADMIN` role assignment with the target and create another one. See [Unassign a group role](/openapi/okta-management/management/tag/RoleAssignmentBGroup/#tag/RoleAssignmentBGroup/operation/unassignRoleFromGroup).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param appName Name of the app definition (the OIN catalog app key name)
	@return ApiUnassignAppTargetToAdminRoleForGroupRequest
*/
func (a *RoleBTargetBGroupAPIService) UnassignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleAssignmentId string, appName string) ApiUnassignAppTargetToAdminRoleForGroupRequest {
	return ApiUnassignAppTargetToAdminRoleForGroupRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		appName:          appName,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) UnassignAppTargetToAdminRoleForGroupExecute(r ApiUnassignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.UnassignAppTargetToAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
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

type ApiUnassignGroupTargetFromGroupAdminRoleRequest struct {
	ctx              context.Context
	ApiService       RoleBTargetBGroupAPI
	groupId          string
	roleAssignmentId string
	targetGroupId    string
	retryCount       int32
}

func (r ApiUnassignGroupTargetFromGroupAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignGroupTargetFromGroupAdminRoleExecute(r)
}

/*
UnassignGroupTargetFromGroupAdminRole Unassign a group role group target

Unassigns a group target from a [`USER_ADMIN`](/openapi/okta-management/guides/roles/#standard-roles), `HELP_DESK_ADMIN`, or `GROUP_MEMBERSHIP_ADMIN` role assignment to a group.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleAssignmentId The `id` of the role assignment
	@param targetGroupId
	@return ApiUnassignGroupTargetFromGroupAdminRoleRequest
*/
func (a *RoleBTargetBGroupAPIService) UnassignGroupTargetFromGroupAdminRole(ctx context.Context, groupId string, roleAssignmentId string, targetGroupId string) ApiUnassignGroupTargetFromGroupAdminRoleRequest {
	return ApiUnassignGroupTargetFromGroupAdminRoleRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		roleAssignmentId: roleAssignmentId,
		targetGroupId:    targetGroupId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *RoleBTargetBGroupAPIService) UnassignGroupTargetFromGroupAdminRoleExecute(r ApiUnassignGroupTargetFromGroupAdminRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleBTargetBGroupAPIService.UnassignGroupTargetFromGroupAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleAssignmentId}/targets/groups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleAssignmentId"+"}", url.PathEscape(parameterToString(r.roleAssignmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", url.PathEscape(parameterToString(r.targetGroupId, "")), -1)

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

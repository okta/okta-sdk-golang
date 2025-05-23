/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"strings"
)


type RoleTargetAPI interface {

	/*
	AssignAllAppsAsTargetToRoleForUser Assign all Apps as Target to Role

	Assigns all Apps as Target to Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@return ApiAssignAllAppsAsTargetToRoleForUserRequest
	*/
	AssignAllAppsAsTargetToRoleForUser(ctx context.Context, userId string, roleId string) ApiAssignAllAppsAsTargetToRoleForUserRequest

	// AssignAllAppsAsTargetToRoleForUserExecute executes the request
	AssignAllAppsAsTargetToRoleForUserExecute(r ApiAssignAllAppsAsTargetToRoleForUserRequest) (*APIResponse, error)

	/*
	AssignAppInstanceTargetToAppAdminRoleForGroup Assign an Application Instance Target to Application Administrator Role

	Assigns App Instance Target to App Administrator Role given to a Group

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appId Application ID
	@return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest
	*/
	AssignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest

	// AssignAppInstanceTargetToAppAdminRoleForGroupExecute executes the request
	AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error)

	/*
	AssignAppInstanceTargetToAppAdminRoleForUser Assign an Application Instance Target to an Application Administrator Role

	Assigns anapplication instance target to appplication administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appId Application ID
	@return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest
	*/
	AssignAppInstanceTargetToAppAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest

	// AssignAppInstanceTargetToAppAdminRoleForUserExecute executes the request
	AssignAppInstanceTargetToAppAdminRoleForUserExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) (*APIResponse, error)

	/*
	AssignAppTargetInstanceRoleForClient Replace an App Instance Target for a Client

	Replaces an App Instance Target by `clientId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appInstanceId `id` of the application instance
	@return ApiAssignAppTargetInstanceRoleForClientRequest
	*/
	AssignAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleId string, appName string, appInstanceId string) ApiAssignAppTargetInstanceRoleForClientRequest

	// AssignAppTargetInstanceRoleForClientExecute executes the request
	AssignAppTargetInstanceRoleForClientExecute(r ApiAssignAppTargetInstanceRoleForClientRequest) (*APIResponse, error)

	/*
	AssignAppTargetRoleToClient Assign an App Target to a Client

	Assigns an Application Target to a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiAssignAppTargetRoleToClientRequest
	*/
	AssignAppTargetRoleToClient(ctx context.Context, clientId string, roleId string, appName string) ApiAssignAppTargetRoleToClientRequest

	// AssignAppTargetRoleToClientExecute executes the request
	AssignAppTargetRoleToClientExecute(r ApiAssignAppTargetRoleToClientRequest) (*APIResponse, error)

	/*
	AssignAppTargetToAdminRoleForGroup Assign an Application Target to Administrator Role

	Assigns an application target to administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiAssignAppTargetToAdminRoleForGroupRequest
	*/
	AssignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string) ApiAssignAppTargetToAdminRoleForGroupRequest

	// AssignAppTargetToAdminRoleForGroupExecute executes the request
	AssignAppTargetToAdminRoleForGroupExecute(r ApiAssignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error)

	/*
	AssignAppTargetToAdminRoleForUser Assign an Application Target to Administrator Role

	Assigns an application target to administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiAssignAppTargetToAdminRoleForUserRequest
	*/
	AssignAppTargetToAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string) ApiAssignAppTargetToAdminRoleForUserRequest

	// AssignAppTargetToAdminRoleForUserExecute executes the request
	AssignAppTargetToAdminRoleForUserExecute(r ApiAssignAppTargetToAdminRoleForUserRequest) (*APIResponse, error)

	/*
	AssignGroupTargetRoleForClient Assign a Group Target to a Client

	Assigns a Group Target to a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param groupId The `id` of the group
	@return ApiAssignGroupTargetRoleForClientRequest
	*/
	AssignGroupTargetRoleForClient(ctx context.Context, clientId string, roleId string, groupId string) ApiAssignGroupTargetRoleForClientRequest

	// AssignGroupTargetRoleForClientExecute executes the request
	AssignGroupTargetRoleForClientExecute(r ApiAssignGroupTargetRoleForClientRequest) (*APIResponse, error)

	/*
	AssignGroupTargetToGroupAdminRole Assign a Group Target to a Group Role

	Assigns a group target to a group role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param targetGroupId
	@return ApiAssignGroupTargetToGroupAdminRoleRequest
	*/
	AssignGroupTargetToGroupAdminRole(ctx context.Context, groupId string, roleId string, targetGroupId string) ApiAssignGroupTargetToGroupAdminRoleRequest

	// AssignGroupTargetToGroupAdminRoleExecute executes the request
	AssignGroupTargetToGroupAdminRoleExecute(r ApiAssignGroupTargetToGroupAdminRoleRequest) (*APIResponse, error)

	/*
	AssignGroupTargetToUserRole Assign a Group Target to Role

	Assigns a Group Target to Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param groupId The `id` of the group
	@return ApiAssignGroupTargetToUserRoleRequest
	*/
	AssignGroupTargetToUserRole(ctx context.Context, userId string, roleId string, groupId string) ApiAssignGroupTargetToUserRoleRequest

	// AssignGroupTargetToUserRoleExecute executes the request
	AssignGroupTargetToUserRoleExecute(r ApiAssignGroupTargetToUserRoleRequest) (*APIResponse, error)

	/*
	ListAppTargetRoleToClient List all App Targets assigned to a Client

	Lists App Targets for Client and Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@return ApiListAppTargetRoleToClientRequest
	*/
	ListAppTargetRoleToClient(ctx context.Context, clientId string, roleId string) ApiListAppTargetRoleToClientRequest

	// ListAppTargetRoleToClientExecute executes the request
	//  @return Client
	ListAppTargetRoleToClientExecute(r ApiListAppTargetRoleToClientRequest) (*Client, *APIResponse, error)

	/*
	ListApplicationTargetsForApplicationAdministratorRoleForGroup List all Application Targets for an Application Administrator Role

	Lists all App targets for an `APP_ADMIN` Role assigned to a Group. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest
	*/
	ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx context.Context, groupId string, roleId string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest

	// ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute executes the request
	//  @return []CatalogApplication
	ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) ([]CatalogApplication, *APIResponse, error)

	/*
	ListApplicationTargetsForApplicationAdministratorRoleForUser List all Application Targets for Application Administrator Role

	Lists all App targets for an `APP_ADMIN` Role assigned to a User. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest
	*/
	ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx context.Context, userId string, roleId string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest

	// ListApplicationTargetsForApplicationAdministratorRoleForUserExecute executes the request
	//  @return []CatalogApplication
	ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) ([]CatalogApplication, *APIResponse, error)

	/*
	ListGroupTargetRoleForClient List all Group Targets for a Client

	Lists all Group Targets by `clientId` and `roleId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@return ApiListGroupTargetRoleForClientRequest
	*/
	ListGroupTargetRoleForClient(ctx context.Context, clientId string, roleId string) ApiListGroupTargetRoleForClientRequest

	// ListGroupTargetRoleForClientExecute executes the request
	//  @return Client
	ListGroupTargetRoleForClientExecute(r ApiListGroupTargetRoleForClientRequest) (*Client, *APIResponse, error)

	/*
	ListGroupTargetsForGroupRole List all Group Targets for a Group Role

	Lists all group targets for a group role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@return ApiListGroupTargetsForGroupRoleRequest
	*/
	ListGroupTargetsForGroupRole(ctx context.Context, groupId string, roleId string) ApiListGroupTargetsForGroupRoleRequest

	// ListGroupTargetsForGroupRoleExecute executes the request
	//  @return []Group
	ListGroupTargetsForGroupRoleExecute(r ApiListGroupTargetsForGroupRoleRequest) ([]Group, *APIResponse, error)

	/*
	ListGroupTargetsForRole List all Group Targets for Role

	Lists all group targets for role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@return ApiListGroupTargetsForRoleRequest
	*/
	ListGroupTargetsForRole(ctx context.Context, userId string, roleId string) ApiListGroupTargetsForRoleRequest

	// ListGroupTargetsForRoleExecute executes the request
	//  @return []Group
	ListGroupTargetsForRoleExecute(r ApiListGroupTargetsForRoleRequest) ([]Group, *APIResponse, error)

	/*
	RemoveAppTargetInstanceRoleForClient Delete an App Instance Target for a Client

	Deletes an App Instance Target from a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appInstanceId `id` of the application instance
	@return ApiRemoveAppTargetInstanceRoleForClientRequest
	*/
	RemoveAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleId string, appName string, appInstanceId string) ApiRemoveAppTargetInstanceRoleForClientRequest

	// RemoveAppTargetInstanceRoleForClientExecute executes the request
	RemoveAppTargetInstanceRoleForClientExecute(r ApiRemoveAppTargetInstanceRoleForClientRequest) (*APIResponse, error)

	/*
	RemoveAppTargetRoleFromClient Remove an App Target from a Client

	Removes an Application Target by `clientId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiRemoveAppTargetRoleFromClientRequest
	*/
	RemoveAppTargetRoleFromClient(ctx context.Context, clientId string, roleId string, appName string) ApiRemoveAppTargetRoleFromClientRequest

	// RemoveAppTargetRoleFromClientExecute executes the request
	RemoveAppTargetRoleFromClientExecute(r ApiRemoveAppTargetRoleFromClientRequest) (*APIResponse, error)

	/*
	RemoveGroupTargetRoleFromClient Delete a Group Target from a Client

	Deletes a Group Target from a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId `client_id` of the app
	@param roleId `id` of the Role
	@param groupId The `id` of the group
	@return ApiRemoveGroupTargetRoleFromClientRequest
	*/
	RemoveGroupTargetRoleFromClient(ctx context.Context, clientId string, roleId string, groupId string) ApiRemoveGroupTargetRoleFromClientRequest

	// RemoveGroupTargetRoleFromClientExecute executes the request
	RemoveGroupTargetRoleFromClientExecute(r ApiRemoveGroupTargetRoleFromClientRequest) (*APIResponse, error)

	/*
	UnassignAppInstanceTargetFromAdminRoleForUser Unassign an Application Instance Target from an Application Administrator Role

	Unassigns an application instance target from an application administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appId Application ID
	@return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest
	*/
	UnassignAppInstanceTargetFromAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string, appId string) ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest

	// UnassignAppInstanceTargetFromAdminRoleForUserExecute executes the request
	UnassignAppInstanceTargetFromAdminRoleForUserExecute(r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) (*APIResponse, error)

	/*
	UnassignAppInstanceTargetToAppAdminRoleForGroup Unassign an Application Instance Target from an Application Administrator Role

	Unassigns an application instance target from application administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@param appId Application ID
	@return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest
	*/
	UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string, appId string) ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest

	// UnassignAppInstanceTargetToAppAdminRoleForGroupExecute executes the request
	UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error)

	/*
	UnassignAppTargetFromAppAdminRoleForUser Unassign an Application Target from an Application Administrator Role

	Unassigns an application target from application administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiUnassignAppTargetFromAppAdminRoleForUserRequest
	*/
	UnassignAppTargetFromAppAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string) ApiUnassignAppTargetFromAppAdminRoleForUserRequest

	// UnassignAppTargetFromAppAdminRoleForUserExecute executes the request
	UnassignAppTargetFromAppAdminRoleForUserExecute(r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) (*APIResponse, error)

	/*
	UnassignAppTargetToAdminRoleForGroup Unassign an Application Target from Application Administrator Role

	Unassigns an application target from application administrator role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param appName Application name for the app type
	@return ApiUnassignAppTargetToAdminRoleForGroupRequest
	*/
	UnassignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string) ApiUnassignAppTargetToAdminRoleForGroupRequest

	// UnassignAppTargetToAdminRoleForGroupExecute executes the request
	UnassignAppTargetToAdminRoleForGroupExecute(r ApiUnassignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error)

	/*
	UnassignGroupTargetFromGroupAdminRole Unassign a Group Target from a Group Role

	Unassigns a group target from a group role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The `id` of the group
	@param roleId `id` of the Role
	@param targetGroupId
	@return ApiUnassignGroupTargetFromGroupAdminRoleRequest
	*/
	UnassignGroupTargetFromGroupAdminRole(ctx context.Context, groupId string, roleId string, targetGroupId string) ApiUnassignGroupTargetFromGroupAdminRoleRequest

	// UnassignGroupTargetFromGroupAdminRoleExecute executes the request
	UnassignGroupTargetFromGroupAdminRoleExecute(r ApiUnassignGroupTargetFromGroupAdminRoleRequest) (*APIResponse, error)

	/*
	UnassignGroupTargetFromUserAdminRole Unassign a Group Target from Role

	Unassigns a Group Target from Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@param roleId `id` of the Role
	@param groupId The `id` of the group
	@return ApiUnassignGroupTargetFromUserAdminRoleRequest
	*/
	UnassignGroupTargetFromUserAdminRole(ctx context.Context, userId string, roleId string, groupId string) ApiUnassignGroupTargetFromUserAdminRoleRequest

	// UnassignGroupTargetFromUserAdminRoleExecute executes the request
	UnassignGroupTargetFromUserAdminRoleExecute(r ApiUnassignGroupTargetFromUserAdminRoleRequest) (*APIResponse, error)
}

// RoleTargetAPIService RoleTargetAPI service
type RoleTargetAPIService service

type ApiAssignAllAppsAsTargetToRoleForUserRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	retryCount int32
}

func (r ApiAssignAllAppsAsTargetToRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAllAppsAsTargetToRoleForUserExecute(r)
}

/*
AssignAllAppsAsTargetToRoleForUser Assign all Apps as Target to Role

Assigns all Apps as Target to Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @return ApiAssignAllAppsAsTargetToRoleForUserRequest
*/
func (a *RoleTargetAPIService) AssignAllAppsAsTargetToRoleForUser(ctx context.Context, userId string, roleId string) ApiAssignAllAppsAsTargetToRoleForUserRequest {
	return ApiAssignAllAppsAsTargetToRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAllAppsAsTargetToRoleForUserExecute(r ApiAssignAllAppsAsTargetToRoleForUserRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAllAppsAsTargetToRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	appName string
	appId string
	retryCount int32
}

func (r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r)
}

/*
AssignAppInstanceTargetToAppAdminRoleForGroup Assign an Application Instance Target to Application Administrator Role

Assigns App Instance Target to App Administrator Role given to a Group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appId Application ID
 @return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest
*/
func (a *RoleTargetAPIService) AssignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest {
	return ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		appName: appName,
		appId: appId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppInstanceTargetToAppAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	appName string
	appId string
	retryCount int32
}

func (r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppInstanceTargetToAppAdminRoleForUserExecute(r)
}

/*
AssignAppInstanceTargetToAppAdminRoleForUser Assign an Application Instance Target to an Application Administrator Role

Assigns anapplication instance target to appplication administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appId Application ID
 @return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest
*/
func (a *RoleTargetAPIService) AssignAppInstanceTargetToAppAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string, appId string) ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest {
	return ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		appName: appName,
		appId: appId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppInstanceTargetToAppAdminRoleForUserExecute(r ApiAssignAppInstanceTargetToAppAdminRoleForUserRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppInstanceTargetToAppAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAssignAppTargetInstanceRoleForClientRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	appName string
	appInstanceId string
	retryCount int32
}

func (r ApiAssignAppTargetInstanceRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetInstanceRoleForClientExecute(r)
}

/*
AssignAppTargetInstanceRoleForClient Replace an App Instance Target for a Client

Replaces an App Instance Target by `clientId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appInstanceId `id` of the application instance
 @return ApiAssignAppTargetInstanceRoleForClientRequest
*/
func (a *RoleTargetAPIService) AssignAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleId string, appName string, appInstanceId string) ApiAssignAppTargetInstanceRoleForClientRequest {
	return ApiAssignAppTargetInstanceRoleForClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		appName: appName,
		appInstanceId: appInstanceId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppTargetInstanceRoleForClientExecute(r ApiAssignAppTargetInstanceRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppTargetInstanceRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}/{appInstanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appInstanceId"+"}", url.PathEscape(parameterToString(r.appInstanceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiAssignAppTargetRoleToClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetRoleToClientExecute(r)
}

/*
AssignAppTargetRoleToClient Assign an App Target to a Client

Assigns an Application Target to a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiAssignAppTargetRoleToClientRequest
*/
func (a *RoleTargetAPIService) AssignAppTargetRoleToClient(ctx context.Context, clientId string, roleId string, appName string) ApiAssignAppTargetRoleToClientRequest {
	return ApiAssignAppTargetRoleToClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppTargetRoleToClientExecute(r ApiAssignAppTargetRoleToClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppTargetRoleToClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiAssignAppTargetToAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetToAdminRoleForGroupExecute(r)
}

/*
AssignAppTargetToAdminRoleForGroup Assign an Application Target to Administrator Role

Assigns an application target to administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiAssignAppTargetToAdminRoleForGroupRequest
*/
func (a *RoleTargetAPIService) AssignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string) ApiAssignAppTargetToAdminRoleForGroupRequest {
	return ApiAssignAppTargetToAdminRoleForGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppTargetToAdminRoleForGroupExecute(r ApiAssignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppTargetToAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiAssignAppTargetToAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignAppTargetToAdminRoleForUserExecute(r)
}

/*
AssignAppTargetToAdminRoleForUser Assign an Application Target to Administrator Role

Assigns an application target to administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiAssignAppTargetToAdminRoleForUserRequest
*/
func (a *RoleTargetAPIService) AssignAppTargetToAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string) ApiAssignAppTargetToAdminRoleForUserRequest {
	return ApiAssignAppTargetToAdminRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignAppTargetToAdminRoleForUserExecute(r ApiAssignAppTargetToAdminRoleForUserRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignAppTargetToAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	groupId string
	retryCount int32
}

func (r ApiAssignGroupTargetRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetRoleForClientExecute(r)
}

/*
AssignGroupTargetRoleForClient Assign a Group Target to a Client

Assigns a Group Target to a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param groupId The `id` of the group
 @return ApiAssignGroupTargetRoleForClientRequest
*/
func (a *RoleTargetAPIService) AssignGroupTargetRoleForClient(ctx context.Context, clientId string, roleId string, groupId string) ApiAssignGroupTargetRoleForClientRequest {
	return ApiAssignGroupTargetRoleForClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		groupId: groupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignGroupTargetRoleForClientExecute(r ApiAssignGroupTargetRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignGroupTargetRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	targetGroupId string
	retryCount int32
}

func (r ApiAssignGroupTargetToGroupAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetToGroupAdminRoleExecute(r)
}

/*
AssignGroupTargetToGroupAdminRole Assign a Group Target to a Group Role

Assigns a group target to a group role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param targetGroupId
 @return ApiAssignGroupTargetToGroupAdminRoleRequest
*/
func (a *RoleTargetAPIService) AssignGroupTargetToGroupAdminRole(ctx context.Context, groupId string, roleId string, targetGroupId string) ApiAssignGroupTargetToGroupAdminRoleRequest {
	return ApiAssignGroupTargetToGroupAdminRoleRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		targetGroupId: targetGroupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignGroupTargetToGroupAdminRoleExecute(r ApiAssignGroupTargetToGroupAdminRoleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignGroupTargetToGroupAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	groupId string
	retryCount int32
}

func (r ApiAssignGroupTargetToUserRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignGroupTargetToUserRoleExecute(r)
}

/*
AssignGroupTargetToUserRole Assign a Group Target to Role

Assigns a Group Target to Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param groupId The `id` of the group
 @return ApiAssignGroupTargetToUserRoleRequest
*/
func (a *RoleTargetAPIService) AssignGroupTargetToUserRole(ctx context.Context, userId string, roleId string, groupId string) ApiAssignGroupTargetToUserRoleRequest {
	return ApiAssignGroupTargetToUserRoleRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		groupId: groupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) AssignGroupTargetToUserRoleExecute(r ApiAssignGroupTargetToUserRoleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.AssignGroupTargetToUserRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination).
func (r ApiListAppTargetRoleToClientRequest) After(after string) ApiListAppTargetRoleToClientRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListAppTargetRoleToClientRequest) Limit(limit int32) ApiListAppTargetRoleToClientRequest {
	r.limit = &limit
	return r
}

func (r ApiListAppTargetRoleToClientRequest) Execute() (*Client, *APIResponse, error) {
	return r.ApiService.ListAppTargetRoleToClientExecute(r)
}

/*
ListAppTargetRoleToClient List all App Targets assigned to a Client

Lists App Targets for Client and Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @return ApiListAppTargetRoleToClientRequest
*/
func (a *RoleTargetAPIService) ListAppTargetRoleToClient(ctx context.Context, clientId string, roleId string) ApiListAppTargetRoleToClientRequest {
	return ApiListAppTargetRoleToClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return Client
func (a *RoleTargetAPIService) ListAppTargetRoleToClientExecute(r ApiListAppTargetRoleToClientRequest) (*Client, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Client
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListAppTargetRoleToClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) After(after string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	r.after = &after
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) Limit(limit int32) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	r.limit = &limit
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) Execute() ([]CatalogApplication, *APIResponse, error) {
	return r.ApiService.ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r)
}

/*
ListApplicationTargetsForApplicationAdministratorRoleForGroup List all Application Targets for an Application Administrator Role

Lists all App targets for an `APP_ADMIN` Role assigned to a Group. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest
*/
func (a *RoleTargetAPIService) ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx context.Context, groupId string, roleId string) ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest {
	return ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []CatalogApplication
func (a *RoleTargetAPIService) ListApplicationTargetsForApplicationAdministratorRoleForGroupExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForGroupRequest) ([]CatalogApplication, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CatalogApplication
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListApplicationTargetsForApplicationAdministratorRoleForGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) After(after string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	r.after = &after
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) Limit(limit int32) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	r.limit = &limit
	return r
}

func (r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) Execute() ([]CatalogApplication, *APIResponse, error) {
	return r.ApiService.ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r)
}

/*
ListApplicationTargetsForApplicationAdministratorRoleForUser List all Application Targets for Application Administrator Role

Lists all App targets for an `APP_ADMIN` Role assigned to a User. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest
*/
func (a *RoleTargetAPIService) ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx context.Context, userId string, roleId string) ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest {
	return ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []CatalogApplication
func (a *RoleTargetAPIService) ListApplicationTargetsForApplicationAdministratorRoleForUserExecute(r ApiListApplicationTargetsForApplicationAdministratorRoleForUserRequest) ([]CatalogApplication, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CatalogApplication
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListApplicationTargetsForApplicationAdministratorRoleForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination).
func (r ApiListGroupTargetRoleForClientRequest) After(after string) ApiListGroupTargetRoleForClientRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListGroupTargetRoleForClientRequest) Limit(limit int32) ApiListGroupTargetRoleForClientRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetRoleForClientRequest) Execute() (*Client, *APIResponse, error) {
	return r.ApiService.ListGroupTargetRoleForClientExecute(r)
}

/*
ListGroupTargetRoleForClient List all Group Targets for a Client

Lists all Group Targets by `clientId` and `roleId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @return ApiListGroupTargetRoleForClientRequest
*/
func (a *RoleTargetAPIService) ListGroupTargetRoleForClient(ctx context.Context, clientId string, roleId string) ApiListGroupTargetRoleForClientRequest {
	return ApiListGroupTargetRoleForClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return Client
func (a *RoleTargetAPIService) ListGroupTargetRoleForClientExecute(r ApiListGroupTargetRoleForClientRequest) (*Client, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Client
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListGroupTargetRoleForClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

func (r ApiListGroupTargetsForGroupRoleRequest) After(after string) ApiListGroupTargetsForGroupRoleRequest {
	r.after = &after
	return r
}

func (r ApiListGroupTargetsForGroupRoleRequest) Limit(limit int32) ApiListGroupTargetsForGroupRoleRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetsForGroupRoleRequest) Execute() ([]Group, *APIResponse, error) {
	return r.ApiService.ListGroupTargetsForGroupRoleExecute(r)
}

/*
ListGroupTargetsForGroupRole List all Group Targets for a Group Role

Lists all group targets for a group role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @return ApiListGroupTargetsForGroupRoleRequest
*/
func (a *RoleTargetAPIService) ListGroupTargetsForGroupRole(ctx context.Context, groupId string, roleId string) ApiListGroupTargetsForGroupRoleRequest {
	return ApiListGroupTargetsForGroupRoleRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []Group
func (a *RoleTargetAPIService) ListGroupTargetsForGroupRoleExecute(r ApiListGroupTargetsForGroupRoleRequest) ([]Group, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Group
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListGroupTargetsForGroupRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	after *string
	limit *int32
	retryCount int32
}

func (r ApiListGroupTargetsForRoleRequest) After(after string) ApiListGroupTargetsForRoleRequest {
	r.after = &after
	return r
}

func (r ApiListGroupTargetsForRoleRequest) Limit(limit int32) ApiListGroupTargetsForRoleRequest {
	r.limit = &limit
	return r
}

func (r ApiListGroupTargetsForRoleRequest) Execute() ([]Group, *APIResponse, error) {
	return r.ApiService.ListGroupTargetsForRoleExecute(r)
}

/*
ListGroupTargetsForRole List all Group Targets for Role

Lists all group targets for role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @return ApiListGroupTargetsForRoleRequest
*/
func (a *RoleTargetAPIService) ListGroupTargetsForRole(ctx context.Context, userId string, roleId string) ApiListGroupTargetsForRoleRequest {
	return ApiListGroupTargetsForRoleRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []Group
func (a *RoleTargetAPIService) ListGroupTargetsForRoleExecute(r ApiListGroupTargetsForRoleRequest) ([]Group, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Group
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.ListGroupTargetsForRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	appName string
	appInstanceId string
	retryCount int32
}

func (r ApiRemoveAppTargetInstanceRoleForClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveAppTargetInstanceRoleForClientExecute(r)
}

/*
RemoveAppTargetInstanceRoleForClient Delete an App Instance Target for a Client

Deletes an App Instance Target from a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appInstanceId `id` of the application instance
 @return ApiRemoveAppTargetInstanceRoleForClientRequest
*/
func (a *RoleTargetAPIService) RemoveAppTargetInstanceRoleForClient(ctx context.Context, clientId string, roleId string, appName string, appInstanceId string) ApiRemoveAppTargetInstanceRoleForClientRequest {
	return ApiRemoveAppTargetInstanceRoleForClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		appName: appName,
		appInstanceId: appInstanceId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) RemoveAppTargetInstanceRoleForClientExecute(r ApiRemoveAppTargetInstanceRoleForClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.RemoveAppTargetInstanceRoleForClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}/{appInstanceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appInstanceId"+"}", url.PathEscape(parameterToString(r.appInstanceId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiRemoveAppTargetRoleFromClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveAppTargetRoleFromClientExecute(r)
}

/*
RemoveAppTargetRoleFromClient Remove an App Target from a Client

Removes an Application Target by `clientId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiRemoveAppTargetRoleFromClientRequest
*/
func (a *RoleTargetAPIService) RemoveAppTargetRoleFromClient(ctx context.Context, clientId string, roleId string, appName string) ApiRemoveAppTargetRoleFromClientRequest {
	return ApiRemoveAppTargetRoleFromClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) RemoveAppTargetRoleFromClientExecute(r ApiRemoveAppTargetRoleFromClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.RemoveAppTargetRoleFromClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	clientId string
	roleId string
	groupId string
	retryCount int32
}

func (r ApiRemoveGroupTargetRoleFromClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RemoveGroupTargetRoleFromClientExecute(r)
}

/*
RemoveGroupTargetRoleFromClient Delete a Group Target from a Client

Deletes a Group Target from a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId `client_id` of the app
 @param roleId `id` of the Role
 @param groupId The `id` of the group
 @return ApiRemoveGroupTargetRoleFromClientRequest
*/
func (a *RoleTargetAPIService) RemoveGroupTargetRoleFromClient(ctx context.Context, clientId string, roleId string, groupId string) ApiRemoveGroupTargetRoleFromClientRequest {
	return ApiRemoveGroupTargetRoleFromClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
		roleId: roleId,
		groupId: groupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) RemoveGroupTargetRoleFromClientExecute(r ApiRemoveGroupTargetRoleFromClientRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.RemoveGroupTargetRoleFromClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/v1/clients/{clientId}/roles/{roleId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	appName string
	appId string
	retryCount int32
}

func (r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppInstanceTargetFromAdminRoleForUserExecute(r)
}

/*
UnassignAppInstanceTargetFromAdminRoleForUser Unassign an Application Instance Target from an Application Administrator Role

Unassigns an application instance target from an application administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appId Application ID
 @return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest
*/
func (a *RoleTargetAPIService) UnassignAppInstanceTargetFromAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string, appId string) ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest {
	return ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		appName: appName,
		appId: appId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignAppInstanceTargetFromAdminRoleForUserExecute(r ApiUnassignAppInstanceTargetFromAdminRoleForUserRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignAppInstanceTargetFromAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest struct {
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	appName string
	appId string
	retryCount int32
}

func (r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r)
}

/*
UnassignAppInstanceTargetToAppAdminRoleForGroup Unassign an Application Instance Target from an Application Administrator Role

Unassigns an application instance target from application administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @param appId Application ID
 @return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest
*/
func (a *RoleTargetAPIService) UnassignAppInstanceTargetToAppAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string, appId string) ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest {
	return ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		appName: appName,
		appId: appId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignAppInstanceTargetToAppAdminRoleForGroupExecute(r ApiUnassignAppInstanceTargetToAppAdminRoleForGroupRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignAppInstanceTargetToAppAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppTargetFromAppAdminRoleForUserExecute(r)
}

/*
UnassignAppTargetFromAppAdminRoleForUser Unassign an Application Target from an Application Administrator Role

Unassigns an application target from application administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiUnassignAppTargetFromAppAdminRoleForUserRequest
*/
func (a *RoleTargetAPIService) UnassignAppTargetFromAppAdminRoleForUser(ctx context.Context, userId string, roleId string, appName string) ApiUnassignAppTargetFromAppAdminRoleForUserRequest {
	return ApiUnassignAppTargetFromAppAdminRoleForUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignAppTargetFromAppAdminRoleForUserExecute(r ApiUnassignAppTargetFromAppAdminRoleForUserRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignAppTargetFromAppAdminRoleForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	appName string
	retryCount int32
}

func (r ApiUnassignAppTargetToAdminRoleForGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignAppTargetToAdminRoleForGroupExecute(r)
}

/*
UnassignAppTargetToAdminRoleForGroup Unassign an Application Target from Application Administrator Role

Unassigns an application target from application administrator role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param appName Application name for the app type
 @return ApiUnassignAppTargetToAdminRoleForGroupRequest
*/
func (a *RoleTargetAPIService) UnassignAppTargetToAdminRoleForGroup(ctx context.Context, groupId string, roleId string, appName string) ApiUnassignAppTargetToAdminRoleForGroupRequest {
	return ApiUnassignAppTargetToAdminRoleForGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		appName: appName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignAppTargetToAdminRoleForGroupExecute(r ApiUnassignAppTargetToAdminRoleForGroupRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignAppTargetToAdminRoleForGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	groupId string
	roleId string
	targetGroupId string
	retryCount int32
}

func (r ApiUnassignGroupTargetFromGroupAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignGroupTargetFromGroupAdminRoleExecute(r)
}

/*
UnassignGroupTargetFromGroupAdminRole Unassign a Group Target from a Group Role

Unassigns a group target from a group role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The `id` of the group
 @param roleId `id` of the Role
 @param targetGroupId
 @return ApiUnassignGroupTargetFromGroupAdminRoleRequest
*/
func (a *RoleTargetAPIService) UnassignGroupTargetFromGroupAdminRole(ctx context.Context, groupId string, roleId string, targetGroupId string) ApiUnassignGroupTargetFromGroupAdminRoleRequest {
	return ApiUnassignGroupTargetFromGroupAdminRoleRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		roleId: roleId,
		targetGroupId: targetGroupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignGroupTargetFromGroupAdminRoleExecute(r ApiUnassignGroupTargetFromGroupAdminRoleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignGroupTargetFromGroupAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
	ctx context.Context
	ApiService RoleTargetAPI
	userId string
	roleId string
	groupId string
	retryCount int32
}

func (r ApiUnassignGroupTargetFromUserAdminRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignGroupTargetFromUserAdminRoleExecute(r)
}

/*
UnassignGroupTargetFromUserAdminRole Unassign a Group Target from Role

Unassigns a Group Target from Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @param roleId `id` of the Role
 @param groupId The `id` of the group
 @return ApiUnassignGroupTargetFromUserAdminRoleRequest
*/
func (a *RoleTargetAPIService) UnassignGroupTargetFromUserAdminRole(ctx context.Context, userId string, roleId string, groupId string) ApiUnassignGroupTargetFromUserAdminRoleRequest {
	return ApiUnassignGroupTargetFromUserAdminRoleRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		roleId: roleId,
		groupId: groupId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *RoleTargetAPIService) UnassignGroupTargetFromUserAdminRoleExecute(r ApiUnassignGroupTargetFromUserAdminRoleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleTargetAPIService.UnassignGroupTargetFromUserAdminRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterToString(r.roleId, "")), -1)
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

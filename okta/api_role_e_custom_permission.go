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

type RoleECustomPermissionAPI interface {

	/*
		CreateRolePermission Create a custom role permission

		Creates a permission (specified by `permissionType`) for a custom role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param roleIdOrLabel `id` or `label` of the role
		@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
		@return ApiCreateRolePermissionRequest
	*/
	CreateRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiCreateRolePermissionRequest

	// CreateRolePermissionExecute executes the request
	CreateRolePermissionExecute(r ApiCreateRolePermissionRequest) (*APIResponse, error)

	/*
		DeleteRolePermission Delete a custom role permission

		Deletes a permission (identified by `permissionType`) from a custom role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param roleIdOrLabel `id` or `label` of the role
		@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
		@return ApiDeleteRolePermissionRequest
	*/
	DeleteRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiDeleteRolePermissionRequest

	// DeleteRolePermissionExecute executes the request
	DeleteRolePermissionExecute(r ApiDeleteRolePermissionRequest) (*APIResponse, error)

	/*
		GetRolePermission Retrieve a custom role permission

		Retrieves a permission (identified by `permissionType`) for a custom role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param roleIdOrLabel `id` or `label` of the role
		@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
		@return ApiGetRolePermissionRequest
	*/
	GetRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiGetRolePermissionRequest

	// GetRolePermissionExecute executes the request
	//  @return Permission
	GetRolePermissionExecute(r ApiGetRolePermissionRequest) (*Permission, *APIResponse, error)

	/*
		ListRolePermissions List all custom role permissions

		Lists all permissions for a custom role by `roleIdOrLabel`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param roleIdOrLabel `id` or `label` of the role
		@return ApiListRolePermissionsRequest
	*/
	ListRolePermissions(ctx context.Context, roleIdOrLabel string) ApiListRolePermissionsRequest

	// ListRolePermissionsExecute executes the request
	//  @return Permissions
	ListRolePermissionsExecute(r ApiListRolePermissionsRequest) (*Permissions, *APIResponse, error)

	/*
		ReplaceRolePermission Replace a custom role permission

		Replaces a permission (specified by `permissionType`) for a custom role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param roleIdOrLabel `id` or `label` of the role
		@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
		@return ApiReplaceRolePermissionRequest
	*/
	ReplaceRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiReplaceRolePermissionRequest

	// ReplaceRolePermissionExecute executes the request
	//  @return Permission
	ReplaceRolePermissionExecute(r ApiReplaceRolePermissionRequest) (*Permission, *APIResponse, error)
}

// RoleECustomPermissionAPIService RoleECustomPermissionAPI service
type RoleECustomPermissionAPIService service

type ApiCreateRolePermissionRequest struct {
	ctx            context.Context
	ApiService     RoleECustomPermissionAPI
	roleIdOrLabel  string
	permissionType string
	instance       *CreateUpdateIamRolePermissionRequest
	retryCount     int32
}

func (r ApiCreateRolePermissionRequest) Instance(instance CreateUpdateIamRolePermissionRequest) ApiCreateRolePermissionRequest {
	r.instance = &instance
	return r
}

func (r ApiCreateRolePermissionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateRolePermissionExecute(r)
}

/*
CreateRolePermission Create a custom role permission

Creates a permission (specified by `permissionType`) for a custom role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleIdOrLabel `id` or `label` of the role
	@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
	@return ApiCreateRolePermissionRequest
*/
func (a *RoleECustomPermissionAPIService) CreateRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiCreateRolePermissionRequest {
	return ApiCreateRolePermissionRequest{
		ApiService:     a,
		ctx:            ctx,
		roleIdOrLabel:  roleIdOrLabel,
		permissionType: permissionType,
		retryCount:     0,
	}
}

// Execute executes the request
func (a *RoleECustomPermissionAPIService) CreateRolePermissionExecute(r ApiCreateRolePermissionRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleECustomPermissionAPIService.CreateRolePermission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionType"+"}", url.PathEscape(parameterToString(r.permissionType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.instance
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiDeleteRolePermissionRequest struct {
	ctx            context.Context
	ApiService     RoleECustomPermissionAPI
	roleIdOrLabel  string
	permissionType string
	retryCount     int32
}

func (r ApiDeleteRolePermissionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteRolePermissionExecute(r)
}

/*
DeleteRolePermission Delete a custom role permission

Deletes a permission (identified by `permissionType`) from a custom role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleIdOrLabel `id` or `label` of the role
	@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
	@return ApiDeleteRolePermissionRequest
*/
func (a *RoleECustomPermissionAPIService) DeleteRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiDeleteRolePermissionRequest {
	return ApiDeleteRolePermissionRequest{
		ApiService:     a,
		ctx:            ctx,
		roleIdOrLabel:  roleIdOrLabel,
		permissionType: permissionType,
		retryCount:     0,
	}
}

// Execute executes the request
func (a *RoleECustomPermissionAPIService) DeleteRolePermissionExecute(r ApiDeleteRolePermissionRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleECustomPermissionAPIService.DeleteRolePermission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionType"+"}", url.PathEscape(parameterToString(r.permissionType, "")), -1)

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

type ApiGetRolePermissionRequest struct {
	ctx            context.Context
	ApiService     RoleECustomPermissionAPI
	roleIdOrLabel  string
	permissionType string
	retryCount     int32
}

func (r ApiGetRolePermissionRequest) Execute() (*Permission, *APIResponse, error) {
	return r.ApiService.GetRolePermissionExecute(r)
}

/*
GetRolePermission Retrieve a custom role permission

Retrieves a permission (identified by `permissionType`) for a custom role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleIdOrLabel `id` or `label` of the role
	@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
	@return ApiGetRolePermissionRequest
*/
func (a *RoleECustomPermissionAPIService) GetRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiGetRolePermissionRequest {
	return ApiGetRolePermissionRequest{
		ApiService:     a,
		ctx:            ctx,
		roleIdOrLabel:  roleIdOrLabel,
		permissionType: permissionType,
		retryCount:     0,
	}
}

// Execute executes the request
//
//	@return Permission
func (a *RoleECustomPermissionAPIService) GetRolePermissionExecute(r ApiGetRolePermissionRequest) (*Permission, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Permission
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleECustomPermissionAPIService.GetRolePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionType"+"}", url.PathEscape(parameterToString(r.permissionType, "")), -1)

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

type ApiListRolePermissionsRequest struct {
	ctx           context.Context
	ApiService    RoleECustomPermissionAPI
	roleIdOrLabel string
	retryCount    int32
}

func (r ApiListRolePermissionsRequest) Execute() (*Permissions, *APIResponse, error) {
	return r.ApiService.ListRolePermissionsExecute(r)
}

/*
ListRolePermissions List all custom role permissions

Lists all permissions for a custom role by `roleIdOrLabel`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleIdOrLabel `id` or `label` of the role
	@return ApiListRolePermissionsRequest
*/
func (a *RoleECustomPermissionAPIService) ListRolePermissions(ctx context.Context, roleIdOrLabel string) ApiListRolePermissionsRequest {
	return ApiListRolePermissionsRequest{
		ApiService:    a,
		ctx:           ctx,
		roleIdOrLabel: roleIdOrLabel,
		retryCount:    0,
	}
}

// Execute executes the request
//
//	@return Permissions
func (a *RoleECustomPermissionAPIService) ListRolePermissionsExecute(r ApiListRolePermissionsRequest) (*Permissions, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Permissions
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleECustomPermissionAPIService.ListRolePermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/roles/{roleIdOrLabel}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)

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

type ApiReplaceRolePermissionRequest struct {
	ctx            context.Context
	ApiService     RoleECustomPermissionAPI
	roleIdOrLabel  string
	permissionType string
	instance       *CreateUpdateIamRolePermissionRequest
	retryCount     int32
}

func (r ApiReplaceRolePermissionRequest) Instance(instance CreateUpdateIamRolePermissionRequest) ApiReplaceRolePermissionRequest {
	r.instance = &instance
	return r
}

func (r ApiReplaceRolePermissionRequest) Execute() (*Permission, *APIResponse, error) {
	return r.ApiService.ReplaceRolePermissionExecute(r)
}

/*
ReplaceRolePermission Replace a custom role permission

Replaces a permission (specified by `permissionType`) for a custom role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleIdOrLabel `id` or `label` of the role
	@param permissionType An Okta [permission](/openapi/okta-management/guides/permissions)
	@return ApiReplaceRolePermissionRequest
*/
func (a *RoleECustomPermissionAPIService) ReplaceRolePermission(ctx context.Context, roleIdOrLabel string, permissionType string) ApiReplaceRolePermissionRequest {
	return ApiReplaceRolePermissionRequest{
		ApiService:     a,
		ctx:            ctx,
		roleIdOrLabel:  roleIdOrLabel,
		permissionType: permissionType,
		retryCount:     0,
	}
}

// Execute executes the request
//
//	@return Permission
func (a *RoleECustomPermissionAPIService) ReplaceRolePermissionExecute(r ApiReplaceRolePermissionRequest) (*Permission, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Permission
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleECustomPermissionAPIService.ReplaceRolePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/roles/{roleIdOrLabel}/permissions/{permissionType}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permissionType"+"}", url.PathEscape(parameterToString(r.permissionType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.instance
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

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

type RoleDResourceSetBindingMemberAPI interface {

	/*
		AddMembersToBinding Add more role resource set binding members

		Adds more members to a role resource set binding

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSetIdOrLabel `id` or `label` of the resource set
		@param roleIdOrLabel `id` or `label` of the role
		@return ApiAddMembersToBindingRequest
	*/
	AddMembersToBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string) ApiAddMembersToBindingRequest

	// AddMembersToBindingExecute executes the request
	//  @return ResourceSetBindingEditResponse
	AddMembersToBindingExecute(r ApiAddMembersToBindingRequest) (*ResourceSetBindingEditResponse, *APIResponse, error)

	/*
		GetMemberOfBinding Retrieve a role resource set binding member

		Retrieves a member (identified by `memberId`) that belongs to a role resource set binding

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSetIdOrLabel `id` or `label` of the resource set
		@param roleIdOrLabel `id` or `label` of the role
		@param memberId `id` of the member
		@return ApiGetMemberOfBindingRequest
	*/
	GetMemberOfBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string, memberId string) ApiGetMemberOfBindingRequest

	// GetMemberOfBindingExecute executes the request
	//  @return ResourceSetBindingMember
	GetMemberOfBindingExecute(r ApiGetMemberOfBindingRequest) (*ResourceSetBindingMember, *APIResponse, error)

	/*
		ListMembersOfBinding List all role resource set binding members

		Lists all members of a role resource set binding with pagination support

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSetIdOrLabel `id` or `label` of the resource set
		@param roleIdOrLabel `id` or `label` of the role
		@return ApiListMembersOfBindingRequest
	*/
	ListMembersOfBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string) ApiListMembersOfBindingRequest

	// ListMembersOfBindingExecute executes the request
	//  @return ResourceSetBindingMembers
	ListMembersOfBindingExecute(r ApiListMembersOfBindingRequest) (*ResourceSetBindingMembers, *APIResponse, error)

	/*
		UnassignMemberFromBinding Unassign a role resource set binding member

		Unassigns a member (identified by `memberId`) from a role resource set binding

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param resourceSetIdOrLabel `id` or `label` of the resource set
		@param roleIdOrLabel `id` or `label` of the role
		@param memberId `id` of the member
		@return ApiUnassignMemberFromBindingRequest
	*/
	UnassignMemberFromBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string, memberId string) ApiUnassignMemberFromBindingRequest

	// UnassignMemberFromBindingExecute executes the request
	UnassignMemberFromBindingExecute(r ApiUnassignMemberFromBindingRequest) (*APIResponse, error)
}

// RoleDResourceSetBindingMemberAPIService RoleDResourceSetBindingMemberAPI service
type RoleDResourceSetBindingMemberAPIService service

type ApiAddMembersToBindingRequest struct {
	ctx                  context.Context
	ApiService           RoleDResourceSetBindingMemberAPI
	resourceSetIdOrLabel string
	roleIdOrLabel        string
	instance             *ResourceSetBindingAddMembersRequest
	retryCount           int32
}

func (r ApiAddMembersToBindingRequest) Instance(instance ResourceSetBindingAddMembersRequest) ApiAddMembersToBindingRequest {
	r.instance = &instance
	return r
}

func (r ApiAddMembersToBindingRequest) Execute() (*ResourceSetBindingEditResponse, *APIResponse, error) {
	return r.ApiService.AddMembersToBindingExecute(r)
}

/*
AddMembersToBinding Add more role resource set binding members

Adds more members to a role resource set binding

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceSetIdOrLabel `id` or `label` of the resource set
	@param roleIdOrLabel `id` or `label` of the role
	@return ApiAddMembersToBindingRequest
*/
func (a *RoleDResourceSetBindingMemberAPIService) AddMembersToBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string) ApiAddMembersToBindingRequest {
	return ApiAddMembersToBindingRequest{
		ApiService:           a,
		ctx:                  ctx,
		resourceSetIdOrLabel: resourceSetIdOrLabel,
		roleIdOrLabel:        roleIdOrLabel,
		retryCount:           0,
	}
}

// Execute executes the request
//
//	@return ResourceSetBindingEditResponse
func (a *RoleDResourceSetBindingMemberAPIService) AddMembersToBindingExecute(r ApiAddMembersToBindingRequest) (*ResourceSetBindingEditResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceSetBindingEditResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleDResourceSetBindingMemberAPIService.AddMembersToBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSetIdOrLabel"+"}", url.PathEscape(parameterToString(r.resourceSetIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.instance == nil {
		return localVarReturnValue, nil, reportError("instance is required and must be specified")
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

type ApiGetMemberOfBindingRequest struct {
	ctx                  context.Context
	ApiService           RoleDResourceSetBindingMemberAPI
	resourceSetIdOrLabel string
	roleIdOrLabel        string
	memberId             string
	retryCount           int32
}

func (r ApiGetMemberOfBindingRequest) Execute() (*ResourceSetBindingMember, *APIResponse, error) {
	return r.ApiService.GetMemberOfBindingExecute(r)
}

/*
GetMemberOfBinding Retrieve a role resource set binding member

Retrieves a member (identified by `memberId`) that belongs to a role resource set binding

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceSetIdOrLabel `id` or `label` of the resource set
	@param roleIdOrLabel `id` or `label` of the role
	@param memberId `id` of the member
	@return ApiGetMemberOfBindingRequest
*/
func (a *RoleDResourceSetBindingMemberAPIService) GetMemberOfBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string, memberId string) ApiGetMemberOfBindingRequest {
	return ApiGetMemberOfBindingRequest{
		ApiService:           a,
		ctx:                  ctx,
		resourceSetIdOrLabel: resourceSetIdOrLabel,
		roleIdOrLabel:        roleIdOrLabel,
		memberId:             memberId,
		retryCount:           0,
	}
}

// Execute executes the request
//
//	@return ResourceSetBindingMember
func (a *RoleDResourceSetBindingMemberAPIService) GetMemberOfBindingExecute(r ApiGetMemberOfBindingRequest) (*ResourceSetBindingMember, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceSetBindingMember
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleDResourceSetBindingMemberAPIService.GetMemberOfBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members/{memberId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSetIdOrLabel"+"}", url.PathEscape(parameterToString(r.resourceSetIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberId"+"}", url.PathEscape(parameterToString(r.memberId, "")), -1)

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

type ApiListMembersOfBindingRequest struct {
	ctx                  context.Context
	ApiService           RoleDResourceSetBindingMemberAPI
	resourceSetIdOrLabel string
	roleIdOrLabel        string
	after                *string
	retryCount           int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header).
func (r ApiListMembersOfBindingRequest) After(after string) ApiListMembersOfBindingRequest {
	r.after = &after
	return r
}

func (r ApiListMembersOfBindingRequest) Execute() (*ResourceSetBindingMembers, *APIResponse, error) {
	return r.ApiService.ListMembersOfBindingExecute(r)
}

/*
ListMembersOfBinding List all role resource set binding members

Lists all members of a role resource set binding with pagination support

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceSetIdOrLabel `id` or `label` of the resource set
	@param roleIdOrLabel `id` or `label` of the role
	@return ApiListMembersOfBindingRequest
*/
func (a *RoleDResourceSetBindingMemberAPIService) ListMembersOfBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string) ApiListMembersOfBindingRequest {
	return ApiListMembersOfBindingRequest{
		ApiService:           a,
		ctx:                  ctx,
		resourceSetIdOrLabel: resourceSetIdOrLabel,
		roleIdOrLabel:        roleIdOrLabel,
		retryCount:           0,
	}
}

// Execute executes the request
//
//	@return ResourceSetBindingMembers
func (a *RoleDResourceSetBindingMemberAPIService) ListMembersOfBindingExecute(r ApiListMembersOfBindingRequest) (*ResourceSetBindingMembers, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourceSetBindingMembers
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleDResourceSetBindingMemberAPIService.ListMembersOfBinding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSetIdOrLabel"+"}", url.PathEscape(parameterToString(r.resourceSetIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
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

type ApiUnassignMemberFromBindingRequest struct {
	ctx                  context.Context
	ApiService           RoleDResourceSetBindingMemberAPI
	resourceSetIdOrLabel string
	roleIdOrLabel        string
	memberId             string
	retryCount           int32
}

func (r ApiUnassignMemberFromBindingRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnassignMemberFromBindingExecute(r)
}

/*
UnassignMemberFromBinding Unassign a role resource set binding member

Unassigns a member (identified by `memberId`) from a role resource set binding

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceSetIdOrLabel `id` or `label` of the resource set
	@param roleIdOrLabel `id` or `label` of the role
	@param memberId `id` of the member
	@return ApiUnassignMemberFromBindingRequest
*/
func (a *RoleDResourceSetBindingMemberAPIService) UnassignMemberFromBinding(ctx context.Context, resourceSetIdOrLabel string, roleIdOrLabel string, memberId string) ApiUnassignMemberFromBindingRequest {
	return ApiUnassignMemberFromBindingRequest{
		ApiService:           a,
		ctx:                  ctx,
		resourceSetIdOrLabel: resourceSetIdOrLabel,
		roleIdOrLabel:        roleIdOrLabel,
		memberId:             memberId,
		retryCount:           0,
	}
}

// Execute executes the request
func (a *RoleDResourceSetBindingMemberAPIService) UnassignMemberFromBindingExecute(r ApiUnassignMemberFromBindingRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleDResourceSetBindingMemberAPIService.UnassignMemberFromBinding")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/iam/resource-sets/{resourceSetIdOrLabel}/bindings/{roleIdOrLabel}/members/{memberId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSetIdOrLabel"+"}", url.PathEscape(parameterToString(r.resourceSetIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleIdOrLabel"+"}", url.PathEscape(parameterToString(r.roleIdOrLabel, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberId"+"}", url.PathEscape(parameterToString(r.memberId, "")), -1)

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

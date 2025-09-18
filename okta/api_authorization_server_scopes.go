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

type AuthorizationServerScopesAPI interface {

	/*
		CreateOAuth2Scope Create a custom token scope

		Creates a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@return ApiCreateOAuth2ScopeRequest
	*/
	CreateOAuth2Scope(ctx context.Context, authServerId string) ApiCreateOAuth2ScopeRequest

	// CreateOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	CreateOAuth2ScopeExecute(r ApiCreateOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error)

	/*
		DeleteOAuth2Scope Delete a custom token scope

		Deletes a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiDeleteOAuth2ScopeRequest
	*/
	DeleteOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiDeleteOAuth2ScopeRequest

	// DeleteOAuth2ScopeExecute executes the request
	DeleteOAuth2ScopeExecute(r ApiDeleteOAuth2ScopeRequest) (*APIResponse, error)

	/*
		GetOAuth2Scope Retrieve a custom token scope

		Retrieves a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiGetOAuth2ScopeRequest
	*/
	GetOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiGetOAuth2ScopeRequest

	// GetOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	GetOAuth2ScopeExecute(r ApiGetOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error)

	/*
		ListOAuth2Scopes List all custom token scopes

		Lists all custom token scopes

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@return ApiListOAuth2ScopesRequest
	*/
	ListOAuth2Scopes(ctx context.Context, authServerId string) ApiListOAuth2ScopesRequest

	// ListOAuth2ScopesExecute executes the request
	//  @return []OAuth2Scope
	ListOAuth2ScopesExecute(r ApiListOAuth2ScopesRequest) ([]OAuth2Scope, *APIResponse, error)

	/*
		ReplaceOAuth2Scope Replace a custom token scope

		Replaces a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiReplaceOAuth2ScopeRequest
	*/
	ReplaceOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiReplaceOAuth2ScopeRequest

	// ReplaceOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	ReplaceOAuth2ScopeExecute(r ApiReplaceOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error)
}

// AuthorizationServerScopesAPIService AuthorizationServerScopesAPI service
type AuthorizationServerScopesAPIService service

type ApiCreateOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	oAuth2Scope  *OAuth2Scope
	retryCount   int32
}

func (r ApiCreateOAuth2ScopeRequest) OAuth2Scope(oAuth2Scope OAuth2Scope) ApiCreateOAuth2ScopeRequest {
	r.oAuth2Scope = &oAuth2Scope
	return r
}

func (r ApiCreateOAuth2ScopeRequest) Execute() (*OAuth2Scope, *APIResponse, error) {
	return r.ApiService.CreateOAuth2ScopeExecute(r)
}

/*
CreateOAuth2Scope Create a custom token scope

Creates a custom token scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@return ApiCreateOAuth2ScopeRequest
*/
func (a *AuthorizationServerScopesAPIService) CreateOAuth2Scope(ctx context.Context, authServerId string) ApiCreateOAuth2ScopeRequest {
	return ApiCreateOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		retryCount:   0,
	}
}

// Execute executes the request
//
//	@return OAuth2Scope
func (a *AuthorizationServerScopesAPIService) CreateOAuth2ScopeExecute(r ApiCreateOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Scope
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.CreateOAuth2Scope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuth2Scope == nil {
		return localVarReturnValue, nil, reportError("oAuth2Scope is required and must be specified")
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
	localVarPostBody = r.oAuth2Scope
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

type ApiDeleteOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	retryCount   int32
}

func (r ApiDeleteOAuth2ScopeRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteOAuth2ScopeExecute(r)
}

/*
DeleteOAuth2Scope Delete a custom token scope

Deletes a custom token scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param scopeId `id` of Scope
	@return ApiDeleteOAuth2ScopeRequest
*/
func (a *AuthorizationServerScopesAPIService) DeleteOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiDeleteOAuth2ScopeRequest {
	return ApiDeleteOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request
func (a *AuthorizationServerScopesAPIService) DeleteOAuth2ScopeExecute(r ApiDeleteOAuth2ScopeRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.DeleteOAuth2Scope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

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

type ApiGetOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	retryCount   int32
}

func (r ApiGetOAuth2ScopeRequest) Execute() (*OAuth2Scope, *APIResponse, error) {
	return r.ApiService.GetOAuth2ScopeExecute(r)
}

/*
GetOAuth2Scope Retrieve a custom token scope

Retrieves a custom token scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param scopeId `id` of Scope
	@return ApiGetOAuth2ScopeRequest
*/
func (a *AuthorizationServerScopesAPIService) GetOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiGetOAuth2ScopeRequest {
	return ApiGetOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request
//
//	@return OAuth2Scope
func (a *AuthorizationServerScopesAPIService) GetOAuth2ScopeExecute(r ApiGetOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Scope
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.GetOAuth2Scope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

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

type ApiListOAuth2ScopesRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	q            *string
	filter       *string
	after        *string
	limit        *int32
	retryCount   int32
}

// Searches the &#x60;name&#x60; of Custom Token Scopes for matching values
func (r ApiListOAuth2ScopesRequest) Q(q string) ApiListOAuth2ScopesRequest {
	r.q = &q
	return r
}

// Filter expression for Custom Token Scopes
func (r ApiListOAuth2ScopesRequest) Filter(filter string) ApiListOAuth2ScopesRequest {
	r.filter = &filter
	return r
}

// Specifies the pagination cursor for the next page of scopes. Treat the after cursor as an opaque value and obtain it through the next link relationship. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListOAuth2ScopesRequest) After(after string) ApiListOAuth2ScopesRequest {
	r.after = &after
	return r
}

// Specifies the number of objects to return per page. If there are multiple pages of results, the Link header contains a &#x60;next&#x60; link that you need to use as an opaque value (follow it, don&#39;t parse it). See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListOAuth2ScopesRequest) Limit(limit int32) ApiListOAuth2ScopesRequest {
	r.limit = &limit
	return r
}

func (r ApiListOAuth2ScopesRequest) Execute() ([]OAuth2Scope, *APIResponse, error) {
	return r.ApiService.ListOAuth2ScopesExecute(r)
}

/*
ListOAuth2Scopes List all custom token scopes

Lists all custom token scopes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@return ApiListOAuth2ScopesRequest
*/
func (a *AuthorizationServerScopesAPIService) ListOAuth2Scopes(ctx context.Context, authServerId string) ApiListOAuth2ScopesRequest {
	return ApiListOAuth2ScopesRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		retryCount:   0,
	}
}

// Execute executes the request
//
//	@return []OAuth2Scope
func (a *AuthorizationServerScopesAPIService) ListOAuth2ScopesExecute(r ApiListOAuth2ScopesRequest) ([]OAuth2Scope, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OAuth2Scope
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.ListOAuth2Scopes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
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

type ApiReplaceOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	oAuth2Scope  *OAuth2Scope
	retryCount   int32
}

func (r ApiReplaceOAuth2ScopeRequest) OAuth2Scope(oAuth2Scope OAuth2Scope) ApiReplaceOAuth2ScopeRequest {
	r.oAuth2Scope = &oAuth2Scope
	return r
}

func (r ApiReplaceOAuth2ScopeRequest) Execute() (*OAuth2Scope, *APIResponse, error) {
	return r.ApiService.ReplaceOAuth2ScopeExecute(r)
}

/*
ReplaceOAuth2Scope Replace a custom token scope

Replaces a custom token scope

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param scopeId `id` of Scope
	@return ApiReplaceOAuth2ScopeRequest
*/
func (a *AuthorizationServerScopesAPIService) ReplaceOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiReplaceOAuth2ScopeRequest {
	return ApiReplaceOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request
//
//	@return OAuth2Scope
func (a *AuthorizationServerScopesAPIService) ReplaceOAuth2ScopeExecute(r ApiReplaceOAuth2ScopeRequest) (*OAuth2Scope, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2Scope
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.ReplaceOAuth2Scope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuth2Scope == nil {
		return localVarReturnValue, nil, reportError("oAuth2Scope is required and must be specified")
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
	localVarPostBody = r.oAuth2Scope
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

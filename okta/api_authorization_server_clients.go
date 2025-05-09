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


type AuthorizationServerClientsAPI interface {

	/*
	GetRefreshTokenForAuthorizationServerAndClient Retrieve a refresh token for a Client

	Retrieves a refresh token for a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param clientId `client_id` of the app
	@param tokenId `id` of Token
	@return ApiGetRefreshTokenForAuthorizationServerAndClientRequest
	*/
	GetRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string, tokenId string) ApiGetRefreshTokenForAuthorizationServerAndClientRequest

	// GetRefreshTokenForAuthorizationServerAndClientExecute executes the request
	//  @return OAuth2RefreshToken
	GetRefreshTokenForAuthorizationServerAndClientExecute(r ApiGetRefreshTokenForAuthorizationServerAndClientRequest) (*OAuth2RefreshToken, *APIResponse, error)

	/*
	ListOAuth2ClientsForAuthorizationServer List all Client resources for an authorization server

	Lists all Client resources for which the specified authorization server has tokens

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@return ApiListOAuth2ClientsForAuthorizationServerRequest
	*/
	ListOAuth2ClientsForAuthorizationServer(ctx context.Context, authServerId string) ApiListOAuth2ClientsForAuthorizationServerRequest

	// ListOAuth2ClientsForAuthorizationServerExecute executes the request
	//  @return []OAuth2Client
	ListOAuth2ClientsForAuthorizationServerExecute(r ApiListOAuth2ClientsForAuthorizationServerRequest) ([]OAuth2Client, *APIResponse, error)

	/*
	ListRefreshTokensForAuthorizationServerAndClient List all refresh tokens for a Client

	Lists all refresh tokens issued by an authorization server for a specific Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param clientId `client_id` of the app
	@return ApiListRefreshTokensForAuthorizationServerAndClientRequest
	*/
	ListRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string) ApiListRefreshTokensForAuthorizationServerAndClientRequest

	// ListRefreshTokensForAuthorizationServerAndClientExecute executes the request
	//  @return []OAuth2RefreshToken
	ListRefreshTokensForAuthorizationServerAndClientExecute(r ApiListRefreshTokensForAuthorizationServerAndClientRequest) ([]OAuth2RefreshToken, *APIResponse, error)

	/*
	RevokeRefreshTokenForAuthorizationServerAndClient Revoke a refresh token for a Client

	Revokes a refresh token for a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param clientId `client_id` of the app
	@param tokenId `id` of Token
	@return ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest
	*/
	RevokeRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string, tokenId string) ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest

	// RevokeRefreshTokenForAuthorizationServerAndClientExecute executes the request
	RevokeRefreshTokenForAuthorizationServerAndClientExecute(r ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest) (*APIResponse, error)

	/*
	RevokeRefreshTokensForAuthorizationServerAndClient Revoke all refresh tokens for a Client

	Revokes all refresh tokens for a Client

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param clientId `client_id` of the app
	@return ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest
	*/
	RevokeRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string) ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest

	// RevokeRefreshTokensForAuthorizationServerAndClientExecute executes the request
	RevokeRefreshTokensForAuthorizationServerAndClientExecute(r ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest) (*APIResponse, error)
}

// AuthorizationServerClientsAPIService AuthorizationServerClientsAPI service
type AuthorizationServerClientsAPIService service

type ApiGetRefreshTokenForAuthorizationServerAndClientRequest struct {
	ctx context.Context
	ApiService AuthorizationServerClientsAPI
	authServerId string
	clientId string
	tokenId string
	expand *string
	retryCount int32
}

// Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute.
func (r ApiGetRefreshTokenForAuthorizationServerAndClientRequest) Expand(expand string) ApiGetRefreshTokenForAuthorizationServerAndClientRequest {
	r.expand = &expand
	return r
}

func (r ApiGetRefreshTokenForAuthorizationServerAndClientRequest) Execute() (*OAuth2RefreshToken, *APIResponse, error) {
	return r.ApiService.GetRefreshTokenForAuthorizationServerAndClientExecute(r)
}

/*
GetRefreshTokenForAuthorizationServerAndClient Retrieve a refresh token for a Client

Retrieves a refresh token for a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param clientId `client_id` of the app
 @param tokenId `id` of Token
 @return ApiGetRefreshTokenForAuthorizationServerAndClientRequest
*/
func (a *AuthorizationServerClientsAPIService) GetRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string, tokenId string) ApiGetRefreshTokenForAuthorizationServerAndClientRequest {
	return ApiGetRefreshTokenForAuthorizationServerAndClientRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		clientId: clientId,
		tokenId: tokenId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return OAuth2RefreshToken
func (a *AuthorizationServerClientsAPIService) GetRefreshTokenForAuthorizationServerAndClientExecute(r ApiGetRefreshTokenForAuthorizationServerAndClientRequest) (*OAuth2RefreshToken, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2RefreshToken
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerClientsAPIService.GetRefreshTokenForAuthorizationServerAndClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", url.PathEscape(parameterToString(r.tokenId, "")), -1)

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

type ApiListOAuth2ClientsForAuthorizationServerRequest struct {
	ctx context.Context
	ApiService AuthorizationServerClientsAPI
	authServerId string
	retryCount int32
}

func (r ApiListOAuth2ClientsForAuthorizationServerRequest) Execute() ([]OAuth2Client, *APIResponse, error) {
	return r.ApiService.ListOAuth2ClientsForAuthorizationServerExecute(r)
}

/*
ListOAuth2ClientsForAuthorizationServer List all Client resources for an authorization server

Lists all Client resources for which the specified authorization server has tokens

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @return ApiListOAuth2ClientsForAuthorizationServerRequest
*/
func (a *AuthorizationServerClientsAPIService) ListOAuth2ClientsForAuthorizationServer(ctx context.Context, authServerId string) ApiListOAuth2ClientsForAuthorizationServerRequest {
	return ApiListOAuth2ClientsForAuthorizationServerRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []OAuth2Client
func (a *AuthorizationServerClientsAPIService) ListOAuth2ClientsForAuthorizationServerExecute(r ApiListOAuth2ClientsForAuthorizationServerRequest) ([]OAuth2Client, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OAuth2Client
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerClientsAPIService.ListOAuth2ClientsForAuthorizationServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/clients"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

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

type ApiListRefreshTokensForAuthorizationServerAndClientRequest struct {
	ctx context.Context
	ApiService AuthorizationServerClientsAPI
	authServerId string
	clientId string
	expand *string
	after *string
	limit *int32
	retryCount int32
}

// Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute.
func (r ApiListRefreshTokensForAuthorizationServerAndClientRequest) Expand(expand string) ApiListRefreshTokensForAuthorizationServerAndClientRequest {
	r.expand = &expand
	return r
}

// Specifies the pagination cursor for the next page of tokens
func (r ApiListRefreshTokensForAuthorizationServerAndClientRequest) After(after string) ApiListRefreshTokensForAuthorizationServerAndClientRequest {
	r.after = &after
	return r
}

// The maximum number of tokens to return (maximum 200)
func (r ApiListRefreshTokensForAuthorizationServerAndClientRequest) Limit(limit int32) ApiListRefreshTokensForAuthorizationServerAndClientRequest {
	r.limit = &limit
	return r
}

func (r ApiListRefreshTokensForAuthorizationServerAndClientRequest) Execute() ([]OAuth2RefreshToken, *APIResponse, error) {
	return r.ApiService.ListRefreshTokensForAuthorizationServerAndClientExecute(r)
}

/*
ListRefreshTokensForAuthorizationServerAndClient List all refresh tokens for a Client

Lists all refresh tokens issued by an authorization server for a specific Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param clientId `client_id` of the app
 @return ApiListRefreshTokensForAuthorizationServerAndClientRequest
*/
func (a *AuthorizationServerClientsAPIService) ListRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string) ApiListRefreshTokensForAuthorizationServerAndClientRequest {
	return ApiListRefreshTokensForAuthorizationServerAndClientRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		clientId: clientId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []OAuth2RefreshToken
func (a *AuthorizationServerClientsAPIService) ListRefreshTokensForAuthorizationServerAndClientExecute(r ApiListRefreshTokensForAuthorizationServerAndClientRequest) ([]OAuth2RefreshToken, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OAuth2RefreshToken
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerClientsAPIService.ListRefreshTokensForAuthorizationServerAndClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
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

type ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest struct {
	ctx context.Context
	ApiService AuthorizationServerClientsAPI
	authServerId string
	clientId string
	tokenId string
	retryCount int32
}

func (r ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RevokeRefreshTokenForAuthorizationServerAndClientExecute(r)
}

/*
RevokeRefreshTokenForAuthorizationServerAndClient Revoke a refresh token for a Client

Revokes a refresh token for a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param clientId `client_id` of the app
 @param tokenId `id` of Token
 @return ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest
*/
func (a *AuthorizationServerClientsAPIService) RevokeRefreshTokenForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string, tokenId string) ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest {
	return ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		clientId: clientId,
		tokenId: tokenId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *AuthorizationServerClientsAPIService) RevokeRefreshTokenForAuthorizationServerAndClientExecute(r ApiRevokeRefreshTokenForAuthorizationServerAndClientRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerClientsAPIService.RevokeRefreshTokenForAuthorizationServerAndClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tokenId"+"}", url.PathEscape(parameterToString(r.tokenId, "")), -1)

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

type ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest struct {
	ctx context.Context
	ApiService AuthorizationServerClientsAPI
	authServerId string
	clientId string
	retryCount int32
}

func (r ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest) Execute() (*APIResponse, error) {
	return r.ApiService.RevokeRefreshTokensForAuthorizationServerAndClientExecute(r)
}

/*
RevokeRefreshTokensForAuthorizationServerAndClient Revoke all refresh tokens for a Client

Revokes all refresh tokens for a Client

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param clientId `client_id` of the app
 @return ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest
*/
func (a *AuthorizationServerClientsAPIService) RevokeRefreshTokensForAuthorizationServerAndClient(ctx context.Context, authServerId string, clientId string) ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest {
	return ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		clientId: clientId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *AuthorizationServerClientsAPIService) RevokeRefreshTokensForAuthorizationServerAndClientExecute(r ApiRevokeRefreshTokensForAuthorizationServerAndClientRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerClientsAPIService.RevokeRefreshTokensForAuthorizationServerAndClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

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

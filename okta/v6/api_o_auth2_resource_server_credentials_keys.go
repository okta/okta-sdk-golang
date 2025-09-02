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

API version: 5.1.0
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


type OAuth2ResourceServerCredentialsKeysAPI interface {

	/*
	ActivateOAuth2ResourceServerJsonWebKey Activate a Custom Authorization Server Public JSON Web Key

	Activates a custom authorization server public JSON web key by key `id`.
> **Note:** You can have only one active key at any given time for the authorization server. When you activate an inactive key, Okta automatically deactivates the current active key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param keyId Unique `id` of the Custom Authorization Server JSON Web Key
	@return ApiActivateOAuth2ResourceServerJsonWebKeyRequest
	*/
	ActivateOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiActivateOAuth2ResourceServerJsonWebKeyRequest

	// ActivateOAuth2ResourceServerJsonWebKeyExecute executes the request
	//  @return OAuth2ResourceServerJsonWebKey
	ActivateOAuth2ResourceServerJsonWebKeyExecute(r ApiActivateOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error)

	/*
	AddOAuth2ResourceServerJsonWebKey Add a JSON Web Key

	Adds a new JSON Web Key to the custom authorization server`s JSON web keys.
> **Note:** This API doesn't allow you to add a key if the existing key doesn't have a `kid`. Use the [Replace an Authorization Server](/openapi/okta-management/management/tag/AuthorizationServer/#tag/AuthorizationServer/operation/replaceAuthorizationServer) operation to update the JWKS or [Delete a Custom Authorization Server Public JSON Web Key](/openapi/okta-management/management/tag/OAuth2ResourceServerCredentialsKeys/#tag/OAuth2ResourceServerCredentialsKeys/operation/deleteOAuth2ResourceServerJsonWebKey) and re-add the key with a `kid`.
> **Note:** This API doesn't allow you to add a key with an ACTIVE status. You need to add an INACTIVE key first, and then ACTIVATE the key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@return ApiAddOAuth2ResourceServerJsonWebKeyRequest
	*/
	AddOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string) ApiAddOAuth2ResourceServerJsonWebKeyRequest

	// AddOAuth2ResourceServerJsonWebKeyExecute executes the request
	//  @return OAuth2ResourceServerJsonWebKey
	AddOAuth2ResourceServerJsonWebKeyExecute(r ApiAddOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error)

	/*
	DeactivateOAuth2ResourceServerJsonWebKey Deactivate a Custom Authorization Server Public JSON Web Key

	Deactivates a custom authorization server public JSON web key by key `id`.
> **Note:** Deactivating the active key isn't allowed if the authorization server has access token encryption enabled. You can activate another key, which makes the current key inactive.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param keyId Unique `id` of the Custom Authorization Server JSON Web Key
	@return ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest
	*/
	DeactivateOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest

	// DeactivateOAuth2ResourceServerJsonWebKeyExecute executes the request
	//  @return OAuth2ResourceServerJsonWebKey
	DeactivateOAuth2ResourceServerJsonWebKeyExecute(r ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error)

	/*
	DeleteOAuth2ResourceServerJsonWebKey Delete a Custom Authorization Server Public JSON Web Key

	Deletes a custom authorization server public JSON web key by key `id`. You can only delete an inactive key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param keyId Unique `id` of the Custom Authorization Server JSON Web Key
	@return ApiDeleteOAuth2ResourceServerJsonWebKeyRequest
	*/
	DeleteOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiDeleteOAuth2ResourceServerJsonWebKeyRequest

	// DeleteOAuth2ResourceServerJsonWebKeyExecute executes the request
	DeleteOAuth2ResourceServerJsonWebKeyExecute(r ApiDeleteOAuth2ResourceServerJsonWebKeyRequest) (*APIResponse, error)

	/*
	GetOAuth2ResourceServerJsonWebKey Retrieve a Custom Authorization Server Public JSON Web Key

	Retrieves a custom authorization server public JSON web key by key `id`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param keyId Unique `id` of the Custom Authorization Server JSON Web Key
	@return ApiGetOAuth2ResourceServerJsonWebKeyRequest
	*/
	GetOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiGetOAuth2ResourceServerJsonWebKeyRequest

	// GetOAuth2ResourceServerJsonWebKeyExecute executes the request
	//  @return OAuth2ResourceServerJsonWebKey
	GetOAuth2ResourceServerJsonWebKeyExecute(r ApiGetOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error)

	/*
	ListOAuth2ResourceServerJsonWebKeys List all Custom Authorization Server Public JSON Web Keys

	Lists all the public keys used by the custom authorization server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@return ApiListOAuth2ResourceServerJsonWebKeysRequest
	*/
	ListOAuth2ResourceServerJsonWebKeys(ctx context.Context, authServerId string) ApiListOAuth2ResourceServerJsonWebKeysRequest

	// ListOAuth2ResourceServerJsonWebKeysExecute executes the request
	//  @return []OAuth2ResourceServerJsonWebKey
	ListOAuth2ResourceServerJsonWebKeysExecute(r ApiListOAuth2ResourceServerJsonWebKeysRequest) ([]OAuth2ResourceServerJsonWebKey, *APIResponse, error)
}

// OAuth2ResourceServerCredentialsKeysAPIService OAuth2ResourceServerCredentialsKeysAPI service
type OAuth2ResourceServerCredentialsKeysAPIService service

type ApiActivateOAuth2ResourceServerJsonWebKeyRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	keyId string
	retryCount int32
}

func (r ApiActivateOAuth2ResourceServerJsonWebKeyRequest) Execute() (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	return r.ApiService.ActivateOAuth2ResourceServerJsonWebKeyExecute(r)
}

/*
ActivateOAuth2ResourceServerJsonWebKey Activate a Custom Authorization Server Public JSON Web Key

Activates a custom authorization server public JSON web key by key `id`.
> **Note:** You can have only one active key at any given time for the authorization server. When you activate an inactive key, Okta automatically deactivates the current active key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param keyId Unique `id` of the Custom Authorization Server JSON Web Key
 @return ApiActivateOAuth2ResourceServerJsonWebKeyRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) ActivateOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiActivateOAuth2ResourceServerJsonWebKeyRequest {
	return ApiActivateOAuth2ResourceServerJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		keyId: keyId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return OAuth2ResourceServerJsonWebKey
func (a *OAuth2ResourceServerCredentialsKeysAPIService) ActivateOAuth2ResourceServerJsonWebKeyExecute(r ApiActivateOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2ResourceServerJsonWebKey
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.ActivateOAuth2ResourceServerJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiAddOAuth2ResourceServerJsonWebKeyRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	oAuth2ResourceServerJsonWebKeyRequestBody *OAuth2ResourceServerJsonWebKeyRequestBody
	retryCount int32
}

func (r ApiAddOAuth2ResourceServerJsonWebKeyRequest) OAuth2ResourceServerJsonWebKeyRequestBody(oAuth2ResourceServerJsonWebKeyRequestBody OAuth2ResourceServerJsonWebKeyRequestBody) ApiAddOAuth2ResourceServerJsonWebKeyRequest {
	r.oAuth2ResourceServerJsonWebKeyRequestBody = &oAuth2ResourceServerJsonWebKeyRequestBody
	return r
}

func (r ApiAddOAuth2ResourceServerJsonWebKeyRequest) Execute() (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	return r.ApiService.AddOAuth2ResourceServerJsonWebKeyExecute(r)
}

/*
AddOAuth2ResourceServerJsonWebKey Add a JSON Web Key

Adds a new JSON Web Key to the custom authorization server`s JSON web keys.
> **Note:** This API doesn't allow you to add a key if the existing key doesn't have a `kid`. Use the [Replace an Authorization Server](/openapi/okta-management/management/tag/AuthorizationServer/#tag/AuthorizationServer/operation/replaceAuthorizationServer) operation to update the JWKS or [Delete a Custom Authorization Server Public JSON Web Key](/openapi/okta-management/management/tag/OAuth2ResourceServerCredentialsKeys/#tag/OAuth2ResourceServerCredentialsKeys/operation/deleteOAuth2ResourceServerJsonWebKey) and re-add the key with a `kid`.
> **Note:** This API doesn't allow you to add a key with an ACTIVE status. You need to add an INACTIVE key first, and then ACTIVATE the key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @return ApiAddOAuth2ResourceServerJsonWebKeyRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) AddOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string) ApiAddOAuth2ResourceServerJsonWebKeyRequest {
	return ApiAddOAuth2ResourceServerJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return OAuth2ResourceServerJsonWebKey
func (a *OAuth2ResourceServerCredentialsKeysAPIService) AddOAuth2ResourceServerJsonWebKeyExecute(r ApiAddOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2ResourceServerJsonWebKey
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.AddOAuth2ResourceServerJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuth2ResourceServerJsonWebKeyRequestBody == nil {
		return localVarReturnValue, nil, reportError("oAuth2ResourceServerJsonWebKeyRequestBody is required and must be specified")
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
	localVarPostBody = r.oAuth2ResourceServerJsonWebKeyRequestBody
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	keyId string
	retryCount int32
}

func (r ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest) Execute() (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	return r.ApiService.DeactivateOAuth2ResourceServerJsonWebKeyExecute(r)
}

/*
DeactivateOAuth2ResourceServerJsonWebKey Deactivate a Custom Authorization Server Public JSON Web Key

Deactivates a custom authorization server public JSON web key by key `id`.
> **Note:** Deactivating the active key isn't allowed if the authorization server has access token encryption enabled. You can activate another key, which makes the current key inactive.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param keyId Unique `id` of the Custom Authorization Server JSON Web Key
 @return ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) DeactivateOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest {
	return ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		keyId: keyId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return OAuth2ResourceServerJsonWebKey
func (a *OAuth2ResourceServerCredentialsKeysAPIService) DeactivateOAuth2ResourceServerJsonWebKeyExecute(r ApiDeactivateOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2ResourceServerJsonWebKey
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.DeactivateOAuth2ResourceServerJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiDeleteOAuth2ResourceServerJsonWebKeyRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	keyId string
	retryCount int32
}

func (r ApiDeleteOAuth2ResourceServerJsonWebKeyRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteOAuth2ResourceServerJsonWebKeyExecute(r)
}

/*
DeleteOAuth2ResourceServerJsonWebKey Delete a Custom Authorization Server Public JSON Web Key

Deletes a custom authorization server public JSON web key by key `id`. You can only delete an inactive key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param keyId Unique `id` of the Custom Authorization Server JSON Web Key
 @return ApiDeleteOAuth2ResourceServerJsonWebKeyRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) DeleteOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiDeleteOAuth2ResourceServerJsonWebKeyRequest {
	return ApiDeleteOAuth2ResourceServerJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		keyId: keyId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *OAuth2ResourceServerCredentialsKeysAPIService) DeleteOAuth2ResourceServerJsonWebKeyExecute(r ApiDeleteOAuth2ResourceServerJsonWebKeyRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.DeleteOAuth2ResourceServerJsonWebKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetOAuth2ResourceServerJsonWebKeyRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	keyId string
	retryCount int32
}

func (r ApiGetOAuth2ResourceServerJsonWebKeyRequest) Execute() (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	return r.ApiService.GetOAuth2ResourceServerJsonWebKeyExecute(r)
}

/*
GetOAuth2ResourceServerJsonWebKey Retrieve a Custom Authorization Server Public JSON Web Key

Retrieves a custom authorization server public JSON web key by key `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param keyId Unique `id` of the Custom Authorization Server JSON Web Key
 @return ApiGetOAuth2ResourceServerJsonWebKeyRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) GetOAuth2ResourceServerJsonWebKey(ctx context.Context, authServerId string, keyId string) ApiGetOAuth2ResourceServerJsonWebKeyRequest {
	return ApiGetOAuth2ResourceServerJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		keyId: keyId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return OAuth2ResourceServerJsonWebKey
func (a *OAuth2ResourceServerCredentialsKeysAPIService) GetOAuth2ResourceServerJsonWebKeyExecute(r ApiGetOAuth2ResourceServerJsonWebKeyRequest) (*OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OAuth2ResourceServerJsonWebKey
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.GetOAuth2ResourceServerJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterToString(r.keyId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiListOAuth2ResourceServerJsonWebKeysRequest struct {
	ctx context.Context
	ApiService OAuth2ResourceServerCredentialsKeysAPI
	authServerId string
	retryCount int32
}

func (r ApiListOAuth2ResourceServerJsonWebKeysRequest) Execute() ([]OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	return r.ApiService.ListOAuth2ResourceServerJsonWebKeysExecute(r)
}

/*
ListOAuth2ResourceServerJsonWebKeys List all Custom Authorization Server Public JSON Web Keys

Lists all the public keys used by the custom authorization server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @return ApiListOAuth2ResourceServerJsonWebKeysRequest
*/
func (a *OAuth2ResourceServerCredentialsKeysAPIService) ListOAuth2ResourceServerJsonWebKeys(ctx context.Context, authServerId string) ApiListOAuth2ResourceServerJsonWebKeysRequest {
	return ApiListOAuth2ResourceServerJsonWebKeysRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []OAuth2ResourceServerJsonWebKey
func (a *OAuth2ResourceServerCredentialsKeysAPIService) ListOAuth2ResourceServerJsonWebKeysExecute(r ApiListOAuth2ResourceServerJsonWebKeysRequest) ([]OAuth2ResourceServerJsonWebKey, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OAuth2ResourceServerJsonWebKey
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth2ResourceServerCredentialsKeysAPIService.ListOAuth2ResourceServerJsonWebKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys"
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
		if localVarHTTPResponse.StatusCode == 401 {
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

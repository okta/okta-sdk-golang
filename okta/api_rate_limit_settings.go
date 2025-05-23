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
)


type RateLimitSettingsAPI interface {

	/*
	GetRateLimitSettingsAdminNotifications Retrieve the Rate Limit Admin Notification Settings

	Retrieves the currently configured Rate Limit Admin Notification Settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsAdminNotificationsRequest
	*/
	GetRateLimitSettingsAdminNotifications(ctx context.Context) ApiGetRateLimitSettingsAdminNotificationsRequest

	// GetRateLimitSettingsAdminNotificationsExecute executes the request
	//  @return RateLimitAdminNotifications
	GetRateLimitSettingsAdminNotificationsExecute(r ApiGetRateLimitSettingsAdminNotificationsRequest) (*RateLimitAdminNotifications, *APIResponse, error)

	/*
	GetRateLimitSettingsPerClient Retrieve the Per-Client Rate Limit Settings

	Retrieves the currently configured Per-Client Rate Limit Settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsPerClientRequest
	*/
	GetRateLimitSettingsPerClient(ctx context.Context) ApiGetRateLimitSettingsPerClientRequest

	// GetRateLimitSettingsPerClientExecute executes the request
	//  @return PerClientRateLimitSettings
	GetRateLimitSettingsPerClientExecute(r ApiGetRateLimitSettingsPerClientRequest) (*PerClientRateLimitSettings, *APIResponse, error)

	/*
	GetRateLimitSettingsWarningThreshold Retrieve the Rate Limit Warning Threshold Percentage

	Retrieves the currently configured threshold for warning notifications when the API's rate limit is exceeded

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRateLimitSettingsWarningThresholdRequest
	*/
	GetRateLimitSettingsWarningThreshold(ctx context.Context) ApiGetRateLimitSettingsWarningThresholdRequest

	// GetRateLimitSettingsWarningThresholdExecute executes the request
	//  @return RateLimitWarningThresholdResponse
	GetRateLimitSettingsWarningThresholdExecute(r ApiGetRateLimitSettingsWarningThresholdRequest) (*RateLimitWarningThresholdResponse, *APIResponse, error)

	/*
	ReplaceRateLimitSettingsAdminNotifications Replace the Rate Limit Admin Notification Settings

	Replaces the Rate Limit Admin Notification Settings and returns the configured properties

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsAdminNotificationsRequest
	*/
	ReplaceRateLimitSettingsAdminNotifications(ctx context.Context) ApiReplaceRateLimitSettingsAdminNotificationsRequest

	// ReplaceRateLimitSettingsAdminNotificationsExecute executes the request
	//  @return RateLimitAdminNotifications
	ReplaceRateLimitSettingsAdminNotificationsExecute(r ApiReplaceRateLimitSettingsAdminNotificationsRequest) (*RateLimitAdminNotifications, *APIResponse, error)

	/*
	ReplaceRateLimitSettingsPerClient Replace the Per-Client Rate Limit Settings

	Replaces the Per-Client Rate Limit Settings and returns the configured properties

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsPerClientRequest
	*/
	ReplaceRateLimitSettingsPerClient(ctx context.Context) ApiReplaceRateLimitSettingsPerClientRequest

	// ReplaceRateLimitSettingsPerClientExecute executes the request
	//  @return PerClientRateLimitSettings
	ReplaceRateLimitSettingsPerClientExecute(r ApiReplaceRateLimitSettingsPerClientRequest) (*PerClientRateLimitSettings, *APIResponse, error)

	/*
	ReplaceRateLimitSettingsWarningThreshold Replace the Rate Limit Warning Threshold Percentage

	Replaces the Rate Limit Warning Threshold Percentage and returns the configured property

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceRateLimitSettingsWarningThresholdRequest
	*/
	ReplaceRateLimitSettingsWarningThreshold(ctx context.Context) ApiReplaceRateLimitSettingsWarningThresholdRequest

	// ReplaceRateLimitSettingsWarningThresholdExecute executes the request
	//  @return RateLimitWarningThresholdResponse
	ReplaceRateLimitSettingsWarningThresholdExecute(r ApiReplaceRateLimitSettingsWarningThresholdRequest) (*RateLimitWarningThresholdResponse, *APIResponse, error)
}

// RateLimitSettingsAPIService RateLimitSettingsAPI service
type RateLimitSettingsAPIService service

type ApiGetRateLimitSettingsAdminNotificationsRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	retryCount int32
}

func (r ApiGetRateLimitSettingsAdminNotificationsRequest) Execute() (*RateLimitAdminNotifications, *APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsAdminNotificationsExecute(r)
}

/*
GetRateLimitSettingsAdminNotifications Retrieve the Rate Limit Admin Notification Settings

Retrieves the currently configured Rate Limit Admin Notification Settings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsAdminNotificationsRequest
*/
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsAdminNotifications(ctx context.Context) ApiGetRateLimitSettingsAdminNotificationsRequest {
	return ApiGetRateLimitSettingsAdminNotificationsRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitAdminNotifications
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsAdminNotificationsExecute(r ApiGetRateLimitSettingsAdminNotificationsRequest) (*RateLimitAdminNotifications, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RateLimitAdminNotifications
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsAdminNotifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/admin-notifications"

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

type ApiGetRateLimitSettingsPerClientRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	retryCount int32
}

func (r ApiGetRateLimitSettingsPerClientRequest) Execute() (*PerClientRateLimitSettings, *APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsPerClientExecute(r)
}

/*
GetRateLimitSettingsPerClient Retrieve the Per-Client Rate Limit Settings

Retrieves the currently configured Per-Client Rate Limit Settings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsPerClientRequest
*/
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsPerClient(ctx context.Context) ApiGetRateLimitSettingsPerClientRequest {
	return ApiGetRateLimitSettingsPerClientRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return PerClientRateLimitSettings
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsPerClientExecute(r ApiGetRateLimitSettingsPerClientRequest) (*PerClientRateLimitSettings, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PerClientRateLimitSettings
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsPerClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/per-client"

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

type ApiGetRateLimitSettingsWarningThresholdRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	retryCount int32
}

func (r ApiGetRateLimitSettingsWarningThresholdRequest) Execute() (*RateLimitWarningThresholdResponse, *APIResponse, error) {
	return r.ApiService.GetRateLimitSettingsWarningThresholdExecute(r)
}

/*
GetRateLimitSettingsWarningThreshold Retrieve the Rate Limit Warning Threshold Percentage

Retrieves the currently configured threshold for warning notifications when the API's rate limit is exceeded

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRateLimitSettingsWarningThresholdRequest
*/
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsWarningThreshold(ctx context.Context) ApiGetRateLimitSettingsWarningThresholdRequest {
	return ApiGetRateLimitSettingsWarningThresholdRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitWarningThresholdResponse
func (a *RateLimitSettingsAPIService) GetRateLimitSettingsWarningThresholdExecute(r ApiGetRateLimitSettingsWarningThresholdRequest) (*RateLimitWarningThresholdResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RateLimitWarningThresholdResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.GetRateLimitSettingsWarningThreshold")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/warning-threshold"

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

type ApiReplaceRateLimitSettingsAdminNotificationsRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	rateLimitAdminNotifications *RateLimitAdminNotifications
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsAdminNotificationsRequest) RateLimitAdminNotifications(rateLimitAdminNotifications RateLimitAdminNotifications) ApiReplaceRateLimitSettingsAdminNotificationsRequest {
	r.rateLimitAdminNotifications = &rateLimitAdminNotifications
	return r
}

func (r ApiReplaceRateLimitSettingsAdminNotificationsRequest) Execute() (*RateLimitAdminNotifications, *APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsAdminNotificationsExecute(r)
}

/*
ReplaceRateLimitSettingsAdminNotifications Replace the Rate Limit Admin Notification Settings

Replaces the Rate Limit Admin Notification Settings and returns the configured properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsAdminNotificationsRequest
*/
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsAdminNotifications(ctx context.Context) ApiReplaceRateLimitSettingsAdminNotificationsRequest {
	return ApiReplaceRateLimitSettingsAdminNotificationsRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitAdminNotifications
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsAdminNotificationsExecute(r ApiReplaceRateLimitSettingsAdminNotificationsRequest) (*RateLimitAdminNotifications, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RateLimitAdminNotifications
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsAdminNotifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/admin-notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rateLimitAdminNotifications == nil {
		return localVarReturnValue, nil, reportError("rateLimitAdminNotifications is required and must be specified")
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
	localVarPostBody = r.rateLimitAdminNotifications
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

type ApiReplaceRateLimitSettingsPerClientRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	perClientRateLimitSettings *PerClientRateLimitSettings
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsPerClientRequest) PerClientRateLimitSettings(perClientRateLimitSettings PerClientRateLimitSettings) ApiReplaceRateLimitSettingsPerClientRequest {
	r.perClientRateLimitSettings = &perClientRateLimitSettings
	return r
}

func (r ApiReplaceRateLimitSettingsPerClientRequest) Execute() (*PerClientRateLimitSettings, *APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsPerClientExecute(r)
}

/*
ReplaceRateLimitSettingsPerClient Replace the Per-Client Rate Limit Settings

Replaces the Per-Client Rate Limit Settings and returns the configured properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsPerClientRequest
*/
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsPerClient(ctx context.Context) ApiReplaceRateLimitSettingsPerClientRequest {
	return ApiReplaceRateLimitSettingsPerClientRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return PerClientRateLimitSettings
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsPerClientExecute(r ApiReplaceRateLimitSettingsPerClientRequest) (*PerClientRateLimitSettings, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PerClientRateLimitSettings
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsPerClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/per-client"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.perClientRateLimitSettings == nil {
		return localVarReturnValue, nil, reportError("perClientRateLimitSettings is required and must be specified")
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
	localVarPostBody = r.perClientRateLimitSettings
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

type ApiReplaceRateLimitSettingsWarningThresholdRequest struct {
	ctx context.Context
	ApiService RateLimitSettingsAPI
	rateLimitWarningThreshold *RateLimitWarningThresholdRequest
	retryCount int32
}

func (r ApiReplaceRateLimitSettingsWarningThresholdRequest) RateLimitWarningThreshold(rateLimitWarningThreshold RateLimitWarningThresholdRequest) ApiReplaceRateLimitSettingsWarningThresholdRequest {
	r.rateLimitWarningThreshold = &rateLimitWarningThreshold
	return r
}

func (r ApiReplaceRateLimitSettingsWarningThresholdRequest) Execute() (*RateLimitWarningThresholdResponse, *APIResponse, error) {
	return r.ApiService.ReplaceRateLimitSettingsWarningThresholdExecute(r)
}

/*
ReplaceRateLimitSettingsWarningThreshold Replace the Rate Limit Warning Threshold Percentage

Replaces the Rate Limit Warning Threshold Percentage and returns the configured property

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceRateLimitSettingsWarningThresholdRequest
*/
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsWarningThreshold(ctx context.Context) ApiReplaceRateLimitSettingsWarningThresholdRequest {
	return ApiReplaceRateLimitSettingsWarningThresholdRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return RateLimitWarningThresholdResponse
func (a *RateLimitSettingsAPIService) ReplaceRateLimitSettingsWarningThresholdExecute(r ApiReplaceRateLimitSettingsWarningThresholdRequest) (*RateLimitWarningThresholdResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RateLimitWarningThresholdResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitSettingsAPIService.ReplaceRateLimitSettingsWarningThreshold")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rate-limit-settings/warning-threshold"

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
	localVarPostBody = r.rateLimitWarningThreshold
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

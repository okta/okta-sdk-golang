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
	"time"
)

type DeviceAccessAPI interface {

	/*
		GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting Retrieve the Desktop MFA Enforce Number Matching Challenge org setting

		Retrieves the status of the Desktop MFA Enforce Number Matching Challenge push notifications feature. That is, whether or not the feature is enabled for your org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest
	*/
	GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx context.Context) ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest

	// GetDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute executes the request
	//  @return DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	GetDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error)

	/*
		GetDesktopMFARecoveryPinOrgSetting Retrieve the Desktop MFA Recovery PIN org setting

		Retrieves the status of the Desktop MFA Recovery PIN feature. That is, whether or not the feature is enabled for your org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetDesktopMFARecoveryPinOrgSettingRequest
	*/
	GetDesktopMFARecoveryPinOrgSetting(ctx context.Context) ApiGetDesktopMFARecoveryPinOrgSettingRequest

	// GetDesktopMFARecoveryPinOrgSettingExecute executes the request
	//  @return DesktopMFARecoveryPinOrgSetting
	GetDesktopMFARecoveryPinOrgSettingExecute(r ApiGetDesktopMFARecoveryPinOrgSettingRequest) (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error)

	/*
		ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting Replace the Desktop MFA Enforce Number Matching Challenge org setting

		Replaces the status of the Desktop MFA Enforce Number Matching Challenge push notifications feature. That is, whether or not the feature is enabled for your org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest
	*/
	ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx context.Context) ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest

	// ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute executes the request
	//  @return DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error)

	/*
		ReplaceDesktopMFARecoveryPinOrgSetting Replace the Desktop MFA Recovery PIN org setting

		Replaces the Desktop MFA Recovery PIN feature for your org

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiReplaceDesktopMFARecoveryPinOrgSettingRequest
	*/
	ReplaceDesktopMFARecoveryPinOrgSetting(ctx context.Context) ApiReplaceDesktopMFARecoveryPinOrgSettingRequest

	// ReplaceDesktopMFARecoveryPinOrgSettingExecute executes the request
	//  @return DesktopMFARecoveryPinOrgSetting
	ReplaceDesktopMFARecoveryPinOrgSettingExecute(r ApiReplaceDesktopMFARecoveryPinOrgSettingRequest) (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error)
}

// DeviceAccessAPIService DeviceAccessAPI service
type DeviceAccessAPIService service

type ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest struct {
	ctx        context.Context
	ApiService DeviceAccessAPI
	retryCount int32
}

func (r ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) Execute() (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error) {
	return r.ApiService.GetDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r)
}

/*
GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting Retrieve the Desktop MFA Enforce Number Matching Challenge org setting

Retrieves the status of the Desktop MFA Enforce Number Matching Challenge push notifications feature. That is, whether or not the feature is enabled for your org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest
*/
func (a *DeviceAccessAPIService) GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx context.Context) ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest {
	return ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return DesktopMFAEnforceNumberMatchingChallengeOrgSetting
func (a *DeviceAccessAPIService) GetDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r ApiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DesktopMFAEnforceNumberMatchingChallengeOrgSetting
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAccessAPIService.GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-access/api/v1/desktop-mfa/enforce-number-matching-challenge-settings"

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

type ApiGetDesktopMFARecoveryPinOrgSettingRequest struct {
	ctx        context.Context
	ApiService DeviceAccessAPI
	retryCount int32
}

func (r ApiGetDesktopMFARecoveryPinOrgSettingRequest) Execute() (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error) {
	return r.ApiService.GetDesktopMFARecoveryPinOrgSettingExecute(r)
}

/*
GetDesktopMFARecoveryPinOrgSetting Retrieve the Desktop MFA Recovery PIN org setting

Retrieves the status of the Desktop MFA Recovery PIN feature. That is, whether or not the feature is enabled for your org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDesktopMFARecoveryPinOrgSettingRequest
*/
func (a *DeviceAccessAPIService) GetDesktopMFARecoveryPinOrgSetting(ctx context.Context) ApiGetDesktopMFARecoveryPinOrgSettingRequest {
	return ApiGetDesktopMFARecoveryPinOrgSettingRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return DesktopMFARecoveryPinOrgSetting
func (a *DeviceAccessAPIService) GetDesktopMFARecoveryPinOrgSettingExecute(r ApiGetDesktopMFARecoveryPinOrgSettingRequest) (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DesktopMFARecoveryPinOrgSetting
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAccessAPIService.GetDesktopMFARecoveryPinOrgSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-access/api/v1/desktop-mfa/recovery-pin-settings"

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

type ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest struct {
	ctx                                                context.Context
	ApiService                                         DeviceAccessAPI
	desktopMFAEnforceNumberMatchingChallengeOrgSetting *DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	retryCount                                         int32
}

func (r ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) DesktopMFAEnforceNumberMatchingChallengeOrgSetting(desktopMFAEnforceNumberMatchingChallengeOrgSetting DesktopMFAEnforceNumberMatchingChallengeOrgSetting) ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest {
	r.desktopMFAEnforceNumberMatchingChallengeOrgSetting = &desktopMFAEnforceNumberMatchingChallengeOrgSetting
	return r
}

func (r ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) Execute() (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error) {
	return r.ApiService.ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r)
}

/*
ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting Replace the Desktop MFA Enforce Number Matching Challenge org setting

Replaces the status of the Desktop MFA Enforce Number Matching Challenge push notifications feature. That is, whether or not the feature is enabled for your org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest
*/
func (a *DeviceAccessAPIService) ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx context.Context) ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest {
	return ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return DesktopMFAEnforceNumberMatchingChallengeOrgSetting
func (a *DeviceAccessAPIService) ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingExecute(r ApiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest) (*DesktopMFAEnforceNumberMatchingChallengeOrgSetting, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DesktopMFAEnforceNumberMatchingChallengeOrgSetting
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAccessAPIService.ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-access/api/v1/desktop-mfa/enforce-number-matching-challenge-settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.desktopMFAEnforceNumberMatchingChallengeOrgSetting == nil {
		return localVarReturnValue, nil, reportError("desktopMFAEnforceNumberMatchingChallengeOrgSetting is required and must be specified")
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
	localVarPostBody = r.desktopMFAEnforceNumberMatchingChallengeOrgSetting
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

type ApiReplaceDesktopMFARecoveryPinOrgSettingRequest struct {
	ctx                             context.Context
	ApiService                      DeviceAccessAPI
	desktopMFARecoveryPinOrgSetting *DesktopMFARecoveryPinOrgSetting
	retryCount                      int32
}

func (r ApiReplaceDesktopMFARecoveryPinOrgSettingRequest) DesktopMFARecoveryPinOrgSetting(desktopMFARecoveryPinOrgSetting DesktopMFARecoveryPinOrgSetting) ApiReplaceDesktopMFARecoveryPinOrgSettingRequest {
	r.desktopMFARecoveryPinOrgSetting = &desktopMFARecoveryPinOrgSetting
	return r
}

func (r ApiReplaceDesktopMFARecoveryPinOrgSettingRequest) Execute() (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error) {
	return r.ApiService.ReplaceDesktopMFARecoveryPinOrgSettingExecute(r)
}

/*
ReplaceDesktopMFARecoveryPinOrgSetting Replace the Desktop MFA Recovery PIN org setting

Replaces the Desktop MFA Recovery PIN feature for your org

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceDesktopMFARecoveryPinOrgSettingRequest
*/
func (a *DeviceAccessAPIService) ReplaceDesktopMFARecoveryPinOrgSetting(ctx context.Context) ApiReplaceDesktopMFARecoveryPinOrgSettingRequest {
	return ApiReplaceDesktopMFARecoveryPinOrgSettingRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return DesktopMFARecoveryPinOrgSetting
func (a *DeviceAccessAPIService) ReplaceDesktopMFARecoveryPinOrgSettingExecute(r ApiReplaceDesktopMFARecoveryPinOrgSettingRequest) (*DesktopMFARecoveryPinOrgSetting, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DesktopMFARecoveryPinOrgSetting
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAccessAPIService.ReplaceDesktopMFARecoveryPinOrgSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-access/api/v1/desktop-mfa/recovery-pin-settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.desktopMFARecoveryPinOrgSetting == nil {
		return localVarReturnValue, nil, reportError("desktopMFARecoveryPinOrgSetting is required and must be specified")
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
	localVarPostBody = r.desktopMFARecoveryPinOrgSetting
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

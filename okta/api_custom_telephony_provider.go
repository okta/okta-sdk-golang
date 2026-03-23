/*
Okta Admin Management API

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

type CustomTelephonyProviderAPI interface {

	/*
		ActivateCustomTelephonyCredential Activate a custom telephony provider

		Activates a custom telephony provider by its ID. You must activate a provider before it can be used.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customTelephonyProviderId The ID of the custom telephony provider
		@return ApiActivateCustomTelephonyCredentialRequest
	*/
	ActivateCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiActivateCustomTelephonyCredentialRequest

	// ActivateCustomTelephonyCredentialExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	ActivateCustomTelephonyCredentialExecute(r ApiActivateCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
		CreateCustomTelephonyProviderCredentials Create a custom telephony provider

		Creates a custom telephony provider with the provided credentials

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateCustomTelephonyProviderCredentialsRequest
	*/
	CreateCustomTelephonyProviderCredentials(ctx context.Context) ApiCreateCustomTelephonyProviderCredentialsRequest

	// CreateCustomTelephonyProviderCredentialsExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	CreateCustomTelephonyProviderCredentialsExecute(r ApiCreateCustomTelephonyProviderCredentialsRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
			DeactivateCustomTelephonyCredential Deactivate a custom telephony provider

			Deactivates a custom telephony provider by its ID. Keep the following points in mind when you deactivate a provider:
		* You must deactivate a provider before deleting it.
		* If you have two telephony providers configured, and both are active, you can only deactivate the secondary provider. The second provider is the one that isn't set as the primary provider.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param customTelephonyProviderId The ID of the custom telephony provider
			@return ApiDeactivateCustomTelephonyCredentialRequest
	*/
	DeactivateCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiDeactivateCustomTelephonyCredentialRequest

	// DeactivateCustomTelephonyCredentialExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	DeactivateCustomTelephonyCredentialExecute(r ApiDeactivateCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
			DeleteCustomTelephonyProviderCredential Delete a custom telephony provider

			Deletes a custom telephony provider by its ID.

		Before you delete a provider, ensure that it is [deactivated](/openapi/okta-management/management/customtelephonyprovider/deactivatecustomtelephonycredential). Consider setting up another telephony provider if you still plan to use telephony in your org. See [Set up an external telephony provider](https://help.okta.com/okta_help.htm?type=oie&id=about-telephony).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param customTelephonyProviderId The ID of the custom telephony provider
			@return ApiDeleteCustomTelephonyProviderCredentialRequest
	*/
	DeleteCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiDeleteCustomTelephonyProviderCredentialRequest

	// DeleteCustomTelephonyProviderCredentialExecute executes the request
	DeleteCustomTelephonyProviderCredentialExecute(r ApiDeleteCustomTelephonyProviderCredentialRequest) (*APIResponse, error)

	/*
		GetCustomTelephonyProviderCredential Retrieve a custom telephony provider

		Retrieves the details of a custom telephony provider by its ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customTelephonyProviderId The ID of the custom telephony provider
		@return ApiGetCustomTelephonyProviderCredentialRequest
	*/
	GetCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiGetCustomTelephonyProviderCredentialRequest

	// GetCustomTelephonyProviderCredentialExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	GetCustomTelephonyProviderCredentialExecute(r ApiGetCustomTelephonyProviderCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
		ListAllCustomTelephonyProviderCredentials List all custom telephony providers

		Lists all custom telephony providers that are configured in your org

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListAllCustomTelephonyProviderCredentialsRequest
	*/
	ListAllCustomTelephonyProviderCredentials(ctx context.Context) ApiListAllCustomTelephonyProviderCredentialsRequest

	// ListAllCustomTelephonyProviderCredentialsExecute executes the request
	//  @return []CustomTelephonyProviderCredentialResponse
	ListAllCustomTelephonyProviderCredentialsExecute(r ApiListAllCustomTelephonyProviderCredentialsRequest) ([]CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
			SendTestCustomTelephonyProviderCredential Send a test message from a custom telephony provider

			Sends a test message (SMS or call) using the specified custom telephony provider to verify that the provider is configured correctly.

		You must provide a valid phone number and country code to send the test message. Send it to a phone number that you have access to so you can confirm that the message was received.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param customTelephonyProviderId The ID of the custom telephony provider
			@return ApiSendTestCustomTelephonyProviderCredentialRequest
	*/
	SendTestCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiSendTestCustomTelephonyProviderCredentialRequest

	// SendTestCustomTelephonyProviderCredentialExecute executes the request
	SendTestCustomTelephonyProviderCredentialExecute(r ApiSendTestCustomTelephonyProviderCredentialRequest) (*APIResponse, error)

	/*
		SetAsPrimaryCustomTelephonyCredential Set a custom telephony provider as a primary telephony provider

		Sets a custom telephony provider as the primary telephony provider for the org. You can only set one provider as a primary provider at a time.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customTelephonyProviderId The ID of the custom telephony provider
		@return ApiSetAsPrimaryCustomTelephonyCredentialRequest
	*/
	SetAsPrimaryCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiSetAsPrimaryCustomTelephonyCredentialRequest

	// SetAsPrimaryCustomTelephonyCredentialExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	SetAsPrimaryCustomTelephonyCredentialExecute(r ApiSetAsPrimaryCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)

	/*
		UpdateCustomTelephonyProviderCredential Update a custom telephony provider credential

		Updates the credentials of an existing custom telephony provider

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customTelephonyProviderId The ID of the custom telephony provider
		@return ApiUpdateCustomTelephonyProviderCredentialRequest
	*/
	UpdateCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiUpdateCustomTelephonyProviderCredentialRequest

	// UpdateCustomTelephonyProviderCredentialExecute executes the request
	//  @return CustomTelephonyProviderCredentialResponse
	UpdateCustomTelephonyProviderCredentialExecute(r ApiUpdateCustomTelephonyProviderCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error)
}

// CustomTelephonyProviderAPIService CustomTelephonyProviderAPI service
type CustomTelephonyProviderAPIService service

type ApiActivateCustomTelephonyCredentialRequest struct {
	ctx                       context.Context
	ApiService                CustomTelephonyProviderAPI
	customTelephonyProviderId string
	retryCount                int32
}

func (r ApiActivateCustomTelephonyCredentialRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.ActivateCustomTelephonyCredentialExecute(r)
}

/*
ActivateCustomTelephonyCredential Activate a custom telephony provider

Activates a custom telephony provider by its ID. You must activate a provider before it can be used.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiActivateCustomTelephonyCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) ActivateCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiActivateCustomTelephonyCredentialRequest {
	return ApiActivateCustomTelephonyCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) ActivateCustomTelephonyCredentialExecute(r ApiActivateCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.ActivateCustomTelephonyCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

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

type ApiCreateCustomTelephonyProviderCredentialsRequest struct {
	ctx                                            context.Context
	ApiService                                     CustomTelephonyProviderAPI
	customTelephonyProviderCredentialCreateRequest *CustomTelephonyProviderCredentialCreateRequest
	retryCount                                     int32
}

func (r ApiCreateCustomTelephonyProviderCredentialsRequest) CustomTelephonyProviderCredentialCreateRequest(customTelephonyProviderCredentialCreateRequest CustomTelephonyProviderCredentialCreateRequest) ApiCreateCustomTelephonyProviderCredentialsRequest {
	r.customTelephonyProviderCredentialCreateRequest = &customTelephonyProviderCredentialCreateRequest
	return r
}

func (r ApiCreateCustomTelephonyProviderCredentialsRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.CreateCustomTelephonyProviderCredentialsExecute(r)
}

/*
CreateCustomTelephonyProviderCredentials Create a custom telephony provider

Creates a custom telephony provider with the provided credentials

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateCustomTelephonyProviderCredentialsRequest
*/
func (a *CustomTelephonyProviderAPIService) CreateCustomTelephonyProviderCredentials(ctx context.Context) ApiCreateCustomTelephonyProviderCredentialsRequest {
	return ApiCreateCustomTelephonyProviderCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) CreateCustomTelephonyProviderCredentialsExecute(r ApiCreateCustomTelephonyProviderCredentialsRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.CreateCustomTelephonyProviderCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customTelephonyProviderCredentialCreateRequest == nil {
		return localVarReturnValue, nil, reportError("customTelephonyProviderCredentialCreateRequest is required and must be specified")
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
	localVarPostBody = r.customTelephonyProviderCredentialCreateRequest
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

type ApiDeactivateCustomTelephonyCredentialRequest struct {
	ctx                       context.Context
	ApiService                CustomTelephonyProviderAPI
	customTelephonyProviderId string
	retryCount                int32
}

func (r ApiDeactivateCustomTelephonyCredentialRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.DeactivateCustomTelephonyCredentialExecute(r)
}

/*
DeactivateCustomTelephonyCredential Deactivate a custom telephony provider

Deactivates a custom telephony provider by its ID. Keep the following points in mind when you deactivate a provider:
* You must deactivate a provider before deleting it.
* If you have two telephony providers configured, and both are active, you can only deactivate the secondary provider. The second provider is the one that isn't set as the primary provider.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiDeactivateCustomTelephonyCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) DeactivateCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiDeactivateCustomTelephonyCredentialRequest {
	return ApiDeactivateCustomTelephonyCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) DeactivateCustomTelephonyCredentialExecute(r ApiDeactivateCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.DeactivateCustomTelephonyCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

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

type ApiDeleteCustomTelephonyProviderCredentialRequest struct {
	ctx                       context.Context
	ApiService                CustomTelephonyProviderAPI
	customTelephonyProviderId string
	retryCount                int32
}

func (r ApiDeleteCustomTelephonyProviderCredentialRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteCustomTelephonyProviderCredentialExecute(r)
}

/*
DeleteCustomTelephonyProviderCredential Delete a custom telephony provider

Deletes a custom telephony provider by its ID.

Before you delete a provider, ensure that it is [deactivated](/openapi/okta-management/management/customtelephonyprovider/deactivatecustomtelephonycredential). Consider setting up another telephony provider if you still plan to use telephony in your org. See [Set up an external telephony provider](https://help.okta.com/okta_help.htm?type=oie&id=about-telephony).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiDeleteCustomTelephonyProviderCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) DeleteCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiDeleteCustomTelephonyProviderCredentialRequest {
	return ApiDeleteCustomTelephonyProviderCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
func (a *CustomTelephonyProviderAPIService) DeleteCustomTelephonyProviderCredentialExecute(r ApiDeleteCustomTelephonyProviderCredentialRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.DeleteCustomTelephonyProviderCredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

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
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiGetCustomTelephonyProviderCredentialRequest struct {
	ctx                       context.Context
	ApiService                CustomTelephonyProviderAPI
	customTelephonyProviderId string
	retryCount                int32
}

func (r ApiGetCustomTelephonyProviderCredentialRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.GetCustomTelephonyProviderCredentialExecute(r)
}

/*
GetCustomTelephonyProviderCredential Retrieve a custom telephony provider

Retrieves the details of a custom telephony provider by its ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiGetCustomTelephonyProviderCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) GetCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiGetCustomTelephonyProviderCredentialRequest {
	return ApiGetCustomTelephonyProviderCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) GetCustomTelephonyProviderCredentialExecute(r ApiGetCustomTelephonyProviderCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.GetCustomTelephonyProviderCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

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

type ApiListAllCustomTelephonyProviderCredentialsRequest struct {
	ctx        context.Context
	ApiService CustomTelephonyProviderAPI
	retryCount int32
}

func (r ApiListAllCustomTelephonyProviderCredentialsRequest) Execute() ([]CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.ListAllCustomTelephonyProviderCredentialsExecute(r)
}

/*
ListAllCustomTelephonyProviderCredentials List all custom telephony providers

Lists all custom telephony providers that are configured in your org

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAllCustomTelephonyProviderCredentialsRequest
*/
func (a *CustomTelephonyProviderAPIService) ListAllCustomTelephonyProviderCredentials(ctx context.Context) ApiListAllCustomTelephonyProviderCredentialsRequest {
	return ApiListAllCustomTelephonyProviderCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) ListAllCustomTelephonyProviderCredentialsExecute(r ApiListAllCustomTelephonyProviderCredentialsRequest) ([]CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.ListAllCustomTelephonyProviderCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers"

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

type ApiSendTestCustomTelephonyProviderCredentialRequest struct {
	ctx                                              context.Context
	ApiService                                       CustomTelephonyProviderAPI
	customTelephonyProviderId                        string
	customTelephonyProviderCredentialSendTestRequest *CustomTelephonyProviderCredentialSendTestRequest
	retryCount                                       int32
}

func (r ApiSendTestCustomTelephonyProviderCredentialRequest) CustomTelephonyProviderCredentialSendTestRequest(customTelephonyProviderCredentialSendTestRequest CustomTelephonyProviderCredentialSendTestRequest) ApiSendTestCustomTelephonyProviderCredentialRequest {
	r.customTelephonyProviderCredentialSendTestRequest = &customTelephonyProviderCredentialSendTestRequest
	return r
}

func (r ApiSendTestCustomTelephonyProviderCredentialRequest) Execute() (*APIResponse, error) {
	return r.ApiService.SendTestCustomTelephonyProviderCredentialExecute(r)
}

/*
SendTestCustomTelephonyProviderCredential Send a test message from a custom telephony provider

Sends a test message (SMS or call) using the specified custom telephony provider to verify that the provider is configured correctly.

You must provide a valid phone number and country code to send the test message. Send it to a phone number that you have access to so you can confirm that the message was received.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiSendTestCustomTelephonyProviderCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) SendTestCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiSendTestCustomTelephonyProviderCredentialRequest {
	return ApiSendTestCustomTelephonyProviderCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
func (a *CustomTelephonyProviderAPIService) SendTestCustomTelephonyProviderCredentialExecute(r ApiSendTestCustomTelephonyProviderCredentialRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.SendTestCustomTelephonyProviderCredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}/test"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customTelephonyProviderCredentialSendTestRequest == nil {
		return nil, reportError("customTelephonyProviderCredentialSendTestRequest is required and must be specified")
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
	localVarPostBody = r.customTelephonyProviderCredentialSendTestRequest
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
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiSetAsPrimaryCustomTelephonyCredentialRequest struct {
	ctx                       context.Context
	ApiService                CustomTelephonyProviderAPI
	customTelephonyProviderId string
	retryCount                int32
}

func (r ApiSetAsPrimaryCustomTelephonyCredentialRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.SetAsPrimaryCustomTelephonyCredentialExecute(r)
}

/*
SetAsPrimaryCustomTelephonyCredential Set a custom telephony provider as a primary telephony provider

Sets a custom telephony provider as the primary telephony provider for the org. You can only set one provider as a primary provider at a time.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiSetAsPrimaryCustomTelephonyCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) SetAsPrimaryCustomTelephonyCredential(ctx context.Context, customTelephonyProviderId string) ApiSetAsPrimaryCustomTelephonyCredentialRequest {
	return ApiSetAsPrimaryCustomTelephonyCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) SetAsPrimaryCustomTelephonyCredentialExecute(r ApiSetAsPrimaryCustomTelephonyCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.SetAsPrimaryCustomTelephonyCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}/setAsPrimary"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

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

type ApiUpdateCustomTelephonyProviderCredentialRequest struct {
	ctx                                            context.Context
	ApiService                                     CustomTelephonyProviderAPI
	customTelephonyProviderId                      string
	customTelephonyProviderCredentialUpdateRequest *CustomTelephonyProviderCredentialUpdateRequest
	retryCount                                     int32
}

func (r ApiUpdateCustomTelephonyProviderCredentialRequest) CustomTelephonyProviderCredentialUpdateRequest(customTelephonyProviderCredentialUpdateRequest CustomTelephonyProviderCredentialUpdateRequest) ApiUpdateCustomTelephonyProviderCredentialRequest {
	r.customTelephonyProviderCredentialUpdateRequest = &customTelephonyProviderCredentialUpdateRequest
	return r
}

func (r ApiUpdateCustomTelephonyProviderCredentialRequest) Execute() (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	return r.ApiService.UpdateCustomTelephonyProviderCredentialExecute(r)
}

/*
UpdateCustomTelephonyProviderCredential Update a custom telephony provider credential

Updates the credentials of an existing custom telephony provider

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customTelephonyProviderId The ID of the custom telephony provider
	@return ApiUpdateCustomTelephonyProviderCredentialRequest
*/
func (a *CustomTelephonyProviderAPIService) UpdateCustomTelephonyProviderCredential(ctx context.Context, customTelephonyProviderId string) ApiUpdateCustomTelephonyProviderCredentialRequest {
	return ApiUpdateCustomTelephonyProviderCredentialRequest{
		ApiService:                a,
		ctx:                       ctx,
		customTelephonyProviderId: customTelephonyProviderId,
		retryCount:                0,
	}
}

// Execute executes the request
//
//	@return CustomTelephonyProviderCredentialResponse
func (a *CustomTelephonyProviderAPIService) UpdateCustomTelephonyProviderCredentialExecute(r ApiUpdateCustomTelephonyProviderCredentialRequest) (*CustomTelephonyProviderCredentialResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomTelephonyProviderCredentialResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomTelephonyProviderAPIService.UpdateCustomTelephonyProviderCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telephony-providers/{customTelephonyProviderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customTelephonyProviderId"+"}", url.PathEscape(parameterToString(r.customTelephonyProviderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customTelephonyProviderCredentialUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("customTelephonyProviderCredentialUpdateRequest is required and must be specified")
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
	localVarPostBody = r.customTelephonyProviderCredentialUpdateRequest
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

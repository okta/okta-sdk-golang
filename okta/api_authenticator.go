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


type AuthenticatorAPI interface {

	/*
	ActivateAuthenticator Activate an Authenticator

	Activates an authenticator by `authenticatorId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@return ApiActivateAuthenticatorRequest
	*/
	ActivateAuthenticator(ctx context.Context, authenticatorId string) ApiActivateAuthenticatorRequest

	// ActivateAuthenticatorExecute executes the request
	//  @return AuthenticatorBase
	ActivateAuthenticatorExecute(r ApiActivateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error)

	/*
	ActivateAuthenticatorMethod Activate an Authenticator Method

	Activates a Method for an Authenticator identified by `authenticatorId` and `methodType`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@param methodType Type of authenticator method
	@return ApiActivateAuthenticatorMethodRequest
	*/
	ActivateAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiActivateAuthenticatorMethodRequest

	// ActivateAuthenticatorMethodExecute executes the request
	//  @return ListAuthenticatorMethods200ResponseInner
	ActivateAuthenticatorMethodExecute(r ApiActivateAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error)

	/*
	CreateAuthenticator Create an Authenticator

	Creates an authenticator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAuthenticatorRequest
	*/
	CreateAuthenticator(ctx context.Context) ApiCreateAuthenticatorRequest

	// CreateAuthenticatorExecute executes the request
	//  @return AuthenticatorBase
	CreateAuthenticatorExecute(r ApiCreateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error)

	/*
	DeactivateAuthenticator Deactivate an Authenticator

	Deactivates an authenticator by `authenticatorId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@return ApiDeactivateAuthenticatorRequest
	*/
	DeactivateAuthenticator(ctx context.Context, authenticatorId string) ApiDeactivateAuthenticatorRequest

	// DeactivateAuthenticatorExecute executes the request
	//  @return AuthenticatorBase
	DeactivateAuthenticatorExecute(r ApiDeactivateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error)

	/*
	DeactivateAuthenticatorMethod Deactivate an Authenticator Method

	Deactivates a Method for an Authenticator identified by `authenticatorId` and `methodType`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@param methodType Type of authenticator method
	@return ApiDeactivateAuthenticatorMethodRequest
	*/
	DeactivateAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiDeactivateAuthenticatorMethodRequest

	// DeactivateAuthenticatorMethodExecute executes the request
	//  @return ListAuthenticatorMethods200ResponseInner
	DeactivateAuthenticatorMethodExecute(r ApiDeactivateAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error)

	/*
	GetAuthenticator Retrieve an Authenticator

	Retrieves an authenticator from your Okta organization by `authenticatorId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@return ApiGetAuthenticatorRequest
	*/
	GetAuthenticator(ctx context.Context, authenticatorId string) ApiGetAuthenticatorRequest

	// GetAuthenticatorExecute executes the request
	//  @return AuthenticatorBase
	GetAuthenticatorExecute(r ApiGetAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error)

	/*
	GetAuthenticatorMethod Retrieve an Authenticator Method

	Retrieves a Method identified by `methodType` of an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@param methodType Type of authenticator method
	@return ApiGetAuthenticatorMethodRequest
	*/
	GetAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiGetAuthenticatorMethodRequest

	// GetAuthenticatorMethodExecute executes the request
	//  @return ListAuthenticatorMethods200ResponseInner
	GetAuthenticatorMethodExecute(r ApiGetAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error)

	/*
	GetWellKnownAppAuthenticatorConfiguration Retrieve the Well-Known App Authenticator Configuration

	Retrieves the well-known app authenticator configuration. Includes an app authenticator's settings, supported methods, and other details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetWellKnownAppAuthenticatorConfigurationRequest
	*/
	GetWellKnownAppAuthenticatorConfiguration(ctx context.Context) ApiGetWellKnownAppAuthenticatorConfigurationRequest

	// GetWellKnownAppAuthenticatorConfigurationExecute executes the request
	//  @return []WellKnownAppAuthenticatorConfiguration
	GetWellKnownAppAuthenticatorConfigurationExecute(r ApiGetWellKnownAppAuthenticatorConfigurationRequest) ([]WellKnownAppAuthenticatorConfiguration, *APIResponse, error)

	/*
	ListAuthenticatorMethods List all Methods of an Authenticator

	Lists all Methods of an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@return ApiListAuthenticatorMethodsRequest
	*/
	ListAuthenticatorMethods(ctx context.Context, authenticatorId string) ApiListAuthenticatorMethodsRequest

	// ListAuthenticatorMethodsExecute executes the request
	//  @return []ListAuthenticatorMethods200ResponseInner
	ListAuthenticatorMethodsExecute(r ApiListAuthenticatorMethodsRequest) ([]ListAuthenticatorMethods200ResponseInner, *APIResponse, error)

	/*
	ListAuthenticators List all Authenticators

	Lists all authenticators

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAuthenticatorsRequest
	*/
	ListAuthenticators(ctx context.Context) ApiListAuthenticatorsRequest

	// ListAuthenticatorsExecute executes the request
	//  @return []ListAuthenticators200ResponseInner
	ListAuthenticatorsExecute(r ApiListAuthenticatorsRequest) ([]ListAuthenticators200ResponseInner, *APIResponse, error)

	/*
	ReplaceAuthenticator Replace an Authenticator

	Replaces the properties for an Authenticator identified by `authenticatorId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@return ApiReplaceAuthenticatorRequest
	*/
	ReplaceAuthenticator(ctx context.Context, authenticatorId string) ApiReplaceAuthenticatorRequest

	// ReplaceAuthenticatorExecute executes the request
	//  @return AuthenticatorBase
	ReplaceAuthenticatorExecute(r ApiReplaceAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error)

	/*
	ReplaceAuthenticatorMethod Replace an Authenticator Method

	Replaces a Method of `methodType` for an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authenticatorId `id` of the Authenticator
	@param methodType Type of authenticator method
	@return ApiReplaceAuthenticatorMethodRequest
	*/
	ReplaceAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiReplaceAuthenticatorMethodRequest

	// ReplaceAuthenticatorMethodExecute executes the request
	//  @return ListAuthenticatorMethods200ResponseInner
	ReplaceAuthenticatorMethodExecute(r ApiReplaceAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error)
}

// AuthenticatorAPIService AuthenticatorAPI service
type AuthenticatorAPIService service

type ApiActivateAuthenticatorRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	retryCount int32
}

func (r ApiActivateAuthenticatorRequest) Execute() (*AuthenticatorBase, *APIResponse, error) {
	return r.ApiService.ActivateAuthenticatorExecute(r)
}

/*
ActivateAuthenticator Activate an Authenticator

Activates an authenticator by `authenticatorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @return ApiActivateAuthenticatorRequest
*/
func (a *AuthenticatorAPIService) ActivateAuthenticator(ctx context.Context, authenticatorId string) ApiActivateAuthenticatorRequest {
	return ApiActivateAuthenticatorRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthenticatorBase
func (a *AuthenticatorAPIService) ActivateAuthenticatorExecute(r ApiActivateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticatorBase
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ActivateAuthenticator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)

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

type ApiActivateAuthenticatorMethodRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	methodType string
	retryCount int32
}

func (r ApiActivateAuthenticatorMethodRequest) Execute() (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	return r.ApiService.ActivateAuthenticatorMethodExecute(r)
}

/*
ActivateAuthenticatorMethod Activate an Authenticator Method

Activates a Method for an Authenticator identified by `authenticatorId` and `methodType`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @param methodType Type of authenticator method
 @return ApiActivateAuthenticatorMethodRequest
*/
func (a *AuthenticatorAPIService) ActivateAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiActivateAuthenticatorMethodRequest {
	return ApiActivateAuthenticatorMethodRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		methodType: methodType,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ListAuthenticatorMethods200ResponseInner
func (a *AuthenticatorAPIService) ActivateAuthenticatorMethodExecute(r ApiActivateAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAuthenticatorMethods200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ActivateAuthenticatorMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"methodType"+"}", url.PathEscape(parameterToString(r.methodType, "")), -1)

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

type ApiCreateAuthenticatorRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticator *AuthenticatorBase
	activate *bool
	retryCount int32
}

func (r ApiCreateAuthenticatorRequest) Authenticator(authenticator AuthenticatorBase) ApiCreateAuthenticatorRequest {
	r.authenticator = &authenticator
	return r
}

// Whether to execute the activation lifecycle operation when Okta creates the authenticator
func (r ApiCreateAuthenticatorRequest) Activate(activate bool) ApiCreateAuthenticatorRequest {
	r.activate = &activate
	return r
}

func (r ApiCreateAuthenticatorRequest) Execute() (*AuthenticatorBase, *APIResponse, error) {
	return r.ApiService.CreateAuthenticatorExecute(r)
}

/*
CreateAuthenticator Create an Authenticator

Creates an authenticator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAuthenticatorRequest
*/
func (a *AuthenticatorAPIService) CreateAuthenticator(ctx context.Context) ApiCreateAuthenticatorRequest {
	return ApiCreateAuthenticatorRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthenticatorBase
func (a *AuthenticatorAPIService) CreateAuthenticatorExecute(r ApiCreateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticatorBase
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.CreateAuthenticator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authenticator == nil {
		return localVarReturnValue, nil, reportError("authenticator is required and must be specified")
	}

	if r.activate != nil {
		localVarQueryParams.Add("activate", parameterToString(*r.activate, ""))
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
	localVarPostBody = r.authenticator
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

type ApiDeactivateAuthenticatorRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	retryCount int32
}

func (r ApiDeactivateAuthenticatorRequest) Execute() (*AuthenticatorBase, *APIResponse, error) {
	return r.ApiService.DeactivateAuthenticatorExecute(r)
}

/*
DeactivateAuthenticator Deactivate an Authenticator

Deactivates an authenticator by `authenticatorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @return ApiDeactivateAuthenticatorRequest
*/
func (a *AuthenticatorAPIService) DeactivateAuthenticator(ctx context.Context, authenticatorId string) ApiDeactivateAuthenticatorRequest {
	return ApiDeactivateAuthenticatorRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthenticatorBase
func (a *AuthenticatorAPIService) DeactivateAuthenticatorExecute(r ApiDeactivateAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticatorBase
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.DeactivateAuthenticator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)

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

type ApiDeactivateAuthenticatorMethodRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	methodType string
	retryCount int32
}

func (r ApiDeactivateAuthenticatorMethodRequest) Execute() (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	return r.ApiService.DeactivateAuthenticatorMethodExecute(r)
}

/*
DeactivateAuthenticatorMethod Deactivate an Authenticator Method

Deactivates a Method for an Authenticator identified by `authenticatorId` and `methodType`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @param methodType Type of authenticator method
 @return ApiDeactivateAuthenticatorMethodRequest
*/
func (a *AuthenticatorAPIService) DeactivateAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiDeactivateAuthenticatorMethodRequest {
	return ApiDeactivateAuthenticatorMethodRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		methodType: methodType,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ListAuthenticatorMethods200ResponseInner
func (a *AuthenticatorAPIService) DeactivateAuthenticatorMethodExecute(r ApiDeactivateAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAuthenticatorMethods200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.DeactivateAuthenticatorMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/methods/{methodType}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"methodType"+"}", url.PathEscape(parameterToString(r.methodType, "")), -1)

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

type ApiGetAuthenticatorRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	retryCount int32
}

func (r ApiGetAuthenticatorRequest) Execute() (*AuthenticatorBase, *APIResponse, error) {
	return r.ApiService.GetAuthenticatorExecute(r)
}

/*
GetAuthenticator Retrieve an Authenticator

Retrieves an authenticator from your Okta organization by `authenticatorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @return ApiGetAuthenticatorRequest
*/
func (a *AuthenticatorAPIService) GetAuthenticator(ctx context.Context, authenticatorId string) ApiGetAuthenticatorRequest {
	return ApiGetAuthenticatorRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthenticatorBase
func (a *AuthenticatorAPIService) GetAuthenticatorExecute(r ApiGetAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticatorBase
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.GetAuthenticator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)

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

type ApiGetAuthenticatorMethodRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	methodType string
	retryCount int32
}

func (r ApiGetAuthenticatorMethodRequest) Execute() (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	return r.ApiService.GetAuthenticatorMethodExecute(r)
}

/*
GetAuthenticatorMethod Retrieve an Authenticator Method

Retrieves a Method identified by `methodType` of an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @param methodType Type of authenticator method
 @return ApiGetAuthenticatorMethodRequest
*/
func (a *AuthenticatorAPIService) GetAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiGetAuthenticatorMethodRequest {
	return ApiGetAuthenticatorMethodRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		methodType: methodType,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ListAuthenticatorMethods200ResponseInner
func (a *AuthenticatorAPIService) GetAuthenticatorMethodExecute(r ApiGetAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAuthenticatorMethods200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.GetAuthenticatorMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/methods/{methodType}"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"methodType"+"}", url.PathEscape(parameterToString(r.methodType, "")), -1)

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

type ApiGetWellKnownAppAuthenticatorConfigurationRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	oauthClientId *string
	retryCount int32
}

// Filters app authenticator configurations by &#x60;oauthClientId&#x60;
func (r ApiGetWellKnownAppAuthenticatorConfigurationRequest) OauthClientId(oauthClientId string) ApiGetWellKnownAppAuthenticatorConfigurationRequest {
	r.oauthClientId = &oauthClientId
	return r
}

func (r ApiGetWellKnownAppAuthenticatorConfigurationRequest) Execute() ([]WellKnownAppAuthenticatorConfiguration, *APIResponse, error) {
	return r.ApiService.GetWellKnownAppAuthenticatorConfigurationExecute(r)
}

/*
GetWellKnownAppAuthenticatorConfiguration Retrieve the Well-Known App Authenticator Configuration

Retrieves the well-known app authenticator configuration. Includes an app authenticator's settings, supported methods, and other details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWellKnownAppAuthenticatorConfigurationRequest
*/
func (a *AuthenticatorAPIService) GetWellKnownAppAuthenticatorConfiguration(ctx context.Context) ApiGetWellKnownAppAuthenticatorConfigurationRequest {
	return ApiGetWellKnownAppAuthenticatorConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []WellKnownAppAuthenticatorConfiguration
func (a *AuthenticatorAPIService) GetWellKnownAppAuthenticatorConfigurationExecute(r ApiGetWellKnownAppAuthenticatorConfigurationRequest) ([]WellKnownAppAuthenticatorConfiguration, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WellKnownAppAuthenticatorConfiguration
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.GetWellKnownAppAuthenticatorConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/.well-known/app-authenticator-configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oauthClientId == nil {
		return localVarReturnValue, nil, reportError("oauthClientId is required and must be specified")
	}

	localVarQueryParams.Add("oauthClientId", parameterToString(*r.oauthClientId, ""))
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

type ApiListAuthenticatorMethodsRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	retryCount int32
}

func (r ApiListAuthenticatorMethodsRequest) Execute() ([]ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	return r.ApiService.ListAuthenticatorMethodsExecute(r)
}

/*
ListAuthenticatorMethods List all Methods of an Authenticator

Lists all Methods of an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @return ApiListAuthenticatorMethodsRequest
*/
func (a *AuthenticatorAPIService) ListAuthenticatorMethods(ctx context.Context, authenticatorId string) ApiListAuthenticatorMethodsRequest {
	return ApiListAuthenticatorMethodsRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []ListAuthenticatorMethods200ResponseInner
func (a *AuthenticatorAPIService) ListAuthenticatorMethodsExecute(r ApiListAuthenticatorMethodsRequest) ([]ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListAuthenticatorMethods200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ListAuthenticatorMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/methods"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)

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

type ApiListAuthenticatorsRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	retryCount int32
}

func (r ApiListAuthenticatorsRequest) Execute() ([]ListAuthenticators200ResponseInner, *APIResponse, error) {
	return r.ApiService.ListAuthenticatorsExecute(r)
}

/*
ListAuthenticators List all Authenticators

Lists all authenticators

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAuthenticatorsRequest
*/
func (a *AuthenticatorAPIService) ListAuthenticators(ctx context.Context) ApiListAuthenticatorsRequest {
	return ApiListAuthenticatorsRequest{
		ApiService: a,
		ctx: ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []ListAuthenticators200ResponseInner
func (a *AuthenticatorAPIService) ListAuthenticatorsExecute(r ApiListAuthenticatorsRequest) ([]ListAuthenticators200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListAuthenticators200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ListAuthenticators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators"

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

type ApiReplaceAuthenticatorRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	authenticator *AuthenticatorBase
	retryCount int32
}

func (r ApiReplaceAuthenticatorRequest) Authenticator(authenticator AuthenticatorBase) ApiReplaceAuthenticatorRequest {
	r.authenticator = &authenticator
	return r
}

func (r ApiReplaceAuthenticatorRequest) Execute() (*AuthenticatorBase, *APIResponse, error) {
	return r.ApiService.ReplaceAuthenticatorExecute(r)
}

/*
ReplaceAuthenticator Replace an Authenticator

Replaces the properties for an Authenticator identified by `authenticatorId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @return ApiReplaceAuthenticatorRequest
*/
func (a *AuthenticatorAPIService) ReplaceAuthenticator(ctx context.Context, authenticatorId string) ApiReplaceAuthenticatorRequest {
	return ApiReplaceAuthenticatorRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthenticatorBase
func (a *AuthenticatorAPIService) ReplaceAuthenticatorExecute(r ApiReplaceAuthenticatorRequest) (*AuthenticatorBase, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticatorBase
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ReplaceAuthenticator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authenticator == nil {
		return localVarReturnValue, nil, reportError("authenticator is required and must be specified")
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
	localVarPostBody = r.authenticator
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

type ApiReplaceAuthenticatorMethodRequest struct {
	ctx context.Context
	ApiService AuthenticatorAPI
	authenticatorId string
	methodType string
	listAuthenticatorMethods200ResponseInner *ListAuthenticatorMethods200ResponseInner
	retryCount int32
}

func (r ApiReplaceAuthenticatorMethodRequest) ListAuthenticatorMethods200ResponseInner(listAuthenticatorMethods200ResponseInner ListAuthenticatorMethods200ResponseInner) ApiReplaceAuthenticatorMethodRequest {
	r.listAuthenticatorMethods200ResponseInner = &listAuthenticatorMethods200ResponseInner
	return r
}

func (r ApiReplaceAuthenticatorMethodRequest) Execute() (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	return r.ApiService.ReplaceAuthenticatorMethodExecute(r)
}

/*
ReplaceAuthenticatorMethod Replace an Authenticator Method

Replaces a Method of `methodType` for an Authenticator identified by `authenticatorId`
> **Note:** <x-lifecycle class="ea"></x-lifecycle>
> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.
> This feature has several limitations when enrolling a security key:
> - Enrollment is currently unsupported on Firefox.
> - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.
> - If prompted during enrollment, users must allow Okta to see the make and model of the security key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authenticatorId `id` of the Authenticator
 @param methodType Type of authenticator method
 @return ApiReplaceAuthenticatorMethodRequest
*/
func (a *AuthenticatorAPIService) ReplaceAuthenticatorMethod(ctx context.Context, authenticatorId string, methodType string) ApiReplaceAuthenticatorMethodRequest {
	return ApiReplaceAuthenticatorMethodRequest{
		ApiService: a,
		ctx: ctx,
		authenticatorId: authenticatorId,
		methodType: methodType,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ListAuthenticatorMethods200ResponseInner
func (a *AuthenticatorAPIService) ReplaceAuthenticatorMethodExecute(r ApiReplaceAuthenticatorMethodRequest) (*ListAuthenticatorMethods200ResponseInner, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAuthenticatorMethods200ResponseInner
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticatorAPIService.ReplaceAuthenticatorMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authenticators/{authenticatorId}/methods/{methodType}"
	localVarPath = strings.Replace(localVarPath, "{"+"authenticatorId"+"}", url.PathEscape(parameterToString(r.authenticatorId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"methodType"+"}", url.PathEscape(parameterToString(r.methodType, "")), -1)

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
	localVarPostBody = r.listAuthenticatorMethods200ResponseInner
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

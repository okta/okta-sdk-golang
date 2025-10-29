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

type IdentityProviderAPI interface {

	/*
		ActivateIdentityProvider Activate an IdP

		Activates an inactive identity provider (IdP)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param idpId `id` of IdP
		@return ApiActivateIdentityProviderRequest
	*/
	ActivateIdentityProvider(ctx context.Context, idpId string) ApiActivateIdentityProviderRequest

	// ActivateIdentityProviderExecute executes the request
	//  @return IdentityProvider
	ActivateIdentityProviderExecute(r ApiActivateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error)

	/*
			CreateIdentityProvider Create an IdP

			Creates a new identity provider (IdP) integration.

		#### SAML 2.0 IdP

		You must first add the IdP's signature certificate to the IdP key store before you can add a SAML 2.0 IdP with a `kid` credential reference.

		Don't use `fromURI` to automatically redirect a user to a particular app after successfully authenticating with a third-party IdP. Instead, use SAML deep links. Using `fromURI` isn't tested or supported. For more information about using deep links when signing users in using an SP-initiated flow, see [Understanding SP-Initiated Login flow](https://developer.okta.com/docs/concepts/saml/#understanding-sp-initiated-login-flow).

		Use SAML deep links to automatically redirect the user to an app after successfully authenticating with a third-party IdP. To use deep links, assemble these three parts into a URL:

		* SP ACS URL<br>
		For example: `https://${yourOktaDomain}/sso/saml2/:idpId`
		* The app to which the user is automatically redirected after successfully authenticating with the IdP <br>
		For example: `/app/:app-location/:appId/sso/saml`
		* Optionally, if the app is an outbound SAML app, you can specify the `relayState` passed to it.<br>
		For example: `?RelayState=:anyUrlEncodedValue`

		The deep link for the above three parts is:<br>
		`https://${yourOktaDomain}/sso/saml2/:idpId/app/:app-location/:appId/sso/saml?RelayState=:anyUrlEncodedValue`

		#### Smart Card X509 IdP

		You must first add the IdP's server certificate to the IdP key store before you can add a Smart Card `X509` IdP with a `kid` credential reference.
		You need to upload the whole trust chain as a single key using the [Key Store API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProviderKeys/#tag/IdentityProviderKeys/operation/createIdentityProviderKey).
		Depending on the information stored in the smart card, select the proper [template](https://developer.okta.com/docs/reference/okta-expression-language/#idp-user-profile) `idpuser.subjectAltNameEmail` or `idpuser.subjectAltNameUpn`.

		#### Identity verification vendors as identity providers

		Identity verification vendors (IDVs) work like IdPs, with a few key differences. IDVs verify your user's identities by requiring them to submit a proof of identity. There are many ways to verify user identities. For example, a proof of identity can be a selfie to determine liveliness or it can be requiring users to submit a photo of their driver's license and matching that information with a database.

		There are three IDVs that you can configure as IdPs in your org by creating an account with the vendor, and then creating an IdP integration. Control how the IDVs verify your users by using [Okta account management policy rules](https://developer.okta.com/docs/guides/okta-account-management-policy/main/).

		* [Persona](https://withpersona.com/)

		* [CLEAR Verified](https://www.clearme.com/)

		* [Incode](https://incode.com/)

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCreateIdentityProviderRequest
	*/
	CreateIdentityProvider(ctx context.Context) ApiCreateIdentityProviderRequest

	// CreateIdentityProviderExecute executes the request
	//  @return IdentityProvider
	CreateIdentityProviderExecute(r ApiCreateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error)

	/*
		DeactivateIdentityProvider Deactivate an IdP

		Deactivates an active identity provider (IdP)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param idpId `id` of IdP
		@return ApiDeactivateIdentityProviderRequest
	*/
	DeactivateIdentityProvider(ctx context.Context, idpId string) ApiDeactivateIdentityProviderRequest

	// DeactivateIdentityProviderExecute executes the request
	//  @return IdentityProvider
	DeactivateIdentityProviderExecute(r ApiDeactivateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error)

	/*
			DeleteIdentityProvider Delete an IdP

			Deletes an identity provider (IdP) integration by `idpId`
		* All existing IdP users are unlinked with the highest order profile source taking precedence for each IdP user.
		* Unlinked users keep their existing authentication provider such as `FEDERATION` or `SOCIAL`.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param idpId `id` of IdP
			@return ApiDeleteIdentityProviderRequest
	*/
	DeleteIdentityProvider(ctx context.Context, idpId string) ApiDeleteIdentityProviderRequest

	// DeleteIdentityProviderExecute executes the request
	DeleteIdentityProviderExecute(r ApiDeleteIdentityProviderRequest) (*APIResponse, error)

	/*
		GetIdentityProvider Retrieve an IdP

		Retrieves an identity provider (IdP) integration by `idpId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param idpId `id` of IdP
		@return ApiGetIdentityProviderRequest
	*/
	GetIdentityProvider(ctx context.Context, idpId string) ApiGetIdentityProviderRequest

	// GetIdentityProviderExecute executes the request
	//  @return IdentityProvider
	GetIdentityProviderExecute(r ApiGetIdentityProviderRequest) (*IdentityProvider, *APIResponse, error)

	/*
		ListIdentityProviders List all IdPs

		Lists all identity provider (IdP) integrations with pagination. A subset of IdPs can be returned that match a supported filter expression or query.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListIdentityProvidersRequest
	*/
	ListIdentityProviders(ctx context.Context) ApiListIdentityProvidersRequest

	// ListIdentityProvidersExecute executes the request
	//  @return []IdentityProvider
	ListIdentityProvidersExecute(r ApiListIdentityProvidersRequest) ([]IdentityProvider, *APIResponse, error)

	/*
		ReplaceIdentityProvider Replace an IdP

		Replaces an identity provider (IdP) integration by `idpId`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param idpId `id` of IdP
		@return ApiReplaceIdentityProviderRequest
	*/
	ReplaceIdentityProvider(ctx context.Context, idpId string) ApiReplaceIdentityProviderRequest

	// ReplaceIdentityProviderExecute executes the request
	//  @return IdentityProvider
	ReplaceIdentityProviderExecute(r ApiReplaceIdentityProviderRequest) (*IdentityProvider, *APIResponse, error)
}

// IdentityProviderAPIService IdentityProviderAPI service
type IdentityProviderAPIService service

type ApiActivateIdentityProviderRequest struct {
	ctx        context.Context
	ApiService IdentityProviderAPI
	idpId      string
	retryCount int32
}

func (r ApiActivateIdentityProviderRequest) Execute() (*IdentityProvider, *APIResponse, error) {
	return r.ApiService.ActivateIdentityProviderExecute(r)
}

/*
ActivateIdentityProvider Activate an IdP

Activates an inactive identity provider (IdP)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param idpId `id` of IdP
	@return ApiActivateIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) ActivateIdentityProvider(ctx context.Context, idpId string) ApiActivateIdentityProviderRequest {
	return ApiActivateIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		idpId:      idpId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return IdentityProvider
func (a *IdentityProviderAPIService) ActivateIdentityProviderExecute(r ApiActivateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.ActivateIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps/{idpId}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"idpId"+"}", url.PathEscape(parameterToString(r.idpId, "")), -1)

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

type ApiCreateIdentityProviderRequest struct {
	ctx              context.Context
	ApiService       IdentityProviderAPI
	identityProvider *IdentityProvider
	retryCount       int32
}

// IdP settings
func (r ApiCreateIdentityProviderRequest) IdentityProvider(identityProvider IdentityProvider) ApiCreateIdentityProviderRequest {
	r.identityProvider = &identityProvider
	return r
}

func (r ApiCreateIdentityProviderRequest) Execute() (*IdentityProvider, *APIResponse, error) {
	return r.ApiService.CreateIdentityProviderExecute(r)
}

/*
CreateIdentityProvider Create an IdP

Creates a new identity provider (IdP) integration.

#### SAML 2.0 IdP

You must first add the IdP's signature certificate to the IdP key store before you can add a SAML 2.0 IdP with a `kid` credential reference.

Don't use `fromURI` to automatically redirect a user to a particular app after successfully authenticating with a third-party IdP. Instead, use SAML deep links. Using `fromURI` isn't tested or supported. For more information about using deep links when signing users in using an SP-initiated flow, see [Understanding SP-Initiated Login flow](https://developer.okta.com/docs/concepts/saml/#understanding-sp-initiated-login-flow).

Use SAML deep links to automatically redirect the user to an app after successfully authenticating with a third-party IdP. To use deep links, assemble these three parts into a URL:

* SP ACS URL<br>
For example: `https://${yourOktaDomain}/sso/saml2/:idpId`
* The app to which the user is automatically redirected after successfully authenticating with the IdP <br>
For example: `/app/:app-location/:appId/sso/saml`
* Optionally, if the app is an outbound SAML app, you can specify the `relayState` passed to it.<br>
For example: `?RelayState=:anyUrlEncodedValue`

The deep link for the above three parts is:<br>
`https://${yourOktaDomain}/sso/saml2/:idpId/app/:app-location/:appId/sso/saml?RelayState=:anyUrlEncodedValue`

#### Smart Card X509 IdP

You must first add the IdP's server certificate to the IdP key store before you can add a Smart Card `X509` IdP with a `kid` credential reference.
You need to upload the whole trust chain as a single key using the [Key Store API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProviderKeys/#tag/IdentityProviderKeys/operation/createIdentityProviderKey).
Depending on the information stored in the smart card, select the proper [template](https://developer.okta.com/docs/reference/okta-expression-language/#idp-user-profile) `idpuser.subjectAltNameEmail` or `idpuser.subjectAltNameUpn`.

#### Identity verification vendors as identity providers

Identity verification vendors (IDVs) work like IdPs, with a few key differences. IDVs verify your user's identities by requiring them to submit a proof of identity. There are many ways to verify user identities. For example, a proof of identity can be a selfie to determine liveliness or it can be requiring users to submit a photo of their driver's license and matching that information with a database.

There are three IDVs that you can configure as IdPs in your org by creating an account with the vendor, and then creating an IdP integration. Control how the IDVs verify your users by using [Okta account management policy rules](https://developer.okta.com/docs/guides/okta-account-management-policy/main/).

* [Persona](https://withpersona.com/)

* [CLEAR Verified](https://www.clearme.com/)

* [Incode](https://incode.com/)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) CreateIdentityProvider(ctx context.Context) ApiCreateIdentityProviderRequest {
	return ApiCreateIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return IdentityProvider
func (a *IdentityProviderAPIService) CreateIdentityProviderExecute(r ApiCreateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.CreateIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identityProvider == nil {
		return localVarReturnValue, nil, reportError("identityProvider is required and must be specified")
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
	localVarPostBody = r.identityProvider
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

type ApiDeactivateIdentityProviderRequest struct {
	ctx        context.Context
	ApiService IdentityProviderAPI
	idpId      string
	retryCount int32
}

func (r ApiDeactivateIdentityProviderRequest) Execute() (*IdentityProvider, *APIResponse, error) {
	return r.ApiService.DeactivateIdentityProviderExecute(r)
}

/*
DeactivateIdentityProvider Deactivate an IdP

Deactivates an active identity provider (IdP)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param idpId `id` of IdP
	@return ApiDeactivateIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) DeactivateIdentityProvider(ctx context.Context, idpId string) ApiDeactivateIdentityProviderRequest {
	return ApiDeactivateIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		idpId:      idpId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return IdentityProvider
func (a *IdentityProviderAPIService) DeactivateIdentityProviderExecute(r ApiDeactivateIdentityProviderRequest) (*IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.DeactivateIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps/{idpId}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"idpId"+"}", url.PathEscape(parameterToString(r.idpId, "")), -1)

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

type ApiDeleteIdentityProviderRequest struct {
	ctx        context.Context
	ApiService IdentityProviderAPI
	idpId      string
	retryCount int32
}

func (r ApiDeleteIdentityProviderRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteIdentityProviderExecute(r)
}

/*
DeleteIdentityProvider Delete an IdP

Deletes an identity provider (IdP) integration by `idpId`
* All existing IdP users are unlinked with the highest order profile source taking precedence for each IdP user.
* Unlinked users keep their existing authentication provider such as `FEDERATION` or `SOCIAL`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param idpId `id` of IdP
	@return ApiDeleteIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) DeleteIdentityProvider(ctx context.Context, idpId string) ApiDeleteIdentityProviderRequest {
	return ApiDeleteIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		idpId:      idpId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *IdentityProviderAPIService) DeleteIdentityProviderExecute(r ApiDeleteIdentityProviderRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.DeleteIdentityProvider")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps/{idpId}"
	localVarPath = strings.Replace(localVarPath, "{"+"idpId"+"}", url.PathEscape(parameterToString(r.idpId, "")), -1)

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

type ApiGetIdentityProviderRequest struct {
	ctx        context.Context
	ApiService IdentityProviderAPI
	idpId      string
	retryCount int32
}

func (r ApiGetIdentityProviderRequest) Execute() (*IdentityProvider, *APIResponse, error) {
	return r.ApiService.GetIdentityProviderExecute(r)
}

/*
GetIdentityProvider Retrieve an IdP

Retrieves an identity provider (IdP) integration by `idpId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param idpId `id` of IdP
	@return ApiGetIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) GetIdentityProvider(ctx context.Context, idpId string) ApiGetIdentityProviderRequest {
	return ApiGetIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		idpId:      idpId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return IdentityProvider
func (a *IdentityProviderAPIService) GetIdentityProviderExecute(r ApiGetIdentityProviderRequest) (*IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.GetIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps/{idpId}"
	localVarPath = strings.Replace(localVarPath, "{"+"idpId"+"}", url.PathEscape(parameterToString(r.idpId, "")), -1)

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

type ApiListIdentityProvidersRequest struct {
	ctx        context.Context
	ApiService IdentityProviderAPI
	q          *string
	after      *string
	limit      *int32
	type_      *string
	retryCount int32
}

// Searches the &#x60;name&#x60; property of IdPs for matching value
func (r ApiListIdentityProvidersRequest) Q(q string) ApiListIdentityProvidersRequest {
	r.q = &q
	return r
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListIdentityProvidersRequest) After(after string) ApiListIdentityProvidersRequest {
	r.after = &after
	return r
}

// A limit on the number of objects to return
func (r ApiListIdentityProvidersRequest) Limit(limit int32) ApiListIdentityProvidersRequest {
	r.limit = &limit
	return r
}

// Filters IdPs by &#x60;type&#x60;
func (r ApiListIdentityProvidersRequest) Type_(type_ string) ApiListIdentityProvidersRequest {
	r.type_ = &type_
	return r
}

func (r ApiListIdentityProvidersRequest) Execute() ([]IdentityProvider, *APIResponse, error) {
	return r.ApiService.ListIdentityProvidersExecute(r)
}

/*
ListIdentityProviders List all IdPs

Lists all identity provider (IdP) integrations with pagination. A subset of IdPs can be returned that match a supported filter expression or query.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListIdentityProvidersRequest
*/
func (a *IdentityProviderAPIService) ListIdentityProviders(ctx context.Context) ApiListIdentityProvidersRequest {
	return ApiListIdentityProvidersRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []IdentityProvider
func (a *IdentityProviderAPIService) ListIdentityProvidersExecute(r ApiListIdentityProvidersRequest) ([]IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.ListIdentityProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
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

type ApiReplaceIdentityProviderRequest struct {
	ctx              context.Context
	ApiService       IdentityProviderAPI
	idpId            string
	identityProvider *IdentityProvider
	retryCount       int32
}

// Updated configuration for the IdP
func (r ApiReplaceIdentityProviderRequest) IdentityProvider(identityProvider IdentityProvider) ApiReplaceIdentityProviderRequest {
	r.identityProvider = &identityProvider
	return r
}

func (r ApiReplaceIdentityProviderRequest) Execute() (*IdentityProvider, *APIResponse, error) {
	return r.ApiService.ReplaceIdentityProviderExecute(r)
}

/*
ReplaceIdentityProvider Replace an IdP

Replaces an identity provider (IdP) integration by `idpId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param idpId `id` of IdP
	@return ApiReplaceIdentityProviderRequest
*/
func (a *IdentityProviderAPIService) ReplaceIdentityProvider(ctx context.Context, idpId string) ApiReplaceIdentityProviderRequest {
	return ApiReplaceIdentityProviderRequest{
		ApiService: a,
		ctx:        ctx,
		idpId:      idpId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return IdentityProvider
func (a *IdentityProviderAPIService) ReplaceIdentityProviderExecute(r ApiReplaceIdentityProviderRequest) (*IdentityProvider, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentityProvider
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProviderAPIService.ReplaceIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/idps/{idpId}"
	localVarPath = strings.Replace(localVarPath, "{"+"idpId"+"}", url.PathEscape(parameterToString(r.idpId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.identityProvider == nil {
		return localVarReturnValue, nil, reportError("identityProvider is required and must be specified")
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
	localVarPostBody = r.identityProvider
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

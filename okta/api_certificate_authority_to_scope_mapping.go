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
	"time"
)

type CertificateAuthorityToScopeMappingAPI interface {

	/*
			ListCertificateAuthoritiesForScope List all Certificate Authorities for a scope

			Lists all Certificate Authorities (CAs) eligible for the requested scope, annotated with the current mapping for that scope (if any).

		Each item represents a CA that's compatible with the requested scope. The `mapping` field is populated when the CA is currently mapped to the scope and is omitted otherwise.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListCertificateAuthoritiesForScopeRequest
	*/
	ListCertificateAuthoritiesForScope(ctx context.Context) ApiListCertificateAuthoritiesForScopeRequest

	// ListCertificateAuthoritiesForScopeExecute executes the request
	//  @return []CertificateAuthorityToScopeMappingResponse
	ListCertificateAuthoritiesForScopeExecute(r ApiListCertificateAuthoritiesForScopeRequest) ([]CertificateAuthorityToScopeMappingResponse, *APIResponse, error)

	/*
			ReplaceCertificateAuthorityMappings Replace all Certificate Authority mappings for a scope

			Replaces the org's Certificate Authority (CA) mappings for the specified scope with the provided list. Only mappings belonging to the given scope are touched — other scopes are unaffected.

		This is a scope-based full reconcile: any existing mappings for the scope that are absent from the request body are deleted, and any CA IDs in the body that aren't already mapped to the scope are inserted. Sending an empty array clears all mappings for the scope without affecting other scopes.

		The response contains the resulting active mappings for the scope. Clients that need the full enriched CA view after an update should issue a subsequent `GET /device-identity/api/v1/certificate-authority-mappings/certificate-authorities`.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiReplaceCertificateAuthorityMappingsRequest
	*/
	ReplaceCertificateAuthorityMappings(ctx context.Context) ApiReplaceCertificateAuthorityMappingsRequest

	// ReplaceCertificateAuthorityMappingsExecute executes the request
	//  @return []CertificateAuthorityToScopeMapping
	ReplaceCertificateAuthorityMappingsExecute(r ApiReplaceCertificateAuthorityMappingsRequest) ([]CertificateAuthorityToScopeMapping, *APIResponse, error)
}

// CertificateAuthorityToScopeMappingAPIService CertificateAuthorityToScopeMappingAPI service
type CertificateAuthorityToScopeMappingAPIService service

type ApiListCertificateAuthoritiesForScopeRequest struct {
	ctx        context.Context
	ApiService CertificateAuthorityToScopeMappingAPI
	scope      *string
	retryCount int32
}

// The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected.
func (r ApiListCertificateAuthoritiesForScopeRequest) Scope(scope string) ApiListCertificateAuthoritiesForScopeRequest {
	r.scope = &scope
	return r
}

func (r ApiListCertificateAuthoritiesForScopeRequest) Execute() ([]CertificateAuthorityToScopeMappingResponse, *APIResponse, error) {
	return r.ApiService.ListCertificateAuthoritiesForScopeExecute(r)
}

/*
ListCertificateAuthoritiesForScope List all Certificate Authorities for a scope

Lists all Certificate Authorities (CAs) eligible for the requested scope, annotated with the current mapping for that scope (if any).

Each item represents a CA that's compatible with the requested scope. The `mapping` field is populated when the CA is currently mapped to the scope and is omitted otherwise.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListCertificateAuthoritiesForScopeRequest
*/
func (a *CertificateAuthorityToScopeMappingAPIService) ListCertificateAuthoritiesForScope(ctx context.Context) ApiListCertificateAuthoritiesForScopeRequest {
	return ApiListCertificateAuthoritiesForScopeRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []CertificateAuthorityToScopeMappingResponse
func (a *CertificateAuthorityToScopeMappingAPIService) ListCertificateAuthoritiesForScopeExecute(r ApiListCertificateAuthoritiesForScopeRequest) ([]CertificateAuthorityToScopeMappingResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CertificateAuthorityToScopeMappingResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAuthorityToScopeMappingAPIService.ListCertificateAuthoritiesForScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-identity/api/v1/certificate-authority-mappings/certificate-authorities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.scope == nil {
		return localVarReturnValue, nil, reportError("scope is required and must be specified")
	}

	localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
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

type ApiReplaceCertificateAuthorityMappingsRequest struct {
	ctx        context.Context
	ApiService CertificateAuthorityToScopeMappingAPI
	scope      *string
	mappings   *[]CertificateAuthorityToScopeMappingRequest
	retryCount int32
}

// The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected.
func (r ApiReplaceCertificateAuthorityMappingsRequest) Scope(scope string) ApiReplaceCertificateAuthorityMappingsRequest {
	r.scope = &scope
	return r
}

func (r ApiReplaceCertificateAuthorityMappingsRequest) Mappings(mappings []CertificateAuthorityToScopeMappingRequest) ApiReplaceCertificateAuthorityMappingsRequest {
	r.mappings = &mappings
	return r
}

func (r ApiReplaceCertificateAuthorityMappingsRequest) Execute() ([]CertificateAuthorityToScopeMapping, *APIResponse, error) {
	return r.ApiService.ReplaceCertificateAuthorityMappingsExecute(r)
}

/*
ReplaceCertificateAuthorityMappings Replace all Certificate Authority mappings for a scope

Replaces the org's Certificate Authority (CA) mappings for the specified scope with the provided list. Only mappings belonging to the given scope are touched — other scopes are unaffected.

This is a scope-based full reconcile: any existing mappings for the scope that are absent from the request body are deleted, and any CA IDs in the body that aren't already mapped to the scope are inserted. Sending an empty array clears all mappings for the scope without affecting other scopes.

The response contains the resulting active mappings for the scope. Clients that need the full enriched CA view after an update should issue a subsequent `GET /device-identity/api/v1/certificate-authority-mappings/certificate-authorities`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReplaceCertificateAuthorityMappingsRequest
*/
func (a *CertificateAuthorityToScopeMappingAPIService) ReplaceCertificateAuthorityMappings(ctx context.Context) ApiReplaceCertificateAuthorityMappingsRequest {
	return ApiReplaceCertificateAuthorityMappingsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []CertificateAuthorityToScopeMapping
func (a *CertificateAuthorityToScopeMappingAPIService) ReplaceCertificateAuthorityMappingsExecute(r ApiReplaceCertificateAuthorityMappingsRequest) ([]CertificateAuthorityToScopeMapping, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CertificateAuthorityToScopeMapping
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateAuthorityToScopeMappingAPIService.ReplaceCertificateAuthorityMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device-identity/api/v1/certificate-authority-mappings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.scope == nil {
		return localVarReturnValue, nil, reportError("scope is required and must be specified")
	}
	if r.mappings == nil {
		return localVarReturnValue, nil, reportError("mappings is required and must be specified")
	}

	localVarQueryParams.Add("scope", parameterToString(*r.scope, ""))
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
	localVarPostBody = r.mappings
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

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


type AuthorizationServerRulesAPI interface {

	/*
	ActivateAuthorizationServerPolicyRule Activate a Policy Rule

	Activates an authorization server policy rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@param ruleId `id` of the Policy Rule
	@return ApiActivateAuthorizationServerPolicyRuleRequest
	*/
	ActivateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiActivateAuthorizationServerPolicyRuleRequest

	// ActivateAuthorizationServerPolicyRuleExecute executes the request
	ActivateAuthorizationServerPolicyRuleExecute(r ApiActivateAuthorizationServerPolicyRuleRequest) (*APIResponse, error)

	/*
	CreateAuthorizationServerPolicyRule Create a Policy Rule

	Creates a policy rule for the specified Custom Authorization Server and Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@return ApiCreateAuthorizationServerPolicyRuleRequest
	*/
	CreateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string) ApiCreateAuthorizationServerPolicyRuleRequest

	// CreateAuthorizationServerPolicyRuleExecute executes the request
	//  @return AuthorizationServerPolicyRule
	CreateAuthorizationServerPolicyRuleExecute(r ApiCreateAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error)

	/*
	DeactivateAuthorizationServerPolicyRule Deactivate a Policy Rule

	Deactivates an authorization server policy rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@param ruleId `id` of the Policy Rule
	@return ApiDeactivateAuthorizationServerPolicyRuleRequest
	*/
	DeactivateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiDeactivateAuthorizationServerPolicyRuleRequest

	// DeactivateAuthorizationServerPolicyRuleExecute executes the request
	DeactivateAuthorizationServerPolicyRuleExecute(r ApiDeactivateAuthorizationServerPolicyRuleRequest) (*APIResponse, error)

	/*
	DeleteAuthorizationServerPolicyRule Delete a Policy Rule

	Deletes a Policy Rule defined in the specified Custom Authorization Server and Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@param ruleId `id` of the Policy Rule
	@return ApiDeleteAuthorizationServerPolicyRuleRequest
	*/
	DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiDeleteAuthorizationServerPolicyRuleRequest

	// DeleteAuthorizationServerPolicyRuleExecute executes the request
	DeleteAuthorizationServerPolicyRuleExecute(r ApiDeleteAuthorizationServerPolicyRuleRequest) (*APIResponse, error)

	/*
	GetAuthorizationServerPolicyRule Retrieve a Policy Rule

	Retrieves a policy rule by `ruleId`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@param ruleId `id` of the Policy Rule
	@return ApiGetAuthorizationServerPolicyRuleRequest
	*/
	GetAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiGetAuthorizationServerPolicyRuleRequest

	// GetAuthorizationServerPolicyRuleExecute executes the request
	//  @return AuthorizationServerPolicyRule
	GetAuthorizationServerPolicyRuleExecute(r ApiGetAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error)

	/*
	ListAuthorizationServerPolicyRules List all Policy Rules

	Lists all policy rules for the specified Custom Authorization Server and Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@return ApiListAuthorizationServerPolicyRulesRequest
	*/
	ListAuthorizationServerPolicyRules(ctx context.Context, authServerId string, policyId string) ApiListAuthorizationServerPolicyRulesRequest

	// ListAuthorizationServerPolicyRulesExecute executes the request
	//  @return []AuthorizationServerPolicyRule
	ListAuthorizationServerPolicyRulesExecute(r ApiListAuthorizationServerPolicyRulesRequest) ([]AuthorizationServerPolicyRule, *APIResponse, error)

	/*
	ReplaceAuthorizationServerPolicyRule Replace a Policy Rule

	Replaces the configuration of the Policy Rule defined in the specified Custom Authorization Server and Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authServerId `id` of the Authorization Server
	@param policyId `id` of the Policy
	@param ruleId `id` of the Policy Rule
	@return ApiReplaceAuthorizationServerPolicyRuleRequest
	*/
	ReplaceAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiReplaceAuthorizationServerPolicyRuleRequest

	// ReplaceAuthorizationServerPolicyRuleExecute executes the request
	//  @return AuthorizationServerPolicyRule
	ReplaceAuthorizationServerPolicyRuleExecute(r ApiReplaceAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error)
}

// AuthorizationServerRulesAPIService AuthorizationServerRulesAPI service
type AuthorizationServerRulesAPIService service

type ApiActivateAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	ruleId string
	retryCount int32
}

func (r ApiActivateAuthorizationServerPolicyRuleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ActivateAuthorizationServerPolicyRuleExecute(r)
}

/*
ActivateAuthorizationServerPolicyRule Activate a Policy Rule

Activates an authorization server policy rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @param ruleId `id` of the Policy Rule
 @return ApiActivateAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) ActivateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiActivateAuthorizationServerPolicyRuleRequest {
	return ApiActivateAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		ruleId: ruleId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *AuthorizationServerRulesAPIService) ActivateAuthorizationServerPolicyRuleExecute(r ApiActivateAuthorizationServerPolicyRuleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.ActivateAuthorizationServerPolicyRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterToString(r.ruleId, "")), -1)

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

type ApiCreateAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	policyRule *AuthorizationServerPolicyRule
	retryCount int32
}

func (r ApiCreateAuthorizationServerPolicyRuleRequest) PolicyRule(policyRule AuthorizationServerPolicyRule) ApiCreateAuthorizationServerPolicyRuleRequest {
	r.policyRule = &policyRule
	return r
}

func (r ApiCreateAuthorizationServerPolicyRuleRequest) Execute() (*AuthorizationServerPolicyRule, *APIResponse, error) {
	return r.ApiService.CreateAuthorizationServerPolicyRuleExecute(r)
}

/*
CreateAuthorizationServerPolicyRule Create a Policy Rule

Creates a policy rule for the specified Custom Authorization Server and Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @return ApiCreateAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) CreateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string) ApiCreateAuthorizationServerPolicyRuleRequest {
	return ApiCreateAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthorizationServerPolicyRule
func (a *AuthorizationServerRulesAPIService) CreateAuthorizationServerPolicyRuleExecute(r ApiCreateAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationServerPolicyRule
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.CreateAuthorizationServerPolicyRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.policyRule == nil {
		return localVarReturnValue, nil, reportError("policyRule is required and must be specified")
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
	localVarPostBody = r.policyRule
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

type ApiDeactivateAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	ruleId string
	retryCount int32
}

func (r ApiDeactivateAuthorizationServerPolicyRuleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeactivateAuthorizationServerPolicyRuleExecute(r)
}

/*
DeactivateAuthorizationServerPolicyRule Deactivate a Policy Rule

Deactivates an authorization server policy rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @param ruleId `id` of the Policy Rule
 @return ApiDeactivateAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) DeactivateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiDeactivateAuthorizationServerPolicyRuleRequest {
	return ApiDeactivateAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		ruleId: ruleId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *AuthorizationServerRulesAPIService) DeactivateAuthorizationServerPolicyRuleExecute(r ApiDeactivateAuthorizationServerPolicyRuleRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.DeactivateAuthorizationServerPolicyRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterToString(r.ruleId, "")), -1)

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

type ApiDeleteAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	ruleId string
	retryCount int32
}

func (r ApiDeleteAuthorizationServerPolicyRuleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteAuthorizationServerPolicyRuleExecute(r)
}

/*
DeleteAuthorizationServerPolicyRule Delete a Policy Rule

Deletes a Policy Rule defined in the specified Custom Authorization Server and Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @param ruleId `id` of the Policy Rule
 @return ApiDeleteAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiDeleteAuthorizationServerPolicyRuleRequest {
	return ApiDeleteAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		ruleId: ruleId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *AuthorizationServerRulesAPIService) DeleteAuthorizationServerPolicyRuleExecute(r ApiDeleteAuthorizationServerPolicyRuleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.DeleteAuthorizationServerPolicyRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterToString(r.ruleId, "")), -1)

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

type ApiGetAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	ruleId string
	retryCount int32
}

func (r ApiGetAuthorizationServerPolicyRuleRequest) Execute() (*AuthorizationServerPolicyRule, *APIResponse, error) {
	return r.ApiService.GetAuthorizationServerPolicyRuleExecute(r)
}

/*
GetAuthorizationServerPolicyRule Retrieve a Policy Rule

Retrieves a policy rule by `ruleId`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @param ruleId `id` of the Policy Rule
 @return ApiGetAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) GetAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiGetAuthorizationServerPolicyRuleRequest {
	return ApiGetAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		ruleId: ruleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthorizationServerPolicyRule
func (a *AuthorizationServerRulesAPIService) GetAuthorizationServerPolicyRuleExecute(r ApiGetAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationServerPolicyRule
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.GetAuthorizationServerPolicyRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterToString(r.ruleId, "")), -1)

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

type ApiListAuthorizationServerPolicyRulesRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	retryCount int32
}

func (r ApiListAuthorizationServerPolicyRulesRequest) Execute() ([]AuthorizationServerPolicyRule, *APIResponse, error) {
	return r.ApiService.ListAuthorizationServerPolicyRulesExecute(r)
}

/*
ListAuthorizationServerPolicyRules List all Policy Rules

Lists all policy rules for the specified Custom Authorization Server and Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @return ApiListAuthorizationServerPolicyRulesRequest
*/
func (a *AuthorizationServerRulesAPIService) ListAuthorizationServerPolicyRules(ctx context.Context, authServerId string, policyId string) ApiListAuthorizationServerPolicyRulesRequest {
	return ApiListAuthorizationServerPolicyRulesRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []AuthorizationServerPolicyRule
func (a *AuthorizationServerRulesAPIService) ListAuthorizationServerPolicyRulesExecute(r ApiListAuthorizationServerPolicyRulesRequest) ([]AuthorizationServerPolicyRule, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuthorizationServerPolicyRule
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.ListAuthorizationServerPolicyRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)

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

type ApiReplaceAuthorizationServerPolicyRuleRequest struct {
	ctx context.Context
	ApiService AuthorizationServerRulesAPI
	authServerId string
	policyId string
	ruleId string
	policyRule *AuthorizationServerPolicyRule
	retryCount int32
}

func (r ApiReplaceAuthorizationServerPolicyRuleRequest) PolicyRule(policyRule AuthorizationServerPolicyRule) ApiReplaceAuthorizationServerPolicyRuleRequest {
	r.policyRule = &policyRule
	return r
}

func (r ApiReplaceAuthorizationServerPolicyRuleRequest) Execute() (*AuthorizationServerPolicyRule, *APIResponse, error) {
	return r.ApiService.ReplaceAuthorizationServerPolicyRuleExecute(r)
}

/*
ReplaceAuthorizationServerPolicyRule Replace a Policy Rule

Replaces the configuration of the Policy Rule defined in the specified Custom Authorization Server and Policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param policyId `id` of the Policy
 @param ruleId `id` of the Policy Rule
 @return ApiReplaceAuthorizationServerPolicyRuleRequest
*/
func (a *AuthorizationServerRulesAPIService) ReplaceAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) ApiReplaceAuthorizationServerPolicyRuleRequest {
	return ApiReplaceAuthorizationServerPolicyRuleRequest{
		ApiService: a,
		ctx: ctx,
		authServerId: authServerId,
		policyId: policyId,
		ruleId: ruleId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return AuthorizationServerPolicyRule
func (a *AuthorizationServerRulesAPIService) ReplaceAuthorizationServerPolicyRuleExecute(r ApiReplaceAuthorizationServerPolicyRuleRequest) (*AuthorizationServerPolicyRule, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationServerPolicyRule
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerRulesAPIService.ReplaceAuthorizationServerPolicyRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyId"+"}", url.PathEscape(parameterToString(r.policyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterToString(r.ruleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.policyRule == nil {
		return localVarReturnValue, nil, reportError("policyRule is required and must be specified")
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
	localVarPostBody = r.policyRule
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

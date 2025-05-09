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


type ApplicationPoliciesAPI interface {

	/*
	AssignApplicationPolicy Assign an application to a Policy

	Assigns an application to an [authentication policy](/openapi/okta-management/management/tag/Policy/), identified by `policyId`.
If the application was previously assigned to another policy, this operation replaces that assignment with the updated policy identified by `policyId`.

> **Note:** When you [merge duplicate authentication policies](https://help.okta.com/okta_help.htm?type=oie&id=ext-merge-auth-policies),
the policy and mapping CRUD operations may be unavailable during the consolidation. When the consolidation is complete, you receive an email.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId Application ID
	@param policyId `id` of the Policy
	@return ApiAssignApplicationPolicyRequest
	*/
	AssignApplicationPolicy(ctx context.Context, appId string, policyId string) ApiAssignApplicationPolicyRequest

	// AssignApplicationPolicyExecute executes the request
	AssignApplicationPolicyExecute(r ApiAssignApplicationPolicyRequest) (*APIResponse, error)
}

// ApplicationPoliciesAPIService ApplicationPoliciesAPI service
type ApplicationPoliciesAPIService service

type ApiAssignApplicationPolicyRequest struct {
	ctx context.Context
	ApiService ApplicationPoliciesAPI
	appId string
	policyId string
	retryCount int32
}

func (r ApiAssignApplicationPolicyRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignApplicationPolicyExecute(r)
}

/*
AssignApplicationPolicy Assign an application to a Policy

Assigns an application to an [authentication policy](/openapi/okta-management/management/tag/Policy/), identified by `policyId`.
If the application was previously assigned to another policy, this operation replaces that assignment with the updated policy identified by `policyId`.

> **Note:** When you [merge duplicate authentication policies](https://help.okta.com/okta_help.htm?type=oie&id=ext-merge-auth-policies),
the policy and mapping CRUD operations may be unavailable during the consolidation. When the consolidation is complete, you receive an email.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId Application ID
 @param policyId `id` of the Policy
 @return ApiAssignApplicationPolicyRequest
*/
func (a *ApplicationPoliciesAPIService) AssignApplicationPolicy(ctx context.Context, appId string, policyId string) ApiAssignApplicationPolicyRequest {
	return ApiAssignApplicationPolicyRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		policyId: policyId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *ApplicationPoliciesAPIService) AssignApplicationPolicyExecute(r ApiAssignApplicationPolicyRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationPoliciesAPIService.AssignApplicationPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appId}/policies/{policyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)
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

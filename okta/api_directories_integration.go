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

type DirectoriesIntegrationAPI interface {

	/*
			GetGroupAttributeQueryResult Retrieve the results of an AD group query

			Retrieves the results of the requested Active Directory (AD) group attributes using the `resultId` returned from the `POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query` call.
		If the operation has expired or if the `resultId` is invalid, returns a `404` status.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param appInstanceId ID of the AD instance in Okta
			@param groupId ID of the Okta group
			@param resultId The unique identifier returned by the initial POST request (`POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query`)
			@return ApiGetGroupAttributeQueryResultRequest
	*/
	GetGroupAttributeQueryResult(ctx context.Context, appInstanceId string, groupId string, resultId string) ApiGetGroupAttributeQueryResultRequest

	// GetGroupAttributeQueryResultExecute executes the request
	//  @return GroupProfileResult
	GetGroupAttributeQueryResultExecute(r ApiGetGroupAttributeQueryResultRequest) (*GroupProfileResult, *APIResponse, error)

	/*
			SubmitGroupAttributeQuery Submit a query for AD Group

			Submits a query search on the on-premises agent to asynchronously fetch specific Active Directory (AD) attributes for a group.
		Returns a `resultId` that is used to poll for the results.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param appInstanceId ID of the AD instance in Okta
			@param groupId ID of the Okta group
			@return ApiSubmitGroupAttributeQueryRequest
	*/
	SubmitGroupAttributeQuery(ctx context.Context, appInstanceId string, groupId string) ApiSubmitGroupAttributeQueryRequest

	// SubmitGroupAttributeQueryExecute executes the request
	//  @return GroupQueryResponse
	SubmitGroupAttributeQueryExecute(r ApiSubmitGroupAttributeQueryRequest) (*GroupQueryResponse, *APIResponse, error)

	/*
			UpdateGroupMembership Update an external directory group membership

			Updates an Active Directory or LDAP  group membership directly in the Active Directory or LDAP server.

		You can add or remove users from groups based on their identity and access requirements. This ensures that changes made to user access in Okta are reflected in AD or LDAP. When you use Okta Access Certifications to revoke a user's membership to an AD or LDAP group, the removal is reflected in AD or LDAP.

		See [AD Bidirectional Group Management](https://help.okta.com/okta_help.htm?type=oie&id=ad-bidirectional-group-mgmt) and [LDAP Bidirectional Group Management](https://help.okta.com/okta_help.htm?type=oie&id=ldap-bidirectional-group-mgmt).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param appInstanceId ID of the Active Directory or LDAP app instance in Okta
			@return ApiUpdateGroupMembershipRequest
	*/
	UpdateGroupMembership(ctx context.Context, appInstanceId string) ApiUpdateGroupMembershipRequest

	// UpdateGroupMembershipExecute executes the request
	UpdateGroupMembershipExecute(r ApiUpdateGroupMembershipRequest) (*APIResponse, error)
}

// DirectoriesIntegrationAPIService DirectoriesIntegrationAPI service
type DirectoriesIntegrationAPIService service

type ApiGetGroupAttributeQueryResultRequest struct {
	ctx           context.Context
	ApiService    DirectoriesIntegrationAPI
	appInstanceId string
	groupId       string
	resultId      string
	retryCount    int32
}

func (r ApiGetGroupAttributeQueryResultRequest) Execute() (*GroupProfileResult, *APIResponse, error) {
	return r.ApiService.GetGroupAttributeQueryResultExecute(r)
}

/*
GetGroupAttributeQueryResult Retrieve the results of an AD group query

Retrieves the results of the requested Active Directory (AD) group attributes using the `resultId` returned from the `POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query` call.
If the operation has expired or if the `resultId` is invalid, returns a `404` status.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appInstanceId ID of the AD instance in Okta
	@param groupId ID of the Okta group
	@param resultId The unique identifier returned by the initial POST request (`POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query`)
	@return ApiGetGroupAttributeQueryResultRequest
*/
func (a *DirectoriesIntegrationAPIService) GetGroupAttributeQueryResult(ctx context.Context, appInstanceId string, groupId string, resultId string) ApiGetGroupAttributeQueryResultRequest {
	return ApiGetGroupAttributeQueryResultRequest{
		ApiService:    a,
		ctx:           ctx,
		appInstanceId: appInstanceId,
		groupId:       groupId,
		resultId:      resultId,
		retryCount:    0,
	}
}

// Execute executes the request
//
//	@return GroupProfileResult
func (a *DirectoriesIntegrationAPIService) GetGroupAttributeQueryResultExecute(r ApiGetGroupAttributeQueryResultRequest) (*GroupProfileResult, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupProfileResult
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoriesIntegrationAPIService.GetGroupAttributeQueryResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/directories/{appInstanceId}/groups/{groupId}/query/{resultId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appInstanceId"+"}", url.PathEscape(parameterToString(r.appInstanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resultId"+"}", url.PathEscape(parameterToString(r.resultId, "")), -1)

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

type ApiSubmitGroupAttributeQueryRequest struct {
	ctx               context.Context
	ApiService        DirectoriesIntegrationAPI
	appInstanceId     string
	groupId           string
	groupQueryRequest *GroupQueryRequest
	retryCount        int32
}

func (r ApiSubmitGroupAttributeQueryRequest) GroupQueryRequest(groupQueryRequest GroupQueryRequest) ApiSubmitGroupAttributeQueryRequest {
	r.groupQueryRequest = &groupQueryRequest
	return r
}

func (r ApiSubmitGroupAttributeQueryRequest) Execute() (*GroupQueryResponse, *APIResponse, error) {
	return r.ApiService.SubmitGroupAttributeQueryExecute(r)
}

/*
SubmitGroupAttributeQuery Submit a query for AD Group

Submits a query search on the on-premises agent to asynchronously fetch specific Active Directory (AD) attributes for a group.
Returns a `resultId` that is used to poll for the results.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appInstanceId ID of the AD instance in Okta
	@param groupId ID of the Okta group
	@return ApiSubmitGroupAttributeQueryRequest
*/
func (a *DirectoriesIntegrationAPIService) SubmitGroupAttributeQuery(ctx context.Context, appInstanceId string, groupId string) ApiSubmitGroupAttributeQueryRequest {
	return ApiSubmitGroupAttributeQueryRequest{
		ApiService:    a,
		ctx:           ctx,
		appInstanceId: appInstanceId,
		groupId:       groupId,
		retryCount:    0,
	}
}

// Execute executes the request
//
//	@return GroupQueryResponse
func (a *DirectoriesIntegrationAPIService) SubmitGroupAttributeQueryExecute(r ApiSubmitGroupAttributeQueryRequest) (*GroupQueryResponse, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupQueryResponse
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoriesIntegrationAPIService.SubmitGroupAttributeQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/directories/{appInstanceId}/groups/{groupId}/query"
	localVarPath = strings.Replace(localVarPath, "{"+"appInstanceId"+"}", url.PathEscape(parameterToString(r.appInstanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupQueryRequest == nil {
		return localVarReturnValue, nil, reportError("groupQueryRequest is required and must be specified")
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
	localVarPostBody = r.groupQueryRequest
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
		if localVarHTTPResponse.StatusCode == 502 {
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
		if localVarHTTPResponse.StatusCode == 504 {
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

type ApiUpdateGroupMembershipRequest struct {
	ctx           context.Context
	ApiService    DirectoriesIntegrationAPI
	appInstanceId string
	agentAction   *AgentAction
	retryCount    int32
}

func (r ApiUpdateGroupMembershipRequest) AgentAction(agentAction AgentAction) ApiUpdateGroupMembershipRequest {
	r.agentAction = &agentAction
	return r
}

func (r ApiUpdateGroupMembershipRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UpdateGroupMembershipExecute(r)
}

/*
UpdateGroupMembership Update an external directory group membership

Updates an Active Directory or LDAP  group membership directly in the Active Directory or LDAP server.

You can add or remove users from groups based on their identity and access requirements. This ensures that changes made to user access in Okta are reflected in AD or LDAP. When you use Okta Access Certifications to revoke a user's membership to an AD or LDAP group, the removal is reflected in AD or LDAP.

See [AD Bidirectional Group Management](https://help.okta.com/okta_help.htm?type=oie&id=ad-bidirectional-group-mgmt) and [LDAP Bidirectional Group Management](https://help.okta.com/okta_help.htm?type=oie&id=ldap-bidirectional-group-mgmt).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appInstanceId ID of the Active Directory or LDAP app instance in Okta
	@return ApiUpdateGroupMembershipRequest
*/
func (a *DirectoriesIntegrationAPIService) UpdateGroupMembership(ctx context.Context, appInstanceId string) ApiUpdateGroupMembershipRequest {
	return ApiUpdateGroupMembershipRequest{
		ApiService:    a,
		ctx:           ctx,
		appInstanceId: appInstanceId,
		retryCount:    0,
	}
}

// Execute executes the request
func (a *DirectoriesIntegrationAPIService) UpdateGroupMembershipExecute(r ApiUpdateGroupMembershipRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoriesIntegrationAPIService.UpdateGroupMembership")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/directories/{appInstanceId}/groups/modify"
	localVarPath = strings.Replace(localVarPath, "{"+"appInstanceId"+"}", url.PathEscape(parameterToString(r.appInstanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentAction == nil {
		return nil, reportError("agentAction is required and must be specified")
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
	localVarPostBody = r.agentAction
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
		if localVarHTTPResponse.StatusCode == 502 {
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
		if localVarHTTPResponse.StatusCode == 504 {
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

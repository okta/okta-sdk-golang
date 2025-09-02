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


type UserLinkedObjectAPI interface {

	/*
	AssignLinkedObjectValueForPrimary Assign a linked object value for primary

	Assigns the first user as the `associated` and the second user as the `primary` for the specified relationship.

If the first user is already associated with a different `primary` for this relationship, the previous link is removed. A linked object relationship can specify only one primary user for an associated user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	@param primaryRelationshipName Name of the `primary` relationship being assigned
	@param primaryUserId User ID to be assigned to the `primary` relationship for the `associated` user
	@return ApiAssignLinkedObjectValueForPrimaryRequest
	*/
	AssignLinkedObjectValueForPrimary(ctx context.Context, userIdOrLogin string, primaryRelationshipName string, primaryUserId string) ApiAssignLinkedObjectValueForPrimaryRequest

	// AssignLinkedObjectValueForPrimaryExecute executes the request
	AssignLinkedObjectValueForPrimaryExecute(r ApiAssignLinkedObjectValueForPrimaryRequest) (*APIResponse, error)

	/*
	DeleteLinkedObjectForUser Delete a linked object value

	Deletes any existing relationship between the `associated` and `primary` user. For the `associated` user, this is specified by the ID. The `primary` name specifies the relationship.

The operation is successful if the relationship is deleted. The operation is also successful if the specified user isn't in the `associated` relationship for any instance of the specified `primary` and thus, no relationship is found.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	@param relationshipName Name of the `primary` or `associated` relationship being queried
	@return ApiDeleteLinkedObjectForUserRequest
	*/
	DeleteLinkedObjectForUser(ctx context.Context, userIdOrLogin string, relationshipName string) ApiDeleteLinkedObjectForUserRequest

	// DeleteLinkedObjectForUserExecute executes the request
	DeleteLinkedObjectForUserExecute(r ApiDeleteLinkedObjectForUserRequest) (*APIResponse, error)

	/*
	ListLinkedObjectsForUser List the primary or all of the associated linked object values

	Lists either the `self` link for the primary user or all associated users in the relationship specified by `relationshipName`. If the specified user isn't associated in any relationship, an empty array is returned.

Use `me` instead of `id` to specify the current session user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
	@param relationshipName Name of the `primary` or `associated` relationship being queried
	@return ApiListLinkedObjectsForUserRequest
	*/
	ListLinkedObjectsForUser(ctx context.Context, userIdOrLogin string, relationshipName string) ApiListLinkedObjectsForUserRequest

	// ListLinkedObjectsForUserExecute executes the request
	//  @return []ResponseLinks
	ListLinkedObjectsForUserExecute(r ApiListLinkedObjectsForUserRequest) ([]ResponseLinks, *APIResponse, error)
}

// UserLinkedObjectAPIService UserLinkedObjectAPI service
type UserLinkedObjectAPIService service

type ApiAssignLinkedObjectValueForPrimaryRequest struct {
	ctx context.Context
	ApiService UserLinkedObjectAPI
	userIdOrLogin string
	primaryRelationshipName string
	primaryUserId string
	retryCount int32
}

func (r ApiAssignLinkedObjectValueForPrimaryRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AssignLinkedObjectValueForPrimaryExecute(r)
}

/*
AssignLinkedObjectValueForPrimary Assign a linked object value for primary

Assigns the first user as the `associated` and the second user as the `primary` for the specified relationship.

If the first user is already associated with a different `primary` for this relationship, the previous link is removed. A linked object relationship can specify only one primary user for an associated user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
 @param primaryRelationshipName Name of the `primary` relationship being assigned
 @param primaryUserId User ID to be assigned to the `primary` relationship for the `associated` user
 @return ApiAssignLinkedObjectValueForPrimaryRequest
*/
func (a *UserLinkedObjectAPIService) AssignLinkedObjectValueForPrimary(ctx context.Context, userIdOrLogin string, primaryRelationshipName string, primaryUserId string) ApiAssignLinkedObjectValueForPrimaryRequest {
	return ApiAssignLinkedObjectValueForPrimaryRequest{
		ApiService: a,
		ctx: ctx,
		userIdOrLogin: userIdOrLogin,
		primaryRelationshipName: primaryRelationshipName,
		primaryUserId: primaryUserId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *UserLinkedObjectAPIService) AssignLinkedObjectValueForPrimaryExecute(r ApiAssignLinkedObjectValueForPrimaryRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserLinkedObjectAPIService.AssignLinkedObjectValueForPrimary")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userIdOrLogin}/linkedObjects/{primaryRelationshipName}/{primaryUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userIdOrLogin"+"}", url.PathEscape(parameterToString(r.userIdOrLogin, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"primaryRelationshipName"+"}", url.PathEscape(parameterToString(r.primaryRelationshipName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"primaryUserId"+"}", url.PathEscape(parameterToString(r.primaryUserId, "")), -1)

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

type ApiDeleteLinkedObjectForUserRequest struct {
	ctx context.Context
	ApiService UserLinkedObjectAPI
	userIdOrLogin string
	relationshipName string
	retryCount int32
}

func (r ApiDeleteLinkedObjectForUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteLinkedObjectForUserExecute(r)
}

/*
DeleteLinkedObjectForUser Delete a linked object value

Deletes any existing relationship between the `associated` and `primary` user. For the `associated` user, this is specified by the ID. The `primary` name specifies the relationship.

The operation is successful if the relationship is deleted. The operation is also successful if the specified user isn't in the `associated` relationship for any instance of the specified `primary` and thus, no relationship is found.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
 @param relationshipName Name of the `primary` or `associated` relationship being queried
 @return ApiDeleteLinkedObjectForUserRequest
*/
func (a *UserLinkedObjectAPIService) DeleteLinkedObjectForUser(ctx context.Context, userIdOrLogin string, relationshipName string) ApiDeleteLinkedObjectForUserRequest {
	return ApiDeleteLinkedObjectForUserRequest{
		ApiService: a,
		ctx: ctx,
		userIdOrLogin: userIdOrLogin,
		relationshipName: relationshipName,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *UserLinkedObjectAPIService) DeleteLinkedObjectForUserExecute(r ApiDeleteLinkedObjectForUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserLinkedObjectAPIService.DeleteLinkedObjectForUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userIdOrLogin"+"}", url.PathEscape(parameterToString(r.userIdOrLogin, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"relationshipName"+"}", url.PathEscape(parameterToString(r.relationshipName, "")), -1)

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

type ApiListLinkedObjectsForUserRequest struct {
	ctx context.Context
	ApiService UserLinkedObjectAPI
	userIdOrLogin string
	relationshipName string
	retryCount int32
}

func (r ApiListLinkedObjectsForUserRequest) Execute() ([]ResponseLinks, *APIResponse, error) {
	return r.ApiService.ListLinkedObjectsForUserExecute(r)
}

/*
ListLinkedObjectsForUser List the primary or all of the associated linked object values

Lists either the `self` link for the primary user or all associated users in the relationship specified by `relationshipName`. If the specified user isn't associated in any relationship, an empty array is returned.

Use `me` instead of `id` to specify the current session user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userIdOrLogin If for the `self` link, this is the ID of the user for whom you want to get the primary user ID. If for the `associated` relation, this is the user ID or login value of the user assigned the associated relationship.  This can be `me` to represent the current session user.
 @param relationshipName Name of the `primary` or `associated` relationship being queried
 @return ApiListLinkedObjectsForUserRequest
*/
func (a *UserLinkedObjectAPIService) ListLinkedObjectsForUser(ctx context.Context, userIdOrLogin string, relationshipName string) ApiListLinkedObjectsForUserRequest {
	return ApiListLinkedObjectsForUserRequest{
		ApiService: a,
		ctx: ctx,
		userIdOrLogin: userIdOrLogin,
		relationshipName: relationshipName,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []ResponseLinks
func (a *UserLinkedObjectAPIService) ListLinkedObjectsForUserExecute(r ApiListLinkedObjectsForUserRequest) ([]ResponseLinks, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ResponseLinks
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserLinkedObjectAPIService.ListLinkedObjectsForUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userIdOrLogin}/linkedObjects/{relationshipName}"
	localVarPath = strings.Replace(localVarPath, "{"+"userIdOrLogin"+"}", url.PathEscape(parameterToString(r.userIdOrLogin, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"relationshipName"+"}", url.PathEscape(parameterToString(r.relationshipName, "")), -1)

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

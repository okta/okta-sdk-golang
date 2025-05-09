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


type SubscriptionAPI interface {

	/*
	GetSubscriptionsNotificationTypeRole Retrieve a Subscription for a Role

	Retrieves a subscription by `notificationType` for a specified Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
	@param notificationType
	@return ApiGetSubscriptionsNotificationTypeRoleRequest
	*/
	GetSubscriptionsNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiGetSubscriptionsNotificationTypeRoleRequest

	// GetSubscriptionsNotificationTypeRoleExecute executes the request
	//  @return Subscription
	GetSubscriptionsNotificationTypeRoleExecute(r ApiGetSubscriptionsNotificationTypeRoleRequest) (*Subscription, *APIResponse, error)

	/*
	GetSubscriptionsNotificationTypeUser Retrieve a Subscription for a User

	Retrieves a subscription by `notificationType` for a specified User. Returns an `AccessDeniedException` message if requests are made for another user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationType
	@param userId ID of an existing Okta user
	@return ApiGetSubscriptionsNotificationTypeUserRequest
	*/
	GetSubscriptionsNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiGetSubscriptionsNotificationTypeUserRequest

	// GetSubscriptionsNotificationTypeUserExecute executes the request
	//  @return Subscription
	GetSubscriptionsNotificationTypeUserExecute(r ApiGetSubscriptionsNotificationTypeUserRequest) (*Subscription, *APIResponse, error)

	/*
	ListSubscriptionsRole List all Subscriptions for a Role

	Lists all subscriptions available to a specified Role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
	@return ApiListSubscriptionsRoleRequest
	*/
	ListSubscriptionsRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter) ApiListSubscriptionsRoleRequest

	// ListSubscriptionsRoleExecute executes the request
	//  @return []Subscription
	ListSubscriptionsRoleExecute(r ApiListSubscriptionsRoleRequest) ([]Subscription, *APIResponse, error)

	/*
	ListSubscriptionsUser List all Subscriptions for a User

	Lists all subscriptions available to a specified User. Returns an `AccessDeniedException` message if requests are made for another user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of an existing Okta user
	@return ApiListSubscriptionsUserRequest
	*/
	ListSubscriptionsUser(ctx context.Context, userId string) ApiListSubscriptionsUserRequest

	// ListSubscriptionsUserExecute executes the request
	//  @return []Subscription
	ListSubscriptionsUserExecute(r ApiListSubscriptionsUserRequest) ([]Subscription, *APIResponse, error)

	/*
	SubscribeByNotificationTypeRole Subscribe a Role to a Specific Notification Type

	Subscribes a Role to a specified notification type. Changes to Role subscriptions override the subscription status of any individual users with the Role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
	@param notificationType
	@return ApiSubscribeByNotificationTypeRoleRequest
	*/
	SubscribeByNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiSubscribeByNotificationTypeRoleRequest

	// SubscribeByNotificationTypeRoleExecute executes the request
	SubscribeByNotificationTypeRoleExecute(r ApiSubscribeByNotificationTypeRoleRequest) (*APIResponse, error)

	/*
	SubscribeByNotificationTypeUser Subscribe a User to a Specific Notification Type

	Subscribes the current user to a specified notification type. Returns an `AccessDeniedException` message if requests are made for another user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationType
	@param userId ID of an existing Okta user
	@return ApiSubscribeByNotificationTypeUserRequest
	*/
	SubscribeByNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiSubscribeByNotificationTypeUserRequest

	// SubscribeByNotificationTypeUserExecute executes the request
	SubscribeByNotificationTypeUserExecute(r ApiSubscribeByNotificationTypeUserRequest) (*APIResponse, error)

	/*
	UnsubscribeByNotificationTypeRole Unsubscribe a Role from a Specific Notification Type

	Unsubscribes a Role from a specified notification type. Changes to Role subscriptions override the subscription status of any individual users with the Role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
	@param notificationType
	@return ApiUnsubscribeByNotificationTypeRoleRequest
	*/
	UnsubscribeByNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiUnsubscribeByNotificationTypeRoleRequest

	// UnsubscribeByNotificationTypeRoleExecute executes the request
	UnsubscribeByNotificationTypeRoleExecute(r ApiUnsubscribeByNotificationTypeRoleRequest) (*APIResponse, error)

	/*
	UnsubscribeByNotificationTypeUser Unsubscribe a User from a Specific Notification Type

	Unsubscribes the current user from a specified notification type. Returns an `AccessDeniedException` message if requests are made for another user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationType
	@param userId ID of an existing Okta user
	@return ApiUnsubscribeByNotificationTypeUserRequest
	*/
	UnsubscribeByNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiUnsubscribeByNotificationTypeUserRequest

	// UnsubscribeByNotificationTypeUserExecute executes the request
	UnsubscribeByNotificationTypeUserExecute(r ApiUnsubscribeByNotificationTypeUserRequest) (*APIResponse, error)
}

// SubscriptionAPIService SubscriptionAPI service
type SubscriptionAPIService service

type ApiGetSubscriptionsNotificationTypeRoleRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	roleRef ListSubscriptionsRoleRoleRefParameter
	notificationType string
	retryCount int32
}

func (r ApiGetSubscriptionsNotificationTypeRoleRequest) Execute() (*Subscription, *APIResponse, error) {
	return r.ApiService.GetSubscriptionsNotificationTypeRoleExecute(r)
}

/*
GetSubscriptionsNotificationTypeRole Retrieve a Subscription for a Role

Retrieves a subscription by `notificationType` for a specified Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
 @param notificationType
 @return ApiGetSubscriptionsNotificationTypeRoleRequest
*/
func (a *SubscriptionAPIService) GetSubscriptionsNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiGetSubscriptionsNotificationTypeRoleRequest {
	return ApiGetSubscriptionsNotificationTypeRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleRef: roleRef,
		notificationType: notificationType,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return Subscription
func (a *SubscriptionAPIService) GetSubscriptionsNotificationTypeRoleExecute(r ApiGetSubscriptionsNotificationTypeRoleRequest) (*Subscription, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subscription
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.GetSubscriptionsNotificationTypeRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/roles/{roleRef}/subscriptions/{notificationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleRef"+"}", url.PathEscape(parameterToString(r.roleRef, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)

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

type ApiGetSubscriptionsNotificationTypeUserRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	notificationType string
	userId string
	retryCount int32
}

func (r ApiGetSubscriptionsNotificationTypeUserRequest) Execute() (*Subscription, *APIResponse, error) {
	return r.ApiService.GetSubscriptionsNotificationTypeUserExecute(r)
}

/*
GetSubscriptionsNotificationTypeUser Retrieve a Subscription for a User

Retrieves a subscription by `notificationType` for a specified User. Returns an `AccessDeniedException` message if requests are made for another user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param notificationType
 @param userId ID of an existing Okta user
 @return ApiGetSubscriptionsNotificationTypeUserRequest
*/
func (a *SubscriptionAPIService) GetSubscriptionsNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiGetSubscriptionsNotificationTypeUserRequest {
	return ApiGetSubscriptionsNotificationTypeUserRequest{
		ApiService: a,
		ctx: ctx,
		notificationType: notificationType,
		userId: userId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return Subscription
func (a *SubscriptionAPIService) GetSubscriptionsNotificationTypeUserExecute(r ApiGetSubscriptionsNotificationTypeUserRequest) (*Subscription, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subscription
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.GetSubscriptionsNotificationTypeUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/subscriptions/{notificationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiListSubscriptionsRoleRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	roleRef ListSubscriptionsRoleRoleRefParameter
	retryCount int32
}

func (r ApiListSubscriptionsRoleRequest) Execute() ([]Subscription, *APIResponse, error) {
	return r.ApiService.ListSubscriptionsRoleExecute(r)
}

/*
ListSubscriptionsRole List all Subscriptions for a Role

Lists all subscriptions available to a specified Role

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
 @return ApiListSubscriptionsRoleRequest
*/
func (a *SubscriptionAPIService) ListSubscriptionsRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter) ApiListSubscriptionsRoleRequest {
	return ApiListSubscriptionsRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleRef: roleRef,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionAPIService) ListSubscriptionsRoleExecute(r ApiListSubscriptionsRoleRequest) ([]Subscription, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Subscription
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.ListSubscriptionsRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/roles/{roleRef}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"roleRef"+"}", url.PathEscape(parameterToString(r.roleRef, "")), -1)

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

type ApiListSubscriptionsUserRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	userId string
	retryCount int32
}

func (r ApiListSubscriptionsUserRequest) Execute() ([]Subscription, *APIResponse, error) {
	return r.ApiService.ListSubscriptionsUserExecute(r)
}

/*
ListSubscriptionsUser List all Subscriptions for a User

Lists all subscriptions available to a specified User. Returns an `AccessDeniedException` message if requests are made for another user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of an existing Okta user
 @return ApiListSubscriptionsUserRequest
*/
func (a *SubscriptionAPIService) ListSubscriptionsUser(ctx context.Context, userId string) ApiListSubscriptionsUserRequest {
	return ApiListSubscriptionsUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *SubscriptionAPIService) ListSubscriptionsUserExecute(r ApiListSubscriptionsUserRequest) ([]Subscription, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Subscription
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err 				 error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.ListSubscriptionsUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiSubscribeByNotificationTypeRoleRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	roleRef ListSubscriptionsRoleRoleRefParameter
	notificationType string
	retryCount int32
}

func (r ApiSubscribeByNotificationTypeRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.SubscribeByNotificationTypeRoleExecute(r)
}

/*
SubscribeByNotificationTypeRole Subscribe a Role to a Specific Notification Type

Subscribes a Role to a specified notification type. Changes to Role subscriptions override the subscription status of any individual users with the Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
 @param notificationType
 @return ApiSubscribeByNotificationTypeRoleRequest
*/
func (a *SubscriptionAPIService) SubscribeByNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiSubscribeByNotificationTypeRoleRequest {
	return ApiSubscribeByNotificationTypeRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleRef: roleRef,
		notificationType: notificationType,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *SubscriptionAPIService) SubscribeByNotificationTypeRoleExecute(r ApiSubscribeByNotificationTypeRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.SubscribeByNotificationTypeRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/roles/{roleRef}/subscriptions/{notificationType}/subscribe"
	localVarPath = strings.Replace(localVarPath, "{"+"roleRef"+"}", url.PathEscape(parameterToString(r.roleRef, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)

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

type ApiSubscribeByNotificationTypeUserRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	notificationType string
	userId string
	retryCount int32
}

func (r ApiSubscribeByNotificationTypeUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.SubscribeByNotificationTypeUserExecute(r)
}

/*
SubscribeByNotificationTypeUser Subscribe a User to a Specific Notification Type

Subscribes the current user to a specified notification type. Returns an `AccessDeniedException` message if requests are made for another user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param notificationType
 @param userId ID of an existing Okta user
 @return ApiSubscribeByNotificationTypeUserRequest
*/
func (a *SubscriptionAPIService) SubscribeByNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiSubscribeByNotificationTypeUserRequest {
	return ApiSubscribeByNotificationTypeUserRequest{
		ApiService: a,
		ctx: ctx,
		notificationType: notificationType,
		userId: userId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *SubscriptionAPIService) SubscribeByNotificationTypeUserExecute(r ApiSubscribeByNotificationTypeUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.SubscribeByNotificationTypeUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/subscriptions/{notificationType}/subscribe"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiUnsubscribeByNotificationTypeRoleRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	roleRef ListSubscriptionsRoleRoleRefParameter
	notificationType string
	retryCount int32
}

func (r ApiUnsubscribeByNotificationTypeRoleRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnsubscribeByNotificationTypeRoleExecute(r)
}

/*
UnsubscribeByNotificationTypeRole Unsubscribe a Role from a Specific Notification Type

Unsubscribes a Role from a specified notification type. Changes to Role subscriptions override the subscription status of any individual users with the Role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roleRef A reference to an existing role. Standard roles require a `roleType`, while Custom Roles require a `roleId`. See [Standard Roles](/openapi/okta-management/guides/roles/#standard-roles).
 @param notificationType
 @return ApiUnsubscribeByNotificationTypeRoleRequest
*/
func (a *SubscriptionAPIService) UnsubscribeByNotificationTypeRole(ctx context.Context, roleRef ListSubscriptionsRoleRoleRefParameter, notificationType string) ApiUnsubscribeByNotificationTypeRoleRequest {
	return ApiUnsubscribeByNotificationTypeRoleRequest{
		ApiService: a,
		ctx: ctx,
		roleRef: roleRef,
		notificationType: notificationType,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *SubscriptionAPIService) UnsubscribeByNotificationTypeRoleExecute(r ApiUnsubscribeByNotificationTypeRoleRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.UnsubscribeByNotificationTypeRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/roles/{roleRef}/subscriptions/{notificationType}/unsubscribe"
	localVarPath = strings.Replace(localVarPath, "{"+"roleRef"+"}", url.PathEscape(parameterToString(r.roleRef, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)

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

type ApiUnsubscribeByNotificationTypeUserRequest struct {
	ctx context.Context
	ApiService SubscriptionAPI
	notificationType string
	userId string
	retryCount int32
}

func (r ApiUnsubscribeByNotificationTypeUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UnsubscribeByNotificationTypeUserExecute(r)
}

/*
UnsubscribeByNotificationTypeUser Unsubscribe a User from a Specific Notification Type

Unsubscribes the current user from a specified notification type. Returns an `AccessDeniedException` message if requests are made for another user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param notificationType
 @param userId ID of an existing Okta user
 @return ApiUnsubscribeByNotificationTypeUserRequest
*/
func (a *SubscriptionAPIService) UnsubscribeByNotificationTypeUser(ctx context.Context, notificationType string, userId string) ApiUnsubscribeByNotificationTypeUserRequest {
	return ApiUnsubscribeByNotificationTypeUserRequest{
		ApiService: a,
		ctx: ctx,
		notificationType: notificationType,
		userId: userId,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *SubscriptionAPIService) UnsubscribeByNotificationTypeUserExecute(r ApiUnsubscribeByNotificationTypeUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAPIService.UnsubscribeByNotificationTypeUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationType"+"}", url.PathEscape(parameterToString(r.notificationType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

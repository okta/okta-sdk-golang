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

type UserAPI interface {

	/*
			CreateUser Create a user

			Creates a new user in your Okta org with or without credentials.<br>
		> **Legal Disclaimer**
		>
		> After a user is added to the Okta directory, they receive an activation email. As part of signing up for this service,
		> you agreed not to use Okta's service/product to spam and/or send unsolicited messages.
		> Please refrain from adding unrelated accounts to the directory as Okta is not responsible for, and disclaims any and all
		> liability associated with, the activation email's content. You, and you alone, bear responsibility for the emails sent to any recipients.

		All responses return the created user. Activation of a user is an asynchronous operation. The system performs group reconciliation during activation and assigns the user to all apps via direct or indirect relationships (group memberships).
		* The user's `transitioningToStatus` property is `ACTIVE` during activation to indicate that the user hasn't completed the asynchronous operation.
		* The user's `status` is `ACTIVE` when the activation process is complete.

		The user is emailed a one-time activation token if activated without a password.

		> **Note:** If the user is assigned to an app that is configured for provisioning, the activation process triggers downstream provisioning to the app.  It is possible for a user to sign in before these apps have been successfully provisioned for the user.

		> **Important:** Do not generate or send a one-time activation token when activating users with an assigned password. Users should sign in with their assigned password.

		For more information about the various scenarios of creating a user listed in the examples, see the [User creation scenarios](/openapi/okta-management/management/tag/User/#user-creation-scenarios) section.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCreateUserRequest
	*/
	CreateUser(ctx context.Context) ApiCreateUserRequest

	// CreateUserExecute executes the request
	//  @return User
	CreateUserExecute(r ApiCreateUserRequest) (*User, *APIResponse, error)

	/*
			DeleteUser Delete a user

			Deletes a user permanently. This operation can only be performed on users that have a `DEPROVISIONED` status.

		> **Warning:** This action can't be recovered!

		This operation on a user that hasn't been deactivated causes that user to be deactivated. A second delete operation is required to delete the user.

		> **Note:** You can also perform user deletion asynchronously. To invoke asynchronous user deletion, pass an HTTP header `Prefer: respond-async` with the request.

		This header is also supported by user deactivation, which is performed if the delete endpoint is invoked on a user that hasn't been deactivated.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
			@return ApiDeleteUserRequest
	*/
	DeleteUser(ctx context.Context, id string) ApiDeleteUserRequest

	// DeleteUserExecute executes the request
	DeleteUserExecute(r ApiDeleteUserRequest) (*APIResponse, error)

	/*
			GetUser Retrieve a user

			Retrieves a user from your Okta org.

		You can substitute `me` for the `id` to fetch the current user linked to an API token or session cookie.
		 * The request returns the user linked to the API token that is specified in the Authorization header, not the user linked to the active session. Details of the admin user who granted the API token is returned.
		 * When the end user has an active Okta session, it is typically a CORS request from the browser. Therefore, it's possible to retrieve the current user without the Authorization header.

		When fetching a user by `login` or `login shortname`, [URL encode](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) the request parameter to ensure that special characters are escaped properly. Logins with a `/` character can only be fetched by `id` due to URL issues with escaping the `/` character. If you don't know a user's ID, you can use the [List all users](/openapi/okta-management/management/tag/User/#tag/User/operation/listUsers) endpoint to find it.

		> **Note:** Some browsers block third-party cookies by default, which disrupts Okta functionality in certain flows. See [Mitigate the impact of third-party cookie deprecation](https://help.okta.com/okta_help.htm?type=oie&id=ext-third-party-cookies).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
			@return ApiGetUserRequest
	*/
	GetUser(ctx context.Context, id string) ApiGetUserRequest

	// GetUserExecute executes the request
	//  @return UserGetSingleton
	GetUserExecute(r ApiGetUserRequest) (*UserGetSingleton, *APIResponse, error)

	/*
		ListUserBlocks List all user blocks

		Lists information about how the user is blocked from accessing their account

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
		@return ApiListUserBlocksRequest
	*/
	ListUserBlocks(ctx context.Context, id string) ApiListUserBlocksRequest

	// ListUserBlocksExecute executes the request
	//  @return []UserBlock
	ListUserBlocksExecute(r ApiListUserBlocksRequest) ([]UserBlock, *APIResponse, error)

	/*
			ListUsers List all users

			Lists users in your org, with pagination in most cases.

		A subset of users can be returned that match a supported filter expression or search criteria. Different results are returned depending on specified queries in the request.

		> **Notes:**
		>  * This operation omits users that have a status of `DEPROVISIONED` in the response. To return all users, use a filter or search query instead.
		>  * The `search` parameter results are sourced from an eventually consistent datasource and may not reflect the latest information.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListUsersRequest
	*/
	ListUsers(ctx context.Context) ApiListUsersRequest

	// ListUsersExecute executes the request
	//  @return []User
	ListUsersExecute(r ApiListUsersRequest) ([]User, *APIResponse, error)

	/*
			ReplaceUser Replace a user

			Replaces a user's profile, credentials, or both using strict-update semantics.

		All profile properties must be specified when updating a user's profile with a `PUT` method. Any property not specified in the request is deleted.
		> **Important:** Don't use a `PUT` method for partial updates.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
			@return ApiReplaceUserRequest
	*/
	ReplaceUser(ctx context.Context, id string) ApiReplaceUserRequest

	// ReplaceUserExecute executes the request
	//  @return User
	ReplaceUserExecute(r ApiReplaceUserRequest) (*User, *APIResponse, error)

	/*
			UpdateUser Update a user

			Updates a user's profile or credentials with partial update semantics.

		> **Important:** Use the `POST` method for partial updates. Unspecified properties are set to null with `PUT`.

		`profile` and `credentials` can be updated independently or together with a single request.
		> **Note**: Currently, the user type of a user can only be changed via a full replacement PUT operation. If the request parameters of a partial update include the type element from the user object,
		the value must match the existing type of the user. Only admins are permitted to change the user type of a user; end users are not allowed to change their own user type.

		> **Note**: To update a current user's profile with partial semantics, the `/api/v1/users/me` endpoint can be invoked.
		>
		> A user can only update profile properties for which the user has write access. Within the profile, if the user tries to update the primary or the secondary email IDs, verification emails are sent to those email IDs, and the fields are updated only upon verification.

		If you are using this endpoint to set a password, it sets a password without validating existing user credentials. This is an administrative operation. For operations that validate credentials, refer to the [Reset password](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/resetPassword), [Start forgot password flow](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/forgotPassword), and [Update password](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/changePassword) endpoints.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
			@return ApiUpdateUserRequest
	*/
	UpdateUser(ctx context.Context, id string) ApiUpdateUserRequest

	// UpdateUserExecute executes the request
	//  @return User
	UpdateUserExecute(r ApiUpdateUserRequest) (*User, *APIResponse, error)
}

// UserAPIService UserAPI service
type UserAPIService service

type ApiCreateUserRequest struct {
	ctx        context.Context
	ApiService UserAPI
	body       *CreateUserRequest
	activate   *bool
	provider   *bool
	nextLogin  *string
	retryCount int32
}

func (r ApiCreateUserRequest) Body(body CreateUserRequest) ApiCreateUserRequest {
	r.body = &body
	return r
}

// Executes an [activation lifecycle](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserLifecycle/#tag/UserLifecycle/operation/activateUser) operation when creating the user
func (r ApiCreateUserRequest) Activate(activate bool) ApiCreateUserRequest {
	r.activate = &activate
	return r
}

// Indicates whether to create a user with a specified authentication provider
func (r ApiCreateUserRequest) Provider(provider bool) ApiCreateUserRequest {
	r.provider = &provider
	return r
}

// With &#x60;activate&#x3D;true&#x60;, if &#x60;nextLogin&#x3D;changePassword&#x60;, a user is created, activated, and the password is set to &#x60;EXPIRED&#x60;. The user must change it the next time they sign in.
func (r ApiCreateUserRequest) NextLogin(nextLogin string) ApiCreateUserRequest {
	r.nextLogin = &nextLogin
	return r
}

func (r ApiCreateUserRequest) Execute() (*User, *APIResponse, error) {
	return r.ApiService.CreateUserExecute(r)
}

/*
CreateUser Create a user

Creates a new user in your Okta org with or without credentials.<br>
> **Legal Disclaimer**
>
> After a user is added to the Okta directory, they receive an activation email. As part of signing up for this service,
> you agreed not to use Okta's service/product to spam and/or send unsolicited messages.
> Please refrain from adding unrelated accounts to the directory as Okta is not responsible for, and disclaims any and all
> liability associated with, the activation email's content. You, and you alone, bear responsibility for the emails sent to any recipients.

All responses return the created user. Activation of a user is an asynchronous operation. The system performs group reconciliation during activation and assigns the user to all apps via direct or indirect relationships (group memberships).
* The user's `transitioningToStatus` property is `ACTIVE` during activation to indicate that the user hasn't completed the asynchronous operation.
* The user's `status` is `ACTIVE` when the activation process is complete.

The user is emailed a one-time activation token if activated without a password.

> **Note:** If the user is assigned to an app that is configured for provisioning, the activation process triggers downstream provisioning to the app.  It is possible for a user to sign in before these apps have been successfully provisioned for the user.

> **Important:** Do not generate or send a one-time activation token when activating users with an assigned password. Users should sign in with their assigned password.

For more information about the various scenarios of creating a user listed in the examples, see the [User creation scenarios](/openapi/okta-management/management/tag/User/#user-creation-scenarios) section.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateUserRequest
*/
func (a *UserAPIService) CreateUser(ctx context.Context) ApiCreateUserRequest {
	return ApiCreateUserRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return User
func (a *UserAPIService) CreateUserExecute(r ApiCreateUserRequest) (*User, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.CreateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.activate != nil {
		localVarQueryParams.Add("activate", parameterToString(*r.activate, ""))
	}
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	if r.nextLogin != nil {
		localVarQueryParams.Add("nextLogin", parameterToString(*r.nextLogin, ""))
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
	localVarPostBody = r.body
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

type ApiDeleteUserRequest struct {
	ctx        context.Context
	ApiService UserAPI
	id         string
	sendEmail  *bool
	prefer     *string
	retryCount int32
}

// Sends a deactivation email to the admin if &#x60;true&#x60;
func (r ApiDeleteUserRequest) SendEmail(sendEmail bool) ApiDeleteUserRequest {
	r.sendEmail = &sendEmail
	return r
}

func (r ApiDeleteUserRequest) Prefer(prefer string) ApiDeleteUserRequest {
	r.prefer = &prefer
	return r
}

func (r ApiDeleteUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteUserExecute(r)
}

/*
DeleteUser Delete a user

Deletes a user permanently. This operation can only be performed on users that have a `DEPROVISIONED` status.

> **Warning:** This action can't be recovered!

This operation on a user that hasn't been deactivated causes that user to be deactivated. A second delete operation is required to delete the user.

> **Note:** You can also perform user deletion asynchronously. To invoke asynchronous user deletion, pass an HTTP header `Prefer: respond-async` with the request.

This header is also supported by user deactivation, which is performed if the delete endpoint is invoked on a user that hasn't been deactivated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	@return ApiDeleteUserRequest
*/
func (a *UserAPIService) DeleteUser(ctx context.Context, id string) ApiDeleteUserRequest {
	return ApiDeleteUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		retryCount: 0,
	}
}

// Execute executes the request
func (a *UserAPIService) DeleteUserExecute(r ApiDeleteUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.DeleteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sendEmail != nil {
		localVarQueryParams.Add("sendEmail", parameterToString(*r.sendEmail, ""))
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
	if r.prefer != nil {
		localVarHeaderParams["Prefer"] = parameterToString(*r.prefer, "")
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

type ApiGetUserRequest struct {
	ctx         context.Context
	ApiService  UserAPI
	id          string
	contentType *string
	expand      *string
	retryCount  int32
}

// Specifies the media type of the resource. Optional &#x60;okta-response&#x60; value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for &#x60;okta-response&#x60;:   * &#x60;omitCredentials&#x60;: Omits the credentials subobject from the response.   * &#x60;omitCredentialsLinks&#x60;: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * &#x60;omitTransitioningToStatus&#x60;: Omits the &#x60;transitioningToStatus&#x60; field from the response.
func (r ApiGetUserRequest) ContentType(contentType string) ApiGetUserRequest {
	r.contentType = &contentType
	return r
}

// An optional parameter to include metadata in the &#x60;_embedded&#x60; attribute. Valid values: &#x60;blocks&#x60; or &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;classification&#x60;.
func (r ApiGetUserRequest) Expand(expand string) ApiGetUserRequest {
	r.expand = &expand
	return r
}

func (r ApiGetUserRequest) Execute() (*UserGetSingleton, *APIResponse, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Retrieve a user

Retrieves a user from your Okta org.

You can substitute `me` for the `id` to fetch the current user linked to an API token or session cookie.
  - The request returns the user linked to the API token that is specified in the Authorization header, not the user linked to the active session. Details of the admin user who granted the API token is returned.
  - When the end user has an active Okta session, it is typically a CORS request from the browser. Therefore, it's possible to retrieve the current user without the Authorization header.

When fetching a user by `login` or `login shortname`, [URL encode](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) the request parameter to ensure that special characters are escaped properly. Logins with a `/` character can only be fetched by `id` due to URL issues with escaping the `/` character. If you don't know a user's ID, you can use the [List all users](/openapi/okta-management/management/tag/User/#tag/User/operation/listUsers) endpoint to find it.

> **Note:** Some browsers block third-party cookies by default, which disrupts Okta functionality in certain flows. See [Mitigate the impact of third-party cookie deprecation](https://help.okta.com/okta_help.htm?type=oie&id=ext-third-party-cookies).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	@return ApiGetUserRequest
*/
func (a *UserAPIService) GetUser(ctx context.Context, id string) ApiGetUserRequest {
	return ApiGetUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return UserGetSingleton
func (a *UserAPIService) GetUserExecute(r ApiGetUserRequest) (*UserGetSingleton, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserGetSingleton
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiListUserBlocksRequest struct {
	ctx        context.Context
	ApiService UserAPI
	id         string
	retryCount int32
}

func (r ApiListUserBlocksRequest) Execute() ([]UserBlock, *APIResponse, error) {
	return r.ApiService.ListUserBlocksExecute(r)
}

/*
ListUserBlocks List all user blocks

Lists information about how the user is blocked from accessing their account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	@return ApiListUserBlocksRequest
*/
func (a *UserAPIService) ListUserBlocks(ctx context.Context, id string) ApiListUserBlocksRequest {
	return ApiListUserBlocksRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []UserBlock
func (a *UserAPIService) ListUserBlocksExecute(r ApiListUserBlocksRequest) ([]UserBlock, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UserBlock
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.ListUserBlocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{id}/blocks"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiListUsersRequest struct {
	ctx         context.Context
	ApiService  UserAPI
	contentType *string
	search      *string
	filter      *string
	q           *string
	after       *string
	limit       *int32
	sortBy      *string
	sortOrder   *string
	fields      *string
	expand      *string
	retryCount  int32
}

// Specifies the media type of the resource. Optional &#x60;okta-response&#x60; value can be included for performance optimization.  Complex DelAuth configurations may degrade performance when fetching specific parts of the response, and passing this parameter can omit these parts, bypassing the bottleneck.  Enum values for &#x60;okta-response&#x60;:   * &#x60;omitCredentials&#x60;: Omits the credentials subobject from the response.   * &#x60;omitCredentialsLinks&#x60;: Omits the following HAL links from the response: Update password, Change recovery question, Start forgot password flow, Reset password, Reset factors, Unlock.   * &#x60;omitTransitioningToStatus&#x60;: Omits the &#x60;transitioningToStatus&#x60; field from the response.
func (r ApiListUsersRequest) ContentType(contentType string) ApiListUsersRequest {
	r.contentType = &contentType
	return r
}

// Searches for users with a supported filtering expression for most properties. Okta recommends using this parameter for optimal search performance.  &gt; **Note:** Using an overly complex or long search query can result in an error.  This operation supports [pagination](https://developer.okta.com/docs/api/#pagination). Use an ID lookup for records that you update to ensure your results contain the latest data. Returned users include those with the &#x60;DEPROVISIONED&#x60; status.  Property names in the search parameter are case sensitive, whereas operators (&#x60;eq&#x60;, &#x60;sw&#x60;, and so on) and string values are case insensitive. Unlike with user logins, diacritical marks are significant in search string values: a search for &#x60;isaac.brock&#x60; finds &#x60;Isaac.Brock&#x60;, but doesn&#39;t find a property whose value is &#x60;isáàc.bröck&#x60;.  This operation requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, &#x60;search&#x3D;profile.department eq \&quot;Engineering\&quot;&#x60; is encoded as &#x60;search&#x3D;profile.department%20eq%20%22Engineering%22&#x60;. If you use the special character &#x60;\&quot;&#x60; within a quoted string, it must also be escaped &#x60;\\&#x60; and encoded. For example, &#x60;search&#x3D;profile.lastName eq \&quot;bob\&quot;smith\&quot;&#x60; is encoded as &#x60;search&#x3D;profile.lastName%20eq%20%22bob%5C%22smith%22&#x60;. See [Special Characters](https://developer.okta.com/docs/api/#special-characters).  This operation searches many properties:   * Any user profile attribute, including custom-defined attributes   * The top-level properties: &#x60;id&#x60;, &#x60;status&#x60;, &#x60;created&#x60;, &#x60;activated&#x60;, &#x60;statusChanged&#x60;, and &#x60;lastUpdated&#x60;   * The [user type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/#tag/UserType/operation/updateUserType) accessed as &#x60;type.id&#x60;  &gt; **Note:** &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; The ability to search by user classification is available as an [Early Access](https://developer.okta.com/docs/api/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature. The &#x60;classification.type&#x60; property cannot be used in conjunction with other search terms. You can search using &#x60;classification.type eq \&quot;LITE\&quot;&#x60; or &#x60;classification.type eq \&quot;STANDARD\&quot;&#x60;.  You can also use &#x60;sortBy&#x60; and &#x60;sortOrder&#x60; parameters. The &#x60;ne&#x60; (not equal) operator isn&#39;t supported, but you can obtain the same result by using &#x60;lt ... or ... gt&#x60;. For example, to see all users except those that have a status of &#x60;STAGED&#x60;, use &#x60;(status lt \&quot;STAGED\&quot; or status gt \&quot;STAGED\&quot;)&#x60;.  You can search properties that are arrays. If any element matches the search term, the entire array (object) is returned. Okta follows the [SCIM Protocol Specification](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) for searching arrays. You can search multiple arrays, multiple values in an array, as well as using the standard logical and filtering operators. See [Filter](https://developer.okta.com/docs/reference/core-okta-api/#filter).  Searches for users can be filtered by the following operators: &#x60;sw&#x60;, &#x60;eq&#x60;, and &#x60;co&#x60;. You can only use &#x60;co&#x60; with these select user profile attributes: &#x60;profile.firstName&#x60;, &#x60;profile.lastName&#x60;, &#x60;profile.email&#x60;, and &#x60;profile.login&#x60;. See [Operators](https://developer.okta.com/docs/api/#operators).
func (r ApiListUsersRequest) Search(search string) ApiListUsersRequest {
	r.search = &search
	return r
}

// Filters users with a supported expression for a subset of properties.  &gt; **Note:** Returned users include those with the &#x60;DEPROVISIONED&#x60; status.  This requires [URL encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). For example, &#x60;filter&#x3D;lastUpdated gt \&quot;2013-06-01T00:00:00.000Z\&quot;&#x60; is encoded as &#x60;filter&#x3D;lastUpdated%20gt%20%222013-06-01T00:00:00.000Z%22&#x60;. Filtering is case-sensitive for property names and query values, while operators are case-insensitive.  Filtering supports the following limited number of properties: &#x60;status&#x60;, &#x60;lastUpdated&#x60;, &#x60;id&#x60;, &#x60;profile.login&#x60;, &#x60;profile.email&#x60;, &#x60;profile.firstName&#x60;, and &#x60;profile.lastName&#x60;.  Additionally, filtering supports only the equal &#x60;eq&#x60; operator from the standard Okta API filtering semantics, except in the case of the &#x60;lastUpdated&#x60; property. This property can also use the inequality operators (&#x60;gt&#x60;, &#x60;ge&#x60;, &#x60;lt&#x60;, and &#x60;le&#x60;). For logical operators, only the logical operators &#x60;and&#x60; and &#x60;or&#x60; are supported. The &#x60;not&#x60; operator isn&#39;t supported. See [Filter](https://developer.okta.com/docs/api/#filter) and [Operators](https://developer.okta.com/docs/api/#operators).
func (r ApiListUsersRequest) Filter(filter string) ApiListUsersRequest {
	r.filter = &filter
	return r
}

// Finds users who match the specified query. This doesn&#39;t support pagination.  &gt; **Note:** For optimal performance, use the &#x60;search&#x60; parameter instead.  Use the &#x60;q&#x60; parameter for simple queries, such as a lookup of users by name when creating a people picker.  The value of &#x60;q&#x60; is matched against &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60;. This performs a &#x60;startsWith&#x60; match, but this is an implementation detail and can change without notice. You don&#39;t need to specify &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60;.  &gt; **Note:** Using the &#x60;q&#x60; parameter in a request omits users that have a status of &#x60;DEPROVISIONED&#x60;. To return all users, use a filter or search query instead.
func (r ApiListUsersRequest) Q(q string) ApiListUsersRequest {
	r.q = &q
	return r
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListUsersRequest) After(after string) ApiListUsersRequest {
	r.after = &after
	return r
}

// Specifies the number of results returned. Defaults to 10 if &#x60;q&#x60; is provided.
func (r ApiListUsersRequest) Limit(limit int32) ApiListUsersRequest {
	r.limit = &limit
	return r
}

// Specifies field to sort by (for search queries only). This can be any single property, for example &#x60;sortBy&#x3D;profile.lastName&#x60;. Users with the same value for the &#x60;sortBy&#x60; property will be ordered by &#x60;id&#x60;.
func (r ApiListUsersRequest) SortBy(sortBy string) ApiListUsersRequest {
	r.sortBy = &sortBy
	return r
}

// Specifies the sort order: &#x60;asc&#x60; or &#x60;desc&#x60; (for search queries only). Sorting is done in ASCII sort order (that is, by ASCII character value), but isn&#39;t case sensitive. &#x60;sortOrder&#x60; is ignored if &#x60;sortBy&#x60; isn&#39;t present.
func (r ApiListUsersRequest) SortOrder(sortOrder string) ApiListUsersRequest {
	r.sortOrder = &sortOrder
	return r
}

// Specifies a select set of user properties to query. Any other properties will be filtered out of the returned users.  Requested fields should be comma-separated. Profile attributes should be contained within a &#x60;profile:({field1}, {field2}, ...)&#x60; directive.
func (r ApiListUsersRequest) Fields(fields string) ApiListUsersRequest {
	r.fields = &fields
	return r
}

// &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;A parameter to include metadata in the &#x60;_embedded&#x60; property. Supported value: &#x60;classification&#x60;.
func (r ApiListUsersRequest) Expand(expand string) ApiListUsersRequest {
	r.expand = &expand
	return r
}

func (r ApiListUsersRequest) Execute() ([]User, *APIResponse, error) {
	return r.ApiService.ListUsersExecute(r)
}

/*
ListUsers List all users

Lists users in your org, with pagination in most cases.

A subset of users can be returned that match a supported filter expression or search criteria. Different results are returned depending on specified queries in the request.

> **Notes:**
>  * This operation omits users that have a status of `DEPROVISIONED` in the response. To return all users, use a filter or search query instead.
>  * The `search` parameter results are sourced from an eventually consistent datasource and may not reflect the latest information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListUsersRequest
*/
func (a *UserAPIService) ListUsers(ctx context.Context) ApiListUsersRequest {
	return ApiListUsersRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return []User
func (a *UserAPIService) ListUsersExecute(r ApiListUsersRequest) ([]User, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []User
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.ListUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sortBy", parameterToString(*r.sortBy, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
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
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
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

type ApiReplaceUserRequest struct {
	ctx        context.Context
	ApiService UserAPI
	id         string
	user       *UpdateUserRequest
	strict     *bool
	ifMatch    *string
	retryCount int32
}

func (r ApiReplaceUserRequest) User(user UpdateUserRequest) ApiReplaceUserRequest {
	r.user = &user
	return r
}

// If &#x60;true&#x60;, validates against minimum age and history password policy
func (r ApiReplaceUserRequest) Strict(strict bool) ApiReplaceUserRequest {
	r.strict = &strict
	return r
}

// The ETag value of the user&#39;s expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags).
func (r ApiReplaceUserRequest) IfMatch(ifMatch string) ApiReplaceUserRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiReplaceUserRequest) Execute() (*User, *APIResponse, error) {
	return r.ApiService.ReplaceUserExecute(r)
}

/*
ReplaceUser Replace a user

Replaces a user's profile, credentials, or both using strict-update semantics.

All profile properties must be specified when updating a user's profile with a `PUT` method. Any property not specified in the request is deleted.
> **Important:** Don't use a `PUT` method for partial updates.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	@return ApiReplaceUserRequest
*/
func (a *UserAPIService) ReplaceUser(ctx context.Context, id string) ApiReplaceUserRequest {
	return ApiReplaceUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return User
func (a *UserAPIService) ReplaceUserExecute(r ApiReplaceUserRequest) (*User, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.ReplaceUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return localVarReturnValue, nil, reportError("user is required and must be specified")
	}

	if r.strict != nil {
		localVarQueryParams.Add("strict", parameterToString(*r.strict, ""))
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
	if r.ifMatch != nil {
		localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	}
	// body params
	localVarPostBody = r.user
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

type ApiUpdateUserRequest struct {
	ctx        context.Context
	ApiService UserAPI
	id         string
	user       *UpdateUserRequest
	strict     *bool
	ifMatch    *string
	retryCount int32
}

func (r ApiUpdateUserRequest) User(user UpdateUserRequest) ApiUpdateUserRequest {
	r.user = &user
	return r
}

// If true, validates against minimum age and history password policy
func (r ApiUpdateUserRequest) Strict(strict bool) ApiUpdateUserRequest {
	r.strict = &strict
	return r
}

// The ETag value of the user&#39;s expected current state. This becomes a conditional request used for concurrency control. See [Conditional Requests and Entity Tags](/#conditional-requests-and-entity-tags).
func (r ApiUpdateUserRequest) IfMatch(ifMatch string) ApiUpdateUserRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUpdateUserRequest) Execute() (*User, *APIResponse, error) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
UpdateUser Update a user

Updates a user's profile or credentials with partial update semantics.

> **Important:** Use the `POST` method for partial updates. Unspecified properties are set to null with `PUT`.

`profile` and `credentials` can be updated independently or together with a single request.
> **Note**: Currently, the user type of a user can only be changed via a full replacement PUT operation. If the request parameters of a partial update include the type element from the user object,
the value must match the existing type of the user. Only admins are permitted to change the user type of a user; end users are not allowed to change their own user type.

> **Note**: To update a current user's profile with partial semantics, the `/api/v1/users/me` endpoint can be invoked.
>
> A user can only update profile properties for which the user has write access. Within the profile, if the user tries to update the primary or the secondary email IDs, verification emails are sent to those email IDs, and the fields are updated only upon verification.

If you are using this endpoint to set a password, it sets a password without validating existing user credentials. This is an administrative operation. For operations that validate credentials, refer to the [Reset password](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/resetPassword), [Start forgot password flow](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/forgotPassword), and [Update password](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserCred/#tag/UserCred/operation/changePassword) endpoints.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user
	@return ApiUpdateUserRequest
*/
func (a *UserAPIService) UpdateUser(ctx context.Context, id string) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return User
func (a *UserAPIService) UpdateUserExecute(r ApiUpdateUserRequest) (*User, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAPIService.UpdateUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return localVarReturnValue, nil, reportError("user is required and must be specified")
	}

	if r.strict != nil {
		localVarQueryParams.Add("strict", parameterToString(*r.strict, ""))
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
	if r.ifMatch != nil {
		localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	}
	// body params
	localVarPostBody = r.user
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

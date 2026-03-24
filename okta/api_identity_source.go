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

type IdentitySourceAPI interface {

	/*
		CreateIdentitySourceGroups Create an identity source group

		Creates a group in an identity source for the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@return ApiCreateIdentitySourceGroupsRequest
	*/
	CreateIdentitySourceGroups(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceGroupsRequest

	// CreateIdentitySourceGroupsExecute executes the request
	//  @return GroupsResponseSchema
	CreateIdentitySourceGroupsExecute(r ApiCreateIdentitySourceGroupsRequest) (*GroupsResponseSchema, *APIResponse, error)

	/*
		CreateIdentitySourceGroupsMemberships Create the memberships for the given identity source group

		Creates the group memberships for the given identity source group

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@return ApiCreateIdentitySourceGroupsMembershipsRequest
	*/
	CreateIdentitySourceGroupsMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiCreateIdentitySourceGroupsMembershipsRequest

	// CreateIdentitySourceGroupsMembershipsExecute executes the request
	CreateIdentitySourceGroupsMembershipsExecute(r ApiCreateIdentitySourceGroupsMembershipsRequest) (*APIResponse, error)

	/*
		CreateIdentitySourceSession Create an identity source session

		Creates an identity source session for the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@return ApiCreateIdentitySourceSessionRequest
	*/
	CreateIdentitySourceSession(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceSessionRequest

	// CreateIdentitySourceSessionExecute executes the request
	//  @return IdentitySourceSession
	CreateIdentitySourceSessionExecute(r ApiCreateIdentitySourceSessionRequest) (*IdentitySourceSession, *APIResponse, error)

	/*
		CreateIdentitySourceUser Create an identity source user

		Creates a user in an identity source for the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@return ApiCreateIdentitySourceUserRequest
	*/
	CreateIdentitySourceUser(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceUserRequest

	// CreateIdentitySourceUserExecute executes the request
	CreateIdentitySourceUserExecute(r ApiCreateIdentitySourceUserRequest) (*APIResponse, error)

	/*
		DeleteIdentitySourceGroup Delete an identity source group

		Deletes a group in an identity source for a given identity source ID and group ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@return ApiDeleteIdentitySourceGroupRequest
	*/
	DeleteIdentitySourceGroup(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiDeleteIdentitySourceGroupRequest

	// DeleteIdentitySourceGroupExecute executes the request
	DeleteIdentitySourceGroupExecute(r ApiDeleteIdentitySourceGroupRequest) (*APIResponse, error)

	/*
		DeleteIdentitySourceGroupMemberships Delete the memberships for the specified identity source group

		Deletes group memberships for the specified identity source group using member external IDs

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@param memberExternalId The external ID of the identity source user
		@return ApiDeleteIdentitySourceGroupMembershipsRequest
	*/
	DeleteIdentitySourceGroupMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string, memberExternalId string) ApiDeleteIdentitySourceGroupMembershipsRequest

	// DeleteIdentitySourceGroupMembershipsExecute executes the request
	DeleteIdentitySourceGroupMembershipsExecute(r ApiDeleteIdentitySourceGroupMembershipsRequest) (*APIResponse, error)

	/*
		DeleteIdentitySourceSession Delete an identity source session

		Deletes an identity source session for a given identity source ID and session Id

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiDeleteIdentitySourceSessionRequest
	*/
	DeleteIdentitySourceSession(ctx context.Context, identitySourceId string, sessionId string) ApiDeleteIdentitySourceSessionRequest

	// DeleteIdentitySourceSessionExecute executes the request
	DeleteIdentitySourceSessionExecute(r ApiDeleteIdentitySourceSessionRequest) (*APIResponse, error)

	/*
		DeleteIdentitySourceUser Delete an identity source user

		Deletes a user in an identity source for the given identity source instance and external ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param externalId The external ID of the user
		@return ApiDeleteIdentitySourceUserRequest
	*/
	DeleteIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiDeleteIdentitySourceUserRequest

	// DeleteIdentitySourceUserExecute executes the request
	DeleteIdentitySourceUserExecute(r ApiDeleteIdentitySourceUserRequest) (*APIResponse, error)

	/*
		GetIdentitySourceGroup Retrieve an identity source group

		Retrieves a group from an identity source for a given identity source ID and group ID or external ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@return ApiGetIdentitySourceGroupRequest
	*/
	GetIdentitySourceGroup(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiGetIdentitySourceGroupRequest

	// GetIdentitySourceGroupExecute executes the request
	//  @return GroupsResponseSchema
	GetIdentitySourceGroupExecute(r ApiGetIdentitySourceGroupRequest) (*GroupsResponseSchema, *APIResponse, error)

	/*
		GetIdentitySourceGroupMemberships Retrieve the memberships for the given identity source group

		Retrieves the group memberships for the given identity source group in the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@return ApiGetIdentitySourceGroupMembershipsRequest
	*/
	GetIdentitySourceGroupMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiGetIdentitySourceGroupMembershipsRequest

	// GetIdentitySourceGroupMembershipsExecute executes the request
	//  @return GroupMembershipsResponseSchema
	GetIdentitySourceGroupMembershipsExecute(r ApiGetIdentitySourceGroupMembershipsRequest) (*GroupMembershipsResponseSchema, *APIResponse, error)

	/*
		GetIdentitySourceSession Retrieve an identity source session

		Retrieves an identity source session for a given identity source ID and session ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiGetIdentitySourceSessionRequest
	*/
	GetIdentitySourceSession(ctx context.Context, identitySourceId string, sessionId string) ApiGetIdentitySourceSessionRequest

	// GetIdentitySourceSessionExecute executes the request
	//  @return IdentitySourceSession
	GetIdentitySourceSessionExecute(r ApiGetIdentitySourceSessionRequest) (*IdentitySourceSession, *APIResponse, error)

	/*
		GetIdentitySourceUser Retrieve an identity source user

		Retrieves a user by external ID in an identity source for the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param externalId The external ID of the user
		@return ApiGetIdentitySourceUserRequest
	*/
	GetIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiGetIdentitySourceUserRequest

	// GetIdentitySourceUserExecute executes the request
	//  @return UserResponseSchema
	GetIdentitySourceUserExecute(r ApiGetIdentitySourceUserRequest) (*UserResponseSchema, *APIResponse, error)

	/*
		ListIdentitySourceSessions List all identity source sessions

		Lists all identity source sessions for the given identity source instance

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@return ApiListIdentitySourceSessionsRequest
	*/
	ListIdentitySourceSessions(ctx context.Context, identitySourceId string) ApiListIdentitySourceSessionsRequest

	// ListIdentitySourceSessionsExecute executes the request
	//  @return []IdentitySourceSession
	ListIdentitySourceSessionsExecute(r ApiListIdentitySourceSessionsRequest) ([]IdentitySourceSession, *APIResponse, error)

	/*
		ReplaceExistingIdentitySourceUser Replace an existing identity source user

		Replaces an existing user for the given identity source instance and external ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param externalId The external ID of the user
		@return ApiReplaceExistingIdentitySourceUserRequest
	*/
	ReplaceExistingIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiReplaceExistingIdentitySourceUserRequest

	// ReplaceExistingIdentitySourceUserExecute executes the request
	//  @return UserResponseSchema
	ReplaceExistingIdentitySourceUserExecute(r ApiReplaceExistingIdentitySourceUserRequest) (*UserResponseSchema, *APIResponse, error)

	/*
		StartImportFromIdentitySource Start the import from the identity source

		Starts the import from the identity source described by the uploaded bulk operations

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiStartImportFromIdentitySourceRequest
	*/
	StartImportFromIdentitySource(ctx context.Context, identitySourceId string, sessionId string) ApiStartImportFromIdentitySourceRequest

	// StartImportFromIdentitySourceExecute executes the request
	//  @return IdentitySourceSession
	StartImportFromIdentitySourceExecute(r ApiStartImportFromIdentitySourceRequest) (*IdentitySourceSession, *APIResponse, error)

	/*
		UpdateIdentitySourceGroups Update an identity source group

		Updates a group to an identity source for the given identity source instance and group ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param groupOrExternalId The Okta group ID or external ID of the identity source group
		@return ApiUpdateIdentitySourceGroupsRequest
	*/
	UpdateIdentitySourceGroups(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiUpdateIdentitySourceGroupsRequest

	// UpdateIdentitySourceGroupsExecute executes the request
	//  @return GroupsResponseSchema
	UpdateIdentitySourceGroupsExecute(r ApiUpdateIdentitySourceGroupsRequest) (*GroupsResponseSchema, *APIResponse, error)

	/*
		UpdateIdentitySourceUsers Update an identity source user

		Updates a user to an identity source for the given identity source instance and external ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param externalId The external ID of the user
		@return ApiUpdateIdentitySourceUsersRequest
	*/
	UpdateIdentitySourceUsers(ctx context.Context, identitySourceId string, externalId string) ApiUpdateIdentitySourceUsersRequest

	// UpdateIdentitySourceUsersExecute executes the request
	//  @return UserResponseSchema
	UpdateIdentitySourceUsersExecute(r ApiUpdateIdentitySourceUsersRequest) (*UserResponseSchema, *APIResponse, error)

	/*
		UploadIdentitySourceDataForDelete Upload the data to be deleted in Okta

		Uploads external IDs of entities that need to be deleted in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceDataForDeleteRequest
	*/
	UploadIdentitySourceDataForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceDataForDeleteRequest

	// UploadIdentitySourceDataForDeleteExecute executes the request
	UploadIdentitySourceDataForDeleteExecute(r ApiUploadIdentitySourceDataForDeleteRequest) (*APIResponse, error)

	/*
		UploadIdentitySourceDataForUpsert Upload the data to be upserted in Okta

		Uploads entities that need to be inserted or updated in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceDataForUpsertRequest
	*/
	UploadIdentitySourceDataForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceDataForUpsertRequest

	// UploadIdentitySourceDataForUpsertExecute executes the request
	UploadIdentitySourceDataForUpsertExecute(r ApiUploadIdentitySourceDataForUpsertRequest) (*APIResponse, error)

	/*
		UploadIdentitySourceGroupMembershipsForDelete Upload the group memberships to be deleted in Okta

		Uploads the group memberships that need to be deleted in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceGroupMembershipsForDeleteRequest
	*/
	UploadIdentitySourceGroupMembershipsForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupMembershipsForDeleteRequest

	// UploadIdentitySourceGroupMembershipsForDeleteExecute executes the request
	UploadIdentitySourceGroupMembershipsForDeleteExecute(r ApiUploadIdentitySourceGroupMembershipsForDeleteRequest) (*APIResponse, error)

	/*
		UploadIdentitySourceGroupMembershipsForUpsert Upload the group memberships to be upserted in Okta

		Uploads the group memberships that need to be inserted or updated in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceGroupMembershipsForUpsertRequest
	*/
	UploadIdentitySourceGroupMembershipsForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupMembershipsForUpsertRequest

	// UploadIdentitySourceGroupMembershipsForUpsertExecute executes the request
	UploadIdentitySourceGroupMembershipsForUpsertExecute(r ApiUploadIdentitySourceGroupMembershipsForUpsertRequest) (*APIResponse, error)

	/*
		UploadIdentitySourceGroupsDataForDelete Upload the group external IDs to be deleted in Okta

		Uploads external IDs of groups that need to be deleted in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceGroupsDataForDeleteRequest
	*/
	UploadIdentitySourceGroupsDataForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupsDataForDeleteRequest

	// UploadIdentitySourceGroupsDataForDeleteExecute executes the request
	UploadIdentitySourceGroupsDataForDeleteExecute(r ApiUploadIdentitySourceGroupsDataForDeleteRequest) (*APIResponse, error)

	/*
		UploadIdentitySourceGroupsForUpsert Upload the group profiles without memberships to be upserted in Okta

		Uploads the group profiles without memberships that need to be inserted or updated in Okta from the identity source for the given session

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param identitySourceId The ID of the identity source for which the session is created
		@param sessionId The ID of the identity source session
		@return ApiUploadIdentitySourceGroupsForUpsertRequest
	*/
	UploadIdentitySourceGroupsForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupsForUpsertRequest

	// UploadIdentitySourceGroupsForUpsertExecute executes the request
	UploadIdentitySourceGroupsForUpsertExecute(r ApiUploadIdentitySourceGroupsForUpsertRequest) (*APIResponse, error)
}

// IdentitySourceAPIService IdentitySourceAPI service
type IdentitySourceAPIService service

type ApiCreateIdentitySourceGroupsRequest struct {
	ctx                 context.Context
	ApiService          IdentitySourceAPI
	identitySourceId    string
	groupsRequestSchema *GroupsRequestSchema
	retryCount          int32
}

func (r ApiCreateIdentitySourceGroupsRequest) GroupsRequestSchema(groupsRequestSchema GroupsRequestSchema) ApiCreateIdentitySourceGroupsRequest {
	r.groupsRequestSchema = &groupsRequestSchema
	return r
}

func (r ApiCreateIdentitySourceGroupsRequest) Execute() (*GroupsResponseSchema, *APIResponse, error) {
	return r.ApiService.CreateIdentitySourceGroupsExecute(r)
}

/*
CreateIdentitySourceGroups Create an identity source group

Creates a group in an identity source for the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@return ApiCreateIdentitySourceGroupsRequest
*/
func (a *IdentitySourceAPIService) CreateIdentitySourceGroups(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceGroupsRequest {
	return ApiCreateIdentitySourceGroupsRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return GroupsResponseSchema
func (a *IdentitySourceAPIService) CreateIdentitySourceGroupsExecute(r ApiCreateIdentitySourceGroupsRequest) (*GroupsResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.CreateIdentitySourceGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)

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
	localVarPostBody = r.groupsRequestSchema
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

type ApiCreateIdentitySourceGroupsMembershipsRequest struct {
	ctx                     context.Context
	ApiService              IdentitySourceAPI
	identitySourceId        string
	groupOrExternalId       string
	membershipRequestSchema *MembershipRequestSchema
	retryCount              int32
}

func (r ApiCreateIdentitySourceGroupsMembershipsRequest) MembershipRequestSchema(membershipRequestSchema MembershipRequestSchema) ApiCreateIdentitySourceGroupsMembershipsRequest {
	r.membershipRequestSchema = &membershipRequestSchema
	return r
}

func (r ApiCreateIdentitySourceGroupsMembershipsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateIdentitySourceGroupsMembershipsExecute(r)
}

/*
CreateIdentitySourceGroupsMemberships Create the memberships for the given identity source group

Creates the group memberships for the given identity source group

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@return ApiCreateIdentitySourceGroupsMembershipsRequest
*/
func (a *IdentitySourceAPIService) CreateIdentitySourceGroupsMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiCreateIdentitySourceGroupsMembershipsRequest {
	return ApiCreateIdentitySourceGroupsMembershipsRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) CreateIdentitySourceGroupsMembershipsExecute(r ApiCreateIdentitySourceGroupsMembershipsRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.CreateIdentitySourceGroupsMemberships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)

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
	localVarPostBody = r.membershipRequestSchema
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

type ApiCreateIdentitySourceSessionRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	retryCount       int32
}

func (r ApiCreateIdentitySourceSessionRequest) Execute() (*IdentitySourceSession, *APIResponse, error) {
	return r.ApiService.CreateIdentitySourceSessionExecute(r)
}

/*
CreateIdentitySourceSession Create an identity source session

Creates an identity source session for the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@return ApiCreateIdentitySourceSessionRequest
*/
func (a *IdentitySourceAPIService) CreateIdentitySourceSession(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceSessionRequest {
	return ApiCreateIdentitySourceSessionRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return IdentitySourceSession
func (a *IdentitySourceAPIService) CreateIdentitySourceSessionExecute(r ApiCreateIdentitySourceSessionRequest) (*IdentitySourceSession, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentitySourceSession
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.CreateIdentitySourceSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)

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

type ApiCreateIdentitySourceUserRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	userRequestSchema *UserRequestSchema
	retryCount        int32
}

func (r ApiCreateIdentitySourceUserRequest) UserRequestSchema(userRequestSchema UserRequestSchema) ApiCreateIdentitySourceUserRequest {
	r.userRequestSchema = &userRequestSchema
	return r
}

func (r ApiCreateIdentitySourceUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateIdentitySourceUserExecute(r)
}

/*
CreateIdentitySourceUser Create an identity source user

Creates a user in an identity source for the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@return ApiCreateIdentitySourceUserRequest
*/
func (a *IdentitySourceAPIService) CreateIdentitySourceUser(ctx context.Context, identitySourceId string) ApiCreateIdentitySourceUserRequest {
	return ApiCreateIdentitySourceUserRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) CreateIdentitySourceUserExecute(r ApiCreateIdentitySourceUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.CreateIdentitySourceUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)

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
	localVarPostBody = r.userRequestSchema
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

type ApiDeleteIdentitySourceGroupRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	groupOrExternalId string
	retryCount        int32
}

func (r ApiDeleteIdentitySourceGroupRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteIdentitySourceGroupExecute(r)
}

/*
DeleteIdentitySourceGroup Delete an identity source group

Deletes a group in an identity source for a given identity source ID and group ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@return ApiDeleteIdentitySourceGroupRequest
*/
func (a *IdentitySourceAPIService) DeleteIdentitySourceGroup(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiDeleteIdentitySourceGroupRequest {
	return ApiDeleteIdentitySourceGroupRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) DeleteIdentitySourceGroupExecute(r ApiDeleteIdentitySourceGroupRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.DeleteIdentitySourceGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)

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

type ApiDeleteIdentitySourceGroupMembershipsRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	groupOrExternalId string
	memberExternalId  string
	retryCount        int32
}

func (r ApiDeleteIdentitySourceGroupMembershipsRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteIdentitySourceGroupMembershipsExecute(r)
}

/*
DeleteIdentitySourceGroupMemberships Delete the memberships for the specified identity source group

Deletes group memberships for the specified identity source group using member external IDs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@param memberExternalId The external ID of the identity source user
	@return ApiDeleteIdentitySourceGroupMembershipsRequest
*/
func (a *IdentitySourceAPIService) DeleteIdentitySourceGroupMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string, memberExternalId string) ApiDeleteIdentitySourceGroupMembershipsRequest {
	return ApiDeleteIdentitySourceGroupMembershipsRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		memberExternalId:  memberExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) DeleteIdentitySourceGroupMembershipsExecute(r ApiDeleteIdentitySourceGroupMembershipsRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.DeleteIdentitySourceGroupMemberships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership/{memberExternalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberExternalId"+"}", url.PathEscape(parameterToString(r.memberExternalId, "")), -1)

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

type ApiDeleteIdentitySourceSessionRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	sessionId        string
	retryCount       int32
}

func (r ApiDeleteIdentitySourceSessionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteIdentitySourceSessionExecute(r)
}

/*
DeleteIdentitySourceSession Delete an identity source session

Deletes an identity source session for a given identity source ID and session Id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiDeleteIdentitySourceSessionRequest
*/
func (a *IdentitySourceAPIService) DeleteIdentitySourceSession(ctx context.Context, identitySourceId string, sessionId string) ApiDeleteIdentitySourceSessionRequest {
	return ApiDeleteIdentitySourceSessionRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) DeleteIdentitySourceSessionExecute(r ApiDeleteIdentitySourceSessionRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.DeleteIdentitySourceSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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

type ApiDeleteIdentitySourceUserRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	externalId       string
	retryCount       int32
}

func (r ApiDeleteIdentitySourceUserRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteIdentitySourceUserExecute(r)
}

/*
DeleteIdentitySourceUser Delete an identity source user

Deletes a user in an identity source for the given identity source instance and external ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param externalId The external ID of the user
	@return ApiDeleteIdentitySourceUserRequest
*/
func (a *IdentitySourceAPIService) DeleteIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiDeleteIdentitySourceUserRequest {
	return ApiDeleteIdentitySourceUserRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		externalId:       externalId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) DeleteIdentitySourceUserExecute(r ApiDeleteIdentitySourceUserRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.DeleteIdentitySourceUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/users/{externalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalId"+"}", url.PathEscape(parameterToString(r.externalId, "")), -1)

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

type ApiGetIdentitySourceGroupRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	groupOrExternalId string
	retryCount        int32
}

func (r ApiGetIdentitySourceGroupRequest) Execute() (*GroupsResponseSchema, *APIResponse, error) {
	return r.ApiService.GetIdentitySourceGroupExecute(r)
}

/*
GetIdentitySourceGroup Retrieve an identity source group

Retrieves a group from an identity source for a given identity source ID and group ID or external ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@return ApiGetIdentitySourceGroupRequest
*/
func (a *IdentitySourceAPIService) GetIdentitySourceGroup(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiGetIdentitySourceGroupRequest {
	return ApiGetIdentitySourceGroupRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
//
//	@return GroupsResponseSchema
func (a *IdentitySourceAPIService) GetIdentitySourceGroupExecute(r ApiGetIdentitySourceGroupRequest) (*GroupsResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.GetIdentitySourceGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)

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

type ApiGetIdentitySourceGroupMembershipsRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	groupOrExternalId string
	after             *string
	limit             *int32
	retryCount        int32
}

// The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header).
func (r ApiGetIdentitySourceGroupMembershipsRequest) After(after string) ApiGetIdentitySourceGroupMembershipsRequest {
	r.after = &after
	return r
}

// Specifies the number of group membership results in a page. Okta recommends using a specific value other than the default or maximum. If your request times out, retry your request with a smaller &#x60;limit&#x60; and [page the results](https://developer.okta.com/docs/api/#pagination).
func (r ApiGetIdentitySourceGroupMembershipsRequest) Limit(limit int32) ApiGetIdentitySourceGroupMembershipsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetIdentitySourceGroupMembershipsRequest) Execute() (*GroupMembershipsResponseSchema, *APIResponse, error) {
	return r.ApiService.GetIdentitySourceGroupMembershipsExecute(r)
}

/*
GetIdentitySourceGroupMemberships Retrieve the memberships for the given identity source group

Retrieves the group memberships for the given identity source group in the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@return ApiGetIdentitySourceGroupMembershipsRequest
*/
func (a *IdentitySourceAPIService) GetIdentitySourceGroupMemberships(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiGetIdentitySourceGroupMembershipsRequest {
	return ApiGetIdentitySourceGroupMembershipsRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
//
//	@return GroupMembershipsResponseSchema
func (a *IdentitySourceAPIService) GetIdentitySourceGroupMembershipsExecute(r ApiGetIdentitySourceGroupMembershipsRequest) (*GroupMembershipsResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupMembershipsResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.GetIdentitySourceGroupMemberships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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

type ApiGetIdentitySourceSessionRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	sessionId        string
	retryCount       int32
}

func (r ApiGetIdentitySourceSessionRequest) Execute() (*IdentitySourceSession, *APIResponse, error) {
	return r.ApiService.GetIdentitySourceSessionExecute(r)
}

/*
GetIdentitySourceSession Retrieve an identity source session

Retrieves an identity source session for a given identity source ID and session ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiGetIdentitySourceSessionRequest
*/
func (a *IdentitySourceAPIService) GetIdentitySourceSession(ctx context.Context, identitySourceId string, sessionId string) ApiGetIdentitySourceSessionRequest {
	return ApiGetIdentitySourceSessionRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return IdentitySourceSession
func (a *IdentitySourceAPIService) GetIdentitySourceSessionExecute(r ApiGetIdentitySourceSessionRequest) (*IdentitySourceSession, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentitySourceSession
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.GetIdentitySourceSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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

type ApiGetIdentitySourceUserRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	externalId       string
	retryCount       int32
}

func (r ApiGetIdentitySourceUserRequest) Execute() (*UserResponseSchema, *APIResponse, error) {
	return r.ApiService.GetIdentitySourceUserExecute(r)
}

/*
GetIdentitySourceUser Retrieve an identity source user

Retrieves a user by external ID in an identity source for the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param externalId The external ID of the user
	@return ApiGetIdentitySourceUserRequest
*/
func (a *IdentitySourceAPIService) GetIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiGetIdentitySourceUserRequest {
	return ApiGetIdentitySourceUserRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		externalId:       externalId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return UserResponseSchema
func (a *IdentitySourceAPIService) GetIdentitySourceUserExecute(r ApiGetIdentitySourceUserRequest) (*UserResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.GetIdentitySourceUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/users/{externalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalId"+"}", url.PathEscape(parameterToString(r.externalId, "")), -1)

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

type ApiListIdentitySourceSessionsRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	retryCount       int32
}

func (r ApiListIdentitySourceSessionsRequest) Execute() ([]IdentitySourceSession, *APIResponse, error) {
	return r.ApiService.ListIdentitySourceSessionsExecute(r)
}

/*
ListIdentitySourceSessions List all identity source sessions

Lists all identity source sessions for the given identity source instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@return ApiListIdentitySourceSessionsRequest
*/
func (a *IdentitySourceAPIService) ListIdentitySourceSessions(ctx context.Context, identitySourceId string) ApiListIdentitySourceSessionsRequest {
	return ApiListIdentitySourceSessionsRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return []IdentitySourceSession
func (a *IdentitySourceAPIService) ListIdentitySourceSessionsExecute(r ApiListIdentitySourceSessionsRequest) ([]IdentitySourceSession, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []IdentitySourceSession
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.ListIdentitySourceSessions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)

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

type ApiReplaceExistingIdentitySourceUserRequest struct {
	ctx               context.Context
	ApiService        IdentitySourceAPI
	identitySourceId  string
	externalId        string
	userRequestSchema *UserRequestSchema
	retryCount        int32
}

func (r ApiReplaceExistingIdentitySourceUserRequest) UserRequestSchema(userRequestSchema UserRequestSchema) ApiReplaceExistingIdentitySourceUserRequest {
	r.userRequestSchema = &userRequestSchema
	return r
}

func (r ApiReplaceExistingIdentitySourceUserRequest) Execute() (*UserResponseSchema, *APIResponse, error) {
	return r.ApiService.ReplaceExistingIdentitySourceUserExecute(r)
}

/*
ReplaceExistingIdentitySourceUser Replace an existing identity source user

Replaces an existing user for the given identity source instance and external ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param externalId The external ID of the user
	@return ApiReplaceExistingIdentitySourceUserRequest
*/
func (a *IdentitySourceAPIService) ReplaceExistingIdentitySourceUser(ctx context.Context, identitySourceId string, externalId string) ApiReplaceExistingIdentitySourceUserRequest {
	return ApiReplaceExistingIdentitySourceUserRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		externalId:       externalId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return UserResponseSchema
func (a *IdentitySourceAPIService) ReplaceExistingIdentitySourceUserExecute(r ApiReplaceExistingIdentitySourceUserRequest) (*UserResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.ReplaceExistingIdentitySourceUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/users/{externalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalId"+"}", url.PathEscape(parameterToString(r.externalId, "")), -1)

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
	localVarPostBody = r.userRequestSchema
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

type ApiStartImportFromIdentitySourceRequest struct {
	ctx              context.Context
	ApiService       IdentitySourceAPI
	identitySourceId string
	sessionId        string
	retryCount       int32
}

func (r ApiStartImportFromIdentitySourceRequest) Execute() (*IdentitySourceSession, *APIResponse, error) {
	return r.ApiService.StartImportFromIdentitySourceExecute(r)
}

/*
StartImportFromIdentitySource Start the import from the identity source

Starts the import from the identity source described by the uploaded bulk operations

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiStartImportFromIdentitySourceRequest
*/
func (a *IdentitySourceAPIService) StartImportFromIdentitySource(ctx context.Context, identitySourceId string, sessionId string) ApiStartImportFromIdentitySourceRequest {
	return ApiStartImportFromIdentitySourceRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return IdentitySourceSession
func (a *IdentitySourceAPIService) StartImportFromIdentitySourceExecute(r ApiStartImportFromIdentitySourceRequest) (*IdentitySourceSession, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdentitySourceSession
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.StartImportFromIdentitySource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/start-import"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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

type ApiUpdateIdentitySourceGroupsRequest struct {
	ctx                 context.Context
	ApiService          IdentitySourceAPI
	identitySourceId    string
	groupOrExternalId   string
	groupsRequestSchema *GroupsRequestSchema
	retryCount          int32
}

func (r ApiUpdateIdentitySourceGroupsRequest) GroupsRequestSchema(groupsRequestSchema GroupsRequestSchema) ApiUpdateIdentitySourceGroupsRequest {
	r.groupsRequestSchema = &groupsRequestSchema
	return r
}

func (r ApiUpdateIdentitySourceGroupsRequest) Execute() (*GroupsResponseSchema, *APIResponse, error) {
	return r.ApiService.UpdateIdentitySourceGroupsExecute(r)
}

/*
UpdateIdentitySourceGroups Update an identity source group

Updates a group to an identity source for the given identity source instance and group ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param groupOrExternalId The Okta group ID or external ID of the identity source group
	@return ApiUpdateIdentitySourceGroupsRequest
*/
func (a *IdentitySourceAPIService) UpdateIdentitySourceGroups(ctx context.Context, identitySourceId string, groupOrExternalId string) ApiUpdateIdentitySourceGroupsRequest {
	return ApiUpdateIdentitySourceGroupsRequest{
		ApiService:        a,
		ctx:               ctx,
		identitySourceId:  identitySourceId,
		groupOrExternalId: groupOrExternalId,
		retryCount:        0,
	}
}

// Execute executes the request
//
//	@return GroupsResponseSchema
func (a *IdentitySourceAPIService) UpdateIdentitySourceGroupsExecute(r ApiUpdateIdentitySourceGroupsRequest) (*GroupsResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UpdateIdentitySourceGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/groups/{groupOrExternalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupOrExternalId"+"}", url.PathEscape(parameterToString(r.groupOrExternalId, "")), -1)

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
	localVarPostBody = r.groupsRequestSchema
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

type ApiUpdateIdentitySourceUsersRequest struct {
	ctx                      context.Context
	ApiService               IdentitySourceAPI
	identitySourceId         string
	externalId               string
	usersUpdateRequestSchema *UsersUpdateRequestSchema
	retryCount               int32
}

func (r ApiUpdateIdentitySourceUsersRequest) UsersUpdateRequestSchema(usersUpdateRequestSchema UsersUpdateRequestSchema) ApiUpdateIdentitySourceUsersRequest {
	r.usersUpdateRequestSchema = &usersUpdateRequestSchema
	return r
}

func (r ApiUpdateIdentitySourceUsersRequest) Execute() (*UserResponseSchema, *APIResponse, error) {
	return r.ApiService.UpdateIdentitySourceUsersExecute(r)
}

/*
UpdateIdentitySourceUsers Update an identity source user

Updates a user to an identity source for the given identity source instance and external ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param externalId The external ID of the user
	@return ApiUpdateIdentitySourceUsersRequest
*/
func (a *IdentitySourceAPIService) UpdateIdentitySourceUsers(ctx context.Context, identitySourceId string, externalId string) ApiUpdateIdentitySourceUsersRequest {
	return ApiUpdateIdentitySourceUsersRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		externalId:       externalId,
		retryCount:       0,
	}
}

// Execute executes the request
//
//	@return UserResponseSchema
func (a *IdentitySourceAPIService) UpdateIdentitySourceUsersExecute(r ApiUpdateIdentitySourceUsersRequest) (*UserResponseSchema, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponseSchema
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UpdateIdentitySourceUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/users/{externalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalId"+"}", url.PathEscape(parameterToString(r.externalId, "")), -1)

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
	localVarPostBody = r.usersUpdateRequestSchema
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

type ApiUploadIdentitySourceDataForDeleteRequest struct {
	ctx                   context.Context
	ApiService            IdentitySourceAPI
	identitySourceId      string
	sessionId             string
	bulkDeleteRequestBody *BulkDeleteRequestBody
	retryCount            int32
}

func (r ApiUploadIdentitySourceDataForDeleteRequest) BulkDeleteRequestBody(bulkDeleteRequestBody BulkDeleteRequestBody) ApiUploadIdentitySourceDataForDeleteRequest {
	r.bulkDeleteRequestBody = &bulkDeleteRequestBody
	return r
}

func (r ApiUploadIdentitySourceDataForDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceDataForDeleteExecute(r)
}

/*
UploadIdentitySourceDataForDelete Upload the data to be deleted in Okta

Uploads external IDs of entities that need to be deleted in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceDataForDeleteRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceDataForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceDataForDeleteRequest {
	return ApiUploadIdentitySourceDataForDeleteRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceDataForDeleteExecute(r ApiUploadIdentitySourceDataForDeleteRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceDataForDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-delete"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkDeleteRequestBody
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

type ApiUploadIdentitySourceDataForUpsertRequest struct {
	ctx                   context.Context
	ApiService            IdentitySourceAPI
	identitySourceId      string
	sessionId             string
	bulkUpsertRequestBody *BulkUpsertRequestBody
	retryCount            int32
}

func (r ApiUploadIdentitySourceDataForUpsertRequest) BulkUpsertRequestBody(bulkUpsertRequestBody BulkUpsertRequestBody) ApiUploadIdentitySourceDataForUpsertRequest {
	r.bulkUpsertRequestBody = &bulkUpsertRequestBody
	return r
}

func (r ApiUploadIdentitySourceDataForUpsertRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceDataForUpsertExecute(r)
}

/*
UploadIdentitySourceDataForUpsert Upload the data to be upserted in Okta

Uploads entities that need to be inserted or updated in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceDataForUpsertRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceDataForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceDataForUpsertRequest {
	return ApiUploadIdentitySourceDataForUpsertRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceDataForUpsertExecute(r ApiUploadIdentitySourceDataForUpsertRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceDataForUpsert")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-upsert"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkUpsertRequestBody
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

type ApiUploadIdentitySourceGroupMembershipsForDeleteRequest struct {
	ctx                                   context.Context
	ApiService                            IdentitySourceAPI
	identitySourceId                      string
	sessionId                             string
	bulkGroupMembershipsDeleteRequestBody *BulkGroupMembershipsDeleteRequestBody
	retryCount                            int32
}

func (r ApiUploadIdentitySourceGroupMembershipsForDeleteRequest) BulkGroupMembershipsDeleteRequestBody(bulkGroupMembershipsDeleteRequestBody BulkGroupMembershipsDeleteRequestBody) ApiUploadIdentitySourceGroupMembershipsForDeleteRequest {
	r.bulkGroupMembershipsDeleteRequestBody = &bulkGroupMembershipsDeleteRequestBody
	return r
}

func (r ApiUploadIdentitySourceGroupMembershipsForDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceGroupMembershipsForDeleteExecute(r)
}

/*
UploadIdentitySourceGroupMembershipsForDelete Upload the group memberships to be deleted in Okta

Uploads the group memberships that need to be deleted in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceGroupMembershipsForDeleteRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupMembershipsForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupMembershipsForDeleteRequest {
	return ApiUploadIdentitySourceGroupMembershipsForDeleteRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupMembershipsForDeleteExecute(r ApiUploadIdentitySourceGroupMembershipsForDeleteRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceGroupMembershipsForDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-delete"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkGroupMembershipsDeleteRequestBody
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

type ApiUploadIdentitySourceGroupMembershipsForUpsertRequest struct {
	ctx                                   context.Context
	ApiService                            IdentitySourceAPI
	identitySourceId                      string
	sessionId                             string
	bulkGroupMembershipsUpsertRequestBody *BulkGroupMembershipsUpsertRequestBody
	retryCount                            int32
}

func (r ApiUploadIdentitySourceGroupMembershipsForUpsertRequest) BulkGroupMembershipsUpsertRequestBody(bulkGroupMembershipsUpsertRequestBody BulkGroupMembershipsUpsertRequestBody) ApiUploadIdentitySourceGroupMembershipsForUpsertRequest {
	r.bulkGroupMembershipsUpsertRequestBody = &bulkGroupMembershipsUpsertRequestBody
	return r
}

func (r ApiUploadIdentitySourceGroupMembershipsForUpsertRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceGroupMembershipsForUpsertExecute(r)
}

/*
UploadIdentitySourceGroupMembershipsForUpsert Upload the group memberships to be upserted in Okta

Uploads the group memberships that need to be inserted or updated in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceGroupMembershipsForUpsertRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupMembershipsForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupMembershipsForUpsertRequest {
	return ApiUploadIdentitySourceGroupMembershipsForUpsertRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupMembershipsForUpsertExecute(r ApiUploadIdentitySourceGroupMembershipsForUpsertRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceGroupMembershipsForUpsert")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-upsert"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkGroupMembershipsUpsertRequestBody
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

type ApiUploadIdentitySourceGroupsDataForDeleteRequest struct {
	ctx                        context.Context
	ApiService                 IdentitySourceAPI
	identitySourceId           string
	sessionId                  string
	bulkGroupDeleteRequestBody *BulkGroupDeleteRequestBody
	retryCount                 int32
}

func (r ApiUploadIdentitySourceGroupsDataForDeleteRequest) BulkGroupDeleteRequestBody(bulkGroupDeleteRequestBody BulkGroupDeleteRequestBody) ApiUploadIdentitySourceGroupsDataForDeleteRequest {
	r.bulkGroupDeleteRequestBody = &bulkGroupDeleteRequestBody
	return r
}

func (r ApiUploadIdentitySourceGroupsDataForDeleteRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceGroupsDataForDeleteExecute(r)
}

/*
UploadIdentitySourceGroupsDataForDelete Upload the group external IDs to be deleted in Okta

Uploads external IDs of groups that need to be deleted in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceGroupsDataForDeleteRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupsDataForDelete(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupsDataForDeleteRequest {
	return ApiUploadIdentitySourceGroupsDataForDeleteRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupsDataForDeleteExecute(r ApiUploadIdentitySourceGroupsDataForDeleteRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceGroupsDataForDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-delete"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkGroupDeleteRequestBody
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

type ApiUploadIdentitySourceGroupsForUpsertRequest struct {
	ctx                        context.Context
	ApiService                 IdentitySourceAPI
	identitySourceId           string
	sessionId                  string
	bulkGroupUpsertRequestBody *BulkGroupUpsertRequestBody
	retryCount                 int32
}

func (r ApiUploadIdentitySourceGroupsForUpsertRequest) BulkGroupUpsertRequestBody(bulkGroupUpsertRequestBody BulkGroupUpsertRequestBody) ApiUploadIdentitySourceGroupsForUpsertRequest {
	r.bulkGroupUpsertRequestBody = &bulkGroupUpsertRequestBody
	return r
}

func (r ApiUploadIdentitySourceGroupsForUpsertRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UploadIdentitySourceGroupsForUpsertExecute(r)
}

/*
UploadIdentitySourceGroupsForUpsert Upload the group profiles without memberships to be upserted in Okta

Uploads the group profiles without memberships that need to be inserted or updated in Okta from the identity source for the given session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param identitySourceId The ID of the identity source for which the session is created
	@param sessionId The ID of the identity source session
	@return ApiUploadIdentitySourceGroupsForUpsertRequest
*/
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupsForUpsert(ctx context.Context, identitySourceId string, sessionId string) ApiUploadIdentitySourceGroupsForUpsertRequest {
	return ApiUploadIdentitySourceGroupsForUpsertRequest{
		ApiService:       a,
		ctx:              ctx,
		identitySourceId: identitySourceId,
		sessionId:        sessionId,
		retryCount:       0,
	}
}

// Execute executes the request
func (a *IdentitySourceAPIService) UploadIdentitySourceGroupsForUpsertExecute(r ApiUploadIdentitySourceGroupsForUpsertRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentitySourceAPIService.UploadIdentitySourceGroupsForUpsert")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-upsert"
	localVarPath = strings.Replace(localVarPath, "{"+"identitySourceId"+"}", url.PathEscape(parameterToString(r.identitySourceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterToString(r.sessionId, "")), -1)

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
	localVarPostBody = r.bulkGroupUpsertRequestBody
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

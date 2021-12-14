# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](UserApi.md#ActivateUser) | **Post** /api/v1/users/{userId}/lifecycle/activate | Activate User
[**AddAllAppsAsTargetToRole**](UserApi.md#AddAllAppsAsTargetToRole) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | 
[**AddApplicationTargetToAdminRoleForUser**](UserApi.md#AddApplicationTargetToAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | 
[**AddApplicationTargetToAppAdminRoleForUser**](UserApi.md#AddApplicationTargetToAppAdminRoleForUser) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Add App Instance Target to App Administrator Role given to a User
[**AddGroupTargetToRole**](UserApi.md#AddGroupTargetToRole) | **Put** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | 
[**AssignRoleToUser**](UserApi.md#AssignRoleToUser) | **Post** /api/v1/users/{userId}/roles | 
[**ChangePassword**](UserApi.md#ChangePassword) | **Post** /api/v1/users/{userId}/credentials/change_password | Change Password
[**ChangeRecoveryQuestion**](UserApi.md#ChangeRecoveryQuestion) | **Post** /api/v1/users/{userId}/credentials/change_recovery_question | Change Recovery Question
[**ClearUserSessions**](UserApi.md#ClearUserSessions) | **Delete** /api/v1/users/{userId}/sessions | 
[**CreateUser**](UserApi.md#CreateUser) | **Post** /api/v1/users | Create User
[**DeactivateOrDeleteUser**](UserApi.md#DeactivateOrDeleteUser) | **Delete** /api/v1/users/{userId} | Delete User
[**DeactivateUser**](UserApi.md#DeactivateUser) | **Post** /api/v1/users/{userId}/lifecycle/deactivate | Deactivate User
[**ExpirePassword**](UserApi.md#ExpirePassword) | **Post** /api/v1/users/{userId}/lifecycle/expire_password | Expire Password
[**ExpirePasswordAndGetTemporaryPassword**](UserApi.md#ExpirePasswordAndGetTemporaryPassword) | **Post** /api/v1/users/{userId}/lifecycle/expire_password_with_temp_password | Expire Password and Set Temporary Password
[**ForgotPassword**](UserApi.md#ForgotPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password | Initiate Forgot Password
[**ForgotPasswordSetNewPassword**](UserApi.md#ForgotPasswordSetNewPassword) | **Post** /api/v1/users/{userId}/credentials/forgot_password_recovery_question | Reset Password with Recovery Question
[**GetLinkedObjectsForUser**](UserApi.md#GetLinkedObjectsForUser) | **Get** /api/v1/users/{userId}/linkedObjects/{relationshipName} | 
[**GetRefreshTokenForUserAndClient**](UserApi.md#GetRefreshTokenForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | 
[**GetUser**](UserApi.md#GetUser) | **Get** /api/v1/users/{userId} | Get User
[**GetUserGrant**](UserApi.md#GetUserGrant) | **Get** /api/v1/users/{userId}/grants/{grantId} | 
[**GetUserRole**](UserApi.md#GetUserRole) | **Get** /api/v1/users/{userId}/roles/{roleId} | 
[**ListAppLinks**](UserApi.md#ListAppLinks) | **Get** /api/v1/users/{userId}/appLinks | Get Assigned App Links
[**ListApplicationTargetsForApplicationAdministratorRoleForUser**](UserApi.md#ListApplicationTargetsForApplicationAdministratorRoleForUser) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps | 
[**ListAssignedRolesForUser**](UserApi.md#ListAssignedRolesForUser) | **Get** /api/v1/users/{userId}/roles | 
[**ListGrantsForUserAndClient**](UserApi.md#ListGrantsForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/grants | 
[**ListGroupTargetsForRole**](UserApi.md#ListGroupTargetsForRole) | **Get** /api/v1/users/{userId}/roles/{roleId}/targets/groups | 
[**ListRefreshTokensForUserAndClient**](UserApi.md#ListRefreshTokensForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens | 
[**ListUserClients**](UserApi.md#ListUserClients) | **Get** /api/v1/users/{userId}/clients | 
[**ListUserGrants**](UserApi.md#ListUserGrants) | **Get** /api/v1/users/{userId}/grants | 
[**ListUserGroups**](UserApi.md#ListUserGroups) | **Get** /api/v1/users/{userId}/groups | Get Member Groups
[**ListUserIdentityProviders**](UserApi.md#ListUserIdentityProviders) | **Get** /api/v1/users/{userId}/idps | Listing IdPs associated with a user
[**ListUsers**](UserApi.md#ListUsers) | **Get** /api/v1/users | List Users
[**PartialUpdateUser**](UserApi.md#PartialUpdateUser) | **Post** /api/v1/users/{userId} | 
[**ReactivateUser**](UserApi.md#ReactivateUser) | **Post** /api/v1/users/{userId}/lifecycle/reactivate | Reactivate User
[**RemoveApplicationTargetFromAdministratorRoleForUser**](UserApi.md#RemoveApplicationTargetFromAdministratorRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Remove App Instance Target to App Administrator Role given to a User
[**RemoveApplicationTargetFromApplicationAdministratorRoleForUser**](UserApi.md#RemoveApplicationTargetFromApplicationAdministratorRoleForUser) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/catalog/apps/{appName} | 
[**RemoveGroupTargetFromRole**](UserApi.md#RemoveGroupTargetFromRole) | **Delete** /api/v1/users/{userId}/roles/{roleId}/targets/groups/{groupId} | 
[**RemoveLinkedObjectForUser**](UserApi.md#RemoveLinkedObjectForUser) | **Delete** /api/v1/users/{userId}/linkedObjects/{relationshipName} | 
[**RemoveRoleFromUser**](UserApi.md#RemoveRoleFromUser) | **Delete** /api/v1/users/{userId}/roles/{roleId} | 
[**ResetFactors**](UserApi.md#ResetFactors) | **Post** /api/v1/users/{userId}/lifecycle/reset_factors | Reset Factors
[**ResetPassword**](UserApi.md#ResetPassword) | **Post** /api/v1/users/{userId}/lifecycle/reset_password | Reset Password
[**RevokeGrantsForUserAndClient**](UserApi.md#RevokeGrantsForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/grants | 
[**RevokeTokenForUserAndClient**](UserApi.md#RevokeTokenForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | 
[**RevokeTokensForUserAndClient**](UserApi.md#RevokeTokensForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens | 
[**RevokeUserGrant**](UserApi.md#RevokeUserGrant) | **Delete** /api/v1/users/{userId}/grants/{grantId} | 
[**RevokeUserGrants**](UserApi.md#RevokeUserGrants) | **Delete** /api/v1/users/{userId}/grants | 
[**SetLinkedObjectForUser**](UserApi.md#SetLinkedObjectForUser) | **Put** /api/v1/users/{associatedUserId}/linkedObjects/{primaryRelationshipName}/{primaryUserId} | Set Linked Object for User
[**SuspendUser**](UserApi.md#SuspendUser) | **Post** /api/v1/users/{userId}/lifecycle/suspend | Suspend User
[**UnlockUser**](UserApi.md#UnlockUser) | **Post** /api/v1/users/{userId}/lifecycle/unlock | Unlock User
[**UnsuspendUser**](UserApi.md#UnsuspendUser) | **Post** /api/v1/users/{userId}/lifecycle/unsuspend | Unsuspend User
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /api/v1/users/{userId} | Update User

# **ActivateUser**
> UserActivationToken ActivateUser(ctx, userId, sendEmail)
Activate User

Activates a user.  This operation can only be performed on users with a `STAGED` status.  Activation of a user is an asynchronous operation. The user will have the `transitioningToStatus` property with a value of `ACTIVE` during activation to indicate that the user hasn't completed the asynchronous operation.  The user will have a status of `ACTIVE` when the activation process is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **sendEmail** | **bool**| Sends an activation email to the user if true | [default to true]

### Return type

[**UserActivationToken**](UserActivationToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAllAppsAsTargetToRole**
> AddAllAppsAsTargetToRole(ctx, userId, roleId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddApplicationTargetToAdminRoleForUser**
> AddApplicationTargetToAdminRoleForUser(ctx, userId, roleId, appName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **appName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddApplicationTargetToAppAdminRoleForUser**
> AddApplicationTargetToAppAdminRoleForUser(ctx, userId, roleId, appName, applicationId)
Add App Instance Target to App Administrator Role given to a User

Add App Instance Target to App Administrator Role given to a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **appName** | **string**|  | 
  **applicationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddGroupTargetToRole**
> AddGroupTargetToRole(ctx, userId, roleId, groupId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignRoleToUser**
> Role AssignRoleToUser(ctx, body, userId, optional)


Assigns a role to a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssignRoleRequest**](AssignRoleRequest.md)|  | 
  **userId** | **string**|  | 
 **optional** | ***UserApiAssignRoleToUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiAssignRoleToUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disableNotifications** | **optional.**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePassword**
> UserCredentials ChangePassword(ctx, body, userId, optional)
Change Password

Changes a user's password by validating the user's current password. This operation can only be performed on users in `STAGED`, `ACTIVE`, `PASSWORD_EXPIRED`, or `RECOVERY` status that have a valid password credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangePasswordRequest**](ChangePasswordRequest.md)|  | 
  **userId** | **string**|  | 
 **optional** | ***UserApiChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiChangePasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strict** | **optional.**|  | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeRecoveryQuestion**
> UserCredentials ChangeRecoveryQuestion(ctx, body, userId)
Change Recovery Question

Changes a user's recovery question & answer credential by validating the user's current password.  This operation can only be performed on users in **STAGED**, **ACTIVE** or **RECOVERY** `status` that have a valid password credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCredentials**](UserCredentials.md)|  | 
  **userId** | **string**|  | 

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearUserSessions**
> ClearUserSessions(ctx, userId, optional)


Removes all active identity provider sessions. This forces the user to authenticate on the next operation. Optionally revokes OpenID Connect and OAuth refresh and access tokens issued to the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiClearUserSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiClearUserSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oauthTokens** | **optional.Bool**| Revoke issued OpenID Connect and OAuth refresh and access tokens | [default to false]

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, body, optional)
Create User

Creates a new user in your Okta organization with or without credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserRequest**](CreateUserRequest.md)|  | 
 **optional** | ***UserApiCreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiCreateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activate** | **optional.**| Executes activation lifecycle operation when creating the user | [default to true]
 **provider** | **optional.**| Indicates whether to create a user with a specified authentication provider | [default to false]
 **nextLogin** | [**optional.Interface of UserNextLogin**](.md)| With activate&#x3D;true, set nextLogin to \&quot;changePassword\&quot; to have the password be EXPIRED, so user must change it the next time they log in. | 

### Return type

[**User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateOrDeleteUser**
> DeactivateOrDeleteUser(ctx, userId, optional)
Delete User

Deletes a user permanently.  This operation can only be performed on users that have a `DEPROVISIONED` status.  **This action cannot be recovered!**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiDeactivateOrDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeactivateOrDeleteUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateUser**
> DeactivateUser(ctx, userId, optional)
Deactivate User

Deactivates a user.  This operation can only be performed on users that do not have a `DEPROVISIONED` status.  Deactivation of a user is an asynchronous operation.  The user will have the `transitioningToStatus` property with a value of `DEPROVISIONED` during deactivation to indicate that the user hasn't completed the asynchronous operation.  The user will have a status of `DEPROVISIONED` when the deactivation process is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiDeactivateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiDeactivateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpirePassword**
> User ExpirePassword(ctx, userId)
Expire Password

This operation transitions the user to the status of `PASSWORD_EXPIRED` so that the user is required to change their password at their next login.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExpirePasswordAndGetTemporaryPassword**
> TempPassword ExpirePasswordAndGetTemporaryPassword(ctx, userId)
Expire Password and Set Temporary Password

This operation transitions the user to the status of `PASSWORD_EXPIRED` so that the user is required to change their password at their next login, and also sets the user's password to a temporary password returned in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**TempPassword**](TempPassword.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForgotPassword**
> ForgotPasswordResponse ForgotPassword(ctx, userId, optional)
Initiate Forgot Password

Initiate forgot password flow. Generates a one-time token (OTT) that can be used to reset a user's password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiForgotPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiForgotPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **optional.Bool**|  | [default to true]

### Return type

[**ForgotPasswordResponse**](ForgotPasswordResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForgotPasswordSetNewPassword**
> UserCredentials ForgotPasswordSetNewPassword(ctx, userId, optional)
Reset Password with Recovery Question

Resets the user's password to the specified password if the provided answer to the recovery question is correct.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiForgotPasswordSetNewPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiForgotPasswordSetNewPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserCredentials**](UserCredentials.md)|  | 
 **sendEmail** | **optional.**|  | [default to true]

### Return type

[**UserCredentials**](UserCredentials.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedObjectsForUser**
> []ResponseLinks GetLinkedObjectsForUser(ctx, userId, relationshipName, optional)


Get linked objects for a user, relationshipName can be a primary or associated relationship name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **relationshipName** | **string**|  | 
 **optional** | ***UserApiGetLinkedObjectsForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetLinkedObjectsForUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to -1]

### Return type

[**[]ResponseLinks**](ResponseLinks.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRefreshTokenForUserAndClient**
> OAuth2RefreshToken GetRefreshTokenForUserAndClient(ctx, userId, clientId, tokenId, optional)


Gets a refresh token issued for the specified User and Client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 
  **tokenId** | **string**|  | 
 **optional** | ***UserApiGetRefreshTokenForUserAndClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetRefreshTokenForUserAndClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]
 **after** | **optional.String**|  | 

### Return type

[**OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, userId)
Get User

Fetches a user from your Okta organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGrant**
> OAuth2ScopeConsentGrant GetUserGrant(ctx, userId, grantId, optional)


Gets a grant for the specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **grantId** | **string**|  | 
 **optional** | ***UserApiGetUserGrantOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiGetUserGrantOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRole**
> Role GetUserRole(ctx, userId, roleId)


Gets role that is assigned to user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppLinks**
> []AppLink ListAppLinks(ctx, userId)
Get Assigned App Links

Fetches appLinks for all direct or indirect (via group membership) assigned applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]AppLink**](AppLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationTargetsForApplicationAdministratorRoleForUser**
> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForUser(ctx, userId, roleId, optional)


Lists all App targets for an `APP_ADMIN` Role assigned to a User. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
 **optional** | ***UserApiListApplicationTargetsForApplicationAdministratorRoleForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListApplicationTargetsForApplicationAdministratorRoleForUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]CatalogApplication**](CatalogApplication.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssignedRolesForUser**
> []Role ListAssignedRolesForUser(ctx, userId, optional)


Lists all roles assigned to a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiListAssignedRolesForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListAssignedRolesForUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 

### Return type

[**[]Role**](Role.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGrantsForUserAndClient**
> []OAuth2ScopeConsentGrant ListGrantsForUserAndClient(ctx, userId, clientId, optional)


Lists all grants for a specified user and client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 
 **optional** | ***UserApiListGrantsForUserAndClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListGrantsForUserAndClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupTargetsForRole**
> []Group ListGroupTargetsForRole(ctx, userId, roleId, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
 **optional** | ***UserApiListGroupTargetsForRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListGroupTargetsForRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRefreshTokensForUserAndClient**
> []OAuth2RefreshToken ListRefreshTokensForUserAndClient(ctx, userId, clientId, optional)


Lists all refresh tokens issued for the specified User and Client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 
 **optional** | ***UserApiListRefreshTokensForUserAndClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListRefreshTokensForUserAndClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserClients**
> []OAuth2Client ListUserClients(ctx, userId)


Lists all client resources for which the specified user has grants or tokens.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserGrants**
> []OAuth2ScopeConsentGrant ListUserGrants(ctx, userId, optional)


Lists all grants for the specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiListUserGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListUserGrantsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeId** | **optional.String**|  | 
 **expand** | **optional.String**|  | 
 **after** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserGroups**
> []Group ListUserGroups(ctx, userId)
Get Member Groups

Fetches the groups of which the user is a member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserIdentityProviders**
> []IdentityProvider ListUserIdentityProviders(ctx, userId)
Listing IdPs associated with a user

Lists the IdPs associated with the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**[]IdentityProvider**](IdentityProvider.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> []User ListUsers(ctx, optional)
List Users

Lists users in your organization with pagination in most cases.  A subset of users can be returned that match a supported filter expression or search criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserApiListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiListUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Finds a user that matches firstName, lastName, and email properties | 
 **after** | **optional.String**| Specifies the pagination cursor for the next page of users | 
 **limit** | **optional.Int32**| Specifies the number of results returned | [default to 10]
 **filter** | **optional.String**| Filters users with a supported expression for a subset of properties | 
 **search** | **optional.String**| Searches for users with a supported filtering  expression for most properties | 
 **sortBy** | **optional.String**|  | 
 **sortOrder** | **optional.String**|  | 

### Return type

[**[]User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateUser**
> User PartialUpdateUser(ctx, body, userId, optional)


Fetch a user by `id`, `login`, or `login shortname` if the short name is unambiguous.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 
  **userId** | **string**|  | 
 **optional** | ***UserApiPartialUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiPartialUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strict** | **optional.**|  | 

### Return type

[**User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateUser**
> UserActivationToken ReactivateUser(ctx, userId, optional)
Reactivate User

Reactivates a user.  This operation can only be performed on users with a `PROVISIONED` status.  This operation restarts the activation workflow if for some reason the user activation was not completed when using the activationToken from [Activate User](#activate-user).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
 **optional** | ***UserApiReactivateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiReactivateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmail** | **optional.Bool**| Sends an activation email to the user if true | [default to false]

### Return type

[**UserActivationToken**](UserActivationToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveApplicationTargetFromAdministratorRoleForUser**
> RemoveApplicationTargetFromAdministratorRoleForUser(ctx, userId, roleId, appName, applicationId)
Remove App Instance Target to App Administrator Role given to a User

Remove App Instance Target to App Administrator Role given to a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **appName** | **string**|  | 
  **applicationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveApplicationTargetFromApplicationAdministratorRoleForUser**
> RemoveApplicationTargetFromApplicationAdministratorRoleForUser(ctx, userId, roleId, appName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **appName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupTargetFromRole**
> RemoveGroupTargetFromRole(ctx, userId, roleId, groupId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLinkedObjectForUser**
> RemoveLinkedObjectForUser(ctx, userId, relationshipName)


Delete linked objects for a user, relationshipName can be ONLY a primary relationship name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **relationshipName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleFromUser**
> RemoveRoleFromUser(ctx, userId, roleId)


Unassigns a role from a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetFactors**
> ResetFactors(ctx, userId)
Reset Factors

This operation resets all factors for the specified user. All MFA factor enrollments returned to the unenrolled state. The user's status remains ACTIVE. This link is present only if the user is currently enrolled in one or more MFA factors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> ResetPasswordToken ResetPassword(ctx, userId, sendEmail)
Reset Password

Generates a one-time token (OTT) that can be used to reset a user's password.  The OTT link can be automatically emailed to the user or returned to the API caller and distributed using a custom flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **sendEmail** | **bool**|  | 

### Return type

[**ResetPasswordToken**](ResetPasswordToken.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeGrantsForUserAndClient**
> RevokeGrantsForUserAndClient(ctx, userId, clientId)


Revokes all grants for the specified user and client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeTokenForUserAndClient**
> RevokeTokenForUserAndClient(ctx, userId, clientId, tokenId)


Revokes the specified refresh token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 
  **tokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeTokensForUserAndClient**
> RevokeTokensForUserAndClient(ctx, userId, clientId)


Revokes all refresh tokens issued for the specified User and Client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **clientId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeUserGrant**
> RevokeUserGrant(ctx, userId, grantId)


Revokes one grant for a specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **grantId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeUserGrants**
> RevokeUserGrants(ctx, userId)


Revokes all grants for a specified user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLinkedObjectForUser**
> SetLinkedObjectForUser(ctx, associatedUserId, primaryRelationshipName, primaryUserId)
Set Linked Object for User

Sets a linked object for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **associatedUserId** | **string**|  | 
  **primaryRelationshipName** | **string**|  | 
  **primaryUserId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendUser**
> SuspendUser(ctx, userId)
Suspend User

Suspends a user.  This operation can only be performed on users with an `ACTIVE` status.  The user will have a status of `SUSPENDED` when the process is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlockUser**
> UnlockUser(ctx, userId)
Unlock User

Unlocks a user with a `LOCKED_OUT` status and returns them to `ACTIVE` status.  Users will be able to login with their current password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsuspendUser**
> UnsuspendUser(ctx, userId)
Unsuspend User

Unsuspends a user and returns them to the `ACTIVE` state.  This operation can only be performed on users that have a `SUSPENDED` status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, body, userId, optional)
Update User

Update a user's profile and/or credentials using strict-update semantics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 
  **userId** | **string**|  | 
 **optional** | ***UserApiUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUpdateUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **strict** | **optional.**|  | 

### Return type

[**User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


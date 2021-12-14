# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGroupRule**](GroupApi.md#ActivateGroupRule) | **Post** /api/v1/groups/rules/{ruleId}/lifecycle/activate | Activate a group Rule
[**AddApplicationInstanceTargetToAppAdminRoleGivenToGroup**](GroupApi.md#AddApplicationInstanceTargetToAppAdminRoleGivenToGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Add App Instance Target to App Administrator Role given to a Group
[**AddApplicationTargetToAdminRoleGivenToGroup**](GroupApi.md#AddApplicationTargetToAdminRoleGivenToGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | 
[**AddGroupTargetToGroupAdministratorRoleForGroup**](GroupApi.md#AddGroupTargetToGroupAdministratorRoleForGroup) | **Put** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Add Group Target for Group Role
[**AddUserToGroup**](GroupApi.md#AddUserToGroup) | **Put** /api/v1/groups/{groupId}/users/{userId} | Add User to Group
[**AssignRoleToGroup**](GroupApi.md#AssignRoleToGroup) | **Post** /api/v1/groups/{groupId}/roles | 
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /api/v1/groups | Add Group
[**CreateGroupRule**](GroupApi.md#CreateGroupRule) | **Post** /api/v1/groups/rules | Create Group Rule
[**DeactivateGroupRule**](GroupApi.md#DeactivateGroupRule) | **Post** /api/v1/groups/rules/{ruleId}/lifecycle/deactivate | Deactivate a group Rule
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /api/v1/groups/{groupId} | Remove Group
[**DeleteGroupRule**](GroupApi.md#DeleteGroupRule) | **Delete** /api/v1/groups/rules/{ruleId} | Delete a group Rule
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /api/v1/groups/{groupId} | List Group Rules
[**GetGroupRule**](GroupApi.md#GetGroupRule) | **Get** /api/v1/groups/rules/{ruleId} | Get Group Rule
[**GetRole**](GroupApi.md#GetRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId} | 
[**ListApplicationTargetsForApplicationAdministratorRoleForGroup**](GroupApi.md#ListApplicationTargetsForApplicationAdministratorRoleForGroup) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps | 
[**ListAssignedApplicationsForGroup**](GroupApi.md#ListAssignedApplicationsForGroup) | **Get** /api/v1/groups/{groupId}/apps | List Assigned Applications
[**ListGroupAssignedRoles**](GroupApi.md#ListGroupAssignedRoles) | **Get** /api/v1/groups/{groupId}/roles | 
[**ListGroupRules**](GroupApi.md#ListGroupRules) | **Get** /api/v1/groups/rules | List Group Rules
[**ListGroupTargetsForGroupRole**](GroupApi.md#ListGroupTargetsForGroupRole) | **Get** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups | List Group Targets for Group Role
[**ListGroupUsers**](GroupApi.md#ListGroupUsers) | **Get** /api/v1/groups/{groupId}/users | List Group Members
[**ListGroups**](GroupApi.md#ListGroups) | **Get** /api/v1/groups | List Groups
[**RemoveApplicationTargetFromAdministratorRoleGivenToGroup**](GroupApi.md#RemoveApplicationTargetFromAdministratorRoleGivenToGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName}/{applicationId} | Remove App Instance Target to App Administrator Role given to a Group
[**RemoveApplicationTargetFromApplicationAdministratorRoleGivenToGroup**](GroupApi.md#RemoveApplicationTargetFromApplicationAdministratorRoleGivenToGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/catalog/apps/{appName} | 
[**RemoveGroupTargetFromGroupAdministratorRoleGivenToGroup**](GroupApi.md#RemoveGroupTargetFromGroupAdministratorRoleGivenToGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId}/targets/groups/{targetGroupId} | Delete Group Target for Group Role
[**RemoveRoleFromGroup**](GroupApi.md#RemoveRoleFromGroup) | **Delete** /api/v1/groups/{groupId}/roles/{roleId} | 
[**RemoveUserFromGroup**](GroupApi.md#RemoveUserFromGroup) | **Delete** /api/v1/groups/{groupId}/users/{userId} | Remove User from Group
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Put** /api/v1/groups/{groupId} | Update Group
[**UpdateGroupRule**](GroupApi.md#UpdateGroupRule) | **Put** /api/v1/groups/rules/{ruleId} | 

# **ActivateGroupRule**
> ActivateGroupRule(ctx, ruleId)
Activate a group Rule

Activates a specific group rule by id from your organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddApplicationInstanceTargetToAppAdminRoleGivenToGroup**
> AddApplicationInstanceTargetToAppAdminRoleGivenToGroup(ctx, groupId, roleId, appName, applicationId)
Add App Instance Target to App Administrator Role given to a Group

Add App Instance Target to App Administrator Role given to a Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
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

# **AddApplicationTargetToAdminRoleGivenToGroup**
> AddApplicationTargetToAdminRoleGivenToGroup(ctx, groupId, roleId, appName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
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

# **AddGroupTargetToGroupAdministratorRoleForGroup**
> AddGroupTargetToGroupAdministratorRoleForGroup(ctx, groupId, roleId, targetGroupId)
Add Group Target for Group Role

Enumerates group targets for a group role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 
  **targetGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUserToGroup**
> AddUserToGroup(ctx, groupId, userId)
Add User to Group

Adds a user to a group with 'OKTA_GROUP' type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignRoleToGroup**
> Role AssignRoleToGroup(ctx, body, groupId, optional)


Assigns a Role to a Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssignRoleRequest**](AssignRoleRequest.md)|  | 
  **groupId** | **string**|  | 
 **optional** | ***GroupApiAssignRoleToGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiAssignRoleToGroupOpts struct
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

# **CreateGroup**
> Group CreateGroup(ctx, body)
Add Group

Adds a new group with `OKTA_GROUP` type to your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroupRule**
> GroupRule CreateGroupRule(ctx, body)
Create Group Rule

Creates a group rule to dynamically add users to the specified group if they match the condition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupRule**](GroupRule.md)|  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateGroupRule**
> DeactivateGroupRule(ctx, ruleId)
Deactivate a group Rule

Deactivates a specific group rule by id from your organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> DeleteGroup(ctx, groupId)
Remove Group

Removes a group with `OKTA_GROUP` type from your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupRule**
> DeleteGroupRule(ctx, ruleId, optional)
Delete a group Rule

Removes a specific group rule by id from your organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 
 **optional** | ***GroupApiDeleteGroupRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiDeleteGroupRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeUsers** | **optional.Bool**| Indicates whether to keep or remove users from groups assigned by this rule. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> Group GetGroup(ctx, groupId)
List Group Rules

Fetches a group from your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupRule**
> GroupRule GetGroupRule(ctx, ruleId, optional)
Get Group Rule

Fetches a specific group rule by id from your organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 
 **optional** | ***GroupApiGetGroupRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiGetGroupRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**|  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, groupId, roleId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplicationTargetsForApplicationAdministratorRoleForGroup**
> []CatalogApplication ListApplicationTargetsForApplicationAdministratorRoleForGroup(ctx, groupId, roleId, optional)


Lists all App targets for an `APP_ADMIN` Role assigned to a Group. This methods return list may include full Applications or Instances. The response for an instance will have an `ID` value, while Application will not have an ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 
 **optional** | ***GroupApiListApplicationTargetsForApplicationAdministratorRoleForGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListApplicationTargetsForApplicationAdministratorRoleForGroupOpts struct
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

# **ListAssignedApplicationsForGroup**
> []Application ListAssignedApplicationsForGroup(ctx, groupId, optional)
List Assigned Applications

Enumerates all applications that are assigned to a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***GroupApiListAssignedApplicationsForGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListAssignedApplicationsForGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.String**| Specifies the pagination cursor for the next page of apps | 
 **limit** | **optional.Int32**| Specifies the number of app results for a page | [default to 20]

### Return type

[**[]Application**](Application.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupAssignedRoles**
> []Role ListGroupAssignedRoles(ctx, groupId, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***GroupApiListGroupAssignedRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListGroupAssignedRolesOpts struct
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

# **ListGroupRules**
> []GroupRule ListGroupRules(ctx, optional)
List Group Rules

Lists all group rules for your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupApiListGroupRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListGroupRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Specifies the number of rule results in a page | [default to 50]
 **after** | **optional.String**| Specifies the pagination cursor for the next page of rules | 
 **search** | **optional.String**| Specifies the keyword to search fules for | 
 **expand** | **optional.String**| If specified as &#x60;groupIdToGroupNameMap&#x60;, then show group names | 

### Return type

[**[]GroupRule**](GroupRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupTargetsForGroupRole**
> []Group ListGroupTargetsForGroupRole(ctx, groupId, roleId, optional)
List Group Targets for Group Role

Enumerates group targets for a group role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 
 **optional** | ***GroupApiListGroupTargetsForGroupRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListGroupTargetsForGroupRoleOpts struct
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

# **ListGroupUsers**
> []User ListGroupUsers(ctx, groupId, optional)
List Group Members

Enumerates all users that are a member of a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
 **optional** | ***GroupApiListGroupUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListGroupUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.String**| Specifies the pagination cursor for the next page of users | 
 **limit** | **optional.Int32**| Specifies the number of user results in a page | [default to 1000]

### Return type

[**[]User**](User.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroups**
> []Group ListGroups(ctx, optional)
List Groups

Enumerates groups in your organization with pagination. A subset of groups can be returned that match a supported filter expression or query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupApiListGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupApiListGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Searches the name property of groups for matching value | 
 **search** | **optional.String**| Filter expression for groups | 
 **after** | **optional.String**| Specifies the pagination cursor for the next page of groups | 
 **limit** | **optional.Int32**| Specifies the number of group results in a page | [default to 10000]
 **expand** | **optional.String**| If specified, it causes additional metadata to be included in the response. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveApplicationTargetFromAdministratorRoleGivenToGroup**
> RemoveApplicationTargetFromAdministratorRoleGivenToGroup(ctx, groupId, roleId, appName, applicationId)
Remove App Instance Target to App Administrator Role given to a Group

Remove App Instance Target to App Administrator Role given to a Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
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

# **RemoveApplicationTargetFromApplicationAdministratorRoleGivenToGroup**
> RemoveApplicationTargetFromApplicationAdministratorRoleGivenToGroup(ctx, groupId, roleId, appName)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
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

# **RemoveGroupTargetFromGroupAdministratorRoleGivenToGroup**
> RemoveGroupTargetFromGroupAdministratorRoleGivenToGroup(ctx, groupId, roleId, targetGroupId)
Delete Group Target for Group Role

remove group target for a group role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 
  **targetGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRoleFromGroup**
> RemoveRoleFromGroup(ctx, groupId, roleId)


Unassigns a Role from a Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserFromGroup**
> RemoveUserFromGroup(ctx, groupId, userId)
Remove User from Group

Removes a user from a group with 'OKTA_GROUP' type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> Group UpdateGroup(ctx, body, groupId)
Update Group

Updates the profile for a group with `OKTA_GROUP` type from your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Group**](Group.md)|  | 
  **groupId** | **string**|  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupRule**
> GroupRule UpdateGroupRule(ctx, body, ruleId)


Updates a group rule. Only `INACTIVE` rules can be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupRule**](GroupRule.md)|  | 
  **ruleId** | **string**|  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


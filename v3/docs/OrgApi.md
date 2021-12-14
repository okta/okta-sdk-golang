# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendOktaSupport**](OrgApi.md#ExtendOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/extend | Extend Okta Support
[**GetOktaCommunicationSettings**](OrgApi.md#GetOktaCommunicationSettings) | **Get** /api/v1/org/privacy/oktaCommunication | Get Okta Communication Settings
[**GetOrgContactTypes**](OrgApi.md#GetOrgContactTypes) | **Get** /api/v1/org/contacts | Get org contact types
[**GetOrgContactUser**](OrgApi.md#GetOrgContactUser) | **Get** /api/v1/org/contacts/{contactType} | Get org contact user
[**GetOrgOktaSupportSettings**](OrgApi.md#GetOrgOktaSupportSettings) | **Get** /api/v1/org/privacy/oktaSupport | Get Okta Support settings
[**GetOrgPreferences**](OrgApi.md#GetOrgPreferences) | **Get** /api/v1/org/preferences | Get org preferences
[**GetOrgSettings**](OrgApi.md#GetOrgSettings) | **Get** /api/v1/org | Get org settings
[**GrantOktaSupport**](OrgApi.md#GrantOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/grant | Grant Okta Support
[**HideOktaUIFooter**](OrgApi.md#HideOktaUIFooter) | **Post** /api/v1/org/preferences/hideEndUserFooter | Show Okta UI Footer
[**OptInUsersToOktaCommunicationEmails**](OrgApi.md#OptInUsersToOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optIn | Opt in all users to Okta Communication emails
[**OptOutUsersFromOktaCommunicationEmails**](OrgApi.md#OptOutUsersFromOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optOut | Opt out all users from Okta Communication emails
[**PartialUpdateOrgSetting**](OrgApi.md#PartialUpdateOrgSetting) | **Post** /api/v1/org | Partial update Org Setting
[**RevokeOktaSupport**](OrgApi.md#RevokeOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/revoke | Extend Okta Support
[**ShowOktaUIFooter**](OrgApi.md#ShowOktaUIFooter) | **Post** /api/v1/org/preferences/showEndUserFooter | Show Okta UI Footer
[**UpdateOrgContactUser**](OrgApi.md#UpdateOrgContactUser) | **Put** /api/v1/org/contacts/{contactType} | Update org contact user
[**UpdateOrgSetting**](OrgApi.md#UpdateOrgSetting) | **Put** /api/v1/org | Update Org setting

# **ExtendOktaSupport**
> OrgOktaSupportSettingsObj ExtendOktaSupport(ctx, )
Extend Okta Support

Extends the length of time that Okta Support can access your org by 24 hours. This means that 24 hours are added to the remaining access time.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOktaCommunicationSettings**
> OrgOktaCommunicationSetting GetOktaCommunicationSettings(ctx, )
Get Okta Communication Settings

Gets Okta Communication Settings of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgContactTypes**
> []OrgContactTypeObj GetOrgContactTypes(ctx, )
Get org contact types

Gets Contact Types of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OrgContactTypeObj**](OrgContactTypeObj.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgContactUser**
> OrgContactUser GetOrgContactUser(ctx, contactType)
Get org contact user

Retrieves the URL of the User associated with the specified Contact Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactType** | **string**|  | 

### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgOktaSupportSettings**
> OrgOktaSupportSettingsObj GetOrgOktaSupportSettings(ctx, )
Get Okta Support settings

Gets Okta Support Settings of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgPreferences**
> OrgPreferences GetOrgPreferences(ctx, )
Get org preferences

Gets preferences of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSettings**
> OrgSetting GetOrgSettings(ctx, )
Get org settings

Get settings of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrantOktaSupport**
> OrgOktaSupportSettingsObj GrantOktaSupport(ctx, )
Grant Okta Support

Enables you to temporarily allow Okta Support to access your org as an administrator for eight hours.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HideOktaUIFooter**
> OrgPreferences HideOktaUIFooter(ctx, )
Show Okta UI Footer

Hide the Okta UI footer for all end users of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptInUsersToOktaCommunicationEmails**
> OrgOktaCommunicationSetting OptInUsersToOktaCommunicationEmails(ctx, )
Opt in all users to Okta Communication emails

Opts in all users of this org to Okta Communication emails.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptOutUsersFromOktaCommunicationEmails**
> OrgOktaCommunicationSetting OptOutUsersFromOktaCommunicationEmails(ctx, )
Opt out all users from Okta Communication emails

Opts out all users of this org from Okta Communication emails.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialUpdateOrgSetting**
> OrgSetting PartialUpdateOrgSetting(ctx, optional)
Partial update Org Setting

Partial update settings of your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrgApiPartialUpdateOrgSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgApiPartialUpdateOrgSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OrgSetting**](OrgSetting.md)|  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeOktaSupport**
> OrgOktaSupportSettingsObj RevokeOktaSupport(ctx, )
Extend Okta Support

Revokes Okta Support access to your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowOktaUIFooter**
> OrgPreferences ShowOktaUIFooter(ctx, )
Show Okta UI Footer

Makes the Okta UI footer visible for all end users of your organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgContactUser**
> OrgContactUser UpdateOrgContactUser(ctx, body, contactType)
Update org contact user

Updates the User associated with the specified Contact Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserIdString**](UserIdString.md)|  | 
  **contactType** | **string**|  | 

### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSetting**
> OrgSetting UpdateOrgSetting(ctx, body)
Update Org setting

Update settings of your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrgSetting**](OrgSetting.md)|  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


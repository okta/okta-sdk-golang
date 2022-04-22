# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCustomization**](BrandApi.md#CreateEmailCustomization) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Create Email Customization
[**DeleteAllCustomizations**](BrandApi.md#DeleteAllCustomizations) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Delete Email Customizations
[**DeleteBrandThemeBackgroundImage**](BrandApi.md#DeleteBrandThemeBackgroundImage) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Deletes a Theme background image
[**DeleteBrandThemeFavicon**](BrandApi.md#DeleteBrandThemeFavicon) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Deletes a Theme favicon. The org then uses the Okta default favicon.
[**DeleteBrandThemeLogo**](BrandApi.md#DeleteBrandThemeLogo) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/logo | Deletes a Theme logo. The org then uses the Okta default logo.
[**DeleteEmailCustomization**](BrandApi.md#DeleteEmailCustomization) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Delete Email Customization
[**GetAllEmailCustomizations**](BrandApi.md#GetAllEmailCustomizations) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | List Email Customizations
[**GetBrand**](BrandApi.md#GetBrand) | **Get** /api/v1/brands/{brandId} | Get Brand
[**GetBrandTheme**](BrandApi.md#GetBrandTheme) | **Get** /api/v1/brands/{brandId}/themes/{themeId} | Get a theme for a brand
[**GetEmailCustomization**](BrandApi.md#GetEmailCustomization) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Get Email Customization
[**GetEmailTemplate**](BrandApi.md#GetEmailTemplate) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName} | Get Email Template
[**GetEmailTemplateDefaultContent**](BrandApi.md#GetEmailTemplateDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content | Get Default Content of Email Template
[**ListBrandThemes**](BrandApi.md#ListBrandThemes) | **Get** /api/v1/brands/{brandId}/themes | Get Brand Themes
[**ListBrands**](BrandApi.md#ListBrands) | **Get** /api/v1/brands | List Brands
[**ListEmailTemplates**](BrandApi.md#ListEmailTemplates) | **Get** /api/v1/brands/{brandId}/templates/email | List Email Templates
[**PreviewEmailCustomizationContent**](BrandApi.md#PreviewEmailCustomizationContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId}/preview | Get Preview Content of Email Customization
[**PreviewEmailDefaultContent**](BrandApi.md#PreviewEmailDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content/preview | Get Preview of Email Template Default Content
[**SendTestEmailFromTemplate**](BrandApi.md#SendTestEmailFromTemplate) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/test | Send Test Email
[**UpdateBrand**](BrandApi.md#UpdateBrand) | **Put** /api/v1/brands/{brandId} | Update Brand
[**UpdateBrandTheme**](BrandApi.md#UpdateBrandTheme) | **Put** /api/v1/brands/{brandId}/themes/{themeId} | Update a theme for a brand
[**UpdateEmailCustomization**](BrandApi.md#UpdateEmailCustomization) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Update Email Customization
[**UploadBrandThemeBackgroundImage**](BrandApi.md#UploadBrandThemeBackgroundImage) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Updates the background image for your Theme
[**UploadBrandThemeFavicon**](BrandApi.md#UploadBrandThemeFavicon) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Updates the favicon for your theme
[**UploadBrandThemeLogo**](BrandApi.md#UploadBrandThemeLogo) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/logo | Update a themes logo

# **CreateEmailCustomization**
> EmailCustomizationResponse CreateEmailCustomization(ctx, brandId, templateName, optional)
Create Email Customization

Create an email customization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 
 **optional** | ***BrandApiCreateEmailCustomizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiCreateEmailCustomizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EmailCustomization**](EmailCustomization.md)|  | 

### Return type

[**EmailCustomizationResponse**](EmailCustomizationResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllCustomizations**
> DeleteAllCustomizations(ctx, brandId, templateName)
Delete Email Customizations

Delete all customizations for an email template. Also known as “Reset to Default”.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandThemeBackgroundImage**
> DeleteBrandThemeBackgroundImage(ctx, brandId, themeId)
Deletes a Theme background image

Deletes a Theme background image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandThemeFavicon**
> DeleteBrandThemeFavicon(ctx, brandId, themeId)
Deletes a Theme favicon. The org then uses the Okta default favicon.

Deletes a Theme favicon. The org then uses the Okta default favicon.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandThemeLogo**
> DeleteBrandThemeLogo(ctx, brandId, themeId)
Deletes a Theme logo. The org then uses the Okta default logo.

Deletes a Theme logo. The org then uses the Okta default logo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailCustomization**
> DeleteEmailCustomization(ctx, brandId, templateName, customizationId)
Delete Email Customization

Delete an email customization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 
  **customizationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEmailCustomizations**
> []EmailCustomizationResponse GetAllEmailCustomizations(ctx, brandId, templateName)
List Email Customizations

List all email customcations for an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 

### Return type

[**[]EmailCustomizationResponse**](EmailCustomizationResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrand**
> Brand GetBrand(ctx, brandId)
Get Brand

Fetches a brand by `brandId`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBrandTheme**
> ThemeResponse GetBrandTheme(ctx, brandId, themeId)
Get a theme for a brand

Fetches a theme for a brand

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCustomization**
> EmailCustomizationResponse GetEmailCustomization(ctx, brandId, templateName, customizationId)
Get Email Customization

Fetch an email customization by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 
  **customizationId** | **string**|  | 

### Return type

[**EmailCustomizationResponse**](EmailCustomizationResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTemplate**
> EmailTemplate GetEmailTemplate(ctx, brandId, templateName)
Get Email Template

Fetch an email template by `templateName`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailTemplateDefaultContent**
> EmailDefaultContent GetEmailTemplateDefaultContent(ctx, brandId, templateName)
Get Default Content of Email Template

Fetch the default content for an email template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 

### Return type

[**EmailDefaultContent**](EmailDefaultContent.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBrandThemes**
> []ThemeResponse ListBrandThemes(ctx, brandId)
Get Brand Themes

List all the themes in your brand

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 

### Return type

[**[]ThemeResponse**](ThemeResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBrands**
> []Brand ListBrands(ctx, )
List Brands

List all the brands in your org.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Brand**](Brand.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEmailTemplates**
> []EmailTemplate ListEmailTemplates(ctx, brandId, optional)
List Email Templates

List email templates in your organization with pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**| brand id | 
 **optional** | ***BrandApiListEmailTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiListEmailTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **optional.String**| Specifies the pagination cursor for the next page of email templates. | 
 **limit** | **optional.Int32**| Specifies the number of results returned (maximum 200) | [default to 20]

### Return type

[**[]EmailTemplate**](EmailTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewEmailCustomizationContent**
> EmailPreview PreviewEmailCustomizationContent(ctx, brandId, templateName, customizationId)
Get Preview Content of Email Customization

Get a preview of an email template customization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 
  **customizationId** | **string**|  | 

### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewEmailDefaultContent**
> EmailPreview PreviewEmailDefaultContent(ctx, brandId, templateName)
Get Preview of Email Template Default Content

Fetch a preview of an email template's default content by populating velocity references with the current user's environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 

### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmailFromTemplate**
> SendTestEmailFromTemplate(ctx, templateName, brandId, optional)
Send Test Email

Send a test email to the current user’s primary and secondary email addresses. The email content is selected based on the following priority: 1. An email customization specifically for the user’s locale. 2. The default language of email customizations. 3. The email template’s default content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateName** | **string**|  | 
  **brandId** | **string**|  | 
 **optional** | ***BrandApiSendTestEmailFromTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiSendTestEmailFromTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EmailTestResponse**](EmailTestResponse.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBrand**
> Brand UpdateBrand(ctx, brandId, optional)
Update Brand

Updates a brand by `brandId`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
 **optional** | ***BrandApiUpdateBrandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiUpdateBrandOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Brand**](Brand.md)|  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBrandTheme**
> ThemeResponse UpdateBrandTheme(ctx, brandId, themeId, optional)
Update a theme for a brand

Updates a theme for a brand

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 
 **optional** | ***BrandApiUpdateBrandThemeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiUpdateBrandThemeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Theme**](Theme.md)|  | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailCustomization**
> EmailCustomizationResponse UpdateEmailCustomization(ctx, brandId, templateName, customizationId, optional)
Update Email Customization

Update an email customization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **templateName** | **string**|  | 
  **customizationId** | **string**|  | 
 **optional** | ***BrandApiUpdateEmailCustomizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandApiUpdateEmailCustomizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of EmailCustomization**](EmailCustomization.md)| Request | 

### Return type

[**EmailCustomizationResponse**](EmailCustomizationResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBrandThemeBackgroundImage**
> ImageUploadResponse UploadBrandThemeBackgroundImage(ctx, brandId, themeId)
Updates the background image for your Theme

Updates the background image for your Theme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBrandThemeFavicon**
> ImageUploadResponse UploadBrandThemeFavicon(ctx, brandId, themeId)
Updates the favicon for your theme

Updates the favicon for your theme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBrandThemeLogo**
> ImageUploadResponse UploadBrandThemeLogo(ctx, brandId, themeId)
Update a themes logo

Updates the logo for your Theme

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **string**|  | 
  **themeId** | **string**|  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[api_token](../README.md#api_token), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


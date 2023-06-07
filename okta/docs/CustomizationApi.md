# \CustomizationApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrand**](CustomizationApi.md#CreateBrand) | **Post** /api/v1/brands | Create a Brand
[**CreateEmailCustomization**](CustomizationApi.md#CreateEmailCustomization) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Create an Email Customization
[**DeleteAllCustomizations**](CustomizationApi.md#DeleteAllCustomizations) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Delete all Email Customizations
[**DeleteBrand**](CustomizationApi.md#DeleteBrand) | **Delete** /api/v1/brands/{brandId} | Delete a brand
[**DeleteBrandThemeBackgroundImage**](CustomizationApi.md#DeleteBrandThemeBackgroundImage) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Delete the Background Image
[**DeleteBrandThemeFavicon**](CustomizationApi.md#DeleteBrandThemeFavicon) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Delete the Favicon
[**DeleteBrandThemeLogo**](CustomizationApi.md#DeleteBrandThemeLogo) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/logo | Delete the Logo
[**DeleteEmailCustomization**](CustomizationApi.md#DeleteEmailCustomization) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Delete an Email Customization
[**GetBrand**](CustomizationApi.md#GetBrand) | **Get** /api/v1/brands/{brandId} | Retrieve a Brand
[**GetBrandTheme**](CustomizationApi.md#GetBrandTheme) | **Get** /api/v1/brands/{brandId}/themes/{themeId} | Retrieve a Theme
[**GetCustomizationPreview**](CustomizationApi.md#GetCustomizationPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId}/preview | Retrieve a Preview of an Email Customization
[**GetCustomizedErrorPage**](CustomizationApi.md#GetCustomizedErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/customized | Retrieve the Customized Error Page
[**GetCustomizedSignInPage**](CustomizationApi.md#GetCustomizedSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/customized | Retrieve the Customized Sign-in Page
[**GetDefaultErrorPage**](CustomizationApi.md#GetDefaultErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/default | Retrieve the Default Error Page
[**GetDefaultSignInPage**](CustomizationApi.md#GetDefaultSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/default | Retrieve the Default Sign-in Page
[**GetEmailCustomization**](CustomizationApi.md#GetEmailCustomization) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Retrieve an Email Customization
[**GetEmailDefaultContent**](CustomizationApi.md#GetEmailDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content | Retrieve an Email Template Default Content
[**GetEmailDefaultPreview**](CustomizationApi.md#GetEmailDefaultPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content/preview | Retrieve a Preview of the Email Template Default Content
[**GetEmailSettings**](CustomizationApi.md#GetEmailSettings) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Retrieve the Email Template Settings
[**GetEmailTemplate**](CustomizationApi.md#GetEmailTemplate) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName} | Retrieve an Email Template
[**GetErrorPage**](CustomizationApi.md#GetErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error | Retrieve the Error Page
[**GetPreviewErrorPage**](CustomizationApi.md#GetPreviewErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/preview | Retrieve the Preview Error Page Preview
[**GetPreviewSignInPage**](CustomizationApi.md#GetPreviewSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/preview | Retrieve the Preview Sign-in Page Preview
[**GetSignInPage**](CustomizationApi.md#GetSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in | Retrieve the Sign-in Page
[**GetSignOutPageSettings**](CustomizationApi.md#GetSignOutPageSettings) | **Get** /api/v1/brands/{brandId}/pages/sign-out/customized | Retrieve the Sign-out Page Settings
[**LinkBrandDomain**](CustomizationApi.md#LinkBrandDomain) | **Post** /api/v1/brands/{brandId}/domains | Link a Brand to a Domain
[**ListAllSignInWidgetVersions**](CustomizationApi.md#ListAllSignInWidgetVersions) | **Get** /api/v1/brands/{brandId}/pages/sign-in/widget-versions | List all Sign-in Widget Versions
[**ListBrandDomains**](CustomizationApi.md#ListBrandDomains) | **Get** /api/v1/brands/{brandId}/domains | List all Domains associated with a Brand
[**ListBrandThemes**](CustomizationApi.md#ListBrandThemes) | **Get** /api/v1/brands/{brandId}/themes | List all Themes
[**ListBrands**](CustomizationApi.md#ListBrands) | **Get** /api/v1/brands | List all Brands
[**ListEmailCustomizations**](CustomizationApi.md#ListEmailCustomizations) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | List all Email Customizations
[**ListEmailTemplates**](CustomizationApi.md#ListEmailTemplates) | **Get** /api/v1/brands/{brandId}/templates/email | List all Email Templates
[**ReplaceBrand**](CustomizationApi.md#ReplaceBrand) | **Put** /api/v1/brands/{brandId} | Replace a Brand
[**ReplaceBrandTheme**](CustomizationApi.md#ReplaceBrandTheme) | **Put** /api/v1/brands/{brandId}/themes/{themeId} | Replace a Theme
[**ReplaceCustomizedErrorPage**](CustomizationApi.md#ReplaceCustomizedErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/customized | Replace the Customized Error Page
[**ReplaceCustomizedSignInPage**](CustomizationApi.md#ReplaceCustomizedSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/customized | Replace the Customized Sign-in Page
[**ReplaceEmailCustomization**](CustomizationApi.md#ReplaceEmailCustomization) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Replace an Email Customization
[**ReplaceEmailSettings**](CustomizationApi.md#ReplaceEmailSettings) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Replace the Email Template Settings
[**ReplacePreviewErrorPage**](CustomizationApi.md#ReplacePreviewErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/preview | Replace the Preview Error Page
[**ReplacePreviewSignInPage**](CustomizationApi.md#ReplacePreviewSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/preview | Replace the Preview Sign-in Page
[**ReplaceSignOutPageSettings**](CustomizationApi.md#ReplaceSignOutPageSettings) | **Put** /api/v1/brands/{brandId}/pages/sign-out/customized | Replace the Sign-out Page Settings
[**ResetCustomizedErrorPage**](CustomizationApi.md#ResetCustomizedErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/customized | Reset the Customized Error Page
[**ResetCustomizedSignInPage**](CustomizationApi.md#ResetCustomizedSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/customized | Reset the Customized Sign-in Page
[**ResetPreviewErrorPage**](CustomizationApi.md#ResetPreviewErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/preview | Reset the Preview Error Page
[**ResetPreviewSignInPage**](CustomizationApi.md#ResetPreviewSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/preview | Reset the Preview Sign-in Page
[**SendTestEmail**](CustomizationApi.md#SendTestEmail) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/test | Send a Test Email
[**UnlinkBrandDomain**](CustomizationApi.md#UnlinkBrandDomain) | **Delete** /api/v1/brands/{brandId}/domains/{domainId} | Unlink a Brand from a Domain
[**UploadBrandThemeBackgroundImage**](CustomizationApi.md#UploadBrandThemeBackgroundImage) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Upload the Background Image
[**UploadBrandThemeFavicon**](CustomizationApi.md#UploadBrandThemeFavicon) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Upload the Favicon
[**UploadBrandThemeLogo**](CustomizationApi.md#UploadBrandThemeLogo) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/logo | Upload the Logo



## CreateBrand

> Brand CreateBrand(ctx).CreateBrandRequest(createBrandRequest).Execute()

Create a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBrandRequest := *openapiclient.NewCreateBrandRequest() // CreateBrandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.CreateBrand(context.Background()).CreateBrandRequest(createBrandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.CreateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.CreateBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBrandRequest** | [**CreateBrandRequest**](CreateBrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailCustomization

> EmailCustomization CreateEmailCustomization(ctx, brandId, templateName).Instance(instance).Execute()

Create an Email Customization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.CreateEmailCustomization(context.Background(), brandId, templateName).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.CreateEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.CreateEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**EmailCustomization**](EmailCustomization.md) |  | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllCustomizations

> DeleteAllCustomizations(ctx, brandId, templateName).Execute()

Delete all Email Customizations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteAllCustomizations(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteAllCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCustomizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrand

> DeleteBrand(ctx, brandId).Execute()

Delete a brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrandThemeBackgroundImage

> DeleteBrandThemeBackgroundImage(ctx, brandId, themeId).Execute()

Delete the Background Image



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeBackgroundImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrandThemeFavicon

> DeleteBrandThemeFavicon(ctx, brandId, themeId).Execute()

Delete the Favicon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeFavicon(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeFaviconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrandThemeLogo

> DeleteBrandThemeLogo(ctx, brandId, themeId).Execute()

Delete the Logo



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteBrandThemeLogo(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandThemeLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailCustomization

> DeleteEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Delete an Email Customization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.DeleteEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.DeleteEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrand

> Brand GetBrand(ctx, brandId).Execute()

Retrieve a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandTheme

> ThemeResponse GetBrandTheme(ctx, brandId, themeId).Execute()

Retrieve a Theme



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetBrandTheme(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomizationPreview

> EmailPreview GetCustomizationPreview(ctx, brandId, templateName, customizationId).Execute()

Retrieve a Preview of an Email Customization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetCustomizationPreview(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetCustomizationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizationPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetCustomizationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomizedErrorPage

> CustomizablePage GetCustomizedErrorPage(ctx, brandId).Execute()

Retrieve the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetCustomizedErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizedErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomizedSignInPage

> SignInPage GetCustomizedSignInPage(ctx, brandId).Execute()

Retrieve the Customized Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetCustomizedSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizedSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultErrorPage

> CustomizablePage GetDefaultErrorPage(ctx, brandId).Execute()

Retrieve the Default Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetDefaultErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetDefaultErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetDefaultErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultSignInPage

> SignInPage GetDefaultSignInPage(ctx, brandId).Execute()

Retrieve the Default Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetDefaultSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetDefaultSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetDefaultSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailCustomization

> EmailCustomization GetEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Retrieve an Email Customization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultContent

> EmailDefaultContent GetEmailDefaultContent(ctx, brandId, templateName).Language(language).Execute()

Retrieve an Email Template Default Content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailDefaultContent(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailDefaultContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultContent`: EmailDefaultContent
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailDefaultContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailDefaultContent**](EmailDefaultContent.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultPreview

> EmailPreview GetEmailDefaultPreview(ctx, brandId, templateName).Language(language).Execute()

Retrieve a Preview of the Email Template Default Content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailDefaultPreview(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailDefaultPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailDefaultPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSettings

> EmailSettings GetEmailSettings(ctx, brandId, templateName).Execute()

Retrieve the Email Template Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailSettings(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSettings`: EmailSettings
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailSettings**](EmailSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplate

> EmailTemplate GetEmailTemplate(ctx, brandId, templateName).Expand(expand).Execute()

Retrieve an Email Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetEmailTemplate(context.Background(), brandId, templateName).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplate`: EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetErrorPage

> PageRoot GetErrorPage(ctx, brandId).Expand(expand).Execute()

Retrieve the Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetErrorPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetErrorPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**PageRoot**](PageRoot.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviewErrorPage

> CustomizablePage GetPreviewErrorPage(ctx, brandId).Execute()

Retrieve the Preview Error Page Preview



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetPreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetPreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetPreviewErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviewSignInPage

> SignInPage GetPreviewSignInPage(ctx, brandId).Execute()

Retrieve the Preview Sign-in Page Preview



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetPreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetPreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetPreviewSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviewSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignInPage

> PageRoot GetSignInPage(ctx, brandId).Expand(expand).Execute()

Retrieve the Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetSignInPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignInPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**PageRoot**](PageRoot.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignOutPageSettings

> HostedPage GetSignOutPageSettings(ctx, brandId).Execute()

Retrieve the Sign-out Page Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.GetSignOutPageSettings(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.GetSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.GetSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignOutPageSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostedPage**](HostedPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkBrandDomain

> BrandDomain LinkBrandDomain(ctx, brandId).CreateBrandDomainRequest(createBrandDomainRequest).Execute()

Link a Brand to a Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    createBrandDomainRequest := *openapiclient.NewCreateBrandDomainRequest() // CreateBrandDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.LinkBrandDomain(context.Background(), brandId).CreateBrandDomainRequest(createBrandDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.LinkBrandDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkBrandDomain`: BrandDomain
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.LinkBrandDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkBrandDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBrandDomainRequest** | [**CreateBrandDomainRequest**](CreateBrandDomainRequest.md) |  | 

### Return type

[**BrandDomain**](BrandDomain.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllSignInWidgetVersions

> []string ListAllSignInWidgetVersions(ctx, brandId).Execute()

List all Sign-in Widget Versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListAllSignInWidgetVersions(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListAllSignInWidgetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllSignInWidgetVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListAllSignInWidgetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllSignInWidgetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandDomains

> []DomainResponse ListBrandDomains(ctx, brandId).Execute()

List all Domains associated with a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListBrandDomains(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListBrandDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandDomains`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListBrandDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandThemes

> []ThemeResponse ListBrandThemes(ctx, brandId).Execute()

List all Themes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListBrandThemes(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListBrandThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandThemes`: []ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListBrandThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandThemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ThemeResponse**](ThemeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrands

> []Brand ListBrands(ctx).Execute()

List all Brands



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListBrands(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrands`: []Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListBrands`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandsRequest struct via the builder pattern


### Return type

[**[]Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailCustomizations

> []EmailCustomization ListEmailCustomizations(ctx, brandId, templateName).After(after).Limit(limit).Execute()

List all Email Customizations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination) for more information. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListEmailCustomizations(context.Background(), brandId, templateName).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListEmailCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailCustomizations`: []EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListEmailCustomizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailCustomizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination) for more information. | 
 **limit** | **int32** | A limit on the number of objects to return. | [default to 20]

### Return type

[**[]EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTemplates

> []EmailTemplate ListEmailTemplates(ctx, brandId).After(after).Limit(limit).Expand(expand).Execute()

List all Email Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination) for more information. (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return. (optional) (default to 20)
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ListEmailTemplates(context.Background(), brandId).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ListEmailTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplates`: []EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination) for more information. | 
 **limit** | **int32** | A limit on the number of objects to return. | [default to 20]
 **expand** | **[]string** | Specifies additional metadata to be included in the response. | 

### Return type

[**[]EmailTemplate**](EmailTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBrand

> Brand ReplaceBrand(ctx, brandId).Brand(brand).Execute()

Replace a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    brand := *openapiclient.NewBrandRequest() // BrandRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceBrand(context.Background(), brandId).Brand(brand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brand** | [**BrandRequest**](BrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBrandTheme

> ThemeResponse ReplaceBrandTheme(ctx, brandId, themeId).Theme(theme).Execute()

Replace a Theme



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.
    theme := *openapiclient.NewTheme() // Theme | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceBrandTheme(context.Background(), brandId, themeId).Theme(theme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBrandThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **theme** | [**Theme**](Theme.md) |  | 

### Return type

[**ThemeResponse**](ThemeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCustomizedErrorPage

> CustomizablePage ReplaceCustomizedErrorPage(ctx, brandId).CustomizablePage(customizablePage).Execute()

Replace the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    customizablePage := *openapiclient.NewCustomizablePage() // CustomizablePage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceCustomizedErrorPage(context.Background(), brandId).CustomizablePage(customizablePage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceCustomizedErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomizedErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customizablePage** | [**CustomizablePage**](CustomizablePage.md) |  | 

### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCustomizedSignInPage

> SignInPage ReplaceCustomizedSignInPage(ctx, brandId).SignInPage(signInPage).Execute()

Replace the Customized Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    signInPage := *openapiclient.NewSignInPage() // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceCustomizedSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceCustomizedSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomizedSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signInPage** | [**SignInPage**](SignInPage.md) |  | 

### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailCustomization

> EmailCustomization ReplaceEmailCustomization(ctx, brandId, templateName, customizationId).Instance(instance).Execute()

Replace an Email Customization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    customizationId := "customizationId_example" // string | The ID of the email customization.
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization | Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceEmailCustomization(context.Background(), brandId, templateName, customizationId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 
**customizationId** | **string** | The ID of the email customization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **instance** | [**EmailCustomization**](EmailCustomization.md) | Request | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailSettings

> ReplaceEmailSettings(ctx, brandId, templateName).EmailSettings(emailSettings).Execute()

Replace the Email Template Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    emailSettings := *openapiclient.NewEmailSettings("Recipients_example") // EmailSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceEmailSettings(context.Background(), brandId, templateName).EmailSettings(emailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailSettings** | [**EmailSettings**](EmailSettings.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePreviewErrorPage

> CustomizablePage ReplacePreviewErrorPage(ctx, brandId).CustomizablePage(customizablePage).Execute()

Replace the Preview Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    customizablePage := *openapiclient.NewCustomizablePage() // CustomizablePage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplacePreviewErrorPage(context.Background(), brandId).CustomizablePage(customizablePage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplacePreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewErrorPage`: CustomizablePage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplacePreviewErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customizablePage** | [**CustomizablePage**](CustomizablePage.md) |  | 

### Return type

[**CustomizablePage**](CustomizablePage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePreviewSignInPage

> SignInPage ReplacePreviewSignInPage(ctx, brandId).SignInPage(signInPage).Execute()

Replace the Preview Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    signInPage := *openapiclient.NewSignInPage() // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplacePreviewSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplacePreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplacePreviewSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePreviewSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signInPage** | [**SignInPage**](SignInPage.md) |  | 

### Return type

[**SignInPage**](SignInPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSignOutPageSettings

> HostedPage ReplaceSignOutPageSettings(ctx, brandId).HostedPage(hostedPage).Execute()

Replace the Sign-out Page Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    hostedPage := *openapiclient.NewHostedPage(openapiclient.HostedPageType("EXTERNALLY_HOSTED")) // HostedPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ReplaceSignOutPageSettings(context.Background(), brandId).HostedPage(hostedPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ReplaceSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.ReplaceSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSignOutPageSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostedPage** | [**HostedPage**](HostedPage.md) |  | 

### Return type

[**HostedPage**](HostedPage.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCustomizedErrorPage

> ResetCustomizedErrorPage(ctx, brandId).Execute()

Reset the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCustomizedErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetCustomizedSignInPage

> ResetCustomizedSignInPage(ctx, brandId).Execute()

Reset the Customized Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetCustomizedSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPreviewErrorPage

> ResetPreviewErrorPage(ctx, brandId).Execute()

Reset the Preview Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetPreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetPreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPreviewSignInPage

> ResetPreviewSignInPage(ctx, brandId).Execute()

Reset the Preview Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.ResetPreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.ResetPreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPreviewSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestEmail

> SendTestEmail(ctx, brandId, templateName).Language(language).Execute()

Send a Test Email



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    templateName := "templateName_example" // string | The name of the email template.
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.SendTestEmail(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.SendTestEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**templateName** | **string** | The name of the email template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkBrandDomain

> UnlinkBrandDomain(ctx, brandId, domainId).Execute()

Unlink a Brand from a Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    domainId := "domainId_example" // string | The ID of the domain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UnlinkBrandDomain(context.Background(), brandId, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UnlinkBrandDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**domainId** | **string** | The ID of the domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkBrandDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeBackgroundImage

> ImageUploadResponse UploadBrandThemeBackgroundImage(ctx, brandId, themeId).File(file).Execute()

Upload the Background Image



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeBackgroundImage(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeBackgroundImage`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeBackgroundImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeBackgroundImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeFavicon

> ImageUploadResponse UploadBrandThemeFavicon(ctx, brandId, themeId).File(file).Execute()

Upload the Favicon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeFavicon(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeFavicon`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeFavicon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeFaviconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBrandThemeLogo

> ImageUploadResponse UploadBrandThemeLogo(ctx, brandId, themeId).File(file).Execute()

Upload the Logo



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand.
    themeId := "themeId_example" // string | The ID of the theme.
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationApi.UploadBrandThemeLogo(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationApi.UploadBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeLogo`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationApi.UploadBrandThemeLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand. | 
**themeId** | **string** | The ID of the theme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBrandThemeLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**ImageUploadResponse**](ImageUploadResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


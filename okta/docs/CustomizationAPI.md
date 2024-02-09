# \CustomizationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrand**](CustomizationAPI.md#CreateBrand) | **Post** /api/v1/brands | Create a Brand
[**CreateEmailCustomization**](CustomizationAPI.md#CreateEmailCustomization) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Create an Email Customization
[**DeleteAllCustomizations**](CustomizationAPI.md#DeleteAllCustomizations) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Delete all Email Customizations
[**DeleteBrand**](CustomizationAPI.md#DeleteBrand) | **Delete** /api/v1/brands/{brandId} | Delete a brand
[**DeleteBrandThemeBackgroundImage**](CustomizationAPI.md#DeleteBrandThemeBackgroundImage) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Delete the Background Image
[**DeleteBrandThemeFavicon**](CustomizationAPI.md#DeleteBrandThemeFavicon) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Delete the Favicon
[**DeleteBrandThemeLogo**](CustomizationAPI.md#DeleteBrandThemeLogo) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/logo | Delete the Logo
[**DeleteCustomizedErrorPage**](CustomizationAPI.md#DeleteCustomizedErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/customized | Delete the Customized Error Page
[**DeleteCustomizedSignInPage**](CustomizationAPI.md#DeleteCustomizedSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/customized | Delete the Customized Sign-in Page
[**DeleteEmailCustomization**](CustomizationAPI.md#DeleteEmailCustomization) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Delete an Email Customization
[**DeletePreviewErrorPage**](CustomizationAPI.md#DeletePreviewErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/preview | Delete the Preview Error Page
[**DeletePreviewSignInPage**](CustomizationAPI.md#DeletePreviewSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/preview | Delete the Preview Sign-in Page
[**GetBrand**](CustomizationAPI.md#GetBrand) | **Get** /api/v1/brands/{brandId} | Retrieve a Brand
[**GetBrandTheme**](CustomizationAPI.md#GetBrandTheme) | **Get** /api/v1/brands/{brandId}/themes/{themeId} | Retrieve a Theme
[**GetCustomizationPreview**](CustomizationAPI.md#GetCustomizationPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId}/preview | Retrieve a Preview of an Email Customization
[**GetCustomizedErrorPage**](CustomizationAPI.md#GetCustomizedErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/customized | Retrieve the Customized Error Page
[**GetCustomizedSignInPage**](CustomizationAPI.md#GetCustomizedSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/customized | Retrieve the Customized Sign-in Page
[**GetDefaultErrorPage**](CustomizationAPI.md#GetDefaultErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/default | Retrieve the Default Error Page
[**GetDefaultSignInPage**](CustomizationAPI.md#GetDefaultSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/default | Retrieve the Default Sign-in Page
[**GetEmailCustomization**](CustomizationAPI.md#GetEmailCustomization) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Retrieve an Email Customization
[**GetEmailDefaultContent**](CustomizationAPI.md#GetEmailDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content | Retrieve an Email Template Default Content
[**GetEmailDefaultPreview**](CustomizationAPI.md#GetEmailDefaultPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content/preview | Retrieve a Preview of the Email Template default content
[**GetEmailSettings**](CustomizationAPI.md#GetEmailSettings) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Retrieve the Email Template Settings
[**GetEmailTemplate**](CustomizationAPI.md#GetEmailTemplate) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName} | Retrieve an Email Template
[**GetErrorPage**](CustomizationAPI.md#GetErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error | Retrieve the Error Page Sub-Resources
[**GetPreviewErrorPage**](CustomizationAPI.md#GetPreviewErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/preview | Retrieve the Preview Error Page Preview
[**GetPreviewSignInPage**](CustomizationAPI.md#GetPreviewSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/preview | Retrieve the Preview Sign-in Page Preview
[**GetSignInPage**](CustomizationAPI.md#GetSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in | Retrieve the Sign-in Page Sub-Resources
[**GetSignOutPageSettings**](CustomizationAPI.md#GetSignOutPageSettings) | **Get** /api/v1/brands/{brandId}/pages/sign-out/customized | Retrieve the Sign-out Page Settings
[**ListAllSignInWidgetVersions**](CustomizationAPI.md#ListAllSignInWidgetVersions) | **Get** /api/v1/brands/{brandId}/pages/sign-in/widget-versions | List all Sign-in Widget Versions
[**ListBrandDomains**](CustomizationAPI.md#ListBrandDomains) | **Get** /api/v1/brands/{brandId}/domains | List all Domains associated with a Brand
[**ListBrandThemes**](CustomizationAPI.md#ListBrandThemes) | **Get** /api/v1/brands/{brandId}/themes | List all Themes
[**ListBrands**](CustomizationAPI.md#ListBrands) | **Get** /api/v1/brands | List all Brands
[**ListEmailCustomizations**](CustomizationAPI.md#ListEmailCustomizations) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | List all Email Customizations
[**ListEmailTemplates**](CustomizationAPI.md#ListEmailTemplates) | **Get** /api/v1/brands/{brandId}/templates/email | List all Email Templates
[**ReplaceBrand**](CustomizationAPI.md#ReplaceBrand) | **Put** /api/v1/brands/{brandId} | Replace a Brand
[**ReplaceBrandTheme**](CustomizationAPI.md#ReplaceBrandTheme) | **Put** /api/v1/brands/{brandId}/themes/{themeId} | Replace a Theme
[**ReplaceCustomizedErrorPage**](CustomizationAPI.md#ReplaceCustomizedErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/customized | Replace the Customized Error Page
[**ReplaceCustomizedSignInPage**](CustomizationAPI.md#ReplaceCustomizedSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/customized | Replace the Customized Sign-in Page
[**ReplaceEmailCustomization**](CustomizationAPI.md#ReplaceEmailCustomization) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Replace an Email Customization
[**ReplaceEmailSettings**](CustomizationAPI.md#ReplaceEmailSettings) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Replace the Email Template Settings
[**ReplacePreviewErrorPage**](CustomizationAPI.md#ReplacePreviewErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/preview | Replace the Preview Error Page
[**ReplacePreviewSignInPage**](CustomizationAPI.md#ReplacePreviewSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/preview | Replace the Preview Sign-in Page
[**ReplaceSignOutPageSettings**](CustomizationAPI.md#ReplaceSignOutPageSettings) | **Put** /api/v1/brands/{brandId}/pages/sign-out/customized | Replace the Sign-out Page Settings
[**SendTestEmail**](CustomizationAPI.md#SendTestEmail) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/test | Send a Test Email
[**UploadBrandThemeBackgroundImage**](CustomizationAPI.md#UploadBrandThemeBackgroundImage) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Upload the Background Image
[**UploadBrandThemeFavicon**](CustomizationAPI.md#UploadBrandThemeFavicon) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Upload the Favicon
[**UploadBrandThemeLogo**](CustomizationAPI.md#UploadBrandThemeLogo) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/logo | Upload the Logo



## CreateBrand

> Brand CreateBrand(ctx).Expand(expand).After(after).Limit(limit).Q(q).CreateBrandRequest(createBrandRequest).Execute()

Create a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
    q := "q_example" // string | Searches the records for matching value (optional)
    createBrandRequest := *openapiclient.NewCreateBrandRequest("Name_example") // CreateBrandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.CreateBrand(context.Background()).Expand(expand).After(after).Limit(limit).Q(q).CreateBrandRequest(createBrandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.CreateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.CreateBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **q** | **string** | Searches the records for matching value | 
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.CreateEmailCustomization(context.Background(), brandId, templateName).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.CreateEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.CreateEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteAllCustomizations(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteAllCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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

> DeleteBrand(ctx, brandId).Expand(expand).Execute()

Delete a brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteBrand(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteBrandThemeFavicon(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteBrandThemeLogo(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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


## DeleteCustomizedErrorPage

> DeleteCustomizedErrorPage(ctx, brandId).Execute()

Delete the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomizedErrorPageRequest struct via the builder pattern


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


## DeleteCustomizedSignInPage

> DeleteCustomizedSignInPage(ctx, brandId).Execute()

Delete the Customized Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomizedSignInPageRequest struct via the builder pattern


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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeleteEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeleteEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

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


## DeletePreviewErrorPage

> DeletePreviewErrorPage(ctx, brandId).Execute()

Delete the Preview Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeletePreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeletePreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePreviewErrorPageRequest struct via the builder pattern


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


## DeletePreviewSignInPage

> DeletePreviewSignInPage(ctx, brandId).Execute()

Delete the Preview Sign-in Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.DeletePreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.DeletePreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePreviewSignInPageRequest struct via the builder pattern


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

> BrandWithEmbedded GetBrand(ctx, brandId).Expand(expand).Execute()

Retrieve a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetBrand(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: BrandWithEmbedded
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**BrandWithEmbedded**](BrandWithEmbedded.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetBrandTheme(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetCustomizationPreview(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetCustomizationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizationPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetCustomizationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

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

> ErrorPage GetCustomizedErrorPage(ctx, brandId).Execute()

Retrieve the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetCustomizedErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizedErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorPage**](ErrorPage.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetCustomizedSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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

> ErrorPage GetDefaultErrorPage(ctx, brandId).Execute()

Retrieve the Default Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetDefaultErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetDefaultErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetDefaultErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorPage**](ErrorPage.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetDefaultSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetDefaultSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetDefaultSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetEmailDefaultContent(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetEmailDefaultContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultContent`: EmailDefaultContent
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetEmailDefaultContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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

Retrieve a Preview of the Email Template default content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetEmailDefaultPreview(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetEmailDefaultPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetEmailDefaultPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetEmailSettings(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSettings`: EmailSettings
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetEmailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetEmailTemplate(context.Background(), brandId, templateName).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplate`: EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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

Retrieve the Error Page Sub-Resources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetErrorPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetErrorPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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

> ErrorPage GetPreviewErrorPage(ctx, brandId).Execute()

Retrieve the Preview Error Page Preview



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetPreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetPreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetPreviewErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorPage**](ErrorPage.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetPreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetPreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetPreviewSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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

Retrieve the Sign-in Page Sub-Resources



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetSignInPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignInPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSignInPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.GetSignOutPageSettings(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.GetSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.GetSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListAllSignInWidgetVersions(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListAllSignInWidgetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllSignInWidgetVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListAllSignInWidgetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListBrandDomains(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListBrandDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandDomains`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListBrandDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListBrandThemes(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListBrandThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandThemes`: []ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListBrandThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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

> []BrandWithEmbedded ListBrands(ctx).Expand(expand).After(after).Limit(limit).Q(q).Execute()

List all Brands



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
    q := "q_example" // string | Searches the records for matching value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListBrands(context.Background()).Expand(expand).After(after).Limit(limit).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrands`: []BrandWithEmbedded
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **q** | **string** | Searches the records for matching value | 

### Return type

[**[]BrandWithEmbedded**](BrandWithEmbedded.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListEmailCustomizations(context.Background(), brandId, templateName).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListEmailCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailCustomizations`: []EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListEmailCustomizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailCustomizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ListEmailTemplates(context.Background(), brandId).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ListEmailTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplates`: []EmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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

> Brand ReplaceBrand(ctx, brandId).Brand(brand).Expand(expand).Execute()

Replace a Brand



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    brand := *openapiclient.NewBrandRequest() // BrandRequest | 
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceBrand(context.Background(), brandId).Brand(brand).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brand** | [**BrandRequest**](BrandRequest.md) |  | 
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme
    theme := *openapiclient.NewTheme() // Theme | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceBrandTheme(context.Background(), brandId, themeId).Theme(theme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceBrandTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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

> ErrorPage ReplaceCustomizedErrorPage(ctx, brandId).ErrorPage(errorPage).Execute()

Replace the Customized Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    errorPage := *openapiclient.NewErrorPage() // ErrorPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceCustomizedErrorPage(context.Background(), brandId).ErrorPage(errorPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceCustomizedErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomizedErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorPage** | [**ErrorPage**](ErrorPage.md) |  | 

### Return type

[**ErrorPage**](ErrorPage.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    signInPage := *openapiclient.NewSignInPage() // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceCustomizedSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceCustomizedSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization | Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceEmailCustomization(context.Background(), brandId, templateName, customizationId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    emailSettings := *openapiclient.NewEmailSettings("Recipients_example") // EmailSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.ReplaceEmailSettings(context.Background(), brandId, templateName).EmailSettings(emailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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

> ErrorPage ReplacePreviewErrorPage(ctx, brandId).ErrorPage(errorPage).Execute()

Replace the Preview Error Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    errorPage := *openapiclient.NewErrorPage() // ErrorPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplacePreviewErrorPage(context.Background(), brandId).ErrorPage(errorPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplacePreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplacePreviewErrorPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePreviewErrorPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errorPage** | [**ErrorPage**](ErrorPage.md) |  | 

### Return type

[**ErrorPage**](ErrorPage.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    signInPage := *openapiclient.NewSignInPage() // SignInPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplacePreviewSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplacePreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplacePreviewSignInPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    hostedPage := *openapiclient.NewHostedPage("Type_example") // HostedPage | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.ReplaceSignOutPageSettings(context.Background(), brandId).HostedPage(hostedPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.ReplaceSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.ReplaceSignOutPageSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomizationAPI.SendTestEmail(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.SendTestEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.UploadBrandThemeBackgroundImage(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.UploadBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeBackgroundImage`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.UploadBrandThemeBackgroundImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.UploadBrandThemeFavicon(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.UploadBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeFavicon`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.UploadBrandThemeFavicon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    brandId := "brandId_example" // string | The ID of the brand
    themeId := "themeId_example" // string | The ID of the theme
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomizationAPI.UploadBrandThemeLogo(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomizationAPI.UploadBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeLogo`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomizationAPI.UploadBrandThemeLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**themeId** | **string** | The ID of the theme | 

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


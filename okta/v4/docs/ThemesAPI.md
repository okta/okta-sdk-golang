# \ThemesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBrandThemeBackgroundImage**](ThemesAPI.md#DeleteBrandThemeBackgroundImage) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Delete the Background Image
[**DeleteBrandThemeFavicon**](ThemesAPI.md#DeleteBrandThemeFavicon) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Delete the Favicon
[**DeleteBrandThemeLogo**](ThemesAPI.md#DeleteBrandThemeLogo) | **Delete** /api/v1/brands/{brandId}/themes/{themeId}/logo | Delete the Logo
[**GetBrandTheme**](ThemesAPI.md#GetBrandTheme) | **Get** /api/v1/brands/{brandId}/themes/{themeId} | Retrieve a Theme
[**ListBrandThemes**](ThemesAPI.md#ListBrandThemes) | **Get** /api/v1/brands/{brandId}/themes | List all Themes
[**ReplaceBrandTheme**](ThemesAPI.md#ReplaceBrandTheme) | **Put** /api/v1/brands/{brandId}/themes/{themeId} | Replace a Theme
[**UploadBrandThemeBackgroundImage**](ThemesAPI.md#UploadBrandThemeBackgroundImage) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/background-image | Upload the Background Image
[**UploadBrandThemeFavicon**](ThemesAPI.md#UploadBrandThemeFavicon) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/favicon | Upload the Favicon
[**UploadBrandThemeLogo**](ThemesAPI.md#UploadBrandThemeLogo) | **Post** /api/v1/brands/{brandId}/themes/{themeId}/logo | Upload the Logo



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
    r, err := apiClient.ThemesAPI.DeleteBrandThemeBackgroundImage(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.DeleteBrandThemeBackgroundImage``: %v\n", err)
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
    r, err := apiClient.ThemesAPI.DeleteBrandThemeFavicon(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.DeleteBrandThemeFavicon``: %v\n", err)
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
    r, err := apiClient.ThemesAPI.DeleteBrandThemeLogo(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.DeleteBrandThemeLogo``: %v\n", err)
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
    resp, r, err := apiClient.ThemesAPI.GetBrandTheme(context.Background(), brandId, themeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.GetBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.GetBrandTheme`: %v\n", resp)
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
    resp, r, err := apiClient.ThemesAPI.ListBrandThemes(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.ListBrandThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandThemes`: []ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.ListBrandThemes`: %v\n", resp)
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
    theme := *openapiclient.NewUpdateThemeRequest("EmailTemplateTouchPointVariant_example", "EndUserDashboardTouchPointVariant_example", "ErrorPageTouchPointVariant_example", "PrimaryColorHex_example", "SecondaryColorHex_example", "SignInPageTouchPointVariant_example") // UpdateThemeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemesAPI.ReplaceBrandTheme(context.Background(), brandId, themeId).Theme(theme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.ReplaceBrandTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrandTheme`: ThemeResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.ReplaceBrandTheme`: %v\n", resp)
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


 **theme** | [**UpdateThemeRequest**](UpdateThemeRequest.md) |  | 

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
    resp, r, err := apiClient.ThemesAPI.UploadBrandThemeBackgroundImage(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.UploadBrandThemeBackgroundImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeBackgroundImage`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.UploadBrandThemeBackgroundImage`: %v\n", resp)
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
    resp, r, err := apiClient.ThemesAPI.UploadBrandThemeFavicon(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.UploadBrandThemeFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeFavicon`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.UploadBrandThemeFavicon`: %v\n", resp)
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
    resp, r, err := apiClient.ThemesAPI.UploadBrandThemeLogo(context.Background(), brandId, themeId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.UploadBrandThemeLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBrandThemeLogo`: ImageUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.UploadBrandThemeLogo`: %v\n", resp)
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


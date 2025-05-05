# \CustomPagesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomizedErrorPage**](CustomPagesAPI.md#DeleteCustomizedErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/customized | Delete the Customized Error Page
[**DeleteCustomizedSignInPage**](CustomPagesAPI.md#DeleteCustomizedSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/customized | Delete the Customized Sign-in Page
[**DeletePreviewErrorPage**](CustomPagesAPI.md#DeletePreviewErrorPage) | **Delete** /api/v1/brands/{brandId}/pages/error/preview | Delete the Preview Error Page
[**DeletePreviewSignInPage**](CustomPagesAPI.md#DeletePreviewSignInPage) | **Delete** /api/v1/brands/{brandId}/pages/sign-in/preview | Delete the Preview Sign-in Page
[**GetCustomizedErrorPage**](CustomPagesAPI.md#GetCustomizedErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/customized | Retrieve the Customized Error Page
[**GetCustomizedSignInPage**](CustomPagesAPI.md#GetCustomizedSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/customized | Retrieve the Customized Sign-in Page
[**GetDefaultErrorPage**](CustomPagesAPI.md#GetDefaultErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/default | Retrieve the Default Error Page
[**GetDefaultSignInPage**](CustomPagesAPI.md#GetDefaultSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/default | Retrieve the Default Sign-in Page
[**GetErrorPage**](CustomPagesAPI.md#GetErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error | Retrieve the Error Page Sub-Resources
[**GetPreviewErrorPage**](CustomPagesAPI.md#GetPreviewErrorPage) | **Get** /api/v1/brands/{brandId}/pages/error/preview | Retrieve the Preview Error Page Preview
[**GetPreviewSignInPage**](CustomPagesAPI.md#GetPreviewSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in/preview | Retrieve the Preview Sign-in Page Preview
[**GetSignInPage**](CustomPagesAPI.md#GetSignInPage) | **Get** /api/v1/brands/{brandId}/pages/sign-in | Retrieve the Sign-in Page Sub-Resources
[**GetSignOutPageSettings**](CustomPagesAPI.md#GetSignOutPageSettings) | **Get** /api/v1/brands/{brandId}/pages/sign-out/customized | Retrieve the Sign-out Page Settings
[**ListAllSignInWidgetVersions**](CustomPagesAPI.md#ListAllSignInWidgetVersions) | **Get** /api/v1/brands/{brandId}/pages/sign-in/widget-versions | List all Sign-in Widget Versions
[**ReplaceCustomizedErrorPage**](CustomPagesAPI.md#ReplaceCustomizedErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/customized | Replace the Customized Error Page
[**ReplaceCustomizedSignInPage**](CustomPagesAPI.md#ReplaceCustomizedSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/customized | Replace the Customized Sign-in Page
[**ReplacePreviewErrorPage**](CustomPagesAPI.md#ReplacePreviewErrorPage) | **Put** /api/v1/brands/{brandId}/pages/error/preview | Replace the Preview Error Page
[**ReplacePreviewSignInPage**](CustomPagesAPI.md#ReplacePreviewSignInPage) | **Put** /api/v1/brands/{brandId}/pages/sign-in/preview | Replace the Preview Sign-in Page
[**ReplaceSignOutPageSettings**](CustomPagesAPI.md#ReplaceSignOutPageSettings) | **Put** /api/v1/brands/{brandId}/pages/sign-out/customized | Replace the Sign-out Page Settings



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
    r, err := apiClient.CustomPagesAPI.DeleteCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.DeleteCustomizedErrorPage``: %v\n", err)
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
    r, err := apiClient.CustomPagesAPI.DeleteCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.DeleteCustomizedSignInPage``: %v\n", err)
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
    r, err := apiClient.CustomPagesAPI.DeletePreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.DeletePreviewErrorPage``: %v\n", err)
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
    r, err := apiClient.CustomPagesAPI.DeletePreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.DeletePreviewSignInPage``: %v\n", err)
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
    resp, r, err := apiClient.CustomPagesAPI.GetCustomizedErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetCustomizedErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetCustomizedSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetCustomizedSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetDefaultErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetDefaultErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetDefaultErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetDefaultSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetDefaultSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetDefaultSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetErrorPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetErrorPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetPreviewErrorPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetPreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetPreviewErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetPreviewSignInPage(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetPreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetPreviewSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetSignInPage(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignInPage`: PageRoot
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.GetSignOutPageSettings(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.GetSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.GetSignOutPageSettings`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ListAllSignInWidgetVersions(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ListAllSignInWidgetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllSignInWidgetVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ListAllSignInWidgetVersions`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ReplaceCustomizedErrorPage(context.Background(), brandId).ErrorPage(errorPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ReplaceCustomizedErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ReplaceCustomizedErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ReplaceCustomizedSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ReplaceCustomizedSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCustomizedSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ReplaceCustomizedSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ReplacePreviewErrorPage(context.Background(), brandId).ErrorPage(errorPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ReplacePreviewErrorPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewErrorPage`: ErrorPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ReplacePreviewErrorPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ReplacePreviewSignInPage(context.Background(), brandId).SignInPage(signInPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ReplacePreviewSignInPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePreviewSignInPage`: SignInPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ReplacePreviewSignInPage`: %v\n", resp)
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
    resp, r, err := apiClient.CustomPagesAPI.ReplaceSignOutPageSettings(context.Background(), brandId).HostedPage(hostedPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesAPI.ReplaceSignOutPageSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSignOutPageSettings`: HostedPage
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesAPI.ReplaceSignOutPageSettings`: %v\n", resp)
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


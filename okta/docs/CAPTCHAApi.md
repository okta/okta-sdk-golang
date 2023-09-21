# \CAPTCHAApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCaptchaInstance**](CAPTCHAApi.md#CreateCaptchaInstance) | **Post** /api/v1/captchas | Create a CAPTCHA instance
[**DeleteCaptchaInstance**](CAPTCHAApi.md#DeleteCaptchaInstance) | **Delete** /api/v1/captchas/{captchaId} | Delete a CAPTCHA Instance
[**DeleteOrgCaptchaSettings**](CAPTCHAApi.md#DeleteOrgCaptchaSettings) | **Delete** /api/v1/org/captcha | Delete the Org-wide CAPTCHA Settings
[**GetCaptchaInstance**](CAPTCHAApi.md#GetCaptchaInstance) | **Get** /api/v1/captchas/{captchaId} | Retrieve a CAPTCHA Instance
[**GetOrgCaptchaSettings**](CAPTCHAApi.md#GetOrgCaptchaSettings) | **Get** /api/v1/org/captcha | Retrieve the Org-wide CAPTCHA Settings
[**ListCaptchaInstances**](CAPTCHAApi.md#ListCaptchaInstances) | **Get** /api/v1/captchas | List all CAPTCHA Instances
[**ReplaceCaptchaInstance**](CAPTCHAApi.md#ReplaceCaptchaInstance) | **Put** /api/v1/captchas/{captchaId} | Replace a CAPTCHA Instance
[**ReplacesOrgCaptchaSettings**](CAPTCHAApi.md#ReplacesOrgCaptchaSettings) | **Put** /api/v1/org/captcha | Replace the Org-wide CAPTCHA Settings
[**UpdateCaptchaInstance**](CAPTCHAApi.md#UpdateCaptchaInstance) | **Post** /api/v1/captchas/{captchaId} | Update a CAPTCHA Instance



## CreateCaptchaInstance

> CAPTCHAInstance CreateCaptchaInstance(ctx).Instance(instance).Execute()

Create a CAPTCHA instance



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
    instance := *openapiclient.NewCAPTCHAInstance() // CAPTCHAInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.CreateCaptchaInstance(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.CreateCaptchaInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCaptchaInstance`: CAPTCHAInstance
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.CreateCaptchaInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCaptchaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**CAPTCHAInstance**](CAPTCHAInstance.md) |  | 

### Return type

[**CAPTCHAInstance**](CAPTCHAInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptchaInstance

> DeleteCaptchaInstance(ctx, captchaId).Execute()

Delete a CAPTCHA Instance



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
    captchaId := "captchaId_example" // string | The unique key used to identify your CAPTCHA instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CAPTCHAApi.DeleteCaptchaInstance(context.Background(), captchaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.DeleteCaptchaInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaId** | **string** | The unique key used to identify your CAPTCHA instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptchaInstanceRequest struct via the builder pattern


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


## DeleteOrgCaptchaSettings

> DeleteOrgCaptchaSettings(ctx).Execute()

Delete the Org-wide CAPTCHA Settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CAPTCHAApi.DeleteOrgCaptchaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.DeleteOrgCaptchaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgCaptchaSettingsRequest struct via the builder pattern


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


## GetCaptchaInstance

> CAPTCHAInstance GetCaptchaInstance(ctx, captchaId).Execute()

Retrieve a CAPTCHA Instance



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
    captchaId := "captchaId_example" // string | The unique key used to identify your CAPTCHA instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.GetCaptchaInstance(context.Background(), captchaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.GetCaptchaInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCaptchaInstance`: CAPTCHAInstance
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.GetCaptchaInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaId** | **string** | The unique key used to identify your CAPTCHA instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CAPTCHAInstance**](CAPTCHAInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgCaptchaSettings

> OrgCAPTCHASettings GetOrgCaptchaSettings(ctx).Execute()

Retrieve the Org-wide CAPTCHA Settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.GetOrgCaptchaSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.GetOrgCaptchaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgCaptchaSettings`: OrgCAPTCHASettings
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.GetOrgCaptchaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCaptchaSettingsRequest struct via the builder pattern


### Return type

[**OrgCAPTCHASettings**](OrgCAPTCHASettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCaptchaInstances

> []CAPTCHAInstance ListCaptchaInstances(ctx).Execute()

List all CAPTCHA Instances



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.ListCaptchaInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.ListCaptchaInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCaptchaInstances`: []CAPTCHAInstance
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.ListCaptchaInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCaptchaInstancesRequest struct via the builder pattern


### Return type

[**[]CAPTCHAInstance**](CAPTCHAInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCaptchaInstance

> CAPTCHAInstance ReplaceCaptchaInstance(ctx, captchaId).Instance(instance).Execute()

Replace a CAPTCHA Instance



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
    captchaId := "captchaId_example" // string | The unique key used to identify your CAPTCHA instance
    instance := *openapiclient.NewCAPTCHAInstance() // CAPTCHAInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.ReplaceCaptchaInstance(context.Background(), captchaId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.ReplaceCaptchaInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCaptchaInstance`: CAPTCHAInstance
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.ReplaceCaptchaInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaId** | **string** | The unique key used to identify your CAPTCHA instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCaptchaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**CAPTCHAInstance**](CAPTCHAInstance.md) |  | 

### Return type

[**CAPTCHAInstance**](CAPTCHAInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacesOrgCaptchaSettings

> OrgCAPTCHASettings ReplacesOrgCaptchaSettings(ctx).OrgCAPTCHASettings(orgCAPTCHASettings).Execute()

Replace the Org-wide CAPTCHA Settings



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
    orgCAPTCHASettings := *openapiclient.NewOrgCAPTCHASettings() // OrgCAPTCHASettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.ReplacesOrgCaptchaSettings(context.Background()).OrgCAPTCHASettings(orgCAPTCHASettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.ReplacesOrgCaptchaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacesOrgCaptchaSettings`: OrgCAPTCHASettings
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.ReplacesOrgCaptchaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplacesOrgCaptchaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgCAPTCHASettings** | [**OrgCAPTCHASettings**](OrgCAPTCHASettings.md) |  | 

### Return type

[**OrgCAPTCHASettings**](OrgCAPTCHASettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCaptchaInstance

> CAPTCHAInstance UpdateCaptchaInstance(ctx, captchaId).Instance(instance).Execute()

Update a CAPTCHA Instance



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
    captchaId := "captchaId_example" // string | The unique key used to identify your CAPTCHA instance
    instance := *openapiclient.NewCAPTCHAInstance() // CAPTCHAInstance | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CAPTCHAApi.UpdateCaptchaInstance(context.Background(), captchaId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CAPTCHAApi.UpdateCaptchaInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCaptchaInstance`: CAPTCHAInstance
    fmt.Fprintf(os.Stdout, "Response from `CAPTCHAApi.UpdateCaptchaInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captchaId** | **string** | The unique key used to identify your CAPTCHA instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCaptchaInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instance** | [**CAPTCHAInstance**](CAPTCHAInstance.md) |  | 

### Return type

[**CAPTCHAInstance**](CAPTCHAInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


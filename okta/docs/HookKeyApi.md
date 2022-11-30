# \HookKeyApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHookKey**](HookKeyApi.md#CreateHookKey) | **Post** /api/v1/hook-keys | Create a key
[**DeleteHookKey**](HookKeyApi.md#DeleteHookKey) | **Delete** /api/v1/hook-keys/{hookKeyId} | Delete a key
[**GetHookKey**](HookKeyApi.md#GetHookKey) | **Get** /api/v1/hook-keys/{hookKeyId} | Retrieve a key
[**GetPublicKey**](HookKeyApi.md#GetPublicKey) | **Get** /api/v1/hook-keys/public/{keyId} | Retrieve a public key
[**ListHookKeys**](HookKeyApi.md#ListHookKeys) | **Get** /api/v1/hook-keys | List all keys
[**ReplaceHookKey**](HookKeyApi.md#ReplaceHookKey) | **Put** /api/v1/hook-keys/{hookKeyId} | Replace a key



## CreateHookKey

> HookKey CreateHookKey(ctx).KeyRequest(keyRequest).Execute()

Create a key



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
    keyRequest := *openapiclient.NewKeyRequest() // KeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyApi.CreateHookKey(context.Background()).KeyRequest(keyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.CreateHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyApi.CreateHookKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyRequest** | [**KeyRequest**](KeyRequest.md) |  | 

### Return type

[**HookKey**](HookKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHookKey

> DeleteHookKey(ctx, hookKeyId).Execute()

Delete a key



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
    hookKeyId := "hookKeyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyApi.DeleteHookKey(context.Background(), hookKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.DeleteHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHookKeyRequest struct via the builder pattern


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


## GetHookKey

> HookKey GetHookKey(ctx, hookKeyId).Execute()

Retrieve a key



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
    hookKeyId := "hookKeyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyApi.GetHookKey(context.Background(), hookKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.GetHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyApi.GetHookKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHookKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HookKey**](HookKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicKey

> JsonWebKey GetPublicKey(ctx, keyId).Execute()

Retrieve a public key



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
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyApi.GetPublicKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.GetPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyApi.GetPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHookKeys

> []HookKey ListHookKeys(ctx).Execute()

List all keys



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
    resp, r, err := apiClient.HookKeyApi.ListHookKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.ListHookKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHookKeys`: []HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyApi.ListHookKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHookKeysRequest struct via the builder pattern


### Return type

[**[]HookKey**](HookKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceHookKey

> HookKey ReplaceHookKey(ctx, hookKeyId).KeyRequest(keyRequest).Execute()

Replace a key



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
    hookKeyId := "hookKeyId_example" // string | 
    keyRequest := *openapiclient.NewKeyRequest() // KeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyApi.ReplaceHookKey(context.Background(), hookKeyId).KeyRequest(keyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyApi.ReplaceHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyApi.ReplaceHookKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceHookKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyRequest** | [**KeyRequest**](KeyRequest.md) |  | 

### Return type

[**HookKey**](HookKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


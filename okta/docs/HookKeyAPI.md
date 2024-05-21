# \HookKeyAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHookKey**](HookKeyAPI.md#CreateHookKey) | **Post** /api/v1/hook-keys | Create a key
[**DeleteHookKey**](HookKeyAPI.md#DeleteHookKey) | **Delete** /api/v1/hook-keys/{hookKeyId} | Delete a key
[**GetHookKey**](HookKeyAPI.md#GetHookKey) | **Get** /api/v1/hook-keys/{hookKeyId} | Retrieve a key
[**GetPublicKey**](HookKeyAPI.md#GetPublicKey) | **Get** /api/v1/hook-keys/public/{publicKeyId} | Retrieve a public key
[**ListHookKeys**](HookKeyAPI.md#ListHookKeys) | **Get** /api/v1/hook-keys | List all keys
[**ReplaceHookKey**](HookKeyAPI.md#ReplaceHookKey) | **Put** /api/v1/hook-keys/{hookKeyId} | Replace a key



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    keyRequest := *openapiclient.NewKeyRequest() // KeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyAPI.CreateHookKey(context.Background()).KeyRequest(keyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.CreateHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyAPI.CreateHookKey`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    hookKeyId := "XreKU5laGwBkjOTehusG" // string | `id` of the Hook Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HookKeyAPI.DeleteHookKey(context.Background(), hookKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.DeleteHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** | &#x60;id&#x60; of the Hook Key | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    hookKeyId := "XreKU5laGwBkjOTehusG" // string | `id` of the Hook Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyAPI.GetHookKey(context.Background(), hookKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.GetHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyAPI.GetHookKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** | &#x60;id&#x60; of the Hook Key | 

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

> JsonWebKey GetPublicKey(ctx, publicKeyId).Execute()

Retrieve a public key



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
    publicKeyId := "FcH2P9Eg7wr0o8N2FuV0" // string | `id` of the Public Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyAPI.GetPublicKey(context.Background(), publicKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.GetPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyAPI.GetPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKeyId** | **string** | &#x60;id&#x60; of the Public Key | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyAPI.ListHookKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.ListHookKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHookKeys`: []HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyAPI.ListHookKeys`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    hookKeyId := "XreKU5laGwBkjOTehusG" // string | `id` of the Hook Key
    keyRequest := *openapiclient.NewKeyRequest() // KeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HookKeyAPI.ReplaceHookKey(context.Background(), hookKeyId).KeyRequest(keyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HookKeyAPI.ReplaceHookKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceHookKey`: HookKey
    fmt.Fprintf(os.Stdout, "Response from `HookKeyAPI.ReplaceHookKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKeyId** | **string** | &#x60;id&#x60; of the Hook Key | 

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


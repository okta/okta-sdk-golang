# \ApiTokenAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiToken**](ApiTokenAPI.md#GetApiToken) | **Get** /api/v1/api-tokens/{apiTokenId} | Retrieve an API Token&#39;s Metadata
[**ListApiTokens**](ApiTokenAPI.md#ListApiTokens) | **Get** /api/v1/api-tokens | List all API Token Metadata
[**RevokeApiToken**](ApiTokenAPI.md#RevokeApiToken) | **Delete** /api/v1/api-tokens/{apiTokenId} | Revoke an API Token
[**RevokeCurrentApiToken**](ApiTokenAPI.md#RevokeCurrentApiToken) | **Delete** /api/v1/api-tokens/current | Revoke the Current API Token
[**UpsertApiToken**](ApiTokenAPI.md#UpsertApiToken) | **Put** /api/v1/api-tokens/{apiTokenId} | Upsert an API Token Network Condition



## GetApiToken

> ApiToken GetApiToken(ctx, apiTokenId).Execute()

Retrieve an API Token's Metadata



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
    apiTokenId := "00Tabcdefg1234567890" // string | id of the API Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiTokenAPI.GetApiToken(context.Background(), apiTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenAPI.GetApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokenAPI.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string** | id of the API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiTokens

> []ApiToken ListApiTokens(ctx).Execute()

List all API Token Metadata



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
    resp, r, err := apiClient.ApiTokenAPI.ListApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenAPI.ListApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiTokens`: []ApiToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokenAPI.ListApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiTokensRequest struct via the builder pattern


### Return type

[**[]ApiToken**](ApiToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeApiToken

> RevokeApiToken(ctx, apiTokenId).Execute()

Revoke an API Token



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
    apiTokenId := "00Tabcdefg1234567890" // string | id of the API Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApiTokenAPI.RevokeApiToken(context.Background(), apiTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenAPI.RevokeApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string** | id of the API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeApiTokenRequest struct via the builder pattern


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


## RevokeCurrentApiToken

> RevokeCurrentApiToken(ctx).Execute()

Revoke the Current API Token



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
    r, err := apiClient.ApiTokenAPI.RevokeCurrentApiToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenAPI.RevokeCurrentApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCurrentApiTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertApiToken

> ApiToken UpsertApiToken(ctx, apiTokenId).ApiTokenUpdate(apiTokenUpdate).Execute()

Upsert an API Token Network Condition



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
    apiTokenId := "00Tabcdefg1234567890" // string | id of the API Token
    apiTokenUpdate := *openapiclient.NewApiTokenUpdate() // ApiTokenUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiTokenAPI.UpsertApiToken(context.Background(), apiTokenId).ApiTokenUpdate(apiTokenUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiTokenAPI.UpsertApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertApiToken`: ApiToken
    fmt.Fprintf(os.Stdout, "Response from `ApiTokenAPI.UpsertApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiTokenId** | **string** | id of the API Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTokenUpdate** | [**ApiTokenUpdate**](ApiTokenUpdate.md) |  | 

### Return type

[**ApiToken**](ApiToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


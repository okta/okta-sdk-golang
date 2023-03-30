# \PrincipalRateLimitApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrincipalRateLimitEntity**](PrincipalRateLimitApi.md#CreatePrincipalRateLimitEntity) | **Post** /api/v1/principal-rate-limits | Create a Principal Rate Limit
[**GetPrincipalRateLimitEntity**](PrincipalRateLimitApi.md#GetPrincipalRateLimitEntity) | **Get** /api/v1/principal-rate-limits/{principalRateLimitId} | Retrieve a Principal Rate Limit
[**ListPrincipalRateLimitEntities**](PrincipalRateLimitApi.md#ListPrincipalRateLimitEntities) | **Get** /api/v1/principal-rate-limits | List all Principal Rate Limits
[**ReplacePrincipalRateLimitEntity**](PrincipalRateLimitApi.md#ReplacePrincipalRateLimitEntity) | **Put** /api/v1/principal-rate-limits/{principalRateLimitId} | Replace a Principal Rate Limit



## CreatePrincipalRateLimitEntity

> PrincipalRateLimitEntity CreatePrincipalRateLimitEntity(ctx).Entity(entity).Execute()

Create a Principal Rate Limit



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
    entity := *openapiclient.NewPrincipalRateLimitEntity("PrincipalId_example", openapiclient.PrincipalType("SSWS_TOKEN")) // PrincipalRateLimitEntity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrincipalRateLimitApi.CreatePrincipalRateLimitEntity(context.Background()).Entity(entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitApi.CreatePrincipalRateLimitEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrincipalRateLimitEntity`: PrincipalRateLimitEntity
    fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitApi.CreatePrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md) |  | 

### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalRateLimitEntity

> PrincipalRateLimitEntity GetPrincipalRateLimitEntity(ctx, principalRateLimitId).Execute()

Retrieve a Principal Rate Limit



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
    principalRateLimitId := "abcd1234" // string | id of the Principal Rate Limit

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrincipalRateLimitApi.GetPrincipalRateLimitEntity(context.Background(), principalRateLimitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitApi.GetPrincipalRateLimitEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrincipalRateLimitEntity`: PrincipalRateLimitEntity
    fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitApi.GetPrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRateLimitId** | **string** | id of the Principal Rate Limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrincipalRateLimitEntities

> []PrincipalRateLimitEntity ListPrincipalRateLimitEntities(ctx).Filter(filter).After(after).Limit(limit).Execute()

List all Principal Rate Limits



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
    filter := "filter_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrincipalRateLimitApi.ListPrincipalRateLimitEntities(context.Background()).Filter(filter).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitApi.ListPrincipalRateLimitEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrincipalRateLimitEntities`: []PrincipalRateLimitEntity
    fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitApi.ListPrincipalRateLimitEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalRateLimitEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePrincipalRateLimitEntity

> PrincipalRateLimitEntity ReplacePrincipalRateLimitEntity(ctx, principalRateLimitId).Entity(entity).Execute()

Replace a Principal Rate Limit



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
    principalRateLimitId := "abcd1234" // string | id of the Principal Rate Limit
    entity := *openapiclient.NewPrincipalRateLimitEntity("PrincipalId_example", openapiclient.PrincipalType("SSWS_TOKEN")) // PrincipalRateLimitEntity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrincipalRateLimitApi.ReplacePrincipalRateLimitEntity(context.Background(), principalRateLimitId).Entity(entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitApi.ReplacePrincipalRateLimitEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePrincipalRateLimitEntity`: PrincipalRateLimitEntity
    fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitApi.ReplacePrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRateLimitId** | **string** | id of the Principal Rate Limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entity** | [**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md) |  | 

### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


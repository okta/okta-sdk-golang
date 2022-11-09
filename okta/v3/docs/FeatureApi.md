# \FeatureApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeature**](FeatureApi.md#GetFeature) | **Get** /api/v1/features/{featureId} | Retrieve a Feature
[**ListFeatureDependencies**](FeatureApi.md#ListFeatureDependencies) | **Get** /api/v1/features/{featureId}/dependencies | List all Dependencies
[**ListFeatureDependents**](FeatureApi.md#ListFeatureDependents) | **Get** /api/v1/features/{featureId}/dependents | List all Dependents
[**ListFeatures**](FeatureApi.md#ListFeatures) | **Get** /api/v1/features | List all Features
[**UpdateFeatureLifecycle**](FeatureApi.md#UpdateFeatureLifecycle) | **Post** /api/v1/features/{featureId}/{lifecycle} | Update a Feature Lifecycle



## GetFeature

> Feature GetFeature(ctx, featureId).Execute()

Retrieve a Feature



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
    featureId := "featureId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureApi.GetFeature(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureApi.GetFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeature`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureApi.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Feature**](Feature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureDependencies

> []Feature ListFeatureDependencies(ctx, featureId).Execute()

List all Dependencies



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
    featureId := "featureId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureApi.ListFeatureDependencies(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureApi.ListFeatureDependencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatureDependencies`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureApi.ListFeatureDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Feature**](Feature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureDependents

> []Feature ListFeatureDependents(ctx, featureId).Execute()

List all Dependents



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
    featureId := "featureId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureApi.ListFeatureDependents(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureApi.ListFeatureDependents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatureDependents`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureApi.ListFeatureDependents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureDependentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Feature**](Feature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> []Feature ListFeatures(ctx).Execute()

List all Features



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
    resp, r, err := apiClient.FeatureApi.ListFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureApi.ListFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatures`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureApi.ListFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


### Return type

[**[]Feature**](Feature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureLifecycle

> Feature UpdateFeatureLifecycle(ctx, featureId, lifecycle).Mode(mode).Execute()

Update a Feature Lifecycle



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
    featureId := "featureId_example" // string | 
    lifecycle := "lifecycle_example" // string | 
    mode := "mode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureApi.UpdateFeatureLifecycle(context.Background(), featureId, lifecycle).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureApi.UpdateFeatureLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureLifecycle`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureApi.UpdateFeatureLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 
**lifecycle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **string** |  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


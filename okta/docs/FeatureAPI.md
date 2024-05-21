# \FeatureAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeature**](FeatureAPI.md#GetFeature) | **Get** /api/v1/features/{featureId} | Retrieve a Feature
[**ListFeatureDependencies**](FeatureAPI.md#ListFeatureDependencies) | **Get** /api/v1/features/{featureId}/dependencies | List all dependencies
[**ListFeatureDependents**](FeatureAPI.md#ListFeatureDependents) | **Get** /api/v1/features/{featureId}/dependents | List all dependents
[**ListFeatures**](FeatureAPI.md#ListFeatures) | **Get** /api/v1/features | List all Features
[**UpdateFeatureLifecycle**](FeatureAPI.md#UpdateFeatureLifecycle) | **Post** /api/v1/features/{featureId}/{lifecycle} | Update a Feature lifecycle



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    featureId := "R5HjqNn1pEqWGy48E9jg" // string | `id` of the feature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureAPI.GetFeature(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureAPI.GetFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeature`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureAPI.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | &#x60;id&#x60; of the feature | 

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

List all dependencies



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
    featureId := "R5HjqNn1pEqWGy48E9jg" // string | `id` of the feature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureAPI.ListFeatureDependencies(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureAPI.ListFeatureDependencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatureDependencies`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureAPI.ListFeatureDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | &#x60;id&#x60; of the feature | 

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

List all dependents



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
    featureId := "R5HjqNn1pEqWGy48E9jg" // string | `id` of the feature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureAPI.ListFeatureDependents(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureAPI.ListFeatureDependents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatureDependents`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureAPI.ListFeatureDependents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | &#x60;id&#x60; of the feature | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureAPI.ListFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureAPI.ListFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatures`: []Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureAPI.ListFeatures`: %v\n", resp)
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

Update a Feature lifecycle



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
    featureId := "R5HjqNn1pEqWGy48E9jg" // string | `id` of the feature
    lifecycle := "lifecycle_example" // string | Whether to `ENABLE` or `DISABLE` the feature
    mode := "mode_example" // string | Indicates if you want to force enable or disable a feature. Supported value is `force`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureAPI.UpdateFeatureLifecycle(context.Background(), featureId, lifecycle).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureAPI.UpdateFeatureLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureLifecycle`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureAPI.UpdateFeatureLifecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | &#x60;id&#x60; of the feature | 
**lifecycle** | **string** | Whether to &#x60;ENABLE&#x60; or &#x60;DISABLE&#x60; the feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **string** | Indicates if you want to force enable or disable a feature. Supported value is &#x60;force&#x60;. | 

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


# \ApplicationFeaturesApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeatureForApplication**](ApplicationFeaturesApi.md#GetFeatureForApplication) | **Get** /api/v1/apps/{appId}/features/{featureName} | Retrieve a Feature
[**ListFeaturesForApplication**](ApplicationFeaturesApi.md#ListFeaturesForApplication) | **Get** /api/v1/apps/{appId}/features | List all Features
[**UpdateFeatureForApplication**](ApplicationFeaturesApi.md#UpdateFeatureForApplication) | **Put** /api/v1/apps/{appId}/features/{featureName} | Update a Feature



## GetFeatureForApplication

> ApplicationFeature GetFeatureForApplication(ctx, appId, featureName).Execute()

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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    featureName := "USER_PROVISIONING" // string | Name of the Feature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFeaturesApi.GetFeatureForApplication(context.Background(), appId, featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFeaturesApi.GetFeatureForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureForApplication`: ApplicationFeature
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFeaturesApi.GetFeatureForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**featureName** | **string** | Name of the Feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationFeature**](ApplicationFeature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeaturesForApplication

> []ApplicationFeature ListFeaturesForApplication(ctx, appId).Execute()

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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFeaturesApi.ListFeaturesForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFeaturesApi.ListFeaturesForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeaturesForApplication`: []ApplicationFeature
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFeaturesApi.ListFeaturesForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApplicationFeature**](ApplicationFeature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureForApplication

> ApplicationFeature UpdateFeatureForApplication(ctx, appId, featureName).CapabilitiesObject(capabilitiesObject).Execute()

Update a Feature



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    featureName := "USER_PROVISIONING" // string | Name of the Feature
    capabilitiesObject := *openapiclient.NewCapabilitiesObject() // CapabilitiesObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFeaturesApi.UpdateFeatureForApplication(context.Background(), appId, featureName).CapabilitiesObject(capabilitiesObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFeaturesApi.UpdateFeatureForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureForApplication`: ApplicationFeature
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFeaturesApi.UpdateFeatureForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**featureName** | **string** | Name of the Feature | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **capabilitiesObject** | [**CapabilitiesObject**](CapabilitiesObject.md) |  | 

### Return type

[**ApplicationFeature**](ApplicationFeature.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ProfileMappingApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileMapping**](ProfileMappingApi.md#GetProfileMapping) | **Get** /api/v1/mappings/{mappingId} | Retrieve a Profile Mapping
[**ListProfileMappings**](ProfileMappingApi.md#ListProfileMappings) | **Get** /api/v1/mappings | List all Profile Mappings
[**UpdateProfileMapping**](ProfileMappingApi.md#UpdateProfileMapping) | **Post** /api/v1/mappings/{mappingId} | Update a Profile Mapping



## GetProfileMapping

> ProfileMapping GetProfileMapping(ctx, mappingId).Execute()

Retrieve a Profile Mapping



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
    mappingId := "mappingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingApi.GetProfileMapping(context.Background(), mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingApi.GetProfileMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileMapping`: ProfileMapping
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingApi.GetProfileMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProfileMapping**](ProfileMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileMappings

> []ProfileMapping ListProfileMappings(ctx).After(after).Limit(limit).SourceId(sourceId).TargetId(targetId).Execute()

List all Profile Mappings



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
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)
    sourceId := "sourceId_example" // string |  (optional)
    targetId := "targetId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingApi.ListProfileMappings(context.Background()).After(after).Limit(limit).SourceId(sourceId).TargetId(targetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingApi.ListProfileMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProfileMappings`: []ProfileMapping
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingApi.ListProfileMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfileMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]
 **sourceId** | **string** |  | 
 **targetId** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]ProfileMapping**](ProfileMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfileMapping

> ProfileMapping UpdateProfileMapping(ctx, mappingId).ProfileMapping(profileMapping).Execute()

Update a Profile Mapping



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
    mappingId := "mappingId_example" // string | 
    profileMapping := *openapiclient.NewProfileMapping() // ProfileMapping | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingApi.UpdateProfileMapping(context.Background(), mappingId).ProfileMapping(profileMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingApi.UpdateProfileMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfileMapping`: ProfileMapping
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingApi.UpdateProfileMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileMapping** | [**ProfileMapping**](ProfileMapping.md) |  | 

### Return type

[**ProfileMapping**](ProfileMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


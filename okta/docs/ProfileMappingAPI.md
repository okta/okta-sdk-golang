# \ProfileMappingAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileMapping**](ProfileMappingAPI.md#GetProfileMapping) | **Get** /api/v1/mappings/{mappingId} | Retrieve a Profile Mapping
[**ListProfileMappings**](ProfileMappingAPI.md#ListProfileMappings) | **Get** /api/v1/mappings | List all Profile Mappings
[**UpdateProfileMapping**](ProfileMappingAPI.md#UpdateProfileMapping) | **Post** /api/v1/mappings/{mappingId} | Update a Profile Mapping



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    mappingId := "cB6u7X8mptebWkffatKA" // string | `id` of the Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingAPI.GetProfileMapping(context.Background(), mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingAPI.GetProfileMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileMapping`: ProfileMapping
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingAPI.GetProfileMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **string** | &#x60;id&#x60; of the Mapping | 

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

> []ListProfileMappings ListProfileMappings(ctx).After(after).Limit(limit).SourceId(sourceId).TargetId(targetId).Execute()

List all Profile Mappings



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
    after := "after_example" // string | Mapping `id` that specifies the pagination cursor for the next page of mappings (optional)
    limit := int32(56) // int32 | Specifies the number of results per page (maximum 200) (optional) (default to 20)
    sourceId := "sourceId_example" // string | The UserType or App Instance `id` that acts as the source of expressions in a mapping. If this parameter is included, all returned mappings have this as their `source.id`. (optional)
    targetId := "targetId_example" // string | The UserType or App Instance `id` that acts as the target of expressions in a mapping. If this parameter is included, all returned mappings have this as their `target.id`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingAPI.ListProfileMappings(context.Background()).After(after).Limit(limit).SourceId(sourceId).TargetId(targetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingAPI.ListProfileMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProfileMappings`: []ListProfileMappings
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingAPI.ListProfileMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfileMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Mapping &#x60;id&#x60; that specifies the pagination cursor for the next page of mappings | 
 **limit** | **int32** | Specifies the number of results per page (maximum 200) | [default to 20]
 **sourceId** | **string** | The UserType or App Instance &#x60;id&#x60; that acts as the source of expressions in a mapping. If this parameter is included, all returned mappings have this as their &#x60;source.id&#x60;. | 
 **targetId** | **string** | The UserType or App Instance &#x60;id&#x60; that acts as the target of expressions in a mapping. If this parameter is included, all returned mappings have this as their &#x60;target.id&#x60;. | 

### Return type

[**[]ListProfileMappings**](ListProfileMappings.md)

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    mappingId := "cB6u7X8mptebWkffatKA" // string | `id` of the Mapping
    profileMapping := *openapiclient.NewProfileMappingRequest(map[string]ProfileMappingProperty{"key": *openapiclient.NewProfileMappingProperty()}) // ProfileMappingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileMappingAPI.UpdateProfileMapping(context.Background(), mappingId).ProfileMapping(profileMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileMappingAPI.UpdateProfileMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfileMapping`: ProfileMapping
    fmt.Fprintf(os.Stdout, "Response from `ProfileMappingAPI.UpdateProfileMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingId** | **string** | &#x60;id&#x60; of the Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileMapping** | [**ProfileMappingRequest**](ProfileMappingRequest.md) |  | 

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


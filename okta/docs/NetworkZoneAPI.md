# \NetworkZoneAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateNetworkZone**](NetworkZoneAPI.md#ActivateNetworkZone) | **Post** /api/v1/zones/{zoneId}/lifecycle/activate | Activate a Network Zone
[**CreateNetworkZone**](NetworkZoneAPI.md#CreateNetworkZone) | **Post** /api/v1/zones | Create a Network Zone
[**DeactivateNetworkZone**](NetworkZoneAPI.md#DeactivateNetworkZone) | **Post** /api/v1/zones/{zoneId}/lifecycle/deactivate | Deactivate a Network Zone
[**DeleteNetworkZone**](NetworkZoneAPI.md#DeleteNetworkZone) | **Delete** /api/v1/zones/{zoneId} | Delete a Network Zone
[**GetNetworkZone**](NetworkZoneAPI.md#GetNetworkZone) | **Get** /api/v1/zones/{zoneId} | Retrieve a Network Zone
[**ListNetworkZones**](NetworkZoneAPI.md#ListNetworkZones) | **Get** /api/v1/zones | List all Network Zones
[**ReplaceNetworkZone**](NetworkZoneAPI.md#ReplaceNetworkZone) | **Put** /api/v1/zones/{zoneId} | Replace a Network Zone



## ActivateNetworkZone

> ListNetworkZones200ResponseInner ActivateNetworkZone(ctx, zoneId).Execute()

Activate a Network Zone



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
    zoneId := "nzowc1U5Jh5xuAK0o0g3" // string | `id` of the Network Zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.ActivateNetworkZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.ActivateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateNetworkZone`: ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.ActivateNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | &#x60;id&#x60; of the Network Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkZone

> ListNetworkZones200ResponseInner CreateNetworkZone(ctx).Zone(zone).Execute()

Create a Network Zone



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
    zone := openapiclient.listNetworkZones_200_response_inner{DynamicNetworkZone: openapiclient.NewDynamicNetworkZone("Name_example", "Type_example")} // ListNetworkZones200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.CreateNetworkZone(context.Background()).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.CreateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkZone`: ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.CreateNetworkZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | [**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md) |  | 

### Return type

[**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateNetworkZone

> ListNetworkZones200ResponseInner DeactivateNetworkZone(ctx, zoneId).Execute()

Deactivate a Network Zone



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
    zoneId := "nzowc1U5Jh5xuAK0o0g3" // string | `id` of the Network Zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.DeactivateNetworkZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.DeactivateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateNetworkZone`: ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.DeactivateNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | &#x60;id&#x60; of the Network Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkZone

> DeleteNetworkZone(ctx, zoneId).Execute()

Delete a Network Zone



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
    zoneId := "nzowc1U5Jh5xuAK0o0g3" // string | `id` of the Network Zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkZoneAPI.DeleteNetworkZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.DeleteNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | &#x60;id&#x60; of the Network Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkZoneRequest struct via the builder pattern


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


## GetNetworkZone

> ListNetworkZones200ResponseInner GetNetworkZone(ctx, zoneId).Execute()

Retrieve a Network Zone



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
    zoneId := "nzowc1U5Jh5xuAK0o0g3" // string | `id` of the Network Zone

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.GetNetworkZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.GetNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkZone`: ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.GetNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | &#x60;id&#x60; of the Network Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkZones

> []ListNetworkZones200ResponseInner ListNetworkZones(ctx).After(after).Limit(limit).Filter(filter).Execute()

List all Network Zones



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
    after := "BlockedIpZones" // string |  (optional)
    limit := int32(5) // int32 |  (optional) (default to -1)
    filter := "id eq "nzowc1U5Jh5xuAK0o0g3"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.ListNetworkZones(context.Background()).After(after).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.ListNetworkZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkZones`: []ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.ListNetworkZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]
 **filter** | **string** |  | 

### Return type

[**[]ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNetworkZone

> ListNetworkZones200ResponseInner ReplaceNetworkZone(ctx, zoneId).Zone(zone).Execute()

Replace a Network Zone



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
    zoneId := "nzowc1U5Jh5xuAK0o0g3" // string | `id` of the Network Zone
    zone := openapiclient.listNetworkZones_200_response_inner{DynamicNetworkZone: openapiclient.NewDynamicNetworkZone("Name_example", "Type_example")} // ListNetworkZones200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.ReplaceNetworkZone(context.Background(), zoneId).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.ReplaceNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNetworkZone`: ListNetworkZones200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.ReplaceNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | &#x60;id&#x60; of the Network Zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | [**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md) |  | 

### Return type

[**ListNetworkZones200ResponseInner**](ListNetworkZones200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


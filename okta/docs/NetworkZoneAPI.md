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

> NetworkZone ActivateNetworkZone(ctx, zoneId).Execute()

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
    // response from `ActivateNetworkZone`: NetworkZone
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

[**NetworkZone**](NetworkZone.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkZone

> NetworkZone CreateNetworkZone(ctx).Zone(zone).Execute()

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
    zone := *openapiclient.NewNetworkZone() // NetworkZone | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.CreateNetworkZone(context.Background()).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.CreateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkZone`: NetworkZone
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.CreateNetworkZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | [**NetworkZone**](NetworkZone.md) |  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateNetworkZone

> NetworkZone DeactivateNetworkZone(ctx, zoneId).Execute()

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
    // response from `DeactivateNetworkZone`: NetworkZone
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

[**NetworkZone**](NetworkZone.md)

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

> NetworkZone GetNetworkZone(ctx, zoneId).Execute()

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
    // response from `GetNetworkZone`: NetworkZone
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

[**NetworkZone**](NetworkZone.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkZones

> []NetworkZone ListNetworkZones(ctx).After(after).Limit(limit).Filter(filter).Execute()

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
    after := "200u7yq5goxNFTiMjW1d7" // string | Specifies the pagination cursor for the next page of network zones (optional)
    limit := int32(5) // int32 | Specifies the number of results for a page (optional) (default to -1)
    filter := "filter=%28id+eq+%22nzowc1U5Jh5xuAK0o0g3%22%29" // string | Filters zones by usage or ID expression (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.ListNetworkZones(context.Background()).After(after).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.ListNetworkZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkZones`: []NetworkZone
    fmt.Fprintf(os.Stdout, "Response from `NetworkZoneAPI.ListNetworkZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Specifies the pagination cursor for the next page of network zones | 
 **limit** | **int32** | Specifies the number of results for a page | [default to -1]
 **filter** | **string** | Filters zones by usage or ID expression | 

### Return type

[**[]NetworkZone**](NetworkZone.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNetworkZone

> NetworkZone ReplaceNetworkZone(ctx, zoneId).Zone(zone).Execute()

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
    zone := *openapiclient.NewNetworkZone() // NetworkZone | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZoneAPI.ReplaceNetworkZone(context.Background(), zoneId).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZoneAPI.ReplaceNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNetworkZone`: NetworkZone
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

 **zone** | [**NetworkZone**](NetworkZone.md) |  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


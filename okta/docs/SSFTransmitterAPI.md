# \SSFTransmitterAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSsfStream**](SSFTransmitterAPI.md#CreateSsfStream) | **Post** /api/v1/ssf/stream | Create an SSF Stream
[**DeleteSsfStream**](SSFTransmitterAPI.md#DeleteSsfStream) | **Delete** /api/v1/ssf/stream | Delete an SSF Stream
[**GetSsfStreams**](SSFTransmitterAPI.md#GetSsfStreams) | **Get** /api/v1/ssf/stream | Retrieve the SSF Stream configuration(s)
[**GetWellknownSsfMetadata**](SSFTransmitterAPI.md#GetWellknownSsfMetadata) | **Get** /.well-known/ssf-configuration | Retrieve the SSF Transmitter metadata
[**ReplaceSsfStream**](SSFTransmitterAPI.md#ReplaceSsfStream) | **Put** /api/v1/ssf/stream | Replace an SSF Stream
[**UpdateSsfStream**](SSFTransmitterAPI.md#UpdateSsfStream) | **Patch** /api/v1/ssf/stream | Update an SSF Stream



## CreateSsfStream

> StreamConfiguration CreateSsfStream(ctx).Instance(instance).Execute()

Create an SSF Stream



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
    instance := *openapiclient.NewStreamConfigurationCreateRequest(*openapiclient.NewStreamConfigurationDelivery("https://example.com/", "Method_example"), []string{"EventsRequested_example"}) // StreamConfigurationCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFTransmitterAPI.CreateSsfStream(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.CreateSsfStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSsfStream`: StreamConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SSFTransmitterAPI.CreateSsfStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSsfStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**StreamConfigurationCreateRequest**](StreamConfigurationCreateRequest.md) |  | 

### Return type

[**StreamConfiguration**](StreamConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSsfStream

> DeleteSsfStream(ctx).StreamId(streamId).Execute()

Delete an SSF Stream



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
    streamId := "esc1k235GIIztAuGK0g5" // string | The ID of the specified SSF Stream configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SSFTransmitterAPI.DeleteSsfStream(context.Background()).StreamId(streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.DeleteSsfStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSsfStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **string** | The ID of the specified SSF Stream configuration | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsfStreams

> GetSsfStreams200Response GetSsfStreams(ctx).StreamId(streamId).Execute()

Retrieve the SSF Stream configuration(s)



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
    streamId := "esc1k235GIIztAuGK0g5" // string | The ID of the specified SSF Stream configuration (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFTransmitterAPI.GetSsfStreams(context.Background()).StreamId(streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.GetSsfStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSsfStreams`: GetSsfStreams200Response
    fmt.Fprintf(os.Stdout, "Response from `SSFTransmitterAPI.GetSsfStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSsfStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamId** | **string** | The ID of the specified SSF Stream configuration | 

### Return type

[**GetSsfStreams200Response**](GetSsfStreams200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellknownSsfMetadata

> WellKnownSSFMetadata GetWellknownSsfMetadata(ctx).Execute()

Retrieve the SSF Transmitter metadata



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
    resp, r, err := apiClient.SSFTransmitterAPI.GetWellknownSsfMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.GetWellknownSsfMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWellknownSsfMetadata`: WellKnownSSFMetadata
    fmt.Fprintf(os.Stdout, "Response from `SSFTransmitterAPI.GetWellknownSsfMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellknownSsfMetadataRequest struct via the builder pattern


### Return type

[**WellKnownSSFMetadata**](WellKnownSSFMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSsfStream

> StreamConfiguration ReplaceSsfStream(ctx).Instance(instance).Execute()

Replace an SSF Stream



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
    instance := *openapiclient.NewStreamConfiguration(*openapiclient.NewStreamConfigurationDelivery("https://example.com/", "Method_example"), []string{"EventsRequested_example"}) // StreamConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFTransmitterAPI.ReplaceSsfStream(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.ReplaceSsfStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSsfStream`: StreamConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SSFTransmitterAPI.ReplaceSsfStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSsfStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**StreamConfiguration**](StreamConfiguration.md) |  | 

### Return type

[**StreamConfiguration**](StreamConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSsfStream

> StreamConfiguration UpdateSsfStream(ctx).Instance(instance).Execute()

Update an SSF Stream



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
    instance := *openapiclient.NewStreamConfiguration(*openapiclient.NewStreamConfigurationDelivery("https://example.com/", "Method_example"), []string{"EventsRequested_example"}) // StreamConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSFTransmitterAPI.UpdateSsfStream(context.Background()).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSFTransmitterAPI.UpdateSsfStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSsfStream`: StreamConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SSFTransmitterAPI.UpdateSsfStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSsfStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**StreamConfiguration**](StreamConfiguration.md) |  | 

### Return type

[**StreamConfiguration**](StreamConfiguration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


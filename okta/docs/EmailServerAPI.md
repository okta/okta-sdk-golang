# \EmailServerAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailServer**](EmailServerAPI.md#CreateEmailServer) | **Post** /api/v1/email-servers | Create a custom SMTP server
[**DeleteEmailServer**](EmailServerAPI.md#DeleteEmailServer) | **Delete** /api/v1/email-servers/{emailServerId} | Delete an SMTP Server configuration
[**GetEmailServer**](EmailServerAPI.md#GetEmailServer) | **Get** /api/v1/email-servers/{emailServerId} | Retrieve an SMTP Server configuration
[**ListEmailServers**](EmailServerAPI.md#ListEmailServers) | **Get** /api/v1/email-servers | List all enrolled SMTP servers
[**TestEmailServer**](EmailServerAPI.md#TestEmailServer) | **Post** /api/v1/email-servers/{emailServerId}/test | Test an SMTP Server configuration
[**UpdateEmailServer**](EmailServerAPI.md#UpdateEmailServer) | **Patch** /api/v1/email-servers/{emailServerId} | Update an SMTP Server configuration



## CreateEmailServer

> EmailServerResponse CreateEmailServer(ctx).EmailServerPost(emailServerPost).Execute()

Create a custom SMTP server



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
    emailServerPost := *openapiclient.NewEmailServerPost("CustomServer1", "192.168.160.1", int32(587), "aUser", "Password_example") // EmailServerPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailServerAPI.CreateEmailServer(context.Background()).EmailServerPost(emailServerPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.CreateEmailServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailServer`: EmailServerResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailServerAPI.CreateEmailServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailServerPost** | [**EmailServerPost**](EmailServerPost.md) |  | 

### Return type

[**EmailServerResponse**](EmailServerResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailServer

> DeleteEmailServer(ctx, emailServerId).Execute()

Delete an SMTP Server configuration



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
    emailServerId := "emailServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EmailServerAPI.DeleteEmailServer(context.Background(), emailServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.DeleteEmailServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailServerRequest struct via the builder pattern


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


## GetEmailServer

> EmailServerListResponse GetEmailServer(ctx, emailServerId).Execute()

Retrieve an SMTP Server configuration



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
    emailServerId := "emailServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailServerAPI.GetEmailServer(context.Background(), emailServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.GetEmailServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailServer`: EmailServerListResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailServerAPI.GetEmailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailServerListResponse**](EmailServerListResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailServers

> EmailServerListResponse ListEmailServers(ctx).Execute()

List all enrolled SMTP servers



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
    resp, r, err := apiClient.EmailServerAPI.ListEmailServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.ListEmailServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailServers`: EmailServerListResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailServerAPI.ListEmailServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailServersRequest struct via the builder pattern


### Return type

[**EmailServerListResponse**](EmailServerListResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestEmailServer

> TestEmailServer(ctx, emailServerId).EmailTestAddresses(emailTestAddresses).Execute()

Test an SMTP Server configuration



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
    emailServerId := "emailServerId_example" // string | 
    emailTestAddresses := *openapiclient.NewEmailTestAddresses("sender@host.com", "receiver@host.com") // EmailTestAddresses |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EmailServerAPI.TestEmailServer(context.Background(), emailServerId).EmailTestAddresses(emailTestAddresses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.TestEmailServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestEmailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailTestAddresses** | [**EmailTestAddresses**](EmailTestAddresses.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailServer

> EmailServerResponse UpdateEmailServer(ctx, emailServerId).EmailServerRequest(emailServerRequest).Execute()

Update an SMTP Server configuration



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
    emailServerId := "emailServerId_example" // string | 
    emailServerRequest := *openapiclient.NewEmailServerRequest() // EmailServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailServerAPI.UpdateEmailServer(context.Background(), emailServerId).EmailServerRequest(emailServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailServerAPI.UpdateEmailServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailServer`: EmailServerResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailServerAPI.UpdateEmailServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailServerRequest** | [**EmailServerRequest**](EmailServerRequest.md) |  | 

### Return type

[**EmailServerResponse**](EmailServerResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


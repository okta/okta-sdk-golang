# \ApplicationConnectionsApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsApi.md#ActivateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default/lifecycle/activate | Activate the default Provisioning Connection
[**DeactivateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsApi.md#DeactivateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default/lifecycle/deactivate | Deactivate the default Provisioning Connection
[**GetDefaultProvisioningConnectionForApplication**](ApplicationConnectionsApi.md#GetDefaultProvisioningConnectionForApplication) | **Get** /api/v1/apps/{appId}/connections/default | Retrieve the default Provisioning Connection
[**UpdateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsApi.md#UpdateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default | Update the default Provisioning Connection



## ActivateDefaultProvisioningConnectionForApplication

> ActivateDefaultProvisioningConnectionForApplication(ctx, appId).Execute()

Activate the default Provisioning Connection



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
    r, err := apiClient.ApplicationConnectionsApi.ActivateDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsApi.ActivateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


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


## DeactivateDefaultProvisioningConnectionForApplication

> DeactivateDefaultProvisioningConnectionForApplication(ctx, appId).Execute()

Deactivate the default Provisioning Connection



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
    r, err := apiClient.ApplicationConnectionsApi.DeactivateDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsApi.DeactivateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


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


## GetDefaultProvisioningConnectionForApplication

> ProvisioningConnection GetDefaultProvisioningConnectionForApplication(ctx, appId).Execute()

Retrieve the default Provisioning Connection



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
    resp, r, err := apiClient.ApplicationConnectionsApi.GetDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsApi.GetDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultProvisioningConnectionForApplication`: ProvisioningConnection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConnectionsApi.GetDefaultProvisioningConnectionForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningConnection**](ProvisioningConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultProvisioningConnectionForApplication

> ProvisioningConnection UpdateDefaultProvisioningConnectionForApplication(ctx, appId).ProvisioningConnectionRequest(provisioningConnectionRequest).Activate(activate).Execute()

Update the default Provisioning Connection



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
    provisioningConnectionRequest := *openapiclient.NewProvisioningConnectionRequest(*openapiclient.NewProvisioningConnectionProfile()) // ProvisioningConnectionRequest | 
    activate := true // bool | Activates the Provisioning Connection (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConnectionsApi.UpdateDefaultProvisioningConnectionForApplication(context.Background(), appId).ProvisioningConnectionRequest(provisioningConnectionRequest).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsApi.UpdateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaultProvisioningConnectionForApplication`: ProvisioningConnection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConnectionsApi.UpdateDefaultProvisioningConnectionForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningConnectionRequest** | [**ProvisioningConnectionRequest**](ProvisioningConnectionRequest.md) |  | 
 **activate** | **bool** | Activates the Provisioning Connection | 

### Return type

[**ProvisioningConnection**](ProvisioningConnection.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


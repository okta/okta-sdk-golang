# \ApplicationConnectionsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsAPI.md#ActivateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default/lifecycle/activate | Activate the default Provisioning Connection
[**DeactivateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsAPI.md#DeactivateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default/lifecycle/deactivate | Deactivate the default Provisioning Connection
[**GetDefaultProvisioningConnectionForApplication**](ApplicationConnectionsAPI.md#GetDefaultProvisioningConnectionForApplication) | **Get** /api/v1/apps/{appId}/connections/default | Retrieve the default Provisioning Connection
[**UpdateDefaultProvisioningConnectionForApplication**](ApplicationConnectionsAPI.md#UpdateDefaultProvisioningConnectionForApplication) | **Post** /api/v1/apps/{appId}/connections/default | Update the default Provisioning Connection



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationConnectionsAPI.ActivateDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsAPI.ActivateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationConnectionsAPI.DeactivateDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsAPI.DeactivateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

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

> GetDefaultProvisioningConnectionForApplication200Response GetDefaultProvisioningConnectionForApplication(ctx, appId).Execute()

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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConnectionsAPI.GetDefaultProvisioningConnectionForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsAPI.GetDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultProvisioningConnectionForApplication`: GetDefaultProvisioningConnectionForApplication200Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConnectionsAPI.GetDefaultProvisioningConnectionForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultProvisioningConnectionForApplication200Response**](GetDefaultProvisioningConnectionForApplication200Response.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultProvisioningConnectionForApplication

> ProvisioningConnection UpdateDefaultProvisioningConnectionForApplication(ctx, appId).UpdateDefaultProvisioningConnectionForApplicationRequest(updateDefaultProvisioningConnectionForApplicationRequest).Activate(activate).Execute()

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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    updateDefaultProvisioningConnectionForApplicationRequest := openapiclient.updateDefaultProvisioningConnectionForApplication_request{ProvisioningConnectionOauthRequest: openapiclient.NewProvisioningConnectionOauthRequest(*openapiclient.NewProvisioningConnectionProfileOauth("AuthScheme_example"))} // UpdateDefaultProvisioningConnectionForApplicationRequest | 
    activate := true // bool | Activates the Provisioning Connection (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationConnectionsAPI.UpdateDefaultProvisioningConnectionForApplication(context.Background(), appId).UpdateDefaultProvisioningConnectionForApplicationRequest(updateDefaultProvisioningConnectionForApplicationRequest).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationConnectionsAPI.UpdateDefaultProvisioningConnectionForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefaultProvisioningConnectionForApplication`: ProvisioningConnection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationConnectionsAPI.UpdateDefaultProvisioningConnectionForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultProvisioningConnectionForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDefaultProvisioningConnectionForApplicationRequest** | [**UpdateDefaultProvisioningConnectionForApplicationRequest**](UpdateDefaultProvisioningConnectionForApplicationRequest.md) |  | 
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


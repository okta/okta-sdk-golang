# \DeviceIntegrationsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDeviceIntegration**](DeviceIntegrationsAPI.md#ActivateDeviceIntegration) | **Post** /api/v1/device-integrations/{deviceIntegrationId}/lifecycle/activate | Activate a device integration
[**DeactivateDeviceIntegration**](DeviceIntegrationsAPI.md#DeactivateDeviceIntegration) | **Post** /api/v1/device-integrations/{deviceIntegrationId}/lifecycle/deactivate | Deactivate a device integration
[**GetDeviceIntegration**](DeviceIntegrationsAPI.md#GetDeviceIntegration) | **Get** /api/v1/device-integrations/{deviceIntegrationId} | Retrieve a device integration
[**ListDeviceIntegrations**](DeviceIntegrationsAPI.md#ListDeviceIntegrations) | **Get** /api/v1/device-integrations | List all device integrations



## ActivateDeviceIntegration

> DeviceIntegrations ActivateDeviceIntegration(ctx, deviceIntegrationId).Execute()

Activate a device integration



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
	deviceIntegrationId := "deviceIntegrationId_example" // string | The ID of the device integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceIntegrationsAPI.ActivateDeviceIntegration(context.Background(), deviceIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceIntegrationsAPI.ActivateDeviceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateDeviceIntegration`: DeviceIntegrations
	fmt.Fprintf(os.Stdout, "Response from `DeviceIntegrationsAPI.ActivateDeviceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIntegrationId** | **string** | The ID of the device integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateDeviceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceIntegrations**](DeviceIntegrations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateDeviceIntegration

> DeviceIntegrations DeactivateDeviceIntegration(ctx, deviceIntegrationId).Execute()

Deactivate a device integration



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
	deviceIntegrationId := "deviceIntegrationId_example" // string | The ID of the device integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceIntegrationsAPI.DeactivateDeviceIntegration(context.Background(), deviceIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceIntegrationsAPI.DeactivateDeviceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateDeviceIntegration`: DeviceIntegrations
	fmt.Fprintf(os.Stdout, "Response from `DeviceIntegrationsAPI.DeactivateDeviceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIntegrationId** | **string** | The ID of the device integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateDeviceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceIntegrations**](DeviceIntegrations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceIntegration

> DeviceIntegrations GetDeviceIntegration(ctx, deviceIntegrationId).Execute()

Retrieve a device integration



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
	deviceIntegrationId := "deviceIntegrationId_example" // string | The ID of the device integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceIntegrationsAPI.GetDeviceIntegration(context.Background(), deviceIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceIntegrationsAPI.GetDeviceIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceIntegration`: DeviceIntegrations
	fmt.Fprintf(os.Stdout, "Response from `DeviceIntegrationsAPI.GetDeviceIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceIntegrationId** | **string** | The ID of the device integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceIntegrations**](DeviceIntegrations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceIntegrations

> []DeviceIntegrations ListDeviceIntegrations(ctx).Execute()

List all device integrations



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
	resp, r, err := apiClient.DeviceIntegrationsAPI.ListDeviceIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceIntegrationsAPI.ListDeviceIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceIntegrations`: []DeviceIntegrations
	fmt.Fprintf(os.Stdout, "Response from `DeviceIntegrationsAPI.ListDeviceIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceIntegrationsRequest struct via the builder pattern


### Return type

[**[]DeviceIntegrations**](DeviceIntegrations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


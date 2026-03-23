# \CustomTelephonyProviderAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCustomTelephonyCredential**](CustomTelephonyProviderAPI.md#ActivateCustomTelephonyCredential) | **Post** /api/v1/telephony-providers/{customTelephonyProviderId}/lifecycle/activate | Activate a custom telephony provider
[**CreateCustomTelephonyProviderCredentials**](CustomTelephonyProviderAPI.md#CreateCustomTelephonyProviderCredentials) | **Post** /api/v1/telephony-providers | Create a custom telephony provider
[**DeactivateCustomTelephonyCredential**](CustomTelephonyProviderAPI.md#DeactivateCustomTelephonyCredential) | **Post** /api/v1/telephony-providers/{customTelephonyProviderId}/lifecycle/deactivate | Deactivate a custom telephony provider
[**DeleteCustomTelephonyProviderCredential**](CustomTelephonyProviderAPI.md#DeleteCustomTelephonyProviderCredential) | **Delete** /api/v1/telephony-providers/{customTelephonyProviderId} | Delete a custom telephony provider
[**GetCustomTelephonyProviderCredential**](CustomTelephonyProviderAPI.md#GetCustomTelephonyProviderCredential) | **Get** /api/v1/telephony-providers/{customTelephonyProviderId} | Retrieve a custom telephony provider
[**ListAllCustomTelephonyProviderCredentials**](CustomTelephonyProviderAPI.md#ListAllCustomTelephonyProviderCredentials) | **Get** /api/v1/telephony-providers | List all custom telephony providers
[**SendTestCustomTelephonyProviderCredential**](CustomTelephonyProviderAPI.md#SendTestCustomTelephonyProviderCredential) | **Post** /api/v1/telephony-providers/{customTelephonyProviderId}/test | Send a test message from a custom telephony provider
[**SetAsPrimaryCustomTelephonyCredential**](CustomTelephonyProviderAPI.md#SetAsPrimaryCustomTelephonyCredential) | **Post** /api/v1/telephony-providers/{customTelephonyProviderId}/setAsPrimary | Set a custom telephony provider as a primary telephony provider
[**UpdateCustomTelephonyProviderCredential**](CustomTelephonyProviderAPI.md#UpdateCustomTelephonyProviderCredential) | **Patch** /api/v1/telephony-providers/{customTelephonyProviderId} | Update a custom telephony provider credential



## ActivateCustomTelephonyCredential

> CustomTelephonyProviderCredentialResponse ActivateCustomTelephonyCredential(ctx, customTelephonyProviderId).Execute()

Activate a custom telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.ActivateCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.ActivateCustomTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateCustomTelephonyCredential`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.ActivateCustomTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateCustomTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomTelephonyProviderCredentials

> CustomTelephonyProviderCredentialResponse CreateCustomTelephonyProviderCredentials(ctx).CustomTelephonyProviderCredentialCreateRequest(customTelephonyProviderCredentialCreateRequest).Execute()

Create a custom telephony provider



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
	customTelephonyProviderCredentialCreateRequest := *openapiclient.NewCustomTelephonyProviderCredentialCreateRequest() // CustomTelephonyProviderCredentialCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.CreateCustomTelephonyProviderCredentials(context.Background()).CustomTelephonyProviderCredentialCreateRequest(customTelephonyProviderCredentialCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.CreateCustomTelephonyProviderCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomTelephonyProviderCredentials`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.CreateCustomTelephonyProviderCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomTelephonyProviderCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customTelephonyProviderCredentialCreateRequest** | [**CustomTelephonyProviderCredentialCreateRequest**](CustomTelephonyProviderCredentialCreateRequest.md) |  | 

### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateCustomTelephonyCredential

> CustomTelephonyProviderCredentialResponse DeactivateCustomTelephonyCredential(ctx, customTelephonyProviderId).Execute()

Deactivate a custom telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.DeactivateCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.DeactivateCustomTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateCustomTelephonyCredential`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.DeactivateCustomTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateCustomTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomTelephonyProviderCredential

> DeleteCustomTelephonyProviderCredential(ctx, customTelephonyProviderId).Execute()

Delete a custom telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomTelephonyProviderAPI.DeleteCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.DeleteCustomTelephonyProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomTelephonyProviderCredentialRequest struct via the builder pattern


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


## GetCustomTelephonyProviderCredential

> CustomTelephonyProviderCredentialResponse GetCustomTelephonyProviderCredential(ctx, customTelephonyProviderId).Execute()

Retrieve a custom telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.GetCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.GetCustomTelephonyProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomTelephonyProviderCredential`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.GetCustomTelephonyProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomTelephonyProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCustomTelephonyProviderCredentials

> []CustomTelephonyProviderCredentialResponse ListAllCustomTelephonyProviderCredentials(ctx).Execute()

List all custom telephony providers



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
	resp, r, err := apiClient.CustomTelephonyProviderAPI.ListAllCustomTelephonyProviderCredentials(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.ListAllCustomTelephonyProviderCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllCustomTelephonyProviderCredentials`: []CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.ListAllCustomTelephonyProviderCredentials`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllCustomTelephonyProviderCredentialsRequest struct via the builder pattern


### Return type

[**[]CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestCustomTelephonyProviderCredential

> SendTestCustomTelephonyProviderCredential(ctx, customTelephonyProviderId).CustomTelephonyProviderCredentialSendTestRequest(customTelephonyProviderCredentialSendTestRequest).Execute()

Send a test message from a custom telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider
	customTelephonyProviderCredentialSendTestRequest := *openapiclient.NewCustomTelephonyProviderCredentialSendTestRequest() // CustomTelephonyProviderCredentialSendTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomTelephonyProviderAPI.SendTestCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).CustomTelephonyProviderCredentialSendTestRequest(customTelephonyProviderCredentialSendTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.SendTestCustomTelephonyProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestCustomTelephonyProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customTelephonyProviderCredentialSendTestRequest** | [**CustomTelephonyProviderCredentialSendTestRequest**](CustomTelephonyProviderCredentialSendTestRequest.md) |  | 

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


## SetAsPrimaryCustomTelephonyCredential

> CustomTelephonyProviderCredentialResponse SetAsPrimaryCustomTelephonyCredential(ctx, customTelephonyProviderId).Execute()

Set a custom telephony provider as a primary telephony provider



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.SetAsPrimaryCustomTelephonyCredential(context.Background(), customTelephonyProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.SetAsPrimaryCustomTelephonyCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAsPrimaryCustomTelephonyCredential`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.SetAsPrimaryCustomTelephonyCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAsPrimaryCustomTelephonyCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomTelephonyProviderCredential

> CustomTelephonyProviderCredentialResponse UpdateCustomTelephonyProviderCredential(ctx, customTelephonyProviderId).CustomTelephonyProviderCredentialUpdateRequest(customTelephonyProviderCredentialUpdateRequest).Execute()

Update a custom telephony provider credential



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
	customTelephonyProviderId := "customTelephonyProviderId_example" // string | The ID of the custom telephony provider
	customTelephonyProviderCredentialUpdateRequest := *openapiclient.NewCustomTelephonyProviderCredentialUpdateRequest() // CustomTelephonyProviderCredentialUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomTelephonyProviderAPI.UpdateCustomTelephonyProviderCredential(context.Background(), customTelephonyProviderId).CustomTelephonyProviderCredentialUpdateRequest(customTelephonyProviderCredentialUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomTelephonyProviderAPI.UpdateCustomTelephonyProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomTelephonyProviderCredential`: CustomTelephonyProviderCredentialResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomTelephonyProviderAPI.UpdateCustomTelephonyProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customTelephonyProviderId** | **string** | The ID of the custom telephony provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomTelephonyProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customTelephonyProviderCredentialUpdateRequest** | [**CustomTelephonyProviderCredentialUpdateRequest**](CustomTelephonyProviderCredentialUpdateRequest.md) |  | 

### Return type

[**CustomTelephonyProviderCredentialResponse**](CustomTelephonyProviderCredentialResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


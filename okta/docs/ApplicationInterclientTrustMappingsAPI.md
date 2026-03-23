# \ApplicationInterclientTrustMappingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterclientTrustMapping**](ApplicationInterclientTrustMappingsAPI.md#CreateInterclientTrustMapping) | **Post** /api/v1/apps/{appId}/interclient-allowed-apps | Create an allowed app mapping for a target app
[**DeleteInterclientTrustMapping**](ApplicationInterclientTrustMappingsAPI.md#DeleteInterclientTrustMapping) | **Delete** /api/v1/apps/{appId}/interclient-allowed-apps/{allowedAppId} | Delete an interclient trust mapping
[**ListInterclientAllowedApplications**](ApplicationInterclientTrustMappingsAPI.md#ListInterclientAllowedApplications) | **Get** /api/v1/apps/{appId}/interclient-allowed-apps | List all allowed apps for a target app
[**ListInterclientTargetApplications**](ApplicationInterclientTrustMappingsAPI.md#ListInterclientTargetApplications) | **Get** /api/v1/apps/{appId}/interclient-target-apps | List all target apps for an allowed app



## CreateInterclientTrustMapping

> InterclientTrustMapping CreateInterclientTrustMapping(ctx, appId).InterclientTrustMappingRequestBody(interclientTrustMappingRequestBody).Execute()

Create an allowed app mapping for a target app



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
	interclientTrustMappingRequestBody := *openapiclient.NewInterclientTrustMappingRequestBody() // InterclientTrustMappingRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationInterclientTrustMappingsAPI.CreateInterclientTrustMapping(context.Background(), appId).InterclientTrustMappingRequestBody(interclientTrustMappingRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInterclientTrustMappingsAPI.CreateInterclientTrustMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInterclientTrustMapping`: InterclientTrustMapping
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInterclientTrustMappingsAPI.CreateInterclientTrustMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterclientTrustMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interclientTrustMappingRequestBody** | [**InterclientTrustMappingRequestBody**](InterclientTrustMappingRequestBody.md) |  | 

### Return type

[**InterclientTrustMapping**](InterclientTrustMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterclientTrustMapping

> DeleteInterclientTrustMapping(ctx, appId, allowedAppId).Execute()

Delete an interclient trust mapping



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
	allowedAppId := "0oa1elyw9EAkUNUrW0g5" // string | App ID of the allowed app instance to delete mapping from the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationInterclientTrustMappingsAPI.DeleteInterclientTrustMapping(context.Background(), appId, allowedAppId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInterclientTrustMappingsAPI.DeleteInterclientTrustMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**allowedAppId** | **string** | App ID of the allowed app instance to delete mapping from the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterclientTrustMappingRequest struct via the builder pattern


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


## ListInterclientAllowedApplications

> []string ListInterclientAllowedApplications(ctx, appId).Execute()

List all allowed apps for a target app



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
	resp, r, err := apiClient.ApplicationInterclientTrustMappingsAPI.ListInterclientAllowedApplications(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInterclientTrustMappingsAPI.ListInterclientAllowedApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInterclientAllowedApplications`: []string
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInterclientTrustMappingsAPI.ListInterclientAllowedApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterclientAllowedApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterclientTargetApplications

> []string ListInterclientTargetApplications(ctx, appId).Execute()

List all target apps for an allowed app



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
	resp, r, err := apiClient.ApplicationInterclientTrustMappingsAPI.ListInterclientTargetApplications(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationInterclientTrustMappingsAPI.ListInterclientTargetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInterclientTargetApplications`: []string
	fmt.Fprintf(os.Stdout, "Response from `ApplicationInterclientTrustMappingsAPI.ListInterclientTargetApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterclientTargetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ServiceAccountAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppServiceAccount**](ServiceAccountAPI.md#CreateAppServiceAccount) | **Post** /privileged-access/api/v1/service-accounts | Create an app service account
[**DeleteAppServiceAccount**](ServiceAccountAPI.md#DeleteAppServiceAccount) | **Delete** /privileged-access/api/v1/service-accounts/{id} | Delete an app service account
[**GetAppServiceAccount**](ServiceAccountAPI.md#GetAppServiceAccount) | **Get** /privileged-access/api/v1/service-accounts/{id} | Retrieve an app service account
[**ListAppServiceAccounts**](ServiceAccountAPI.md#ListAppServiceAccounts) | **Get** /privileged-access/api/v1/service-accounts | List all app service accounts
[**UpdateAppServiceAccount**](ServiceAccountAPI.md#UpdateAppServiceAccount) | **Patch** /privileged-access/api/v1/service-accounts/{id} | Update an existing app service account



## CreateAppServiceAccount

> AppServiceAccount CreateAppServiceAccount(ctx).Body(body).Execute()

Create an app service account



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
	body := *openapiclient.NewAppServiceAccount("orn:okta:idp:00o1n8sbwArJ7OQRw406:apps:salesforce:0oa1gjh63g214q0Hq0g4", "salesforce Prod-5 account", "testuser-salesforce-5@example.com") // AppServiceAccount | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.CreateAppServiceAccount(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.CreateAppServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppServiceAccount`: AppServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.CreateAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppServiceAccount**](AppServiceAccount.md) |  | 

### Return type

[**AppServiceAccount**](AppServiceAccount.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppServiceAccount

> DeleteAppServiceAccount(ctx, id).Execute()

Delete an app service account



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
	id := "id_example" // string | ID of an existing service account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountAPI.DeleteAppServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.DeleteAppServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppServiceAccountRequest struct via the builder pattern


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


## GetAppServiceAccount

> AppServiceAccount GetAppServiceAccount(ctx, id).Execute()

Retrieve an app service account



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
	id := "id_example" // string | ID of an existing service account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.GetAppServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.GetAppServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppServiceAccount`: AppServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.GetAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppServiceAccount**](AppServiceAccount.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceAccounts

> []AppServiceAccount ListAppServiceAccounts(ctx).Limit(limit).After(after).Match(match).Execute()

List all app service accounts



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
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). (optional)
	match := "salesforce" // string | Searches for app service accounts where the account name (`name`), username (`username`), app instance label (`containerInstanceName`), or OIN app key name (`containerGlobalName`) contains the given value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.ListAppServiceAccounts(context.Background()).Limit(limit).After(after).Match(match).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.ListAppServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppServiceAccounts`: []AppServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.ListAppServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
 **match** | **string** | Searches for app service accounts where the account name (&#x60;name&#x60;), username (&#x60;username&#x60;), app instance label (&#x60;containerInstanceName&#x60;), or OIN app key name (&#x60;containerGlobalName&#x60;) contains the given value | 

### Return type

[**[]AppServiceAccount**](AppServiceAccount.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppServiceAccount

> AppServiceAccount UpdateAppServiceAccount(ctx, id).Body(body).Execute()

Update an existing app service account



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
	id := "id_example" // string | ID of an existing service account
	body := *openapiclient.NewAppServiceAccountForUpdate() // AppServiceAccountForUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.UpdateAppServiceAccount(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.UpdateAppServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppServiceAccount`: AppServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.UpdateAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AppServiceAccountForUpdate**](AppServiceAccountForUpdate.md) |  | 

### Return type

[**AppServiceAccount**](AppServiceAccount.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


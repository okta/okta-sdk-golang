# \OktaManagedUserAccountAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOktaManagedUserAccount**](OktaManagedUserAccountAPI.md#CreateOktaManagedUserAccount) | **Post** /privileged-access/api/v1/okta-service-accounts | Create an Okta managed user account
[**DeleteOktaManagedUserAccount**](OktaManagedUserAccountAPI.md#DeleteOktaManagedUserAccount) | **Delete** /privileged-access/api/v1/okta-service-accounts/{id} | Delete an Okta managed user account
[**GetOktaManagedUserAccount**](OktaManagedUserAccountAPI.md#GetOktaManagedUserAccount) | **Get** /privileged-access/api/v1/okta-service-accounts/{id} | Retrieve an Okta managed user account
[**ListOktaManagedUserAccounts**](OktaManagedUserAccountAPI.md#ListOktaManagedUserAccounts) | **Get** /privileged-access/api/v1/okta-service-accounts | List all Okta managed user accounts
[**UpdateOktaManagedUserAccount**](OktaManagedUserAccountAPI.md#UpdateOktaManagedUserAccount) | **Patch** /privileged-access/api/v1/okta-service-accounts/{id} | Update an Okta managed user account



## CreateOktaManagedUserAccount

> OktaManagedUserAccountResponse CreateOktaManagedUserAccount(ctx).Body(body).Execute()

Create an Okta managed user account



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
	body := *openapiclient.NewOktaManagedUserAccountRequest("AD Integrations Admin", "00u11s48P9zGW8yqm0g5") // OktaManagedUserAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OktaManagedUserAccountAPI.CreateOktaManagedUserAccount(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaManagedUserAccountAPI.CreateOktaManagedUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOktaManagedUserAccount`: OktaManagedUserAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `OktaManagedUserAccountAPI.CreateOktaManagedUserAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOktaManagedUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OktaManagedUserAccountRequest**](OktaManagedUserAccountRequest.md) |  | 

### Return type

[**OktaManagedUserAccountResponse**](OktaManagedUserAccountResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOktaManagedUserAccount

> DeleteOktaManagedUserAccount(ctx, id).Execute()

Delete an Okta managed user account



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
	id := "a747a818-a4c4-4446-8a87-704216495a08" // string | ID of an existing Okta managed user account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OktaManagedUserAccountAPI.DeleteOktaManagedUserAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaManagedUserAccountAPI.DeleteOktaManagedUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing Okta managed user account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOktaManagedUserAccountRequest struct via the builder pattern


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


## GetOktaManagedUserAccount

> OktaManagedUserAccountResponse GetOktaManagedUserAccount(ctx, id).Execute()

Retrieve an Okta managed user account



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
	id := "a747a818-a4c4-4446-8a87-704216495a08" // string | ID of an existing Okta managed user account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OktaManagedUserAccountAPI.GetOktaManagedUserAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaManagedUserAccountAPI.GetOktaManagedUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOktaManagedUserAccount`: OktaManagedUserAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `OktaManagedUserAccountAPI.GetOktaManagedUserAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing Okta managed user account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOktaManagedUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OktaManagedUserAccountResponse**](OktaManagedUserAccountResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOktaManagedUserAccounts

> []OktaManagedUserAccountResponse ListOktaManagedUserAccounts(ctx).Limit(limit).After(after).Match(match).Execute()

List all Okta managed user accounts



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
	match := "admin" // string | Searches for Okta managed user accounts where the account name (`name`) or username (`username`) contains the given value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OktaManagedUserAccountAPI.ListOktaManagedUserAccounts(context.Background()).Limit(limit).After(after).Match(match).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaManagedUserAccountAPI.ListOktaManagedUserAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOktaManagedUserAccounts`: []OktaManagedUserAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `OktaManagedUserAccountAPI.ListOktaManagedUserAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOktaManagedUserAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination) and [Link header](https://developer.okta.com/docs/api/#link-header). | 
 **match** | **string** | Searches for Okta managed user accounts where the account name (&#x60;name&#x60;) or username (&#x60;username&#x60;) contains the given value | 

### Return type

[**[]OktaManagedUserAccountResponse**](OktaManagedUserAccountResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOktaManagedUserAccount

> OktaManagedUserAccountResponse UpdateOktaManagedUserAccount(ctx, id).Body(body).Execute()

Update an Okta managed user account



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
	id := "a747a818-a4c4-4446-8a87-704216495a08" // string | ID of an existing Okta managed user account
	body := *openapiclient.NewOktaManagedUserAccountForUpdate() // OktaManagedUserAccountForUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OktaManagedUserAccountAPI.UpdateOktaManagedUserAccount(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaManagedUserAccountAPI.UpdateOktaManagedUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOktaManagedUserAccount`: OktaManagedUserAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `OktaManagedUserAccountAPI.UpdateOktaManagedUserAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an existing Okta managed user account | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOktaManagedUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OktaManagedUserAccountForUpdate**](OktaManagedUserAccountForUpdate.md) |  | 

### Return type

[**OktaManagedUserAccountResponse**](OktaManagedUserAccountResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


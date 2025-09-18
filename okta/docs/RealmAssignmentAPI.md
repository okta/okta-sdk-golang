# \RealmAssignmentAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateRealmAssignment**](RealmAssignmentAPI.md#ActivateRealmAssignment) | **Post** /api/v1/realm-assignments/{assignmentId}/lifecycle/activate | Activate a realm assignment
[**CreateRealmAssignment**](RealmAssignmentAPI.md#CreateRealmAssignment) | **Post** /api/v1/realm-assignments | Create a realm assignment
[**DeactivateRealmAssignment**](RealmAssignmentAPI.md#DeactivateRealmAssignment) | **Post** /api/v1/realm-assignments/{assignmentId}/lifecycle/deactivate | Deactivate a realm assignment
[**DeleteRealmAssignment**](RealmAssignmentAPI.md#DeleteRealmAssignment) | **Delete** /api/v1/realm-assignments/{assignmentId} | Delete a realm assignment
[**ExecuteRealmAssignment**](RealmAssignmentAPI.md#ExecuteRealmAssignment) | **Post** /api/v1/realm-assignments/operations | Execute a realm assignment
[**GetRealmAssignment**](RealmAssignmentAPI.md#GetRealmAssignment) | **Get** /api/v1/realm-assignments/{assignmentId} | Retrieve a realm assignment
[**ListRealmAssignmentOperations**](RealmAssignmentAPI.md#ListRealmAssignmentOperations) | **Get** /api/v1/realm-assignments/operations | List all realm assignment operations
[**ListRealmAssignments**](RealmAssignmentAPI.md#ListRealmAssignments) | **Get** /api/v1/realm-assignments | List all realm assignments
[**ReplaceRealmAssignment**](RealmAssignmentAPI.md#ReplaceRealmAssignment) | **Put** /api/v1/realm-assignments/{assignmentId} | Replace a realm assignment



## ActivateRealmAssignment

> ActivateRealmAssignment(ctx, assignmentId).Execute()

Activate a realm assignment



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
	assignmentId := "rul2jy7jLUlnO3ng00g4" // string | ID of the realm assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmAssignmentAPI.ActivateRealmAssignment(context.Background(), assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.ActivateRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **string** | ID of the realm assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateRealmAssignmentRequest struct via the builder pattern


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


## CreateRealmAssignment

> RealmAssignment CreateRealmAssignment(ctx).Body(body).Execute()

Create a realm assignment



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
	body := *openapiclient.NewCreateRealmAssignmentRequest() // CreateRealmAssignmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.CreateRealmAssignment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.CreateRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRealmAssignment`: RealmAssignment
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.CreateRealmAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRealmAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRealmAssignmentRequest**](CreateRealmAssignmentRequest.md) |  | 

### Return type

[**RealmAssignment**](RealmAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateRealmAssignment

> DeactivateRealmAssignment(ctx, assignmentId).Execute()

Deactivate a realm assignment



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
	assignmentId := "rul2jy7jLUlnO3ng00g4" // string | ID of the realm assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmAssignmentAPI.DeactivateRealmAssignment(context.Background(), assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.DeactivateRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **string** | ID of the realm assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateRealmAssignmentRequest struct via the builder pattern


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


## DeleteRealmAssignment

> DeleteRealmAssignment(ctx, assignmentId).Execute()

Delete a realm assignment



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
	assignmentId := "rul2jy7jLUlnO3ng00g4" // string | ID of the realm assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmAssignmentAPI.DeleteRealmAssignment(context.Background(), assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.DeleteRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **string** | ID of the realm assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRealmAssignmentRequest struct via the builder pattern


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


## ExecuteRealmAssignment

> OperationResponse ExecuteRealmAssignment(ctx).Body(body).Execute()

Execute a realm assignment



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
	body := *openapiclient.NewOperationRequest() // OperationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.ExecuteRealmAssignment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.ExecuteRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteRealmAssignment`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.ExecuteRealmAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteRealmAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OperationRequest**](OperationRequest.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealmAssignment

> RealmAssignment GetRealmAssignment(ctx, assignmentId).Execute()

Retrieve a realm assignment



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
	assignmentId := "rul2jy7jLUlnO3ng00g4" // string | ID of the realm assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.GetRealmAssignment(context.Background(), assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.GetRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealmAssignment`: RealmAssignment
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.GetRealmAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **string** | ID of the realm assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealmAssignment**](RealmAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealmAssignmentOperations

> []OperationResponse ListRealmAssignmentOperations(ctx).Limit(limit).After(after).Execute()

List all realm assignment operations



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.ListRealmAssignmentOperations(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.ListRealmAssignmentOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRealmAssignmentOperations`: []OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.ListRealmAssignmentOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRealmAssignmentOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**[]OperationResponse**](OperationResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealmAssignments

> []RealmAssignment ListRealmAssignments(ctx).Limit(limit).After(after).Execute()

List all realm assignments



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
	after := "after_example" // string | The cursor used for pagination. It represents the priority of the last realm assignment returned in the previous fetch operation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.ListRealmAssignments(context.Background()).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.ListRealmAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRealmAssignments`: []RealmAssignment
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.ListRealmAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRealmAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **after** | **string** | The cursor used for pagination. It represents the priority of the last realm assignment returned in the previous fetch operation. | 

### Return type

[**[]RealmAssignment**](RealmAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRealmAssignment

> RealmAssignment ReplaceRealmAssignment(ctx, assignmentId).Body(body).Execute()

Replace a realm assignment



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
	assignmentId := "rul2jy7jLUlnO3ng00g4" // string | ID of the realm assignment
	body := *openapiclient.NewUpdateRealmAssignmentRequest() // UpdateRealmAssignmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmAssignmentAPI.ReplaceRealmAssignment(context.Background(), assignmentId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmAssignmentAPI.ReplaceRealmAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRealmAssignment`: RealmAssignment
	fmt.Fprintf(os.Stdout, "Response from `RealmAssignmentAPI.ReplaceRealmAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assignmentId** | **string** | ID of the realm assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRealmAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateRealmAssignmentRequest**](UpdateRealmAssignmentRequest.md) |  | 

### Return type

[**RealmAssignment**](RealmAssignment.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


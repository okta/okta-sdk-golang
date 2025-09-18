# \UserClassificationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserClassification**](UserClassificationAPI.md#GetUserClassification) | **Get** /api/v1/users/{userId}/classification | Retrieve a user&#39;s classification
[**ReplaceUserClassification**](UserClassificationAPI.md#ReplaceUserClassification) | **Put** /api/v1/users/{userId}/classification | Replace the user&#39;s classification



## GetUserClassification

> UserClassification GetUserClassification(ctx, userId).Execute()

Retrieve a user's classification



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserClassificationAPI.GetUserClassification(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserClassificationAPI.GetUserClassification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserClassification`: UserClassification
	fmt.Fprintf(os.Stdout, "Response from `UserClassificationAPI.GetUserClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserClassification**](UserClassification.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUserClassification

> UserClassification ReplaceUserClassification(ctx, userId).ReplaceUserClassification(replaceUserClassification).Execute()

Replace the user's classification



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	replaceUserClassification := *openapiclient.NewReplaceUserClassification() // ReplaceUserClassification | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserClassificationAPI.ReplaceUserClassification(context.Background(), userId).ReplaceUserClassification(replaceUserClassification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserClassificationAPI.ReplaceUserClassification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceUserClassification`: UserClassification
	fmt.Fprintf(os.Stdout, "Response from `UserClassificationAPI.ReplaceUserClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceUserClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceUserClassification** | [**ReplaceUserClassification**](ReplaceUserClassification.md) |  | 

### Return type

[**UserClassification**](UserClassification.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


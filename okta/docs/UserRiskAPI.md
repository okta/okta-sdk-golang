# \UserRiskAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserRisk**](UserRiskAPI.md#GetUserRisk) | **Get** /api/v1/users/{userId}/risk | Retrieve the user&#39;s risk
[**UpsertUserRisk**](UserRiskAPI.md#UpsertUserRisk) | **Put** /api/v1/users/{userId}/risk | Upsert the user&#39;s risk



## GetUserRisk

> GetUserRisk200Response GetUserRisk(ctx, userId).Execute()

Retrieve the user's risk



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
	resp, r, err := apiClient.UserRiskAPI.GetUserRisk(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRiskAPI.GetUserRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRisk`: GetUserRisk200Response
	fmt.Fprintf(os.Stdout, "Response from `UserRiskAPI.GetUserRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserRisk200Response**](GetUserRisk200Response.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUserRisk

> UserRiskPutResponse UpsertUserRisk(ctx, userId).UserRiskRequest(userRiskRequest).Execute()

Upsert the user's risk



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
	userRiskRequest := *openapiclient.NewUserRiskRequest() // UserRiskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserRiskAPI.UpsertUserRisk(context.Background(), userId).UserRiskRequest(userRiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserRiskAPI.UpsertUserRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertUserRisk`: UserRiskPutResponse
	fmt.Fprintf(os.Stdout, "Response from `UserRiskAPI.UpsertUserRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUserRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRiskRequest** | [**UserRiskRequest**](UserRiskRequest.md) |  | 

### Return type

[**UserRiskPutResponse**](UserRiskPutResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ApplicationSSOFederatedClaimsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFederatedClaim**](ApplicationSSOFederatedClaimsAPI.md#CreateFederatedClaim) | **Post** /api/v1/apps/{appId}/federated-claims | Create a federated claim
[**DeleteFederatedClaim**](ApplicationSSOFederatedClaimsAPI.md#DeleteFederatedClaim) | **Delete** /api/v1/apps/{appId}/federated-claims/{claimId} | Delete a federated claim
[**GetFederatedClaim**](ApplicationSSOFederatedClaimsAPI.md#GetFederatedClaim) | **Get** /api/v1/apps/{appId}/federated-claims/{claimId} | Retrieve a federated claim
[**ListFederatedClaims**](ApplicationSSOFederatedClaimsAPI.md#ListFederatedClaims) | **Get** /api/v1/apps/{appId}/federated-claims | List all configured federated claims
[**ReplaceFederatedClaim**](ApplicationSSOFederatedClaimsAPI.md#ReplaceFederatedClaim) | **Put** /api/v1/apps/{appId}/federated-claims/{claimId} | Replace a federated claim



## CreateFederatedClaim

> FederatedClaim CreateFederatedClaim(ctx, appId).FederatedClaimRequestBody(federatedClaimRequestBody).Execute()

Create a federated claim



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
	federatedClaimRequestBody := *openapiclient.NewFederatedClaimRequestBody() // FederatedClaimRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOFederatedClaimsAPI.CreateFederatedClaim(context.Background(), appId).FederatedClaimRequestBody(federatedClaimRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOFederatedClaimsAPI.CreateFederatedClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFederatedClaim`: FederatedClaim
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOFederatedClaimsAPI.CreateFederatedClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFederatedClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **federatedClaimRequestBody** | [**FederatedClaimRequestBody**](FederatedClaimRequestBody.md) |  | 

### Return type

[**FederatedClaim**](FederatedClaim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFederatedClaim

> DeleteFederatedClaim(ctx, appId, claimId).Execute()

Delete a federated claim



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
	claimId := "ofc2f4zrZbs8nUa7p0g4" // string | The unique `id` of the federated claim

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationSSOFederatedClaimsAPI.DeleteFederatedClaim(context.Background(), appId, claimId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOFederatedClaimsAPI.DeleteFederatedClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**claimId** | **string** | The unique &#x60;id&#x60; of the federated claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFederatedClaimRequest struct via the builder pattern


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


## GetFederatedClaim

> FederatedClaimRequestBody GetFederatedClaim(ctx, appId, claimId).Execute()

Retrieve a federated claim



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
	claimId := "ofc2f4zrZbs8nUa7p0g4" // string | The unique `id` of the federated claim

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOFederatedClaimsAPI.GetFederatedClaim(context.Background(), appId, claimId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOFederatedClaimsAPI.GetFederatedClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFederatedClaim`: FederatedClaimRequestBody
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOFederatedClaimsAPI.GetFederatedClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**claimId** | **string** | The unique &#x60;id&#x60; of the federated claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFederatedClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FederatedClaimRequestBody**](FederatedClaimRequestBody.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFederatedClaims

> []FederatedClaim ListFederatedClaims(ctx, appId).Execute()

List all configured federated claims



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
	resp, r, err := apiClient.ApplicationSSOFederatedClaimsAPI.ListFederatedClaims(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOFederatedClaimsAPI.ListFederatedClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFederatedClaims`: []FederatedClaim
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOFederatedClaimsAPI.ListFederatedClaims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFederatedClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FederatedClaim**](FederatedClaim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFederatedClaim

> FederatedClaim ReplaceFederatedClaim(ctx, appId, claimId).FederatedClaim(federatedClaim).Execute()

Replace a federated claim



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
	claimId := "ofc2f4zrZbs8nUa7p0g4" // string | The unique `id` of the federated claim
	federatedClaim := *openapiclient.NewFederatedClaim() // FederatedClaim |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOFederatedClaimsAPI.ReplaceFederatedClaim(context.Background(), appId, claimId).FederatedClaim(federatedClaim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOFederatedClaimsAPI.ReplaceFederatedClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFederatedClaim`: FederatedClaim
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOFederatedClaimsAPI.ReplaceFederatedClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**claimId** | **string** | The unique &#x60;id&#x60; of the federated claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFederatedClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **federatedClaim** | [**FederatedClaim**](FederatedClaim.md) |  | 

### Return type

[**FederatedClaim**](FederatedClaim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


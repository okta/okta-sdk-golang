# \PrincipalRateLimitAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrincipalRateLimitEntity**](PrincipalRateLimitAPI.md#CreatePrincipalRateLimitEntity) | **Post** /api/v1/principal-rate-limits | Create a principal rate limit
[**GetPrincipalRateLimitEntity**](PrincipalRateLimitAPI.md#GetPrincipalRateLimitEntity) | **Get** /api/v1/principal-rate-limits/{principalRateLimitId} | Retrieve a principal rate limit
[**ListPrincipalRateLimitEntities**](PrincipalRateLimitAPI.md#ListPrincipalRateLimitEntities) | **Get** /api/v1/principal-rate-limits | List all principal rate limits
[**ReplacePrincipalRateLimitEntity**](PrincipalRateLimitAPI.md#ReplacePrincipalRateLimitEntity) | **Put** /api/v1/principal-rate-limits/{principalRateLimitId} | Replace a principal rate limit



## CreatePrincipalRateLimitEntity

> PrincipalRateLimitEntity CreatePrincipalRateLimitEntity(ctx).Entity(entity).Execute()

Create a principal rate limit



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
	entity := *openapiclient.NewPrincipalRateLimitEntity("PrincipalId_example", "PrincipalType_example") // PrincipalRateLimitEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalRateLimitAPI.CreatePrincipalRateLimitEntity(context.Background()).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitAPI.CreatePrincipalRateLimitEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrincipalRateLimitEntity`: PrincipalRateLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitAPI.CreatePrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md) |  | 

### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalRateLimitEntity

> PrincipalRateLimitEntity GetPrincipalRateLimitEntity(ctx, principalRateLimitId).Execute()

Retrieve a principal rate limit



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
	principalRateLimitId := "0oacamvryxiyMqgiY1d7" // string | ID of the principal rate limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalRateLimitAPI.GetPrincipalRateLimitEntity(context.Background(), principalRateLimitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitAPI.GetPrincipalRateLimitEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalRateLimitEntity`: PrincipalRateLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitAPI.GetPrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRateLimitId** | **string** | ID of the principal rate limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrincipalRateLimitEntities

> []PrincipalRateLimitEntity ListPrincipalRateLimitEntities(ctx).Filter(filter).After(after).Limit(limit).Execute()

List all principal rate limits



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
	filter := "filter_example" // string | Filters the list of principal rate limit entities by the provided principal type (`principalType`). For example, `filter=principalType eq \"SSWS_TOKEN\"` or `filter=principalType eq \"OAUTH_CLIENT\"`.
	after := "after_example" // string | The cursor to use for pagination. It's an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | Specifies the number of items to return in a single response page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalRateLimitAPI.ListPrincipalRateLimitEntities(context.Background()).Filter(filter).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitAPI.ListPrincipalRateLimitEntities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrincipalRateLimitEntities`: []PrincipalRateLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitAPI.ListPrincipalRateLimitEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalRateLimitEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filters the list of principal rate limit entities by the provided principal type (&#x60;principalType&#x60;). For example, &#x60;filter&#x3D;principalType eq \&quot;SSWS_TOKEN\&quot;&#x60; or &#x60;filter&#x3D;principalType eq \&quot;OAUTH_CLIENT\&quot;&#x60;. | 
 **after** | **string** | The cursor to use for pagination. It&#39;s an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of items to return in a single response page. | [default to 20]

### Return type

[**[]PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePrincipalRateLimitEntity

> PrincipalRateLimitEntity ReplacePrincipalRateLimitEntity(ctx, principalRateLimitId).Entity(entity).Execute()

Replace a principal rate limit



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
	principalRateLimitId := "0oacamvryxiyMqgiY1d7" // string | ID of the principal rate limit
	entity := *openapiclient.NewPrincipalRateLimitEntity("PrincipalId_example", "PrincipalType_example") // PrincipalRateLimitEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalRateLimitAPI.ReplacePrincipalRateLimitEntity(context.Background(), principalRateLimitId).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalRateLimitAPI.ReplacePrincipalRateLimitEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplacePrincipalRateLimitEntity`: PrincipalRateLimitEntity
	fmt.Fprintf(os.Stdout, "Response from `PrincipalRateLimitAPI.ReplacePrincipalRateLimitEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalRateLimitId** | **string** | ID of the principal rate limit | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePrincipalRateLimitEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entity** | [**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md) |  | 

### Return type

[**PrincipalRateLimitEntity**](PrincipalRateLimitEntity.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


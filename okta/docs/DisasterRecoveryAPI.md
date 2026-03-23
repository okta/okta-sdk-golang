# \DisasterRecoveryAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDRStatus**](DisasterRecoveryAPI.md#GetDRStatus) | **Get** /api/v1/dr/status | Retrieve the disaster recovery status for all domains
[**GetDRStatusForDomain**](DisasterRecoveryAPI.md#GetDRStatusForDomain) | **Get** /api/v1/dr/status/{domain} | Retrieve the disaster recovery status for a domain
[**StartOrgFailback**](DisasterRecoveryAPI.md#StartOrgFailback) | **Post** /api/v1/dr/failback | Start the failback of your org
[**StartOrgFailover**](DisasterRecoveryAPI.md#StartOrgFailover) | **Post** /api/v1/dr/failover | Start the failover of your org



## GetDRStatus

> GetDRStatus200Response GetDRStatus(ctx).Execute()

Retrieve the disaster recovery status for all domains



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
	resp, r, err := apiClient.DisasterRecoveryAPI.GetDRStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.GetDRStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDRStatus`: GetDRStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.GetDRStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDRStatusRequest struct via the builder pattern


### Return type

[**GetDRStatus200Response**](GetDRStatus200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDRStatusForDomain

> GetDRStatus200Response GetDRStatusForDomain(ctx, domain).Execute()

Retrieve the disaster recovery status for a domain



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
	domain := "domain_example" // string | The Okta domain name of your org or one of your custom domains

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.GetDRStatusForDomain(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.GetDRStatusForDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDRStatusForDomain`: GetDRStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.GetDRStatusForDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The Okta domain name of your org or one of your custom domains | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDRStatusForDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDRStatus200Response**](GetDRStatus200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOrgFailback

> StartOrgFailback200Response StartOrgFailback(ctx).StartOrgFailbackRequest(startOrgFailbackRequest).Execute()

Start the failback of your org



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
	startOrgFailbackRequest := openapiclient.startOrgFailback_request{FailbackRequestSchema: openapiclient.NewFailbackRequestSchema()} // StartOrgFailbackRequest | The request body is optional. You can specify a domain to failback, an empty object (`{}`), or no payload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.StartOrgFailback(context.Background()).StartOrgFailbackRequest(startOrgFailbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.StartOrgFailback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOrgFailback`: StartOrgFailback200Response
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.StartOrgFailback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartOrgFailbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startOrgFailbackRequest** | [**StartOrgFailbackRequest**](StartOrgFailbackRequest.md) | The request body is optional. You can specify a domain to failback, an empty object (&#x60;{}&#x60;), or no payload. | 

### Return type

[**StartOrgFailback200Response**](StartOrgFailback200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOrgFailover

> StartOrgFailover200Response StartOrgFailover(ctx).StartOrgFailoverRequest(startOrgFailoverRequest).Execute()

Start the failover of your org



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
	startOrgFailoverRequest := openapiclient.startOrgFailover_request{FailoverRequestSchema: openapiclient.NewFailoverRequestSchema()} // StartOrgFailoverRequest | The request body is optional. You can specify a domain to failover, an empty object (`{}`), or no payload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.StartOrgFailover(context.Background()).StartOrgFailoverRequest(startOrgFailoverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.StartOrgFailover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOrgFailover`: StartOrgFailover200Response
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.StartOrgFailover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartOrgFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startOrgFailoverRequest** | [**StartOrgFailoverRequest**](StartOrgFailoverRequest.md) | The request body is optional. You can specify a domain to failover, an empty object (&#x60;{}&#x60;), or no payload. | 

### Return type

[**StartOrgFailover200Response**](StartOrgFailover200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


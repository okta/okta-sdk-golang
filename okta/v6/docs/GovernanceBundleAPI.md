# \GovernanceBundleAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGovernanceBundle**](GovernanceBundleAPI.md#CreateGovernanceBundle) | **Post** /api/v1/iam/governance/bundles | Create a governance bundle for the Admin Console in RAMP
[**DeleteGovernanceBundle**](GovernanceBundleAPI.md#DeleteGovernanceBundle) | **Delete** /api/v1/iam/governance/bundles/{bundleId} | Delete a governance bundle from RAMP
[**GetGovernanceBundle**](GovernanceBundleAPI.md#GetGovernanceBundle) | **Get** /api/v1/iam/governance/bundles/{bundleId} | Retrieve a governance bundle from RAMP
[**GetOptInStatus**](GovernanceBundleAPI.md#GetOptInStatus) | **Get** /api/v1/iam/governance/optIn | Retrieve the opt-in status from RAMP
[**ListBundleEntitlementValues**](GovernanceBundleAPI.md#ListBundleEntitlementValues) | **Get** /api/v1/iam/governance/bundles/{bundleId}/entitlements/{entitlementId}/values | List all entitlement values for a bundle entitlement
[**ListBundleEntitlements**](GovernanceBundleAPI.md#ListBundleEntitlements) | **Get** /api/v1/iam/governance/bundles/{bundleId}/entitlements | List all entitlements for a governance bundle
[**ListGovernanceBundles**](GovernanceBundleAPI.md#ListGovernanceBundles) | **Get** /api/v1/iam/governance/bundles | List all governance bundles for the Admin Console
[**OptIn**](GovernanceBundleAPI.md#OptIn) | **Post** /api/v1/iam/governance/optIn | Opt in the Admin Console to RAMP
[**OptOut**](GovernanceBundleAPI.md#OptOut) | **Post** /api/v1/iam/governance/optOut | Opt out the Admin Console from RAMP
[**ReplaceGovernanceBundle**](GovernanceBundleAPI.md#ReplaceGovernanceBundle) | **Put** /api/v1/iam/governance/bundles/{bundleId} | Replace a governance bundle in RAMP



## CreateGovernanceBundle

> GovernanceBundle CreateGovernanceBundle(ctx).GovernanceBundleCreateRequest(governanceBundleCreateRequest).Execute()

Create a governance bundle for the Admin Console in RAMP



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
	governanceBundleCreateRequest := *openapiclient.NewGovernanceBundleCreateRequest() // GovernanceBundleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.CreateGovernanceBundle(context.Background()).GovernanceBundleCreateRequest(governanceBundleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.CreateGovernanceBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGovernanceBundle`: GovernanceBundle
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.CreateGovernanceBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGovernanceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **governanceBundleCreateRequest** | [**GovernanceBundleCreateRequest**](GovernanceBundleCreateRequest.md) |  | 

### Return type

[**GovernanceBundle**](GovernanceBundle.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGovernanceBundle

> DeleteGovernanceBundle(ctx, bundleId).Execute()

Delete a governance bundle from RAMP



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
	bundleId := "enbllojq9J9J105DL1d6" // string | The `id` of a bundle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GovernanceBundleAPI.DeleteGovernanceBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.DeleteGovernanceBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The &#x60;id&#x60; of a bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGovernanceBundleRequest struct via the builder pattern


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


## GetGovernanceBundle

> GovernanceBundle GetGovernanceBundle(ctx, bundleId).Execute()

Retrieve a governance bundle from RAMP



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
	bundleId := "enbllojq9J9J105DL1d6" // string | The `id` of a bundle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.GetGovernanceBundle(context.Background(), bundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.GetGovernanceBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGovernanceBundle`: GovernanceBundle
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.GetGovernanceBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The &#x60;id&#x60; of a bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGovernanceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GovernanceBundle**](GovernanceBundle.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptInStatus

> OptInStatusResponse GetOptInStatus(ctx).Execute()

Retrieve the opt-in status from RAMP



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
	resp, r, err := apiClient.GovernanceBundleAPI.GetOptInStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.GetOptInStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptInStatus`: OptInStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.GetOptInStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptInStatusRequest struct via the builder pattern


### Return type

[**OptInStatusResponse**](OptInStatusResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundleEntitlementValues

> EntitlementValuesResponse ListBundleEntitlementValues(ctx, bundleId, entitlementId).After(after).Limit(limit).Execute()

List all entitlement values for a bundle entitlement



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
	bundleId := "enbllojq9J9J105DL1d6" // string | The `id` of a bundle
	entitlementId := "ent4rg7fltWSgrlDT8g6" // string | The `id` of a bundle entitlement
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.ListBundleEntitlementValues(context.Background(), bundleId, entitlementId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.ListBundleEntitlementValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBundleEntitlementValues`: EntitlementValuesResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.ListBundleEntitlementValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The &#x60;id&#x60; of a bundle | 
**entitlementId** | **string** | The &#x60;id&#x60; of a bundle entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBundleEntitlementValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**EntitlementValuesResponse**](EntitlementValuesResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundleEntitlements

> BundleEntitlementsResponse ListBundleEntitlements(ctx, bundleId).After(after).Limit(limit).Execute()

List all entitlements for a governance bundle



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
	bundleId := "enbllojq9J9J105DL1d6" // string | The `id` of a bundle
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.ListBundleEntitlements(context.Background(), bundleId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.ListBundleEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBundleEntitlements`: BundleEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.ListBundleEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The &#x60;id&#x60; of a bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBundleEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**BundleEntitlementsResponse**](BundleEntitlementsResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGovernanceBundles

> GovernanceBundlesResponse ListGovernanceBundles(ctx).After(after).Limit(limit).Execute()

List all governance bundles for the Admin Console



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.ListGovernanceBundles(context.Background()).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.ListGovernanceBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGovernanceBundles`: GovernanceBundlesResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.ListGovernanceBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGovernanceBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**GovernanceBundlesResponse**](GovernanceBundlesResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptIn

> OptInStatusResponse OptIn(ctx).Execute()

Opt in the Admin Console to RAMP



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
	resp, r, err := apiClient.GovernanceBundleAPI.OptIn(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.OptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptIn`: OptInStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.OptIn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptInRequest struct via the builder pattern


### Return type

[**OptInStatusResponse**](OptInStatusResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOut

> OptInStatusResponse OptOut(ctx).Execute()

Opt out the Admin Console from RAMP



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
	resp, r, err := apiClient.GovernanceBundleAPI.OptOut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.OptOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptOut`: OptInStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.OptOut`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptOutRequest struct via the builder pattern


### Return type

[**OptInStatusResponse**](OptInStatusResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceGovernanceBundle

> GovernanceBundle ReplaceGovernanceBundle(ctx, bundleId).GovernanceBundleUpdateRequest(governanceBundleUpdateRequest).Execute()

Replace a governance bundle in RAMP



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
	bundleId := "enbllojq9J9J105DL1d6" // string | The `id` of a bundle
	governanceBundleUpdateRequest := *openapiclient.NewGovernanceBundleUpdateRequest() // GovernanceBundleUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceBundleAPI.ReplaceGovernanceBundle(context.Background(), bundleId).GovernanceBundleUpdateRequest(governanceBundleUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceBundleAPI.ReplaceGovernanceBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceGovernanceBundle`: GovernanceBundle
	fmt.Fprintf(os.Stdout, "Response from `GovernanceBundleAPI.ReplaceGovernanceBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | **string** | The &#x60;id&#x60; of a bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGovernanceBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **governanceBundleUpdateRequest** | [**GovernanceBundleUpdateRequest**](GovernanceBundleUpdateRequest.md) |  | 

### Return type

[**GovernanceBundle**](GovernanceBundle.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


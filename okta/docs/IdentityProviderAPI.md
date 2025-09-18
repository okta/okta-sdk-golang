# \IdentityProviderAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateIdentityProvider**](IdentityProviderAPI.md#ActivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/activate | Activate an IdP
[**CreateIdentityProvider**](IdentityProviderAPI.md#CreateIdentityProvider) | **Post** /api/v1/idps | Create an IdP
[**DeactivateIdentityProvider**](IdentityProviderAPI.md#DeactivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/deactivate | Deactivate an IdP
[**DeleteIdentityProvider**](IdentityProviderAPI.md#DeleteIdentityProvider) | **Delete** /api/v1/idps/{idpId} | Delete an IdP
[**GetIdentityProvider**](IdentityProviderAPI.md#GetIdentityProvider) | **Get** /api/v1/idps/{idpId} | Retrieve an IdP
[**ListIdentityProviders**](IdentityProviderAPI.md#ListIdentityProviders) | **Get** /api/v1/idps | List all IdPs
[**ReplaceIdentityProvider**](IdentityProviderAPI.md#ReplaceIdentityProvider) | **Put** /api/v1/idps/{idpId} | Replace an IdP



## ActivateIdentityProvider

> IdentityProvider ActivateIdentityProvider(ctx, idpId).Execute()

Activate an IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.ActivateIdentityProvider(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ActivateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateIdentityProvider`: IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ActivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentityProvider

> IdentityProvider CreateIdentityProvider(ctx).IdentityProvider(identityProvider).Execute()

Create an IdP



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
	identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | IdP settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.CreateIdentityProvider(context.Background()).IdentityProvider(identityProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CreateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentityProvider`: IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) | IdP settings | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateIdentityProvider

> IdentityProvider DeactivateIdentityProvider(ctx, idpId).Execute()

Deactivate an IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.DeactivateIdentityProvider(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeactivateIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateIdentityProvider`: IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.DeactivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, idpId).Execute()

Delete an IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderAPI.DeleteIdentityProvider(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeleteIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


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


## GetIdentityProvider

> IdentityProvider GetIdentityProvider(ctx, idpId).Execute()

Retrieve an IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProvider(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProvider`: IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviders

> []IdentityProvider ListIdentityProviders(ctx).Q(q).After(after).Limit(limit).Type_(type_).Execute()

List all IdPs



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
	q := "Example SAML" // string | Searches the `name` property of IdPs for matching value (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
	type_ := "type__example" // string | Filters IdPs by `type` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.ListIdentityProviders(context.Background()).Q(q).After(after).Limit(limit).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListIdentityProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityProviders`: []IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Searches the &#x60;name&#x60; property of IdPs for matching value | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **type_** | **string** | Filters IdPs by &#x60;type&#x60; | 

### Return type

[**[]IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIdentityProvider

> IdentityProvider ReplaceIdentityProvider(ctx, idpId).IdentityProvider(identityProvider).Execute()

Replace an IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | Updated configuration for the IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderAPI.ReplaceIdentityProvider(context.Background(), idpId).IdentityProvider(identityProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ReplaceIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceIdentityProvider`: IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ReplaceIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) | Updated configuration for the IdP | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


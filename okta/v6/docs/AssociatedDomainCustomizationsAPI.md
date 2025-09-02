# \AssociatedDomainCustomizationsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllWellKnownURIs**](AssociatedDomainCustomizationsAPI.md#GetAllWellKnownURIs) | **Get** /api/v1/brands/{brandId}/well-known-uris | Retrieve all the well-known URIs
[**GetAppleAppSiteAssociationWellKnownURI**](AssociatedDomainCustomizationsAPI.md#GetAppleAppSiteAssociationWellKnownURI) | **Get** /.well-known/apple-app-site-association | Retrieve the customized apple-app-site-association URI content
[**GetAssetLinksWellKnownURI**](AssociatedDomainCustomizationsAPI.md#GetAssetLinksWellKnownURI) | **Get** /.well-known/assetlinks.json | Retrieve the customized assetlinks.json URI content
[**GetBrandWellKnownURI**](AssociatedDomainCustomizationsAPI.md#GetBrandWellKnownURI) | **Get** /api/v1/brands/{brandId}/well-known-uris/{path}/customized | Retrieve the customized content of the specified well-known URI
[**GetRootBrandWellKnownURI**](AssociatedDomainCustomizationsAPI.md#GetRootBrandWellKnownURI) | **Get** /api/v1/brands/{brandId}/well-known-uris/{path} | Retrieve the well-known URI of a specific brand
[**GetWebAuthnWellKnownURI**](AssociatedDomainCustomizationsAPI.md#GetWebAuthnWellKnownURI) | **Get** /.well-known/webauthn | Retrieve the customized webauthn URI content
[**ReplaceBrandWellKnownURI**](AssociatedDomainCustomizationsAPI.md#ReplaceBrandWellKnownURI) | **Put** /api/v1/brands/{brandId}/well-known-uris/{path}/customized | Replace the customized well-known URI of the specific path



## GetAllWellKnownURIs

> WellKnownURIsRoot GetAllWellKnownURIs(ctx, brandId).Expand(expand).Execute()

Retrieve all the well-known URIs



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
	brandId := "brandId_example" // string | The ID of the brand
	expand := []string{"Expand_example"} // []string | Specifies additional metadata to include in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetAllWellKnownURIs(context.Background(), brandId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetAllWellKnownURIs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWellKnownURIs`: WellKnownURIsRoot
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetAllWellKnownURIs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWellKnownURIsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to include in the response | 

### Return type

[**WellKnownURIsRoot**](WellKnownURIsRoot.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppleAppSiteAssociationWellKnownURI

> map[string]interface{} GetAppleAppSiteAssociationWellKnownURI(ctx).Execute()

Retrieve the customized apple-app-site-association URI content



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
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetAppleAppSiteAssociationWellKnownURI(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetAppleAppSiteAssociationWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppleAppSiteAssociationWellKnownURI`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetAppleAppSiteAssociationWellKnownURI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppleAppSiteAssociationWellKnownURIRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetLinksWellKnownURI

> []string GetAssetLinksWellKnownURI(ctx).Execute()

Retrieve the customized assetlinks.json URI content



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
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetAssetLinksWellKnownURI(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetAssetLinksWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetLinksWellKnownURI`: []string
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetAssetLinksWellKnownURI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetLinksWellKnownURIRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandWellKnownURI

> WellKnownURIObjectResponse GetBrandWellKnownURI(ctx, brandId, path).Execute()

Retrieve the customized content of the specified well-known URI



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
	brandId := "brandId_example" // string | The ID of the brand
	path := "path_example" // string | The path of the well-known URI

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetBrandWellKnownURI(context.Background(), brandId, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetBrandWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandWellKnownURI`: WellKnownURIObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetBrandWellKnownURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**path** | **string** | The path of the well-known URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandWellKnownURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WellKnownURIObjectResponse**](WellKnownURIObjectResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootBrandWellKnownURI

> WellKnownURIObjectResponse GetRootBrandWellKnownURI(ctx, brandId, path).Expand(expand).Execute()

Retrieve the well-known URI of a specific brand



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
	brandId := "brandId_example" // string | The ID of the brand
	path := "path_example" // string | The path of the well-known URI
	expand := []string{"Expand_example"} // []string | Specifies additional metadata to include in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetRootBrandWellKnownURI(context.Background(), brandId, path).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetRootBrandWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootBrandWellKnownURI`: WellKnownURIObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetRootBrandWellKnownURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**path** | **string** | The path of the well-known URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootBrandWellKnownURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Specifies additional metadata to include in the response | 

### Return type

[**WellKnownURIObjectResponse**](WellKnownURIObjectResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebAuthnWellKnownURI

> map[string]interface{} GetWebAuthnWellKnownURI(ctx).Execute()

Retrieve the customized webauthn URI content



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
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.GetWebAuthnWellKnownURI(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.GetWebAuthnWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebAuthnWellKnownURI`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.GetWebAuthnWellKnownURI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebAuthnWellKnownURIRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBrandWellKnownURI

> WellKnownURIObjectResponse ReplaceBrandWellKnownURI(ctx, brandId, path).WellKnownURIRequest(wellKnownURIRequest).Execute()

Replace the customized well-known URI of the specific path



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
	brandId := "brandId_example" // string | The ID of the brand
	path := "path_example" // string | The path of the well-known URI
	wellKnownURIRequest := *openapiclient.NewWellKnownURIRequest(map[string]interface{}(123)) // WellKnownURIRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociatedDomainCustomizationsAPI.ReplaceBrandWellKnownURI(context.Background(), brandId, path).WellKnownURIRequest(wellKnownURIRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociatedDomainCustomizationsAPI.ReplaceBrandWellKnownURI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceBrandWellKnownURI`: WellKnownURIObjectResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociatedDomainCustomizationsAPI.ReplaceBrandWellKnownURI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**path** | **string** | The path of the well-known URI | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBrandWellKnownURIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wellKnownURIRequest** | [**WellKnownURIRequest**](WellKnownURIRequest.md) |  | 

### Return type

[**WellKnownURIObjectResponse**](WellKnownURIObjectResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


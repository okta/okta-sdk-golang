# \BrandsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrand**](BrandsAPI.md#CreateBrand) | **Post** /api/v1/brands | Create a Brand
[**DeleteBrand**](BrandsAPI.md#DeleteBrand) | **Delete** /api/v1/brands/{brandId} | Delete a brand
[**GetBrand**](BrandsAPI.md#GetBrand) | **Get** /api/v1/brands/{brandId} | Retrieve a Brand
[**ListBrandDomains**](BrandsAPI.md#ListBrandDomains) | **Get** /api/v1/brands/{brandId}/domains | List all Domains associated with a Brand
[**ListBrands**](BrandsAPI.md#ListBrands) | **Get** /api/v1/brands | List all Brands
[**ReplaceBrand**](BrandsAPI.md#ReplaceBrand) | **Put** /api/v1/brands/{brandId} | Replace a Brand



## CreateBrand

> Brand CreateBrand(ctx).CreateBrandRequest(createBrandRequest).Execute()

Create a Brand



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
    createBrandRequest := *openapiclient.NewCreateBrandRequest("Name_example") // CreateBrandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandsAPI.CreateBrand(context.Background()).CreateBrandRequest(createBrandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.CreateBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.CreateBrand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBrandRequest** | [**CreateBrandRequest**](CreateBrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrand

> DeleteBrand(ctx, brandId).Execute()

Delete a brand



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BrandsAPI.DeleteBrand(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.DeleteBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


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


## GetBrand

> BrandWithEmbedded GetBrand(ctx, brandId).Expand(expand).Execute()

Retrieve a Brand



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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandsAPI.GetBrand(context.Background(), brandId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: BrandWithEmbedded
    fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**BrandWithEmbedded**](BrandWithEmbedded.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandDomains

> []DomainResponse ListBrandDomains(ctx, brandId).Execute()

List all Domains associated with a Brand



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandsAPI.ListBrandDomains(context.Background(), brandId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.ListBrandDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandDomains`: []DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.ListBrandDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainResponse**](DomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrands

> []BrandWithEmbedded ListBrands(ctx).Expand(expand).After(after).Limit(limit).Q(q).Execute()

List all Brands



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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
    q := "q_example" // string | Searches the records for matching value (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandsAPI.ListBrands(context.Background()).Expand(expand).After(after).Limit(limit).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.ListBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrands`: []BrandWithEmbedded
    fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.ListBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **q** | **string** | Searches the records for matching value | 

### Return type

[**[]BrandWithEmbedded**](BrandWithEmbedded.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBrand

> Brand ReplaceBrand(ctx, brandId).Brand(brand).Execute()

Replace a Brand



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
    brand := *openapiclient.NewBrandRequest("Name_example") // BrandRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandsAPI.ReplaceBrand(context.Background(), brandId).Brand(brand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.ReplaceBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBrand`: Brand
    fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.ReplaceBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brand** | [**BrandRequest**](BrandRequest.md) |  | 

### Return type

[**Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


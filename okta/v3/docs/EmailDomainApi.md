# \EmailDomainApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailDomain**](EmailDomainApi.md#CreateEmailDomain) | **Post** /api/v1/email-domains | Create an Email Domain
[**DeleteEmailDomain**](EmailDomainApi.md#DeleteEmailDomain) | **Delete** /api/v1/email-domains/{emailDomainId} | Delete an Email Domain
[**GetEmailDomain**](EmailDomainApi.md#GetEmailDomain) | **Get** /api/v1/email-domains/{emailDomainId} | Retrieve a Email Domain
[**ListEmailDomainBrands**](EmailDomainApi.md#ListEmailDomainBrands) | **Get** /api/v1/email-domains/{emailDomainId}/brands | List all brands linked to an email domain
[**ListEmailDomains**](EmailDomainApi.md#ListEmailDomains) | **Get** /api/v1/email-domains | List all email domains
[**ReplaceEmailDomain**](EmailDomainApi.md#ReplaceEmailDomain) | **Put** /api/v1/email-domains/{emailDomainId} | Replace an Email Domain
[**VerifyEmailDomain**](EmailDomainApi.md#VerifyEmailDomain) | **Post** /api/v1/email-domains/{emailDomainId}/verify | Verify an Email Domain



## CreateEmailDomain

> EmailDomainResponse CreateEmailDomain(ctx).EmailDomain(emailDomain).Execute()

Create an Email Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomain := *openapiclient.NewEmailDomain("Domain_example", "DisplayName_example", "UserName_example") // EmailDomain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.CreateEmailDomain(context.Background()).EmailDomain(emailDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.CreateEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailDomain`: EmailDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.CreateEmailDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailDomain** | [**EmailDomain**](EmailDomain.md) |  | 

### Return type

[**EmailDomainResponse**](EmailDomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailDomain

> DeleteEmailDomain(ctx, emailDomainId).Execute()

Delete an Email Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.DeleteEmailDomain(context.Background(), emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.DeleteEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailDomainRequest struct via the builder pattern


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


## GetEmailDomain

> EmailDomainResponse GetEmailDomain(ctx, emailDomainId).Execute()

Retrieve a Email Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.GetEmailDomain(context.Background(), emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.GetEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDomain`: EmailDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.GetEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailDomainResponse**](EmailDomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailDomainBrands

> []Brand ListEmailDomainBrands(ctx, emailDomainId).Execute()

List all brands linked to an email domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.ListEmailDomainBrands(context.Background(), emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.ListEmailDomainBrands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailDomainBrands`: []Brand
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.ListEmailDomainBrands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailDomainBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Brand**](Brand.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailDomains

> EmailDomainListResponse ListEmailDomains(ctx).Execute()

List all email domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.ListEmailDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.ListEmailDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailDomains`: EmailDomainListResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.ListEmailDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailDomainsRequest struct via the builder pattern


### Return type

[**EmailDomainListResponse**](EmailDomainListResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailDomain

> EmailDomainResponse ReplaceEmailDomain(ctx, emailDomainId).UpdateEmailDomain(updateEmailDomain).Execute()

Replace an Email Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomainId := "emailDomainId_example" // string | 
    updateEmailDomain := *openapiclient.NewUpdateEmailDomain("DisplayName_example", "UserName_example") // UpdateEmailDomain | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.ReplaceEmailDomain(context.Background(), emailDomainId).UpdateEmailDomain(updateEmailDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.ReplaceEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEmailDomain`: EmailDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.ReplaceEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEmailDomain** | [**UpdateEmailDomain**](UpdateEmailDomain.md) |  | 

### Return type

[**EmailDomainResponse**](EmailDomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailDomain

> EmailDomainResponse VerifyEmailDomain(ctx, emailDomainId).Execute()

Verify an Email Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    emailDomainId := "emailDomainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.VerifyEmailDomain(context.Background(), emailDomainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.VerifyEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEmailDomain`: EmailDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.VerifyEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailDomainResponse**](EmailDomainResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


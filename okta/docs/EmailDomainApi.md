# \EmailDomainApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailDomain**](EmailDomainApi.md#CreateEmailDomain) | **Post** /api/v1/email-domains | Create an Email Domain
[**DeleteEmailDomain**](EmailDomainApi.md#DeleteEmailDomain) | **Delete** /api/v1/email-domains/{emailDomainId} | Delete an Email Domain
[**GetEmailDomain**](EmailDomainApi.md#GetEmailDomain) | **Get** /api/v1/email-domains/{emailDomainId} | Retrieve an Email Domain
[**ListEmailDomains**](EmailDomainApi.md#ListEmailDomains) | **Get** /api/v1/email-domains | List all Email Domains
[**ReplaceEmailDomain**](EmailDomainApi.md#ReplaceEmailDomain) | **Put** /api/v1/email-domains/{emailDomainId} | Replace an Email Domain
[**VerifyEmailDomain**](EmailDomainApi.md#VerifyEmailDomain) | **Post** /api/v1/email-domains/{emailDomainId}/verify | Verify an Email Domain



## CreateEmailDomain

> EmailDomainResponse CreateEmailDomain(ctx).EmailDomain(emailDomain).Expand(expand).Execute()

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
    emailDomain := *openapiclient.NewEmailDomain("BrandId_example", "Domain_example", "DisplayName_example", "UserName_example") // EmailDomain | 
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.CreateEmailDomain(context.Background()).EmailDomain(emailDomain).Expand(expand).Execute()
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
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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

> DeleteEmailDomain(ctx, emailDomainId).Expand(expand).Execute()

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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.DeleteEmailDomain(context.Background(), emailDomainId).Expand(expand).Execute()
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

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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

> EmailDomainResponseWithEmbedded GetEmailDomain(ctx, emailDomainId).Expand(expand).Execute()

Retrieve an Email Domain



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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.GetEmailDomain(context.Background(), emailDomainId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.GetEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDomain`: EmailDomainResponseWithEmbedded
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

 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**EmailDomainResponseWithEmbedded**](EmailDomainResponseWithEmbedded.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailDomains

> []EmailDomainResponseWithEmbedded ListEmailDomains(ctx).Expand(expand).Execute()

List all Email Domains



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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.ListEmailDomains(context.Background()).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailDomainApi.ListEmailDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailDomains`: []EmailDomainResponseWithEmbedded
    fmt.Fprintf(os.Stdout, "Response from `EmailDomainApi.ListEmailDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**[]EmailDomainResponseWithEmbedded**](EmailDomainResponseWithEmbedded.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailDomain

> EmailDomainResponse ReplaceEmailDomain(ctx, emailDomainId).UpdateEmailDomain(updateEmailDomain).Expand(expand).Execute()

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
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailDomainApi.ReplaceEmailDomain(context.Background(), emailDomainId).UpdateEmailDomain(updateEmailDomain).Expand(expand).Execute()
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
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

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


# \ApplicationTokensAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOAuth2TokenForApplication**](ApplicationTokensAPI.md#GetOAuth2TokenForApplication) | **Get** /api/v1/apps/{appId}/tokens/{tokenId} | Retrieve an application Token
[**ListOAuth2TokensForApplication**](ApplicationTokensAPI.md#ListOAuth2TokensForApplication) | **Get** /api/v1/apps/{appId}/tokens | List all application refresh Tokens
[**RevokeOAuth2TokenForApplication**](ApplicationTokensAPI.md#RevokeOAuth2TokenForApplication) | **Delete** /api/v1/apps/{appId}/tokens/{tokenId} | Revoke an application Token
[**RevokeOAuth2TokensForApplication**](ApplicationTokensAPI.md#RevokeOAuth2TokensForApplication) | **Delete** /api/v1/apps/{appId}/tokens | Revoke all application Tokens



## GetOAuth2TokenForApplication

> OAuth2RefreshToken GetOAuth2TokenForApplication(ctx, appId, tokenId).Expand(expand).Execute()

Retrieve an application Token



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
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
    expand := "scope" // string | An optional parameter to return scope details in the `_embedded` property. Valid value: `scope` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTokensAPI.GetOAuth2TokenForApplication(context.Background(), appId, tokenId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensAPI.GetOAuth2TokenForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2TokenForApplication`: OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTokensAPI.GetOAuth2TokenForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2TokenForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | An optional parameter to return scope details in the &#x60;_embedded&#x60; property. Valid value: &#x60;scope&#x60; | 

### Return type

[**OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2TokensForApplication

> []OAuth2RefreshToken ListOAuth2TokensForApplication(ctx, appId).Expand(expand).After(after).Limit(limit).Execute()

List all application refresh Tokens



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
    expand := "scope" // string | An optional parameter to return scope details in the `_embedded` property. Valid value: `scope` (optional)
    after := "16275000448691" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTokensAPI.ListOAuth2TokensForApplication(context.Background(), appId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensAPI.ListOAuth2TokensForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2TokensForApplication`: []OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTokensAPI.ListOAuth2TokensForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2TokensForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | An optional parameter to return scope details in the &#x60;_embedded&#x60; property. Valid value: &#x60;scope&#x60; | 
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the next link relationship. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**[]OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOAuth2TokenForApplication

> RevokeOAuth2TokenForApplication(ctx, appId, tokenId).Execute()

Revoke an application Token



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
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationTokensAPI.RevokeOAuth2TokenForApplication(context.Background(), appId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensAPI.RevokeOAuth2TokenForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOAuth2TokenForApplicationRequest struct via the builder pattern


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


## RevokeOAuth2TokensForApplication

> RevokeOAuth2TokensForApplication(ctx, appId).Execute()

Revoke all application Tokens



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
    r, err := apiClient.ApplicationTokensAPI.RevokeOAuth2TokensForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensAPI.RevokeOAuth2TokensForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOAuth2TokensForApplicationRequest struct via the builder pattern


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


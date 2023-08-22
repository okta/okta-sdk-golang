# \ApplicationTokensApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOAuth2TokenForApplication**](ApplicationTokensApi.md#GetOAuth2TokenForApplication) | **Get** /api/v1/apps/{appId}/tokens/{tokenId} | Retrieve an OAuth 2.0 Token
[**ListOAuth2TokensForApplication**](ApplicationTokensApi.md#ListOAuth2TokensForApplication) | **Get** /api/v1/apps/{appId}/tokens | List all OAuth 2.0 Tokens
[**RevokeOAuth2TokenForApplication**](ApplicationTokensApi.md#RevokeOAuth2TokenForApplication) | **Delete** /api/v1/apps/{appId}/tokens/{tokenId} | Revoke an OAuth 2.0 Token
[**RevokeOAuth2TokensForApplication**](ApplicationTokensApi.md#RevokeOAuth2TokensForApplication) | **Delete** /api/v1/apps/{appId}/tokens | Revoke all OAuth 2.0 Tokens



## GetOAuth2TokenForApplication

> OAuth2Token GetOAuth2TokenForApplication(ctx, appId, tokenId).Expand(expand).Execute()

Retrieve an OAuth 2.0 Token



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTokensApi.GetOAuth2TokenForApplication(context.Background(), appId, tokenId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensApi.GetOAuth2TokenForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2TokenForApplication`: OAuth2Token
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTokensApi.GetOAuth2TokenForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2TokenForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2TokensForApplication

> []OAuth2Token ListOAuth2TokensForApplication(ctx, appId).Expand(expand).After(after).Limit(limit).Execute()

List all OAuth 2.0 Tokens



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTokensApi.ListOAuth2TokensForApplication(context.Background(), appId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensApi.ListOAuth2TokensForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2TokensForApplication`: []OAuth2Token
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTokensApi.ListOAuth2TokensForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2TokensForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**[]OAuth2Token**](OAuth2Token.md)

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

Revoke an OAuth 2.0 Token



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationTokensApi.RevokeOAuth2TokenForApplication(context.Background(), appId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensApi.RevokeOAuth2TokenForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 
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

Revoke all OAuth 2.0 Tokens



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | ID of the Application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationTokensApi.RevokeOAuth2TokensForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTokensApi.RevokeOAuth2TokensForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

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


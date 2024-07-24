# \AuthorizationServerScopesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuth2Scope**](AuthorizationServerScopesAPI.md#CreateOAuth2Scope) | **Post** /api/v1/authorizationServers/{authServerId}/scopes | Create a Custom Token Scope
[**DeleteOAuth2Scope**](AuthorizationServerScopesAPI.md#DeleteOAuth2Scope) | **Delete** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Delete a Custom Token Scope
[**GetOAuth2Scope**](AuthorizationServerScopesAPI.md#GetOAuth2Scope) | **Get** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Retrieve a Custom Token Scope
[**ListOAuth2Scopes**](AuthorizationServerScopesAPI.md#ListOAuth2Scopes) | **Get** /api/v1/authorizationServers/{authServerId}/scopes | List all Custom Token Scopes
[**ReplaceOAuth2Scope**](AuthorizationServerScopesAPI.md#ReplaceOAuth2Scope) | **Put** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Replace a Custom Token Scope



## CreateOAuth2Scope

> OAuth2Scope CreateOAuth2Scope(ctx, authServerId).OAuth2Scope(oAuth2Scope).Execute()

Create a Custom Token Scope



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
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerScopesAPI.CreateOAuth2Scope(context.Background(), authServerId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerScopesAPI.CreateOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerScopesAPI.CreateOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2Scope** | [**OAuth2Scope**](OAuth2Scope.md) |  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuth2Scope

> DeleteOAuth2Scope(ctx, authServerId, scopeId).Execute()

Delete a Custom Token Scope



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
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerScopesAPI.DeleteOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerScopesAPI.DeleteOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ScopeRequest struct via the builder pattern


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


## GetOAuth2Scope

> OAuth2Scope GetOAuth2Scope(ctx, authServerId, scopeId).Execute()

Retrieve a Custom Token Scope



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
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerScopesAPI.GetOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerScopesAPI.GetOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerScopesAPI.GetOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2Scopes

> []OAuth2Scope ListOAuth2Scopes(ctx, authServerId).Q(q).Filter(filter).After(after).Limit(limit).Execute()

List all Custom Token Scopes



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
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    q := "q_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerScopesAPI.ListOAuth2Scopes(context.Background(), authServerId).Q(q).Filter(filter).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerScopesAPI.ListOAuth2Scopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2Scopes`: []OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerScopesAPI.ListOAuth2Scopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 
 **filter** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

### Return type

[**[]OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOAuth2Scope

> OAuth2Scope ReplaceOAuth2Scope(ctx, authServerId, scopeId).OAuth2Scope(oAuth2Scope).Execute()

Replace a Custom Token Scope



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
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerScopesAPI.ReplaceOAuth2Scope(context.Background(), authServerId, scopeId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerScopesAPI.ReplaceOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerScopesAPI.ReplaceOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oAuth2Scope** | [**OAuth2Scope**](OAuth2Scope.md) |  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


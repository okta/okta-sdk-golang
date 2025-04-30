# \AuthorizationServerClaimsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuth2Claim**](AuthorizationServerClaimsAPI.md#CreateOAuth2Claim) | **Post** /api/v1/authorizationServers/{authServerId}/claims | Create a custom token Claim
[**DeleteOAuth2Claim**](AuthorizationServerClaimsAPI.md#DeleteOAuth2Claim) | **Delete** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Delete a custom token Claim
[**GetOAuth2Claim**](AuthorizationServerClaimsAPI.md#GetOAuth2Claim) | **Get** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Retrieve a custom token Claim
[**ListOAuth2Claims**](AuthorizationServerClaimsAPI.md#ListOAuth2Claims) | **Get** /api/v1/authorizationServers/{authServerId}/claims | List all custom token Claims
[**ReplaceOAuth2Claim**](AuthorizationServerClaimsAPI.md#ReplaceOAuth2Claim) | **Put** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Replace a custom token Claim



## CreateOAuth2Claim

> OAuth2Claim CreateOAuth2Claim(ctx, authServerId).OAuth2Claim(oAuth2Claim).Execute()

Create a custom token Claim



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
    oAuth2Claim := *openapiclient.NewOAuth2Claim() // OAuth2Claim | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClaimsAPI.CreateOAuth2Claim(context.Background(), authServerId).OAuth2Claim(oAuth2Claim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClaimsAPI.CreateOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClaimsAPI.CreateOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2Claim** | [**OAuth2Claim**](OAuth2Claim.md) |  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuth2Claim

> DeleteOAuth2Claim(ctx, authServerId, claimId).Execute()

Delete a custom token Claim



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
    claimId := "hNJ3Uk76xLagWkGx5W3N" // string | `id` of Claim

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerClaimsAPI.DeleteOAuth2Claim(context.Background(), authServerId, claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClaimsAPI.DeleteOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**claimId** | **string** | &#x60;id&#x60; of Claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ClaimRequest struct via the builder pattern


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


## GetOAuth2Claim

> OAuth2Claim GetOAuth2Claim(ctx, authServerId, claimId).Execute()

Retrieve a custom token Claim



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
    claimId := "hNJ3Uk76xLagWkGx5W3N" // string | `id` of Claim

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClaimsAPI.GetOAuth2Claim(context.Background(), authServerId, claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClaimsAPI.GetOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClaimsAPI.GetOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**claimId** | **string** | &#x60;id&#x60; of Claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2Claims

> []OAuth2Claim ListOAuth2Claims(ctx, authServerId).Execute()

List all custom token Claims



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClaimsAPI.ListOAuth2Claims(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClaimsAPI.ListOAuth2Claims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2Claims`: []OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClaimsAPI.ListOAuth2Claims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOAuth2Claim

> OAuth2Claim ReplaceOAuth2Claim(ctx, authServerId, claimId).OAuth2Claim(oAuth2Claim).Execute()

Replace a custom token Claim



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
    claimId := "hNJ3Uk76xLagWkGx5W3N" // string | `id` of Claim
    oAuth2Claim := *openapiclient.NewOAuth2Claim() // OAuth2Claim | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClaimsAPI.ReplaceOAuth2Claim(context.Background(), authServerId, claimId).OAuth2Claim(oAuth2Claim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClaimsAPI.ReplaceOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClaimsAPI.ReplaceOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**claimId** | **string** | &#x60;id&#x60; of Claim | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oAuth2Claim** | [**OAuth2Claim**](OAuth2Claim.md) |  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


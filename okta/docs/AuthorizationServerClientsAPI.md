# \AuthorizationServerClientsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerClientsAPI.md#GetRefreshTokenForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Retrieve a refresh token for a Client
[**ListOAuth2ClientsForAuthorizationServer**](AuthorizationServerClientsAPI.md#ListOAuth2ClientsForAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId}/clients | List all Client resources for an authorization server
[**ListRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerClientsAPI.md#ListRefreshTokensForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | List all refresh tokens for a Client
[**RevokeRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerClientsAPI.md#RevokeRefreshTokenForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Revoke a refresh token for a Client
[**RevokeRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerClientsAPI.md#RevokeRefreshTokensForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | Revoke all refresh tokens for a Client



## GetRefreshTokenForAuthorizationServerAndClient

> OAuth2RefreshToken GetRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId).Expand(expand).Execute()

Retrieve a refresh token for a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
    expand := "expand_example" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClientsAPI.GetRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClientsAPI.GetRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefreshTokenForAuthorizationServerAndClient`: OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClientsAPI.GetRefreshTokenForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefreshTokenForAuthorizationServerAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 

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


## ListOAuth2ClientsForAuthorizationServer

> []OAuth2Client ListOAuth2ClientsForAuthorizationServer(ctx, authServerId).Execute()

List all Client resources for an authorization server



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
    resp, r, err := apiClient.AuthorizationServerClientsAPI.ListOAuth2ClientsForAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClientsAPI.ListOAuth2ClientsForAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2ClientsForAuthorizationServer`: []OAuth2Client
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClientsAPI.ListOAuth2ClientsForAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ClientsForAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefreshTokensForAuthorizationServerAndClient

> []OAuth2RefreshToken ListRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all refresh tokens for a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    expand := "expand_example" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of tokens (optional)
    limit := int32(56) // int32 | The maximum number of tokens to return (maximum 200) (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerClientsAPI.ListRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClientsAPI.ListRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRefreshTokensForAuthorizationServerAndClient`: []OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerClientsAPI.ListRefreshTokensForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRefreshTokensForAuthorizationServerAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 
 **after** | **string** | Specifies the pagination cursor for the next page of tokens | 
 **limit** | **int32** | The maximum number of tokens to return (maximum 200) | [default to -1]

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


## RevokeRefreshTokenForAuthorizationServerAndClient

> RevokeRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId).Execute()

Revoke a refresh token for a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerClientsAPI.RevokeRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClientsAPI.RevokeRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRefreshTokenForAuthorizationServerAndClientRequest struct via the builder pattern


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


## RevokeRefreshTokensForAuthorizationServerAndClient

> RevokeRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId).Execute()

Revoke all refresh tokens for a Client



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
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerClientsAPI.RevokeRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerClientsAPI.RevokeRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRefreshTokensForAuthorizationServerAndClientRequest struct via the builder pattern


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


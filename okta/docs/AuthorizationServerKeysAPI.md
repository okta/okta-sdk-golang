# \AuthorizationServerKeysAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuthorizationServerKeys**](AuthorizationServerKeysAPI.md#ListAuthorizationServerKeys) | **Get** /api/v1/authorizationServers/{authServerId}/credentials/keys | List all Credential Keys
[**RotateAuthorizationServerKeys**](AuthorizationServerKeysAPI.md#RotateAuthorizationServerKeys) | **Post** /api/v1/authorizationServers/{authServerId}/credentials/lifecycle/keyRotate | Rotate all Credential Keys



## ListAuthorizationServerKeys

> []AuthorizationServerJsonWebKey ListAuthorizationServerKeys(ctx, authServerId).Execute()

List all Credential Keys



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
    resp, r, err := apiClient.AuthorizationServerKeysAPI.ListAuthorizationServerKeys(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerKeysAPI.ListAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerKeys`: []AuthorizationServerJsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerKeysAPI.ListAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServerKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthorizationServerJsonWebKey**](AuthorizationServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateAuthorizationServerKeys

> []AuthorizationServerJsonWebKey RotateAuthorizationServerKeys(ctx, authServerId).Use(use).Execute()

Rotate all Credential Keys



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
    use := *openapiclient.NewJwkUse() // JwkUse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerKeysAPI.RotateAuthorizationServerKeys(context.Background(), authServerId).Use(use).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerKeysAPI.RotateAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAuthorizationServerKeys`: []AuthorizationServerJsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerKeysAPI.RotateAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAuthorizationServerKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **use** | [**JwkUse**](JwkUse.md) |  | 

### Return type

[**[]AuthorizationServerJsonWebKey**](AuthorizationServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


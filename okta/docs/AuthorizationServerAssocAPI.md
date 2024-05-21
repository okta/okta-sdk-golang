# \AuthorizationServerAssocAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssociatedServers**](AuthorizationServerAssocAPI.md#CreateAssociatedServers) | **Post** /api/v1/authorizationServers/{authServerId}/associatedServers | Create an associated Authorization Server
[**DeleteAssociatedServer**](AuthorizationServerAssocAPI.md#DeleteAssociatedServer) | **Delete** /api/v1/authorizationServers/{authServerId}/associatedServers/{associatedServerId} | Delete an associated Authorization Server
[**ListAssociatedServersByTrustedType**](AuthorizationServerAssocAPI.md#ListAssociatedServersByTrustedType) | **Get** /api/v1/authorizationServers/{authServerId}/associatedServers | List all associated Authorization Servers



## CreateAssociatedServers

> []AuthorizationServer CreateAssociatedServers(ctx, authServerId).AssociatedServerMediated(associatedServerMediated).Execute()

Create an associated Authorization Server



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
    associatedServerMediated := *openapiclient.NewAssociatedServerMediated() // AssociatedServerMediated | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAssocAPI.CreateAssociatedServers(context.Background(), authServerId).AssociatedServerMediated(associatedServerMediated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAssocAPI.CreateAssociatedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssociatedServers`: []AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAssocAPI.CreateAssociatedServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssociatedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associatedServerMediated** | [**AssociatedServerMediated**](AssociatedServerMediated.md) |  | 

### Return type

[**[]AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssociatedServer

> DeleteAssociatedServer(ctx, authServerId, associatedServerId).Execute()

Delete an associated Authorization Server



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
    associatedServerId := "aus6xt9jKPmCyn6kg0g4" // string | `id` of the associated Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAssocAPI.DeleteAssociatedServer(context.Background(), authServerId, associatedServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAssocAPI.DeleteAssociatedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**associatedServerId** | **string** | &#x60;id&#x60; of the associated Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssociatedServerRequest struct via the builder pattern


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


## ListAssociatedServersByTrustedType

> []AuthorizationServer ListAssociatedServersByTrustedType(ctx, authServerId).Trusted(trusted).Q(q).Limit(limit).After(after).Execute()

List all associated Authorization Servers



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
    trusted := true // bool | Searches trusted authorization servers when `true` or searches untrusted authorization servers when `false` (optional)
    q := "customasone" // string | Searches for the name or audience of the associated authorization servers (optional)
    limit := int32(56) // int32 | Specifies the number of results for a page (optional) (default to 200)
    after := "after_example" // string | Specifies the pagination cursor for the next page of the associated authorization servers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAssocAPI.ListAssociatedServersByTrustedType(context.Background(), authServerId).Trusted(trusted).Q(q).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAssocAPI.ListAssociatedServersByTrustedType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssociatedServersByTrustedType`: []AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAssocAPI.ListAssociatedServersByTrustedType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssociatedServersByTrustedTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trusted** | **bool** | Searches trusted authorization servers when &#x60;true&#x60; or searches untrusted authorization servers when &#x60;false&#x60; | 
 **q** | **string** | Searches for the name or audience of the associated authorization servers | 
 **limit** | **int32** | Specifies the number of results for a page | [default to 200]
 **after** | **string** | Specifies the pagination cursor for the next page of the associated authorization servers | 

### Return type

[**[]AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


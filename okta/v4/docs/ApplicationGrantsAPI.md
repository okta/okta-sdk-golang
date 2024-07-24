# \ApplicationGrantsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScopeConsentGrant**](ApplicationGrantsAPI.md#GetScopeConsentGrant) | **Get** /api/v1/apps/{appId}/grants/{grantId} | Retrieve an app Grant
[**GrantConsentToScope**](ApplicationGrantsAPI.md#GrantConsentToScope) | **Post** /api/v1/apps/{appId}/grants | Grant consent to scope
[**ListScopeConsentGrants**](ApplicationGrantsAPI.md#ListScopeConsentGrants) | **Get** /api/v1/apps/{appId}/grants | List all app Grants
[**RevokeScopeConsentGrant**](ApplicationGrantsAPI.md#RevokeScopeConsentGrant) | **Delete** /api/v1/apps/{appId}/grants/{grantId} | Revoke an app Grant



## GetScopeConsentGrant

> OAuth2ScopeConsentGrant GetScopeConsentGrant(ctx, appId, grantId).Expand(expand).Execute()

Retrieve an app Grant



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
    grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID
    expand := "scope" // string | An optional parameter to return scope details in the `_embedded` property. Valid value: `scope` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGrantsAPI.GetScopeConsentGrant(context.Background(), appId, grantId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGrantsAPI.GetScopeConsentGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScopeConsentGrant`: OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGrantsAPI.GetScopeConsentGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScopeConsentGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | An optional parameter to return scope details in the &#x60;_embedded&#x60; property. Valid value: &#x60;scope&#x60; | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantConsentToScope

> OAuth2ScopeConsentGrant GrantConsentToScope(ctx, appId).OAuth2ScopeConsentGrant(oAuth2ScopeConsentGrant).Execute()

Grant consent to scope



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
    oAuth2ScopeConsentGrant := *openapiclient.NewOAuth2ScopeConsentGrant("https://my_test_okta_org.oktapreview.com", "okta.users.read") // OAuth2ScopeConsentGrant | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGrantsAPI.GrantConsentToScope(context.Background(), appId).OAuth2ScopeConsentGrant(oAuth2ScopeConsentGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGrantsAPI.GrantConsentToScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantConsentToScope`: OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGrantsAPI.GrantConsentToScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantConsentToScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2ScopeConsentGrant** | [**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md) |  | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScopeConsentGrants

> []OAuth2ScopeConsentGrant ListScopeConsentGrants(ctx, appId).Expand(expand).Execute()

List all app Grants



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGrantsAPI.ListScopeConsentGrants(context.Background(), appId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGrantsAPI.ListScopeConsentGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScopeConsentGrants`: []OAuth2ScopeConsentGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGrantsAPI.ListScopeConsentGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScopeConsentGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | An optional parameter to return scope details in the &#x60;_embedded&#x60; property. Valid value: &#x60;scope&#x60; | 

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeScopeConsentGrant

> RevokeScopeConsentGrant(ctx, appId, grantId).Execute()

Revoke an app Grant



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
    grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationGrantsAPI.RevokeScopeConsentGrant(context.Background(), appId, grantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGrantsAPI.RevokeScopeConsentGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeScopeConsentGrantRequest struct via the builder pattern


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


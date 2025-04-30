# \RealmAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRealm**](RealmAPI.md#CreateRealm) | **Post** /api/v1/realms | Create a Realm
[**DeleteRealm**](RealmAPI.md#DeleteRealm) | **Delete** /api/v1/realms/{realmId} | Delete a Realm
[**GetRealm**](RealmAPI.md#GetRealm) | **Get** /api/v1/realms/{realmId} | Retrieve a Realm
[**ListRealms**](RealmAPI.md#ListRealms) | **Get** /api/v1/realms | List all Realms
[**ReplaceRealm**](RealmAPI.md#ReplaceRealm) | **Put** /api/v1/realms/{realmId} | Replace the realm profile



## CreateRealm

> Realm CreateRealm(ctx).Body(body).Execute()

Create a Realm



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
    body := *openapiclient.NewCreateRealmRequest() // CreateRealmRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealmAPI.CreateRealm(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealmAPI.CreateRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRealm`: Realm
    fmt.Fprintf(os.Stdout, "Response from `RealmAPI.CreateRealm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateRealmRequest**](CreateRealmRequest.md) |  | 

### Return type

[**Realm**](Realm.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRealm

> DeleteRealm(ctx, realmId).Execute()

Delete a Realm



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
    realmId := "vvrcFogtKCrK9aYq3fgV" // string | `id` of the Realm

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RealmAPI.DeleteRealm(context.Background(), realmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealmAPI.DeleteRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | &#x60;id&#x60; of the Realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRealmRequest struct via the builder pattern


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


## GetRealm

> Realm GetRealm(ctx, realmId).Execute()

Retrieve a Realm



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
    realmId := "vvrcFogtKCrK9aYq3fgV" // string | `id` of the Realm

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealmAPI.GetRealm(context.Background(), realmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealmAPI.GetRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRealm`: Realm
    fmt.Fprintf(os.Stdout, "Response from `RealmAPI.GetRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | &#x60;id&#x60; of the Realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Realm**](Realm.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRealms

> []Realm ListRealms(ctx).Limit(limit).After(after).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()

List all Realms



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
    limit := int32(56) // int32 | Specifies the number of results returned. Defaults to 10 if `search` is provided. (optional) (default to 200)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    search := "search_example" // string | Searches for Realms with a supported filtering expression for most properties (optional)
    sortBy := "profile.name" // string | Specifies field to sort by and can be any single property (for search queries only). (optional)
    sortOrder := "sortOrder_example" // string | Specifies sort order `asc` or `desc` (for search queries only). This parameter is ignored if `sortBy` isn't present. (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealmAPI.ListRealms(context.Background()).Limit(limit).After(after).Search(search).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealmAPI.ListRealms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRealms`: []Realm
    fmt.Fprintf(os.Stdout, "Response from `RealmAPI.ListRealms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRealmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Specifies the number of results returned. Defaults to 10 if &#x60;search&#x60; is provided. | [default to 200]
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **search** | **string** | Searches for Realms with a supported filtering expression for most properties | 
 **sortBy** | **string** | Specifies field to sort by and can be any single property (for search queries only). | 
 **sortOrder** | **string** | Specifies sort order &#x60;asc&#x60; or &#x60;desc&#x60; (for search queries only). This parameter is ignored if &#x60;sortBy&#x60; isn&#39;t present. | [default to &quot;asc&quot;]

### Return type

[**[]Realm**](Realm.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRealm

> Realm ReplaceRealm(ctx, realmId).Body(body).Execute()

Replace the realm profile



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
    realmId := "vvrcFogtKCrK9aYq3fgV" // string | `id` of the Realm
    body := *openapiclient.NewUpdateRealmRequest() // UpdateRealmRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RealmAPI.ReplaceRealm(context.Background(), realmId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealmAPI.ReplaceRealm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceRealm`: Realm
    fmt.Fprintf(os.Stdout, "Response from `RealmAPI.ReplaceRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | &#x60;id&#x60; of the Realm | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateRealmRequest**](UpdateRealmRequest.md) |  | 

### Return type

[**Realm**](Realm.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ApplicationApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateApplication**](ApplicationApi.md#ActivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/activate | Activate an Application
[**CreateApplication**](ApplicationApi.md#CreateApplication) | **Post** /api/v1/apps | Create an Application
[**DeactivateApplication**](ApplicationApi.md#DeactivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/deactivate | Deactivate an Application
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /api/v1/apps/{appId} | Delete an Application
[**GetApplication**](ApplicationApi.md#GetApplication) | **Get** /api/v1/apps/{appId} | Retrieve an Application
[**ListApplications**](ApplicationApi.md#ListApplications) | **Get** /api/v1/apps | List all Applications
[**ReplaceApplication**](ApplicationApi.md#ReplaceApplication) | **Put** /api/v1/apps/{appId} | Replace an Application



## ActivateApplication

> ActivateApplication(ctx, appId).Execute()

Activate an Application



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
    r, err := apiClient.ApplicationApi.ActivateApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.ActivateApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiActivateApplicationRequest struct via the builder pattern


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


## CreateApplication

> ListApplications200ResponseInner CreateApplication(ctx).Application(application).Activate(activate).OktaAccessGatewayAgent(oktaAccessGatewayAgent).Execute()

Create an Application



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
    application := openapiclient.listApplications_200_response_inner{AutoLoginApplication: openapiclient.NewAutoLoginApplication()} // ListApplications200ResponseInner | 
    activate := true // bool | Executes activation lifecycle operation when creating the app (optional) (default to true)
    oktaAccessGatewayAgent := "oktaAccessGatewayAgent_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.CreateApplication(context.Background()).Application(application).Activate(activate).OktaAccessGatewayAgent(oktaAccessGatewayAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: ListApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**ListApplications200ResponseInner**](ListApplications200ResponseInner.md) |  | 
 **activate** | **bool** | Executes activation lifecycle operation when creating the app | [default to true]
 **oktaAccessGatewayAgent** | **string** |  | 

### Return type

[**ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateApplication

> DeactivateApplication(ctx, appId).Execute()

Deactivate an Application



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
    r, err := apiClient.ApplicationApi.DeactivateApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.DeactivateApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeactivateApplicationRequest struct via the builder pattern


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


## DeleteApplication

> DeleteApplication(ctx, appId).Execute()

Delete an Application



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
    r, err := apiClient.ApplicationApi.DeleteApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.DeleteApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## GetApplication

> ListApplications200ResponseInner GetApplication(ctx, appId).Expand(expand).Execute()

Retrieve an Application



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.GetApplication(context.Background(), appId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ListApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 

### Return type

[**ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> []ListApplications200ResponseInner ListApplications(ctx).Q(q).After(after).Limit(limit).Filter(filter).Expand(expand).IncludeNonDeleted(includeNonDeleted).Execute()

List all Applications



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
    q := "q_example" // string |  (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of apps (optional)
    limit := int32(56) // int32 | Specifies the number of results for a page (optional) (default to -1)
    filter := "filter_example" // string | Filters apps by status, user.id, group.id or credentials.signing.kid expression (optional)
    expand := "expand_example" // string | Traverses users link relationship and optionally embeds Application User resource (optional)
    includeNonDeleted := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.ListApplications(context.Background()).Q(q).After(after).Limit(limit).Filter(filter).Expand(expand).IncludeNonDeleted(includeNonDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.ListApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplications`: []ListApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **after** | **string** | Specifies the pagination cursor for the next page of apps | 
 **limit** | **int32** | Specifies the number of results for a page | [default to -1]
 **filter** | **string** | Filters apps by status, user.id, group.id or credentials.signing.kid expression | 
 **expand** | **string** | Traverses users link relationship and optionally embeds Application User resource | 
 **includeNonDeleted** | **bool** |  | [default to false]

### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceApplication

> ListApplications200ResponseInner ReplaceApplication(ctx, appId).Application(application).Execute()

Replace an Application



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
    application := openapiclient.listApplications_200_response_inner{AutoLoginApplication: openapiclient.NewAutoLoginApplication()} // ListApplications200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationApi.ReplaceApplication(context.Background(), appId).Application(application).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationApi.ReplaceApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceApplication`: ListApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ApplicationApi.ReplaceApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | [**ListApplications200ResponseInner**](ListApplications200ResponseInner.md) |  | 

### Return type

[**ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


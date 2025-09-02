# \ApplicationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateApplication**](ApplicationAPI.md#ActivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/activate | Activate an application
[**CreateApplication**](ApplicationAPI.md#CreateApplication) | **Post** /api/v1/apps | Create an application
[**DeactivateApplication**](ApplicationAPI.md#DeactivateApplication) | **Post** /api/v1/apps/{appId}/lifecycle/deactivate | Deactivate an application
[**DeleteApplication**](ApplicationAPI.md#DeleteApplication) | **Delete** /api/v1/apps/{appId} | Delete an application
[**GetApplication**](ApplicationAPI.md#GetApplication) | **Get** /api/v1/apps/{appId} | Retrieve an application
[**ListApplications**](ApplicationAPI.md#ListApplications) | **Get** /api/v1/apps | List all applications
[**ReplaceApplication**](ApplicationAPI.md#ReplaceApplication) | **Put** /api/v1/apps/{appId} | Replace an application



## ActivateApplication

> ActivateApplication(ctx, appId).Execute()

Activate an application



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
	r, err := apiClient.ApplicationAPI.ActivateApplication(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ActivateApplication``: %v\n", err)
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

Create an application



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
	application := openapiclient.listApplications_200_response_inner{AutoLoginApplication: openapiclient.NewAutoLoginApplication("Label_example", "SignOnMode_example")} // ListApplications200ResponseInner | 
	activate := true // bool | Executes activation lifecycle operation when creating the app (optional) (default to true)
	oktaAccessGatewayAgent := "oktaAccessGatewayAgent_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.CreateApplication(context.Background()).Application(application).Activate(activate).OktaAccessGatewayAgent(oktaAccessGatewayAgent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CreateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplication`: ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.CreateApplication`: %v\n", resp)
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

Deactivate an application



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
	r, err := apiClient.ApplicationAPI.DeactivateApplication(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeactivateApplication``: %v\n", err)
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

Delete an application



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
	r, err := apiClient.ApplicationAPI.DeleteApplication(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.DeleteApplication``: %v\n", err)
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

Retrieve an application



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
	expand := "user/0oa1gjh63g214q0Hq0g4" // string | An optional query parameter to return the specified [Application User](/openapi/okta-management/management/tag/ApplicationUsers/) in the `_embedded` property. Valid value: `expand=user/{userId}` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplication(context.Background(), appId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | An optional query parameter to return the specified [Application User](/openapi/okta-management/management/tag/ApplicationUsers/) in the &#x60;_embedded&#x60; property. Valid value: &#x60;expand&#x3D;user/{userId}&#x60; | 

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

> []ListApplications200ResponseInner ListApplications(ctx).Q(q).After(after).UseOptimization(useOptimization).Limit(limit).Filter(filter).Expand(expand).IncludeNonDeleted(includeNonDeleted).Execute()

List all applications



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
	q := "Okta" // string | Searches for apps with `name` or `label` properties that starts with the `q` value using the `startsWith` operation (optional)
	after := "16278919418571" // string | Specifies the [pagination](/#pagination) cursor for the next page of results. Treat this as an opaque value obtained through the `next` link relationship. (optional)
	useOptimization := true // bool | Specifies whether to use query optimization. If you specify `useOptimization=true` in the request query, the response contains a subset of app instance properties. (optional) (default to false)
	limit := int32(56) // int32 | Specifies the number of results per page (optional) (default to -1)
	filter := "status%20eq%20%22ACTIVE%22" // string | Filters apps by `status`, `user.id`, `group.id`, `credentials.signing.kid` or `name` expression that supports the `eq` operator (optional)
	expand := "user/0oa1gjh63g214q0Hq0g4" // string | An optional parameter used for link expansion to embed more resources in the response. Only supports `expand=user/{userId}` and must be used with the `user.id eq \"{userId}\"` filter query for the same user. Returns the assigned [application user](/openapi/okta-management/management/tag/ApplicationUsers/) in the `_embedded` property. (optional)
	includeNonDeleted := true // bool | Specifies whether to include non-active, but not deleted apps in the results (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ListApplications(context.Background()).Q(q).After(after).UseOptimization(useOptimization).Limit(limit).Filter(filter).Expand(expand).IncludeNonDeleted(includeNonDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Searches for apps with &#x60;name&#x60; or &#x60;label&#x60; properties that starts with the &#x60;q&#x60; value using the &#x60;startsWith&#x60; operation | 
 **after** | **string** | Specifies the [pagination](/#pagination) cursor for the next page of results. Treat this as an opaque value obtained through the &#x60;next&#x60; link relationship. | 
 **useOptimization** | **bool** | Specifies whether to use query optimization. If you specify &#x60;useOptimization&#x3D;true&#x60; in the request query, the response contains a subset of app instance properties. | [default to false]
 **limit** | **int32** | Specifies the number of results per page | [default to -1]
 **filter** | **string** | Filters apps by &#x60;status&#x60;, &#x60;user.id&#x60;, &#x60;group.id&#x60;, &#x60;credentials.signing.kid&#x60; or &#x60;name&#x60; expression that supports the &#x60;eq&#x60; operator | 
 **expand** | **string** | An optional parameter used for link expansion to embed more resources in the response. Only supports &#x60;expand&#x3D;user/{userId}&#x60; and must be used with the &#x60;user.id eq \&quot;{userId}\&quot;&#x60; filter query for the same user. Returns the assigned [application user](/openapi/okta-management/management/tag/ApplicationUsers/) in the &#x60;_embedded&#x60; property. | 
 **includeNonDeleted** | **bool** | Specifies whether to include non-active, but not deleted apps in the results | [default to false]

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

Replace an application



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
	application := openapiclient.listApplications_200_response_inner{AutoLoginApplication: openapiclient.NewAutoLoginApplication("Label_example", "SignOnMode_example")} // ListApplications200ResponseInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.ReplaceApplication(context.Background(), appId).Application(application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.ReplaceApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceApplication`: ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.ReplaceApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

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


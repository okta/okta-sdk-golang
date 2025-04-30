# \BehaviorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBehaviorDetectionRule**](BehaviorAPI.md#ActivateBehaviorDetectionRule) | **Post** /api/v1/behaviors/{behaviorId}/lifecycle/activate | Activate a Behavior Detection Rule
[**CreateBehaviorDetectionRule**](BehaviorAPI.md#CreateBehaviorDetectionRule) | **Post** /api/v1/behaviors | Create a Behavior Detection Rule
[**DeactivateBehaviorDetectionRule**](BehaviorAPI.md#DeactivateBehaviorDetectionRule) | **Post** /api/v1/behaviors/{behaviorId}/lifecycle/deactivate | Deactivate a Behavior Detection Rule
[**DeleteBehaviorDetectionRule**](BehaviorAPI.md#DeleteBehaviorDetectionRule) | **Delete** /api/v1/behaviors/{behaviorId} | Delete a Behavior Detection Rule
[**GetBehaviorDetectionRule**](BehaviorAPI.md#GetBehaviorDetectionRule) | **Get** /api/v1/behaviors/{behaviorId} | Retrieve a Behavior Detection Rule
[**ListBehaviorDetectionRules**](BehaviorAPI.md#ListBehaviorDetectionRules) | **Get** /api/v1/behaviors | List all Behavior Detection Rules
[**ReplaceBehaviorDetectionRule**](BehaviorAPI.md#ReplaceBehaviorDetectionRule) | **Put** /api/v1/behaviors/{behaviorId} | Replace a Behavior Detection Rule



## ActivateBehaviorDetectionRule

> ListBehaviorDetectionRules200ResponseInner ActivateBehaviorDetectionRule(ctx, behaviorId).Execute()

Activate a Behavior Detection Rule



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
    behaviorId := "abcd1234" // string | id of the Behavior Detection Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.ActivateBehaviorDetectionRule(context.Background(), behaviorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.ActivateBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateBehaviorDetectionRule`: ListBehaviorDetectionRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.ActivateBehaviorDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**behaviorId** | **string** | id of the Behavior Detection Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateBehaviorDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBehaviorDetectionRule

> BehaviorRule CreateBehaviorDetectionRule(ctx).Rule(rule).Execute()

Create a Behavior Detection Rule



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
    rule := openapiclient.listBehaviorDetectionRules_200_response_inner{BehaviorRuleAnomalousDevice: openapiclient.NewBehaviorRuleAnomalousDevice("Name_example", "Type_example")} // ListBehaviorDetectionRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.CreateBehaviorDetectionRule(context.Background()).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.CreateBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBehaviorDetectionRule`: BehaviorRule
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.CreateBehaviorDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBehaviorDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rule** | [**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md) |  | 

### Return type

[**BehaviorRule**](BehaviorRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateBehaviorDetectionRule

> ListBehaviorDetectionRules200ResponseInner DeactivateBehaviorDetectionRule(ctx, behaviorId).Execute()

Deactivate a Behavior Detection Rule



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
    behaviorId := "abcd1234" // string | id of the Behavior Detection Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.DeactivateBehaviorDetectionRule(context.Background(), behaviorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.DeactivateBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateBehaviorDetectionRule`: ListBehaviorDetectionRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.DeactivateBehaviorDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**behaviorId** | **string** | id of the Behavior Detection Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateBehaviorDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBehaviorDetectionRule

> DeleteBehaviorDetectionRule(ctx, behaviorId).Execute()

Delete a Behavior Detection Rule



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
    behaviorId := "abcd1234" // string | id of the Behavior Detection Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BehaviorAPI.DeleteBehaviorDetectionRule(context.Background(), behaviorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.DeleteBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**behaviorId** | **string** | id of the Behavior Detection Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBehaviorDetectionRuleRequest struct via the builder pattern


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


## GetBehaviorDetectionRule

> ListBehaviorDetectionRules200ResponseInner GetBehaviorDetectionRule(ctx, behaviorId).Execute()

Retrieve a Behavior Detection Rule



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
    behaviorId := "abcd1234" // string | id of the Behavior Detection Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.GetBehaviorDetectionRule(context.Background(), behaviorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.GetBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBehaviorDetectionRule`: ListBehaviorDetectionRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.GetBehaviorDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**behaviorId** | **string** | id of the Behavior Detection Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBehaviorDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBehaviorDetectionRules

> []ListBehaviorDetectionRules200ResponseInner ListBehaviorDetectionRules(ctx).Execute()

List all Behavior Detection Rules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.ListBehaviorDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.ListBehaviorDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBehaviorDetectionRules`: []ListBehaviorDetectionRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.ListBehaviorDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBehaviorDetectionRulesRequest struct via the builder pattern


### Return type

[**[]ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBehaviorDetectionRule

> ListBehaviorDetectionRules200ResponseInner ReplaceBehaviorDetectionRule(ctx, behaviorId).Rule(rule).Execute()

Replace a Behavior Detection Rule



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
    behaviorId := "abcd1234" // string | id of the Behavior Detection Rule
    rule := openapiclient.listBehaviorDetectionRules_200_response_inner{BehaviorRuleAnomalousDevice: openapiclient.NewBehaviorRuleAnomalousDevice("Name_example", "Type_example")} // ListBehaviorDetectionRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BehaviorAPI.ReplaceBehaviorDetectionRule(context.Background(), behaviorId).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BehaviorAPI.ReplaceBehaviorDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceBehaviorDetectionRule`: ListBehaviorDetectionRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BehaviorAPI.ReplaceBehaviorDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**behaviorId** | **string** | id of the Behavior Detection Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBehaviorDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rule** | [**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md) |  | 

### Return type

[**ListBehaviorDetectionRules200ResponseInner**](ListBehaviorDetectionRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


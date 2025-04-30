# \AuthorizationServerRulesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#ActivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a Policy Rule
[**CreateAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#CreateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | Create a Policy Rule
[**DeactivateAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#DeactivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a Policy Rule
[**DeleteAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#DeleteAuthorizationServerPolicyRule) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Delete a Policy Rule
[**GetAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#GetAuthorizationServerPolicyRule) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Retrieve a Policy Rule
[**ListAuthorizationServerPolicyRules**](AuthorizationServerRulesAPI.md#ListAuthorizationServerPolicyRules) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | List all Policy Rules
[**ReplaceAuthorizationServerPolicyRule**](AuthorizationServerRulesAPI.md#ReplaceAuthorizationServerPolicyRule) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Replace a Policy Rule



## ActivateAuthorizationServerPolicyRule

> ActivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Activate a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerRulesAPI.ActivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.ActivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## CreateAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule CreateAuthorizationServerPolicyRule(ctx, authServerId, policyId).PolicyRule(policyRule).Execute()

Create a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerRulesAPI.CreateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.CreateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerRulesAPI.CreateAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRule** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md) |  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAuthorizationServerPolicyRule

> DeactivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Deactivate a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerRulesAPI.DeactivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.DeactivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## DeleteAuthorizationServerPolicyRule

> DeleteAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Delete a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerRulesAPI.DeleteAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.DeleteAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## GetAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule GetAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Retrieve a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerRulesAPI.GetAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.GetAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerRulesAPI.GetAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationServerPolicyRules

> []AuthorizationServerPolicyRule ListAuthorizationServerPolicyRules(ctx, authServerId, policyId).Execute()

List all Policy Rules



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerRulesAPI.ListAuthorizationServerPolicyRules(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.ListAuthorizationServerPolicyRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerPolicyRules`: []AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerRulesAPI.ListAuthorizationServerPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServerPolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule ReplaceAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).PolicyRule(policyRule).Execute()

Replace a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerRulesAPI.ReplaceAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerRulesAPI.ReplaceAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerRulesAPI.ReplaceAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **policyRule** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md) |  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


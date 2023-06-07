# \PolicyApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePolicy**](PolicyApi.md#ActivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/activate | Activate a Policy
[**ActivatePolicyRule**](PolicyApi.md#ActivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a Policy Rule
[**ClonePolicy**](PolicyApi.md#ClonePolicy) | **Post** /api/v1/policies/{policyId}/clone | Clone an existing policy
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Post** /api/v1/policies | Create a Policy
[**CreatePolicyRule**](PolicyApi.md#CreatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules | Create a Policy Rule
[**DeactivatePolicy**](PolicyApi.md#DeactivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/deactivate | Deactivate a Policy
[**DeactivatePolicyRule**](PolicyApi.md#DeactivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a Policy Rule
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /api/v1/policies/{policyId} | Delete a Policy
[**DeletePolicyRule**](PolicyApi.md#DeletePolicyRule) | **Delete** /api/v1/policies/{policyId}/rules/{ruleId} | Delete a Policy Rule
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /api/v1/policies/{policyId} | Retrieve a Policy
[**GetPolicyRule**](PolicyApi.md#GetPolicyRule) | **Get** /api/v1/policies/{policyId}/rules/{ruleId} | Retrieve a Policy Rule
[**ListPolicies**](PolicyApi.md#ListPolicies) | **Get** /api/v1/policies | List all Policies
[**ListPolicyRules**](PolicyApi.md#ListPolicyRules) | **Get** /api/v1/policies/{policyId}/rules | List all Policy Rules
[**ReplacePolicy**](PolicyApi.md#ReplacePolicy) | **Put** /api/v1/policies/{policyId} | Replace a Policy
[**ReplacePolicyRule**](PolicyApi.md#ReplacePolicyRule) | **Put** /api/v1/policies/{policyId}/rules/{ruleId} | Replace a Policy Rule



## ActivatePolicy

> ActivatePolicy(ctx, policyId).Execute()

Activate a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ActivatePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ActivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRequest struct via the builder pattern


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


## ActivatePolicyRule

> ActivatePolicyRule(ctx, policyId, ruleId).Execute()

Activate a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ActivatePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ActivatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRuleRequest struct via the builder pattern


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


## ClonePolicy

> ListPolicies200ResponseInner ClonePolicy(ctx, policyId).Execute()

Clone an existing policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ClonePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ClonePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClonePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ClonePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> ListPolicies200ResponseInner CreatePolicy(ctx).Policy(policy).Activate(activate).Execute()

Create a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policy := openapiclient.listPolicies_200_response_inner{AccessPolicy: openapiclient.NewAccessPolicy()} // ListPolicies200ResponseInner | 
    activate := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.CreatePolicy(context.Background()).Policy(policy).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md) |  | 
 **activate** | **bool** |  | [default to true]

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyRule

> ListPolicyRules200ResponseInner CreatePolicyRule(ctx, policyId).PolicyRule(policyRule).Execute()

Create a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.CreatePolicyRule(context.Background(), policyId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePolicy

> DeactivatePolicy(ctx, policyId).Execute()

Deactivate a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.DeactivatePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeactivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRequest struct via the builder pattern


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


## DeactivatePolicyRule

> DeactivatePolicyRule(ctx, policyId, ruleId).Execute()

Deactivate a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.DeactivatePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeactivatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRuleRequest struct via the builder pattern


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


## DeletePolicy

> DeletePolicy(ctx, policyId).Execute()

Delete a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.DeletePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## DeletePolicyRule

> DeletePolicyRule(ctx, policyId, ruleId).Execute()

Delete a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.DeletePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRuleRequest struct via the builder pattern


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


## GetPolicy

> ListPolicies200ResponseInner GetPolicy(ctx, policyId).Expand(expand).Execute()

Retrieve a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    expand := "expand_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicy(context.Background(), policyId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | [default to &quot;&quot;]

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyRule

> ListPolicyRules200ResponseInner GetPolicyRule(ctx, policyId, ruleId).Execute()

Retrieve a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []ListPolicies200ResponseInner ListPolicies(ctx).Type_(type_).Status(status).Expand(expand).Execute()

List all Policies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | 
    status := "status_example" // string |  (optional)
    expand := "expand_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicies(context.Background()).Type_(type_).Status(status).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: []ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **status** | **string** |  | 
 **expand** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyRules

> []ListPolicyRules200ResponseInner ListPolicyRules(ctx, policyId).Execute()

List all Policy Rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicyRules(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicyRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyRules`: []ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicy

> ListPolicies200ResponseInner ReplacePolicy(ctx, policyId).Policy(policy).Execute()

Replace a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    policy := openapiclient.listPolicies_200_response_inner{AccessPolicy: openapiclient.NewAccessPolicy()} // ListPolicies200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ReplacePolicy(context.Background(), policyId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReplacePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReplacePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md) |  | 

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicyRule

> ListPolicyRules200ResponseInner ReplacePolicyRule(ctx, policyId, ruleId).PolicyRule(policyRule).Execute()

Replace a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 
    policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ReplacePolicyRule(context.Background(), policyId, ruleId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReplacePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReplacePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


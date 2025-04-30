# \ApplicationPoliciesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignApplicationPolicy**](ApplicationPoliciesAPI.md#AssignApplicationPolicy) | **Put** /api/v1/apps/{appId}/policies/{policyId} | Assign an application to a Policy



## AssignApplicationPolicy

> AssignApplicationPolicy(ctx, appId, policyId).Execute()

Assign an application to a Policy



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationPoliciesAPI.AssignApplicationPolicy(context.Background(), appId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPoliciesAPI.AssignApplicationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignApplicationPolicyRequest struct via the builder pattern


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


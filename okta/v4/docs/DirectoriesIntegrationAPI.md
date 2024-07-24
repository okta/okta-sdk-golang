# \DirectoriesIntegrationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateADGroupMembership**](DirectoriesIntegrationAPI.md#UpdateADGroupMembership) | **Post** /api/v1/directories/{appInstanceId}/groups/modify | Update an AD Group membership



## UpdateADGroupMembership

> UpdateADGroupMembership(ctx, appInstanceId).AgentAction(agentAction).Execute()

Update an AD Group membership



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
    appInstanceId := "appInstanceId_example" // string | ID of the AD AppInstance in Okta
    agentAction := *openapiclient.NewAgentAction() // AgentAction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DirectoriesIntegrationAPI.UpdateADGroupMembership(context.Background(), appInstanceId).AgentAction(agentAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoriesIntegrationAPI.UpdateADGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** | ID of the AD AppInstance in Okta | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateADGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentAction** | [**AgentAction**](AgentAction.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


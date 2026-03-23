# \DirectoriesIntegrationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupAttributeQueryResult**](DirectoriesIntegrationAPI.md#GetGroupAttributeQueryResult) | **Get** /api/v1/directories/{appInstanceId}/groups/{groupId}/query/{resultId} | Retrieve the results of an AD group query
[**SubmitGroupAttributeQuery**](DirectoriesIntegrationAPI.md#SubmitGroupAttributeQuery) | **Post** /api/v1/directories/{appInstanceId}/groups/{groupId}/query | Submit a query for AD Group
[**UpdateGroupMembership**](DirectoriesIntegrationAPI.md#UpdateGroupMembership) | **Post** /api/v1/directories/{appInstanceId}/groups/modify | Update an external directory group membership



## GetGroupAttributeQueryResult

> GroupProfileResult GetGroupAttributeQueryResult(ctx, appInstanceId, groupId, resultId).Execute()

Retrieve the results of an AD group query



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
	appInstanceId := "00a1xucgTZFrziXg10g4" // string | ID of the AD instance in Okta
	groupId := "00g1xucgTZFrziXg10g4" // string | ID of the Okta group
	resultId := "resultId_example" // string | The unique identifier returned by the initial POST request (`POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query`)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoriesIntegrationAPI.GetGroupAttributeQueryResult(context.Background(), appInstanceId, groupId, resultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoriesIntegrationAPI.GetGroupAttributeQueryResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupAttributeQueryResult`: GroupProfileResult
	fmt.Fprintf(os.Stdout, "Response from `DirectoriesIntegrationAPI.GetGroupAttributeQueryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** | ID of the AD instance in Okta | 
**groupId** | **string** | ID of the Okta group | 
**resultId** | **string** | The unique identifier returned by the initial POST request (&#x60;POST /api/v1/directories/{appInstanceId}/groups/{groupId}/query&#x60;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAttributeQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GroupProfileResult**](GroupProfileResult.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitGroupAttributeQuery

> GroupQueryResponse SubmitGroupAttributeQuery(ctx, appInstanceId, groupId).GroupQueryRequest(groupQueryRequest).Execute()

Submit a query for AD Group



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
	appInstanceId := "00a1xucgTZFrziXg10g4" // string | ID of the AD instance in Okta
	groupId := "00g1xucgTZFrziXg10g4" // string | ID of the Okta group
	groupQueryRequest := *openapiclient.NewGroupQueryRequest([]string{"Attributes_example"}) // GroupQueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectoriesIntegrationAPI.SubmitGroupAttributeQuery(context.Background(), appInstanceId, groupId).GroupQueryRequest(groupQueryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoriesIntegrationAPI.SubmitGroupAttributeQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitGroupAttributeQuery`: GroupQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectoriesIntegrationAPI.SubmitGroupAttributeQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** | ID of the AD instance in Okta | 
**groupId** | **string** | ID of the Okta group | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitGroupAttributeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupQueryRequest** | [**GroupQueryRequest**](GroupQueryRequest.md) |  | 

### Return type

[**GroupQueryResponse**](GroupQueryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMembership

> UpdateGroupMembership(ctx, appInstanceId).AgentAction(agentAction).Execute()

Update an external directory group membership



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
	appInstanceId := "appInstanceId_example" // string | ID of the Active Directory or LDAP app instance in Okta
	agentAction := *openapiclient.NewAgentAction("Id_example", *openapiclient.NewParameters("Action_example", "member", []string{"Values_example"})) // AgentAction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DirectoriesIntegrationAPI.UpdateGroupMembership(context.Background(), appInstanceId).AgentAction(agentAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectoriesIntegrationAPI.UpdateGroupMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appInstanceId** | **string** | ID of the Active Directory or LDAP app instance in Okta | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMembershipRequest struct via the builder pattern


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


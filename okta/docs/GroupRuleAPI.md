# \GroupRuleAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGroupRule**](GroupRuleAPI.md#ActivateGroupRule) | **Post** /api/v1/groups/rules/{groupRuleId}/lifecycle/activate | Activate a group rule
[**CreateGroupRule**](GroupRuleAPI.md#CreateGroupRule) | **Post** /api/v1/groups/rules | Create a group rule
[**DeactivateGroupRule**](GroupRuleAPI.md#DeactivateGroupRule) | **Post** /api/v1/groups/rules/{groupRuleId}/lifecycle/deactivate | Deactivate a group rule
[**DeleteGroupRule**](GroupRuleAPI.md#DeleteGroupRule) | **Delete** /api/v1/groups/rules/{groupRuleId} | Delete a group rule
[**GetGroupRule**](GroupRuleAPI.md#GetGroupRule) | **Get** /api/v1/groups/rules/{groupRuleId} | Retrieve a group rule
[**ListGroupRules**](GroupRuleAPI.md#ListGroupRules) | **Get** /api/v1/groups/rules | List all group rules
[**ReplaceGroupRule**](GroupRuleAPI.md#ReplaceGroupRule) | **Put** /api/v1/groups/rules/{groupRuleId} | Replace a group rule



## ActivateGroupRule

> ActivateGroupRule(ctx, groupRuleId).Execute()

Activate a group rule



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
	groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupRuleAPI.ActivateGroupRule(context.Background(), groupRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.ActivateGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateGroupRuleRequest struct via the builder pattern


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


## CreateGroupRule

> GroupRule CreateGroupRule(ctx).GroupRule(groupRule).Execute()

Create a group rule



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
	groupRule := *openapiclient.NewCreateGroupRuleRequest() // CreateGroupRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupRuleAPI.CreateGroupRule(context.Background()).GroupRule(groupRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.CreateGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupRule`: GroupRule
	fmt.Fprintf(os.Stdout, "Response from `GroupRuleAPI.CreateGroupRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRule** | [**CreateGroupRuleRequest**](CreateGroupRuleRequest.md) |  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateGroupRule

> DeactivateGroupRule(ctx, groupRuleId).Execute()

Deactivate a group rule



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
	groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupRuleAPI.DeactivateGroupRule(context.Background(), groupRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.DeactivateGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateGroupRuleRequest struct via the builder pattern


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


## DeleteGroupRule

> DeleteGroupRule(ctx, groupRuleId).RemoveUsers(removeUsers).Execute()

Delete a group rule



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
	groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
	removeUsers := true // bool | If set to `true`, removes users from groups assigned by this rule (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupRuleAPI.DeleteGroupRule(context.Background(), groupRuleId).RemoveUsers(removeUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.DeleteGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeUsers** | **bool** | If set to &#x60;true&#x60;, removes users from groups assigned by this rule | [default to false]

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


## GetGroupRule

> GroupRule GetGroupRule(ctx, groupRuleId).Expand(expand).Execute()

Retrieve a group rule



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
	groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
	expand := "expand_example" // string | If specified as `groupIdToGroupNameMap`, then show group names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupRuleAPI.GetGroupRule(context.Background(), groupRuleId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.GetGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupRule`: GroupRule
	fmt.Fprintf(os.Stdout, "Response from `GroupRuleAPI.GetGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | If specified as &#x60;groupIdToGroupNameMap&#x60;, then show group names | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupRules

> []GroupRule ListGroupRules(ctx).Limit(limit).After(after).Search(search).Expand(expand).Execute()

List all group rules



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
	limit := int32(56) // int32 | Specifies the number of rule results in a page (optional) (default to 50)
	after := "after_example" // string | Specifies the pagination cursor for the next page of rules (optional)
	search := "search_example" // string | Specifies the keyword to search rules for (optional)
	expand := "expand_example" // string | If specified as `groupIdToGroupNameMap`, then displays group names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupRuleAPI.ListGroupRules(context.Background()).Limit(limit).After(after).Search(search).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.ListGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGroupRules`: []GroupRule
	fmt.Fprintf(os.Stdout, "Response from `GroupRuleAPI.ListGroupRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Specifies the number of rule results in a page | [default to 50]
 **after** | **string** | Specifies the pagination cursor for the next page of rules | 
 **search** | **string** | Specifies the keyword to search rules for | 
 **expand** | **string** | If specified as &#x60;groupIdToGroupNameMap&#x60;, then displays group names | 

### Return type

[**[]GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceGroupRule

> GroupRule ReplaceGroupRule(ctx, groupRuleId).GroupRule(groupRule).Execute()

Replace a group rule



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
	groupRuleId := "0pr3f7zMZZHPgUoWO0g4" // string | The `id` of the group rule
	groupRule := *openapiclient.NewGroupRule() // GroupRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupRuleAPI.ReplaceGroupRule(context.Background(), groupRuleId).GroupRule(groupRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupRuleAPI.ReplaceGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceGroupRule`: GroupRule
	fmt.Fprintf(os.Stdout, "Response from `GroupRuleAPI.ReplaceGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupRuleId** | **string** | The &#x60;id&#x60; of the group rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRule** | [**GroupRule**](GroupRule.md) |  | 

### Return type

[**GroupRule**](GroupRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


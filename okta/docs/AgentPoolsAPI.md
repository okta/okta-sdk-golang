# \AgentPoolsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAgentPoolsUpdate**](AgentPoolsAPI.md#ActivateAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/activate | Activate an Agent Pool update
[**CreateAgentPoolsUpdate**](AgentPoolsAPI.md#CreateAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates | Create an Agent Pool update
[**DeactivateAgentPoolsUpdate**](AgentPoolsAPI.md#DeactivateAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/deactivate | Deactivate an Agent Pool update
[**DeleteAgentPoolsUpdate**](AgentPoolsAPI.md#DeleteAgentPoolsUpdate) | **Delete** /api/v1/agentPools/{poolId}/updates/{updateId} | Delete an Agent Pool update
[**GetAgentPoolsUpdateInstance**](AgentPoolsAPI.md#GetAgentPoolsUpdateInstance) | **Get** /api/v1/agentPools/{poolId}/updates/{updateId} | Retrieve an Agent Pool update by id
[**GetAgentPoolsUpdateSettings**](AgentPoolsAPI.md#GetAgentPoolsUpdateSettings) | **Get** /api/v1/agentPools/{poolId}/updates/settings | Retrieve an Agent Pool update&#39;s settings
[**ListAgentPools**](AgentPoolsAPI.md#ListAgentPools) | **Get** /api/v1/agentPools | List all Agent Pools
[**ListAgentPoolsUpdates**](AgentPoolsAPI.md#ListAgentPoolsUpdates) | **Get** /api/v1/agentPools/{poolId}/updates | List all Agent Pool updates
[**PauseAgentPoolsUpdate**](AgentPoolsAPI.md#PauseAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/pause | Pause an Agent Pool update
[**ResumeAgentPoolsUpdate**](AgentPoolsAPI.md#ResumeAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/resume | Resume an Agent Pool update
[**RetryAgentPoolsUpdate**](AgentPoolsAPI.md#RetryAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/retry | Retry an Agent Pool update
[**StopAgentPoolsUpdate**](AgentPoolsAPI.md#StopAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId}/stop | Stop an Agent Pool update
[**UpdateAgentPoolsUpdate**](AgentPoolsAPI.md#UpdateAgentPoolsUpdate) | **Post** /api/v1/agentPools/{poolId}/updates/{updateId} | Update an Agent Pool update by id
[**UpdateAgentPoolsUpdateSettings**](AgentPoolsAPI.md#UpdateAgentPoolsUpdateSettings) | **Post** /api/v1/agentPools/{poolId}/updates/settings | Update an Agent Pool update settings



## ActivateAgentPoolsUpdate

> AgentPoolUpdate ActivateAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Activate an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.ActivateAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.ActivateAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.ActivateAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAgentPoolsUpdate

> AgentPoolUpdate CreateAgentPoolsUpdate(ctx, poolId).AgentPoolUpdate(agentPoolUpdate).Execute()

Create an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    agentPoolUpdate := *openapiclient.NewAgentPoolUpdate() // AgentPoolUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.CreateAgentPoolsUpdate(context.Background(), poolId).AgentPoolUpdate(agentPoolUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.CreateAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.CreateAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPoolUpdate** | [**AgentPoolUpdate**](AgentPoolUpdate.md) |  | 

### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAgentPoolsUpdate

> AgentPoolUpdate DeactivateAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Deactivate an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.DeactivateAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.DeactivateAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.DeactivateAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentPoolsUpdate

> DeleteAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Delete an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgentPoolsAPI.DeleteAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.DeleteAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentPoolsUpdateRequest struct via the builder pattern


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


## GetAgentPoolsUpdateInstance

> AgentPoolUpdate GetAgentPoolsUpdateInstance(ctx, poolId, updateId).Execute()

Retrieve an Agent Pool update by id



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateInstance(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.GetAgentPoolsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentPoolsUpdateInstance`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.GetAgentPoolsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentPoolsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentPoolsUpdateSettings

> AgentPoolUpdateSetting GetAgentPoolsUpdateSettings(ctx, poolId).Execute()

Retrieve an Agent Pool update's settings



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateSettings(context.Background(), poolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.GetAgentPoolsUpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentPoolsUpdateSettings`: AgentPoolUpdateSetting
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.GetAgentPoolsUpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentPoolsUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentPoolUpdateSetting**](AgentPoolUpdateSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentPools

> []AgentPool ListAgentPools(ctx).LimitPerPoolType(limitPerPoolType).PoolType(poolType).After(after).Execute()

List all Agent Pools



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
    limitPerPoolType := int32(56) // int32 | Maximum number of AgentPools being returned (optional) (default to 5)
    poolType := "poolType_example" // string | Agent type to search for (optional)
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.ListAgentPools(context.Background()).LimitPerPoolType(limitPerPoolType).PoolType(poolType).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.ListAgentPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAgentPools`: []AgentPool
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.ListAgentPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAgentPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limitPerPoolType** | **int32** | Maximum number of AgentPools being returned | [default to 5]
 **poolType** | **string** | Agent type to search for | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 

### Return type

[**[]AgentPool**](AgentPool.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentPoolsUpdates

> []AgentPoolUpdate ListAgentPoolsUpdates(ctx, poolId).Scheduled(scheduled).Execute()

List all Agent Pool updates



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    scheduled := true // bool | Scope the list only to scheduled or ad-hoc updates. If the parameter is not provided we will return the whole list of updates. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.ListAgentPoolsUpdates(context.Background(), poolId).Scheduled(scheduled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.ListAgentPoolsUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAgentPoolsUpdates`: []AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.ListAgentPoolsUpdates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentPoolsUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduled** | **bool** | Scope the list only to scheduled or ad-hoc updates. If the parameter is not provided we will return the whole list of updates. | 

### Return type

[**[]AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseAgentPoolsUpdate

> AgentPoolUpdate PauseAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Pause an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.PauseAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.PauseAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.PauseAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeAgentPoolsUpdate

> AgentPoolUpdate ResumeAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Resume an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.ResumeAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.ResumeAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.ResumeAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryAgentPoolsUpdate

> AgentPoolUpdate RetryAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Retry an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.RetryAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.RetryAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetryAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.RetryAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAgentPoolsUpdate

> AgentPoolUpdate StopAgentPoolsUpdate(ctx, poolId, updateId).Execute()

Stop an Agent Pool update



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.StopAgentPoolsUpdate(context.Background(), poolId, updateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.StopAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.StopAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentPoolsUpdate

> AgentPoolUpdate UpdateAgentPoolsUpdate(ctx, poolId, updateId).AgentPoolUpdate(agentPoolUpdate).Execute()

Update an Agent Pool update by id



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    updateId := "updateId_example" // string | Id of the update
    agentPoolUpdate := *openapiclient.NewAgentPoolUpdate() // AgentPoolUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdate(context.Background(), poolId, updateId).AgentPoolUpdate(agentPoolUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.UpdateAgentPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentPoolsUpdate`: AgentPoolUpdate
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.UpdateAgentPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 
**updateId** | **string** | Id of the update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentPoolUpdate** | [**AgentPoolUpdate**](AgentPoolUpdate.md) |  | 

### Return type

[**AgentPoolUpdate**](AgentPoolUpdate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentPoolsUpdateSettings

> AgentPoolUpdateSetting UpdateAgentPoolsUpdateSettings(ctx, poolId).AgentPoolUpdateSetting(agentPoolUpdateSetting).Execute()

Update an Agent Pool update settings



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
    poolId := "poolId_example" // string | Id of the agent pool for which the settings will apply
    agentPoolUpdateSetting := *openapiclient.NewAgentPoolUpdateSetting() // AgentPoolUpdateSetting | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdateSettings(context.Background(), poolId).AgentPoolUpdateSetting(agentPoolUpdateSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentPoolsAPI.UpdateAgentPoolsUpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentPoolsUpdateSettings`: AgentPoolUpdateSetting
    fmt.Fprintf(os.Stdout, "Response from `AgentPoolsAPI.UpdateAgentPoolsUpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string** | Id of the agent pool for which the settings will apply | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentPoolsUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPoolUpdateSetting** | [**AgentPoolUpdateSetting**](AgentPoolUpdateSetting.md) |  | 

### Return type

[**AgentPoolUpdateSetting**](AgentPoolUpdateSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


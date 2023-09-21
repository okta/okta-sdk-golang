# \ApiServiceIntegrationsApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateApiServiceIntegrationInstanceSecret**](ApiServiceIntegrationsApi.md#ActivateApiServiceIntegrationInstanceSecret) | **Post** /integrations/api/v1/api-services/{apiServiceId}/credentials/secrets/{secretId}/lifecycle/activate | Activate an API Service Integration instance Secret
[**CreateApiServiceIntegrationInstance**](ApiServiceIntegrationsApi.md#CreateApiServiceIntegrationInstance) | **Post** /integrations/api/v1/api-services | Create an API Service Integration instance
[**CreateApiServiceIntegrationInstanceSecret**](ApiServiceIntegrationsApi.md#CreateApiServiceIntegrationInstanceSecret) | **Post** /integrations/api/v1/api-services/{apiServiceId}/credentials/secrets | Create an API Service Integration instance Secret
[**DeactivateApiServiceIntegrationInstanceSecret**](ApiServiceIntegrationsApi.md#DeactivateApiServiceIntegrationInstanceSecret) | **Post** /integrations/api/v1/api-services/{apiServiceId}/credentials/secrets/{secretId}/lifecycle/deactivate | Deactivate an API Service Integration instance Secret
[**DeleteApiServiceIntegrationInstance**](ApiServiceIntegrationsApi.md#DeleteApiServiceIntegrationInstance) | **Delete** /integrations/api/v1/api-services/{apiServiceId} | Delete an API Service Integration instance
[**DeleteApiServiceIntegrationInstanceSecret**](ApiServiceIntegrationsApi.md#DeleteApiServiceIntegrationInstanceSecret) | **Delete** /integrations/api/v1/api-services/{apiServiceId}/credentials/secrets/{secretId} | Delete an API Service Integration instance Secret
[**GetApiServiceIntegrationInstance**](ApiServiceIntegrationsApi.md#GetApiServiceIntegrationInstance) | **Get** /integrations/api/v1/api-services/{apiServiceId} | Retrieve an API Service Integration instance
[**ListApiServiceIntegrationInstanceSecrets**](ApiServiceIntegrationsApi.md#ListApiServiceIntegrationInstanceSecrets) | **Get** /integrations/api/v1/api-services/{apiServiceId}/credentials/secrets | List all API Service Integration instance Secrets
[**ListApiServiceIntegrationInstances**](ApiServiceIntegrationsApi.md#ListApiServiceIntegrationInstances) | **Get** /integrations/api/v1/api-services | List all API Service Integration instances



## ActivateApiServiceIntegrationInstanceSecret

> APIServiceIntegrationInstanceSecret ActivateApiServiceIntegrationInstanceSecret(ctx, apiServiceId, secretId).Execute()

Activate an API Service Integration instance Secret



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance
    secretId := "ocs2f4zrZbs8nUa7p0g4" // string | `id` of the API Service Integration instance Secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.ActivateApiServiceIntegrationInstanceSecret(context.Background(), apiServiceId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.ActivateApiServiceIntegrationInstanceSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateApiServiceIntegrationInstanceSecret`: APIServiceIntegrationInstanceSecret
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.ActivateApiServiceIntegrationInstanceSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 
**secretId** | **string** | &#x60;id&#x60; of the API Service Integration instance Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateApiServiceIntegrationInstanceSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**APIServiceIntegrationInstanceSecret**](APIServiceIntegrationInstanceSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiServiceIntegrationInstance

> PostAPIServiceIntegrationInstance CreateApiServiceIntegrationInstance(ctx).PostAPIServiceIntegrationInstanceRequest(postAPIServiceIntegrationInstanceRequest).Execute()

Create an API Service Integration instance



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
    postAPIServiceIntegrationInstanceRequest := *openapiclient.NewPostAPIServiceIntegrationInstanceRequest([]string{"GrantedScopes_example"}, "my_app_cie") // PostAPIServiceIntegrationInstanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstance(context.Background()).PostAPIServiceIntegrationInstanceRequest(postAPIServiceIntegrationInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiServiceIntegrationInstance`: PostAPIServiceIntegrationInstance
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiServiceIntegrationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postAPIServiceIntegrationInstanceRequest** | [**PostAPIServiceIntegrationInstanceRequest**](PostAPIServiceIntegrationInstanceRequest.md) |  | 

### Return type

[**PostAPIServiceIntegrationInstance**](PostAPIServiceIntegrationInstance.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiServiceIntegrationInstanceSecret

> APIServiceIntegrationInstanceSecret CreateApiServiceIntegrationInstanceSecret(ctx, apiServiceId).Execute()

Create an API Service Integration instance Secret



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstanceSecret(context.Background(), apiServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstanceSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiServiceIntegrationInstanceSecret`: APIServiceIntegrationInstanceSecret
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.CreateApiServiceIntegrationInstanceSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiServiceIntegrationInstanceSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIServiceIntegrationInstanceSecret**](APIServiceIntegrationInstanceSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateApiServiceIntegrationInstanceSecret

> APIServiceIntegrationInstanceSecret DeactivateApiServiceIntegrationInstanceSecret(ctx, apiServiceId, secretId).Execute()

Deactivate an API Service Integration instance Secret



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance
    secretId := "ocs2f4zrZbs8nUa7p0g4" // string | `id` of the API Service Integration instance Secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.DeactivateApiServiceIntegrationInstanceSecret(context.Background(), apiServiceId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.DeactivateApiServiceIntegrationInstanceSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateApiServiceIntegrationInstanceSecret`: APIServiceIntegrationInstanceSecret
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.DeactivateApiServiceIntegrationInstanceSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 
**secretId** | **string** | &#x60;id&#x60; of the API Service Integration instance Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateApiServiceIntegrationInstanceSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**APIServiceIntegrationInstanceSecret**](APIServiceIntegrationInstanceSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiServiceIntegrationInstance

> DeleteApiServiceIntegrationInstance(ctx, apiServiceId).Execute()

Delete an API Service Integration instance



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApiServiceIntegrationsApi.DeleteApiServiceIntegrationInstance(context.Background(), apiServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.DeleteApiServiceIntegrationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiServiceIntegrationInstanceRequest struct via the builder pattern


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


## DeleteApiServiceIntegrationInstanceSecret

> DeleteApiServiceIntegrationInstanceSecret(ctx, apiServiceId, secretId).Execute()

Delete an API Service Integration instance Secret



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance
    secretId := "ocs2f4zrZbs8nUa7p0g4" // string | `id` of the API Service Integration instance Secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApiServiceIntegrationsApi.DeleteApiServiceIntegrationInstanceSecret(context.Background(), apiServiceId, secretId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.DeleteApiServiceIntegrationInstanceSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 
**secretId** | **string** | &#x60;id&#x60; of the API Service Integration instance Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiServiceIntegrationInstanceSecretRequest struct via the builder pattern


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


## GetApiServiceIntegrationInstance

> APIServiceIntegrationInstance GetApiServiceIntegrationInstance(ctx, apiServiceId).Execute()

Retrieve an API Service Integration instance



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.GetApiServiceIntegrationInstance(context.Background(), apiServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.GetApiServiceIntegrationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiServiceIntegrationInstance`: APIServiceIntegrationInstance
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.GetApiServiceIntegrationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiServiceIntegrationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**APIServiceIntegrationInstance**](APIServiceIntegrationInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiServiceIntegrationInstanceSecrets

> []APIServiceIntegrationInstanceSecret ListApiServiceIntegrationInstanceSecrets(ctx, apiServiceId).Execute()

List all API Service Integration instance Secrets



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
    apiServiceId := "000lr2rLjZ6NsGn1P0g3" // string | `id` of the API Service Integration instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.ListApiServiceIntegrationInstanceSecrets(context.Background(), apiServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.ListApiServiceIntegrationInstanceSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiServiceIntegrationInstanceSecrets`: []APIServiceIntegrationInstanceSecret
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.ListApiServiceIntegrationInstanceSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiServiceId** | **string** | &#x60;id&#x60; of the API Service Integration instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiServiceIntegrationInstanceSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]APIServiceIntegrationInstanceSecret**](APIServiceIntegrationInstanceSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiServiceIntegrationInstances

> []APIServiceIntegrationInstance ListApiServiceIntegrationInstances(ctx).After(after).Execute()

List all API Service Integration instances



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
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination) for more information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiServiceIntegrationsApi.ListApiServiceIntegrationInstances(context.Background()).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiServiceIntegrationsApi.ListApiServiceIntegrationInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiServiceIntegrationInstances`: []APIServiceIntegrationInstance
    fmt.Fprintf(os.Stdout, "Response from `ApiServiceIntegrationsApi.ListApiServiceIntegrationInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiServiceIntegrationInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination) for more information. | 

### Return type

[**[]APIServiceIntegrationInstance**](APIServiceIntegrationInstance.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


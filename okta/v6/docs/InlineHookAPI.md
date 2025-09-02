# \InlineHookAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateInlineHook**](InlineHookAPI.md#ActivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/activate | Activate an inline hook
[**CreateInlineHook**](InlineHookAPI.md#CreateInlineHook) | **Post** /api/v1/inlineHooks | Create an inline hook
[**DeactivateInlineHook**](InlineHookAPI.md#DeactivateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/lifecycle/deactivate | Deactivate an inline hook
[**DeleteInlineHook**](InlineHookAPI.md#DeleteInlineHook) | **Delete** /api/v1/inlineHooks/{inlineHookId} | Delete an inline hook
[**ExecuteInlineHook**](InlineHookAPI.md#ExecuteInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId}/execute | Execute an inline hook
[**GetInlineHook**](InlineHookAPI.md#GetInlineHook) | **Get** /api/v1/inlineHooks/{inlineHookId} | Retrieve an inline hook
[**ListInlineHooks**](InlineHookAPI.md#ListInlineHooks) | **Get** /api/v1/inlineHooks | List all inline hooks
[**ReplaceInlineHook**](InlineHookAPI.md#ReplaceInlineHook) | **Put** /api/v1/inlineHooks/{inlineHookId} | Replace an inline hook
[**UpdateInlineHook**](InlineHookAPI.md#UpdateInlineHook) | **Post** /api/v1/inlineHooks/{inlineHookId} | Update an inline hook



## ActivateInlineHook

> InlineHook ActivateInlineHook(ctx, inlineHookId).Execute()

Activate an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.ActivateInlineHook(context.Background(), inlineHookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ActivateInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateInlineHook`: InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ActivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInlineHook

> InlineHookCreateResponse CreateInlineHook(ctx).InlineHookCreate(inlineHookCreate).Execute()

Create an inline hook



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
	inlineHookCreate := *openapiclient.NewInlineHookCreate() // InlineHookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.CreateInlineHook(context.Background()).InlineHookCreate(inlineHookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.CreateInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInlineHook`: InlineHookCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.CreateInlineHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineHookCreate** | [**InlineHookCreate**](InlineHookCreate.md) |  | 

### Return type

[**InlineHookCreateResponse**](InlineHookCreateResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateInlineHook

> InlineHook DeactivateInlineHook(ctx, inlineHookId).Execute()

Deactivate an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.DeactivateInlineHook(context.Background(), inlineHookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.DeactivateInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateInlineHook`: InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.DeactivateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInlineHook

> DeleteInlineHook(ctx, inlineHookId).Execute()

Delete an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InlineHookAPI.DeleteInlineHook(context.Background(), inlineHookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.DeleteInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInlineHookRequest struct via the builder pattern


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


## ExecuteInlineHook

> ExecuteInlineHook200Response ExecuteInlineHook(ctx, inlineHookId).PayloadData(payloadData).Execute()

Execute an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook
	payloadData := openapiclient.executeInlineHook_request{PasswordImportRequestExecute: openapiclient.NewPasswordImportRequestExecute()} // ExecuteInlineHookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.ExecuteInlineHook(context.Background(), inlineHookId).PayloadData(payloadData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ExecuteInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteInlineHook`: ExecuteInlineHook200Response
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ExecuteInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payloadData** | [**ExecuteInlineHookRequest**](ExecuteInlineHookRequest.md) |  | 

### Return type

[**ExecuteInlineHook200Response**](ExecuteInlineHook200Response.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInlineHook

> InlineHook GetInlineHook(ctx, inlineHookId).Execute()

Retrieve an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.GetInlineHook(context.Background(), inlineHookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.GetInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInlineHook`: InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.GetInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInlineHooks

> []InlineHook ListInlineHooks(ctx).Type_(type_).Execute()

List all inline hooks



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
	type_ := "type__example" // string | One of the supported inline hook types (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.ListInlineHooks(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ListInlineHooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInlineHooks`: []InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ListInlineHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInlineHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | One of the supported inline hook types | 

### Return type

[**[]InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceInlineHook

> InlineHook ReplaceInlineHook(ctx, inlineHookId).InlineHook(inlineHook).Execute()

Replace an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook
	inlineHook := *openapiclient.NewInlineHookReplace() // InlineHookReplace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.ReplaceInlineHook(context.Background(), inlineHookId).InlineHook(inlineHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.ReplaceInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceInlineHook`: InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.ReplaceInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineHook** | [**InlineHookReplace**](InlineHookReplace.md) |  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInlineHook

> InlineHook UpdateInlineHook(ctx, inlineHookId).InlineHook(inlineHook).Execute()

Update an inline hook



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
	inlineHookId := "Y7Rzrd4g4xj6WdKzrBHH" // string | `id` of the inline hook
	inlineHook := *openapiclient.NewInlineHookReplace() // InlineHookReplace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InlineHookAPI.UpdateInlineHook(context.Background(), inlineHookId).InlineHook(inlineHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InlineHookAPI.UpdateInlineHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInlineHook`: InlineHook
	fmt.Fprintf(os.Stdout, "Response from `InlineHookAPI.UpdateInlineHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineHookId** | **string** | &#x60;id&#x60; of the inline hook | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInlineHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineHook** | [**InlineHookReplace**](InlineHookReplace.md) |  | 

### Return type

[**InlineHook**](InlineHook.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


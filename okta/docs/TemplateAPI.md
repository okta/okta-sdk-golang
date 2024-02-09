# \TemplateAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsTemplate**](TemplateAPI.md#CreateSmsTemplate) | **Post** /api/v1/templates/sms | Create an SMS Template
[**DeleteSmsTemplate**](TemplateAPI.md#DeleteSmsTemplate) | **Delete** /api/v1/templates/sms/{templateId} | Delete an SMS Template
[**GetSmsTemplate**](TemplateAPI.md#GetSmsTemplate) | **Get** /api/v1/templates/sms/{templateId} | Retrieve an SMS Template
[**ListSmsTemplates**](TemplateAPI.md#ListSmsTemplates) | **Get** /api/v1/templates/sms | List all SMS Templates
[**ReplaceSmsTemplate**](TemplateAPI.md#ReplaceSmsTemplate) | **Put** /api/v1/templates/sms/{templateId} | Replace an SMS Template
[**UpdateSmsTemplate**](TemplateAPI.md#UpdateSmsTemplate) | **Post** /api/v1/templates/sms/{templateId} | Update an SMS Template



## CreateSmsTemplate

> SmsTemplate CreateSmsTemplate(ctx).SmsTemplate(smsTemplate).Execute()

Create an SMS Template



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
    smsTemplate := *openapiclient.NewSmsTemplate() // SmsTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.CreateSmsTemplate(context.Background()).SmsTemplate(smsTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.CreateSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmsTemplate`: SmsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.CreateSmsTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsTemplate** | [**SmsTemplate**](SmsTemplate.md) |  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmsTemplate

> DeleteSmsTemplate(ctx, templateId).Execute()

Delete an SMS Template



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
    templateId := "6NQUJ5yR3bpgEiYmq8IC" // string | `id` of the Template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplateAPI.DeleteSmsTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.DeleteSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | &#x60;id&#x60; of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsTemplateRequest struct via the builder pattern


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


## GetSmsTemplate

> SmsTemplate GetSmsTemplate(ctx, templateId).Execute()

Retrieve an SMS Template



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
    templateId := "6NQUJ5yR3bpgEiYmq8IC" // string | `id` of the Template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.GetSmsTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.GetSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmsTemplate`: SmsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.GetSmsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | &#x60;id&#x60; of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsTemplates

> []SmsTemplate ListSmsTemplates(ctx).TemplateType(templateType).Execute()

List all SMS Templates



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
    templateType := "templateType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.ListSmsTemplates(context.Background()).TemplateType(templateType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.ListSmsTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmsTemplates`: []SmsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.ListSmsTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSmsTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateType** | **string** |  | 

### Return type

[**[]SmsTemplate**](SmsTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSmsTemplate

> SmsTemplate ReplaceSmsTemplate(ctx, templateId).SmsTemplate(smsTemplate).Execute()

Replace an SMS Template



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
    templateId := "6NQUJ5yR3bpgEiYmq8IC" // string | `id` of the Template
    smsTemplate := *openapiclient.NewSmsTemplate() // SmsTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.ReplaceSmsTemplate(context.Background(), templateId).SmsTemplate(smsTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.ReplaceSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSmsTemplate`: SmsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.ReplaceSmsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | &#x60;id&#x60; of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsTemplate** | [**SmsTemplate**](SmsTemplate.md) |  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmsTemplate

> SmsTemplate UpdateSmsTemplate(ctx, templateId).SmsTemplate(smsTemplate).Execute()

Update an SMS Template



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
    templateId := "6NQUJ5yR3bpgEiYmq8IC" // string | `id` of the Template
    smsTemplate := *openapiclient.NewSmsTemplate() // SmsTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateAPI.UpdateSmsTemplate(context.Background(), templateId).SmsTemplate(smsTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAPI.UpdateSmsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmsTemplate`: SmsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplateAPI.UpdateSmsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | &#x60;id&#x60; of the Template | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsTemplate** | [**SmsTemplate**](SmsTemplate.md) |  | 

### Return type

[**SmsTemplate**](SmsTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


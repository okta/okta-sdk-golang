# \CustomTemplatesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCustomization**](CustomTemplatesAPI.md#CreateEmailCustomization) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Create an Email Customization
[**DeleteAllCustomizations**](CustomTemplatesAPI.md#DeleteAllCustomizations) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | Delete all Email Customizations
[**DeleteEmailCustomization**](CustomTemplatesAPI.md#DeleteEmailCustomization) | **Delete** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Delete an Email Customization
[**GetCustomizationPreview**](CustomTemplatesAPI.md#GetCustomizationPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId}/preview | Retrieve a Preview of an Email Customization
[**GetEmailCustomization**](CustomTemplatesAPI.md#GetEmailCustomization) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Retrieve an Email Customization
[**GetEmailDefaultContent**](CustomTemplatesAPI.md#GetEmailDefaultContent) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content | Retrieve an Email Template Default Content
[**GetEmailDefaultPreview**](CustomTemplatesAPI.md#GetEmailDefaultPreview) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/default-content/preview | Retrieve a Preview of the Email Template default content
[**GetEmailSettings**](CustomTemplatesAPI.md#GetEmailSettings) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Retrieve the Email Template Settings
[**GetEmailTemplate**](CustomTemplatesAPI.md#GetEmailTemplate) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName} | Retrieve an Email Template
[**ListEmailCustomizations**](CustomTemplatesAPI.md#ListEmailCustomizations) | **Get** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations | List all Email Customizations
[**ListEmailTemplates**](CustomTemplatesAPI.md#ListEmailTemplates) | **Get** /api/v1/brands/{brandId}/templates/email | List all Email Templates
[**ReplaceEmailCustomization**](CustomTemplatesAPI.md#ReplaceEmailCustomization) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/customizations/{customizationId} | Replace an Email Customization
[**ReplaceEmailSettings**](CustomTemplatesAPI.md#ReplaceEmailSettings) | **Put** /api/v1/brands/{brandId}/templates/email/{templateName}/settings | Replace the Email Template Settings
[**SendTestEmail**](CustomTemplatesAPI.md#SendTestEmail) | **Post** /api/v1/brands/{brandId}/templates/email/{templateName}/test | Send a Test Email



## CreateEmailCustomization

> EmailCustomization CreateEmailCustomization(ctx, brandId, templateName).Instance(instance).Execute()

Create an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.CreateEmailCustomization(context.Background(), brandId, templateName).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.CreateEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.CreateEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instance** | [**EmailCustomization**](EmailCustomization.md) |  | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllCustomizations

> DeleteAllCustomizations(ctx, brandId, templateName).Execute()

Delete all Email Customizations



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomTemplatesAPI.DeleteAllCustomizations(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.DeleteAllCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCustomizationsRequest struct via the builder pattern


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


## DeleteEmailCustomization

> DeleteEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Delete an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomTemplatesAPI.DeleteEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.DeleteEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailCustomizationRequest struct via the builder pattern


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


## GetCustomizationPreview

> EmailPreview GetCustomizationPreview(ctx, brandId, templateName, customizationId).Execute()

Retrieve a Preview of an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetCustomizationPreview(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetCustomizationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomizationPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetCustomizationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomizationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailCustomization

> EmailCustomization GetEmailCustomization(ctx, brandId, templateName, customizationId).Execute()

Retrieve an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetEmailCustomization(context.Background(), brandId, templateName, customizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultContent

> EmailDefaultContent GetEmailDefaultContent(ctx, brandId, templateName).Language(language).Execute()

Retrieve an Email Template Default Content



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetEmailDefaultContent(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetEmailDefaultContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultContent`: EmailDefaultContent
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetEmailDefaultContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailDefaultContent**](EmailDefaultContent.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailDefaultPreview

> EmailPreview GetEmailDefaultPreview(ctx, brandId, templateName).Language(language).Execute()

Retrieve a Preview of the Email Template default content



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetEmailDefaultPreview(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetEmailDefaultPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailDefaultPreview`: EmailPreview
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetEmailDefaultPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailDefaultPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

### Return type

[**EmailPreview**](EmailPreview.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSettings

> EmailSettingsResponse GetEmailSettings(ctx, brandId, templateName).Execute()

Retrieve the Email Template Settings



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetEmailSettings(context.Background(), brandId, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailSettings`: EmailSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetEmailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailSettingsResponse**](EmailSettingsResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailTemplate

> EmailTemplateResponse GetEmailTemplate(ctx, brandId, templateName).Expand(expand).Execute()

Retrieve an Email Template



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.GetEmailTemplate(context.Background(), brandId, templateName).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.GetEmailTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmailTemplate`: EmailTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.GetEmailTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**EmailTemplateResponse**](EmailTemplateResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailCustomizations

> []EmailCustomization ListEmailCustomizations(ctx, brandId, templateName).After(after).Limit(limit).Execute()

List all Email Customizations



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.ListEmailCustomizations(context.Background(), brandId, templateName).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.ListEmailCustomizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailCustomizations`: []EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.ListEmailCustomizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailCustomizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

### Return type

[**[]EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEmailTemplates

> []EmailTemplateResponse ListEmailTemplates(ctx, brandId).After(after).Limit(limit).Expand(expand).Execute()

List all Email Templates



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
    brandId := "brandId_example" // string | The ID of the brand
    after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](/#pagination). (optional)
    limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
    expand := []string{"Expand_example"} // []string | Specifies additional metadata to be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.ListEmailTemplates(context.Background(), brandId).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.ListEmailTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailTemplates`: []EmailTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.ListEmailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **expand** | **[]string** | Specifies additional metadata to be included in the response | 

### Return type

[**[]EmailTemplateResponse**](EmailTemplateResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailCustomization

> EmailCustomization ReplaceEmailCustomization(ctx, brandId, templateName, customizationId).Instance(instance).Execute()

Replace an Email Customization



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    customizationId := "customizationId_example" // string | The ID of the email customization
    instance := *openapiclient.NewEmailCustomization("Body_example", "Subject_example", "Language_example") // EmailCustomization | Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.ReplaceEmailCustomization(context.Background(), brandId, templateName, customizationId).Instance(instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.ReplaceEmailCustomization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEmailCustomization`: EmailCustomization
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.ReplaceEmailCustomization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 
**customizationId** | **string** | The ID of the email customization | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEmailCustomizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **instance** | [**EmailCustomization**](EmailCustomization.md) | Request | 

### Return type

[**EmailCustomization**](EmailCustomization.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEmailSettings

> EmailSettings ReplaceEmailSettings(ctx, brandId, templateName).EmailSettings(emailSettings).Execute()

Replace the Email Template Settings



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    emailSettings := *openapiclient.NewEmailSettings("Recipients_example") // EmailSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomTemplatesAPI.ReplaceEmailSettings(context.Background(), brandId, templateName).EmailSettings(emailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.ReplaceEmailSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceEmailSettings`: EmailSettings
    fmt.Fprintf(os.Stdout, "Response from `CustomTemplatesAPI.ReplaceEmailSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEmailSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailSettings** | [**EmailSettings**](EmailSettings.md) |  | 

### Return type

[**EmailSettings**](EmailSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestEmail

> SendTestEmail(ctx, brandId, templateName).Language(language).Execute()

Send a Test Email



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
    brandId := "brandId_example" // string | The ID of the brand
    templateName := "templateName_example" // string | The name of the email template
    language := "language_example" // string | The language to use for the email. Defaults to the current user's language if unspecified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomTemplatesAPI.SendTestEmail(context.Background(), brandId, templateName).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomTemplatesAPI.SendTestEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** | The ID of the brand | 
**templateName** | **string** | The name of the email template | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** | The language to use for the email. Defaults to the current user&#39;s language if unspecified. | 

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


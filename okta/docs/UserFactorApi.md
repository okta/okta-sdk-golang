# \UserFactorApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](UserFactorApi.md#ActivateFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/lifecycle/activate | Activate a Factor
[**EnrollFactor**](UserFactorApi.md#EnrollFactor) | **Post** /api/v1/users/{userId}/factors | Enroll a Factor
[**GetFactor**](UserFactorApi.md#GetFactor) | **Get** /api/v1/users/{userId}/factors/{factorId} | Retrieve a Factor
[**GetFactorTransactionStatus**](UserFactorApi.md#GetFactorTransactionStatus) | **Get** /api/v1/users/{userId}/factors/{factorId}/transactions/{transactionId} | Retrieve a Factor Transaction Status
[**ListFactors**](UserFactorApi.md#ListFactors) | **Get** /api/v1/users/{userId}/factors | List all Factors
[**ListSupportedFactors**](UserFactorApi.md#ListSupportedFactors) | **Get** /api/v1/users/{userId}/factors/catalog | List all Supported Factors
[**ListSupportedSecurityQuestions**](UserFactorApi.md#ListSupportedSecurityQuestions) | **Get** /api/v1/users/{userId}/factors/questions | List all Supported Security Questions
[**ResendEnrollFactor**](UserFactorApi.md#ResendEnrollFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/resend | Resend a factor enrollment
[**UnenrollFactor**](UserFactorApi.md#UnenrollFactor) | **Delete** /api/v1/users/{userId}/factors/{factorId} | Unenroll a Factor
[**VerifyFactor**](UserFactorApi.md#VerifyFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/verify | Verify an MFA Factor



## ActivateFactor

> ListFactors200ResponseInner ActivateFactor(ctx, userId, factorId).Body(body).Execute()

Activate a Factor



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor
    body := *openapiclient.NewActivateFactorRequest() // ActivateFactorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.ActivateFactor(context.Background(), userId, factorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.ActivateFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.ActivateFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ActivateFactorRequest**](ActivateFactorRequest.md) |  | 

### Return type

[**ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollFactor

> ListFactors200ResponseInner EnrollFactor(ctx, userId).Body(body).UpdatePhone(updatePhone).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).Activate(activate).Execute()

Enroll a Factor



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
    userId := "userId_example" // string | 
    body := openapiclient.listFactors_200_response_inner{CallUserFactor: openapiclient.NewCallUserFactor()} // ListFactors200ResponseInner | Factor
    updatePhone := true // bool |  (optional) (default to false)
    templateId := "templateId_example" // string | id of SMS template (only for SMS factor) (optional)
    tokenLifetimeSeconds := int32(56) // int32 |  (optional) (default to 300)
    activate := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.EnrollFactor(context.Background(), userId).Body(body).UpdatePhone(updatePhone).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.EnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.EnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ListFactors200ResponseInner**](ListFactors200ResponseInner.md) | Factor | 
 **updatePhone** | **bool** |  | [default to false]
 **templateId** | **string** | id of SMS template (only for SMS factor) | 
 **tokenLifetimeSeconds** | **int32** |  | [default to 300]
 **activate** | **bool** |  | [default to false]

### Return type

[**ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFactor

> ListFactors200ResponseInner GetFactor(ctx, userId, factorId).Execute()

Retrieve a Factor



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.GetFactor(context.Background(), userId, factorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.GetFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.GetFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFactorTransactionStatus

> VerifyUserFactorResponse GetFactorTransactionStatus(ctx, userId, factorId, transactionId).Execute()

Retrieve a Factor Transaction Status



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor
    transactionId := "gPAQcN3NDjSGOCAeG2Jv" // string | `id` of the Transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.GetFactorTransactionStatus(context.Background(), userId, factorId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.GetFactorTransactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFactorTransactionStatus`: VerifyUserFactorResponse
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.GetFactorTransactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 
**transactionId** | **string** | &#x60;id&#x60; of the Transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorTransactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VerifyUserFactorResponse**](VerifyUserFactorResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFactors

> []ListFactors200ResponseInner ListFactors(ctx, userId).Execute()

List all Factors



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.ListFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.ListFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFactors`: []ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.ListFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedFactors

> []ListFactors200ResponseInner ListSupportedFactors(ctx, userId).Execute()

List all Supported Factors



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.ListSupportedFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.ListSupportedFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedFactors`: []ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.ListSupportedFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedFactorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedSecurityQuestions

> []SecurityQuestion ListSupportedSecurityQuestions(ctx, userId).Execute()

List all Supported Security Questions



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.ListSupportedSecurityQuestions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.ListSupportedSecurityQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedSecurityQuestions`: []SecurityQuestion
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.ListSupportedSecurityQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedSecurityQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SecurityQuestion**](SecurityQuestion.md)

### Authorization

[apiToken](../README.md#apiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEnrollFactor

> ListFactors200ResponseInner ResendEnrollFactor(ctx, userId, factorId).ListFactors200ResponseInner(listFactors200ResponseInner).TemplateId(templateId).Execute()

Resend a factor enrollment



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor
    listFactors200ResponseInner := openapiclient.listFactors_200_response_inner{CallUserFactor: openapiclient.NewCallUserFactor()} // ListFactors200ResponseInner | Factor
    templateId := "templateId_example" // string | ID of SMS template (only for SMS factor) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.ResendEnrollFactor(context.Background(), userId, factorId).ListFactors200ResponseInner(listFactors200ResponseInner).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.ResendEnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendEnrollFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.ResendEnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **listFactors200ResponseInner** | [**ListFactors200ResponseInner**](ListFactors200ResponseInner.md) | Factor | 
 **templateId** | **string** | ID of SMS template (only for SMS factor) | 

### Return type

[**ListFactors200ResponseInner**](ListFactors200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnenrollFactor

> UnenrollFactor(ctx, userId, factorId).RemoveRecoveryEnrollment(removeRecoveryEnrollment).Execute()

Unenroll a Factor



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor
    removeRecoveryEnrollment := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserFactorApi.UnenrollFactor(context.Background(), userId, factorId).RemoveRecoveryEnrollment(removeRecoveryEnrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.UnenrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeRecoveryEnrollment** | **bool** |  | [default to false]

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


## VerifyFactor

> VerifyUserFactorResponse VerifyFactor(ctx, userId, factorId).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).XForwardedFor(xForwardedFor).UserAgent(userAgent).AcceptLanguage(acceptLanguage).Body(body).Execute()

Verify an MFA Factor



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
    userId := "userId_example" // string | 
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | `id` of the Factor
    templateId := "templateId_example" // string |  (optional)
    tokenLifetimeSeconds := int32(56) // int32 |  (optional) (default to 300)
    xForwardedFor := "xForwardedFor_example" // string |  (optional)
    userAgent := "userAgent_example" // string |  (optional)
    acceptLanguage := "acceptLanguage_example" // string |  (optional)
    body := *openapiclient.NewVerifyFactorRequest() // VerifyFactorRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorApi.VerifyFactor(context.Background(), userId, factorId).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).XForwardedFor(xForwardedFor).UserAgent(userAgent).AcceptLanguage(acceptLanguage).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorApi.VerifyFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyFactor`: VerifyUserFactorResponse
    fmt.Fprintf(os.Stdout, "Response from `UserFactorApi.VerifyFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**factorId** | **string** | &#x60;id&#x60; of the Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateId** | **string** |  | 
 **tokenLifetimeSeconds** | **int32** |  | [default to 300]
 **xForwardedFor** | **string** |  | 
 **userAgent** | **string** |  | 
 **acceptLanguage** | **string** |  | 
 **body** | [**VerifyFactorRequest**](VerifyFactorRequest.md) |  | 

### Return type

[**VerifyUserFactorResponse**](VerifyUserFactorResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


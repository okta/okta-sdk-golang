# \UserFactorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](UserFactorAPI.md#ActivateFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/lifecycle/activate | Activate a Factor
[**EnrollFactor**](UserFactorAPI.md#EnrollFactor) | **Post** /api/v1/users/{userId}/factors | Enroll a Factor
[**GetFactor**](UserFactorAPI.md#GetFactor) | **Get** /api/v1/users/{userId}/factors/{factorId} | Retrieve a Factor
[**GetFactorTransactionStatus**](UserFactorAPI.md#GetFactorTransactionStatus) | **Get** /api/v1/users/{userId}/factors/{factorId}/transactions/{transactionId} | Retrieve a Factor transaction status
[**ListFactors**](UserFactorAPI.md#ListFactors) | **Get** /api/v1/users/{userId}/factors | List all enrolled Factors
[**ListSupportedFactors**](UserFactorAPI.md#ListSupportedFactors) | **Get** /api/v1/users/{userId}/factors/catalog | List all supported Factors
[**ListSupportedSecurityQuestions**](UserFactorAPI.md#ListSupportedSecurityQuestions) | **Get** /api/v1/users/{userId}/factors/questions | List all supported Security Questions
[**ResendEnrollFactor**](UserFactorAPI.md#ResendEnrollFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/resend | Resend a Factor enrollment
[**UnenrollFactor**](UserFactorAPI.md#UnenrollFactor) | **Delete** /api/v1/users/{userId}/factors/{factorId} | Unenroll a Factor
[**VerifyFactor**](UserFactorAPI.md#VerifyFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/verify | Verify a Factor



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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor
    body := *openapiclient.NewUserFactorActivateRequest() // UserFactorActivateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ActivateFactor(context.Background(), userId, factorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ActivateFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ActivateFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UserFactorActivateRequest**](UserFactorActivateRequest.md) |  | 

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
    userId := "userId_example" // string | ID of an existing Okta user
    body := openapiclient.listFactors_200_response_inner{UserFactorCall: openapiclient.NewUserFactorCall()} // ListFactors200ResponseInner | Factor
    updatePhone := true // bool | If `true`, indicates that you'll update the `phoneNumber`. Only used for `sms` Factors that are pending activation. (optional) (default to false)
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. (optional)
    tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
    activate := true // bool | If `true`, the `sms` Factor is immediately activated as part of the enrollment. An activation text message isn't sent to the device. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.EnrollFactor(context.Background(), userId).Body(body).UpdatePhone(updatePhone).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.EnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.EnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ListFactors200ResponseInner**](ListFactors200ResponseInner.md) | Factor | 
 **updatePhone** | **bool** | If &#x60;true&#x60;, indicates that you&#39;ll update the &#x60;phoneNumber&#x60;. Only used for &#x60;sms&#x60; Factors that are pending activation. | [default to false]
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **activate** | **bool** | If &#x60;true&#x60;, the &#x60;sms&#x60; Factor is immediately activated as part of the enrollment. An activation text message isn&#39;t sent to the device. | [default to false]

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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.GetFactor(context.Background(), userId, factorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.GetFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.GetFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 

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

> UserFactorVerifyResponse GetFactorTransactionStatus(ctx, userId, factorId, transactionId).Execute()

Retrieve a Factor transaction status



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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor
    transactionId := "gPAQcN3NDjSGOCAeG2Jv" // string | ID of an existing Factor verification transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.GetFactorTransactionStatus(context.Background(), userId, factorId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.GetFactorTransactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFactorTransactionStatus`: UserFactorVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.GetFactorTransactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 
**transactionId** | **string** | ID of an existing Factor verification transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorTransactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserFactorVerifyResponse**](UserFactorVerifyResponse.md)

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

List all enrolled Factors



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ListFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ListFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFactors`: []ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ListFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

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

List all supported Factors



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ListSupportedFactors(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ListSupportedFactors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedFactors`: []ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ListSupportedFactors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

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

> []UserFactorSecurityQuestionProfile ListSupportedSecurityQuestions(ctx, userId).Execute()

List all supported Security Questions



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
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ListSupportedSecurityQuestions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ListSupportedSecurityQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedSecurityQuestions`: []UserFactorSecurityQuestionProfile
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ListSupportedSecurityQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedSecurityQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserFactorSecurityQuestionProfile**](UserFactorSecurityQuestionProfile.md)

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

Resend a Factor enrollment



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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor
    listFactors200ResponseInner := openapiclient.listFactors_200_response_inner{UserFactorCall: openapiclient.NewUserFactorCall()} // ListFactors200ResponseInner | 
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ResendEnrollFactor(context.Background(), userId, factorId).ListFactors200ResponseInner(listFactors200ResponseInner).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ResendEnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendEnrollFactor`: ListFactors200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ResendEnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **listFactors200ResponseInner** | [**ListFactors200ResponseInner**](ListFactors200ResponseInner.md) |  | 
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. | 

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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor
    removeRecoveryEnrollment := true // bool | If `true`, removes the the phone number as both a recovery method and a Factor. Only used for `sms` and `call` Factors. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserFactorAPI.UnenrollFactor(context.Background(), userId, factorId).RemoveRecoveryEnrollment(removeRecoveryEnrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.UnenrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeRecoveryEnrollment** | **bool** | If &#x60;true&#x60;, removes the the phone number as both a recovery method and a Factor. Only used for &#x60;sms&#x60; and &#x60;call&#x60; Factors. | [default to false]

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

> UserFactorVerifyResponse VerifyFactor(ctx, userId, factorId).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).XForwardedFor(xForwardedFor).UserAgent(userAgent).AcceptLanguage(acceptLanguage).Body(body).Execute()

Verify a Factor



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
    userId := "userId_example" // string | ID of an existing Okta user
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing User Factor
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. (optional)
    tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
    xForwardedFor := "xForwardedFor_example" // string | Public IP address for the user agent (optional)
    userAgent := "userAgent_example" // string | Type of user agent detected when the request is made (optional)
    acceptLanguage := "acceptLanguage_example" // string | Sets a two-letter language code that defines a localized message to send. Only used by the `sms` Factor.  * If the language code doesn't exist in the SMS template, the message uses the default template. * If the `templateId` doesn't exist, the message is sent using the default template. (optional)
    body := *openapiclient.NewUserFactorVerifyRequest() // UserFactorVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.VerifyFactor(context.Background(), userId, factorId).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).XForwardedFor(xForwardedFor).UserAgent(userAgent).AcceptLanguage(acceptLanguage).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.VerifyFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyFactor`: UserFactorVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.VerifyFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing User Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **xForwardedFor** | **string** | Public IP address for the user agent | 
 **userAgent** | **string** | Type of user agent detected when the request is made | 
 **acceptLanguage** | **string** | Sets a two-letter language code that defines a localized message to send. Only used by the &#x60;sms&#x60; Factor.  * If the language code doesn&#39;t exist in the SMS template, the message uses the default template. * If the &#x60;templateId&#x60; doesn&#39;t exist, the message is sent using the default template. | 
 **body** | [**UserFactorVerifyRequest**](UserFactorVerifyRequest.md) |  | 

### Return type

[**UserFactorVerifyResponse**](UserFactorVerifyResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


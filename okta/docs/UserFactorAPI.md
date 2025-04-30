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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

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
**factorId** | **string** | ID of an existing user Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

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

> ListFactors200ResponseInner EnrollFactor(ctx, userId).Body(body).UpdatePhone(updatePhone).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).Activate(activate).AcceptLanguage(acceptLanguage).Execute()

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
    updatePhone := true // bool | If `true`, indicates you are replacing the currently registered phone number for the specified user. This parameter is ignored if the existing phone number is used by an activated Factor. (optional) (default to false)
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. If the provided ID doesn't exist, the default template is used instead. (optional)
    tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
    activate := true // bool | If `true`, the `sms` Factor is immediately activated as part of the enrollment. An activation text message isn't sent to the device. (optional) (default to false)
    acceptLanguage := "fr" // string | An ISO 639-1 two-letter language code that defines a localized message to send. Only used by `sms` Factors. If a localized message doesn't exist or the `templateId` is incorrect, the default template is used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.EnrollFactor(context.Background(), userId).Body(body).UpdatePhone(updatePhone).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).Activate(activate).AcceptLanguage(acceptLanguage).Execute()
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
 **updatePhone** | **bool** | If &#x60;true&#x60;, indicates you are replacing the currently registered phone number for the specified user. This parameter is ignored if the existing phone number is used by an activated Factor. | [default to false]
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. If the provided ID doesn&#39;t exist, the default template is used instead. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **activate** | **bool** | If &#x60;true&#x60;, the &#x60;sms&#x60; Factor is immediately activated as part of the enrollment. An activation text message isn&#39;t sent to the device. | [default to false]
 **acceptLanguage** | **string** | An ISO 639-1 two-letter language code that defines a localized message to send. Only used by &#x60;sms&#x60; Factors. If a localized message doesn&#39;t exist or the &#x60;templateId&#x60; is incorrect, the default template is used instead. | 

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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor

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
**factorId** | **string** | ID of an existing user Factor | 

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

> GetFactorTransactionStatus200Response GetFactorTransactionStatus(ctx, userId, factorId, transactionId).Execute()

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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor
    transactionId := "gPAQcN3NDjSGOCAeG2Jv" // string | ID of an existing Factor verification transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.GetFactorTransactionStatus(context.Background(), userId, factorId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.GetFactorTransactionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFactorTransactionStatus`: GetFactorTransactionStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.GetFactorTransactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing user Factor | 
**transactionId** | **string** | ID of an existing Factor verification transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorTransactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFactorTransactionStatus200Response**](GetFactorTransactionStatus200Response.md)

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

> []UserFactorSupported ListSupportedFactors(ctx, userId).Execute()

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
    // response from `ListSupportedFactors`: []UserFactorSupported
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

[**[]UserFactorSupported**](UserFactorSupported.md)

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

> ResendEnrollFactorRequest ResendEnrollFactor(ctx, userId, factorId).ResendEnrollFactorRequest(resendEnrollFactorRequest).TemplateId(templateId).Execute()

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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor
    resendEnrollFactorRequest := openapiclient.resendEnrollFactor_request{UserFactorCall: openapiclient.NewUserFactorCall()} // ResendEnrollFactorRequest | 
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserFactorAPI.ResendEnrollFactor(context.Background(), userId, factorId).ResendEnrollFactorRequest(resendEnrollFactorRequest).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ResendEnrollFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendEnrollFactor`: ResendEnrollFactorRequest
    fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ResendEnrollFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing user Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resendEnrollFactorRequest** | [**ResendEnrollFactorRequest**](ResendEnrollFactorRequest.md) |  | 
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. | 

### Return type

[**ResendEnrollFactorRequest**](ResendEnrollFactorRequest.md)

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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor
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
**factorId** | **string** | ID of an existing user Factor | 

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
    factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user Factor
    templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by `sms` Factors. (optional)
    tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
    xForwardedFor := "xForwardedFor_example" // string | Public IP address for the user agent (optional)
    userAgent := "userAgent_example" // string | Type of user agent detected when the request is made. Required to verify `push` Factors. (optional)
    acceptLanguage := "fr" // string | An ISO 639-1 two-letter language code that defines a localized message to send. Only used by `sms` Factors. If a localized message doesn't exist or the `templateId` is incorrect, the default template is used instead. (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} | Some Factors (`call`, `email`, `push`, `sms`, `u2f`, and `webauthn`) must first issue a challenge before you can verify the Factor. Do this by making a request without a body. After a challenge is issued, make another request to verify the Factor. (optional)

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
**factorId** | **string** | ID of an existing user Factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). Only used by &#x60;sms&#x60; Factors. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **xForwardedFor** | **string** | Public IP address for the user agent | 
 **userAgent** | **string** | Type of user agent detected when the request is made. Required to verify &#x60;push&#x60; Factors. | 
 **acceptLanguage** | **string** | An ISO 639-1 two-letter language code that defines a localized message to send. Only used by &#x60;sms&#x60; Factors. If a localized message doesn&#39;t exist or the &#x60;templateId&#x60; is incorrect, the default template is used instead. | 
 **body** | **map[string]interface{}** | Some Factors (&#x60;call&#x60;, &#x60;email&#x60;, &#x60;push&#x60;, &#x60;sms&#x60;, &#x60;u2f&#x60;, and &#x60;webauthn&#x60;) must first issue a challenge before you can verify the Factor. Do this by making a request without a body. After a challenge is issued, make another request to verify the Factor. | 

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


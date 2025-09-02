# \UserFactorAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateFactor**](UserFactorAPI.md#ActivateFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/lifecycle/activate | Activate a factor
[**EnrollFactor**](UserFactorAPI.md#EnrollFactor) | **Post** /api/v1/users/{userId}/factors | Enroll a factor
[**GetFactor**](UserFactorAPI.md#GetFactor) | **Get** /api/v1/users/{userId}/factors/{factorId} | Retrieve a factor
[**GetFactorTransactionStatus**](UserFactorAPI.md#GetFactorTransactionStatus) | **Get** /api/v1/users/{userId}/factors/{factorId}/transactions/{transactionId} | Retrieve a factor transaction status
[**GetYubikeyOtpTokenById**](UserFactorAPI.md#GetYubikeyOtpTokenById) | **Get** /api/v1/org/factors/yubikey_token/tokens/{tokenId} | Retrieve a YubiKey OTP token
[**ListFactors**](UserFactorAPI.md#ListFactors) | **Get** /api/v1/users/{userId}/factors | List all enrolled factors
[**ListSupportedFactors**](UserFactorAPI.md#ListSupportedFactors) | **Get** /api/v1/users/{userId}/factors/catalog | List all supported factors
[**ListSupportedSecurityQuestions**](UserFactorAPI.md#ListSupportedSecurityQuestions) | **Get** /api/v1/users/{userId}/factors/questions | List all supported security questions
[**ListYubikeyOtpTokens**](UserFactorAPI.md#ListYubikeyOtpTokens) | **Get** /api/v1/org/factors/yubikey_token/tokens | List all YubiKey OTP tokens
[**ResendEnrollFactor**](UserFactorAPI.md#ResendEnrollFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/resend | Resend a factor enrollment
[**UnenrollFactor**](UserFactorAPI.md#UnenrollFactor) | **Delete** /api/v1/users/{userId}/factors/{factorId} | Unenroll a factor
[**UploadYubikeyOtpTokenSeed**](UserFactorAPI.md#UploadYubikeyOtpTokenSeed) | **Post** /api/v1/org/factors/yubikey_token/tokens | Upload a YubiKey OTP seed
[**VerifyFactor**](UserFactorAPI.md#VerifyFactor) | **Post** /api/v1/users/{userId}/factors/{factorId}/verify | Verify a factor



## ActivateFactor

> ActivateFactor200Response ActivateFactor(ctx, userId, factorId).Body(body).Execute()

Activate a factor



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor
	body := openapiclient.UserFactorActivateRequest{Call: openapiclient.NewCall()} // UserFactorActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFactorAPI.ActivateFactor(context.Background(), userId, factorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ActivateFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateFactor`: ActivateFactor200Response
	fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ActivateFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing user factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UserFactorActivateRequest**](UserFactorActivateRequest.md) |  | 

### Return type

[**ActivateFactor200Response**](ActivateFactor200Response.md)

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

Enroll a factor



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	body := openapiclient.listFactors_200_response_inner{UserFactorCall: openapiclient.NewUserFactorCall()} // ListFactors200ResponseInner | Factor
	updatePhone := true // bool | If `true`, indicates that you are replacing the currently registered phone number for the specified user. This parameter is ignored if the existing phone number is used by an activated factor. (optional) (default to false)
	templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by `sms` factors. If the provided ID doesn't exist, the default template is used instead. (optional)
	tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
	activate := true // bool | If `true`, the factor is immediately activated as part of the enrollment. An activation process isn't required. Currently auto-activation is supported by `sms`, `call`, `email` and `token:hotp` (Custom TOTP) factors. (optional) (default to false)
	acceptLanguage := "fr" // string | An ISO 639-1 two-letter language code that defines a localized message to send. This parameter is only used by `sms` factors. If a localized message doesn't exist or the `templateId` is incorrect, the default template is used instead. (optional)

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
 **updatePhone** | **bool** | If &#x60;true&#x60;, indicates that you are replacing the currently registered phone number for the specified user. This parameter is ignored if the existing phone number is used by an activated factor. | [default to false]
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by &#x60;sms&#x60; factors. If the provided ID doesn&#39;t exist, the default template is used instead. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **activate** | **bool** | If &#x60;true&#x60;, the factor is immediately activated as part of the enrollment. An activation process isn&#39;t required. Currently auto-activation is supported by &#x60;sms&#x60;, &#x60;call&#x60;, &#x60;email&#x60; and &#x60;token:hotp&#x60; (Custom TOTP) factors. | [default to false]
 **acceptLanguage** | **string** | An ISO 639-1 two-letter language code that defines a localized message to send. This parameter is only used by &#x60;sms&#x60; factors. If a localized message doesn&#39;t exist or the &#x60;templateId&#x60; is incorrect, the default template is used instead. | 

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

> UserFactor GetFactor(ctx, userId, factorId).Execute()

Retrieve a factor



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFactorAPI.GetFactor(context.Background(), userId, factorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.GetFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFactor`: UserFactor
	fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.GetFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**factorId** | **string** | ID of an existing user factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserFactor**](UserFactor.md)

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

Retrieve a factor transaction status



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor
	transactionId := "gPAQcN3NDjSGOCAeG2Jv" // string | ID of an existing factor verification transaction

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
**factorId** | **string** | ID of an existing user factor | 
**transactionId** | **string** | ID of an existing factor verification transaction | 

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


## GetYubikeyOtpTokenById

> UserFactorYubikeyOtpToken GetYubikeyOtpTokenById(ctx, tokenId).Execute()

Retrieve a YubiKey OTP token



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
	tokenId := "ykkxdtCA1fKVxyu6R0g3" // string | ID of a YubiKey token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFactorAPI.GetYubikeyOtpTokenById(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.GetYubikeyOtpTokenById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYubikeyOtpTokenById`: UserFactorYubikeyOtpToken
	fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.GetYubikeyOtpTokenById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | ID of a YubiKey token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYubikeyOtpTokenByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserFactorYubikeyOtpToken**](UserFactorYubikeyOtpToken.md)

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

List all enrolled factors



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

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

List all supported factors



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

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

List all supported security questions



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

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


## ListYubikeyOtpTokens

> []UserFactorYubikeyOtpToken ListYubikeyOtpTokens(ctx).After(after).Expand(expand).Filter(filter).ForDownload(forDownload).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

List all YubiKey OTP tokens



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
	after := "after_example" // string | Specifies the pagination cursor for the next page of tokens (optional)
	expand := "expand_example" // string | Embeds the [user](/openapi/okta-management/management/tag/User/) resource if the YubiKey token is assigned to a user and `expand` is set to `user` (optional)
	filter := "filter_example" // string | The expression used to filter tokens (optional)
	forDownload := true // bool | Returns tokens in a CSV to download instead of in the response. When you use this query parameter, the `limit` default changes to 1000. (optional) (default to false)
	limit := int32(56) // int32 | Specifies the number of results per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | The value of how the tokens are sorted (optional)
	sortOrder := "sortOrder_example" // string | Specifies the sort order, either `ASC` or `DESC` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFactorAPI.ListYubikeyOtpTokens(context.Background()).After(after).Expand(expand).Filter(filter).ForDownload(forDownload).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.ListYubikeyOtpTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListYubikeyOtpTokens`: []UserFactorYubikeyOtpToken
	fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.ListYubikeyOtpTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListYubikeyOtpTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Specifies the pagination cursor for the next page of tokens | 
 **expand** | **string** | Embeds the [user](/openapi/okta-management/management/tag/User/) resource if the YubiKey token is assigned to a user and &#x60;expand&#x60; is set to &#x60;user&#x60; | 
 **filter** | **string** | The expression used to filter tokens | 
 **forDownload** | **bool** | Returns tokens in a CSV to download instead of in the response. When you use this query parameter, the &#x60;limit&#x60; default changes to 1000. | [default to false]
 **limit** | **int32** | Specifies the number of results per page | [default to 20]
 **sortBy** | **string** | The value of how the tokens are sorted | 
 **sortOrder** | **string** | Specifies the sort order, either &#x60;ASC&#x60; or &#x60;DESC&#x60; | 

### Return type

[**[]UserFactorYubikeyOtpToken**](UserFactorYubikeyOtpToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEnrollFactor

> ResendEnrollFactorRequest ResendEnrollFactor(ctx, userId, factorId).ResendEnrollFactorRequest(resendEnrollFactorRequest).TemplateId(templateId).Execute()

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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor
	resendEnrollFactorRequest := openapiclient.resendEnrollFactor_request{UserFactorCall: openapiclient.NewUserFactorCall()} // ResendEnrollFactorRequest | 
	templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by `sms` factors. (optional)

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
**factorId** | **string** | ID of an existing user factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEnrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resendEnrollFactorRequest** | [**ResendEnrollFactorRequest**](ResendEnrollFactorRequest.md) |  | 
 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by &#x60;sms&#x60; factors. | 

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

Unenroll a factor



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor
	removeRecoveryEnrollment := true // bool | If `true`, removes the phone number as both a recovery method and a factor. This parameter is only used for the `sms` and `call` factors. (optional) (default to false)

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
**factorId** | **string** | ID of an existing user factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeRecoveryEnrollment** | **bool** | If &#x60;true&#x60;, removes the phone number as both a recovery method and a factor. This parameter is only used for the &#x60;sms&#x60; and &#x60;call&#x60; factors. | [default to false]

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


## UploadYubikeyOtpTokenSeed

> UserFactorYubikeyOtpToken UploadYubikeyOtpTokenSeed(ctx).UploadYubikeyOtpTokenSeedRequest(uploadYubikeyOtpTokenSeedRequest).After(after).Expand(expand).Filter(filter).ForDownload(forDownload).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

Upload a YubiKey OTP seed



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
	uploadYubikeyOtpTokenSeedRequest := *openapiclient.NewUploadYubikeyOtpTokenSeedRequest() // UploadYubikeyOtpTokenSeedRequest | 
	after := "after_example" // string | Specifies the pagination cursor for the next page of tokens (optional)
	expand := "expand_example" // string | Embeds the [user](/openapi/okta-management/management/tag/User/) resource if the YubiKey token is assigned to a user and `expand` is set to `user` (optional)
	filter := "filter_example" // string | The expression used to filter tokens (optional)
	forDownload := true // bool | Returns tokens in a CSV to download instead of in the response. When you use this query parameter, the `limit` default changes to 1000. (optional) (default to false)
	limit := int32(56) // int32 | Specifies the number of results per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | The value of how the tokens are sorted (optional)
	sortOrder := "sortOrder_example" // string | Specifies the sort order, either `ASC` or `DESC` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFactorAPI.UploadYubikeyOtpTokenSeed(context.Background()).UploadYubikeyOtpTokenSeedRequest(uploadYubikeyOtpTokenSeedRequest).After(after).Expand(expand).Filter(filter).ForDownload(forDownload).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFactorAPI.UploadYubikeyOtpTokenSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadYubikeyOtpTokenSeed`: UserFactorYubikeyOtpToken
	fmt.Fprintf(os.Stdout, "Response from `UserFactorAPI.UploadYubikeyOtpTokenSeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadYubikeyOtpTokenSeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadYubikeyOtpTokenSeedRequest** | [**UploadYubikeyOtpTokenSeedRequest**](UploadYubikeyOtpTokenSeedRequest.md) |  | 
 **after** | **string** | Specifies the pagination cursor for the next page of tokens | 
 **expand** | **string** | Embeds the [user](/openapi/okta-management/management/tag/User/) resource if the YubiKey token is assigned to a user and &#x60;expand&#x60; is set to &#x60;user&#x60; | 
 **filter** | **string** | The expression used to filter tokens | 
 **forDownload** | **bool** | Returns tokens in a CSV to download instead of in the response. When you use this query parameter, the &#x60;limit&#x60; default changes to 1000. | [default to false]
 **limit** | **int32** | Specifies the number of results per page | [default to 20]
 **sortBy** | **string** | The value of how the tokens are sorted | 
 **sortOrder** | **string** | Specifies the sort order, either &#x60;ASC&#x60; or &#x60;DESC&#x60; | 

### Return type

[**UserFactorYubikeyOtpToken**](UserFactorYubikeyOtpToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyFactor

> UserFactorVerifyResponse VerifyFactor(ctx, userId, factorId).TemplateId(templateId).TokenLifetimeSeconds(tokenLifetimeSeconds).XForwardedFor(xForwardedFor).UserAgent(userAgent).AcceptLanguage(acceptLanguage).Body(body).Execute()

Verify a factor



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	factorId := "zAgrsaBe0wVGRugDYtdv" // string | ID of an existing user factor
	templateId := "cstk2flOtuCMDJK4b0g3" // string | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by `sms` factors. (optional)
	tokenLifetimeSeconds := int32(56) // int32 | Defines how long the token remains valid (optional) (default to 300)
	xForwardedFor := "xForwardedFor_example" // string | Public IP address for the user agent (optional)
	userAgent := "userAgent_example" // string | Type of user agent detected when the request is made. Required to verify `push` factors. (optional)
	acceptLanguage := "fr" // string | An ISO 639-1 two-letter language code that defines a localized message to send. This parameter is only used by `sms` factors. If a localized message doesn't exist or the `templateId` is incorrect, the default template is used instead. (optional)
	body := openapiclient.UserFactorVerifyRequest{Call1: openapiclient.NewCall1()} // UserFactorVerifyRequest | Verifies an OTP for a factor. Some factors (`call`, `email`, `push`, `sms`, `u2f`, and `webauthn`) must first issue a challenge before you can verify the factor. Do this by making a request without a body. After a challenge is issued, make another request to verify the factor.  > **Note:** > Unlike standard push challenges that don't require a request body, a number matching [`push`](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor!path=2/useNumberMatchingChallenge&t=request) challenge requires a request body. `useNumberMatchingChallenge` must be set to `true`. > When a number matching challenge is issued for an Okta Verify `push` factor enrollment, a `correctAnswer` challenge object is returned in the [`_embedded`](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor!c=200&path=_embedded&t=response) object. (optional)

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
**factorId** | **string** | ID of an existing user factor | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateId** | **string** | ID of an existing custom SMS template. See the [SMS Templates API](../Template). This parameter is only used by &#x60;sms&#x60; factors. | 
 **tokenLifetimeSeconds** | **int32** | Defines how long the token remains valid | [default to 300]
 **xForwardedFor** | **string** | Public IP address for the user agent | 
 **userAgent** | **string** | Type of user agent detected when the request is made. Required to verify &#x60;push&#x60; factors. | 
 **acceptLanguage** | **string** | An ISO 639-1 two-letter language code that defines a localized message to send. This parameter is only used by &#x60;sms&#x60; factors. If a localized message doesn&#39;t exist or the &#x60;templateId&#x60; is incorrect, the default template is used instead. | 
 **body** | [**UserFactorVerifyRequest**](UserFactorVerifyRequest.md) | Verifies an OTP for a factor. Some factors (&#x60;call&#x60;, &#x60;email&#x60;, &#x60;push&#x60;, &#x60;sms&#x60;, &#x60;u2f&#x60;, and &#x60;webauthn&#x60;) must first issue a challenge before you can verify the factor. Do this by making a request without a body. After a challenge is issued, make another request to verify the factor.  &gt; **Note:** &gt; Unlike standard push challenges that don&#39;t require a request body, a number matching [&#x60;push&#x60;](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor!path&#x3D;2/useNumberMatchingChallenge&amp;t&#x3D;request) challenge requires a request body. &#x60;useNumberMatchingChallenge&#x60; must be set to &#x60;true&#x60;. &gt; When a number matching challenge is issued for an Okta Verify &#x60;push&#x60; factor enrollment, a &#x60;correctAnswer&#x60; challenge object is returned in the [&#x60;_embedded&#x60;](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor!c&#x3D;200&amp;path&#x3D;_embedded&amp;t&#x3D;response) object. | 

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


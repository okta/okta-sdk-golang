# \WebAuthnPreregistrationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePreregistrationEnrollment**](WebAuthnPreregistrationAPI.md#ActivatePreregistrationEnrollment) | **Post** /webauthn-registration/api/v1/activate | Activate a Preregistered WebAuthn Factor
[**EnrollPreregistrationEnrollment**](WebAuthnPreregistrationAPI.md#EnrollPreregistrationEnrollment) | **Post** /webauthn-registration/api/v1/enroll | Enroll a Preregistered WebAuthn Factor
[**GenerateFulfillmentRequest**](WebAuthnPreregistrationAPI.md#GenerateFulfillmentRequest) | **Post** /webauthn-registration/api/v1/initiate-fulfillment-request | Generate a Fulfillment Request



## ActivatePreregistrationEnrollment

> EnrollmentActivationResponse ActivatePreregistrationEnrollment(ctx).Body(body).Execute()

Activate a Preregistered WebAuthn Factor



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
    body := *openapiclient.NewEnrollmentActivationRequest() // EnrollmentActivationRequest | Enrollment Activation Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebAuthnPreregistrationAPI.ActivatePreregistrationEnrollment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebAuthnPreregistrationAPI.ActivatePreregistrationEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivatePreregistrationEnrollment`: EnrollmentActivationResponse
    fmt.Fprintf(os.Stdout, "Response from `WebAuthnPreregistrationAPI.ActivatePreregistrationEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivatePreregistrationEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EnrollmentActivationRequest**](EnrollmentActivationRequest.md) | Enrollment Activation Request | 

### Return type

[**EnrollmentActivationResponse**](EnrollmentActivationResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnrollPreregistrationEnrollment

> EnrollmentInitializationResponse EnrollPreregistrationEnrollment(ctx).Body(body).Execute()

Enroll a Preregistered WebAuthn Factor



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
    body := *openapiclient.NewEnrollmentInitializationRequest() // EnrollmentInitializationRequest | Enrollment Initialization Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebAuthnPreregistrationAPI.EnrollPreregistrationEnrollment(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebAuthnPreregistrationAPI.EnrollPreregistrationEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnrollPreregistrationEnrollment`: EnrollmentInitializationResponse
    fmt.Fprintf(os.Stdout, "Response from `WebAuthnPreregistrationAPI.EnrollPreregistrationEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnrollPreregistrationEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EnrollmentInitializationRequest**](EnrollmentInitializationRequest.md) | Enrollment Initialization Request | 

### Return type

[**EnrollmentInitializationResponse**](EnrollmentInitializationResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateFulfillmentRequest

> GenerateFulfillmentRequest(ctx).Body(body).Execute()

Generate a Fulfillment Request



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
    body := *openapiclient.NewFulfillmentRequest() // FulfillmentRequest | Fulfillment Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebAuthnPreregistrationAPI.GenerateFulfillmentRequest(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebAuthnPreregistrationAPI.GenerateFulfillmentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateFulfillmentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FulfillmentRequest**](FulfillmentRequest.md) | Fulfillment Request | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


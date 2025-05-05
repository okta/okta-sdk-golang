# \ApplicationCredentialsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneApplicationKey**](ApplicationCredentialsAPI.md#CloneApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/{keyId}/clone | Clone a Key Credential
[**GenerateApplicationKey**](ApplicationCredentialsAPI.md#GenerateApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/generate | Generate a Key Credential
[**GenerateCsrForApplication**](ApplicationCredentialsAPI.md#GenerateCsrForApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs | Generate a Certificate Signing Request
[**GetApplicationKey**](ApplicationCredentialsAPI.md#GetApplicationKey) | **Get** /api/v1/apps/{appId}/credentials/keys/{keyId} | Retrieve a Key Credential
[**GetCsrForApplication**](ApplicationCredentialsAPI.md#GetCsrForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Retrieve a Certificate Signing Request
[**ListApplicationKeys**](ApplicationCredentialsAPI.md#ListApplicationKeys) | **Get** /api/v1/apps/{appId}/credentials/keys | List all Key Credentials
[**ListCsrsForApplication**](ApplicationCredentialsAPI.md#ListCsrsForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs | List all Certificate Signing Requests
[**PublishCsrFromApplication**](ApplicationCredentialsAPI.md#PublishCsrFromApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs/{csrId}/lifecycle/publish | Publish a Certificate Signing Request
[**RevokeCsrFromApplication**](ApplicationCredentialsAPI.md#RevokeCsrFromApplication) | **Delete** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Revoke a Certificate Signing Request



## CloneApplicationKey

> JsonWebKey CloneApplicationKey(ctx, appId, keyId).TargetAid(targetAid).Execute()

Clone a Key Credential



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    keyId := "sjP9eiETijYz110VkhHN" // string | ID of the Key Credential for the application
    targetAid := "targetAid_example" // string | Unique key of the target Application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.CloneApplicationKey(context.Background(), appId, keyId).TargetAid(targetAid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.CloneApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneApplicationKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.CloneApplicationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | ID of the Key Credential for the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetAid** | **string** | Unique key of the target Application | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateApplicationKey

> JsonWebKey GenerateApplicationKey(ctx, appId).ValidityYears(validityYears).Execute()

Generate a Key Credential



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    validityYears := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.GenerateApplicationKey(context.Background(), appId).ValidityYears(validityYears).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.GenerateApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateApplicationKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.GenerateApplicationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityYears** | **int32** |  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCsrForApplication

> Csr GenerateCsrForApplication(ctx, appId).Metadata(metadata).Execute()

Generate a Certificate Signing Request



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    metadata := *openapiclient.NewCsrMetadata() // CsrMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.GenerateCsrForApplication(context.Background(), appId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.GenerateCsrForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCsrForApplication`: Csr
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.GenerateCsrForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCsrForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadata** | [**CsrMetadata**](CsrMetadata.md) |  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationKey

> JsonWebKey GetApplicationKey(ctx, appId, keyId).Execute()

Retrieve a Key Credential



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    keyId := "sjP9eiETijYz110VkhHN" // string | ID of the Key Credential for the application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.GetApplicationKey(context.Background(), appId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.GetApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.GetApplicationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | ID of the Key Credential for the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsrForApplication

> Csr GetCsrForApplication(ctx, appId, csrId).Execute()

Retrieve a Certificate Signing Request



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    csrId := "fd7x1h7uTcZFx22rU1f7" // string | `id` of the CSR

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.GetCsrForApplication(context.Background(), appId, csrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.GetCsrForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrForApplication`: Csr
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.GetCsrForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**csrId** | **string** | &#x60;id&#x60; of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationKeys

> []JsonWebKey ListApplicationKeys(ctx, appId).Execute()

List all Key Credentials



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.ListApplicationKeys(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.ListApplicationKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.ListApplicationKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCsrsForApplication

> []Csr ListCsrsForApplication(ctx, appId).Execute()

List all Certificate Signing Requests



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.ListCsrsForApplication(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.ListCsrsForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCsrsForApplication`: []Csr
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.ListCsrsForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCsrsForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishCsrFromApplication

> JsonWebKey PublishCsrFromApplication(ctx, appId, csrId).Body(body).Execute()

Publish a Certificate Signing Request



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    csrId := "fd7x1h7uTcZFx22rU1f7" // string | `id` of the CSR
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCredentialsAPI.PublishCsrFromApplication(context.Background(), appId, csrId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.PublishCsrFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishCsrFromApplication`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialsAPI.PublishCsrFromApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**csrId** | **string** | &#x60;id&#x60; of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishCsrFromApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-x509-ca-cert, application/pkix-cert, application/x-pem-file
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCsrFromApplication

> RevokeCsrFromApplication(ctx, appId, csrId).Execute()

Revoke a Certificate Signing Request



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
    appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
    csrId := "fd7x1h7uTcZFx22rU1f7" // string | `id` of the CSR

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationCredentialsAPI.RevokeCsrFromApplication(context.Background(), appId, csrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialsAPI.RevokeCsrFromApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**csrId** | **string** | &#x60;id&#x60; of the CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCsrFromApplicationRequest struct via the builder pattern


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


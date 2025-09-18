# \ApplicationSSOCredentialKeyAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneApplicationKey**](ApplicationSSOCredentialKeyAPI.md#CloneApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/{keyId}/clone | Clone a key credential
[**GenerateApplicationKey**](ApplicationSSOCredentialKeyAPI.md#GenerateApplicationKey) | **Post** /api/v1/apps/{appId}/credentials/keys/generate | Generate a key credential
[**GenerateCsrForApplication**](ApplicationSSOCredentialKeyAPI.md#GenerateCsrForApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs | Generate a certificate signing request
[**GetApplicationKey**](ApplicationSSOCredentialKeyAPI.md#GetApplicationKey) | **Get** /api/v1/apps/{appId}/credentials/keys/{keyId} | Retrieve a key credential
[**GetCsrForApplication**](ApplicationSSOCredentialKeyAPI.md#GetCsrForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Retrieve a certificate signing request
[**ListApplicationKeys**](ApplicationSSOCredentialKeyAPI.md#ListApplicationKeys) | **Get** /api/v1/apps/{appId}/credentials/keys | List all key credentials
[**ListCsrsForApplication**](ApplicationSSOCredentialKeyAPI.md#ListCsrsForApplication) | **Get** /api/v1/apps/{appId}/credentials/csrs | List all certificate signing requests
[**PublishCsrFromApplication**](ApplicationSSOCredentialKeyAPI.md#PublishCsrFromApplication) | **Post** /api/v1/apps/{appId}/credentials/csrs/{csrId}/lifecycle/publish | Publish a certificate signing request
[**RevokeCsrFromApplication**](ApplicationSSOCredentialKeyAPI.md#RevokeCsrFromApplication) | **Delete** /api/v1/apps/{appId}/credentials/csrs/{csrId} | Revoke a certificate signing request



## CloneApplicationKey

> JsonWebKey CloneApplicationKey(ctx, appId, keyId).TargetAid(targetAid).Execute()

Clone a key credential



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
	targetAid := "0ouuytCAJSSDELFTUIDS" // string | Unique key of the target Application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.CloneApplicationKey(context.Background(), appId, keyId).TargetAid(targetAid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.CloneApplicationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneApplicationKey`: JsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.CloneApplicationKey`: %v\n", resp)
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

Generate a key credential



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
	validityYears := int32(5) // int32 | Expiry years of the Application Key Credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.GenerateApplicationKey(context.Background(), appId).ValidityYears(validityYears).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.GenerateApplicationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateApplicationKey`: JsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.GenerateApplicationKey`: %v\n", resp)
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

 **validityYears** | **int32** | Expiry years of the Application Key Credential | 

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

> string GenerateCsrForApplication(ctx, appId).Metadata(metadata).Execute()

Generate a certificate signing request



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.GenerateCsrForApplication(context.Background(), appId).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.GenerateCsrForApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCsrForApplication`: string
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.GenerateCsrForApplication`: %v\n", resp)
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

**string**

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pkcs10, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationKey

> JsonWebKey GetApplicationKey(ctx, appId, keyId).Execute()

Retrieve a key credential



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.GetApplicationKey(context.Background(), appId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.GetApplicationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationKey`: JsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.GetApplicationKey`: %v\n", resp)
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

Retrieve a certificate signing request



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.GetCsrForApplication(context.Background(), appId, csrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.GetCsrForApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsrForApplication`: Csr
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.GetCsrForApplication`: %v\n", resp)
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
- **Accept**: application/json, application/pkcs10

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationKeys

> []JsonWebKey ListApplicationKeys(ctx, appId).Execute()

List all key credentials



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.ListApplicationKeys(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.ListApplicationKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationKeys`: []JsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.ListApplicationKeys`: %v\n", resp)
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

List all certificate signing requests



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.ListCsrsForApplication(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.ListCsrsForApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCsrsForApplication`: []Csr
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.ListCsrsForApplication`: %v\n", resp)
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

Publish a certificate signing request



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
	resp, r, err := apiClient.ApplicationSSOCredentialKeyAPI.PublishCsrFromApplication(context.Background(), appId, csrId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.PublishCsrFromApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishCsrFromApplication`: JsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOCredentialKeyAPI.PublishCsrFromApplication`: %v\n", resp)
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

Revoke a certificate signing request



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
	r, err := apiClient.ApplicationSSOCredentialKeyAPI.RevokeCsrFromApplication(context.Background(), appId, csrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOCredentialKeyAPI.RevokeCsrFromApplication``: %v\n", err)
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


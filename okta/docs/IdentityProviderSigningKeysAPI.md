# \IdentityProviderSigningKeysAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneIdentityProviderKey**](IdentityProviderSigningKeysAPI.md#CloneIdentityProviderKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/{kid}/clone | Clone a signing key credential for IdP
[**GenerateCsrForIdentityProvider**](IdentityProviderSigningKeysAPI.md#GenerateCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs | Generate a certificate signing request
[**GenerateIdentityProviderSigningKey**](IdentityProviderSigningKeysAPI.md#GenerateIdentityProviderSigningKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/generate | Generate a new signing key credential for IdP
[**GetCsrForIdentityProvider**](IdentityProviderSigningKeysAPI.md#GetCsrForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId} | Retrieve a certificate signing request
[**GetIdentityProviderSigningKey**](IdentityProviderSigningKeysAPI.md#GetIdentityProviderSigningKey) | **Get** /api/v1/idps/{idpId}/credentials/keys/{kid} | Retrieve a signing key credential for IdP
[**ListActiveIdentityProviderSigningKey**](IdentityProviderSigningKeysAPI.md#ListActiveIdentityProviderSigningKey) | **Get** /api/v1/idps/{idpId}/credentials/keys/active | List the active signing key credential for IdP
[**ListCsrsForIdentityProvider**](IdentityProviderSigningKeysAPI.md#ListCsrsForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs | List all certificate signing requests
[**ListIdentityProviderSigningKeys**](IdentityProviderSigningKeysAPI.md#ListIdentityProviderSigningKeys) | **Get** /api/v1/idps/{idpId}/credentials/keys | List all signing key credentials for IdP
[**PublishCsrForIdentityProvider**](IdentityProviderSigningKeysAPI.md#PublishCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId}/lifecycle/publish | Publish a certificate signing request
[**RevokeCsrForIdentityProvider**](IdentityProviderSigningKeysAPI.md#RevokeCsrForIdentityProvider) | **Delete** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId} | Revoke a certificate signing request



## CloneIdentityProviderKey

> IdPKeyCredential CloneIdentityProviderKey(ctx, idpId, kid).TargetIdpId(targetIdpId).Execute()

Clone a signing key credential for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	kid := "KmMo85SSsU7TZzOShcGb" // string | Unique `id` of the IdP key credential
	targetIdpId := "targetIdpId_example" // string | `id` of the target IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.CloneIdentityProviderKey(context.Background(), idpId, kid).TargetIdpId(targetIdpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.CloneIdentityProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneIdentityProviderKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.CloneIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**kid** | **string** | Unique &#x60;id&#x60; of the IdP key credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetIdpId** | **string** | &#x60;id&#x60; of the target IdP | 

### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCsrForIdentityProvider

> IdPCsr GenerateCsrForIdentityProvider(ctx, idpId).Metadata(metadata).Execute()

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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	metadata := *openapiclient.NewCsrMetadata() // CsrMetadata | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.GenerateCsrForIdentityProvider(context.Background(), idpId).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.GenerateCsrForIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCsrForIdentityProvider`: IdPCsr
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.GenerateCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadata** | [**CsrMetadata**](CsrMetadata.md) |  | 

### Return type

[**IdPCsr**](IdPCsr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/pkcs10

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateIdentityProviderSigningKey

> IdPKeyCredential GenerateIdentityProviderSigningKey(ctx, idpId).ValidityYears(validityYears).Execute()

Generate a new signing key credential for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	validityYears := int32(56) // int32 | expiry of the IdP key credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.GenerateIdentityProviderSigningKey(context.Background(), idpId).ValidityYears(validityYears).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.GenerateIdentityProviderSigningKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateIdentityProviderSigningKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.GenerateIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateIdentityProviderSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityYears** | **int32** | expiry of the IdP key credential | 

### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsrForIdentityProvider

> IdPCsr GetCsrForIdentityProvider(ctx, idpId, idpCsrId).Execute()

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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.GetCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.GetCsrForIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsrForIdentityProvider`: IdPCsr
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.GetCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdPCsr**](IdPCsr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pkcs10

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviderSigningKey

> IdPKeyCredential GetIdentityProviderSigningKey(ctx, idpId, kid).Execute()

Retrieve a signing key credential for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	kid := "KmMo85SSsU7TZzOShcGb" // string | Unique `id` of the IdP key credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.GetIdentityProviderSigningKey(context.Background(), idpId, kid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.GetIdentityProviderSigningKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProviderSigningKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.GetIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**kid** | **string** | Unique &#x60;id&#x60; of the IdP key credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveIdentityProviderSigningKey

> []IdPKeyCredential ListActiveIdentityProviderSigningKey(ctx, idpId).Execute()

List the active signing key credential for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.ListActiveIdentityProviderSigningKey(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.ListActiveIdentityProviderSigningKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActiveIdentityProviderSigningKey`: []IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.ListActiveIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActiveIdentityProviderSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCsrsForIdentityProvider

> []IdPCsr ListCsrsForIdentityProvider(ctx, idpId).Execute()

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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.ListCsrsForIdentityProvider(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.ListCsrsForIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCsrsForIdentityProvider`: []IdPCsr
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.ListCsrsForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCsrsForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdPCsr**](IdPCsr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviderSigningKeys

> []IdPKeyCredential ListIdentityProviderSigningKeys(ctx, idpId).Execute()

List all signing key credentials for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.ListIdentityProviderSigningKeys(context.Background(), idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.ListIdentityProviderSigningKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityProviderSigningKeys`: []IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.ListIdentityProviderSigningKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderSigningKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishCsrForIdentityProvider

> IdPKeyCredential PublishCsrForIdentityProvider(ctx, idpId, idpCsrId).Body(body).Execute()

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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderSigningKeysAPI.PublishCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.PublishCsrForIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishCsrForIdentityProvider`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderSigningKeysAPI.PublishCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/pkix-cert, application/x-x509-ca-cert, application/x-pem-file
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCsrForIdentityProvider

> RevokeCsrForIdentityProvider(ctx, idpId, idpCsrId).Execute()

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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderSigningKeysAPI.RevokeCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderSigningKeysAPI.RevokeCsrForIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCsrForIdentityProviderRequest struct via the builder pattern


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


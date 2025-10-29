# \IdentityProviderKeysAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProviderKey**](IdentityProviderKeysAPI.md#CreateIdentityProviderKey) | **Post** /api/v1/idps/credentials/keys | Create an IdP key credential
[**DeleteIdentityProviderKey**](IdentityProviderKeysAPI.md#DeleteIdentityProviderKey) | **Delete** /api/v1/idps/credentials/keys/{kid} | Delete an IdP key credential
[**GetIdentityProviderKey**](IdentityProviderKeysAPI.md#GetIdentityProviderKey) | **Get** /api/v1/idps/credentials/keys/{kid} | Retrieve an IdP key credential
[**ListIdentityProviderKeys**](IdentityProviderKeysAPI.md#ListIdentityProviderKeys) | **Get** /api/v1/idps/credentials/keys | List all IdP key credentials
[**ReplaceIdentityProviderKey**](IdentityProviderKeysAPI.md#ReplaceIdentityProviderKey) | **Put** /api/v1/idps/credentials/keys/{kid} | Replace an IdP key credential



## CreateIdentityProviderKey

> IdPKeyCredential CreateIdentityProviderKey(ctx).JsonWebKey(jsonWebKey).Execute()

Create an IdP key credential



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
	jsonWebKey := *openapiclient.NewIdPCertificateCredential([]string{"MIIDnjCCAoagAwIBAgIGAVG3MN+PMA0GCSqGSIb3DQEBBQUAMIGPMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzENMAsGA1UECgwET2t0YTEUMBIGA1UECwwLU1NPUHJvdmlkZXIxEDAOBgNVBAMMB2V4YW1wbGUxHDAaBgkqhkiG9w0BCQEWDWluZm9Ab2t0YS5jb20wHhcNMTUxMjE4MjIyMjMyWhcNMjUxMjE4MjIyMzMyWjCBjzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDTALBgNVBAoMBE9rdGExFDASBgNVBAsMC1NTT1Byb3ZpZGVyMRAwDgYDVQQDDAdleGFtcGxlMRwwGgYJKoZIhvcNAQkBFg1pbmZvQG9rdGEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtcnyvuVCrsFEKCwHDenS3Ocjed8eWDv3zLtD2K/iZfE8BMj2wpTfn6Ry8zCYey3mWlKdxIybnV9amrujGRnE0ab6Q16v9D6RlFQLOG6dwqoRKuZy33Uyg8PGdEudZjGbWuKCqqXEp+UKALJHV+k4wWeVH8g5d1n3KyR2TVajVJpCrPhLFmq1Il4G/IUnPe4MvjXqB6CpKkog1+ThWsItPRJPAM+RweFHXq7KfChXsYE7Mmfuly8sDQlvBmQyxZnFHVuiPfCvGHJjpvHy11YlHdOjfgqHRvZbmo30+y0X/oY/yV4YEJ00LL6eJWU4wi7ViY3HP6/VCdRjHoRdr5L/DwIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQCzzhOFkvyYLNFj2WDcq1YqD4sBy1iCia9QpRH3rjQvMKDwQDYWbi6EdOX0TQ/IYR7UWGj+2pXd6v0t33lYtoKocp/4lUvT3tfBnWZ5KnObi+J2uY2teUqoYkASN7F+GRPVOuMVoVgm05ss8tuMb2dLc9vsx93sDt+XlMTv/2qi5VPwaDtqduKkzwW9lUfn4xIMkTiVvCpe0X2HneD2Bpuao3/U8Rk0uiPfq6TooWaoW3kjsmErhEAs9bA7xuqo1KKY9CdHcFhkSsMhoeaZylZHtzbnoipUlQKSLMdJQiiYZQ0bYL83/Ta9fulr1EERICMFt3GUmtYaZZKHpWSfdJp9"}) // IdPCertificateCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderKeysAPI.CreateIdentityProviderKey(context.Background()).JsonWebKey(jsonWebKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderKeysAPI.CreateIdentityProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentityProviderKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderKeysAPI.CreateIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonWebKey** | [**IdPCertificateCredential**](IdPCertificateCredential.md) |  | 

### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProviderKey

> DeleteIdentityProviderKey(ctx, kid).Execute()

Delete an IdP key credential



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
	kid := "KmMo85SSsU7TZzOShcGb" // string | Unique `id` of the IdP key credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderKeysAPI.DeleteIdentityProviderKey(context.Background(), kid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderKeysAPI.DeleteIdentityProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kid** | **string** | Unique &#x60;id&#x60; of the IdP key credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderKeyRequest struct via the builder pattern


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


## GetIdentityProviderKey

> IdPKeyCredential GetIdentityProviderKey(ctx, kid).Execute()

Retrieve an IdP key credential



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
	kid := "KmMo85SSsU7TZzOShcGb" // string | Unique `id` of the IdP key credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderKeysAPI.GetIdentityProviderKey(context.Background(), kid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderKeysAPI.GetIdentityProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProviderKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderKeysAPI.GetIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kid** | **string** | Unique &#x60;id&#x60; of the IdP key credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderKeyRequest struct via the builder pattern


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


## ListIdentityProviderKeys

> []IdPKeyCredential ListIdentityProviderKeys(ctx).After(after).Limit(limit).Execute()

List all IdP key credentials



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
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderKeysAPI.ListIdentityProviderKeys(context.Background()).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderKeysAPI.ListIdentityProviderKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityProviderKeys`: []IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderKeysAPI.ListIdentityProviderKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]

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


## ReplaceIdentityProviderKey

> IdPKeyCredential ReplaceIdentityProviderKey(ctx, kid).IdPKeyCredential(idPKeyCredential).Execute()

Replace an IdP key credential



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
	kid := "KmMo85SSsU7TZzOShcGb" // string | Unique `id` of the IdP key credential
	idPKeyCredential := *openapiclient.NewIdPKeyCredential() // IdPKeyCredential | Updated IdP key credential

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderKeysAPI.ReplaceIdentityProviderKey(context.Background(), kid).IdPKeyCredential(idPKeyCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderKeysAPI.ReplaceIdentityProviderKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceIdentityProviderKey`: IdPKeyCredential
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderKeysAPI.ReplaceIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kid** | **string** | Unique &#x60;id&#x60; of the IdP key credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idPKeyCredential** | [**IdPKeyCredential**](IdPKeyCredential.md) | Updated IdP key credential | 

### Return type

[**IdPKeyCredential**](IdPKeyCredential.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


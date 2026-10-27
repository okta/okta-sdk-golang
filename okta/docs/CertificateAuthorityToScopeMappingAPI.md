# \CertificateAuthorityToScopeMappingAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCertificateAuthoritiesForScope**](CertificateAuthorityToScopeMappingAPI.md#ListCertificateAuthoritiesForScope) | **Get** /device-identity/api/v1/certificate-authority-mappings/certificate-authorities | List all Certificate Authorities for a scope
[**ReplaceCertificateAuthorityMappings**](CertificateAuthorityToScopeMappingAPI.md#ReplaceCertificateAuthorityMappings) | **Put** /device-identity/api/v1/certificate-authority-mappings | Replace all Certificate Authority mappings for a scope



## ListCertificateAuthoritiesForScope

> []CertificateAuthorityToScopeMappingResponse ListCertificateAuthoritiesForScope(ctx).Scope(scope).Execute()

List all Certificate Authorities for a scope



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
	scope := "scope_example" // string | The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAuthorityToScopeMappingAPI.ListCertificateAuthoritiesForScope(context.Background()).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityToScopeMappingAPI.ListCertificateAuthoritiesForScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertificateAuthoritiesForScope`: []CertificateAuthorityToScopeMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityToScopeMappingAPI.ListCertificateAuthoritiesForScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificateAuthoritiesForScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected. | 

### Return type

[**[]CertificateAuthorityToScopeMappingResponse**](CertificateAuthorityToScopeMappingResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCertificateAuthorityMappings

> []CertificateAuthorityToScopeMapping ReplaceCertificateAuthorityMappings(ctx).Scope(scope).Mappings(mappings).Execute()

Replace all Certificate Authority mappings for a scope



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
	scope := "scope_example" // string | The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected.
	mappings := []openapiclient.CertificateAuthorityToScopeMappingRequest{*openapiclient.NewCertificateAuthorityToScopeMappingRequest("tci00000000000000001")} // []CertificateAuthorityToScopeMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateAuthorityToScopeMappingAPI.ReplaceCertificateAuthorityMappings(context.Background()).Scope(scope).Mappings(mappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityToScopeMappingAPI.ReplaceCertificateAuthorityMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCertificateAuthorityMappings`: []CertificateAuthorityToScopeMapping
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityToScopeMappingAPI.ReplaceCertificateAuthorityMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCertificateAuthorityMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The Certificate Authority scope. Only Certificate Authorities and mappings for this scope are affected. | 
 **mappings** | [**[]CertificateAuthorityToScopeMappingRequest**](CertificateAuthorityToScopeMappingRequest.md) |  | 

### Return type

[**[]CertificateAuthorityToScopeMapping**](CertificateAuthorityToScopeMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


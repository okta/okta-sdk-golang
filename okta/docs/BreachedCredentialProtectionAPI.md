# \BreachedCredentialProtectionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBreachedCredentialProtectionConfiguration**](BreachedCredentialProtectionAPI.md#GetBreachedCredentialProtectionConfiguration) | **Get** /api/v1/breached-credential-protection/configuration | Retrieve the breached credential protection configuration
[**ReplaceBreachedCredentialProtectionConfiguration**](BreachedCredentialProtectionAPI.md#ReplaceBreachedCredentialProtectionConfiguration) | **Put** /api/v1/breached-credential-protection/configuration | Replace the breached credential protection configuration



## GetBreachedCredentialProtectionConfiguration

> BreachedCredentialProtectionConfiguration GetBreachedCredentialProtectionConfiguration(ctx).Execute()

Retrieve the breached credential protection configuration



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BreachedCredentialProtectionAPI.GetBreachedCredentialProtectionConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BreachedCredentialProtectionAPI.GetBreachedCredentialProtectionConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBreachedCredentialProtectionConfiguration`: BreachedCredentialProtectionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BreachedCredentialProtectionAPI.GetBreachedCredentialProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBreachedCredentialProtectionConfigurationRequest struct via the builder pattern


### Return type

[**BreachedCredentialProtectionConfiguration**](BreachedCredentialProtectionConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBreachedCredentialProtectionConfiguration

> BreachedCredentialProtectionConfiguration ReplaceBreachedCredentialProtectionConfiguration(ctx).BreachedCredentialProtectionConfiguration(breachedCredentialProtectionConfiguration).Execute()

Replace the breached credential protection configuration



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
	breachedCredentialProtectionConfiguration := *openapiclient.NewBreachedCredentialProtectionConfigurationRequest("DetectionMethod_example") // BreachedCredentialProtectionConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BreachedCredentialProtectionAPI.ReplaceBreachedCredentialProtectionConfiguration(context.Background()).BreachedCredentialProtectionConfiguration(breachedCredentialProtectionConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BreachedCredentialProtectionAPI.ReplaceBreachedCredentialProtectionConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceBreachedCredentialProtectionConfiguration`: BreachedCredentialProtectionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BreachedCredentialProtectionAPI.ReplaceBreachedCredentialProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBreachedCredentialProtectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breachedCredentialProtectionConfiguration** | [**BreachedCredentialProtectionConfigurationRequest**](BreachedCredentialProtectionConfigurationRequest.md) |  | 

### Return type

[**BreachedCredentialProtectionConfiguration**](BreachedCredentialProtectionConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


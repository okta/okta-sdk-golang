# \DeviceAccessAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting**](DeviceAccessAPI.md#GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting) | **Get** /device-access/api/v1/desktop-mfa/enforce-number-matching-challenge-settings | Retrieve the Desktop MFA Enforce Number Matching Challenge org setting
[**GetDesktopMFARecoveryPinOrgSetting**](DeviceAccessAPI.md#GetDesktopMFARecoveryPinOrgSetting) | **Get** /device-access/api/v1/desktop-mfa/recovery-pin-settings | Retrieve the Desktop MFA Recovery PIN org setting
[**ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting**](DeviceAccessAPI.md#ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting) | **Put** /device-access/api/v1/desktop-mfa/enforce-number-matching-challenge-settings | Replace the Desktop MFA Enforce Number Matching Challenge org setting
[**ReplaceDesktopMFARecoveryPinOrgSetting**](DeviceAccessAPI.md#ReplaceDesktopMFARecoveryPinOrgSetting) | **Put** /device-access/api/v1/desktop-mfa/recovery-pin-settings | Replace the Desktop MFA Recovery PIN org setting



## GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting

> DesktopMFAEnforceNumberMatchingChallengeOrgSetting GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx).Execute()

Retrieve the Desktop MFA Enforce Number Matching Challenge org setting



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
	resp, r, err := apiClient.DeviceAccessAPI.GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAccessAPI.GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting`: DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	fmt.Fprintf(os.Stdout, "Response from `DeviceAccessAPI.GetDesktopMFAEnforceNumberMatchingChallengeOrgSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest struct via the builder pattern


### Return type

[**DesktopMFAEnforceNumberMatchingChallengeOrgSetting**](DesktopMFAEnforceNumberMatchingChallengeOrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopMFARecoveryPinOrgSetting

> DesktopMFARecoveryPinOrgSetting GetDesktopMFARecoveryPinOrgSetting(ctx).Execute()

Retrieve the Desktop MFA Recovery PIN org setting



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
	resp, r, err := apiClient.DeviceAccessAPI.GetDesktopMFARecoveryPinOrgSetting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAccessAPI.GetDesktopMFARecoveryPinOrgSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDesktopMFARecoveryPinOrgSetting`: DesktopMFARecoveryPinOrgSetting
	fmt.Fprintf(os.Stdout, "Response from `DeviceAccessAPI.GetDesktopMFARecoveryPinOrgSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopMFARecoveryPinOrgSettingRequest struct via the builder pattern


### Return type

[**DesktopMFARecoveryPinOrgSetting**](DesktopMFARecoveryPinOrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting

> DesktopMFAEnforceNumberMatchingChallengeOrgSetting ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting(ctx).DesktopMFAEnforceNumberMatchingChallengeOrgSetting(desktopMFAEnforceNumberMatchingChallengeOrgSetting).Execute()

Replace the Desktop MFA Enforce Number Matching Challenge org setting



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
	desktopMFAEnforceNumberMatchingChallengeOrgSetting := *openapiclient.NewDesktopMFAEnforceNumberMatchingChallengeOrgSetting() // DesktopMFAEnforceNumberMatchingChallengeOrgSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAccessAPI.ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting(context.Background()).DesktopMFAEnforceNumberMatchingChallengeOrgSetting(desktopMFAEnforceNumberMatchingChallengeOrgSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAccessAPI.ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting`: DesktopMFAEnforceNumberMatchingChallengeOrgSetting
	fmt.Fprintf(os.Stdout, "Response from `DeviceAccessAPI.ReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDesktopMFAEnforceNumberMatchingChallengeOrgSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **desktopMFAEnforceNumberMatchingChallengeOrgSetting** | [**DesktopMFAEnforceNumberMatchingChallengeOrgSetting**](DesktopMFAEnforceNumberMatchingChallengeOrgSetting.md) |  | 

### Return type

[**DesktopMFAEnforceNumberMatchingChallengeOrgSetting**](DesktopMFAEnforceNumberMatchingChallengeOrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDesktopMFARecoveryPinOrgSetting

> DesktopMFARecoveryPinOrgSetting ReplaceDesktopMFARecoveryPinOrgSetting(ctx).DesktopMFARecoveryPinOrgSetting(desktopMFARecoveryPinOrgSetting).Execute()

Replace the Desktop MFA Recovery PIN org setting



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
	desktopMFARecoveryPinOrgSetting := *openapiclient.NewDesktopMFARecoveryPinOrgSetting() // DesktopMFARecoveryPinOrgSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAccessAPI.ReplaceDesktopMFARecoveryPinOrgSetting(context.Background()).DesktopMFARecoveryPinOrgSetting(desktopMFARecoveryPinOrgSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAccessAPI.ReplaceDesktopMFARecoveryPinOrgSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDesktopMFARecoveryPinOrgSetting`: DesktopMFARecoveryPinOrgSetting
	fmt.Fprintf(os.Stdout, "Response from `DeviceAccessAPI.ReplaceDesktopMFARecoveryPinOrgSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDesktopMFARecoveryPinOrgSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **desktopMFARecoveryPinOrgSetting** | [**DesktopMFARecoveryPinOrgSetting**](DesktopMFARecoveryPinOrgSetting.md) |  | 

### Return type

[**DesktopMFARecoveryPinOrgSetting**](DesktopMFARecoveryPinOrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OrgSettingGeneralAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSettings**](OrgSettingGeneralAPI.md#GetOrgSettings) | **Get** /api/v1/org | Retrieve the Org general settings
[**ReplaceOrgSettings**](OrgSettingGeneralAPI.md#ReplaceOrgSettings) | **Put** /api/v1/org | Replace the Org general settings
[**UpdateOrgSettings**](OrgSettingGeneralAPI.md#UpdateOrgSettings) | **Post** /api/v1/org | Update the Org general settings



## GetOrgSettings

> OrgSetting GetOrgSettings(ctx).Execute()

Retrieve the Org general settings



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
	resp, r, err := apiClient.OrgSettingGeneralAPI.GetOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingGeneralAPI.GetOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSettings`: OrgSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingGeneralAPI.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrgSettings

> OrgSetting ReplaceOrgSettings(ctx).OrgSetting(orgSetting).Execute()

Replace the Org general settings



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
	orgSetting := *openapiclient.NewOrgSetting() // OrgSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingGeneralAPI.ReplaceOrgSettings(context.Background()).OrgSetting(orgSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingGeneralAPI.ReplaceOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceOrgSettings`: OrgSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingGeneralAPI.ReplaceOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> OrgSetting UpdateOrgSettings(ctx).OrgSetting(orgSetting).Execute()

Update the Org general settings



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
	orgSetting := *openapiclient.NewOrgSetting() // OrgSetting |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingGeneralAPI.UpdateOrgSettings(context.Background()).OrgSetting(orgSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingGeneralAPI.UpdateOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSettings`: OrgSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingGeneralAPI.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \OktaPersonalSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPersonalAppsExportBlockList**](OktaPersonalSettingsAPI.md#ListPersonalAppsExportBlockList) | **Get** /okta-personal-settings/api/v1/export-blocklists | List all blocked email domains
[**ReplaceBlockedEmailDomains**](OktaPersonalSettingsAPI.md#ReplaceBlockedEmailDomains) | **Put** /okta-personal-settings/api/v1/export-blocklists | Replace the blocked email domains
[**ReplaceOktaPersonalAdminSettings**](OktaPersonalSettingsAPI.md#ReplaceOktaPersonalAdminSettings) | **Put** /okta-personal-settings/api/v1/edit-feature | Replace the Okta Personal admin settings



## ListPersonalAppsExportBlockList

> PersonalAppsBlockList ListPersonalAppsExportBlockList(ctx).Execute()

List all blocked email domains



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
	resp, r, err := apiClient.OktaPersonalSettingsAPI.ListPersonalAppsExportBlockList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaPersonalSettingsAPI.ListPersonalAppsExportBlockList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPersonalAppsExportBlockList`: PersonalAppsBlockList
	fmt.Fprintf(os.Stdout, "Response from `OktaPersonalSettingsAPI.ListPersonalAppsExportBlockList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalAppsExportBlockListRequest struct via the builder pattern


### Return type

[**PersonalAppsBlockList**](PersonalAppsBlockList.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBlockedEmailDomains

> ReplaceBlockedEmailDomains(ctx).PersonalAppsBlockList(personalAppsBlockList).Execute()

Replace the blocked email domains



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
	personalAppsBlockList := *openapiclient.NewPersonalAppsBlockList() // PersonalAppsBlockList | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OktaPersonalSettingsAPI.ReplaceBlockedEmailDomains(context.Background()).PersonalAppsBlockList(personalAppsBlockList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaPersonalSettingsAPI.ReplaceBlockedEmailDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBlockedEmailDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personalAppsBlockList** | [**PersonalAppsBlockList**](PersonalAppsBlockList.md) |  | 

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


## ReplaceOktaPersonalAdminSettings

> ReplaceOktaPersonalAdminSettings(ctx).OktaPersonalAdminFeatureSettings(oktaPersonalAdminFeatureSettings).Execute()

Replace the Okta Personal admin settings



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
	oktaPersonalAdminFeatureSettings := *openapiclient.NewOktaPersonalAdminFeatureSettings() // OktaPersonalAdminFeatureSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OktaPersonalSettingsAPI.ReplaceOktaPersonalAdminSettings(context.Background()).OktaPersonalAdminFeatureSettings(oktaPersonalAdminFeatureSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OktaPersonalSettingsAPI.ReplaceOktaPersonalAdminSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOktaPersonalAdminSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oktaPersonalAdminFeatureSettings** | [**OktaPersonalAdminFeatureSettings**](OktaPersonalAdminFeatureSettings.md) |  | 

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


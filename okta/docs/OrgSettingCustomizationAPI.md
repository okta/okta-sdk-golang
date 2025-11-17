# \OrgSettingCustomizationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgPreferences**](OrgSettingCustomizationAPI.md#GetOrgPreferences) | **Get** /api/v1/org/preferences | Retrieve the org preferences
[**SetOrgHideOktaUIFooter**](OrgSettingCustomizationAPI.md#SetOrgHideOktaUIFooter) | **Post** /api/v1/org/preferences/hideEndUserFooter | Set the hide dashboard footer preference
[**SetOrgShowOktaUIFooter**](OrgSettingCustomizationAPI.md#SetOrgShowOktaUIFooter) | **Post** /api/v1/org/preferences/showEndUserFooter | Set the show dashboard footer preference



## GetOrgPreferences

> OrgPreferences GetOrgPreferences(ctx).Execute()

Retrieve the org preferences



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
	resp, r, err := apiClient.OrgSettingCustomizationAPI.GetOrgPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCustomizationAPI.GetOrgPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgPreferences`: OrgPreferences
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCustomizationAPI.GetOrgPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPreferencesRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrgHideOktaUIFooter

> OrgPreferences SetOrgHideOktaUIFooter(ctx).Execute()

Set the hide dashboard footer preference



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
	resp, r, err := apiClient.OrgSettingCustomizationAPI.SetOrgHideOktaUIFooter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCustomizationAPI.SetOrgHideOktaUIFooter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrgHideOktaUIFooter`: OrgPreferences
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCustomizationAPI.SetOrgHideOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrgHideOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrgShowOktaUIFooter

> OrgPreferences SetOrgShowOktaUIFooter(ctx).Execute()

Set the show dashboard footer preference



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
	resp, r, err := apiClient.OrgSettingCustomizationAPI.SetOrgShowOktaUIFooter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCustomizationAPI.SetOrgShowOktaUIFooter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrgShowOktaUIFooter`: OrgPreferences
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCustomizationAPI.SetOrgShowOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrgShowOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


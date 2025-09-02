# \OrgSettingAdminAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignClientPrivilegesSetting**](OrgSettingAdminAPI.md#AssignClientPrivilegesSetting) | **Put** /api/v1/org/settings/clientPrivilegesSetting | Assign the default public client app role setting
[**GetAutoAssignAdminAppSetting**](OrgSettingAdminAPI.md#GetAutoAssignAdminAppSetting) | **Get** /api/v1/org/settings/autoAssignAdminAppSetting | Retrieve the Okta Admin Console assignment setting
[**GetClientPrivilegesSetting**](OrgSettingAdminAPI.md#GetClientPrivilegesSetting) | **Get** /api/v1/org/settings/clientPrivilegesSetting | Retrieve the default public client app role setting
[**GetThirdPartyAdminSetting**](OrgSettingAdminAPI.md#GetThirdPartyAdminSetting) | **Get** /api/v1/org/orgSettings/thirdPartyAdminSetting | Retrieve the org third-party admin setting
[**UpdateAutoAssignAdminAppSetting**](OrgSettingAdminAPI.md#UpdateAutoAssignAdminAppSetting) | **Post** /api/v1/org/settings/autoAssignAdminAppSetting | Update the Okta Admin Console assignment setting
[**UpdateThirdPartyAdminSetting**](OrgSettingAdminAPI.md#UpdateThirdPartyAdminSetting) | **Post** /api/v1/org/orgSettings/thirdPartyAdminSetting | Update the org third-party admin setting



## AssignClientPrivilegesSetting

> ClientPrivilegesSetting AssignClientPrivilegesSetting(ctx).ClientPrivilegesSetting(clientPrivilegesSetting).Execute()

Assign the default public client app role setting



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
	clientPrivilegesSetting := *openapiclient.NewClientPrivilegesSetting() // ClientPrivilegesSetting |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingAdminAPI.AssignClientPrivilegesSetting(context.Background()).ClientPrivilegesSetting(clientPrivilegesSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.AssignClientPrivilegesSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignClientPrivilegesSetting`: ClientPrivilegesSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.AssignClientPrivilegesSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignClientPrivilegesSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientPrivilegesSetting** | [**ClientPrivilegesSetting**](ClientPrivilegesSetting.md) |  | 

### Return type

[**ClientPrivilegesSetting**](ClientPrivilegesSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoAssignAdminAppSetting

> AutoAssignAdminAppSetting GetAutoAssignAdminAppSetting(ctx).Execute()

Retrieve the Okta Admin Console assignment setting



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
	resp, r, err := apiClient.OrgSettingAdminAPI.GetAutoAssignAdminAppSetting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.GetAutoAssignAdminAppSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoAssignAdminAppSetting`: AutoAssignAdminAppSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.GetAutoAssignAdminAppSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoAssignAdminAppSettingRequest struct via the builder pattern


### Return type

[**AutoAssignAdminAppSetting**](AutoAssignAdminAppSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientPrivilegesSetting

> ClientPrivilegesSetting GetClientPrivilegesSetting(ctx).Execute()

Retrieve the default public client app role setting



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
	resp, r, err := apiClient.OrgSettingAdminAPI.GetClientPrivilegesSetting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.GetClientPrivilegesSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientPrivilegesSetting`: ClientPrivilegesSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.GetClientPrivilegesSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientPrivilegesSettingRequest struct via the builder pattern


### Return type

[**ClientPrivilegesSetting**](ClientPrivilegesSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThirdPartyAdminSetting

> ThirdPartyAdminSetting GetThirdPartyAdminSetting(ctx).Execute()

Retrieve the org third-party admin setting



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
	resp, r, err := apiClient.OrgSettingAdminAPI.GetThirdPartyAdminSetting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.GetThirdPartyAdminSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThirdPartyAdminSetting`: ThirdPartyAdminSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.GetThirdPartyAdminSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyAdminSettingRequest struct via the builder pattern


### Return type

[**ThirdPartyAdminSetting**](ThirdPartyAdminSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoAssignAdminAppSetting

> AutoAssignAdminAppSetting UpdateAutoAssignAdminAppSetting(ctx).AutoAssignAdminAppSetting(autoAssignAdminAppSetting).Execute()

Update the Okta Admin Console assignment setting



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
	autoAssignAdminAppSetting := *openapiclient.NewAutoAssignAdminAppSetting() // AutoAssignAdminAppSetting |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingAdminAPI.UpdateAutoAssignAdminAppSetting(context.Background()).AutoAssignAdminAppSetting(autoAssignAdminAppSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.UpdateAutoAssignAdminAppSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoAssignAdminAppSetting`: AutoAssignAdminAppSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.UpdateAutoAssignAdminAppSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoAssignAdminAppSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoAssignAdminAppSetting** | [**AutoAssignAdminAppSetting**](AutoAssignAdminAppSetting.md) |  | 

### Return type

[**AutoAssignAdminAppSetting**](AutoAssignAdminAppSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThirdPartyAdminSetting

> ThirdPartyAdminSetting UpdateThirdPartyAdminSetting(ctx).ThirdPartyAdminSetting(thirdPartyAdminSetting).Execute()

Update the org third-party admin setting



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
	thirdPartyAdminSetting := *openapiclient.NewThirdPartyAdminSetting() // ThirdPartyAdminSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingAdminAPI.UpdateThirdPartyAdminSetting(context.Background()).ThirdPartyAdminSetting(thirdPartyAdminSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAdminAPI.UpdateThirdPartyAdminSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateThirdPartyAdminSetting`: ThirdPartyAdminSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingAdminAPI.UpdateThirdPartyAdminSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThirdPartyAdminSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyAdminSetting** | [**ThirdPartyAdminSetting**](ThirdPartyAdminSetting.md) |  | 

### Return type

[**ThirdPartyAdminSetting**](ThirdPartyAdminSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


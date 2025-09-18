# \OrgSettingCommunicationAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOktaCommunicationSettings**](OrgSettingCommunicationAPI.md#GetOktaCommunicationSettings) | **Get** /api/v1/org/privacy/oktaCommunication | Retrieve the Okta communication settings
[**OptInUsersToOktaCommunicationEmails**](OrgSettingCommunicationAPI.md#OptInUsersToOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optIn | Opt in to Okta user communication emails
[**OptOutUsersFromOktaCommunicationEmails**](OrgSettingCommunicationAPI.md#OptOutUsersFromOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optOut | Opt out of Okta user communication emails



## GetOktaCommunicationSettings

> OrgOktaCommunicationSetting GetOktaCommunicationSettings(ctx).Execute()

Retrieve the Okta communication settings



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
	resp, r, err := apiClient.OrgSettingCommunicationAPI.GetOktaCommunicationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCommunicationAPI.GetOktaCommunicationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOktaCommunicationSettings`: OrgOktaCommunicationSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCommunicationAPI.GetOktaCommunicationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOktaCommunicationSettingsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptInUsersToOktaCommunicationEmails

> OrgOktaCommunicationSetting OptInUsersToOktaCommunicationEmails(ctx).Execute()

Opt in to Okta user communication emails



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
	resp, r, err := apiClient.OrgSettingCommunicationAPI.OptInUsersToOktaCommunicationEmails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCommunicationAPI.OptInUsersToOktaCommunicationEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptInUsersToOktaCommunicationEmails`: OrgOktaCommunicationSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCommunicationAPI.OptInUsersToOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptInUsersToOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOutUsersFromOktaCommunicationEmails

> OrgOktaCommunicationSetting OptOutUsersFromOktaCommunicationEmails(ctx).Execute()

Opt out of Okta user communication emails



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
	resp, r, err := apiClient.OrgSettingCommunicationAPI.OptOutUsersFromOktaCommunicationEmails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingCommunicationAPI.OptOutUsersFromOktaCommunicationEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptOutUsersFromOktaCommunicationEmails`: OrgOktaCommunicationSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingCommunicationAPI.OptOutUsersFromOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptOutUsersFromOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


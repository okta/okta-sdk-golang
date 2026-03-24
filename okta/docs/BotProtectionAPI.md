# \BotProtectionAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBotProtectionConfiguration**](BotProtectionAPI.md#GetBotProtectionConfiguration) | **Get** /api/v1/bot-protection/configuration | Retrieve the bot protection configuration
[**UpdateBotProtectionConfiguration**](BotProtectionAPI.md#UpdateBotProtectionConfiguration) | **Post** /api/v1/bot-protection/configuration | Update the bot protection configuration



## GetBotProtectionConfiguration

> BotProtectionConfiguration GetBotProtectionConfiguration(ctx).Execute()

Retrieve the bot protection configuration



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
	resp, r, err := apiClient.BotProtectionAPI.GetBotProtectionConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotProtectionAPI.GetBotProtectionConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotProtectionConfiguration`: BotProtectionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BotProtectionAPI.GetBotProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotProtectionConfigurationRequest struct via the builder pattern


### Return type

[**BotProtectionConfiguration**](BotProtectionConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBotProtectionConfiguration

> BotProtectionConfiguration UpdateBotProtectionConfiguration(ctx).BotProtectionConfiguration(botProtectionConfiguration).Execute()

Update the bot protection configuration



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
	botProtectionConfiguration := *openapiclient.NewBotProtectionConfiguration("Level_example", "Mode_example") // BotProtectionConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotProtectionAPI.UpdateBotProtectionConfiguration(context.Background()).BotProtectionConfiguration(botProtectionConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotProtectionAPI.UpdateBotProtectionConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBotProtectionConfiguration`: BotProtectionConfiguration
	fmt.Fprintf(os.Stdout, "Response from `BotProtectionAPI.UpdateBotProtectionConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBotProtectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botProtectionConfiguration** | [**BotProtectionConfiguration**](BotProtectionConfiguration.md) |  | 

### Return type

[**BotProtectionConfiguration**](BotProtectionConfiguration.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \UserSessionsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevokeUserSessions**](UserSessionsAPI.md#RevokeUserSessions) | **Delete** /api/v1/users/{userId}/sessions | Revoke all user sessions



## RevokeUserSessions

> RevokeUserSessions(ctx, userId).OauthTokens(oauthTokens).ForgetDevices(forgetDevices).Execute()

Revoke all user sessions



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
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	oauthTokens := true // bool | Revokes issued OpenID Connect and OAuth refresh and access tokens (optional) (default to false)
	forgetDevices := true // bool | Clears the user's remembered factors for all devices. > **Note:** This parameter defaults to false in Classic Engine. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserSessionsAPI.RevokeUserSessions(context.Background(), userId).OauthTokens(oauthTokens).ForgetDevices(forgetDevices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionsAPI.RevokeUserSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oauthTokens** | **bool** | Revokes issued OpenID Connect and OAuth refresh and access tokens | [default to false]
 **forgetDevices** | **bool** | Clears the user&#39;s remembered factors for all devices. &gt; **Note:** This parameter defaults to false in Classic Engine. | [default to true]

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


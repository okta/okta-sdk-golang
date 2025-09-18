# \UserOAuthAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRefreshTokenForUserAndClient**](UserOAuthAPI.md#GetRefreshTokenForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | Retrieve a refresh token for a client
[**ListRefreshTokensForUserAndClient**](UserOAuthAPI.md#ListRefreshTokensForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/tokens | List all refresh tokens for a client
[**RevokeTokenForUserAndClient**](UserOAuthAPI.md#RevokeTokenForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens/{tokenId} | Revoke a token for a client
[**RevokeTokensForUserAndClient**](UserOAuthAPI.md#RevokeTokensForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/tokens | Revoke all refresh tokens for a client



## GetRefreshTokenForUserAndClient

> OAuth2RefreshToken GetRefreshTokenForUserAndClient(ctx, userId, clientId, tokenId).Expand(expand).Execute()

Retrieve a refresh token for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
	expand := "scope" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOAuthAPI.GetRefreshTokenForUserAndClient(context.Background(), userId, clientId, tokenId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOAuthAPI.GetRefreshTokenForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRefreshTokenForUserAndClient`: OAuth2RefreshToken
	fmt.Fprintf(os.Stdout, "Response from `UserOAuthAPI.GetRefreshTokenForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | Client app ID | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefreshTokenForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 

### Return type

[**OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefreshTokensForUserAndClient

> []OAuth2RefreshToken ListRefreshTokensForUserAndClient(ctx, userId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all refresh tokens for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	expand := "scope" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | Specifies the number of tokens to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOAuthAPI.ListRefreshTokensForUserAndClient(context.Background(), userId, clientId).Expand(expand).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOAuthAPI.ListRefreshTokensForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRefreshTokensForUserAndClient`: []OAuth2RefreshToken
	fmt.Fprintf(os.Stdout, "Response from `UserOAuthAPI.ListRefreshTokensForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | Client app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRefreshTokensForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of tokens to return | [default to 20]

### Return type

[**[]OAuth2RefreshToken**](OAuth2RefreshToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeTokenForUserAndClient

> RevokeTokenForUserAndClient(ctx, userId, clientId, tokenId).Execute()

Revoke a token for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID
	tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserOAuthAPI.RevokeTokenForUserAndClient(context.Background(), userId, clientId, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOAuthAPI.RevokeTokenForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | Client app ID | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## RevokeTokensForUserAndClient

> RevokeTokensForUserAndClient(ctx, userId, clientId).Execute()

Revoke all refresh tokens for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | Client app ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserOAuthAPI.RevokeTokensForUserAndClient(context.Background(), userId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOAuthAPI.RevokeTokensForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | Client app ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokensForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


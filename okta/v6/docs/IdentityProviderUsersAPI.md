# \IdentityProviderUsersAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentityProviderApplicationUser**](IdentityProviderUsersAPI.md#GetIdentityProviderApplicationUser) | **Get** /api/v1/idps/{idpId}/users/{userId} | Retrieve a user for IdP
[**LinkUserToIdentityProvider**](IdentityProviderUsersAPI.md#LinkUserToIdentityProvider) | **Post** /api/v1/idps/{idpId}/users/{userId} | Link a user to IdP
[**ListIdentityProviderApplicationUsers**](IdentityProviderUsersAPI.md#ListIdentityProviderApplicationUsers) | **Get** /api/v1/idps/{idpId}/users | List all users for IdP
[**ListSocialAuthTokens**](IdentityProviderUsersAPI.md#ListSocialAuthTokens) | **Get** /api/v1/idps/{idpId}/users/{userId}/credentials/tokens | List all tokens from OIDC IdP
[**ListUserIdentityProviders**](IdentityProviderUsersAPI.md#ListUserIdentityProviders) | **Get** /api/v1/users/{id}/idps | List all IdPs for user
[**UnlinkUserFromIdentityProvider**](IdentityProviderUsersAPI.md#UnlinkUserFromIdentityProvider) | **Delete** /api/v1/idps/{idpId}/users/{userId} | Unlink a user from IdP



## GetIdentityProviderApplicationUser

> IdentityProviderApplicationUser GetIdentityProviderApplicationUser(ctx, idpId, userId).Execute()

Retrieve a user for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderUsersAPI.GetIdentityProviderApplicationUser(context.Background(), idpId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.GetIdentityProviderApplicationUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityProviderApplicationUser`: IdentityProviderApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderUsersAPI.GetIdentityProviderApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderApplicationUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkUserToIdentityProvider

> IdentityProviderApplicationUser LinkUserToIdentityProvider(ctx, idpId, userId).UserIdentityProviderLinkRequest(userIdentityProviderLinkRequest).Execute()

Link a user to IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user
	userIdentityProviderLinkRequest := *openapiclient.NewUserIdentityProviderLinkRequest() // UserIdentityProviderLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderUsersAPI.LinkUserToIdentityProvider(context.Background(), idpId, userId).UserIdentityProviderLinkRequest(userIdentityProviderLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.LinkUserToIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkUserToIdentityProvider`: IdentityProviderApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderUsersAPI.LinkUserToIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkUserToIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userIdentityProviderLinkRequest** | [**UserIdentityProviderLinkRequest**](UserIdentityProviderLinkRequest.md) |  | 

### Return type

[**IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviderApplicationUsers

> []IdentityProviderApplicationUser ListIdentityProviderApplicationUsers(ctx, idpId).Q(q).After(after).Limit(limit).Expand(expand).Execute()

List all users for IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	q := "q_example" // string | Searches the records for matching value (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | A limit on the number of objects to return (optional) (default to 20)
	expand := "user" // string | Expand user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderUsersAPI.ListIdentityProviderApplicationUsers(context.Background(), idpId).Q(q).After(after).Limit(limit).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.ListIdentityProviderApplicationUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityProviderApplicationUsers`: []IdentityProviderApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderUsersAPI.ListIdentityProviderApplicationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderApplicationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Searches the records for matching value | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | A limit on the number of objects to return | [default to 20]
 **expand** | **string** | Expand user data | 

### Return type

[**[]IdentityProviderApplicationUser**](IdentityProviderApplicationUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSocialAuthTokens

> []SocialAuthToken ListSocialAuthTokens(ctx, idpId, userId).Execute()

List all tokens from OIDC IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderUsersAPI.ListSocialAuthTokens(context.Background(), idpId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.ListSocialAuthTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSocialAuthTokens`: []SocialAuthToken
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderUsersAPI.ListSocialAuthTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSocialAuthTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SocialAuthToken**](SocialAuthToken.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserIdentityProviders

> []IdentityProvider ListUserIdentityProviders(ctx, id).Execute()

List all IdPs for user



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
	id := "id_example" // string | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProviderUsersAPI.ListUserIdentityProviders(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.ListUserIdentityProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserIdentityProviders`: []IdentityProvider
	fmt.Fprintf(os.Stdout, "Response from `IdentityProviderUsersAPI.ListUserIdentityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An ID, login, or login shortname (as long as the shortname is unambiguous) of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkUserFromIdentityProvider

> UnlinkUserFromIdentityProvider(ctx, idpId, userId).Execute()

Unlink a user from IdP



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
	idpId := "0oa62bfdjnK55Z5x80h7" // string | `id` of IdP
	userId := "00ub0oNGTSWTBKOLGLNR" // string | ID of an existing Okta user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProviderUsersAPI.UnlinkUserFromIdentityProvider(context.Background(), idpId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderUsersAPI.UnlinkUserFromIdentityProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkUserFromIdentityProviderRequest struct via the builder pattern


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


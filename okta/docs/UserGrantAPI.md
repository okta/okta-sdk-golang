# \UserGrantAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserGrant**](UserGrantAPI.md#GetUserGrant) | **Get** /api/v1/users/{userId}/grants/{grantId} | Retrieve a user grant
[**ListGrantsForUserAndClient**](UserGrantAPI.md#ListGrantsForUserAndClient) | **Get** /api/v1/users/{userId}/clients/{clientId}/grants | List all grants for a client
[**ListUserGrants**](UserGrantAPI.md#ListUserGrants) | **Get** /api/v1/users/{userId}/grants | List all user grants
[**RevokeGrantsForUserAndClient**](UserGrantAPI.md#RevokeGrantsForUserAndClient) | **Delete** /api/v1/users/{userId}/clients/{clientId}/grants | Revoke all grants for a client
[**RevokeUserGrant**](UserGrantAPI.md#RevokeUserGrant) | **Delete** /api/v1/users/{userId}/grants/{grantId} | Revoke a user grant
[**RevokeUserGrants**](UserGrantAPI.md#RevokeUserGrants) | **Delete** /api/v1/users/{userId}/grants | Revoke all user grants



## GetUserGrant

> OAuth2ScopeConsentGrant GetUserGrant(ctx, userId, grantId).Expand(expand).Execute()

Retrieve a user grant



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
	grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID
	expand := "scope" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserGrantAPI.GetUserGrant(context.Background(), userId, grantId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.GetUserGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGrant`: OAuth2ScopeConsentGrant
	fmt.Fprintf(os.Stdout, "Response from `UserGrantAPI.GetUserGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 

### Return type

[**OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrantsForUserAndClient

> []OAuth2ScopeConsentGrant ListGrantsForUserAndClient(ctx, userId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all grants for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
	expand := "expand_example" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | Specifies the number of tokens to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserGrantAPI.ListGrantsForUserAndClient(context.Background(), userId, clientId).Expand(expand).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.ListGrantsForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrantsForUserAndClient`: []OAuth2ScopeConsentGrant
	fmt.Fprintf(os.Stdout, "Response from `UserGrantAPI.ListGrantsForUserAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsForUserAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of tokens to return | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGrants

> []OAuth2ScopeConsentGrant ListUserGrants(ctx, userId).ScopeId(scopeId).Expand(expand).After(after).Limit(limit).Execute()

List all user grants



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
	scopeId := "scopeId_example" // string | The scope ID to filter on (optional)
	expand := "scope" // string | Valid value: `scope`. If specified, scope details are included in the `_embedded` attribute. (optional)
	after := "after_example" // string | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the `Link` response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | Specifies the number of grants to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserGrantAPI.ListUserGrants(context.Background(), userId).ScopeId(scopeId).Expand(expand).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.ListUserGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserGrants`: []OAuth2ScopeConsentGrant
	fmt.Fprintf(os.Stdout, "Response from `UserGrantAPI.ListUserGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeId** | **string** | The scope ID to filter on | 
 **expand** | **string** | Valid value: &#x60;scope&#x60;. If specified, scope details are included in the &#x60;_embedded&#x60; attribute. | 
 **after** | **string** | The cursor to use for pagination. It is an opaque string that specifies your current location in the list and is obtained from the &#x60;Link&#x60; response header. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | Specifies the number of grants to return | [default to 20]

### Return type

[**[]OAuth2ScopeConsentGrant**](OAuth2ScopeConsentGrant.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeGrantsForUserAndClient

> RevokeGrantsForUserAndClient(ctx, userId, clientId).Execute()

Revoke all grants for a client



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
	clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserGrantAPI.RevokeGrantsForUserAndClient(context.Background(), userId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.RevokeGrantsForUserAndClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeGrantsForUserAndClientRequest struct via the builder pattern


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


## RevokeUserGrant

> RevokeUserGrant(ctx, userId, grantId).Execute()

Revoke a user grant



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
	grantId := "iJoqkwx50mrgX4T9LcaH" // string | Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserGrantAPI.RevokeUserGrant(context.Background(), userId, grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.RevokeUserGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of an existing Okta user | 
**grantId** | **string** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserGrantRequest struct via the builder pattern


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


## RevokeUserGrants

> RevokeUserGrants(ctx, userId).Execute()

Revoke all user grants



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserGrantAPI.RevokeUserGrants(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserGrantAPI.RevokeUserGrants``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokeUserGrantsRequest struct via the builder pattern


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


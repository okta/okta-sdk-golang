# \OAuth2ResourceServerCredentialsKeysAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerCredentialsKeysAPI.md#ActivateOAuth2ResourceServerJsonWebKey) | **Post** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}/lifecycle/activate | Activate a Custom Authorization Server Public JSON Web Key
[**AddOAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerCredentialsKeysAPI.md#AddOAuth2ResourceServerJsonWebKey) | **Post** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys | Add a JSON Web Key
[**DeactivateOAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerCredentialsKeysAPI.md#DeactivateOAuth2ResourceServerJsonWebKey) | **Post** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId}/lifecycle/deactivate | Deactivate a Custom Authorization Server Public JSON Web Key
[**DeleteOAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerCredentialsKeysAPI.md#DeleteOAuth2ResourceServerJsonWebKey) | **Delete** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId} | Delete a Custom Authorization Server Public JSON Web Key
[**GetOAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerCredentialsKeysAPI.md#GetOAuth2ResourceServerJsonWebKey) | **Get** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys/{keyId} | Retrieve a Custom Authorization Server Public JSON Web Key
[**ListOAuth2ResourceServerJsonWebKeys**](OAuth2ResourceServerCredentialsKeysAPI.md#ListOAuth2ResourceServerJsonWebKeys) | **Get** /api/v1/authorizationServers/{authServerId}/resourceservercredentials/keys | List all Custom Authorization Server Public JSON Web Keys



## ActivateOAuth2ResourceServerJsonWebKey

> OAuth2ResourceServerJsonWebKey ActivateOAuth2ResourceServerJsonWebKey(ctx, authServerId, keyId).Execute()

Activate a Custom Authorization Server Public JSON Web Key



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.ActivateOAuth2ResourceServerJsonWebKey(context.Background(), authServerId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.ActivateOAuth2ResourceServerJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateOAuth2ResourceServerJsonWebKey`: OAuth2ResourceServerJsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ResourceServerCredentialsKeysAPI.ActivateOAuth2ResourceServerJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOAuth2ResourceServerJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOAuth2ResourceServerJsonWebKey

> OAuth2ResourceServerJsonWebKey AddOAuth2ResourceServerJsonWebKey(ctx, authServerId).OAuth2ResourceServerJsonWebKeyRequestBody(oAuth2ResourceServerJsonWebKeyRequestBody).Execute()

Add a JSON Web Key



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
	oAuth2ResourceServerJsonWebKeyRequestBody := *openapiclient.NewOAuth2ResourceServerJsonWebKeyRequestBody() // OAuth2ResourceServerJsonWebKeyRequestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.AddOAuth2ResourceServerJsonWebKey(context.Background(), authServerId).OAuth2ResourceServerJsonWebKeyRequestBody(oAuth2ResourceServerJsonWebKeyRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.AddOAuth2ResourceServerJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOAuth2ResourceServerJsonWebKey`: OAuth2ResourceServerJsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ResourceServerCredentialsKeysAPI.AddOAuth2ResourceServerJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOAuth2ResourceServerJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2ResourceServerJsonWebKeyRequestBody** | [**OAuth2ResourceServerJsonWebKeyRequestBody**](OAuth2ResourceServerJsonWebKeyRequestBody.md) |  | 

### Return type

[**OAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOAuth2ResourceServerJsonWebKey

> OAuth2ResourceServerJsonWebKey DeactivateOAuth2ResourceServerJsonWebKey(ctx, authServerId, keyId).Execute()

Deactivate a Custom Authorization Server Public JSON Web Key



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.DeactivateOAuth2ResourceServerJsonWebKey(context.Background(), authServerId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.DeactivateOAuth2ResourceServerJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateOAuth2ResourceServerJsonWebKey`: OAuth2ResourceServerJsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ResourceServerCredentialsKeysAPI.DeactivateOAuth2ResourceServerJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOAuth2ResourceServerJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuth2ResourceServerJsonWebKey

> DeleteOAuth2ResourceServerJsonWebKey(ctx, authServerId, keyId).Execute()

Delete a Custom Authorization Server Public JSON Web Key



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.DeleteOAuth2ResourceServerJsonWebKey(context.Background(), authServerId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.DeleteOAuth2ResourceServerJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ResourceServerJsonWebKeyRequest struct via the builder pattern


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


## GetOAuth2ResourceServerJsonWebKey

> OAuth2ResourceServerJsonWebKey GetOAuth2ResourceServerJsonWebKey(ctx, authServerId, keyId).Execute()

Retrieve a Custom Authorization Server Public JSON Web Key



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.GetOAuth2ResourceServerJsonWebKey(context.Background(), authServerId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.GetOAuth2ResourceServerJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuth2ResourceServerJsonWebKey`: OAuth2ResourceServerJsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ResourceServerCredentialsKeysAPI.GetOAuth2ResourceServerJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ResourceServerJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2ResourceServerJsonWebKeys

> []OAuth2ResourceServerJsonWebKey ListOAuth2ResourceServerJsonWebKeys(ctx, authServerId).Execute()

List all Custom Authorization Server Public JSON Web Keys



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
	authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth2ResourceServerCredentialsKeysAPI.ListOAuth2ResourceServerJsonWebKeys(context.Background(), authServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ResourceServerCredentialsKeysAPI.ListOAuth2ResourceServerJsonWebKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOAuth2ResourceServerJsonWebKeys`: []OAuth2ResourceServerJsonWebKey
	fmt.Fprintf(os.Stdout, "Response from `OAuth2ResourceServerCredentialsKeysAPI.ListOAuth2ResourceServerJsonWebKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ResourceServerJsonWebKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2ResourceServerJsonWebKey**](OAuth2ResourceServerJsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


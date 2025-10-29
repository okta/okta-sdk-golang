# \ApplicationSSOPublicKeysAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateOAuth2ClientJsonWebKey**](ApplicationSSOPublicKeysAPI.md#ActivateOAuth2ClientJsonWebKey) | **Post** /api/v1/apps/{appId}/credentials/jwks/{keyId}/lifecycle/activate | Activate an OAuth 2.0 client JSON Web Key
[**ActivateOAuth2ClientSecret**](ApplicationSSOPublicKeysAPI.md#ActivateOAuth2ClientSecret) | **Post** /api/v1/apps/{appId}/credentials/secrets/{secretId}/lifecycle/activate | Activate an OAuth 2.0 client secret
[**AddJwk**](ApplicationSSOPublicKeysAPI.md#AddJwk) | **Post** /api/v1/apps/{appId}/credentials/jwks | Add a JSON Web Key
[**CreateOAuth2ClientSecret**](ApplicationSSOPublicKeysAPI.md#CreateOAuth2ClientSecret) | **Post** /api/v1/apps/{appId}/credentials/secrets | Create an OAuth 2.0 client secret
[**DeactivateOAuth2ClientJsonWebKey**](ApplicationSSOPublicKeysAPI.md#DeactivateOAuth2ClientJsonWebKey) | **Post** /api/v1/apps/{appId}/credentials/jwks/{keyId}/lifecycle/deactivate | Deactivate an OAuth 2.0 client JSON Web Key
[**DeactivateOAuth2ClientSecret**](ApplicationSSOPublicKeysAPI.md#DeactivateOAuth2ClientSecret) | **Post** /api/v1/apps/{appId}/credentials/secrets/{secretId}/lifecycle/deactivate | Deactivate an OAuth 2.0 client secret
[**DeleteOAuth2ClientSecret**](ApplicationSSOPublicKeysAPI.md#DeleteOAuth2ClientSecret) | **Delete** /api/v1/apps/{appId}/credentials/secrets/{secretId} | Delete an OAuth 2.0 client secret
[**Deletejwk**](ApplicationSSOPublicKeysAPI.md#Deletejwk) | **Delete** /api/v1/apps/{appId}/credentials/jwks/{keyId} | Delete an OAuth 2.0 client JSON Web Key
[**GetJwk**](ApplicationSSOPublicKeysAPI.md#GetJwk) | **Get** /api/v1/apps/{appId}/credentials/jwks/{keyId} | Retrieve an OAuth 2.0 client JSON Web Key
[**GetOAuth2ClientSecret**](ApplicationSSOPublicKeysAPI.md#GetOAuth2ClientSecret) | **Get** /api/v1/apps/{appId}/credentials/secrets/{secretId} | Retrieve an OAuth 2.0 client secret
[**ListJwk**](ApplicationSSOPublicKeysAPI.md#ListJwk) | **Get** /api/v1/apps/{appId}/credentials/jwks | List all the OAuth 2.0 client JSON Web Keys
[**ListOAuth2ClientSecrets**](ApplicationSSOPublicKeysAPI.md#ListOAuth2ClientSecrets) | **Get** /api/v1/apps/{appId}/credentials/secrets | List all OAuth 2.0 client secrets



## ActivateOAuth2ClientJsonWebKey

> ListJwk200ResponseInner ActivateOAuth2ClientJsonWebKey(ctx, appId, keyId).Execute()

Activate an OAuth 2.0 client JSON Web Key



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientJsonWebKey(context.Background(), appId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateOAuth2ClientJsonWebKey`: ListJwk200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOAuth2ClientJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListJwk200ResponseInner**](ListJwk200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivateOAuth2ClientSecret

> OAuth2ClientSecret ActivateOAuth2ClientSecret(ctx, appId, secretId).Execute()

Activate an OAuth 2.0 client secret



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	secretId := "ocs2f4zrZbs8nUa7p0g4" // string | Unique `id` of the OAuth 2.0 Client Secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientSecret(context.Background(), appId, secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateOAuth2ClientSecret`: OAuth2ClientSecret
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.ActivateOAuth2ClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**secretId** | **string** | Unique &#x60;id&#x60; of the OAuth 2.0 Client Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateOAuth2ClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ClientSecret**](OAuth2ClientSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddJwk

> ListJwk200ResponseInner AddJwk(ctx, appId).AddJwkRequest(addJwkRequest).Execute()

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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	addJwkRequest := openapiclient.addJwk_request{OAuth2ClientJsonEncryptionKeyRequest: openapiclient.NewOAuth2ClientJsonEncryptionKeyRequest()} // AddJwkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.AddJwk(context.Background(), appId).AddJwkRequest(addJwkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.AddJwk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddJwk`: ListJwk200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.AddJwk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddJwkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addJwkRequest** | [**AddJwkRequest**](AddJwkRequest.md) |  | 

### Return type

[**ListJwk200ResponseInner**](ListJwk200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuth2ClientSecret

> OAuth2ClientSecret CreateOAuth2ClientSecret(ctx, appId).OAuth2ClientSecretRequestBody(oAuth2ClientSecretRequestBody).Execute()

Create an OAuth 2.0 client secret



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	oAuth2ClientSecretRequestBody := *openapiclient.NewOAuth2ClientSecretRequestBody() // OAuth2ClientSecretRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.CreateOAuth2ClientSecret(context.Background(), appId).OAuth2ClientSecretRequestBody(oAuth2ClientSecretRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.CreateOAuth2ClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOAuth2ClientSecret`: OAuth2ClientSecret
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.CreateOAuth2ClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2ClientSecretRequestBody** | [**OAuth2ClientSecretRequestBody**](OAuth2ClientSecretRequestBody.md) |  | 

### Return type

[**OAuth2ClientSecret**](OAuth2ClientSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOAuth2ClientJsonWebKey

> OAuth2ClientJsonSigningKeyResponse DeactivateOAuth2ClientJsonWebKey(ctx, appId, keyId).Execute()

Deactivate an OAuth 2.0 client JSON Web Key



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientJsonWebKey(context.Background(), appId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientJsonWebKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateOAuth2ClientJsonWebKey`: OAuth2ClientJsonSigningKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientJsonWebKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOAuth2ClientJsonWebKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ClientJsonSigningKeyResponse**](OAuth2ClientJsonSigningKeyResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateOAuth2ClientSecret

> OAuth2ClientSecret DeactivateOAuth2ClientSecret(ctx, appId, secretId).Execute()

Deactivate an OAuth 2.0 client secret



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	secretId := "ocs2f4zrZbs8nUa7p0g4" // string | Unique `id` of the OAuth 2.0 Client Secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientSecret(context.Background(), appId, secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateOAuth2ClientSecret`: OAuth2ClientSecret
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.DeactivateOAuth2ClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**secretId** | **string** | Unique &#x60;id&#x60; of the OAuth 2.0 Client Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateOAuth2ClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ClientSecret**](OAuth2ClientSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuth2ClientSecret

> DeleteOAuth2ClientSecret(ctx, appId, secretId).Execute()

Delete an OAuth 2.0 client secret



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	secretId := "ocs2f4zrZbs8nUa7p0g4" // string | Unique `id` of the OAuth 2.0 Client Secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationSSOPublicKeysAPI.DeleteOAuth2ClientSecret(context.Background(), appId, secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.DeleteOAuth2ClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**secretId** | **string** | Unique &#x60;id&#x60; of the OAuth 2.0 Client Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ClientSecretRequest struct via the builder pattern


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


## Deletejwk

> Deletejwk(ctx, appId, keyId).Execute()

Delete an OAuth 2.0 client JSON Web Key



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationSSOPublicKeysAPI.Deletejwk(context.Background(), appId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.Deletejwk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletejwkRequest struct via the builder pattern


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


## GetJwk

> GetJwk200Response GetJwk(ctx, appId, keyId).Execute()

Retrieve an OAuth 2.0 client JSON Web Key



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	keyId := "apk2f4zrZbs8nUa7p0g4" // string | Unique `id` of the Custom Authorization Server JSON Web Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.GetJwk(context.Background(), appId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.GetJwk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJwk`: GetJwk200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.GetJwk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**keyId** | **string** | Unique &#x60;id&#x60; of the Custom Authorization Server JSON Web Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetJwk200Response**](GetJwk200Response.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuth2ClientSecret

> OAuth2ClientSecret GetOAuth2ClientSecret(ctx, appId, secretId).Execute()

Retrieve an OAuth 2.0 client secret



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	secretId := "ocs2f4zrZbs8nUa7p0g4" // string | Unique `id` of the OAuth 2.0 Client Secret

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.GetOAuth2ClientSecret(context.Background(), appId, secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.GetOAuth2ClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuth2ClientSecret`: OAuth2ClientSecret
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.GetOAuth2ClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 
**secretId** | **string** | Unique &#x60;id&#x60; of the OAuth 2.0 Client Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2ClientSecret**](OAuth2ClientSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJwk

> []ListJwk200ResponseInner ListJwk(ctx, appId).Execute()

List all the OAuth 2.0 client JSON Web Keys



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.ListJwk(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.ListJwk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJwk`: []ListJwk200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.ListJwk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJwkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListJwk200ResponseInner**](ListJwk200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2ClientSecrets

> []OAuth2ClientSecret ListOAuth2ClientSecrets(ctx, appId).Execute()

List all OAuth 2.0 client secrets



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationSSOPublicKeysAPI.ListOAuth2ClientSecrets(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSSOPublicKeysAPI.ListOAuth2ClientSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOAuth2ClientSecrets`: []OAuth2ClientSecret
	fmt.Fprintf(os.Stdout, "Response from `ApplicationSSOPublicKeysAPI.ListOAuth2ClientSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ClientSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2ClientSecret**](OAuth2ClientSecret.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


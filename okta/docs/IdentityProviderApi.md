# \IdentityProviderApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateIdentityProvider**](IdentityProviderApi.md#ActivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/activate | Activate an Identity Provider
[**CloneIdentityProviderKey**](IdentityProviderApi.md#CloneIdentityProviderKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/{keyId}/clone | Clone a Signing Credential Key
[**CreateIdentityProvider**](IdentityProviderApi.md#CreateIdentityProvider) | **Post** /api/v1/idps | Create an Identity Provider
[**CreateIdentityProviderKey**](IdentityProviderApi.md#CreateIdentityProviderKey) | **Post** /api/v1/idps/credentials/keys | Create an X.509 Certificate Public Key
[**DeactivateIdentityProvider**](IdentityProviderApi.md#DeactivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/deactivate | Deactivate an Identity Provider
[**DeleteIdentityProvider**](IdentityProviderApi.md#DeleteIdentityProvider) | **Delete** /api/v1/idps/{idpId} | Delete an Identity Provider
[**DeleteIdentityProviderKey**](IdentityProviderApi.md#DeleteIdentityProviderKey) | **Delete** /api/v1/idps/credentials/keys/{keyId} | Delete a Signing Credential Key
[**GenerateCsrForIdentityProvider**](IdentityProviderApi.md#GenerateCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs | Generate a Certificate Signing Request
[**GenerateIdentityProviderSigningKey**](IdentityProviderApi.md#GenerateIdentityProviderSigningKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/generate | Generate a new Signing Credential Key
[**GetCsrForIdentityProvider**](IdentityProviderApi.md#GetCsrForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs/{csrId} | Retrieve a Certificate Signing Request
[**GetIdentityProvider**](IdentityProviderApi.md#GetIdentityProvider) | **Get** /api/v1/idps/{idpId} | Retrieve an Identity Provider
[**GetIdentityProviderApplicationUser**](IdentityProviderApi.md#GetIdentityProviderApplicationUser) | **Get** /api/v1/idps/{idpId}/users/{userId} | Retrieve a User
[**GetIdentityProviderKey**](IdentityProviderApi.md#GetIdentityProviderKey) | **Get** /api/v1/idps/credentials/keys/{keyId} | Retrieve an Credential Key
[**GetIdentityProviderSigningKey**](IdentityProviderApi.md#GetIdentityProviderSigningKey) | **Get** /api/v1/idps/{idpId}/credentials/keys/{keyId} | Retrieve a Signing Credential Key
[**LinkUserToIdentityProvider**](IdentityProviderApi.md#LinkUserToIdentityProvider) | **Post** /api/v1/idps/{idpId}/users/{userId} | Link a User to a Social IdP
[**ListCsrsForIdentityProvider**](IdentityProviderApi.md#ListCsrsForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs | List all Certificate Signing Requests
[**ListIdentityProviderApplicationUsers**](IdentityProviderApi.md#ListIdentityProviderApplicationUsers) | **Get** /api/v1/idps/{idpId}/users | List all Users
[**ListIdentityProviderKeys**](IdentityProviderApi.md#ListIdentityProviderKeys) | **Get** /api/v1/idps/credentials/keys | List all Credential Keys
[**ListIdentityProviderSigningKeys**](IdentityProviderApi.md#ListIdentityProviderSigningKeys) | **Get** /api/v1/idps/{idpId}/credentials/keys | List all Signing Credential Keys
[**ListIdentityProviders**](IdentityProviderApi.md#ListIdentityProviders) | **Get** /api/v1/idps | List all Identity Providers
[**ListSocialAuthTokens**](IdentityProviderApi.md#ListSocialAuthTokens) | **Get** /api/v1/idps/{idpId}/users/{userId}/credentials/tokens | List all Tokens from a OIDC Identity Provider
[**PublishCsrForIdentityProvider**](IdentityProviderApi.md#PublishCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs/{csrId}/lifecycle/publish | Publish a Certificate Signing Request
[**ReplaceIdentityProvider**](IdentityProviderApi.md#ReplaceIdentityProvider) | **Put** /api/v1/idps/{idpId} | Replace an Identity Provider
[**RevokeCsrForIdentityProvider**](IdentityProviderApi.md#RevokeCsrForIdentityProvider) | **Delete** /api/v1/idps/{idpId}/credentials/csrs/{csrId} | Revoke a Certificate Signing Request
[**UnlinkUserFromIdentityProvider**](IdentityProviderApi.md#UnlinkUserFromIdentityProvider) | **Delete** /api/v1/idps/{idpId}/users/{userId} | Unlink a User from IdP



## ActivateIdentityProvider

> IdentityProvider ActivateIdentityProvider(ctx, idpId).Execute()

Activate an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ActivateIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ActivateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ActivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneIdentityProviderKey

> JsonWebKey CloneIdentityProviderKey(ctx, idpId, keyId).TargetIdpId(targetIdpId).Execute()

Clone a Signing Credential Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    keyId := "keyId_example" // string | 
    targetIdpId := "targetIdpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.CloneIdentityProviderKey(context.Background(), idpId, keyId).TargetIdpId(targetIdpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.CloneIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.CloneIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetIdpId** | **string** |  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentityProvider

> IdentityProvider CreateIdentityProvider(ctx).IdentityProvider(identityProvider).Execute()

Create an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.CreateIdentityProvider(context.Background()).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.CreateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.CreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) |  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentityProviderKey

> JsonWebKey CreateIdentityProviderKey(ctx).JsonWebKey(jsonWebKey).Execute()

Create an X.509 Certificate Public Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jsonWebKey := *openapiclient.NewJsonWebKey() // JsonWebKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.CreateIdentityProviderKey(context.Background()).JsonWebKey(jsonWebKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.CreateIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.CreateIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonWebKey** | [**JsonWebKey**](JsonWebKey.md) |  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateIdentityProvider

> IdentityProvider DeactivateIdentityProvider(ctx, idpId).Execute()

Deactivate an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.DeactivateIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DeactivateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.DeactivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, idpId).Execute()

Delete an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.DeleteIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DeleteIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


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


## DeleteIdentityProviderKey

> DeleteIdentityProviderKey(ctx, keyId).Execute()

Delete a Signing Credential Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.DeleteIdentityProviderKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.DeleteIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderKeyRequest struct via the builder pattern


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


## GenerateCsrForIdentityProvider

> Csr GenerateCsrForIdentityProvider(ctx, idpId).Metadata(metadata).Execute()

Generate a Certificate Signing Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    metadata := *openapiclient.NewCsrMetadata() // CsrMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GenerateCsrForIdentityProvider(context.Background(), idpId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GenerateCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCsrForIdentityProvider`: Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GenerateCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadata** | [**CsrMetadata**](CsrMetadata.md) |  | 

### Return type

[**Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateIdentityProviderSigningKey

> JsonWebKey GenerateIdentityProviderSigningKey(ctx, idpId).ValidityYears(validityYears).Execute()

Generate a new Signing Credential Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    validityYears := int32(56) // int32 | expiry of the IdP Key Credential

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GenerateIdentityProviderSigningKey(context.Background(), idpId).ValidityYears(validityYears).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GenerateIdentityProviderSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateIdentityProviderSigningKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GenerateIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateIdentityProviderSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityYears** | **int32** | expiry of the IdP Key Credential | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsrForIdentityProvider

> Csr GetCsrForIdentityProvider(ctx, idpId, csrId).Execute()

Retrieve a Certificate Signing Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    csrId := "csrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetCsrForIdentityProvider(context.Background(), idpId, csrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrForIdentityProvider`: Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**csrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProvider

> IdentityProvider GetIdentityProvider(ctx, idpId).Execute()

Retrieve an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviderApplicationUser

> IdentityProviderApplicationUser GetIdentityProviderApplicationUser(ctx, idpId, userId).Execute()

Retrieve a User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProviderApplicationUser(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProviderApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderApplicationUser`: IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProviderApplicationUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**userId** | **string** |  | 

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


## GetIdentityProviderKey

> JsonWebKey GetIdentityProviderKey(ctx, keyId).Execute()

Retrieve an Credential Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProviderKey(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviderSigningKey

> JsonWebKey GetIdentityProviderSigningKey(ctx, idpId, keyId).Execute()

Retrieve a Signing Credential Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.GetIdentityProviderSigningKey(context.Background(), idpId, keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.GetIdentityProviderSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderSigningKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.GetIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProviderSigningKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JsonWebKey**](JsonWebKey.md)

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

Link a User to a Social IdP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    userId := "userId_example" // string | 
    userIdentityProviderLinkRequest := *openapiclient.NewUserIdentityProviderLinkRequest() // UserIdentityProviderLinkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.LinkUserToIdentityProvider(context.Background(), idpId, userId).UserIdentityProviderLinkRequest(userIdentityProviderLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.LinkUserToIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkUserToIdentityProvider`: IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.LinkUserToIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**userId** | **string** |  | 

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


## ListCsrsForIdentityProvider

> []Csr ListCsrsForIdentityProvider(ctx, idpId).Execute()

List all Certificate Signing Requests



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListCsrsForIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListCsrsForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCsrsForIdentityProvider`: []Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListCsrsForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCsrsForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Csr**](Csr.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviderApplicationUsers

> []IdentityProviderApplicationUser ListIdentityProviderApplicationUsers(ctx, idpId).Execute()

List all Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListIdentityProviderApplicationUsers(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListIdentityProviderApplicationUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderApplicationUsers`: []IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListIdentityProviderApplicationUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderApplicationUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListIdentityProviderKeys

> []JsonWebKey ListIdentityProviderKeys(ctx).After(after).Limit(limit).Execute()

List all Credential Keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    after := "after_example" // string | Specifies the pagination cursor for the next page of keys (optional)
    limit := int32(56) // int32 | Specifies the number of key results in a page (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListIdentityProviderKeys(context.Background()).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListIdentityProviderKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListIdentityProviderKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Specifies the pagination cursor for the next page of keys | 
 **limit** | **int32** | Specifies the number of key results in a page | [default to 20]

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviderSigningKeys

> []JsonWebKey ListIdentityProviderSigningKeys(ctx, idpId).Execute()

List all Signing Credential Keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListIdentityProviderSigningKeys(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListIdentityProviderSigningKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderSigningKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListIdentityProviderSigningKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProviderSigningKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityProviders

> []IdentityProvider ListIdentityProviders(ctx).Q(q).After(after).Limit(limit).Type_(type_).Execute()

List all Identity Providers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | Searches the name property of IdPs for matching value (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of IdPs (optional)
    limit := int32(56) // int32 | Specifies the number of IdP results in a page (optional) (default to 20)
    type_ := "type__example" // string | Filters IdPs by type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListIdentityProviders(context.Background()).Q(q).After(after).Limit(limit).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviders`: []IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Searches the name property of IdPs for matching value | 
 **after** | **string** | Specifies the pagination cursor for the next page of IdPs | 
 **limit** | **int32** | Specifies the number of IdP results in a page | [default to 20]
 **type_** | **string** | Filters IdPs by type | 

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


## ListSocialAuthTokens

> []SocialAuthToken ListSocialAuthTokens(ctx, idpId, userId).Execute()

List all Tokens from a OIDC Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ListSocialAuthTokens(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ListSocialAuthTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSocialAuthTokens`: []SocialAuthToken
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ListSocialAuthTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**userId** | **string** |  | 

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


## PublishCsrForIdentityProvider

> JsonWebKey PublishCsrForIdentityProvider(ctx, idpId, csrId).Body(body).Execute()

Publish a Certificate Signing Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    csrId := "csrId_example" // string | 
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.PublishCsrForIdentityProvider(context.Background(), idpId, csrId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.PublishCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishCsrForIdentityProvider`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.PublishCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**csrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishCsrForIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

[**JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/x-x509-ca-cert, application/pkix-cert, application/x-pem-file
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIdentityProvider

> IdentityProvider ReplaceIdentityProvider(ctx, idpId).IdentityProvider(identityProvider).Execute()

Replace an Identity Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.ReplaceIdentityProvider(context.Background(), idpId).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.ReplaceIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderApi.ReplaceIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) |  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeCsrForIdentityProvider

> RevokeCsrForIdentityProvider(ctx, idpId, csrId).Execute()

Revoke a Certificate Signing Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    csrId := "csrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.RevokeCsrForIdentityProvider(context.Background(), idpId, csrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.RevokeCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**csrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeCsrForIdentityProviderRequest struct via the builder pattern


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


## UnlinkUserFromIdentityProvider

> UnlinkUserFromIdentityProvider(ctx, idpId, userId).Execute()

Unlink a User from IdP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    idpId := "idpId_example" // string | 
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderApi.UnlinkUserFromIdentityProvider(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderApi.UnlinkUserFromIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** |  | 
**userId** | **string** |  | 

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


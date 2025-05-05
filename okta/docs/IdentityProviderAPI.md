# \IdentityProviderAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateIdentityProvider**](IdentityProviderAPI.md#ActivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/activate | Activate an Identity Provider
[**CloneIdentityProviderKey**](IdentityProviderAPI.md#CloneIdentityProviderKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/{idpKeyId}/clone | Clone a Signing Credential Key
[**CreateIdentityProvider**](IdentityProviderAPI.md#CreateIdentityProvider) | **Post** /api/v1/idps | Create an Identity Provider
[**CreateIdentityProviderKey**](IdentityProviderAPI.md#CreateIdentityProviderKey) | **Post** /api/v1/idps/credentials/keys | Create an X.509 Certificate Public Key
[**DeactivateIdentityProvider**](IdentityProviderAPI.md#DeactivateIdentityProvider) | **Post** /api/v1/idps/{idpId}/lifecycle/deactivate | Deactivate an Identity Provider
[**DeleteIdentityProvider**](IdentityProviderAPI.md#DeleteIdentityProvider) | **Delete** /api/v1/idps/{idpId} | Delete an Identity Provider
[**DeleteIdentityProviderKey**](IdentityProviderAPI.md#DeleteIdentityProviderKey) | **Delete** /api/v1/idps/credentials/keys/{idpKeyId} | Delete a Signing Credential Key
[**GenerateCsrForIdentityProvider**](IdentityProviderAPI.md#GenerateCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs | Generate a Certificate Signing Request
[**GenerateIdentityProviderSigningKey**](IdentityProviderAPI.md#GenerateIdentityProviderSigningKey) | **Post** /api/v1/idps/{idpId}/credentials/keys/generate | Generate a new Signing Credential Key
[**GetCsrForIdentityProvider**](IdentityProviderAPI.md#GetCsrForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId} | Retrieve a Certificate Signing Request
[**GetIdentityProvider**](IdentityProviderAPI.md#GetIdentityProvider) | **Get** /api/v1/idps/{idpId} | Retrieve an Identity Provider
[**GetIdentityProviderApplicationUser**](IdentityProviderAPI.md#GetIdentityProviderApplicationUser) | **Get** /api/v1/idps/{idpId}/users/{userId} | Retrieve a User
[**GetIdentityProviderKey**](IdentityProviderAPI.md#GetIdentityProviderKey) | **Get** /api/v1/idps/credentials/keys/{idpKeyId} | Retrieve an Credential Key
[**GetIdentityProviderSigningKey**](IdentityProviderAPI.md#GetIdentityProviderSigningKey) | **Get** /api/v1/idps/{idpId}/credentials/keys/{idpKeyId} | Retrieve a Signing Credential Key
[**LinkUserToIdentityProvider**](IdentityProviderAPI.md#LinkUserToIdentityProvider) | **Post** /api/v1/idps/{idpId}/users/{userId} | Link a User to a Social IdP
[**ListCsrsForIdentityProvider**](IdentityProviderAPI.md#ListCsrsForIdentityProvider) | **Get** /api/v1/idps/{idpId}/credentials/csrs | List all Certificate Signing Requests
[**ListIdentityProviderApplicationUsers**](IdentityProviderAPI.md#ListIdentityProviderApplicationUsers) | **Get** /api/v1/idps/{idpId}/users | List all Users
[**ListIdentityProviderKeys**](IdentityProviderAPI.md#ListIdentityProviderKeys) | **Get** /api/v1/idps/credentials/keys | List all Credential Keys
[**ListIdentityProviderSigningKeys**](IdentityProviderAPI.md#ListIdentityProviderSigningKeys) | **Get** /api/v1/idps/{idpId}/credentials/keys | List all Signing Credential Keys
[**ListIdentityProviders**](IdentityProviderAPI.md#ListIdentityProviders) | **Get** /api/v1/idps | List all Identity Providers
[**ListSocialAuthTokens**](IdentityProviderAPI.md#ListSocialAuthTokens) | **Get** /api/v1/idps/{idpId}/users/{userId}/credentials/tokens | List all Tokens from a OIDC Identity Provider
[**PublishCsrForIdentityProvider**](IdentityProviderAPI.md#PublishCsrForIdentityProvider) | **Post** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId}/lifecycle/publish | Publish a Certificate Signing Request
[**ReplaceIdentityProvider**](IdentityProviderAPI.md#ReplaceIdentityProvider) | **Put** /api/v1/idps/{idpId} | Replace an Identity Provider
[**RevokeCsrForIdentityProvider**](IdentityProviderAPI.md#RevokeCsrForIdentityProvider) | **Delete** /api/v1/idps/{idpId}/credentials/csrs/{idpCsrId} | Revoke a Certificate Signing Request
[**UnlinkUserFromIdentityProvider**](IdentityProviderAPI.md#UnlinkUserFromIdentityProvider) | **Delete** /api/v1/idps/{idpId}/users/{userId} | Unlink a User from IdP



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ActivateIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ActivateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ActivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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

> JsonWebKey CloneIdentityProviderKey(ctx, idpId, idpKeyId).TargetIdpId(targetIdpId).Execute()

Clone a Signing Credential Key



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    idpKeyId := "KmMo85SSsU7TZzOShcGb" // string | `id` of IdP Key
    targetIdpId := "targetIdpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.CloneIdentityProviderKey(context.Background(), idpId, idpKeyId).TargetIdpId(targetIdpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CloneIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CloneIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpKeyId** | **string** | &#x60;id&#x60; of IdP Key | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.CreateIdentityProvider(context.Background()).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CreateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CreateIdentityProvider`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    jsonWebKey := *openapiclient.NewJsonWebKey() // JsonWebKey | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.CreateIdentityProviderKey(context.Background()).JsonWebKey(jsonWebKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.CreateIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.CreateIdentityProviderKey`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.DeactivateIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeactivateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.DeactivateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityProviderAPI.DeleteIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeleteIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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

> DeleteIdentityProviderKey(ctx, idpKeyId).Execute()

Delete a Signing Credential Key



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
    idpKeyId := "KmMo85SSsU7TZzOShcGb" // string | `id` of IdP Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityProviderAPI.DeleteIdentityProviderKey(context.Background(), idpKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.DeleteIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpKeyId** | **string** | &#x60;id&#x60; of IdP Key | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    metadata := *openapiclient.NewCsrMetadata() // CsrMetadata | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GenerateCsrForIdentityProvider(context.Background(), idpId).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GenerateCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateCsrForIdentityProvider`: Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GenerateCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    validityYears := int32(56) // int32 | expiry of the IdP Key Credential

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GenerateIdentityProviderSigningKey(context.Background(), idpId).ValidityYears(validityYears).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GenerateIdentityProviderSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateIdentityProviderSigningKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GenerateIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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

> Csr GetCsrForIdentityProvider(ctx, idpId, idpCsrId).Execute()

Retrieve a Certificate Signing Request



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GetCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrForIdentityProvider`: Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProviderApplicationUser(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProviderApplicationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderApplicationUser`: IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProviderApplicationUser`: %v\n", resp)
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


## GetIdentityProviderKey

> JsonWebKey GetIdentityProviderKey(ctx, idpKeyId).Execute()

Retrieve an Credential Key



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
    idpKeyId := "KmMo85SSsU7TZzOShcGb" // string | `id` of IdP Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProviderKey(context.Background(), idpKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProviderKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProviderKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpKeyId** | **string** | &#x60;id&#x60; of IdP Key | 

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

> JsonWebKey GetIdentityProviderSigningKey(ctx, idpId, idpKeyId).Execute()

Retrieve a Signing Credential Key



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    idpKeyId := "KmMo85SSsU7TZzOShcGb" // string | `id` of IdP Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.GetIdentityProviderSigningKey(context.Background(), idpId, idpKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.GetIdentityProviderSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviderSigningKey`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.GetIdentityProviderSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpKeyId** | **string** | &#x60;id&#x60; of IdP Key | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    userId := "userId_example" // string | ID of an existing Okta user
    userIdentityProviderLinkRequest := *openapiclient.NewUserIdentityProviderLinkRequest() // UserIdentityProviderLinkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.LinkUserToIdentityProvider(context.Background(), idpId, userId).UserIdentityProviderLinkRequest(userIdentityProviderLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.LinkUserToIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkUserToIdentityProvider`: IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.LinkUserToIdentityProvider`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListCsrsForIdentityProvider(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListCsrsForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCsrsForIdentityProvider`: []Csr
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListCsrsForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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

> []IdentityProviderApplicationUser ListIdentityProviderApplicationUsers(ctx, idpId).Q(q).After(after).Limit(limit).Expand(expand).Execute()

List all Users



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    q := "q_example" // string | Searches the name property of IdPs for matching value (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of IdPs (optional)
    limit := int32(56) // int32 | Specifies the number of IdP results in a page (optional) (default to 20)
    expand := "expand_example" // string | Expand user data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListIdentityProviderApplicationUsers(context.Background(), idpId).Q(q).After(after).Limit(limit).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListIdentityProviderApplicationUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderApplicationUsers`: []IdentityProviderApplicationUser
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListIdentityProviderApplicationUsers`: %v\n", resp)
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

 **q** | **string** | Searches the name property of IdPs for matching value | 
 **after** | **string** | Specifies the pagination cursor for the next page of IdPs | 
 **limit** | **int32** | Specifies the number of IdP results in a page | [default to 20]
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    after := "after_example" // string | Specifies the pagination cursor for the next page of keys (optional)
    limit := int32(56) // int32 | Specifies the number of key results in a page (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListIdentityProviderKeys(context.Background()).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListIdentityProviderKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListIdentityProviderKeys`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListIdentityProviderSigningKeys(context.Background(), idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListIdentityProviderSigningKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviderSigningKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListIdentityProviderSigningKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    q := "q_example" // string | Searches the name property of IdPs for matching value (optional)
    after := "after_example" // string | Specifies the pagination cursor for the next page of IdPs (optional)
    limit := int32(56) // int32 | Specifies the number of IdP results in a page (optional) (default to 20)
    type_ := "type__example" // string | Filters IdPs by type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListIdentityProviders(context.Background()).Q(q).After(after).Limit(limit).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityProviders`: []IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListIdentityProviders`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ListSocialAuthTokens(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ListSocialAuthTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSocialAuthTokens`: []SocialAuthToken
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ListSocialAuthTokens`: %v\n", resp)
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


## PublishCsrForIdentityProvider

> JsonWebKey PublishCsrForIdentityProvider(ctx, idpId, idpCsrId).Body(body).Execute()

Publish a Certificate Signing Request



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.PublishCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.PublishCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishCsrForIdentityProvider`: JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.PublishCsrForIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    identityProvider := *openapiclient.NewIdentityProvider() // IdentityProvider | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAPI.ReplaceIdentityProvider(context.Background(), idpId).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.ReplaceIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAPI.ReplaceIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 

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

> RevokeCsrForIdentityProvider(ctx, idpId, idpCsrId).Execute()

Revoke a Certificate Signing Request



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
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    idpCsrId := "1uEhyE65oV3H6KM9gYcN" // string | `id` of the IdP CSR

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityProviderAPI.RevokeCsrForIdentityProvider(context.Background(), idpId, idpCsrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.RevokeCsrForIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idpId** | **string** | &#x60;id&#x60; of IdP | 
**idpCsrId** | **string** | &#x60;id&#x60; of the IdP CSR | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    idpId := "SVHoAOh0l8cPQkVX1LRl" // string | `id` of IdP
    userId := "userId_example" // string | ID of an existing Okta user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityProviderAPI.UnlinkUserFromIdentityProvider(context.Background(), idpId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAPI.UnlinkUserFromIdentityProvider``: %v\n", err)
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


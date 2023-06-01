# \AuthorizationServerApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthorizationServer**](AuthorizationServerApi.md#ActivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/activate | Activate an Authorization Server
[**ActivateAuthorizationServerPolicy**](AuthorizationServerApi.md#ActivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/activate | Activate a Policy
[**ActivateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#ActivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a Policy Rule
[**CreateAuthorizationServer**](AuthorizationServerApi.md#CreateAuthorizationServer) | **Post** /api/v1/authorizationServers | Create an Authorization Server
[**CreateAuthorizationServerPolicy**](AuthorizationServerApi.md#CreateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies | Create a Policy
[**CreateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#CreateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | Create a Policy Rule
[**CreateOAuth2Claim**](AuthorizationServerApi.md#CreateOAuth2Claim) | **Post** /api/v1/authorizationServers/{authServerId}/claims | Create a Custom Token Claim
[**CreateOAuth2Scope**](AuthorizationServerApi.md#CreateOAuth2Scope) | **Post** /api/v1/authorizationServers/{authServerId}/scopes | Create a Custom Token Scope
[**DeactivateAuthorizationServer**](AuthorizationServerApi.md#DeactivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/deactivate | Deactivate an Authorization Server
[**DeactivateAuthorizationServerPolicy**](AuthorizationServerApi.md#DeactivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/deactivate | Deactivate a Policy
[**DeactivateAuthorizationServerPolicyRule**](AuthorizationServerApi.md#DeactivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a Policy Rule
[**DeleteAuthorizationServer**](AuthorizationServerApi.md#DeleteAuthorizationServer) | **Delete** /api/v1/authorizationServers/{authServerId} | Delete an Authorization Server
[**DeleteAuthorizationServerPolicy**](AuthorizationServerApi.md#DeleteAuthorizationServerPolicy) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Delete a Policy
[**DeleteAuthorizationServerPolicyRule**](AuthorizationServerApi.md#DeleteAuthorizationServerPolicyRule) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Delete a Policy Rule
[**DeleteOAuth2Claim**](AuthorizationServerApi.md#DeleteOAuth2Claim) | **Delete** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Delete a Custom Token Claim
[**DeleteOAuth2Scope**](AuthorizationServerApi.md#DeleteOAuth2Scope) | **Delete** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Delete a Custom Token Scope
[**GetAuthorizationServer**](AuthorizationServerApi.md#GetAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId} | Retrieve an Authorization Server
[**GetAuthorizationServerPolicy**](AuthorizationServerApi.md#GetAuthorizationServerPolicy) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Retrieve a Policy
[**GetAuthorizationServerPolicyRule**](AuthorizationServerApi.md#GetAuthorizationServerPolicyRule) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Retrieve a Policy Rule
[**GetOAuth2Claim**](AuthorizationServerApi.md#GetOAuth2Claim) | **Get** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Retrieve a Custom Token Claim
[**GetOAuth2Scope**](AuthorizationServerApi.md#GetOAuth2Scope) | **Get** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Retrieve a Custom Token Scope
[**GetRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerApi.md#GetRefreshTokenForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Retrieve a Refresh Token for a Client
[**ListAuthorizationServerKeys**](AuthorizationServerApi.md#ListAuthorizationServerKeys) | **Get** /api/v1/authorizationServers/{authServerId}/credentials/keys | List all Credential Keys
[**ListAuthorizationServerPolicies**](AuthorizationServerApi.md#ListAuthorizationServerPolicies) | **Get** /api/v1/authorizationServers/{authServerId}/policies | List all Policies
[**ListAuthorizationServerPolicyRules**](AuthorizationServerApi.md#ListAuthorizationServerPolicyRules) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | List all Policy Rules
[**ListAuthorizationServers**](AuthorizationServerApi.md#ListAuthorizationServers) | **Get** /api/v1/authorizationServers | List all Authorization Servers
[**ListOAuth2Claims**](AuthorizationServerApi.md#ListOAuth2Claims) | **Get** /api/v1/authorizationServers/{authServerId}/claims | List all Custom Token Claims
[**ListOAuth2ClientsForAuthorizationServer**](AuthorizationServerApi.md#ListOAuth2ClientsForAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId}/clients | List all Clients
[**ListOAuth2Scopes**](AuthorizationServerApi.md#ListOAuth2Scopes) | **Get** /api/v1/authorizationServers/{authServerId}/scopes | List all Custom Token Scopes
[**ListRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerApi.md#ListRefreshTokensForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | List all Refresh Tokens for a Client
[**ReplaceAuthorizationServer**](AuthorizationServerApi.md#ReplaceAuthorizationServer) | **Put** /api/v1/authorizationServers/{authServerId} | Replace an Authorization Server
[**ReplaceAuthorizationServerPolicy**](AuthorizationServerApi.md#ReplaceAuthorizationServerPolicy) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Replace a Policy
[**ReplaceAuthorizationServerPolicyRule**](AuthorizationServerApi.md#ReplaceAuthorizationServerPolicyRule) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Replace a Policy Rule
[**ReplaceOAuth2Claim**](AuthorizationServerApi.md#ReplaceOAuth2Claim) | **Put** /api/v1/authorizationServers/{authServerId}/claims/{claimId} | Replace a Custom Token Claim
[**ReplaceOAuth2Scope**](AuthorizationServerApi.md#ReplaceOAuth2Scope) | **Put** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Replace a Custom Token Scope
[**RevokeRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerApi.md#RevokeRefreshTokenForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Revoke a Refresh Token for a Client
[**RevokeRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerApi.md#RevokeRefreshTokensForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | Revoke all Refresh Tokens for a Client
[**RotateAuthorizationServerKeys**](AuthorizationServerApi.md#RotateAuthorizationServerKeys) | **Post** /api/v1/authorizationServers/{authServerId}/credentials/lifecycle/keyRotate | Rotate all Credential Keys



## ActivateAuthorizationServer

> ActivateAuthorizationServer(ctx, authServerId).Execute()

Activate an Authorization Server



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ActivateAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ActivateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthorizationServerRequest struct via the builder pattern


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


## ActivateAuthorizationServerPolicy

> ActivateAuthorizationServerPolicy(ctx, authServerId, policyId).Execute()

Activate a Policy



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ActivateAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ActivateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthorizationServerPolicyRequest struct via the builder pattern


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


## ActivateAuthorizationServerPolicyRule

> ActivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Activate a Policy Rule



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ActivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ActivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## CreateAuthorizationServer

> AuthorizationServer CreateAuthorizationServer(ctx).AuthorizationServer(authorizationServer).Execute()

Create an Authorization Server



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
    authorizationServer := *openapiclient.NewAuthorizationServer() // AuthorizationServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.CreateAuthorizationServer(context.Background()).AuthorizationServer(authorizationServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.CreateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.CreateAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationServer** | [**AuthorizationServer**](AuthorizationServer.md) |  | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationServerPolicy

> AuthorizationServerPolicy CreateAuthorizationServerPolicy(ctx, authServerId).Policy(policy).Execute()

Create a Policy



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
    authServerId := "authServerId_example" // string | 
    policy := *openapiclient.NewAuthorizationServerPolicy() // AuthorizationServerPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.CreateAuthorizationServerPolicy(context.Background(), authServerId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.CreateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.CreateAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**AuthorizationServerPolicy**](AuthorizationServerPolicy.md) |  | 

### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule CreateAuthorizationServerPolicyRule(ctx, policyId, authServerId).PolicyRule(policyRule).Execute()

Create a Policy Rule



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
    policyId := "policyId_example" // string | 
    authServerId := "authServerId_example" // string | 
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.CreateAuthorizationServerPolicyRule(context.Background(), policyId, authServerId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.CreateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.CreateAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRule** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md) |  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuth2Claim

> OAuth2Claim CreateOAuth2Claim(ctx, authServerId).OAuth2Claim(oAuth2Claim).Execute()

Create a Custom Token Claim



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
    authServerId := "authServerId_example" // string | 
    oAuth2Claim := *openapiclient.NewOAuth2Claim() // OAuth2Claim | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.CreateOAuth2Claim(context.Background(), authServerId).OAuth2Claim(oAuth2Claim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.CreateOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.CreateOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2Claim** | [**OAuth2Claim**](OAuth2Claim.md) |  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuth2Scope

> OAuth2Scope CreateOAuth2Scope(ctx, authServerId).OAuth2Scope(oAuth2Scope).Execute()

Create a Custom Token Scope



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
    authServerId := "authServerId_example" // string | 
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.CreateOAuth2Scope(context.Background(), authServerId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.CreateOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.CreateOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2Scope** | [**OAuth2Scope**](OAuth2Scope.md) |  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateAuthorizationServer

> DeactivateAuthorizationServer(ctx, authServerId).Execute()

Deactivate an Authorization Server



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeactivateAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeactivateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthorizationServerRequest struct via the builder pattern


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


## DeactivateAuthorizationServerPolicy

> DeactivateAuthorizationServerPolicy(ctx, authServerId, policyId).Execute()

Deactivate a Policy



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeactivateAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeactivateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthorizationServerPolicyRequest struct via the builder pattern


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


## DeactivateAuthorizationServerPolicyRule

> DeactivateAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Deactivate a Policy Rule



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeactivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeactivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## DeleteAuthorizationServer

> DeleteAuthorizationServer(ctx, authServerId).Execute()

Delete an Authorization Server



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeleteAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeleteAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerRequest struct via the builder pattern


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


## DeleteAuthorizationServerPolicy

> DeleteAuthorizationServerPolicy(ctx, authServerId, policyId).Execute()

Delete a Policy



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeleteAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeleteAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerPolicyRequest struct via the builder pattern


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


## DeleteAuthorizationServerPolicyRule

> DeleteAuthorizationServerPolicyRule(ctx, policyId, authServerId, ruleId).Execute()

Delete a Policy Rule



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
    policyId := "policyId_example" // string | 
    authServerId := "authServerId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeleteAuthorizationServerPolicyRule(context.Background(), policyId, authServerId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeleteAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**authServerId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationServerPolicyRuleRequest struct via the builder pattern


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


## DeleteOAuth2Claim

> DeleteOAuth2Claim(ctx, authServerId, claimId).Execute()

Delete a Custom Token Claim



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
    authServerId := "authServerId_example" // string | 
    claimId := "claimId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeleteOAuth2Claim(context.Background(), authServerId, claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeleteOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**claimId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ClaimRequest struct via the builder pattern


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


## DeleteOAuth2Scope

> DeleteOAuth2Scope(ctx, authServerId, scopeId).Execute()

Delete a Custom Token Scope



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
    authServerId := "authServerId_example" // string | 
    scopeId := "scopeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.DeleteOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.DeleteOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**scopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuth2ScopeRequest struct via the builder pattern


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


## GetAuthorizationServer

> AuthorizationServer GetAuthorizationServer(ctx, authServerId).Execute()

Retrieve an Authorization Server



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationServerPolicy

> AuthorizationServerPolicy GetAuthorizationServerPolicy(ctx, authServerId, policyId).Execute()

Retrieve a Policy



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule GetAuthorizationServerPolicyRule(ctx, policyId, authServerId, ruleId).Execute()

Retrieve a Policy Rule



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
    policyId := "policyId_example" // string | 
    authServerId := "authServerId_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetAuthorizationServerPolicyRule(context.Background(), policyId, authServerId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**authServerId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuth2Claim

> OAuth2Claim GetOAuth2Claim(ctx, authServerId, claimId).Execute()

Retrieve a Custom Token Claim



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
    authServerId := "authServerId_example" // string | 
    claimId := "claimId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetOAuth2Claim(context.Background(), authServerId, claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**claimId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuth2Scope

> OAuth2Scope GetOAuth2Scope(ctx, authServerId, scopeId).Execute()

Retrieve a Custom Token Scope



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
    authServerId := "authServerId_example" // string | 
    scopeId := "scopeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**scopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefreshTokenForAuthorizationServerAndClient

> OAuth2RefreshToken GetRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId).Expand(expand).Execute()

Retrieve a Refresh Token for a Client



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
    authServerId := "authServerId_example" // string | 
    clientId := "clientId_example" // string | 
    tokenId := "tokenId_example" // string | 
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.GetRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.GetRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefreshTokenForAuthorizationServerAndClient`: OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.GetRefreshTokenForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**clientId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefreshTokenForAuthorizationServerAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 

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


## ListAuthorizationServerKeys

> []JsonWebKey ListAuthorizationServerKeys(ctx, authServerId).Execute()

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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListAuthorizationServerKeys(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServerKeysRequest struct via the builder pattern


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


## ListAuthorizationServerPolicies

> []AuthorizationServerPolicy ListAuthorizationServerPolicies(ctx, authServerId).Execute()

List all Policies



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListAuthorizationServerPolicies(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListAuthorizationServerPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerPolicies`: []AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListAuthorizationServerPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServerPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationServerPolicyRules

> []AuthorizationServerPolicyRule ListAuthorizationServerPolicyRules(ctx, policyId, authServerId).Execute()

List all Policy Rules



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
    policyId := "policyId_example" // string | 
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListAuthorizationServerPolicyRules(context.Background(), policyId, authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListAuthorizationServerPolicyRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerPolicyRules`: []AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListAuthorizationServerPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServerPolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationServers

> []AuthorizationServer ListAuthorizationServers(ctx).Q(q).Limit(limit).After(after).Execute()

List all Authorization Servers



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
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 200)
    after := "after_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListAuthorizationServers(context.Background()).Q(q).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListAuthorizationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServers`: []AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListAuthorizationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** |  | [default to 200]
 **after** | **string** |  | 

### Return type

[**[]AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2Claims

> []OAuth2Claim ListOAuth2Claims(ctx, authServerId).Execute()

List all Custom Token Claims



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListOAuth2Claims(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListOAuth2Claims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2Claims`: []OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListOAuth2Claims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ClaimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2ClientsForAuthorizationServer

> []OAuth2Client ListOAuth2ClientsForAuthorizationServer(ctx, authServerId).Execute()

List all Clients



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
    authServerId := "authServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListOAuth2ClientsForAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListOAuth2ClientsForAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2ClientsForAuthorizationServer`: []OAuth2Client
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListOAuth2ClientsForAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ClientsForAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOAuth2Scopes

> []OAuth2Scope ListOAuth2Scopes(ctx, authServerId).Q(q).Filter(filter).Cursor(cursor).Limit(limit).Execute()

List all Custom Token Scopes



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
    authServerId := "authServerId_example" // string | 
    q := "q_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListOAuth2Scopes(context.Background(), authServerId).Q(q).Filter(filter).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListOAuth2Scopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2Scopes`: []OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListOAuth2Scopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOAuth2ScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 
 **filter** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

### Return type

[**[]OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRefreshTokensForAuthorizationServerAndClient

> []OAuth2RefreshToken ListRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId).Expand(expand).After(after).Limit(limit).Execute()

List all Refresh Tokens for a Client



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
    authServerId := "authServerId_example" // string | 
    clientId := "clientId_example" // string | 
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ListRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ListRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRefreshTokensForAuthorizationServerAndClient`: []OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ListRefreshTokensForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRefreshTokensForAuthorizationServerAndClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to -1]

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


## ReplaceAuthorizationServer

> AuthorizationServer ReplaceAuthorizationServer(ctx, authServerId).AuthorizationServer(authorizationServer).Execute()

Replace an Authorization Server



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
    authServerId := "authServerId_example" // string | 
    authorizationServer := *openapiclient.NewAuthorizationServer() // AuthorizationServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ReplaceAuthorizationServer(context.Background(), authServerId).AuthorizationServer(authorizationServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ReplaceAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ReplaceAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthorizationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationServer** | [**AuthorizationServer**](AuthorizationServer.md) |  | 

### Return type

[**AuthorizationServer**](AuthorizationServer.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthorizationServerPolicy

> AuthorizationServerPolicy ReplaceAuthorizationServerPolicy(ctx, authServerId, policyId).Policy(policy).Execute()

Replace a Policy



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
    authServerId := "authServerId_example" // string | 
    policyId := "policyId_example" // string | 
    policy := *openapiclient.NewAuthorizationServerPolicy() // AuthorizationServerPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ReplaceAuthorizationServerPolicy(context.Background(), authServerId, policyId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ReplaceAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ReplaceAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthorizationServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policy** | [**AuthorizationServerPolicy**](AuthorizationServerPolicy.md) |  | 

### Return type

[**AuthorizationServerPolicy**](AuthorizationServerPolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAuthorizationServerPolicyRule

> AuthorizationServerPolicyRule ReplaceAuthorizationServerPolicyRule(ctx, policyId, authServerId, ruleId).PolicyRule(policyRule).Execute()

Replace a Policy Rule



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
    policyId := "policyId_example" // string | 
    authServerId := "authServerId_example" // string | 
    ruleId := "ruleId_example" // string | 
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ReplaceAuthorizationServerPolicyRule(context.Background(), policyId, authServerId, ruleId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ReplaceAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ReplaceAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 
**authServerId** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAuthorizationServerPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **policyRule** | [**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md) |  | 

### Return type

[**AuthorizationServerPolicyRule**](AuthorizationServerPolicyRule.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOAuth2Claim

> OAuth2Claim ReplaceOAuth2Claim(ctx, authServerId, claimId).OAuth2Claim(oAuth2Claim).Execute()

Replace a Custom Token Claim



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
    authServerId := "authServerId_example" // string | 
    claimId := "claimId_example" // string | 
    oAuth2Claim := *openapiclient.NewOAuth2Claim() // OAuth2Claim | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ReplaceOAuth2Claim(context.Background(), authServerId, claimId).OAuth2Claim(oAuth2Claim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ReplaceOAuth2Claim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOAuth2Claim`: OAuth2Claim
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ReplaceOAuth2Claim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**claimId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOAuth2ClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oAuth2Claim** | [**OAuth2Claim**](OAuth2Claim.md) |  | 

### Return type

[**OAuth2Claim**](OAuth2Claim.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOAuth2Scope

> OAuth2Scope ReplaceOAuth2Scope(ctx, authServerId, scopeId).OAuth2Scope(oAuth2Scope).Execute()

Replace a Custom Token Scope



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
    authServerId := "authServerId_example" // string | 
    scopeId := "scopeId_example" // string | 
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.ReplaceOAuth2Scope(context.Background(), authServerId, scopeId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.ReplaceOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.ReplaceOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**scopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOAuth2ScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oAuth2Scope** | [**OAuth2Scope**](OAuth2Scope.md) |  | 

### Return type

[**OAuth2Scope**](OAuth2Scope.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeRefreshTokenForAuthorizationServerAndClient

> RevokeRefreshTokenForAuthorizationServerAndClient(ctx, authServerId, clientId, tokenId).Execute()

Revoke a Refresh Token for a Client



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
    authServerId := "authServerId_example" // string | 
    clientId := "clientId_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.RevokeRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.RevokeRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**clientId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRefreshTokenForAuthorizationServerAndClientRequest struct via the builder pattern


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


## RevokeRefreshTokensForAuthorizationServerAndClient

> RevokeRefreshTokensForAuthorizationServerAndClient(ctx, authServerId, clientId).Execute()

Revoke all Refresh Tokens for a Client



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
    authServerId := "authServerId_example" // string | 
    clientId := "clientId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.RevokeRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.RevokeRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRefreshTokensForAuthorizationServerAndClientRequest struct via the builder pattern


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


## RotateAuthorizationServerKeys

> []JsonWebKey RotateAuthorizationServerKeys(ctx, authServerId).Use(use).Execute()

Rotate all Credential Keys



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
    authServerId := "authServerId_example" // string | 
    use := *openapiclient.NewJwkUse() // JwkUse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerApi.RotateAuthorizationServerKeys(context.Background(), authServerId).Use(use).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerApi.RotateAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAuthorizationServerKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerApi.RotateAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAuthorizationServerKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **use** | [**JwkUse**](JwkUse.md) |  | 

### Return type

[**[]JsonWebKey**](JsonWebKey.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


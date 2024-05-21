# \AuthorizationServerAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAuthorizationServer**](AuthorizationServerAPI.md#ActivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/activate | Activate an Authorization Server
[**ActivateAuthorizationServerPolicy**](AuthorizationServerAPI.md#ActivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/activate | Activate a Policy
[**ActivateAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#ActivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a Policy Rule
[**CreateAuthorizationServer**](AuthorizationServerAPI.md#CreateAuthorizationServer) | **Post** /api/v1/authorizationServers | Create an Authorization Server
[**CreateAuthorizationServerPolicy**](AuthorizationServerAPI.md#CreateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies | Create a Policy
[**CreateAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#CreateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | Create a Policy Rule
[**CreateOAuth2Scope**](AuthorizationServerAPI.md#CreateOAuth2Scope) | **Post** /api/v1/authorizationServers/{authServerId}/scopes | Create a Custom Token Scope
[**DeactivateAuthorizationServer**](AuthorizationServerAPI.md#DeactivateAuthorizationServer) | **Post** /api/v1/authorizationServers/{authServerId}/lifecycle/deactivate | Deactivate an Authorization Server
[**DeactivateAuthorizationServerPolicy**](AuthorizationServerAPI.md#DeactivateAuthorizationServerPolicy) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/lifecycle/deactivate | Deactivate a Policy
[**DeactivateAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#DeactivateAuthorizationServerPolicyRule) | **Post** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a Policy Rule
[**DeleteAuthorizationServer**](AuthorizationServerAPI.md#DeleteAuthorizationServer) | **Delete** /api/v1/authorizationServers/{authServerId} | Delete an Authorization Server
[**DeleteAuthorizationServerPolicy**](AuthorizationServerAPI.md#DeleteAuthorizationServerPolicy) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Delete a Policy
[**DeleteAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#DeleteAuthorizationServerPolicyRule) | **Delete** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Delete a Policy Rule
[**DeleteOAuth2Scope**](AuthorizationServerAPI.md#DeleteOAuth2Scope) | **Delete** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Delete a Custom Token Scope
[**GetAuthorizationServer**](AuthorizationServerAPI.md#GetAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId} | Retrieve an Authorization Server
[**GetAuthorizationServerPolicy**](AuthorizationServerAPI.md#GetAuthorizationServerPolicy) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Retrieve a Policy
[**GetAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#GetAuthorizationServerPolicyRule) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Retrieve a Policy Rule
[**GetOAuth2Scope**](AuthorizationServerAPI.md#GetOAuth2Scope) | **Get** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Retrieve a Custom Token Scope
[**GetRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerAPI.md#GetRefreshTokenForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Retrieve a Refresh Token for a Client
[**ListAuthorizationServerKeys**](AuthorizationServerAPI.md#ListAuthorizationServerKeys) | **Get** /api/v1/authorizationServers/{authServerId}/credentials/keys | List all Credential Keys
[**ListAuthorizationServerPolicies**](AuthorizationServerAPI.md#ListAuthorizationServerPolicies) | **Get** /api/v1/authorizationServers/{authServerId}/policies | List all Policies
[**ListAuthorizationServerPolicyRules**](AuthorizationServerAPI.md#ListAuthorizationServerPolicyRules) | **Get** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules | List all Policy Rules
[**ListAuthorizationServers**](AuthorizationServerAPI.md#ListAuthorizationServers) | **Get** /api/v1/authorizationServers | List all Authorization Servers
[**ListOAuth2ClientsForAuthorizationServer**](AuthorizationServerAPI.md#ListOAuth2ClientsForAuthorizationServer) | **Get** /api/v1/authorizationServers/{authServerId}/clients | List all Clients
[**ListOAuth2Scopes**](AuthorizationServerAPI.md#ListOAuth2Scopes) | **Get** /api/v1/authorizationServers/{authServerId}/scopes | List all Custom Token Scopes
[**ListRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerAPI.md#ListRefreshTokensForAuthorizationServerAndClient) | **Get** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | List all Refresh Tokens for a Client
[**ReplaceAuthorizationServer**](AuthorizationServerAPI.md#ReplaceAuthorizationServer) | **Put** /api/v1/authorizationServers/{authServerId} | Replace an Authorization Server
[**ReplaceAuthorizationServerPolicy**](AuthorizationServerAPI.md#ReplaceAuthorizationServerPolicy) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId} | Replace a Policy
[**ReplaceAuthorizationServerPolicyRule**](AuthorizationServerAPI.md#ReplaceAuthorizationServerPolicyRule) | **Put** /api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId} | Replace a Policy Rule
[**ReplaceOAuth2Scope**](AuthorizationServerAPI.md#ReplaceOAuth2Scope) | **Put** /api/v1/authorizationServers/{authServerId}/scopes/{scopeId} | Replace a Custom Token Scope
[**RevokeRefreshTokenForAuthorizationServerAndClient**](AuthorizationServerAPI.md#RevokeRefreshTokenForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens/{tokenId} | Revoke a Refresh Token for a Client
[**RevokeRefreshTokensForAuthorizationServerAndClient**](AuthorizationServerAPI.md#RevokeRefreshTokensForAuthorizationServerAndClient) | **Delete** /api/v1/authorizationServers/{authServerId}/clients/{clientId}/tokens | Revoke all Refresh Tokens for a Client
[**RotateAuthorizationServerKeys**](AuthorizationServerAPI.md#RotateAuthorizationServerKeys) | **Post** /api/v1/authorizationServers/{authServerId}/credentials/lifecycle/keyRotate | Rotate all Credential Keys



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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.ActivateAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ActivateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.ActivateAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ActivateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.ActivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ActivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authorizationServer := *openapiclient.NewAuthorizationServer() // AuthorizationServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.CreateAuthorizationServer(context.Background()).AuthorizationServer(authorizationServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.CreateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.CreateAuthorizationServer`: %v\n", resp)
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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policy := *openapiclient.NewAuthorizationServerPolicy() // AuthorizationServerPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.CreateAuthorizationServerPolicy(context.Background(), authServerId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.CreateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.CreateAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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

> AuthorizationServerPolicyRule CreateAuthorizationServerPolicyRule(ctx, authServerId, policyId).PolicyRule(policyRule).Execute()

Create a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.CreateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.CreateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.CreateAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.CreateOAuth2Scope(context.Background(), authServerId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.CreateOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.CreateOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeactivateAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeactivateAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeactivateAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeactivateAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeactivateAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeactivateAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeleteAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeleteAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeleteAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeleteAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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

> DeleteAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Delete a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeleteAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeleteAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.DeleteOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.DeleteOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.GetAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.GetAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.GetAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.GetAuthorizationServerPolicy(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.GetAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.GetAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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

> AuthorizationServerPolicyRule GetAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).Execute()

Retrieve a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.GetAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.GetAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.GetAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.GetOAuth2Scope(context.Background(), authServerId, scopeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.GetOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.GetOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.GetRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.GetRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefreshTokenForAuthorizationServerAndClient`: OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.GetRefreshTokenForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListAuthorizationServerKeys(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListAuthorizationServerPolicies(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListAuthorizationServerPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerPolicies`: []AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListAuthorizationServerPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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

> []AuthorizationServerPolicyRule ListAuthorizationServerPolicyRules(ctx, authServerId, policyId).Execute()

List all Policy Rules



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListAuthorizationServerPolicyRules(context.Background(), authServerId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListAuthorizationServerPolicyRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServerPolicyRules`: []AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListAuthorizationServerPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    q := "customasone" // string | Searches the `name` and `audiences` of authorization servers for matching values (optional)
    limit := int32(56) // int32 | Specifies the number of authorization server results on a page. Maximum value: 200 (optional) (default to 200)
    after := "after_example" // string | Specifies the pagination cursor for the next page of authorization servers. Treat as an opaque value and obtain through the next link relationship. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListAuthorizationServers(context.Background()).Q(q).Limit(limit).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListAuthorizationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizationServers`: []AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListAuthorizationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Searches the &#x60;name&#x60; and &#x60;audiences&#x60; of authorization servers for matching values | 
 **limit** | **int32** | Specifies the number of authorization server results on a page. Maximum value: 200 | [default to 200]
 **after** | **string** | Specifies the pagination cursor for the next page of authorization servers. Treat as an opaque value and obtain through the next link relationship. | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListOAuth2ClientsForAuthorizationServer(context.Background(), authServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListOAuth2ClientsForAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2ClientsForAuthorizationServer`: []OAuth2Client
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListOAuth2ClientsForAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    q := "q_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListOAuth2Scopes(context.Background(), authServerId).Q(q).Filter(filter).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListOAuth2Scopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOAuth2Scopes`: []OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListOAuth2Scopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    expand := "expand_example" // string |  (optional)
    after := "after_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ListRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Expand(expand).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ListRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRefreshTokensForAuthorizationServerAndClient`: []OAuth2RefreshToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ListRefreshTokensForAuthorizationServerAndClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    authorizationServer := *openapiclient.NewAuthorizationServer() // AuthorizationServer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ReplaceAuthorizationServer(context.Background(), authServerId).AuthorizationServer(authorizationServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ReplaceAuthorizationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServer`: AuthorizationServer
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ReplaceAuthorizationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policy := *openapiclient.NewAuthorizationServerPolicy() // AuthorizationServerPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ReplaceAuthorizationServerPolicy(context.Background(), authServerId, policyId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ReplaceAuthorizationServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServerPolicy`: AuthorizationServerPolicy
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ReplaceAuthorizationServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

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

> AuthorizationServerPolicyRule ReplaceAuthorizationServerPolicyRule(ctx, authServerId, policyId, ruleId).PolicyRule(policyRule).Execute()

Replace a Policy Rule



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
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule
    policyRule := *openapiclient.NewAuthorizationServerPolicyRule() // AuthorizationServerPolicyRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ReplaceAuthorizationServerPolicyRule(context.Background(), authServerId, policyId, ruleId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ReplaceAuthorizationServerPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceAuthorizationServerPolicyRule`: AuthorizationServerPolicyRule
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ReplaceAuthorizationServerPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    scopeId := "0TMRpCWXRKFjP7HiPFNM" // string | `id` of Scope
    oAuth2Scope := *openapiclient.NewOAuth2Scope() // OAuth2Scope | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.ReplaceOAuth2Scope(context.Background(), authServerId, scopeId).OAuth2Scope(oAuth2Scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.ReplaceOAuth2Scope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOAuth2Scope`: OAuth2Scope
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.ReplaceOAuth2Scope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**scopeId** | **string** | &#x60;id&#x60; of Scope | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app
    tokenId := "sHHSth53yJAyNSTQKDJZ" // string | `id` of Token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.RevokeRefreshTokenForAuthorizationServerAndClient(context.Background(), authServerId, clientId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.RevokeRefreshTokenForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 
**tokenId** | **string** | &#x60;id&#x60; of Token | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    clientId := "52Uy4BUWVBOjFItcg2jWsmnd83Ad8dD" // string | `client_id` of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationServerAPI.RevokeRefreshTokensForAuthorizationServerAndClient(context.Background(), authServerId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.RevokeRefreshTokensForAuthorizationServerAndClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 
**clientId** | **string** | &#x60;client_id&#x60; of the app | 

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
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    authServerId := "GeGRTEr7f3yu2n7grw22" // string | `id` of the Authorization Server
    use := *openapiclient.NewJwkUse() // JwkUse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationServerAPI.RotateAuthorizationServerKeys(context.Background(), authServerId).Use(use).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationServerAPI.RotateAuthorizationServerKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAuthorizationServerKeys`: []JsonWebKey
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationServerAPI.RotateAuthorizationServerKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authServerId** | **string** | &#x60;id&#x60; of the Authorization Server | 

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


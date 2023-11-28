# \OrgSettingAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignClientPrivilegesSetting**](OrgSettingAPI.md#AssignClientPrivilegesSetting) | **Put** /api/v1/org/settings/clientPrivilegesSetting | Assign the Super Admin role to a public client app
[**BulkRemoveEmailAddressBounces**](OrgSettingAPI.md#BulkRemoveEmailAddressBounces) | **Post** /api/v1/org/email/bounces/remove-list | Remove Emails from Email Provider Bounce List
[**ExtendOktaSupport**](OrgSettingAPI.md#ExtendOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/extend | Extend Okta Support Access
[**GetClientPrivilegesSetting**](OrgSettingAPI.md#GetClientPrivilegesSetting) | **Get** /api/v1/org/settings/clientPrivilegesSetting | Retrieve the Org settings to assign the Super Admin role
[**GetOktaCommunicationSettings**](OrgSettingAPI.md#GetOktaCommunicationSettings) | **Get** /api/v1/org/privacy/oktaCommunication | Retrieve the Okta Communication Settings
[**GetOrgContactTypes**](OrgSettingAPI.md#GetOrgContactTypes) | **Get** /api/v1/org/contacts | Retrieve the Org Contact Types
[**GetOrgContactUser**](OrgSettingAPI.md#GetOrgContactUser) | **Get** /api/v1/org/contacts/{contactType} | Retrieve the User of the Contact Type
[**GetOrgOktaSupportSettings**](OrgSettingAPI.md#GetOrgOktaSupportSettings) | **Get** /api/v1/org/privacy/oktaSupport | Retrieve the Okta Support Settings
[**GetOrgPreferences**](OrgSettingAPI.md#GetOrgPreferences) | **Get** /api/v1/org/preferences | Retrieve the Org Preferences
[**GetOrgSettings**](OrgSettingAPI.md#GetOrgSettings) | **Get** /api/v1/org | Retrieve the Org Settings
[**GetThirdPartyAdminSetting**](OrgSettingAPI.md#GetThirdPartyAdminSetting) | **Get** /api/v1/org/orgSettings/thirdPartyAdminSetting | Retrieve the Org Third-Party Admin setting
[**GetWellknownOrgMetadata**](OrgSettingAPI.md#GetWellknownOrgMetadata) | **Get** /.well-known/okta-organization | Retrieve the Well-Known Org Metadata
[**GrantOktaSupport**](OrgSettingAPI.md#GrantOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/grant | Grant Okta Support Access to your Org
[**OptInUsersToOktaCommunicationEmails**](OrgSettingAPI.md#OptInUsersToOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optIn | Opt in all Users to Okta Communication emails
[**OptOutUsersFromOktaCommunicationEmails**](OrgSettingAPI.md#OptOutUsersFromOktaCommunicationEmails) | **Post** /api/v1/org/privacy/oktaCommunication/optOut | Opt out all Users from Okta Communication emails
[**ReplaceOrgContactUser**](OrgSettingAPI.md#ReplaceOrgContactUser) | **Put** /api/v1/org/contacts/{contactType} | Replace the User of the Contact Type
[**ReplaceOrgSettings**](OrgSettingAPI.md#ReplaceOrgSettings) | **Put** /api/v1/org | Replace the Org Settings
[**RevokeOktaSupport**](OrgSettingAPI.md#RevokeOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/revoke | Revoke Okta Support Access
[**UpdateOrgHideOktaUIFooter**](OrgSettingAPI.md#UpdateOrgHideOktaUIFooter) | **Post** /api/v1/org/preferences/hideEndUserFooter | Update the Preference to Hide the Okta Dashboard Footer
[**UpdateOrgSettings**](OrgSettingAPI.md#UpdateOrgSettings) | **Post** /api/v1/org | Update the Org Settings
[**UpdateOrgShowOktaUIFooter**](OrgSettingAPI.md#UpdateOrgShowOktaUIFooter) | **Post** /api/v1/org/preferences/showEndUserFooter | Update the Preference to Show the Okta Dashboard Footer
[**UpdateThirdPartyAdminSetting**](OrgSettingAPI.md#UpdateThirdPartyAdminSetting) | **Post** /api/v1/org/orgSettings/thirdPartyAdminSetting | Update the Org Third-Party Admin setting
[**UploadOrgLogo**](OrgSettingAPI.md#UploadOrgLogo) | **Post** /api/v1/org/logo | Upload the Org Logo



## AssignClientPrivilegesSetting

> ClientPrivilegesSetting AssignClientPrivilegesSetting(ctx).ClientPrivilegesSetting(clientPrivilegesSetting).Execute()

Assign the Super Admin role to a public client app



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
    clientPrivilegesSetting := *openapiclient.NewClientPrivilegesSetting() // ClientPrivilegesSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.AssignClientPrivilegesSetting(context.Background()).ClientPrivilegesSetting(clientPrivilegesSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.AssignClientPrivilegesSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignClientPrivilegesSetting`: ClientPrivilegesSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.AssignClientPrivilegesSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignClientPrivilegesSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientPrivilegesSetting** | [**ClientPrivilegesSetting**](ClientPrivilegesSetting.md) |  | 

### Return type

[**ClientPrivilegesSetting**](ClientPrivilegesSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRemoveEmailAddressBounces

> BouncesRemoveListResult BulkRemoveEmailAddressBounces(ctx).BouncesRemoveListObj(bouncesRemoveListObj).Execute()

Remove Emails from Email Provider Bounce List



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
    bouncesRemoveListObj := *openapiclient.NewBouncesRemoveListObj() // BouncesRemoveListObj |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.BulkRemoveEmailAddressBounces(context.Background()).BouncesRemoveListObj(bouncesRemoveListObj).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.BulkRemoveEmailAddressBounces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkRemoveEmailAddressBounces`: BouncesRemoveListResult
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.BulkRemoveEmailAddressBounces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRemoveEmailAddressBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bouncesRemoveListObj** | [**BouncesRemoveListObj**](BouncesRemoveListObj.md) |  | 

### Return type

[**BouncesRemoveListResult**](BouncesRemoveListResult.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendOktaSupport

> OrgOktaSupportSettingsObj ExtendOktaSupport(ctx).Execute()

Extend Okta Support Access



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.ExtendOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.ExtendOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtendOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.ExtendOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExtendOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientPrivilegesSetting

> ClientPrivilegesSetting GetClientPrivilegesSetting(ctx).Execute()

Retrieve the Org settings to assign the Super Admin role



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetClientPrivilegesSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetClientPrivilegesSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientPrivilegesSetting`: ClientPrivilegesSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetClientPrivilegesSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientPrivilegesSettingRequest struct via the builder pattern


### Return type

[**ClientPrivilegesSetting**](ClientPrivilegesSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOktaCommunicationSettings

> OrgOktaCommunicationSetting GetOktaCommunicationSettings(ctx).Execute()

Retrieve the Okta Communication Settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOktaCommunicationSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOktaCommunicationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOktaCommunicationSettings`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOktaCommunicationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOktaCommunicationSettingsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgContactTypes

> []OrgContactTypeObj GetOrgContactTypes(ctx).Execute()

Retrieve the Org Contact Types



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOrgContactTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOrgContactTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgContactTypes`: []OrgContactTypeObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOrgContactTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgContactTypesRequest struct via the builder pattern


### Return type

[**[]OrgContactTypeObj**](OrgContactTypeObj.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgContactUser

> OrgContactUser GetOrgContactUser(ctx, contactType).Execute()

Retrieve the User of the Contact Type



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
    contactType := "contactType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOrgContactUser(context.Background(), contactType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOrgContactUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgContactUser`: OrgContactUser
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOrgContactUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgContactUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgOktaSupportSettings

> OrgOktaSupportSettingsObj GetOrgOktaSupportSettings(ctx).Execute()

Retrieve the Okta Support Settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOrgOktaSupportSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOrgOktaSupportSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgOktaSupportSettings`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOrgOktaSupportSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgOktaSupportSettingsRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgPreferences

> OrgPreferences GetOrgPreferences(ctx).Execute()

Retrieve the Org Preferences



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOrgPreferences(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOrgPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgPreferences`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOrgPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPreferencesRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSettings

> OrgSetting GetOrgSettings(ctx).Execute()

Retrieve the Org Settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetOrgSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgSettings`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThirdPartyAdminSetting

> ThirdPartyAdminSetting GetThirdPartyAdminSetting(ctx).Execute()

Retrieve the Org Third-Party Admin setting



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetThirdPartyAdminSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetThirdPartyAdminSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThirdPartyAdminSetting`: ThirdPartyAdminSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetThirdPartyAdminSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyAdminSettingRequest struct via the builder pattern


### Return type

[**ThirdPartyAdminSetting**](ThirdPartyAdminSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWellknownOrgMetadata

> WellKnownOrgMetadata GetWellknownOrgMetadata(ctx).Execute()

Retrieve the Well-Known Org Metadata



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GetWellknownOrgMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GetWellknownOrgMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWellknownOrgMetadata`: WellKnownOrgMetadata
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GetWellknownOrgMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWellknownOrgMetadataRequest struct via the builder pattern


### Return type

[**WellKnownOrgMetadata**](WellKnownOrgMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantOktaSupport

> OrgOktaSupportSettingsObj GrantOktaSupport(ctx).Execute()

Grant Okta Support Access to your Org



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.GrantOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.GrantOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.GrantOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGrantOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptInUsersToOktaCommunicationEmails

> OrgOktaCommunicationSetting OptInUsersToOktaCommunicationEmails(ctx).Execute()

Opt in all Users to Okta Communication emails



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.OptInUsersToOktaCommunicationEmails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.OptInUsersToOktaCommunicationEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptInUsersToOktaCommunicationEmails`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.OptInUsersToOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptInUsersToOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOutUsersFromOktaCommunicationEmails

> OrgOktaCommunicationSetting OptOutUsersFromOktaCommunicationEmails(ctx).Execute()

Opt out all Users from Okta Communication emails



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.OptOutUsersFromOktaCommunicationEmails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.OptOutUsersFromOktaCommunicationEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptOutUsersFromOktaCommunicationEmails`: OrgOktaCommunicationSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.OptOutUsersFromOktaCommunicationEmails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptOutUsersFromOktaCommunicationEmailsRequest struct via the builder pattern


### Return type

[**OrgOktaCommunicationSetting**](OrgOktaCommunicationSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrgContactUser

> OrgContactUser ReplaceOrgContactUser(ctx, contactType).OrgContactUser(orgContactUser).Execute()

Replace the User of the Contact Type



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
    contactType := "contactType_example" // string | 
    orgContactUser := *openapiclient.NewOrgContactUser() // OrgContactUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.ReplaceOrgContactUser(context.Background(), contactType).OrgContactUser(orgContactUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.ReplaceOrgContactUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOrgContactUser`: OrgContactUser
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.ReplaceOrgContactUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrgContactUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgContactUser** | [**OrgContactUser**](OrgContactUser.md) |  | 

### Return type

[**OrgContactUser**](OrgContactUser.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrgSettings

> OrgSetting ReplaceOrgSettings(ctx).OrgSetting(orgSetting).Execute()

Replace the Org Settings



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
    orgSetting := *openapiclient.NewOrgSetting() // OrgSetting | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.ReplaceOrgSettings(context.Background()).OrgSetting(orgSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.ReplaceOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceOrgSettings`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.ReplaceOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOktaSupport

> OrgOktaSupportSettingsObj RevokeOktaSupport(ctx).Execute()

Revoke Okta Support Access



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.RevokeOktaSupport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.RevokeOktaSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeOktaSupport`: OrgOktaSupportSettingsObj
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.RevokeOktaSupport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOktaSupportRequest struct via the builder pattern


### Return type

[**OrgOktaSupportSettingsObj**](OrgOktaSupportSettingsObj.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgHideOktaUIFooter

> OrgPreferences UpdateOrgHideOktaUIFooter(ctx).Execute()

Update the Preference to Hide the Okta Dashboard Footer



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.UpdateOrgHideOktaUIFooter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.UpdateOrgHideOktaUIFooter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgHideOktaUIFooter`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.UpdateOrgHideOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgHideOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> OrgSetting UpdateOrgSettings(ctx).OrgSetting(orgSetting).Execute()

Update the Org Settings



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
    orgSetting := *openapiclient.NewOrgSetting() // OrgSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.UpdateOrgSettings(context.Background()).OrgSetting(orgSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.UpdateOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgSettings`: OrgSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSetting** | [**OrgSetting**](OrgSetting.md) |  | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgShowOktaUIFooter

> OrgPreferences UpdateOrgShowOktaUIFooter(ctx).Execute()

Update the Preference to Show the Okta Dashboard Footer



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.UpdateOrgShowOktaUIFooter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.UpdateOrgShowOktaUIFooter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgShowOktaUIFooter`: OrgPreferences
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.UpdateOrgShowOktaUIFooter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgShowOktaUIFooterRequest struct via the builder pattern


### Return type

[**OrgPreferences**](OrgPreferences.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateThirdPartyAdminSetting

> ThirdPartyAdminSetting UpdateThirdPartyAdminSetting(ctx).Execute()

Update the Org Third-Party Admin setting



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgSettingAPI.UpdateThirdPartyAdminSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.UpdateThirdPartyAdminSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateThirdPartyAdminSetting`: ThirdPartyAdminSetting
    fmt.Fprintf(os.Stdout, "Response from `OrgSettingAPI.UpdateThirdPartyAdminSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateThirdPartyAdminSettingRequest struct via the builder pattern


### Return type

[**ThirdPartyAdminSetting**](ThirdPartyAdminSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadOrgLogo

> UploadOrgLogo(ctx).File(file).Execute()

Upload the Org Logo



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
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrgSettingAPI.UploadOrgLogo(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingAPI.UploadOrgLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadOrgLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


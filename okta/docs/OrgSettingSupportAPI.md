# \OrgSettingSupportAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendOktaSupport**](OrgSettingSupportAPI.md#ExtendOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/extend | Extend Okta Support access
[**GetAerialConsent**](OrgSettingSupportAPI.md#GetAerialConsent) | **Get** /api/v1/org/privacy/aerial | Retrieve Okta Aerial consent for your org
[**GetOrgOktaSupportSettings**](OrgSettingSupportAPI.md#GetOrgOktaSupportSettings) | **Get** /api/v1/org/privacy/oktaSupport | Retrieve the Okta Support settings
[**GrantAerialConsent**](OrgSettingSupportAPI.md#GrantAerialConsent) | **Post** /api/v1/org/privacy/aerial/grant | Grant Okta Aerial access to your org
[**GrantOktaSupport**](OrgSettingSupportAPI.md#GrantOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/grant | Grant Okta Support access
[**ListOktaSupportCases**](OrgSettingSupportAPI.md#ListOktaSupportCases) | **Get** /api/v1/org/privacy/oktaSupport/cases | List all Okta Support cases
[**RevokeAerialConsent**](OrgSettingSupportAPI.md#RevokeAerialConsent) | **Post** /api/v1/org/privacy/aerial/revoke | Revoke Okta Aerial access to your org
[**RevokeOktaSupport**](OrgSettingSupportAPI.md#RevokeOktaSupport) | **Post** /api/v1/org/privacy/oktaSupport/revoke | Revoke Okta Support access
[**UpdateOktaSupportCase**](OrgSettingSupportAPI.md#UpdateOktaSupportCase) | **Patch** /api/v1/org/privacy/oktaSupport/cases/{caseNumber} | Update an Okta Support case



## ExtendOktaSupport

> ExtendOktaSupport(ctx).Execute()

Extend Okta Support access



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
	r, err := apiClient.OrgSettingSupportAPI.ExtendOktaSupport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.ExtendOktaSupport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExtendOktaSupportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAerialConsent

> OrgAerialConsentDetails GetAerialConsent(ctx).Execute()

Retrieve Okta Aerial consent for your org



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
	resp, r, err := apiClient.OrgSettingSupportAPI.GetAerialConsent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.GetAerialConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAerialConsent`: OrgAerialConsentDetails
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.GetAerialConsent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAerialConsentRequest struct via the builder pattern


### Return type

[**OrgAerialConsentDetails**](OrgAerialConsentDetails.md)

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

Retrieve the Okta Support settings



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
	resp, r, err := apiClient.OrgSettingSupportAPI.GetOrgOktaSupportSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.GetOrgOktaSupportSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgOktaSupportSettings`: OrgOktaSupportSettingsObj
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.GetOrgOktaSupportSettings`: %v\n", resp)
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


## GrantAerialConsent

> OrgAerialConsentDetails GrantAerialConsent(ctx).OrgAerialConsent(orgAerialConsent).Execute()

Grant Okta Aerial access to your org



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
	orgAerialConsent := *openapiclient.NewOrgAerialConsent("AccountId_example") // OrgAerialConsent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingSupportAPI.GrantAerialConsent(context.Background()).OrgAerialConsent(orgAerialConsent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.GrantAerialConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantAerialConsent`: OrgAerialConsentDetails
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.GrantAerialConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGrantAerialConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgAerialConsent** | [**OrgAerialConsent**](OrgAerialConsent.md) |  | 

### Return type

[**OrgAerialConsentDetails**](OrgAerialConsentDetails.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantOktaSupport

> GrantOktaSupport(ctx).Execute()

Grant Okta Support access



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
	r, err := apiClient.OrgSettingSupportAPI.GrantOktaSupport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.GrantOktaSupport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGrantOktaSupportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOktaSupportCases

> OktaSupportCases ListOktaSupportCases(ctx).Execute()

List all Okta Support cases



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
	resp, r, err := apiClient.OrgSettingSupportAPI.ListOktaSupportCases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.ListOktaSupportCases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOktaSupportCases`: OktaSupportCases
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.ListOktaSupportCases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOktaSupportCasesRequest struct via the builder pattern


### Return type

[**OktaSupportCases**](OktaSupportCases.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAerialConsent

> OrgAerialConsentRevoked RevokeAerialConsent(ctx).OrgAerialConsent(orgAerialConsent).Execute()

Revoke Okta Aerial access to your org



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
	orgAerialConsent := *openapiclient.NewOrgAerialConsent("AccountId_example") // OrgAerialConsent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingSupportAPI.RevokeAerialConsent(context.Background()).OrgAerialConsent(orgAerialConsent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.RevokeAerialConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeAerialConsent`: OrgAerialConsentRevoked
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.RevokeAerialConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAerialConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgAerialConsent** | [**OrgAerialConsent**](OrgAerialConsent.md) |  | 

### Return type

[**OrgAerialConsentRevoked**](OrgAerialConsentRevoked.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOktaSupport

> RevokeOktaSupport(ctx).Execute()

Revoke Okta Support access



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
	r, err := apiClient.OrgSettingSupportAPI.RevokeOktaSupport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.RevokeOktaSupport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOktaSupportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOktaSupportCase

> OktaSupportCase UpdateOktaSupportCase(ctx, caseNumber).OktaSupportCase(oktaSupportCase).Execute()

Update an Okta Support case



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
	caseNumber := "00000144" // string | 
	oktaSupportCase := *openapiclient.NewOktaSupportCase() // OktaSupportCase |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingSupportAPI.UpdateOktaSupportCase(context.Background(), caseNumber).OktaSupportCase(oktaSupportCase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingSupportAPI.UpdateOktaSupportCase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOktaSupportCase`: OktaSupportCase
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingSupportAPI.UpdateOktaSupportCase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOktaSupportCaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oktaSupportCase** | [**OktaSupportCase**](OktaSupportCase.md) |  | 

### Return type

[**OktaSupportCase**](OktaSupportCase.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


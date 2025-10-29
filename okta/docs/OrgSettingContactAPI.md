# \OrgSettingContactAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgContactUser**](OrgSettingContactAPI.md#GetOrgContactUser) | **Get** /api/v1/org/contacts/{contactType} | Retrieve the contact type user
[**ListOrgContactTypes**](OrgSettingContactAPI.md#ListOrgContactTypes) | **Get** /api/v1/org/contacts | List all org contact types
[**ReplaceOrgContactUser**](OrgSettingContactAPI.md#ReplaceOrgContactUser) | **Put** /api/v1/org/contacts/{contactType} | Replace the contact type user



## GetOrgContactUser

> OrgContactUser GetOrgContactUser(ctx, contactType).Execute()

Retrieve the contact type user



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
	contactType := "BILLING" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingContactAPI.GetOrgContactUser(context.Background(), contactType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingContactAPI.GetOrgContactUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgContactUser`: OrgContactUser
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingContactAPI.GetOrgContactUser`: %v\n", resp)
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


## ListOrgContactTypes

> []ListOrgContactTypes200ResponseInner ListOrgContactTypes(ctx).Execute()

List all org contact types



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
	resp, r, err := apiClient.OrgSettingContactAPI.ListOrgContactTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingContactAPI.ListOrgContactTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgContactTypes`: []ListOrgContactTypes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingContactAPI.ListOrgContactTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgContactTypesRequest struct via the builder pattern


### Return type

[**[]ListOrgContactTypes200ResponseInner**](ListOrgContactTypes200ResponseInner.md)

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

Replace the contact type user



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
	contactType := "BILLING" // string | 
	orgContactUser := *openapiclient.NewOrgContactUser() // OrgContactUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgSettingContactAPI.ReplaceOrgContactUser(context.Background(), contactType).OrgContactUser(orgContactUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgSettingContactAPI.ReplaceOrgContactUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceOrgContactUser`: OrgContactUser
	fmt.Fprintf(os.Stdout, "Response from `OrgSettingContactAPI.ReplaceOrgContactUser`: %v\n", resp)
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


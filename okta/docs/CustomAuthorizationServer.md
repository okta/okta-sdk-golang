# CustomAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerUrl** | **string** | Issuer URL for the authorization server | 
**Logo** | Pointer to **string** | Image URL for the authorization server | [optional] 
**Name** | **string** | Display name of the authorization server | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the authorization server | 
**ResourceIndicator** | Pointer to **string** | Audience value to use when requesting tokens | [optional] 
**Links** | [**CustomAuthorizationServerLinks**](CustomAuthorizationServerLinks.md) |  | 

## Methods

### NewCustomAuthorizationServer

`func NewCustomAuthorizationServer(issuerUrl string, name string, orn string, links CustomAuthorizationServerLinks, ) *CustomAuthorizationServer`

NewCustomAuthorizationServer instantiates a new CustomAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAuthorizationServerWithDefaults

`func NewCustomAuthorizationServerWithDefaults() *CustomAuthorizationServer`

NewCustomAuthorizationServerWithDefaults instantiates a new CustomAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerUrl

`func (o *CustomAuthorizationServer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CustomAuthorizationServer) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CustomAuthorizationServer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.


### GetLogo

`func (o *CustomAuthorizationServer) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CustomAuthorizationServer) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CustomAuthorizationServer) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CustomAuthorizationServer) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *CustomAuthorizationServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAuthorizationServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAuthorizationServer) SetName(v string)`

SetName sets Name field to given value.


### GetOrn

`func (o *CustomAuthorizationServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *CustomAuthorizationServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *CustomAuthorizationServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetResourceIndicator

`func (o *CustomAuthorizationServer) GetResourceIndicator() string`

GetResourceIndicator returns the ResourceIndicator field if non-nil, zero value otherwise.

### GetResourceIndicatorOk

`func (o *CustomAuthorizationServer) GetResourceIndicatorOk() (*string, bool)`

GetResourceIndicatorOk returns a tuple with the ResourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIndicator

`func (o *CustomAuthorizationServer) SetResourceIndicator(v string)`

SetResourceIndicator sets ResourceIndicator field to given value.

### HasResourceIndicator

`func (o *CustomAuthorizationServer) HasResourceIndicator() bool`

HasResourceIndicator returns a boolean if a field has been set.

### GetLinks

`func (o *CustomAuthorizationServer) GetLinks() CustomAuthorizationServerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomAuthorizationServer) GetLinksOk() (*CustomAuthorizationServerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomAuthorizationServer) SetLinks(v CustomAuthorizationServerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



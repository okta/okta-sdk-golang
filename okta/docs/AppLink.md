# AppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to [**HrefObjectAppLink**](HrefObjectAppLink.md) |  | [optional] 
**Logo** | Pointer to [**HrefObjectLogoLink**](HrefObjectLogoLink.md) |  | [optional] 

## Methods

### NewAppLink

`func NewAppLink() *AppLink`

NewAppLink instantiates a new AppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkWithDefaults

`func NewAppLinkWithDefaults() *AppLink`

NewAppLinkWithDefaults instantiates a new AppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *AppLink) GetLogin() HrefObjectAppLink`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AppLink) GetLoginOk() (*HrefObjectAppLink, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AppLink) SetLogin(v HrefObjectAppLink)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AppLink) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLogo

`func (o *AppLink) GetLogo() HrefObjectLogoLink`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AppLink) GetLogoOk() (*HrefObjectLogoLink, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AppLink) SetLogo(v HrefObjectLogoLink)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AppLink) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ManagedConnectionAppInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to **string** | Image URL for the app logo | [optional] 
**Name** | **string** | Display name of the app | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the app instance | 
**Links** | [**CustomAuthorizationServerLinks**](CustomAuthorizationServerLinks.md) |  | 

## Methods

### NewManagedConnectionAppInstance

`func NewManagedConnectionAppInstance(name string, orn string, links CustomAuthorizationServerLinks, ) *ManagedConnectionAppInstance`

NewManagedConnectionAppInstance instantiates a new ManagedConnectionAppInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionAppInstanceWithDefaults

`func NewManagedConnectionAppInstanceWithDefaults() *ManagedConnectionAppInstance`

NewManagedConnectionAppInstanceWithDefaults instantiates a new ManagedConnectionAppInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *ManagedConnectionAppInstance) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ManagedConnectionAppInstance) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ManagedConnectionAppInstance) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ManagedConnectionAppInstance) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *ManagedConnectionAppInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedConnectionAppInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedConnectionAppInstance) SetName(v string)`

SetName sets Name field to given value.


### GetOrn

`func (o *ManagedConnectionAppInstance) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ManagedConnectionAppInstance) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ManagedConnectionAppInstance) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetLinks

`func (o *ManagedConnectionAppInstance) GetLinks() CustomAuthorizationServerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManagedConnectionAppInstance) GetLinksOk() (*CustomAuthorizationServerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManagedConnectionAppInstance) SetLinks(v CustomAuthorizationServerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



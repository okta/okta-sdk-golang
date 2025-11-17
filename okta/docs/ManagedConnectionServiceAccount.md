# ManagedConnectionServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the service account | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the service account | 
**Links** | [**CustomAuthorizationServerLinks**](CustomAuthorizationServerLinks.md) |  | 

## Methods

### NewManagedConnectionServiceAccount

`func NewManagedConnectionServiceAccount(name string, orn string, links CustomAuthorizationServerLinks, ) *ManagedConnectionServiceAccount`

NewManagedConnectionServiceAccount instantiates a new ManagedConnectionServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionServiceAccountWithDefaults

`func NewManagedConnectionServiceAccountWithDefaults() *ManagedConnectionServiceAccount`

NewManagedConnectionServiceAccountWithDefaults instantiates a new ManagedConnectionServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagedConnectionServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedConnectionServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedConnectionServiceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetOrn

`func (o *ManagedConnectionServiceAccount) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ManagedConnectionServiceAccount) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ManagedConnectionServiceAccount) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetLinks

`func (o *ManagedConnectionServiceAccount) GetLinks() CustomAuthorizationServerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManagedConnectionServiceAccount) GetLinksOk() (*CustomAuthorizationServerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManagedConnectionServiceAccount) SetLinks(v CustomAuthorizationServerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



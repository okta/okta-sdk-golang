# ManagedConnectionVaultedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description of the secret | [optional] 
**Name** | **string** | Display name of the secret | 
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the vaulted secret | 
**Path** | Pointer to **string** | Secret path in Okta Privileged Access (OPA) | [optional] 
**Links** | [**CustomAuthorizationServerLinks**](CustomAuthorizationServerLinks.md) |  | 

## Methods

### NewManagedConnectionVaultedSecret

`func NewManagedConnectionVaultedSecret(name string, orn string, links CustomAuthorizationServerLinks, ) *ManagedConnectionVaultedSecret`

NewManagedConnectionVaultedSecret instantiates a new ManagedConnectionVaultedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionVaultedSecretWithDefaults

`func NewManagedConnectionVaultedSecretWithDefaults() *ManagedConnectionVaultedSecret`

NewManagedConnectionVaultedSecretWithDefaults instantiates a new ManagedConnectionVaultedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManagedConnectionVaultedSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedConnectionVaultedSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedConnectionVaultedSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedConnectionVaultedSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ManagedConnectionVaultedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedConnectionVaultedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedConnectionVaultedSecret) SetName(v string)`

SetName sets Name field to given value.


### GetOrn

`func (o *ManagedConnectionVaultedSecret) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ManagedConnectionVaultedSecret) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ManagedConnectionVaultedSecret) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetPath

`func (o *ManagedConnectionVaultedSecret) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManagedConnectionVaultedSecret) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManagedConnectionVaultedSecret) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManagedConnectionVaultedSecret) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLinks

`func (o *ManagedConnectionVaultedSecret) GetLinks() CustomAuthorizationServerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManagedConnectionVaultedSecret) GetLinksOk() (*CustomAuthorizationServerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManagedConnectionVaultedSecret) SetLinks(v CustomAuthorizationServerLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



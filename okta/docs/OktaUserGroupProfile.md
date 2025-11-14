# OktaUserGroupProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the group | [optional] 
**Name** | Pointer to **string** | Name of the group | [optional] 
**ObjectClass** | Pointer to **string** | The object class type | [optional] [readonly] 

## Methods

### NewOktaUserGroupProfile

`func NewOktaUserGroupProfile() *OktaUserGroupProfile`

NewOktaUserGroupProfile instantiates a new OktaUserGroupProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaUserGroupProfileWithDefaults

`func NewOktaUserGroupProfileWithDefaults() *OktaUserGroupProfile`

NewOktaUserGroupProfileWithDefaults instantiates a new OktaUserGroupProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OktaUserGroupProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OktaUserGroupProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OktaUserGroupProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OktaUserGroupProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OktaUserGroupProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OktaUserGroupProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OktaUserGroupProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OktaUserGroupProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectClass

`func (o *OktaUserGroupProfile) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *OktaUserGroupProfile) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *OktaUserGroupProfile) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *OktaUserGroupProfile) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



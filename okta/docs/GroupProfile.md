# GroupProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Windows group | [optional] 
**Name** | Pointer to **string** | Name of the Windows group | [optional] 
**ObjectClass** | Pointer to **string** | The object class type | [optional] [readonly] 
**Dn** | Pointer to **string** | The distinguished name of the Windows group | [optional] 
**ExternalId** | Pointer to **string** | Base-64 encoded GUID (&#x60;objectGUID&#x60;) of the Windows group | [optional] 
**SamAccountName** | Pointer to **string** | Pre-Windows 2000 name of the Windows group | [optional] 
**WindowsDomainQualifiedName** | Pointer to **string** | Fully qualified name of the Windows group | [optional] 

## Methods

### NewGroupProfile

`func NewGroupProfile() *GroupProfile`

NewGroupProfile instantiates a new GroupProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupProfileWithDefaults

`func NewGroupProfileWithDefaults() *GroupProfile`

NewGroupProfileWithDefaults instantiates a new GroupProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GroupProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectClass

`func (o *GroupProfile) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *GroupProfile) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *GroupProfile) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *GroupProfile) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetDn

`func (o *GroupProfile) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *GroupProfile) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *GroupProfile) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *GroupProfile) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetExternalId

`func (o *GroupProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GroupProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GroupProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GroupProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSamAccountName

`func (o *GroupProfile) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *GroupProfile) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *GroupProfile) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *GroupProfile) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetWindowsDomainQualifiedName

`func (o *GroupProfile) GetWindowsDomainQualifiedName() string`

GetWindowsDomainQualifiedName returns the WindowsDomainQualifiedName field if non-nil, zero value otherwise.

### GetWindowsDomainQualifiedNameOk

`func (o *GroupProfile) GetWindowsDomainQualifiedNameOk() (*string, bool)`

GetWindowsDomainQualifiedNameOk returns a tuple with the WindowsDomainQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainQualifiedName

`func (o *GroupProfile) SetWindowsDomainQualifiedName(v string)`

SetWindowsDomainQualifiedName sets WindowsDomainQualifiedName field to given value.

### HasWindowsDomainQualifiedName

`func (o *GroupProfile) HasWindowsDomainQualifiedName() bool`

HasWindowsDomainQualifiedName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



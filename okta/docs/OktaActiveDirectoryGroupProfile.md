# OktaActiveDirectoryGroupProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Windows group | [optional] 
**Dn** | Pointer to **string** | The distinguished name of the Windows group | [optional] 
**ExternalId** | Pointer to **string** | Base-64 encoded GUID (&#x60;objectGUID&#x60;) of the Windows group | [optional] 
**Name** | Pointer to **string** | Name of the Windows group | [optional] 
**ObjectClass** | Pointer to **string** | The object class type | [optional] [readonly] 
**SamAccountName** | Pointer to **string** | Pre-Windows 2000 name of the Windows group | [optional] 
**WindowsDomainQualifiedName** | Pointer to **string** | Fully qualified name of the Windows group | [optional] 

## Methods

### NewOktaActiveDirectoryGroupProfile

`func NewOktaActiveDirectoryGroupProfile() *OktaActiveDirectoryGroupProfile`

NewOktaActiveDirectoryGroupProfile instantiates a new OktaActiveDirectoryGroupProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaActiveDirectoryGroupProfileWithDefaults

`func NewOktaActiveDirectoryGroupProfileWithDefaults() *OktaActiveDirectoryGroupProfile`

NewOktaActiveDirectoryGroupProfileWithDefaults instantiates a new OktaActiveDirectoryGroupProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OktaActiveDirectoryGroupProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OktaActiveDirectoryGroupProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OktaActiveDirectoryGroupProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OktaActiveDirectoryGroupProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDn

`func (o *OktaActiveDirectoryGroupProfile) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *OktaActiveDirectoryGroupProfile) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *OktaActiveDirectoryGroupProfile) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *OktaActiveDirectoryGroupProfile) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetExternalId

`func (o *OktaActiveDirectoryGroupProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OktaActiveDirectoryGroupProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OktaActiveDirectoryGroupProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OktaActiveDirectoryGroupProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *OktaActiveDirectoryGroupProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OktaActiveDirectoryGroupProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OktaActiveDirectoryGroupProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OktaActiveDirectoryGroupProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectClass

`func (o *OktaActiveDirectoryGroupProfile) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *OktaActiveDirectoryGroupProfile) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *OktaActiveDirectoryGroupProfile) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *OktaActiveDirectoryGroupProfile) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetSamAccountName

`func (o *OktaActiveDirectoryGroupProfile) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *OktaActiveDirectoryGroupProfile) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *OktaActiveDirectoryGroupProfile) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *OktaActiveDirectoryGroupProfile) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### GetWindowsDomainQualifiedName

`func (o *OktaActiveDirectoryGroupProfile) GetWindowsDomainQualifiedName() string`

GetWindowsDomainQualifiedName returns the WindowsDomainQualifiedName field if non-nil, zero value otherwise.

### GetWindowsDomainQualifiedNameOk

`func (o *OktaActiveDirectoryGroupProfile) GetWindowsDomainQualifiedNameOk() (*string, bool)`

GetWindowsDomainQualifiedNameOk returns a tuple with the WindowsDomainQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomainQualifiedName

`func (o *OktaActiveDirectoryGroupProfile) SetWindowsDomainQualifiedName(v string)`

SetWindowsDomainQualifiedName sets WindowsDomainQualifiedName field to given value.

### HasWindowsDomainQualifiedName

`func (o *OktaActiveDirectoryGroupProfile) HasWindowsDomainQualifiedName() bool`

HasWindowsDomainQualifiedName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



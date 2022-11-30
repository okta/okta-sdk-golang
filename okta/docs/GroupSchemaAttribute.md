# GroupSchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enum** | Pointer to **[]string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**ExternalNamespace** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**UserSchemaAttributeItems**](UserSchemaAttributeItems.md) |  | [optional] 
**Master** | Pointer to [**UserSchemaAttributeMaster**](UserSchemaAttributeMaster.md) |  | [optional] 
**MaxLength** | Pointer to **int32** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**Mutability** | Pointer to **string** |  | [optional] 
**OneOf** | Pointer to [**[]UserSchemaAttributeEnum**](UserSchemaAttributeEnum.md) |  | [optional] 
**Permissions** | Pointer to [**[]UserSchemaAttributePermission**](UserSchemaAttributePermission.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to [**UserSchemaAttributeScope**](UserSchemaAttributeScope.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UserSchemaAttributeType**](UserSchemaAttributeType.md) |  | [optional] 
**Union** | Pointer to [**UserSchemaAttributeUnion**](UserSchemaAttributeUnion.md) |  | [optional] 
**Unique** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupSchemaAttribute

`func NewGroupSchemaAttribute() *GroupSchemaAttribute`

NewGroupSchemaAttribute instantiates a new GroupSchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSchemaAttributeWithDefaults

`func NewGroupSchemaAttributeWithDefaults() *GroupSchemaAttribute`

NewGroupSchemaAttributeWithDefaults instantiates a new GroupSchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupSchemaAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupSchemaAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupSchemaAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupSchemaAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnum

`func (o *GroupSchemaAttribute) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *GroupSchemaAttribute) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *GroupSchemaAttribute) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *GroupSchemaAttribute) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetExternalName

`func (o *GroupSchemaAttribute) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GroupSchemaAttribute) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GroupSchemaAttribute) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GroupSchemaAttribute) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetExternalNamespace

`func (o *GroupSchemaAttribute) GetExternalNamespace() string`

GetExternalNamespace returns the ExternalNamespace field if non-nil, zero value otherwise.

### GetExternalNamespaceOk

`func (o *GroupSchemaAttribute) GetExternalNamespaceOk() (*string, bool)`

GetExternalNamespaceOk returns a tuple with the ExternalNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNamespace

`func (o *GroupSchemaAttribute) SetExternalNamespace(v string)`

SetExternalNamespace sets ExternalNamespace field to given value.

### HasExternalNamespace

`func (o *GroupSchemaAttribute) HasExternalNamespace() bool`

HasExternalNamespace returns a boolean if a field has been set.

### GetItems

`func (o *GroupSchemaAttribute) GetItems() UserSchemaAttributeItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupSchemaAttribute) GetItemsOk() (*UserSchemaAttributeItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupSchemaAttribute) SetItems(v UserSchemaAttributeItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupSchemaAttribute) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaster

`func (o *GroupSchemaAttribute) GetMaster() UserSchemaAttributeMaster`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *GroupSchemaAttribute) GetMasterOk() (*UserSchemaAttributeMaster, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *GroupSchemaAttribute) SetMaster(v UserSchemaAttributeMaster)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *GroupSchemaAttribute) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetMaxLength

`func (o *GroupSchemaAttribute) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *GroupSchemaAttribute) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *GroupSchemaAttribute) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *GroupSchemaAttribute) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *GroupSchemaAttribute) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *GroupSchemaAttribute) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *GroupSchemaAttribute) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *GroupSchemaAttribute) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMutability

`func (o *GroupSchemaAttribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *GroupSchemaAttribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *GroupSchemaAttribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *GroupSchemaAttribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetOneOf

`func (o *GroupSchemaAttribute) GetOneOf() []UserSchemaAttributeEnum`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *GroupSchemaAttribute) GetOneOfOk() (*[]UserSchemaAttributeEnum, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *GroupSchemaAttribute) SetOneOf(v []UserSchemaAttributeEnum)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *GroupSchemaAttribute) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### GetPermissions

`func (o *GroupSchemaAttribute) GetPermissions() []UserSchemaAttributePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupSchemaAttribute) GetPermissionsOk() (*[]UserSchemaAttributePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupSchemaAttribute) SetPermissions(v []UserSchemaAttributePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupSchemaAttribute) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRequired

`func (o *GroupSchemaAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GroupSchemaAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GroupSchemaAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GroupSchemaAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetScope

`func (o *GroupSchemaAttribute) GetScope() UserSchemaAttributeScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GroupSchemaAttribute) GetScopeOk() (*UserSchemaAttributeScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GroupSchemaAttribute) SetScope(v UserSchemaAttributeScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GroupSchemaAttribute) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTitle

`func (o *GroupSchemaAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupSchemaAttribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupSchemaAttribute) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupSchemaAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *GroupSchemaAttribute) GetType() UserSchemaAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupSchemaAttribute) GetTypeOk() (*UserSchemaAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupSchemaAttribute) SetType(v UserSchemaAttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *GroupSchemaAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnion

`func (o *GroupSchemaAttribute) GetUnion() UserSchemaAttributeUnion`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *GroupSchemaAttribute) GetUnionOk() (*UserSchemaAttributeUnion, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *GroupSchemaAttribute) SetUnion(v UserSchemaAttributeUnion)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *GroupSchemaAttribute) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetUnique

`func (o *GroupSchemaAttribute) GetUnique() string`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *GroupSchemaAttribute) GetUniqueOk() (*string, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *GroupSchemaAttribute) SetUnique(v string)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *GroupSchemaAttribute) HasUnique() bool`

HasUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



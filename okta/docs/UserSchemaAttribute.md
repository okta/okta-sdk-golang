# UserSchemaAttribute

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
**Pattern** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]UserSchemaAttributePermission**](UserSchemaAttributePermission.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to [**UserSchemaAttributeScope**](UserSchemaAttributeScope.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UserSchemaAttributeType**](UserSchemaAttributeType.md) |  | [optional] 
**Union** | Pointer to [**UserSchemaAttributeUnion**](UserSchemaAttributeUnion.md) |  | [optional] 
**Unique** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSchemaAttribute

`func NewUserSchemaAttribute() *UserSchemaAttribute`

NewUserSchemaAttribute instantiates a new UserSchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaAttributeWithDefaults

`func NewUserSchemaAttributeWithDefaults() *UserSchemaAttribute`

NewUserSchemaAttributeWithDefaults instantiates a new UserSchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserSchemaAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSchemaAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSchemaAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserSchemaAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnum

`func (o *UserSchemaAttribute) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *UserSchemaAttribute) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *UserSchemaAttribute) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *UserSchemaAttribute) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetExternalName

`func (o *UserSchemaAttribute) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *UserSchemaAttribute) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *UserSchemaAttribute) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *UserSchemaAttribute) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetExternalNamespace

`func (o *UserSchemaAttribute) GetExternalNamespace() string`

GetExternalNamespace returns the ExternalNamespace field if non-nil, zero value otherwise.

### GetExternalNamespaceOk

`func (o *UserSchemaAttribute) GetExternalNamespaceOk() (*string, bool)`

GetExternalNamespaceOk returns a tuple with the ExternalNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNamespace

`func (o *UserSchemaAttribute) SetExternalNamespace(v string)`

SetExternalNamespace sets ExternalNamespace field to given value.

### HasExternalNamespace

`func (o *UserSchemaAttribute) HasExternalNamespace() bool`

HasExternalNamespace returns a boolean if a field has been set.

### GetItems

`func (o *UserSchemaAttribute) GetItems() UserSchemaAttributeItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSchemaAttribute) GetItemsOk() (*UserSchemaAttributeItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSchemaAttribute) SetItems(v UserSchemaAttributeItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserSchemaAttribute) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaster

`func (o *UserSchemaAttribute) GetMaster() UserSchemaAttributeMaster`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *UserSchemaAttribute) GetMasterOk() (*UserSchemaAttributeMaster, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *UserSchemaAttribute) SetMaster(v UserSchemaAttributeMaster)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *UserSchemaAttribute) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetMaxLength

`func (o *UserSchemaAttribute) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *UserSchemaAttribute) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *UserSchemaAttribute) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *UserSchemaAttribute) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *UserSchemaAttribute) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *UserSchemaAttribute) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *UserSchemaAttribute) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *UserSchemaAttribute) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMutability

`func (o *UserSchemaAttribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *UserSchemaAttribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *UserSchemaAttribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *UserSchemaAttribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetOneOf

`func (o *UserSchemaAttribute) GetOneOf() []UserSchemaAttributeEnum`

GetOneOf returns the OneOf field if non-nil, zero value otherwise.

### GetOneOfOk

`func (o *UserSchemaAttribute) GetOneOfOk() (*[]UserSchemaAttributeEnum, bool)`

GetOneOfOk returns a tuple with the OneOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneOf

`func (o *UserSchemaAttribute) SetOneOf(v []UserSchemaAttributeEnum)`

SetOneOf sets OneOf field to given value.

### HasOneOf

`func (o *UserSchemaAttribute) HasOneOf() bool`

HasOneOf returns a boolean if a field has been set.

### GetPattern

`func (o *UserSchemaAttribute) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *UserSchemaAttribute) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *UserSchemaAttribute) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *UserSchemaAttribute) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPermissions

`func (o *UserSchemaAttribute) GetPermissions() []UserSchemaAttributePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserSchemaAttribute) GetPermissionsOk() (*[]UserSchemaAttributePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserSchemaAttribute) SetPermissions(v []UserSchemaAttributePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserSchemaAttribute) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRequired

`func (o *UserSchemaAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UserSchemaAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UserSchemaAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UserSchemaAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetScope

`func (o *UserSchemaAttribute) GetScope() UserSchemaAttributeScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UserSchemaAttribute) GetScopeOk() (*UserSchemaAttributeScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UserSchemaAttribute) SetScope(v UserSchemaAttributeScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UserSchemaAttribute) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTitle

`func (o *UserSchemaAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserSchemaAttribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserSchemaAttribute) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserSchemaAttribute) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *UserSchemaAttribute) GetType() UserSchemaAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSchemaAttribute) GetTypeOk() (*UserSchemaAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSchemaAttribute) SetType(v UserSchemaAttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *UserSchemaAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnion

`func (o *UserSchemaAttribute) GetUnion() UserSchemaAttributeUnion`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *UserSchemaAttribute) GetUnionOk() (*UserSchemaAttributeUnion, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *UserSchemaAttribute) SetUnion(v UserSchemaAttributeUnion)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *UserSchemaAttribute) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetUnique

`func (o *UserSchemaAttribute) GetUnique() string`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *UserSchemaAttribute) GetUniqueOk() (*string, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *UserSchemaAttribute) SetUnique(v string)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *UserSchemaAttribute) HasUnique() bool`

HasUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



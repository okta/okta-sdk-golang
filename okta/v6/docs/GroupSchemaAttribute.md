# GroupSchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the property | [optional] 
**Enum** | Pointer to [**[]GroupSchemaAttributeEnumInner**](GroupSchemaAttributeEnumInner.md) | Enumerated value of the property.  The value of the property is limited to one of the values specified in the enum definition. The list of values for the enum must consist of unique elements. | [optional] 
**ExternalName** | Pointer to **string** | Name of the property as it exists in an external application | [optional] 
**ExternalNamespace** | Pointer to **string** | Namespace from the external application | [optional] 
**Format** | Pointer to **string** | Identifies the type of data represented by the string | [optional] 
**Items** | Pointer to [**UserSchemaAttributeItems**](UserSchemaAttributeItems.md) |  | [optional] 
**Master** | Pointer to [**NullableUserSchemaAttributeMaster**](UserSchemaAttributeMaster.md) | Identifies where the property is mastered | [optional] 
**MaxLength** | Pointer to **NullableInt32** | Maximum character length of a string property | [optional] 
**MinLength** | Pointer to **NullableInt32** | Minimum character length of a string property | [optional] 
**Mutability** | Pointer to **string** | Defines the mutability of the property | [optional] 
**OneOf** | Pointer to [**[]UserSchemaAttributeEnum**](UserSchemaAttributeEnum.md) | Non-empty array of valid JSON schemas.  The &#x60;oneOf&#x60; key is only supported in conjunction with &#x60;enum&#x60; and provides a mechanism to return a display name for the &#x60;enum&#x60; value.&lt;br&gt; Each schema has the following format:  &#x60;&#x60;&#x60; {   \&quot;const\&quot;: \&quot;enumValue\&quot;,   \&quot;title\&quot;: \&quot;display name\&quot; } &#x60;&#x60;&#x60;  When &#x60;enum&#x60; is used in conjunction with &#x60;oneOf&#x60;, you must keep the set of enumerated values and their order.&lt;br&gt; For example:  &#x60;&#x60;&#x60; \&quot;enum\&quot;: [\&quot;S\&quot;,\&quot;M\&quot;,\&quot;L\&quot;,\&quot;XL\&quot;], \&quot;oneOf\&quot;: [     {\&quot;const\&quot;: \&quot;S\&quot;, \&quot;title\&quot;: \&quot;Small\&quot;},     {\&quot;const\&quot;: \&quot;M\&quot;, \&quot;title\&quot;: \&quot;Medium\&quot;},     {\&quot;const\&quot;: \&quot;L\&quot;, \&quot;title\&quot;: \&quot;Large\&quot;},     {\&quot;const\&quot;: \&quot;XL\&quot;, \&quot;title\&quot;: \&quot;Extra Large\&quot;}   ] &#x60;&#x60;&#x60; | [optional] 
**Permissions** | Pointer to [**[]UserSchemaAttributePermission**](UserSchemaAttributePermission.md) | Access control permissions for the property | [optional] 
**Required** | Pointer to **NullableBool** | Determines whether the property is required | [optional] 
**Scope** | Pointer to **string** | Determines whether a group attribute can be set at the individual or group level | [optional] 
**Title** | Pointer to **string** | User-defined display name for the property | [optional] 
**Type** | Pointer to **string** | Type of property | [optional] 
**Unique** | Pointer to **NullableBool** | Determines whether property values must be unique | [optional] 

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

`func (o *GroupSchemaAttribute) GetEnum() []GroupSchemaAttributeEnumInner`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *GroupSchemaAttribute) GetEnumOk() (*[]GroupSchemaAttributeEnumInner, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *GroupSchemaAttribute) SetEnum(v []GroupSchemaAttributeEnumInner)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *GroupSchemaAttribute) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *GroupSchemaAttribute) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *GroupSchemaAttribute) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
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

### GetFormat

`func (o *GroupSchemaAttribute) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GroupSchemaAttribute) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GroupSchemaAttribute) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GroupSchemaAttribute) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

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

### SetMasterNil

`func (o *GroupSchemaAttribute) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *GroupSchemaAttribute) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
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

### SetMaxLengthNil

`func (o *GroupSchemaAttribute) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *GroupSchemaAttribute) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
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

### SetMinLengthNil

`func (o *GroupSchemaAttribute) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *GroupSchemaAttribute) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
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

### SetOneOfNil

`func (o *GroupSchemaAttribute) SetOneOfNil(b bool)`

 SetOneOfNil sets the value for OneOf to be an explicit nil

### UnsetOneOf
`func (o *GroupSchemaAttribute) UnsetOneOf()`

UnsetOneOf ensures that no value is present for OneOf, not even an explicit nil
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

### SetPermissionsNil

`func (o *GroupSchemaAttribute) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GroupSchemaAttribute) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
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

### SetRequiredNil

`func (o *GroupSchemaAttribute) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *GroupSchemaAttribute) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetScope

`func (o *GroupSchemaAttribute) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GroupSchemaAttribute) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GroupSchemaAttribute) SetScope(v string)`

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

`func (o *GroupSchemaAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupSchemaAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupSchemaAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupSchemaAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *GroupSchemaAttribute) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *GroupSchemaAttribute) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *GroupSchemaAttribute) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *GroupSchemaAttribute) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### SetUniqueNil

`func (o *GroupSchemaAttribute) SetUniqueNil(b bool)`

 SetUniqueNil sets the value for Unique to be an explicit nil

### UnsetUnique
`func (o *GroupSchemaAttribute) UnsetUnique()`

UnsetUnique ensures that no value is present for Unique, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



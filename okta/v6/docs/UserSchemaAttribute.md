# UserSchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **interface{}** | If specified, assigns the value as the default value for the custom attribute. This is a nullable property. If you don&#39;t specify a value for this custom attribute during user creation or update, the &#x60;default&#x60; is used instead of setting the value to &#x60;null&#x60; or empty. | [optional] 
**Description** | Pointer to **string** | Description of the property | [optional] 
**Enum** | Pointer to [**[]GroupSchemaAttributeEnumInner**](GroupSchemaAttributeEnumInner.md) | Enumerated value of the property.  The value of the property is limited to one of the values specified in the enum definition. The list of values for the enum must consist of unique elements. | [optional] 
**ExternalName** | Pointer to **string** | Name of the property as it exists in an external application  **NOTE**: When you add a custom property, only Identity Provider app user schemas require &#x60;externalName&#x60; to be included in the request body. If an existing custom Identity Provider app user schema property has an empty &#x60;externalName&#x60;, requests aren&#39;t allowed to update other properties until the &#x60;externalName&#x60; is defined. | [optional] 
**ExternalNamespace** | Pointer to **string** | Namespace from the external application | [optional] 
**Format** | Pointer to **string** | Identifies the type of data represented by the string | [optional] 
**Master** | Pointer to [**NullableUserSchemaAttributeMaster**](UserSchemaAttributeMaster.md) | Identifies where the property is mastered | [optional] 
**MaxLength** | Pointer to **NullableInt32** | Maximum character length of a string property | [optional] 
**MinLength** | Pointer to **NullableInt32** | Minimum character length of a string property | [optional] 
**Mutability** | Pointer to **string** | Defines the mutability of the property | [optional] 
**OneOf** | Pointer to [**[]UserSchemaAttributeEnum**](UserSchemaAttributeEnum.md) | Non-empty array of valid JSON schemas.  The &#x60;oneOf&#x60; key is only supported in conjunction with &#x60;enum&#x60; and provides a mechanism to return a display name for the &#x60;enum&#x60; value.&lt;br&gt; Each schema has the following format:  &#x60;&#x60;&#x60; {   \&quot;const\&quot;: \&quot;enumValue\&quot;,   \&quot;title\&quot;: \&quot;display name\&quot; } &#x60;&#x60;&#x60;  When &#x60;enum&#x60; is used in conjunction with &#x60;oneOf&#x60;, you must keep the set of enumerated values and their order.&lt;br&gt; For example:  &#x60;&#x60;&#x60; \&quot;enum\&quot;: [\&quot;S\&quot;,\&quot;M\&quot;,\&quot;L\&quot;,\&quot;XL\&quot;], \&quot;oneOf\&quot;: [     {\&quot;const\&quot;: \&quot;S\&quot;, \&quot;title\&quot;: \&quot;Small\&quot;},     {\&quot;const\&quot;: \&quot;M\&quot;, \&quot;title\&quot;: \&quot;Medium\&quot;},     {\&quot;const\&quot;: \&quot;L\&quot;, \&quot;title\&quot;: \&quot;Large\&quot;},     {\&quot;const\&quot;: \&quot;XL\&quot;, \&quot;title\&quot;: \&quot;Extra Large\&quot;}   ] &#x60;&#x60;&#x60; | [optional] 
**Pattern** | Pointer to **string** | For &#x60;string&#x60; property types, specifies the regular expression used to validate the property | [optional] 
**Permissions** | Pointer to [**[]UserSchemaAttributePermission**](UserSchemaAttributePermission.md) | Access control permissions for the property | [optional] 
**Required** | Pointer to **NullableBool** | Determines whether the property is required | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | User-defined display name for the property | [optional] 
**Type** | Pointer to **string** | Type of property | [optional] 
**Unique** | Pointer to **NullableBool** | Determines whether property values must be unique | [optional] 

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

### GetDefault

`func (o *UserSchemaAttribute) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UserSchemaAttribute) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UserSchemaAttribute) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UserSchemaAttribute) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *UserSchemaAttribute) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *UserSchemaAttribute) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
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

`func (o *UserSchemaAttribute) GetEnum() []GroupSchemaAttributeEnumInner`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *UserSchemaAttribute) GetEnumOk() (*[]GroupSchemaAttributeEnumInner, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *UserSchemaAttribute) SetEnum(v []GroupSchemaAttributeEnumInner)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *UserSchemaAttribute) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### SetEnumNil

`func (o *UserSchemaAttribute) SetEnumNil(b bool)`

 SetEnumNil sets the value for Enum to be an explicit nil

### UnsetEnum
`func (o *UserSchemaAttribute) UnsetEnum()`

UnsetEnum ensures that no value is present for Enum, not even an explicit nil
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

### GetFormat

`func (o *UserSchemaAttribute) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UserSchemaAttribute) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UserSchemaAttribute) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UserSchemaAttribute) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

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

### SetMasterNil

`func (o *UserSchemaAttribute) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *UserSchemaAttribute) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
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

### SetMaxLengthNil

`func (o *UserSchemaAttribute) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *UserSchemaAttribute) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
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

### SetMinLengthNil

`func (o *UserSchemaAttribute) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *UserSchemaAttribute) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
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

### SetOneOfNil

`func (o *UserSchemaAttribute) SetOneOfNil(b bool)`

 SetOneOfNil sets the value for OneOf to be an explicit nil

### UnsetOneOf
`func (o *UserSchemaAttribute) UnsetOneOf()`

UnsetOneOf ensures that no value is present for OneOf, not even an explicit nil
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

### SetPermissionsNil

`func (o *UserSchemaAttribute) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UserSchemaAttribute) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
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

### SetRequiredNil

`func (o *UserSchemaAttribute) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *UserSchemaAttribute) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetScope

`func (o *UserSchemaAttribute) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UserSchemaAttribute) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UserSchemaAttribute) SetScope(v string)`

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

`func (o *UserSchemaAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSchemaAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSchemaAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSchemaAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *UserSchemaAttribute) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *UserSchemaAttribute) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *UserSchemaAttribute) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *UserSchemaAttribute) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### SetUniqueNil

`func (o *UserSchemaAttribute) SetUniqueNil(b bool)`

 SetUniqueNil sets the value for Unique to be an explicit nil

### UnsetUnique
`func (o *UserSchemaAttribute) UnsetUnique()`

UnsetUnique ensures that no value is present for Unique, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



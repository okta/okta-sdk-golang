/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupSchemaAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSchemaAttribute{}

// GroupSchemaAttribute struct for GroupSchemaAttribute
type GroupSchemaAttribute struct {
	// Description of the property
	Description *string `json:"description,omitempty"`
	// Enumerated value of the property.  The value of the property is limited to one of the values specified in the enum definition. The list of values for the enum must consist of unique elements.
	Enum []GroupSchemaAttributeEnumInner `json:"enum,omitempty"`
	// Name of the property as it exists in an external application
	ExternalName *string `json:"externalName,omitempty"`
	// Namespace from the external application
	ExternalNamespace *string `json:"externalNamespace,omitempty"`
	// Identifies the type of data represented by the string
	Format *string                   `json:"format,omitempty"`
	Items  *UserSchemaAttributeItems `json:"items,omitempty"`
	// Identifies where the property is mastered
	Master NullableUserSchemaAttributeMaster `json:"master,omitempty"`
	// Maximum character length of a string property
	MaxLength NullableInt32 `json:"maxLength,omitempty"`
	// Minimum character length of a string property
	MinLength NullableInt32 `json:"minLength,omitempty"`
	// Defines the mutability of the property
	Mutability *string `json:"mutability,omitempty"`
	// Non-empty array of valid JSON schemas.  The `oneOf` key is only supported in conjunction with `enum` and provides a mechanism to return a display name for the `enum` value.<br> Each schema has the following format:  ``` {   \"const\": \"enumValue\",   \"title\": \"display name\" } ```  When `enum` is used in conjunction with `oneOf`, you must keep the set of enumerated values and their order.<br> For example:  ``` \"enum\": [\"S\",\"M\",\"L\",\"XL\"], \"oneOf\": [     {\"const\": \"S\", \"title\": \"Small\"},     {\"const\": \"M\", \"title\": \"Medium\"},     {\"const\": \"L\", \"title\": \"Large\"},     {\"const\": \"XL\", \"title\": \"Extra Large\"}   ] ```
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	// Access control permissions for the property
	Permissions []UserSchemaAttributePermission `json:"permissions,omitempty"`
	// Determines whether the property is required
	Required NullableBool `json:"required,omitempty"`
	// Determines whether a group attribute can be set at the individual or group level
	Scope *string `json:"scope,omitempty"`
	// User-defined display name for the property
	Title *string `json:"title,omitempty"`
	// Type of property
	Type *string `json:"type,omitempty"`
	// Determines whether property values must be unique
	Unique               NullableBool `json:"unique,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchemaAttribute GroupSchemaAttribute

// NewGroupSchemaAttribute instantiates a new GroupSchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchemaAttribute() *GroupSchemaAttribute {
	this := GroupSchemaAttribute{}
	return &this
}

// NewGroupSchemaAttributeWithDefaults instantiates a new GroupSchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaAttributeWithDefaults() *GroupSchemaAttribute {
	this := GroupSchemaAttribute{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupSchemaAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetEnum() []GroupSchemaAttributeEnumInner {
	if o == nil {
		var ret []GroupSchemaAttributeEnumInner
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetEnumOk() ([]GroupSchemaAttributeEnumInner, bool) {
	if o == nil || IsNil(o.Enum) {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasEnum() bool {
	if o != nil && !IsNil(o.Enum) {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []GroupSchemaAttributeEnumInner and assigns it to the Enum field.
func (o *GroupSchemaAttribute) SetEnum(v []GroupSchemaAttributeEnumInner) {
	o.Enum = v
}

// GetExternalName returns the ExternalName field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetExternalName() string {
	if o == nil || IsNil(o.ExternalName) {
		var ret string
		return ret
	}
	return *o.ExternalName
}

// GetExternalNameOk returns a tuple with the ExternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetExternalNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalName) {
		return nil, false
	}
	return o.ExternalName, true
}

// HasExternalName returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasExternalName() bool {
	if o != nil && !IsNil(o.ExternalName) {
		return true
	}

	return false
}

// SetExternalName gets a reference to the given string and assigns it to the ExternalName field.
func (o *GroupSchemaAttribute) SetExternalName(v string) {
	o.ExternalName = &v
}

// GetExternalNamespace returns the ExternalNamespace field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetExternalNamespace() string {
	if o == nil || IsNil(o.ExternalNamespace) {
		var ret string
		return ret
	}
	return *o.ExternalNamespace
}

// GetExternalNamespaceOk returns a tuple with the ExternalNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetExternalNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalNamespace) {
		return nil, false
	}
	return o.ExternalNamespace, true
}

// HasExternalNamespace returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasExternalNamespace() bool {
	if o != nil && !IsNil(o.ExternalNamespace) {
		return true
	}

	return false
}

// SetExternalNamespace gets a reference to the given string and assigns it to the ExternalNamespace field.
func (o *GroupSchemaAttribute) SetExternalNamespace(v string) {
	o.ExternalNamespace = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *GroupSchemaAttribute) SetFormat(v string) {
	o.Format = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetItems() UserSchemaAttributeItems {
	if o == nil || IsNil(o.Items) {
		var ret UserSchemaAttributeItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetItemsOk() (*UserSchemaAttributeItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given UserSchemaAttributeItems and assigns it to the Items field.
func (o *GroupSchemaAttribute) SetItems(v UserSchemaAttributeItems) {
	o.Items = &v
}

// GetMaster returns the Master field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetMaster() UserSchemaAttributeMaster {
	if o == nil || IsNil(o.Master.Get()) {
		var ret UserSchemaAttributeMaster
		return ret
	}
	return *o.Master.Get()
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetMasterOk() (*UserSchemaAttributeMaster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Master.Get(), o.Master.IsSet()
}

// HasMaster returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMaster() bool {
	if o != nil && o.Master.IsSet() {
		return true
	}

	return false
}

// SetMaster gets a reference to the given NullableUserSchemaAttributeMaster and assigns it to the Master field.
func (o *GroupSchemaAttribute) SetMaster(v UserSchemaAttributeMaster) {
	o.Master.Set(&v)
}

// SetMasterNil sets the value for Master to be an explicit nil
func (o *GroupSchemaAttribute) SetMasterNil() {
	o.Master.Set(nil)
}

// UnsetMaster ensures that no value is present for Master, not even an explicit nil
func (o *GroupSchemaAttribute) UnsetMaster() {
	o.Master.Unset()
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetMaxLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *GroupSchemaAttribute) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}

// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *GroupSchemaAttribute) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *GroupSchemaAttribute) UnsetMaxLength() {
	o.MaxLength.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength.Get()) {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *GroupSchemaAttribute) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}

// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *GroupSchemaAttribute) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *GroupSchemaAttribute) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetMutability() string {
	if o == nil || IsNil(o.Mutability) {
		var ret string
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetMutabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Mutability) {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMutability() bool {
	if o != nil && !IsNil(o.Mutability) {
		return true
	}

	return false
}

// SetMutability gets a reference to the given string and assigns it to the Mutability field.
func (o *GroupSchemaAttribute) SetMutability(v string) {
	o.Mutability = &v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetOneOf() []UserSchemaAttributeEnum {
	if o == nil {
		var ret []UserSchemaAttributeEnum
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetOneOfOk() ([]UserSchemaAttributeEnum, bool) {
	if o == nil || IsNil(o.OneOf) {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasOneOf() bool {
	if o != nil && !IsNil(o.OneOf) {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []UserSchemaAttributeEnum and assigns it to the OneOf field.
func (o *GroupSchemaAttribute) SetOneOf(v []UserSchemaAttributeEnum) {
	o.OneOf = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetPermissions() []UserSchemaAttributePermission {
	if o == nil {
		var ret []UserSchemaAttributePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetPermissionsOk() ([]UserSchemaAttributePermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSchemaAttributePermission and assigns it to the Permissions field.
func (o *GroupSchemaAttribute) SetPermissions(v []UserSchemaAttributePermission) {
	o.Permissions = v
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *GroupSchemaAttribute) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil
func (o *GroupSchemaAttribute) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *GroupSchemaAttribute) UnsetRequired() {
	o.Required.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GroupSchemaAttribute) SetScope(v string) {
	o.Scope = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupSchemaAttribute) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupSchemaAttribute) SetType(v string) {
	o.Type = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupSchemaAttribute) GetUnique() bool {
	if o == nil || IsNil(o.Unique.Get()) {
		var ret bool
		return ret
	}
	return *o.Unique.Get()
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupSchemaAttribute) GetUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unique.Get(), o.Unique.IsSet()
}

// HasUnique returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasUnique() bool {
	if o != nil && o.Unique.IsSet() {
		return true
	}

	return false
}

// SetUnique gets a reference to the given NullableBool and assigns it to the Unique field.
func (o *GroupSchemaAttribute) SetUnique(v bool) {
	o.Unique.Set(&v)
}

// SetUniqueNil sets the value for Unique to be an explicit nil
func (o *GroupSchemaAttribute) SetUniqueNil() {
	o.Unique.Set(nil)
}

// UnsetUnique ensures that no value is present for Unique, not even an explicit nil
func (o *GroupSchemaAttribute) UnsetUnique() {
	o.Unique.Unset()
}

func (o GroupSchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSchemaAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if !IsNil(o.ExternalName) {
		toSerialize["externalName"] = o.ExternalName
	}
	if !IsNil(o.ExternalNamespace) {
		toSerialize["externalNamespace"] = o.ExternalNamespace
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if o.Master.IsSet() {
		toSerialize["master"] = o.Master.Get()
	}
	if o.MaxLength.IsSet() {
		toSerialize["maxLength"] = o.MaxLength.Get()
	}
	if o.MinLength.IsSet() {
		toSerialize["minLength"] = o.MinLength.Get()
	}
	if !IsNil(o.Mutability) {
		toSerialize["mutability"] = o.Mutability
	}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Unique.IsSet() {
		toSerialize["unique"] = o.Unique.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupSchemaAttribute) UnmarshalJSON(data []byte) (err error) {
	varGroupSchemaAttribute := _GroupSchemaAttribute{}

	err = json.Unmarshal(data, &varGroupSchemaAttribute)

	if err != nil {
		return err
	}

	*o = GroupSchemaAttribute(varGroupSchemaAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "enum")
		delete(additionalProperties, "externalName")
		delete(additionalProperties, "externalNamespace")
		delete(additionalProperties, "format")
		delete(additionalProperties, "items")
		delete(additionalProperties, "master")
		delete(additionalProperties, "maxLength")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "mutability")
		delete(additionalProperties, "oneOf")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "required")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "unique")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupSchemaAttribute struct {
	value *GroupSchemaAttribute
	isSet bool
}

func (v NullableGroupSchemaAttribute) Get() *GroupSchemaAttribute {
	return v.value
}

func (v *NullableGroupSchemaAttribute) Set(val *GroupSchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchemaAttribute(val *GroupSchemaAttribute) *NullableGroupSchemaAttribute {
	return &NullableGroupSchemaAttribute{value: val, isSet: true}
}

func (v NullableGroupSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

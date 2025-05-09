/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserSchemaAttribute struct for UserSchemaAttribute
type UserSchemaAttribute struct {
	Description *string `json:"description,omitempty"`
	Enum []string `json:"enum,omitempty"`
	ExternalName *string `json:"externalName,omitempty"`
	ExternalNamespace *string `json:"externalNamespace,omitempty"`
	Items *UserSchemaAttributeItems `json:"items,omitempty"`
	Master *UserSchemaAttributeMaster `json:"master,omitempty"`
	MaxLength NullableInt32 `json:"maxLength,omitempty"`
	MinLength NullableInt32 `json:"minLength,omitempty"`
	Mutability *string `json:"mutability,omitempty"`
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	Permissions []UserSchemaAttributePermission `json:"permissions,omitempty"`
	Required *bool `json:"required,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Union *string `json:"union,omitempty"`
	Unique *string `json:"unique,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaAttribute UserSchemaAttribute

// NewUserSchemaAttribute instantiates a new UserSchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaAttribute() *UserSchemaAttribute {
	this := UserSchemaAttribute{}
	return &this
}

// NewUserSchemaAttributeWithDefaults instantiates a new UserSchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaAttributeWithDefaults() *UserSchemaAttribute {
	this := UserSchemaAttribute{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserSchemaAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetEnumOk() ([]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *UserSchemaAttribute) SetEnum(v []string) {
	o.Enum = v
}

// GetExternalName returns the ExternalName field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetExternalName() string {
	if o == nil || o.ExternalName == nil {
		var ret string
		return ret
	}
	return *o.ExternalName
}

// GetExternalNameOk returns a tuple with the ExternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetExternalNameOk() (*string, bool) {
	if o == nil || o.ExternalName == nil {
		return nil, false
	}
	return o.ExternalName, true
}

// HasExternalName returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasExternalName() bool {
	if o != nil && o.ExternalName != nil {
		return true
	}

	return false
}

// SetExternalName gets a reference to the given string and assigns it to the ExternalName field.
func (o *UserSchemaAttribute) SetExternalName(v string) {
	o.ExternalName = &v
}

// GetExternalNamespace returns the ExternalNamespace field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetExternalNamespace() string {
	if o == nil || o.ExternalNamespace == nil {
		var ret string
		return ret
	}
	return *o.ExternalNamespace
}

// GetExternalNamespaceOk returns a tuple with the ExternalNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetExternalNamespaceOk() (*string, bool) {
	if o == nil || o.ExternalNamespace == nil {
		return nil, false
	}
	return o.ExternalNamespace, true
}

// HasExternalNamespace returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasExternalNamespace() bool {
	if o != nil && o.ExternalNamespace != nil {
		return true
	}

	return false
}

// SetExternalNamespace gets a reference to the given string and assigns it to the ExternalNamespace field.
func (o *UserSchemaAttribute) SetExternalNamespace(v string) {
	o.ExternalNamespace = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetItems() UserSchemaAttributeItems {
	if o == nil || o.Items == nil {
		var ret UserSchemaAttributeItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetItemsOk() (*UserSchemaAttributeItems, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given UserSchemaAttributeItems and assigns it to the Items field.
func (o *UserSchemaAttribute) SetItems(v UserSchemaAttributeItems) {
	o.Items = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetMaster() UserSchemaAttributeMaster {
	if o == nil || o.Master == nil {
		var ret UserSchemaAttributeMaster
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetMasterOk() (*UserSchemaAttributeMaster, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given UserSchemaAttributeMaster and assigns it to the Master field.
func (o *UserSchemaAttribute) SetMaster(v UserSchemaAttributeMaster) {
	o.Master = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSchemaAttribute) GetMaxLength() int32 {
	if o == nil || o.MaxLength.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength.Get()
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSchemaAttribute) GetMaxLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLength.Get(), o.MaxLength.IsSet()
}

// HasMaxLength returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasMaxLength() bool {
	if o != nil && o.MaxLength.IsSet() {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given NullableInt32 and assigns it to the MaxLength field.
func (o *UserSchemaAttribute) SetMaxLength(v int32) {
	o.MaxLength.Set(&v)
}
// SetMaxLengthNil sets the value for MaxLength to be an explicit nil
func (o *UserSchemaAttribute) SetMaxLengthNil() {
	o.MaxLength.Set(nil)
}

// UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
func (o *UserSchemaAttribute) UnsetMaxLength() {
	o.MaxLength.Unset()
}

// GetMinLength returns the MinLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSchemaAttribute) GetMinLength() int32 {
	if o == nil || o.MinLength.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinLength.Get()
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSchemaAttribute) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinLength.Get(), o.MinLength.IsSet()
}

// HasMinLength returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasMinLength() bool {
	if o != nil && o.MinLength.IsSet() {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given NullableInt32 and assigns it to the MinLength field.
func (o *UserSchemaAttribute) SetMinLength(v int32) {
	o.MinLength.Set(&v)
}
// SetMinLengthNil sets the value for MinLength to be an explicit nil
func (o *UserSchemaAttribute) SetMinLengthNil() {
	o.MinLength.Set(nil)
}

// UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
func (o *UserSchemaAttribute) UnsetMinLength() {
	o.MinLength.Unset()
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetMutability() string {
	if o == nil || o.Mutability == nil {
		var ret string
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetMutabilityOk() (*string, bool) {
	if o == nil || o.Mutability == nil {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasMutability() bool {
	if o != nil && o.Mutability != nil {
		return true
	}

	return false
}

// SetMutability gets a reference to the given string and assigns it to the Mutability field.
func (o *UserSchemaAttribute) SetMutability(v string) {
	o.Mutability = &v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetOneOf() []UserSchemaAttributeEnum {
	if o == nil || o.OneOf == nil {
		var ret []UserSchemaAttributeEnum
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetOneOfOk() ([]UserSchemaAttributeEnum, bool) {
	if o == nil || o.OneOf == nil {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasOneOf() bool {
	if o != nil && o.OneOf != nil {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []UserSchemaAttributeEnum and assigns it to the OneOf field.
func (o *UserSchemaAttribute) SetOneOf(v []UserSchemaAttributeEnum) {
	o.OneOf = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *UserSchemaAttribute) SetPattern(v string) {
	o.Pattern = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetPermissions() []UserSchemaAttributePermission {
	if o == nil || o.Permissions == nil {
		var ret []UserSchemaAttributePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetPermissionsOk() ([]UserSchemaAttributePermission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSchemaAttributePermission and assigns it to the Permissions field.
func (o *UserSchemaAttribute) SetPermissions(v []UserSchemaAttributePermission) {
	o.Permissions = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UserSchemaAttribute) SetRequired(v bool) {
	o.Required = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UserSchemaAttribute) SetScope(v string) {
	o.Scope = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserSchemaAttribute) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSchemaAttribute) SetType(v string) {
	o.Type = &v
}

// GetUnion returns the Union field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetUnion() string {
	if o == nil || o.Union == nil {
		var ret string
		return ret
	}
	return *o.Union
}

// GetUnionOk returns a tuple with the Union field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetUnionOk() (*string, bool) {
	if o == nil || o.Union == nil {
		return nil, false
	}
	return o.Union, true
}

// HasUnion returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasUnion() bool {
	if o != nil && o.Union != nil {
		return true
	}

	return false
}

// SetUnion gets a reference to the given string and assigns it to the Union field.
func (o *UserSchemaAttribute) SetUnion(v string) {
	o.Union = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *UserSchemaAttribute) GetUnique() string {
	if o == nil || o.Unique == nil {
		var ret string
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttribute) GetUniqueOk() (*string, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *UserSchemaAttribute) HasUnique() bool {
	if o != nil && o.Unique != nil {
		return true
	}

	return false
}

// SetUnique gets a reference to the given string and assigns it to the Unique field.
func (o *UserSchemaAttribute) SetUnique(v string) {
	o.Unique = &v
}

func (o UserSchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.ExternalName != nil {
		toSerialize["externalName"] = o.ExternalName
	}
	if o.ExternalNamespace != nil {
		toSerialize["externalNamespace"] = o.ExternalNamespace
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if o.MaxLength.IsSet() {
		toSerialize["maxLength"] = o.MaxLength.Get()
	}
	if o.MinLength.IsSet() {
		toSerialize["minLength"] = o.MinLength.Get()
	}
	if o.Mutability != nil {
		toSerialize["mutability"] = o.Mutability
	}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Union != nil {
		toSerialize["union"] = o.Union
	}
	if o.Unique != nil {
		toSerialize["unique"] = o.Unique
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaAttribute := _UserSchemaAttribute{}

	err = json.Unmarshal(bytes, &varUserSchemaAttribute)
	if err == nil {
		*o = UserSchemaAttribute(varUserSchemaAttribute)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "enum")
		delete(additionalProperties, "externalName")
		delete(additionalProperties, "externalNamespace")
		delete(additionalProperties, "items")
		delete(additionalProperties, "master")
		delete(additionalProperties, "maxLength")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "mutability")
		delete(additionalProperties, "oneOf")
		delete(additionalProperties, "pattern")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "required")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "union")
		delete(additionalProperties, "unique")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaAttribute struct {
	value *UserSchemaAttribute
	isSet bool
}

func (v NullableUserSchemaAttribute) Get() *UserSchemaAttribute {
	return v.value
}

func (v *NullableUserSchemaAttribute) Set(val *UserSchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttribute(val *UserSchemaAttribute) *NullableUserSchemaAttribute {
	return &NullableUserSchemaAttribute{value: val, isSet: true}
}

func (v NullableUserSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


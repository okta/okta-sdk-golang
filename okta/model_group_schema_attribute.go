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

// GroupSchemaAttribute struct for GroupSchemaAttribute
type GroupSchemaAttribute struct {
	Description *string `json:"description,omitempty"`
	Enum []string `json:"enum,omitempty"`
	ExternalName *string `json:"externalName,omitempty"`
	ExternalNamespace *string `json:"externalNamespace,omitempty"`
	Items *UserSchemaAttributeItems `json:"items,omitempty"`
	Master *UserSchemaAttributeMaster `json:"master,omitempty"`
	MaxLength *int32 `json:"maxLength,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
	Mutability *string `json:"mutability,omitempty"`
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	Permissions []UserSchemaAttributePermission `json:"permissions,omitempty"`
	Required *bool `json:"required,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Union *string `json:"union,omitempty"`
	Unique *string `json:"unique,omitempty"`
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
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupSchemaAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetEnumOk() ([]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasEnum() bool {
	if o != nil && o.Enum != nil {
		return true
	}

	return false
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *GroupSchemaAttribute) SetEnum(v []string) {
	o.Enum = v
}

// GetExternalName returns the ExternalName field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetExternalName() string {
	if o == nil || o.ExternalName == nil {
		var ret string
		return ret
	}
	return *o.ExternalName
}

// GetExternalNameOk returns a tuple with the ExternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetExternalNameOk() (*string, bool) {
	if o == nil || o.ExternalName == nil {
		return nil, false
	}
	return o.ExternalName, true
}

// HasExternalName returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasExternalName() bool {
	if o != nil && o.ExternalName != nil {
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
	if o == nil || o.ExternalNamespace == nil {
		var ret string
		return ret
	}
	return *o.ExternalNamespace
}

// GetExternalNamespaceOk returns a tuple with the ExternalNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetExternalNamespaceOk() (*string, bool) {
	if o == nil || o.ExternalNamespace == nil {
		return nil, false
	}
	return o.ExternalNamespace, true
}

// HasExternalNamespace returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasExternalNamespace() bool {
	if o != nil && o.ExternalNamespace != nil {
		return true
	}

	return false
}

// SetExternalNamespace gets a reference to the given string and assigns it to the ExternalNamespace field.
func (o *GroupSchemaAttribute) SetExternalNamespace(v string) {
	o.ExternalNamespace = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetItems() UserSchemaAttributeItems {
	if o == nil || o.Items == nil {
		var ret UserSchemaAttributeItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetItemsOk() (*UserSchemaAttributeItems, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given UserSchemaAttributeItems and assigns it to the Items field.
func (o *GroupSchemaAttribute) SetItems(v UserSchemaAttributeItems) {
	o.Items = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetMaster() UserSchemaAttributeMaster {
	if o == nil || o.Master == nil {
		var ret UserSchemaAttributeMaster
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetMasterOk() (*UserSchemaAttributeMaster, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given UserSchemaAttributeMaster and assigns it to the Master field.
func (o *GroupSchemaAttribute) SetMaster(v UserSchemaAttributeMaster) {
	o.Master = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetMaxLength() int32 {
	if o == nil || o.MaxLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetMaxLengthOk() (*int32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *GroupSchemaAttribute) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *GroupSchemaAttribute) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetMutability() string {
	if o == nil || o.Mutability == nil {
		var ret string
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetMutabilityOk() (*string, bool) {
	if o == nil || o.Mutability == nil {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasMutability() bool {
	if o != nil && o.Mutability != nil {
		return true
	}

	return false
}

// SetMutability gets a reference to the given string and assigns it to the Mutability field.
func (o *GroupSchemaAttribute) SetMutability(v string) {
	o.Mutability = &v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetOneOf() []UserSchemaAttributeEnum {
	if o == nil || o.OneOf == nil {
		var ret []UserSchemaAttributeEnum
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetOneOfOk() ([]UserSchemaAttributeEnum, bool) {
	if o == nil || o.OneOf == nil {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasOneOf() bool {
	if o != nil && o.OneOf != nil {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []UserSchemaAttributeEnum and assigns it to the OneOf field.
func (o *GroupSchemaAttribute) SetOneOf(v []UserSchemaAttributeEnum) {
	o.OneOf = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetPermissions() []UserSchemaAttributePermission {
	if o == nil || o.Permissions == nil {
		var ret []UserSchemaAttributePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetPermissionsOk() ([]UserSchemaAttributePermission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []UserSchemaAttributePermission and assigns it to the Permissions field.
func (o *GroupSchemaAttribute) SetPermissions(v []UserSchemaAttributePermission) {
	o.Permissions = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *GroupSchemaAttribute) SetRequired(v bool) {
	o.Required = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasScope() bool {
	if o != nil && o.Scope != nil {
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
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasTitle() bool {
	if o != nil && o.Title != nil {
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
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupSchemaAttribute) SetType(v string) {
	o.Type = &v
}

// GetUnion returns the Union field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetUnion() string {
	if o == nil || o.Union == nil {
		var ret string
		return ret
	}
	return *o.Union
}

// GetUnionOk returns a tuple with the Union field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetUnionOk() (*string, bool) {
	if o == nil || o.Union == nil {
		return nil, false
	}
	return o.Union, true
}

// HasUnion returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasUnion() bool {
	if o != nil && o.Union != nil {
		return true
	}

	return false
}

// SetUnion gets a reference to the given string and assigns it to the Union field.
func (o *GroupSchemaAttribute) SetUnion(v string) {
	o.Union = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *GroupSchemaAttribute) GetUnique() string {
	if o == nil || o.Unique == nil {
		var ret string
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchemaAttribute) GetUniqueOk() (*string, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *GroupSchemaAttribute) HasUnique() bool {
	if o != nil && o.Unique != nil {
		return true
	}

	return false
}

// SetUnique gets a reference to the given string and assigns it to the Unique field.
func (o *GroupSchemaAttribute) SetUnique(v string) {
	o.Unique = &v
}

func (o GroupSchemaAttribute) MarshalJSON() ([]byte, error) {
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
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.Mutability != nil {
		toSerialize["mutability"] = o.Mutability
	}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
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

func (o *GroupSchemaAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varGroupSchemaAttribute := _GroupSchemaAttribute{}

	err = json.Unmarshal(bytes, &varGroupSchemaAttribute)
	if err == nil {
		*o = GroupSchemaAttribute(varGroupSchemaAttribute)
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


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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the EntitlementTypesInnerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementTypesInnerAttributes{}

// EntitlementTypesInnerAttributes Attributes for the entitlement type
type EntitlementTypesInnerAttributes struct {
	// A boolean value to indicate if this entitlement type is required for the user
	Required *bool `json:"required,omitempty"`
	// A boolean value to indicate if a user can have multiple entitlements of this type
	Multivalued          *bool `json:"multivalued,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementTypesInnerAttributes EntitlementTypesInnerAttributes

// NewEntitlementTypesInnerAttributes instantiates a new EntitlementTypesInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementTypesInnerAttributes() *EntitlementTypesInnerAttributes {
	this := EntitlementTypesInnerAttributes{}
	var required bool = false
	this.Required = &required
	var multivalued bool = false
	this.Multivalued = &multivalued
	return &this
}

// NewEntitlementTypesInnerAttributesWithDefaults instantiates a new EntitlementTypesInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementTypesInnerAttributesWithDefaults() *EntitlementTypesInnerAttributes {
	this := EntitlementTypesInnerAttributes{}
	var required bool = false
	this.Required = &required
	var multivalued bool = false
	this.Multivalued = &multivalued
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EntitlementTypesInnerAttributes) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInnerAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EntitlementTypesInnerAttributes) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EntitlementTypesInnerAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetMultivalued returns the Multivalued field value if set, zero value otherwise.
func (o *EntitlementTypesInnerAttributes) GetMultivalued() bool {
	if o == nil || IsNil(o.Multivalued) {
		var ret bool
		return ret
	}
	return *o.Multivalued
}

// GetMultivaluedOk returns a tuple with the Multivalued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInnerAttributes) GetMultivaluedOk() (*bool, bool) {
	if o == nil || IsNil(o.Multivalued) {
		return nil, false
	}
	return o.Multivalued, true
}

// HasMultivalued returns a boolean if a field has been set.
func (o *EntitlementTypesInnerAttributes) HasMultivalued() bool {
	if o != nil && !IsNil(o.Multivalued) {
		return true
	}

	return false
}

// SetMultivalued gets a reference to the given bool and assigns it to the Multivalued field.
func (o *EntitlementTypesInnerAttributes) SetMultivalued(v bool) {
	o.Multivalued = &v
}

func (o EntitlementTypesInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementTypesInnerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Multivalued) {
		toSerialize["multivalued"] = o.Multivalued
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementTypesInnerAttributes) UnmarshalJSON(data []byte) (err error) {
	varEntitlementTypesInnerAttributes := _EntitlementTypesInnerAttributes{}

	err = json.Unmarshal(data, &varEntitlementTypesInnerAttributes)

	if err != nil {
		return err
	}

	*o = EntitlementTypesInnerAttributes(varEntitlementTypesInnerAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "required")
		delete(additionalProperties, "multivalued")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementTypesInnerAttributes struct {
	value *EntitlementTypesInnerAttributes
	isSet bool
}

func (v NullableEntitlementTypesInnerAttributes) Get() *EntitlementTypesInnerAttributes {
	return v.value
}

func (v *NullableEntitlementTypesInnerAttributes) Set(val *EntitlementTypesInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementTypesInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementTypesInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementTypesInnerAttributes(val *EntitlementTypesInnerAttributes) *NullableEntitlementTypesInnerAttributes {
	return &NullableEntitlementTypesInnerAttributes{value: val, isSet: true}
}

func (v NullableEntitlementTypesInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementTypesInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

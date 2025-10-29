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

// checks if the SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner{}

// SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner struct for SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner
type SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner struct {
	Attributes *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes `json:"attributes,omitempty"`
	// The actual value of the attribute
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner

// NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner {
	this := SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner{}
	return &this
}

// NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerWithDefaults instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerWithDefaults() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner {
	this := SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) GetAttributes() SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) GetAttributesOk() (*SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes and assigns it to the Attributes field.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) SetAttributes(v SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) {
	o.Attributes = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) SetValue(v string) {
	o.Value = &v
}

func (o SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner := _SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner(varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner struct {
	value *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) Get() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) Set(val *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner(val *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner {
	return &NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

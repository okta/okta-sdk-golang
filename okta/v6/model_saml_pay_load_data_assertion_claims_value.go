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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// SAMLPayLoadDataAssertionClaimsValue struct for SAMLPayLoadDataAssertionClaimsValue
type SAMLPayLoadDataAssertionClaimsValue struct {
	Attributes *SAMLPayLoadDataAssertionClaimsValueAttributes `json:"attributes,omitempty"`
	AttributeValues []SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner `json:"attributeValues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionClaimsValue SAMLPayLoadDataAssertionClaimsValue

// NewSAMLPayLoadDataAssertionClaimsValue instantiates a new SAMLPayLoadDataAssertionClaimsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionClaimsValue() *SAMLPayLoadDataAssertionClaimsValue {
	this := SAMLPayLoadDataAssertionClaimsValue{}
	return &this
}

// NewSAMLPayLoadDataAssertionClaimsValueWithDefaults instantiates a new SAMLPayLoadDataAssertionClaimsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionClaimsValueWithDefaults() *SAMLPayLoadDataAssertionClaimsValue {
	this := SAMLPayLoadDataAssertionClaimsValue{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValue) GetAttributes() SAMLPayLoadDataAssertionClaimsValueAttributes {
	if o == nil || o.Attributes == nil {
		var ret SAMLPayLoadDataAssertionClaimsValueAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValue) GetAttributesOk() (*SAMLPayLoadDataAssertionClaimsValueAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValue) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SAMLPayLoadDataAssertionClaimsValueAttributes and assigns it to the Attributes field.
func (o *SAMLPayLoadDataAssertionClaimsValue) SetAttributes(v SAMLPayLoadDataAssertionClaimsValueAttributes) {
	o.Attributes = &v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValue) GetAttributeValues() []SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner {
	if o == nil || o.AttributeValues == nil {
		var ret []SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValue) GetAttributeValuesOk() ([]SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner, bool) {
	if o == nil || o.AttributeValues == nil {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValue) HasAttributeValues() bool {
	if o != nil && o.AttributeValues != nil {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner and assigns it to the AttributeValues field.
func (o *SAMLPayLoadDataAssertionClaimsValue) SetAttributeValues(v []SAMLPayLoadDataAssertionClaimsValueAttributeValuesInner) {
	o.AttributeValues = v
}

func (o SAMLPayLoadDataAssertionClaimsValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.AttributeValues != nil {
		toSerialize["attributeValues"] = o.AttributeValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SAMLPayLoadDataAssertionClaimsValue) UnmarshalJSON(bytes []byte) (err error) {
	varSAMLPayLoadDataAssertionClaimsValue := _SAMLPayLoadDataAssertionClaimsValue{}

	err = json.Unmarshal(bytes, &varSAMLPayLoadDataAssertionClaimsValue)
	if err == nil {
		*o = SAMLPayLoadDataAssertionClaimsValue(varSAMLPayLoadDataAssertionClaimsValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "attributeValues")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSAMLPayLoadDataAssertionClaimsValue struct {
	value *SAMLPayLoadDataAssertionClaimsValue
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionClaimsValue) Get() *SAMLPayLoadDataAssertionClaimsValue {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValue) Set(val *SAMLPayLoadDataAssertionClaimsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionClaimsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionClaimsValue(val *SAMLPayLoadDataAssertionClaimsValue) *NullableSAMLPayLoadDataAssertionClaimsValue {
	return &NullableSAMLPayLoadDataAssertionClaimsValue{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionClaimsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


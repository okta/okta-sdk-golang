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

// checks if the SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes{}

// SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes struct for SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes
type SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes struct {
	// Used to derive the type of the attribute
	XsiType              *string `json:"xsi:type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes

// NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes {
	this := SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes{}
	return &this
}

// NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributesWithDefaults instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributesWithDefaults() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes {
	this := SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes{}
	return &this
}

// GetXsiType returns the XsiType field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) GetXsiType() string {
	if o == nil || IsNil(o.XsiType) {
		var ret string
		return ret
	}
	return *o.XsiType
}

// GetXsiTypeOk returns a tuple with the XsiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) GetXsiTypeOk() (*string, bool) {
	if o == nil || IsNil(o.XsiType) {
		return nil, false
	}
	return o.XsiType, true
}

// HasXsiType returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) HasXsiType() bool {
	if o != nil && !IsNil(o.XsiType) {
		return true
	}

	return false
}

// SetXsiType gets a reference to the given string and assigns it to the XsiType field.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) SetXsiType(v string) {
	o.XsiType = &v
}

func (o SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.XsiType) {
		toSerialize["xsi:type"] = o.XsiType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes := _SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes(varSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "xsi:type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes struct {
	value *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) Get() *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) Set(val *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes(val *SAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes {
	return &NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributeValuesInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

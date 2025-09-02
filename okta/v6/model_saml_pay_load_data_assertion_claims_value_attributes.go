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

// SAMLPayLoadDataAssertionClaimsValueAttributes struct for SAMLPayLoadDataAssertionClaimsValueAttributes
type SAMLPayLoadDataAssertionClaimsValueAttributes struct {
	// Indicates how to interpret the attribute name
	NameFormat *string `json:"NameFormat,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionClaimsValueAttributes SAMLPayLoadDataAssertionClaimsValueAttributes

// NewSAMLPayLoadDataAssertionClaimsValueAttributes instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionClaimsValueAttributes() *SAMLPayLoadDataAssertionClaimsValueAttributes {
	this := SAMLPayLoadDataAssertionClaimsValueAttributes{}
	return &this
}

// NewSAMLPayLoadDataAssertionClaimsValueAttributesWithDefaults instantiates a new SAMLPayLoadDataAssertionClaimsValueAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionClaimsValueAttributesWithDefaults() *SAMLPayLoadDataAssertionClaimsValueAttributes {
	this := SAMLPayLoadDataAssertionClaimsValueAttributes{}
	return &this
}

// GetNameFormat returns the NameFormat field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributes) GetNameFormat() string {
	if o == nil || o.NameFormat == nil {
		var ret string
		return ret
	}
	return *o.NameFormat
}

// GetNameFormatOk returns a tuple with the NameFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributes) GetNameFormatOk() (*string, bool) {
	if o == nil || o.NameFormat == nil {
		return nil, false
	}
	return o.NameFormat, true
}

// HasNameFormat returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributes) HasNameFormat() bool {
	if o != nil && o.NameFormat != nil {
		return true
	}

	return false
}

// SetNameFormat gets a reference to the given string and assigns it to the NameFormat field.
func (o *SAMLPayLoadDataAssertionClaimsValueAttributes) SetNameFormat(v string) {
	o.NameFormat = &v
}

func (o SAMLPayLoadDataAssertionClaimsValueAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NameFormat != nil {
		toSerialize["NameFormat"] = o.NameFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SAMLPayLoadDataAssertionClaimsValueAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varSAMLPayLoadDataAssertionClaimsValueAttributes := _SAMLPayLoadDataAssertionClaimsValueAttributes{}

	err = json.Unmarshal(bytes, &varSAMLPayLoadDataAssertionClaimsValueAttributes)
	if err == nil {
		*o = SAMLPayLoadDataAssertionClaimsValueAttributes(varSAMLPayLoadDataAssertionClaimsValueAttributes)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "NameFormat")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSAMLPayLoadDataAssertionClaimsValueAttributes struct {
	value *SAMLPayLoadDataAssertionClaimsValueAttributes
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributes) Get() *SAMLPayLoadDataAssertionClaimsValueAttributes {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributes) Set(val *SAMLPayLoadDataAssertionClaimsValueAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionClaimsValueAttributes(val *SAMLPayLoadDataAssertionClaimsValueAttributes) *NullableSAMLPayLoadDataAssertionClaimsValueAttributes {
	return &NullableSAMLPayLoadDataAssertionClaimsValueAttributes{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionClaimsValueAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionClaimsValueAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


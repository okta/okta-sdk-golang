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

// checks if the SAMLPayLoadDataAssertionSubjectConfirmation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionSubjectConfirmation{}

// SAMLPayLoadDataAssertionSubjectConfirmation struct for SAMLPayLoadDataAssertionSubjectConfirmation
type SAMLPayLoadDataAssertionSubjectConfirmation struct {
	// Used to indicate how the authorization server confirmed the SAML assertion
	Method               *string                                          `json:"method,omitempty"`
	Data                 *SAMLPayLoadDataAssertionSubjectConfirmationData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionSubjectConfirmation SAMLPayLoadDataAssertionSubjectConfirmation

// NewSAMLPayLoadDataAssertionSubjectConfirmation instantiates a new SAMLPayLoadDataAssertionSubjectConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionSubjectConfirmation() *SAMLPayLoadDataAssertionSubjectConfirmation {
	this := SAMLPayLoadDataAssertionSubjectConfirmation{}
	return &this
}

// NewSAMLPayLoadDataAssertionSubjectConfirmationWithDefaults instantiates a new SAMLPayLoadDataAssertionSubjectConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionSubjectConfirmationWithDefaults() *SAMLPayLoadDataAssertionSubjectConfirmation {
	this := SAMLPayLoadDataAssertionSubjectConfirmation{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) SetMethod(v string) {
	o.Method = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetData() SAMLPayLoadDataAssertionSubjectConfirmationData {
	if o == nil || IsNil(o.Data) {
		var ret SAMLPayLoadDataAssertionSubjectConfirmationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetDataOk() (*SAMLPayLoadDataAssertionSubjectConfirmationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SAMLPayLoadDataAssertionSubjectConfirmationData and assigns it to the Data field.
func (o *SAMLPayLoadDataAssertionSubjectConfirmation) SetData(v SAMLPayLoadDataAssertionSubjectConfirmationData) {
	o.Data = &v
}

func (o SAMLPayLoadDataAssertionSubjectConfirmation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionSubjectConfirmation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionSubjectConfirmation) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionSubjectConfirmation := _SAMLPayLoadDataAssertionSubjectConfirmation{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionSubjectConfirmation)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionSubjectConfirmation(varSAMLPayLoadDataAssertionSubjectConfirmation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionSubjectConfirmation struct {
	value *SAMLPayLoadDataAssertionSubjectConfirmation
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmation) Get() *SAMLPayLoadDataAssertionSubjectConfirmation {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmation) Set(val *SAMLPayLoadDataAssertionSubjectConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionSubjectConfirmation(val *SAMLPayLoadDataAssertionSubjectConfirmation) *NullableSAMLPayLoadDataAssertionSubjectConfirmation {
	return &NullableSAMLPayLoadDataAssertionSubjectConfirmation{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

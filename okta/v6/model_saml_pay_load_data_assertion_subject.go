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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SAMLPayLoadDataAssertionSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionSubject{}

// SAMLPayLoadDataAssertionSubject Provides a JSON representation of the `<saml:Subject>` element of the SAML assertion
type SAMLPayLoadDataAssertionSubject struct {
	// The unique identifier of the user
	NameId *string `json:"nameId,omitempty"`
	// Indicates how to interpret the attribute name
	NameFormat           *string                                      `json:"nameFormat,omitempty"`
	Confirmation         *SAMLPayLoadDataAssertionSubjectConfirmation `json:"confirmation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionSubject SAMLPayLoadDataAssertionSubject

// NewSAMLPayLoadDataAssertionSubject instantiates a new SAMLPayLoadDataAssertionSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionSubject() *SAMLPayLoadDataAssertionSubject {
	this := SAMLPayLoadDataAssertionSubject{}
	return &this
}

// NewSAMLPayLoadDataAssertionSubjectWithDefaults instantiates a new SAMLPayLoadDataAssertionSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionSubjectWithDefaults() *SAMLPayLoadDataAssertionSubject {
	this := SAMLPayLoadDataAssertionSubject{}
	return &this
}

// GetNameId returns the NameId field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubject) GetNameId() string {
	if o == nil || IsNil(o.NameId) {
		var ret string
		return ret
	}
	return *o.NameId
}

// GetNameIdOk returns a tuple with the NameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubject) GetNameIdOk() (*string, bool) {
	if o == nil || IsNil(o.NameId) {
		return nil, false
	}
	return o.NameId, true
}

// HasNameId returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubject) HasNameId() bool {
	if o != nil && !IsNil(o.NameId) {
		return true
	}

	return false
}

// SetNameId gets a reference to the given string and assigns it to the NameId field.
func (o *SAMLPayLoadDataAssertionSubject) SetNameId(v string) {
	o.NameId = &v
}

// GetNameFormat returns the NameFormat field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubject) GetNameFormat() string {
	if o == nil || IsNil(o.NameFormat) {
		var ret string
		return ret
	}
	return *o.NameFormat
}

// GetNameFormatOk returns a tuple with the NameFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubject) GetNameFormatOk() (*string, bool) {
	if o == nil || IsNil(o.NameFormat) {
		return nil, false
	}
	return o.NameFormat, true
}

// HasNameFormat returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubject) HasNameFormat() bool {
	if o != nil && !IsNil(o.NameFormat) {
		return true
	}

	return false
}

// SetNameFormat gets a reference to the given string and assigns it to the NameFormat field.
func (o *SAMLPayLoadDataAssertionSubject) SetNameFormat(v string) {
	o.NameFormat = &v
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubject) GetConfirmation() SAMLPayLoadDataAssertionSubjectConfirmation {
	if o == nil || IsNil(o.Confirmation) {
		var ret SAMLPayLoadDataAssertionSubjectConfirmation
		return ret
	}
	return *o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubject) GetConfirmationOk() (*SAMLPayLoadDataAssertionSubjectConfirmation, bool) {
	if o == nil || IsNil(o.Confirmation) {
		return nil, false
	}
	return o.Confirmation, true
}

// HasConfirmation returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubject) HasConfirmation() bool {
	if o != nil && !IsNil(o.Confirmation) {
		return true
	}

	return false
}

// SetConfirmation gets a reference to the given SAMLPayLoadDataAssertionSubjectConfirmation and assigns it to the Confirmation field.
func (o *SAMLPayLoadDataAssertionSubject) SetConfirmation(v SAMLPayLoadDataAssertionSubjectConfirmation) {
	o.Confirmation = &v
}

func (o SAMLPayLoadDataAssertionSubject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NameId) {
		toSerialize["nameId"] = o.NameId
	}
	if !IsNil(o.NameFormat) {
		toSerialize["nameFormat"] = o.NameFormat
	}
	if !IsNil(o.Confirmation) {
		toSerialize["confirmation"] = o.Confirmation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionSubject) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionSubject := _SAMLPayLoadDataAssertionSubject{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionSubject)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionSubject(varSAMLPayLoadDataAssertionSubject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nameId")
		delete(additionalProperties, "nameFormat")
		delete(additionalProperties, "confirmation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionSubject struct {
	value *SAMLPayLoadDataAssertionSubject
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionSubject) Get() *SAMLPayLoadDataAssertionSubject {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionSubject) Set(val *SAMLPayLoadDataAssertionSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionSubject(val *SAMLPayLoadDataAssertionSubject) *NullableSAMLPayLoadDataAssertionSubject {
	return &NullableSAMLPayLoadDataAssertionSubject{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the SAMLPayLoadDataAssertionSubjectConfirmationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionSubjectConfirmationData{}

// SAMLPayLoadDataAssertionSubjectConfirmationData struct for SAMLPayLoadDataAssertionSubjectConfirmationData
type SAMLPayLoadDataAssertionSubjectConfirmationData struct {
	// The token endpoint URL of the authorization server
	Recipient            *string `json:"recipient,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionSubjectConfirmationData SAMLPayLoadDataAssertionSubjectConfirmationData

// NewSAMLPayLoadDataAssertionSubjectConfirmationData instantiates a new SAMLPayLoadDataAssertionSubjectConfirmationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionSubjectConfirmationData() *SAMLPayLoadDataAssertionSubjectConfirmationData {
	this := SAMLPayLoadDataAssertionSubjectConfirmationData{}
	return &this
}

// NewSAMLPayLoadDataAssertionSubjectConfirmationDataWithDefaults instantiates a new SAMLPayLoadDataAssertionSubjectConfirmationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionSubjectConfirmationDataWithDefaults() *SAMLPayLoadDataAssertionSubjectConfirmationData {
	this := SAMLPayLoadDataAssertionSubjectConfirmationData{}
	return &this
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionSubjectConfirmationData) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmationData) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionSubjectConfirmationData) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *SAMLPayLoadDataAssertionSubjectConfirmationData) SetRecipient(v string) {
	o.Recipient = &v
}

func (o SAMLPayLoadDataAssertionSubjectConfirmationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionSubjectConfirmationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionSubjectConfirmationData) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionSubjectConfirmationData := _SAMLPayLoadDataAssertionSubjectConfirmationData{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionSubjectConfirmationData)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionSubjectConfirmationData(varSAMLPayLoadDataAssertionSubjectConfirmationData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionSubjectConfirmationData struct {
	value *SAMLPayLoadDataAssertionSubjectConfirmationData
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmationData) Get() *SAMLPayLoadDataAssertionSubjectConfirmationData {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmationData) Set(val *SAMLPayLoadDataAssertionSubjectConfirmationData) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmationData) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionSubjectConfirmationData(val *SAMLPayLoadDataAssertionSubjectConfirmationData) *NullableSAMLPayLoadDataAssertionSubjectConfirmationData {
	return &NullableSAMLPayLoadDataAssertionSubjectConfirmationData{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionSubjectConfirmationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionSubjectConfirmationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

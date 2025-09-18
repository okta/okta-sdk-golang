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

// checks if the U2f1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &U2f1{}

// U2f1 Verifies a `u2f` factor challenge by posting a signed assertion using the challenge `nonce`
type U2f1 struct {
	// Base64-encoded client data from the U2F token
	ClientData *string `json:"clientData,omitempty"`
	// Base64-encoded signature data from the U2F token
	SignatureData        interface{} `json:"signatureData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _U2f1 U2f1

// NewU2f1 instantiates a new U2f1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewU2f1() *U2f1 {
	this := U2f1{}
	return &this
}

// NewU2f1WithDefaults instantiates a new U2f1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewU2f1WithDefaults() *U2f1 {
	this := U2f1{}
	return &this
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *U2f1) GetClientData() string {
	if o == nil || IsNil(o.ClientData) {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *U2f1) GetClientDataOk() (*string, bool) {
	if o == nil || IsNil(o.ClientData) {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *U2f1) HasClientData() bool {
	if o != nil && !IsNil(o.ClientData) {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *U2f1) SetClientData(v string) {
	o.ClientData = &v
}

// GetSignatureData returns the SignatureData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *U2f1) GetSignatureData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SignatureData
}

// GetSignatureDataOk returns a tuple with the SignatureData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *U2f1) GetSignatureDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SignatureData) {
		return nil, false
	}
	return &o.SignatureData, true
}

// HasSignatureData returns a boolean if a field has been set.
func (o *U2f1) HasSignatureData() bool {
	if o != nil && !IsNil(o.SignatureData) {
		return true
	}

	return false
}

// SetSignatureData gets a reference to the given interface{} and assigns it to the SignatureData field.
func (o *U2f1) SetSignatureData(v interface{}) {
	o.SignatureData = v
}

func (o U2f1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o U2f1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientData) {
		toSerialize["clientData"] = o.ClientData
	}
	if o.SignatureData != nil {
		toSerialize["signatureData"] = o.SignatureData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *U2f1) UnmarshalJSON(data []byte) (err error) {
	varU2f1 := _U2f1{}

	err = json.Unmarshal(data, &varU2f1)

	if err != nil {
		return err
	}

	*o = U2f1(varU2f1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientData")
		delete(additionalProperties, "signatureData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableU2f1 struct {
	value *U2f1
	isSet bool
}

func (v NullableU2f1) Get() *U2f1 {
	return v.value
}

func (v *NullableU2f1) Set(val *U2f1) {
	v.value = val
	v.isSet = true
}

func (v NullableU2f1) IsSet() bool {
	return v.isSet
}

func (v *NullableU2f1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableU2f1(val *U2f1) *NullableU2f1 {
	return &NullableU2f1{value: val, isSet: true}
}

func (v NullableU2f1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableU2f1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

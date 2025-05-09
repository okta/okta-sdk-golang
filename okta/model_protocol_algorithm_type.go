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

// ProtocolAlgorithmType struct for ProtocolAlgorithmType
type ProtocolAlgorithmType struct {
	Signature *ProtocolAlgorithmTypeSignature `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolAlgorithmType ProtocolAlgorithmType

// NewProtocolAlgorithmType instantiates a new ProtocolAlgorithmType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolAlgorithmType() *ProtocolAlgorithmType {
	this := ProtocolAlgorithmType{}
	return &this
}

// NewProtocolAlgorithmTypeWithDefaults instantiates a new ProtocolAlgorithmType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolAlgorithmTypeWithDefaults() *ProtocolAlgorithmType {
	this := ProtocolAlgorithmType{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ProtocolAlgorithmType) GetSignature() ProtocolAlgorithmTypeSignature {
	if o == nil || o.Signature == nil {
		var ret ProtocolAlgorithmTypeSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolAlgorithmType) GetSignatureOk() (*ProtocolAlgorithmTypeSignature, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ProtocolAlgorithmType) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given ProtocolAlgorithmTypeSignature and assigns it to the Signature field.
func (o *ProtocolAlgorithmType) SetSignature(v ProtocolAlgorithmTypeSignature) {
	o.Signature = &v
}

func (o ProtocolAlgorithmType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolAlgorithmType) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolAlgorithmType := _ProtocolAlgorithmType{}

	err = json.Unmarshal(bytes, &varProtocolAlgorithmType)
	if err == nil {
		*o = ProtocolAlgorithmType(varProtocolAlgorithmType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolAlgorithmType struct {
	value *ProtocolAlgorithmType
	isSet bool
}

func (v NullableProtocolAlgorithmType) Get() *ProtocolAlgorithmType {
	return v.value
}

func (v *NullableProtocolAlgorithmType) Set(val *ProtocolAlgorithmType) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolAlgorithmType) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolAlgorithmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolAlgorithmType(val *ProtocolAlgorithmType) *NullableProtocolAlgorithmType {
	return &NullableProtocolAlgorithmType{value: val, isSet: true}
}

func (v NullableProtocolAlgorithmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolAlgorithmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


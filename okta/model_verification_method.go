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

// checks if the VerificationMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationMethod{}

// VerificationMethod The method used to verify a user
type VerificationMethod struct {
	// Verification method type
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerificationMethod VerificationMethod

// NewVerificationMethod instantiates a new VerificationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationMethod() *VerificationMethod {
	this := VerificationMethod{}
	return &this
}

// NewVerificationMethodWithDefaults instantiates a new VerificationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationMethodWithDefaults() *VerificationMethod {
	this := VerificationMethod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VerificationMethod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMethod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VerificationMethod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VerificationMethod) SetType(v string) {
	o.Type = &v
}

func (o VerificationMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerificationMethod) UnmarshalJSON(data []byte) (err error) {
	varVerificationMethod := _VerificationMethod{}

	err = json.Unmarshal(data, &varVerificationMethod)

	if err != nil {
		return err
	}

	*o = VerificationMethod(varVerificationMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationMethod struct {
	value *VerificationMethod
	isSet bool
}

func (v NullableVerificationMethod) Get() *VerificationMethod {
	return v.value
}

func (v *NullableVerificationMethod) Set(val *VerificationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationMethod(val *VerificationMethod) *NullableVerificationMethod {
	return &NullableVerificationMethod{value: val, isSet: true}
}

func (v NullableVerificationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

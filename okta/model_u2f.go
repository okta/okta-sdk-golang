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

// checks if the U2f type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &U2f{}

// U2f Activates a `u2f` factor with the specified client and registration information from the U2F token
type U2f struct {
	// Base64-encoded client data from the U2F token
	ClientData *string `json:"clientData,omitempty"`
	// Base64-encoded registration data from the U2F token
	RegistrationData     *string `json:"registrationData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _U2f U2f

// NewU2f instantiates a new U2f object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewU2f() *U2f {
	this := U2f{}
	return &this
}

// NewU2fWithDefaults instantiates a new U2f object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewU2fWithDefaults() *U2f {
	this := U2f{}
	return &this
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *U2f) GetClientData() string {
	if o == nil || IsNil(o.ClientData) {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *U2f) GetClientDataOk() (*string, bool) {
	if o == nil || IsNil(o.ClientData) {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *U2f) HasClientData() bool {
	if o != nil && !IsNil(o.ClientData) {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *U2f) SetClientData(v string) {
	o.ClientData = &v
}

// GetRegistrationData returns the RegistrationData field value if set, zero value otherwise.
func (o *U2f) GetRegistrationData() string {
	if o == nil || IsNil(o.RegistrationData) {
		var ret string
		return ret
	}
	return *o.RegistrationData
}

// GetRegistrationDataOk returns a tuple with the RegistrationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *U2f) GetRegistrationDataOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationData) {
		return nil, false
	}
	return o.RegistrationData, true
}

// HasRegistrationData returns a boolean if a field has been set.
func (o *U2f) HasRegistrationData() bool {
	if o != nil && !IsNil(o.RegistrationData) {
		return true
	}

	return false
}

// SetRegistrationData gets a reference to the given string and assigns it to the RegistrationData field.
func (o *U2f) SetRegistrationData(v string) {
	o.RegistrationData = &v
}

func (o U2f) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o U2f) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientData) {
		toSerialize["clientData"] = o.ClientData
	}
	if !IsNil(o.RegistrationData) {
		toSerialize["registrationData"] = o.RegistrationData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *U2f) UnmarshalJSON(data []byte) (err error) {
	varU2f := _U2f{}

	err = json.Unmarshal(data, &varU2f)

	if err != nil {
		return err
	}

	*o = U2f(varU2f)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientData")
		delete(additionalProperties, "registrationData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableU2f struct {
	value *U2f
	isSet bool
}

func (v NullableU2f) Get() *U2f {
	return v.value
}

func (v *NullableU2f) Set(val *U2f) {
	v.value = val
	v.isSet = true
}

func (v NullableU2f) IsSet() bool {
	return v.isSet
}

func (v *NullableU2f) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableU2f(val *U2f) *NullableU2f {
	return &NullableU2f{value: val, isSet: true}
}

func (v NullableU2f) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableU2f) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

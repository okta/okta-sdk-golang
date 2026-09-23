/*
Okta Admin Management API

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

// checks if the UserIdentificationSecurityMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationSecurityMethods{}

// UserIdentificationSecurityMethods <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the security methods that can be used to identify the user when signing in
type UserIdentificationSecurityMethods struct {
	Fastpass             *UserIdentificationFastpassSettings `json:"fastpass,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationSecurityMethods UserIdentificationSecurityMethods

// NewUserIdentificationSecurityMethods instantiates a new UserIdentificationSecurityMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationSecurityMethods() *UserIdentificationSecurityMethods {
	this := UserIdentificationSecurityMethods{}
	return &this
}

// NewUserIdentificationSecurityMethodsWithDefaults instantiates a new UserIdentificationSecurityMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationSecurityMethodsWithDefaults() *UserIdentificationSecurityMethods {
	this := UserIdentificationSecurityMethods{}
	return &this
}

// GetFastpass returns the Fastpass field value if set, zero value otherwise.
func (o *UserIdentificationSecurityMethods) GetFastpass() UserIdentificationFastpassSettings {
	if o == nil || IsNil(o.Fastpass) {
		var ret UserIdentificationFastpassSettings
		return ret
	}
	return *o.Fastpass
}

// GetFastpassOk returns a tuple with the Fastpass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationSecurityMethods) GetFastpassOk() (*UserIdentificationFastpassSettings, bool) {
	if o == nil || IsNil(o.Fastpass) {
		return nil, false
	}
	return o.Fastpass, true
}

// HasFastpass returns a boolean if a field has been set.
func (o *UserIdentificationSecurityMethods) HasFastpass() bool {
	if o != nil && !IsNil(o.Fastpass) {
		return true
	}

	return false
}

// SetFastpass gets a reference to the given UserIdentificationFastpassSettings and assigns it to the Fastpass field.
func (o *UserIdentificationSecurityMethods) SetFastpass(v UserIdentificationFastpassSettings) {
	o.Fastpass = &v
}

func (o UserIdentificationSecurityMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationSecurityMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fastpass) {
		toSerialize["fastpass"] = o.Fastpass
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationSecurityMethods) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationSecurityMethods := _UserIdentificationSecurityMethods{}

	err = json.Unmarshal(data, &varUserIdentificationSecurityMethods)

	if err != nil {
		return err
	}

	*o = UserIdentificationSecurityMethods(varUserIdentificationSecurityMethods)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fastpass")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationSecurityMethods struct {
	value *UserIdentificationSecurityMethods
	isSet bool
}

func (v NullableUserIdentificationSecurityMethods) Get() *UserIdentificationSecurityMethods {
	return v.value
}

func (v *NullableUserIdentificationSecurityMethods) Set(val *UserIdentificationSecurityMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationSecurityMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationSecurityMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationSecurityMethods(val *UserIdentificationSecurityMethods) *NullableUserIdentificationSecurityMethods {
	return &NullableUserIdentificationSecurityMethods{value: val, isSet: true}
}

func (v NullableUserIdentificationSecurityMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationSecurityMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

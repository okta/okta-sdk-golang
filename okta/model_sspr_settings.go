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

// checks if the SsprSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsprSettings{}

// SsprSettings <x-lifecycle class=\"oie\"></x-lifecycle> Settings that change how a user can recover their password
type SsprSettings struct {
	// Allows a user to recover their password with their email address when they have not enrolled the email authenticator
	AllowRecoveryEmailWithoutEnrollment *bool `json:"allowRecoveryEmailWithoutEnrollment,omitempty"`
	AdditionalProperties                map[string]interface{}
}

type _SsprSettings SsprSettings

// NewSsprSettings instantiates a new SsprSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprSettings() *SsprSettings {
	this := SsprSettings{}
	return &this
}

// NewSsprSettingsWithDefaults instantiates a new SsprSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprSettingsWithDefaults() *SsprSettings {
	this := SsprSettings{}
	return &this
}

// GetAllowRecoveryEmailWithoutEnrollment returns the AllowRecoveryEmailWithoutEnrollment field value if set, zero value otherwise.
func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollment() bool {
	if o == nil || IsNil(o.AllowRecoveryEmailWithoutEnrollment) {
		var ret bool
		return ret
	}
	return *o.AllowRecoveryEmailWithoutEnrollment
}

// GetAllowRecoveryEmailWithoutEnrollmentOk returns a tuple with the AllowRecoveryEmailWithoutEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRecoveryEmailWithoutEnrollment) {
		return nil, false
	}
	return o.AllowRecoveryEmailWithoutEnrollment, true
}

// HasAllowRecoveryEmailWithoutEnrollment returns a boolean if a field has been set.
func (o *SsprSettings) HasAllowRecoveryEmailWithoutEnrollment() bool {
	if o != nil && !IsNil(o.AllowRecoveryEmailWithoutEnrollment) {
		return true
	}

	return false
}

// SetAllowRecoveryEmailWithoutEnrollment gets a reference to the given bool and assigns it to the AllowRecoveryEmailWithoutEnrollment field.
func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollment(v bool) {
	o.AllowRecoveryEmailWithoutEnrollment = &v
}

func (o SsprSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsprSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowRecoveryEmailWithoutEnrollment) {
		toSerialize["allowRecoveryEmailWithoutEnrollment"] = o.AllowRecoveryEmailWithoutEnrollment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsprSettings) UnmarshalJSON(data []byte) (err error) {
	varSsprSettings := _SsprSettings{}

	err = json.Unmarshal(data, &varSsprSettings)

	if err != nil {
		return err
	}

	*o = SsprSettings(varSsprSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowRecoveryEmailWithoutEnrollment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsprSettings struct {
	value *SsprSettings
	isSet bool
}

func (v NullableSsprSettings) Get() *SsprSettings {
	return v.value
}

func (v *NullableSsprSettings) Set(val *SsprSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprSettings(val *SsprSettings) *NullableSsprSettings {
	return &NullableSsprSettings{value: val, isSet: true}
}

func (v NullableSsprSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

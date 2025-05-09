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

// MultifactorEnrollmentPolicyAuthenticatorSettings struct for MultifactorEnrollmentPolicyAuthenticatorSettings
type MultifactorEnrollmentPolicyAuthenticatorSettings struct {
	Constraints NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints `json:"constraints,omitempty"`
	Enroll *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll `json:"enroll,omitempty"`
	Key *string `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicyAuthenticatorSettings MultifactorEnrollmentPolicyAuthenticatorSettings

// NewMultifactorEnrollmentPolicyAuthenticatorSettings instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicyAuthenticatorSettings() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	this := MultifactorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsWithDefaults instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsWithDefaults() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	this := MultifactorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetConstraints() MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetConstraintsOk() (*MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableMultifactorEnrollmentPolicyAuthenticatorSettingsConstraints and assigns it to the Constraints field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetConstraints(v MultifactorEnrollmentPolicyAuthenticatorSettingsConstraints) {
	o.Constraints.Set(&v)
}
// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetEnroll() MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	if o == nil || o.Enroll == nil {
		var ret MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetEnrollOk() (*MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll and assigns it to the Enroll field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetEnroll(v MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) {
	o.Enroll = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) SetKey(v string) {
	o.Key = &v
}

func (o MultifactorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MultifactorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicyAuthenticatorSettings := _MultifactorEnrollmentPolicyAuthenticatorSettings{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyAuthenticatorSettings)
	if err == nil {
		*o = MultifactorEnrollmentPolicyAuthenticatorSettings(varMultifactorEnrollmentPolicyAuthenticatorSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMultifactorEnrollmentPolicyAuthenticatorSettings struct {
	value *MultifactorEnrollmentPolicyAuthenticatorSettings
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Get() *MultifactorEnrollmentPolicyAuthenticatorSettings {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Set(val *MultifactorEnrollmentPolicyAuthenticatorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAuthenticatorSettings(val *MultifactorEnrollmentPolicyAuthenticatorSettings) *NullableMultifactorEnrollmentPolicyAuthenticatorSettings {
	return &NullableMultifactorEnrollmentPolicyAuthenticatorSettings{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


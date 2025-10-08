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

// checks if the PasswordPolicyRecoveryFactorSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryFactorSettings{}

// PasswordPolicyRecoveryFactorSettings struct for PasswordPolicyRecoveryFactorSettings
type PasswordPolicyRecoveryFactorSettings struct {
	// Whether or not the factor is active
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryFactorSettings PasswordPolicyRecoveryFactorSettings

// NewPasswordPolicyRecoveryFactorSettings instantiates a new PasswordPolicyRecoveryFactorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryFactorSettings() *PasswordPolicyRecoveryFactorSettings {
	this := PasswordPolicyRecoveryFactorSettings{}
	return &this
}

// NewPasswordPolicyRecoveryFactorSettingsWithDefaults instantiates a new PasswordPolicyRecoveryFactorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryFactorSettingsWithDefaults() *PasswordPolicyRecoveryFactorSettings {
	this := PasswordPolicyRecoveryFactorSettings{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryFactorSettings) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactorSettings) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactorSettings) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PasswordPolicyRecoveryFactorSettings) SetStatus(v string) {
	o.Status = &v
}

func (o PasswordPolicyRecoveryFactorSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryFactorSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryFactorSettings) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryFactorSettings := _PasswordPolicyRecoveryFactorSettings{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryFactorSettings)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryFactorSettings(varPasswordPolicyRecoveryFactorSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryFactorSettings struct {
	value *PasswordPolicyRecoveryFactorSettings
	isSet bool
}

func (v NullablePasswordPolicyRecoveryFactorSettings) Get() *PasswordPolicyRecoveryFactorSettings {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryFactorSettings) Set(val *PasswordPolicyRecoveryFactorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryFactorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryFactorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryFactorSettings(val *PasswordPolicyRecoveryFactorSettings) *NullablePasswordPolicyRecoveryFactorSettings {
	return &NullablePasswordPolicyRecoveryFactorSettings{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryFactorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryFactorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

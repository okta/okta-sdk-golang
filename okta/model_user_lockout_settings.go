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

// UserLockoutSettings struct for UserLockoutSettings
type UserLockoutSettings struct {
	// Prevents brute-force lockout from unknown devices for the password authenticator.
	PreventBruteForceLockoutFromUnknownDevices *bool `json:"preventBruteForceLockoutFromUnknownDevices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserLockoutSettings UserLockoutSettings

// NewUserLockoutSettings instantiates a new UserLockoutSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLockoutSettings() *UserLockoutSettings {
	this := UserLockoutSettings{}
	var preventBruteForceLockoutFromUnknownDevices bool = false
	this.PreventBruteForceLockoutFromUnknownDevices = &preventBruteForceLockoutFromUnknownDevices
	return &this
}

// NewUserLockoutSettingsWithDefaults instantiates a new UserLockoutSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLockoutSettingsWithDefaults() *UserLockoutSettings {
	this := UserLockoutSettings{}
	var preventBruteForceLockoutFromUnknownDevices bool = false
	this.PreventBruteForceLockoutFromUnknownDevices = &preventBruteForceLockoutFromUnknownDevices
	return &this
}

// GetPreventBruteForceLockoutFromUnknownDevices returns the PreventBruteForceLockoutFromUnknownDevices field value if set, zero value otherwise.
func (o *UserLockoutSettings) GetPreventBruteForceLockoutFromUnknownDevices() bool {
	if o == nil || o.PreventBruteForceLockoutFromUnknownDevices == nil {
		var ret bool
		return ret
	}
	return *o.PreventBruteForceLockoutFromUnknownDevices
}

// GetPreventBruteForceLockoutFromUnknownDevicesOk returns a tuple with the PreventBruteForceLockoutFromUnknownDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLockoutSettings) GetPreventBruteForceLockoutFromUnknownDevicesOk() (*bool, bool) {
	if o == nil || o.PreventBruteForceLockoutFromUnknownDevices == nil {
		return nil, false
	}
	return o.PreventBruteForceLockoutFromUnknownDevices, true
}

// HasPreventBruteForceLockoutFromUnknownDevices returns a boolean if a field has been set.
func (o *UserLockoutSettings) HasPreventBruteForceLockoutFromUnknownDevices() bool {
	if o != nil && o.PreventBruteForceLockoutFromUnknownDevices != nil {
		return true
	}

	return false
}

// SetPreventBruteForceLockoutFromUnknownDevices gets a reference to the given bool and assigns it to the PreventBruteForceLockoutFromUnknownDevices field.
func (o *UserLockoutSettings) SetPreventBruteForceLockoutFromUnknownDevices(v bool) {
	o.PreventBruteForceLockoutFromUnknownDevices = &v
}

func (o UserLockoutSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PreventBruteForceLockoutFromUnknownDevices != nil {
		toSerialize["preventBruteForceLockoutFromUnknownDevices"] = o.PreventBruteForceLockoutFromUnknownDevices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserLockoutSettings) UnmarshalJSON(bytes []byte) (err error) {
	varUserLockoutSettings := _UserLockoutSettings{}

	err = json.Unmarshal(bytes, &varUserLockoutSettings)
	if err == nil {
		*o = UserLockoutSettings(varUserLockoutSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "preventBruteForceLockoutFromUnknownDevices")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserLockoutSettings struct {
	value *UserLockoutSettings
	isSet bool
}

func (v NullableUserLockoutSettings) Get() *UserLockoutSettings {
	return v.value
}

func (v *NullableUserLockoutSettings) Set(val *UserLockoutSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLockoutSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLockoutSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLockoutSettings(val *UserLockoutSettings) *NullableUserLockoutSettings {
	return &NullableUserLockoutSettings{value: val, isSet: true}
}

func (v NullableUserLockoutSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLockoutSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// PerClientRateLimitSettings 
type PerClientRateLimitSettings struct {
	DefaultMode string `json:"defaultMode"`
	UseCaseModeOverrides *PerClientRateLimitSettingsUseCaseModeOverrides `json:"useCaseModeOverrides,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PerClientRateLimitSettings PerClientRateLimitSettings

// NewPerClientRateLimitSettings instantiates a new PerClientRateLimitSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerClientRateLimitSettings(defaultMode string) *PerClientRateLimitSettings {
	this := PerClientRateLimitSettings{}
	this.DefaultMode = defaultMode
	return &this
}

// NewPerClientRateLimitSettingsWithDefaults instantiates a new PerClientRateLimitSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerClientRateLimitSettingsWithDefaults() *PerClientRateLimitSettings {
	this := PerClientRateLimitSettings{}
	return &this
}

// GetDefaultMode returns the DefaultMode field value
func (o *PerClientRateLimitSettings) GetDefaultMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettings) GetDefaultModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultMode, true
}

// SetDefaultMode sets field value
func (o *PerClientRateLimitSettings) SetDefaultMode(v string) {
	o.DefaultMode = v
}

// GetUseCaseModeOverrides returns the UseCaseModeOverrides field value if set, zero value otherwise.
func (o *PerClientRateLimitSettings) GetUseCaseModeOverrides() PerClientRateLimitSettingsUseCaseModeOverrides {
	if o == nil || o.UseCaseModeOverrides == nil {
		var ret PerClientRateLimitSettingsUseCaseModeOverrides
		return ret
	}
	return *o.UseCaseModeOverrides
}

// GetUseCaseModeOverridesOk returns a tuple with the UseCaseModeOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettings) GetUseCaseModeOverridesOk() (*PerClientRateLimitSettingsUseCaseModeOverrides, bool) {
	if o == nil || o.UseCaseModeOverrides == nil {
		return nil, false
	}
	return o.UseCaseModeOverrides, true
}

// HasUseCaseModeOverrides returns a boolean if a field has been set.
func (o *PerClientRateLimitSettings) HasUseCaseModeOverrides() bool {
	if o != nil && o.UseCaseModeOverrides != nil {
		return true
	}

	return false
}

// SetUseCaseModeOverrides gets a reference to the given PerClientRateLimitSettingsUseCaseModeOverrides and assigns it to the UseCaseModeOverrides field.
func (o *PerClientRateLimitSettings) SetUseCaseModeOverrides(v PerClientRateLimitSettingsUseCaseModeOverrides) {
	o.UseCaseModeOverrides = &v
}

func (o PerClientRateLimitSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["defaultMode"] = o.DefaultMode
	}
	if o.UseCaseModeOverrides != nil {
		toSerialize["useCaseModeOverrides"] = o.UseCaseModeOverrides
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PerClientRateLimitSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPerClientRateLimitSettings := _PerClientRateLimitSettings{}

	err = json.Unmarshal(bytes, &varPerClientRateLimitSettings)
	if err == nil {
		*o = PerClientRateLimitSettings(varPerClientRateLimitSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "defaultMode")
		delete(additionalProperties, "useCaseModeOverrides")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePerClientRateLimitSettings struct {
	value *PerClientRateLimitSettings
	isSet bool
}

func (v NullablePerClientRateLimitSettings) Get() *PerClientRateLimitSettings {
	return v.value
}

func (v *NullablePerClientRateLimitSettings) Set(val *PerClientRateLimitSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePerClientRateLimitSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePerClientRateLimitSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerClientRateLimitSettings(val *PerClientRateLimitSettings) *NullablePerClientRateLimitSettings {
	return &NullablePerClientRateLimitSettings{value: val, isSet: true}
}

func (v NullablePerClientRateLimitSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerClientRateLimitSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


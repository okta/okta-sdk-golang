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

// AdminConsoleSettings Settings specific to the Okta Admin Console
type AdminConsoleSettings struct {
	// The maximum idle time before the Okta Admin Console session expires. Must be no more than 12 hours.
	SessionIdleTimeoutMinutes *int32 `json:"sessionIdleTimeoutMinutes,omitempty"`
	// The absolute maximum session lifetime of the Okta Admin Console. Must be no more than 7 days.
	SessionMaxLifetimeMinutes *int32 `json:"sessionMaxLifetimeMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdminConsoleSettings AdminConsoleSettings

// NewAdminConsoleSettings instantiates a new AdminConsoleSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminConsoleSettings() *AdminConsoleSettings {
	this := AdminConsoleSettings{}
	var sessionIdleTimeoutMinutes int32 = 15
	this.SessionIdleTimeoutMinutes = &sessionIdleTimeoutMinutes
	var sessionMaxLifetimeMinutes int32 = 720
	this.SessionMaxLifetimeMinutes = &sessionMaxLifetimeMinutes
	return &this
}

// NewAdminConsoleSettingsWithDefaults instantiates a new AdminConsoleSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminConsoleSettingsWithDefaults() *AdminConsoleSettings {
	this := AdminConsoleSettings{}
	var sessionIdleTimeoutMinutes int32 = 15
	this.SessionIdleTimeoutMinutes = &sessionIdleTimeoutMinutes
	var sessionMaxLifetimeMinutes int32 = 720
	this.SessionMaxLifetimeMinutes = &sessionMaxLifetimeMinutes
	return &this
}

// GetSessionIdleTimeoutMinutes returns the SessionIdleTimeoutMinutes field value if set, zero value otherwise.
func (o *AdminConsoleSettings) GetSessionIdleTimeoutMinutes() int32 {
	if o == nil || o.SessionIdleTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.SessionIdleTimeoutMinutes
}

// GetSessionIdleTimeoutMinutesOk returns a tuple with the SessionIdleTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminConsoleSettings) GetSessionIdleTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.SessionIdleTimeoutMinutes == nil {
		return nil, false
	}
	return o.SessionIdleTimeoutMinutes, true
}

// HasSessionIdleTimeoutMinutes returns a boolean if a field has been set.
func (o *AdminConsoleSettings) HasSessionIdleTimeoutMinutes() bool {
	if o != nil && o.SessionIdleTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetSessionIdleTimeoutMinutes gets a reference to the given int32 and assigns it to the SessionIdleTimeoutMinutes field.
func (o *AdminConsoleSettings) SetSessionIdleTimeoutMinutes(v int32) {
	o.SessionIdleTimeoutMinutes = &v
}

// GetSessionMaxLifetimeMinutes returns the SessionMaxLifetimeMinutes field value if set, zero value otherwise.
func (o *AdminConsoleSettings) GetSessionMaxLifetimeMinutes() int32 {
	if o == nil || o.SessionMaxLifetimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.SessionMaxLifetimeMinutes
}

// GetSessionMaxLifetimeMinutesOk returns a tuple with the SessionMaxLifetimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminConsoleSettings) GetSessionMaxLifetimeMinutesOk() (*int32, bool) {
	if o == nil || o.SessionMaxLifetimeMinutes == nil {
		return nil, false
	}
	return o.SessionMaxLifetimeMinutes, true
}

// HasSessionMaxLifetimeMinutes returns a boolean if a field has been set.
func (o *AdminConsoleSettings) HasSessionMaxLifetimeMinutes() bool {
	if o != nil && o.SessionMaxLifetimeMinutes != nil {
		return true
	}

	return false
}

// SetSessionMaxLifetimeMinutes gets a reference to the given int32 and assigns it to the SessionMaxLifetimeMinutes field.
func (o *AdminConsoleSettings) SetSessionMaxLifetimeMinutes(v int32) {
	o.SessionMaxLifetimeMinutes = &v
}

func (o AdminConsoleSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionIdleTimeoutMinutes != nil {
		toSerialize["sessionIdleTimeoutMinutes"] = o.SessionIdleTimeoutMinutes
	}
	if o.SessionMaxLifetimeMinutes != nil {
		toSerialize["sessionMaxLifetimeMinutes"] = o.SessionMaxLifetimeMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdminConsoleSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAdminConsoleSettings := _AdminConsoleSettings{}

	err = json.Unmarshal(bytes, &varAdminConsoleSettings)
	if err == nil {
		*o = AdminConsoleSettings(varAdminConsoleSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "sessionIdleTimeoutMinutes")
		delete(additionalProperties, "sessionMaxLifetimeMinutes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAdminConsoleSettings struct {
	value *AdminConsoleSettings
	isSet bool
}

func (v NullableAdminConsoleSettings) Get() *AdminConsoleSettings {
	return v.value
}

func (v *NullableAdminConsoleSettings) Set(val *AdminConsoleSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminConsoleSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminConsoleSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminConsoleSettings(val *AdminConsoleSettings) *NullableAdminConsoleSettings {
	return &NullableAdminConsoleSettings{value: val, isSet: true}
}

func (v NullableAdminConsoleSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminConsoleSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


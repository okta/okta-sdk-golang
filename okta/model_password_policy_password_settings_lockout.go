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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyPasswordSettingsLockout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyPasswordSettingsLockout{}

// PasswordPolicyPasswordSettingsLockout Lockout settings
type PasswordPolicyPasswordSettingsLockout struct {
	// Specifies the time interval (in minutes) a locked account remains locked before it is automatically unlocked: `0` indicates no limit
	AutoUnlockMinutes *int32 `json:"autoUnlockMinutes,omitempty"`
	// Specifies the number of times Users can attempt to sign in to their accounts with an invalid password before their accounts are locked: `0` indicates no limit
	MaxAttempts *int32 `json:"maxAttempts,omitempty"`
	// Indicates if the User should be informed when their account is locked
	ShowLockoutFailures *bool `json:"showLockoutFailures,omitempty"`
	// How the user is notified when their account becomes locked. The only acceptable values are `[]` and `['EMAIL']`.
	UserLockoutNotificationChannels []string `json:"userLockoutNotificationChannels,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _PasswordPolicyPasswordSettingsLockout PasswordPolicyPasswordSettingsLockout

// NewPasswordPolicyPasswordSettingsLockout instantiates a new PasswordPolicyPasswordSettingsLockout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettingsLockout() *PasswordPolicyPasswordSettingsLockout {
	this := PasswordPolicyPasswordSettingsLockout{}
	var autoUnlockMinutes int32 = 0
	this.AutoUnlockMinutes = &autoUnlockMinutes
	var maxAttempts int32 = 10
	this.MaxAttempts = &maxAttempts
	var showLockoutFailures bool = false
	this.ShowLockoutFailures = &showLockoutFailures
	return &this
}

// NewPasswordPolicyPasswordSettingsLockoutWithDefaults instantiates a new PasswordPolicyPasswordSettingsLockout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsLockoutWithDefaults() *PasswordPolicyPasswordSettingsLockout {
	this := PasswordPolicyPasswordSettingsLockout{}
	var autoUnlockMinutes int32 = 0
	this.AutoUnlockMinutes = &autoUnlockMinutes
	var maxAttempts int32 = 10
	this.MaxAttempts = &maxAttempts
	var showLockoutFailures bool = false
	this.ShowLockoutFailures = &showLockoutFailures
	return &this
}

// GetAutoUnlockMinutes returns the AutoUnlockMinutes field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsLockout) GetAutoUnlockMinutes() int32 {
	if o == nil || IsNil(o.AutoUnlockMinutes) {
		var ret int32
		return ret
	}
	return *o.AutoUnlockMinutes
}

// GetAutoUnlockMinutesOk returns a tuple with the AutoUnlockMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsLockout) GetAutoUnlockMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoUnlockMinutes) {
		return nil, false
	}
	return o.AutoUnlockMinutes, true
}

// HasAutoUnlockMinutes returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsLockout) HasAutoUnlockMinutes() bool {
	if o != nil && !IsNil(o.AutoUnlockMinutes) {
		return true
	}

	return false
}

// SetAutoUnlockMinutes gets a reference to the given int32 and assigns it to the AutoUnlockMinutes field.
func (o *PasswordPolicyPasswordSettingsLockout) SetAutoUnlockMinutes(v int32) {
	o.AutoUnlockMinutes = &v
}

// GetMaxAttempts returns the MaxAttempts field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsLockout) GetMaxAttempts() int32 {
	if o == nil || IsNil(o.MaxAttempts) {
		var ret int32
		return ret
	}
	return *o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsLockout) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAttempts) {
		return nil, false
	}
	return o.MaxAttempts, true
}

// HasMaxAttempts returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsLockout) HasMaxAttempts() bool {
	if o != nil && !IsNil(o.MaxAttempts) {
		return true
	}

	return false
}

// SetMaxAttempts gets a reference to the given int32 and assigns it to the MaxAttempts field.
func (o *PasswordPolicyPasswordSettingsLockout) SetMaxAttempts(v int32) {
	o.MaxAttempts = &v
}

// GetShowLockoutFailures returns the ShowLockoutFailures field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsLockout) GetShowLockoutFailures() bool {
	if o == nil || IsNil(o.ShowLockoutFailures) {
		var ret bool
		return ret
	}
	return *o.ShowLockoutFailures
}

// GetShowLockoutFailuresOk returns a tuple with the ShowLockoutFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsLockout) GetShowLockoutFailuresOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowLockoutFailures) {
		return nil, false
	}
	return o.ShowLockoutFailures, true
}

// HasShowLockoutFailures returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsLockout) HasShowLockoutFailures() bool {
	if o != nil && !IsNil(o.ShowLockoutFailures) {
		return true
	}

	return false
}

// SetShowLockoutFailures gets a reference to the given bool and assigns it to the ShowLockoutFailures field.
func (o *PasswordPolicyPasswordSettingsLockout) SetShowLockoutFailures(v bool) {
	o.ShowLockoutFailures = &v
}

// GetUserLockoutNotificationChannels returns the UserLockoutNotificationChannels field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsLockout) GetUserLockoutNotificationChannels() []string {
	if o == nil || IsNil(o.UserLockoutNotificationChannels) {
		var ret []string
		return ret
	}
	return o.UserLockoutNotificationChannels
}

// GetUserLockoutNotificationChannelsOk returns a tuple with the UserLockoutNotificationChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsLockout) GetUserLockoutNotificationChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserLockoutNotificationChannels) {
		return nil, false
	}
	return o.UserLockoutNotificationChannels, true
}

// HasUserLockoutNotificationChannels returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsLockout) HasUserLockoutNotificationChannels() bool {
	if o != nil && !IsNil(o.UserLockoutNotificationChannels) {
		return true
	}

	return false
}

// SetUserLockoutNotificationChannels gets a reference to the given []string and assigns it to the UserLockoutNotificationChannels field.
func (o *PasswordPolicyPasswordSettingsLockout) SetUserLockoutNotificationChannels(v []string) {
	o.UserLockoutNotificationChannels = v
}

func (o PasswordPolicyPasswordSettingsLockout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyPasswordSettingsLockout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUnlockMinutes) {
		toSerialize["autoUnlockMinutes"] = o.AutoUnlockMinutes
	}
	if !IsNil(o.MaxAttempts) {
		toSerialize["maxAttempts"] = o.MaxAttempts
	}
	if !IsNil(o.ShowLockoutFailures) {
		toSerialize["showLockoutFailures"] = o.ShowLockoutFailures
	}
	if !IsNil(o.UserLockoutNotificationChannels) {
		toSerialize["userLockoutNotificationChannels"] = o.UserLockoutNotificationChannels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyPasswordSettingsLockout) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyPasswordSettingsLockout := _PasswordPolicyPasswordSettingsLockout{}

	err = json.Unmarshal(data, &varPasswordPolicyPasswordSettingsLockout)

	if err != nil {
		return err
	}

	*o = PasswordPolicyPasswordSettingsLockout(varPasswordPolicyPasswordSettingsLockout)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoUnlockMinutes")
		delete(additionalProperties, "maxAttempts")
		delete(additionalProperties, "showLockoutFailures")
		delete(additionalProperties, "userLockoutNotificationChannels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyPasswordSettingsLockout struct {
	value *PasswordPolicyPasswordSettingsLockout
	isSet bool
}

func (v NullablePasswordPolicyPasswordSettingsLockout) Get() *PasswordPolicyPasswordSettingsLockout {
	return v.value
}

func (v *NullablePasswordPolicyPasswordSettingsLockout) Set(val *PasswordPolicyPasswordSettingsLockout) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyPasswordSettingsLockout) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyPasswordSettingsLockout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyPasswordSettingsLockout(val *PasswordPolicyPasswordSettingsLockout) *NullablePasswordPolicyPasswordSettingsLockout {
	return &NullablePasswordPolicyPasswordSettingsLockout{value: val, isSet: true}
}

func (v NullablePasswordPolicyPasswordSettingsLockout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyPasswordSettingsLockout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

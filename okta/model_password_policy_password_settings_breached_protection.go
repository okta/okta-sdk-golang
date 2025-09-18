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

// checks if the PasswordPolicyPasswordSettingsBreachedProtection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyPasswordSettingsBreachedProtection{}

// PasswordPolicyPasswordSettingsBreachedProtection Breached Protection settings
type PasswordPolicyPasswordSettingsBreachedProtection struct {
	// The `id` of the workflow that runs when a breached password is found during a sign-in attempt.
	DelegatedWorkflowId NullableString `json:"delegatedWorkflowId,omitempty"`
	// Specifies the number of days after a breached password is found during a sign-in attempt that the user's password should expire. Valid values: 0 through 10. If set to 0, it happens immediately.
	ExpireAfterDays NullableInt32 `json:"expireAfterDays,omitempty"`
	// (Optional, default is false) If true, you must also specify a value for `expireAfterDays`. When enabled, the user's session(s) are terminated immediately the first time the user's credentials are detected as part of a breach.
	LogoutEnabled        NullableBool `json:"logoutEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyPasswordSettingsBreachedProtection PasswordPolicyPasswordSettingsBreachedProtection

// NewPasswordPolicyPasswordSettingsBreachedProtection instantiates a new PasswordPolicyPasswordSettingsBreachedProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettingsBreachedProtection() *PasswordPolicyPasswordSettingsBreachedProtection {
	this := PasswordPolicyPasswordSettingsBreachedProtection{}
	var logoutEnabled bool = false
	this.LogoutEnabled = *NewNullableBool(&logoutEnabled)
	return &this
}

// NewPasswordPolicyPasswordSettingsBreachedProtectionWithDefaults instantiates a new PasswordPolicyPasswordSettingsBreachedProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsBreachedProtectionWithDefaults() *PasswordPolicyPasswordSettingsBreachedProtection {
	this := PasswordPolicyPasswordSettingsBreachedProtection{}
	var logoutEnabled bool = false
	this.LogoutEnabled = *NewNullableBool(&logoutEnabled)
	return &this
}

// GetDelegatedWorkflowId returns the DelegatedWorkflowId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetDelegatedWorkflowId() string {
	if o == nil || IsNil(o.DelegatedWorkflowId.Get()) {
		var ret string
		return ret
	}
	return *o.DelegatedWorkflowId.Get()
}

// GetDelegatedWorkflowIdOk returns a tuple with the DelegatedWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetDelegatedWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelegatedWorkflowId.Get(), o.DelegatedWorkflowId.IsSet()
}

// HasDelegatedWorkflowId returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasDelegatedWorkflowId() bool {
	if o != nil && o.DelegatedWorkflowId.IsSet() {
		return true
	}

	return false
}

// SetDelegatedWorkflowId gets a reference to the given NullableString and assigns it to the DelegatedWorkflowId field.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetDelegatedWorkflowId(v string) {
	o.DelegatedWorkflowId.Set(&v)
}

// SetDelegatedWorkflowIdNil sets the value for DelegatedWorkflowId to be an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetDelegatedWorkflowIdNil() {
	o.DelegatedWorkflowId.Set(nil)
}

// UnsetDelegatedWorkflowId ensures that no value is present for DelegatedWorkflowId, not even an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetDelegatedWorkflowId() {
	o.DelegatedWorkflowId.Unset()
}

// GetExpireAfterDays returns the ExpireAfterDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetExpireAfterDays() int32 {
	if o == nil || IsNil(o.ExpireAfterDays.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpireAfterDays.Get()
}

// GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetExpireAfterDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireAfterDays.Get(), o.ExpireAfterDays.IsSet()
}

// HasExpireAfterDays returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasExpireAfterDays() bool {
	if o != nil && o.ExpireAfterDays.IsSet() {
		return true
	}

	return false
}

// SetExpireAfterDays gets a reference to the given NullableInt32 and assigns it to the ExpireAfterDays field.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetExpireAfterDays(v int32) {
	o.ExpireAfterDays.Set(&v)
}

// SetExpireAfterDaysNil sets the value for ExpireAfterDays to be an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetExpireAfterDaysNil() {
	o.ExpireAfterDays.Set(nil)
}

// UnsetExpireAfterDays ensures that no value is present for ExpireAfterDays, not even an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetExpireAfterDays() {
	o.ExpireAfterDays.Unset()
}

// GetLogoutEnabled returns the LogoutEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetLogoutEnabled() bool {
	if o == nil || IsNil(o.LogoutEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.LogoutEnabled.Get()
}

// GetLogoutEnabledOk returns a tuple with the LogoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetLogoutEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoutEnabled.Get(), o.LogoutEnabled.IsSet()
}

// HasLogoutEnabled returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasLogoutEnabled() bool {
	if o != nil && o.LogoutEnabled.IsSet() {
		return true
	}

	return false
}

// SetLogoutEnabled gets a reference to the given NullableBool and assigns it to the LogoutEnabled field.
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetLogoutEnabled(v bool) {
	o.LogoutEnabled.Set(&v)
}

// SetLogoutEnabledNil sets the value for LogoutEnabled to be an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetLogoutEnabledNil() {
	o.LogoutEnabled.Set(nil)
}

// UnsetLogoutEnabled ensures that no value is present for LogoutEnabled, not even an explicit nil
func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetLogoutEnabled() {
	o.LogoutEnabled.Unset()
}

func (o PasswordPolicyPasswordSettingsBreachedProtection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyPasswordSettingsBreachedProtection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegatedWorkflowId.IsSet() {
		toSerialize["delegatedWorkflowId"] = o.DelegatedWorkflowId.Get()
	}
	if o.ExpireAfterDays.IsSet() {
		toSerialize["expireAfterDays"] = o.ExpireAfterDays.Get()
	}
	if o.LogoutEnabled.IsSet() {
		toSerialize["logoutEnabled"] = o.LogoutEnabled.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyPasswordSettingsBreachedProtection := _PasswordPolicyPasswordSettingsBreachedProtection{}

	err = json.Unmarshal(data, &varPasswordPolicyPasswordSettingsBreachedProtection)

	if err != nil {
		return err
	}

	*o = PasswordPolicyPasswordSettingsBreachedProtection(varPasswordPolicyPasswordSettingsBreachedProtection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegatedWorkflowId")
		delete(additionalProperties, "expireAfterDays")
		delete(additionalProperties, "logoutEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyPasswordSettingsBreachedProtection struct {
	value *PasswordPolicyPasswordSettingsBreachedProtection
	isSet bool
}

func (v NullablePasswordPolicyPasswordSettingsBreachedProtection) Get() *PasswordPolicyPasswordSettingsBreachedProtection {
	return v.value
}

func (v *NullablePasswordPolicyPasswordSettingsBreachedProtection) Set(val *PasswordPolicyPasswordSettingsBreachedProtection) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyPasswordSettingsBreachedProtection) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyPasswordSettingsBreachedProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyPasswordSettingsBreachedProtection(val *PasswordPolicyPasswordSettingsBreachedProtection) *NullablePasswordPolicyPasswordSettingsBreachedProtection {
	return &NullablePasswordPolicyPasswordSettingsBreachedProtection{value: val, isSet: true}
}

func (v NullablePasswordPolicyPasswordSettingsBreachedProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyPasswordSettingsBreachedProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

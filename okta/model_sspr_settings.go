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

// SsprSettings <x-lifecycle class=\"oie\"></x-lifecycle> Describes other behavioral settings that apply to the self-service password reset (SSPR) process
type SsprSettings struct {
	// <x-lifecycle class=\"oie\"></x-lifecycle> Controls whether a user can receive a password recovery email at their profile email address even if they haven't enrolled email as an authenticator.  * `true`: Email appears as a recovery option during password reset and account unlock for users without an enrolled email authenticator. The recovery OTP is sent to the email address on the user's profile. * `false` or `null`: Email recovery requires an enrolled email authenticator. This is the default behavior for Identity Engine orgs.  For orgs that have migrated from Classic Engine to Identity Engine, `allowRecoveryEmailWithoutEnrollment` is set to `true`. It's set to `true` by the migration action to preserve Classic Engine behavior where recovery emails are always sent to the profile email address regardless of the user's enrollment state.  > **Note:** This property is only available if the Email auto-enrollment and recovery control [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled.
	AllowRecoveryEmailWithoutEnrollment NullableBool `json:"allowRecoveryEmailWithoutEnrollment,omitempty"`
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

// GetAllowRecoveryEmailWithoutEnrollment returns the AllowRecoveryEmailWithoutEnrollment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollment() bool {
	if o == nil || IsNil(o.AllowRecoveryEmailWithoutEnrollment.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowRecoveryEmailWithoutEnrollment.Get()
}

// GetAllowRecoveryEmailWithoutEnrollmentOk returns a tuple with the AllowRecoveryEmailWithoutEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowRecoveryEmailWithoutEnrollment.Get(), o.AllowRecoveryEmailWithoutEnrollment.IsSet()
}

// HasAllowRecoveryEmailWithoutEnrollment returns a boolean if a field has been set.
func (o *SsprSettings) HasAllowRecoveryEmailWithoutEnrollment() bool {
	if o != nil && o.AllowRecoveryEmailWithoutEnrollment.IsSet() {
		return true
	}

	return false
}

// SetAllowRecoveryEmailWithoutEnrollment gets a reference to the given NullableBool and assigns it to the AllowRecoveryEmailWithoutEnrollment field.
func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollment(v bool) {
	o.AllowRecoveryEmailWithoutEnrollment.Set(&v)
}

// SetAllowRecoveryEmailWithoutEnrollmentNil sets the value for AllowRecoveryEmailWithoutEnrollment to be an explicit nil
func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollmentNil() {
	o.AllowRecoveryEmailWithoutEnrollment.Set(nil)
}

// UnsetAllowRecoveryEmailWithoutEnrollment ensures that no value is present for AllowRecoveryEmailWithoutEnrollment, not even an explicit nil
func (o *SsprSettings) UnsetAllowRecoveryEmailWithoutEnrollment() {
	o.AllowRecoveryEmailWithoutEnrollment.Unset()
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
	if o.AllowRecoveryEmailWithoutEnrollment.IsSet() {
		toSerialize["allowRecoveryEmailWithoutEnrollment"] = o.AllowRecoveryEmailWithoutEnrollment.Get()
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

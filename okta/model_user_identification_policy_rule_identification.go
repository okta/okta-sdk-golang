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

// checks if the UserIdentificationPolicyRuleIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationPolicyRuleIdentification{}

// UserIdentificationPolicyRuleIdentification <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies how user experience is identified when signing in
type UserIdentificationPolicyRuleIdentification struct {
	Settings             *UserIdentificationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationPolicyRuleIdentification UserIdentificationPolicyRuleIdentification

// NewUserIdentificationPolicyRuleIdentification instantiates a new UserIdentificationPolicyRuleIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationPolicyRuleIdentification() *UserIdentificationPolicyRuleIdentification {
	this := UserIdentificationPolicyRuleIdentification{}
	return &this
}

// NewUserIdentificationPolicyRuleIdentificationWithDefaults instantiates a new UserIdentificationPolicyRuleIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationPolicyRuleIdentificationWithDefaults() *UserIdentificationPolicyRuleIdentification {
	this := UserIdentificationPolicyRuleIdentification{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *UserIdentificationPolicyRuleIdentification) GetSettings() UserIdentificationSettings {
	if o == nil || IsNil(o.Settings) {
		var ret UserIdentificationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationPolicyRuleIdentification) GetSettingsOk() (*UserIdentificationSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UserIdentificationPolicyRuleIdentification) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given UserIdentificationSettings and assigns it to the Settings field.
func (o *UserIdentificationPolicyRuleIdentification) SetSettings(v UserIdentificationSettings) {
	o.Settings = &v
}

func (o UserIdentificationPolicyRuleIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationPolicyRuleIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationPolicyRuleIdentification) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationPolicyRuleIdentification := _UserIdentificationPolicyRuleIdentification{}

	err = json.Unmarshal(data, &varUserIdentificationPolicyRuleIdentification)

	if err != nil {
		return err
	}

	*o = UserIdentificationPolicyRuleIdentification(varUserIdentificationPolicyRuleIdentification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationPolicyRuleIdentification struct {
	value *UserIdentificationPolicyRuleIdentification
	isSet bool
}

func (v NullableUserIdentificationPolicyRuleIdentification) Get() *UserIdentificationPolicyRuleIdentification {
	return v.value
}

func (v *NullableUserIdentificationPolicyRuleIdentification) Set(val *UserIdentificationPolicyRuleIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationPolicyRuleIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationPolicyRuleIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationPolicyRuleIdentification(val *UserIdentificationPolicyRuleIdentification) *NullableUserIdentificationPolicyRuleIdentification {
	return &NullableUserIdentificationPolicyRuleIdentification{value: val, isSet: true}
}

func (v NullableUserIdentificationPolicyRuleIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationPolicyRuleIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// PasswordPolicyDelegationSettings struct for PasswordPolicyDelegationSettings
type PasswordPolicyDelegationSettings struct {
	Options *PasswordPolicyDelegationSettingsOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyDelegationSettings PasswordPolicyDelegationSettings

// NewPasswordPolicyDelegationSettings instantiates a new PasswordPolicyDelegationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyDelegationSettings() *PasswordPolicyDelegationSettings {
	this := PasswordPolicyDelegationSettings{}
	return &this
}

// NewPasswordPolicyDelegationSettingsWithDefaults instantiates a new PasswordPolicyDelegationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyDelegationSettingsWithDefaults() *PasswordPolicyDelegationSettings {
	this := PasswordPolicyDelegationSettings{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PasswordPolicyDelegationSettings) GetOptions() PasswordPolicyDelegationSettingsOptions {
	if o == nil || o.Options == nil {
		var ret PasswordPolicyDelegationSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyDelegationSettings) GetOptionsOk() (*PasswordPolicyDelegationSettingsOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PasswordPolicyDelegationSettings) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given PasswordPolicyDelegationSettingsOptions and assigns it to the Options field.
func (o *PasswordPolicyDelegationSettings) SetOptions(v PasswordPolicyDelegationSettingsOptions) {
	o.Options = &v
}

func (o PasswordPolicyDelegationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyDelegationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyDelegationSettings := _PasswordPolicyDelegationSettings{}

	err = json.Unmarshal(bytes, &varPasswordPolicyDelegationSettings)
	if err == nil {
		*o = PasswordPolicyDelegationSettings(varPasswordPolicyDelegationSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyDelegationSettings struct {
	value *PasswordPolicyDelegationSettings
	isSet bool
}

func (v NullablePasswordPolicyDelegationSettings) Get() *PasswordPolicyDelegationSettings {
	return v.value
}

func (v *NullablePasswordPolicyDelegationSettings) Set(val *PasswordPolicyDelegationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyDelegationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyDelegationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyDelegationSettings(val *PasswordPolicyDelegationSettings) *NullablePasswordPolicyDelegationSettings {
	return &NullablePasswordPolicyDelegationSettings{value: val, isSet: true}
}

func (v NullablePasswordPolicyDelegationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyDelegationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


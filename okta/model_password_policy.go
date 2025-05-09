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
	"reflect"
	"strings"
)

// PasswordPolicy struct for PasswordPolicy
type PasswordPolicy struct {
	Policy
	Conditions *PasswordPolicyConditions `json:"conditions,omitempty"`
	Settings *PasswordPolicySettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicy PasswordPolicy

// NewPasswordPolicy instantiates a new PasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicy() *PasswordPolicy {
	this := PasswordPolicy{}
	return &this
}

// NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyWithDefaults() *PasswordPolicy {
	this := PasswordPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PasswordPolicy) GetConditions() PasswordPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret PasswordPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetConditionsOk() (*PasswordPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PasswordPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PasswordPolicyConditions and assigns it to the Conditions field.
func (o *PasswordPolicy) SetConditions(v PasswordPolicyConditions) {
	o.Conditions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *PasswordPolicy) GetSettings() PasswordPolicySettings {
	if o == nil || o.Settings == nil {
		var ret PasswordPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetSettingsOk() (*PasswordPolicySettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PasswordPolicy) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given PasswordPolicySettings and assigns it to the Settings field.
func (o *PasswordPolicy) SetSettings(v PasswordPolicySettings) {
	o.Settings = &v
}

func (o PasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type PasswordPolicyWithoutEmbeddedStruct struct {
		Conditions *PasswordPolicyConditions `json:"conditions,omitempty"`
		Settings *PasswordPolicySettings `json:"settings,omitempty"`
	}

	varPasswordPolicyWithoutEmbeddedStruct := PasswordPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPasswordPolicyWithoutEmbeddedStruct)
	if err == nil {
		varPasswordPolicy := _PasswordPolicy{}
		varPasswordPolicy.Conditions = varPasswordPolicyWithoutEmbeddedStruct.Conditions
		varPasswordPolicy.Settings = varPasswordPolicyWithoutEmbeddedStruct.Settings
		*o = PasswordPolicy(varPasswordPolicy)
	} else {
		return err
	}

	varPasswordPolicy := _PasswordPolicy{}

	err = json.Unmarshal(bytes, &varPasswordPolicy)
	if err == nil {
		o.Policy = varPasswordPolicy.Policy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectPolicy := reflect.ValueOf(o.Policy)
		for i := 0; i < reflectPolicy.Type().NumField(); i++ {
			t := reflectPolicy.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicy struct {
	value *PasswordPolicy
	isSet bool
}

func (v NullablePasswordPolicy) Get() *PasswordPolicy {
	return v.value
}

func (v *NullablePasswordPolicy) Set(val *PasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicy(val *PasswordPolicy) *NullablePasswordPolicy {
	return &NullablePasswordPolicy{value: val, isSet: true}
}

func (v NullablePasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


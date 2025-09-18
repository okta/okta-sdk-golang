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
	"fmt"
	"reflect"
	"strings"
)

// checks if the PasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicy{}

// PasswordPolicy struct for PasswordPolicy
type PasswordPolicy struct {
	Policy
	Conditions           *PasswordPolicyConditions `json:"conditions,omitempty"`
	Settings             *PasswordPolicySettings   `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicy PasswordPolicy

// NewPasswordPolicy instantiates a new PasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicy(name string, type_ string) *PasswordPolicy {
	this := PasswordPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
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
	if o == nil || IsNil(o.Conditions) {
		var ret PasswordPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetConditionsOk() (*PasswordPolicyConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PasswordPolicy) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
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
	if o == nil || IsNil(o.Settings) {
		var ret PasswordPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetSettingsOk() (*PasswordPolicySettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PasswordPolicy) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given PasswordPolicySettings and assigns it to the Settings field.
func (o *PasswordPolicy) SetSettings(v PasswordPolicySettings) {
	o.Settings = &v
}

func (o PasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type PasswordPolicyWithoutEmbeddedStruct struct {
		Conditions *PasswordPolicyConditions `json:"conditions,omitempty"`
		Settings   *PasswordPolicySettings   `json:"settings,omitempty"`
	}

	varPasswordPolicyWithoutEmbeddedStruct := PasswordPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPasswordPolicyWithoutEmbeddedStruct)
	if err == nil {
		varPasswordPolicy := _PasswordPolicy{}
		varPasswordPolicy.Conditions = varPasswordPolicyWithoutEmbeddedStruct.Conditions
		varPasswordPolicy.Settings = varPasswordPolicyWithoutEmbeddedStruct.Settings
		*o = PasswordPolicy(varPasswordPolicy)
	} else {
		return err
	}

	varPasswordPolicy := _PasswordPolicy{}

	err = json.Unmarshal(data, &varPasswordPolicy)
	if err == nil {
		o.Policy = varPasswordPolicy.Policy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

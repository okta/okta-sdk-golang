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

// MultifactorEnrollmentPolicy struct for MultifactorEnrollmentPolicy
type MultifactorEnrollmentPolicy struct {
	Policy
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	Settings *MultifactorEnrollmentPolicySettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicy MultifactorEnrollmentPolicy

// NewMultifactorEnrollmentPolicy instantiates a new MultifactorEnrollmentPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicy() *MultifactorEnrollmentPolicy {
	this := MultifactorEnrollmentPolicy{}
	return &this
}

// NewMultifactorEnrollmentPolicyWithDefaults instantiates a new MultifactorEnrollmentPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyWithDefaults() *MultifactorEnrollmentPolicy {
	this := MultifactorEnrollmentPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicy) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicy) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *MultifactorEnrollmentPolicy) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicy) GetSettings() MultifactorEnrollmentPolicySettings {
	if o == nil || o.Settings == nil {
		var ret MultifactorEnrollmentPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicy) GetSettingsOk() (*MultifactorEnrollmentPolicySettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicy) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given MultifactorEnrollmentPolicySettings and assigns it to the Settings field.
func (o *MultifactorEnrollmentPolicy) SetSettings(v MultifactorEnrollmentPolicySettings) {
	o.Settings = &v
}

func (o MultifactorEnrollmentPolicy) MarshalJSON() ([]byte, error) {
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

func (o *MultifactorEnrollmentPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type MultifactorEnrollmentPolicyWithoutEmbeddedStruct struct {
		Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
		Settings *MultifactorEnrollmentPolicySettings `json:"settings,omitempty"`
	}

	varMultifactorEnrollmentPolicyWithoutEmbeddedStruct := MultifactorEnrollmentPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyWithoutEmbeddedStruct)
	if err == nil {
		varMultifactorEnrollmentPolicy := _MultifactorEnrollmentPolicy{}
		varMultifactorEnrollmentPolicy.Conditions = varMultifactorEnrollmentPolicyWithoutEmbeddedStruct.Conditions
		varMultifactorEnrollmentPolicy.Settings = varMultifactorEnrollmentPolicyWithoutEmbeddedStruct.Settings
		*o = MultifactorEnrollmentPolicy(varMultifactorEnrollmentPolicy)
	} else {
		return err
	}

	varMultifactorEnrollmentPolicy := _MultifactorEnrollmentPolicy{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicy)
	if err == nil {
		o.Policy = varMultifactorEnrollmentPolicy.Policy
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

type NullableMultifactorEnrollmentPolicy struct {
	value *MultifactorEnrollmentPolicy
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicy) Get() *MultifactorEnrollmentPolicy {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicy) Set(val *MultifactorEnrollmentPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicy(val *MultifactorEnrollmentPolicy) *NullableMultifactorEnrollmentPolicy {
	return &NullableMultifactorEnrollmentPolicy{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


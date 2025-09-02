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

// AuthenticatorEnrollmentPolicy struct for AuthenticatorEnrollmentPolicy
type AuthenticatorEnrollmentPolicy struct {
	Policy
	Conditions *AuthenticatorEnrollmentPolicyConditions `json:"conditions,omitempty"`
	Settings *AuthenticatorEnrollmentPolicySettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicy AuthenticatorEnrollmentPolicy

// NewAuthenticatorEnrollmentPolicy instantiates a new AuthenticatorEnrollmentPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicy(name string, type_ string) *AuthenticatorEnrollmentPolicy {
	this := AuthenticatorEnrollmentPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
	return &this
}

// NewAuthenticatorEnrollmentPolicyWithDefaults instantiates a new AuthenticatorEnrollmentPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyWithDefaults() *AuthenticatorEnrollmentPolicy {
	this := AuthenticatorEnrollmentPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicy) GetConditions() AuthenticatorEnrollmentPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret AuthenticatorEnrollmentPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicy) GetConditionsOk() (*AuthenticatorEnrollmentPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthenticatorEnrollmentPolicyConditions and assigns it to the Conditions field.
func (o *AuthenticatorEnrollmentPolicy) SetConditions(v AuthenticatorEnrollmentPolicyConditions) {
	o.Conditions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicy) GetSettings() AuthenticatorEnrollmentPolicySettings {
	if o == nil || o.Settings == nil {
		var ret AuthenticatorEnrollmentPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicy) GetSettingsOk() (*AuthenticatorEnrollmentPolicySettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicy) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AuthenticatorEnrollmentPolicySettings and assigns it to the Settings field.
func (o *AuthenticatorEnrollmentPolicy) SetSettings(v AuthenticatorEnrollmentPolicySettings) {
	o.Settings = &v
}

func (o AuthenticatorEnrollmentPolicy) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorEnrollmentPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorEnrollmentPolicyWithoutEmbeddedStruct struct {
		Conditions *AuthenticatorEnrollmentPolicyConditions `json:"conditions,omitempty"`
		Settings *AuthenticatorEnrollmentPolicySettings `json:"settings,omitempty"`
	}

	varAuthenticatorEnrollmentPolicyWithoutEmbeddedStruct := AuthenticatorEnrollmentPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicyWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorEnrollmentPolicy := _AuthenticatorEnrollmentPolicy{}
		varAuthenticatorEnrollmentPolicy.Conditions = varAuthenticatorEnrollmentPolicyWithoutEmbeddedStruct.Conditions
		varAuthenticatorEnrollmentPolicy.Settings = varAuthenticatorEnrollmentPolicyWithoutEmbeddedStruct.Settings
		*o = AuthenticatorEnrollmentPolicy(varAuthenticatorEnrollmentPolicy)
	} else {
		return err
	}

	varAuthenticatorEnrollmentPolicy := _AuthenticatorEnrollmentPolicy{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicy)
	if err == nil {
		o.Policy = varAuthenticatorEnrollmentPolicy.Policy
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

type NullableAuthenticatorEnrollmentPolicy struct {
	value *AuthenticatorEnrollmentPolicy
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicy) Get() *AuthenticatorEnrollmentPolicy {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicy) Set(val *AuthenticatorEnrollmentPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicy(val *AuthenticatorEnrollmentPolicy) *NullableAuthenticatorEnrollmentPolicy {
	return &NullableAuthenticatorEnrollmentPolicy{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


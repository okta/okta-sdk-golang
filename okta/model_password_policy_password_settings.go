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

// PasswordPolicyPasswordSettings struct for PasswordPolicyPasswordSettings
type PasswordPolicyPasswordSettings struct {
	Age *PasswordPolicyPasswordSettingsAge `json:"age,omitempty"`
	Complexity *PasswordPolicyPasswordSettingsComplexity `json:"complexity,omitempty"`
	Lockout *PasswordPolicyPasswordSettingsLockout `json:"lockout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyPasswordSettings PasswordPolicyPasswordSettings

// NewPasswordPolicyPasswordSettings instantiates a new PasswordPolicyPasswordSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettings() *PasswordPolicyPasswordSettings {
	this := PasswordPolicyPasswordSettings{}
	return &this
}

// NewPasswordPolicyPasswordSettingsWithDefaults instantiates a new PasswordPolicyPasswordSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsWithDefaults() *PasswordPolicyPasswordSettings {
	this := PasswordPolicyPasswordSettings{}
	return &this
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettings) GetAge() PasswordPolicyPasswordSettingsAge {
	if o == nil || o.Age == nil {
		var ret PasswordPolicyPasswordSettingsAge
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettings) GetAgeOk() (*PasswordPolicyPasswordSettingsAge, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettings) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given PasswordPolicyPasswordSettingsAge and assigns it to the Age field.
func (o *PasswordPolicyPasswordSettings) SetAge(v PasswordPolicyPasswordSettingsAge) {
	o.Age = &v
}

// GetComplexity returns the Complexity field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettings) GetComplexity() PasswordPolicyPasswordSettingsComplexity {
	if o == nil || o.Complexity == nil {
		var ret PasswordPolicyPasswordSettingsComplexity
		return ret
	}
	return *o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettings) GetComplexityOk() (*PasswordPolicyPasswordSettingsComplexity, bool) {
	if o == nil || o.Complexity == nil {
		return nil, false
	}
	return o.Complexity, true
}

// HasComplexity returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettings) HasComplexity() bool {
	if o != nil && o.Complexity != nil {
		return true
	}

	return false
}

// SetComplexity gets a reference to the given PasswordPolicyPasswordSettingsComplexity and assigns it to the Complexity field.
func (o *PasswordPolicyPasswordSettings) SetComplexity(v PasswordPolicyPasswordSettingsComplexity) {
	o.Complexity = &v
}

// GetLockout returns the Lockout field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettings) GetLockout() PasswordPolicyPasswordSettingsLockout {
	if o == nil || o.Lockout == nil {
		var ret PasswordPolicyPasswordSettingsLockout
		return ret
	}
	return *o.Lockout
}

// GetLockoutOk returns a tuple with the Lockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettings) GetLockoutOk() (*PasswordPolicyPasswordSettingsLockout, bool) {
	if o == nil || o.Lockout == nil {
		return nil, false
	}
	return o.Lockout, true
}

// HasLockout returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettings) HasLockout() bool {
	if o != nil && o.Lockout != nil {
		return true
	}

	return false
}

// SetLockout gets a reference to the given PasswordPolicyPasswordSettingsLockout and assigns it to the Lockout field.
func (o *PasswordPolicyPasswordSettings) SetLockout(v PasswordPolicyPasswordSettingsLockout) {
	o.Lockout = &v
}

func (o PasswordPolicyPasswordSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Age != nil {
		toSerialize["age"] = o.Age
	}
	if o.Complexity != nil {
		toSerialize["complexity"] = o.Complexity
	}
	if o.Lockout != nil {
		toSerialize["lockout"] = o.Lockout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyPasswordSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyPasswordSettings := _PasswordPolicyPasswordSettings{}

	err = json.Unmarshal(bytes, &varPasswordPolicyPasswordSettings)
	if err == nil {
		*o = PasswordPolicyPasswordSettings(varPasswordPolicyPasswordSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "age")
		delete(additionalProperties, "complexity")
		delete(additionalProperties, "lockout")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyPasswordSettings struct {
	value *PasswordPolicyPasswordSettings
	isSet bool
}

func (v NullablePasswordPolicyPasswordSettings) Get() *PasswordPolicyPasswordSettings {
	return v.value
}

func (v *NullablePasswordPolicyPasswordSettings) Set(val *PasswordPolicyPasswordSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyPasswordSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyPasswordSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyPasswordSettings(val *PasswordPolicyPasswordSettings) *NullablePasswordPolicyPasswordSettings {
	return &NullablePasswordPolicyPasswordSettings{value: val, isSet: true}
}

func (v NullablePasswordPolicyPasswordSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyPasswordSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


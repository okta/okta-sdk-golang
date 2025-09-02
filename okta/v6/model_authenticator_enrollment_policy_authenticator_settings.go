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
)

// AuthenticatorEnrollmentPolicyAuthenticatorSettings struct for AuthenticatorEnrollmentPolicyAuthenticatorSettings
type AuthenticatorEnrollmentPolicyAuthenticatorSettings struct {
	Constraints NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints `json:"constraints,omitempty"`
	Enroll *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll `json:"enroll,omitempty"`
	// A label that identifies the authenticator
	Key *string `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyAuthenticatorSettings AuthenticatorEnrollmentPolicyAuthenticatorSettings

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettings instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettings() *AuthenticatorEnrollmentPolicyAuthenticatorSettings {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettings {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettings{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetConstraints() AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetConstraintsOk() (*AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints and assigns it to the Constraints field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetConstraints(v AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) {
	o.Constraints.Set(&v)
}
// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetEnroll() AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll {
	if o == nil || o.Enroll == nil {
		var ret AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetEnrollOk() (*AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll and assigns it to the Enroll field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetEnroll(v AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) {
	o.Enroll = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetKey(v string) {
	o.Key = &v
}

func (o AuthenticatorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorEnrollmentPolicyAuthenticatorSettings := _AuthenticatorEnrollmentPolicyAuthenticatorSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorEnrollmentPolicyAuthenticatorSettings)
	if err == nil {
		*o = AuthenticatorEnrollmentPolicyAuthenticatorSettings(varAuthenticatorEnrollmentPolicyAuthenticatorSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings struct {
	value *AuthenticatorEnrollmentPolicyAuthenticatorSettings
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) Get() *AuthenticatorEnrollmentPolicyAuthenticatorSettings {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) Set(val *AuthenticatorEnrollmentPolicyAuthenticatorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyAuthenticatorSettings(val *AuthenticatorEnrollmentPolicyAuthenticatorSettings) *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings {
	return &NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


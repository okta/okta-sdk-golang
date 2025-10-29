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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyRecoveryFactors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoveryFactors{}

// PasswordPolicyRecoveryFactors Settings for the factors that can be used for recovery
type PasswordPolicyRecoveryFactors struct {
	OktaCall             *PasswordPolicyRecoveryFactorSettings `json:"okta_call,omitempty"`
	OktaEmail            *PasswordPolicyRecoveryEmail          `json:"okta_email,omitempty"`
	OktaSms              *PasswordPolicyRecoveryFactorSettings `json:"okta_sms,omitempty"`
	RecoveryQuestion     *PasswordPolicyRecoveryQuestion       `json:"recovery_question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryFactors PasswordPolicyRecoveryFactors

// NewPasswordPolicyRecoveryFactors instantiates a new PasswordPolicyRecoveryFactors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryFactors() *PasswordPolicyRecoveryFactors {
	this := PasswordPolicyRecoveryFactors{}
	return &this
}

// NewPasswordPolicyRecoveryFactorsWithDefaults instantiates a new PasswordPolicyRecoveryFactors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryFactorsWithDefaults() *PasswordPolicyRecoveryFactors {
	this := PasswordPolicyRecoveryFactors{}
	return &this
}

// GetOktaCall returns the OktaCall field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryFactors) GetOktaCall() PasswordPolicyRecoveryFactorSettings {
	if o == nil || IsNil(o.OktaCall) {
		var ret PasswordPolicyRecoveryFactorSettings
		return ret
	}
	return *o.OktaCall
}

// GetOktaCallOk returns a tuple with the OktaCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaCallOk() (*PasswordPolicyRecoveryFactorSettings, bool) {
	if o == nil || IsNil(o.OktaCall) {
		return nil, false
	}
	return o.OktaCall, true
}

// HasOktaCall returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaCall() bool {
	if o != nil && !IsNil(o.OktaCall) {
		return true
	}

	return false
}

// SetOktaCall gets a reference to the given PasswordPolicyRecoveryFactorSettings and assigns it to the OktaCall field.
func (o *PasswordPolicyRecoveryFactors) SetOktaCall(v PasswordPolicyRecoveryFactorSettings) {
	o.OktaCall = &v
}

// GetOktaEmail returns the OktaEmail field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryFactors) GetOktaEmail() PasswordPolicyRecoveryEmail {
	if o == nil || IsNil(o.OktaEmail) {
		var ret PasswordPolicyRecoveryEmail
		return ret
	}
	return *o.OktaEmail
}

// GetOktaEmailOk returns a tuple with the OktaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaEmailOk() (*PasswordPolicyRecoveryEmail, bool) {
	if o == nil || IsNil(o.OktaEmail) {
		return nil, false
	}
	return o.OktaEmail, true
}

// HasOktaEmail returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaEmail() bool {
	if o != nil && !IsNil(o.OktaEmail) {
		return true
	}

	return false
}

// SetOktaEmail gets a reference to the given PasswordPolicyRecoveryEmail and assigns it to the OktaEmail field.
func (o *PasswordPolicyRecoveryFactors) SetOktaEmail(v PasswordPolicyRecoveryEmail) {
	o.OktaEmail = &v
}

// GetOktaSms returns the OktaSms field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryFactors) GetOktaSms() PasswordPolicyRecoveryFactorSettings {
	if o == nil || IsNil(o.OktaSms) {
		var ret PasswordPolicyRecoveryFactorSettings
		return ret
	}
	return *o.OktaSms
}

// GetOktaSmsOk returns a tuple with the OktaSms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaSmsOk() (*PasswordPolicyRecoveryFactorSettings, bool) {
	if o == nil || IsNil(o.OktaSms) {
		return nil, false
	}
	return o.OktaSms, true
}

// HasOktaSms returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaSms() bool {
	if o != nil && !IsNil(o.OktaSms) {
		return true
	}

	return false
}

// SetOktaSms gets a reference to the given PasswordPolicyRecoveryFactorSettings and assigns it to the OktaSms field.
func (o *PasswordPolicyRecoveryFactors) SetOktaSms(v PasswordPolicyRecoveryFactorSettings) {
	o.OktaSms = &v
}

// GetRecoveryQuestion returns the RecoveryQuestion field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryFactors) GetRecoveryQuestion() PasswordPolicyRecoveryQuestion {
	if o == nil || IsNil(o.RecoveryQuestion) {
		var ret PasswordPolicyRecoveryQuestion
		return ret
	}
	return *o.RecoveryQuestion
}

// GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetRecoveryQuestionOk() (*PasswordPolicyRecoveryQuestion, bool) {
	if o == nil || IsNil(o.RecoveryQuestion) {
		return nil, false
	}
	return o.RecoveryQuestion, true
}

// HasRecoveryQuestion returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasRecoveryQuestion() bool {
	if o != nil && !IsNil(o.RecoveryQuestion) {
		return true
	}

	return false
}

// SetRecoveryQuestion gets a reference to the given PasswordPolicyRecoveryQuestion and assigns it to the RecoveryQuestion field.
func (o *PasswordPolicyRecoveryFactors) SetRecoveryQuestion(v PasswordPolicyRecoveryQuestion) {
	o.RecoveryQuestion = &v
}

func (o PasswordPolicyRecoveryFactors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoveryFactors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OktaCall) {
		toSerialize["okta_call"] = o.OktaCall
	}
	if !IsNil(o.OktaEmail) {
		toSerialize["okta_email"] = o.OktaEmail
	}
	if !IsNil(o.OktaSms) {
		toSerialize["okta_sms"] = o.OktaSms
	}
	if !IsNil(o.RecoveryQuestion) {
		toSerialize["recovery_question"] = o.RecoveryQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoveryFactors) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoveryFactors := _PasswordPolicyRecoveryFactors{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoveryFactors)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoveryFactors(varPasswordPolicyRecoveryFactors)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "okta_call")
		delete(additionalProperties, "okta_email")
		delete(additionalProperties, "okta_sms")
		delete(additionalProperties, "recovery_question")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoveryFactors struct {
	value *PasswordPolicyRecoveryFactors
	isSet bool
}

func (v NullablePasswordPolicyRecoveryFactors) Get() *PasswordPolicyRecoveryFactors {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryFactors) Set(val *PasswordPolicyRecoveryFactors) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryFactors) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryFactors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryFactors(val *PasswordPolicyRecoveryFactors) *NullablePasswordPolicyRecoveryFactors {
	return &NullablePasswordPolicyRecoveryFactors{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryFactors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryFactors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

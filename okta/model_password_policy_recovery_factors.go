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

// PasswordPolicyRecoveryFactors struct for PasswordPolicyRecoveryFactors
type PasswordPolicyRecoveryFactors struct {
	OktaCall *PasswordPolicyRecoveryFactorSettings `json:"okta_call,omitempty"`
	OktaEmail *PasswordPolicyRecoveryEmail `json:"okta_email,omitempty"`
	OktaSms *PasswordPolicyRecoveryFactorSettings `json:"okta_sms,omitempty"`
	RecoveryQuestion *PasswordPolicyRecoveryQuestion `json:"recovery_question,omitempty"`
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
	if o == nil || o.OktaCall == nil {
		var ret PasswordPolicyRecoveryFactorSettings
		return ret
	}
	return *o.OktaCall
}

// GetOktaCallOk returns a tuple with the OktaCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaCallOk() (*PasswordPolicyRecoveryFactorSettings, bool) {
	if o == nil || o.OktaCall == nil {
		return nil, false
	}
	return o.OktaCall, true
}

// HasOktaCall returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaCall() bool {
	if o != nil && o.OktaCall != nil {
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
	if o == nil || o.OktaEmail == nil {
		var ret PasswordPolicyRecoveryEmail
		return ret
	}
	return *o.OktaEmail
}

// GetOktaEmailOk returns a tuple with the OktaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaEmailOk() (*PasswordPolicyRecoveryEmail, bool) {
	if o == nil || o.OktaEmail == nil {
		return nil, false
	}
	return o.OktaEmail, true
}

// HasOktaEmail returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaEmail() bool {
	if o != nil && o.OktaEmail != nil {
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
	if o == nil || o.OktaSms == nil {
		var ret PasswordPolicyRecoveryFactorSettings
		return ret
	}
	return *o.OktaSms
}

// GetOktaSmsOk returns a tuple with the OktaSms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetOktaSmsOk() (*PasswordPolicyRecoveryFactorSettings, bool) {
	if o == nil || o.OktaSms == nil {
		return nil, false
	}
	return o.OktaSms, true
}

// HasOktaSms returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasOktaSms() bool {
	if o != nil && o.OktaSms != nil {
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
	if o == nil || o.RecoveryQuestion == nil {
		var ret PasswordPolicyRecoveryQuestion
		return ret
	}
	return *o.RecoveryQuestion
}

// GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryFactors) GetRecoveryQuestionOk() (*PasswordPolicyRecoveryQuestion, bool) {
	if o == nil || o.RecoveryQuestion == nil {
		return nil, false
	}
	return o.RecoveryQuestion, true
}

// HasRecoveryQuestion returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryFactors) HasRecoveryQuestion() bool {
	if o != nil && o.RecoveryQuestion != nil {
		return true
	}

	return false
}

// SetRecoveryQuestion gets a reference to the given PasswordPolicyRecoveryQuestion and assigns it to the RecoveryQuestion field.
func (o *PasswordPolicyRecoveryFactors) SetRecoveryQuestion(v PasswordPolicyRecoveryQuestion) {
	o.RecoveryQuestion = &v
}

func (o PasswordPolicyRecoveryFactors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OktaCall != nil {
		toSerialize["okta_call"] = o.OktaCall
	}
	if o.OktaEmail != nil {
		toSerialize["okta_email"] = o.OktaEmail
	}
	if o.OktaSms != nil {
		toSerialize["okta_sms"] = o.OktaSms
	}
	if o.RecoveryQuestion != nil {
		toSerialize["recovery_question"] = o.RecoveryQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRecoveryFactors) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyRecoveryFactors := _PasswordPolicyRecoveryFactors{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRecoveryFactors)
	if err == nil {
		*o = PasswordPolicyRecoveryFactors(varPasswordPolicyRecoveryFactors)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "okta_call")
		delete(additionalProperties, "okta_email")
		delete(additionalProperties, "okta_sms")
		delete(additionalProperties, "recovery_question")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


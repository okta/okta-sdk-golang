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

// AttackProtectionAuthenticatorSettings struct for AttackProtectionAuthenticatorSettings
type AttackProtectionAuthenticatorSettings struct {
	// If true, requires users to verify a possession factor before verifying a knowledge factor when the assurance requires two-factor authentication (2FA).
	VerifyKnowledgeSecondWhen2faRequired *bool `json:"verifyKnowledgeSecondWhen2faRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttackProtectionAuthenticatorSettings AttackProtectionAuthenticatorSettings

// NewAttackProtectionAuthenticatorSettings instantiates a new AttackProtectionAuthenticatorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackProtectionAuthenticatorSettings() *AttackProtectionAuthenticatorSettings {
	this := AttackProtectionAuthenticatorSettings{}
	var verifyKnowledgeSecondWhen2faRequired bool = false
	this.VerifyKnowledgeSecondWhen2faRequired = &verifyKnowledgeSecondWhen2faRequired
	return &this
}

// NewAttackProtectionAuthenticatorSettingsWithDefaults instantiates a new AttackProtectionAuthenticatorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackProtectionAuthenticatorSettingsWithDefaults() *AttackProtectionAuthenticatorSettings {
	this := AttackProtectionAuthenticatorSettings{}
	var verifyKnowledgeSecondWhen2faRequired bool = false
	this.VerifyKnowledgeSecondWhen2faRequired = &verifyKnowledgeSecondWhen2faRequired
	return &this
}

// GetVerifyKnowledgeSecondWhen2faRequired returns the VerifyKnowledgeSecondWhen2faRequired field value if set, zero value otherwise.
func (o *AttackProtectionAuthenticatorSettings) GetVerifyKnowledgeSecondWhen2faRequired() bool {
	if o == nil || o.VerifyKnowledgeSecondWhen2faRequired == nil {
		var ret bool
		return ret
	}
	return *o.VerifyKnowledgeSecondWhen2faRequired
}

// GetVerifyKnowledgeSecondWhen2faRequiredOk returns a tuple with the VerifyKnowledgeSecondWhen2faRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackProtectionAuthenticatorSettings) GetVerifyKnowledgeSecondWhen2faRequiredOk() (*bool, bool) {
	if o == nil || o.VerifyKnowledgeSecondWhen2faRequired == nil {
		return nil, false
	}
	return o.VerifyKnowledgeSecondWhen2faRequired, true
}

// HasVerifyKnowledgeSecondWhen2faRequired returns a boolean if a field has been set.
func (o *AttackProtectionAuthenticatorSettings) HasVerifyKnowledgeSecondWhen2faRequired() bool {
	if o != nil && o.VerifyKnowledgeSecondWhen2faRequired != nil {
		return true
	}

	return false
}

// SetVerifyKnowledgeSecondWhen2faRequired gets a reference to the given bool and assigns it to the VerifyKnowledgeSecondWhen2faRequired field.
func (o *AttackProtectionAuthenticatorSettings) SetVerifyKnowledgeSecondWhen2faRequired(v bool) {
	o.VerifyKnowledgeSecondWhen2faRequired = &v
}

func (o AttackProtectionAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VerifyKnowledgeSecondWhen2faRequired != nil {
		toSerialize["verifyKnowledgeSecondWhen2faRequired"] = o.VerifyKnowledgeSecondWhen2faRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AttackProtectionAuthenticatorSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAttackProtectionAuthenticatorSettings := _AttackProtectionAuthenticatorSettings{}

	err = json.Unmarshal(bytes, &varAttackProtectionAuthenticatorSettings)
	if err == nil {
		*o = AttackProtectionAuthenticatorSettings(varAttackProtectionAuthenticatorSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "verifyKnowledgeSecondWhen2faRequired")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAttackProtectionAuthenticatorSettings struct {
	value *AttackProtectionAuthenticatorSettings
	isSet bool
}

func (v NullableAttackProtectionAuthenticatorSettings) Get() *AttackProtectionAuthenticatorSettings {
	return v.value
}

func (v *NullableAttackProtectionAuthenticatorSettings) Set(val *AttackProtectionAuthenticatorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackProtectionAuthenticatorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackProtectionAuthenticatorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackProtectionAuthenticatorSettings(val *AttackProtectionAuthenticatorSettings) *NullableAttackProtectionAuthenticatorSettings {
	return &NullableAttackProtectionAuthenticatorSettings{value: val, isSet: true}
}

func (v NullableAttackProtectionAuthenticatorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackProtectionAuthenticatorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyRecoverySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRecoverySettings{}

// PasswordPolicyRecoverySettings Specifies the password recovery settings for the policy > **Note:** With Identity Engine, you can specify recovery factors inside the password policy rule instead of in the policy settings object. Recovery factors for the rule are defined inside the [`selfServicePasswordReset` action](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Policy/#tag/Policy/operation/createPolicyRule!path=1/actions/selfServicePasswordReset&t=request).
type PasswordPolicyRecoverySettings struct {
	Factors              *PasswordPolicyRecoveryFactors `json:"factors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoverySettings PasswordPolicyRecoverySettings

// NewPasswordPolicyRecoverySettings instantiates a new PasswordPolicyRecoverySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoverySettings() *PasswordPolicyRecoverySettings {
	this := PasswordPolicyRecoverySettings{}
	return &this
}

// NewPasswordPolicyRecoverySettingsWithDefaults instantiates a new PasswordPolicyRecoverySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoverySettingsWithDefaults() *PasswordPolicyRecoverySettings {
	this := PasswordPolicyRecoverySettings{}
	return &this
}

// GetFactors returns the Factors field value if set, zero value otherwise.
func (o *PasswordPolicyRecoverySettings) GetFactors() PasswordPolicyRecoveryFactors {
	if o == nil || IsNil(o.Factors) {
		var ret PasswordPolicyRecoveryFactors
		return ret
	}
	return *o.Factors
}

// GetFactorsOk returns a tuple with the Factors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoverySettings) GetFactorsOk() (*PasswordPolicyRecoveryFactors, bool) {
	if o == nil || IsNil(o.Factors) {
		return nil, false
	}
	return o.Factors, true
}

// HasFactors returns a boolean if a field has been set.
func (o *PasswordPolicyRecoverySettings) HasFactors() bool {
	if o != nil && !IsNil(o.Factors) {
		return true
	}

	return false
}

// SetFactors gets a reference to the given PasswordPolicyRecoveryFactors and assigns it to the Factors field.
func (o *PasswordPolicyRecoverySettings) SetFactors(v PasswordPolicyRecoveryFactors) {
	o.Factors = &v
}

func (o PasswordPolicyRecoverySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRecoverySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Factors) {
		toSerialize["factors"] = o.Factors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRecoverySettings) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRecoverySettings := _PasswordPolicyRecoverySettings{}

	err = json.Unmarshal(data, &varPasswordPolicyRecoverySettings)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRecoverySettings(varPasswordPolicyRecoverySettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRecoverySettings struct {
	value *PasswordPolicyRecoverySettings
	isSet bool
}

func (v NullablePasswordPolicyRecoverySettings) Get() *PasswordPolicyRecoverySettings {
	return v.value
}

func (v *NullablePasswordPolicyRecoverySettings) Set(val *PasswordPolicyRecoverySettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoverySettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoverySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoverySettings(val *PasswordPolicyRecoverySettings) *NullablePasswordPolicyRecoverySettings {
	return &NullablePasswordPolicyRecoverySettings{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoverySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoverySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

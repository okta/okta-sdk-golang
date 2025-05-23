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

// PasswordPolicySettings struct for PasswordPolicySettings
type PasswordPolicySettings struct {
	Delegation *PasswordPolicyDelegationSettings `json:"delegation,omitempty"`
	Password *PasswordPolicyPasswordSettings `json:"password,omitempty"`
	Recovery *PasswordPolicyRecoverySettings `json:"recovery,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicySettings PasswordPolicySettings

// NewPasswordPolicySettings instantiates a new PasswordPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicySettings() *PasswordPolicySettings {
	this := PasswordPolicySettings{}
	return &this
}

// NewPasswordPolicySettingsWithDefaults instantiates a new PasswordPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicySettingsWithDefaults() *PasswordPolicySettings {
	this := PasswordPolicySettings{}
	return &this
}

// GetDelegation returns the Delegation field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetDelegation() PasswordPolicyDelegationSettings {
	if o == nil || o.Delegation == nil {
		var ret PasswordPolicyDelegationSettings
		return ret
	}
	return *o.Delegation
}

// GetDelegationOk returns a tuple with the Delegation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetDelegationOk() (*PasswordPolicyDelegationSettings, bool) {
	if o == nil || o.Delegation == nil {
		return nil, false
	}
	return o.Delegation, true
}

// HasDelegation returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasDelegation() bool {
	if o != nil && o.Delegation != nil {
		return true
	}

	return false
}

// SetDelegation gets a reference to the given PasswordPolicyDelegationSettings and assigns it to the Delegation field.
func (o *PasswordPolicySettings) SetDelegation(v PasswordPolicyDelegationSettings) {
	o.Delegation = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetPassword() PasswordPolicyPasswordSettings {
	if o == nil || o.Password == nil {
		var ret PasswordPolicyPasswordSettings
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetPasswordOk() (*PasswordPolicyPasswordSettings, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordPolicyPasswordSettings and assigns it to the Password field.
func (o *PasswordPolicySettings) SetPassword(v PasswordPolicyPasswordSettings) {
	o.Password = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *PasswordPolicySettings) GetRecovery() PasswordPolicyRecoverySettings {
	if o == nil || o.Recovery == nil {
		var ret PasswordPolicyRecoverySettings
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicySettings) GetRecoveryOk() (*PasswordPolicyRecoverySettings, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *PasswordPolicySettings) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given PasswordPolicyRecoverySettings and assigns it to the Recovery field.
func (o *PasswordPolicySettings) SetRecovery(v PasswordPolicyRecoverySettings) {
	o.Recovery = &v
}

func (o PasswordPolicySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delegation != nil {
		toSerialize["delegation"] = o.Delegation
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Recovery != nil {
		toSerialize["recovery"] = o.Recovery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicySettings) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicySettings := _PasswordPolicySettings{}

	err = json.Unmarshal(bytes, &varPasswordPolicySettings)
	if err == nil {
		*o = PasswordPolicySettings(varPasswordPolicySettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "delegation")
		delete(additionalProperties, "password")
		delete(additionalProperties, "recovery")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicySettings struct {
	value *PasswordPolicySettings
	isSet bool
}

func (v NullablePasswordPolicySettings) Get() *PasswordPolicySettings {
	return v.value
}

func (v *NullablePasswordPolicySettings) Set(val *PasswordPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicySettings(val *PasswordPolicySettings) *NullablePasswordPolicySettings {
	return &NullablePasswordPolicySettings{value: val, isSet: true}
}

func (v NullablePasswordPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


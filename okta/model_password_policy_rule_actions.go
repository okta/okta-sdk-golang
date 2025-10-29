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

// checks if the PasswordPolicyRuleActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyRuleActions{}

// PasswordPolicyRuleActions struct for PasswordPolicyRuleActions
type PasswordPolicyRuleActions struct {
	PasswordChange           *PasswordPolicyRuleAction       `json:"passwordChange,omitempty"`
	SelfServicePasswordReset *SelfServicePasswordResetAction `json:"selfServicePasswordReset,omitempty"`
	SelfServiceUnlock        *PasswordPolicyRuleAction       `json:"selfServiceUnlock,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _PasswordPolicyRuleActions PasswordPolicyRuleActions

// NewPasswordPolicyRuleActions instantiates a new PasswordPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRuleActions() *PasswordPolicyRuleActions {
	this := PasswordPolicyRuleActions{}
	return &this
}

// NewPasswordPolicyRuleActionsWithDefaults instantiates a new PasswordPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRuleActionsWithDefaults() *PasswordPolicyRuleActions {
	this := PasswordPolicyRuleActions{}
	return &this
}

// GetPasswordChange returns the PasswordChange field value if set, zero value otherwise.
func (o *PasswordPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction {
	if o == nil || IsNil(o.PasswordChange) {
		var ret PasswordPolicyRuleAction
		return ret
	}
	return *o.PasswordChange
}

// GetPasswordChangeOk returns a tuple with the PasswordChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool) {
	if o == nil || IsNil(o.PasswordChange) {
		return nil, false
	}
	return o.PasswordChange, true
}

// HasPasswordChange returns a boolean if a field has been set.
func (o *PasswordPolicyRuleActions) HasPasswordChange() bool {
	if o != nil && !IsNil(o.PasswordChange) {
		return true
	}

	return false
}

// SetPasswordChange gets a reference to the given PasswordPolicyRuleAction and assigns it to the PasswordChange field.
func (o *PasswordPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction) {
	o.PasswordChange = &v
}

// GetSelfServicePasswordReset returns the SelfServicePasswordReset field value if set, zero value otherwise.
func (o *PasswordPolicyRuleActions) GetSelfServicePasswordReset() SelfServicePasswordResetAction {
	if o == nil || IsNil(o.SelfServicePasswordReset) {
		var ret SelfServicePasswordResetAction
		return ret
	}
	return *o.SelfServicePasswordReset
}

// GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleActions) GetSelfServicePasswordResetOk() (*SelfServicePasswordResetAction, bool) {
	if o == nil || IsNil(o.SelfServicePasswordReset) {
		return nil, false
	}
	return o.SelfServicePasswordReset, true
}

// HasSelfServicePasswordReset returns a boolean if a field has been set.
func (o *PasswordPolicyRuleActions) HasSelfServicePasswordReset() bool {
	if o != nil && !IsNil(o.SelfServicePasswordReset) {
		return true
	}

	return false
}

// SetSelfServicePasswordReset gets a reference to the given SelfServicePasswordResetAction and assigns it to the SelfServicePasswordReset field.
func (o *PasswordPolicyRuleActions) SetSelfServicePasswordReset(v SelfServicePasswordResetAction) {
	o.SelfServicePasswordReset = &v
}

// GetSelfServiceUnlock returns the SelfServiceUnlock field value if set, zero value otherwise.
func (o *PasswordPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction {
	if o == nil || IsNil(o.SelfServiceUnlock) {
		var ret PasswordPolicyRuleAction
		return ret
	}
	return *o.SelfServiceUnlock
}

// GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool) {
	if o == nil || IsNil(o.SelfServiceUnlock) {
		return nil, false
	}
	return o.SelfServiceUnlock, true
}

// HasSelfServiceUnlock returns a boolean if a field has been set.
func (o *PasswordPolicyRuleActions) HasSelfServiceUnlock() bool {
	if o != nil && !IsNil(o.SelfServiceUnlock) {
		return true
	}

	return false
}

// SetSelfServiceUnlock gets a reference to the given PasswordPolicyRuleAction and assigns it to the SelfServiceUnlock field.
func (o *PasswordPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction) {
	o.SelfServiceUnlock = &v
}

func (o PasswordPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyRuleActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PasswordChange) {
		toSerialize["passwordChange"] = o.PasswordChange
	}
	if !IsNil(o.SelfServicePasswordReset) {
		toSerialize["selfServicePasswordReset"] = o.SelfServicePasswordReset
	}
	if !IsNil(o.SelfServiceUnlock) {
		toSerialize["selfServiceUnlock"] = o.SelfServiceUnlock
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyRuleActions) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyRuleActions := _PasswordPolicyRuleActions{}

	err = json.Unmarshal(data, &varPasswordPolicyRuleActions)

	if err != nil {
		return err
	}

	*o = PasswordPolicyRuleActions(varPasswordPolicyRuleActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "passwordChange")
		delete(additionalProperties, "selfServicePasswordReset")
		delete(additionalProperties, "selfServiceUnlock")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyRuleActions struct {
	value *PasswordPolicyRuleActions
	isSet bool
}

func (v NullablePasswordPolicyRuleActions) Get() *PasswordPolicyRuleActions {
	return v.value
}

func (v *NullablePasswordPolicyRuleActions) Set(val *PasswordPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRuleActions(val *PasswordPolicyRuleActions) *NullablePasswordPolicyRuleActions {
	return &NullablePasswordPolicyRuleActions{value: val, isSet: true}
}

func (v NullablePasswordPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

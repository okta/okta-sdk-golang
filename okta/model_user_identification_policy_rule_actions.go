/*
Okta Admin Management API

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

// checks if the UserIdentificationPolicyRuleActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationPolicyRuleActions{}

// UserIdentificationPolicyRuleActions <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies actions to be taken, or operations that may be allowed, if the rule conditions are satisfied
type UserIdentificationPolicyRuleActions struct {
	UserIdentification   *UserIdentificationPolicyRuleIdentification `json:"userIdentification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationPolicyRuleActions UserIdentificationPolicyRuleActions

// NewUserIdentificationPolicyRuleActions instantiates a new UserIdentificationPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationPolicyRuleActions() *UserIdentificationPolicyRuleActions {
	this := UserIdentificationPolicyRuleActions{}
	return &this
}

// NewUserIdentificationPolicyRuleActionsWithDefaults instantiates a new UserIdentificationPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationPolicyRuleActionsWithDefaults() *UserIdentificationPolicyRuleActions {
	this := UserIdentificationPolicyRuleActions{}
	return &this
}

// GetUserIdentification returns the UserIdentification field value if set, zero value otherwise.
func (o *UserIdentificationPolicyRuleActions) GetUserIdentification() UserIdentificationPolicyRuleIdentification {
	if o == nil || IsNil(o.UserIdentification) {
		var ret UserIdentificationPolicyRuleIdentification
		return ret
	}
	return *o.UserIdentification
}

// GetUserIdentificationOk returns a tuple with the UserIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationPolicyRuleActions) GetUserIdentificationOk() (*UserIdentificationPolicyRuleIdentification, bool) {
	if o == nil || IsNil(o.UserIdentification) {
		return nil, false
	}
	return o.UserIdentification, true
}

// HasUserIdentification returns a boolean if a field has been set.
func (o *UserIdentificationPolicyRuleActions) HasUserIdentification() bool {
	if o != nil && !IsNil(o.UserIdentification) {
		return true
	}

	return false
}

// SetUserIdentification gets a reference to the given UserIdentificationPolicyRuleIdentification and assigns it to the UserIdentification field.
func (o *UserIdentificationPolicyRuleActions) SetUserIdentification(v UserIdentificationPolicyRuleIdentification) {
	o.UserIdentification = &v
}

func (o UserIdentificationPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationPolicyRuleActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserIdentification) {
		toSerialize["userIdentification"] = o.UserIdentification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationPolicyRuleActions) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationPolicyRuleActions := _UserIdentificationPolicyRuleActions{}

	err = json.Unmarshal(data, &varUserIdentificationPolicyRuleActions)

	if err != nil {
		return err
	}

	*o = UserIdentificationPolicyRuleActions(varUserIdentificationPolicyRuleActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userIdentification")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationPolicyRuleActions struct {
	value *UserIdentificationPolicyRuleActions
	isSet bool
}

func (v NullableUserIdentificationPolicyRuleActions) Get() *UserIdentificationPolicyRuleActions {
	return v.value
}

func (v *NullableUserIdentificationPolicyRuleActions) Set(val *UserIdentificationPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationPolicyRuleActions(val *UserIdentificationPolicyRuleActions) *NullableUserIdentificationPolicyRuleActions {
	return &NullableUserIdentificationPolicyRuleActions{value: val, isSet: true}
}

func (v NullableUserIdentificationPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

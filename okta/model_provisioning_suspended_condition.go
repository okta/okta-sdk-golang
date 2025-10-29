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

// checks if the ProvisioningSuspendedCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningSuspendedCondition{}

// ProvisioningSuspendedCondition Behavior for a previously suspended IdP user during authentication. Not supported with OIDC IdPs.
type ProvisioningSuspendedCondition struct {
	// Specifies the action during authentication when an IdP user is linked to a previously suspended Okta user
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningSuspendedCondition ProvisioningSuspendedCondition

// NewProvisioningSuspendedCondition instantiates a new ProvisioningSuspendedCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningSuspendedCondition() *ProvisioningSuspendedCondition {
	this := ProvisioningSuspendedCondition{}
	return &this
}

// NewProvisioningSuspendedConditionWithDefaults instantiates a new ProvisioningSuspendedCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningSuspendedConditionWithDefaults() *ProvisioningSuspendedCondition {
	this := ProvisioningSuspendedCondition{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ProvisioningSuspendedCondition) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningSuspendedCondition) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ProvisioningSuspendedCondition) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ProvisioningSuspendedCondition) SetAction(v string) {
	o.Action = &v
}

func (o ProvisioningSuspendedCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningSuspendedCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningSuspendedCondition) UnmarshalJSON(data []byte) (err error) {
	varProvisioningSuspendedCondition := _ProvisioningSuspendedCondition{}

	err = json.Unmarshal(data, &varProvisioningSuspendedCondition)

	if err != nil {
		return err
	}

	*o = ProvisioningSuspendedCondition(varProvisioningSuspendedCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningSuspendedCondition struct {
	value *ProvisioningSuspendedCondition
	isSet bool
}

func (v NullableProvisioningSuspendedCondition) Get() *ProvisioningSuspendedCondition {
	return v.value
}

func (v *NullableProvisioningSuspendedCondition) Set(val *ProvisioningSuspendedCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningSuspendedCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningSuspendedCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningSuspendedCondition(val *ProvisioningSuspendedCondition) *NullableProvisioningSuspendedCondition {
	return &NullableProvisioningSuspendedCondition{value: val, isSet: true}
}

func (v NullableProvisioningSuspendedCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningSuspendedCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

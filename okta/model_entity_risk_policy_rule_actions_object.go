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

// checks if the EntityRiskPolicyRuleActionsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRiskPolicyRuleActionsObject{}

// EntityRiskPolicyRuleActionsObject struct for EntityRiskPolicyRuleActionsObject
type EntityRiskPolicyRuleActionsObject struct {
	Action               *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleActionsObject EntityRiskPolicyRuleActionsObject

// NewEntityRiskPolicyRuleActionsObject instantiates a new EntityRiskPolicyRuleActionsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleActionsObject() *EntityRiskPolicyRuleActionsObject {
	this := EntityRiskPolicyRuleActionsObject{}
	return &this
}

// NewEntityRiskPolicyRuleActionsObjectWithDefaults instantiates a new EntityRiskPolicyRuleActionsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleActionsObjectWithDefaults() *EntityRiskPolicyRuleActionsObject {
	this := EntityRiskPolicyRuleActionsObject{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleActionsObject) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleActionsObject) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleActionsObject) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EntityRiskPolicyRuleActionsObject) SetAction(v string) {
	o.Action = &v
}

func (o EntityRiskPolicyRuleActionsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRiskPolicyRuleActionsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityRiskPolicyRuleActionsObject) UnmarshalJSON(data []byte) (err error) {
	varEntityRiskPolicyRuleActionsObject := _EntityRiskPolicyRuleActionsObject{}

	err = json.Unmarshal(data, &varEntityRiskPolicyRuleActionsObject)

	if err != nil {
		return err
	}

	*o = EntityRiskPolicyRuleActionsObject(varEntityRiskPolicyRuleActionsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityRiskPolicyRuleActionsObject struct {
	value *EntityRiskPolicyRuleActionsObject
	isSet bool
}

func (v NullableEntityRiskPolicyRuleActionsObject) Get() *EntityRiskPolicyRuleActionsObject {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleActionsObject) Set(val *EntityRiskPolicyRuleActionsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleActionsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleActionsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleActionsObject(val *EntityRiskPolicyRuleActionsObject) *NullableEntityRiskPolicyRuleActionsObject {
	return &NullableEntityRiskPolicyRuleActionsObject{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleActionsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleActionsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

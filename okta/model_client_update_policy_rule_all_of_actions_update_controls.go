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

// checks if the ClientUpdatePolicyRuleAllOfActionsUpdateControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientUpdatePolicyRuleAllOfActionsUpdateControls{}

// ClientUpdatePolicyRuleAllOfActionsUpdateControls The object that contains the `actions` array
type ClientUpdatePolicyRuleAllOfActionsUpdateControls struct {
	// The `updateControls` object's `actions` array supports one action type that determines if an Okta Verify client version update is allowed.
	Actions              []ClientUpdatePolicyRuleActionsObject `json:"actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientUpdatePolicyRuleAllOfActionsUpdateControls ClientUpdatePolicyRuleAllOfActionsUpdateControls

// NewClientUpdatePolicyRuleAllOfActionsUpdateControls instantiates a new ClientUpdatePolicyRuleAllOfActionsUpdateControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientUpdatePolicyRuleAllOfActionsUpdateControls() *ClientUpdatePolicyRuleAllOfActionsUpdateControls {
	this := ClientUpdatePolicyRuleAllOfActionsUpdateControls{}
	return &this
}

// NewClientUpdatePolicyRuleAllOfActionsUpdateControlsWithDefaults instantiates a new ClientUpdatePolicyRuleAllOfActionsUpdateControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientUpdatePolicyRuleAllOfActionsUpdateControlsWithDefaults() *ClientUpdatePolicyRuleAllOfActionsUpdateControls {
	this := ClientUpdatePolicyRuleAllOfActionsUpdateControls{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ClientUpdatePolicyRuleAllOfActionsUpdateControls) GetActions() []ClientUpdatePolicyRuleActionsObject {
	if o == nil || IsNil(o.Actions) {
		var ret []ClientUpdatePolicyRuleActionsObject
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUpdatePolicyRuleAllOfActionsUpdateControls) GetActionsOk() ([]ClientUpdatePolicyRuleActionsObject, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ClientUpdatePolicyRuleAllOfActionsUpdateControls) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ClientUpdatePolicyRuleActionsObject and assigns it to the Actions field.
func (o *ClientUpdatePolicyRuleAllOfActionsUpdateControls) SetActions(v []ClientUpdatePolicyRuleActionsObject) {
	o.Actions = v
}

func (o ClientUpdatePolicyRuleAllOfActionsUpdateControls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientUpdatePolicyRuleAllOfActionsUpdateControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientUpdatePolicyRuleAllOfActionsUpdateControls) UnmarshalJSON(data []byte) (err error) {
	varClientUpdatePolicyRuleAllOfActionsUpdateControls := _ClientUpdatePolicyRuleAllOfActionsUpdateControls{}

	err = json.Unmarshal(data, &varClientUpdatePolicyRuleAllOfActionsUpdateControls)

	if err != nil {
		return err
	}

	*o = ClientUpdatePolicyRuleAllOfActionsUpdateControls(varClientUpdatePolicyRuleAllOfActionsUpdateControls)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientUpdatePolicyRuleAllOfActionsUpdateControls struct {
	value *ClientUpdatePolicyRuleAllOfActionsUpdateControls
	isSet bool
}

func (v NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) Get() *ClientUpdatePolicyRuleAllOfActionsUpdateControls {
	return v.value
}

func (v *NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) Set(val *ClientUpdatePolicyRuleAllOfActionsUpdateControls) {
	v.value = val
	v.isSet = true
}

func (v NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) IsSet() bool {
	return v.isSet
}

func (v *NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientUpdatePolicyRuleAllOfActionsUpdateControls(val *ClientUpdatePolicyRuleAllOfActionsUpdateControls) *NullableClientUpdatePolicyRuleAllOfActionsUpdateControls {
	return &NullableClientUpdatePolicyRuleAllOfActionsUpdateControls{value: val, isSet: true}
}

func (v NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientUpdatePolicyRuleAllOfActionsUpdateControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

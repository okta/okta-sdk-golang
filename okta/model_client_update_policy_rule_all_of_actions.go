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

// checks if the ClientUpdatePolicyRuleAllOfActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientUpdatePolicyRuleAllOfActions{}

// ClientUpdatePolicyRuleAllOfActions The action to take based on the risk event
type ClientUpdatePolicyRuleAllOfActions struct {
	UpdateControls       *ClientUpdatePolicyRuleAllOfActionsUpdateControls `json:"updateControls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientUpdatePolicyRuleAllOfActions ClientUpdatePolicyRuleAllOfActions

// NewClientUpdatePolicyRuleAllOfActions instantiates a new ClientUpdatePolicyRuleAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientUpdatePolicyRuleAllOfActions() *ClientUpdatePolicyRuleAllOfActions {
	this := ClientUpdatePolicyRuleAllOfActions{}
	return &this
}

// NewClientUpdatePolicyRuleAllOfActionsWithDefaults instantiates a new ClientUpdatePolicyRuleAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientUpdatePolicyRuleAllOfActionsWithDefaults() *ClientUpdatePolicyRuleAllOfActions {
	this := ClientUpdatePolicyRuleAllOfActions{}
	return &this
}

// GetUpdateControls returns the UpdateControls field value if set, zero value otherwise.
func (o *ClientUpdatePolicyRuleAllOfActions) GetUpdateControls() ClientUpdatePolicyRuleAllOfActionsUpdateControls {
	if o == nil || IsNil(o.UpdateControls) {
		var ret ClientUpdatePolicyRuleAllOfActionsUpdateControls
		return ret
	}
	return *o.UpdateControls
}

// GetUpdateControlsOk returns a tuple with the UpdateControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientUpdatePolicyRuleAllOfActions) GetUpdateControlsOk() (*ClientUpdatePolicyRuleAllOfActionsUpdateControls, bool) {
	if o == nil || IsNil(o.UpdateControls) {
		return nil, false
	}
	return o.UpdateControls, true
}

// HasUpdateControls returns a boolean if a field has been set.
func (o *ClientUpdatePolicyRuleAllOfActions) HasUpdateControls() bool {
	if o != nil && !IsNil(o.UpdateControls) {
		return true
	}

	return false
}

// SetUpdateControls gets a reference to the given ClientUpdatePolicyRuleAllOfActionsUpdateControls and assigns it to the UpdateControls field.
func (o *ClientUpdatePolicyRuleAllOfActions) SetUpdateControls(v ClientUpdatePolicyRuleAllOfActionsUpdateControls) {
	o.UpdateControls = &v
}

func (o ClientUpdatePolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientUpdatePolicyRuleAllOfActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdateControls) {
		toSerialize["updateControls"] = o.UpdateControls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientUpdatePolicyRuleAllOfActions) UnmarshalJSON(data []byte) (err error) {
	varClientUpdatePolicyRuleAllOfActions := _ClientUpdatePolicyRuleAllOfActions{}

	err = json.Unmarshal(data, &varClientUpdatePolicyRuleAllOfActions)

	if err != nil {
		return err
	}

	*o = ClientUpdatePolicyRuleAllOfActions(varClientUpdatePolicyRuleAllOfActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "updateControls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientUpdatePolicyRuleAllOfActions struct {
	value *ClientUpdatePolicyRuleAllOfActions
	isSet bool
}

func (v NullableClientUpdatePolicyRuleAllOfActions) Get() *ClientUpdatePolicyRuleAllOfActions {
	return v.value
}

func (v *NullableClientUpdatePolicyRuleAllOfActions) Set(val *ClientUpdatePolicyRuleAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullableClientUpdatePolicyRuleAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullableClientUpdatePolicyRuleAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientUpdatePolicyRuleAllOfActions(val *ClientUpdatePolicyRuleAllOfActions) *NullableClientUpdatePolicyRuleAllOfActions {
	return &NullableClientUpdatePolicyRuleAllOfActions{value: val, isSet: true}
}

func (v NullableClientUpdatePolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientUpdatePolicyRuleAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the AvailableActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableActions{}

// AvailableActions struct for AvailableActions
type AvailableActions struct {
	Actions              []AvailableAction `json:"actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AvailableActions AvailableActions

// NewAvailableActions instantiates a new AvailableActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableActions() *AvailableActions {
	this := AvailableActions{}
	return &this
}

// NewAvailableActionsWithDefaults instantiates a new AvailableActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableActionsWithDefaults() *AvailableActions {
	this := AvailableActions{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AvailableActions) GetActions() []AvailableAction {
	if o == nil || IsNil(o.Actions) {
		var ret []AvailableAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableActions) GetActionsOk() ([]AvailableAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AvailableActions) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []AvailableAction and assigns it to the Actions field.
func (o *AvailableActions) SetActions(v []AvailableAction) {
	o.Actions = v
}

func (o AvailableActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailableActions) UnmarshalJSON(data []byte) (err error) {
	varAvailableActions := _AvailableActions{}

	err = json.Unmarshal(data, &varAvailableActions)

	if err != nil {
		return err
	}

	*o = AvailableActions(varAvailableActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailableActions struct {
	value *AvailableActions
	isSet bool
}

func (v NullableAvailableActions) Get() *AvailableActions {
	return v.value
}

func (v *NullableAvailableActions) Set(val *AvailableActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableActions(val *AvailableActions) *NullableAvailableActions {
	return &NullableAvailableActions{value: val, isSet: true}
}

func (v NullableAvailableActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

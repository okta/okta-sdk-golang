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

// checks if the SubmissionActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionActions{}

// SubmissionActions struct for SubmissionActions
type SubmissionActions struct {
	Actions              []SubmissionAction `json:"actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionActions SubmissionActions

// NewSubmissionActions instantiates a new SubmissionActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionActions() *SubmissionActions {
	this := SubmissionActions{}
	return &this
}

// NewSubmissionActionsWithDefaults instantiates a new SubmissionActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionActionsWithDefaults() *SubmissionActions {
	this := SubmissionActions{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *SubmissionActions) GetActions() []SubmissionAction {
	if o == nil || IsNil(o.Actions) {
		var ret []SubmissionAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionActions) GetActionsOk() ([]SubmissionAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *SubmissionActions) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []SubmissionAction and assigns it to the Actions field.
func (o *SubmissionActions) SetActions(v []SubmissionAction) {
	o.Actions = v
}

func (o SubmissionActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmissionActions) UnmarshalJSON(data []byte) (err error) {
	varSubmissionActions := _SubmissionActions{}

	err = json.Unmarshal(data, &varSubmissionActions)

	if err != nil {
		return err
	}

	*o = SubmissionActions(varSubmissionActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmissionActions struct {
	value *SubmissionActions
	isSet bool
}

func (v NullableSubmissionActions) Get() *SubmissionActions {
	return v.value
}

func (v *NullableSubmissionActions) Set(val *SubmissionActions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionActions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionActions(val *SubmissionActions) *NullableSubmissionActions {
	return &NullableSubmissionActions{value: val, isSet: true}
}

func (v NullableSubmissionActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

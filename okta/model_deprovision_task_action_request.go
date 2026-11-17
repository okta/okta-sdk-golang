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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DeprovisionTaskActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprovisionTaskActionRequest{}

// DeprovisionTaskActionRequest Specify exactly one of `taskIds` or `appInstanceIds`.
type DeprovisionTaskActionRequest struct {
	// Action to perform on a task. Only the actions listed in a task's `eligibleActions` can be performed on it. If you specify `taskIds` and a task doesn't allow the action, the request is rejected with a `400` error. If you specify `appInstanceIds`, only the tasks that allow the action are acted on. Tasks that no longer allow the action when the request is processed are skipped.
	Action string `json:"action"`
	// Acts on every currently eligible task across these app instances (maximum 10 by default)
	AppInstanceIds []string `json:"appInstanceIds,omitempty"`
	// Explicit IDs of the tasks to act on (maximum 50 by default)
	TaskIds []string `json:"taskIds,omitempty"`
}

type _DeprovisionTaskActionRequest DeprovisionTaskActionRequest

// NewDeprovisionTaskActionRequest instantiates a new DeprovisionTaskActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprovisionTaskActionRequest(action string) *DeprovisionTaskActionRequest {
	this := DeprovisionTaskActionRequest{}
	this.Action = action
	return &this
}

// NewDeprovisionTaskActionRequestWithDefaults instantiates a new DeprovisionTaskActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprovisionTaskActionRequestWithDefaults() *DeprovisionTaskActionRequest {
	this := DeprovisionTaskActionRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *DeprovisionTaskActionRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTaskActionRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *DeprovisionTaskActionRequest) SetAction(v string) {
	o.Action = v
}

// GetAppInstanceIds returns the AppInstanceIds field value if set, zero value otherwise.
func (o *DeprovisionTaskActionRequest) GetAppInstanceIds() []string {
	if o == nil || IsNil(o.AppInstanceIds) {
		var ret []string
		return ret
	}
	return o.AppInstanceIds
}

// GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprovisionTaskActionRequest) GetAppInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppInstanceIds) {
		return nil, false
	}
	return o.AppInstanceIds, true
}

// HasAppInstanceIds returns a boolean if a field has been set.
func (o *DeprovisionTaskActionRequest) HasAppInstanceIds() bool {
	if o != nil && !IsNil(o.AppInstanceIds) {
		return true
	}

	return false
}

// SetAppInstanceIds gets a reference to the given []string and assigns it to the AppInstanceIds field.
func (o *DeprovisionTaskActionRequest) SetAppInstanceIds(v []string) {
	o.AppInstanceIds = v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *DeprovisionTaskActionRequest) GetTaskIds() []string {
	if o == nil || IsNil(o.TaskIds) {
		var ret []string
		return ret
	}
	return o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprovisionTaskActionRequest) GetTaskIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TaskIds) {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *DeprovisionTaskActionRequest) HasTaskIds() bool {
	if o != nil && !IsNil(o.TaskIds) {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *DeprovisionTaskActionRequest) SetTaskIds(v []string) {
	o.TaskIds = v
}

func (o DeprovisionTaskActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprovisionTaskActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.AppInstanceIds) {
		toSerialize["appInstanceIds"] = o.AppInstanceIds
	}
	if !IsNil(o.TaskIds) {
		toSerialize["taskIds"] = o.TaskIds
	}
	return toSerialize, nil
}

func (o *DeprovisionTaskActionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeprovisionTaskActionRequest := _DeprovisionTaskActionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeprovisionTaskActionRequest)

	if err != nil {
		return err
	}

	*o = DeprovisionTaskActionRequest(varDeprovisionTaskActionRequest)

	return err
}

type NullableDeprovisionTaskActionRequest struct {
	value *DeprovisionTaskActionRequest
	isSet bool
}

func (v NullableDeprovisionTaskActionRequest) Get() *DeprovisionTaskActionRequest {
	return v.value
}

func (v *NullableDeprovisionTaskActionRequest) Set(val *DeprovisionTaskActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprovisionTaskActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprovisionTaskActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprovisionTaskActionRequest(val *DeprovisionTaskActionRequest) *NullableDeprovisionTaskActionRequest {
	return &NullableDeprovisionTaskActionRequest{value: val, isSet: true}
}

func (v NullableDeprovisionTaskActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprovisionTaskActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

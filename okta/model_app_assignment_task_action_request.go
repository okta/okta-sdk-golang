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

// checks if the AppAssignmentTaskActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAssignmentTaskActionRequest{}

// AppAssignmentTaskActionRequest Specify exactly one of `taskIds` or `appInstanceIds`.
type AppAssignmentTaskActionRequest struct {
	// Action to perform on an app assignment task. Only the actions listed in a task's `eligibleActions` can be performed on it. If you specify `taskIds` and a task doesn't allow the action, the request is rejected with a `400` error. If you specify `appInstanceIds`, only the tasks that allow the action are acted on. Tasks that no longer allow the action when the request is processed are skipped.
	Action string `json:"action"`
	// Acts on every currently eligible task across these app instances (maximum 10 by default)
	AppInstanceIds []string `json:"appInstanceIds,omitempty"`
	// Explicit IDs of the tasks to act on (maximum 50 by default)
	TaskIds []string `json:"taskIds,omitempty"`
}

type _AppAssignmentTaskActionRequest AppAssignmentTaskActionRequest

// NewAppAssignmentTaskActionRequest instantiates a new AppAssignmentTaskActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAssignmentTaskActionRequest(action string) *AppAssignmentTaskActionRequest {
	this := AppAssignmentTaskActionRequest{}
	this.Action = action
	return &this
}

// NewAppAssignmentTaskActionRequestWithDefaults instantiates a new AppAssignmentTaskActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAssignmentTaskActionRequestWithDefaults() *AppAssignmentTaskActionRequest {
	this := AppAssignmentTaskActionRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *AppAssignmentTaskActionRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTaskActionRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *AppAssignmentTaskActionRequest) SetAction(v string) {
	o.Action = v
}

// GetAppInstanceIds returns the AppInstanceIds field value if set, zero value otherwise.
func (o *AppAssignmentTaskActionRequest) GetAppInstanceIds() []string {
	if o == nil || IsNil(o.AppInstanceIds) {
		var ret []string
		return ret
	}
	return o.AppInstanceIds
}

// GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAssignmentTaskActionRequest) GetAppInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppInstanceIds) {
		return nil, false
	}
	return o.AppInstanceIds, true
}

// HasAppInstanceIds returns a boolean if a field has been set.
func (o *AppAssignmentTaskActionRequest) HasAppInstanceIds() bool {
	if o != nil && !IsNil(o.AppInstanceIds) {
		return true
	}

	return false
}

// SetAppInstanceIds gets a reference to the given []string and assigns it to the AppInstanceIds field.
func (o *AppAssignmentTaskActionRequest) SetAppInstanceIds(v []string) {
	o.AppInstanceIds = v
}

// GetTaskIds returns the TaskIds field value if set, zero value otherwise.
func (o *AppAssignmentTaskActionRequest) GetTaskIds() []string {
	if o == nil || IsNil(o.TaskIds) {
		var ret []string
		return ret
	}
	return o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAssignmentTaskActionRequest) GetTaskIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TaskIds) {
		return nil, false
	}
	return o.TaskIds, true
}

// HasTaskIds returns a boolean if a field has been set.
func (o *AppAssignmentTaskActionRequest) HasTaskIds() bool {
	if o != nil && !IsNil(o.TaskIds) {
		return true
	}

	return false
}

// SetTaskIds gets a reference to the given []string and assigns it to the TaskIds field.
func (o *AppAssignmentTaskActionRequest) SetTaskIds(v []string) {
	o.TaskIds = v
}

func (o AppAssignmentTaskActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAssignmentTaskActionRequest) ToMap() (map[string]interface{}, error) {
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

func (o *AppAssignmentTaskActionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppAssignmentTaskActionRequest := _AppAssignmentTaskActionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppAssignmentTaskActionRequest)

	if err != nil {
		return err
	}

	*o = AppAssignmentTaskActionRequest(varAppAssignmentTaskActionRequest)

	return err
}

type NullableAppAssignmentTaskActionRequest struct {
	value *AppAssignmentTaskActionRequest
	isSet bool
}

func (v NullableAppAssignmentTaskActionRequest) Get() *AppAssignmentTaskActionRequest {
	return v.value
}

func (v *NullableAppAssignmentTaskActionRequest) Set(val *AppAssignmentTaskActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAssignmentTaskActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAssignmentTaskActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAssignmentTaskActionRequest(val *AppAssignmentTaskActionRequest) *NullableAppAssignmentTaskActionRequest {
	return &NullableAppAssignmentTaskActionRequest{value: val, isSet: true}
}

func (v NullableAppAssignmentTaskActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAssignmentTaskActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

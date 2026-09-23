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
	"time"
)

// checks if the ScriptInvokeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptInvokeResponse{}

// ScriptInvokeResponse Response returned when a remote script is successfully invoked.
type ScriptInvokeResponse struct {
	// Process ID of the launched process on the agent host
	Pid *string `json:"pid,omitempty"`
	// ISO-8601 timestamp of when the process was started
	StartTime            *time.Time `json:"startTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScriptInvokeResponse ScriptInvokeResponse

// NewScriptInvokeResponse instantiates a new ScriptInvokeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptInvokeResponse() *ScriptInvokeResponse {
	this := ScriptInvokeResponse{}
	return &this
}

// NewScriptInvokeResponseWithDefaults instantiates a new ScriptInvokeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptInvokeResponseWithDefaults() *ScriptInvokeResponse {
	this := ScriptInvokeResponse{}
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *ScriptInvokeResponse) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptInvokeResponse) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *ScriptInvokeResponse) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *ScriptInvokeResponse) SetPid(v string) {
	o.Pid = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ScriptInvokeResponse) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptInvokeResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ScriptInvokeResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ScriptInvokeResponse) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o ScriptInvokeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptInvokeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScriptInvokeResponse) UnmarshalJSON(data []byte) (err error) {
	varScriptInvokeResponse := _ScriptInvokeResponse{}

	err = json.Unmarshal(data, &varScriptInvokeResponse)

	if err != nil {
		return err
	}

	*o = ScriptInvokeResponse(varScriptInvokeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pid")
		delete(additionalProperties, "startTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScriptInvokeResponse struct {
	value *ScriptInvokeResponse
	isSet bool
}

func (v NullableScriptInvokeResponse) Get() *ScriptInvokeResponse {
	return v.value
}

func (v *NullableScriptInvokeResponse) Set(val *ScriptInvokeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptInvokeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptInvokeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptInvokeResponse(val *ScriptInvokeResponse) *NullableScriptInvokeResponse {
	return &NullableScriptInvokeResponse{value: val, isSet: true}
}

func (v NullableScriptInvokeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptInvokeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ScheduledUserLifecycleAction struct for ScheduledUserLifecycleAction
type ScheduledUserLifecycleAction struct {
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledUserLifecycleAction ScheduledUserLifecycleAction

// NewScheduledUserLifecycleAction instantiates a new ScheduledUserLifecycleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledUserLifecycleAction() *ScheduledUserLifecycleAction {
	this := ScheduledUserLifecycleAction{}
	return &this
}

// NewScheduledUserLifecycleActionWithDefaults instantiates a new ScheduledUserLifecycleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledUserLifecycleActionWithDefaults() *ScheduledUserLifecycleAction {
	this := ScheduledUserLifecycleAction{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScheduledUserLifecycleAction) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledUserLifecycleAction) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScheduledUserLifecycleAction) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScheduledUserLifecycleAction) SetStatus(v string) {
	o.Status = &v
}

func (o ScheduledUserLifecycleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduledUserLifecycleAction) UnmarshalJSON(bytes []byte) (err error) {
	varScheduledUserLifecycleAction := _ScheduledUserLifecycleAction{}

	err = json.Unmarshal(bytes, &varScheduledUserLifecycleAction)
	if err == nil {
		*o = ScheduledUserLifecycleAction(varScheduledUserLifecycleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableScheduledUserLifecycleAction struct {
	value *ScheduledUserLifecycleAction
	isSet bool
}

func (v NullableScheduledUserLifecycleAction) Get() *ScheduledUserLifecycleAction {
	return v.value
}

func (v *NullableScheduledUserLifecycleAction) Set(val *ScheduledUserLifecycleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledUserLifecycleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledUserLifecycleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledUserLifecycleAction(val *ScheduledUserLifecycleAction) *NullableScheduledUserLifecycleAction {
	return &NullableScheduledUserLifecycleAction{value: val, isSet: true}
}

func (v NullableScheduledUserLifecycleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledUserLifecycleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


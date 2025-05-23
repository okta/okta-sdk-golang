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

// UpdateRealmAssignmentRequest struct for UpdateRealmAssignmentRequest
type UpdateRealmAssignmentRequest struct {
	Actions *Actions `json:"actions,omitempty"`
	Conditions *Conditions `json:"conditions,omitempty"`
	Name *string `json:"name,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRealmAssignmentRequest UpdateRealmAssignmentRequest

// NewUpdateRealmAssignmentRequest instantiates a new UpdateRealmAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRealmAssignmentRequest() *UpdateRealmAssignmentRequest {
	this := UpdateRealmAssignmentRequest{}
	return &this
}

// NewUpdateRealmAssignmentRequestWithDefaults instantiates a new UpdateRealmAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRealmAssignmentRequestWithDefaults() *UpdateRealmAssignmentRequest {
	this := UpdateRealmAssignmentRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *UpdateRealmAssignmentRequest) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRealmAssignmentRequest) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *UpdateRealmAssignmentRequest) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *UpdateRealmAssignmentRequest) SetActions(v Actions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *UpdateRealmAssignmentRequest) GetConditions() Conditions {
	if o == nil || o.Conditions == nil {
		var ret Conditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRealmAssignmentRequest) GetConditionsOk() (*Conditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *UpdateRealmAssignmentRequest) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given Conditions and assigns it to the Conditions field.
func (o *UpdateRealmAssignmentRequest) SetConditions(v Conditions) {
	o.Conditions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRealmAssignmentRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRealmAssignmentRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRealmAssignmentRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRealmAssignmentRequest) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *UpdateRealmAssignmentRequest) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRealmAssignmentRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *UpdateRealmAssignmentRequest) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *UpdateRealmAssignmentRequest) SetPriority(v int32) {
	o.Priority = &v
}

func (o UpdateRealmAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateRealmAssignmentRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateRealmAssignmentRequest := _UpdateRealmAssignmentRequest{}

	err = json.Unmarshal(bytes, &varUpdateRealmAssignmentRequest)
	if err == nil {
		*o = UpdateRealmAssignmentRequest(varUpdateRealmAssignmentRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUpdateRealmAssignmentRequest struct {
	value *UpdateRealmAssignmentRequest
	isSet bool
}

func (v NullableUpdateRealmAssignmentRequest) Get() *UpdateRealmAssignmentRequest {
	return v.value
}

func (v *NullableUpdateRealmAssignmentRequest) Set(val *UpdateRealmAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRealmAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRealmAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRealmAssignmentRequest(val *UpdateRealmAssignmentRequest) *NullableUpdateRealmAssignmentRequest {
	return &NullableUpdateRealmAssignmentRequest{value: val, isSet: true}
}

func (v NullableUpdateRealmAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRealmAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


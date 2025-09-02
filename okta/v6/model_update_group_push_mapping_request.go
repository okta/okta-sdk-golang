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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// UpdateGroupPushMappingRequest struct for UpdateGroupPushMappingRequest
type UpdateGroupPushMappingRequest struct {
	// The status of the group push mapping
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _UpdateGroupPushMappingRequest UpdateGroupPushMappingRequest

// NewUpdateGroupPushMappingRequest instantiates a new UpdateGroupPushMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupPushMappingRequest(status string) *UpdateGroupPushMappingRequest {
	this := UpdateGroupPushMappingRequest{}
	this.Status = status
	return &this
}

// NewUpdateGroupPushMappingRequestWithDefaults instantiates a new UpdateGroupPushMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupPushMappingRequestWithDefaults() *UpdateGroupPushMappingRequest {
	this := UpdateGroupPushMappingRequest{}
	var status string = "ACTIVE"
	this.Status = status
	return &this
}

// GetStatus returns the Status field value
func (o *UpdateGroupPushMappingRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupPushMappingRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateGroupPushMappingRequest) SetStatus(v string) {
	o.Status = v
}

func (o UpdateGroupPushMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateGroupPushMappingRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateGroupPushMappingRequest := _UpdateGroupPushMappingRequest{}

	err = json.Unmarshal(bytes, &varUpdateGroupPushMappingRequest)
	if err == nil {
		*o = UpdateGroupPushMappingRequest(varUpdateGroupPushMappingRequest)
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

type NullableUpdateGroupPushMappingRequest struct {
	value *UpdateGroupPushMappingRequest
	isSet bool
}

func (v NullableUpdateGroupPushMappingRequest) Get() *UpdateGroupPushMappingRequest {
	return v.value
}

func (v *NullableUpdateGroupPushMappingRequest) Set(val *UpdateGroupPushMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupPushMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupPushMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupPushMappingRequest(val *UpdateGroupPushMappingRequest) *NullableUpdateGroupPushMappingRequest {
	return &NullableUpdateGroupPushMappingRequest{value: val, isSet: true}
}

func (v NullableUpdateGroupPushMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupPushMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


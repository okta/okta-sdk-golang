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
	"fmt"
)

// checks if the UpdateGroupPushMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupPushMappingRequest{}

// UpdateGroupPushMappingRequest struct for UpdateGroupPushMappingRequest
type UpdateGroupPushMappingRequest struct {
	// The status of the group push mapping
	Status               string `json:"status"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupPushMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateGroupPushMappingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varUpdateGroupPushMappingRequest := _UpdateGroupPushMappingRequest{}

	err = json.Unmarshal(data, &varUpdateGroupPushMappingRequest)

	if err != nil {
		return err
	}

	*o = UpdateGroupPushMappingRequest(varUpdateGroupPushMappingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
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

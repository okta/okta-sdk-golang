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
	"fmt"
)

// checks if the ExecutionAcceptedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionAcceptedResponse{}

// ExecutionAcceptedResponse struct for ExecutionAcceptedResponse
type ExecutionAcceptedResponse struct {
	// ID used to retrieve the execution status
	ExecutionId string `json:"executionId"`
	// Status of the accepted execution
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionAcceptedResponse ExecutionAcceptedResponse

// NewExecutionAcceptedResponse instantiates a new ExecutionAcceptedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionAcceptedResponse(executionId string, status string) *ExecutionAcceptedResponse {
	this := ExecutionAcceptedResponse{}
	this.ExecutionId = executionId
	this.Status = status
	return &this
}

// NewExecutionAcceptedResponseWithDefaults instantiates a new ExecutionAcceptedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionAcceptedResponseWithDefaults() *ExecutionAcceptedResponse {
	this := ExecutionAcceptedResponse{}
	return &this
}

// GetExecutionId returns the ExecutionId field value
func (o *ExecutionAcceptedResponse) GetExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value
// and a boolean to check if the value has been set.
func (o *ExecutionAcceptedResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionId, true
}

// SetExecutionId sets field value
func (o *ExecutionAcceptedResponse) SetExecutionId(v string) {
	o.ExecutionId = v
}

// GetStatus returns the Status field value
func (o *ExecutionAcceptedResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExecutionAcceptedResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExecutionAcceptedResponse) SetStatus(v string) {
	o.Status = v
}

func (o ExecutionAcceptedResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionAcceptedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["executionId"] = o.ExecutionId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionAcceptedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"executionId",
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

	varExecutionAcceptedResponse := _ExecutionAcceptedResponse{}

	err = json.Unmarshal(data, &varExecutionAcceptedResponse)

	if err != nil {
		return err
	}

	*o = ExecutionAcceptedResponse(varExecutionAcceptedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionAcceptedResponse struct {
	value *ExecutionAcceptedResponse
	isSet bool
}

func (v NullableExecutionAcceptedResponse) Get() *ExecutionAcceptedResponse {
	return v.value
}

func (v *NullableExecutionAcceptedResponse) Set(val *ExecutionAcceptedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionAcceptedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionAcceptedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionAcceptedResponse(val *ExecutionAcceptedResponse) *NullableExecutionAcceptedResponse {
	return &NullableExecutionAcceptedResponse{value: val, isSet: true}
}

func (v NullableExecutionAcceptedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionAcceptedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

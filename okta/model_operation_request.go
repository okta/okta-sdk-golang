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

// OperationRequest struct for OperationRequest
type OperationRequest struct {
	AssignmentId *string `json:"assignmentId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationRequest OperationRequest

// NewOperationRequest instantiates a new OperationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationRequest() *OperationRequest {
	this := OperationRequest{}
	return &this
}

// NewOperationRequestWithDefaults instantiates a new OperationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationRequestWithDefaults() *OperationRequest {
	this := OperationRequest{}
	return &this
}

// GetAssignmentId returns the AssignmentId field value if set, zero value otherwise.
func (o *OperationRequest) GetAssignmentId() string {
	if o == nil || o.AssignmentId == nil {
		var ret string
		return ret
	}
	return *o.AssignmentId
}

// GetAssignmentIdOk returns a tuple with the AssignmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationRequest) GetAssignmentIdOk() (*string, bool) {
	if o == nil || o.AssignmentId == nil {
		return nil, false
	}
	return o.AssignmentId, true
}

// HasAssignmentId returns a boolean if a field has been set.
func (o *OperationRequest) HasAssignmentId() bool {
	if o != nil && o.AssignmentId != nil {
		return true
	}

	return false
}

// SetAssignmentId gets a reference to the given string and assigns it to the AssignmentId field.
func (o *OperationRequest) SetAssignmentId(v string) {
	o.AssignmentId = &v
}

func (o OperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignmentId != nil {
		toSerialize["assignmentId"] = o.AssignmentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varOperationRequest := _OperationRequest{}

	err = json.Unmarshal(bytes, &varOperationRequest)
	if err == nil {
		*o = OperationRequest(varOperationRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignmentId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOperationRequest struct {
	value *OperationRequest
	isSet bool
}

func (v NullableOperationRequest) Get() *OperationRequest {
	return v.value
}

func (v *NullableOperationRequest) Set(val *OperationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationRequest(val *OperationRequest) *NullableOperationRequest {
	return &NullableOperationRequest{value: val, isSet: true}
}

func (v NullableOperationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateIamRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIamRoleRequest{}

// UpdateIamRoleRequest struct for UpdateIamRoleRequest
type UpdateIamRoleRequest struct {
	// Description of the role
	Description string `json:"description"`
	// Unique label for the role
	Label                string `json:"label"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIamRoleRequest UpdateIamRoleRequest

// NewUpdateIamRoleRequest instantiates a new UpdateIamRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIamRoleRequest(description string, label string) *UpdateIamRoleRequest {
	this := UpdateIamRoleRequest{}
	this.Description = description
	this.Label = label
	return &this
}

// NewUpdateIamRoleRequestWithDefaults instantiates a new UpdateIamRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIamRoleRequestWithDefaults() *UpdateIamRoleRequest {
	this := UpdateIamRoleRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *UpdateIamRoleRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateIamRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateIamRoleRequest) SetDescription(v string) {
	o.Description = v
}

// GetLabel returns the Label field value
func (o *UpdateIamRoleRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *UpdateIamRoleRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *UpdateIamRoleRequest) SetLabel(v string) {
	o.Label = v
}

func (o UpdateIamRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIamRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["label"] = o.Label

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIamRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"label",
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

	varUpdateIamRoleRequest := _UpdateIamRoleRequest{}

	err = json.Unmarshal(data, &varUpdateIamRoleRequest)

	if err != nil {
		return err
	}

	*o = UpdateIamRoleRequest(varUpdateIamRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIamRoleRequest struct {
	value *UpdateIamRoleRequest
	isSet bool
}

func (v NullableUpdateIamRoleRequest) Get() *UpdateIamRoleRequest {
	return v.value
}

func (v *NullableUpdateIamRoleRequest) Set(val *UpdateIamRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIamRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIamRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIamRoleRequest(val *UpdateIamRoleRequest) *NullableUpdateIamRoleRequest {
	return &NullableUpdateIamRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateIamRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIamRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

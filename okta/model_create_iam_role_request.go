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

// checks if the CreateIamRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIamRoleRequest{}

// CreateIamRoleRequest struct for CreateIamRoleRequest
type CreateIamRoleRequest struct {
	// Description of the role
	Description string `json:"description"`
	// Unique label for the role
	Label string `json:"label"`
	// Array of permissions that the role grants. See [Permissions](/openapi/okta-management/guides/permissions).
	Permissions          []string `json:"permissions"`
	AdditionalProperties map[string]interface{}
}

type _CreateIamRoleRequest CreateIamRoleRequest

// NewCreateIamRoleRequest instantiates a new CreateIamRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIamRoleRequest(description string, label string, permissions []string) *CreateIamRoleRequest {
	this := CreateIamRoleRequest{}
	this.Description = description
	this.Label = label
	this.Permissions = permissions
	return &this
}

// NewCreateIamRoleRequestWithDefaults instantiates a new CreateIamRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIamRoleRequestWithDefaults() *CreateIamRoleRequest {
	this := CreateIamRoleRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateIamRoleRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateIamRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateIamRoleRequest) SetDescription(v string) {
	o.Description = v
}

// GetLabel returns the Label field value
func (o *CreateIamRoleRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateIamRoleRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateIamRoleRequest) SetLabel(v string) {
	o.Label = v
}

// GetPermissions returns the Permissions field value
func (o *CreateIamRoleRequest) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *CreateIamRoleRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *CreateIamRoleRequest) SetPermissions(v []string) {
	o.Permissions = v
}

func (o CreateIamRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIamRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["label"] = o.Label
	toSerialize["permissions"] = o.Permissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateIamRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"label",
		"permissions",
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

	varCreateIamRoleRequest := _CreateIamRoleRequest{}

	err = json.Unmarshal(data, &varCreateIamRoleRequest)

	if err != nil {
		return err
	}

	*o = CreateIamRoleRequest(varCreateIamRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateIamRoleRequest struct {
	value *CreateIamRoleRequest
	isSet bool
}

func (v NullableCreateIamRoleRequest) Get() *CreateIamRoleRequest {
	return v.value
}

func (v *NullableCreateIamRoleRequest) Set(val *CreateIamRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIamRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIamRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIamRoleRequest(val *CreateIamRoleRequest) *NullableCreateIamRoleRequest {
	return &NullableCreateIamRoleRequest{value: val, isSet: true}
}

func (v NullableCreateIamRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIamRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

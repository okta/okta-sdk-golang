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

// CreateUpdateIamRolePermissionRequest struct for CreateUpdateIamRolePermissionRequest
type CreateUpdateIamRolePermissionRequest struct {
	// Conditions for further restricting a permission
	Conditions map[string]interface{} `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUpdateIamRolePermissionRequest CreateUpdateIamRolePermissionRequest

// NewCreateUpdateIamRolePermissionRequest instantiates a new CreateUpdateIamRolePermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateIamRolePermissionRequest() *CreateUpdateIamRolePermissionRequest {
	this := CreateUpdateIamRolePermissionRequest{}
	return &this
}

// NewCreateUpdateIamRolePermissionRequestWithDefaults instantiates a new CreateUpdateIamRolePermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateIamRolePermissionRequestWithDefaults() *CreateUpdateIamRolePermissionRequest {
	this := CreateUpdateIamRolePermissionRequest{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUpdateIamRolePermissionRequest) GetConditions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUpdateIamRolePermissionRequest) GetConditionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CreateUpdateIamRolePermissionRequest) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given map[string]interface{} and assigns it to the Conditions field.
func (o *CreateUpdateIamRolePermissionRequest) SetConditions(v map[string]interface{}) {
	o.Conditions = v
}

func (o CreateUpdateIamRolePermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUpdateIamRolePermissionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUpdateIamRolePermissionRequest := _CreateUpdateIamRolePermissionRequest{}

	err = json.Unmarshal(bytes, &varCreateUpdateIamRolePermissionRequest)
	if err == nil {
		*o = CreateUpdateIamRolePermissionRequest(varCreateUpdateIamRolePermissionRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateUpdateIamRolePermissionRequest struct {
	value *CreateUpdateIamRolePermissionRequest
	isSet bool
}

func (v NullableCreateUpdateIamRolePermissionRequest) Get() *CreateUpdateIamRolePermissionRequest {
	return v.value
}

func (v *NullableCreateUpdateIamRolePermissionRequest) Set(val *CreateUpdateIamRolePermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateIamRolePermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateIamRolePermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateIamRolePermissionRequest(val *CreateUpdateIamRolePermissionRequest) *NullableCreateUpdateIamRolePermissionRequest {
	return &NullableCreateUpdateIamRolePermissionRequest{value: val, isSet: true}
}

func (v NullableCreateUpdateIamRolePermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateIamRolePermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


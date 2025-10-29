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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ResourceSetResourcePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetResourcePutRequest{}

// ResourceSetResourcePutRequest struct for ResourceSetResourcePutRequest
type ResourceSetResourcePutRequest struct {
	Conditions           *ResourceConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResourcePutRequest ResourceSetResourcePutRequest

// NewResourceSetResourcePutRequest instantiates a new ResourceSetResourcePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResourcePutRequest() *ResourceSetResourcePutRequest {
	this := ResourceSetResourcePutRequest{}
	return &this
}

// NewResourceSetResourcePutRequestWithDefaults instantiates a new ResourceSetResourcePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourcePutRequestWithDefaults() *ResourceSetResourcePutRequest {
	this := ResourceSetResourcePutRequest{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ResourceSetResourcePutRequest) GetConditions() ResourceConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret ResourceConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcePutRequest) GetConditionsOk() (*ResourceConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ResourceSetResourcePutRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given ResourceConditions and assigns it to the Conditions field.
func (o *ResourceSetResourcePutRequest) SetConditions(v ResourceConditions) {
	o.Conditions = &v
}

func (o ResourceSetResourcePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetResourcePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetResourcePutRequest) UnmarshalJSON(data []byte) (err error) {
	varResourceSetResourcePutRequest := _ResourceSetResourcePutRequest{}

	err = json.Unmarshal(data, &varResourceSetResourcePutRequest)

	if err != nil {
		return err
	}

	*o = ResourceSetResourcePutRequest(varResourceSetResourcePutRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetResourcePutRequest struct {
	value *ResourceSetResourcePutRequest
	isSet bool
}

func (v NullableResourceSetResourcePutRequest) Get() *ResourceSetResourcePutRequest {
	return v.value
}

func (v *NullableResourceSetResourcePutRequest) Set(val *ResourceSetResourcePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResourcePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResourcePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResourcePutRequest(val *ResourceSetResourcePutRequest) *NullableResourceSetResourcePutRequest {
	return &NullableResourceSetResourcePutRequest{value: val, isSet: true}
}

func (v NullableResourceSetResourcePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResourcePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

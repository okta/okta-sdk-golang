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

// checks if the CimdClientEntityBindingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntityBindingRequest{}

// CimdClientEntityBindingRequest Request body for creating a CIMD client entity binding
type CimdClientEntityBindingRequest struct {
	// Unique identifier of the entity to bind to the CIMD client entity
	BoundEntityId string `json:"boundEntityId"`
	// The type of resource that's bound to the CIMD client entity
	BoundEntityType      string `json:"boundEntityType"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntityBindingRequest CimdClientEntityBindingRequest

// NewCimdClientEntityBindingRequest instantiates a new CimdClientEntityBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntityBindingRequest(boundEntityId string, boundEntityType string) *CimdClientEntityBindingRequest {
	this := CimdClientEntityBindingRequest{}
	this.BoundEntityId = boundEntityId
	this.BoundEntityType = boundEntityType
	return &this
}

// NewCimdClientEntityBindingRequestWithDefaults instantiates a new CimdClientEntityBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityBindingRequestWithDefaults() *CimdClientEntityBindingRequest {
	this := CimdClientEntityBindingRequest{}
	return &this
}

// GetBoundEntityId returns the BoundEntityId field value
func (o *CimdClientEntityBindingRequest) GetBoundEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoundEntityId
}

// GetBoundEntityIdOk returns a tuple with the BoundEntityId field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBindingRequest) GetBoundEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundEntityId, true
}

// SetBoundEntityId sets field value
func (o *CimdClientEntityBindingRequest) SetBoundEntityId(v string) {
	o.BoundEntityId = v
}

// GetBoundEntityType returns the BoundEntityType field value
func (o *CimdClientEntityBindingRequest) GetBoundEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoundEntityType
}

// GetBoundEntityTypeOk returns a tuple with the BoundEntityType field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBindingRequest) GetBoundEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundEntityType, true
}

// SetBoundEntityType sets field value
func (o *CimdClientEntityBindingRequest) SetBoundEntityType(v string) {
	o.BoundEntityType = v
}

func (o CimdClientEntityBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntityBindingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boundEntityId"] = o.BoundEntityId
	toSerialize["boundEntityType"] = o.BoundEntityType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdClientEntityBindingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boundEntityId",
		"boundEntityType",
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

	varCimdClientEntityBindingRequest := _CimdClientEntityBindingRequest{}

	err = json.Unmarshal(data, &varCimdClientEntityBindingRequest)

	if err != nil {
		return err
	}

	*o = CimdClientEntityBindingRequest(varCimdClientEntityBindingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "boundEntityId")
		delete(additionalProperties, "boundEntityType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntityBindingRequest struct {
	value *CimdClientEntityBindingRequest
	isSet bool
}

func (v NullableCimdClientEntityBindingRequest) Get() *CimdClientEntityBindingRequest {
	return v.value
}

func (v *NullableCimdClientEntityBindingRequest) Set(val *CimdClientEntityBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntityBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntityBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntityBindingRequest(val *CimdClientEntityBindingRequest) *NullableCimdClientEntityBindingRequest {
	return &NullableCimdClientEntityBindingRequest{value: val, isSet: true}
}

func (v NullableCimdClientEntityBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntityBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

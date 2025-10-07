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

// checks if the ProfileMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileMappingRequest{}

// ProfileMappingRequest The updated request body properties
type ProfileMappingRequest struct {
	Properties           ProfileMappingProperty `json:"properties"`
	AdditionalProperties map[string]interface{}
}

type _ProfileMappingRequest ProfileMappingRequest

// NewProfileMappingRequest instantiates a new ProfileMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMappingRequest(properties ProfileMappingProperty) *ProfileMappingRequest {
	this := ProfileMappingRequest{}
	this.Properties = properties
	return &this
}

// NewProfileMappingRequestWithDefaults instantiates a new ProfileMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMappingRequestWithDefaults() *ProfileMappingRequest {
	this := ProfileMappingRequest{}
	return &this
}

// GetProperties returns the Properties field value
func (o *ProfileMappingRequest) GetProperties() ProfileMappingProperty {
	if o == nil {
		var ret ProfileMappingProperty
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ProfileMappingRequest) GetPropertiesOk() (ProfileMappingProperty, bool) {
	if o == nil {
		return ProfileMappingProperty{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *ProfileMappingRequest) SetProperties(v ProfileMappingProperty) {
	o.Properties = v
}

func (o ProfileMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileMappingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"properties",
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

	varProfileMappingRequest := _ProfileMappingRequest{}

	err = json.Unmarshal(data, &varProfileMappingRequest)

	if err != nil {
		return err
	}

	*o = ProfileMappingRequest(varProfileMappingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileMappingRequest struct {
	value *ProfileMappingRequest
	isSet bool
}

func (v NullableProfileMappingRequest) Get() *ProfileMappingRequest {
	return v.value
}

func (v *NullableProfileMappingRequest) Set(val *ProfileMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMappingRequest(val *ProfileMappingRequest) *NullableProfileMappingRequest {
	return &NullableProfileMappingRequest{value: val, isSet: true}
}

func (v NullableProfileMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

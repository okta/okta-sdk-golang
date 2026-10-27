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

// checks if the DeviceCapabilitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCapabilitiesResponse{}

// DeviceCapabilitiesResponse struct for DeviceCapabilitiesResponse
type DeviceCapabilitiesResponse struct {
	// One entry per capability. An entry is omitted entirely if your org's plan doesn't include that capability, or if the device's platform doesn't support it.
	Capabilities         []DeviceCapability `json:"capabilities"`
	AdditionalProperties map[string]interface{}
}

type _DeviceCapabilitiesResponse DeviceCapabilitiesResponse

// NewDeviceCapabilitiesResponse instantiates a new DeviceCapabilitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCapabilitiesResponse(capabilities []DeviceCapability) *DeviceCapabilitiesResponse {
	this := DeviceCapabilitiesResponse{}
	this.Capabilities = capabilities
	return &this
}

// NewDeviceCapabilitiesResponseWithDefaults instantiates a new DeviceCapabilitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCapabilitiesResponseWithDefaults() *DeviceCapabilitiesResponse {
	this := DeviceCapabilitiesResponse{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *DeviceCapabilitiesResponse) GetCapabilities() []DeviceCapability {
	if o == nil {
		var ret []DeviceCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *DeviceCapabilitiesResponse) GetCapabilitiesOk() ([]DeviceCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *DeviceCapabilitiesResponse) SetCapabilities(v []DeviceCapability) {
	o.Capabilities = v
}

func (o DeviceCapabilitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCapabilitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capabilities"] = o.Capabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceCapabilitiesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capabilities",
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

	varDeviceCapabilitiesResponse := _DeviceCapabilitiesResponse{}

	err = json.Unmarshal(data, &varDeviceCapabilitiesResponse)

	if err != nil {
		return err
	}

	*o = DeviceCapabilitiesResponse(varDeviceCapabilitiesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceCapabilitiesResponse struct {
	value *DeviceCapabilitiesResponse
	isSet bool
}

func (v NullableDeviceCapabilitiesResponse) Get() *DeviceCapabilitiesResponse {
	return v.value
}

func (v *NullableDeviceCapabilitiesResponse) Set(val *DeviceCapabilitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCapabilitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCapabilitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCapabilitiesResponse(val *DeviceCapabilitiesResponse) *NullableDeviceCapabilitiesResponse {
	return &NullableDeviceCapabilitiesResponse{value: val, isSet: true}
}

func (v NullableDeviceCapabilitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCapabilitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

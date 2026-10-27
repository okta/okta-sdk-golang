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

// checks if the DeviceCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCapability{}

// DeviceCapability struct for DeviceCapability
type DeviceCapability struct {
	// Capability identifier
	Capability string `json:"capability"`
	// List of possible reason codes for unmet preconditions. It's only present when the `status` is `NOT_READY`.
	Reasons []string `json:"reasons,omitempty"`
	// Readiness of the capability
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DeviceCapability DeviceCapability

// NewDeviceCapability instantiates a new DeviceCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCapability(capability string, status string) *DeviceCapability {
	this := DeviceCapability{}
	this.Capability = capability
	this.Status = status
	return &this
}

// NewDeviceCapabilityWithDefaults instantiates a new DeviceCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCapabilityWithDefaults() *DeviceCapability {
	this := DeviceCapability{}
	return &this
}

// GetCapability returns the Capability field value
func (o *DeviceCapability) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *DeviceCapability) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *DeviceCapability) SetCapability(v string) {
	o.Capability = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *DeviceCapability) GetReasons() []string {
	if o == nil || IsNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCapability) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *DeviceCapability) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *DeviceCapability) SetReasons(v []string) {
	o.Reasons = v
}

// GetStatus returns the Status field value
func (o *DeviceCapability) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeviceCapability) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeviceCapability) SetStatus(v string) {
	o.Status = v
}

func (o DeviceCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capability"] = o.Capability
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceCapability) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capability",
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

	varDeviceCapability := _DeviceCapability{}

	err = json.Unmarshal(data, &varDeviceCapability)

	if err != nil {
		return err
	}

	*o = DeviceCapability(varDeviceCapability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capability")
		delete(additionalProperties, "reasons")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceCapability struct {
	value *DeviceCapability
	isSet bool
}

func (v NullableDeviceCapability) Get() *DeviceCapability {
	return v.value
}

func (v *NullableDeviceCapability) Set(val *DeviceCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCapability(val *DeviceCapability) *NullableDeviceCapability {
	return &NullableDeviceCapability{value: val, isSet: true}
}

func (v NullableDeviceCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the DevicePreRegistrationEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePreRegistrationEntry{}

// DevicePreRegistrationEntry struct for DevicePreRegistrationEntry
type DevicePreRegistrationEntry struct {
	// Device platform for pre-registration
	Platform string `json:"platform"`
	// Registration grants to associate with this device
	RegistrationGrants []RegistrationGrantRequest `json:"registrationGrants"`
	// The serial number of the device
	SerialNumber         string `json:"serialNumber"`
	AdditionalProperties map[string]interface{}
}

type _DevicePreRegistrationEntry DevicePreRegistrationEntry

// NewDevicePreRegistrationEntry instantiates a new DevicePreRegistrationEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePreRegistrationEntry(platform string, registrationGrants []RegistrationGrantRequest, serialNumber string) *DevicePreRegistrationEntry {
	this := DevicePreRegistrationEntry{}
	this.Platform = platform
	this.RegistrationGrants = registrationGrants
	this.SerialNumber = serialNumber
	return &this
}

// NewDevicePreRegistrationEntryWithDefaults instantiates a new DevicePreRegistrationEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePreRegistrationEntryWithDefaults() *DevicePreRegistrationEntry {
	this := DevicePreRegistrationEntry{}
	return &this
}

// GetPlatform returns the Platform field value
func (o *DevicePreRegistrationEntry) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DevicePreRegistrationEntry) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *DevicePreRegistrationEntry) SetPlatform(v string) {
	o.Platform = v
}

// GetRegistrationGrants returns the RegistrationGrants field value
func (o *DevicePreRegistrationEntry) GetRegistrationGrants() []RegistrationGrantRequest {
	if o == nil {
		var ret []RegistrationGrantRequest
		return ret
	}

	return o.RegistrationGrants
}

// GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field value
// and a boolean to check if the value has been set.
func (o *DevicePreRegistrationEntry) GetRegistrationGrantsOk() ([]RegistrationGrantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationGrants, true
}

// SetRegistrationGrants sets field value
func (o *DevicePreRegistrationEntry) SetRegistrationGrants(v []RegistrationGrantRequest) {
	o.RegistrationGrants = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *DevicePreRegistrationEntry) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *DevicePreRegistrationEntry) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *DevicePreRegistrationEntry) SetSerialNumber(v string) {
	o.SerialNumber = v
}

func (o DevicePreRegistrationEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePreRegistrationEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["platform"] = o.Platform
	toSerialize["registrationGrants"] = o.RegistrationGrants
	toSerialize["serialNumber"] = o.SerialNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePreRegistrationEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"platform",
		"registrationGrants",
		"serialNumber",
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

	varDevicePreRegistrationEntry := _DevicePreRegistrationEntry{}

	err = json.Unmarshal(data, &varDevicePreRegistrationEntry)

	if err != nil {
		return err
	}

	*o = DevicePreRegistrationEntry(varDevicePreRegistrationEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "platform")
		delete(additionalProperties, "registrationGrants")
		delete(additionalProperties, "serialNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePreRegistrationEntry struct {
	value *DevicePreRegistrationEntry
	isSet bool
}

func (v NullableDevicePreRegistrationEntry) Get() *DevicePreRegistrationEntry {
	return v.value
}

func (v *NullableDevicePreRegistrationEntry) Set(val *DevicePreRegistrationEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePreRegistrationEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePreRegistrationEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePreRegistrationEntry(val *DevicePreRegistrationEntry) *NullableDevicePreRegistrationEntry {
	return &NullableDevicePreRegistrationEntry{value: val, isSet: true}
}

func (v NullableDevicePreRegistrationEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePreRegistrationEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

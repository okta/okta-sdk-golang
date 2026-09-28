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

// checks if the DeviceEnrollmentAuthenticatorProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceEnrollmentAuthenticatorProfile{}

// DeviceEnrollmentAuthenticatorProfile struct for DeviceEnrollmentAuthenticatorProfile
type DeviceEnrollmentAuthenticatorProfile struct {
	// The display name of the device
	DeviceName string `json:"deviceName"`
	// The unique identifier of the user associated with the enrollment
	UserId               string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _DeviceEnrollmentAuthenticatorProfile DeviceEnrollmentAuthenticatorProfile

// NewDeviceEnrollmentAuthenticatorProfile instantiates a new DeviceEnrollmentAuthenticatorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentAuthenticatorProfile(deviceName string, userId string) *DeviceEnrollmentAuthenticatorProfile {
	this := DeviceEnrollmentAuthenticatorProfile{}
	this.DeviceName = deviceName
	this.UserId = userId
	return &this
}

// NewDeviceEnrollmentAuthenticatorProfileWithDefaults instantiates a new DeviceEnrollmentAuthenticatorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentAuthenticatorProfileWithDefaults() *DeviceEnrollmentAuthenticatorProfile {
	this := DeviceEnrollmentAuthenticatorProfile{}
	return &this
}

// GetDeviceName returns the DeviceName field value
func (o *DeviceEnrollmentAuthenticatorProfile) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticatorProfile) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value
func (o *DeviceEnrollmentAuthenticatorProfile) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetUserId returns the UserId field value
func (o *DeviceEnrollmentAuthenticatorProfile) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticatorProfile) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeviceEnrollmentAuthenticatorProfile) SetUserId(v string) {
	o.UserId = v
}

func (o DeviceEnrollmentAuthenticatorProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceEnrollmentAuthenticatorProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceName"] = o.DeviceName
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceEnrollmentAuthenticatorProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceName",
		"userId",
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

	varDeviceEnrollmentAuthenticatorProfile := _DeviceEnrollmentAuthenticatorProfile{}

	err = json.Unmarshal(data, &varDeviceEnrollmentAuthenticatorProfile)

	if err != nil {
		return err
	}

	*o = DeviceEnrollmentAuthenticatorProfile(varDeviceEnrollmentAuthenticatorProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceName")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceEnrollmentAuthenticatorProfile struct {
	value *DeviceEnrollmentAuthenticatorProfile
	isSet bool
}

func (v NullableDeviceEnrollmentAuthenticatorProfile) Get() *DeviceEnrollmentAuthenticatorProfile {
	return v.value
}

func (v *NullableDeviceEnrollmentAuthenticatorProfile) Set(val *DeviceEnrollmentAuthenticatorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentAuthenticatorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentAuthenticatorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentAuthenticatorProfile(val *DeviceEnrollmentAuthenticatorProfile) *NullableDeviceEnrollmentAuthenticatorProfile {
	return &NullableDeviceEnrollmentAuthenticatorProfile{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentAuthenticatorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentAuthenticatorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

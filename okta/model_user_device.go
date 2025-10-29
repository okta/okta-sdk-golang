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
	"time"
)

// checks if the UserDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDevice{}

// UserDevice struct for UserDevice
type UserDevice struct {
	// Timestamp when the device was created
	Created *time.Time `json:"created,omitempty"`
	Device  Device     `json:"device,omitempty"`
	// Unique key for the user device link
	DeviceUserId         *string `json:"deviceUserId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserDevice UserDevice

// NewUserDevice instantiates a new UserDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDevice() *UserDevice {
	this := UserDevice{}
	return &this
}

// NewUserDeviceWithDefaults instantiates a new UserDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDeviceWithDefaults() *UserDevice {
	this := UserDevice{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserDevice) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDevice) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserDevice) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserDevice) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *UserDevice) GetDevice() Device {
	if o == nil || IsNil(o.Device) {
		var ret Device
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDevice) GetDeviceOk() (Device, bool) {
	if o == nil || IsNil(o.Device) {
		return Device{}, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *UserDevice) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *UserDevice) SetDevice(v Device) {
	o.Device = v
}

// GetDeviceUserId returns the DeviceUserId field value if set, zero value otherwise.
func (o *UserDevice) GetDeviceUserId() string {
	if o == nil || IsNil(o.DeviceUserId) {
		var ret string
		return ret
	}
	return *o.DeviceUserId
}

// GetDeviceUserIdOk returns a tuple with the DeviceUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDevice) GetDeviceUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceUserId) {
		return nil, false
	}
	return o.DeviceUserId, true
}

// HasDeviceUserId returns a boolean if a field has been set.
func (o *UserDevice) HasDeviceUserId() bool {
	if o != nil && !IsNil(o.DeviceUserId) {
		return true
	}

	return false
}

// SetDeviceUserId gets a reference to the given string and assigns it to the DeviceUserId field.
func (o *UserDevice) SetDeviceUserId(v string) {
	o.DeviceUserId = &v
}

func (o UserDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.DeviceUserId) {
		toSerialize["deviceUserId"] = o.DeviceUserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserDevice) UnmarshalJSON(data []byte) (err error) {
	varUserDevice := _UserDevice{}

	err = json.Unmarshal(data, &varUserDevice)

	if err != nil {
		return err
	}

	*o = UserDevice(varUserDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "device")
		delete(additionalProperties, "deviceUserId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserDevice struct {
	value *UserDevice
	isSet bool
}

func (v NullableUserDevice) Get() *UserDevice {
	return v.value
}

func (v *NullableUserDevice) Set(val *UserDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDevice(val *UserDevice) *NullableUserDevice {
	return &NullableUserDevice{value: val, isSet: true}
}

func (v NullableUserDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

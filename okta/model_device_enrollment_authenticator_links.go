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
)

// checks if the DeviceEnrollmentAuthenticatorLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceEnrollmentAuthenticatorLinks{}

// DeviceEnrollmentAuthenticatorLinks struct for DeviceEnrollmentAuthenticatorLinks
type DeviceEnrollmentAuthenticatorLinks struct {
	Device               *HrefObject `json:"device,omitempty"`
	Self                 *HrefObject `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceEnrollmentAuthenticatorLinks DeviceEnrollmentAuthenticatorLinks

// NewDeviceEnrollmentAuthenticatorLinks instantiates a new DeviceEnrollmentAuthenticatorLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentAuthenticatorLinks() *DeviceEnrollmentAuthenticatorLinks {
	this := DeviceEnrollmentAuthenticatorLinks{}
	return &this
}

// NewDeviceEnrollmentAuthenticatorLinksWithDefaults instantiates a new DeviceEnrollmentAuthenticatorLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentAuthenticatorLinksWithDefaults() *DeviceEnrollmentAuthenticatorLinks {
	this := DeviceEnrollmentAuthenticatorLinks{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DeviceEnrollmentAuthenticatorLinks) GetDevice() HrefObject {
	if o == nil || IsNil(o.Device) {
		var ret HrefObject
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticatorLinks) GetDeviceOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DeviceEnrollmentAuthenticatorLinks) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given HrefObject and assigns it to the Device field.
func (o *DeviceEnrollmentAuthenticatorLinks) SetDevice(v HrefObject) {
	o.Device = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DeviceEnrollmentAuthenticatorLinks) GetSelf() HrefObject {
	if o == nil || IsNil(o.Self) {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticatorLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DeviceEnrollmentAuthenticatorLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *DeviceEnrollmentAuthenticatorLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

func (o DeviceEnrollmentAuthenticatorLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceEnrollmentAuthenticatorLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceEnrollmentAuthenticatorLinks) UnmarshalJSON(data []byte) (err error) {
	varDeviceEnrollmentAuthenticatorLinks := _DeviceEnrollmentAuthenticatorLinks{}

	err = json.Unmarshal(data, &varDeviceEnrollmentAuthenticatorLinks)

	if err != nil {
		return err
	}

	*o = DeviceEnrollmentAuthenticatorLinks(varDeviceEnrollmentAuthenticatorLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceEnrollmentAuthenticatorLinks struct {
	value *DeviceEnrollmentAuthenticatorLinks
	isSet bool
}

func (v NullableDeviceEnrollmentAuthenticatorLinks) Get() *DeviceEnrollmentAuthenticatorLinks {
	return v.value
}

func (v *NullableDeviceEnrollmentAuthenticatorLinks) Set(val *DeviceEnrollmentAuthenticatorLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentAuthenticatorLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentAuthenticatorLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentAuthenticatorLinks(val *DeviceEnrollmentAuthenticatorLinks) *NullableDeviceEnrollmentAuthenticatorLinks {
	return &NullableDeviceEnrollmentAuthenticatorLinks{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentAuthenticatorLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentAuthenticatorLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the DeviceListAllOfEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceListAllOfEmbedded{}

// DeviceListAllOfEmbedded List of associated users for the device if the `expand=user` query parameter is specified in the request. Use `expand=userSummary` to get only a summary of each associated user for the device.
type DeviceListAllOfEmbedded struct {
	// Users for the device
	Users                []DeviceUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceListAllOfEmbedded DeviceListAllOfEmbedded

// NewDeviceListAllOfEmbedded instantiates a new DeviceListAllOfEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceListAllOfEmbedded() *DeviceListAllOfEmbedded {
	this := DeviceListAllOfEmbedded{}
	return &this
}

// NewDeviceListAllOfEmbeddedWithDefaults instantiates a new DeviceListAllOfEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceListAllOfEmbeddedWithDefaults() *DeviceListAllOfEmbedded {
	this := DeviceListAllOfEmbedded{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *DeviceListAllOfEmbedded) GetUsers() []DeviceUser {
	if o == nil || IsNil(o.Users) {
		var ret []DeviceUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceListAllOfEmbedded) GetUsersOk() ([]DeviceUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *DeviceListAllOfEmbedded) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []DeviceUser and assigns it to the Users field.
func (o *DeviceListAllOfEmbedded) SetUsers(v []DeviceUser) {
	o.Users = v
}

func (o DeviceListAllOfEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceListAllOfEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceListAllOfEmbedded) UnmarshalJSON(data []byte) (err error) {
	varDeviceListAllOfEmbedded := _DeviceListAllOfEmbedded{}

	err = json.Unmarshal(data, &varDeviceListAllOfEmbedded)

	if err != nil {
		return err
	}

	*o = DeviceListAllOfEmbedded(varDeviceListAllOfEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceListAllOfEmbedded struct {
	value *DeviceListAllOfEmbedded
	isSet bool
}

func (v NullableDeviceListAllOfEmbedded) Get() *DeviceListAllOfEmbedded {
	return v.value
}

func (v *NullableDeviceListAllOfEmbedded) Set(val *DeviceListAllOfEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceListAllOfEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceListAllOfEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceListAllOfEmbedded(val *DeviceListAllOfEmbedded) *NullableDeviceListAllOfEmbedded {
	return &NullableDeviceListAllOfEmbedded{value: val, isSet: true}
}

func (v NullableDeviceListAllOfEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceListAllOfEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

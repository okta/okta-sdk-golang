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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DeviceUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceUser{}

// DeviceUser struct for DeviceUser
type DeviceUser struct {
	// Timestamp when device was created
	Created *string `json:"created,omitempty"`
	// The management status of the device
	ManagementStatus *string `json:"managementStatus,omitempty"`
	// Screen lock type of the device
	ScreenLockType       *string `json:"screenLockType,omitempty"`
	User                 *User   `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceUser DeviceUser

// NewDeviceUser instantiates a new DeviceUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUser() *DeviceUser {
	this := DeviceUser{}
	return &this
}

// NewDeviceUserWithDefaults instantiates a new DeviceUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUserWithDefaults() *DeviceUser {
	this := DeviceUser{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeviceUser) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUser) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeviceUser) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *DeviceUser) SetCreated(v string) {
	o.Created = &v
}

// GetManagementStatus returns the ManagementStatus field value if set, zero value otherwise.
func (o *DeviceUser) GetManagementStatus() string {
	if o == nil || IsNil(o.ManagementStatus) {
		var ret string
		return ret
	}
	return *o.ManagementStatus
}

// GetManagementStatusOk returns a tuple with the ManagementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUser) GetManagementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementStatus) {
		return nil, false
	}
	return o.ManagementStatus, true
}

// HasManagementStatus returns a boolean if a field has been set.
func (o *DeviceUser) HasManagementStatus() bool {
	if o != nil && !IsNil(o.ManagementStatus) {
		return true
	}

	return false
}

// SetManagementStatus gets a reference to the given string and assigns it to the ManagementStatus field.
func (o *DeviceUser) SetManagementStatus(v string) {
	o.ManagementStatus = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *DeviceUser) GetScreenLockType() string {
	if o == nil || IsNil(o.ScreenLockType) {
		var ret string
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUser) GetScreenLockTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScreenLockType) {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *DeviceUser) HasScreenLockType() bool {
	if o != nil && !IsNil(o.ScreenLockType) {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given string and assigns it to the ScreenLockType field.
func (o *DeviceUser) SetScreenLockType(v string) {
	o.ScreenLockType = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DeviceUser) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUser) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DeviceUser) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *DeviceUser) SetUser(v User) {
	o.User = &v
}

func (o DeviceUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.ManagementStatus) {
		toSerialize["managementStatus"] = o.ManagementStatus
	}
	if !IsNil(o.ScreenLockType) {
		toSerialize["screenLockType"] = o.ScreenLockType
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceUser) UnmarshalJSON(data []byte) (err error) {
	varDeviceUser := _DeviceUser{}

	err = json.Unmarshal(data, &varDeviceUser)

	if err != nil {
		return err
	}

	*o = DeviceUser(varDeviceUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "managementStatus")
		delete(additionalProperties, "screenLockType")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceUser struct {
	value *DeviceUser
	isSet bool
}

func (v NullableDeviceUser) Get() *DeviceUser {
	return v.value
}

func (v *NullableDeviceUser) Set(val *DeviceUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUser(val *DeviceUser) *NullableDeviceUser {
	return &NullableDeviceUser{value: val, isSet: true}
}

func (v NullableDeviceUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

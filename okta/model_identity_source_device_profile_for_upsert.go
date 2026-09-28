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

// checks if the IdentitySourceDeviceProfileForUpsert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySourceDeviceProfileForUpsert{}

// IdentitySourceDeviceProfileForUpsert Device profile attributes for upsert operations within an identity source session
type IdentitySourceDeviceProfileForUpsert struct {
	// Display name of the device
	DisplayName string `json:"displayName"`
	// OS platform of the device
	Platform string `json:"platform"`
	// Serial number of the device
	SerialNumber         string `json:"serialNumber"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceDeviceProfileForUpsert IdentitySourceDeviceProfileForUpsert

// NewIdentitySourceDeviceProfileForUpsert instantiates a new IdentitySourceDeviceProfileForUpsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceDeviceProfileForUpsert(displayName string, platform string, serialNumber string) *IdentitySourceDeviceProfileForUpsert {
	this := IdentitySourceDeviceProfileForUpsert{}
	this.DisplayName = displayName
	this.Platform = platform
	this.SerialNumber = serialNumber
	return &this
}

// NewIdentitySourceDeviceProfileForUpsertWithDefaults instantiates a new IdentitySourceDeviceProfileForUpsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceDeviceProfileForUpsertWithDefaults() *IdentitySourceDeviceProfileForUpsert {
	this := IdentitySourceDeviceProfileForUpsert{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *IdentitySourceDeviceProfileForUpsert) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *IdentitySourceDeviceProfileForUpsert) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *IdentitySourceDeviceProfileForUpsert) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPlatform returns the Platform field value
func (o *IdentitySourceDeviceProfileForUpsert) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *IdentitySourceDeviceProfileForUpsert) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *IdentitySourceDeviceProfileForUpsert) SetPlatform(v string) {
	o.Platform = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *IdentitySourceDeviceProfileForUpsert) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *IdentitySourceDeviceProfileForUpsert) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *IdentitySourceDeviceProfileForUpsert) SetSerialNumber(v string) {
	o.SerialNumber = v
}

func (o IdentitySourceDeviceProfileForUpsert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySourceDeviceProfileForUpsert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["platform"] = o.Platform
	toSerialize["serialNumber"] = o.SerialNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySourceDeviceProfileForUpsert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"platform",
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

	varIdentitySourceDeviceProfileForUpsert := _IdentitySourceDeviceProfileForUpsert{}

	err = json.Unmarshal(data, &varIdentitySourceDeviceProfileForUpsert)

	if err != nil {
		return err
	}

	*o = IdentitySourceDeviceProfileForUpsert(varIdentitySourceDeviceProfileForUpsert)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "serialNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySourceDeviceProfileForUpsert struct {
	value *IdentitySourceDeviceProfileForUpsert
	isSet bool
}

func (v NullableIdentitySourceDeviceProfileForUpsert) Get() *IdentitySourceDeviceProfileForUpsert {
	return v.value
}

func (v *NullableIdentitySourceDeviceProfileForUpsert) Set(val *IdentitySourceDeviceProfileForUpsert) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceDeviceProfileForUpsert) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceDeviceProfileForUpsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceDeviceProfileForUpsert(val *IdentitySourceDeviceProfileForUpsert) *NullableIdentitySourceDeviceProfileForUpsert {
	return &NullableIdentitySourceDeviceProfileForUpsert{value: val, isSet: true}
}

func (v NullableIdentitySourceDeviceProfileForUpsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceDeviceProfileForUpsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

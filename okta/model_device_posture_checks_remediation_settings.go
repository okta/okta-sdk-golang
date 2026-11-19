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

// checks if the DevicePostureChecksRemediationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecksRemediationSettings{}

// DevicePostureChecksRemediationSettings Represents the remediation instructions shown to the end user when the device posture check fails
type DevicePostureChecksRemediationSettings struct {
	Link                 DevicePostureChecksRemediationSettingsLink    `json:"link"`
	Message              DevicePostureChecksRemediationSettingsMessage `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecksRemediationSettings DevicePostureChecksRemediationSettings

// NewDevicePostureChecksRemediationSettings instantiates a new DevicePostureChecksRemediationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecksRemediationSettings(link DevicePostureChecksRemediationSettingsLink, message DevicePostureChecksRemediationSettingsMessage) *DevicePostureChecksRemediationSettings {
	this := DevicePostureChecksRemediationSettings{}
	this.Link = link
	this.Message = message
	return &this
}

// NewDevicePostureChecksRemediationSettingsWithDefaults instantiates a new DevicePostureChecksRemediationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksRemediationSettingsWithDefaults() *DevicePostureChecksRemediationSettings {
	this := DevicePostureChecksRemediationSettings{}
	return &this
}

// GetLink returns the Link field value
func (o *DevicePostureChecksRemediationSettings) GetLink() DevicePostureChecksRemediationSettingsLink {
	if o == nil {
		var ret DevicePostureChecksRemediationSettingsLink
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettings) GetLinkOk() (*DevicePostureChecksRemediationSettingsLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *DevicePostureChecksRemediationSettings) SetLink(v DevicePostureChecksRemediationSettingsLink) {
	o.Link = v
}

// GetMessage returns the Message field value
func (o *DevicePostureChecksRemediationSettings) GetMessage() DevicePostureChecksRemediationSettingsMessage {
	if o == nil {
		var ret DevicePostureChecksRemediationSettingsMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettings) GetMessageOk() (*DevicePostureChecksRemediationSettingsMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DevicePostureChecksRemediationSettings) SetMessage(v DevicePostureChecksRemediationSettingsMessage) {
	o.Message = v
}

func (o DevicePostureChecksRemediationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureChecksRemediationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link"] = o.Link
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecksRemediationSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"link",
		"message",
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

	varDevicePostureChecksRemediationSettings := _DevicePostureChecksRemediationSettings{}

	err = json.Unmarshal(data, &varDevicePostureChecksRemediationSettings)

	if err != nil {
		return err
	}

	*o = DevicePostureChecksRemediationSettings(varDevicePostureChecksRemediationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "link")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureChecksRemediationSettings struct {
	value *DevicePostureChecksRemediationSettings
	isSet bool
}

func (v NullableDevicePostureChecksRemediationSettings) Get() *DevicePostureChecksRemediationSettings {
	return v.value
}

func (v *NullableDevicePostureChecksRemediationSettings) Set(val *DevicePostureChecksRemediationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureChecksRemediationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureChecksRemediationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureChecksRemediationSettings(val *DevicePostureChecksRemediationSettings) *NullableDevicePostureChecksRemediationSettings {
	return &NullableDevicePostureChecksRemediationSettings{value: val, isSet: true}
}

func (v NullableDevicePostureChecksRemediationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureChecksRemediationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

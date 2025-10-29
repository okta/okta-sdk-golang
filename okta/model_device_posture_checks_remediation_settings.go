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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DevicePostureChecksRemediationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecksRemediationSettings{}

// DevicePostureChecksRemediationSettings Represents the remediation instructions shown to the end user when the device posture check fails
type DevicePostureChecksRemediationSettings struct {
	Link                 *DevicePostureChecksRemediationSettingsLink    `json:"link,omitempty"`
	Message              *DevicePostureChecksRemediationSettingsMessage `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecksRemediationSettings DevicePostureChecksRemediationSettings

// NewDevicePostureChecksRemediationSettings instantiates a new DevicePostureChecksRemediationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecksRemediationSettings() *DevicePostureChecksRemediationSettings {
	this := DevicePostureChecksRemediationSettings{}
	return &this
}

// NewDevicePostureChecksRemediationSettingsWithDefaults instantiates a new DevicePostureChecksRemediationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksRemediationSettingsWithDefaults() *DevicePostureChecksRemediationSettings {
	this := DevicePostureChecksRemediationSettings{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettings) GetLink() DevicePostureChecksRemediationSettingsLink {
	if o == nil || IsNil(o.Link) {
		var ret DevicePostureChecksRemediationSettingsLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettings) GetLinkOk() (*DevicePostureChecksRemediationSettingsLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettings) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given DevicePostureChecksRemediationSettingsLink and assigns it to the Link field.
func (o *DevicePostureChecksRemediationSettings) SetLink(v DevicePostureChecksRemediationSettingsLink) {
	o.Link = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettings) GetMessage() DevicePostureChecksRemediationSettingsMessage {
	if o == nil || IsNil(o.Message) {
		var ret DevicePostureChecksRemediationSettingsMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettings) GetMessageOk() (*DevicePostureChecksRemediationSettingsMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettings) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given DevicePostureChecksRemediationSettingsMessage and assigns it to the Message field.
func (o *DevicePostureChecksRemediationSettings) SetMessage(v DevicePostureChecksRemediationSettingsMessage) {
	o.Message = &v
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
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecksRemediationSettings) UnmarshalJSON(data []byte) (err error) {
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

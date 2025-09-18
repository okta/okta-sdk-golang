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

// checks if the DevicePostureChecksRemediationSettingsMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecksRemediationSettingsMessage{}

// DevicePostureChecksRemediationSettingsMessage struct for DevicePostureChecksRemediationSettingsMessage
type DevicePostureChecksRemediationSettingsMessage struct {
	// Default i18n key for the message. This property is only relevant if type is set to `BUILTIN`. If type is set to `CUSTOM`, this field is ignored.
	DefaultI18nKey *string `json:"defaultI18nKey,omitempty"`
	// Custom text for the message
	CustomText           *string `json:"customText,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecksRemediationSettingsMessage DevicePostureChecksRemediationSettingsMessage

// NewDevicePostureChecksRemediationSettingsMessage instantiates a new DevicePostureChecksRemediationSettingsMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecksRemediationSettingsMessage() *DevicePostureChecksRemediationSettingsMessage {
	this := DevicePostureChecksRemediationSettingsMessage{}
	return &this
}

// NewDevicePostureChecksRemediationSettingsMessageWithDefaults instantiates a new DevicePostureChecksRemediationSettingsMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksRemediationSettingsMessageWithDefaults() *DevicePostureChecksRemediationSettingsMessage {
	this := DevicePostureChecksRemediationSettingsMessage{}
	return &this
}

// GetDefaultI18nKey returns the DefaultI18nKey field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettingsMessage) GetDefaultI18nKey() string {
	if o == nil || IsNil(o.DefaultI18nKey) {
		var ret string
		return ret
	}
	return *o.DefaultI18nKey
}

// GetDefaultI18nKeyOk returns a tuple with the DefaultI18nKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettingsMessage) GetDefaultI18nKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultI18nKey) {
		return nil, false
	}
	return o.DefaultI18nKey, true
}

// HasDefaultI18nKey returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettingsMessage) HasDefaultI18nKey() bool {
	if o != nil && !IsNil(o.DefaultI18nKey) {
		return true
	}

	return false
}

// SetDefaultI18nKey gets a reference to the given string and assigns it to the DefaultI18nKey field.
func (o *DevicePostureChecksRemediationSettingsMessage) SetDefaultI18nKey(v string) {
	o.DefaultI18nKey = &v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettingsMessage) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettingsMessage) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettingsMessage) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *DevicePostureChecksRemediationSettingsMessage) SetCustomText(v string) {
	o.CustomText = &v
}

func (o DevicePostureChecksRemediationSettingsMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureChecksRemediationSettingsMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultI18nKey) {
		toSerialize["defaultI18nKey"] = o.DefaultI18nKey
	}
	if !IsNil(o.CustomText) {
		toSerialize["customText"] = o.CustomText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecksRemediationSettingsMessage) UnmarshalJSON(data []byte) (err error) {
	varDevicePostureChecksRemediationSettingsMessage := _DevicePostureChecksRemediationSettingsMessage{}

	err = json.Unmarshal(data, &varDevicePostureChecksRemediationSettingsMessage)

	if err != nil {
		return err
	}

	*o = DevicePostureChecksRemediationSettingsMessage(varDevicePostureChecksRemediationSettingsMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultI18nKey")
		delete(additionalProperties, "customText")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureChecksRemediationSettingsMessage struct {
	value *DevicePostureChecksRemediationSettingsMessage
	isSet bool
}

func (v NullableDevicePostureChecksRemediationSettingsMessage) Get() *DevicePostureChecksRemediationSettingsMessage {
	return v.value
}

func (v *NullableDevicePostureChecksRemediationSettingsMessage) Set(val *DevicePostureChecksRemediationSettingsMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureChecksRemediationSettingsMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureChecksRemediationSettingsMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureChecksRemediationSettingsMessage(val *DevicePostureChecksRemediationSettingsMessage) *NullableDevicePostureChecksRemediationSettingsMessage {
	return &NullableDevicePostureChecksRemediationSettingsMessage{value: val, isSet: true}
}

func (v NullableDevicePostureChecksRemediationSettingsMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureChecksRemediationSettingsMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

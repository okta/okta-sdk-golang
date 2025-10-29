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

// checks if the DevicePostureChecksRemediationSettingsLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicePostureChecksRemediationSettingsLink{}

// DevicePostureChecksRemediationSettingsLink struct for DevicePostureChecksRemediationSettingsLink
type DevicePostureChecksRemediationSettingsLink struct {
	// Default URL for the link. This property is only relevant if type is set to `BUILTIN`. If type is set to `CUSTOM`, this field is ignored.
	DefaultUrl *string `json:"defaultUrl,omitempty"`
	// Custom URL for the link
	CustomUrl            *string `json:"customUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureChecksRemediationSettingsLink DevicePostureChecksRemediationSettingsLink

// NewDevicePostureChecksRemediationSettingsLink instantiates a new DevicePostureChecksRemediationSettingsLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureChecksRemediationSettingsLink() *DevicePostureChecksRemediationSettingsLink {
	this := DevicePostureChecksRemediationSettingsLink{}
	return &this
}

// NewDevicePostureChecksRemediationSettingsLinkWithDefaults instantiates a new DevicePostureChecksRemediationSettingsLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureChecksRemediationSettingsLinkWithDefaults() *DevicePostureChecksRemediationSettingsLink {
	this := DevicePostureChecksRemediationSettingsLink{}
	return &this
}

// GetDefaultUrl returns the DefaultUrl field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettingsLink) GetDefaultUrl() string {
	if o == nil || IsNil(o.DefaultUrl) {
		var ret string
		return ret
	}
	return *o.DefaultUrl
}

// GetDefaultUrlOk returns a tuple with the DefaultUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettingsLink) GetDefaultUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultUrl) {
		return nil, false
	}
	return o.DefaultUrl, true
}

// HasDefaultUrl returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettingsLink) HasDefaultUrl() bool {
	if o != nil && !IsNil(o.DefaultUrl) {
		return true
	}

	return false
}

// SetDefaultUrl gets a reference to the given string and assigns it to the DefaultUrl field.
func (o *DevicePostureChecksRemediationSettingsLink) SetDefaultUrl(v string) {
	o.DefaultUrl = &v
}

// GetCustomUrl returns the CustomUrl field value if set, zero value otherwise.
func (o *DevicePostureChecksRemediationSettingsLink) GetCustomUrl() string {
	if o == nil || IsNil(o.CustomUrl) {
		var ret string
		return ret
	}
	return *o.CustomUrl
}

// GetCustomUrlOk returns a tuple with the CustomUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureChecksRemediationSettingsLink) GetCustomUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomUrl) {
		return nil, false
	}
	return o.CustomUrl, true
}

// HasCustomUrl returns a boolean if a field has been set.
func (o *DevicePostureChecksRemediationSettingsLink) HasCustomUrl() bool {
	if o != nil && !IsNil(o.CustomUrl) {
		return true
	}

	return false
}

// SetCustomUrl gets a reference to the given string and assigns it to the CustomUrl field.
func (o *DevicePostureChecksRemediationSettingsLink) SetCustomUrl(v string) {
	o.CustomUrl = &v
}

func (o DevicePostureChecksRemediationSettingsLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicePostureChecksRemediationSettingsLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultUrl) {
		toSerialize["defaultUrl"] = o.DefaultUrl
	}
	if !IsNil(o.CustomUrl) {
		toSerialize["customUrl"] = o.CustomUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DevicePostureChecksRemediationSettingsLink) UnmarshalJSON(data []byte) (err error) {
	varDevicePostureChecksRemediationSettingsLink := _DevicePostureChecksRemediationSettingsLink{}

	err = json.Unmarshal(data, &varDevicePostureChecksRemediationSettingsLink)

	if err != nil {
		return err
	}

	*o = DevicePostureChecksRemediationSettingsLink(varDevicePostureChecksRemediationSettingsLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultUrl")
		delete(additionalProperties, "customUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicePostureChecksRemediationSettingsLink struct {
	value *DevicePostureChecksRemediationSettingsLink
	isSet bool
}

func (v NullableDevicePostureChecksRemediationSettingsLink) Get() *DevicePostureChecksRemediationSettingsLink {
	return v.value
}

func (v *NullableDevicePostureChecksRemediationSettingsLink) Set(val *DevicePostureChecksRemediationSettingsLink) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureChecksRemediationSettingsLink) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureChecksRemediationSettingsLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureChecksRemediationSettingsLink(val *DevicePostureChecksRemediationSettingsLink) *NullableDevicePostureChecksRemediationSettingsLink {
	return &NullableDevicePostureChecksRemediationSettingsLink{value: val, isSet: true}
}

func (v NullableDevicePostureChecksRemediationSettingsLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureChecksRemediationSettingsLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

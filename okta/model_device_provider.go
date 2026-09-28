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

// checks if the DeviceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceProvider{}

// DeviceProvider struct for DeviceProvider
type DeviceProvider struct {
	Payload *OktaVerifyPayload `json:"payload,omitempty"`
	// The version of the payload schema
	PayloadVersion *string `json:"payloadVersion,omitempty"`
	// The name of the provider
	Provider             *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceProvider DeviceProvider

// NewDeviceProvider instantiates a new DeviceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceProvider() *DeviceProvider {
	this := DeviceProvider{}
	return &this
}

// NewDeviceProviderWithDefaults instantiates a new DeviceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceProviderWithDefaults() *DeviceProvider {
	this := DeviceProvider{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *DeviceProvider) GetPayload() OktaVerifyPayload {
	if o == nil || IsNil(o.Payload) {
		var ret OktaVerifyPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProvider) GetPayloadOk() (*OktaVerifyPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *DeviceProvider) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given OktaVerifyPayload and assigns it to the Payload field.
func (o *DeviceProvider) SetPayload(v OktaVerifyPayload) {
	o.Payload = &v
}

// GetPayloadVersion returns the PayloadVersion field value if set, zero value otherwise.
func (o *DeviceProvider) GetPayloadVersion() string {
	if o == nil || IsNil(o.PayloadVersion) {
		var ret string
		return ret
	}
	return *o.PayloadVersion
}

// GetPayloadVersionOk returns a tuple with the PayloadVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProvider) GetPayloadVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadVersion) {
		return nil, false
	}
	return o.PayloadVersion, true
}

// HasPayloadVersion returns a boolean if a field has been set.
func (o *DeviceProvider) HasPayloadVersion() bool {
	if o != nil && !IsNil(o.PayloadVersion) {
		return true
	}

	return false
}

// SetPayloadVersion gets a reference to the given string and assigns it to the PayloadVersion field.
func (o *DeviceProvider) SetPayloadVersion(v string) {
	o.PayloadVersion = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *DeviceProvider) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProvider) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *DeviceProvider) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *DeviceProvider) SetProvider(v string) {
	o.Provider = &v
}

func (o DeviceProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadVersion) {
		toSerialize["payloadVersion"] = o.PayloadVersion
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceProvider) UnmarshalJSON(data []byte) (err error) {
	varDeviceProvider := _DeviceProvider{}

	err = json.Unmarshal(data, &varDeviceProvider)

	if err != nil {
		return err
	}

	*o = DeviceProvider(varDeviceProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadVersion")
		delete(additionalProperties, "provider")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceProvider struct {
	value *DeviceProvider
	isSet bool
}

func (v NullableDeviceProvider) Get() *DeviceProvider {
	return v.value
}

func (v *NullableDeviceProvider) Set(val *DeviceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceProvider(val *DeviceProvider) *NullableDeviceProvider {
	return &NullableDeviceProvider{value: val, isSet: true}
}

func (v NullableDeviceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

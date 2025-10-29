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
	"fmt"
)

// checks if the ApplicationSettingsNotificationsVpn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSettingsNotificationsVpn{}

// ApplicationSettingsNotificationsVpn Sends customizable messages with conditions to end users when a VPN connection is required
type ApplicationSettingsNotificationsVpn struct {
	// An optional URL to a help page to assist your end users in signing in to your company VPN
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A VPN requirement message that's displayed to users
	Message              *string                                    `json:"message,omitempty"`
	Network              ApplicationSettingsNotificationsVpnNetwork `json:"network"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettingsNotificationsVpn ApplicationSettingsNotificationsVpn

// NewApplicationSettingsNotificationsVpn instantiates a new ApplicationSettingsNotificationsVpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettingsNotificationsVpn(network ApplicationSettingsNotificationsVpnNetwork) *ApplicationSettingsNotificationsVpn {
	this := ApplicationSettingsNotificationsVpn{}
	this.Network = network
	return &this
}

// NewApplicationSettingsNotificationsVpnWithDefaults instantiates a new ApplicationSettingsNotificationsVpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSettingsNotificationsVpnWithDefaults() *ApplicationSettingsNotificationsVpn {
	this := ApplicationSettingsNotificationsVpn{}
	return &this
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpn) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpn) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *ApplicationSettingsNotificationsVpn) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpn) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpn) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplicationSettingsNotificationsVpn) SetMessage(v string) {
	o.Message = &v
}

// GetNetwork returns the Network field value
func (o *ApplicationSettingsNotificationsVpn) GetNetwork() ApplicationSettingsNotificationsVpnNetwork {
	if o == nil {
		var ret ApplicationSettingsNotificationsVpnNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetNetworkOk() (*ApplicationSettingsNotificationsVpnNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *ApplicationSettingsNotificationsVpn) SetNetwork(v ApplicationSettingsNotificationsVpnNetwork) {
	o.Network = v
}

func (o ApplicationSettingsNotificationsVpn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSettingsNotificationsVpn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HelpUrl) {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["network"] = o.Network

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationSettingsNotificationsVpn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"network",
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

	varApplicationSettingsNotificationsVpn := _ApplicationSettingsNotificationsVpn{}

	err = json.Unmarshal(data, &varApplicationSettingsNotificationsVpn)

	if err != nil {
		return err
	}

	*o = ApplicationSettingsNotificationsVpn(varApplicationSettingsNotificationsVpn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "helpUrl")
		delete(additionalProperties, "message")
		delete(additionalProperties, "network")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationSettingsNotificationsVpn struct {
	value *ApplicationSettingsNotificationsVpn
	isSet bool
}

func (v NullableApplicationSettingsNotificationsVpn) Get() *ApplicationSettingsNotificationsVpn {
	return v.value
}

func (v *NullableApplicationSettingsNotificationsVpn) Set(val *ApplicationSettingsNotificationsVpn) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSettingsNotificationsVpn) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSettingsNotificationsVpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSettingsNotificationsVpn(val *ApplicationSettingsNotificationsVpn) *NullableApplicationSettingsNotificationsVpn {
	return &NullableApplicationSettingsNotificationsVpn{value: val, isSet: true}
}

func (v NullableApplicationSettingsNotificationsVpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSettingsNotificationsVpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ApplicationSettingsNotificationsVpn struct for ApplicationSettingsNotificationsVpn
type ApplicationSettingsNotificationsVpn struct {
	HelpUrl *string `json:"helpUrl,omitempty"`
	Message *string `json:"message,omitempty"`
	Network *ApplicationSettingsNotificationsVpnNetwork `json:"network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettingsNotificationsVpn ApplicationSettingsNotificationsVpn

// NewApplicationSettingsNotificationsVpn instantiates a new ApplicationSettingsNotificationsVpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettingsNotificationsVpn() *ApplicationSettingsNotificationsVpn {
	this := ApplicationSettingsNotificationsVpn{}
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
	if o == nil || o.HelpUrl == nil {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetHelpUrlOk() (*string, bool) {
	if o == nil || o.HelpUrl == nil {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpn) HasHelpUrl() bool {
	if o != nil && o.HelpUrl != nil {
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
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpn) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplicationSettingsNotificationsVpn) SetMessage(v string) {
	o.Message = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpn) GetNetwork() ApplicationSettingsNotificationsVpnNetwork {
	if o == nil || o.Network == nil {
		var ret ApplicationSettingsNotificationsVpnNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpn) GetNetworkOk() (*ApplicationSettingsNotificationsVpnNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpn) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ApplicationSettingsNotificationsVpnNetwork and assigns it to the Network field.
func (o *ApplicationSettingsNotificationsVpn) SetNetwork(v ApplicationSettingsNotificationsVpnNetwork) {
	o.Network = &v
}

func (o ApplicationSettingsNotificationsVpn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HelpUrl != nil {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationSettingsNotificationsVpn) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationSettingsNotificationsVpn := _ApplicationSettingsNotificationsVpn{}

	err = json.Unmarshal(bytes, &varApplicationSettingsNotificationsVpn)
	if err == nil {
		*o = ApplicationSettingsNotificationsVpn(varApplicationSettingsNotificationsVpn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "helpUrl")
		delete(additionalProperties, "message")
		delete(additionalProperties, "network")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


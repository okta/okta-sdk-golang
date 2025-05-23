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

// ApplicationSettingsNotifications struct for ApplicationSettingsNotifications
type ApplicationSettingsNotifications struct {
	Vpn *ApplicationSettingsNotificationsVpn `json:"vpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettingsNotifications ApplicationSettingsNotifications

// NewApplicationSettingsNotifications instantiates a new ApplicationSettingsNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettingsNotifications() *ApplicationSettingsNotifications {
	this := ApplicationSettingsNotifications{}
	return &this
}

// NewApplicationSettingsNotificationsWithDefaults instantiates a new ApplicationSettingsNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSettingsNotificationsWithDefaults() *ApplicationSettingsNotifications {
	this := ApplicationSettingsNotifications{}
	return &this
}

// GetVpn returns the Vpn field value if set, zero value otherwise.
func (o *ApplicationSettingsNotifications) GetVpn() ApplicationSettingsNotificationsVpn {
	if o == nil || o.Vpn == nil {
		var ret ApplicationSettingsNotificationsVpn
		return ret
	}
	return *o.Vpn
}

// GetVpnOk returns a tuple with the Vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotifications) GetVpnOk() (*ApplicationSettingsNotificationsVpn, bool) {
	if o == nil || o.Vpn == nil {
		return nil, false
	}
	return o.Vpn, true
}

// HasVpn returns a boolean if a field has been set.
func (o *ApplicationSettingsNotifications) HasVpn() bool {
	if o != nil && o.Vpn != nil {
		return true
	}

	return false
}

// SetVpn gets a reference to the given ApplicationSettingsNotificationsVpn and assigns it to the Vpn field.
func (o *ApplicationSettingsNotifications) SetVpn(v ApplicationSettingsNotificationsVpn) {
	o.Vpn = &v
}

func (o ApplicationSettingsNotifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vpn != nil {
		toSerialize["vpn"] = o.Vpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationSettingsNotifications) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationSettingsNotifications := _ApplicationSettingsNotifications{}

	err = json.Unmarshal(bytes, &varApplicationSettingsNotifications)
	if err == nil {
		*o = ApplicationSettingsNotifications(varApplicationSettingsNotifications)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "vpn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationSettingsNotifications struct {
	value *ApplicationSettingsNotifications
	isSet bool
}

func (v NullableApplicationSettingsNotifications) Get() *ApplicationSettingsNotifications {
	return v.value
}

func (v *NullableApplicationSettingsNotifications) Set(val *ApplicationSettingsNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSettingsNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSettingsNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSettingsNotifications(val *ApplicationSettingsNotifications) *NullableApplicationSettingsNotifications {
	return &NullableApplicationSettingsNotifications{value: val, isSet: true}
}

func (v NullableApplicationSettingsNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSettingsNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


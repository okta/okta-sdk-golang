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

// checks if the ApplicationSettingsNotificationsVpnNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSettingsNotificationsVpnNetwork{}

// ApplicationSettingsNotificationsVpnNetwork Defines network zones for VPN notification
type ApplicationSettingsNotificationsVpnNetwork struct {
	// Specifies the VPN connection details required to access the app
	Connection *string `json:"connection,omitempty"`
	// Defines the IP addresses or network ranges that are excluded from the VPN requirement
	Exclude []string `json:"exclude,omitempty"`
	// Defines the IP addresses or network ranges that are required to use the VPN
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettingsNotificationsVpnNetwork ApplicationSettingsNotificationsVpnNetwork

// NewApplicationSettingsNotificationsVpnNetwork instantiates a new ApplicationSettingsNotificationsVpnNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettingsNotificationsVpnNetwork() *ApplicationSettingsNotificationsVpnNetwork {
	this := ApplicationSettingsNotificationsVpnNetwork{}
	return &this
}

// NewApplicationSettingsNotificationsVpnNetworkWithDefaults instantiates a new ApplicationSettingsNotificationsVpnNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSettingsNotificationsVpnNetworkWithDefaults() *ApplicationSettingsNotificationsVpnNetwork {
	this := ApplicationSettingsNotificationsVpnNetwork{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetConnection() string {
	if o == nil || IsNil(o.Connection) {
		var ret string
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetConnectionOk() (*string, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given string and assigns it to the Connection field.
func (o *ApplicationSettingsNotificationsVpnNetwork) SetConnection(v string) {
	o.Connection = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *ApplicationSettingsNotificationsVpnNetwork) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ApplicationSettingsNotificationsVpnNetwork) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *ApplicationSettingsNotificationsVpnNetwork) SetInclude(v []string) {
	o.Include = v
}

func (o ApplicationSettingsNotificationsVpnNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSettingsNotificationsVpnNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationSettingsNotificationsVpnNetwork) UnmarshalJSON(data []byte) (err error) {
	varApplicationSettingsNotificationsVpnNetwork := _ApplicationSettingsNotificationsVpnNetwork{}

	err = json.Unmarshal(data, &varApplicationSettingsNotificationsVpnNetwork)

	if err != nil {
		return err
	}

	*o = ApplicationSettingsNotificationsVpnNetwork(varApplicationSettingsNotificationsVpnNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connection")
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationSettingsNotificationsVpnNetwork struct {
	value *ApplicationSettingsNotificationsVpnNetwork
	isSet bool
}

func (v NullableApplicationSettingsNotificationsVpnNetwork) Get() *ApplicationSettingsNotificationsVpnNetwork {
	return v.value
}

func (v *NullableApplicationSettingsNotificationsVpnNetwork) Set(val *ApplicationSettingsNotificationsVpnNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSettingsNotificationsVpnNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSettingsNotificationsVpnNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSettingsNotificationsVpnNetwork(val *ApplicationSettingsNotificationsVpnNetwork) *NullableApplicationSettingsNotificationsVpnNetwork {
	return &NullableApplicationSettingsNotificationsVpnNetwork{value: val, isSet: true}
}

func (v NullableApplicationSettingsNotificationsVpnNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSettingsNotificationsVpnNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ScimScimServerConfig SCIM server schema configuration
type ScimScimServerConfig struct {
	Patch *ScimScimServerConfigPatch `json:"patch,omitempty"`
	ChangePassword *ScimScimServerConfigChangePassword `json:"changePassword,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScimScimServerConfig ScimScimServerConfig

// NewScimScimServerConfig instantiates a new ScimScimServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimScimServerConfig() *ScimScimServerConfig {
	this := ScimScimServerConfig{}
	return &this
}

// NewScimScimServerConfigWithDefaults instantiates a new ScimScimServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimScimServerConfigWithDefaults() *ScimScimServerConfig {
	this := ScimScimServerConfig{}
	return &this
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *ScimScimServerConfig) GetPatch() ScimScimServerConfigPatch {
	if o == nil || o.Patch == nil {
		var ret ScimScimServerConfigPatch
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimScimServerConfig) GetPatchOk() (*ScimScimServerConfigPatch, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *ScimScimServerConfig) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given ScimScimServerConfigPatch and assigns it to the Patch field.
func (o *ScimScimServerConfig) SetPatch(v ScimScimServerConfigPatch) {
	o.Patch = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *ScimScimServerConfig) GetChangePassword() ScimScimServerConfigChangePassword {
	if o == nil || o.ChangePassword == nil {
		var ret ScimScimServerConfigChangePassword
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimScimServerConfig) GetChangePasswordOk() (*ScimScimServerConfigChangePassword, bool) {
	if o == nil || o.ChangePassword == nil {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *ScimScimServerConfig) HasChangePassword() bool {
	if o != nil && o.ChangePassword != nil {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given ScimScimServerConfigChangePassword and assigns it to the ChangePassword field.
func (o *ScimScimServerConfig) SetChangePassword(v ScimScimServerConfigChangePassword) {
	o.ChangePassword = &v
}

func (o ScimScimServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	if o.ChangePassword != nil {
		toSerialize["changePassword"] = o.ChangePassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScimScimServerConfig) UnmarshalJSON(bytes []byte) (err error) {
	varScimScimServerConfig := _ScimScimServerConfig{}

	err = json.Unmarshal(bytes, &varScimScimServerConfig)
	if err == nil {
		*o = ScimScimServerConfig(varScimScimServerConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "patch")
		delete(additionalProperties, "changePassword")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableScimScimServerConfig struct {
	value *ScimScimServerConfig
	isSet bool
}

func (v NullableScimScimServerConfig) Get() *ScimScimServerConfig {
	return v.value
}

func (v *NullableScimScimServerConfig) Set(val *ScimScimServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableScimScimServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableScimScimServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimScimServerConfig(val *ScimScimServerConfig) *NullableScimScimServerConfig {
	return &NullableScimScimServerConfig{value: val, isSet: true}
}

func (v NullableScimScimServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimScimServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


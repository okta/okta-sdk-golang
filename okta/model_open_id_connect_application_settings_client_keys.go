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

// OpenIdConnectApplicationSettingsClientKeys struct for OpenIdConnectApplicationSettingsClientKeys
type OpenIdConnectApplicationSettingsClientKeys struct {
	Keys []SchemasJsonWebKey `json:"keys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationSettingsClientKeys OpenIdConnectApplicationSettingsClientKeys

// NewOpenIdConnectApplicationSettingsClientKeys instantiates a new OpenIdConnectApplicationSettingsClientKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationSettingsClientKeys() *OpenIdConnectApplicationSettingsClientKeys {
	this := OpenIdConnectApplicationSettingsClientKeys{}
	return &this
}

// NewOpenIdConnectApplicationSettingsClientKeysWithDefaults instantiates a new OpenIdConnectApplicationSettingsClientKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationSettingsClientKeysWithDefaults() *OpenIdConnectApplicationSettingsClientKeys {
	this := OpenIdConnectApplicationSettingsClientKeys{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClientKeys) GetKeys() []SchemasJsonWebKey {
	if o == nil || o.Keys == nil {
		var ret []SchemasJsonWebKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClientKeys) GetKeysOk() ([]SchemasJsonWebKey, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClientKeys) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []SchemasJsonWebKey and assigns it to the Keys field.
func (o *OpenIdConnectApplicationSettingsClientKeys) SetKeys(v []SchemasJsonWebKey) {
	o.Keys = v
}

func (o OpenIdConnectApplicationSettingsClientKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenIdConnectApplicationSettingsClientKeys) UnmarshalJSON(bytes []byte) (err error) {
	varOpenIdConnectApplicationSettingsClientKeys := _OpenIdConnectApplicationSettingsClientKeys{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationSettingsClientKeys)
	if err == nil {
		*o = OpenIdConnectApplicationSettingsClientKeys(varOpenIdConnectApplicationSettingsClientKeys)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOpenIdConnectApplicationSettingsClientKeys struct {
	value *OpenIdConnectApplicationSettingsClientKeys
	isSet bool
}

func (v NullableOpenIdConnectApplicationSettingsClientKeys) Get() *OpenIdConnectApplicationSettingsClientKeys {
	return v.value
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeys) Set(val *OpenIdConnectApplicationSettingsClientKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationSettingsClientKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationSettingsClientKeys(val *OpenIdConnectApplicationSettingsClientKeys) *NullableOpenIdConnectApplicationSettingsClientKeys {
	return &NullableOpenIdConnectApplicationSettingsClientKeys{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationSettingsClientKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationSettingsClientKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


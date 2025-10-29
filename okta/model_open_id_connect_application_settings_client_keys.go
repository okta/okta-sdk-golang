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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OpenIdConnectApplicationSettingsClientKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIdConnectApplicationSettingsClientKeys{}

// OpenIdConnectApplicationSettingsClientKeys A [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) for validating JWTs presented to Okta or for encrypting ID tokens minted by Okta for the client
type OpenIdConnectApplicationSettingsClientKeys struct {
	Keys                 []ListJwk200ResponseInner `json:"keys,omitempty"`
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
func (o *OpenIdConnectApplicationSettingsClientKeys) GetKeys() []ListJwk200ResponseInner {
	if o == nil || IsNil(o.Keys) {
		var ret []ListJwk200ResponseInner
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClientKeys) GetKeysOk() ([]ListJwk200ResponseInner, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClientKeys) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []ListJwk200ResponseInner and assigns it to the Keys field.
func (o *OpenIdConnectApplicationSettingsClientKeys) SetKeys(v []ListJwk200ResponseInner) {
	o.Keys = v
}

func (o OpenIdConnectApplicationSettingsClientKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIdConnectApplicationSettingsClientKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenIdConnectApplicationSettingsClientKeys) UnmarshalJSON(data []byte) (err error) {
	varOpenIdConnectApplicationSettingsClientKeys := _OpenIdConnectApplicationSettingsClientKeys{}

	err = json.Unmarshal(data, &varOpenIdConnectApplicationSettingsClientKeys)

	if err != nil {
		return err
	}

	*o = OpenIdConnectApplicationSettingsClientKeys(varOpenIdConnectApplicationSettingsClientKeys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
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

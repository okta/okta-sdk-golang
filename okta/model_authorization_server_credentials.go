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

// checks if the AuthorizationServerCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServerCredentials{}

// AuthorizationServerCredentials struct for AuthorizationServerCredentials
type AuthorizationServerCredentials struct {
	Signing              *AuthorizationServerCredentialsSigningConfig `json:"signing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerCredentials AuthorizationServerCredentials

// NewAuthorizationServerCredentials instantiates a new AuthorizationServerCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerCredentials() *AuthorizationServerCredentials {
	this := AuthorizationServerCredentials{}
	return &this
}

// NewAuthorizationServerCredentialsWithDefaults instantiates a new AuthorizationServerCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerCredentialsWithDefaults() *AuthorizationServerCredentials {
	this := AuthorizationServerCredentials{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *AuthorizationServerCredentials) GetSigning() AuthorizationServerCredentialsSigningConfig {
	if o == nil || IsNil(o.Signing) {
		var ret AuthorizationServerCredentialsSigningConfig
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentials) GetSigningOk() (*AuthorizationServerCredentialsSigningConfig, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *AuthorizationServerCredentials) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given AuthorizationServerCredentialsSigningConfig and assigns it to the Signing field.
func (o *AuthorizationServerCredentials) SetSigning(v AuthorizationServerCredentialsSigningConfig) {
	o.Signing = &v
}

func (o AuthorizationServerCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServerCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationServerCredentials) UnmarshalJSON(data []byte) (err error) {
	varAuthorizationServerCredentials := _AuthorizationServerCredentials{}

	err = json.Unmarshal(data, &varAuthorizationServerCredentials)

	if err != nil {
		return err
	}

	*o = AuthorizationServerCredentials(varAuthorizationServerCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationServerCredentials struct {
	value *AuthorizationServerCredentials
	isSet bool
}

func (v NullableAuthorizationServerCredentials) Get() *AuthorizationServerCredentials {
	return v.value
}

func (v *NullableAuthorizationServerCredentials) Set(val *AuthorizationServerCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerCredentials(val *AuthorizationServerCredentials) *NullableAuthorizationServerCredentials {
	return &NullableAuthorizationServerCredentials{value: val, isSet: true}
}

func (v NullableAuthorizationServerCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

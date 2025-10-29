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

// checks if the AuthenticatorIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorIdentity{}

// AuthenticatorIdentity Represents a particular authenticator serving as a constraint on a method
type AuthenticatorIdentity struct {
	Key                  *string `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorIdentity AuthenticatorIdentity

// NewAuthenticatorIdentity instantiates a new AuthenticatorIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorIdentity() *AuthenticatorIdentity {
	this := AuthenticatorIdentity{}
	return &this
}

// NewAuthenticatorIdentityWithDefaults instantiates a new AuthenticatorIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorIdentityWithDefaults() *AuthenticatorIdentity {
	this := AuthenticatorIdentity{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuthenticatorIdentity) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorIdentity) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuthenticatorIdentity) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuthenticatorIdentity) SetKey(v string) {
	o.Key = &v
}

func (o AuthenticatorIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorIdentity) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorIdentity := _AuthenticatorIdentity{}

	err = json.Unmarshal(data, &varAuthenticatorIdentity)

	if err != nil {
		return err
	}

	*o = AuthenticatorIdentity(varAuthenticatorIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorIdentity struct {
	value *AuthenticatorIdentity
	isSet bool
}

func (v NullableAuthenticatorIdentity) Get() *AuthenticatorIdentity {
	return v.value
}

func (v *NullableAuthenticatorIdentity) Set(val *AuthenticatorIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorIdentity(val *AuthenticatorIdentity) *NullableAuthenticatorIdentity {
	return &NullableAuthenticatorIdentity{value: val, isSet: true}
}

func (v NullableAuthenticatorIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

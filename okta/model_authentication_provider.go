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

// checks if the AuthenticationProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationProvider{}

// AuthenticationProvider Specifies the authentication provider that validates the user's password credential. The user's current provider is managed by the **Delegated Authentication** settings for your org. The provider object is **read-only**.
type AuthenticationProvider struct {
	// The name of the authentication provider
	Name *string `json:"name,omitempty"`
	// The type of authentication provider
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationProvider AuthenticationProvider

// NewAuthenticationProvider instantiates a new AuthenticationProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationProvider() *AuthenticationProvider {
	this := AuthenticationProvider{}
	return &this
}

// NewAuthenticationProviderWithDefaults instantiates a new AuthenticationProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationProviderWithDefaults() *AuthenticationProvider {
	this := AuthenticationProvider{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationProvider) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticationProvider) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationProvider) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticationProvider) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticationProvider) SetType(v string) {
	o.Type = &v
}

func (o AuthenticationProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticationProvider) UnmarshalJSON(data []byte) (err error) {
	varAuthenticationProvider := _AuthenticationProvider{}

	err = json.Unmarshal(data, &varAuthenticationProvider)

	if err != nil {
		return err
	}

	*o = AuthenticationProvider(varAuthenticationProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticationProvider struct {
	value *AuthenticationProvider
	isSet bool
}

func (v NullableAuthenticationProvider) Get() *AuthenticationProvider {
	return v.value
}

func (v *NullableAuthenticationProvider) Set(val *AuthenticationProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationProvider(val *AuthenticationProvider) *NullableAuthenticationProvider {
	return &NullableAuthenticationProvider{value: val, isSet: true}
}

func (v NullableAuthenticationProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

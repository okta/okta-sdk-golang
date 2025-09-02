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

// AuthenticationProviderWritable Specifies the authentication provider that validates the user password credential. The user's current provider is managed by the **Delegated Authentication** settings in your org. See [Create user with authentication provider](/openapi/okta-management/management/tag/User/#create-user-with-authentication-provider).
type AuthenticationProviderWritable struct {
	// The name of the authentication provider
	Name *string `json:"name,omitempty"`
	// The type of authentication provider
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationProviderWritable AuthenticationProviderWritable

// NewAuthenticationProviderWritable instantiates a new AuthenticationProviderWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationProviderWritable() *AuthenticationProviderWritable {
	this := AuthenticationProviderWritable{}
	return &this
}

// NewAuthenticationProviderWritableWithDefaults instantiates a new AuthenticationProviderWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationProviderWritableWithDefaults() *AuthenticationProviderWritable {
	this := AuthenticationProviderWritable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationProviderWritable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationProviderWritable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationProviderWritable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationProviderWritable) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticationProviderWritable) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationProviderWritable) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticationProviderWritable) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticationProviderWritable) SetType(v string) {
	o.Type = &v
}

func (o AuthenticationProviderWritable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationProviderWritable) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationProviderWritable := _AuthenticationProviderWritable{}

	err = json.Unmarshal(bytes, &varAuthenticationProviderWritable)
	if err == nil {
		*o = AuthenticationProviderWritable(varAuthenticationProviderWritable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticationProviderWritable struct {
	value *AuthenticationProviderWritable
	isSet bool
}

func (v NullableAuthenticationProviderWritable) Get() *AuthenticationProviderWritable {
	return v.value
}

func (v *NullableAuthenticationProviderWritable) Set(val *AuthenticationProviderWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationProviderWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationProviderWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationProviderWritable(val *AuthenticationProviderWritable) *NullableAuthenticationProviderWritable {
	return &NullableAuthenticationProviderWritable{value: val, isSet: true}
}

func (v NullableAuthenticationProviderWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationProviderWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


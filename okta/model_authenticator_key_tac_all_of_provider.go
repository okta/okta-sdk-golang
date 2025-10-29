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

// checks if the AuthenticatorKeyTacAllOfProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyTacAllOfProvider{}

// AuthenticatorKeyTacAllOfProvider <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Settings for the TAC authenticator
type AuthenticatorKeyTacAllOfProvider struct {
	// Provider type
	Type                 *string                                        `json:"type,omitempty"`
	Configuration        *AuthenticatorKeyTacAllOfProviderConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyTacAllOfProvider AuthenticatorKeyTacAllOfProvider

// NewAuthenticatorKeyTacAllOfProvider instantiates a new AuthenticatorKeyTacAllOfProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyTacAllOfProvider() *AuthenticatorKeyTacAllOfProvider {
	this := AuthenticatorKeyTacAllOfProvider{}
	return &this
}

// NewAuthenticatorKeyTacAllOfProviderWithDefaults instantiates a new AuthenticatorKeyTacAllOfProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyTacAllOfProviderWithDefaults() *AuthenticatorKeyTacAllOfProvider {
	this := AuthenticatorKeyTacAllOfProvider{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProvider) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProvider) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProvider) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorKeyTacAllOfProvider) SetType(v string) {
	o.Type = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProvider) GetConfiguration() AuthenticatorKeyTacAllOfProviderConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret AuthenticatorKeyTacAllOfProviderConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProvider) GetConfigurationOk() (*AuthenticatorKeyTacAllOfProviderConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProvider) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given AuthenticatorKeyTacAllOfProviderConfiguration and assigns it to the Configuration field.
func (o *AuthenticatorKeyTacAllOfProvider) SetConfiguration(v AuthenticatorKeyTacAllOfProviderConfiguration) {
	o.Configuration = &v
}

func (o AuthenticatorKeyTacAllOfProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyTacAllOfProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyTacAllOfProvider) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorKeyTacAllOfProvider := _AuthenticatorKeyTacAllOfProvider{}

	err = json.Unmarshal(data, &varAuthenticatorKeyTacAllOfProvider)

	if err != nil {
		return err
	}

	*o = AuthenticatorKeyTacAllOfProvider(varAuthenticatorKeyTacAllOfProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorKeyTacAllOfProvider struct {
	value *AuthenticatorKeyTacAllOfProvider
	isSet bool
}

func (v NullableAuthenticatorKeyTacAllOfProvider) Get() *AuthenticatorKeyTacAllOfProvider {
	return v.value
}

func (v *NullableAuthenticatorKeyTacAllOfProvider) Set(val *AuthenticatorKeyTacAllOfProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyTacAllOfProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyTacAllOfProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyTacAllOfProvider(val *AuthenticatorKeyTacAllOfProvider) *NullableAuthenticatorKeyTacAllOfProvider {
	return &NullableAuthenticatorKeyTacAllOfProvider{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyTacAllOfProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyTacAllOfProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

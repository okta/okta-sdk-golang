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

// AuthenticatorKeyCustomAppAllOfProvider struct for AuthenticatorKeyCustomAppAllOfProvider
type AuthenticatorKeyCustomAppAllOfProvider struct {
	// Provider type
	Type *string `json:"type,omitempty"`
	Configuration *AuthenticatorKeyCustomAppAllOfProviderConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomAppAllOfProvider AuthenticatorKeyCustomAppAllOfProvider

// NewAuthenticatorKeyCustomAppAllOfProvider instantiates a new AuthenticatorKeyCustomAppAllOfProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomAppAllOfProvider() *AuthenticatorKeyCustomAppAllOfProvider {
	this := AuthenticatorKeyCustomAppAllOfProvider{}
	return &this
}

// NewAuthenticatorKeyCustomAppAllOfProviderWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppAllOfProviderWithDefaults() *AuthenticatorKeyCustomAppAllOfProvider {
	this := AuthenticatorKeyCustomAppAllOfProvider{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProvider) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProvider) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorKeyCustomAppAllOfProvider) SetType(v string) {
	o.Type = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProvider) GetConfiguration() AuthenticatorKeyCustomAppAllOfProviderConfiguration {
	if o == nil || o.Configuration == nil {
		var ret AuthenticatorKeyCustomAppAllOfProviderConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProvider) GetConfigurationOk() (*AuthenticatorKeyCustomAppAllOfProviderConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProvider) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given AuthenticatorKeyCustomAppAllOfProviderConfiguration and assigns it to the Configuration field.
func (o *AuthenticatorKeyCustomAppAllOfProvider) SetConfiguration(v AuthenticatorKeyCustomAppAllOfProviderConfiguration) {
	o.Configuration = &v
}

func (o AuthenticatorKeyCustomAppAllOfProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyCustomAppAllOfProvider) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyCustomAppAllOfProvider := _AuthenticatorKeyCustomAppAllOfProvider{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyCustomAppAllOfProvider)
	if err == nil {
		*o = AuthenticatorKeyCustomAppAllOfProvider(varAuthenticatorKeyCustomAppAllOfProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyCustomAppAllOfProvider struct {
	value *AuthenticatorKeyCustomAppAllOfProvider
	isSet bool
}

func (v NullableAuthenticatorKeyCustomAppAllOfProvider) Get() *AuthenticatorKeyCustomAppAllOfProvider {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProvider) Set(val *AuthenticatorKeyCustomAppAllOfProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomAppAllOfProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomAppAllOfProvider(val *AuthenticatorKeyCustomAppAllOfProvider) *NullableAuthenticatorKeyCustomAppAllOfProvider {
	return &NullableAuthenticatorKeyCustomAppAllOfProvider{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomAppAllOfProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


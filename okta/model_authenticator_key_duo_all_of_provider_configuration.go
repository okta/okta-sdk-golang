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

// AuthenticatorKeyDuoAllOfProviderConfiguration struct for AuthenticatorKeyDuoAllOfProviderConfiguration
type AuthenticatorKeyDuoAllOfProviderConfiguration struct {
	// The Duo Security API hostname
	Host *string `json:"host,omitempty"`
	// The Duo Security integration key
	IntegrationKey *string `json:"integrationKey,omitempty"`
	// The Duo Security secret key
	SecretKey *string `json:"secretKey,omitempty"`
	UserNameTemplate *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate `json:"userNameTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyDuoAllOfProviderConfiguration AuthenticatorKeyDuoAllOfProviderConfiguration

// NewAuthenticatorKeyDuoAllOfProviderConfiguration instantiates a new AuthenticatorKeyDuoAllOfProviderConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyDuoAllOfProviderConfiguration() *AuthenticatorKeyDuoAllOfProviderConfiguration {
	this := AuthenticatorKeyDuoAllOfProviderConfiguration{}
	return &this
}

// NewAuthenticatorKeyDuoAllOfProviderConfigurationWithDefaults instantiates a new AuthenticatorKeyDuoAllOfProviderConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyDuoAllOfProviderConfigurationWithDefaults() *AuthenticatorKeyDuoAllOfProviderConfiguration {
	this := AuthenticatorKeyDuoAllOfProviderConfiguration{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetHost(v string) {
	o.Host = &v
}

// GetIntegrationKey returns the IntegrationKey field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetIntegrationKey() string {
	if o == nil || o.IntegrationKey == nil {
		var ret string
		return ret
	}
	return *o.IntegrationKey
}

// GetIntegrationKeyOk returns a tuple with the IntegrationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetIntegrationKeyOk() (*string, bool) {
	if o == nil || o.IntegrationKey == nil {
		return nil, false
	}
	return o.IntegrationKey, true
}

// HasIntegrationKey returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasIntegrationKey() bool {
	if o != nil && o.IntegrationKey != nil {
		return true
	}

	return false
}

// SetIntegrationKey gets a reference to the given string and assigns it to the IntegrationKey field.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetIntegrationKey(v string) {
	o.IntegrationKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetUserNameTemplate returns the UserNameTemplate field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetUserNameTemplate() AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate {
	if o == nil || o.UserNameTemplate == nil {
		var ret AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetUserNameTemplateOk() (*AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate, bool) {
	if o == nil || o.UserNameTemplate == nil {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasUserNameTemplate() bool {
	if o != nil && o.UserNameTemplate != nil {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate and assigns it to the UserNameTemplate field.
func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetUserNameTemplate(v AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) {
	o.UserNameTemplate = &v
}

func (o AuthenticatorKeyDuoAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.IntegrationKey != nil {
		toSerialize["integrationKey"] = o.IntegrationKey
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if o.UserNameTemplate != nil {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyDuoAllOfProviderConfiguration := _AuthenticatorKeyDuoAllOfProviderConfiguration{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyDuoAllOfProviderConfiguration)
	if err == nil {
		*o = AuthenticatorKeyDuoAllOfProviderConfiguration(varAuthenticatorKeyDuoAllOfProviderConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "integrationKey")
		delete(additionalProperties, "secretKey")
		delete(additionalProperties, "userNameTemplate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyDuoAllOfProviderConfiguration struct {
	value *AuthenticatorKeyDuoAllOfProviderConfiguration
	isSet bool
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfiguration) Get() *AuthenticatorKeyDuoAllOfProviderConfiguration {
	return v.value
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfiguration) Set(val *AuthenticatorKeyDuoAllOfProviderConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyDuoAllOfProviderConfiguration(val *AuthenticatorKeyDuoAllOfProviderConfiguration) *NullableAuthenticatorKeyDuoAllOfProviderConfiguration {
	return &NullableAuthenticatorKeyDuoAllOfProviderConfiguration{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


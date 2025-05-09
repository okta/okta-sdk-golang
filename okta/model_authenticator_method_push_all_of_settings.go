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

// AuthenticatorMethodPushAllOfSettings struct for AuthenticatorMethodPushAllOfSettings
type AuthenticatorMethodPushAllOfSettings struct {
	Algorithms []string `json:"algorithms,omitempty"`
	// Indicates whether you must use a hardware key store
	KeyProtection *string `json:"keyProtection,omitempty"`
	TransactionTypes []string `json:"transactionTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodPushAllOfSettings AuthenticatorMethodPushAllOfSettings

// NewAuthenticatorMethodPushAllOfSettings instantiates a new AuthenticatorMethodPushAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodPushAllOfSettings() *AuthenticatorMethodPushAllOfSettings {
	this := AuthenticatorMethodPushAllOfSettings{}
	return &this
}

// NewAuthenticatorMethodPushAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodPushAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodPushAllOfSettingsWithDefaults() *AuthenticatorMethodPushAllOfSettings {
	this := AuthenticatorMethodPushAllOfSettings{}
	return &this
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *AuthenticatorMethodPushAllOfSettings) GetAlgorithms() []string {
	if o == nil || o.Algorithms == nil {
		var ret []string
		return ret
	}
	return o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodPushAllOfSettings) GetAlgorithmsOk() ([]string, bool) {
	if o == nil || o.Algorithms == nil {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *AuthenticatorMethodPushAllOfSettings) HasAlgorithms() bool {
	if o != nil && o.Algorithms != nil {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given []string and assigns it to the Algorithms field.
func (o *AuthenticatorMethodPushAllOfSettings) SetAlgorithms(v []string) {
	o.Algorithms = v
}

// GetKeyProtection returns the KeyProtection field value if set, zero value otherwise.
func (o *AuthenticatorMethodPushAllOfSettings) GetKeyProtection() string {
	if o == nil || o.KeyProtection == nil {
		var ret string
		return ret
	}
	return *o.KeyProtection
}

// GetKeyProtectionOk returns a tuple with the KeyProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodPushAllOfSettings) GetKeyProtectionOk() (*string, bool) {
	if o == nil || o.KeyProtection == nil {
		return nil, false
	}
	return o.KeyProtection, true
}

// HasKeyProtection returns a boolean if a field has been set.
func (o *AuthenticatorMethodPushAllOfSettings) HasKeyProtection() bool {
	if o != nil && o.KeyProtection != nil {
		return true
	}

	return false
}

// SetKeyProtection gets a reference to the given string and assigns it to the KeyProtection field.
func (o *AuthenticatorMethodPushAllOfSettings) SetKeyProtection(v string) {
	o.KeyProtection = &v
}

// GetTransactionTypes returns the TransactionTypes field value if set, zero value otherwise.
func (o *AuthenticatorMethodPushAllOfSettings) GetTransactionTypes() []string {
	if o == nil || o.TransactionTypes == nil {
		var ret []string
		return ret
	}
	return o.TransactionTypes
}

// GetTransactionTypesOk returns a tuple with the TransactionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodPushAllOfSettings) GetTransactionTypesOk() ([]string, bool) {
	if o == nil || o.TransactionTypes == nil {
		return nil, false
	}
	return o.TransactionTypes, true
}

// HasTransactionTypes returns a boolean if a field has been set.
func (o *AuthenticatorMethodPushAllOfSettings) HasTransactionTypes() bool {
	if o != nil && o.TransactionTypes != nil {
		return true
	}

	return false
}

// SetTransactionTypes gets a reference to the given []string and assigns it to the TransactionTypes field.
func (o *AuthenticatorMethodPushAllOfSettings) SetTransactionTypes(v []string) {
	o.TransactionTypes = v
}

func (o AuthenticatorMethodPushAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithms != nil {
		toSerialize["algorithms"] = o.Algorithms
	}
	if o.KeyProtection != nil {
		toSerialize["keyProtection"] = o.KeyProtection
	}
	if o.TransactionTypes != nil {
		toSerialize["transactionTypes"] = o.TransactionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodPushAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodPushAllOfSettings := _AuthenticatorMethodPushAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodPushAllOfSettings)
	if err == nil {
		*o = AuthenticatorMethodPushAllOfSettings(varAuthenticatorMethodPushAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "keyProtection")
		delete(additionalProperties, "transactionTypes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodPushAllOfSettings struct {
	value *AuthenticatorMethodPushAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorMethodPushAllOfSettings) Get() *AuthenticatorMethodPushAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorMethodPushAllOfSettings) Set(val *AuthenticatorMethodPushAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodPushAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodPushAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodPushAllOfSettings(val *AuthenticatorMethodPushAllOfSettings) *NullableAuthenticatorMethodPushAllOfSettings {
	return &NullableAuthenticatorMethodPushAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodPushAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodPushAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


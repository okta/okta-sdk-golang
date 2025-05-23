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

// SupportedMethodsSettings struct for SupportedMethodsSettings
type SupportedMethodsSettings struct {
	// Indicates whether you must use a hardware key store
	KeyProtection *string `json:"keyProtection,omitempty"`
	// The encryption algorithm for this authenticator method
	Algorithms []string `json:"algorithms,omitempty"`
	// The transaction type for this authenticator method
	TransactionTypes []string `json:"transactionTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SupportedMethodsSettings SupportedMethodsSettings

// NewSupportedMethodsSettings instantiates a new SupportedMethodsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedMethodsSettings() *SupportedMethodsSettings {
	this := SupportedMethodsSettings{}
	return &this
}

// NewSupportedMethodsSettingsWithDefaults instantiates a new SupportedMethodsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedMethodsSettingsWithDefaults() *SupportedMethodsSettings {
	this := SupportedMethodsSettings{}
	return &this
}

// GetKeyProtection returns the KeyProtection field value if set, zero value otherwise.
func (o *SupportedMethodsSettings) GetKeyProtection() string {
	if o == nil || o.KeyProtection == nil {
		var ret string
		return ret
	}
	return *o.KeyProtection
}

// GetKeyProtectionOk returns a tuple with the KeyProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetKeyProtectionOk() (*string, bool) {
	if o == nil || o.KeyProtection == nil {
		return nil, false
	}
	return o.KeyProtection, true
}

// HasKeyProtection returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasKeyProtection() bool {
	if o != nil && o.KeyProtection != nil {
		return true
	}

	return false
}

// SetKeyProtection gets a reference to the given string and assigns it to the KeyProtection field.
func (o *SupportedMethodsSettings) SetKeyProtection(v string) {
	o.KeyProtection = &v
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *SupportedMethodsSettings) GetAlgorithms() []string {
	if o == nil || o.Algorithms == nil {
		var ret []string
		return ret
	}
	return o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetAlgorithmsOk() ([]string, bool) {
	if o == nil || o.Algorithms == nil {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasAlgorithms() bool {
	if o != nil && o.Algorithms != nil {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given []string and assigns it to the Algorithms field.
func (o *SupportedMethodsSettings) SetAlgorithms(v []string) {
	o.Algorithms = v
}

// GetTransactionTypes returns the TransactionTypes field value if set, zero value otherwise.
func (o *SupportedMethodsSettings) GetTransactionTypes() []string {
	if o == nil || o.TransactionTypes == nil {
		var ret []string
		return ret
	}
	return o.TransactionTypes
}

// GetTransactionTypesOk returns a tuple with the TransactionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetTransactionTypesOk() ([]string, bool) {
	if o == nil || o.TransactionTypes == nil {
		return nil, false
	}
	return o.TransactionTypes, true
}

// HasTransactionTypes returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasTransactionTypes() bool {
	if o != nil && o.TransactionTypes != nil {
		return true
	}

	return false
}

// SetTransactionTypes gets a reference to the given []string and assigns it to the TransactionTypes field.
func (o *SupportedMethodsSettings) SetTransactionTypes(v []string) {
	o.TransactionTypes = v
}

func (o SupportedMethodsSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyProtection != nil {
		toSerialize["keyProtection"] = o.KeyProtection
	}
	if o.Algorithms != nil {
		toSerialize["algorithms"] = o.Algorithms
	}
	if o.TransactionTypes != nil {
		toSerialize["transactionTypes"] = o.TransactionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SupportedMethodsSettings) UnmarshalJSON(bytes []byte) (err error) {
	varSupportedMethodsSettings := _SupportedMethodsSettings{}

	err = json.Unmarshal(bytes, &varSupportedMethodsSettings)
	if err == nil {
		*o = SupportedMethodsSettings(varSupportedMethodsSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "keyProtection")
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "transactionTypes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSupportedMethodsSettings struct {
	value *SupportedMethodsSettings
	isSet bool
}

func (v NullableSupportedMethodsSettings) Get() *SupportedMethodsSettings {
	return v.value
}

func (v *NullableSupportedMethodsSettings) Set(val *SupportedMethodsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedMethodsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedMethodsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedMethodsSettings(val *SupportedMethodsSettings) *NullableSupportedMethodsSettings {
	return &NullableSupportedMethodsSettings{value: val, isSet: true}
}

func (v NullableSupportedMethodsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedMethodsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


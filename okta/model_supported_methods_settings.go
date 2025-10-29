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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SupportedMethodsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedMethodsSettings{}

// SupportedMethodsSettings struct for SupportedMethodsSettings
type SupportedMethodsSettings struct {
	// Indicates whether you must use a hardware key store
	KeyProtection *string `json:"keyProtection,omitempty"`
	// The encryption algorithm for this authenticator method
	Algorithms []string `json:"algorithms,omitempty"`
	// The transaction type for this authenticator method
	TransactionTypes     []string `json:"transactionTypes,omitempty"`
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
	if o == nil || IsNil(o.KeyProtection) {
		var ret string
		return ret
	}
	return *o.KeyProtection
}

// GetKeyProtectionOk returns a tuple with the KeyProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetKeyProtectionOk() (*string, bool) {
	if o == nil || IsNil(o.KeyProtection) {
		return nil, false
	}
	return o.KeyProtection, true
}

// HasKeyProtection returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasKeyProtection() bool {
	if o != nil && !IsNil(o.KeyProtection) {
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
	if o == nil || IsNil(o.Algorithms) {
		var ret []string
		return ret
	}
	return o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetAlgorithmsOk() ([]string, bool) {
	if o == nil || IsNil(o.Algorithms) {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasAlgorithms() bool {
	if o != nil && !IsNil(o.Algorithms) {
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
	if o == nil || IsNil(o.TransactionTypes) {
		var ret []string
		return ret
	}
	return o.TransactionTypes
}

// GetTransactionTypesOk returns a tuple with the TransactionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethodsSettings) GetTransactionTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.TransactionTypes) {
		return nil, false
	}
	return o.TransactionTypes, true
}

// HasTransactionTypes returns a boolean if a field has been set.
func (o *SupportedMethodsSettings) HasTransactionTypes() bool {
	if o != nil && !IsNil(o.TransactionTypes) {
		return true
	}

	return false
}

// SetTransactionTypes gets a reference to the given []string and assigns it to the TransactionTypes field.
func (o *SupportedMethodsSettings) SetTransactionTypes(v []string) {
	o.TransactionTypes = v
}

func (o SupportedMethodsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedMethodsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyProtection) {
		toSerialize["keyProtection"] = o.KeyProtection
	}
	if !IsNil(o.Algorithms) {
		toSerialize["algorithms"] = o.Algorithms
	}
	if !IsNil(o.TransactionTypes) {
		toSerialize["transactionTypes"] = o.TransactionTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SupportedMethodsSettings) UnmarshalJSON(data []byte) (err error) {
	varSupportedMethodsSettings := _SupportedMethodsSettings{}

	err = json.Unmarshal(data, &varSupportedMethodsSettings)

	if err != nil {
		return err
	}

	*o = SupportedMethodsSettings(varSupportedMethodsSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keyProtection")
		delete(additionalProperties, "algorithms")
		delete(additionalProperties, "transactionTypes")
		o.AdditionalProperties = additionalProperties
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

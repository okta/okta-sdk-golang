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
	"fmt"
)

// AuthenticatorKeyTacAllOfProviderConfiguration Define the configuration settings of the TAC
type AuthenticatorKeyTacAllOfProviderConfiguration struct {
	// Minimum time-to-live (TTL) of the TAC in minutes. The `minTtl` indicates the minimum amount of time that a TAC is valid. The `minTtl` must be less than the `maxTtl`.
	MinTtl float32 `json:"minTtl"`
	// Maximum TTL of the TAC in minutes. The `maxTtl` indicates the maximum amount of time that a TAC is valid. The `maxTtl` must be greater than the `minTtl`.
	MaxTtl float32 `json:"maxTtl"`
	// The default TTL in minutes when you create a TAC. The `defaultTtl` indicates the actual amount of time that a TAC is valid before it expires. The `defaultTtl` must be greater than the `minTtl` and less than the `maxTtl`.
	DefaultTtl float32 `json:"defaultTtl"`
	// Defines the number of characters in a TAC. For example, a `length` of `16` means that the TAC is 16 characters.
	Length float32 `json:"length"`
	Complexity AuthenticatorKeyTacAllOfProviderConfigurationComplexity `json:"complexity"`
	// Indicates whether a TAC can be used multiple times. If set to `true`, the TAC can be used multiple times until it expires.
	MultiUseAllowed *bool `json:"multiUseAllowed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyTacAllOfProviderConfiguration AuthenticatorKeyTacAllOfProviderConfiguration

// NewAuthenticatorKeyTacAllOfProviderConfiguration instantiates a new AuthenticatorKeyTacAllOfProviderConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyTacAllOfProviderConfiguration(minTtl float32, maxTtl float32, defaultTtl float32, length float32, complexity AuthenticatorKeyTacAllOfProviderConfigurationComplexity) *AuthenticatorKeyTacAllOfProviderConfiguration {
	this := AuthenticatorKeyTacAllOfProviderConfiguration{}
	this.MinTtl = minTtl
	this.MaxTtl = maxTtl
	this.DefaultTtl = defaultTtl
	this.Length = length
	this.Complexity = complexity
	return &this
}

// NewAuthenticatorKeyTacAllOfProviderConfigurationWithDefaults instantiates a new AuthenticatorKeyTacAllOfProviderConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyTacAllOfProviderConfigurationWithDefaults() *AuthenticatorKeyTacAllOfProviderConfiguration {
	this := AuthenticatorKeyTacAllOfProviderConfiguration{}
	var defaultTtl float32 = 120
	this.DefaultTtl = defaultTtl
	return &this
}

// GetMinTtl returns the MinTtl field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMinTtl() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinTtl
}

// GetMinTtlOk returns a tuple with the MinTtl field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMinTtlOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinTtl, true
}

// SetMinTtl sets field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMinTtl(v float32) {
	o.MinTtl = v
}

// GetMaxTtl returns the MaxTtl field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMaxTtl() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMaxTtlOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTtl, true
}

// SetMaxTtl sets field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMaxTtl(v float32) {
	o.MaxTtl = v
}

// GetDefaultTtl returns the DefaultTtl field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetDefaultTtl() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetDefaultTtlOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTtl, true
}

// SetDefaultTtl sets field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetDefaultTtl(v float32) {
	o.DefaultTtl = v
}

// GetLength returns the Length field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetLength(v float32) {
	o.Length = v
}

// GetComplexity returns the Complexity field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetComplexity() AuthenticatorKeyTacAllOfProviderConfigurationComplexity {
	if o == nil {
		var ret AuthenticatorKeyTacAllOfProviderConfigurationComplexity
		return ret
	}

	return o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetComplexityOk() (*AuthenticatorKeyTacAllOfProviderConfigurationComplexity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Complexity, true
}

// SetComplexity sets field value
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetComplexity(v AuthenticatorKeyTacAllOfProviderConfigurationComplexity) {
	o.Complexity = v
}

// GetMultiUseAllowed returns the MultiUseAllowed field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMultiUseAllowed() bool {
	if o == nil || o.MultiUseAllowed == nil {
		var ret bool
		return ret
	}
	return *o.MultiUseAllowed
}

// GetMultiUseAllowedOk returns a tuple with the MultiUseAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMultiUseAllowedOk() (*bool, bool) {
	if o == nil || o.MultiUseAllowed == nil {
		return nil, false
	}
	return o.MultiUseAllowed, true
}

// HasMultiUseAllowed returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) HasMultiUseAllowed() bool {
	if o != nil && o.MultiUseAllowed != nil {
		return true
	}

	return false
}

// SetMultiUseAllowed gets a reference to the given bool and assigns it to the MultiUseAllowed field.
func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMultiUseAllowed(v bool) {
	o.MultiUseAllowed = &v
}

func (o AuthenticatorKeyTacAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["minTtl"] = o.MinTtl
	}
	if true {
		toSerialize["maxTtl"] = o.MaxTtl
	}
	if true {
		toSerialize["defaultTtl"] = o.DefaultTtl
	}
	if true {
		toSerialize["length"] = o.Length
	}
	if true {
		toSerialize["complexity"] = o.Complexity
	}
	if o.MultiUseAllowed != nil {
		toSerialize["multiUseAllowed"] = o.MultiUseAllowed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyTacAllOfProviderConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyTacAllOfProviderConfiguration := _AuthenticatorKeyTacAllOfProviderConfiguration{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyTacAllOfProviderConfiguration)
	if err == nil {
		*o = AuthenticatorKeyTacAllOfProviderConfiguration(varAuthenticatorKeyTacAllOfProviderConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "minTtl")
		delete(additionalProperties, "maxTtl")
		delete(additionalProperties, "defaultTtl")
		delete(additionalProperties, "length")
		delete(additionalProperties, "complexity")
		delete(additionalProperties, "multiUseAllowed")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyTacAllOfProviderConfiguration struct {
	value *AuthenticatorKeyTacAllOfProviderConfiguration
	isSet bool
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfiguration) Get() *AuthenticatorKeyTacAllOfProviderConfiguration {
	return v.value
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfiguration) Set(val *AuthenticatorKeyTacAllOfProviderConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyTacAllOfProviderConfiguration(val *AuthenticatorKeyTacAllOfProviderConfiguration) *NullableAuthenticatorKeyTacAllOfProviderConfiguration {
	return &NullableAuthenticatorKeyTacAllOfProviderConfiguration{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


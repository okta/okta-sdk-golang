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

// checks if the AuthenticatorKeyTacAllOfProviderConfigurationComplexity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyTacAllOfProviderConfigurationComplexity{}

// AuthenticatorKeyTacAllOfProviderConfigurationComplexity Define the complexity of the TAC
type AuthenticatorKeyTacAllOfProviderConfigurationComplexity struct {
	// Use numbers in the TAC. `numbers` is always `true` for the TAC authenticator.
	Numbers *bool `json:"numbers,omitempty"`
	// Use letters in the TAC
	Letters *bool `json:"letters,omitempty"`
	// Use special characters in the TAC
	SpecialCharacters    *bool `json:"specialCharacters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyTacAllOfProviderConfigurationComplexity AuthenticatorKeyTacAllOfProviderConfigurationComplexity

// NewAuthenticatorKeyTacAllOfProviderConfigurationComplexity instantiates a new AuthenticatorKeyTacAllOfProviderConfigurationComplexity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyTacAllOfProviderConfigurationComplexity() *AuthenticatorKeyTacAllOfProviderConfigurationComplexity {
	this := AuthenticatorKeyTacAllOfProviderConfigurationComplexity{}
	return &this
}

// NewAuthenticatorKeyTacAllOfProviderConfigurationComplexityWithDefaults instantiates a new AuthenticatorKeyTacAllOfProviderConfigurationComplexity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyTacAllOfProviderConfigurationComplexityWithDefaults() *AuthenticatorKeyTacAllOfProviderConfigurationComplexity {
	this := AuthenticatorKeyTacAllOfProviderConfigurationComplexity{}
	return &this
}

// GetNumbers returns the Numbers field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetNumbers() bool {
	if o == nil || IsNil(o.Numbers) {
		var ret bool
		return ret
	}
	return *o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetNumbersOk() (*bool, bool) {
	if o == nil || IsNil(o.Numbers) {
		return nil, false
	}
	return o.Numbers, true
}

// HasNumbers returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) HasNumbers() bool {
	if o != nil && !IsNil(o.Numbers) {
		return true
	}

	return false
}

// SetNumbers gets a reference to the given bool and assigns it to the Numbers field.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) SetNumbers(v bool) {
	o.Numbers = &v
}

// GetLetters returns the Letters field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetLetters() bool {
	if o == nil || IsNil(o.Letters) {
		var ret bool
		return ret
	}
	return *o.Letters
}

// GetLettersOk returns a tuple with the Letters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetLettersOk() (*bool, bool) {
	if o == nil || IsNil(o.Letters) {
		return nil, false
	}
	return o.Letters, true
}

// HasLetters returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) HasLetters() bool {
	if o != nil && !IsNil(o.Letters) {
		return true
	}

	return false
}

// SetLetters gets a reference to the given bool and assigns it to the Letters field.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) SetLetters(v bool) {
	o.Letters = &v
}

// GetSpecialCharacters returns the SpecialCharacters field value if set, zero value otherwise.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetSpecialCharacters() bool {
	if o == nil || IsNil(o.SpecialCharacters) {
		var ret bool
		return ret
	}
	return *o.SpecialCharacters
}

// GetSpecialCharactersOk returns a tuple with the SpecialCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) GetSpecialCharactersOk() (*bool, bool) {
	if o == nil || IsNil(o.SpecialCharacters) {
		return nil, false
	}
	return o.SpecialCharacters, true
}

// HasSpecialCharacters returns a boolean if a field has been set.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) HasSpecialCharacters() bool {
	if o != nil && !IsNil(o.SpecialCharacters) {
		return true
	}

	return false
}

// SetSpecialCharacters gets a reference to the given bool and assigns it to the SpecialCharacters field.
func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) SetSpecialCharacters(v bool) {
	o.SpecialCharacters = &v
}

func (o AuthenticatorKeyTacAllOfProviderConfigurationComplexity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyTacAllOfProviderConfigurationComplexity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numbers) {
		toSerialize["numbers"] = o.Numbers
	}
	if !IsNil(o.Letters) {
		toSerialize["letters"] = o.Letters
	}
	if !IsNil(o.SpecialCharacters) {
		toSerialize["specialCharacters"] = o.SpecialCharacters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorKeyTacAllOfProviderConfigurationComplexity := _AuthenticatorKeyTacAllOfProviderConfigurationComplexity{}

	err = json.Unmarshal(data, &varAuthenticatorKeyTacAllOfProviderConfigurationComplexity)

	if err != nil {
		return err
	}

	*o = AuthenticatorKeyTacAllOfProviderConfigurationComplexity(varAuthenticatorKeyTacAllOfProviderConfigurationComplexity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "letters")
		delete(additionalProperties, "specialCharacters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity struct {
	value *AuthenticatorKeyTacAllOfProviderConfigurationComplexity
	isSet bool
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) Get() *AuthenticatorKeyTacAllOfProviderConfigurationComplexity {
	return v.value
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) Set(val *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity(val *AuthenticatorKeyTacAllOfProviderConfigurationComplexity) *NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity {
	return &NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyTacAllOfProviderConfigurationComplexity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

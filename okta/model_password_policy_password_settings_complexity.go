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

// PasswordPolicyPasswordSettingsComplexity struct for PasswordPolicyPasswordSettingsComplexity
type PasswordPolicyPasswordSettingsComplexity struct {
	Dictionary *PasswordDictionary `json:"dictionary,omitempty"`
	ExcludeAttributes []string `json:"excludeAttributes,omitempty"`
	ExcludeUsername *bool `json:"excludeUsername,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
	MinLowerCase *int32 `json:"minLowerCase,omitempty"`
	MinNumber *int32 `json:"minNumber,omitempty"`
	MinSymbol *int32 `json:"minSymbol,omitempty"`
	MinUpperCase *int32 `json:"minUpperCase,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyPasswordSettingsComplexity PasswordPolicyPasswordSettingsComplexity

// NewPasswordPolicyPasswordSettingsComplexity instantiates a new PasswordPolicyPasswordSettingsComplexity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyPasswordSettingsComplexity() *PasswordPolicyPasswordSettingsComplexity {
	this := PasswordPolicyPasswordSettingsComplexity{}
	var excludeUsername bool = true
	this.ExcludeUsername = &excludeUsername
	return &this
}

// NewPasswordPolicyPasswordSettingsComplexityWithDefaults instantiates a new PasswordPolicyPasswordSettingsComplexity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsComplexityWithDefaults() *PasswordPolicyPasswordSettingsComplexity {
	this := PasswordPolicyPasswordSettingsComplexity{}
	var excludeUsername bool = true
	this.ExcludeUsername = &excludeUsername
	return &this
}

// GetDictionary returns the Dictionary field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionary() PasswordDictionary {
	if o == nil || o.Dictionary == nil {
		var ret PasswordDictionary
		return ret
	}
	return *o.Dictionary
}

// GetDictionaryOk returns a tuple with the Dictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionaryOk() (*PasswordDictionary, bool) {
	if o == nil || o.Dictionary == nil {
		return nil, false
	}
	return o.Dictionary, true
}

// HasDictionary returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasDictionary() bool {
	if o != nil && o.Dictionary != nil {
		return true
	}

	return false
}

// SetDictionary gets a reference to the given PasswordDictionary and assigns it to the Dictionary field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetDictionary(v PasswordDictionary) {
	o.Dictionary = &v
}

// GetExcludeAttributes returns the ExcludeAttributes field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeAttributes() []string {
	if o == nil || o.ExcludeAttributes == nil {
		var ret []string
		return ret
	}
	return o.ExcludeAttributes
}

// GetExcludeAttributesOk returns a tuple with the ExcludeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeAttributesOk() ([]string, bool) {
	if o == nil || o.ExcludeAttributes == nil {
		return nil, false
	}
	return o.ExcludeAttributes, true
}

// HasExcludeAttributes returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeAttributes() bool {
	if o != nil && o.ExcludeAttributes != nil {
		return true
	}

	return false
}

// SetExcludeAttributes gets a reference to the given []string and assigns it to the ExcludeAttributes field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetExcludeAttributes(v []string) {
	o.ExcludeAttributes = v
}

// GetExcludeUsername returns the ExcludeUsername field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeUsername() bool {
	if o == nil || o.ExcludeUsername == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsername
}

// GetExcludeUsernameOk returns a tuple with the ExcludeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeUsernameOk() (*bool, bool) {
	if o == nil || o.ExcludeUsername == nil {
		return nil, false
	}
	return o.ExcludeUsername, true
}

// HasExcludeUsername returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeUsername() bool {
	if o != nil && o.ExcludeUsername != nil {
		return true
	}

	return false
}

// SetExcludeUsername gets a reference to the given bool and assigns it to the ExcludeUsername field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetExcludeUsername(v bool) {
	o.ExcludeUsername = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMinLowerCase returns the MinLowerCase field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLowerCase() int32 {
	if o == nil || o.MinLowerCase == nil {
		var ret int32
		return ret
	}
	return *o.MinLowerCase
}

// GetMinLowerCaseOk returns a tuple with the MinLowerCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLowerCaseOk() (*int32, bool) {
	if o == nil || o.MinLowerCase == nil {
		return nil, false
	}
	return o.MinLowerCase, true
}

// HasMinLowerCase returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLowerCase() bool {
	if o != nil && o.MinLowerCase != nil {
		return true
	}

	return false
}

// SetMinLowerCase gets a reference to the given int32 and assigns it to the MinLowerCase field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinLowerCase(v int32) {
	o.MinLowerCase = &v
}

// GetMinNumber returns the MinNumber field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinNumber() int32 {
	if o == nil || o.MinNumber == nil {
		var ret int32
		return ret
	}
	return *o.MinNumber
}

// GetMinNumberOk returns a tuple with the MinNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinNumberOk() (*int32, bool) {
	if o == nil || o.MinNumber == nil {
		return nil, false
	}
	return o.MinNumber, true
}

// HasMinNumber returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinNumber() bool {
	if o != nil && o.MinNumber != nil {
		return true
	}

	return false
}

// SetMinNumber gets a reference to the given int32 and assigns it to the MinNumber field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinNumber(v int32) {
	o.MinNumber = &v
}

// GetMinSymbol returns the MinSymbol field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinSymbol() int32 {
	if o == nil || o.MinSymbol == nil {
		var ret int32
		return ret
	}
	return *o.MinSymbol
}

// GetMinSymbolOk returns a tuple with the MinSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinSymbolOk() (*int32, bool) {
	if o == nil || o.MinSymbol == nil {
		return nil, false
	}
	return o.MinSymbol, true
}

// HasMinSymbol returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinSymbol() bool {
	if o != nil && o.MinSymbol != nil {
		return true
	}

	return false
}

// SetMinSymbol gets a reference to the given int32 and assigns it to the MinSymbol field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinSymbol(v int32) {
	o.MinSymbol = &v
}

// GetMinUpperCase returns the MinUpperCase field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinUpperCase() int32 {
	if o == nil || o.MinUpperCase == nil {
		var ret int32
		return ret
	}
	return *o.MinUpperCase
}

// GetMinUpperCaseOk returns a tuple with the MinUpperCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinUpperCaseOk() (*int32, bool) {
	if o == nil || o.MinUpperCase == nil {
		return nil, false
	}
	return o.MinUpperCase, true
}

// HasMinUpperCase returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinUpperCase() bool {
	if o != nil && o.MinUpperCase != nil {
		return true
	}

	return false
}

// SetMinUpperCase gets a reference to the given int32 and assigns it to the MinUpperCase field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinUpperCase(v int32) {
	o.MinUpperCase = &v
}

func (o PasswordPolicyPasswordSettingsComplexity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dictionary != nil {
		toSerialize["dictionary"] = o.Dictionary
	}
	if o.ExcludeAttributes != nil {
		toSerialize["excludeAttributes"] = o.ExcludeAttributes
	}
	if o.ExcludeUsername != nil {
		toSerialize["excludeUsername"] = o.ExcludeUsername
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.MinLowerCase != nil {
		toSerialize["minLowerCase"] = o.MinLowerCase
	}
	if o.MinNumber != nil {
		toSerialize["minNumber"] = o.MinNumber
	}
	if o.MinSymbol != nil {
		toSerialize["minSymbol"] = o.MinSymbol
	}
	if o.MinUpperCase != nil {
		toSerialize["minUpperCase"] = o.MinUpperCase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyPasswordSettingsComplexity) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyPasswordSettingsComplexity := _PasswordPolicyPasswordSettingsComplexity{}

	err = json.Unmarshal(bytes, &varPasswordPolicyPasswordSettingsComplexity)
	if err == nil {
		*o = PasswordPolicyPasswordSettingsComplexity(varPasswordPolicyPasswordSettingsComplexity)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "dictionary")
		delete(additionalProperties, "excludeAttributes")
		delete(additionalProperties, "excludeUsername")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "minLowerCase")
		delete(additionalProperties, "minNumber")
		delete(additionalProperties, "minSymbol")
		delete(additionalProperties, "minUpperCase")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyPasswordSettingsComplexity struct {
	value *PasswordPolicyPasswordSettingsComplexity
	isSet bool
}

func (v NullablePasswordPolicyPasswordSettingsComplexity) Get() *PasswordPolicyPasswordSettingsComplexity {
	return v.value
}

func (v *NullablePasswordPolicyPasswordSettingsComplexity) Set(val *PasswordPolicyPasswordSettingsComplexity) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyPasswordSettingsComplexity) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyPasswordSettingsComplexity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyPasswordSettingsComplexity(val *PasswordPolicyPasswordSettingsComplexity) *NullablePasswordPolicyPasswordSettingsComplexity {
	return &NullablePasswordPolicyPasswordSettingsComplexity{value: val, isSet: true}
}

func (v NullablePasswordPolicyPasswordSettingsComplexity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyPasswordSettingsComplexity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


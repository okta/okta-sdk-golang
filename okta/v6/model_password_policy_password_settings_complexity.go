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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PasswordPolicyPasswordSettingsComplexity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyPasswordSettingsComplexity{}

// PasswordPolicyPasswordSettingsComplexity Complexity settings
type PasswordPolicyPasswordSettingsComplexity struct {
	Dictionary *PasswordDictionary `json:"dictionary,omitempty"`
	// The User profile attributes whose values must be excluded from the password: currently only supports `firstName` and `lastName`
	ExcludeAttributes []string `json:"excludeAttributes,omitempty"`
	// Indicates if the Username must be excluded from the password
	ExcludeUsername *bool `json:"excludeUsername,omitempty"`
	// Minimum password length
	MinLength *int32 `json:"minLength,omitempty"`
	// Indicates if a password must contain at least one lower case letter: `0` indicates no, `1` indicates yes
	MinLowerCase *int32 `json:"minLowerCase,omitempty"`
	// Indicates if a password must contain at least one number: `0` indicates no, `1` indicates yes
	MinNumber *int32 `json:"minNumber,omitempty"`
	// Indicates if a password must contain at least one symbol (For example: !@#$%^&*): `0` indicates no, `1` indicates yes
	MinSymbol *int32 `json:"minSymbol,omitempty"`
	// Indicates if a password must contain at least one upper case letter: `0` indicates no, `1` indicates yes
	MinUpperCase *int32 `json:"minUpperCase,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Use an [Expression Language](https://developer.okta.com/docs/reference/okta-expression-language-in-identity-engine/) expression to block a word from being used in a password. You can only block one word per expression. Use the `OR` operator to connect multiple expressions to block multiple words.
	OelStatement         *string `json:"oelStatement,omitempty"`
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
	var minLength int32 = 8
	this.MinLength = &minLength
	var minLowerCase int32 = 1
	this.MinLowerCase = &minLowerCase
	var minNumber int32 = 1
	this.MinNumber = &minNumber
	var minSymbol int32 = 1
	this.MinSymbol = &minSymbol
	var minUpperCase int32 = 1
	this.MinUpperCase = &minUpperCase
	return &this
}

// NewPasswordPolicyPasswordSettingsComplexityWithDefaults instantiates a new PasswordPolicyPasswordSettingsComplexity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyPasswordSettingsComplexityWithDefaults() *PasswordPolicyPasswordSettingsComplexity {
	this := PasswordPolicyPasswordSettingsComplexity{}
	var excludeUsername bool = true
	this.ExcludeUsername = &excludeUsername
	var minLength int32 = 8
	this.MinLength = &minLength
	var minLowerCase int32 = 1
	this.MinLowerCase = &minLowerCase
	var minNumber int32 = 1
	this.MinNumber = &minNumber
	var minSymbol int32 = 1
	this.MinSymbol = &minSymbol
	var minUpperCase int32 = 1
	this.MinUpperCase = &minUpperCase
	return &this
}

// GetDictionary returns the Dictionary field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionary() PasswordDictionary {
	if o == nil || IsNil(o.Dictionary) {
		var ret PasswordDictionary
		return ret
	}
	return *o.Dictionary
}

// GetDictionaryOk returns a tuple with the Dictionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetDictionaryOk() (*PasswordDictionary, bool) {
	if o == nil || IsNil(o.Dictionary) {
		return nil, false
	}
	return o.Dictionary, true
}

// HasDictionary returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasDictionary() bool {
	if o != nil && !IsNil(o.Dictionary) {
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
	if o == nil || IsNil(o.ExcludeAttributes) {
		var ret []string
		return ret
	}
	return o.ExcludeAttributes
}

// GetExcludeAttributesOk returns a tuple with the ExcludeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeAttributes) {
		return nil, false
	}
	return o.ExcludeAttributes, true
}

// HasExcludeAttributes returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeAttributes() bool {
	if o != nil && !IsNil(o.ExcludeAttributes) {
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
	if o == nil || IsNil(o.ExcludeUsername) {
		var ret bool
		return ret
	}
	return *o.ExcludeUsername
}

// GetExcludeUsernameOk returns a tuple with the ExcludeUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetExcludeUsernameOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUsername) {
		return nil, false
	}
	return o.ExcludeUsername, true
}

// HasExcludeUsername returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasExcludeUsername() bool {
	if o != nil && !IsNil(o.ExcludeUsername) {
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
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
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
	if o == nil || IsNil(o.MinLowerCase) {
		var ret int32
		return ret
	}
	return *o.MinLowerCase
}

// GetMinLowerCaseOk returns a tuple with the MinLowerCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinLowerCaseOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLowerCase) {
		return nil, false
	}
	return o.MinLowerCase, true
}

// HasMinLowerCase returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinLowerCase() bool {
	if o != nil && !IsNil(o.MinLowerCase) {
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
	if o == nil || IsNil(o.MinNumber) {
		var ret int32
		return ret
	}
	return *o.MinNumber
}

// GetMinNumberOk returns a tuple with the MinNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.MinNumber) {
		return nil, false
	}
	return o.MinNumber, true
}

// HasMinNumber returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinNumber() bool {
	if o != nil && !IsNil(o.MinNumber) {
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
	if o == nil || IsNil(o.MinSymbol) {
		var ret int32
		return ret
	}
	return *o.MinSymbol
}

// GetMinSymbolOk returns a tuple with the MinSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinSymbolOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSymbol) {
		return nil, false
	}
	return o.MinSymbol, true
}

// HasMinSymbol returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinSymbol() bool {
	if o != nil && !IsNil(o.MinSymbol) {
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
	if o == nil || IsNil(o.MinUpperCase) {
		var ret int32
		return ret
	}
	return *o.MinUpperCase
}

// GetMinUpperCaseOk returns a tuple with the MinUpperCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetMinUpperCaseOk() (*int32, bool) {
	if o == nil || IsNil(o.MinUpperCase) {
		return nil, false
	}
	return o.MinUpperCase, true
}

// HasMinUpperCase returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasMinUpperCase() bool {
	if o != nil && !IsNil(o.MinUpperCase) {
		return true
	}

	return false
}

// SetMinUpperCase gets a reference to the given int32 and assigns it to the MinUpperCase field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetMinUpperCase(v int32) {
	o.MinUpperCase = &v
}

// GetOelStatement returns the OelStatement field value if set, zero value otherwise.
func (o *PasswordPolicyPasswordSettingsComplexity) GetOelStatement() string {
	if o == nil || IsNil(o.OelStatement) {
		var ret string
		return ret
	}
	return *o.OelStatement
}

// GetOelStatementOk returns a tuple with the OelStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) GetOelStatementOk() (*string, bool) {
	if o == nil || IsNil(o.OelStatement) {
		return nil, false
	}
	return o.OelStatement, true
}

// HasOelStatement returns a boolean if a field has been set.
func (o *PasswordPolicyPasswordSettingsComplexity) HasOelStatement() bool {
	if o != nil && !IsNil(o.OelStatement) {
		return true
	}

	return false
}

// SetOelStatement gets a reference to the given string and assigns it to the OelStatement field.
func (o *PasswordPolicyPasswordSettingsComplexity) SetOelStatement(v string) {
	o.OelStatement = &v
}

func (o PasswordPolicyPasswordSettingsComplexity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyPasswordSettingsComplexity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dictionary) {
		toSerialize["dictionary"] = o.Dictionary
	}
	if !IsNil(o.ExcludeAttributes) {
		toSerialize["excludeAttributes"] = o.ExcludeAttributes
	}
	if !IsNil(o.ExcludeUsername) {
		toSerialize["excludeUsername"] = o.ExcludeUsername
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.MinLowerCase) {
		toSerialize["minLowerCase"] = o.MinLowerCase
	}
	if !IsNil(o.MinNumber) {
		toSerialize["minNumber"] = o.MinNumber
	}
	if !IsNil(o.MinSymbol) {
		toSerialize["minSymbol"] = o.MinSymbol
	}
	if !IsNil(o.MinUpperCase) {
		toSerialize["minUpperCase"] = o.MinUpperCase
	}
	if !IsNil(o.OelStatement) {
		toSerialize["oelStatement"] = o.OelStatement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyPasswordSettingsComplexity) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyPasswordSettingsComplexity := _PasswordPolicyPasswordSettingsComplexity{}

	err = json.Unmarshal(data, &varPasswordPolicyPasswordSettingsComplexity)

	if err != nil {
		return err
	}

	*o = PasswordPolicyPasswordSettingsComplexity(varPasswordPolicyPasswordSettingsComplexity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dictionary")
		delete(additionalProperties, "excludeAttributes")
		delete(additionalProperties, "excludeUsername")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "minLowerCase")
		delete(additionalProperties, "minNumber")
		delete(additionalProperties, "minSymbol")
		delete(additionalProperties, "minUpperCase")
		delete(additionalProperties, "oelStatement")
		o.AdditionalProperties = additionalProperties
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

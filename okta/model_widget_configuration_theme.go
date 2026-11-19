/*
Okta Admin Management API

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

// checks if the WidgetConfigurationTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationTheme{}

// WidgetConfigurationTheme struct for WidgetConfigurationTheme
type WidgetConfigurationTheme struct {
	Tokens *WidgetConfigurationThemeTokens `json:"tokens,omitempty"`
}

// NewWidgetConfigurationTheme instantiates a new WidgetConfigurationTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationTheme() *WidgetConfigurationTheme {
	this := WidgetConfigurationTheme{}
	return &this
}

// NewWidgetConfigurationThemeWithDefaults instantiates a new WidgetConfigurationTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationThemeWithDefaults() *WidgetConfigurationTheme {
	this := WidgetConfigurationTheme{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *WidgetConfigurationTheme) GetTokens() WidgetConfigurationThemeTokens {
	if o == nil || IsNil(o.Tokens) {
		var ret WidgetConfigurationThemeTokens
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationTheme) GetTokensOk() (*WidgetConfigurationThemeTokens, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *WidgetConfigurationTheme) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given WidgetConfigurationThemeTokens and assigns it to the Tokens field.
func (o *WidgetConfigurationTheme) SetTokens(v WidgetConfigurationThemeTokens) {
	o.Tokens = &v
}

func (o WidgetConfigurationTheme) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	return toSerialize, nil
}

type NullableWidgetConfigurationTheme struct {
	value *WidgetConfigurationTheme
	isSet bool
}

func (v NullableWidgetConfigurationTheme) Get() *WidgetConfigurationTheme {
	return v.value
}

func (v *NullableWidgetConfigurationTheme) Set(val *WidgetConfigurationTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationTheme(val *WidgetConfigurationTheme) *NullableWidgetConfigurationTheme {
	return &NullableWidgetConfigurationTheme{value: val, isSet: true}
}

func (v NullableWidgetConfigurationTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

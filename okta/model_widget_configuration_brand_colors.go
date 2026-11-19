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

// checks if the WidgetConfigurationBrandColors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationBrandColors{}

// WidgetConfigurationBrandColors Brand color used to generate the widget's color palette
type WidgetConfigurationBrandColors struct {
	// Primary brand color (hex value, e.g. \"#1662dd\")
	PrimaryColor *string `json:"primaryColor,omitempty"`
}

// NewWidgetConfigurationBrandColors instantiates a new WidgetConfigurationBrandColors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationBrandColors() *WidgetConfigurationBrandColors {
	this := WidgetConfigurationBrandColors{}
	return &this
}

// NewWidgetConfigurationBrandColorsWithDefaults instantiates a new WidgetConfigurationBrandColors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationBrandColorsWithDefaults() *WidgetConfigurationBrandColors {
	this := WidgetConfigurationBrandColors{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *WidgetConfigurationBrandColors) GetPrimaryColor() string {
	if o == nil || IsNil(o.PrimaryColor) {
		var ret string
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationBrandColors) GetPrimaryColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryColor) {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *WidgetConfigurationBrandColors) HasPrimaryColor() bool {
	if o != nil && !IsNil(o.PrimaryColor) {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given string and assigns it to the PrimaryColor field.
func (o *WidgetConfigurationBrandColors) SetPrimaryColor(v string) {
	o.PrimaryColor = &v
}

func (o WidgetConfigurationBrandColors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationBrandColors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryColor) {
		toSerialize["primaryColor"] = o.PrimaryColor
	}
	return toSerialize, nil
}

type NullableWidgetConfigurationBrandColors struct {
	value *WidgetConfigurationBrandColors
	isSet bool
}

func (v NullableWidgetConfigurationBrandColors) Get() *WidgetConfigurationBrandColors {
	return v.value
}

func (v *NullableWidgetConfigurationBrandColors) Set(val *WidgetConfigurationBrandColors) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationBrandColors) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationBrandColors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationBrandColors(val *WidgetConfigurationBrandColors) *NullableWidgetConfigurationBrandColors {
	return &NullableWidgetConfigurationBrandColors{value: val, isSet: true}
}

func (v NullableWidgetConfigurationBrandColors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationBrandColors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

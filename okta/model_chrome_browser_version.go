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

// checks if the ChromeBrowserVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChromeBrowserVersion{}

// ChromeBrowserVersion Current version of the Chrome Browser
type ChromeBrowserVersion struct {
	Minimum              *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChromeBrowserVersion ChromeBrowserVersion

// NewChromeBrowserVersion instantiates a new ChromeBrowserVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChromeBrowserVersion() *ChromeBrowserVersion {
	this := ChromeBrowserVersion{}
	return &this
}

// NewChromeBrowserVersionWithDefaults instantiates a new ChromeBrowserVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChromeBrowserVersionWithDefaults() *ChromeBrowserVersion {
	this := ChromeBrowserVersion{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *ChromeBrowserVersion) GetMinimum() string {
	if o == nil || IsNil(o.Minimum) {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChromeBrowserVersion) GetMinimumOk() (*string, bool) {
	if o == nil || IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *ChromeBrowserVersion) HasMinimum() bool {
	if o != nil && !IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *ChromeBrowserVersion) SetMinimum(v string) {
	o.Minimum = &v
}

func (o ChromeBrowserVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChromeBrowserVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChromeBrowserVersion) UnmarshalJSON(data []byte) (err error) {
	varChromeBrowserVersion := _ChromeBrowserVersion{}

	err = json.Unmarshal(data, &varChromeBrowserVersion)

	if err != nil {
		return err
	}

	*o = ChromeBrowserVersion(varChromeBrowserVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChromeBrowserVersion struct {
	value *ChromeBrowserVersion
	isSet bool
}

func (v NullableChromeBrowserVersion) Get() *ChromeBrowserVersion {
	return v.value
}

func (v *NullableChromeBrowserVersion) Set(val *ChromeBrowserVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableChromeBrowserVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableChromeBrowserVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChromeBrowserVersion(val *ChromeBrowserVersion) *NullableChromeBrowserVersion {
	return &NullableChromeBrowserVersion{value: val, isSet: true}
}

func (v NullableChromeBrowserVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChromeBrowserVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

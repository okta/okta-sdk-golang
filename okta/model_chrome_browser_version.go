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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

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
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChromeBrowserVersion) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *ChromeBrowserVersion) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *ChromeBrowserVersion) SetMinimum(v string) {
	o.Minimum = &v
}

func (o ChromeBrowserVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChromeBrowserVersion) UnmarshalJSON(bytes []byte) (err error) {
	varChromeBrowserVersion := _ChromeBrowserVersion{}

	err = json.Unmarshal(bytes, &varChromeBrowserVersion)
	if err == nil {
		*o = ChromeBrowserVersion(varChromeBrowserVersion)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

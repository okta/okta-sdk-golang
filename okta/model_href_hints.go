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

// HrefHints Describes allowed HTTP verbs for the `href`
type HrefHints struct {
	Allow []string `json:"allow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HrefHints HrefHints

// NewHrefHints instantiates a new HrefHints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefHints() *HrefHints {
	this := HrefHints{}
	return &this
}

// NewHrefHintsWithDefaults instantiates a new HrefHints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefHintsWithDefaults() *HrefHints {
	this := HrefHints{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *HrefHints) GetAllow() []string {
	if o == nil || o.Allow == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefHints) GetAllowOk() ([]string, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *HrefHints) HasAllow() bool {
	if o != nil && o.Allow != nil {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *HrefHints) SetAllow(v []string) {
	o.Allow = v
}

func (o HrefHints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HrefHints) UnmarshalJSON(bytes []byte) (err error) {
	varHrefHints := _HrefHints{}

	err = json.Unmarshal(bytes, &varHrefHints)
	if err == nil {
		*o = HrefHints(varHrefHints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allow")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableHrefHints struct {
	value *HrefHints
	isSet bool
}

func (v NullableHrefHints) Get() *HrefHints {
	return v.value
}

func (v *NullableHrefHints) Set(val *HrefHints) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefHints) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefHints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefHints(val *HrefHints) *NullableHrefHints {
	return &NullableHrefHints{value: val, isSet: true}
}

func (v NullableHrefHints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefHints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


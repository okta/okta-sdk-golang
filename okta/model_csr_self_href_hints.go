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

// checks if the CsrSelfHrefHints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsrSelfHrefHints{}

// CsrSelfHrefHints Describes allowed HTTP verbs for the `href`
type CsrSelfHrefHints struct {
	Allow                []string `json:"allow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CsrSelfHrefHints CsrSelfHrefHints

// NewCsrSelfHrefHints instantiates a new CsrSelfHrefHints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrSelfHrefHints() *CsrSelfHrefHints {
	this := CsrSelfHrefHints{}
	return &this
}

// NewCsrSelfHrefHintsWithDefaults instantiates a new CsrSelfHrefHints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrSelfHrefHintsWithDefaults() *CsrSelfHrefHints {
	this := CsrSelfHrefHints{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *CsrSelfHrefHints) GetAllow() []string {
	if o == nil || IsNil(o.Allow) {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrSelfHrefHints) GetAllowOk() ([]string, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *CsrSelfHrefHints) HasAllow() bool {
	if o != nil && !IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *CsrSelfHrefHints) SetAllow(v []string) {
	o.Allow = v
}

func (o CsrSelfHrefHints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsrSelfHrefHints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CsrSelfHrefHints) UnmarshalJSON(data []byte) (err error) {
	varCsrSelfHrefHints := _CsrSelfHrefHints{}

	err = json.Unmarshal(data, &varCsrSelfHrefHints)

	if err != nil {
		return err
	}

	*o = CsrSelfHrefHints(varCsrSelfHrefHints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCsrSelfHrefHints struct {
	value *CsrSelfHrefHints
	isSet bool
}

func (v NullableCsrSelfHrefHints) Get() *CsrSelfHrefHints {
	return v.value
}

func (v *NullableCsrSelfHrefHints) Set(val *CsrSelfHrefHints) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrSelfHrefHints) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrSelfHrefHints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrSelfHrefHints(val *CsrSelfHrefHints) *NullableCsrSelfHrefHints {
	return &NullableCsrSelfHrefHints{value: val, isSet: true}
}

func (v NullableCsrSelfHrefHints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrSelfHrefHints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

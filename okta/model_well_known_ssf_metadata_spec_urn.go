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

// checks if the WellKnownSSFMetadataSpecUrn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownSSFMetadataSpecUrn{}

// WellKnownSSFMetadataSpecUrn struct for WellKnownSSFMetadataSpecUrn
type WellKnownSSFMetadataSpecUrn struct {
	// The URN that describes the specification of the protocol being used
	SpecUrn              *string `json:"spec_urn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownSSFMetadataSpecUrn WellKnownSSFMetadataSpecUrn

// NewWellKnownSSFMetadataSpecUrn instantiates a new WellKnownSSFMetadataSpecUrn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownSSFMetadataSpecUrn() *WellKnownSSFMetadataSpecUrn {
	this := WellKnownSSFMetadataSpecUrn{}
	return &this
}

// NewWellKnownSSFMetadataSpecUrnWithDefaults instantiates a new WellKnownSSFMetadataSpecUrn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownSSFMetadataSpecUrnWithDefaults() *WellKnownSSFMetadataSpecUrn {
	this := WellKnownSSFMetadataSpecUrn{}
	return &this
}

// GetSpecUrn returns the SpecUrn field value if set, zero value otherwise.
func (o *WellKnownSSFMetadataSpecUrn) GetSpecUrn() string {
	if o == nil || IsNil(o.SpecUrn) {
		var ret string
		return ret
	}
	return *o.SpecUrn
}

// GetSpecUrnOk returns a tuple with the SpecUrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownSSFMetadataSpecUrn) GetSpecUrnOk() (*string, bool) {
	if o == nil || IsNil(o.SpecUrn) {
		return nil, false
	}
	return o.SpecUrn, true
}

// HasSpecUrn returns a boolean if a field has been set.
func (o *WellKnownSSFMetadataSpecUrn) HasSpecUrn() bool {
	if o != nil && !IsNil(o.SpecUrn) {
		return true
	}

	return false
}

// SetSpecUrn gets a reference to the given string and assigns it to the SpecUrn field.
func (o *WellKnownSSFMetadataSpecUrn) SetSpecUrn(v string) {
	o.SpecUrn = &v
}

func (o WellKnownSSFMetadataSpecUrn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownSSFMetadataSpecUrn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpecUrn) {
		toSerialize["spec_urn"] = o.SpecUrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownSSFMetadataSpecUrn) UnmarshalJSON(data []byte) (err error) {
	varWellKnownSSFMetadataSpecUrn := _WellKnownSSFMetadataSpecUrn{}

	err = json.Unmarshal(data, &varWellKnownSSFMetadataSpecUrn)

	if err != nil {
		return err
	}

	*o = WellKnownSSFMetadataSpecUrn(varWellKnownSSFMetadataSpecUrn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "spec_urn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownSSFMetadataSpecUrn struct {
	value *WellKnownSSFMetadataSpecUrn
	isSet bool
}

func (v NullableWellKnownSSFMetadataSpecUrn) Get() *WellKnownSSFMetadataSpecUrn {
	return v.value
}

func (v *NullableWellKnownSSFMetadataSpecUrn) Set(val *WellKnownSSFMetadataSpecUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownSSFMetadataSpecUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownSSFMetadataSpecUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownSSFMetadataSpecUrn(val *WellKnownSSFMetadataSpecUrn) *NullableWellKnownSSFMetadataSpecUrn {
	return &NullableWellKnownSSFMetadataSpecUrn{value: val, isSet: true}
}

func (v NullableWellKnownSSFMetadataSpecUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownSSFMetadataSpecUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

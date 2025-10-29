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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ReplaceUserClassification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceUserClassification{}

// ReplaceUserClassification struct for ReplaceUserClassification
type ReplaceUserClassification struct {
	// The type of user classification
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReplaceUserClassification ReplaceUserClassification

// NewReplaceUserClassification instantiates a new ReplaceUserClassification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceUserClassification() *ReplaceUserClassification {
	this := ReplaceUserClassification{}
	return &this
}

// NewReplaceUserClassificationWithDefaults instantiates a new ReplaceUserClassification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceUserClassificationWithDefaults() *ReplaceUserClassification {
	this := ReplaceUserClassification{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReplaceUserClassification) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceUserClassification) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReplaceUserClassification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReplaceUserClassification) SetType(v string) {
	o.Type = &v
}

func (o ReplaceUserClassification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceUserClassification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReplaceUserClassification) UnmarshalJSON(data []byte) (err error) {
	varReplaceUserClassification := _ReplaceUserClassification{}

	err = json.Unmarshal(data, &varReplaceUserClassification)

	if err != nil {
		return err
	}

	*o = ReplaceUserClassification(varReplaceUserClassification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplaceUserClassification struct {
	value *ReplaceUserClassification
	isSet bool
}

func (v NullableReplaceUserClassification) Get() *ReplaceUserClassification {
	return v.value
}

func (v *NullableReplaceUserClassification) Set(val *ReplaceUserClassification) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceUserClassification) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceUserClassification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceUserClassification(val *ReplaceUserClassification) *NullableReplaceUserClassification {
	return &NullableReplaceUserClassification{value: val, isSet: true}
}

func (v NullableReplaceUserClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceUserClassification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

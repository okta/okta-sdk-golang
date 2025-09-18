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

// checks if the ByDateTimeExpiry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ByDateTimeExpiry{}

// ByDateTimeExpiry struct for ByDateTimeExpiry
type ByDateTimeExpiry struct {
	AdditionalProperties map[string]interface{}
}

type _ByDateTimeExpiry ByDateTimeExpiry

// NewByDateTimeExpiry instantiates a new ByDateTimeExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewByDateTimeExpiry() *ByDateTimeExpiry {
	this := ByDateTimeExpiry{}
	return &this
}

// NewByDateTimeExpiryWithDefaults instantiates a new ByDateTimeExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewByDateTimeExpiryWithDefaults() *ByDateTimeExpiry {
	this := ByDateTimeExpiry{}
	return &this
}

func (o ByDateTimeExpiry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ByDateTimeExpiry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ByDateTimeExpiry) UnmarshalJSON(data []byte) (err error) {
	varByDateTimeExpiry := _ByDateTimeExpiry{}

	err = json.Unmarshal(data, &varByDateTimeExpiry)

	if err != nil {
		return err
	}

	*o = ByDateTimeExpiry(varByDateTimeExpiry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableByDateTimeExpiry struct {
	value *ByDateTimeExpiry
	isSet bool
}

func (v NullableByDateTimeExpiry) Get() *ByDateTimeExpiry {
	return v.value
}

func (v *NullableByDateTimeExpiry) Set(val *ByDateTimeExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableByDateTimeExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableByDateTimeExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableByDateTimeExpiry(val *ByDateTimeExpiry) *NullableByDateTimeExpiry {
	return &NullableByDateTimeExpiry{value: val, isSet: true}
}

func (v NullableByDateTimeExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableByDateTimeExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

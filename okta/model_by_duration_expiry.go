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

// checks if the ByDurationExpiry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ByDurationExpiry{}

// ByDurationExpiry A time duration specified as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). Must be between 1 and 180 days inclusive.
type ByDurationExpiry struct {
	AdditionalProperties map[string]interface{}
}

type _ByDurationExpiry ByDurationExpiry

// NewByDurationExpiry instantiates a new ByDurationExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewByDurationExpiry() *ByDurationExpiry {
	this := ByDurationExpiry{}
	return &this
}

// NewByDurationExpiryWithDefaults instantiates a new ByDurationExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewByDurationExpiryWithDefaults() *ByDurationExpiry {
	this := ByDurationExpiry{}
	return &this
}

func (o ByDurationExpiry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ByDurationExpiry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ByDurationExpiry) UnmarshalJSON(data []byte) (err error) {
	varByDurationExpiry := _ByDurationExpiry{}

	err = json.Unmarshal(data, &varByDurationExpiry)

	if err != nil {
		return err
	}

	*o = ByDurationExpiry(varByDurationExpiry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableByDurationExpiry struct {
	value *ByDurationExpiry
	isSet bool
}

func (v NullableByDurationExpiry) Get() *ByDurationExpiry {
	return v.value
}

func (v *NullableByDurationExpiry) Set(val *ByDurationExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableByDurationExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableByDurationExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableByDurationExpiry(val *ByDurationExpiry) *NullableByDurationExpiry {
	return &NullableByDurationExpiry{value: val, isSet: true}
}

func (v NullableByDurationExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableByDurationExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
	"reflect"
	"strings"
)

// checks if the BySkipCountAuthenticatorGracePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BySkipCountAuthenticatorGracePeriod{}

// BySkipCountAuthenticatorGracePeriod struct for BySkipCountAuthenticatorGracePeriod
type BySkipCountAuthenticatorGracePeriod struct {
	EnrollmentPolicyAuthenticatorGracePeriod
	// The number of times the user can skip enrolling the corresponding authenticator before it becomes mandatory.
	SkipCount            *int32 `json:"skipCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BySkipCountAuthenticatorGracePeriod BySkipCountAuthenticatorGracePeriod

// NewBySkipCountAuthenticatorGracePeriod instantiates a new BySkipCountAuthenticatorGracePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBySkipCountAuthenticatorGracePeriod() *BySkipCountAuthenticatorGracePeriod {
	this := BySkipCountAuthenticatorGracePeriod{}
	return &this
}

// NewBySkipCountAuthenticatorGracePeriodWithDefaults instantiates a new BySkipCountAuthenticatorGracePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBySkipCountAuthenticatorGracePeriodWithDefaults() *BySkipCountAuthenticatorGracePeriod {
	this := BySkipCountAuthenticatorGracePeriod{}
	return &this
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *BySkipCountAuthenticatorGracePeriod) GetSkipCount() int32 {
	if o == nil || IsNil(o.SkipCount) {
		var ret int32
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BySkipCountAuthenticatorGracePeriod) GetSkipCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SkipCount) {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *BySkipCountAuthenticatorGracePeriod) HasSkipCount() bool {
	if o != nil && !IsNil(o.SkipCount) {
		return true
	}

	return false
}

// SetSkipCount gets a reference to the given int32 and assigns it to the SkipCount field.
func (o *BySkipCountAuthenticatorGracePeriod) SetSkipCount(v int32) {
	o.SkipCount = &v
}

func (o BySkipCountAuthenticatorGracePeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BySkipCountAuthenticatorGracePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEnrollmentPolicyAuthenticatorGracePeriod, errEnrollmentPolicyAuthenticatorGracePeriod := json.Marshal(o.EnrollmentPolicyAuthenticatorGracePeriod)
	if errEnrollmentPolicyAuthenticatorGracePeriod != nil {
		return map[string]interface{}{}, errEnrollmentPolicyAuthenticatorGracePeriod
	}
	errEnrollmentPolicyAuthenticatorGracePeriod = json.Unmarshal([]byte(serializedEnrollmentPolicyAuthenticatorGracePeriod), &toSerialize)
	if errEnrollmentPolicyAuthenticatorGracePeriod != nil {
		return map[string]interface{}{}, errEnrollmentPolicyAuthenticatorGracePeriod
	}
	if !IsNil(o.SkipCount) {
		toSerialize["skipCount"] = o.SkipCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BySkipCountAuthenticatorGracePeriod) UnmarshalJSON(data []byte) (err error) {
	type BySkipCountAuthenticatorGracePeriodWithoutEmbeddedStruct struct {
		// The number of times the user can skip enrolling the corresponding authenticator before it becomes mandatory.
		SkipCount *int32 `json:"skipCount,omitempty"`
	}

	varBySkipCountAuthenticatorGracePeriodWithoutEmbeddedStruct := BySkipCountAuthenticatorGracePeriodWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBySkipCountAuthenticatorGracePeriodWithoutEmbeddedStruct)
	if err == nil {
		varBySkipCountAuthenticatorGracePeriod := _BySkipCountAuthenticatorGracePeriod{}
		varBySkipCountAuthenticatorGracePeriod.SkipCount = varBySkipCountAuthenticatorGracePeriodWithoutEmbeddedStruct.SkipCount
		*o = BySkipCountAuthenticatorGracePeriod(varBySkipCountAuthenticatorGracePeriod)
	} else {
		return err
	}

	varBySkipCountAuthenticatorGracePeriod := _BySkipCountAuthenticatorGracePeriod{}

	err = json.Unmarshal(data, &varBySkipCountAuthenticatorGracePeriod)
	if err == nil {
		o.EnrollmentPolicyAuthenticatorGracePeriod = varBySkipCountAuthenticatorGracePeriod.EnrollmentPolicyAuthenticatorGracePeriod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "skipCount")

		// remove fields from embedded structs
		reflectEnrollmentPolicyAuthenticatorGracePeriod := reflect.ValueOf(o.EnrollmentPolicyAuthenticatorGracePeriod)
		for i := 0; i < reflectEnrollmentPolicyAuthenticatorGracePeriod.Type().NumField(); i++ {
			t := reflectEnrollmentPolicyAuthenticatorGracePeriod.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBySkipCountAuthenticatorGracePeriod struct {
	value *BySkipCountAuthenticatorGracePeriod
	isSet bool
}

func (v NullableBySkipCountAuthenticatorGracePeriod) Get() *BySkipCountAuthenticatorGracePeriod {
	return v.value
}

func (v *NullableBySkipCountAuthenticatorGracePeriod) Set(val *BySkipCountAuthenticatorGracePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableBySkipCountAuthenticatorGracePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableBySkipCountAuthenticatorGracePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBySkipCountAuthenticatorGracePeriod(val *BySkipCountAuthenticatorGracePeriod) *NullableBySkipCountAuthenticatorGracePeriod {
	return &NullableBySkipCountAuthenticatorGracePeriod{value: val, isSet: true}
}

func (v NullableBySkipCountAuthenticatorGracePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBySkipCountAuthenticatorGracePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

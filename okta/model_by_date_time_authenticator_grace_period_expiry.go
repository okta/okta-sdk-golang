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
	"reflect"
	"strings"
)

// checks if the ByDateTimeAuthenticatorGracePeriodExpiry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ByDateTimeAuthenticatorGracePeriodExpiry{}

// ByDateTimeAuthenticatorGracePeriodExpiry struct for ByDateTimeAuthenticatorGracePeriodExpiry
type ByDateTimeAuthenticatorGracePeriodExpiry struct {
	EnrollmentPolicyAuthenticatorGracePeriod
	// The expiry date for a `BY_DATE_TIME` grace period type. Valid format: `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`  For example, `2025-01-01T18:30:45.000Z`
	Expiry               *string `json:"expiry,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ByDateTimeAuthenticatorGracePeriodExpiry ByDateTimeAuthenticatorGracePeriodExpiry

// NewByDateTimeAuthenticatorGracePeriodExpiry instantiates a new ByDateTimeAuthenticatorGracePeriodExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewByDateTimeAuthenticatorGracePeriodExpiry() *ByDateTimeAuthenticatorGracePeriodExpiry {
	this := ByDateTimeAuthenticatorGracePeriodExpiry{}
	return &this
}

// NewByDateTimeAuthenticatorGracePeriodExpiryWithDefaults instantiates a new ByDateTimeAuthenticatorGracePeriodExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewByDateTimeAuthenticatorGracePeriodExpiryWithDefaults() *ByDateTimeAuthenticatorGracePeriodExpiry {
	this := ByDateTimeAuthenticatorGracePeriodExpiry{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ByDateTimeAuthenticatorGracePeriodExpiry) GetExpiry() string {
	if o == nil || IsNil(o.Expiry) {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ByDateTimeAuthenticatorGracePeriodExpiry) GetExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ByDateTimeAuthenticatorGracePeriodExpiry) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *ByDateTimeAuthenticatorGracePeriodExpiry) SetExpiry(v string) {
	o.Expiry = &v
}

func (o ByDateTimeAuthenticatorGracePeriodExpiry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ByDateTimeAuthenticatorGracePeriodExpiry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEnrollmentPolicyAuthenticatorGracePeriod, errEnrollmentPolicyAuthenticatorGracePeriod := json.Marshal(o.EnrollmentPolicyAuthenticatorGracePeriod)
	if errEnrollmentPolicyAuthenticatorGracePeriod != nil {
		return map[string]interface{}{}, errEnrollmentPolicyAuthenticatorGracePeriod
	}
	errEnrollmentPolicyAuthenticatorGracePeriod = json.Unmarshal([]byte(serializedEnrollmentPolicyAuthenticatorGracePeriod), &toSerialize)
	if errEnrollmentPolicyAuthenticatorGracePeriod != nil {
		return map[string]interface{}{}, errEnrollmentPolicyAuthenticatorGracePeriod
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ByDateTimeAuthenticatorGracePeriodExpiry) UnmarshalJSON(data []byte) (err error) {
	type ByDateTimeAuthenticatorGracePeriodExpiryWithoutEmbeddedStruct struct {
		// The expiry date for a `BY_DATE_TIME` grace period type. Valid format: `yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`  For example, `2025-01-01T18:30:45.000Z`
		Expiry *string `json:"expiry,omitempty"`
	}

	varByDateTimeAuthenticatorGracePeriodExpiryWithoutEmbeddedStruct := ByDateTimeAuthenticatorGracePeriodExpiryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varByDateTimeAuthenticatorGracePeriodExpiryWithoutEmbeddedStruct)
	if err == nil {
		varByDateTimeAuthenticatorGracePeriodExpiry := _ByDateTimeAuthenticatorGracePeriodExpiry{}
		varByDateTimeAuthenticatorGracePeriodExpiry.Expiry = varByDateTimeAuthenticatorGracePeriodExpiryWithoutEmbeddedStruct.Expiry
		*o = ByDateTimeAuthenticatorGracePeriodExpiry(varByDateTimeAuthenticatorGracePeriodExpiry)
	} else {
		return err
	}

	varByDateTimeAuthenticatorGracePeriodExpiry := _ByDateTimeAuthenticatorGracePeriodExpiry{}

	err = json.Unmarshal(data, &varByDateTimeAuthenticatorGracePeriodExpiry)
	if err == nil {
		o.EnrollmentPolicyAuthenticatorGracePeriod = varByDateTimeAuthenticatorGracePeriodExpiry.EnrollmentPolicyAuthenticatorGracePeriod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiry")

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

type NullableByDateTimeAuthenticatorGracePeriodExpiry struct {
	value *ByDateTimeAuthenticatorGracePeriodExpiry
	isSet bool
}

func (v NullableByDateTimeAuthenticatorGracePeriodExpiry) Get() *ByDateTimeAuthenticatorGracePeriodExpiry {
	return v.value
}

func (v *NullableByDateTimeAuthenticatorGracePeriodExpiry) Set(val *ByDateTimeAuthenticatorGracePeriodExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableByDateTimeAuthenticatorGracePeriodExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableByDateTimeAuthenticatorGracePeriodExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableByDateTimeAuthenticatorGracePeriodExpiry(val *ByDateTimeAuthenticatorGracePeriodExpiry) *NullableByDateTimeAuthenticatorGracePeriodExpiry {
	return &NullableByDateTimeAuthenticatorGracePeriodExpiry{value: val, isSet: true}
}

func (v NullableByDateTimeAuthenticatorGracePeriodExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableByDateTimeAuthenticatorGracePeriodExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

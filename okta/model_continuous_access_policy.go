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
	"reflect"
	"strings"
)

// ContinuousAccessPolicy struct for ContinuousAccessPolicy
type ContinuousAccessPolicy struct {
	Policy
	// Policy conditions aren't supported for this policy type.
	Conditions NullableString `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessPolicy ContinuousAccessPolicy

// NewContinuousAccessPolicy instantiates a new ContinuousAccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessPolicy() *ContinuousAccessPolicy {
	this := ContinuousAccessPolicy{}
	return &this
}

// NewContinuousAccessPolicyWithDefaults instantiates a new ContinuousAccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessPolicyWithDefaults() *ContinuousAccessPolicy {
	this := ContinuousAccessPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContinuousAccessPolicy) GetConditions() string {
	if o == nil || o.Conditions.Get() == nil {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContinuousAccessPolicy) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *ContinuousAccessPolicy) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *ContinuousAccessPolicy) SetConditions(v string) {
	o.Conditions.Set(&v)
}
// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *ContinuousAccessPolicy) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *ContinuousAccessPolicy) UnsetConditions() {
	o.Conditions.Unset()
}

func (o ContinuousAccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type ContinuousAccessPolicyWithoutEmbeddedStruct struct {
		// Policy conditions aren't supported for this policy type.
		Conditions NullableString `json:"conditions,omitempty"`
	}

	varContinuousAccessPolicyWithoutEmbeddedStruct := ContinuousAccessPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicyWithoutEmbeddedStruct)
	if err == nil {
		varContinuousAccessPolicy := _ContinuousAccessPolicy{}
		varContinuousAccessPolicy.Conditions = varContinuousAccessPolicyWithoutEmbeddedStruct.Conditions
		*o = ContinuousAccessPolicy(varContinuousAccessPolicy)
	} else {
		return err
	}

	varContinuousAccessPolicy := _ContinuousAccessPolicy{}

	err = json.Unmarshal(bytes, &varContinuousAccessPolicy)
	if err == nil {
		o.Policy = varContinuousAccessPolicy.Policy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")

		// remove fields from embedded structs
		reflectPolicy := reflect.ValueOf(o.Policy)
		for i := 0; i < reflectPolicy.Type().NumField(); i++ {
			t := reflectPolicy.Type().Field(i)

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
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessPolicy struct {
	value *ContinuousAccessPolicy
	isSet bool
}

func (v NullableContinuousAccessPolicy) Get() *ContinuousAccessPolicy {
	return v.value
}

func (v *NullableContinuousAccessPolicy) Set(val *ContinuousAccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessPolicy(val *ContinuousAccessPolicy) *NullableContinuousAccessPolicy {
	return &NullableContinuousAccessPolicy{value: val, isSet: true}
}

func (v NullableContinuousAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


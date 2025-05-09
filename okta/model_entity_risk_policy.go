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

// EntityRiskPolicy struct for EntityRiskPolicy
type EntityRiskPolicy struct {
	Policy
	// Policy conditions aren't supported for this policy types.
	Conditions NullableString `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicy EntityRiskPolicy

// NewEntityRiskPolicy instantiates a new EntityRiskPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicy() *EntityRiskPolicy {
	this := EntityRiskPolicy{}
	return &this
}

// NewEntityRiskPolicyWithDefaults instantiates a new EntityRiskPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyWithDefaults() *EntityRiskPolicy {
	this := EntityRiskPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityRiskPolicy) GetConditions() string {
	if o == nil || o.Conditions.Get() == nil {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityRiskPolicy) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *EntityRiskPolicy) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *EntityRiskPolicy) SetConditions(v string) {
	o.Conditions.Set(&v)
}
// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *EntityRiskPolicy) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *EntityRiskPolicy) UnsetConditions() {
	o.Conditions.Unset()
}

func (o EntityRiskPolicy) MarshalJSON() ([]byte, error) {
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

func (o *EntityRiskPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type EntityRiskPolicyWithoutEmbeddedStruct struct {
		// Policy conditions aren't supported for this policy types.
		Conditions NullableString `json:"conditions,omitempty"`
	}

	varEntityRiskPolicyWithoutEmbeddedStruct := EntityRiskPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyWithoutEmbeddedStruct)
	if err == nil {
		varEntityRiskPolicy := _EntityRiskPolicy{}
		varEntityRiskPolicy.Conditions = varEntityRiskPolicyWithoutEmbeddedStruct.Conditions
		*o = EntityRiskPolicy(varEntityRiskPolicy)
	} else {
		return err
	}

	varEntityRiskPolicy := _EntityRiskPolicy{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicy)
	if err == nil {
		o.Policy = varEntityRiskPolicy.Policy
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

type NullableEntityRiskPolicy struct {
	value *EntityRiskPolicy
	isSet bool
}

func (v NullableEntityRiskPolicy) Get() *EntityRiskPolicy {
	return v.value
}

func (v *NullableEntityRiskPolicy) Set(val *EntityRiskPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicy(val *EntityRiskPolicy) *NullableEntityRiskPolicy {
	return &NullableEntityRiskPolicy{value: val, isSet: true}
}

func (v NullableEntityRiskPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


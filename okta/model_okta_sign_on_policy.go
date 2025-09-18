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
	"fmt"
	"reflect"
	"strings"
)

// checks if the OktaSignOnPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSignOnPolicy{}

// OktaSignOnPolicy struct for OktaSignOnPolicy
type OktaSignOnPolicy struct {
	Policy
	Conditions           *OktaSignOnPolicyConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicy OktaSignOnPolicy

// NewOktaSignOnPolicy instantiates a new OktaSignOnPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicy(name string, type_ string) *OktaSignOnPolicy {
	this := OktaSignOnPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
	return &this
}

// NewOktaSignOnPolicyWithDefaults instantiates a new OktaSignOnPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyWithDefaults() *OktaSignOnPolicy {
	this := OktaSignOnPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OktaSignOnPolicy) GetConditions() OktaSignOnPolicyConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret OktaSignOnPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicy) GetConditionsOk() (*OktaSignOnPolicyConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OktaSignOnPolicy) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given OktaSignOnPolicyConditions and assigns it to the Conditions field.
func (o *OktaSignOnPolicy) SetConditions(v OktaSignOnPolicyConditions) {
	o.Conditions = &v
}

func (o OktaSignOnPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSignOnPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSignOnPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type OktaSignOnPolicyWithoutEmbeddedStruct struct {
		Conditions *OktaSignOnPolicyConditions `json:"conditions,omitempty"`
	}

	varOktaSignOnPolicyWithoutEmbeddedStruct := OktaSignOnPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOktaSignOnPolicyWithoutEmbeddedStruct)
	if err == nil {
		varOktaSignOnPolicy := _OktaSignOnPolicy{}
		varOktaSignOnPolicy.Conditions = varOktaSignOnPolicyWithoutEmbeddedStruct.Conditions
		*o = OktaSignOnPolicy(varOktaSignOnPolicy)
	} else {
		return err
	}

	varOktaSignOnPolicy := _OktaSignOnPolicy{}

	err = json.Unmarshal(data, &varOktaSignOnPolicy)
	if err == nil {
		o.Policy = varOktaSignOnPolicy.Policy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableOktaSignOnPolicy struct {
	value *OktaSignOnPolicy
	isSet bool
}

func (v NullableOktaSignOnPolicy) Get() *OktaSignOnPolicy {
	return v.value
}

func (v *NullableOktaSignOnPolicy) Set(val *OktaSignOnPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicy(val *OktaSignOnPolicy) *NullableOktaSignOnPolicy {
	return &NullableOktaSignOnPolicy{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

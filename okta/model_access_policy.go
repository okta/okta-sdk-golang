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

// AccessPolicy struct for AccessPolicy
type AccessPolicy struct {
	Policy
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicy AccessPolicy

// NewAccessPolicy instantiates a new AccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicy() *AccessPolicy {
	this := AccessPolicy{}
	return &this
}

// NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyWithDefaults() *AccessPolicy {
	this := AccessPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AccessPolicy) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AccessPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *AccessPolicy) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

func (o AccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return []byte{}, errPolicy
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type AccessPolicyWithoutEmbeddedStruct struct {
		Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	}

	varAccessPolicyWithoutEmbeddedStruct := AccessPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAccessPolicyWithoutEmbeddedStruct)
	if err == nil {
		varAccessPolicy := _AccessPolicy{}
		varAccessPolicy.Conditions = varAccessPolicyWithoutEmbeddedStruct.Conditions
		*o = AccessPolicy(varAccessPolicy)
	} else {
		return err
	}

	varAccessPolicy := _AccessPolicy{}

	err = json.Unmarshal(bytes, &varAccessPolicy)
	if err == nil {
		o.Policy = varAccessPolicy.Policy
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

type NullableAccessPolicy struct {
	value *AccessPolicy
	isSet bool
}

func (v NullableAccessPolicy) Get() *AccessPolicy {
	return v.value
}

func (v *NullableAccessPolicy) Set(val *AccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicy(val *AccessPolicy) *NullableAccessPolicy {
	return &NullableAccessPolicy{value: val, isSet: true}
}

func (v NullableAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


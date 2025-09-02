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

// IdpDiscoveryPolicy struct for IdpDiscoveryPolicy
type IdpDiscoveryPolicy struct {
	Policy
	// Policy conditions aren't supported for this policy type.
	Conditions NullableString `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpDiscoveryPolicy IdpDiscoveryPolicy

// NewIdpDiscoveryPolicy instantiates a new IdpDiscoveryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpDiscoveryPolicy(name string, type_ string) *IdpDiscoveryPolicy {
	this := IdpDiscoveryPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
	return &this
}

// NewIdpDiscoveryPolicyWithDefaults instantiates a new IdpDiscoveryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpDiscoveryPolicyWithDefaults() *IdpDiscoveryPolicy {
	this := IdpDiscoveryPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpDiscoveryPolicy) GetConditions() string {
	if o == nil || o.Conditions.Get() == nil {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpDiscoveryPolicy) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *IdpDiscoveryPolicy) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *IdpDiscoveryPolicy) SetConditions(v string) {
	o.Conditions.Set(&v)
}
// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *IdpDiscoveryPolicy) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *IdpDiscoveryPolicy) UnsetConditions() {
	o.Conditions.Unset()
}

func (o IdpDiscoveryPolicy) MarshalJSON() ([]byte, error) {
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

func (o *IdpDiscoveryPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type IdpDiscoveryPolicyWithoutEmbeddedStruct struct {
		// Policy conditions aren't supported for this policy type.
		Conditions NullableString `json:"conditions,omitempty"`
	}

	varIdpDiscoveryPolicyWithoutEmbeddedStruct := IdpDiscoveryPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIdpDiscoveryPolicyWithoutEmbeddedStruct)
	if err == nil {
		varIdpDiscoveryPolicy := _IdpDiscoveryPolicy{}
		varIdpDiscoveryPolicy.Conditions = varIdpDiscoveryPolicyWithoutEmbeddedStruct.Conditions
		*o = IdpDiscoveryPolicy(varIdpDiscoveryPolicy)
	} else {
		return err
	}

	varIdpDiscoveryPolicy := _IdpDiscoveryPolicy{}

	err = json.Unmarshal(bytes, &varIdpDiscoveryPolicy)
	if err == nil {
		o.Policy = varIdpDiscoveryPolicy.Policy
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

type NullableIdpDiscoveryPolicy struct {
	value *IdpDiscoveryPolicy
	isSet bool
}

func (v NullableIdpDiscoveryPolicy) Get() *IdpDiscoveryPolicy {
	return v.value
}

func (v *NullableIdpDiscoveryPolicy) Set(val *IdpDiscoveryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpDiscoveryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpDiscoveryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpDiscoveryPolicy(val *IdpDiscoveryPolicy) *NullableIdpDiscoveryPolicy {
	return &NullableIdpDiscoveryPolicy{value: val, isSet: true}
}

func (v NullableIdpDiscoveryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpDiscoveryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


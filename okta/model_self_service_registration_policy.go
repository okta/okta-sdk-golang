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
	"fmt"
	"reflect"
	"strings"
)

// checks if the SelfServiceRegistrationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceRegistrationPolicy{}

// SelfServiceRegistrationPolicy <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Represents a Self-service Registration Policy. Use this policy to control self-service registration attempts based on network zones and email domain checks. Use the [Map a resource to a policy](https://developer.okta.com/docs/api/openapi/okta-management/management/tags/policy/other/mapresourcetopolicy) API to link this policy to a Profile Enrollment policy.
type SelfServiceRegistrationPolicy struct {
	Policy
	// Policy conditions aren't supported for this policy type.
	Conditions           NullableString `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceRegistrationPolicy SelfServiceRegistrationPolicy

// NewSelfServiceRegistrationPolicy instantiates a new SelfServiceRegistrationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationPolicy(name string, type_ string) *SelfServiceRegistrationPolicy {
	this := SelfServiceRegistrationPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
	return &this
}

// NewSelfServiceRegistrationPolicyWithDefaults instantiates a new SelfServiceRegistrationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationPolicyWithDefaults() *SelfServiceRegistrationPolicy {
	this := SelfServiceRegistrationPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelfServiceRegistrationPolicy) GetConditions() string {
	if o == nil || IsNil(o.Conditions.Get()) {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelfServiceRegistrationPolicy) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *SelfServiceRegistrationPolicy) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *SelfServiceRegistrationPolicy) SetConditions(v string) {
	o.Conditions.Set(&v)
}

// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *SelfServiceRegistrationPolicy) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *SelfServiceRegistrationPolicy) UnsetConditions() {
	o.Conditions.Unset()
}

func (o SelfServiceRegistrationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceRegistrationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceRegistrationPolicy) UnmarshalJSON(data []byte) (err error) {
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

	type SelfServiceRegistrationPolicyWithoutEmbeddedStruct struct {
		// Policy conditions aren't supported for this policy type.
		Conditions NullableString `json:"conditions,omitempty"`
	}

	varSelfServiceRegistrationPolicyWithoutEmbeddedStruct := SelfServiceRegistrationPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSelfServiceRegistrationPolicy := _SelfServiceRegistrationPolicy{}
		varSelfServiceRegistrationPolicy.Conditions = varSelfServiceRegistrationPolicyWithoutEmbeddedStruct.Conditions
		*o = SelfServiceRegistrationPolicy(varSelfServiceRegistrationPolicy)
	} else {
		return err
	}

	varSelfServiceRegistrationPolicy := _SelfServiceRegistrationPolicy{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicy)
	if err == nil {
		o.Policy = varSelfServiceRegistrationPolicy.Policy
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

type NullableSelfServiceRegistrationPolicy struct {
	value *SelfServiceRegistrationPolicy
	isSet bool
}

func (v NullableSelfServiceRegistrationPolicy) Get() *SelfServiceRegistrationPolicy {
	return v.value
}

func (v *NullableSelfServiceRegistrationPolicy) Set(val *SelfServiceRegistrationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationPolicy(val *SelfServiceRegistrationPolicy) *NullableSelfServiceRegistrationPolicy {
	return &NullableSelfServiceRegistrationPolicy{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

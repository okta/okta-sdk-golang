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
	"fmt"
)

// checks if the PolicyContextGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyContextGroups{}

// PolicyContextGroups An array of Group IDs for the simulate operation. Only user IDs or Group IDs are allowed, not both.
type PolicyContextGroups struct {
	Ids                  []string `json:"ids"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContextGroups PolicyContextGroups

// NewPolicyContextGroups instantiates a new PolicyContextGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContextGroups(ids []string) *PolicyContextGroups {
	this := PolicyContextGroups{}
	this.Ids = ids
	return &this
}

// NewPolicyContextGroupsWithDefaults instantiates a new PolicyContextGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextGroupsWithDefaults() *PolicyContextGroups {
	this := PolicyContextGroups{}
	return &this
}

// GetIds returns the Ids field value
func (o *PolicyContextGroups) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *PolicyContextGroups) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *PolicyContextGroups) SetIds(v []string) {
	o.Ids = v
}

func (o PolicyContextGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyContextGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyContextGroups) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
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

	varPolicyContextGroups := _PolicyContextGroups{}

	err = json.Unmarshal(data, &varPolicyContextGroups)

	if err != nil {
		return err
	}

	*o = PolicyContextGroups(varPolicyContextGroups)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyContextGroups struct {
	value *PolicyContextGroups
	isSet bool
}

func (v NullablePolicyContextGroups) Get() *PolicyContextGroups {
	return v.value
}

func (v *NullablePolicyContextGroups) Set(val *PolicyContextGroups) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContextGroups) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContextGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContextGroups(val *PolicyContextGroups) *NullablePolicyContextGroups {
	return &NullablePolicyContextGroups{value: val, isSet: true}
}

func (v NullablePolicyContextGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContextGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

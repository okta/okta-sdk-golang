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
)

// checks if the PolicyContextRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyContextRisk{}

// PolicyContextRisk The risk rule condition level
type PolicyContextRisk struct {
	Level                *string `json:"level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContextRisk PolicyContextRisk

// NewPolicyContextRisk instantiates a new PolicyContextRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContextRisk() *PolicyContextRisk {
	this := PolicyContextRisk{}
	return &this
}

// NewPolicyContextRiskWithDefaults instantiates a new PolicyContextRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextRiskWithDefaults() *PolicyContextRisk {
	this := PolicyContextRisk{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *PolicyContextRisk) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextRisk) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *PolicyContextRisk) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *PolicyContextRisk) SetLevel(v string) {
	o.Level = &v
}

func (o PolicyContextRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyContextRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyContextRisk) UnmarshalJSON(data []byte) (err error) {
	varPolicyContextRisk := _PolicyContextRisk{}

	err = json.Unmarshal(data, &varPolicyContextRisk)

	if err != nil {
		return err
	}

	*o = PolicyContextRisk(varPolicyContextRisk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyContextRisk struct {
	value *PolicyContextRisk
	isSet bool
}

func (v NullablePolicyContextRisk) Get() *PolicyContextRisk {
	return v.value
}

func (v *NullablePolicyContextRisk) Set(val *PolicyContextRisk) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContextRisk) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContextRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContextRisk(val *PolicyContextRisk) *NullablePolicyContextRisk {
	return &NullablePolicyContextRisk{value: val, isSet: true}
}

func (v NullablePolicyContextRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContextRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

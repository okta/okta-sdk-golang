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
)

// DevicePolicyRuleConditionAssurance struct for DevicePolicyRuleConditionAssurance
type DevicePolicyRuleConditionAssurance struct {
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePolicyRuleConditionAssurance DevicePolicyRuleConditionAssurance

// NewDevicePolicyRuleConditionAssurance instantiates a new DevicePolicyRuleConditionAssurance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePolicyRuleConditionAssurance() *DevicePolicyRuleConditionAssurance {
	this := DevicePolicyRuleConditionAssurance{}
	return &this
}

// NewDevicePolicyRuleConditionAssuranceWithDefaults instantiates a new DevicePolicyRuleConditionAssurance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePolicyRuleConditionAssuranceWithDefaults() *DevicePolicyRuleConditionAssurance {
	this := DevicePolicyRuleConditionAssurance{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *DevicePolicyRuleConditionAssurance) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePolicyRuleConditionAssurance) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *DevicePolicyRuleConditionAssurance) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *DevicePolicyRuleConditionAssurance) SetInclude(v []string) {
	o.Include = v
}

func (o DevicePolicyRuleConditionAssurance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DevicePolicyRuleConditionAssurance) UnmarshalJSON(bytes []byte) (err error) {
	varDevicePolicyRuleConditionAssurance := _DevicePolicyRuleConditionAssurance{}

	err = json.Unmarshal(bytes, &varDevicePolicyRuleConditionAssurance)
	if err == nil {
		*o = DevicePolicyRuleConditionAssurance(varDevicePolicyRuleConditionAssurance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDevicePolicyRuleConditionAssurance struct {
	value *DevicePolicyRuleConditionAssurance
	isSet bool
}

func (v NullableDevicePolicyRuleConditionAssurance) Get() *DevicePolicyRuleConditionAssurance {
	return v.value
}

func (v *NullableDevicePolicyRuleConditionAssurance) Set(val *DevicePolicyRuleConditionAssurance) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePolicyRuleConditionAssurance) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePolicyRuleConditionAssurance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePolicyRuleConditionAssurance(val *DevicePolicyRuleConditionAssurance) *NullableDevicePolicyRuleConditionAssurance {
	return &NullableDevicePolicyRuleConditionAssurance{value: val, isSet: true}
}

func (v NullableDevicePolicyRuleConditionAssurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePolicyRuleConditionAssurance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


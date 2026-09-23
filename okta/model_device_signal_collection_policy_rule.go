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
	"reflect"
	"strings"
)

// checks if the DeviceSignalCollectionPolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSignalCollectionPolicyRule{}

// DeviceSignalCollectionPolicyRule struct for DeviceSignalCollectionPolicyRule
type DeviceSignalCollectionPolicyRule struct {
	PolicyRule
	AdditionalProperties map[string]interface{}
}

type _DeviceSignalCollectionPolicyRule DeviceSignalCollectionPolicyRule

// NewDeviceSignalCollectionPolicyRule instantiates a new DeviceSignalCollectionPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPolicyRule() *DeviceSignalCollectionPolicyRule {
	this := DeviceSignalCollectionPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewDeviceSignalCollectionPolicyRuleWithDefaults instantiates a new DeviceSignalCollectionPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPolicyRuleWithDefaults() *DeviceSignalCollectionPolicyRule {
	this := DeviceSignalCollectionPolicyRule{}
	return &this
}

func (o DeviceSignalCollectionPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSignalCollectionPolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return map[string]interface{}{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return map[string]interface{}{}, errPolicyRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSignalCollectionPolicyRule) UnmarshalJSON(data []byte) (err error) {
	type DeviceSignalCollectionPolicyRuleWithoutEmbeddedStruct struct {
	}

	varDeviceSignalCollectionPolicyRuleWithoutEmbeddedStruct := DeviceSignalCollectionPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varDeviceSignalCollectionPolicyRule := _DeviceSignalCollectionPolicyRule{}
		*o = DeviceSignalCollectionPolicyRule(varDeviceSignalCollectionPolicyRule)
	} else {
		return err
	}

	varDeviceSignalCollectionPolicyRule := _DeviceSignalCollectionPolicyRule{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicyRule)
	if err == nil {
		o.PolicyRule = varDeviceSignalCollectionPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectPolicyRule := reflect.ValueOf(o.PolicyRule)
		for i := 0; i < reflectPolicyRule.Type().NumField(); i++ {
			t := reflectPolicyRule.Type().Field(i)

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

type NullableDeviceSignalCollectionPolicyRule struct {
	value *DeviceSignalCollectionPolicyRule
	isSet bool
}

func (v NullableDeviceSignalCollectionPolicyRule) Get() *DeviceSignalCollectionPolicyRule {
	return v.value
}

func (v *NullableDeviceSignalCollectionPolicyRule) Set(val *DeviceSignalCollectionPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPolicyRule(val *DeviceSignalCollectionPolicyRule) *NullableDeviceSignalCollectionPolicyRule {
	return &NullableDeviceSignalCollectionPolicyRule{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// BehaviorRuleAnomalousDevice struct for BehaviorRuleAnomalousDevice
type BehaviorRuleAnomalousDevice struct {
	BehaviorRule
	Settings *BehaviorRuleSettingsAnomalousDevice `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleAnomalousDevice BehaviorRuleAnomalousDevice

// NewBehaviorRuleAnomalousDevice instantiates a new BehaviorRuleAnomalousDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleAnomalousDevice(name string, type_ string) *BehaviorRuleAnomalousDevice {
	this := BehaviorRuleAnomalousDevice{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewBehaviorRuleAnomalousDeviceWithDefaults instantiates a new BehaviorRuleAnomalousDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleAnomalousDeviceWithDefaults() *BehaviorRuleAnomalousDevice {
	this := BehaviorRuleAnomalousDevice{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BehaviorRuleAnomalousDevice) GetSettings() BehaviorRuleSettingsAnomalousDevice {
	if o == nil || o.Settings == nil {
		var ret BehaviorRuleSettingsAnomalousDevice
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleAnomalousDevice) GetSettingsOk() (*BehaviorRuleSettingsAnomalousDevice, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BehaviorRuleAnomalousDevice) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given BehaviorRuleSettingsAnomalousDevice and assigns it to the Settings field.
func (o *BehaviorRuleAnomalousDevice) SetSettings(v BehaviorRuleSettingsAnomalousDevice) {
	o.Settings = &v
}

func (o BehaviorRuleAnomalousDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBehaviorRule, errBehaviorRule := json.Marshal(o.BehaviorRule)
	if errBehaviorRule != nil {
		return []byte{}, errBehaviorRule
	}
	errBehaviorRule = json.Unmarshal([]byte(serializedBehaviorRule), &toSerialize)
	if errBehaviorRule != nil {
		return []byte{}, errBehaviorRule
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleAnomalousDevice) UnmarshalJSON(bytes []byte) (err error) {
	type BehaviorRuleAnomalousDeviceWithoutEmbeddedStruct struct {
		Settings *BehaviorRuleSettingsAnomalousDevice `json:"settings,omitempty"`
	}

	varBehaviorRuleAnomalousDeviceWithoutEmbeddedStruct := BehaviorRuleAnomalousDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBehaviorRuleAnomalousDevice := _BehaviorRuleAnomalousDevice{}
		varBehaviorRuleAnomalousDevice.Settings = varBehaviorRuleAnomalousDeviceWithoutEmbeddedStruct.Settings
		*o = BehaviorRuleAnomalousDevice(varBehaviorRuleAnomalousDevice)
	} else {
		return err
	}

	varBehaviorRuleAnomalousDevice := _BehaviorRuleAnomalousDevice{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousDevice)
	if err == nil {
		o.BehaviorRule = varBehaviorRuleAnomalousDevice.BehaviorRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectBehaviorRule := reflect.ValueOf(o.BehaviorRule)
		for i := 0; i < reflectBehaviorRule.Type().NumField(); i++ {
			t := reflectBehaviorRule.Type().Field(i)

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

type NullableBehaviorRuleAnomalousDevice struct {
	value *BehaviorRuleAnomalousDevice
	isSet bool
}

func (v NullableBehaviorRuleAnomalousDevice) Get() *BehaviorRuleAnomalousDevice {
	return v.value
}

func (v *NullableBehaviorRuleAnomalousDevice) Set(val *BehaviorRuleAnomalousDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleAnomalousDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleAnomalousDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleAnomalousDevice(val *BehaviorRuleAnomalousDevice) *NullableBehaviorRuleAnomalousDevice {
	return &NullableBehaviorRuleAnomalousDevice{value: val, isSet: true}
}

func (v NullableBehaviorRuleAnomalousDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleAnomalousDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


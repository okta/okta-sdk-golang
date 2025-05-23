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

// BehaviorRuleAnomalousIP struct for BehaviorRuleAnomalousIP
type BehaviorRuleAnomalousIP struct {
	BehaviorRule
	Settings *BehaviorRuleSettingsAnomalousIP `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleAnomalousIP BehaviorRuleAnomalousIP

// NewBehaviorRuleAnomalousIP instantiates a new BehaviorRuleAnomalousIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleAnomalousIP(name string, type_ string) *BehaviorRuleAnomalousIP {
	this := BehaviorRuleAnomalousIP{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewBehaviorRuleAnomalousIPWithDefaults instantiates a new BehaviorRuleAnomalousIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleAnomalousIPWithDefaults() *BehaviorRuleAnomalousIP {
	this := BehaviorRuleAnomalousIP{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BehaviorRuleAnomalousIP) GetSettings() BehaviorRuleSettingsAnomalousIP {
	if o == nil || o.Settings == nil {
		var ret BehaviorRuleSettingsAnomalousIP
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleAnomalousIP) GetSettingsOk() (*BehaviorRuleSettingsAnomalousIP, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BehaviorRuleAnomalousIP) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given BehaviorRuleSettingsAnomalousIP and assigns it to the Settings field.
func (o *BehaviorRuleAnomalousIP) SetSettings(v BehaviorRuleSettingsAnomalousIP) {
	o.Settings = &v
}

func (o BehaviorRuleAnomalousIP) MarshalJSON() ([]byte, error) {
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

func (o *BehaviorRuleAnomalousIP) UnmarshalJSON(bytes []byte) (err error) {
	type BehaviorRuleAnomalousIPWithoutEmbeddedStruct struct {
		Settings *BehaviorRuleSettingsAnomalousIP `json:"settings,omitempty"`
	}

	varBehaviorRuleAnomalousIPWithoutEmbeddedStruct := BehaviorRuleAnomalousIPWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousIPWithoutEmbeddedStruct)
	if err == nil {
		varBehaviorRuleAnomalousIP := _BehaviorRuleAnomalousIP{}
		varBehaviorRuleAnomalousIP.Settings = varBehaviorRuleAnomalousIPWithoutEmbeddedStruct.Settings
		*o = BehaviorRuleAnomalousIP(varBehaviorRuleAnomalousIP)
	} else {
		return err
	}

	varBehaviorRuleAnomalousIP := _BehaviorRuleAnomalousIP{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousIP)
	if err == nil {
		o.BehaviorRule = varBehaviorRuleAnomalousIP.BehaviorRule
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

type NullableBehaviorRuleAnomalousIP struct {
	value *BehaviorRuleAnomalousIP
	isSet bool
}

func (v NullableBehaviorRuleAnomalousIP) Get() *BehaviorRuleAnomalousIP {
	return v.value
}

func (v *NullableBehaviorRuleAnomalousIP) Set(val *BehaviorRuleAnomalousIP) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleAnomalousIP) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleAnomalousIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleAnomalousIP(val *BehaviorRuleAnomalousIP) *NullableBehaviorRuleAnomalousIP {
	return &NullableBehaviorRuleAnomalousIP{value: val, isSet: true}
}

func (v NullableBehaviorRuleAnomalousIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleAnomalousIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


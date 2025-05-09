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

// BehaviorRuleAnomalousLocation struct for BehaviorRuleAnomalousLocation
type BehaviorRuleAnomalousLocation struct {
	BehaviorRule
	Settings *BehaviorRuleSettingsAnomalousLocation `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleAnomalousLocation BehaviorRuleAnomalousLocation

// NewBehaviorRuleAnomalousLocation instantiates a new BehaviorRuleAnomalousLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleAnomalousLocation(name string, type_ string) *BehaviorRuleAnomalousLocation {
	this := BehaviorRuleAnomalousLocation{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewBehaviorRuleAnomalousLocationWithDefaults instantiates a new BehaviorRuleAnomalousLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleAnomalousLocationWithDefaults() *BehaviorRuleAnomalousLocation {
	this := BehaviorRuleAnomalousLocation{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BehaviorRuleAnomalousLocation) GetSettings() BehaviorRuleSettingsAnomalousLocation {
	if o == nil || o.Settings == nil {
		var ret BehaviorRuleSettingsAnomalousLocation
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleAnomalousLocation) GetSettingsOk() (*BehaviorRuleSettingsAnomalousLocation, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BehaviorRuleAnomalousLocation) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given BehaviorRuleSettingsAnomalousLocation and assigns it to the Settings field.
func (o *BehaviorRuleAnomalousLocation) SetSettings(v BehaviorRuleSettingsAnomalousLocation) {
	o.Settings = &v
}

func (o BehaviorRuleAnomalousLocation) MarshalJSON() ([]byte, error) {
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

func (o *BehaviorRuleAnomalousLocation) UnmarshalJSON(bytes []byte) (err error) {
	type BehaviorRuleAnomalousLocationWithoutEmbeddedStruct struct {
		Settings *BehaviorRuleSettingsAnomalousLocation `json:"settings,omitempty"`
	}

	varBehaviorRuleAnomalousLocationWithoutEmbeddedStruct := BehaviorRuleAnomalousLocationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousLocationWithoutEmbeddedStruct)
	if err == nil {
		varBehaviorRuleAnomalousLocation := _BehaviorRuleAnomalousLocation{}
		varBehaviorRuleAnomalousLocation.Settings = varBehaviorRuleAnomalousLocationWithoutEmbeddedStruct.Settings
		*o = BehaviorRuleAnomalousLocation(varBehaviorRuleAnomalousLocation)
	} else {
		return err
	}

	varBehaviorRuleAnomalousLocation := _BehaviorRuleAnomalousLocation{}

	err = json.Unmarshal(bytes, &varBehaviorRuleAnomalousLocation)
	if err == nil {
		o.BehaviorRule = varBehaviorRuleAnomalousLocation.BehaviorRule
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

type NullableBehaviorRuleAnomalousLocation struct {
	value *BehaviorRuleAnomalousLocation
	isSet bool
}

func (v NullableBehaviorRuleAnomalousLocation) Get() *BehaviorRuleAnomalousLocation {
	return v.value
}

func (v *NullableBehaviorRuleAnomalousLocation) Set(val *BehaviorRuleAnomalousLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleAnomalousLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleAnomalousLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleAnomalousLocation(val *BehaviorRuleAnomalousLocation) *NullableBehaviorRuleAnomalousLocation {
	return &NullableBehaviorRuleAnomalousLocation{value: val, isSet: true}
}

func (v NullableBehaviorRuleAnomalousLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleAnomalousLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


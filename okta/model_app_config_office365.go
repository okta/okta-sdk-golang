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

// checks if the AppConfigOffice365 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConfigOffice365{}

// AppConfigOffice365 struct for AppConfigOffice365
type AppConfigOffice365 struct {
	AppConfig
	// Indicates that administrative roles in Office 365 can be assigned to a group
	AssignableToRole     bool `json:"assignableToRole"`
	AdditionalProperties map[string]interface{}
}

type _AppConfigOffice365 AppConfigOffice365

// NewAppConfigOffice365 instantiates a new AppConfigOffice365 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigOffice365(assignableToRole bool, type_ string) *AppConfigOffice365 {
	this := AppConfigOffice365{}
	this.Type = type_
	return &this
}

// NewAppConfigOffice365WithDefaults instantiates a new AppConfigOffice365 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigOffice365WithDefaults() *AppConfigOffice365 {
	this := AppConfigOffice365{}
	return &this
}

// GetAssignableToRole returns the AssignableToRole field value
func (o *AppConfigOffice365) GetAssignableToRole() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssignableToRole
}

// GetAssignableToRoleOk returns a tuple with the AssignableToRole field value
// and a boolean to check if the value has been set.
func (o *AppConfigOffice365) GetAssignableToRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignableToRole, true
}

// SetAssignableToRole sets field value
func (o *AppConfigOffice365) SetAssignableToRole(v bool) {
	o.AssignableToRole = v
}

func (o AppConfigOffice365) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConfigOffice365) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAppConfig, errAppConfig := json.Marshal(o.AppConfig)
	if errAppConfig != nil {
		return map[string]interface{}{}, errAppConfig
	}
	errAppConfig = json.Unmarshal([]byte(serializedAppConfig), &toSerialize)
	if errAppConfig != nil {
		return map[string]interface{}{}, errAppConfig
	}
	toSerialize["assignableToRole"] = o.AssignableToRole

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppConfigOffice365) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignableToRole",
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

	type AppConfigOffice365WithoutEmbeddedStruct struct {
		// Indicates that administrative roles in Office 365 can be assigned to a group
		AssignableToRole bool `json:"assignableToRole"`
	}

	varAppConfigOffice365WithoutEmbeddedStruct := AppConfigOffice365WithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAppConfigOffice365WithoutEmbeddedStruct)
	if err == nil {
		varAppConfigOffice365 := _AppConfigOffice365{}
		varAppConfigOffice365.AssignableToRole = varAppConfigOffice365WithoutEmbeddedStruct.AssignableToRole
		*o = AppConfigOffice365(varAppConfigOffice365)
	} else {
		return err
	}

	varAppConfigOffice365 := _AppConfigOffice365{}

	err = json.Unmarshal(data, &varAppConfigOffice365)
	if err == nil {
		o.AppConfig = varAppConfigOffice365.AppConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignableToRole")

		// remove fields from embedded structs
		reflectAppConfig := reflect.ValueOf(o.AppConfig)
		for i := 0; i < reflectAppConfig.Type().NumField(); i++ {
			t := reflectAppConfig.Type().Field(i)

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

type NullableAppConfigOffice365 struct {
	value *AppConfigOffice365
	isSet bool
}

func (v NullableAppConfigOffice365) Get() *AppConfigOffice365 {
	return v.value
}

func (v *NullableAppConfigOffice365) Set(val *AppConfigOffice365) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigOffice365) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigOffice365) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigOffice365(val *AppConfigOffice365) *NullableAppConfigOffice365 {
	return &NullableAppConfigOffice365{value: val, isSet: true}
}

func (v NullableAppConfigOffice365) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigOffice365) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

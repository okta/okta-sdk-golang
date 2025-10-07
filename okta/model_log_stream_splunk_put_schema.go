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

// checks if the LogStreamSplunkPutSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStreamSplunkPutSchema{}

// LogStreamSplunkPutSchema struct for LogStreamSplunkPutSchema
type LogStreamSplunkPutSchema struct {
	LogStreamPutSchema
	Settings             LogStreamSettingsSplunkPut `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSplunkPutSchema LogStreamSplunkPutSchema

// NewLogStreamSplunkPutSchema instantiates a new LogStreamSplunkPutSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSplunkPutSchema(settings LogStreamSettingsSplunkPut, name string, type_ string) *LogStreamSplunkPutSchema {
	this := LogStreamSplunkPutSchema{}
	this.Name = name
	this.Type = type_
	this.Settings = settings
	return &this
}

// NewLogStreamSplunkPutSchemaWithDefaults instantiates a new LogStreamSplunkPutSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSplunkPutSchemaWithDefaults() *LogStreamSplunkPutSchema {
	this := LogStreamSplunkPutSchema{}
	return &this
}

// GetSettings returns the Settings field value
func (o *LogStreamSplunkPutSchema) GetSettings() LogStreamSettingsSplunkPut {
	if o == nil {
		var ret LogStreamSettingsSplunkPut
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LogStreamSplunkPutSchema) GetSettingsOk() (*LogStreamSettingsSplunkPut, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LogStreamSplunkPutSchema) SetSettings(v LogStreamSettingsSplunkPut) {
	o.Settings = v
}

func (o LogStreamSplunkPutSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStreamSplunkPutSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLogStreamPutSchema, errLogStreamPutSchema := json.Marshal(o.LogStreamPutSchema)
	if errLogStreamPutSchema != nil {
		return map[string]interface{}{}, errLogStreamPutSchema
	}
	errLogStreamPutSchema = json.Unmarshal([]byte(serializedLogStreamPutSchema), &toSerialize)
	if errLogStreamPutSchema != nil {
		return map[string]interface{}{}, errLogStreamPutSchema
	}
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogStreamSplunkPutSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"settings",
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

	type LogStreamSplunkPutSchemaWithoutEmbeddedStruct struct {
		Settings LogStreamSettingsSplunkPut `json:"settings"`
	}

	varLogStreamSplunkPutSchemaWithoutEmbeddedStruct := LogStreamSplunkPutSchemaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varLogStreamSplunkPutSchemaWithoutEmbeddedStruct)
	if err == nil {
		varLogStreamSplunkPutSchema := _LogStreamSplunkPutSchema{}
		varLogStreamSplunkPutSchema.Settings = varLogStreamSplunkPutSchemaWithoutEmbeddedStruct.Settings
		*o = LogStreamSplunkPutSchema(varLogStreamSplunkPutSchema)
	} else {
		return err
	}

	varLogStreamSplunkPutSchema := _LogStreamSplunkPutSchema{}

	err = json.Unmarshal(data, &varLogStreamSplunkPutSchema)
	if err == nil {
		o.LogStreamPutSchema = varLogStreamSplunkPutSchema.LogStreamPutSchema
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectLogStreamPutSchema := reflect.ValueOf(o.LogStreamPutSchema)
		for i := 0; i < reflectLogStreamPutSchema.Type().NumField(); i++ {
			t := reflectLogStreamPutSchema.Type().Field(i)

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

type NullableLogStreamSplunkPutSchema struct {
	value *LogStreamSplunkPutSchema
	isSet bool
}

func (v NullableLogStreamSplunkPutSchema) Get() *LogStreamSplunkPutSchema {
	return v.value
}

func (v *NullableLogStreamSplunkPutSchema) Set(val *LogStreamSplunkPutSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSplunkPutSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSplunkPutSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSplunkPutSchema(val *LogStreamSplunkPutSchema) *NullableLogStreamSplunkPutSchema {
	return &NullableLogStreamSplunkPutSchema{value: val, isSet: true}
}

func (v NullableLogStreamSplunkPutSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSplunkPutSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

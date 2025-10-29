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

// checks if the LogStreamAwsPutSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStreamAwsPutSchema{}

// LogStreamAwsPutSchema struct for LogStreamAwsPutSchema
type LogStreamAwsPutSchema struct {
	LogStreamPutSchema
	Settings             LogStreamSettingsAws `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamAwsPutSchema LogStreamAwsPutSchema

// NewLogStreamAwsPutSchema instantiates a new LogStreamAwsPutSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamAwsPutSchema(settings LogStreamSettingsAws, name string, type_ string) *LogStreamAwsPutSchema {
	this := LogStreamAwsPutSchema{}
	this.Name = name
	this.Type = type_
	this.Settings = settings
	return &this
}

// NewLogStreamAwsPutSchemaWithDefaults instantiates a new LogStreamAwsPutSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamAwsPutSchemaWithDefaults() *LogStreamAwsPutSchema {
	this := LogStreamAwsPutSchema{}
	return &this
}

// GetSettings returns the Settings field value
func (o *LogStreamAwsPutSchema) GetSettings() LogStreamSettingsAws {
	if o == nil {
		var ret LogStreamSettingsAws
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LogStreamAwsPutSchema) GetSettingsOk() (*LogStreamSettingsAws, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LogStreamAwsPutSchema) SetSettings(v LogStreamSettingsAws) {
	o.Settings = v
}

func (o LogStreamAwsPutSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStreamAwsPutSchema) ToMap() (map[string]interface{}, error) {
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

func (o *LogStreamAwsPutSchema) UnmarshalJSON(data []byte) (err error) {
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

	type LogStreamAwsPutSchemaWithoutEmbeddedStruct struct {
		Settings LogStreamSettingsAws `json:"settings"`
	}

	varLogStreamAwsPutSchemaWithoutEmbeddedStruct := LogStreamAwsPutSchemaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varLogStreamAwsPutSchemaWithoutEmbeddedStruct)
	if err == nil {
		varLogStreamAwsPutSchema := _LogStreamAwsPutSchema{}
		varLogStreamAwsPutSchema.Settings = varLogStreamAwsPutSchemaWithoutEmbeddedStruct.Settings
		*o = LogStreamAwsPutSchema(varLogStreamAwsPutSchema)
	} else {
		return err
	}

	varLogStreamAwsPutSchema := _LogStreamAwsPutSchema{}

	err = json.Unmarshal(data, &varLogStreamAwsPutSchema)
	if err == nil {
		o.LogStreamPutSchema = varLogStreamAwsPutSchema.LogStreamPutSchema
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

type NullableLogStreamAwsPutSchema struct {
	value *LogStreamAwsPutSchema
	isSet bool
}

func (v NullableLogStreamAwsPutSchema) Get() *LogStreamAwsPutSchema {
	return v.value
}

func (v *NullableLogStreamAwsPutSchema) Set(val *LogStreamAwsPutSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamAwsPutSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamAwsPutSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamAwsPutSchema(val *LogStreamAwsPutSchema) *NullableLogStreamAwsPutSchema {
	return &NullableLogStreamAwsPutSchema{value: val, isSet: true}
}

func (v NullableLogStreamAwsPutSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamAwsPutSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

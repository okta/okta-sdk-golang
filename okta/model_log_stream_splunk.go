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
	"time"
)

// LogStreamSplunk struct for LogStreamSplunk
type LogStreamSplunk struct {
	LogStream
	Settings             LogStreamSettingsSplunk `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSplunk LogStreamSplunk

// NewLogStreamSplunk instantiates a new LogStreamSplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSplunk(settings LogStreamSettingsSplunk, created time.Time, id string, lastUpdated time.Time, name string, status string, type_ string, links LogStreamLinksSelfAndLifecycle) *LogStreamSplunk {
	this := LogStreamSplunk{}
	this.Created = created
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Name = name
	this.Status = status
	this.Type = type_
	this.Links = links
	this.Settings = settings
	return &this
}

// NewLogStreamSplunkWithDefaults instantiates a new LogStreamSplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSplunkWithDefaults() *LogStreamSplunk {
	this := LogStreamSplunk{}
	return &this
}

// GetSettings returns the Settings field value
func (o *LogStreamSplunk) GetSettings() LogStreamSettingsSplunk {
	if o == nil {
		var ret LogStreamSettingsSplunk
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LogStreamSplunk) GetSettingsOk() (*LogStreamSettingsSplunk, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LogStreamSplunk) SetSettings(v LogStreamSettingsSplunk) {
	o.Settings = v
}

func (o LogStreamSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedLogStream, errLogStream := json.Marshal(o.LogStream)
	if errLogStream != nil {
		return []byte{}, errLogStream
	}
	errLogStream = json.Unmarshal([]byte(serializedLogStream), &toSerialize)
	if errLogStream != nil {
		return []byte{}, errLogStream
	}
	if true {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamSplunk) UnmarshalJSON(bytes []byte) (err error) {
	type LogStreamSplunkWithoutEmbeddedStruct struct {
		Settings LogStreamSettingsSplunk `json:"settings"`
	}

	varLogStreamSplunkWithoutEmbeddedStruct := LogStreamSplunkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLogStreamSplunkWithoutEmbeddedStruct)
	if err == nil {
		varLogStreamSplunk := _LogStreamSplunk{}
		varLogStreamSplunk.Settings = varLogStreamSplunkWithoutEmbeddedStruct.Settings
		*o = LogStreamSplunk(varLogStreamSplunk)
	} else {
		return err
	}

	varLogStreamSplunk := _LogStreamSplunk{}

	err = json.Unmarshal(bytes, &varLogStreamSplunk)
	if err == nil {
		o.LogStream = varLogStreamSplunk.LogStream
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectLogStream := reflect.ValueOf(o.LogStream)
		for i := 0; i < reflectLogStream.Type().NumField(); i++ {
			t := reflectLogStream.Type().Field(i)

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

type NullableLogStreamSplunk struct {
	value *LogStreamSplunk
	isSet bool
}

func (v NullableLogStreamSplunk) Get() *LogStreamSplunk {
	return v.value
}

func (v *NullableLogStreamSplunk) Set(val *LogStreamSplunk) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSplunk) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSplunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSplunk(val *LogStreamSplunk) *NullableLogStreamSplunk {
	return &NullableLogStreamSplunk{value: val, isSet: true}
}

func (v NullableLogStreamSplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

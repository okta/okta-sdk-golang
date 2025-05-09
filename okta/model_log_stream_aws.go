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

// LogStreamAws struct for LogStreamAws
type LogStreamAws struct {
	LogStream
	Settings             LogStreamSettingsAws `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamAws LogStreamAws

// NewLogStreamAws instantiates a new LogStreamAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamAws(settings LogStreamSettingsAws, created time.Time, id string, lastUpdated time.Time, name string, status string, type_ string, links LogStreamLinksSelfAndLifecycle) *LogStreamAws {
	this := LogStreamAws{}
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

// NewLogStreamAwsWithDefaults instantiates a new LogStreamAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamAwsWithDefaults() *LogStreamAws {
	this := LogStreamAws{}
	return &this
}

// GetSettings returns the Settings field value
func (o *LogStreamAws) GetSettings() LogStreamSettingsAws {
	if o == nil {
		var ret LogStreamSettingsAws
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LogStreamAws) GetSettingsOk() (*LogStreamSettingsAws, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LogStreamAws) SetSettings(v LogStreamSettingsAws) {
	o.Settings = v
}

func (o LogStreamAws) MarshalJSON() ([]byte, error) {
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

func (o *LogStreamAws) UnmarshalJSON(bytes []byte) (err error) {
	type LogStreamAwsWithoutEmbeddedStruct struct {
		Settings LogStreamSettingsAws `json:"settings"`
	}

	varLogStreamAwsWithoutEmbeddedStruct := LogStreamAwsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLogStreamAwsWithoutEmbeddedStruct)
	if err == nil {
		varLogStreamAws := _LogStreamAws{}
		varLogStreamAws.Settings = varLogStreamAwsWithoutEmbeddedStruct.Settings
		*o = LogStreamAws(varLogStreamAws)
	} else {
		return err
	}

	varLogStreamAws := _LogStreamAws{}

	err = json.Unmarshal(bytes, &varLogStreamAws)
	if err == nil {
		o.LogStream = varLogStreamAws.LogStream
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

type NullableLogStreamAws struct {
	value *LogStreamAws
	isSet bool
}

func (v NullableLogStreamAws) Get() *LogStreamAws {
	return v.value
}

func (v *NullableLogStreamAws) Set(val *LogStreamAws) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamAws) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamAws(val *LogStreamAws) *NullableLogStreamAws {
	return &NullableLogStreamAws{value: val, isSet: true}
}

func (v NullableLogStreamAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

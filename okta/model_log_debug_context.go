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

// LogDebugContext struct for LogDebugContext
type LogDebugContext struct {
	DebugData map[string]interface{} `json:"debugData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogDebugContext LogDebugContext

// NewLogDebugContext instantiates a new LogDebugContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDebugContext() *LogDebugContext {
	this := LogDebugContext{}
	return &this
}

// NewLogDebugContextWithDefaults instantiates a new LogDebugContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDebugContextWithDefaults() *LogDebugContext {
	this := LogDebugContext{}
	return &this
}

// GetDebugData returns the DebugData field value if set, zero value otherwise.
func (o *LogDebugContext) GetDebugData() map[string]interface{} {
	if o == nil || o.DebugData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DebugData
}

// GetDebugDataOk returns a tuple with the DebugData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDebugContext) GetDebugDataOk() (map[string]interface{}, bool) {
	if o == nil || o.DebugData == nil {
		return nil, false
	}
	return o.DebugData, true
}

// HasDebugData returns a boolean if a field has been set.
func (o *LogDebugContext) HasDebugData() bool {
	if o != nil && o.DebugData != nil {
		return true
	}

	return false
}

// SetDebugData gets a reference to the given map[string]interface{} and assigns it to the DebugData field.
func (o *LogDebugContext) SetDebugData(v map[string]interface{}) {
	o.DebugData = v
}

func (o LogDebugContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DebugData != nil {
		toSerialize["debugData"] = o.DebugData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogDebugContext) UnmarshalJSON(bytes []byte) (err error) {
	varLogDebugContext := _LogDebugContext{}

	err = json.Unmarshal(bytes, &varLogDebugContext)
	if err == nil {
		*o = LogDebugContext(varLogDebugContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "debugData")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogDebugContext struct {
	value *LogDebugContext
	isSet bool
}

func (v NullableLogDebugContext) Get() *LogDebugContext {
	return v.value
}

func (v *NullableLogDebugContext) Set(val *LogDebugContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDebugContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDebugContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDebugContext(val *LogDebugContext) *NullableLogDebugContext {
	return &NullableLogDebugContext{value: val, isSet: true}
}

func (v NullableLogDebugContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDebugContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


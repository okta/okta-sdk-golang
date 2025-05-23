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

// InlineHookChannelHttp struct for InlineHookChannelHttp
type InlineHookChannelHttp struct {
	InlineHookChannel
	Config *InlineHookChannelConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelHttp InlineHookChannelHttp

// NewInlineHookChannelHttp instantiates a new InlineHookChannelHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelHttp() *InlineHookChannelHttp {
	this := InlineHookChannelHttp{}
	return &this
}

// NewInlineHookChannelHttpWithDefaults instantiates a new InlineHookChannelHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelHttpWithDefaults() *InlineHookChannelHttp {
	this := InlineHookChannelHttp{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InlineHookChannelHttp) GetConfig() InlineHookChannelConfig {
	if o == nil || o.Config == nil {
		var ret InlineHookChannelConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelHttp) GetConfigOk() (*InlineHookChannelConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InlineHookChannelHttp) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InlineHookChannelConfig and assigns it to the Config field.
func (o *InlineHookChannelHttp) SetConfig(v InlineHookChannelConfig) {
	o.Config = &v
}

func (o InlineHookChannelHttp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInlineHookChannel, errInlineHookChannel := json.Marshal(o.InlineHookChannel)
	if errInlineHookChannel != nil {
		return []byte{}, errInlineHookChannel
	}
	errInlineHookChannel = json.Unmarshal([]byte(serializedInlineHookChannel), &toSerialize)
	if errInlineHookChannel != nil {
		return []byte{}, errInlineHookChannel
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookChannelHttp) UnmarshalJSON(bytes []byte) (err error) {
	type InlineHookChannelHttpWithoutEmbeddedStruct struct {
		Config *InlineHookChannelConfig `json:"config,omitempty"`
	}

	varInlineHookChannelHttpWithoutEmbeddedStruct := InlineHookChannelHttpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInlineHookChannelHttpWithoutEmbeddedStruct)
	if err == nil {
		varInlineHookChannelHttp := _InlineHookChannelHttp{}
		varInlineHookChannelHttp.Config = varInlineHookChannelHttpWithoutEmbeddedStruct.Config
		*o = InlineHookChannelHttp(varInlineHookChannelHttp)
	} else {
		return err
	}

	varInlineHookChannelHttp := _InlineHookChannelHttp{}

	err = json.Unmarshal(bytes, &varInlineHookChannelHttp)
	if err == nil {
		o.InlineHookChannel = varInlineHookChannelHttp.InlineHookChannel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "config")

		// remove fields from embedded structs
		reflectInlineHookChannel := reflect.ValueOf(o.InlineHookChannel)
		for i := 0; i < reflectInlineHookChannel.Type().NumField(); i++ {
			t := reflectInlineHookChannel.Type().Field(i)

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

type NullableInlineHookChannelHttp struct {
	value *InlineHookChannelHttp
	isSet bool
}

func (v NullableInlineHookChannelHttp) Get() *InlineHookChannelHttp {
	return v.value
}

func (v *NullableInlineHookChannelHttp) Set(val *InlineHookChannelHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelHttp(val *InlineHookChannelHttp) *NullableInlineHookChannelHttp {
	return &NullableInlineHookChannelHttp{value: val, isSet: true}
}

func (v NullableInlineHookChannelHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// InlineHookChannelOAuth struct for InlineHookChannelOAuth
type InlineHookChannelOAuth struct {
	InlineHookChannel
	Config *InlineHookOAuthChannelConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelOAuth InlineHookChannelOAuth

// NewInlineHookChannelOAuth instantiates a new InlineHookChannelOAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelOAuth() *InlineHookChannelOAuth {
	this := InlineHookChannelOAuth{}
	return &this
}

// NewInlineHookChannelOAuthWithDefaults instantiates a new InlineHookChannelOAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelOAuthWithDefaults() *InlineHookChannelOAuth {
	this := InlineHookChannelOAuth{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InlineHookChannelOAuth) GetConfig() InlineHookOAuthChannelConfig {
	if o == nil || o.Config == nil {
		var ret InlineHookOAuthChannelConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelOAuth) GetConfigOk() (*InlineHookOAuthChannelConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InlineHookChannelOAuth) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InlineHookOAuthChannelConfig and assigns it to the Config field.
func (o *InlineHookChannelOAuth) SetConfig(v InlineHookOAuthChannelConfig) {
	o.Config = &v
}

func (o InlineHookChannelOAuth) MarshalJSON() ([]byte, error) {
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

func (o *InlineHookChannelOAuth) UnmarshalJSON(bytes []byte) (err error) {
	type InlineHookChannelOAuthWithoutEmbeddedStruct struct {
		Config *InlineHookOAuthChannelConfig `json:"config,omitempty"`
	}

	varInlineHookChannelOAuthWithoutEmbeddedStruct := InlineHookChannelOAuthWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInlineHookChannelOAuthWithoutEmbeddedStruct)
	if err == nil {
		varInlineHookChannelOAuth := _InlineHookChannelOAuth{}
		varInlineHookChannelOAuth.Config = varInlineHookChannelOAuthWithoutEmbeddedStruct.Config
		*o = InlineHookChannelOAuth(varInlineHookChannelOAuth)
	} else {
		return err
	}

	varInlineHookChannelOAuth := _InlineHookChannelOAuth{}

	err = json.Unmarshal(bytes, &varInlineHookChannelOAuth)
	if err == nil {
		o.InlineHookChannel = varInlineHookChannelOAuth.InlineHookChannel
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

type NullableInlineHookChannelOAuth struct {
	value *InlineHookChannelOAuth
	isSet bool
}

func (v NullableInlineHookChannelOAuth) Get() *InlineHookChannelOAuth {
	return v.value
}

func (v *NullableInlineHookChannelOAuth) Set(val *InlineHookChannelOAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelOAuth(val *InlineHookChannelOAuth) *NullableInlineHookChannelOAuth {
	return &NullableInlineHookChannelOAuth{value: val, isSet: true}
}

func (v NullableInlineHookChannelOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


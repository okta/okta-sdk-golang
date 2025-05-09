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

// InlineHookOAuthChannelConfig struct for InlineHookOAuthChannelConfig
type InlineHookOAuthChannelConfig struct {
	AuthType *string `json:"authType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookOAuthChannelConfig InlineHookOAuthChannelConfig

// NewInlineHookOAuthChannelConfig instantiates a new InlineHookOAuthChannelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookOAuthChannelConfig() *InlineHookOAuthChannelConfig {
	this := InlineHookOAuthChannelConfig{}
	return &this
}

// NewInlineHookOAuthChannelConfigWithDefaults instantiates a new InlineHookOAuthChannelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookOAuthChannelConfigWithDefaults() *InlineHookOAuthChannelConfig {
	this := InlineHookOAuthChannelConfig{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *InlineHookOAuthChannelConfig) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthChannelConfig) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *InlineHookOAuthChannelConfig) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *InlineHookOAuthChannelConfig) SetAuthType(v string) {
	o.AuthType = &v
}

func (o InlineHookOAuthChannelConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthType != nil {
		toSerialize["authType"] = o.AuthType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookOAuthChannelConfig) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookOAuthChannelConfig := _InlineHookOAuthChannelConfig{}

	err = json.Unmarshal(bytes, &varInlineHookOAuthChannelConfig)
	if err == nil {
		*o = InlineHookOAuthChannelConfig(varInlineHookOAuthChannelConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookOAuthChannelConfig struct {
	value *InlineHookOAuthChannelConfig
	isSet bool
}

func (v NullableInlineHookOAuthChannelConfig) Get() *InlineHookOAuthChannelConfig {
	return v.value
}

func (v *NullableInlineHookOAuthChannelConfig) Set(val *InlineHookOAuthChannelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookOAuthChannelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookOAuthChannelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookOAuthChannelConfig(val *InlineHookOAuthChannelConfig) *NullableInlineHookOAuthChannelConfig {
	return &NullableInlineHookOAuthChannelConfig{value: val, isSet: true}
}

func (v NullableInlineHookOAuthChannelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookOAuthChannelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


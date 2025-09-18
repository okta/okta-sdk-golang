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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the AppConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConfig{}

// AppConfig Additional app configuration for group push mappings. Currently only required for Active Directory.
type AppConfig struct {
	// The type of the app configuration
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AppConfig AppConfig

// NewAppConfig instantiates a new AppConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfig(type_ string) *AppConfig {
	this := AppConfig{}
	this.Type = type_
	return &this
}

// NewAppConfigWithDefaults instantiates a new AppConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigWithDefaults() *AppConfig {
	this := AppConfig{}
	return &this
}

// GetType returns the Type field value
func (o *AppConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppConfig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppConfig) SetType(v string) {
	o.Type = v
}

func (o AppConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAppConfig := _AppConfig{}

	err = json.Unmarshal(data, &varAppConfig)

	if err != nil {
		return err
	}

	*o = AppConfig(varAppConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppConfig struct {
	value *AppConfig
	isSet bool
}

func (v NullableAppConfig) Get() *AppConfig {
	return v.value
}

func (v *NullableAppConfig) Set(val *AppConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfig(val *AppConfig) *NullableAppConfig {
	return &NullableAppConfig{value: val, isSet: true}
}

func (v NullableAppConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

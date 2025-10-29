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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the FCMPushProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FCMPushProvider{}

// FCMPushProvider struct for FCMPushProvider
type FCMPushProvider struct {
	PushProvider
	Configuration        *FCMConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FCMPushProvider FCMPushProvider

// NewFCMPushProvider instantiates a new FCMPushProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMPushProvider() *FCMPushProvider {
	this := FCMPushProvider{}
	return &this
}

// NewFCMPushProviderWithDefaults instantiates a new FCMPushProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMPushProviderWithDefaults() *FCMPushProvider {
	this := FCMPushProvider{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *FCMPushProvider) GetConfiguration() FCMConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret FCMConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMPushProvider) GetConfigurationOk() (*FCMConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *FCMPushProvider) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given FCMConfiguration and assigns it to the Configuration field.
func (o *FCMPushProvider) SetConfiguration(v FCMConfiguration) {
	o.Configuration = &v
}

func (o FCMPushProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FCMPushProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPushProvider, errPushProvider := json.Marshal(o.PushProvider)
	if errPushProvider != nil {
		return map[string]interface{}{}, errPushProvider
	}
	errPushProvider = json.Unmarshal([]byte(serializedPushProvider), &toSerialize)
	if errPushProvider != nil {
		return map[string]interface{}{}, errPushProvider
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FCMPushProvider) UnmarshalJSON(data []byte) (err error) {
	type FCMPushProviderWithoutEmbeddedStruct struct {
		Configuration *FCMConfiguration `json:"configuration,omitempty"`
	}

	varFCMPushProviderWithoutEmbeddedStruct := FCMPushProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFCMPushProviderWithoutEmbeddedStruct)
	if err == nil {
		varFCMPushProvider := _FCMPushProvider{}
		varFCMPushProvider.Configuration = varFCMPushProviderWithoutEmbeddedStruct.Configuration
		*o = FCMPushProvider(varFCMPushProvider)
	} else {
		return err
	}

	varFCMPushProvider := _FCMPushProvider{}

	err = json.Unmarshal(data, &varFCMPushProvider)
	if err == nil {
		o.PushProvider = varFCMPushProvider.PushProvider
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")

		// remove fields from embedded structs
		reflectPushProvider := reflect.ValueOf(o.PushProvider)
		for i := 0; i < reflectPushProvider.Type().NumField(); i++ {
			t := reflectPushProvider.Type().Field(i)

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

type NullableFCMPushProvider struct {
	value *FCMPushProvider
	isSet bool
}

func (v NullableFCMPushProvider) Get() *FCMPushProvider {
	return v.value
}

func (v *NullableFCMPushProvider) Set(val *FCMPushProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMPushProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMPushProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMPushProvider(val *FCMPushProvider) *NullableFCMPushProvider {
	return &NullableFCMPushProvider{value: val, isSet: true}
}

func (v NullableFCMPushProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMPushProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

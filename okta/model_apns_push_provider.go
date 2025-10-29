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

// checks if the APNSPushProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APNSPushProvider{}

// APNSPushProvider struct for APNSPushProvider
type APNSPushProvider struct {
	PushProvider
	Configuration        *APNSConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APNSPushProvider APNSPushProvider

// NewAPNSPushProvider instantiates a new APNSPushProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPNSPushProvider() *APNSPushProvider {
	this := APNSPushProvider{}
	return &this
}

// NewAPNSPushProviderWithDefaults instantiates a new APNSPushProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPNSPushProviderWithDefaults() *APNSPushProvider {
	this := APNSPushProvider{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *APNSPushProvider) GetConfiguration() APNSConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret APNSConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APNSPushProvider) GetConfigurationOk() (*APNSConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *APNSPushProvider) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given APNSConfiguration and assigns it to the Configuration field.
func (o *APNSPushProvider) SetConfiguration(v APNSConfiguration) {
	o.Configuration = &v
}

func (o APNSPushProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APNSPushProvider) ToMap() (map[string]interface{}, error) {
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

func (o *APNSPushProvider) UnmarshalJSON(data []byte) (err error) {
	type APNSPushProviderWithoutEmbeddedStruct struct {
		Configuration *APNSConfiguration `json:"configuration,omitempty"`
	}

	varAPNSPushProviderWithoutEmbeddedStruct := APNSPushProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAPNSPushProviderWithoutEmbeddedStruct)
	if err == nil {
		varAPNSPushProvider := _APNSPushProvider{}
		varAPNSPushProvider.Configuration = varAPNSPushProviderWithoutEmbeddedStruct.Configuration
		*o = APNSPushProvider(varAPNSPushProvider)
	} else {
		return err
	}

	varAPNSPushProvider := _APNSPushProvider{}

	err = json.Unmarshal(data, &varAPNSPushProvider)
	if err == nil {
		o.PushProvider = varAPNSPushProvider.PushProvider
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

type NullableAPNSPushProvider struct {
	value *APNSPushProvider
	isSet bool
}

func (v NullableAPNSPushProvider) Get() *APNSPushProvider {
	return v.value
}

func (v *NullableAPNSPushProvider) Set(val *APNSPushProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAPNSPushProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAPNSPushProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPNSPushProvider(val *APNSPushProvider) *NullableAPNSPushProvider {
	return &NullableAPNSPushProvider{value: val, isSet: true}
}

func (v NullableAPNSPushProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPNSPushProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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
)

// checks if the CustomAuthSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAuthSettings{}

// CustomAuthSettings Set of AIPs used for authType `CUSTOM`
type CustomAuthSettings struct {
	AppInstanceProperties []AppInstanceProperty `json:"appInstanceProperties,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _CustomAuthSettings CustomAuthSettings

// NewCustomAuthSettings instantiates a new CustomAuthSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAuthSettings() *CustomAuthSettings {
	this := CustomAuthSettings{}
	return &this
}

// NewCustomAuthSettingsWithDefaults instantiates a new CustomAuthSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAuthSettingsWithDefaults() *CustomAuthSettings {
	this := CustomAuthSettings{}
	return &this
}

// GetAppInstanceProperties returns the AppInstanceProperties field value if set, zero value otherwise.
func (o *CustomAuthSettings) GetAppInstanceProperties() []AppInstanceProperty {
	if o == nil || IsNil(o.AppInstanceProperties) {
		var ret []AppInstanceProperty
		return ret
	}
	return o.AppInstanceProperties
}

// GetAppInstancePropertiesOk returns a tuple with the AppInstanceProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAuthSettings) GetAppInstancePropertiesOk() ([]AppInstanceProperty, bool) {
	if o == nil || IsNil(o.AppInstanceProperties) {
		return nil, false
	}
	return o.AppInstanceProperties, true
}

// HasAppInstanceProperties returns a boolean if a field has been set.
func (o *CustomAuthSettings) HasAppInstanceProperties() bool {
	if o != nil && !IsNil(o.AppInstanceProperties) {
		return true
	}

	return false
}

// SetAppInstanceProperties gets a reference to the given []AppInstanceProperty and assigns it to the AppInstanceProperties field.
func (o *CustomAuthSettings) SetAppInstanceProperties(v []AppInstanceProperty) {
	o.AppInstanceProperties = v
}

func (o CustomAuthSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAuthSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppInstanceProperties) {
		toSerialize["appInstanceProperties"] = o.AppInstanceProperties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAuthSettings) UnmarshalJSON(data []byte) (err error) {
	varCustomAuthSettings := _CustomAuthSettings{}

	err = json.Unmarshal(data, &varCustomAuthSettings)

	if err != nil {
		return err
	}

	*o = CustomAuthSettings(varCustomAuthSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceProperties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAuthSettings struct {
	value *CustomAuthSettings
	isSet bool
}

func (v NullableCustomAuthSettings) Get() *CustomAuthSettings {
	return v.value
}

func (v *NullableCustomAuthSettings) Set(val *CustomAuthSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAuthSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAuthSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAuthSettings(val *CustomAuthSettings) *NullableCustomAuthSettings {
	return &NullableCustomAuthSettings{value: val, isSet: true}
}

func (v NullableCustomAuthSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAuthSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

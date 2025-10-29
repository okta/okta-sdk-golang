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

// checks if the TenantSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantSettings{}

// TenantSettings struct for TenantSettings
type TenantSettings struct {
	AppInstanceProperties []AppInstanceProperty `json:"appInstanceProperties,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _TenantSettings TenantSettings

// NewTenantSettings instantiates a new TenantSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantSettings() *TenantSettings {
	this := TenantSettings{}
	return &this
}

// NewTenantSettingsWithDefaults instantiates a new TenantSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantSettingsWithDefaults() *TenantSettings {
	this := TenantSettings{}
	return &this
}

// GetAppInstanceProperties returns the AppInstanceProperties field value if set, zero value otherwise.
func (o *TenantSettings) GetAppInstanceProperties() []AppInstanceProperty {
	if o == nil || IsNil(o.AppInstanceProperties) {
		var ret []AppInstanceProperty
		return ret
	}
	return o.AppInstanceProperties
}

// GetAppInstancePropertiesOk returns a tuple with the AppInstanceProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantSettings) GetAppInstancePropertiesOk() ([]AppInstanceProperty, bool) {
	if o == nil || IsNil(o.AppInstanceProperties) {
		return nil, false
	}
	return o.AppInstanceProperties, true
}

// HasAppInstanceProperties returns a boolean if a field has been set.
func (o *TenantSettings) HasAppInstanceProperties() bool {
	if o != nil && !IsNil(o.AppInstanceProperties) {
		return true
	}

	return false
}

// SetAppInstanceProperties gets a reference to the given []AppInstanceProperty and assigns it to the AppInstanceProperties field.
func (o *TenantSettings) SetAppInstanceProperties(v []AppInstanceProperty) {
	o.AppInstanceProperties = v
}

func (o TenantSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppInstanceProperties) {
		toSerialize["appInstanceProperties"] = o.AppInstanceProperties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantSettings) UnmarshalJSON(data []byte) (err error) {
	varTenantSettings := _TenantSettings{}

	err = json.Unmarshal(data, &varTenantSettings)

	if err != nil {
		return err
	}

	*o = TenantSettings(varTenantSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceProperties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantSettings struct {
	value *TenantSettings
	isSet bool
}

func (v NullableTenantSettings) Get() *TenantSettings {
	return v.value
}

func (v *NullableTenantSettings) Set(val *TenantSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantSettings(val *TenantSettings) *NullableTenantSettings {
	return &NullableTenantSettings{value: val, isSet: true}
}

func (v NullableTenantSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

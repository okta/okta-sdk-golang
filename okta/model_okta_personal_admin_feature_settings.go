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

// checks if the OktaPersonalAdminFeatureSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaPersonalAdminFeatureSettings{}

// OktaPersonalAdminFeatureSettings Defines a list of Okta Personal settings that can be enabled or disabled for the org
type OktaPersonalAdminFeatureSettings struct {
	// Allow entry points for an Okta Personal account in a Workforce org
	EnableEnduserEntryPoints *bool `json:"enableEnduserEntryPoints,omitempty"`
	// Allow users to migrate apps from a Workforce account to an Okta Personal account
	EnableExportApps     *bool `json:"enableExportApps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaPersonalAdminFeatureSettings OktaPersonalAdminFeatureSettings

// NewOktaPersonalAdminFeatureSettings instantiates a new OktaPersonalAdminFeatureSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaPersonalAdminFeatureSettings() *OktaPersonalAdminFeatureSettings {
	this := OktaPersonalAdminFeatureSettings{}
	return &this
}

// NewOktaPersonalAdminFeatureSettingsWithDefaults instantiates a new OktaPersonalAdminFeatureSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaPersonalAdminFeatureSettingsWithDefaults() *OktaPersonalAdminFeatureSettings {
	this := OktaPersonalAdminFeatureSettings{}
	return &this
}

// GetEnableEnduserEntryPoints returns the EnableEnduserEntryPoints field value if set, zero value otherwise.
func (o *OktaPersonalAdminFeatureSettings) GetEnableEnduserEntryPoints() bool {
	if o == nil || IsNil(o.EnableEnduserEntryPoints) {
		var ret bool
		return ret
	}
	return *o.EnableEnduserEntryPoints
}

// GetEnableEnduserEntryPointsOk returns a tuple with the EnableEnduserEntryPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaPersonalAdminFeatureSettings) GetEnableEnduserEntryPointsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEnduserEntryPoints) {
		return nil, false
	}
	return o.EnableEnduserEntryPoints, true
}

// HasEnableEnduserEntryPoints returns a boolean if a field has been set.
func (o *OktaPersonalAdminFeatureSettings) HasEnableEnduserEntryPoints() bool {
	if o != nil && !IsNil(o.EnableEnduserEntryPoints) {
		return true
	}

	return false
}

// SetEnableEnduserEntryPoints gets a reference to the given bool and assigns it to the EnableEnduserEntryPoints field.
func (o *OktaPersonalAdminFeatureSettings) SetEnableEnduserEntryPoints(v bool) {
	o.EnableEnduserEntryPoints = &v
}

// GetEnableExportApps returns the EnableExportApps field value if set, zero value otherwise.
func (o *OktaPersonalAdminFeatureSettings) GetEnableExportApps() bool {
	if o == nil || IsNil(o.EnableExportApps) {
		var ret bool
		return ret
	}
	return *o.EnableExportApps
}

// GetEnableExportAppsOk returns a tuple with the EnableExportApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaPersonalAdminFeatureSettings) GetEnableExportAppsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExportApps) {
		return nil, false
	}
	return o.EnableExportApps, true
}

// HasEnableExportApps returns a boolean if a field has been set.
func (o *OktaPersonalAdminFeatureSettings) HasEnableExportApps() bool {
	if o != nil && !IsNil(o.EnableExportApps) {
		return true
	}

	return false
}

// SetEnableExportApps gets a reference to the given bool and assigns it to the EnableExportApps field.
func (o *OktaPersonalAdminFeatureSettings) SetEnableExportApps(v bool) {
	o.EnableExportApps = &v
}

func (o OktaPersonalAdminFeatureSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaPersonalAdminFeatureSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableEnduserEntryPoints) {
		toSerialize["enableEnduserEntryPoints"] = o.EnableEnduserEntryPoints
	}
	if !IsNil(o.EnableExportApps) {
		toSerialize["enableExportApps"] = o.EnableExportApps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaPersonalAdminFeatureSettings) UnmarshalJSON(data []byte) (err error) {
	varOktaPersonalAdminFeatureSettings := _OktaPersonalAdminFeatureSettings{}

	err = json.Unmarshal(data, &varOktaPersonalAdminFeatureSettings)

	if err != nil {
		return err
	}

	*o = OktaPersonalAdminFeatureSettings(varOktaPersonalAdminFeatureSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enableEnduserEntryPoints")
		delete(additionalProperties, "enableExportApps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaPersonalAdminFeatureSettings struct {
	value *OktaPersonalAdminFeatureSettings
	isSet bool
}

func (v NullableOktaPersonalAdminFeatureSettings) Get() *OktaPersonalAdminFeatureSettings {
	return v.value
}

func (v *NullableOktaPersonalAdminFeatureSettings) Set(val *OktaPersonalAdminFeatureSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaPersonalAdminFeatureSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaPersonalAdminFeatureSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaPersonalAdminFeatureSettings(val *OktaPersonalAdminFeatureSettings) *NullableOktaPersonalAdminFeatureSettings {
	return &NullableOktaPersonalAdminFeatureSettings{value: val, isSet: true}
}

func (v NullableOktaPersonalAdminFeatureSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaPersonalAdminFeatureSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

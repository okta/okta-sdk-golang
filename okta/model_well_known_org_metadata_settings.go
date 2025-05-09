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

// WellKnownOrgMetadataSettings struct for WellKnownOrgMetadataSettings
type WellKnownOrgMetadataSettings struct {
	AnalyticsCollectionEnabled *bool `json:"analyticsCollectionEnabled,omitempty"`
	BugReportingEnabled *bool `json:"bugReportingEnabled,omitempty"`
	// Whether the legacy Okta Mobile application is enabled for the org
	OmEnabled *bool `json:"omEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownOrgMetadataSettings WellKnownOrgMetadataSettings

// NewWellKnownOrgMetadataSettings instantiates a new WellKnownOrgMetadataSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownOrgMetadataSettings() *WellKnownOrgMetadataSettings {
	this := WellKnownOrgMetadataSettings{}
	return &this
}

// NewWellKnownOrgMetadataSettingsWithDefaults instantiates a new WellKnownOrgMetadataSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownOrgMetadataSettingsWithDefaults() *WellKnownOrgMetadataSettings {
	this := WellKnownOrgMetadataSettings{}
	return &this
}

// GetAnalyticsCollectionEnabled returns the AnalyticsCollectionEnabled field value if set, zero value otherwise.
func (o *WellKnownOrgMetadataSettings) GetAnalyticsCollectionEnabled() bool {
	if o == nil || o.AnalyticsCollectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AnalyticsCollectionEnabled
}

// GetAnalyticsCollectionEnabledOk returns a tuple with the AnalyticsCollectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataSettings) GetAnalyticsCollectionEnabledOk() (*bool, bool) {
	if o == nil || o.AnalyticsCollectionEnabled == nil {
		return nil, false
	}
	return o.AnalyticsCollectionEnabled, true
}

// HasAnalyticsCollectionEnabled returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataSettings) HasAnalyticsCollectionEnabled() bool {
	if o != nil && o.AnalyticsCollectionEnabled != nil {
		return true
	}

	return false
}

// SetAnalyticsCollectionEnabled gets a reference to the given bool and assigns it to the AnalyticsCollectionEnabled field.
func (o *WellKnownOrgMetadataSettings) SetAnalyticsCollectionEnabled(v bool) {
	o.AnalyticsCollectionEnabled = &v
}

// GetBugReportingEnabled returns the BugReportingEnabled field value if set, zero value otherwise.
func (o *WellKnownOrgMetadataSettings) GetBugReportingEnabled() bool {
	if o == nil || o.BugReportingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BugReportingEnabled
}

// GetBugReportingEnabledOk returns a tuple with the BugReportingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataSettings) GetBugReportingEnabledOk() (*bool, bool) {
	if o == nil || o.BugReportingEnabled == nil {
		return nil, false
	}
	return o.BugReportingEnabled, true
}

// HasBugReportingEnabled returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataSettings) HasBugReportingEnabled() bool {
	if o != nil && o.BugReportingEnabled != nil {
		return true
	}

	return false
}

// SetBugReportingEnabled gets a reference to the given bool and assigns it to the BugReportingEnabled field.
func (o *WellKnownOrgMetadataSettings) SetBugReportingEnabled(v bool) {
	o.BugReportingEnabled = &v
}

// GetOmEnabled returns the OmEnabled field value if set, zero value otherwise.
func (o *WellKnownOrgMetadataSettings) GetOmEnabled() bool {
	if o == nil || o.OmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OmEnabled
}

// GetOmEnabledOk returns a tuple with the OmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataSettings) GetOmEnabledOk() (*bool, bool) {
	if o == nil || o.OmEnabled == nil {
		return nil, false
	}
	return o.OmEnabled, true
}

// HasOmEnabled returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataSettings) HasOmEnabled() bool {
	if o != nil && o.OmEnabled != nil {
		return true
	}

	return false
}

// SetOmEnabled gets a reference to the given bool and assigns it to the OmEnabled field.
func (o *WellKnownOrgMetadataSettings) SetOmEnabled(v bool) {
	o.OmEnabled = &v
}

func (o WellKnownOrgMetadataSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyticsCollectionEnabled != nil {
		toSerialize["analyticsCollectionEnabled"] = o.AnalyticsCollectionEnabled
	}
	if o.BugReportingEnabled != nil {
		toSerialize["bugReportingEnabled"] = o.BugReportingEnabled
	}
	if o.OmEnabled != nil {
		toSerialize["omEnabled"] = o.OmEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownOrgMetadataSettings) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownOrgMetadataSettings := _WellKnownOrgMetadataSettings{}

	err = json.Unmarshal(bytes, &varWellKnownOrgMetadataSettings)
	if err == nil {
		*o = WellKnownOrgMetadataSettings(varWellKnownOrgMetadataSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "analyticsCollectionEnabled")
		delete(additionalProperties, "bugReportingEnabled")
		delete(additionalProperties, "omEnabled")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownOrgMetadataSettings struct {
	value *WellKnownOrgMetadataSettings
	isSet bool
}

func (v NullableWellKnownOrgMetadataSettings) Get() *WellKnownOrgMetadataSettings {
	return v.value
}

func (v *NullableWellKnownOrgMetadataSettings) Set(val *WellKnownOrgMetadataSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownOrgMetadataSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownOrgMetadataSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownOrgMetadataSettings(val *WellKnownOrgMetadataSettings) *NullableWellKnownOrgMetadataSettings {
	return &NullableWellKnownOrgMetadataSettings{value: val, isSet: true}
}

func (v NullableWellKnownOrgMetadataSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownOrgMetadataSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


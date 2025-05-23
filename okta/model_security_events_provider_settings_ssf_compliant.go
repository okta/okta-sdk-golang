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

// SecurityEventsProviderSettingsSSFCompliant Security Events Provider with well-known URL setting
type SecurityEventsProviderSettingsSSFCompliant struct {
	// The published well-known URL of the Security Events Provider (the SSF transmitter)
	WellKnownUrl string `json:"well_known_url"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventsProviderSettingsSSFCompliant SecurityEventsProviderSettingsSSFCompliant

// NewSecurityEventsProviderSettingsSSFCompliant instantiates a new SecurityEventsProviderSettingsSSFCompliant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventsProviderSettingsSSFCompliant(wellKnownUrl string) *SecurityEventsProviderSettingsSSFCompliant {
	this := SecurityEventsProviderSettingsSSFCompliant{}
	this.WellKnownUrl = wellKnownUrl
	return &this
}

// NewSecurityEventsProviderSettingsSSFCompliantWithDefaults instantiates a new SecurityEventsProviderSettingsSSFCompliant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventsProviderSettingsSSFCompliantWithDefaults() *SecurityEventsProviderSettingsSSFCompliant {
	this := SecurityEventsProviderSettingsSSFCompliant{}
	return &this
}

// GetWellKnownUrl returns the WellKnownUrl field value
func (o *SecurityEventsProviderSettingsSSFCompliant) GetWellKnownUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WellKnownUrl
}

// GetWellKnownUrlOk returns a tuple with the WellKnownUrl field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderSettingsSSFCompliant) GetWellKnownUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WellKnownUrl, true
}

// SetWellKnownUrl sets field value
func (o *SecurityEventsProviderSettingsSSFCompliant) SetWellKnownUrl(v string) {
	o.WellKnownUrl = v
}

func (o SecurityEventsProviderSettingsSSFCompliant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["well_known_url"] = o.WellKnownUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventsProviderSettingsSSFCompliant) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventsProviderSettingsSSFCompliant := _SecurityEventsProviderSettingsSSFCompliant{}

	err = json.Unmarshal(bytes, &varSecurityEventsProviderSettingsSSFCompliant)
	if err == nil {
		*o = SecurityEventsProviderSettingsSSFCompliant(varSecurityEventsProviderSettingsSSFCompliant)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "well_known_url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventsProviderSettingsSSFCompliant struct {
	value *SecurityEventsProviderSettingsSSFCompliant
	isSet bool
}

func (v NullableSecurityEventsProviderSettingsSSFCompliant) Get() *SecurityEventsProviderSettingsSSFCompliant {
	return v.value
}

func (v *NullableSecurityEventsProviderSettingsSSFCompliant) Set(val *SecurityEventsProviderSettingsSSFCompliant) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderSettingsSSFCompliant) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderSettingsSSFCompliant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderSettingsSSFCompliant(val *SecurityEventsProviderSettingsSSFCompliant) *NullableSecurityEventsProviderSettingsSSFCompliant {
	return &NullableSecurityEventsProviderSettingsSSFCompliant{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderSettingsSSFCompliant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderSettingsSSFCompliant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


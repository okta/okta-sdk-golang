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
	"fmt"
)

// checks if the SecurityEventsProviderSettingsSSFCompliant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityEventsProviderSettingsSSFCompliant{}

// SecurityEventsProviderSettingsSSFCompliant Security Events Provider with well-known URL setting
type SecurityEventsProviderSettingsSSFCompliant struct {
	// The published well-known URL of the Security Events Provider (the SSF transmitter)
	WellKnownUrl         string `json:"well_known_url"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityEventsProviderSettingsSSFCompliant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["well_known_url"] = o.WellKnownUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityEventsProviderSettingsSSFCompliant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"well_known_url",
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

	varSecurityEventsProviderSettingsSSFCompliant := _SecurityEventsProviderSettingsSSFCompliant{}

	err = json.Unmarshal(data, &varSecurityEventsProviderSettingsSSFCompliant)

	if err != nil {
		return err
	}

	*o = SecurityEventsProviderSettingsSSFCompliant(varSecurityEventsProviderSettingsSSFCompliant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "well_known_url")
		o.AdditionalProperties = additionalProperties
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

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
	"fmt"
)


//model_oneof.mustache
// SecurityEventsProviderRequestSettings - Information about the Security Events Provider for signal ingestion
type SecurityEventsProviderRequestSettings struct {
	SecurityEventsProviderSettingsNonSSFCompliant *SecurityEventsProviderSettingsNonSSFCompliant
	SecurityEventsProviderSettingsSSFCompliant *SecurityEventsProviderSettingsSSFCompliant
}

// SecurityEventsProviderSettingsNonSSFCompliantAsSecurityEventsProviderRequestSettings is a convenience function that returns SecurityEventsProviderSettingsNonSSFCompliant wrapped in SecurityEventsProviderRequestSettings
func SecurityEventsProviderSettingsNonSSFCompliantAsSecurityEventsProviderRequestSettings(v *SecurityEventsProviderSettingsNonSSFCompliant) SecurityEventsProviderRequestSettings {
	return SecurityEventsProviderRequestSettings{
		SecurityEventsProviderSettingsNonSSFCompliant: v,
	}
}

// SecurityEventsProviderSettingsSSFCompliantAsSecurityEventsProviderRequestSettings is a convenience function that returns SecurityEventsProviderSettingsSSFCompliant wrapped in SecurityEventsProviderRequestSettings
func SecurityEventsProviderSettingsSSFCompliantAsSecurityEventsProviderRequestSettings(v *SecurityEventsProviderSettingsSSFCompliant) SecurityEventsProviderRequestSettings {
	return SecurityEventsProviderRequestSettings{
		SecurityEventsProviderSettingsSSFCompliant: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *SecurityEventsProviderRequestSettings) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into SecurityEventsProviderSettingsNonSSFCompliant
        err = json.Unmarshal(data, &dst.SecurityEventsProviderSettingsNonSSFCompliant)
        if err == nil {
                jsonSecurityEventsProviderSettingsNonSSFCompliant, _ := json.Marshal(dst.SecurityEventsProviderSettingsNonSSFCompliant)
                if string(jsonSecurityEventsProviderSettingsNonSSFCompliant) == "{}" { // empty struct
                        dst.SecurityEventsProviderSettingsNonSSFCompliant = nil
                } else {
                        match++
                }
        } else {
                dst.SecurityEventsProviderSettingsNonSSFCompliant = nil
        }

        // try to unmarshal data into SecurityEventsProviderSettingsSSFCompliant
        err = json.Unmarshal(data, &dst.SecurityEventsProviderSettingsSSFCompliant)
        if err == nil {
                jsonSecurityEventsProviderSettingsSSFCompliant, _ := json.Marshal(dst.SecurityEventsProviderSettingsSSFCompliant)
                if string(jsonSecurityEventsProviderSettingsSSFCompliant) == "{}" { // empty struct
                        dst.SecurityEventsProviderSettingsSSFCompliant = nil
                } else {
                        match++
                }
        } else {
                dst.SecurityEventsProviderSettingsSSFCompliant = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.SecurityEventsProviderSettingsNonSSFCompliant = nil
                dst.SecurityEventsProviderSettingsSSFCompliant = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(SecurityEventsProviderRequestSettings)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(SecurityEventsProviderRequestSettings)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SecurityEventsProviderRequestSettings) MarshalJSON() ([]byte, error) {
	if src.SecurityEventsProviderSettingsNonSSFCompliant != nil {
		return json.Marshal(&src.SecurityEventsProviderSettingsNonSSFCompliant)
	}

	if src.SecurityEventsProviderSettingsSSFCompliant != nil {
		return json.Marshal(&src.SecurityEventsProviderSettingsSSFCompliant)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SecurityEventsProviderRequestSettings) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SecurityEventsProviderSettingsNonSSFCompliant != nil {
		return obj.SecurityEventsProviderSettingsNonSSFCompliant
	}

	if obj.SecurityEventsProviderSettingsSSFCompliant != nil {
		return obj.SecurityEventsProviderSettingsSSFCompliant
	}

	// all schemas are nil
	return nil
}

type NullableSecurityEventsProviderRequestSettings struct {
	value *SecurityEventsProviderRequestSettings
	isSet bool
}

func (v NullableSecurityEventsProviderRequestSettings) Get() *SecurityEventsProviderRequestSettings {
	return v.value
}

func (v *NullableSecurityEventsProviderRequestSettings) Set(val *SecurityEventsProviderRequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderRequestSettings(val *SecurityEventsProviderRequestSettings) *NullableSecurityEventsProviderRequestSettings {
	return &NullableSecurityEventsProviderRequestSettings{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



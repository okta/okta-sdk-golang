/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the OSAccountWindows type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSAccountWindows{}

// OSAccountWindows struct for OSAccountWindows
type OSAccountWindows struct {
	OSAccount
	Profile              WindowsOSAccountProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _OSAccountWindows OSAccountWindows

// NewOSAccountWindows instantiates a new OSAccountWindows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSAccountWindows(profile WindowsOSAccountProfile, created time.Time, deviceId string, id string, lastUpdated time.Time, platform string, links OSAccountLinks) *OSAccountWindows {
	this := OSAccountWindows{}
	this.Created = created
	this.DeviceId = deviceId
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Platform = platform
	this.Links = links
	this.Profile = profile
	return &this
}

// NewOSAccountWindowsWithDefaults instantiates a new OSAccountWindows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSAccountWindowsWithDefaults() *OSAccountWindows {
	this := OSAccountWindows{}
	return &this
}

// GetProfile returns the Profile field value
func (o *OSAccountWindows) GetProfile() WindowsOSAccountProfile {
	if o == nil {
		var ret WindowsOSAccountProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OSAccountWindows) GetProfileOk() (*WindowsOSAccountProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OSAccountWindows) SetProfile(v WindowsOSAccountProfile) {
	o.Profile = v
}

func (o OSAccountWindows) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSAccountWindows) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOSAccount, errOSAccount := json.Marshal(o.OSAccount)
	if errOSAccount != nil {
		return map[string]interface{}{}, errOSAccount
	}
	errOSAccount = json.Unmarshal([]byte(serializedOSAccount), &toSerialize)
	if errOSAccount != nil {
		return map[string]interface{}{}, errOSAccount
	}
	toSerialize["profile"] = o.Profile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSAccountWindows) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
		"created",
		"deviceId",
		"id",
		"lastUpdated",
		"platform",
		"_links",
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

	type OSAccountWindowsWithoutEmbeddedStruct struct {
		Profile WindowsOSAccountProfile `json:"profile"`
	}

	varOSAccountWindowsWithoutEmbeddedStruct := OSAccountWindowsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOSAccountWindowsWithoutEmbeddedStruct)
	if err == nil {
		varOSAccountWindows := _OSAccountWindows{}
		varOSAccountWindows.Profile = varOSAccountWindowsWithoutEmbeddedStruct.Profile
		*o = OSAccountWindows(varOSAccountWindows)
	} else {
		return err
	}

	varOSAccountWindows := _OSAccountWindows{}

	err = json.Unmarshal(data, &varOSAccountWindows)
	if err == nil {
		o.OSAccount = varOSAccountWindows.OSAccount
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")

		// remove fields from embedded structs
		reflectOSAccount := reflect.ValueOf(o.OSAccount)
		for i := 0; i < reflectOSAccount.Type().NumField(); i++ {
			t := reflectOSAccount.Type().Field(i)

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

type NullableOSAccountWindows struct {
	value *OSAccountWindows
	isSet bool
}

func (v NullableOSAccountWindows) Get() *OSAccountWindows {
	return v.value
}

func (v *NullableOSAccountWindows) Set(val *OSAccountWindows) {
	v.value = val
	v.isSet = true
}

func (v NullableOSAccountWindows) IsSet() bool {
	return v.isSet
}

func (v *NullableOSAccountWindows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSAccountWindows(val *OSAccountWindows) *NullableOSAccountWindows {
	return &NullableOSAccountWindows{value: val, isSet: true}
}

func (v NullableOSAccountWindows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSAccountWindows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

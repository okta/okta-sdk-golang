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

// checks if the OSAccountMacOS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSAccountMacOS{}

// OSAccountMacOS struct for OSAccountMacOS
type OSAccountMacOS struct {
	OSAccount
	Profile              MacOSAccountProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _OSAccountMacOS OSAccountMacOS

// NewOSAccountMacOS instantiates a new OSAccountMacOS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSAccountMacOS(profile MacOSAccountProfile, created time.Time, deviceId string, id string, lastUpdated time.Time, platform string, links OSAccountLinks) *OSAccountMacOS {
	this := OSAccountMacOS{}
	this.Created = created
	this.DeviceId = deviceId
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Platform = platform
	this.Links = links
	this.Profile = profile
	return &this
}

// NewOSAccountMacOSWithDefaults instantiates a new OSAccountMacOS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSAccountMacOSWithDefaults() *OSAccountMacOS {
	this := OSAccountMacOS{}
	return &this
}

// GetProfile returns the Profile field value
func (o *OSAccountMacOS) GetProfile() MacOSAccountProfile {
	if o == nil {
		var ret MacOSAccountProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OSAccountMacOS) GetProfileOk() (*MacOSAccountProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OSAccountMacOS) SetProfile(v MacOSAccountProfile) {
	o.Profile = v
}

func (o OSAccountMacOS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSAccountMacOS) ToMap() (map[string]interface{}, error) {
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

func (o *OSAccountMacOS) UnmarshalJSON(data []byte) (err error) {
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

	type OSAccountMacOSWithoutEmbeddedStruct struct {
		Profile MacOSAccountProfile `json:"profile"`
	}

	varOSAccountMacOSWithoutEmbeddedStruct := OSAccountMacOSWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOSAccountMacOSWithoutEmbeddedStruct)
	if err == nil {
		varOSAccountMacOS := _OSAccountMacOS{}
		varOSAccountMacOS.Profile = varOSAccountMacOSWithoutEmbeddedStruct.Profile
		*o = OSAccountMacOS(varOSAccountMacOS)
	} else {
		return err
	}

	varOSAccountMacOS := _OSAccountMacOS{}

	err = json.Unmarshal(data, &varOSAccountMacOS)
	if err == nil {
		o.OSAccount = varOSAccountMacOS.OSAccount
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

type NullableOSAccountMacOS struct {
	value *OSAccountMacOS
	isSet bool
}

func (v NullableOSAccountMacOS) Get() *OSAccountMacOS {
	return v.value
}

func (v *NullableOSAccountMacOS) Set(val *OSAccountMacOS) {
	v.value = val
	v.isSet = true
}

func (v NullableOSAccountMacOS) IsSet() bool {
	return v.isSet
}

func (v *NullableOSAccountMacOS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSAccountMacOS(val *OSAccountMacOS) *NullableOSAccountMacOS {
	return &NullableOSAccountMacOS{value: val, isSet: true}
}

func (v NullableOSAccountMacOS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSAccountMacOS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

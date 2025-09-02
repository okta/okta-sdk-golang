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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// UserFactorPush struct for UserFactorPush
type UserFactorPush struct {
	UserFactor
	Profile *UserFactorPushProfile `json:"profile,omitempty"`
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPush UserFactorPush

// NewUserFactorPush instantiates a new UserFactorPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPush() *UserFactorPush {
	this := UserFactorPush{}
	return &this
}

// NewUserFactorPushWithDefaults instantiates a new UserFactorPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushWithDefaults() *UserFactorPush {
	this := UserFactorPush{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorPush) GetProfile() UserFactorPushProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorPushProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPush) GetProfileOk() (*UserFactorPushProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorPush) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorPushProfile and assigns it to the Profile field.
func (o *UserFactorPush) SetProfile(v UserFactorPushProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorPush) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPush) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorPush) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorPush) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorPush) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPush) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorPushWithoutEmbeddedStruct struct {
		Profile *UserFactorPushProfile `json:"profile,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}

	varUserFactorPushWithoutEmbeddedStruct := UserFactorPushWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorPushWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPush := _UserFactorPush{}
		varUserFactorPush.Profile = varUserFactorPushWithoutEmbeddedStruct.Profile
		varUserFactorPush.Provider = varUserFactorPushWithoutEmbeddedStruct.Provider
		*o = UserFactorPush(varUserFactorPush)
	} else {
		return err
	}

	varUserFactorPush := _UserFactorPush{}

	err = json.Unmarshal(bytes, &varUserFactorPush)
	if err == nil {
		o.UserFactor = varUserFactorPush.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		delete(additionalProperties, "provider")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

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
	} else {
		return err
	}

	return err
}

type NullableUserFactorPush struct {
	value *UserFactorPush
	isSet bool
}

func (v NullableUserFactorPush) Get() *UserFactorPush {
	return v.value
}

func (v *NullableUserFactorPush) Set(val *UserFactorPush) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPush) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPush(val *UserFactorPush) *NullableUserFactorPush {
	return &NullableUserFactorPush{value: val, isSet: true}
}

func (v NullableUserFactorPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


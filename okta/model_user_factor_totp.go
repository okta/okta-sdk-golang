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
	"reflect"
	"strings"
)

// UserFactorTOTP struct for UserFactorTOTP
type UserFactorTOTP struct {
	UserFactor
	FactorType interface{} `json:"factorType,omitempty"`
	Profile *UserFactorTOTPProfile `json:"profile,omitempty"`
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTOTP UserFactorTOTP

// NewUserFactorTOTP instantiates a new UserFactorTOTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTOTP() *UserFactorTOTP {
	this := UserFactorTOTP{}
	return &this
}

// NewUserFactorTOTPWithDefaults instantiates a new UserFactorTOTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTOTPWithDefaults() *UserFactorTOTP {
	this := UserFactorTOTP{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorTOTP) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorTOTP) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorTOTP) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorTOTP) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorTOTP) GetProfile() UserFactorTOTPProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorTOTPProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTOTP) GetProfileOk() (*UserFactorTOTPProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorTOTP) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorTOTPProfile and assigns it to the Profile field.
func (o *UserFactorTOTP) SetProfile(v UserFactorTOTPProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorTOTP) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTOTP) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorTOTP) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorTOTP) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorTOTP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
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

func (o *UserFactorTOTP) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorTOTPWithoutEmbeddedStruct struct {
		FactorType interface{} `json:"factorType,omitempty"`
		Profile *UserFactorTOTPProfile `json:"profile,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}

	varUserFactorTOTPWithoutEmbeddedStruct := UserFactorTOTPWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorTOTPWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorTOTP := _UserFactorTOTP{}
		varUserFactorTOTP.FactorType = varUserFactorTOTPWithoutEmbeddedStruct.FactorType
		varUserFactorTOTP.Profile = varUserFactorTOTPWithoutEmbeddedStruct.Profile
		varUserFactorTOTP.Provider = varUserFactorTOTPWithoutEmbeddedStruct.Provider
		*o = UserFactorTOTP(varUserFactorTOTP)
	} else {
		return err
	}

	varUserFactorTOTP := _UserFactorTOTP{}

	err = json.Unmarshal(bytes, &varUserFactorTOTP)
	if err == nil {
		o.UserFactor = varUserFactorTOTP.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "factorType")
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

type NullableUserFactorTOTP struct {
	value *UserFactorTOTP
	isSet bool
}

func (v NullableUserFactorTOTP) Get() *UserFactorTOTP {
	return v.value
}

func (v *NullableUserFactorTOTP) Set(val *UserFactorTOTP) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTOTP) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTOTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTOTP(val *UserFactorTOTP) *NullableUserFactorTOTP {
	return &NullableUserFactorTOTP{value: val, isSet: true}
}

func (v NullableUserFactorTOTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTOTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


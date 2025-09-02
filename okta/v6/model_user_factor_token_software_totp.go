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

// UserFactorTokenSoftwareTOTP struct for UserFactorTokenSoftwareTOTP
type UserFactorTokenSoftwareTOTP struct {
	UserFactor
	FactorType interface{} `json:"factorType,omitempty"`
	Profile *UserFactorTokenProfile `json:"profile,omitempty"`
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenSoftwareTOTP UserFactorTokenSoftwareTOTP

// NewUserFactorTokenSoftwareTOTP instantiates a new UserFactorTokenSoftwareTOTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenSoftwareTOTP() *UserFactorTokenSoftwareTOTP {
	this := UserFactorTokenSoftwareTOTP{}
	return &this
}

// NewUserFactorTokenSoftwareTOTPWithDefaults instantiates a new UserFactorTokenSoftwareTOTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenSoftwareTOTPWithDefaults() *UserFactorTokenSoftwareTOTP {
	this := UserFactorTokenSoftwareTOTP{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorTokenSoftwareTOTP) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorTokenSoftwareTOTP) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorTokenSoftwareTOTP) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorTokenSoftwareTOTP) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorTokenSoftwareTOTP) GetProfile() UserFactorTokenProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorTokenProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenSoftwareTOTP) GetProfileOk() (*UserFactorTokenProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorTokenSoftwareTOTP) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorTokenProfile and assigns it to the Profile field.
func (o *UserFactorTokenSoftwareTOTP) SetProfile(v UserFactorTokenProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorTokenSoftwareTOTP) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenSoftwareTOTP) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorTokenSoftwareTOTP) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorTokenSoftwareTOTP) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorTokenSoftwareTOTP) MarshalJSON() ([]byte, error) {
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

func (o *UserFactorTokenSoftwareTOTP) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorTokenSoftwareTOTPWithoutEmbeddedStruct struct {
		FactorType interface{} `json:"factorType,omitempty"`
		Profile *UserFactorTokenProfile `json:"profile,omitempty"`
		Provider *string `json:"provider,omitempty"`
	}

	varUserFactorTokenSoftwareTOTPWithoutEmbeddedStruct := UserFactorTokenSoftwareTOTPWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorTokenSoftwareTOTPWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorTokenSoftwareTOTP := _UserFactorTokenSoftwareTOTP{}
		varUserFactorTokenSoftwareTOTP.FactorType = varUserFactorTokenSoftwareTOTPWithoutEmbeddedStruct.FactorType
		varUserFactorTokenSoftwareTOTP.Profile = varUserFactorTokenSoftwareTOTPWithoutEmbeddedStruct.Profile
		varUserFactorTokenSoftwareTOTP.Provider = varUserFactorTokenSoftwareTOTPWithoutEmbeddedStruct.Provider
		*o = UserFactorTokenSoftwareTOTP(varUserFactorTokenSoftwareTOTP)
	} else {
		return err
	}

	varUserFactorTokenSoftwareTOTP := _UserFactorTokenSoftwareTOTP{}

	err = json.Unmarshal(bytes, &varUserFactorTokenSoftwareTOTP)
	if err == nil {
		o.UserFactor = varUserFactorTokenSoftwareTOTP.UserFactor
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

type NullableUserFactorTokenSoftwareTOTP struct {
	value *UserFactorTokenSoftwareTOTP
	isSet bool
}

func (v NullableUserFactorTokenSoftwareTOTP) Get() *UserFactorTokenSoftwareTOTP {
	return v.value
}

func (v *NullableUserFactorTokenSoftwareTOTP) Set(val *UserFactorTokenSoftwareTOTP) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenSoftwareTOTP) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenSoftwareTOTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenSoftwareTOTP(val *UserFactorTokenSoftwareTOTP) *NullableUserFactorTokenSoftwareTOTP {
	return &NullableUserFactorTokenSoftwareTOTP{value: val, isSet: true}
}

func (v NullableUserFactorTokenSoftwareTOTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenSoftwareTOTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


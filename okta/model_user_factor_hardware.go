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

// UserFactorHardware struct for UserFactorHardware
type UserFactorHardware struct {
	UserFactor
	FactorType interface{} `json:"factorType,omitempty"`
	Profile *UserFactorHardwareProfile `json:"profile,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Verify *UserFactorHardwareAllOfVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorHardware UserFactorHardware

// NewUserFactorHardware instantiates a new UserFactorHardware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorHardware() *UserFactorHardware {
	this := UserFactorHardware{}
	return &this
}

// NewUserFactorHardwareWithDefaults instantiates a new UserFactorHardware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorHardwareWithDefaults() *UserFactorHardware {
	this := UserFactorHardware{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorHardware) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorHardware) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorHardware) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorHardware) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorHardware) GetProfile() UserFactorHardwareProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorHardwareProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorHardware) GetProfileOk() (*UserFactorHardwareProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorHardware) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorHardwareProfile and assigns it to the Profile field.
func (o *UserFactorHardware) SetProfile(v UserFactorHardwareProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorHardware) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorHardware) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorHardware) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorHardware) SetProvider(v string) {
	o.Provider = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorHardware) GetVerify() UserFactorHardwareAllOfVerify {
	if o == nil || o.Verify == nil {
		var ret UserFactorHardwareAllOfVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorHardware) GetVerifyOk() (*UserFactorHardwareAllOfVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorHardware) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given UserFactorHardwareAllOfVerify and assigns it to the Verify field.
func (o *UserFactorHardware) SetVerify(v UserFactorHardwareAllOfVerify) {
	o.Verify = &v
}

func (o UserFactorHardware) MarshalJSON() ([]byte, error) {
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
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorHardware) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorHardwareWithoutEmbeddedStruct struct {
		FactorType interface{} `json:"factorType,omitempty"`
		Profile *UserFactorHardwareProfile `json:"profile,omitempty"`
		Provider *string `json:"provider,omitempty"`
		Verify *UserFactorHardwareAllOfVerify `json:"verify,omitempty"`
	}

	varUserFactorHardwareWithoutEmbeddedStruct := UserFactorHardwareWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorHardwareWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorHardware := _UserFactorHardware{}
		varUserFactorHardware.FactorType = varUserFactorHardwareWithoutEmbeddedStruct.FactorType
		varUserFactorHardware.Profile = varUserFactorHardwareWithoutEmbeddedStruct.Profile
		varUserFactorHardware.Provider = varUserFactorHardwareWithoutEmbeddedStruct.Provider
		varUserFactorHardware.Verify = varUserFactorHardwareWithoutEmbeddedStruct.Verify
		*o = UserFactorHardware(varUserFactorHardware)
	} else {
		return err
	}

	varUserFactorHardware := _UserFactorHardware{}

	err = json.Unmarshal(bytes, &varUserFactorHardware)
	if err == nil {
		o.UserFactor = varUserFactorHardware.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "factorType")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "verify")

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

type NullableUserFactorHardware struct {
	value *UserFactorHardware
	isSet bool
}

func (v NullableUserFactorHardware) Get() *UserFactorHardware {
	return v.value
}

func (v *NullableUserFactorHardware) Set(val *UserFactorHardware) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorHardware) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorHardware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorHardware(val *UserFactorHardware) *NullableUserFactorHardware {
	return &NullableUserFactorHardware{value: val, isSet: true}
}

func (v NullableUserFactorHardware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorHardware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


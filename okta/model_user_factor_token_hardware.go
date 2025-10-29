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
	"reflect"
	"strings"
)

// checks if the UserFactorTokenHardware type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorTokenHardware{}

// UserFactorTokenHardware struct for UserFactorTokenHardware
type UserFactorTokenHardware struct {
	UserFactor
	FactorType           interface{}                         `json:"factorType,omitempty"`
	Profile              *UserFactorTokenProfile             `json:"profile,omitempty"`
	Provider             *string                             `json:"provider,omitempty"`
	Verify               *UserFactorTokenHardwareAllOfVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenHardware UserFactorTokenHardware

// NewUserFactorTokenHardware instantiates a new UserFactorTokenHardware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenHardware() *UserFactorTokenHardware {
	this := UserFactorTokenHardware{}
	return &this
}

// NewUserFactorTokenHardwareWithDefaults instantiates a new UserFactorTokenHardware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenHardwareWithDefaults() *UserFactorTokenHardware {
	this := UserFactorTokenHardware{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorTokenHardware) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorTokenHardware) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorTokenHardware) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorTokenHardware) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorTokenHardware) GetProfile() UserFactorTokenProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserFactorTokenProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHardware) GetProfileOk() (*UserFactorTokenProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorTokenHardware) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorTokenProfile and assigns it to the Profile field.
func (o *UserFactorTokenHardware) SetProfile(v UserFactorTokenProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorTokenHardware) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHardware) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorTokenHardware) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorTokenHardware) SetProvider(v string) {
	o.Provider = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorTokenHardware) GetVerify() UserFactorTokenHardwareAllOfVerify {
	if o == nil || IsNil(o.Verify) {
		var ret UserFactorTokenHardwareAllOfVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHardware) GetVerifyOk() (*UserFactorTokenHardwareAllOfVerify, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorTokenHardware) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given UserFactorTokenHardwareAllOfVerify and assigns it to the Verify field.
func (o *UserFactorTokenHardware) SetVerify(v UserFactorTokenHardwareAllOfVerify) {
	o.Verify = &v
}

func (o UserFactorTokenHardware) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorTokenHardware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return map[string]interface{}{}, errUserFactor
	}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Verify) {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorTokenHardware) UnmarshalJSON(data []byte) (err error) {
	type UserFactorTokenHardwareWithoutEmbeddedStruct struct {
		FactorType interface{}                         `json:"factorType,omitempty"`
		Profile    *UserFactorTokenProfile             `json:"profile,omitempty"`
		Provider   *string                             `json:"provider,omitempty"`
		Verify     *UserFactorTokenHardwareAllOfVerify `json:"verify,omitempty"`
	}

	varUserFactorTokenHardwareWithoutEmbeddedStruct := UserFactorTokenHardwareWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserFactorTokenHardwareWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorTokenHardware := _UserFactorTokenHardware{}
		varUserFactorTokenHardware.FactorType = varUserFactorTokenHardwareWithoutEmbeddedStruct.FactorType
		varUserFactorTokenHardware.Profile = varUserFactorTokenHardwareWithoutEmbeddedStruct.Profile
		varUserFactorTokenHardware.Provider = varUserFactorTokenHardwareWithoutEmbeddedStruct.Provider
		varUserFactorTokenHardware.Verify = varUserFactorTokenHardwareWithoutEmbeddedStruct.Verify
		*o = UserFactorTokenHardware(varUserFactorTokenHardware)
	} else {
		return err
	}

	varUserFactorTokenHardware := _UserFactorTokenHardware{}

	err = json.Unmarshal(data, &varUserFactorTokenHardware)
	if err == nil {
		o.UserFactor = varUserFactorTokenHardware.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableUserFactorTokenHardware struct {
	value *UserFactorTokenHardware
	isSet bool
}

func (v NullableUserFactorTokenHardware) Get() *UserFactorTokenHardware {
	return v.value
}

func (v *NullableUserFactorTokenHardware) Set(val *UserFactorTokenHardware) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenHardware) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenHardware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenHardware(val *UserFactorTokenHardware) *NullableUserFactorTokenHardware {
	return &NullableUserFactorTokenHardware{value: val, isSet: true}
}

func (v NullableUserFactorTokenHardware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenHardware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

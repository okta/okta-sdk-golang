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

// checks if the UserFactorSecurityQuestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorSecurityQuestion{}

// UserFactorSecurityQuestion struct for UserFactorSecurityQuestion
type UserFactorSecurityQuestion struct {
	UserFactor
	FactorType           interface{}                        `json:"factorType,omitempty"`
	Profile              *UserFactorSecurityQuestionProfile `json:"profile,omitempty"`
	Provider             *string                            `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSecurityQuestion UserFactorSecurityQuestion

// NewUserFactorSecurityQuestion instantiates a new UserFactorSecurityQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSecurityQuestion() *UserFactorSecurityQuestion {
	this := UserFactorSecurityQuestion{}
	return &this
}

// NewUserFactorSecurityQuestionWithDefaults instantiates a new UserFactorSecurityQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSecurityQuestionWithDefaults() *UserFactorSecurityQuestion {
	this := UserFactorSecurityQuestion{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorSecurityQuestion) GetFactorType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorSecurityQuestion) GetFactorTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return &o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestion) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given interface{} and assigns it to the FactorType field.
func (o *UserFactorSecurityQuestion) SetFactorType(v interface{}) {
	o.FactorType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorSecurityQuestion) GetProfile() UserFactorSecurityQuestionProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserFactorSecurityQuestionProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSecurityQuestion) GetProfileOk() (*UserFactorSecurityQuestionProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestion) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorSecurityQuestionProfile and assigns it to the Profile field.
func (o *UserFactorSecurityQuestion) SetProfile(v UserFactorSecurityQuestionProfile) {
	o.Profile = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorSecurityQuestion) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSecurityQuestion) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorSecurityQuestion) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorSecurityQuestion) SetProvider(v string) {
	o.Provider = &v
}

func (o UserFactorSecurityQuestion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorSecurityQuestion) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorSecurityQuestion) UnmarshalJSON(data []byte) (err error) {
	type UserFactorSecurityQuestionWithoutEmbeddedStruct struct {
		FactorType interface{}                        `json:"factorType,omitempty"`
		Profile    *UserFactorSecurityQuestionProfile `json:"profile,omitempty"`
		Provider   *string                            `json:"provider,omitempty"`
	}

	varUserFactorSecurityQuestionWithoutEmbeddedStruct := UserFactorSecurityQuestionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserFactorSecurityQuestionWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorSecurityQuestion := _UserFactorSecurityQuestion{}
		varUserFactorSecurityQuestion.FactorType = varUserFactorSecurityQuestionWithoutEmbeddedStruct.FactorType
		varUserFactorSecurityQuestion.Profile = varUserFactorSecurityQuestionWithoutEmbeddedStruct.Profile
		varUserFactorSecurityQuestion.Provider = varUserFactorSecurityQuestionWithoutEmbeddedStruct.Provider
		*o = UserFactorSecurityQuestion(varUserFactorSecurityQuestion)
	} else {
		return err
	}

	varUserFactorSecurityQuestion := _UserFactorSecurityQuestion{}

	err = json.Unmarshal(data, &varUserFactorSecurityQuestion)
	if err == nil {
		o.UserFactor = varUserFactorSecurityQuestion.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableUserFactorSecurityQuestion struct {
	value *UserFactorSecurityQuestion
	isSet bool
}

func (v NullableUserFactorSecurityQuestion) Get() *UserFactorSecurityQuestion {
	return v.value
}

func (v *NullableUserFactorSecurityQuestion) Set(val *UserFactorSecurityQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSecurityQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSecurityQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSecurityQuestion(val *UserFactorSecurityQuestion) *NullableUserFactorSecurityQuestion {
	return &NullableUserFactorSecurityQuestion{value: val, isSet: true}
}

func (v NullableUserFactorSecurityQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSecurityQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

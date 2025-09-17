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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the UserRiskLevelExists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskLevelExists{}

// UserRiskLevelExists struct for UserRiskLevelExists
type UserRiskLevelExists struct {
	UserRiskGetResponse
	// Describes the risk level for the user
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRiskLevelExists UserRiskLevelExists

// NewUserRiskLevelExists instantiates a new UserRiskLevelExists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskLevelExists() *UserRiskLevelExists {
	this := UserRiskLevelExists{}
	return &this
}

// NewUserRiskLevelExistsWithDefaults instantiates a new UserRiskLevelExists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskLevelExistsWithDefaults() *UserRiskLevelExists {
	this := UserRiskLevelExists{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UserRiskLevelExists) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskLevelExists) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UserRiskLevelExists) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UserRiskLevelExists) SetReason(v string) {
	o.Reason = &v
}

func (o UserRiskLevelExists) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskLevelExists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserRiskGetResponse, errUserRiskGetResponse := json.Marshal(o.UserRiskGetResponse)
	if errUserRiskGetResponse != nil {
		return map[string]interface{}{}, errUserRiskGetResponse
	}
	errUserRiskGetResponse = json.Unmarshal([]byte(serializedUserRiskGetResponse), &toSerialize)
	if errUserRiskGetResponse != nil {
		return map[string]interface{}{}, errUserRiskGetResponse
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRiskLevelExists) UnmarshalJSON(data []byte) (err error) {
	type UserRiskLevelExistsWithoutEmbeddedStruct struct {
		// Describes the risk level for the user
		Reason *string `json:"reason,omitempty"`
	}

	varUserRiskLevelExistsWithoutEmbeddedStruct := UserRiskLevelExistsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserRiskLevelExistsWithoutEmbeddedStruct)
	if err == nil {
		varUserRiskLevelExists := _UserRiskLevelExists{}
		varUserRiskLevelExists.Reason = varUserRiskLevelExistsWithoutEmbeddedStruct.Reason
		*o = UserRiskLevelExists(varUserRiskLevelExists)
	} else {
		return err
	}

	varUserRiskLevelExists := _UserRiskLevelExists{}

	err = json.Unmarshal(data, &varUserRiskLevelExists)
	if err == nil {
		o.UserRiskGetResponse = varUserRiskLevelExists.UserRiskGetResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reason")

		// remove fields from embedded structs
		reflectUserRiskGetResponse := reflect.ValueOf(o.UserRiskGetResponse)
		for i := 0; i < reflectUserRiskGetResponse.Type().NumField(); i++ {
			t := reflectUserRiskGetResponse.Type().Field(i)

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

type NullableUserRiskLevelExists struct {
	value *UserRiskLevelExists
	isSet bool
}

func (v NullableUserRiskLevelExists) Get() *UserRiskLevelExists {
	return v.value
}

func (v *NullableUserRiskLevelExists) Set(val *UserRiskLevelExists) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskLevelExists) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskLevelExists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskLevelExists(val *UserRiskLevelExists) *NullableUserRiskLevelExists {
	return &NullableUserRiskLevelExists{value: val, isSet: true}
}

func (v NullableUserRiskLevelExists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskLevelExists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

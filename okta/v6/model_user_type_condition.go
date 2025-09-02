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
	"fmt"
)

// UserTypeCondition <x-lifecycle class=\"oie\"></x-lifecycle> Specifies which user types to include and/or exclude
type UserTypeCondition struct {
	// The user types to exclude
	Exclude []string `json:"exclude"`
	// The user types to include
	Include []string `json:"include"`
	AdditionalProperties map[string]interface{}
}

type _UserTypeCondition UserTypeCondition

// NewUserTypeCondition instantiates a new UserTypeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypeCondition(exclude []string, include []string) *UserTypeCondition {
	this := UserTypeCondition{}
	this.Exclude = exclude
	this.Include = include
	return &this
}

// NewUserTypeConditionWithDefaults instantiates a new UserTypeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeConditionWithDefaults() *UserTypeCondition {
	this := UserTypeCondition{}
	return &this
}

// GetExclude returns the Exclude field value
func (o *UserTypeCondition) GetExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value
// and a boolean to check if the value has been set.
func (o *UserTypeCondition) GetExcludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exclude, true
}

// SetExclude sets field value
func (o *UserTypeCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value
func (o *UserTypeCondition) GetInclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *UserTypeCondition) GetIncludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *UserTypeCondition) SetInclude(v []string) {
	o.Include = v
}

func (o UserTypeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exclude"] = o.Exclude
	}
	if true {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserTypeCondition) UnmarshalJSON(bytes []byte) (err error) {
	varUserTypeCondition := _UserTypeCondition{}

	err = json.Unmarshal(bytes, &varUserTypeCondition)
	if err == nil {
		*o = UserTypeCondition(varUserTypeCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserTypeCondition struct {
	value *UserTypeCondition
	isSet bool
}

func (v NullableUserTypeCondition) Get() *UserTypeCondition {
	return v.value
}

func (v *NullableUserTypeCondition) Set(val *UserTypeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypeCondition(val *UserTypeCondition) *NullableUserTypeCondition {
	return &NullableUserTypeCondition{value: val, isSet: true}
}

func (v NullableUserTypeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


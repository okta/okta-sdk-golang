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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCondition{}

// UserCondition Specifies a set of users to be included or excluded
type UserCondition struct {
	// Users to be excluded
	Exclude []string `json:"exclude,omitempty"`
	// Users to be included
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCondition UserCondition

// NewUserCondition instantiates a new UserCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCondition() *UserCondition {
	this := UserCondition{}
	return &this
}

// NewUserConditionWithDefaults instantiates a new UserCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConditionWithDefaults() *UserCondition {
	this := UserCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *UserCondition) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *UserCondition) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *UserCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *UserCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *UserCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *UserCondition) SetInclude(v []string) {
	o.Include = v
}

func (o UserCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCondition) UnmarshalJSON(data []byte) (err error) {
	varUserCondition := _UserCondition{}

	err = json.Unmarshal(data, &varUserCondition)

	if err != nil {
		return err
	}

	*o = UserCondition(varUserCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCondition struct {
	value *UserCondition
	isSet bool
}

func (v NullableUserCondition) Get() *UserCondition {
	return v.value
}

func (v *NullableUserCondition) Set(val *UserCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCondition(val *UserCondition) *NullableUserCondition {
	return &NullableUserCondition{value: val, isSet: true}
}

func (v NullableUserCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

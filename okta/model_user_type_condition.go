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
)

// UserTypeCondition struct for UserTypeCondition
type UserTypeCondition struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserTypeCondition UserTypeCondition

// NewUserTypeCondition instantiates a new UserTypeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypeCondition() *UserTypeCondition {
	this := UserTypeCondition{}
	return &this
}

// NewUserTypeConditionWithDefaults instantiates a new UserTypeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeConditionWithDefaults() *UserTypeCondition {
	this := UserTypeCondition{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *UserTypeCondition) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypeCondition) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *UserTypeCondition) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *UserTypeCondition) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *UserTypeCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypeCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *UserTypeCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *UserTypeCondition) SetInclude(v []string) {
	o.Include = v
}

func (o UserTypeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.Include != nil {
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


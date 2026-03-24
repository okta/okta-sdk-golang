/*
Okta Admin Management API

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

// checks if the GroupProfileResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupProfileResult{}

// GroupProfileResult struct for GroupProfileResult
type GroupProfileResult struct {
	// The ID of the group
	Id *string `json:"id,omitempty"`
	// Map of requested attributes and their values
	Profile              map[string]interface{} `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupProfileResult GroupProfileResult

// NewGroupProfileResult instantiates a new GroupProfileResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupProfileResult() *GroupProfileResult {
	this := GroupProfileResult{}
	return &this
}

// NewGroupProfileResultWithDefaults instantiates a new GroupProfileResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupProfileResultWithDefaults() *GroupProfileResult {
	this := GroupProfileResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupProfileResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProfileResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupProfileResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupProfileResult) SetId(v string) {
	o.Id = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GroupProfileResult) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProfileResult) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GroupProfileResult) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *GroupProfileResult) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o GroupProfileResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupProfileResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupProfileResult) UnmarshalJSON(data []byte) (err error) {
	varGroupProfileResult := _GroupProfileResult{}

	err = json.Unmarshal(data, &varGroupProfileResult)

	if err != nil {
		return err
	}

	*o = GroupProfileResult(varGroupProfileResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupProfileResult struct {
	value *GroupProfileResult
	isSet bool
}

func (v NullableGroupProfileResult) Get() *GroupProfileResult {
	return v.value
}

func (v *NullableGroupProfileResult) Set(val *GroupProfileResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupProfileResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupProfileResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupProfileResult(val *GroupProfileResult) *NullableGroupProfileResult {
	return &NullableGroupProfileResult{value: val, isSet: true}
}

func (v NullableGroupProfileResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupProfileResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

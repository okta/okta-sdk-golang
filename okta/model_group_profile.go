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

// GroupProfile struct for GroupProfile
type GroupProfile struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupProfile GroupProfile

// NewGroupProfile instantiates a new GroupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupProfile() *GroupProfile {
	this := GroupProfile{}
	return &this
}

// NewGroupProfileWithDefaults instantiates a new GroupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupProfileWithDefaults() *GroupProfile {
	this := GroupProfile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupProfile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupProfile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupProfile) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupProfile) SetName(v string) {
	o.Name = &v
}

func (o GroupProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupProfile) UnmarshalJSON(bytes []byte) (err error) {
	varGroupProfile := _GroupProfile{}

	err = json.Unmarshal(bytes, &varGroupProfile)
	if err == nil {
		*o = GroupProfile(varGroupProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupProfile struct {
	value *GroupProfile
	isSet bool
}

func (v NullableGroupProfile) Get() *GroupProfile {
	return v.value
}

func (v *NullableGroupProfile) Set(val *GroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupProfile(val *GroupProfile) *NullableGroupProfile {
	return &NullableGroupProfile{value: val, isSet: true}
}

func (v NullableGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


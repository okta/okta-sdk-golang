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

// UserTypePutRequest struct for UserTypePutRequest
type UserTypePutRequest struct {
	// The human-readable description of the User Type
	Description string `json:"description"`
	// The human-readable name of the User Type
	DisplayName string `json:"displayName"`
	// The name of the existing type
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _UserTypePutRequest UserTypePutRequest

// NewUserTypePutRequest instantiates a new UserTypePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypePutRequest(description string, displayName string, name string) *UserTypePutRequest {
	this := UserTypePutRequest{}
	this.Description = description
	this.DisplayName = displayName
	this.Name = name
	return &this
}

// NewUserTypePutRequestWithDefaults instantiates a new UserTypePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypePutRequestWithDefaults() *UserTypePutRequest {
	this := UserTypePutRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *UserTypePutRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UserTypePutRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UserTypePutRequest) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value
func (o *UserTypePutRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *UserTypePutRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *UserTypePutRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetName returns the Name field value
func (o *UserTypePutRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserTypePutRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserTypePutRequest) SetName(v string) {
	o.Name = v
}

func (o UserTypePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserTypePutRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUserTypePutRequest := _UserTypePutRequest{}

	err = json.Unmarshal(bytes, &varUserTypePutRequest)
	if err == nil {
		*o = UserTypePutRequest(varUserTypePutRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserTypePutRequest struct {
	value *UserTypePutRequest
	isSet bool
}

func (v NullableUserTypePutRequest) Get() *UserTypePutRequest {
	return v.value
}

func (v *NullableUserTypePutRequest) Set(val *UserTypePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypePutRequest(val *UserTypePutRequest) *NullableUserTypePutRequest {
	return &NullableUserTypePutRequest{value: val, isSet: true}
}

func (v NullableUserTypePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


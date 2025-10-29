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
	"fmt"
)

// checks if the UserTypePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTypePutRequest{}

// UserTypePutRequest struct for UserTypePutRequest
type UserTypePutRequest struct {
	// The human-readable description of the user type
	Description string `json:"description"`
	// The human-readable name of the user type
	DisplayName string `json:"displayName"`
	// The name of the existing type
	Name                 string `json:"name"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTypePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserTypePutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"displayName",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserTypePutRequest := _UserTypePutRequest{}

	err = json.Unmarshal(data, &varUserTypePutRequest)

	if err != nil {
		return err
	}

	*o = UserTypePutRequest(varUserTypePutRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
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

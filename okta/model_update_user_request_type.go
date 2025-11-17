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

// checks if the UpdateUserRequestType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserRequestType{}

// UpdateUserRequestType The ID of the user type. Add this value if you want to create a user with a non-default [User Type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/). The user type determines which [schema](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Schema/) applies to that user. After a user has been created, the user can only be assigned a different user type by an admin through a full replacement (`PUT`) operation.
type UpdateUserRequestType struct {
	// The ID of the user type
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserRequestType UpdateUserRequestType

// NewUpdateUserRequestType instantiates a new UpdateUserRequestType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequestType() *UpdateUserRequestType {
	this := UpdateUserRequestType{}
	return &this
}

// NewUpdateUserRequestTypeWithDefaults instantiates a new UpdateUserRequestType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestTypeWithDefaults() *UpdateUserRequestType {
	this := UpdateUserRequestType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateUserRequestType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequestType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateUserRequestType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateUserRequestType) SetId(v string) {
	o.Id = &v
}

func (o UpdateUserRequestType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserRequestType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateUserRequestType) UnmarshalJSON(data []byte) (err error) {
	varUpdateUserRequestType := _UpdateUserRequestType{}

	err = json.Unmarshal(data, &varUpdateUserRequestType)

	if err != nil {
		return err
	}

	*o = UpdateUserRequestType(varUpdateUserRequestType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserRequestType struct {
	value *UpdateUserRequestType
	isSet bool
}

func (v NullableUpdateUserRequestType) Get() *UpdateUserRequestType {
	return v.value
}

func (v *NullableUpdateUserRequestType) Set(val *UpdateUserRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequestType(val *UpdateUserRequestType) *NullableUpdateUserRequestType {
	return &NullableUpdateUserRequestType{value: val, isSet: true}
}

func (v NullableUpdateUserRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// CreateUserRequestType The ID of the user type. Add this value if you want to create a user with a non-default [user type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/). The user type determines which [schema](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Schema/) applies to that user. After a user has been created, the user can  only be assigned a different user type by an administrator through a full replacement (`PUT`) operation.
type CreateUserRequestType struct {
	// The ID of the user type
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserRequestType CreateUserRequestType

// NewCreateUserRequestType instantiates a new CreateUserRequestType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequestType() *CreateUserRequestType {
	this := CreateUserRequestType{}
	return &this
}

// NewCreateUserRequestTypeWithDefaults instantiates a new CreateUserRequestType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestTypeWithDefaults() *CreateUserRequestType {
	this := CreateUserRequestType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateUserRequestType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequestType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateUserRequestType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateUserRequestType) SetId(v string) {
	o.Id = &v
}

func (o CreateUserRequestType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUserRequestType) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUserRequestType := _CreateUserRequestType{}

	err = json.Unmarshal(bytes, &varCreateUserRequestType)
	if err == nil {
		*o = CreateUserRequestType(varCreateUserRequestType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateUserRequestType struct {
	value *CreateUserRequestType
	isSet bool
}

func (v NullableCreateUserRequestType) Get() *CreateUserRequestType {
	return v.value
}

func (v *NullableCreateUserRequestType) Set(val *CreateUserRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequestType(val *CreateUserRequestType) *NullableCreateUserRequestType {
	return &NullableCreateUserRequestType{value: val, isSet: true}
}

func (v NullableCreateUserRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

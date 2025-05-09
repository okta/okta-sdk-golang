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

// UserTypePostRequest struct for UserTypePostRequest
type UserTypePostRequest struct {
	// The updated human-readable description of the User Type
	Description *string `json:"description,omitempty"`
	// The updated human-readable display name for the User Type
	DisplayName *string `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserTypePostRequest UserTypePostRequest

// NewUserTypePostRequest instantiates a new UserTypePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypePostRequest() *UserTypePostRequest {
	this := UserTypePostRequest{}
	return &this
}

// NewUserTypePostRequestWithDefaults instantiates a new UserTypePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypePostRequestWithDefaults() *UserTypePostRequest {
	this := UserTypePostRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserTypePostRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypePostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserTypePostRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserTypePostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserTypePostRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypePostRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserTypePostRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserTypePostRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o UserTypePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserTypePostRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUserTypePostRequest := _UserTypePostRequest{}

	err = json.Unmarshal(bytes, &varUserTypePostRequest)
	if err == nil {
		*o = UserTypePostRequest(varUserTypePostRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserTypePostRequest struct {
	value *UserTypePostRequest
	isSet bool
}

func (v NullableUserTypePostRequest) Get() *UserTypePostRequest {
	return v.value
}

func (v *NullableUserTypePostRequest) Set(val *UserTypePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypePostRequest(val *UserTypePostRequest) *NullableUserTypePostRequest {
	return &NullableUserTypePostRequest{value: val, isSet: true}
}

func (v NullableUserTypePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


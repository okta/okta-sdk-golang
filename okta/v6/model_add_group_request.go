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
)

// AddGroupRequest struct for AddGroupRequest
type AddGroupRequest struct {
	Profile *OktaUserGroupProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddGroupRequest AddGroupRequest

// NewAddGroupRequest instantiates a new AddGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroupRequest() *AddGroupRequest {
	this := AddGroupRequest{}
	return &this
}

// NewAddGroupRequestWithDefaults instantiates a new AddGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroupRequestWithDefaults() *AddGroupRequest {
	this := AddGroupRequest{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AddGroupRequest) GetProfile() OktaUserGroupProfile {
	if o == nil || o.Profile == nil {
		var ret OktaUserGroupProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroupRequest) GetProfileOk() (*OktaUserGroupProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AddGroupRequest) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given OktaUserGroupProfile and assigns it to the Profile field.
func (o *AddGroupRequest) SetProfile(v OktaUserGroupProfile) {
	o.Profile = &v
}

func (o AddGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AddGroupRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAddGroupRequest := _AddGroupRequest{}

	err = json.Unmarshal(bytes, &varAddGroupRequest)
	if err == nil {
		*o = AddGroupRequest(varAddGroupRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAddGroupRequest struct {
	value *AddGroupRequest
	isSet bool
}

func (v NullableAddGroupRequest) Get() *AddGroupRequest {
	return v.value
}

func (v *NullableAddGroupRequest) Set(val *AddGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroupRequest(val *AddGroupRequest) *NullableAddGroupRequest {
	return &NullableAddGroupRequest{value: val, isSet: true}
}

func (v NullableAddGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


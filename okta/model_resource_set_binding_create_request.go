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

// ResourceSetBindingCreateRequest struct for ResourceSetBindingCreateRequest
type ResourceSetBindingCreateRequest struct {
	Members []string `json:"members,omitempty"`
	// Unique key for the role
	Role *string `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingCreateRequest ResourceSetBindingCreateRequest

// NewResourceSetBindingCreateRequest instantiates a new ResourceSetBindingCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingCreateRequest() *ResourceSetBindingCreateRequest {
	this := ResourceSetBindingCreateRequest{}
	return &this
}

// NewResourceSetBindingCreateRequestWithDefaults instantiates a new ResourceSetBindingCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingCreateRequestWithDefaults() *ResourceSetBindingCreateRequest {
	this := ResourceSetBindingCreateRequest{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ResourceSetBindingCreateRequest) GetMembers() []string {
	if o == nil || o.Members == nil {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingCreateRequest) GetMembersOk() ([]string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ResourceSetBindingCreateRequest) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *ResourceSetBindingCreateRequest) SetMembers(v []string) {
	o.Members = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ResourceSetBindingCreateRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingCreateRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ResourceSetBindingCreateRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ResourceSetBindingCreateRequest) SetRole(v string) {
	o.Role = &v
}

func (o ResourceSetBindingCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingCreateRequest := _ResourceSetBindingCreateRequest{}

	err = json.Unmarshal(bytes, &varResourceSetBindingCreateRequest)
	if err == nil {
		*o = ResourceSetBindingCreateRequest(varResourceSetBindingCreateRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "members")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingCreateRequest struct {
	value *ResourceSetBindingCreateRequest
	isSet bool
}

func (v NullableResourceSetBindingCreateRequest) Get() *ResourceSetBindingCreateRequest {
	return v.value
}

func (v *NullableResourceSetBindingCreateRequest) Set(val *ResourceSetBindingCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingCreateRequest(val *ResourceSetBindingCreateRequest) *NullableResourceSetBindingCreateRequest {
	return &NullableResourceSetBindingCreateRequest{value: val, isSet: true}
}

func (v NullableResourceSetBindingCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


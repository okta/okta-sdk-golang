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

// ResourceSetBindingAddMembersRequest struct for ResourceSetBindingAddMembersRequest
type ResourceSetBindingAddMembersRequest struct {
	Additions []string `json:"additions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingAddMembersRequest ResourceSetBindingAddMembersRequest

// NewResourceSetBindingAddMembersRequest instantiates a new ResourceSetBindingAddMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingAddMembersRequest() *ResourceSetBindingAddMembersRequest {
	this := ResourceSetBindingAddMembersRequest{}
	return &this
}

// NewResourceSetBindingAddMembersRequestWithDefaults instantiates a new ResourceSetBindingAddMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingAddMembersRequestWithDefaults() *ResourceSetBindingAddMembersRequest {
	this := ResourceSetBindingAddMembersRequest{}
	return &this
}

// GetAdditions returns the Additions field value if set, zero value otherwise.
func (o *ResourceSetBindingAddMembersRequest) GetAdditions() []string {
	if o == nil || o.Additions == nil {
		var ret []string
		return ret
	}
	return o.Additions
}

// GetAdditionsOk returns a tuple with the Additions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingAddMembersRequest) GetAdditionsOk() ([]string, bool) {
	if o == nil || o.Additions == nil {
		return nil, false
	}
	return o.Additions, true
}

// HasAdditions returns a boolean if a field has been set.
func (o *ResourceSetBindingAddMembersRequest) HasAdditions() bool {
	if o != nil && o.Additions != nil {
		return true
	}

	return false
}

// SetAdditions gets a reference to the given []string and assigns it to the Additions field.
func (o *ResourceSetBindingAddMembersRequest) SetAdditions(v []string) {
	o.Additions = v
}

func (o ResourceSetBindingAddMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Additions != nil {
		toSerialize["additions"] = o.Additions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingAddMembersRequest) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingAddMembersRequest := _ResourceSetBindingAddMembersRequest{}

	err = json.Unmarshal(bytes, &varResourceSetBindingAddMembersRequest)
	if err == nil {
		*o = ResourceSetBindingAddMembersRequest(varResourceSetBindingAddMembersRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "additions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingAddMembersRequest struct {
	value *ResourceSetBindingAddMembersRequest
	isSet bool
}

func (v NullableResourceSetBindingAddMembersRequest) Get() *ResourceSetBindingAddMembersRequest {
	return v.value
}

func (v *NullableResourceSetBindingAddMembersRequest) Set(val *ResourceSetBindingAddMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingAddMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingAddMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingAddMembersRequest(val *ResourceSetBindingAddMembersRequest) *NullableResourceSetBindingAddMembersRequest {
	return &NullableResourceSetBindingAddMembersRequest{value: val, isSet: true}
}

func (v NullableResourceSetBindingAddMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingAddMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


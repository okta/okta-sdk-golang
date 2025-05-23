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

// ResourceSetBindingMembers struct for ResourceSetBindingMembers
type ResourceSetBindingMembers struct {
	Members []ResourceSetBindingMember `json:"members,omitempty"`
	Links *ResourceSetBindingMembersLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingMembers ResourceSetBindingMembers

// NewResourceSetBindingMembers instantiates a new ResourceSetBindingMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingMembers() *ResourceSetBindingMembers {
	this := ResourceSetBindingMembers{}
	return &this
}

// NewResourceSetBindingMembersWithDefaults instantiates a new ResourceSetBindingMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingMembersWithDefaults() *ResourceSetBindingMembers {
	this := ResourceSetBindingMembers{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ResourceSetBindingMembers) GetMembers() []ResourceSetBindingMember {
	if o == nil || o.Members == nil {
		var ret []ResourceSetBindingMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembers) GetMembersOk() ([]ResourceSetBindingMember, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ResourceSetBindingMembers) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ResourceSetBindingMember and assigns it to the Members field.
func (o *ResourceSetBindingMembers) SetMembers(v []ResourceSetBindingMember) {
	o.Members = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetBindingMembers) GetLinks() ResourceSetBindingMembersLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetBindingMembersLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingMembers) GetLinksOk() (*ResourceSetBindingMembersLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetBindingMembers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetBindingMembersLinks and assigns it to the Links field.
func (o *ResourceSetBindingMembers) SetLinks(v ResourceSetBindingMembersLinks) {
	o.Links = &v
}

func (o ResourceSetBindingMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingMembers) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingMembers := _ResourceSetBindingMembers{}

	err = json.Unmarshal(bytes, &varResourceSetBindingMembers)
	if err == nil {
		*o = ResourceSetBindingMembers(varResourceSetBindingMembers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "members")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetBindingMembers struct {
	value *ResourceSetBindingMembers
	isSet bool
}

func (v NullableResourceSetBindingMembers) Get() *ResourceSetBindingMembers {
	return v.value
}

func (v *NullableResourceSetBindingMembers) Set(val *ResourceSetBindingMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingMembers(val *ResourceSetBindingMembers) *NullableResourceSetBindingMembers {
	return &NullableResourceSetBindingMembers{value: val, isSet: true}
}

func (v NullableResourceSetBindingMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


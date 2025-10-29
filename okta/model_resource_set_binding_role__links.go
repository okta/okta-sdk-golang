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

// checks if the ResourceSetBindingRoleLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetBindingRoleLinks{}

// ResourceSetBindingRoleLinks struct for ResourceSetBindingRoleLinks
type ResourceSetBindingRoleLinks struct {
	Self                 *HrefObjectSelfLink    `json:"self,omitempty"`
	Members              *HrefObjectMembersLink `json:"members,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingRoleLinks ResourceSetBindingRoleLinks

// NewResourceSetBindingRoleLinks instantiates a new ResourceSetBindingRoleLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingRoleLinks() *ResourceSetBindingRoleLinks {
	this := ResourceSetBindingRoleLinks{}
	return &this
}

// NewResourceSetBindingRoleLinksWithDefaults instantiates a new ResourceSetBindingRoleLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingRoleLinksWithDefaults() *ResourceSetBindingRoleLinks {
	this := ResourceSetBindingRoleLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingRoleLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRoleLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingRoleLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingRoleLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ResourceSetBindingRoleLinks) GetMembers() HrefObjectMembersLink {
	if o == nil || IsNil(o.Members) {
		var ret HrefObjectMembersLink
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRoleLinks) GetMembersOk() (*HrefObjectMembersLink, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ResourceSetBindingRoleLinks) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given HrefObjectMembersLink and assigns it to the Members field.
func (o *ResourceSetBindingRoleLinks) SetMembers(v HrefObjectMembersLink) {
	o.Members = &v
}

func (o ResourceSetBindingRoleLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetBindingRoleLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetBindingRoleLinks) UnmarshalJSON(data []byte) (err error) {
	varResourceSetBindingRoleLinks := _ResourceSetBindingRoleLinks{}

	err = json.Unmarshal(data, &varResourceSetBindingRoleLinks)

	if err != nil {
		return err
	}

	*o = ResourceSetBindingRoleLinks(varResourceSetBindingRoleLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetBindingRoleLinks struct {
	value *ResourceSetBindingRoleLinks
	isSet bool
}

func (v NullableResourceSetBindingRoleLinks) Get() *ResourceSetBindingRoleLinks {
	return v.value
}

func (v *NullableResourceSetBindingRoleLinks) Set(val *ResourceSetBindingRoleLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingRoleLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingRoleLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingRoleLinks(val *ResourceSetBindingRoleLinks) *NullableResourceSetBindingRoleLinks {
	return &NullableResourceSetBindingRoleLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingRoleLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingRoleLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

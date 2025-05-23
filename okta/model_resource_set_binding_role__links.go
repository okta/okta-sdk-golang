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

// ResourceSetBindingRoleLinks struct for ResourceSetBindingRoleLinks
type ResourceSetBindingRoleLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Members *HrefObject `json:"members,omitempty"`
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
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRoleLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingRoleLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingRoleLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ResourceSetBindingRoleLinks) GetMembers() HrefObject {
	if o == nil || o.Members == nil {
		var ret HrefObject
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingRoleLinks) GetMembersOk() (*HrefObject, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ResourceSetBindingRoleLinks) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given HrefObject and assigns it to the Members field.
func (o *ResourceSetBindingRoleLinks) SetMembers(v HrefObject) {
	o.Members = &v
}

func (o ResourceSetBindingRoleLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetBindingRoleLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetBindingRoleLinks := _ResourceSetBindingRoleLinks{}

	err = json.Unmarshal(bytes, &varResourceSetBindingRoleLinks)
	if err == nil {
		*o = ResourceSetBindingRoleLinks(varResourceSetBindingRoleLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


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

// RoleGovernanceResourcesLinks struct for RoleGovernanceResourcesLinks
type RoleGovernanceResourcesLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Next *HrefObject `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernanceResourcesLinks RoleGovernanceResourcesLinks

// NewRoleGovernanceResourcesLinks instantiates a new RoleGovernanceResourcesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernanceResourcesLinks() *RoleGovernanceResourcesLinks {
	this := RoleGovernanceResourcesLinks{}
	return &this
}

// NewRoleGovernanceResourcesLinksWithDefaults instantiates a new RoleGovernanceResourcesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceResourcesLinksWithDefaults() *RoleGovernanceResourcesLinks {
	this := RoleGovernanceResourcesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *RoleGovernanceResourcesLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResourcesLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *RoleGovernanceResourcesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *RoleGovernanceResourcesLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *RoleGovernanceResourcesLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceResourcesLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *RoleGovernanceResourcesLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *RoleGovernanceResourcesLinks) SetNext(v HrefObject) {
	o.Next = &v
}

func (o RoleGovernanceResourcesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleGovernanceResourcesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varRoleGovernanceResourcesLinks := _RoleGovernanceResourcesLinks{}

	err = json.Unmarshal(bytes, &varRoleGovernanceResourcesLinks)
	if err == nil {
		*o = RoleGovernanceResourcesLinks(varRoleGovernanceResourcesLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleGovernanceResourcesLinks struct {
	value *RoleGovernanceResourcesLinks
	isSet bool
}

func (v NullableRoleGovernanceResourcesLinks) Get() *RoleGovernanceResourcesLinks {
	return v.value
}

func (v *NullableRoleGovernanceResourcesLinks) Set(val *RoleGovernanceResourcesLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernanceResourcesLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernanceResourcesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernanceResourcesLinks(val *RoleGovernanceResourcesLinks) *NullableRoleGovernanceResourcesLinks {
	return &NullableRoleGovernanceResourcesLinks{value: val, isSet: true}
}

func (v NullableRoleGovernanceResourcesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernanceResourcesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


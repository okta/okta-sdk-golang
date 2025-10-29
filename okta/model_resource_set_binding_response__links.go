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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ResourceSetBindingResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetBindingResponseLinks{}

// ResourceSetBindingResponseLinks struct for ResourceSetBindingResponseLinks
type ResourceSetBindingResponseLinks struct {
	Self                 *HrefObjectSelfLink        `json:"self,omitempty"`
	ResourceSet          *HrefObjectResourceSetLink `json:"resource-set,omitempty"`
	Members              *HrefObjectMembersLink     `json:"members,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetBindingResponseLinks ResourceSetBindingResponseLinks

// NewResourceSetBindingResponseLinks instantiates a new ResourceSetBindingResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetBindingResponseLinks() *ResourceSetBindingResponseLinks {
	this := ResourceSetBindingResponseLinks{}
	return &this
}

// NewResourceSetBindingResponseLinksWithDefaults instantiates a new ResourceSetBindingResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetBindingResponseLinksWithDefaults() *ResourceSetBindingResponseLinks {
	this := ResourceSetBindingResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSetBindingResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetResourceSet() HrefObjectResourceSetLink {
	if o == nil || IsNil(o.ResourceSet) {
		var ret HrefObjectResourceSetLink
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool) {
	if o == nil || IsNil(o.ResourceSet) {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasResourceSet() bool {
	if o != nil && !IsNil(o.ResourceSet) {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObjectResourceSetLink and assigns it to the ResourceSet field.
func (o *ResourceSetBindingResponseLinks) SetResourceSet(v HrefObjectResourceSetLink) {
	o.ResourceSet = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ResourceSetBindingResponseLinks) GetMembers() HrefObjectMembersLink {
	if o == nil || IsNil(o.Members) {
		var ret HrefObjectMembersLink
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetBindingResponseLinks) GetMembersOk() (*HrefObjectMembersLink, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ResourceSetBindingResponseLinks) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given HrefObjectMembersLink and assigns it to the Members field.
func (o *ResourceSetBindingResponseLinks) SetMembers(v HrefObjectMembersLink) {
	o.Members = &v
}

func (o ResourceSetBindingResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetBindingResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.ResourceSet) {
		toSerialize["resource-set"] = o.ResourceSet
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetBindingResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varResourceSetBindingResponseLinks := _ResourceSetBindingResponseLinks{}

	err = json.Unmarshal(data, &varResourceSetBindingResponseLinks)

	if err != nil {
		return err
	}

	*o = ResourceSetBindingResponseLinks(varResourceSetBindingResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resource-set")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetBindingResponseLinks struct {
	value *ResourceSetBindingResponseLinks
	isSet bool
}

func (v NullableResourceSetBindingResponseLinks) Get() *ResourceSetBindingResponseLinks {
	return v.value
}

func (v *NullableResourceSetBindingResponseLinks) Set(val *ResourceSetBindingResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetBindingResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetBindingResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetBindingResponseLinks(val *ResourceSetBindingResponseLinks) *NullableResourceSetBindingResponseLinks {
	return &NullableResourceSetBindingResponseLinks{value: val, isSet: true}
}

func (v NullableResourceSetBindingResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetBindingResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

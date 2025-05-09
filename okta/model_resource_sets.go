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

// ResourceSets struct for ResourceSets
type ResourceSets struct {
	ResourceSets []ResourceSet `json:"resource-sets,omitempty"`
	Links *LinksNext `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSets ResourceSets

// NewResourceSets instantiates a new ResourceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSets() *ResourceSets {
	this := ResourceSets{}
	return &this
}

// NewResourceSetsWithDefaults instantiates a new ResourceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetsWithDefaults() *ResourceSets {
	this := ResourceSets{}
	return &this
}

// GetResourceSets returns the ResourceSets field value if set, zero value otherwise.
func (o *ResourceSets) GetResourceSets() []ResourceSet {
	if o == nil || o.ResourceSets == nil {
		var ret []ResourceSet
		return ret
	}
	return o.ResourceSets
}

// GetResourceSetsOk returns a tuple with the ResourceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSets) GetResourceSetsOk() ([]ResourceSet, bool) {
	if o == nil || o.ResourceSets == nil {
		return nil, false
	}
	return o.ResourceSets, true
}

// HasResourceSets returns a boolean if a field has been set.
func (o *ResourceSets) HasResourceSets() bool {
	if o != nil && o.ResourceSets != nil {
		return true
	}

	return false
}

// SetResourceSets gets a reference to the given []ResourceSet and assigns it to the ResourceSets field.
func (o *ResourceSets) SetResourceSets(v []ResourceSet) {
	o.ResourceSets = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSets) GetLinks() LinksNext {
	if o == nil || o.Links == nil {
		var ret LinksNext
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSets) GetLinksOk() (*LinksNext, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSets) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksNext and assigns it to the Links field.
func (o *ResourceSets) SetLinks(v LinksNext) {
	o.Links = &v
}

func (o ResourceSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceSets != nil {
		toSerialize["resource-sets"] = o.ResourceSets
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSets) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSets := _ResourceSets{}

	err = json.Unmarshal(bytes, &varResourceSets)
	if err == nil {
		*o = ResourceSets(varResourceSets)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resource-sets")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSets struct {
	value *ResourceSets
	isSet bool
}

func (v NullableResourceSets) Get() *ResourceSets {
	return v.value
}

func (v *NullableResourceSets) Set(val *ResourceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSets(val *ResourceSets) *NullableResourceSets {
	return &NullableResourceSets{value: val, isSet: true}
}

func (v NullableResourceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


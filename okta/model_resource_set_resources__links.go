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

// ResourceSetResourcesLinks struct for ResourceSetResourcesLinks
type ResourceSetResourcesLinks struct {
	Next *HrefObject `json:"next,omitempty"`
	ResourceSet *HrefObject `json:"resource-set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResourcesLinks ResourceSetResourcesLinks

// NewResourceSetResourcesLinks instantiates a new ResourceSetResourcesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResourcesLinks() *ResourceSetResourcesLinks {
	this := ResourceSetResourcesLinks{}
	return &this
}

// NewResourceSetResourcesLinksWithDefaults instantiates a new ResourceSetResourcesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourcesLinksWithDefaults() *ResourceSetResourcesLinks {
	this := ResourceSetResourcesLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResourceSetResourcesLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcesLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceSetResourcesLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *ResourceSetResourcesLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetResourceSet returns the ResourceSet field value if set, zero value otherwise.
func (o *ResourceSetResourcesLinks) GetResourceSet() HrefObject {
	if o == nil || o.ResourceSet == nil {
		var ret HrefObject
		return ret
	}
	return *o.ResourceSet
}

// GetResourceSetOk returns a tuple with the ResourceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcesLinks) GetResourceSetOk() (*HrefObject, bool) {
	if o == nil || o.ResourceSet == nil {
		return nil, false
	}
	return o.ResourceSet, true
}

// HasResourceSet returns a boolean if a field has been set.
func (o *ResourceSetResourcesLinks) HasResourceSet() bool {
	if o != nil && o.ResourceSet != nil {
		return true
	}

	return false
}

// SetResourceSet gets a reference to the given HrefObject and assigns it to the ResourceSet field.
func (o *ResourceSetResourcesLinks) SetResourceSet(v HrefObject) {
	o.ResourceSet = &v
}

func (o ResourceSetResourcesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.ResourceSet != nil {
		toSerialize["resource-set"] = o.ResourceSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetResourcesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetResourcesLinks := _ResourceSetResourcesLinks{}

	err = json.Unmarshal(bytes, &varResourceSetResourcesLinks)
	if err == nil {
		*o = ResourceSetResourcesLinks(varResourceSetResourcesLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "next")
		delete(additionalProperties, "resource-set")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetResourcesLinks struct {
	value *ResourceSetResourcesLinks
	isSet bool
}

func (v NullableResourceSetResourcesLinks) Get() *ResourceSetResourcesLinks {
	return v.value
}

func (v *NullableResourceSetResourcesLinks) Set(val *ResourceSetResourcesLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResourcesLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResourcesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResourcesLinks(val *ResourceSetResourcesLinks) *NullableResourceSetResourcesLinks {
	return &NullableResourceSetResourcesLinks{value: val, isSet: true}
}

func (v NullableResourceSetResourcesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResourcesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


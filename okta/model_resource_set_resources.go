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

// ResourceSetResources struct for ResourceSetResources
type ResourceSetResources struct {
	Resources []ResourceSetResource `json:"resources,omitempty"`
	Links *ResourceSetResourcesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResources ResourceSetResources

// NewResourceSetResources instantiates a new ResourceSetResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResources() *ResourceSetResources {
	this := ResourceSetResources{}
	return &this
}

// NewResourceSetResourcesWithDefaults instantiates a new ResourceSetResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourcesWithDefaults() *ResourceSetResources {
	this := ResourceSetResources{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceSetResources) GetResources() []ResourceSetResource {
	if o == nil || o.Resources == nil {
		var ret []ResourceSetResource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResources) GetResourcesOk() ([]ResourceSetResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceSetResources) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourceSetResource and assigns it to the Resources field.
func (o *ResourceSetResources) SetResources(v []ResourceSetResource) {
	o.Resources = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetResources) GetLinks() ResourceSetResourcesLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSetResourcesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResources) GetLinksOk() (*ResourceSetResourcesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetResources) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetResourcesLinks and assigns it to the Links field.
func (o *ResourceSetResources) SetLinks(v ResourceSetResourcesLinks) {
	o.Links = &v
}

func (o ResourceSetResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetResources) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetResources := _ResourceSetResources{}

	err = json.Unmarshal(bytes, &varResourceSetResources)
	if err == nil {
		*o = ResourceSetResources(varResourceSetResources)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resources")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSetResources struct {
	value *ResourceSetResources
	isSet bool
}

func (v NullableResourceSetResources) Get() *ResourceSetResources {
	return v.value
}

func (v *NullableResourceSetResources) Set(val *ResourceSetResources) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResources(val *ResourceSetResources) *NullableResourceSetResources {
	return &NullableResourceSetResources{value: val, isSet: true}
}

func (v NullableResourceSetResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


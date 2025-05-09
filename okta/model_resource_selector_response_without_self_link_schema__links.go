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

// ResourceSelectorResponseWithoutSelfLinkSchemaLinks struct for ResourceSelectorResponseWithoutSelfLinkSchemaLinks
type ResourceSelectorResponseWithoutSelfLinkSchemaLinks struct {
	Resources *HrefObject `json:"resources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorResponseWithoutSelfLinkSchemaLinks ResourceSelectorResponseWithoutSelfLinkSchemaLinks

// NewResourceSelectorResponseWithoutSelfLinkSchemaLinks instantiates a new ResourceSelectorResponseWithoutSelfLinkSchemaLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorResponseWithoutSelfLinkSchemaLinks() *ResourceSelectorResponseWithoutSelfLinkSchemaLinks {
	this := ResourceSelectorResponseWithoutSelfLinkSchemaLinks{}
	return &this
}

// NewResourceSelectorResponseWithoutSelfLinkSchemaLinksWithDefaults instantiates a new ResourceSelectorResponseWithoutSelfLinkSchemaLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorResponseWithoutSelfLinkSchemaLinksWithDefaults() *ResourceSelectorResponseWithoutSelfLinkSchemaLinks {
	this := ResourceSelectorResponseWithoutSelfLinkSchemaLinks{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) GetResources() HrefObject {
	if o == nil || o.Resources == nil {
		var ret HrefObject
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) GetResourcesOk() (*HrefObject, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given HrefObject and assigns it to the Resources field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) SetResources(v HrefObject) {
	o.Resources = &v
}

func (o ResourceSelectorResponseWithoutSelfLinkSchemaLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorResponseWithoutSelfLinkSchemaLinks := _ResourceSelectorResponseWithoutSelfLinkSchemaLinks{}

	err = json.Unmarshal(bytes, &varResourceSelectorResponseWithoutSelfLinkSchemaLinks)
	if err == nil {
		*o = ResourceSelectorResponseWithoutSelfLinkSchemaLinks(varResourceSelectorResponseWithoutSelfLinkSchemaLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks struct {
	value *ResourceSelectorResponseWithoutSelfLinkSchemaLinks
	isSet bool
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) Get() *ResourceSelectorResponseWithoutSelfLinkSchemaLinks {
	return v.value
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) Set(val *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorResponseWithoutSelfLinkSchemaLinks(val *ResourceSelectorResponseWithoutSelfLinkSchemaLinks) *NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks {
	return &NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks{value: val, isSet: true}
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchemaLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


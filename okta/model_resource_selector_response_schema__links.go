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

// ResourceSelectorResponseSchemaLinks struct for ResourceSelectorResponseSchemaLinks
type ResourceSelectorResponseSchemaLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Resources *HrefObject `json:"resources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorResponseSchemaLinks ResourceSelectorResponseSchemaLinks

// NewResourceSelectorResponseSchemaLinks instantiates a new ResourceSelectorResponseSchemaLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorResponseSchemaLinks() *ResourceSelectorResponseSchemaLinks {
	this := ResourceSelectorResponseSchemaLinks{}
	return &this
}

// NewResourceSelectorResponseSchemaLinksWithDefaults instantiates a new ResourceSelectorResponseSchemaLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorResponseSchemaLinksWithDefaults() *ResourceSelectorResponseSchemaLinks {
	this := ResourceSelectorResponseSchemaLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchemaLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchemaLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchemaLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ResourceSelectorResponseSchemaLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchemaLinks) GetResources() HrefObject {
	if o == nil || o.Resources == nil {
		var ret HrefObject
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchemaLinks) GetResourcesOk() (*HrefObject, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchemaLinks) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given HrefObject and assigns it to the Resources field.
func (o *ResourceSelectorResponseSchemaLinks) SetResources(v HrefObject) {
	o.Resources = &v
}

func (o ResourceSelectorResponseSchemaLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorResponseSchemaLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorResponseSchemaLinks := _ResourceSelectorResponseSchemaLinks{}

	err = json.Unmarshal(bytes, &varResourceSelectorResponseSchemaLinks)
	if err == nil {
		*o = ResourceSelectorResponseSchemaLinks(varResourceSelectorResponseSchemaLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorResponseSchemaLinks struct {
	value *ResourceSelectorResponseSchemaLinks
	isSet bool
}

func (v NullableResourceSelectorResponseSchemaLinks) Get() *ResourceSelectorResponseSchemaLinks {
	return v.value
}

func (v *NullableResourceSelectorResponseSchemaLinks) Set(val *ResourceSelectorResponseSchemaLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorResponseSchemaLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorResponseSchemaLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorResponseSchemaLinks(val *ResourceSelectorResponseSchemaLinks) *NullableResourceSelectorResponseSchemaLinks {
	return &NullableResourceSelectorResponseSchemaLinks{value: val, isSet: true}
}

func (v NullableResourceSelectorResponseSchemaLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorResponseSchemaLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


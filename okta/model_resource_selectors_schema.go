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

// ResourceSelectorsSchema struct for ResourceSelectorsSchema
type ResourceSelectorsSchema struct {
	ResourceSelectors []ResourceSelectorResponseWithoutSelfLinkSchema `json:"resourceSelectors,omitempty"`
	Links *LinksNext `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorsSchema ResourceSelectorsSchema

// NewResourceSelectorsSchema instantiates a new ResourceSelectorsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorsSchema() *ResourceSelectorsSchema {
	this := ResourceSelectorsSchema{}
	return &this
}

// NewResourceSelectorsSchemaWithDefaults instantiates a new ResourceSelectorsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorsSchemaWithDefaults() *ResourceSelectorsSchema {
	this := ResourceSelectorsSchema{}
	return &this
}

// GetResourceSelectors returns the ResourceSelectors field value if set, zero value otherwise.
func (o *ResourceSelectorsSchema) GetResourceSelectors() []ResourceSelectorResponseWithoutSelfLinkSchema {
	if o == nil || o.ResourceSelectors == nil {
		var ret []ResourceSelectorResponseWithoutSelfLinkSchema
		return ret
	}
	return o.ResourceSelectors
}

// GetResourceSelectorsOk returns a tuple with the ResourceSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorsSchema) GetResourceSelectorsOk() ([]ResourceSelectorResponseWithoutSelfLinkSchema, bool) {
	if o == nil || o.ResourceSelectors == nil {
		return nil, false
	}
	return o.ResourceSelectors, true
}

// HasResourceSelectors returns a boolean if a field has been set.
func (o *ResourceSelectorsSchema) HasResourceSelectors() bool {
	if o != nil && o.ResourceSelectors != nil {
		return true
	}

	return false
}

// SetResourceSelectors gets a reference to the given []ResourceSelectorResponseWithoutSelfLinkSchema and assigns it to the ResourceSelectors field.
func (o *ResourceSelectorsSchema) SetResourceSelectors(v []ResourceSelectorResponseWithoutSelfLinkSchema) {
	o.ResourceSelectors = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSelectorsSchema) GetLinks() LinksNext {
	if o == nil || o.Links == nil {
		var ret LinksNext
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorsSchema) GetLinksOk() (*LinksNext, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSelectorsSchema) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksNext and assigns it to the Links field.
func (o *ResourceSelectorsSchema) SetLinks(v LinksNext) {
	o.Links = &v
}

func (o ResourceSelectorsSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceSelectors != nil {
		toSerialize["resourceSelectors"] = o.ResourceSelectors
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorsSchema) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorsSchema := _ResourceSelectorsSchema{}

	err = json.Unmarshal(bytes, &varResourceSelectorsSchema)
	if err == nil {
		*o = ResourceSelectorsSchema(varResourceSelectorsSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceSelectors")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorsSchema struct {
	value *ResourceSelectorsSchema
	isSet bool
}

func (v NullableResourceSelectorsSchema) Get() *ResourceSelectorsSchema {
	return v.value
}

func (v *NullableResourceSelectorsSchema) Set(val *ResourceSelectorsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorsSchema(val *ResourceSelectorsSchema) *NullableResourceSelectorsSchema {
	return &NullableResourceSelectorsSchema{value: val, isSet: true}
}

func (v NullableResourceSelectorsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


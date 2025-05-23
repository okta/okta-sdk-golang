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

// ResourceSelectorResponseWithoutSelfLinkSchema struct for ResourceSelectorResponseWithoutSelfLinkSchema
type ResourceSelectorResponseWithoutSelfLinkSchema struct {
	// Description of the Resource Selector
	Description *string `json:"description,omitempty"`
	// Unique key for the Resource Selector
	Id *string `json:"id,omitempty"`
	// Name of the Resource Selector
	Name *string `json:"name,omitempty"`
	// An Okta resource name
	Orn *string `json:"orn,omitempty"`
	Links *ResourceSelectorResponseWithoutSelfLinkSchemaLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorResponseWithoutSelfLinkSchema ResourceSelectorResponseWithoutSelfLinkSchema

// NewResourceSelectorResponseWithoutSelfLinkSchema instantiates a new ResourceSelectorResponseWithoutSelfLinkSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorResponseWithoutSelfLinkSchema() *ResourceSelectorResponseWithoutSelfLinkSchema {
	this := ResourceSelectorResponseWithoutSelfLinkSchema{}
	return &this
}

// NewResourceSelectorResponseWithoutSelfLinkSchemaWithDefaults instantiates a new ResourceSelectorResponseWithoutSelfLinkSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorResponseWithoutSelfLinkSchemaWithDefaults() *ResourceSelectorResponseWithoutSelfLinkSchema {
	this := ResourceSelectorResponseWithoutSelfLinkSchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) SetName(v string) {
	o.Name = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetOrn() string {
	if o == nil || o.Orn == nil {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetOrnOk() (*string, bool) {
	if o == nil || o.Orn == nil {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) HasOrn() bool {
	if o != nil && o.Orn != nil {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) SetOrn(v string) {
	o.Orn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetLinks() ResourceSelectorResponseWithoutSelfLinkSchemaLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSelectorResponseWithoutSelfLinkSchemaLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) GetLinksOk() (*ResourceSelectorResponseWithoutSelfLinkSchemaLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSelectorResponseWithoutSelfLinkSchemaLinks and assigns it to the Links field.
func (o *ResourceSelectorResponseWithoutSelfLinkSchema) SetLinks(v ResourceSelectorResponseWithoutSelfLinkSchemaLinks) {
	o.Links = &v
}

func (o ResourceSelectorResponseWithoutSelfLinkSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Orn != nil {
		toSerialize["orn"] = o.Orn
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorResponseWithoutSelfLinkSchema) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorResponseWithoutSelfLinkSchema := _ResourceSelectorResponseWithoutSelfLinkSchema{}

	err = json.Unmarshal(bytes, &varResourceSelectorResponseWithoutSelfLinkSchema)
	if err == nil {
		*o = ResourceSelectorResponseWithoutSelfLinkSchema(varResourceSelectorResponseWithoutSelfLinkSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorResponseWithoutSelfLinkSchema struct {
	value *ResourceSelectorResponseWithoutSelfLinkSchema
	isSet bool
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchema) Get() *ResourceSelectorResponseWithoutSelfLinkSchema {
	return v.value
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchema) Set(val *ResourceSelectorResponseWithoutSelfLinkSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorResponseWithoutSelfLinkSchema(val *ResourceSelectorResponseWithoutSelfLinkSchema) *NullableResourceSelectorResponseWithoutSelfLinkSchema {
	return &NullableResourceSelectorResponseWithoutSelfLinkSchema{value: val, isSet: true}
}

func (v NullableResourceSelectorResponseWithoutSelfLinkSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorResponseWithoutSelfLinkSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


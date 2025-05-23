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

// ResourceSelectorResponseSchema struct for ResourceSelectorResponseSchema
type ResourceSelectorResponseSchema struct {
	// Description of the Resource Selector
	Description *string `json:"description,omitempty"`
	// Unique key for the Resource Selector
	Id *string `json:"id,omitempty"`
	// Name of the Resource Selector
	Name *string `json:"name,omitempty"`
	// An Okta resource name
	Orn *string `json:"orn,omitempty"`
	Links *ResourceSelectorResponseSchemaLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorResponseSchema ResourceSelectorResponseSchema

// NewResourceSelectorResponseSchema instantiates a new ResourceSelectorResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorResponseSchema() *ResourceSelectorResponseSchema {
	this := ResourceSelectorResponseSchema{}
	return &this
}

// NewResourceSelectorResponseSchemaWithDefaults instantiates a new ResourceSelectorResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorResponseSchemaWithDefaults() *ResourceSelectorResponseSchema {
	this := ResourceSelectorResponseSchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceSelectorResponseSchema) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceSelectorResponseSchema) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceSelectorResponseSchema) SetName(v string) {
	o.Name = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchema) GetOrn() string {
	if o == nil || o.Orn == nil {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchema) GetOrnOk() (*string, bool) {
	if o == nil || o.Orn == nil {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchema) HasOrn() bool {
	if o != nil && o.Orn != nil {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ResourceSelectorResponseSchema) SetOrn(v string) {
	o.Orn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSelectorResponseSchema) GetLinks() ResourceSelectorResponseSchemaLinks {
	if o == nil || o.Links == nil {
		var ret ResourceSelectorResponseSchemaLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorResponseSchema) GetLinksOk() (*ResourceSelectorResponseSchemaLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSelectorResponseSchema) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSelectorResponseSchemaLinks and assigns it to the Links field.
func (o *ResourceSelectorResponseSchema) SetLinks(v ResourceSelectorResponseSchemaLinks) {
	o.Links = &v
}

func (o ResourceSelectorResponseSchema) MarshalJSON() ([]byte, error) {
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

func (o *ResourceSelectorResponseSchema) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorResponseSchema := _ResourceSelectorResponseSchema{}

	err = json.Unmarshal(bytes, &varResourceSelectorResponseSchema)
	if err == nil {
		*o = ResourceSelectorResponseSchema(varResourceSelectorResponseSchema)
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

type NullableResourceSelectorResponseSchema struct {
	value *ResourceSelectorResponseSchema
	isSet bool
}

func (v NullableResourceSelectorResponseSchema) Get() *ResourceSelectorResponseSchema {
	return v.value
}

func (v *NullableResourceSelectorResponseSchema) Set(val *ResourceSelectorResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorResponseSchema(val *ResourceSelectorResponseSchema) *NullableResourceSelectorResponseSchema {
	return &NullableResourceSelectorResponseSchema{value: val, isSet: true}
}

func (v NullableResourceSelectorResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


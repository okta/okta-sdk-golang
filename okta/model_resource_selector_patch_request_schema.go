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

// ResourceSelectorPatchRequestSchema struct for ResourceSelectorPatchRequestSchema
type ResourceSelectorPatchRequestSchema struct {
	// Description of the Resource Selector
	Description *string `json:"description,omitempty"`
	// SCIM filter of the Resource Selector
	Filter *string `json:"filter,omitempty"`
	// Name of the Resource Selector
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorPatchRequestSchema ResourceSelectorPatchRequestSchema

// NewResourceSelectorPatchRequestSchema instantiates a new ResourceSelectorPatchRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorPatchRequestSchema() *ResourceSelectorPatchRequestSchema {
	this := ResourceSelectorPatchRequestSchema{}
	return &this
}

// NewResourceSelectorPatchRequestSchemaWithDefaults instantiates a new ResourceSelectorPatchRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorPatchRequestSchemaWithDefaults() *ResourceSelectorPatchRequestSchema {
	this := ResourceSelectorPatchRequestSchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceSelectorPatchRequestSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorPatchRequestSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceSelectorPatchRequestSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceSelectorPatchRequestSchema) SetDescription(v string) {
	o.Description = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ResourceSelectorPatchRequestSchema) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorPatchRequestSchema) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ResourceSelectorPatchRequestSchema) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ResourceSelectorPatchRequestSchema) SetFilter(v string) {
	o.Filter = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceSelectorPatchRequestSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorPatchRequestSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceSelectorPatchRequestSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceSelectorPatchRequestSchema) SetName(v string) {
	o.Name = &v
}

func (o ResourceSelectorPatchRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorPatchRequestSchema) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorPatchRequestSchema := _ResourceSelectorPatchRequestSchema{}

	err = json.Unmarshal(bytes, &varResourceSelectorPatchRequestSchema)
	if err == nil {
		*o = ResourceSelectorPatchRequestSchema(varResourceSelectorPatchRequestSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorPatchRequestSchema struct {
	value *ResourceSelectorPatchRequestSchema
	isSet bool
}

func (v NullableResourceSelectorPatchRequestSchema) Get() *ResourceSelectorPatchRequestSchema {
	return v.value
}

func (v *NullableResourceSelectorPatchRequestSchema) Set(val *ResourceSelectorPatchRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorPatchRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorPatchRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorPatchRequestSchema(val *ResourceSelectorPatchRequestSchema) *NullableResourceSelectorPatchRequestSchema {
	return &NullableResourceSelectorPatchRequestSchema{value: val, isSet: true}
}

func (v NullableResourceSelectorPatchRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorPatchRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


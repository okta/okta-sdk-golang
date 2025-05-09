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

// ResourceSelectorCreateRequestSchema struct for ResourceSelectorCreateRequestSchema
type ResourceSelectorCreateRequestSchema struct {
	// Description of the Resource Selector
	Description *string `json:"description,omitempty"`
	// SCIM filter of the Resource Selector
	Filter *string `json:"filter,omitempty"`
	// Name of the Resource Selector
	Name *string `json:"name,omitempty"`
	// Schema of the Resource Selector
	Schema *string `json:"schema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSelectorCreateRequestSchema ResourceSelectorCreateRequestSchema

// NewResourceSelectorCreateRequestSchema instantiates a new ResourceSelectorCreateRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSelectorCreateRequestSchema() *ResourceSelectorCreateRequestSchema {
	this := ResourceSelectorCreateRequestSchema{}
	return &this
}

// NewResourceSelectorCreateRequestSchemaWithDefaults instantiates a new ResourceSelectorCreateRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSelectorCreateRequestSchemaWithDefaults() *ResourceSelectorCreateRequestSchema {
	this := ResourceSelectorCreateRequestSchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceSelectorCreateRequestSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorCreateRequestSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceSelectorCreateRequestSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceSelectorCreateRequestSchema) SetDescription(v string) {
	o.Description = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ResourceSelectorCreateRequestSchema) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorCreateRequestSchema) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ResourceSelectorCreateRequestSchema) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *ResourceSelectorCreateRequestSchema) SetFilter(v string) {
	o.Filter = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceSelectorCreateRequestSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorCreateRequestSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceSelectorCreateRequestSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceSelectorCreateRequestSchema) SetName(v string) {
	o.Name = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ResourceSelectorCreateRequestSchema) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSelectorCreateRequestSchema) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ResourceSelectorCreateRequestSchema) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ResourceSelectorCreateRequestSchema) SetSchema(v string) {
	o.Schema = &v
}

func (o ResourceSelectorCreateRequestSchema) MarshalJSON() ([]byte, error) {
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
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSelectorCreateRequestSchema) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSelectorCreateRequestSchema := _ResourceSelectorCreateRequestSchema{}

	err = json.Unmarshal(bytes, &varResourceSelectorCreateRequestSchema)
	if err == nil {
		*o = ResourceSelectorCreateRequestSchema(varResourceSelectorCreateRequestSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "name")
		delete(additionalProperties, "schema")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceSelectorCreateRequestSchema struct {
	value *ResourceSelectorCreateRequestSchema
	isSet bool
}

func (v NullableResourceSelectorCreateRequestSchema) Get() *ResourceSelectorCreateRequestSchema {
	return v.value
}

func (v *NullableResourceSelectorCreateRequestSchema) Set(val *ResourceSelectorCreateRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSelectorCreateRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSelectorCreateRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSelectorCreateRequestSchema(val *ResourceSelectorCreateRequestSchema) *NullableResourceSelectorCreateRequestSchema {
	return &NullableResourceSelectorCreateRequestSchema{value: val, isSet: true}
}

func (v NullableResourceSelectorCreateRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSelectorCreateRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


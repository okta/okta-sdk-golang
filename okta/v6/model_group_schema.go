/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GroupSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupSchema{}

// GroupSchema struct for GroupSchema
type GroupSchema struct {
	// JSON schema version identifier
	Schema *string `json:"$schema,omitempty"`
	// Timestamp when the schema was created
	Created     *string                 `json:"created,omitempty"`
	Definitions *GroupSchemaDefinitions `json:"definitions,omitempty"`
	// Description for the schema
	Description *string `json:"description,omitempty"`
	// URI of group schema
	Id *string `json:"id,omitempty"`
	// Timestamp when the schema was last updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// Name of the schema
	Name       *string               `json:"name,omitempty"`
	Properties *UserSchemaProperties `json:"properties,omitempty"`
	// User-defined display name for the schema
	Title *string `json:"title,omitempty"`
	// Type of [root schema](https://tools.ietf.org/html/draft-zyp-json-schema-04#section-3.4)
	Type                 *string    `json:"type,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupSchema GroupSchema

// NewGroupSchema instantiates a new GroupSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSchema() *GroupSchema {
	this := GroupSchema{}
	return &this
}

// NewGroupSchemaWithDefaults instantiates a new GroupSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSchemaWithDefaults() *GroupSchema {
	this := GroupSchema{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GroupSchema) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GroupSchema) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *GroupSchema) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GroupSchema) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GroupSchema) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *GroupSchema) SetCreated(v string) {
	o.Created = &v
}

// GetDefinitions returns the Definitions field value if set, zero value otherwise.
func (o *GroupSchema) GetDefinitions() GroupSchemaDefinitions {
	if o == nil || IsNil(o.Definitions) {
		var ret GroupSchemaDefinitions
		return ret
	}
	return *o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetDefinitionsOk() (*GroupSchemaDefinitions, bool) {
	if o == nil || IsNil(o.Definitions) {
		return nil, false
	}
	return o.Definitions, true
}

// HasDefinitions returns a boolean if a field has been set.
func (o *GroupSchema) HasDefinitions() bool {
	if o != nil && !IsNil(o.Definitions) {
		return true
	}

	return false
}

// SetDefinitions gets a reference to the given GroupSchemaDefinitions and assigns it to the Definitions field.
func (o *GroupSchema) SetDefinitions(v GroupSchemaDefinitions) {
	o.Definitions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupSchema) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupSchema) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GroupSchema) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GroupSchema) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *GroupSchema) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupSchema) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GroupSchema) GetProperties() UserSchemaProperties {
	if o == nil || IsNil(o.Properties) {
		var ret UserSchemaProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetPropertiesOk() (*UserSchemaProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GroupSchema) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given UserSchemaProperties and assigns it to the Properties field.
func (o *GroupSchema) SetProperties(v UserSchemaProperties) {
	o.Properties = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GroupSchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GroupSchema) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupSchema) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GroupSchema) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSchema) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GroupSchema) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *GroupSchema) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o GroupSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Definitions) {
		toSerialize["definitions"] = o.Definitions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupSchema) UnmarshalJSON(data []byte) (err error) {
	varGroupSchema := _GroupSchema{}

	err = json.Unmarshal(data, &varGroupSchema)

	if err != nil {
		return err
	}

	*o = GroupSchema(varGroupSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "$schema")
		delete(additionalProperties, "created")
		delete(additionalProperties, "definitions")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupSchema struct {
	value *GroupSchema
	isSet bool
}

func (v NullableGroupSchema) Get() *GroupSchema {
	return v.value
}

func (v *NullableGroupSchema) Set(val *GroupSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupSchema(val *GroupSchema) *NullableGroupSchema {
	return &NullableGroupSchema{value: val, isSet: true}
}

func (v NullableGroupSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// UserSchema struct for UserSchema
type UserSchema struct {
	Schema *string `json:"$schema,omitempty"`
	Created *string `json:"created,omitempty"`
	Definitions *UserSchemaDefinitions `json:"definitions,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Properties *UserSchemaProperties `json:"properties,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchema UserSchema

// NewUserSchema instantiates a new UserSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchema() *UserSchema {
	this := UserSchema{}
	return &this
}

// NewUserSchemaWithDefaults instantiates a new UserSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaWithDefaults() *UserSchema {
	this := UserSchema{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *UserSchema) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *UserSchema) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *UserSchema) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserSchema) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserSchema) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *UserSchema) SetCreated(v string) {
	o.Created = &v
}

// GetDefinitions returns the Definitions field value if set, zero value otherwise.
func (o *UserSchema) GetDefinitions() UserSchemaDefinitions {
	if o == nil || o.Definitions == nil {
		var ret UserSchemaDefinitions
		return ret
	}
	return *o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetDefinitionsOk() (*UserSchemaDefinitions, bool) {
	if o == nil || o.Definitions == nil {
		return nil, false
	}
	return o.Definitions, true
}

// HasDefinitions returns a boolean if a field has been set.
func (o *UserSchema) HasDefinitions() bool {
	if o != nil && o.Definitions != nil {
		return true
	}

	return false
}

// SetDefinitions gets a reference to the given UserSchemaDefinitions and assigns it to the Definitions field.
func (o *UserSchema) SetDefinitions(v UserSchemaDefinitions) {
	o.Definitions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserSchema) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserSchema) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserSchema) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *UserSchema) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSchema) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserSchema) GetProperties() UserSchemaProperties {
	if o == nil || o.Properties == nil {
		var ret UserSchemaProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetPropertiesOk() (*UserSchemaProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserSchema) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given UserSchemaProperties and assigns it to the Properties field.
func (o *UserSchema) SetProperties(v UserSchemaProperties) {
	o.Properties = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserSchema) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchema) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserSchema) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSchema) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchema) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSchema) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserSchema) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserSchema) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *UserSchema) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o UserSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Definitions != nil {
		toSerialize["definitions"] = o.Definitions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchema) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchema := _UserSchema{}

	err = json.Unmarshal(bytes, &varUserSchema)
	if err == nil {
		*o = UserSchema(varUserSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "$schema")
		delete(additionalProperties, "created")
		delete(additionalProperties, "definitions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchema struct {
	value *UserSchema
	isSet bool
}

func (v NullableUserSchema) Get() *UserSchema {
	return v.value
}

func (v *NullableUserSchema) Set(val *UserSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchema(val *UserSchema) *NullableUserSchema {
	return &NullableUserSchema{value: val, isSet: true}
}

func (v NullableUserSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


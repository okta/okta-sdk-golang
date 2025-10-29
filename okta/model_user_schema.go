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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSchema{}

// UserSchema struct for UserSchema
type UserSchema struct {
	// JSON schema version identifier
	Schema *string `json:"$schema,omitempty"`
	// Timestamp when the schema was created
	Created *string `json:"created,omitempty"`
	// User profile subschemas  The profile object for a user is defined by a composite schema of base and custom properties using a JSON path to reference subschemas. The `#base` properties are defined and versioned by Okta, while `#custom` properties are extensible. Custom property names for the profile object must be unique and can't conflict with a property name defined in the `#base` subschema.
	Definitions *UserSchemaDefinitions `json:"definitions,omitempty"`
	// URI of user schema
	Id *string `json:"id,omitempty"`
	// Timestamp when the schema was last updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// Name of the schema
	Name *string `json:"name,omitempty"`
	// User Object Properties
	Properties *UserSchemaProperties `json:"properties,omitempty"`
	// User-defined display name for the schema
	Title *string `json:"title,omitempty"`
	// Type of [root schema](https://tools.ietf.org/html/draft-zyp-json-schema-04#section-3.4)
	Type                 *string    `json:"type,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
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
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *UserSchema) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
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
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserSchema) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
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
	if o == nil || IsNil(o.Definitions) {
		var ret UserSchemaDefinitions
		return ret
	}
	return *o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetDefinitionsOk() (*UserSchemaDefinitions, bool) {
	if o == nil || IsNil(o.Definitions) {
		return nil, false
	}
	return o.Definitions, true
}

// HasDefinitions returns a boolean if a field has been set.
func (o *UserSchema) HasDefinitions() bool {
	if o != nil && !IsNil(o.Definitions) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserSchema) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.Properties) {
		var ret UserSchemaProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetPropertiesOk() (*UserSchemaProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserSchema) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
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
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchema) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserSchema) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *UserSchema) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o UserSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSchema) ToMap() (map[string]interface{}, error) {
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

func (o *UserSchema) UnmarshalJSON(data []byte) (err error) {
	varUserSchema := _UserSchema{}

	err = json.Unmarshal(data, &varUserSchema)

	if err != nil {
		return err
	}

	*o = UserSchema(varUserSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

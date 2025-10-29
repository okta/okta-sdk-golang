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

// checks if the LogStreamSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStreamSchema{}

// LogStreamSchema struct for LogStreamSchema
type LogStreamSchema struct {
	// JSON schema version identifier
	Schema *string `json:"$schema,omitempty"`
	// A collection of error messages for individual properties in the schema. Okta implements a subset of [ajv-errors](https://github.com/ajv-validator/ajv-errors).
	ErrorMessage map[string]interface{} `json:"errorMessage,omitempty"`
	// URI of log stream schema
	Id *string `json:"id,omitempty"`
	// Non-empty array of valid JSON schemas.  Okta only supports `oneOf` for specifying display names for an `enum`. Each schema has the following format:  ``` {   \"const\": \"enumValue\",   \"title\": \"display name\" } ```
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	// For `string` log stream schema property type, specifies the regular expression used to validate the property
	Pattern *string `json:"pattern,omitempty"`
	// log stream schema properties object
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Required properties for this log stream schema object
	Required []string `json:"required,omitempty"`
	// Name of the log streaming integration
	Title *string `json:"title,omitempty"`
	// Type of log stream schema property
	Type                 *string    `json:"type,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSchema LogStreamSchema

// NewLogStreamSchema instantiates a new LogStreamSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSchema() *LogStreamSchema {
	this := LogStreamSchema{}
	return &this
}

// NewLogStreamSchemaWithDefaults instantiates a new LogStreamSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSchemaWithDefaults() *LogStreamSchema {
	this := LogStreamSchema{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *LogStreamSchema) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *LogStreamSchema) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *LogStreamSchema) SetSchema(v string) {
	o.Schema = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *LogStreamSchema) GetErrorMessage() map[string]interface{} {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetErrorMessageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return map[string]interface{}{}, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *LogStreamSchema) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given map[string]interface{} and assigns it to the ErrorMessage field.
func (o *LogStreamSchema) SetErrorMessage(v map[string]interface{}) {
	o.ErrorMessage = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogStreamSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogStreamSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogStreamSchema) SetId(v string) {
	o.Id = &v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogStreamSchema) GetOneOf() []UserSchemaAttributeEnum {
	if o == nil {
		var ret []UserSchemaAttributeEnum
		return ret
	}
	return o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogStreamSchema) GetOneOfOk() ([]UserSchemaAttributeEnum, bool) {
	if o == nil || IsNil(o.OneOf) {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *LogStreamSchema) HasOneOf() bool {
	if o != nil && !IsNil(o.OneOf) {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given []UserSchemaAttributeEnum and assigns it to the OneOf field.
func (o *LogStreamSchema) SetOneOf(v []UserSchemaAttributeEnum) {
	o.OneOf = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *LogStreamSchema) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *LogStreamSchema) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *LogStreamSchema) SetPattern(v string) {
	o.Pattern = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *LogStreamSchema) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *LogStreamSchema) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *LogStreamSchema) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *LogStreamSchema) GetRequired() []string {
	if o == nil || IsNil(o.Required) {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *LogStreamSchema) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *LogStreamSchema) SetRequired(v []string) {
	o.Required = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LogStreamSchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LogStreamSchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LogStreamSchema) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogStreamSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogStreamSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogStreamSchema) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LogStreamSchema) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LogStreamSchema) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *LogStreamSchema) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o LogStreamSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStreamSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.OneOf != nil {
		toSerialize["oneOf"] = o.OneOf
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
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

func (o *LogStreamSchema) UnmarshalJSON(data []byte) (err error) {
	varLogStreamSchema := _LogStreamSchema{}

	err = json.Unmarshal(data, &varLogStreamSchema)

	if err != nil {
		return err
	}

	*o = LogStreamSchema(varLogStreamSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "$schema")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "id")
		delete(additionalProperties, "oneOf")
		delete(additionalProperties, "pattern")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "required")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogStreamSchema struct {
	value *LogStreamSchema
	isSet bool
}

func (v NullableLogStreamSchema) Get() *LogStreamSchema {
	return v.value
}

func (v *NullableLogStreamSchema) Set(val *LogStreamSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSchema(val *LogStreamSchema) *NullableLogStreamSchema {
	return &NullableLogStreamSchema{value: val, isSet: true}
}

func (v NullableLogStreamSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

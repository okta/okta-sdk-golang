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

// LogStreamSchema struct for LogStreamSchema
type LogStreamSchema struct {
	Schema *string `json:"$schema,omitempty"`
	Created *string `json:"created,omitempty"`
	ErrorMessage map[string]interface{} `json:"errorMessage,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required []string `json:"required,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
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
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *LogStreamSchema) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *LogStreamSchema) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *LogStreamSchema) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *LogStreamSchema) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *LogStreamSchema) SetCreated(v string) {
	o.Created = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *LogStreamSchema) GetErrorMessage() map[string]interface{} {
	if o == nil || o.ErrorMessage == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetErrorMessageOk() (map[string]interface{}, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *LogStreamSchema) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
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
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogStreamSchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogStreamSchema) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *LogStreamSchema) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *LogStreamSchema) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *LogStreamSchema) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogStreamSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogStreamSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogStreamSchema) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *LogStreamSchema) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *LogStreamSchema) HasProperties() bool {
	if o != nil && o.Properties != nil {
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
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetRequiredOk() ([]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *LogStreamSchema) HasRequired() bool {
	if o != nil && o.Required != nil {
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
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LogStreamSchema) HasTitle() bool {
	if o != nil && o.Title != nil {
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
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogStreamSchema) HasType() bool {
	if o != nil && o.Type != nil {
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
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamSchema) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LogStreamSchema) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *LogStreamSchema) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o LogStreamSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
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
	if o.Required != nil {
		toSerialize["required"] = o.Required
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

func (o *LogStreamSchema) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSchema := _LogStreamSchema{}

	err = json.Unmarshal(bytes, &varLogStreamSchema)
	if err == nil {
		*o = LogStreamSchema(varLogStreamSchema)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "$schema")
		delete(additionalProperties, "created")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "required")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


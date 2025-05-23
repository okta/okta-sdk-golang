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
	"time"
)

// UISchemasResponseObject struct for UISchemasResponseObject
type UISchemasResponseObject struct {
	// Timestamp when the UI Schema was created (ISO-86001)
	Created time.Time `json:"created"`
	// Unique identifier for the UI Schema
	Id string `json:"id"`
	// Timestamp when the UI Schema was last modified (ISO-86001)
	LastUpdated time.Time `json:"lastUpdated"`
	UiSchema UISchemaObject `json:"uiSchema"`
	Links LinksSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _UISchemasResponseObject UISchemasResponseObject

// NewUISchemasResponseObject instantiates a new UISchemasResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUISchemasResponseObject(created time.Time, id string, lastUpdated time.Time, uiSchema UISchemaObject, links LinksSelf) *UISchemasResponseObject {
	this := UISchemasResponseObject{}
	this.Created = created
	this.Id = id
	this.LastUpdated = lastUpdated
	this.UiSchema = uiSchema
	this.Links = links
	return &this
}

// NewUISchemasResponseObjectWithDefaults instantiates a new UISchemasResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUISchemasResponseObjectWithDefaults() *UISchemasResponseObject {
	this := UISchemasResponseObject{}
	return &this
}

// GetCreated returns the Created field value
func (o *UISchemasResponseObject) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UISchemasResponseObject) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UISchemasResponseObject) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *UISchemasResponseObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UISchemasResponseObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UISchemasResponseObject) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *UISchemasResponseObject) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *UISchemasResponseObject) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *UISchemasResponseObject) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetUiSchema returns the UiSchema field value
func (o *UISchemasResponseObject) GetUiSchema() UISchemaObject {
	if o == nil {
		var ret UISchemaObject
		return ret
	}

	return o.UiSchema
}

// GetUiSchemaOk returns a tuple with the UiSchema field value
// and a boolean to check if the value has been set.
func (o *UISchemasResponseObject) GetUiSchemaOk() (*UISchemaObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiSchema, true
}

// SetUiSchema sets field value
func (o *UISchemasResponseObject) SetUiSchema(v UISchemaObject) {
	o.UiSchema = v
}

// GetLinks returns the Links field value
func (o *UISchemasResponseObject) GetLinks() LinksSelf {
	if o == nil {
		var ret LinksSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UISchemasResponseObject) GetLinksOk() (*LinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UISchemasResponseObject) SetLinks(v LinksSelf) {
	o.Links = v
}

func (o UISchemasResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["uiSchema"] = o.UiSchema
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UISchemasResponseObject) UnmarshalJSON(bytes []byte) (err error) {
	varUISchemasResponseObject := _UISchemasResponseObject{}

	err = json.Unmarshal(bytes, &varUISchemasResponseObject)
	if err == nil {
		*o = UISchemasResponseObject(varUISchemasResponseObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "uiSchema")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUISchemasResponseObject struct {
	value *UISchemasResponseObject
	isSet bool
}

func (v NullableUISchemasResponseObject) Get() *UISchemasResponseObject {
	return v.value
}

func (v *NullableUISchemasResponseObject) Set(val *UISchemasResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUISchemasResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUISchemasResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUISchemasResponseObject(val *UISchemasResponseObject) *NullableUISchemasResponseObject {
	return &NullableUISchemasResponseObject{value: val, isSet: true}
}

func (v NullableUISchemasResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUISchemasResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


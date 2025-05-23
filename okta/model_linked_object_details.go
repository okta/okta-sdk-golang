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

// LinkedObjectDetails struct for LinkedObjectDetails
type LinkedObjectDetails struct {
	// Description of the `primary` or the `associated` relationship
	Description *string `json:"description,omitempty"`
	// API name of the `primary` or the `associated` link
	Name string `json:"name"`
	// Display name of the `primary` or the `associated` link
	Title string `json:"title"`
	// The object type for this relationship
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _LinkedObjectDetails LinkedObjectDetails

// NewLinkedObjectDetails instantiates a new LinkedObjectDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedObjectDetails(name string, title string, type_ string) *LinkedObjectDetails {
	this := LinkedObjectDetails{}
	this.Name = name
	this.Title = title
	this.Type = type_
	return &this
}

// NewLinkedObjectDetailsWithDefaults instantiates a new LinkedObjectDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedObjectDetailsWithDefaults() *LinkedObjectDetails {
	this := LinkedObjectDetails{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LinkedObjectDetails) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObjectDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LinkedObjectDetails) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LinkedObjectDetails) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *LinkedObjectDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkedObjectDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkedObjectDetails) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *LinkedObjectDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *LinkedObjectDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *LinkedObjectDetails) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *LinkedObjectDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinkedObjectDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkedObjectDetails) SetType(v string) {
	o.Type = v
}

func (o LinkedObjectDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkedObjectDetails) UnmarshalJSON(bytes []byte) (err error) {
	varLinkedObjectDetails := _LinkedObjectDetails{}

	err = json.Unmarshal(bytes, &varLinkedObjectDetails)
	if err == nil {
		*o = LinkedObjectDetails(varLinkedObjectDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinkedObjectDetails struct {
	value *LinkedObjectDetails
	isSet bool
}

func (v NullableLinkedObjectDetails) Get() *LinkedObjectDetails {
	return v.value
}

func (v *NullableLinkedObjectDetails) Set(val *LinkedObjectDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedObjectDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedObjectDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedObjectDetails(val *LinkedObjectDetails) *NullableLinkedObjectDetails {
	return &NullableLinkedObjectDetails{value: val, isSet: true}
}

func (v NullableLinkedObjectDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedObjectDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


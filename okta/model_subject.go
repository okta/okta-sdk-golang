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

// Subject struct for Subject
type Subject struct {
	// The user identifier
	Format *string `json:"format,omitempty"`
	// ID of the user
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Subject Subject

// NewSubject instantiates a new Subject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubject() *Subject {
	this := Subject{}
	return &this
}

// NewSubjectWithDefaults instantiates a new Subject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectWithDefaults() *Subject {
	this := Subject{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *Subject) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subject) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *Subject) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *Subject) SetFormat(v string) {
	o.Format = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subject) SetId(v string) {
	o.Id = &v
}

func (o Subject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Subject) UnmarshalJSON(bytes []byte) (err error) {
	varSubject := _Subject{}

	err = json.Unmarshal(bytes, &varSubject)
	if err == nil {
		*o = Subject(varSubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSubject struct {
	value *Subject
	isSet bool
}

func (v NullableSubject) Get() *Subject {
	return v.value
}

func (v *NullableSubject) Set(val *Subject) {
	v.value = val
	v.isSet = true
}

func (v NullableSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubject(val *Subject) *NullableSubject {
	return &NullableSubject{value: val, isSet: true}
}

func (v NullableSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the AppResourceHrefObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppResourceHrefObject{}

// AppResourceHrefObject struct for AppResourceHrefObject
type AppResourceHrefObject struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	// Link name
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppResourceHrefObject AppResourceHrefObject

// NewAppResourceHrefObject instantiates a new AppResourceHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppResourceHrefObject() *AppResourceHrefObject {
	this := AppResourceHrefObject{}
	return &this
}

// NewAppResourceHrefObjectWithDefaults instantiates a new AppResourceHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppResourceHrefObjectWithDefaults() *AppResourceHrefObject {
	this := AppResourceHrefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AppResourceHrefObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppResourceHrefObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AppResourceHrefObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AppResourceHrefObject) SetHref(v string) {
	o.Href = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AppResourceHrefObject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppResourceHrefObject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AppResourceHrefObject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AppResourceHrefObject) SetTitle(v string) {
	o.Title = &v
}

func (o AppResourceHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppResourceHrefObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppResourceHrefObject) UnmarshalJSON(data []byte) (err error) {
	varAppResourceHrefObject := _AppResourceHrefObject{}

	err = json.Unmarshal(data, &varAppResourceHrefObject)

	if err != nil {
		return err
	}

	*o = AppResourceHrefObject(varAppResourceHrefObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppResourceHrefObject struct {
	value *AppResourceHrefObject
	isSet bool
}

func (v NullableAppResourceHrefObject) Get() *AppResourceHrefObject {
	return v.value
}

func (v *NullableAppResourceHrefObject) Set(val *AppResourceHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAppResourceHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAppResourceHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppResourceHrefObject(val *AppResourceHrefObject) *NullableAppResourceHrefObject {
	return &NullableAppResourceHrefObject{value: val, isSet: true}
}

func (v NullableAppResourceHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppResourceHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

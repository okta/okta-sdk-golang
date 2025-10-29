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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserResourceHrefObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResourceHrefObject{}

// UserResourceHrefObject struct for UserResourceHrefObject
type UserResourceHrefObject struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	// Link name
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserResourceHrefObject UserResourceHrefObject

// NewUserResourceHrefObject instantiates a new UserResourceHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResourceHrefObject() *UserResourceHrefObject {
	this := UserResourceHrefObject{}
	return &this
}

// NewUserResourceHrefObjectWithDefaults instantiates a new UserResourceHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResourceHrefObjectWithDefaults() *UserResourceHrefObject {
	this := UserResourceHrefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UserResourceHrefObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceHrefObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UserResourceHrefObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UserResourceHrefObject) SetHref(v string) {
	o.Href = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserResourceHrefObject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourceHrefObject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserResourceHrefObject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserResourceHrefObject) SetTitle(v string) {
	o.Title = &v
}

func (o UserResourceHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResourceHrefObject) ToMap() (map[string]interface{}, error) {
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

func (o *UserResourceHrefObject) UnmarshalJSON(data []byte) (err error) {
	varUserResourceHrefObject := _UserResourceHrefObject{}

	err = json.Unmarshal(data, &varUserResourceHrefObject)

	if err != nil {
		return err
	}

	*o = UserResourceHrefObject(varUserResourceHrefObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserResourceHrefObject struct {
	value *UserResourceHrefObject
	isSet bool
}

func (v NullableUserResourceHrefObject) Get() *UserResourceHrefObject {
	return v.value
}

func (v *NullableUserResourceHrefObject) Set(val *UserResourceHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResourceHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResourceHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResourceHrefObject(val *UserResourceHrefObject) *NullableUserResourceHrefObject {
	return &NullableUserResourceHrefObject{value: val, isSet: true}
}

func (v NullableUserResourceHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResourceHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

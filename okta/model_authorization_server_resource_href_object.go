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

// checks if the AuthorizationServerResourceHrefObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServerResourceHrefObject{}

// AuthorizationServerResourceHrefObject struct for AuthorizationServerResourceHrefObject
type AuthorizationServerResourceHrefObject struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	// Link name
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerResourceHrefObject AuthorizationServerResourceHrefObject

// NewAuthorizationServerResourceHrefObject instantiates a new AuthorizationServerResourceHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerResourceHrefObject() *AuthorizationServerResourceHrefObject {
	this := AuthorizationServerResourceHrefObject{}
	return &this
}

// NewAuthorizationServerResourceHrefObjectWithDefaults instantiates a new AuthorizationServerResourceHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerResourceHrefObjectWithDefaults() *AuthorizationServerResourceHrefObject {
	this := AuthorizationServerResourceHrefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AuthorizationServerResourceHrefObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerResourceHrefObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AuthorizationServerResourceHrefObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AuthorizationServerResourceHrefObject) SetHref(v string) {
	o.Href = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AuthorizationServerResourceHrefObject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerResourceHrefObject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AuthorizationServerResourceHrefObject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AuthorizationServerResourceHrefObject) SetTitle(v string) {
	o.Title = &v
}

func (o AuthorizationServerResourceHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServerResourceHrefObject) ToMap() (map[string]interface{}, error) {
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

func (o *AuthorizationServerResourceHrefObject) UnmarshalJSON(data []byte) (err error) {
	varAuthorizationServerResourceHrefObject := _AuthorizationServerResourceHrefObject{}

	err = json.Unmarshal(data, &varAuthorizationServerResourceHrefObject)

	if err != nil {
		return err
	}

	*o = AuthorizationServerResourceHrefObject(varAuthorizationServerResourceHrefObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationServerResourceHrefObject struct {
	value *AuthorizationServerResourceHrefObject
	isSet bool
}

func (v NullableAuthorizationServerResourceHrefObject) Get() *AuthorizationServerResourceHrefObject {
	return v.value
}

func (v *NullableAuthorizationServerResourceHrefObject) Set(val *AuthorizationServerResourceHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerResourceHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerResourceHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerResourceHrefObject(val *AuthorizationServerResourceHrefObject) *NullableAuthorizationServerResourceHrefObject {
	return &NullableAuthorizationServerResourceHrefObject{value: val, isSet: true}
}

func (v NullableAuthorizationServerResourceHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerResourceHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

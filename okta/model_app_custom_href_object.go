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

// AppCustomHrefObject struct for AppCustomHrefObject
type AppCustomHrefObject struct {
	Hints *AppCustomHrefObjectHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Title *string `json:"title,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppCustomHrefObject AppCustomHrefObject

// NewAppCustomHrefObject instantiates a new AppCustomHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomHrefObject(href string) *AppCustomHrefObject {
	this := AppCustomHrefObject{}
	this.Href = href
	return &this
}

// NewAppCustomHrefObjectWithDefaults instantiates a new AppCustomHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomHrefObjectWithDefaults() *AppCustomHrefObject {
	this := AppCustomHrefObject{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *AppCustomHrefObject) GetHints() AppCustomHrefObjectHints {
	if o == nil || o.Hints == nil {
		var ret AppCustomHrefObjectHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomHrefObject) GetHintsOk() (*AppCustomHrefObjectHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *AppCustomHrefObject) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given AppCustomHrefObjectHints and assigns it to the Hints field.
func (o *AppCustomHrefObject) SetHints(v AppCustomHrefObjectHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *AppCustomHrefObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *AppCustomHrefObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *AppCustomHrefObject) SetHref(v string) {
	o.Href = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AppCustomHrefObject) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomHrefObject) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AppCustomHrefObject) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AppCustomHrefObject) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppCustomHrefObject) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomHrefObject) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppCustomHrefObject) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppCustomHrefObject) SetType(v string) {
	o.Type = &v
}

func (o AppCustomHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppCustomHrefObject) UnmarshalJSON(bytes []byte) (err error) {
	varAppCustomHrefObject := _AppCustomHrefObject{}

	err = json.Unmarshal(bytes, &varAppCustomHrefObject)
	if err == nil {
		*o = AppCustomHrefObject(varAppCustomHrefObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppCustomHrefObject struct {
	value *AppCustomHrefObject
	isSet bool
}

func (v NullableAppCustomHrefObject) Get() *AppCustomHrefObject {
	return v.value
}

func (v *NullableAppCustomHrefObject) Set(val *AppCustomHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomHrefObject(val *AppCustomHrefObject) *NullableAppCustomHrefObject {
	return &NullableAppCustomHrefObject{value: val, isSet: true}
}

func (v NullableAppCustomHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


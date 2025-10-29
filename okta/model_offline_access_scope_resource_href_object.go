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

// checks if the OfflineAccessScopeResourceHrefObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfflineAccessScopeResourceHrefObject{}

// OfflineAccessScopeResourceHrefObject struct for OfflineAccessScopeResourceHrefObject
type OfflineAccessScopeResourceHrefObject struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	// Link name
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OfflineAccessScopeResourceHrefObject OfflineAccessScopeResourceHrefObject

// NewOfflineAccessScopeResourceHrefObject instantiates a new OfflineAccessScopeResourceHrefObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfflineAccessScopeResourceHrefObject() *OfflineAccessScopeResourceHrefObject {
	this := OfflineAccessScopeResourceHrefObject{}
	return &this
}

// NewOfflineAccessScopeResourceHrefObjectWithDefaults instantiates a new OfflineAccessScopeResourceHrefObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfflineAccessScopeResourceHrefObjectWithDefaults() *OfflineAccessScopeResourceHrefObject {
	this := OfflineAccessScopeResourceHrefObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OfflineAccessScopeResourceHrefObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAccessScopeResourceHrefObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OfflineAccessScopeResourceHrefObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OfflineAccessScopeResourceHrefObject) SetHref(v string) {
	o.Href = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OfflineAccessScopeResourceHrefObject) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAccessScopeResourceHrefObject) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OfflineAccessScopeResourceHrefObject) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OfflineAccessScopeResourceHrefObject) SetTitle(v string) {
	o.Title = &v
}

func (o OfflineAccessScopeResourceHrefObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfflineAccessScopeResourceHrefObject) ToMap() (map[string]interface{}, error) {
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

func (o *OfflineAccessScopeResourceHrefObject) UnmarshalJSON(data []byte) (err error) {
	varOfflineAccessScopeResourceHrefObject := _OfflineAccessScopeResourceHrefObject{}

	err = json.Unmarshal(data, &varOfflineAccessScopeResourceHrefObject)

	if err != nil {
		return err
	}

	*o = OfflineAccessScopeResourceHrefObject(varOfflineAccessScopeResourceHrefObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOfflineAccessScopeResourceHrefObject struct {
	value *OfflineAccessScopeResourceHrefObject
	isSet bool
}

func (v NullableOfflineAccessScopeResourceHrefObject) Get() *OfflineAccessScopeResourceHrefObject {
	return v.value
}

func (v *NullableOfflineAccessScopeResourceHrefObject) Set(val *OfflineAccessScopeResourceHrefObject) {
	v.value = val
	v.isSet = true
}

func (v NullableOfflineAccessScopeResourceHrefObject) IsSet() bool {
	return v.isSet
}

func (v *NullableOfflineAccessScopeResourceHrefObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfflineAccessScopeResourceHrefObject(val *OfflineAccessScopeResourceHrefObject) *NullableOfflineAccessScopeResourceHrefObject {
	return &NullableOfflineAccessScopeResourceHrefObject{value: val, isSet: true}
}

func (v NullableOfflineAccessScopeResourceHrefObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfflineAccessScopeResourceHrefObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

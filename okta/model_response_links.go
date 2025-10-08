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

// checks if the ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLinks{}

// ResponseLinks Link objects
type ResponseLinks struct {
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseLinks ResponseLinks

// NewResponseLinks instantiates a new ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLinks() *ResponseLinks {
	this := ResponseLinks{}
	return &this
}

// NewResponseLinksWithDefaults instantiates a new ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLinksWithDefaults() *ResponseLinks {
	this := ResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResponseLinks) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLinks) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResponseLinks) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ResponseLinks) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varResponseLinks := _ResponseLinks{}

	err = json.Unmarshal(data, &varResponseLinks)

	if err != nil {
		return err
	}

	*o = ResponseLinks(varResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseLinks struct {
	value *ResponseLinks
	isSet bool
}

func (v NullableResponseLinks) Get() *ResponseLinks {
	return v.value
}

func (v *NullableResponseLinks) Set(val *ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLinks(val *ResponseLinks) *NullableResponseLinks {
	return &NullableResponseLinks{value: val, isSet: true}
}

func (v NullableResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

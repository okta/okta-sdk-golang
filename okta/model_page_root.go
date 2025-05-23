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

// PageRoot struct for PageRoot
type PageRoot struct {
	Embedded *PageRootEmbedded `json:"_embedded,omitempty"`
	Links *PageRootLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PageRoot PageRoot

// NewPageRoot instantiates a new PageRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRoot() *PageRoot {
	this := PageRoot{}
	return &this
}

// NewPageRootWithDefaults instantiates a new PageRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRootWithDefaults() *PageRoot {
	this := PageRoot{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *PageRoot) GetEmbedded() PageRootEmbedded {
	if o == nil || o.Embedded == nil {
		var ret PageRootEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRoot) GetEmbeddedOk() (*PageRootEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *PageRoot) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given PageRootEmbedded and assigns it to the Embedded field.
func (o *PageRoot) SetEmbedded(v PageRootEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PageRoot) GetLinks() PageRootLinks {
	if o == nil || o.Links == nil {
		var ret PageRootLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRoot) GetLinksOk() (*PageRootLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PageRoot) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PageRootLinks and assigns it to the Links field.
func (o *PageRoot) SetLinks(v PageRootLinks) {
	o.Links = &v
}

func (o PageRoot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PageRoot) UnmarshalJSON(bytes []byte) (err error) {
	varPageRoot := _PageRoot{}

	err = json.Unmarshal(bytes, &varPageRoot)
	if err == nil {
		*o = PageRoot(varPageRoot)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePageRoot struct {
	value *PageRoot
	isSet bool
}

func (v NullablePageRoot) Get() *PageRoot {
	return v.value
}

func (v *NullablePageRoot) Set(val *PageRoot) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRoot) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRoot(val *PageRoot) *NullablePageRoot {
	return &NullablePageRoot{value: val, isSet: true}
}

func (v NullablePageRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


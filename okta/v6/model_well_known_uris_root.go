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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// WellKnownURIsRoot struct for WellKnownURIsRoot
type WellKnownURIsRoot struct {
	Embedded *WellKnownURIsRootEmbedded `json:"_embedded,omitempty"`
	Links *WellKnownURIsRootLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownURIsRoot WellKnownURIsRoot

// NewWellKnownURIsRoot instantiates a new WellKnownURIsRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIsRoot() *WellKnownURIsRoot {
	this := WellKnownURIsRoot{}
	return &this
}

// NewWellKnownURIsRootWithDefaults instantiates a new WellKnownURIsRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIsRootWithDefaults() *WellKnownURIsRoot {
	this := WellKnownURIsRoot{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *WellKnownURIsRoot) GetEmbedded() WellKnownURIsRootEmbedded {
	if o == nil || o.Embedded == nil {
		var ret WellKnownURIsRootEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRoot) GetEmbeddedOk() (*WellKnownURIsRootEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *WellKnownURIsRoot) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given WellKnownURIsRootEmbedded and assigns it to the Embedded field.
func (o *WellKnownURIsRoot) SetEmbedded(v WellKnownURIsRootEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WellKnownURIsRoot) GetLinks() WellKnownURIsRootLinks {
	if o == nil || o.Links == nil {
		var ret WellKnownURIsRootLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRoot) GetLinksOk() (*WellKnownURIsRootLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WellKnownURIsRoot) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WellKnownURIsRootLinks and assigns it to the Links field.
func (o *WellKnownURIsRoot) SetLinks(v WellKnownURIsRootLinks) {
	o.Links = &v
}

func (o WellKnownURIsRoot) MarshalJSON() ([]byte, error) {
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

func (o *WellKnownURIsRoot) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownURIsRoot := _WellKnownURIsRoot{}

	err = json.Unmarshal(bytes, &varWellKnownURIsRoot)
	if err == nil {
		*o = WellKnownURIsRoot(varWellKnownURIsRoot)
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

type NullableWellKnownURIsRoot struct {
	value *WellKnownURIsRoot
	isSet bool
}

func (v NullableWellKnownURIsRoot) Get() *WellKnownURIsRoot {
	return v.value
}

func (v *NullableWellKnownURIsRoot) Set(val *WellKnownURIsRoot) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIsRoot) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIsRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIsRoot(val *WellKnownURIsRoot) *NullableWellKnownURIsRoot {
	return &NullableWellKnownURIsRoot{value: val, isSet: true}
}

func (v NullableWellKnownURIsRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIsRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


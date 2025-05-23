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

// LinksNext Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. Use the `LinksNext` object for dynamic discovery of related resources and lifecycle operations.
type LinksNext struct {
	Next *HrefObject `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksNext LinksNext

// NewLinksNext instantiates a new LinksNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksNext() *LinksNext {
	this := LinksNext{}
	return &this
}

// NewLinksNextWithDefaults instantiates a new LinksNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksNextWithDefaults() *LinksNext {
	this := LinksNext{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *LinksNext) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksNext) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *LinksNext) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *LinksNext) SetNext(v HrefObject) {
	o.Next = &v
}

func (o LinksNext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksNext) UnmarshalJSON(bytes []byte) (err error) {
	varLinksNext := _LinksNext{}

	err = json.Unmarshal(bytes, &varLinksNext)
	if err == nil {
		*o = LinksNext(varLinksNext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksNext struct {
	value *LinksNext
	isSet bool
}

func (v NullableLinksNext) Get() *LinksNext {
	return v.value
}

func (v *NullableLinksNext) Set(val *LinksNext) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksNext) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksNext(val *LinksNext) *NullableLinksNext {
	return &NullableLinksNext{value: val, isSet: true}
}

func (v NullableLinksNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


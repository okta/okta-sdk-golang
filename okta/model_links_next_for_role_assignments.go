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

// checks if the LinksNextForRoleAssignments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksNextForRoleAssignments{}

// LinksNextForRoleAssignments Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification.
type LinksNextForRoleAssignments struct {
	// The next page of results if [pagination](#pagination) is required
	Next                 *HrefObject `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksNextForRoleAssignments LinksNextForRoleAssignments

// NewLinksNextForRoleAssignments instantiates a new LinksNextForRoleAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksNextForRoleAssignments() *LinksNextForRoleAssignments {
	this := LinksNextForRoleAssignments{}
	return &this
}

// NewLinksNextForRoleAssignmentsWithDefaults instantiates a new LinksNextForRoleAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksNextForRoleAssignmentsWithDefaults() *LinksNextForRoleAssignments {
	this := LinksNextForRoleAssignments{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *LinksNextForRoleAssignments) GetNext() HrefObject {
	if o == nil || IsNil(o.Next) {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksNextForRoleAssignments) GetNextOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *LinksNextForRoleAssignments) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *LinksNextForRoleAssignments) SetNext(v HrefObject) {
	o.Next = &v
}

func (o LinksNextForRoleAssignments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksNextForRoleAssignments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksNextForRoleAssignments) UnmarshalJSON(data []byte) (err error) {
	varLinksNextForRoleAssignments := _LinksNextForRoleAssignments{}

	err = json.Unmarshal(data, &varLinksNextForRoleAssignments)

	if err != nil {
		return err
	}

	*o = LinksNextForRoleAssignments(varLinksNextForRoleAssignments)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksNextForRoleAssignments struct {
	value *LinksNextForRoleAssignments
	isSet bool
}

func (v NullableLinksNextForRoleAssignments) Get() *LinksNextForRoleAssignments {
	return v.value
}

func (v *NullableLinksNextForRoleAssignments) Set(val *LinksNextForRoleAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksNextForRoleAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksNextForRoleAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksNextForRoleAssignments(val *LinksNextForRoleAssignments) *NullableLinksNextForRoleAssignments {
	return &NullableLinksNextForRoleAssignments{value: val, isSet: true}
}

func (v NullableLinksNextForRoleAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksNextForRoleAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

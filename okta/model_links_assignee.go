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

// checks if the LinksAssignee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksAssignee{}

// LinksAssignee Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification.
type LinksAssignee struct {
	Assignee             *HrefObjectAssigneeLink `json:"assignee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksAssignee LinksAssignee

// NewLinksAssignee instantiates a new LinksAssignee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksAssignee() *LinksAssignee {
	this := LinksAssignee{}
	return &this
}

// NewLinksAssigneeWithDefaults instantiates a new LinksAssignee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksAssigneeWithDefaults() *LinksAssignee {
	this := LinksAssignee{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *LinksAssignee) GetAssignee() HrefObjectAssigneeLink {
	if o == nil || IsNil(o.Assignee) {
		var ret HrefObjectAssigneeLink
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAssignee) GetAssigneeOk() (*HrefObjectAssigneeLink, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *LinksAssignee) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given HrefObjectAssigneeLink and assigns it to the Assignee field.
func (o *LinksAssignee) SetAssignee(v HrefObjectAssigneeLink) {
	o.Assignee = &v
}

func (o LinksAssignee) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksAssignee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksAssignee) UnmarshalJSON(data []byte) (err error) {
	varLinksAssignee := _LinksAssignee{}

	err = json.Unmarshal(data, &varLinksAssignee)

	if err != nil {
		return err
	}

	*o = LinksAssignee(varLinksAssignee)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksAssignee struct {
	value *LinksAssignee
	isSet bool
}

func (v NullableLinksAssignee) Get() *LinksAssignee {
	return v.value
}

func (v *NullableLinksAssignee) Set(val *LinksAssignee) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksAssignee) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksAssignee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksAssignee(val *LinksAssignee) *NullableLinksAssignee {
	return &NullableLinksAssignee{value: val, isSet: true}
}

func (v NullableLinksAssignee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksAssignee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

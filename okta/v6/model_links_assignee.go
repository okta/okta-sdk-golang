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

// LinksAssignee Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification.
type LinksAssignee struct {
	Assignee *HrefObjectAssigneeLink `json:"assignee,omitempty"`
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
	if o == nil || o.Assignee == nil {
		var ret HrefObjectAssigneeLink
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAssignee) GetAssigneeOk() (*HrefObjectAssigneeLink, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *LinksAssignee) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given HrefObjectAssigneeLink and assigns it to the Assignee field.
func (o *LinksAssignee) SetAssignee(v HrefObjectAssigneeLink) {
	o.Assignee = &v
}

func (o LinksAssignee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksAssignee) UnmarshalJSON(bytes []byte) (err error) {
	varLinksAssignee := _LinksAssignee{}

	err = json.Unmarshal(bytes, &varLinksAssignee)
	if err == nil {
		*o = LinksAssignee(varLinksAssignee)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignee")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


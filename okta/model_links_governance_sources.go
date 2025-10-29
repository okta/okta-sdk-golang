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

// checks if the LinksGovernanceSources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksGovernanceSources{}

// LinksGovernanceSources Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the sources using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification.
type LinksGovernanceSources struct {
	Assignee             *HrefObjectUserLink `json:"assignee,omitempty"`
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksGovernanceSources LinksGovernanceSources

// NewLinksGovernanceSources instantiates a new LinksGovernanceSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksGovernanceSources() *LinksGovernanceSources {
	this := LinksGovernanceSources{}
	return &this
}

// NewLinksGovernanceSourcesWithDefaults instantiates a new LinksGovernanceSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksGovernanceSourcesWithDefaults() *LinksGovernanceSources {
	this := LinksGovernanceSources{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *LinksGovernanceSources) GetAssignee() HrefObjectUserLink {
	if o == nil || IsNil(o.Assignee) {
		var ret HrefObjectUserLink
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksGovernanceSources) GetAssigneeOk() (*HrefObjectUserLink, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *LinksGovernanceSources) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given HrefObjectUserLink and assigns it to the Assignee field.
func (o *LinksGovernanceSources) SetAssignee(v HrefObjectUserLink) {
	o.Assignee = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksGovernanceSources) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksGovernanceSources) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksGovernanceSources) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksGovernanceSources) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o LinksGovernanceSources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksGovernanceSources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksGovernanceSources) UnmarshalJSON(data []byte) (err error) {
	varLinksGovernanceSources := _LinksGovernanceSources{}

	err = json.Unmarshal(data, &varLinksGovernanceSources)

	if err != nil {
		return err
	}

	*o = LinksGovernanceSources(varLinksGovernanceSources)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksGovernanceSources struct {
	value *LinksGovernanceSources
	isSet bool
}

func (v NullableLinksGovernanceSources) Get() *LinksGovernanceSources {
	return v.value
}

func (v *NullableLinksGovernanceSources) Set(val *LinksGovernanceSources) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksGovernanceSources) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksGovernanceSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksGovernanceSources(val *LinksGovernanceSources) *NullableLinksGovernanceSources {
	return &NullableLinksGovernanceSources{value: val, isSet: true}
}

func (v NullableLinksGovernanceSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksGovernanceSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

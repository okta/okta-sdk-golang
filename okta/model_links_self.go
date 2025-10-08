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

// checks if the LinksSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksSelf{}

// LinksSelf Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type LinksSelf struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelf LinksSelf

// NewLinksSelf instantiates a new LinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelf() *LinksSelf {
	this := LinksSelf{}
	return &this
}

// NewLinksSelfWithDefaults instantiates a new LinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfWithDefaults() *LinksSelf {
	this := LinksSelf{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelf) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelf) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelf) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelf) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o LinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksSelf) UnmarshalJSON(data []byte) (err error) {
	varLinksSelf := _LinksSelf{}

	err = json.Unmarshal(data, &varLinksSelf)

	if err != nil {
		return err
	}

	*o = LinksSelf(varLinksSelf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksSelf struct {
	value *LinksSelf
	isSet bool
}

func (v NullableLinksSelf) Get() *LinksSelf {
	return v.value
}

func (v *NullableLinksSelf) Set(val *LinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelf(val *LinksSelf) *NullableLinksSelf {
	return &NullableLinksSelf{value: val, isSet: true}
}

func (v NullableLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

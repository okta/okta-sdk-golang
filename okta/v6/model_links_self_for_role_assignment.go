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

// LinksSelfForRoleAssignment Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources.
type LinksSelfForRoleAssignment struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelfForRoleAssignment LinksSelfForRoleAssignment

// NewLinksSelfForRoleAssignment instantiates a new LinksSelfForRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelfForRoleAssignment() *LinksSelfForRoleAssignment {
	this := LinksSelfForRoleAssignment{}
	return &this
}

// NewLinksSelfForRoleAssignmentWithDefaults instantiates a new LinksSelfForRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfForRoleAssignmentWithDefaults() *LinksSelfForRoleAssignment {
	this := LinksSelfForRoleAssignment{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelfForRoleAssignment) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfForRoleAssignment) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfForRoleAssignment) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelfForRoleAssignment) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o LinksSelfForRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksSelfForRoleAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varLinksSelfForRoleAssignment := _LinksSelfForRoleAssignment{}

	err = json.Unmarshal(bytes, &varLinksSelfForRoleAssignment)
	if err == nil {
		*o = LinksSelfForRoleAssignment(varLinksSelfForRoleAssignment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksSelfForRoleAssignment struct {
	value *LinksSelfForRoleAssignment
	isSet bool
}

func (v NullableLinksSelfForRoleAssignment) Get() *LinksSelfForRoleAssignment {
	return v.value
}

func (v *NullableLinksSelfForRoleAssignment) Set(val *LinksSelfForRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelfForRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelfForRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelfForRoleAssignment(val *LinksSelfForRoleAssignment) *NullableLinksSelfForRoleAssignment {
	return &NullableLinksSelfForRoleAssignment{value: val, isSet: true}
}

func (v NullableLinksSelfForRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelfForRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


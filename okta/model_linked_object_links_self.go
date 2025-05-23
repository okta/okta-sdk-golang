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

// LinkedObjectLinksSelf Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type LinkedObjectLinksSelf struct {
	Self *LinkedHrefObject `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkedObjectLinksSelf LinkedObjectLinksSelf

// NewLinkedObjectLinksSelf instantiates a new LinkedObjectLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedObjectLinksSelf() *LinkedObjectLinksSelf {
	this := LinkedObjectLinksSelf{}
	return &this
}

// NewLinkedObjectLinksSelfWithDefaults instantiates a new LinkedObjectLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedObjectLinksSelfWithDefaults() *LinkedObjectLinksSelf {
	this := LinkedObjectLinksSelf{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinkedObjectLinksSelf) GetSelf() LinkedHrefObject {
	if o == nil || o.Self == nil {
		var ret LinkedHrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedObjectLinksSelf) GetSelfOk() (*LinkedHrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinkedObjectLinksSelf) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given LinkedHrefObject and assigns it to the Self field.
func (o *LinkedObjectLinksSelf) SetSelf(v LinkedHrefObject) {
	o.Self = &v
}

func (o LinkedObjectLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkedObjectLinksSelf) UnmarshalJSON(bytes []byte) (err error) {
	varLinkedObjectLinksSelf := _LinkedObjectLinksSelf{}

	err = json.Unmarshal(bytes, &varLinkedObjectLinksSelf)
	if err == nil {
		*o = LinkedObjectLinksSelf(varLinkedObjectLinksSelf)
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

type NullableLinkedObjectLinksSelf struct {
	value *LinkedObjectLinksSelf
	isSet bool
}

func (v NullableLinkedObjectLinksSelf) Get() *LinkedObjectLinksSelf {
	return v.value
}

func (v *NullableLinkedObjectLinksSelf) Set(val *LinkedObjectLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedObjectLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedObjectLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedObjectLinksSelf(val *LinkedObjectLinksSelf) *NullableLinkedObjectLinksSelf {
	return &NullableLinkedObjectLinksSelf{value: val, isSet: true}
}

func (v NullableLinkedObjectLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedObjectLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


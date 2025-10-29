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

// checks if the LinksUserRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksUserRef{}

// LinksUserRef struct for LinksUserRef
type LinksUserRef struct {
	User                 *HrefObjectUserLink `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksUserRef LinksUserRef

// NewLinksUserRef instantiates a new LinksUserRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksUserRef() *LinksUserRef {
	this := LinksUserRef{}
	return &this
}

// NewLinksUserRefWithDefaults instantiates a new LinksUserRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksUserRefWithDefaults() *LinksUserRef {
	this := LinksUserRef{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LinksUserRef) GetUser() HrefObjectUserLink {
	if o == nil || IsNil(o.User) {
		var ret HrefObjectUserLink
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksUserRef) GetUserOk() (*HrefObjectUserLink, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LinksUserRef) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObjectUserLink and assigns it to the User field.
func (o *LinksUserRef) SetUser(v HrefObjectUserLink) {
	o.User = &v
}

func (o LinksUserRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksUserRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksUserRef) UnmarshalJSON(data []byte) (err error) {
	varLinksUserRef := _LinksUserRef{}

	err = json.Unmarshal(data, &varLinksUserRef)

	if err != nil {
		return err
	}

	*o = LinksUserRef(varLinksUserRef)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksUserRef struct {
	value *LinksUserRef
	isSet bool
}

func (v NullableLinksUserRef) Get() *LinksUserRef {
	return v.value
}

func (v *NullableLinksUserRef) Set(val *LinksUserRef) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksUserRef) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksUserRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksUserRef(val *LinksUserRef) *NullableLinksUserRef {
	return &NullableLinksUserRef{value: val, isSet: true}
}

func (v NullableLinksUserRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksUserRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

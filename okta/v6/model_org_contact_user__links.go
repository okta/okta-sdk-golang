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

// checks if the OrgContactUserLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgContactUserLinks{}

// OrgContactUserLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the contact type user object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgContactUserLinks struct {
	User                 *HrefObjectUserLink `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgContactUserLinks OrgContactUserLinks

// NewOrgContactUserLinks instantiates a new OrgContactUserLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgContactUserLinks() *OrgContactUserLinks {
	this := OrgContactUserLinks{}
	return &this
}

// NewOrgContactUserLinksWithDefaults instantiates a new OrgContactUserLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgContactUserLinksWithDefaults() *OrgContactUserLinks {
	this := OrgContactUserLinks{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrgContactUserLinks) GetUser() HrefObjectUserLink {
	if o == nil || IsNil(o.User) {
		var ret HrefObjectUserLink
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContactUserLinks) GetUserOk() (*HrefObjectUserLink, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrgContactUserLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObjectUserLink and assigns it to the User field.
func (o *OrgContactUserLinks) SetUser(v HrefObjectUserLink) {
	o.User = &v
}

func (o OrgContactUserLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgContactUserLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgContactUserLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgContactUserLinks := _OrgContactUserLinks{}

	err = json.Unmarshal(data, &varOrgContactUserLinks)

	if err != nil {
		return err
	}

	*o = OrgContactUserLinks(varOrgContactUserLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgContactUserLinks struct {
	value *OrgContactUserLinks
	isSet bool
}

func (v NullableOrgContactUserLinks) Get() *OrgContactUserLinks {
	return v.value
}

func (v *NullableOrgContactUserLinks) Set(val *OrgContactUserLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgContactUserLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgContactUserLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgContactUserLinks(val *OrgContactUserLinks) *NullableOrgContactUserLinks {
	return &NullableOrgContactUserLinks{value: val, isSet: true}
}

func (v NullableOrgContactUserLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgContactUserLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the OrgContactUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgContactUser{}

// OrgContactUser struct for OrgContactUser
type OrgContactUser struct {
	// Contact user ID
	UserId               *string              `json:"userId,omitempty"`
	Links                *OrgContactUserLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgContactUser OrgContactUser

// NewOrgContactUser instantiates a new OrgContactUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgContactUser() *OrgContactUser {
	this := OrgContactUser{}
	return &this
}

// NewOrgContactUserWithDefaults instantiates a new OrgContactUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgContactUserWithDefaults() *OrgContactUser {
	this := OrgContactUser{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OrgContactUser) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContactUser) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OrgContactUser) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OrgContactUser) SetUserId(v string) {
	o.UserId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgContactUser) GetLinks() OrgContactUserLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgContactUserLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContactUser) GetLinksOk() (*OrgContactUserLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgContactUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgContactUserLinks and assigns it to the Links field.
func (o *OrgContactUser) SetLinks(v OrgContactUserLinks) {
	o.Links = &v
}

func (o OrgContactUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgContactUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgContactUser) UnmarshalJSON(data []byte) (err error) {
	varOrgContactUser := _OrgContactUser{}

	err = json.Unmarshal(data, &varOrgContactUser)

	if err != nil {
		return err
	}

	*o = OrgContactUser(varOrgContactUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgContactUser struct {
	value *OrgContactUser
	isSet bool
}

func (v NullableOrgContactUser) Get() *OrgContactUser {
	return v.value
}

func (v *NullableOrgContactUser) Set(val *OrgContactUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgContactUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgContactUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgContactUser(val *OrgContactUser) *NullableOrgContactUser {
	return &NullableOrgContactUser{value: val, isSet: true}
}

func (v NullableOrgContactUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgContactUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

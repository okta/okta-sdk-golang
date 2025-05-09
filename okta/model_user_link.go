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

// UserLink struct for UserLink
type UserLink struct {
	User *HrefObjectUserLink `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserLink UserLink

// NewUserLink instantiates a new UserLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLink() *UserLink {
	this := UserLink{}
	return &this
}

// NewUserLinkWithDefaults instantiates a new UserLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLinkWithDefaults() *UserLink {
	this := UserLink{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserLink) GetUser() HrefObjectUserLink {
	if o == nil || o.User == nil {
		var ret HrefObjectUserLink
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLink) GetUserOk() (*HrefObjectUserLink, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserLink) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObjectUserLink and assigns it to the User field.
func (o *UserLink) SetUser(v HrefObjectUserLink) {
	o.User = &v
}

func (o UserLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserLink) UnmarshalJSON(bytes []byte) (err error) {
	varUserLink := _UserLink{}

	err = json.Unmarshal(bytes, &varUserLink)
	if err == nil {
		*o = UserLink(varUserLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserLink struct {
	value *UserLink
	isSet bool
}

func (v NullableUserLink) Get() *UserLink {
	return v.value
}

func (v *NullableUserLink) Set(val *UserLink) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLink) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLink(val *UserLink) *NullableUserLink {
	return &NullableUserLink{value: val, isSet: true}
}

func (v NullableUserLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


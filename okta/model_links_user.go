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

// LinksUser struct for LinksUser
type LinksUser struct {
	User *LinksUserUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksUser LinksUser

// NewLinksUser instantiates a new LinksUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksUser() *LinksUser {
	this := LinksUser{}
	return &this
}

// NewLinksUserWithDefaults instantiates a new LinksUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksUserWithDefaults() *LinksUser {
	this := LinksUser{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LinksUser) GetUser() LinksUserUser {
	if o == nil || o.User == nil {
		var ret LinksUserUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksUser) GetUserOk() (*LinksUserUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LinksUser) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given LinksUserUser and assigns it to the User field.
func (o *LinksUser) SetUser(v LinksUserUser) {
	o.User = &v
}

func (o LinksUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksUser) UnmarshalJSON(bytes []byte) (err error) {
	varLinksUser := _LinksUser{}

	err = json.Unmarshal(bytes, &varLinksUser)
	if err == nil {
		*o = LinksUser(varLinksUser)
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

type NullableLinksUser struct {
	value *LinksUser
	isSet bool
}

func (v NullableLinksUser) Get() *LinksUser {
	return v.value
}

func (v *NullableLinksUser) Set(val *LinksUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksUser(val *LinksUser) *NullableLinksUser {
	return &NullableLinksUser{value: val, isSet: true}
}

func (v NullableLinksUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


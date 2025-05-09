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

// UserFactorActivateResponseLinks struct for UserFactorActivateResponseLinks
type UserFactorActivateResponseLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	User *LinksUserUser `json:"user,omitempty"`
	Verify *LinksVerifyVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorActivateResponseLinks UserFactorActivateResponseLinks

// NewUserFactorActivateResponseLinks instantiates a new UserFactorActivateResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorActivateResponseLinks() *UserFactorActivateResponseLinks {
	this := UserFactorActivateResponseLinks{}
	return &this
}

// NewUserFactorActivateResponseLinksWithDefaults instantiates a new UserFactorActivateResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorActivateResponseLinksWithDefaults() *UserFactorActivateResponseLinks {
	this := UserFactorActivateResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *UserFactorActivateResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *UserFactorActivateResponseLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *UserFactorActivateResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserFactorActivateResponseLinks) GetUser() LinksUserUser {
	if o == nil || o.User == nil {
		var ret LinksUserUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateResponseLinks) GetUserOk() (*LinksUserUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserFactorActivateResponseLinks) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given LinksUserUser and assigns it to the User field.
func (o *UserFactorActivateResponseLinks) SetUser(v LinksUserUser) {
	o.User = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorActivateResponseLinks) GetVerify() LinksVerifyVerify {
	if o == nil || o.Verify == nil {
		var ret LinksVerifyVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateResponseLinks) GetVerifyOk() (*LinksVerifyVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorActivateResponseLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given LinksVerifyVerify and assigns it to the Verify field.
func (o *UserFactorActivateResponseLinks) SetVerify(v LinksVerifyVerify) {
	o.Verify = &v
}

func (o UserFactorActivateResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorActivateResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorActivateResponseLinks := _UserFactorActivateResponseLinks{}

	err = json.Unmarshal(bytes, &varUserFactorActivateResponseLinks)
	if err == nil {
		*o = UserFactorActivateResponseLinks(varUserFactorActivateResponseLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "user")
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorActivateResponseLinks struct {
	value *UserFactorActivateResponseLinks
	isSet bool
}

func (v NullableUserFactorActivateResponseLinks) Get() *UserFactorActivateResponseLinks {
	return v.value
}

func (v *NullableUserFactorActivateResponseLinks) Set(val *UserFactorActivateResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorActivateResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorActivateResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorActivateResponseLinks(val *UserFactorActivateResponseLinks) *NullableUserFactorActivateResponseLinks {
	return &NullableUserFactorActivateResponseLinks{value: val, isSet: true}
}

func (v NullableUserFactorActivateResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorActivateResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


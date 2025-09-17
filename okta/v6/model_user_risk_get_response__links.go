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

// checks if the UserRiskGetResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskGetResponseLinks{}

// UserRiskGetResponseLinks struct for UserRiskGetResponseLinks
type UserRiskGetResponseLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	User                 *HrefObjectUserLink `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRiskGetResponseLinks UserRiskGetResponseLinks

// NewUserRiskGetResponseLinks instantiates a new UserRiskGetResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskGetResponseLinks() *UserRiskGetResponseLinks {
	this := UserRiskGetResponseLinks{}
	return &this
}

// NewUserRiskGetResponseLinksWithDefaults instantiates a new UserRiskGetResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskGetResponseLinksWithDefaults() *UserRiskGetResponseLinks {
	this := UserRiskGetResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *UserRiskGetResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskGetResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *UserRiskGetResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *UserRiskGetResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserRiskGetResponseLinks) GetUser() HrefObjectUserLink {
	if o == nil || IsNil(o.User) {
		var ret HrefObjectUserLink
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskGetResponseLinks) GetUserOk() (*HrefObjectUserLink, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserRiskGetResponseLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObjectUserLink and assigns it to the User field.
func (o *UserRiskGetResponseLinks) SetUser(v HrefObjectUserLink) {
	o.User = &v
}

func (o UserRiskGetResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskGetResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRiskGetResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varUserRiskGetResponseLinks := _UserRiskGetResponseLinks{}

	err = json.Unmarshal(data, &varUserRiskGetResponseLinks)

	if err != nil {
		return err
	}

	*o = UserRiskGetResponseLinks(varUserRiskGetResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRiskGetResponseLinks struct {
	value *UserRiskGetResponseLinks
	isSet bool
}

func (v NullableUserRiskGetResponseLinks) Get() *UserRiskGetResponseLinks {
	return v.value
}

func (v *NullableUserRiskGetResponseLinks) Set(val *UserRiskGetResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskGetResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskGetResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskGetResponseLinks(val *UserRiskGetResponseLinks) *NullableUserRiskGetResponseLinks {
	return &NullableUserRiskGetResponseLinks{value: val, isSet: true}
}

func (v NullableUserRiskGetResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskGetResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

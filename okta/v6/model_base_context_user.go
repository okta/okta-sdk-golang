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
	"time"
)

// BaseContextUser Identifies the Okta user that the token was generated to authenticate and provides details of their Okta user profile
type BaseContextUser struct {
	// The unique identifier for the user
	Id *string `json:"id,omitempty"`
	// The timestamp when the user's password was last updated
	PasswordChanged *time.Time `json:"passwordChanged,omitempty"`
	Profile *BaseContextUserProfile `json:"profile,omitempty"`
	Links *BaseContextUserLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseContextUser BaseContextUser

// NewBaseContextUser instantiates a new BaseContextUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContextUser() *BaseContextUser {
	this := BaseContextUser{}
	return &this
}

// NewBaseContextUserWithDefaults instantiates a new BaseContextUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContextUserWithDefaults() *BaseContextUser {
	this := BaseContextUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseContextUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseContextUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseContextUser) SetId(v string) {
	o.Id = &v
}

// GetPasswordChanged returns the PasswordChanged field value if set, zero value otherwise.
func (o *BaseContextUser) GetPasswordChanged() time.Time {
	if o == nil || o.PasswordChanged == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordChanged
}

// GetPasswordChangedOk returns a tuple with the PasswordChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUser) GetPasswordChangedOk() (*time.Time, bool) {
	if o == nil || o.PasswordChanged == nil {
		return nil, false
	}
	return o.PasswordChanged, true
}

// HasPasswordChanged returns a boolean if a field has been set.
func (o *BaseContextUser) HasPasswordChanged() bool {
	if o != nil && o.PasswordChanged != nil {
		return true
	}

	return false
}

// SetPasswordChanged gets a reference to the given time.Time and assigns it to the PasswordChanged field.
func (o *BaseContextUser) SetPasswordChanged(v time.Time) {
	o.PasswordChanged = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *BaseContextUser) GetProfile() BaseContextUserProfile {
	if o == nil || o.Profile == nil {
		var ret BaseContextUserProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUser) GetProfileOk() (*BaseContextUserProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *BaseContextUser) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given BaseContextUserProfile and assigns it to the Profile field.
func (o *BaseContextUser) SetProfile(v BaseContextUserProfile) {
	o.Profile = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BaseContextUser) GetLinks() BaseContextUserLinks {
	if o == nil || o.Links == nil {
		var ret BaseContextUserLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUser) GetLinksOk() (*BaseContextUserLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BaseContextUser) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BaseContextUserLinks and assigns it to the Links field.
func (o *BaseContextUser) SetLinks(v BaseContextUserLinks) {
	o.Links = &v
}

func (o BaseContextUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PasswordChanged != nil {
		toSerialize["passwordChanged"] = o.PasswordChanged
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseContextUser) UnmarshalJSON(bytes []byte) (err error) {
	varBaseContextUser := _BaseContextUser{}

	err = json.Unmarshal(bytes, &varBaseContextUser)
	if err == nil {
		*o = BaseContextUser(varBaseContextUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "passwordChanged")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBaseContextUser struct {
	value *BaseContextUser
	isSet bool
}

func (v NullableBaseContextUser) Get() *BaseContextUser {
	return v.value
}

func (v *NullableBaseContextUser) Set(val *BaseContextUser) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContextUser) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContextUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContextUser(val *BaseContextUser) *NullableBaseContextUser {
	return &NullableBaseContextUser{value: val, isSet: true}
}

func (v NullableBaseContextUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContextUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


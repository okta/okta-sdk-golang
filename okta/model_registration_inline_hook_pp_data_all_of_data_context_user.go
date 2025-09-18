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
	"time"
)

// checks if the RegistrationInlineHookPPDataAllOfDataContextUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookPPDataAllOfDataContextUser{}

// RegistrationInlineHookPPDataAllOfDataContextUser struct for RegistrationInlineHookPPDataAllOfDataContextUser
type RegistrationInlineHookPPDataAllOfDataContextUser struct {
	// The last time the user's password was updated
	PasswordChanged *time.Time            `json:"passwordChanged,omitempty"`
	Links           *BaseContextUserLinks `json:"_links,omitempty"`
	// The user to update's current attributes
	Profile              map[string]interface{} `json:"profile,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookPPDataAllOfDataContextUser RegistrationInlineHookPPDataAllOfDataContextUser

// NewRegistrationInlineHookPPDataAllOfDataContextUser instantiates a new RegistrationInlineHookPPDataAllOfDataContextUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookPPDataAllOfDataContextUser() *RegistrationInlineHookPPDataAllOfDataContextUser {
	this := RegistrationInlineHookPPDataAllOfDataContextUser{}
	return &this
}

// NewRegistrationInlineHookPPDataAllOfDataContextUserWithDefaults instantiates a new RegistrationInlineHookPPDataAllOfDataContextUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookPPDataAllOfDataContextUserWithDefaults() *RegistrationInlineHookPPDataAllOfDataContextUser {
	this := RegistrationInlineHookPPDataAllOfDataContextUser{}
	return &this
}

// GetPasswordChanged returns the PasswordChanged field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetPasswordChanged() time.Time {
	if o == nil || IsNil(o.PasswordChanged) {
		var ret time.Time
		return ret
	}
	return *o.PasswordChanged
}

// GetPasswordChangedOk returns a tuple with the PasswordChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetPasswordChangedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PasswordChanged) {
		return nil, false
	}
	return o.PasswordChanged, true
}

// HasPasswordChanged returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) HasPasswordChanged() bool {
	if o != nil && !IsNil(o.PasswordChanged) {
		return true
	}

	return false
}

// SetPasswordChanged gets a reference to the given time.Time and assigns it to the PasswordChanged field.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) SetPasswordChanged(v time.Time) {
	o.PasswordChanged = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetLinks() BaseContextUserLinks {
	if o == nil || IsNil(o.Links) {
		var ret BaseContextUserLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetLinksOk() (*BaseContextUserLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BaseContextUserLinks and assigns it to the Links field.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) SetLinks(v BaseContextUserLinks) {
	o.Links = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrationInlineHookPPDataAllOfDataContextUser) SetId(v string) {
	o.Id = &v
}

func (o RegistrationInlineHookPPDataAllOfDataContextUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookPPDataAllOfDataContextUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PasswordChanged) {
		toSerialize["passwordChanged"] = o.PasswordChanged
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookPPDataAllOfDataContextUser) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookPPDataAllOfDataContextUser := _RegistrationInlineHookPPDataAllOfDataContextUser{}

	err = json.Unmarshal(data, &varRegistrationInlineHookPPDataAllOfDataContextUser)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookPPDataAllOfDataContextUser(varRegistrationInlineHookPPDataAllOfDataContextUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "passwordChanged")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookPPDataAllOfDataContextUser struct {
	value *RegistrationInlineHookPPDataAllOfDataContextUser
	isSet bool
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContextUser) Get() *RegistrationInlineHookPPDataAllOfDataContextUser {
	return v.value
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContextUser) Set(val *RegistrationInlineHookPPDataAllOfDataContextUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContextUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContextUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookPPDataAllOfDataContextUser(val *RegistrationInlineHookPPDataAllOfDataContextUser) *NullableRegistrationInlineHookPPDataAllOfDataContextUser {
	return &NullableRegistrationInlineHookPPDataAllOfDataContextUser{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContextUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContextUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

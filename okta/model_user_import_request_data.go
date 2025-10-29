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

// checks if the UserImportRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequestData{}

// UserImportRequestData struct for UserImportRequestData
type UserImportRequestData struct {
	Action               *UserImportRequestDataAction  `json:"action,omitempty"`
	AppUser              *UserImportRequestDataAppUser `json:"appUser,omitempty"`
	Context              *UserImportRequestDataContext `json:"context,omitempty"`
	User                 *UserImportRequestDataUser    `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestData UserImportRequestData

// NewUserImportRequestData instantiates a new UserImportRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestData() *UserImportRequestData {
	this := UserImportRequestData{}
	return &this
}

// NewUserImportRequestDataWithDefaults instantiates a new UserImportRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataWithDefaults() *UserImportRequestData {
	this := UserImportRequestData{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserImportRequestData) GetAction() UserImportRequestDataAction {
	if o == nil || IsNil(o.Action) {
		var ret UserImportRequestDataAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestData) GetActionOk() (*UserImportRequestDataAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserImportRequestData) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given UserImportRequestDataAction and assigns it to the Action field.
func (o *UserImportRequestData) SetAction(v UserImportRequestDataAction) {
	o.Action = &v
}

// GetAppUser returns the AppUser field value if set, zero value otherwise.
func (o *UserImportRequestData) GetAppUser() UserImportRequestDataAppUser {
	if o == nil || IsNil(o.AppUser) {
		var ret UserImportRequestDataAppUser
		return ret
	}
	return *o.AppUser
}

// GetAppUserOk returns a tuple with the AppUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestData) GetAppUserOk() (*UserImportRequestDataAppUser, bool) {
	if o == nil || IsNil(o.AppUser) {
		return nil, false
	}
	return o.AppUser, true
}

// HasAppUser returns a boolean if a field has been set.
func (o *UserImportRequestData) HasAppUser() bool {
	if o != nil && !IsNil(o.AppUser) {
		return true
	}

	return false
}

// SetAppUser gets a reference to the given UserImportRequestDataAppUser and assigns it to the AppUser field.
func (o *UserImportRequestData) SetAppUser(v UserImportRequestDataAppUser) {
	o.AppUser = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UserImportRequestData) GetContext() UserImportRequestDataContext {
	if o == nil || IsNil(o.Context) {
		var ret UserImportRequestDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestData) GetContextOk() (*UserImportRequestDataContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UserImportRequestData) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given UserImportRequestDataContext and assigns it to the Context field.
func (o *UserImportRequestData) SetContext(v UserImportRequestDataContext) {
	o.Context = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserImportRequestData) GetUser() UserImportRequestDataUser {
	if o == nil || IsNil(o.User) {
		var ret UserImportRequestDataUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestData) GetUserOk() (*UserImportRequestDataUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserImportRequestData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserImportRequestDataUser and assigns it to the User field.
func (o *UserImportRequestData) SetUser(v UserImportRequestDataUser) {
	o.User = &v
}

func (o UserImportRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AppUser) {
		toSerialize["appUser"] = o.AppUser
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequestData) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequestData := _UserImportRequestData{}

	err = json.Unmarshal(data, &varUserImportRequestData)

	if err != nil {
		return err
	}

	*o = UserImportRequestData(varUserImportRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "appUser")
		delete(additionalProperties, "context")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequestData struct {
	value *UserImportRequestData
	isSet bool
}

func (v NullableUserImportRequestData) Get() *UserImportRequestData {
	return v.value
}

func (v *NullableUserImportRequestData) Set(val *UserImportRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestData(val *UserImportRequestData) *NullableUserImportRequestData {
	return &NullableUserImportRequestData{value: val, isSet: true}
}

func (v NullableUserImportRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

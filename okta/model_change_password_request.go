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

// ChangePasswordRequest struct for ChangePasswordRequest
type ChangePasswordRequest struct {
	NewPassword *PasswordCredential `json:"newPassword,omitempty"`
	OldPassword *PasswordCredential `json:"oldPassword,omitempty"`
	RevokeSessions *bool `json:"revokeSessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangePasswordRequest ChangePasswordRequest

// NewChangePasswordRequest instantiates a new ChangePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangePasswordRequest() *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	return &this
}

// NewChangePasswordRequestWithDefaults instantiates a new ChangePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangePasswordRequestWithDefaults() *ChangePasswordRequest {
	this := ChangePasswordRequest{}
	return &this
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *ChangePasswordRequest) GetNewPassword() PasswordCredential {
	if o == nil || o.NewPassword == nil {
		var ret PasswordCredential
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetNewPasswordOk() (*PasswordCredential, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *ChangePasswordRequest) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given PasswordCredential and assigns it to the NewPassword field.
func (o *ChangePasswordRequest) SetNewPassword(v PasswordCredential) {
	o.NewPassword = &v
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise.
func (o *ChangePasswordRequest) GetOldPassword() PasswordCredential {
	if o == nil || o.OldPassword == nil {
		var ret PasswordCredential
		return ret
	}
	return *o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetOldPasswordOk() (*PasswordCredential, bool) {
	if o == nil || o.OldPassword == nil {
		return nil, false
	}
	return o.OldPassword, true
}

// HasOldPassword returns a boolean if a field has been set.
func (o *ChangePasswordRequest) HasOldPassword() bool {
	if o != nil && o.OldPassword != nil {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given PasswordCredential and assigns it to the OldPassword field.
func (o *ChangePasswordRequest) SetOldPassword(v PasswordCredential) {
	o.OldPassword = &v
}

// GetRevokeSessions returns the RevokeSessions field value if set, zero value otherwise.
func (o *ChangePasswordRequest) GetRevokeSessions() bool {
	if o == nil || o.RevokeSessions == nil {
		var ret bool
		return ret
	}
	return *o.RevokeSessions
}

// GetRevokeSessionsOk returns a tuple with the RevokeSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangePasswordRequest) GetRevokeSessionsOk() (*bool, bool) {
	if o == nil || o.RevokeSessions == nil {
		return nil, false
	}
	return o.RevokeSessions, true
}

// HasRevokeSessions returns a boolean if a field has been set.
func (o *ChangePasswordRequest) HasRevokeSessions() bool {
	if o != nil && o.RevokeSessions != nil {
		return true
	}

	return false
}

// SetRevokeSessions gets a reference to the given bool and assigns it to the RevokeSessions field.
func (o *ChangePasswordRequest) SetRevokeSessions(v bool) {
	o.RevokeSessions = &v
}

func (o ChangePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewPassword != nil {
		toSerialize["newPassword"] = o.NewPassword
	}
	if o.OldPassword != nil {
		toSerialize["oldPassword"] = o.OldPassword
	}
	if o.RevokeSessions != nil {
		toSerialize["revokeSessions"] = o.RevokeSessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChangePasswordRequest) UnmarshalJSON(bytes []byte) (err error) {
	varChangePasswordRequest := _ChangePasswordRequest{}

	err = json.Unmarshal(bytes, &varChangePasswordRequest)
	if err == nil {
		*o = ChangePasswordRequest(varChangePasswordRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "newPassword")
		delete(additionalProperties, "oldPassword")
		delete(additionalProperties, "revokeSessions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableChangePasswordRequest struct {
	value *ChangePasswordRequest
	isSet bool
}

func (v NullableChangePasswordRequest) Get() *ChangePasswordRequest {
	return v.value
}

func (v *NullableChangePasswordRequest) Set(val *ChangePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangePasswordRequest(val *ChangePasswordRequest) *NullableChangePasswordRequest {
	return &NullableChangePasswordRequest{value: val, isSet: true}
}

func (v NullableChangePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// UserCredentials struct for UserCredentials
type UserCredentials struct {
	Password *PasswordCredential `json:"password,omitempty"`
	Provider *AuthenticationProvider `json:"provider,omitempty"`
	RecoveryQuestion *RecoveryQuestionCredential `json:"recovery_question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCredentials UserCredentials

// NewUserCredentials instantiates a new UserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredentials() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// NewUserCredentialsWithDefaults instantiates a new UserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialsWithDefaults() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCredentials) GetPassword() PasswordCredential {
	if o == nil || o.Password == nil {
		var ret PasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetPasswordOk() (*PasswordCredential, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordCredential and assigns it to the Password field.
func (o *UserCredentials) SetPassword(v PasswordCredential) {
	o.Password = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserCredentials) GetProvider() AuthenticationProvider {
	if o == nil || o.Provider == nil {
		var ret AuthenticationProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetProviderOk() (*AuthenticationProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserCredentials) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticationProvider and assigns it to the Provider field.
func (o *UserCredentials) SetProvider(v AuthenticationProvider) {
	o.Provider = &v
}

// GetRecoveryQuestion returns the RecoveryQuestion field value if set, zero value otherwise.
func (o *UserCredentials) GetRecoveryQuestion() RecoveryQuestionCredential {
	if o == nil || o.RecoveryQuestion == nil {
		var ret RecoveryQuestionCredential
		return ret
	}
	return *o.RecoveryQuestion
}

// GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool) {
	if o == nil || o.RecoveryQuestion == nil {
		return nil, false
	}
	return o.RecoveryQuestion, true
}

// HasRecoveryQuestion returns a boolean if a field has been set.
func (o *UserCredentials) HasRecoveryQuestion() bool {
	if o != nil && o.RecoveryQuestion != nil {
		return true
	}

	return false
}

// SetRecoveryQuestion gets a reference to the given RecoveryQuestionCredential and assigns it to the RecoveryQuestion field.
func (o *UserCredentials) SetRecoveryQuestion(v RecoveryQuestionCredential) {
	o.RecoveryQuestion = &v
}

func (o UserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.RecoveryQuestion != nil {
		toSerialize["recovery_question"] = o.RecoveryQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varUserCredentials := _UserCredentials{}

	err = json.Unmarshal(bytes, &varUserCredentials)
	if err == nil {
		*o = UserCredentials(varUserCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "recovery_question")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserCredentials struct {
	value *UserCredentials
	isSet bool
}

func (v NullableUserCredentials) Get() *UserCredentials {
	return v.value
}

func (v *NullableUserCredentials) Set(val *UserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredentials(val *UserCredentials) *NullableUserCredentials {
	return &NullableUserCredentials{value: val, isSet: true}
}

func (v NullableUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


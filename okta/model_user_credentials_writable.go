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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserCredentialsWritable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCredentialsWritable{}

// UserCredentialsWritable Specifies primary authentication and recovery credentials for a user. Credential types and requirements vary depending on the provider and security policy of the org.
type UserCredentialsWritable struct {
	Password             *PasswordCredential             `json:"password,omitempty"`
	Provider             *AuthenticationProviderWritable `json:"provider,omitempty"`
	RecoveryQuestion     *RecoveryQuestionCredential     `json:"recovery_question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserCredentialsWritable UserCredentialsWritable

// NewUserCredentialsWritable instantiates a new UserCredentialsWritable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredentialsWritable() *UserCredentialsWritable {
	this := UserCredentialsWritable{}
	return &this
}

// NewUserCredentialsWritableWithDefaults instantiates a new UserCredentialsWritable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialsWritableWithDefaults() *UserCredentialsWritable {
	this := UserCredentialsWritable{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserCredentialsWritable) GetPassword() PasswordCredential {
	if o == nil || IsNil(o.Password) {
		var ret PasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentialsWritable) GetPasswordOk() (*PasswordCredential, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCredentialsWritable) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordCredential and assigns it to the Password field.
func (o *UserCredentialsWritable) SetPassword(v PasswordCredential) {
	o.Password = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserCredentialsWritable) GetProvider() AuthenticationProviderWritable {
	if o == nil || IsNil(o.Provider) {
		var ret AuthenticationProviderWritable
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentialsWritable) GetProviderOk() (*AuthenticationProviderWritable, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserCredentialsWritable) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AuthenticationProviderWritable and assigns it to the Provider field.
func (o *UserCredentialsWritable) SetProvider(v AuthenticationProviderWritable) {
	o.Provider = &v
}

// GetRecoveryQuestion returns the RecoveryQuestion field value if set, zero value otherwise.
func (o *UserCredentialsWritable) GetRecoveryQuestion() RecoveryQuestionCredential {
	if o == nil || IsNil(o.RecoveryQuestion) {
		var ret RecoveryQuestionCredential
		return ret
	}
	return *o.RecoveryQuestion
}

// GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentialsWritable) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool) {
	if o == nil || IsNil(o.RecoveryQuestion) {
		return nil, false
	}
	return o.RecoveryQuestion, true
}

// HasRecoveryQuestion returns a boolean if a field has been set.
func (o *UserCredentialsWritable) HasRecoveryQuestion() bool {
	if o != nil && !IsNil(o.RecoveryQuestion) {
		return true
	}

	return false
}

// SetRecoveryQuestion gets a reference to the given RecoveryQuestionCredential and assigns it to the RecoveryQuestion field.
func (o *UserCredentialsWritable) SetRecoveryQuestion(v RecoveryQuestionCredential) {
	o.RecoveryQuestion = &v
}

func (o UserCredentialsWritable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCredentialsWritable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.RecoveryQuestion) {
		toSerialize["recovery_question"] = o.RecoveryQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserCredentialsWritable) UnmarshalJSON(data []byte) (err error) {
	varUserCredentialsWritable := _UserCredentialsWritable{}

	err = json.Unmarshal(data, &varUserCredentialsWritable)

	if err != nil {
		return err
	}

	*o = UserCredentialsWritable(varUserCredentialsWritable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "recovery_question")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCredentialsWritable struct {
	value *UserCredentialsWritable
	isSet bool
}

func (v NullableUserCredentialsWritable) Get() *UserCredentialsWritable {
	return v.value
}

func (v *NullableUserCredentialsWritable) Set(val *UserCredentialsWritable) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredentialsWritable) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredentialsWritable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredentialsWritable(val *UserCredentialsWritable) *NullableUserCredentialsWritable {
	return &NullableUserCredentialsWritable{value: val, isSet: true}
}

func (v NullableUserCredentialsWritable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredentialsWritable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

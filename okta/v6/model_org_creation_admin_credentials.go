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
)

// OrgCreationAdminCredentials Specifies primary authentication and recovery credentials for a user. Credential types and requirements vary depending on the provider and security policy of the org.
type OrgCreationAdminCredentials struct {
	Password *OrgCreationAdminCredentialsPassword `json:"password,omitempty"`
	RecoveryQuestion *RecoveryQuestionCredential `json:"recovery_question,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCreationAdminCredentials OrgCreationAdminCredentials

// NewOrgCreationAdminCredentials instantiates a new OrgCreationAdminCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCreationAdminCredentials() *OrgCreationAdminCredentials {
	this := OrgCreationAdminCredentials{}
	return &this
}

// NewOrgCreationAdminCredentialsWithDefaults instantiates a new OrgCreationAdminCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCreationAdminCredentialsWithDefaults() *OrgCreationAdminCredentials {
	this := OrgCreationAdminCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OrgCreationAdminCredentials) GetPassword() OrgCreationAdminCredentialsPassword {
	if o == nil || o.Password == nil {
		var ret OrgCreationAdminCredentialsPassword
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCreationAdminCredentials) GetPasswordOk() (*OrgCreationAdminCredentialsPassword, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OrgCreationAdminCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given OrgCreationAdminCredentialsPassword and assigns it to the Password field.
func (o *OrgCreationAdminCredentials) SetPassword(v OrgCreationAdminCredentialsPassword) {
	o.Password = &v
}

// GetRecoveryQuestion returns the RecoveryQuestion field value if set, zero value otherwise.
func (o *OrgCreationAdminCredentials) GetRecoveryQuestion() RecoveryQuestionCredential {
	if o == nil || o.RecoveryQuestion == nil {
		var ret RecoveryQuestionCredential
		return ret
	}
	return *o.RecoveryQuestion
}

// GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCreationAdminCredentials) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool) {
	if o == nil || o.RecoveryQuestion == nil {
		return nil, false
	}
	return o.RecoveryQuestion, true
}

// HasRecoveryQuestion returns a boolean if a field has been set.
func (o *OrgCreationAdminCredentials) HasRecoveryQuestion() bool {
	if o != nil && o.RecoveryQuestion != nil {
		return true
	}

	return false
}

// SetRecoveryQuestion gets a reference to the given RecoveryQuestionCredential and assigns it to the RecoveryQuestion field.
func (o *OrgCreationAdminCredentials) SetRecoveryQuestion(v RecoveryQuestionCredential) {
	o.RecoveryQuestion = &v
}

func (o OrgCreationAdminCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RecoveryQuestion != nil {
		toSerialize["recovery_question"] = o.RecoveryQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgCreationAdminCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varOrgCreationAdminCredentials := _OrgCreationAdminCredentials{}

	err = json.Unmarshal(bytes, &varOrgCreationAdminCredentials)
	if err == nil {
		*o = OrgCreationAdminCredentials(varOrgCreationAdminCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "recovery_question")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgCreationAdminCredentials struct {
	value *OrgCreationAdminCredentials
	isSet bool
}

func (v NullableOrgCreationAdminCredentials) Get() *OrgCreationAdminCredentials {
	return v.value
}

func (v *NullableOrgCreationAdminCredentials) Set(val *OrgCreationAdminCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCreationAdminCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCreationAdminCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCreationAdminCredentials(val *OrgCreationAdminCredentials) *NullableOrgCreationAdminCredentials {
	return &NullableOrgCreationAdminCredentials{value: val, isSet: true}
}

func (v NullableOrgCreationAdminCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCreationAdminCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


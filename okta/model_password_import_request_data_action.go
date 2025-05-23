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

// PasswordImportRequestDataAction This object specifies the default action Okta is set to take. Okta takes this action if your external service sends an empty HTTP 204 response. You can override the default action by returning a commands object in your response specifying the action to take.
type PasswordImportRequestDataAction struct {
	// The status of the user credential, either `UNVERIFIED` or `VERIFIED`
	Credential *string `json:"credential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequestDataAction PasswordImportRequestDataAction

// NewPasswordImportRequestDataAction instantiates a new PasswordImportRequestDataAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequestDataAction() *PasswordImportRequestDataAction {
	this := PasswordImportRequestDataAction{}
	var credential string = "UNVERIFIED"
	this.Credential = &credential
	return &this
}

// NewPasswordImportRequestDataActionWithDefaults instantiates a new PasswordImportRequestDataAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestDataActionWithDefaults() *PasswordImportRequestDataAction {
	this := PasswordImportRequestDataAction{}
	var credential string = "UNVERIFIED"
	this.Credential = &credential
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *PasswordImportRequestDataAction) GetCredential() string {
	if o == nil || o.Credential == nil {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestDataAction) GetCredentialOk() (*string, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *PasswordImportRequestDataAction) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *PasswordImportRequestDataAction) SetCredential(v string) {
	o.Credential = &v
}

func (o PasswordImportRequestDataAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportRequestDataAction) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportRequestDataAction := _PasswordImportRequestDataAction{}

	err = json.Unmarshal(bytes, &varPasswordImportRequestDataAction)
	if err == nil {
		*o = PasswordImportRequestDataAction(varPasswordImportRequestDataAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credential")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportRequestDataAction struct {
	value *PasswordImportRequestDataAction
	isSet bool
}

func (v NullablePasswordImportRequestDataAction) Get() *PasswordImportRequestDataAction {
	return v.value
}

func (v *NullablePasswordImportRequestDataAction) Set(val *PasswordImportRequestDataAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequestDataAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequestDataAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequestDataAction(val *PasswordImportRequestDataAction) *NullablePasswordImportRequestDataAction {
	return &NullablePasswordImportRequestDataAction{value: val, isSet: true}
}

func (v NullablePasswordImportRequestDataAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequestDataAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


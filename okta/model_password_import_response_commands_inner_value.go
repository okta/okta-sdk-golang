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

// checks if the PasswordImportResponseCommandsInnerValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordImportResponseCommandsInnerValue{}

// PasswordImportResponseCommandsInnerValue The parameter value of the command. * To indicate that the supplied credentials are valid, supply a type property set to `com.okta.action.update` together with a value property set to `{\"credential\": \"VERIFIED\"}`. * To indicate that the supplied credentials are invalid, supply a type property set to `com.okta.action.update` together with a value property set to `{\"credential\": \"UNVERIFIED\"}`. Alternatively, you can send an empty response (`204`). By default, the `data.action.credential` is always set to `UNVERIFIED`.
type PasswordImportResponseCommandsInnerValue struct {
	Credential           *string `json:"credential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportResponseCommandsInnerValue PasswordImportResponseCommandsInnerValue

// NewPasswordImportResponseCommandsInnerValue instantiates a new PasswordImportResponseCommandsInnerValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportResponseCommandsInnerValue() *PasswordImportResponseCommandsInnerValue {
	this := PasswordImportResponseCommandsInnerValue{}
	return &this
}

// NewPasswordImportResponseCommandsInnerValueWithDefaults instantiates a new PasswordImportResponseCommandsInnerValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportResponseCommandsInnerValueWithDefaults() *PasswordImportResponseCommandsInnerValue {
	this := PasswordImportResponseCommandsInnerValue{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *PasswordImportResponseCommandsInnerValue) GetCredential() string {
	if o == nil || IsNil(o.Credential) {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportResponseCommandsInnerValue) GetCredentialOk() (*string, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *PasswordImportResponseCommandsInnerValue) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *PasswordImportResponseCommandsInnerValue) SetCredential(v string) {
	o.Credential = &v
}

func (o PasswordImportResponseCommandsInnerValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordImportResponseCommandsInnerValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordImportResponseCommandsInnerValue) UnmarshalJSON(data []byte) (err error) {
	varPasswordImportResponseCommandsInnerValue := _PasswordImportResponseCommandsInnerValue{}

	err = json.Unmarshal(data, &varPasswordImportResponseCommandsInnerValue)

	if err != nil {
		return err
	}

	*o = PasswordImportResponseCommandsInnerValue(varPasswordImportResponseCommandsInnerValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credential")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordImportResponseCommandsInnerValue struct {
	value *PasswordImportResponseCommandsInnerValue
	isSet bool
}

func (v NullablePasswordImportResponseCommandsInnerValue) Get() *PasswordImportResponseCommandsInnerValue {
	return v.value
}

func (v *NullablePasswordImportResponseCommandsInnerValue) Set(val *PasswordImportResponseCommandsInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportResponseCommandsInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportResponseCommandsInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportResponseCommandsInnerValue(val *PasswordImportResponseCommandsInnerValue) *NullablePasswordImportResponseCommandsInnerValue {
	return &NullablePasswordImportResponseCommandsInnerValue{value: val, isSet: true}
}

func (v NullablePasswordImportResponseCommandsInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportResponseCommandsInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

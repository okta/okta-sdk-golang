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

// SecurePasswordStoreApplicationSettingsApplication struct for SecurePasswordStoreApplicationSettingsApplication
type SecurePasswordStoreApplicationSettingsApplication struct {
	OptionalField1 *string `json:"optionalField1,omitempty"`
	OptionalField1Value *string `json:"optionalField1Value,omitempty"`
	OptionalField2 *string `json:"optionalField2,omitempty"`
	OptionalField2Value *string `json:"optionalField2Value,omitempty"`
	OptionalField3 *string `json:"optionalField3,omitempty"`
	OptionalField3Value *string `json:"optionalField3Value,omitempty"`
	PasswordField *string `json:"passwordField,omitempty"`
	Url *string `json:"url,omitempty"`
	UsernameField *string `json:"usernameField,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurePasswordStoreApplicationSettingsApplication SecurePasswordStoreApplicationSettingsApplication

// NewSecurePasswordStoreApplicationSettingsApplication instantiates a new SecurePasswordStoreApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurePasswordStoreApplicationSettingsApplication() *SecurePasswordStoreApplicationSettingsApplication {
	this := SecurePasswordStoreApplicationSettingsApplication{}
	return &this
}

// NewSecurePasswordStoreApplicationSettingsApplicationWithDefaults instantiates a new SecurePasswordStoreApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurePasswordStoreApplicationSettingsApplicationWithDefaults() *SecurePasswordStoreApplicationSettingsApplication {
	this := SecurePasswordStoreApplicationSettingsApplication{}
	return &this
}

// GetOptionalField1 returns the OptionalField1 field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1() string {
	if o == nil || o.OptionalField1 == nil {
		var ret string
		return ret
	}
	return *o.OptionalField1
}

// GetOptionalField1Ok returns a tuple with the OptionalField1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1Ok() (*string, bool) {
	if o == nil || o.OptionalField1 == nil {
		return nil, false
	}
	return o.OptionalField1, true
}

// HasOptionalField1 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField1() bool {
	if o != nil && o.OptionalField1 != nil {
		return true
	}

	return false
}

// SetOptionalField1 gets a reference to the given string and assigns it to the OptionalField1 field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField1(v string) {
	o.OptionalField1 = &v
}

// GetOptionalField1Value returns the OptionalField1Value field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1Value() string {
	if o == nil || o.OptionalField1Value == nil {
		var ret string
		return ret
	}
	return *o.OptionalField1Value
}

// GetOptionalField1ValueOk returns a tuple with the OptionalField1Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1ValueOk() (*string, bool) {
	if o == nil || o.OptionalField1Value == nil {
		return nil, false
	}
	return o.OptionalField1Value, true
}

// HasOptionalField1Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField1Value() bool {
	if o != nil && o.OptionalField1Value != nil {
		return true
	}

	return false
}

// SetOptionalField1Value gets a reference to the given string and assigns it to the OptionalField1Value field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField1Value(v string) {
	o.OptionalField1Value = &v
}

// GetOptionalField2 returns the OptionalField2 field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2() string {
	if o == nil || o.OptionalField2 == nil {
		var ret string
		return ret
	}
	return *o.OptionalField2
}

// GetOptionalField2Ok returns a tuple with the OptionalField2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2Ok() (*string, bool) {
	if o == nil || o.OptionalField2 == nil {
		return nil, false
	}
	return o.OptionalField2, true
}

// HasOptionalField2 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField2() bool {
	if o != nil && o.OptionalField2 != nil {
		return true
	}

	return false
}

// SetOptionalField2 gets a reference to the given string and assigns it to the OptionalField2 field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField2(v string) {
	o.OptionalField2 = &v
}

// GetOptionalField2Value returns the OptionalField2Value field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2Value() string {
	if o == nil || o.OptionalField2Value == nil {
		var ret string
		return ret
	}
	return *o.OptionalField2Value
}

// GetOptionalField2ValueOk returns a tuple with the OptionalField2Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2ValueOk() (*string, bool) {
	if o == nil || o.OptionalField2Value == nil {
		return nil, false
	}
	return o.OptionalField2Value, true
}

// HasOptionalField2Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField2Value() bool {
	if o != nil && o.OptionalField2Value != nil {
		return true
	}

	return false
}

// SetOptionalField2Value gets a reference to the given string and assigns it to the OptionalField2Value field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField2Value(v string) {
	o.OptionalField2Value = &v
}

// GetOptionalField3 returns the OptionalField3 field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3() string {
	if o == nil || o.OptionalField3 == nil {
		var ret string
		return ret
	}
	return *o.OptionalField3
}

// GetOptionalField3Ok returns a tuple with the OptionalField3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3Ok() (*string, bool) {
	if o == nil || o.OptionalField3 == nil {
		return nil, false
	}
	return o.OptionalField3, true
}

// HasOptionalField3 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField3() bool {
	if o != nil && o.OptionalField3 != nil {
		return true
	}

	return false
}

// SetOptionalField3 gets a reference to the given string and assigns it to the OptionalField3 field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField3(v string) {
	o.OptionalField3 = &v
}

// GetOptionalField3Value returns the OptionalField3Value field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3Value() string {
	if o == nil || o.OptionalField3Value == nil {
		var ret string
		return ret
	}
	return *o.OptionalField3Value
}

// GetOptionalField3ValueOk returns a tuple with the OptionalField3Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3ValueOk() (*string, bool) {
	if o == nil || o.OptionalField3Value == nil {
		return nil, false
	}
	return o.OptionalField3Value, true
}

// HasOptionalField3Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField3Value() bool {
	if o != nil && o.OptionalField3Value != nil {
		return true
	}

	return false
}

// SetOptionalField3Value gets a reference to the given string and assigns it to the OptionalField3Value field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField3Value(v string) {
	o.OptionalField3Value = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetUrl(v string) {
	o.Url = &v
}

// GetUsernameField returns the UsernameField field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUsernameField() string {
	if o == nil || o.UsernameField == nil {
		var ret string
		return ret
	}
	return *o.UsernameField
}

// GetUsernameFieldOk returns a tuple with the UsernameField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUsernameFieldOk() (*string, bool) {
	if o == nil || o.UsernameField == nil {
		return nil, false
	}
	return o.UsernameField, true
}

// HasUsernameField returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasUsernameField() bool {
	if o != nil && o.UsernameField != nil {
		return true
	}

	return false
}

// SetUsernameField gets a reference to the given string and assigns it to the UsernameField field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetUsernameField(v string) {
	o.UsernameField = &v
}

func (o SecurePasswordStoreApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptionalField1 != nil {
		toSerialize["optionalField1"] = o.OptionalField1
	}
	if o.OptionalField1Value != nil {
		toSerialize["optionalField1Value"] = o.OptionalField1Value
	}
	if o.OptionalField2 != nil {
		toSerialize["optionalField2"] = o.OptionalField2
	}
	if o.OptionalField2Value != nil {
		toSerialize["optionalField2Value"] = o.OptionalField2Value
	}
	if o.OptionalField3 != nil {
		toSerialize["optionalField3"] = o.OptionalField3
	}
	if o.OptionalField3Value != nil {
		toSerialize["optionalField3Value"] = o.OptionalField3Value
	}
	if o.PasswordField != nil {
		toSerialize["passwordField"] = o.PasswordField
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UsernameField != nil {
		toSerialize["usernameField"] = o.UsernameField
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurePasswordStoreApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varSecurePasswordStoreApplicationSettingsApplication := _SecurePasswordStoreApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varSecurePasswordStoreApplicationSettingsApplication)
	if err == nil {
		*o = SecurePasswordStoreApplicationSettingsApplication(varSecurePasswordStoreApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "optionalField1")
		delete(additionalProperties, "optionalField1Value")
		delete(additionalProperties, "optionalField2")
		delete(additionalProperties, "optionalField2Value")
		delete(additionalProperties, "optionalField3")
		delete(additionalProperties, "optionalField3Value")
		delete(additionalProperties, "passwordField")
		delete(additionalProperties, "url")
		delete(additionalProperties, "usernameField")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurePasswordStoreApplicationSettingsApplication struct {
	value *SecurePasswordStoreApplicationSettingsApplication
	isSet bool
}

func (v NullableSecurePasswordStoreApplicationSettingsApplication) Get() *SecurePasswordStoreApplicationSettingsApplication {
	return v.value
}

func (v *NullableSecurePasswordStoreApplicationSettingsApplication) Set(val *SecurePasswordStoreApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurePasswordStoreApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurePasswordStoreApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurePasswordStoreApplicationSettingsApplication(val *SecurePasswordStoreApplicationSettingsApplication) *NullableSecurePasswordStoreApplicationSettingsApplication {
	return &NullableSecurePasswordStoreApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableSecurePasswordStoreApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurePasswordStoreApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


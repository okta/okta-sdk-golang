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
	"fmt"
)

// checks if the SecurePasswordStoreApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurePasswordStoreApplicationSettingsApplication{}

// SecurePasswordStoreApplicationSettingsApplication struct for SecurePasswordStoreApplicationSettingsApplication
type SecurePasswordStoreApplicationSettingsApplication struct {
	// Name of the optional parameter in the sign-in form
	OptionalField1 *string `json:"optionalField1,omitempty"`
	// Name of the optional value in the sign-in form
	OptionalField1Value *string `json:"optionalField1Value,omitempty"`
	// Name of the optional parameter in the sign-in form
	OptionalField2 *string `json:"optionalField2,omitempty"`
	// Name of the optional value in the sign-in form
	OptionalField2Value *string `json:"optionalField2Value,omitempty"`
	// Name of the optional parameter in the sign-in form
	OptionalField3 *string `json:"optionalField3,omitempty"`
	// Name of the optional value in the sign-in form
	OptionalField3Value *string `json:"optionalField3Value,omitempty"`
	// CSS selector for the **Password** field in the sign-in form
	PasswordField string `json:"passwordField"`
	// The URL of the sign-in page for this app
	Url string `json:"url"`
	// CSS selector for the **Username** field in the sign-in form
	UsernameField        string `json:"usernameField"`
	AdditionalProperties map[string]interface{}
}

type _SecurePasswordStoreApplicationSettingsApplication SecurePasswordStoreApplicationSettingsApplication

// NewSecurePasswordStoreApplicationSettingsApplication instantiates a new SecurePasswordStoreApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurePasswordStoreApplicationSettingsApplication(passwordField string, url string, usernameField string) *SecurePasswordStoreApplicationSettingsApplication {
	this := SecurePasswordStoreApplicationSettingsApplication{}
	this.PasswordField = passwordField
	this.Url = url
	this.UsernameField = usernameField
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
	if o == nil || IsNil(o.OptionalField1) {
		var ret string
		return ret
	}
	return *o.OptionalField1
}

// GetOptionalField1Ok returns a tuple with the OptionalField1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1Ok() (*string, bool) {
	if o == nil || IsNil(o.OptionalField1) {
		return nil, false
	}
	return o.OptionalField1, true
}

// HasOptionalField1 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField1() bool {
	if o != nil && !IsNil(o.OptionalField1) {
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
	if o == nil || IsNil(o.OptionalField1Value) {
		var ret string
		return ret
	}
	return *o.OptionalField1Value
}

// GetOptionalField1ValueOk returns a tuple with the OptionalField1Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField1ValueOk() (*string, bool) {
	if o == nil || IsNil(o.OptionalField1Value) {
		return nil, false
	}
	return o.OptionalField1Value, true
}

// HasOptionalField1Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField1Value() bool {
	if o != nil && !IsNil(o.OptionalField1Value) {
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
	if o == nil || IsNil(o.OptionalField2) {
		var ret string
		return ret
	}
	return *o.OptionalField2
}

// GetOptionalField2Ok returns a tuple with the OptionalField2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2Ok() (*string, bool) {
	if o == nil || IsNil(o.OptionalField2) {
		return nil, false
	}
	return o.OptionalField2, true
}

// HasOptionalField2 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField2() bool {
	if o != nil && !IsNil(o.OptionalField2) {
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
	if o == nil || IsNil(o.OptionalField2Value) {
		var ret string
		return ret
	}
	return *o.OptionalField2Value
}

// GetOptionalField2ValueOk returns a tuple with the OptionalField2Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField2ValueOk() (*string, bool) {
	if o == nil || IsNil(o.OptionalField2Value) {
		return nil, false
	}
	return o.OptionalField2Value, true
}

// HasOptionalField2Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField2Value() bool {
	if o != nil && !IsNil(o.OptionalField2Value) {
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
	if o == nil || IsNil(o.OptionalField3) {
		var ret string
		return ret
	}
	return *o.OptionalField3
}

// GetOptionalField3Ok returns a tuple with the OptionalField3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3Ok() (*string, bool) {
	if o == nil || IsNil(o.OptionalField3) {
		return nil, false
	}
	return o.OptionalField3, true
}

// HasOptionalField3 returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField3() bool {
	if o != nil && !IsNil(o.OptionalField3) {
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
	if o == nil || IsNil(o.OptionalField3Value) {
		var ret string
		return ret
	}
	return *o.OptionalField3Value
}

// GetOptionalField3ValueOk returns a tuple with the OptionalField3Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetOptionalField3ValueOk() (*string, bool) {
	if o == nil || IsNil(o.OptionalField3Value) {
		return nil, false
	}
	return o.OptionalField3Value, true
}

// HasOptionalField3Value returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) HasOptionalField3Value() bool {
	if o != nil && !IsNil(o.OptionalField3Value) {
		return true
	}

	return false
}

// SetOptionalField3Value gets a reference to the given string and assigns it to the OptionalField3Value field.
func (o *SecurePasswordStoreApplicationSettingsApplication) SetOptionalField3Value(v string) {
	o.OptionalField3Value = &v
}

// GetPasswordField returns the PasswordField field value
func (o *SecurePasswordStoreApplicationSettingsApplication) GetPasswordField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetPasswordFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordField, true
}

// SetPasswordField sets field value
func (o *SecurePasswordStoreApplicationSettingsApplication) SetPasswordField(v string) {
	o.PasswordField = v
}

// GetUrl returns the Url field value
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SecurePasswordStoreApplicationSettingsApplication) SetUrl(v string) {
	o.Url = v
}

// GetUsernameField returns the UsernameField field value
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUsernameField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameField
}

// GetUsernameFieldOk returns a tuple with the UsernameField field value
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplicationSettingsApplication) GetUsernameFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameField, true
}

// SetUsernameField sets field value
func (o *SecurePasswordStoreApplicationSettingsApplication) SetUsernameField(v string) {
	o.UsernameField = v
}

func (o SecurePasswordStoreApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurePasswordStoreApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionalField1) {
		toSerialize["optionalField1"] = o.OptionalField1
	}
	if !IsNil(o.OptionalField1Value) {
		toSerialize["optionalField1Value"] = o.OptionalField1Value
	}
	if !IsNil(o.OptionalField2) {
		toSerialize["optionalField2"] = o.OptionalField2
	}
	if !IsNil(o.OptionalField2Value) {
		toSerialize["optionalField2Value"] = o.OptionalField2Value
	}
	if !IsNil(o.OptionalField3) {
		toSerialize["optionalField3"] = o.OptionalField3
	}
	if !IsNil(o.OptionalField3Value) {
		toSerialize["optionalField3Value"] = o.OptionalField3Value
	}
	toSerialize["passwordField"] = o.PasswordField
	toSerialize["url"] = o.Url
	toSerialize["usernameField"] = o.UsernameField

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurePasswordStoreApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"passwordField",
		"url",
		"usernameField",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSecurePasswordStoreApplicationSettingsApplication := _SecurePasswordStoreApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varSecurePasswordStoreApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = SecurePasswordStoreApplicationSettingsApplication(varSecurePasswordStoreApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

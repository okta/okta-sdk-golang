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

// PasswordPolicyRecoveryEmail struct for PasswordPolicyRecoveryEmail
type PasswordPolicyRecoveryEmail struct {
	Properties *PasswordPolicyRecoveryEmailProperties `json:"properties,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyRecoveryEmail PasswordPolicyRecoveryEmail

// NewPasswordPolicyRecoveryEmail instantiates a new PasswordPolicyRecoveryEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRecoveryEmail() *PasswordPolicyRecoveryEmail {
	this := PasswordPolicyRecoveryEmail{}
	return &this
}

// NewPasswordPolicyRecoveryEmailWithDefaults instantiates a new PasswordPolicyRecoveryEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRecoveryEmailWithDefaults() *PasswordPolicyRecoveryEmail {
	this := PasswordPolicyRecoveryEmail{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryEmail) GetProperties() PasswordPolicyRecoveryEmailProperties {
	if o == nil || o.Properties == nil {
		var ret PasswordPolicyRecoveryEmailProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryEmail) GetPropertiesOk() (*PasswordPolicyRecoveryEmailProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryEmail) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given PasswordPolicyRecoveryEmailProperties and assigns it to the Properties field.
func (o *PasswordPolicyRecoveryEmail) SetProperties(v PasswordPolicyRecoveryEmailProperties) {
	o.Properties = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PasswordPolicyRecoveryEmail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRecoveryEmail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PasswordPolicyRecoveryEmail) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PasswordPolicyRecoveryEmail) SetStatus(v string) {
	o.Status = &v
}

func (o PasswordPolicyRecoveryEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyRecoveryEmail) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyRecoveryEmail := _PasswordPolicyRecoveryEmail{}

	err = json.Unmarshal(bytes, &varPasswordPolicyRecoveryEmail)
	if err == nil {
		*o = PasswordPolicyRecoveryEmail(varPasswordPolicyRecoveryEmail)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "properties")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyRecoveryEmail struct {
	value *PasswordPolicyRecoveryEmail
	isSet bool
}

func (v NullablePasswordPolicyRecoveryEmail) Get() *PasswordPolicyRecoveryEmail {
	return v.value
}

func (v *NullablePasswordPolicyRecoveryEmail) Set(val *PasswordPolicyRecoveryEmail) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRecoveryEmail) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRecoveryEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRecoveryEmail(val *PasswordPolicyRecoveryEmail) *NullablePasswordPolicyRecoveryEmail {
	return &NullablePasswordPolicyRecoveryEmail{value: val, isSet: true}
}

func (v NullablePasswordPolicyRecoveryEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRecoveryEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


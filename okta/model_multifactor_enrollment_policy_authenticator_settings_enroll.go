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

// MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll struct for MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll
type MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll struct {
	Self *string `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll() *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	this := MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll{}
	return &this
}

// NewMultifactorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults instantiates a new MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults() *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	this := MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) SetSelf(v string) {
	o.Self = &v
}

func (o MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll := _MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll)
	if err == nil {
		*o = MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll(varMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll struct {
	value *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) Get() *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) Set(val *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll(val *MultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll {
	return &NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAuthenticatorSettingsEnroll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


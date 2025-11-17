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

// checks if the EnrollmentPolicyAuthenticatorGracePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentPolicyAuthenticatorGracePeriod{}

// EnrollmentPolicyAuthenticatorGracePeriod <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the time period required to complete an authenticator enrollment or setup
type EnrollmentPolicyAuthenticatorGracePeriod struct {
	// Grace period type
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentPolicyAuthenticatorGracePeriod EnrollmentPolicyAuthenticatorGracePeriod

// NewEnrollmentPolicyAuthenticatorGracePeriod instantiates a new EnrollmentPolicyAuthenticatorGracePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentPolicyAuthenticatorGracePeriod() *EnrollmentPolicyAuthenticatorGracePeriod {
	this := EnrollmentPolicyAuthenticatorGracePeriod{}
	return &this
}

// NewEnrollmentPolicyAuthenticatorGracePeriodWithDefaults instantiates a new EnrollmentPolicyAuthenticatorGracePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentPolicyAuthenticatorGracePeriodWithDefaults() *EnrollmentPolicyAuthenticatorGracePeriod {
	this := EnrollmentPolicyAuthenticatorGracePeriod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnrollmentPolicyAuthenticatorGracePeriod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentPolicyAuthenticatorGracePeriod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnrollmentPolicyAuthenticatorGracePeriod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EnrollmentPolicyAuthenticatorGracePeriod) SetType(v string) {
	o.Type = &v
}

func (o EnrollmentPolicyAuthenticatorGracePeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentPolicyAuthenticatorGracePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnrollmentPolicyAuthenticatorGracePeriod) UnmarshalJSON(data []byte) (err error) {
	varEnrollmentPolicyAuthenticatorGracePeriod := _EnrollmentPolicyAuthenticatorGracePeriod{}

	err = json.Unmarshal(data, &varEnrollmentPolicyAuthenticatorGracePeriod)

	if err != nil {
		return err
	}

	*o = EnrollmentPolicyAuthenticatorGracePeriod(varEnrollmentPolicyAuthenticatorGracePeriod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnrollmentPolicyAuthenticatorGracePeriod struct {
	value *EnrollmentPolicyAuthenticatorGracePeriod
	isSet bool
}

func (v NullableEnrollmentPolicyAuthenticatorGracePeriod) Get() *EnrollmentPolicyAuthenticatorGracePeriod {
	return v.value
}

func (v *NullableEnrollmentPolicyAuthenticatorGracePeriod) Set(val *EnrollmentPolicyAuthenticatorGracePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentPolicyAuthenticatorGracePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentPolicyAuthenticatorGracePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentPolicyAuthenticatorGracePeriod(val *EnrollmentPolicyAuthenticatorGracePeriod) *NullableEnrollmentPolicyAuthenticatorGracePeriod {
	return &NullableEnrollmentPolicyAuthenticatorGracePeriod{value: val, isSet: true}
}

func (v NullableEnrollmentPolicyAuthenticatorGracePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentPolicyAuthenticatorGracePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

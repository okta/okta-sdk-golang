/*
Okta Admin Management API

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
	"fmt"
)

// checks if the SelfServiceRegistrationPolicyRuleRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceRegistrationPolicyRuleRegistration{}

// SelfServiceRegistrationPolicyRuleRegistration <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies whether to allow or deny the self-service registration attempt
type SelfServiceRegistrationPolicyRuleRegistration struct {
	Access               string `json:"access"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceRegistrationPolicyRuleRegistration SelfServiceRegistrationPolicyRuleRegistration

// NewSelfServiceRegistrationPolicyRuleRegistration instantiates a new SelfServiceRegistrationPolicyRuleRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationPolicyRuleRegistration(access string) *SelfServiceRegistrationPolicyRuleRegistration {
	this := SelfServiceRegistrationPolicyRuleRegistration{}
	this.Access = access
	return &this
}

// NewSelfServiceRegistrationPolicyRuleRegistrationWithDefaults instantiates a new SelfServiceRegistrationPolicyRuleRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationPolicyRuleRegistrationWithDefaults() *SelfServiceRegistrationPolicyRuleRegistration {
	this := SelfServiceRegistrationPolicyRuleRegistration{}
	return &this
}

// GetAccess returns the Access field value
func (o *SelfServiceRegistrationPolicyRuleRegistration) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationPolicyRuleRegistration) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *SelfServiceRegistrationPolicyRuleRegistration) SetAccess(v string) {
	o.Access = v
}

func (o SelfServiceRegistrationPolicyRuleRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceRegistrationPolicyRuleRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceRegistrationPolicyRuleRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access",
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

	varSelfServiceRegistrationPolicyRuleRegistration := _SelfServiceRegistrationPolicyRuleRegistration{}

	err = json.Unmarshal(data, &varSelfServiceRegistrationPolicyRuleRegistration)

	if err != nil {
		return err
	}

	*o = SelfServiceRegistrationPolicyRuleRegistration(varSelfServiceRegistrationPolicyRuleRegistration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceRegistrationPolicyRuleRegistration struct {
	value *SelfServiceRegistrationPolicyRuleRegistration
	isSet bool
}

func (v NullableSelfServiceRegistrationPolicyRuleRegistration) Get() *SelfServiceRegistrationPolicyRuleRegistration {
	return v.value
}

func (v *NullableSelfServiceRegistrationPolicyRuleRegistration) Set(val *SelfServiceRegistrationPolicyRuleRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationPolicyRuleRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationPolicyRuleRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationPolicyRuleRegistration(val *SelfServiceRegistrationPolicyRuleRegistration) *NullableSelfServiceRegistrationPolicyRuleRegistration {
	return &NullableSelfServiceRegistrationPolicyRuleRegistration{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationPolicyRuleRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationPolicyRuleRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

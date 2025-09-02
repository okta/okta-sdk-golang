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

// OktaSignOnPolicyConditions struct for OktaSignOnPolicyConditions
type OktaSignOnPolicyConditions struct {
	People *AuthenticatorEnrollmentPolicyConditionsAllOfPeople `json:"people,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyConditions OktaSignOnPolicyConditions

// NewOktaSignOnPolicyConditions instantiates a new OktaSignOnPolicyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyConditions() *OktaSignOnPolicyConditions {
	this := OktaSignOnPolicyConditions{}
	return &this
}

// NewOktaSignOnPolicyConditionsWithDefaults instantiates a new OktaSignOnPolicyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyConditionsWithDefaults() *OktaSignOnPolicyConditions {
	this := OktaSignOnPolicyConditions{}
	return &this
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *OktaSignOnPolicyConditions) GetPeople() AuthenticatorEnrollmentPolicyConditionsAllOfPeople {
	if o == nil || o.People == nil {
		var ret AuthenticatorEnrollmentPolicyConditionsAllOfPeople
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyConditions) GetPeopleOk() (*AuthenticatorEnrollmentPolicyConditionsAllOfPeople, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *OktaSignOnPolicyConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given AuthenticatorEnrollmentPolicyConditionsAllOfPeople and assigns it to the People field.
func (o *OktaSignOnPolicyConditions) SetPeople(v AuthenticatorEnrollmentPolicyConditionsAllOfPeople) {
	o.People = &v
}

func (o OktaSignOnPolicyConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.People != nil {
		toSerialize["people"] = o.People
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSignOnPolicyConditions) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSignOnPolicyConditions := _OktaSignOnPolicyConditions{}

	err = json.Unmarshal(bytes, &varOktaSignOnPolicyConditions)
	if err == nil {
		*o = OktaSignOnPolicyConditions(varOktaSignOnPolicyConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "people")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaSignOnPolicyConditions struct {
	value *OktaSignOnPolicyConditions
	isSet bool
}

func (v NullableOktaSignOnPolicyConditions) Get() *OktaSignOnPolicyConditions {
	return v.value
}

func (v *NullableOktaSignOnPolicyConditions) Set(val *OktaSignOnPolicyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyConditions(val *OktaSignOnPolicyConditions) *NullableOktaSignOnPolicyConditions {
	return &NullableOktaSignOnPolicyConditions{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


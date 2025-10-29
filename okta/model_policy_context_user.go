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
	"fmt"
)

// checks if the PolicyContextUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyContextUser{}

// PolicyContextUser The user ID for the simulate operation. Only user IDs or Group IDs are allowed, not both.
type PolicyContextUser struct {
	// The unique ID number for the user.
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContextUser PolicyContextUser

// NewPolicyContextUser instantiates a new PolicyContextUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContextUser(id string) *PolicyContextUser {
	this := PolicyContextUser{}
	this.Id = id
	return &this
}

// NewPolicyContextUserWithDefaults instantiates a new PolicyContextUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextUserWithDefaults() *PolicyContextUser {
	this := PolicyContextUser{}
	return &this
}

// GetId returns the Id field value
func (o *PolicyContextUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PolicyContextUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PolicyContextUser) SetId(v string) {
	o.Id = v
}

func (o PolicyContextUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyContextUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyContextUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPolicyContextUser := _PolicyContextUser{}

	err = json.Unmarshal(data, &varPolicyContextUser)

	if err != nil {
		return err
	}

	*o = PolicyContextUser(varPolicyContextUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyContextUser struct {
	value *PolicyContextUser
	isSet bool
}

func (v NullablePolicyContextUser) Get() *PolicyContextUser {
	return v.value
}

func (v *NullablePolicyContextUser) Set(val *PolicyContextUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContextUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContextUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContextUser(val *PolicyContextUser) *NullablePolicyContextUser {
	return &NullablePolicyContextUser{value: val, isSet: true}
}

func (v NullablePolicyContextUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContextUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

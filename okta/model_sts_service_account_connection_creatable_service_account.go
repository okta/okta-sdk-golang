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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the STSServiceAccountConnectionCreatableServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STSServiceAccountConnectionCreatableServiceAccount{}

// STSServiceAccountConnectionCreatableServiceAccount Reference to a service account in [ORN](/openapi/okta-management/guides/roles/#okta-resource-name-orn) format
type STSServiceAccountConnectionCreatableServiceAccount struct {
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the service account
	Orn string `json:"orn"`
}

type _STSServiceAccountConnectionCreatableServiceAccount STSServiceAccountConnectionCreatableServiceAccount

// NewSTSServiceAccountConnectionCreatableServiceAccount instantiates a new STSServiceAccountConnectionCreatableServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTSServiceAccountConnectionCreatableServiceAccount(orn string) *STSServiceAccountConnectionCreatableServiceAccount {
	this := STSServiceAccountConnectionCreatableServiceAccount{}
	this.Orn = orn
	return &this
}

// NewSTSServiceAccountConnectionCreatableServiceAccountWithDefaults instantiates a new STSServiceAccountConnectionCreatableServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTSServiceAccountConnectionCreatableServiceAccountWithDefaults() *STSServiceAccountConnectionCreatableServiceAccount {
	this := STSServiceAccountConnectionCreatableServiceAccount{}
	return &this
}

// GetOrn returns the Orn field value
func (o *STSServiceAccountConnectionCreatableServiceAccount) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnectionCreatableServiceAccount) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *STSServiceAccountConnectionCreatableServiceAccount) SetOrn(v string) {
	o.Orn = v
}

func (o STSServiceAccountConnectionCreatableServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STSServiceAccountConnectionCreatableServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orn"] = o.Orn
	return toSerialize, nil
}

func (o *STSServiceAccountConnectionCreatableServiceAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orn",
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

	varSTSServiceAccountConnectionCreatableServiceAccount := _STSServiceAccountConnectionCreatableServiceAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSTSServiceAccountConnectionCreatableServiceAccount)

	if err != nil {
		return err
	}

	*o = STSServiceAccountConnectionCreatableServiceAccount(varSTSServiceAccountConnectionCreatableServiceAccount)

	return err
}

type NullableSTSServiceAccountConnectionCreatableServiceAccount struct {
	value *STSServiceAccountConnectionCreatableServiceAccount
	isSet bool
}

func (v NullableSTSServiceAccountConnectionCreatableServiceAccount) Get() *STSServiceAccountConnectionCreatableServiceAccount {
	return v.value
}

func (v *NullableSTSServiceAccountConnectionCreatableServiceAccount) Set(val *STSServiceAccountConnectionCreatableServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSTSServiceAccountConnectionCreatableServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSTSServiceAccountConnectionCreatableServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTSServiceAccountConnectionCreatableServiceAccount(val *STSServiceAccountConnectionCreatableServiceAccount) *NullableSTSServiceAccountConnectionCreatableServiceAccount {
	return &NullableSTSServiceAccountConnectionCreatableServiceAccount{value: val, isSet: true}
}

func (v NullableSTSServiceAccountConnectionCreatableServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTSServiceAccountConnectionCreatableServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

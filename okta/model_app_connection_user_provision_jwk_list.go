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

// checks if the AppConnectionUserProvisionJWKList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConnectionUserProvisionJWKList{}

// AppConnectionUserProvisionJWKList struct for AppConnectionUserProvisionJWKList
type AppConnectionUserProvisionJWKList struct {
	Keys                 []JsonWebKey `json:"keys"`
	AdditionalProperties map[string]interface{}
}

type _AppConnectionUserProvisionJWKList AppConnectionUserProvisionJWKList

// NewAppConnectionUserProvisionJWKList instantiates a new AppConnectionUserProvisionJWKList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConnectionUserProvisionJWKList(keys []JsonWebKey) *AppConnectionUserProvisionJWKList {
	this := AppConnectionUserProvisionJWKList{}
	this.Keys = keys
	return &this
}

// NewAppConnectionUserProvisionJWKListWithDefaults instantiates a new AppConnectionUserProvisionJWKList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConnectionUserProvisionJWKListWithDefaults() *AppConnectionUserProvisionJWKList {
	this := AppConnectionUserProvisionJWKList{}
	return &this
}

// GetKeys returns the Keys field value
func (o *AppConnectionUserProvisionJWKList) GetKeys() []JsonWebKey {
	if o == nil {
		var ret []JsonWebKey
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *AppConnectionUserProvisionJWKList) GetKeysOk() ([]JsonWebKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *AppConnectionUserProvisionJWKList) SetKeys(v []JsonWebKey) {
	o.Keys = v
}

func (o AppConnectionUserProvisionJWKList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConnectionUserProvisionJWKList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keys"] = o.Keys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppConnectionUserProvisionJWKList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keys",
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

	varAppConnectionUserProvisionJWKList := _AppConnectionUserProvisionJWKList{}

	err = json.Unmarshal(data, &varAppConnectionUserProvisionJWKList)

	if err != nil {
		return err
	}

	*o = AppConnectionUserProvisionJWKList(varAppConnectionUserProvisionJWKList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppConnectionUserProvisionJWKList struct {
	value *AppConnectionUserProvisionJWKList
	isSet bool
}

func (v NullableAppConnectionUserProvisionJWKList) Get() *AppConnectionUserProvisionJWKList {
	return v.value
}

func (v *NullableAppConnectionUserProvisionJWKList) Set(val *AppConnectionUserProvisionJWKList) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConnectionUserProvisionJWKList) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConnectionUserProvisionJWKList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConnectionUserProvisionJWKList(val *AppConnectionUserProvisionJWKList) *NullableAppConnectionUserProvisionJWKList {
	return &NullableAppConnectionUserProvisionJWKList{value: val, isSet: true}
}

func (v NullableAppConnectionUserProvisionJWKList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConnectionUserProvisionJWKList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

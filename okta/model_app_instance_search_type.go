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

// checks if the AppInstanceSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInstanceSearchType{}

// AppInstanceSearchType An app integration type and the number of app instances of that type in the org.
type AppInstanceSearchType struct {
	// Number of app instances of this type in the org. For admins with scoped app access, the count only includes apps they can view.
	Count int32 `json:"count"`
	// App Instance type refers to the integration type of the instance.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AppInstanceSearchType AppInstanceSearchType

// NewAppInstanceSearchType instantiates a new AppInstanceSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstanceSearchType(count int32, type_ string) *AppInstanceSearchType {
	this := AppInstanceSearchType{}
	this.Count = count
	this.Type = type_
	return &this
}

// NewAppInstanceSearchTypeWithDefaults instantiates a new AppInstanceSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstanceSearchTypeWithDefaults() *AppInstanceSearchType {
	this := AppInstanceSearchType{}
	return &this
}

// GetCount returns the Count field value
func (o *AppInstanceSearchType) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AppInstanceSearchType) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AppInstanceSearchType) SetCount(v int32) {
	o.Count = v
}

// GetType returns the Type field value
func (o *AppInstanceSearchType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppInstanceSearchType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppInstanceSearchType) SetType(v string) {
	o.Type = v
}

func (o AppInstanceSearchType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInstanceSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppInstanceSearchType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"type",
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

	varAppInstanceSearchType := _AppInstanceSearchType{}

	err = json.Unmarshal(data, &varAppInstanceSearchType)

	if err != nil {
		return err
	}

	*o = AppInstanceSearchType(varAppInstanceSearchType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInstanceSearchType struct {
	value *AppInstanceSearchType
	isSet bool
}

func (v NullableAppInstanceSearchType) Get() *AppInstanceSearchType {
	return v.value
}

func (v *NullableAppInstanceSearchType) Set(val *AppInstanceSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstanceSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstanceSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstanceSearchType(val *AppInstanceSearchType) *NullableAppInstanceSearchType {
	return &NullableAppInstanceSearchType{value: val, isSet: true}
}

func (v NullableAppInstanceSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstanceSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

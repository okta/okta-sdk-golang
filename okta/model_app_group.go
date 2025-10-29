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

// checks if the AppGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppGroup{}

// AppGroup struct for AppGroup
type AppGroup struct {
	// The external ID of the app group whose members might be privileged app users
	ExternalId string `json:"externalId"`
	// The name of the app group whose members might be privileged app users
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AppGroup AppGroup

// NewAppGroup instantiates a new AppGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppGroup(externalId string, name string) *AppGroup {
	this := AppGroup{}
	this.ExternalId = externalId
	this.Name = name
	return &this
}

// NewAppGroupWithDefaults instantiates a new AppGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppGroupWithDefaults() *AppGroup {
	this := AppGroup{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *AppGroup) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AppGroup) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AppGroup) SetExternalId(v string) {
	o.ExternalId = v
}

// GetName returns the Name field value
func (o *AppGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppGroup) SetName(v string) {
	o.Name = v
}

func (o AppGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"name",
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

	varAppGroup := _AppGroup{}

	err = json.Unmarshal(data, &varAppGroup)

	if err != nil {
		return err
	}

	*o = AppGroup(varAppGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppGroup struct {
	value *AppGroup
	isSet bool
}

func (v NullableAppGroup) Get() *AppGroup {
	return v.value
}

func (v *NullableAppGroup) Set(val *AppGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAppGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAppGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppGroup(val *AppGroup) *NullableAppGroup {
	return &NullableAppGroup{value: val, isSet: true}
}

func (v NullableAppGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

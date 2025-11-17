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

// checks if the AppInstanceProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInstanceProperty{}

// AppInstanceProperty struct for AppInstanceProperty
type AppInstanceProperty struct {
	Label                string `json:"label"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AppInstanceProperty AppInstanceProperty

// NewAppInstanceProperty instantiates a new AppInstanceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInstanceProperty(label string, name string) *AppInstanceProperty {
	this := AppInstanceProperty{}
	this.Label = label
	this.Name = name
	return &this
}

// NewAppInstancePropertyWithDefaults instantiates a new AppInstanceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInstancePropertyWithDefaults() *AppInstanceProperty {
	this := AppInstanceProperty{}
	return &this
}

// GetLabel returns the Label field value
func (o *AppInstanceProperty) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AppInstanceProperty) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AppInstanceProperty) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *AppInstanceProperty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppInstanceProperty) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppInstanceProperty) SetName(v string) {
	o.Name = v
}

func (o AppInstanceProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInstanceProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppInstanceProperty) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
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

	varAppInstanceProperty := _AppInstanceProperty{}

	err = json.Unmarshal(data, &varAppInstanceProperty)

	if err != nil {
		return err
	}

	*o = AppInstanceProperty(varAppInstanceProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppInstanceProperty struct {
	value *AppInstanceProperty
	isSet bool
}

func (v NullableAppInstanceProperty) Get() *AppInstanceProperty {
	return v.value
}

func (v *NullableAppInstanceProperty) Set(val *AppInstanceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInstanceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInstanceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInstanceProperty(val *AppInstanceProperty) *NullableAppInstanceProperty {
	return &NullableAppInstanceProperty{value: val, isSet: true}
}

func (v NullableAppInstanceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInstanceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

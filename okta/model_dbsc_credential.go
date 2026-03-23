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

// checks if the DbscCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbscCredential{}

// DbscCredential struct for DbscCredential
type DbscCredential struct {
	// Cookie attributes
	Attributes string `json:"attributes"`
	// Cookie name
	Name string `json:"name"`
	// Credential type
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DbscCredential DbscCredential

// NewDbscCredential instantiates a new DbscCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbscCredential(attributes string, name string, type_ string) *DbscCredential {
	this := DbscCredential{}
	this.Attributes = attributes
	this.Name = name
	this.Type = type_
	return &this
}

// NewDbscCredentialWithDefaults instantiates a new DbscCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbscCredentialWithDefaults() *DbscCredential {
	this := DbscCredential{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *DbscCredential) GetAttributes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DbscCredential) GetAttributesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DbscCredential) SetAttributes(v string) {
	o.Attributes = v
}

// GetName returns the Name field value
func (o *DbscCredential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DbscCredential) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DbscCredential) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DbscCredential) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DbscCredential) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DbscCredential) SetType(v string) {
	o.Type = v
}

func (o DbscCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbscCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DbscCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
		"name",
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

	varDbscCredential := _DbscCredential{}

	err = json.Unmarshal(data, &varDbscCredential)

	if err != nil {
		return err
	}

	*o = DbscCredential(varDbscCredential)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDbscCredential struct {
	value *DbscCredential
	isSet bool
}

func (v NullableDbscCredential) Get() *DbscCredential {
	return v.value
}

func (v *NullableDbscCredential) Set(val *DbscCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableDbscCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableDbscCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbscCredential(val *DbscCredential) *NullableDbscCredential {
	return &NullableDbscCredential{value: val, isSet: true}
}

func (v NullableDbscCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbscCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// SamlClaimsInner struct for SamlClaimsInner
type SamlClaimsInner struct {
	// The attribute name
	Name *string `json:"name,omitempty"`
	// The Okta values inserted in the attribute statement
	Values []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlClaimsInner SamlClaimsInner

// NewSamlClaimsInner instantiates a new SamlClaimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlClaimsInner() *SamlClaimsInner {
	this := SamlClaimsInner{}
	return &this
}

// NewSamlClaimsInnerWithDefaults instantiates a new SamlClaimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlClaimsInnerWithDefaults() *SamlClaimsInner {
	this := SamlClaimsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlClaimsInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlClaimsInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlClaimsInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlClaimsInner) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SamlClaimsInner) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlClaimsInner) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SamlClaimsInner) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SamlClaimsInner) SetValues(v []string) {
	o.Values = v
}

func (o SamlClaimsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlClaimsInner) UnmarshalJSON(bytes []byte) (err error) {
	varSamlClaimsInner := _SamlClaimsInner{}

	err = json.Unmarshal(bytes, &varSamlClaimsInner)
	if err == nil {
		*o = SamlClaimsInner(varSamlClaimsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlClaimsInner struct {
	value *SamlClaimsInner
	isSet bool
}

func (v NullableSamlClaimsInner) Get() *SamlClaimsInner {
	return v.value
}

func (v *NullableSamlClaimsInner) Set(val *SamlClaimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlClaimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlClaimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlClaimsInner(val *SamlClaimsInner) *NullableSamlClaimsInner {
	return &NullableSamlClaimsInner{value: val, isSet: true}
}

func (v NullableSamlClaimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlClaimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


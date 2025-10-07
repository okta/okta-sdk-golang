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
)

// checks if the LinksFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksFactor{}

// LinksFactor struct for LinksFactor
type LinksFactor struct {
	// Link to the factor resource
	Factor               *HrefObject `json:"factor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksFactor LinksFactor

// NewLinksFactor instantiates a new LinksFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksFactor() *LinksFactor {
	this := LinksFactor{}
	return &this
}

// NewLinksFactorWithDefaults instantiates a new LinksFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksFactorWithDefaults() *LinksFactor {
	this := LinksFactor{}
	return &this
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *LinksFactor) GetFactor() HrefObject {
	if o == nil || IsNil(o.Factor) {
		var ret HrefObject
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksFactor) GetFactorOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *LinksFactor) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given HrefObject and assigns it to the Factor field.
func (o *LinksFactor) SetFactor(v HrefObject) {
	o.Factor = &v
}

func (o LinksFactor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksFactor) UnmarshalJSON(data []byte) (err error) {
	varLinksFactor := _LinksFactor{}

	err = json.Unmarshal(data, &varLinksFactor)

	if err != nil {
		return err
	}

	*o = LinksFactor(varLinksFactor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksFactor struct {
	value *LinksFactor
	isSet bool
}

func (v NullableLinksFactor) Get() *LinksFactor {
	return v.value
}

func (v *NullableLinksFactor) Set(val *LinksFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksFactor(val *LinksFactor) *NullableLinksFactor {
	return &NullableLinksFactor{value: val, isSet: true}
}

func (v NullableLinksFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

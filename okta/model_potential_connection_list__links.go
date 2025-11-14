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

// checks if the PotentialConnectionListLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PotentialConnectionListLinks{}

// PotentialConnectionListLinks Links available for the potential connection list
type PotentialConnectionListLinks struct {
	Self                 HrefObjectSelfLink  `json:"self"`
	Next                 *HrefObjectNextLink `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PotentialConnectionListLinks PotentialConnectionListLinks

// NewPotentialConnectionListLinks instantiates a new PotentialConnectionListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPotentialConnectionListLinks(self HrefObjectSelfLink) *PotentialConnectionListLinks {
	this := PotentialConnectionListLinks{}
	this.Self = self
	return &this
}

// NewPotentialConnectionListLinksWithDefaults instantiates a new PotentialConnectionListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPotentialConnectionListLinksWithDefaults() *PotentialConnectionListLinks {
	this := PotentialConnectionListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *PotentialConnectionListLinks) GetSelf() HrefObjectSelfLink {
	if o == nil {
		var ret HrefObjectSelfLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *PotentialConnectionListLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *PotentialConnectionListLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PotentialConnectionListLinks) GetNext() HrefObjectNextLink {
	if o == nil || IsNil(o.Next) {
		var ret HrefObjectNextLink
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PotentialConnectionListLinks) GetNextOk() (*HrefObjectNextLink, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PotentialConnectionListLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObjectNextLink and assigns it to the Next field.
func (o *PotentialConnectionListLinks) SetNext(v HrefObjectNextLink) {
	o.Next = &v
}

func (o PotentialConnectionListLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PotentialConnectionListLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PotentialConnectionListLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varPotentialConnectionListLinks := _PotentialConnectionListLinks{}

	err = json.Unmarshal(data, &varPotentialConnectionListLinks)

	if err != nil {
		return err
	}

	*o = PotentialConnectionListLinks(varPotentialConnectionListLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePotentialConnectionListLinks struct {
	value *PotentialConnectionListLinks
	isSet bool
}

func (v NullablePotentialConnectionListLinks) Get() *PotentialConnectionListLinks {
	return v.value
}

func (v *NullablePotentialConnectionListLinks) Set(val *PotentialConnectionListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePotentialConnectionListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePotentialConnectionListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePotentialConnectionListLinks(val *PotentialConnectionListLinks) *NullablePotentialConnectionListLinks {
	return &NullablePotentialConnectionListLinks{value: val, isSet: true}
}

func (v NullablePotentialConnectionListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePotentialConnectionListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

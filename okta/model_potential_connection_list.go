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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the PotentialConnectionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PotentialConnectionList{}

// PotentialConnectionList struct for PotentialConnectionList
type PotentialConnectionList struct {
	// Potential connections that can be established
	Data                 []PotentialConnection        `json:"data"`
	Links                PotentialConnectionListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _PotentialConnectionList PotentialConnectionList

// NewPotentialConnectionList instantiates a new PotentialConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPotentialConnectionList(data []PotentialConnection, links PotentialConnectionListLinks) *PotentialConnectionList {
	this := PotentialConnectionList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPotentialConnectionListWithDefaults instantiates a new PotentialConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPotentialConnectionListWithDefaults() *PotentialConnectionList {
	this := PotentialConnectionList{}
	return &this
}

// GetData returns the Data field value
func (o *PotentialConnectionList) GetData() []PotentialConnection {
	if o == nil {
		var ret []PotentialConnection
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PotentialConnectionList) GetDataOk() ([]PotentialConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PotentialConnectionList) SetData(v []PotentialConnection) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *PotentialConnectionList) GetLinks() PotentialConnectionListLinks {
	if o == nil {
		var ret PotentialConnectionListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PotentialConnectionList) GetLinksOk() (*PotentialConnectionListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PotentialConnectionList) SetLinks(v PotentialConnectionListLinks) {
	o.Links = v
}

func (o PotentialConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PotentialConnectionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PotentialConnectionList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
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

	varPotentialConnectionList := _PotentialConnectionList{}

	err = json.Unmarshal(data, &varPotentialConnectionList)

	if err != nil {
		return err
	}

	*o = PotentialConnectionList(varPotentialConnectionList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePotentialConnectionList struct {
	value *PotentialConnectionList
	isSet bool
}

func (v NullablePotentialConnectionList) Get() *PotentialConnectionList {
	return v.value
}

func (v *NullablePotentialConnectionList) Set(val *PotentialConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePotentialConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePotentialConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePotentialConnectionList(val *PotentialConnectionList) *NullablePotentialConnectionList {
	return &NullablePotentialConnectionList{value: val, isSet: true}
}

func (v NullablePotentialConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePotentialConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

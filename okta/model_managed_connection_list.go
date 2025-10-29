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

// checks if the ManagedConnectionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedConnectionList{}

// ManagedConnectionList struct for ManagedConnectionList
type ManagedConnectionList struct {
	// All connections the agent has established
	Data                 []ManagedConnection        `json:"data"`
	Links                ManagedConnectionListLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ManagedConnectionList ManagedConnectionList

// NewManagedConnectionList instantiates a new ManagedConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedConnectionList(data []ManagedConnection, links ManagedConnectionListLinks) *ManagedConnectionList {
	this := ManagedConnectionList{}
	this.Data = data
	this.Links = links
	return &this
}

// NewManagedConnectionListWithDefaults instantiates a new ManagedConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedConnectionListWithDefaults() *ManagedConnectionList {
	this := ManagedConnectionList{}
	return &this
}

// GetData returns the Data field value
func (o *ManagedConnectionList) GetData() []ManagedConnection {
	if o == nil {
		var ret []ManagedConnection
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionList) GetDataOk() ([]ManagedConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ManagedConnectionList) SetData(v []ManagedConnection) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ManagedConnectionList) GetLinks() ManagedConnectionListLinks {
	if o == nil {
		var ret ManagedConnectionListLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionList) GetLinksOk() (*ManagedConnectionListLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ManagedConnectionList) SetLinks(v ManagedConnectionListLinks) {
	o.Links = v
}

func (o ManagedConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedConnectionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedConnectionList) UnmarshalJSON(data []byte) (err error) {
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

	varManagedConnectionList := _ManagedConnectionList{}

	err = json.Unmarshal(data, &varManagedConnectionList)

	if err != nil {
		return err
	}

	*o = ManagedConnectionList(varManagedConnectionList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedConnectionList struct {
	value *ManagedConnectionList
	isSet bool
}

func (v NullableManagedConnectionList) Get() *ManagedConnectionList {
	return v.value
}

func (v *NullableManagedConnectionList) Set(val *ManagedConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnectionList(val *ManagedConnectionList) *NullableManagedConnectionList {
	return &NullableManagedConnectionList{value: val, isSet: true}
}

func (v NullableManagedConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

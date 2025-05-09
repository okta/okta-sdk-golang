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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PolicyContextZones The zone ID under the network rule condition.
type PolicyContextZones struct {
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContextZones PolicyContextZones

// NewPolicyContextZones instantiates a new PolicyContextZones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContextZones() *PolicyContextZones {
	this := PolicyContextZones{}
	return &this
}

// NewPolicyContextZonesWithDefaults instantiates a new PolicyContextZones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextZonesWithDefaults() *PolicyContextZones {
	this := PolicyContextZones{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *PolicyContextZones) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextZones) GetIdsOk() ([]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *PolicyContextZones) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *PolicyContextZones) SetIds(v []string) {
	o.Ids = v
}

func (o PolicyContextZones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyContextZones) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyContextZones := _PolicyContextZones{}

	err = json.Unmarshal(bytes, &varPolicyContextZones)
	if err == nil {
		*o = PolicyContextZones(varPolicyContextZones)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyContextZones struct {
	value *PolicyContextZones
	isSet bool
}

func (v NullablePolicyContextZones) Get() *PolicyContextZones {
	return v.value
}

func (v *NullablePolicyContextZones) Set(val *PolicyContextZones) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContextZones) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContextZones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContextZones(val *PolicyContextZones) *NullablePolicyContextZones {
	return &NullablePolicyContextZones{value: val, isSet: true}
}

func (v NullablePolicyContextZones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContextZones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


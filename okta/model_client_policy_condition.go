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

// ClientPolicyCondition Specifies which clients are included in the Policy
type ClientPolicyCondition struct {
	// Which clients are included in the Policy
	Include []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientPolicyCondition ClientPolicyCondition

// NewClientPolicyCondition instantiates a new ClientPolicyCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPolicyCondition() *ClientPolicyCondition {
	this := ClientPolicyCondition{}
	return &this
}

// NewClientPolicyConditionWithDefaults instantiates a new ClientPolicyCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPolicyConditionWithDefaults() *ClientPolicyCondition {
	this := ClientPolicyCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ClientPolicyCondition) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ClientPolicyCondition) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *ClientPolicyCondition) SetInclude(v []string) {
	o.Include = v
}

func (o ClientPolicyCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientPolicyCondition) UnmarshalJSON(bytes []byte) (err error) {
	varClientPolicyCondition := _ClientPolicyCondition{}

	err = json.Unmarshal(bytes, &varClientPolicyCondition)
	if err == nil {
		*o = ClientPolicyCondition(varClientPolicyCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableClientPolicyCondition struct {
	value *ClientPolicyCondition
	isSet bool
}

func (v NullableClientPolicyCondition) Get() *ClientPolicyCondition {
	return v.value
}

func (v *NullableClientPolicyCondition) Set(val *ClientPolicyCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPolicyCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPolicyCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPolicyCondition(val *ClientPolicyCondition) *NullableClientPolicyCondition {
	return &NullableClientPolicyCondition{value: val, isSet: true}
}

func (v NullableClientPolicyCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPolicyCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


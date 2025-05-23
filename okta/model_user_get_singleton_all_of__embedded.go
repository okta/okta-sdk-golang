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

// UserGetSingletonAllOfEmbedded The embedded resources related to the object if the `expand` query parameter is specified
type UserGetSingletonAllOfEmbedded struct {
	// A list of access block details for the user account
	Blocks []UserBlock `json:"blocks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserGetSingletonAllOfEmbedded UserGetSingletonAllOfEmbedded

// NewUserGetSingletonAllOfEmbedded instantiates a new UserGetSingletonAllOfEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetSingletonAllOfEmbedded() *UserGetSingletonAllOfEmbedded {
	this := UserGetSingletonAllOfEmbedded{}
	return &this
}

// NewUserGetSingletonAllOfEmbeddedWithDefaults instantiates a new UserGetSingletonAllOfEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetSingletonAllOfEmbeddedWithDefaults() *UserGetSingletonAllOfEmbedded {
	this := UserGetSingletonAllOfEmbedded{}
	return &this
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *UserGetSingletonAllOfEmbedded) GetBlocks() []UserBlock {
	if o == nil || o.Blocks == nil {
		var ret []UserBlock
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGetSingletonAllOfEmbedded) GetBlocksOk() ([]UserBlock, bool) {
	if o == nil || o.Blocks == nil {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *UserGetSingletonAllOfEmbedded) HasBlocks() bool {
	if o != nil && o.Blocks != nil {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []UserBlock and assigns it to the Blocks field.
func (o *UserGetSingletonAllOfEmbedded) SetBlocks(v []UserBlock) {
	o.Blocks = v
}

func (o UserGetSingletonAllOfEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Blocks != nil {
		toSerialize["blocks"] = o.Blocks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserGetSingletonAllOfEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varUserGetSingletonAllOfEmbedded := _UserGetSingletonAllOfEmbedded{}

	err = json.Unmarshal(bytes, &varUserGetSingletonAllOfEmbedded)
	if err == nil {
		*o = UserGetSingletonAllOfEmbedded(varUserGetSingletonAllOfEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "blocks")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserGetSingletonAllOfEmbedded struct {
	value *UserGetSingletonAllOfEmbedded
	isSet bool
}

func (v NullableUserGetSingletonAllOfEmbedded) Get() *UserGetSingletonAllOfEmbedded {
	return v.value
}

func (v *NullableUserGetSingletonAllOfEmbedded) Set(val *UserGetSingletonAllOfEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetSingletonAllOfEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetSingletonAllOfEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetSingletonAllOfEmbedded(val *UserGetSingletonAllOfEmbedded) *NullableUserGetSingletonAllOfEmbedded {
	return &NullableUserGetSingletonAllOfEmbedded{value: val, isSet: true}
}

func (v NullableUserGetSingletonAllOfEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetSingletonAllOfEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


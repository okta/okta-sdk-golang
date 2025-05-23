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

// UserBlock The description of the access block
type UserBlock struct {
	// The devices that the block applies to
	AppliesTo *string `json:"appliesTo,omitempty"`
	// Type of access block
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserBlock UserBlock

// NewUserBlock instantiates a new UserBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBlock() *UserBlock {
	this := UserBlock{}
	return &this
}

// NewUserBlockWithDefaults instantiates a new UserBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBlockWithDefaults() *UserBlock {
	this := UserBlock{}
	return &this
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *UserBlock) GetAppliesTo() string {
	if o == nil || o.AppliesTo == nil {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBlock) GetAppliesToOk() (*string, bool) {
	if o == nil || o.AppliesTo == nil {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *UserBlock) HasAppliesTo() bool {
	if o != nil && o.AppliesTo != nil {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *UserBlock) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserBlock) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserBlock) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserBlock) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserBlock) SetType(v string) {
	o.Type = &v
}

func (o UserBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppliesTo != nil {
		toSerialize["appliesTo"] = o.AppliesTo
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserBlock) UnmarshalJSON(bytes []byte) (err error) {
	varUserBlock := _UserBlock{}

	err = json.Unmarshal(bytes, &varUserBlock)
	if err == nil {
		*o = UserBlock(varUserBlock)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appliesTo")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserBlock struct {
	value *UserBlock
	isSet bool
}

func (v NullableUserBlock) Get() *UserBlock {
	return v.value
}

func (v *NullableUserBlock) Set(val *UserBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBlock(val *UserBlock) *NullableUserBlock {
	return &NullableUserBlock{value: val, isSet: true}
}

func (v NullableUserBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


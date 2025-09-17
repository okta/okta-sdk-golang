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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ResourceServerJsonWebKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceServerJsonWebKeys{}

// ResourceServerJsonWebKeys <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>A [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) for encrypting JWTs minted by the custom authorization server
type ResourceServerJsonWebKeys struct {
	Keys                 []ResourceServerJsonWebKey `json:"keys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceServerJsonWebKeys ResourceServerJsonWebKeys

// NewResourceServerJsonWebKeys instantiates a new ResourceServerJsonWebKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceServerJsonWebKeys() *ResourceServerJsonWebKeys {
	this := ResourceServerJsonWebKeys{}
	return &this
}

// NewResourceServerJsonWebKeysWithDefaults instantiates a new ResourceServerJsonWebKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceServerJsonWebKeysWithDefaults() *ResourceServerJsonWebKeys {
	this := ResourceServerJsonWebKeys{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKeys) GetKeys() []ResourceServerJsonWebKey {
	if o == nil || IsNil(o.Keys) {
		var ret []ResourceServerJsonWebKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKeys) GetKeysOk() ([]ResourceServerJsonWebKey, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKeys) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []ResourceServerJsonWebKey and assigns it to the Keys field.
func (o *ResourceServerJsonWebKeys) SetKeys(v []ResourceServerJsonWebKey) {
	o.Keys = v
}

func (o ResourceServerJsonWebKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceServerJsonWebKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceServerJsonWebKeys) UnmarshalJSON(data []byte) (err error) {
	varResourceServerJsonWebKeys := _ResourceServerJsonWebKeys{}

	err = json.Unmarshal(data, &varResourceServerJsonWebKeys)

	if err != nil {
		return err
	}

	*o = ResourceServerJsonWebKeys(varResourceServerJsonWebKeys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceServerJsonWebKeys struct {
	value *ResourceServerJsonWebKeys
	isSet bool
}

func (v NullableResourceServerJsonWebKeys) Get() *ResourceServerJsonWebKeys {
	return v.value
}

func (v *NullableResourceServerJsonWebKeys) Set(val *ResourceServerJsonWebKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceServerJsonWebKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceServerJsonWebKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceServerJsonWebKeys(val *ResourceServerJsonWebKeys) *NullableResourceServerJsonWebKeys {
	return &NullableResourceServerJsonWebKeys{value: val, isSet: true}
}

func (v NullableResourceServerJsonWebKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceServerJsonWebKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

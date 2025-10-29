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
)

// checks if the BaseContextUserLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseContextUserLinks{}

// BaseContextUserLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of the user. These links are used to discover what groups the user is a part of and what factors they have enrolled.
type BaseContextUserLinks struct {
	// URL to retrieve the individual user's group memberships
	Groups *HrefObject `json:"groups,omitempty"`
	// URL to retrieve individual user's factor enrollments
	Factors              *HrefObject `json:"factors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseContextUserLinks BaseContextUserLinks

// NewBaseContextUserLinks instantiates a new BaseContextUserLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContextUserLinks() *BaseContextUserLinks {
	this := BaseContextUserLinks{}
	return &this
}

// NewBaseContextUserLinksWithDefaults instantiates a new BaseContextUserLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContextUserLinksWithDefaults() *BaseContextUserLinks {
	this := BaseContextUserLinks{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *BaseContextUserLinks) GetGroups() HrefObject {
	if o == nil || IsNil(o.Groups) {
		var ret HrefObject
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserLinks) GetGroupsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *BaseContextUserLinks) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given HrefObject and assigns it to the Groups field.
func (o *BaseContextUserLinks) SetGroups(v HrefObject) {
	o.Groups = &v
}

// GetFactors returns the Factors field value if set, zero value otherwise.
func (o *BaseContextUserLinks) GetFactors() HrefObject {
	if o == nil || IsNil(o.Factors) {
		var ret HrefObject
		return ret
	}
	return *o.Factors
}

// GetFactorsOk returns a tuple with the Factors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserLinks) GetFactorsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Factors) {
		return nil, false
	}
	return o.Factors, true
}

// HasFactors returns a boolean if a field has been set.
func (o *BaseContextUserLinks) HasFactors() bool {
	if o != nil && !IsNil(o.Factors) {
		return true
	}

	return false
}

// SetFactors gets a reference to the given HrefObject and assigns it to the Factors field.
func (o *BaseContextUserLinks) SetFactors(v HrefObject) {
	o.Factors = &v
}

func (o BaseContextUserLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseContextUserLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Factors) {
		toSerialize["factors"] = o.Factors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseContextUserLinks) UnmarshalJSON(data []byte) (err error) {
	varBaseContextUserLinks := _BaseContextUserLinks{}

	err = json.Unmarshal(data, &varBaseContextUserLinks)

	if err != nil {
		return err
	}

	*o = BaseContextUserLinks(varBaseContextUserLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "factors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseContextUserLinks struct {
	value *BaseContextUserLinks
	isSet bool
}

func (v NullableBaseContextUserLinks) Get() *BaseContextUserLinks {
	return v.value
}

func (v *NullableBaseContextUserLinks) Set(val *BaseContextUserLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContextUserLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContextUserLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContextUserLinks(val *BaseContextUserLinks) *NullableBaseContextUserLinks {
	return &NullableBaseContextUserLinks{value: val, isSet: true}
}

func (v NullableBaseContextUserLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContextUserLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

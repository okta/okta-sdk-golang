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

// checks if the BulkGroupMembershipsDeleteRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkGroupMembershipsDeleteRequestBody{}

// BulkGroupMembershipsDeleteRequestBody struct for BulkGroupMembershipsDeleteRequestBody
type BulkGroupMembershipsDeleteRequestBody struct {
	// Array of group memberships that need to be deleted in Okta
	Memberships          []IdentitySourceGroupMembershipsDeleteProfileInner `json:"memberships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkGroupMembershipsDeleteRequestBody BulkGroupMembershipsDeleteRequestBody

// NewBulkGroupMembershipsDeleteRequestBody instantiates a new BulkGroupMembershipsDeleteRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGroupMembershipsDeleteRequestBody() *BulkGroupMembershipsDeleteRequestBody {
	this := BulkGroupMembershipsDeleteRequestBody{}
	return &this
}

// NewBulkGroupMembershipsDeleteRequestBodyWithDefaults instantiates a new BulkGroupMembershipsDeleteRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGroupMembershipsDeleteRequestBodyWithDefaults() *BulkGroupMembershipsDeleteRequestBody {
	this := BulkGroupMembershipsDeleteRequestBody{}
	return &this
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *BulkGroupMembershipsDeleteRequestBody) GetMemberships() []IdentitySourceGroupMembershipsDeleteProfileInner {
	if o == nil || IsNil(o.Memberships) {
		var ret []IdentitySourceGroupMembershipsDeleteProfileInner
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupMembershipsDeleteRequestBody) GetMembershipsOk() ([]IdentitySourceGroupMembershipsDeleteProfileInner, bool) {
	if o == nil || IsNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *BulkGroupMembershipsDeleteRequestBody) HasMemberships() bool {
	if o != nil && !IsNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []IdentitySourceGroupMembershipsDeleteProfileInner and assigns it to the Memberships field.
func (o *BulkGroupMembershipsDeleteRequestBody) SetMemberships(v []IdentitySourceGroupMembershipsDeleteProfileInner) {
	o.Memberships = v
}

func (o BulkGroupMembershipsDeleteRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkGroupMembershipsDeleteRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkGroupMembershipsDeleteRequestBody) UnmarshalJSON(data []byte) (err error) {
	varBulkGroupMembershipsDeleteRequestBody := _BulkGroupMembershipsDeleteRequestBody{}

	err = json.Unmarshal(data, &varBulkGroupMembershipsDeleteRequestBody)

	if err != nil {
		return err
	}

	*o = BulkGroupMembershipsDeleteRequestBody(varBulkGroupMembershipsDeleteRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "memberships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkGroupMembershipsDeleteRequestBody struct {
	value *BulkGroupMembershipsDeleteRequestBody
	isSet bool
}

func (v NullableBulkGroupMembershipsDeleteRequestBody) Get() *BulkGroupMembershipsDeleteRequestBody {
	return v.value
}

func (v *NullableBulkGroupMembershipsDeleteRequestBody) Set(val *BulkGroupMembershipsDeleteRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGroupMembershipsDeleteRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGroupMembershipsDeleteRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGroupMembershipsDeleteRequestBody(val *BulkGroupMembershipsDeleteRequestBody) *NullableBulkGroupMembershipsDeleteRequestBody {
	return &NullableBulkGroupMembershipsDeleteRequestBody{value: val, isSet: true}
}

func (v NullableBulkGroupMembershipsDeleteRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGroupMembershipsDeleteRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

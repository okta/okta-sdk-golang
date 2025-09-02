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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// BulkGroupMembershipsUpsertRequestBody struct for BulkGroupMembershipsUpsertRequestBody
type BulkGroupMembershipsUpsertRequestBody struct {
	// Array of group memberships that need to be inserted or updated in Okta
	Memberships []IdentitySourceGroupMembershipsUpsertProfileInner `json:"memberships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkGroupMembershipsUpsertRequestBody BulkGroupMembershipsUpsertRequestBody

// NewBulkGroupMembershipsUpsertRequestBody instantiates a new BulkGroupMembershipsUpsertRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGroupMembershipsUpsertRequestBody() *BulkGroupMembershipsUpsertRequestBody {
	this := BulkGroupMembershipsUpsertRequestBody{}
	return &this
}

// NewBulkGroupMembershipsUpsertRequestBodyWithDefaults instantiates a new BulkGroupMembershipsUpsertRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGroupMembershipsUpsertRequestBodyWithDefaults() *BulkGroupMembershipsUpsertRequestBody {
	this := BulkGroupMembershipsUpsertRequestBody{}
	return &this
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *BulkGroupMembershipsUpsertRequestBody) GetMemberships() []IdentitySourceGroupMembershipsUpsertProfileInner {
	if o == nil || o.Memberships == nil {
		var ret []IdentitySourceGroupMembershipsUpsertProfileInner
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupMembershipsUpsertRequestBody) GetMembershipsOk() ([]IdentitySourceGroupMembershipsUpsertProfileInner, bool) {
	if o == nil || o.Memberships == nil {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *BulkGroupMembershipsUpsertRequestBody) HasMemberships() bool {
	if o != nil && o.Memberships != nil {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []IdentitySourceGroupMembershipsUpsertProfileInner and assigns it to the Memberships field.
func (o *BulkGroupMembershipsUpsertRequestBody) SetMemberships(v []IdentitySourceGroupMembershipsUpsertProfileInner) {
	o.Memberships = v
}

func (o BulkGroupMembershipsUpsertRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Memberships != nil {
		toSerialize["memberships"] = o.Memberships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkGroupMembershipsUpsertRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varBulkGroupMembershipsUpsertRequestBody := _BulkGroupMembershipsUpsertRequestBody{}

	err = json.Unmarshal(bytes, &varBulkGroupMembershipsUpsertRequestBody)
	if err == nil {
		*o = BulkGroupMembershipsUpsertRequestBody(varBulkGroupMembershipsUpsertRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "memberships")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBulkGroupMembershipsUpsertRequestBody struct {
	value *BulkGroupMembershipsUpsertRequestBody
	isSet bool
}

func (v NullableBulkGroupMembershipsUpsertRequestBody) Get() *BulkGroupMembershipsUpsertRequestBody {
	return v.value
}

func (v *NullableBulkGroupMembershipsUpsertRequestBody) Set(val *BulkGroupMembershipsUpsertRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGroupMembershipsUpsertRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGroupMembershipsUpsertRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGroupMembershipsUpsertRequestBody(val *BulkGroupMembershipsUpsertRequestBody) *NullableBulkGroupMembershipsUpsertRequestBody {
	return &NullableBulkGroupMembershipsUpsertRequestBody{value: val, isSet: true}
}

func (v NullableBulkGroupMembershipsUpsertRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGroupMembershipsUpsertRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


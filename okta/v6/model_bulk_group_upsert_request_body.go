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

// BulkGroupUpsertRequestBody struct for BulkGroupUpsertRequestBody
type BulkGroupUpsertRequestBody struct {
	// Array of group profiles that needs to be inserted or updated in Okta
	Profiles []BulkGroupUpsertRequestBodyProfilesInner `json:"profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkGroupUpsertRequestBody BulkGroupUpsertRequestBody

// NewBulkGroupUpsertRequestBody instantiates a new BulkGroupUpsertRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGroupUpsertRequestBody() *BulkGroupUpsertRequestBody {
	this := BulkGroupUpsertRequestBody{}
	return &this
}

// NewBulkGroupUpsertRequestBodyWithDefaults instantiates a new BulkGroupUpsertRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGroupUpsertRequestBodyWithDefaults() *BulkGroupUpsertRequestBody {
	this := BulkGroupUpsertRequestBody{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *BulkGroupUpsertRequestBody) GetProfiles() []BulkGroupUpsertRequestBodyProfilesInner {
	if o == nil || o.Profiles == nil {
		var ret []BulkGroupUpsertRequestBodyProfilesInner
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupUpsertRequestBody) GetProfilesOk() ([]BulkGroupUpsertRequestBodyProfilesInner, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BulkGroupUpsertRequestBody) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []BulkGroupUpsertRequestBodyProfilesInner and assigns it to the Profiles field.
func (o *BulkGroupUpsertRequestBody) SetProfiles(v []BulkGroupUpsertRequestBodyProfilesInner) {
	o.Profiles = v
}

func (o BulkGroupUpsertRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkGroupUpsertRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varBulkGroupUpsertRequestBody := _BulkGroupUpsertRequestBody{}

	err = json.Unmarshal(bytes, &varBulkGroupUpsertRequestBody)
	if err == nil {
		*o = BulkGroupUpsertRequestBody(varBulkGroupUpsertRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBulkGroupUpsertRequestBody struct {
	value *BulkGroupUpsertRequestBody
	isSet bool
}

func (v NullableBulkGroupUpsertRequestBody) Get() *BulkGroupUpsertRequestBody {
	return v.value
}

func (v *NullableBulkGroupUpsertRequestBody) Set(val *BulkGroupUpsertRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGroupUpsertRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGroupUpsertRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGroupUpsertRequestBody(val *BulkGroupUpsertRequestBody) *NullableBulkGroupUpsertRequestBody {
	return &NullableBulkGroupUpsertRequestBody{value: val, isSet: true}
}

func (v NullableBulkGroupUpsertRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGroupUpsertRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


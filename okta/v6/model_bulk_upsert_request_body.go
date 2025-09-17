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

// checks if the BulkUpsertRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpsertRequestBody{}

// BulkUpsertRequestBody struct for BulkUpsertRequestBody
type BulkUpsertRequestBody struct {
	// The type of data to upsert into the session. Currently, only `USERS` is supported.
	EntityType *string `json:"entityType,omitempty"`
	// Array of user profiles to be uploaded
	Profiles             []BulkUpsertRequestBodyProfilesInner `json:"profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkUpsertRequestBody BulkUpsertRequestBody

// NewBulkUpsertRequestBody instantiates a new BulkUpsertRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpsertRequestBody() *BulkUpsertRequestBody {
	this := BulkUpsertRequestBody{}
	return &this
}

// NewBulkUpsertRequestBodyWithDefaults instantiates a new BulkUpsertRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpsertRequestBodyWithDefaults() *BulkUpsertRequestBody {
	this := BulkUpsertRequestBody{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BulkUpsertRequestBody) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpsertRequestBody) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BulkUpsertRequestBody) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BulkUpsertRequestBody) SetEntityType(v string) {
	o.EntityType = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *BulkUpsertRequestBody) GetProfiles() []BulkUpsertRequestBodyProfilesInner {
	if o == nil || IsNil(o.Profiles) {
		var ret []BulkUpsertRequestBodyProfilesInner
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpsertRequestBody) GetProfilesOk() ([]BulkUpsertRequestBodyProfilesInner, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BulkUpsertRequestBody) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []BulkUpsertRequestBodyProfilesInner and assigns it to the Profiles field.
func (o *BulkUpsertRequestBody) SetProfiles(v []BulkUpsertRequestBodyProfilesInner) {
	o.Profiles = v
}

func (o BulkUpsertRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpsertRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkUpsertRequestBody) UnmarshalJSON(data []byte) (err error) {
	varBulkUpsertRequestBody := _BulkUpsertRequestBody{}

	err = json.Unmarshal(data, &varBulkUpsertRequestBody)

	if err != nil {
		return err
	}

	*o = BulkUpsertRequestBody(varBulkUpsertRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkUpsertRequestBody struct {
	value *BulkUpsertRequestBody
	isSet bool
}

func (v NullableBulkUpsertRequestBody) Get() *BulkUpsertRequestBody {
	return v.value
}

func (v *NullableBulkUpsertRequestBody) Set(val *BulkUpsertRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpsertRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpsertRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpsertRequestBody(val *BulkUpsertRequestBody) *NullableBulkUpsertRequestBody {
	return &NullableBulkUpsertRequestBody{value: val, isSet: true}
}

func (v NullableBulkUpsertRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpsertRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

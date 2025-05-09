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

// BulkDeleteRequestBody struct for BulkDeleteRequestBody
type BulkDeleteRequestBody struct {
	EntityType *string `json:"entityType,omitempty"`
	Profiles []IdentitySourceUserProfileForDelete `json:"profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkDeleteRequestBody BulkDeleteRequestBody

// NewBulkDeleteRequestBody instantiates a new BulkDeleteRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkDeleteRequestBody() *BulkDeleteRequestBody {
	this := BulkDeleteRequestBody{}
	return &this
}

// NewBulkDeleteRequestBodyWithDefaults instantiates a new BulkDeleteRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkDeleteRequestBodyWithDefaults() *BulkDeleteRequestBody {
	this := BulkDeleteRequestBody{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BulkDeleteRequestBody) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDeleteRequestBody) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BulkDeleteRequestBody) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BulkDeleteRequestBody) SetEntityType(v string) {
	o.EntityType = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *BulkDeleteRequestBody) GetProfiles() []IdentitySourceUserProfileForDelete {
	if o == nil || o.Profiles == nil {
		var ret []IdentitySourceUserProfileForDelete
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkDeleteRequestBody) GetProfilesOk() ([]IdentitySourceUserProfileForDelete, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BulkDeleteRequestBody) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []IdentitySourceUserProfileForDelete and assigns it to the Profiles field.
func (o *BulkDeleteRequestBody) SetProfiles(v []IdentitySourceUserProfileForDelete) {
	o.Profiles = v
}

func (o BulkDeleteRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkDeleteRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varBulkDeleteRequestBody := _BulkDeleteRequestBody{}

	err = json.Unmarshal(bytes, &varBulkDeleteRequestBody)
	if err == nil {
		*o = BulkDeleteRequestBody(varBulkDeleteRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBulkDeleteRequestBody struct {
	value *BulkDeleteRequestBody
	isSet bool
}

func (v NullableBulkDeleteRequestBody) Get() *BulkDeleteRequestBody {
	return v.value
}

func (v *NullableBulkDeleteRequestBody) Set(val *BulkDeleteRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkDeleteRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkDeleteRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkDeleteRequestBody(val *BulkDeleteRequestBody) *NullableBulkDeleteRequestBody {
	return &NullableBulkDeleteRequestBody{value: val, isSet: true}
}

func (v NullableBulkDeleteRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkDeleteRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


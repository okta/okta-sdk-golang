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

// checks if the BulkGroupUpsertRequestBodyProfilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkGroupUpsertRequestBodyProfilesInner{}

// BulkGroupUpsertRequestBodyProfilesInner struct for BulkGroupUpsertRequestBodyProfilesInner
type BulkGroupUpsertRequestBodyProfilesInner struct {
	// The external ID of the group that needs to be created or updated in Okta
	ExternalId           *string                              `json:"externalId,omitempty"`
	Profile              *IdentitySourceGroupProfileForUpsert `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkGroupUpsertRequestBodyProfilesInner BulkGroupUpsertRequestBodyProfilesInner

// NewBulkGroupUpsertRequestBodyProfilesInner instantiates a new BulkGroupUpsertRequestBodyProfilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGroupUpsertRequestBodyProfilesInner() *BulkGroupUpsertRequestBodyProfilesInner {
	this := BulkGroupUpsertRequestBodyProfilesInner{}
	return &this
}

// NewBulkGroupUpsertRequestBodyProfilesInnerWithDefaults instantiates a new BulkGroupUpsertRequestBodyProfilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGroupUpsertRequestBodyProfilesInnerWithDefaults() *BulkGroupUpsertRequestBodyProfilesInner {
	this := BulkGroupUpsertRequestBodyProfilesInner{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *BulkGroupUpsertRequestBodyProfilesInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupUpsertRequestBodyProfilesInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *BulkGroupUpsertRequestBodyProfilesInner) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *BulkGroupUpsertRequestBodyProfilesInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *BulkGroupUpsertRequestBodyProfilesInner) GetProfile() IdentitySourceGroupProfileForUpsert {
	if o == nil || IsNil(o.Profile) {
		var ret IdentitySourceGroupProfileForUpsert
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupUpsertRequestBodyProfilesInner) GetProfileOk() (*IdentitySourceGroupProfileForUpsert, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *BulkGroupUpsertRequestBodyProfilesInner) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given IdentitySourceGroupProfileForUpsert and assigns it to the Profile field.
func (o *BulkGroupUpsertRequestBodyProfilesInner) SetProfile(v IdentitySourceGroupProfileForUpsert) {
	o.Profile = &v
}

func (o BulkGroupUpsertRequestBodyProfilesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkGroupUpsertRequestBodyProfilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkGroupUpsertRequestBodyProfilesInner) UnmarshalJSON(data []byte) (err error) {
	varBulkGroupUpsertRequestBodyProfilesInner := _BulkGroupUpsertRequestBodyProfilesInner{}

	err = json.Unmarshal(data, &varBulkGroupUpsertRequestBodyProfilesInner)

	if err != nil {
		return err
	}

	*o = BulkGroupUpsertRequestBodyProfilesInner(varBulkGroupUpsertRequestBodyProfilesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkGroupUpsertRequestBodyProfilesInner struct {
	value *BulkGroupUpsertRequestBodyProfilesInner
	isSet bool
}

func (v NullableBulkGroupUpsertRequestBodyProfilesInner) Get() *BulkGroupUpsertRequestBodyProfilesInner {
	return v.value
}

func (v *NullableBulkGroupUpsertRequestBodyProfilesInner) Set(val *BulkGroupUpsertRequestBodyProfilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGroupUpsertRequestBodyProfilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGroupUpsertRequestBodyProfilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGroupUpsertRequestBodyProfilesInner(val *BulkGroupUpsertRequestBodyProfilesInner) *NullableBulkGroupUpsertRequestBodyProfilesInner {
	return &NullableBulkGroupUpsertRequestBodyProfilesInner{value: val, isSet: true}
}

func (v NullableBulkGroupUpsertRequestBodyProfilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGroupUpsertRequestBodyProfilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

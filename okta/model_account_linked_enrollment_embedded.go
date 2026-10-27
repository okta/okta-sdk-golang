/*
Okta Admin Management API

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
	"time"
)

// checks if the AccountLinkedEnrollmentEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountLinkedEnrollmentEmbedded{}

// AccountLinkedEnrollmentEmbedded Embedded data for a linked enrollment
type AccountLinkedEnrollmentEmbedded struct {
	// Timestamp of the last successful authentication
	LastAuthenticatedOn *time.Time `json:"lastAuthenticatedOn,omitempty"`
	// Factors that are enrolled for offline authentication on this enrollment. `offlineFactors` is only returned for `desktop_mfa` enrollments when `offlineFactors` is requested with the `expand` query parameter.
	OfflineFactors       []OfflineFactor `json:"offlineFactors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountLinkedEnrollmentEmbedded AccountLinkedEnrollmentEmbedded

// NewAccountLinkedEnrollmentEmbedded instantiates a new AccountLinkedEnrollmentEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinkedEnrollmentEmbedded() *AccountLinkedEnrollmentEmbedded {
	this := AccountLinkedEnrollmentEmbedded{}
	return &this
}

// NewAccountLinkedEnrollmentEmbeddedWithDefaults instantiates a new AccountLinkedEnrollmentEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinkedEnrollmentEmbeddedWithDefaults() *AccountLinkedEnrollmentEmbedded {
	this := AccountLinkedEnrollmentEmbedded{}
	return &this
}

// GetLastAuthenticatedOn returns the LastAuthenticatedOn field value if set, zero value otherwise.
func (o *AccountLinkedEnrollmentEmbedded) GetLastAuthenticatedOn() time.Time {
	if o == nil || IsNil(o.LastAuthenticatedOn) {
		var ret time.Time
		return ret
	}
	return *o.LastAuthenticatedOn
}

// GetLastAuthenticatedOnOk returns a tuple with the LastAuthenticatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollmentEmbedded) GetLastAuthenticatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAuthenticatedOn) {
		return nil, false
	}
	return o.LastAuthenticatedOn, true
}

// HasLastAuthenticatedOn returns a boolean if a field has been set.
func (o *AccountLinkedEnrollmentEmbedded) HasLastAuthenticatedOn() bool {
	if o != nil && !IsNil(o.LastAuthenticatedOn) {
		return true
	}

	return false
}

// SetLastAuthenticatedOn gets a reference to the given time.Time and assigns it to the LastAuthenticatedOn field.
func (o *AccountLinkedEnrollmentEmbedded) SetLastAuthenticatedOn(v time.Time) {
	o.LastAuthenticatedOn = &v
}

// GetOfflineFactors returns the OfflineFactors field value if set, zero value otherwise.
func (o *AccountLinkedEnrollmentEmbedded) GetOfflineFactors() []OfflineFactor {
	if o == nil || IsNil(o.OfflineFactors) {
		var ret []OfflineFactor
		return ret
	}
	return o.OfflineFactors
}

// GetOfflineFactorsOk returns a tuple with the OfflineFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollmentEmbedded) GetOfflineFactorsOk() ([]OfflineFactor, bool) {
	if o == nil || IsNil(o.OfflineFactors) {
		return nil, false
	}
	return o.OfflineFactors, true
}

// HasOfflineFactors returns a boolean if a field has been set.
func (o *AccountLinkedEnrollmentEmbedded) HasOfflineFactors() bool {
	if o != nil && !IsNil(o.OfflineFactors) {
		return true
	}

	return false
}

// SetOfflineFactors gets a reference to the given []OfflineFactor and assigns it to the OfflineFactors field.
func (o *AccountLinkedEnrollmentEmbedded) SetOfflineFactors(v []OfflineFactor) {
	o.OfflineFactors = v
}

func (o AccountLinkedEnrollmentEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountLinkedEnrollmentEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastAuthenticatedOn) {
		toSerialize["lastAuthenticatedOn"] = o.LastAuthenticatedOn
	}
	if !IsNil(o.OfflineFactors) {
		toSerialize["offlineFactors"] = o.OfflineFactors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountLinkedEnrollmentEmbedded) UnmarshalJSON(data []byte) (err error) {
	varAccountLinkedEnrollmentEmbedded := _AccountLinkedEnrollmentEmbedded{}

	err = json.Unmarshal(data, &varAccountLinkedEnrollmentEmbedded)

	if err != nil {
		return err
	}

	*o = AccountLinkedEnrollmentEmbedded(varAccountLinkedEnrollmentEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lastAuthenticatedOn")
		delete(additionalProperties, "offlineFactors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountLinkedEnrollmentEmbedded struct {
	value *AccountLinkedEnrollmentEmbedded
	isSet bool
}

func (v NullableAccountLinkedEnrollmentEmbedded) Get() *AccountLinkedEnrollmentEmbedded {
	return v.value
}

func (v *NullableAccountLinkedEnrollmentEmbedded) Set(val *AccountLinkedEnrollmentEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinkedEnrollmentEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinkedEnrollmentEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinkedEnrollmentEmbedded(val *AccountLinkedEnrollmentEmbedded) *NullableAccountLinkedEnrollmentEmbedded {
	return &NullableAccountLinkedEnrollmentEmbedded{value: val, isSet: true}
}

func (v NullableAccountLinkedEnrollmentEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinkedEnrollmentEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

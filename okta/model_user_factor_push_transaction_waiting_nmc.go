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
	"reflect"
	"strings"
)

// checks if the UserFactorPushTransactionWaitingNMC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorPushTransactionWaitingNMC{}

// UserFactorPushTransactionWaitingNMC struct for UserFactorPushTransactionWaitingNMC
type UserFactorPushTransactionWaitingNMC struct {
	UserFactorPushTransaction
	Profile              *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
	Embedded             *NumberFactorChallengeEmbeddedLinks            `json:"_embedded,omitempty"`
	Links                *UserFactorPushTransactionWaitingNMCAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionWaitingNMC UserFactorPushTransactionWaitingNMC

// NewUserFactorPushTransactionWaitingNMC instantiates a new UserFactorPushTransactionWaitingNMC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionWaitingNMC() *UserFactorPushTransactionWaitingNMC {
	this := UserFactorPushTransactionWaitingNMC{}
	return &this
}

// NewUserFactorPushTransactionWaitingNMCWithDefaults instantiates a new UserFactorPushTransactionWaitingNMC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWaitingNMCWithDefaults() *UserFactorPushTransactionWaitingNMC {
	this := UserFactorPushTransactionWaitingNMC{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNMC) GetProfile() UserFactorPushTransactionRejectedAllOfProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserFactorPushTransactionRejectedAllOfProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNMC) GetProfileOk() (*UserFactorPushTransactionRejectedAllOfProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNMC) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorPushTransactionRejectedAllOfProfile and assigns it to the Profile field.
func (o *UserFactorPushTransactionWaitingNMC) SetProfile(v UserFactorPushTransactionRejectedAllOfProfile) {
	o.Profile = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNMC) GetEmbedded() NumberFactorChallengeEmbeddedLinks {
	if o == nil || IsNil(o.Embedded) {
		var ret NumberFactorChallengeEmbeddedLinks
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNMC) GetEmbeddedOk() (*NumberFactorChallengeEmbeddedLinks, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNMC) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given NumberFactorChallengeEmbeddedLinks and assigns it to the Embedded field.
func (o *UserFactorPushTransactionWaitingNMC) SetEmbedded(v NumberFactorChallengeEmbeddedLinks) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNMC) GetLinks() UserFactorPushTransactionWaitingNMCAllOfLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserFactorPushTransactionWaitingNMCAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNMC) GetLinksOk() (*UserFactorPushTransactionWaitingNMCAllOfLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNMC) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorPushTransactionWaitingNMCAllOfLinks and assigns it to the Links field.
func (o *UserFactorPushTransactionWaitingNMC) SetLinks(v UserFactorPushTransactionWaitingNMCAllOfLinks) {
	o.Links = &v
}

func (o UserFactorPushTransactionWaitingNMC) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorPushTransactionWaitingNMC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactorPushTransaction, errUserFactorPushTransaction := json.Marshal(o.UserFactorPushTransaction)
	if errUserFactorPushTransaction != nil {
		return map[string]interface{}{}, errUserFactorPushTransaction
	}
	errUserFactorPushTransaction = json.Unmarshal([]byte(serializedUserFactorPushTransaction), &toSerialize)
	if errUserFactorPushTransaction != nil {
		return map[string]interface{}{}, errUserFactorPushTransaction
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorPushTransactionWaitingNMC) UnmarshalJSON(data []byte) (err error) {
	type UserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct struct {
		Profile  *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
		Embedded *NumberFactorChallengeEmbeddedLinks            `json:"_embedded,omitempty"`
		Links    *UserFactorPushTransactionWaitingNMCAllOfLinks `json:"_links,omitempty"`
	}

	varUserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct := UserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPushTransactionWaitingNMC := _UserFactorPushTransactionWaitingNMC{}
		varUserFactorPushTransactionWaitingNMC.Profile = varUserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct.Profile
		varUserFactorPushTransactionWaitingNMC.Embedded = varUserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct.Embedded
		varUserFactorPushTransactionWaitingNMC.Links = varUserFactorPushTransactionWaitingNMCWithoutEmbeddedStruct.Links
		*o = UserFactorPushTransactionWaitingNMC(varUserFactorPushTransactionWaitingNMC)
	} else {
		return err
	}

	varUserFactorPushTransactionWaitingNMC := _UserFactorPushTransactionWaitingNMC{}

	err = json.Unmarshal(data, &varUserFactorPushTransactionWaitingNMC)
	if err == nil {
		o.UserFactorPushTransaction = varUserFactorPushTransactionWaitingNMC.UserFactorPushTransaction
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")

		// remove fields from embedded structs
		reflectUserFactorPushTransaction := reflect.ValueOf(o.UserFactorPushTransaction)
		for i := 0; i < reflectUserFactorPushTransaction.Type().NumField(); i++ {
			t := reflectUserFactorPushTransaction.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorPushTransactionWaitingNMC struct {
	value *UserFactorPushTransactionWaitingNMC
	isSet bool
}

func (v NullableUserFactorPushTransactionWaitingNMC) Get() *UserFactorPushTransactionWaitingNMC {
	return v.value
}

func (v *NullableUserFactorPushTransactionWaitingNMC) Set(val *UserFactorPushTransactionWaitingNMC) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionWaitingNMC) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionWaitingNMC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionWaitingNMC(val *UserFactorPushTransactionWaitingNMC) *NullableUserFactorPushTransactionWaitingNMC {
	return &NullableUserFactorPushTransactionWaitingNMC{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionWaitingNMC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionWaitingNMC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

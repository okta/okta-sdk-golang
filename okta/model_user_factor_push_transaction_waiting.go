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
	"reflect"
	"strings"
)

// UserFactorPushTransactionWaiting struct for UserFactorPushTransactionWaiting
type UserFactorPushTransactionWaiting struct {
	UserFactorPushTransaction
	Profile *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
	Links *UserFactorPushTransactionWaitingAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionWaiting UserFactorPushTransactionWaiting

// NewUserFactorPushTransactionWaiting instantiates a new UserFactorPushTransactionWaiting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionWaiting() *UserFactorPushTransactionWaiting {
	this := UserFactorPushTransactionWaiting{}
	return &this
}

// NewUserFactorPushTransactionWaitingWithDefaults instantiates a new UserFactorPushTransactionWaiting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWaitingWithDefaults() *UserFactorPushTransactionWaiting {
	this := UserFactorPushTransactionWaiting{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaiting) GetProfile() UserFactorPushTransactionRejectedAllOfProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorPushTransactionRejectedAllOfProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaiting) GetProfileOk() (*UserFactorPushTransactionRejectedAllOfProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaiting) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorPushTransactionRejectedAllOfProfile and assigns it to the Profile field.
func (o *UserFactorPushTransactionWaiting) SetProfile(v UserFactorPushTransactionRejectedAllOfProfile) {
	o.Profile = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaiting) GetLinks() UserFactorPushTransactionWaitingAllOfLinks {
	if o == nil || o.Links == nil {
		var ret UserFactorPushTransactionWaitingAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaiting) GetLinksOk() (*UserFactorPushTransactionWaitingAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaiting) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorPushTransactionWaitingAllOfLinks and assigns it to the Links field.
func (o *UserFactorPushTransactionWaiting) SetLinks(v UserFactorPushTransactionWaitingAllOfLinks) {
	o.Links = &v
}

func (o UserFactorPushTransactionWaiting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactorPushTransaction, errUserFactorPushTransaction := json.Marshal(o.UserFactorPushTransaction)
	if errUserFactorPushTransaction != nil {
		return []byte{}, errUserFactorPushTransaction
	}
	errUserFactorPushTransaction = json.Unmarshal([]byte(serializedUserFactorPushTransaction), &toSerialize)
	if errUserFactorPushTransaction != nil {
		return []byte{}, errUserFactorPushTransaction
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushTransactionWaiting) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorPushTransactionWaitingWithoutEmbeddedStruct struct {
		Profile *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
		Links *UserFactorPushTransactionWaitingAllOfLinks `json:"_links,omitempty"`
	}

	varUserFactorPushTransactionWaitingWithoutEmbeddedStruct := UserFactorPushTransactionWaitingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionWaitingWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPushTransactionWaiting := _UserFactorPushTransactionWaiting{}
		varUserFactorPushTransactionWaiting.Profile = varUserFactorPushTransactionWaitingWithoutEmbeddedStruct.Profile
		varUserFactorPushTransactionWaiting.Links = varUserFactorPushTransactionWaitingWithoutEmbeddedStruct.Links
		*o = UserFactorPushTransactionWaiting(varUserFactorPushTransactionWaiting)
	} else {
		return err
	}

	varUserFactorPushTransactionWaiting := _UserFactorPushTransactionWaiting{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionWaiting)
	if err == nil {
		o.UserFactorPushTransaction = varUserFactorPushTransactionWaiting.UserFactorPushTransaction
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
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
	} else {
		return err
	}

	return err
}

type NullableUserFactorPushTransactionWaiting struct {
	value *UserFactorPushTransactionWaiting
	isSet bool
}

func (v NullableUserFactorPushTransactionWaiting) Get() *UserFactorPushTransactionWaiting {
	return v.value
}

func (v *NullableUserFactorPushTransactionWaiting) Set(val *UserFactorPushTransactionWaiting) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionWaiting) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionWaiting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionWaiting(val *UserFactorPushTransactionWaiting) *NullableUserFactorPushTransactionWaiting {
	return &NullableUserFactorPushTransactionWaiting{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionWaiting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionWaiting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


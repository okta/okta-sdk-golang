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

// UserFactorPushTransactionTimeout struct for UserFactorPushTransactionTimeout
type UserFactorPushTransactionTimeout struct {
	UserFactorPushTransaction
	Profile *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
	Links *UserFactorPushTransactionRejectedAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionTimeout UserFactorPushTransactionTimeout

// NewUserFactorPushTransactionTimeout instantiates a new UserFactorPushTransactionTimeout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionTimeout() *UserFactorPushTransactionTimeout {
	this := UserFactorPushTransactionTimeout{}
	return &this
}

// NewUserFactorPushTransactionTimeoutWithDefaults instantiates a new UserFactorPushTransactionTimeout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionTimeoutWithDefaults() *UserFactorPushTransactionTimeout {
	this := UserFactorPushTransactionTimeout{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorPushTransactionTimeout) GetProfile() UserFactorPushTransactionRejectedAllOfProfile {
	if o == nil || o.Profile == nil {
		var ret UserFactorPushTransactionRejectedAllOfProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeout) GetProfileOk() (*UserFactorPushTransactionRejectedAllOfProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeout) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserFactorPushTransactionRejectedAllOfProfile and assigns it to the Profile field.
func (o *UserFactorPushTransactionTimeout) SetProfile(v UserFactorPushTransactionRejectedAllOfProfile) {
	o.Profile = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorPushTransactionTimeout) GetLinks() UserFactorPushTransactionRejectedAllOfLinks {
	if o == nil || o.Links == nil {
		var ret UserFactorPushTransactionRejectedAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeout) GetLinksOk() (*UserFactorPushTransactionRejectedAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeout) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorPushTransactionRejectedAllOfLinks and assigns it to the Links field.
func (o *UserFactorPushTransactionTimeout) SetLinks(v UserFactorPushTransactionRejectedAllOfLinks) {
	o.Links = &v
}

func (o UserFactorPushTransactionTimeout) MarshalJSON() ([]byte, error) {
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

func (o *UserFactorPushTransactionTimeout) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorPushTransactionTimeoutWithoutEmbeddedStruct struct {
		Profile *UserFactorPushTransactionRejectedAllOfProfile `json:"profile,omitempty"`
		Links *UserFactorPushTransactionRejectedAllOfLinks `json:"_links,omitempty"`
	}

	varUserFactorPushTransactionTimeoutWithoutEmbeddedStruct := UserFactorPushTransactionTimeoutWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionTimeoutWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPushTransactionTimeout := _UserFactorPushTransactionTimeout{}
		varUserFactorPushTransactionTimeout.Profile = varUserFactorPushTransactionTimeoutWithoutEmbeddedStruct.Profile
		varUserFactorPushTransactionTimeout.Links = varUserFactorPushTransactionTimeoutWithoutEmbeddedStruct.Links
		*o = UserFactorPushTransactionTimeout(varUserFactorPushTransactionTimeout)
	} else {
		return err
	}

	varUserFactorPushTransactionTimeout := _UserFactorPushTransactionTimeout{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionTimeout)
	if err == nil {
		o.UserFactorPushTransaction = varUserFactorPushTransactionTimeout.UserFactorPushTransaction
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

type NullableUserFactorPushTransactionTimeout struct {
	value *UserFactorPushTransactionTimeout
	isSet bool
}

func (v NullableUserFactorPushTransactionTimeout) Get() *UserFactorPushTransactionTimeout {
	return v.value
}

func (v *NullableUserFactorPushTransactionTimeout) Set(val *UserFactorPushTransactionTimeout) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionTimeout) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionTimeout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionTimeout(val *UserFactorPushTransactionTimeout) *NullableUserFactorPushTransactionTimeout {
	return &NullableUserFactorPushTransactionTimeout{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionTimeout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// UserFactorPushTransactionRejectedAllOfLinks struct for UserFactorPushTransactionRejectedAllOfLinks
type UserFactorPushTransactionRejectedAllOfLinks struct {
	UserFactor
	Verify *LinksVerifyVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionRejectedAllOfLinks UserFactorPushTransactionRejectedAllOfLinks

// NewUserFactorPushTransactionRejectedAllOfLinks instantiates a new UserFactorPushTransactionRejectedAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionRejectedAllOfLinks() *UserFactorPushTransactionRejectedAllOfLinks {
	this := UserFactorPushTransactionRejectedAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionRejectedAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionRejectedAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionRejectedAllOfLinksWithDefaults() *UserFactorPushTransactionRejectedAllOfLinks {
	this := UserFactorPushTransactionRejectedAllOfLinks{}
	return &this
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetVerify() LinksVerifyVerify {
	if o == nil || o.Verify == nil {
		var ret LinksVerifyVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetVerifyOk() (*LinksVerifyVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given LinksVerifyVerify and assigns it to the Verify field.
func (o *UserFactorPushTransactionRejectedAllOfLinks) SetVerify(v LinksVerifyVerify) {
	o.Verify = &v
}

func (o UserFactorPushTransactionRejectedAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushTransactionRejectedAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	type UserFactorPushTransactionRejectedAllOfLinksWithoutEmbeddedStruct struct {
		Verify *LinksVerifyVerify `json:"verify,omitempty"`
	}

	varUserFactorPushTransactionRejectedAllOfLinksWithoutEmbeddedStruct := UserFactorPushTransactionRejectedAllOfLinksWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionRejectedAllOfLinksWithoutEmbeddedStruct)
	if err == nil {
		varUserFactorPushTransactionRejectedAllOfLinks := _UserFactorPushTransactionRejectedAllOfLinks{}
		varUserFactorPushTransactionRejectedAllOfLinks.Verify = varUserFactorPushTransactionRejectedAllOfLinksWithoutEmbeddedStruct.Verify
		*o = UserFactorPushTransactionRejectedAllOfLinks(varUserFactorPushTransactionRejectedAllOfLinks)
	} else {
		return err
	}

	varUserFactorPushTransactionRejectedAllOfLinks := _UserFactorPushTransactionRejectedAllOfLinks{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionRejectedAllOfLinks)
	if err == nil {
		o.UserFactor = varUserFactorPushTransactionRejectedAllOfLinks.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "verify")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

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

type NullableUserFactorPushTransactionRejectedAllOfLinks struct {
	value *UserFactorPushTransactionRejectedAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) Get() *UserFactorPushTransactionRejectedAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) Set(val *UserFactorPushTransactionRejectedAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionRejectedAllOfLinks(val *UserFactorPushTransactionRejectedAllOfLinks) *NullableUserFactorPushTransactionRejectedAllOfLinks {
	return &NullableUserFactorPushTransactionRejectedAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


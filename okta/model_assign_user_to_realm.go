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

// AssignUserToRealm struct for AssignUserToRealm
type AssignUserToRealm struct {
	RealmId *string `json:"realmId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignUserToRealm AssignUserToRealm

// NewAssignUserToRealm instantiates a new AssignUserToRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignUserToRealm() *AssignUserToRealm {
	this := AssignUserToRealm{}
	return &this
}

// NewAssignUserToRealmWithDefaults instantiates a new AssignUserToRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignUserToRealmWithDefaults() *AssignUserToRealm {
	this := AssignUserToRealm{}
	return &this
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *AssignUserToRealm) GetRealmId() string {
	if o == nil || o.RealmId == nil {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignUserToRealm) GetRealmIdOk() (*string, bool) {
	if o == nil || o.RealmId == nil {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *AssignUserToRealm) HasRealmId() bool {
	if o != nil && o.RealmId != nil {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *AssignUserToRealm) SetRealmId(v string) {
	o.RealmId = &v
}

func (o AssignUserToRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RealmId != nil {
		toSerialize["realmId"] = o.RealmId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssignUserToRealm) UnmarshalJSON(bytes []byte) (err error) {
	varAssignUserToRealm := _AssignUserToRealm{}

	err = json.Unmarshal(bytes, &varAssignUserToRealm)
	if err == nil {
		*o = AssignUserToRealm(varAssignUserToRealm)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "realmId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAssignUserToRealm struct {
	value *AssignUserToRealm
	isSet bool
}

func (v NullableAssignUserToRealm) Get() *AssignUserToRealm {
	return v.value
}

func (v *NullableAssignUserToRealm) Set(val *AssignUserToRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignUserToRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignUserToRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignUserToRealm(val *AssignUserToRealm) *NullableAssignUserToRealm {
	return &NullableAssignUserToRealm{value: val, isSet: true}
}

func (v NullableAssignUserToRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignUserToRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// RealmProfile struct for RealmProfile
type RealmProfile struct {
	// Name of a Realm
	Name string `json:"name"`
	// Used to store partner users. This must be set to Partner to access Okta's external partner portal.
	RealmType *string `json:"realmType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmProfile RealmProfile

// NewRealmProfile instantiates a new RealmProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmProfile(name string) *RealmProfile {
	this := RealmProfile{}
	this.Name = name
	return &this
}

// NewRealmProfileWithDefaults instantiates a new RealmProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmProfileWithDefaults() *RealmProfile {
	this := RealmProfile{}
	return &this
}

// GetName returns the Name field value
func (o *RealmProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RealmProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RealmProfile) SetName(v string) {
	o.Name = v
}

// GetRealmType returns the RealmType field value if set, zero value otherwise.
func (o *RealmProfile) GetRealmType() string {
	if o == nil || o.RealmType == nil {
		var ret string
		return ret
	}
	return *o.RealmType
}

// GetRealmTypeOk returns a tuple with the RealmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmProfile) GetRealmTypeOk() (*string, bool) {
	if o == nil || o.RealmType == nil {
		return nil, false
	}
	return o.RealmType, true
}

// HasRealmType returns a boolean if a field has been set.
func (o *RealmProfile) HasRealmType() bool {
	if o != nil && o.RealmType != nil {
		return true
	}

	return false
}

// SetRealmType gets a reference to the given string and assigns it to the RealmType field.
func (o *RealmProfile) SetRealmType(v string) {
	o.RealmType = &v
}

func (o RealmProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RealmType != nil {
		toSerialize["realmType"] = o.RealmType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RealmProfile) UnmarshalJSON(bytes []byte) (err error) {
	varRealmProfile := _RealmProfile{}

	err = json.Unmarshal(bytes, &varRealmProfile)
	if err == nil {
		*o = RealmProfile(varRealmProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "realmType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRealmProfile struct {
	value *RealmProfile
	isSet bool
}

func (v NullableRealmProfile) Get() *RealmProfile {
	return v.value
}

func (v *NullableRealmProfile) Set(val *RealmProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmProfile(val *RealmProfile) *NullableRealmProfile {
	return &NullableRealmProfile{value: val, isSet: true}
}

func (v NullableRealmProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


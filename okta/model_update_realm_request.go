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

// checks if the UpdateRealmRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRealmRequest{}

// UpdateRealmRequest struct for UpdateRealmRequest
type UpdateRealmRequest struct {
	Profile              *RealmProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRealmRequest UpdateRealmRequest

// NewUpdateRealmRequest instantiates a new UpdateRealmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRealmRequest() *UpdateRealmRequest {
	this := UpdateRealmRequest{}
	return &this
}

// NewUpdateRealmRequestWithDefaults instantiates a new UpdateRealmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRealmRequestWithDefaults() *UpdateRealmRequest {
	this := UpdateRealmRequest{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UpdateRealmRequest) GetProfile() RealmProfile {
	if o == nil || IsNil(o.Profile) {
		var ret RealmProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRealmRequest) GetProfileOk() (*RealmProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UpdateRealmRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given RealmProfile and assigns it to the Profile field.
func (o *UpdateRealmRequest) SetProfile(v RealmProfile) {
	o.Profile = &v
}

func (o UpdateRealmRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRealmRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRealmRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateRealmRequest := _UpdateRealmRequest{}

	err = json.Unmarshal(data, &varUpdateRealmRequest)

	if err != nil {
		return err
	}

	*o = UpdateRealmRequest(varUpdateRealmRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRealmRequest struct {
	value *UpdateRealmRequest
	isSet bool
}

func (v NullableUpdateRealmRequest) Get() *UpdateRealmRequest {
	return v.value
}

func (v *NullableUpdateRealmRequest) Set(val *UpdateRealmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRealmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRealmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRealmRequest(val *UpdateRealmRequest) *NullableUpdateRealmRequest {
	return &NullableUpdateRealmRequest{value: val, isSet: true}
}

func (v NullableUpdateRealmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRealmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// CreateRealmRequest struct for CreateRealmRequest
type CreateRealmRequest struct {
	Profile *RealmProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRealmRequest CreateRealmRequest

// NewCreateRealmRequest instantiates a new CreateRealmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRealmRequest() *CreateRealmRequest {
	this := CreateRealmRequest{}
	return &this
}

// NewCreateRealmRequestWithDefaults instantiates a new CreateRealmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRealmRequestWithDefaults() *CreateRealmRequest {
	this := CreateRealmRequest{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *CreateRealmRequest) GetProfile() RealmProfile {
	if o == nil || o.Profile == nil {
		var ret RealmProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRealmRequest) GetProfileOk() (*RealmProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CreateRealmRequest) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given RealmProfile and assigns it to the Profile field.
func (o *CreateRealmRequest) SetProfile(v RealmProfile) {
	o.Profile = &v
}

func (o CreateRealmRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateRealmRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateRealmRequest := _CreateRealmRequest{}

	err = json.Unmarshal(bytes, &varCreateRealmRequest)
	if err == nil {
		*o = CreateRealmRequest(varCreateRealmRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateRealmRequest struct {
	value *CreateRealmRequest
	isSet bool
}

func (v NullableCreateRealmRequest) Get() *CreateRealmRequest {
	return v.value
}

func (v *NullableCreateRealmRequest) Set(val *CreateRealmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRealmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRealmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRealmRequest(val *CreateRealmRequest) *NullableCreateRealmRequest {
	return &NullableCreateRealmRequest{value: val, isSet: true}
}

func (v NullableCreateRealmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRealmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


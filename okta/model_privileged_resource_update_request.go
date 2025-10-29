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
)

// checks if the PrivilegedResourceUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegedResourceUpdateRequest{}

// PrivilegedResourceUpdateRequest Update request for a privileged resource
type PrivilegedResourceUpdateRequest struct {
	// Specific profile properties for the privileged resource
	Profile map[string]interface{} `json:"profile,omitempty"`
	// The username associated with the privileged resource
	UserName             *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceUpdateRequest PrivilegedResourceUpdateRequest

// NewPrivilegedResourceUpdateRequest instantiates a new PrivilegedResourceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceUpdateRequest() *PrivilegedResourceUpdateRequest {
	this := PrivilegedResourceUpdateRequest{}
	return &this
}

// NewPrivilegedResourceUpdateRequestWithDefaults instantiates a new PrivilegedResourceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceUpdateRequestWithDefaults() *PrivilegedResourceUpdateRequest {
	this := PrivilegedResourceUpdateRequest{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PrivilegedResourceUpdateRequest) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceUpdateRequest) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PrivilegedResourceUpdateRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *PrivilegedResourceUpdateRequest) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *PrivilegedResourceUpdateRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceUpdateRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PrivilegedResourceUpdateRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PrivilegedResourceUpdateRequest) SetUserName(v string) {
	o.UserName = &v
}

func (o PrivilegedResourceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegedResourceUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivilegedResourceUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varPrivilegedResourceUpdateRequest := _PrivilegedResourceUpdateRequest{}

	err = json.Unmarshal(data, &varPrivilegedResourceUpdateRequest)

	if err != nil {
		return err
	}

	*o = PrivilegedResourceUpdateRequest(varPrivilegedResourceUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrivilegedResourceUpdateRequest struct {
	value *PrivilegedResourceUpdateRequest
	isSet bool
}

func (v NullablePrivilegedResourceUpdateRequest) Get() *PrivilegedResourceUpdateRequest {
	return v.value
}

func (v *NullablePrivilegedResourceUpdateRequest) Set(val *PrivilegedResourceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceUpdateRequest(val *PrivilegedResourceUpdateRequest) *NullablePrivilegedResourceUpdateRequest {
	return &NullablePrivilegedResourceUpdateRequest{value: val, isSet: true}
}

func (v NullablePrivilegedResourceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PrivilegedResourceUpdateRequest Update request for a privileged resource
type PrivilegedResourceUpdateRequest struct {
	// Specific profile properties for the privileged resource
	Profile map[string]interface{} `json:"profile,omitempty"`
	// The username associated with the privileged resource
	UserName *string `json:"userName,omitempty"`
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
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceUpdateRequest) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PrivilegedResourceUpdateRequest) HasProfile() bool {
	if o != nil && o.Profile != nil {
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
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceUpdateRequest) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PrivilegedResourceUpdateRequest) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PrivilegedResourceUpdateRequest) SetUserName(v string) {
	o.UserName = &v
}

func (o PrivilegedResourceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResourceUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPrivilegedResourceUpdateRequest := _PrivilegedResourceUpdateRequest{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceUpdateRequest)
	if err == nil {
		*o = PrivilegedResourceUpdateRequest(varPrivilegedResourceUpdateRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


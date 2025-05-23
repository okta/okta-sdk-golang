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

// UpdateEmailDomain struct for UpdateEmailDomain
type UpdateEmailDomain struct {
	DisplayName string `json:"displayName"`
	UserName string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _UpdateEmailDomain UpdateEmailDomain

// NewUpdateEmailDomain instantiates a new UpdateEmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEmailDomain(displayName string, userName string) *UpdateEmailDomain {
	this := UpdateEmailDomain{}
	this.DisplayName = displayName
	this.UserName = userName
	return &this
}

// NewUpdateEmailDomainWithDefaults instantiates a new UpdateEmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEmailDomainWithDefaults() *UpdateEmailDomain {
	this := UpdateEmailDomain{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *UpdateEmailDomain) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *UpdateEmailDomain) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *UpdateEmailDomain) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUserName returns the UserName field value
func (o *UpdateEmailDomain) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *UpdateEmailDomain) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *UpdateEmailDomain) SetUserName(v string) {
	o.UserName = v
}

func (o UpdateEmailDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateEmailDomain) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateEmailDomain := _UpdateEmailDomain{}

	err = json.Unmarshal(bytes, &varUpdateEmailDomain)
	if err == nil {
		*o = UpdateEmailDomain(varUpdateEmailDomain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUpdateEmailDomain struct {
	value *UpdateEmailDomain
	isSet bool
}

func (v NullableUpdateEmailDomain) Get() *UpdateEmailDomain {
	return v.value
}

func (v *NullableUpdateEmailDomain) Set(val *UpdateEmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEmailDomain(val *UpdateEmailDomain) *NullableUpdateEmailDomain {
	return &NullableUpdateEmailDomain{value: val, isSet: true}
}

func (v NullableUpdateEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


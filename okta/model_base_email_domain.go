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

// BaseEmailDomain struct for BaseEmailDomain
type BaseEmailDomain struct {
	DisplayName string `json:"displayName"`
	UserName string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _BaseEmailDomain BaseEmailDomain

// NewBaseEmailDomain instantiates a new BaseEmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEmailDomain(displayName string, userName string) *BaseEmailDomain {
	this := BaseEmailDomain{}
	this.DisplayName = displayName
	this.UserName = userName
	return &this
}

// NewBaseEmailDomainWithDefaults instantiates a new BaseEmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEmailDomainWithDefaults() *BaseEmailDomain {
	this := BaseEmailDomain{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *BaseEmailDomain) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *BaseEmailDomain) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *BaseEmailDomain) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUserName returns the UserName field value
func (o *BaseEmailDomain) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *BaseEmailDomain) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *BaseEmailDomain) SetUserName(v string) {
	o.UserName = v
}

func (o BaseEmailDomain) MarshalJSON() ([]byte, error) {
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

func (o *BaseEmailDomain) UnmarshalJSON(bytes []byte) (err error) {
	varBaseEmailDomain := _BaseEmailDomain{}

	err = json.Unmarshal(bytes, &varBaseEmailDomain)
	if err == nil {
		*o = BaseEmailDomain(varBaseEmailDomain)
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

type NullableBaseEmailDomain struct {
	value *BaseEmailDomain
	isSet bool
}

func (v NullableBaseEmailDomain) Get() *BaseEmailDomain {
	return v.value
}

func (v *NullableBaseEmailDomain) Set(val *BaseEmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEmailDomain(val *BaseEmailDomain) *NullableBaseEmailDomain {
	return &NullableBaseEmailDomain{value: val, isSet: true}
}

func (v NullableBaseEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


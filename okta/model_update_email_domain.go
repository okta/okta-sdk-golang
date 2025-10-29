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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateEmailDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEmailDomain{}

// UpdateEmailDomain struct for UpdateEmailDomain
type UpdateEmailDomain struct {
	DisplayName          string `json:"displayName"`
	UserName             string `json:"userName"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEmailDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["userName"] = o.UserName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateEmailDomain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"userName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateEmailDomain := _UpdateEmailDomain{}

	err = json.Unmarshal(data, &varUpdateEmailDomain)

	if err != nil {
		return err
	}

	*o = UpdateEmailDomain(varUpdateEmailDomain)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
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

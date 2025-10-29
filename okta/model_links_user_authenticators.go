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

// checks if the LinksUserAuthenticators type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksUserAuthenticators{}

// LinksUserAuthenticators struct for LinksUserAuthenticators
type LinksUserAuthenticators struct {
	// Returns information about the specified user
	User                 *HrefObject `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksUserAuthenticators LinksUserAuthenticators

// NewLinksUserAuthenticators instantiates a new LinksUserAuthenticators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksUserAuthenticators() *LinksUserAuthenticators {
	this := LinksUserAuthenticators{}
	return &this
}

// NewLinksUserAuthenticatorsWithDefaults instantiates a new LinksUserAuthenticators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksUserAuthenticatorsWithDefaults() *LinksUserAuthenticators {
	this := LinksUserAuthenticators{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *LinksUserAuthenticators) GetUser() HrefObject {
	if o == nil || IsNil(o.User) {
		var ret HrefObject
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksUserAuthenticators) GetUserOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *LinksUserAuthenticators) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObject and assigns it to the User field.
func (o *LinksUserAuthenticators) SetUser(v HrefObject) {
	o.User = &v
}

func (o LinksUserAuthenticators) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksUserAuthenticators) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksUserAuthenticators) UnmarshalJSON(data []byte) (err error) {
	varLinksUserAuthenticators := _LinksUserAuthenticators{}

	err = json.Unmarshal(data, &varLinksUserAuthenticators)

	if err != nil {
		return err
	}

	*o = LinksUserAuthenticators(varLinksUserAuthenticators)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksUserAuthenticators struct {
	value *LinksUserAuthenticators
	isSet bool
}

func (v NullableLinksUserAuthenticators) Get() *LinksUserAuthenticators {
	return v.value
}

func (v *NullableLinksUserAuthenticators) Set(val *LinksUserAuthenticators) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksUserAuthenticators) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksUserAuthenticators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksUserAuthenticators(val *LinksUserAuthenticators) *NullableLinksUserAuthenticators {
	return &NullableLinksUserAuthenticators{value: val, isSet: true}
}

func (v NullableLinksUserAuthenticators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksUserAuthenticators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

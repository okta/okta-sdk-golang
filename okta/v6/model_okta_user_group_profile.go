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

// OktaUserGroupProfile Profile for any group that is not imported from Active Directory. Specifies the standard and custom profile properties for a group.  The `objectClass` for these groups is `okta:user_group`.
type OktaUserGroupProfile struct {
	// Description of the group
	Description *string `json:"description,omitempty"`
	// Name of the group
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaUserGroupProfile OktaUserGroupProfile

// NewOktaUserGroupProfile instantiates a new OktaUserGroupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaUserGroupProfile() *OktaUserGroupProfile {
	this := OktaUserGroupProfile{}
	return &this
}

// NewOktaUserGroupProfileWithDefaults instantiates a new OktaUserGroupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaUserGroupProfileWithDefaults() *OktaUserGroupProfile {
	this := OktaUserGroupProfile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OktaUserGroupProfile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserGroupProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OktaUserGroupProfile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OktaUserGroupProfile) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OktaUserGroupProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserGroupProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OktaUserGroupProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OktaUserGroupProfile) SetName(v string) {
	o.Name = &v
}

func (o OktaUserGroupProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaUserGroupProfile) UnmarshalJSON(bytes []byte) (err error) {
	varOktaUserGroupProfile := _OktaUserGroupProfile{}

	err = json.Unmarshal(bytes, &varOktaUserGroupProfile)
	if err == nil {
		*o = OktaUserGroupProfile(varOktaUserGroupProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaUserGroupProfile struct {
	value *OktaUserGroupProfile
	isSet bool
}

func (v NullableOktaUserGroupProfile) Get() *OktaUserGroupProfile {
	return v.value
}

func (v *NullableOktaUserGroupProfile) Set(val *OktaUserGroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaUserGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaUserGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaUserGroupProfile(val *OktaUserGroupProfile) *NullableOktaUserGroupProfile {
	return &NullableOktaUserGroupProfile{value: val, isSet: true}
}

func (v NullableOktaUserGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaUserGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


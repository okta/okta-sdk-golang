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
)

// checks if the CapabilitiesUpdateObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesUpdateObject{}

// CapabilitiesUpdateObject Determines whether updates to a user's profile are pushed to the app
type CapabilitiesUpdateObject struct {
	LifecycleDeactivate  *LifecycleDeactivateSettingObject `json:"lifecycleDeactivate,omitempty"`
	Password             *PasswordSettingObject            `json:"password,omitempty"`
	Profile              *ProfileSettingObject             `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesUpdateObject CapabilitiesUpdateObject

// NewCapabilitiesUpdateObject instantiates a new CapabilitiesUpdateObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesUpdateObject() *CapabilitiesUpdateObject {
	this := CapabilitiesUpdateObject{}
	return &this
}

// NewCapabilitiesUpdateObjectWithDefaults instantiates a new CapabilitiesUpdateObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesUpdateObjectWithDefaults() *CapabilitiesUpdateObject {
	this := CapabilitiesUpdateObject{}
	return &this
}

// GetLifecycleDeactivate returns the LifecycleDeactivate field value if set, zero value otherwise.
func (o *CapabilitiesUpdateObject) GetLifecycleDeactivate() LifecycleDeactivateSettingObject {
	if o == nil || IsNil(o.LifecycleDeactivate) {
		var ret LifecycleDeactivateSettingObject
		return ret
	}
	return *o.LifecycleDeactivate
}

// GetLifecycleDeactivateOk returns a tuple with the LifecycleDeactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesUpdateObject) GetLifecycleDeactivateOk() (*LifecycleDeactivateSettingObject, bool) {
	if o == nil || IsNil(o.LifecycleDeactivate) {
		return nil, false
	}
	return o.LifecycleDeactivate, true
}

// HasLifecycleDeactivate returns a boolean if a field has been set.
func (o *CapabilitiesUpdateObject) HasLifecycleDeactivate() bool {
	if o != nil && !IsNil(o.LifecycleDeactivate) {
		return true
	}

	return false
}

// SetLifecycleDeactivate gets a reference to the given LifecycleDeactivateSettingObject and assigns it to the LifecycleDeactivate field.
func (o *CapabilitiesUpdateObject) SetLifecycleDeactivate(v LifecycleDeactivateSettingObject) {
	o.LifecycleDeactivate = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CapabilitiesUpdateObject) GetPassword() PasswordSettingObject {
	if o == nil || IsNil(o.Password) {
		var ret PasswordSettingObject
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesUpdateObject) GetPasswordOk() (*PasswordSettingObject, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CapabilitiesUpdateObject) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordSettingObject and assigns it to the Password field.
func (o *CapabilitiesUpdateObject) SetPassword(v PasswordSettingObject) {
	o.Password = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *CapabilitiesUpdateObject) GetProfile() ProfileSettingObject {
	if o == nil || IsNil(o.Profile) {
		var ret ProfileSettingObject
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesUpdateObject) GetProfileOk() (*ProfileSettingObject, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CapabilitiesUpdateObject) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ProfileSettingObject and assigns it to the Profile field.
func (o *CapabilitiesUpdateObject) SetProfile(v ProfileSettingObject) {
	o.Profile = &v
}

func (o CapabilitiesUpdateObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesUpdateObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifecycleDeactivate) {
		toSerialize["lifecycleDeactivate"] = o.LifecycleDeactivate
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesUpdateObject) UnmarshalJSON(data []byte) (err error) {
	varCapabilitiesUpdateObject := _CapabilitiesUpdateObject{}

	err = json.Unmarshal(data, &varCapabilitiesUpdateObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesUpdateObject(varCapabilitiesUpdateObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lifecycleDeactivate")
		delete(additionalProperties, "password")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitiesUpdateObject struct {
	value *CapabilitiesUpdateObject
	isSet bool
}

func (v NullableCapabilitiesUpdateObject) Get() *CapabilitiesUpdateObject {
	return v.value
}

func (v *NullableCapabilitiesUpdateObject) Set(val *CapabilitiesUpdateObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesUpdateObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesUpdateObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesUpdateObject(val *CapabilitiesUpdateObject) *NullableCapabilitiesUpdateObject {
	return &NullableCapabilitiesUpdateObject{value: val, isSet: true}
}

func (v NullableCapabilitiesUpdateObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesUpdateObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

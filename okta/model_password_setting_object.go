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

// PasswordSettingObject Determines whether Okta creates and pushes a password in the application for each assigned user
type PasswordSettingObject struct {
	// Determines whether a change in a user's password also updates the user's password in the application
	Change *string `json:"change,omitempty"`
	// Determines whether the generated password is the user's Okta password or a randomly generated password
	Seed *string `json:"seed,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordSettingObject PasswordSettingObject

// NewPasswordSettingObject instantiates a new PasswordSettingObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordSettingObject() *PasswordSettingObject {
	this := PasswordSettingObject{}
	var change string = "KEEP_EXISTING"
	this.Change = &change
	var seed string = "RANDOM"
	this.Seed = &seed
	return &this
}

// NewPasswordSettingObjectWithDefaults instantiates a new PasswordSettingObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordSettingObjectWithDefaults() *PasswordSettingObject {
	this := PasswordSettingObject{}
	var change string = "KEEP_EXISTING"
	this.Change = &change
	var seed string = "RANDOM"
	this.Seed = &seed
	return &this
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *PasswordSettingObject) GetChange() string {
	if o == nil || o.Change == nil {
		var ret string
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingObject) GetChangeOk() (*string, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *PasswordSettingObject) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given string and assigns it to the Change field.
func (o *PasswordSettingObject) SetChange(v string) {
	o.Change = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *PasswordSettingObject) GetSeed() string {
	if o == nil || o.Seed == nil {
		var ret string
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingObject) GetSeedOk() (*string, bool) {
	if o == nil || o.Seed == nil {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *PasswordSettingObject) HasSeed() bool {
	if o != nil && o.Seed != nil {
		return true
	}

	return false
}

// SetSeed gets a reference to the given string and assigns it to the Seed field.
func (o *PasswordSettingObject) SetSeed(v string) {
	o.Seed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PasswordSettingObject) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordSettingObject) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PasswordSettingObject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PasswordSettingObject) SetStatus(v string) {
	o.Status = &v
}

func (o PasswordSettingObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.Seed != nil {
		toSerialize["seed"] = o.Seed
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordSettingObject) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordSettingObject := _PasswordSettingObject{}

	err = json.Unmarshal(bytes, &varPasswordSettingObject)
	if err == nil {
		*o = PasswordSettingObject(varPasswordSettingObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "change")
		delete(additionalProperties, "seed")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordSettingObject struct {
	value *PasswordSettingObject
	isSet bool
}

func (v NullablePasswordSettingObject) Get() *PasswordSettingObject {
	return v.value
}

func (v *NullablePasswordSettingObject) Set(val *PasswordSettingObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordSettingObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordSettingObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordSettingObject(val *PasswordSettingObject) *NullablePasswordSettingObject {
	return &NullablePasswordSettingObject{value: val, isSet: true}
}

func (v NullablePasswordSettingObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordSettingObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


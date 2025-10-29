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

// checks if the ProfileSettingObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSettingObject{}

// ProfileSettingObject This setting determines whether a user in the app gets updated when they're updated in Okta.  If enabled, Okta updates a user's attributes in the app when the app is assigned. Future changes made to the Okta user's profile automatically overwrite the corresponding attribute value in the app.
type ProfileSettingObject struct {
	// Setting status
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileSettingObject ProfileSettingObject

// NewProfileSettingObject instantiates a new ProfileSettingObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSettingObject() *ProfileSettingObject {
	this := ProfileSettingObject{}
	var status string = "DISABLED"
	this.Status = &status
	return &this
}

// NewProfileSettingObjectWithDefaults instantiates a new ProfileSettingObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSettingObjectWithDefaults() *ProfileSettingObject {
	this := ProfileSettingObject{}
	var status string = "DISABLED"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProfileSettingObject) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSettingObject) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProfileSettingObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProfileSettingObject) SetStatus(v string) {
	o.Status = &v
}

func (o ProfileSettingObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSettingObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileSettingObject) UnmarshalJSON(data []byte) (err error) {
	varProfileSettingObject := _ProfileSettingObject{}

	err = json.Unmarshal(data, &varProfileSettingObject)

	if err != nil {
		return err
	}

	*o = ProfileSettingObject(varProfileSettingObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileSettingObject struct {
	value *ProfileSettingObject
	isSet bool
}

func (v NullableProfileSettingObject) Get() *ProfileSettingObject {
	return v.value
}

func (v *NullableProfileSettingObject) Set(val *ProfileSettingObject) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSettingObject) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSettingObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSettingObject(val *ProfileSettingObject) *NullableProfileSettingObject {
	return &NullableProfileSettingObject{value: val, isSet: true}
}

func (v NullableProfileSettingObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSettingObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

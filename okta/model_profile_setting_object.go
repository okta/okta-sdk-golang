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

// ProfileSettingObject This setting determines whether a user in the application gets updated when they're updated in Okta.  If enabled, Okta updates a user's attributes in the application when the application is assigned. Future changes made to the Okta user's profile automatically overwrite the corresponding attribute value in the application. 
type ProfileSettingObject struct {
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileSettingObject ProfileSettingObject

// NewProfileSettingObject instantiates a new ProfileSettingObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSettingObject() *ProfileSettingObject {
	this := ProfileSettingObject{}
	return &this
}

// NewProfileSettingObjectWithDefaults instantiates a new ProfileSettingObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSettingObjectWithDefaults() *ProfileSettingObject {
	this := ProfileSettingObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProfileSettingObject) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSettingObject) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProfileSettingObject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProfileSettingObject) SetStatus(v string) {
	o.Status = &v
}

func (o ProfileSettingObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileSettingObject) UnmarshalJSON(bytes []byte) (err error) {
	varProfileSettingObject := _ProfileSettingObject{}

	err = json.Unmarshal(bytes, &varProfileSettingObject)
	if err == nil {
		*o = ProfileSettingObject(varProfileSettingObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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


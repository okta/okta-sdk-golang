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

// AutoAssignAdminAppSetting The org setting that automatically assigns the Okta Admin Console when an admin role is assigned
type AutoAssignAdminAppSetting struct {
	AutoAssignAdminAppSetting *bool `json:"autoAssignAdminAppSetting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoAssignAdminAppSetting AutoAssignAdminAppSetting

// NewAutoAssignAdminAppSetting instantiates a new AutoAssignAdminAppSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoAssignAdminAppSetting() *AutoAssignAdminAppSetting {
	this := AutoAssignAdminAppSetting{}
	return &this
}

// NewAutoAssignAdminAppSettingWithDefaults instantiates a new AutoAssignAdminAppSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoAssignAdminAppSettingWithDefaults() *AutoAssignAdminAppSetting {
	this := AutoAssignAdminAppSetting{}
	return &this
}

// GetAutoAssignAdminAppSetting returns the AutoAssignAdminAppSetting field value if set, zero value otherwise.
func (o *AutoAssignAdminAppSetting) GetAutoAssignAdminAppSetting() bool {
	if o == nil || o.AutoAssignAdminAppSetting == nil {
		var ret bool
		return ret
	}
	return *o.AutoAssignAdminAppSetting
}

// GetAutoAssignAdminAppSettingOk returns a tuple with the AutoAssignAdminAppSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoAssignAdminAppSetting) GetAutoAssignAdminAppSettingOk() (*bool, bool) {
	if o == nil || o.AutoAssignAdminAppSetting == nil {
		return nil, false
	}
	return o.AutoAssignAdminAppSetting, true
}

// HasAutoAssignAdminAppSetting returns a boolean if a field has been set.
func (o *AutoAssignAdminAppSetting) HasAutoAssignAdminAppSetting() bool {
	if o != nil && o.AutoAssignAdminAppSetting != nil {
		return true
	}

	return false
}

// SetAutoAssignAdminAppSetting gets a reference to the given bool and assigns it to the AutoAssignAdminAppSetting field.
func (o *AutoAssignAdminAppSetting) SetAutoAssignAdminAppSetting(v bool) {
	o.AutoAssignAdminAppSetting = &v
}

func (o AutoAssignAdminAppSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoAssignAdminAppSetting != nil {
		toSerialize["autoAssignAdminAppSetting"] = o.AutoAssignAdminAppSetting
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutoAssignAdminAppSetting) UnmarshalJSON(bytes []byte) (err error) {
	varAutoAssignAdminAppSetting := _AutoAssignAdminAppSetting{}

	err = json.Unmarshal(bytes, &varAutoAssignAdminAppSetting)
	if err == nil {
		*o = AutoAssignAdminAppSetting(varAutoAssignAdminAppSetting)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "autoAssignAdminAppSetting")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAutoAssignAdminAppSetting struct {
	value *AutoAssignAdminAppSetting
	isSet bool
}

func (v NullableAutoAssignAdminAppSetting) Get() *AutoAssignAdminAppSetting {
	return v.value
}

func (v *NullableAutoAssignAdminAppSetting) Set(val *AutoAssignAdminAppSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoAssignAdminAppSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoAssignAdminAppSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoAssignAdminAppSetting(val *AutoAssignAdminAppSetting) *NullableAutoAssignAdminAppSetting {
	return &NullableAutoAssignAdminAppSetting{value: val, isSet: true}
}

func (v NullableAutoAssignAdminAppSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoAssignAdminAppSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


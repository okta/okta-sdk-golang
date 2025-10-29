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

// checks if the ThirdPartyAdminSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyAdminSetting{}

// ThirdPartyAdminSetting The third-party admin setting
type ThirdPartyAdminSetting struct {
	// Indicates if the third-party admin functionality is enabled
	ThirdPartyAdmin      *bool `json:"thirdPartyAdmin,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThirdPartyAdminSetting ThirdPartyAdminSetting

// NewThirdPartyAdminSetting instantiates a new ThirdPartyAdminSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyAdminSetting() *ThirdPartyAdminSetting {
	this := ThirdPartyAdminSetting{}
	return &this
}

// NewThirdPartyAdminSettingWithDefaults instantiates a new ThirdPartyAdminSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyAdminSettingWithDefaults() *ThirdPartyAdminSetting {
	this := ThirdPartyAdminSetting{}
	return &this
}

// GetThirdPartyAdmin returns the ThirdPartyAdmin field value if set, zero value otherwise.
func (o *ThirdPartyAdminSetting) GetThirdPartyAdmin() bool {
	if o == nil || IsNil(o.ThirdPartyAdmin) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyAdmin
}

// GetThirdPartyAdminOk returns a tuple with the ThirdPartyAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAdminSetting) GetThirdPartyAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdPartyAdmin) {
		return nil, false
	}
	return o.ThirdPartyAdmin, true
}

// HasThirdPartyAdmin returns a boolean if a field has been set.
func (o *ThirdPartyAdminSetting) HasThirdPartyAdmin() bool {
	if o != nil && !IsNil(o.ThirdPartyAdmin) {
		return true
	}

	return false
}

// SetThirdPartyAdmin gets a reference to the given bool and assigns it to the ThirdPartyAdmin field.
func (o *ThirdPartyAdminSetting) SetThirdPartyAdmin(v bool) {
	o.ThirdPartyAdmin = &v
}

func (o ThirdPartyAdminSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyAdminSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThirdPartyAdmin) {
		toSerialize["thirdPartyAdmin"] = o.ThirdPartyAdmin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThirdPartyAdminSetting) UnmarshalJSON(data []byte) (err error) {
	varThirdPartyAdminSetting := _ThirdPartyAdminSetting{}

	err = json.Unmarshal(data, &varThirdPartyAdminSetting)

	if err != nil {
		return err
	}

	*o = ThirdPartyAdminSetting(varThirdPartyAdminSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "thirdPartyAdmin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThirdPartyAdminSetting struct {
	value *ThirdPartyAdminSetting
	isSet bool
}

func (v NullableThirdPartyAdminSetting) Get() *ThirdPartyAdminSetting {
	return v.value
}

func (v *NullableThirdPartyAdminSetting) Set(val *ThirdPartyAdminSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyAdminSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyAdminSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyAdminSetting(val *ThirdPartyAdminSetting) *NullableThirdPartyAdminSetting {
	return &NullableThirdPartyAdminSetting{value: val, isSet: true}
}

func (v NullableThirdPartyAdminSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyAdminSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

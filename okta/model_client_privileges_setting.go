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

// ClientPrivilegesSetting The org setting that assigns the super admin role by default to a public client app
type ClientPrivilegesSetting struct {
	ClientPrivilegesSetting *bool `json:"clientPrivilegesSetting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientPrivilegesSetting ClientPrivilegesSetting

// NewClientPrivilegesSetting instantiates a new ClientPrivilegesSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPrivilegesSetting() *ClientPrivilegesSetting {
	this := ClientPrivilegesSetting{}
	return &this
}

// NewClientPrivilegesSettingWithDefaults instantiates a new ClientPrivilegesSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPrivilegesSettingWithDefaults() *ClientPrivilegesSetting {
	this := ClientPrivilegesSetting{}
	return &this
}

// GetClientPrivilegesSetting returns the ClientPrivilegesSetting field value if set, zero value otherwise.
func (o *ClientPrivilegesSetting) GetClientPrivilegesSetting() bool {
	if o == nil || o.ClientPrivilegesSetting == nil {
		var ret bool
		return ret
	}
	return *o.ClientPrivilegesSetting
}

// GetClientPrivilegesSettingOk returns a tuple with the ClientPrivilegesSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPrivilegesSetting) GetClientPrivilegesSettingOk() (*bool, bool) {
	if o == nil || o.ClientPrivilegesSetting == nil {
		return nil, false
	}
	return o.ClientPrivilegesSetting, true
}

// HasClientPrivilegesSetting returns a boolean if a field has been set.
func (o *ClientPrivilegesSetting) HasClientPrivilegesSetting() bool {
	if o != nil && o.ClientPrivilegesSetting != nil {
		return true
	}

	return false
}

// SetClientPrivilegesSetting gets a reference to the given bool and assigns it to the ClientPrivilegesSetting field.
func (o *ClientPrivilegesSetting) SetClientPrivilegesSetting(v bool) {
	o.ClientPrivilegesSetting = &v
}

func (o ClientPrivilegesSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientPrivilegesSetting != nil {
		toSerialize["clientPrivilegesSetting"] = o.ClientPrivilegesSetting
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientPrivilegesSetting) UnmarshalJSON(bytes []byte) (err error) {
	varClientPrivilegesSetting := _ClientPrivilegesSetting{}

	err = json.Unmarshal(bytes, &varClientPrivilegesSetting)
	if err == nil {
		*o = ClientPrivilegesSetting(varClientPrivilegesSetting)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "clientPrivilegesSetting")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableClientPrivilegesSetting struct {
	value *ClientPrivilegesSetting
	isSet bool
}

func (v NullableClientPrivilegesSetting) Get() *ClientPrivilegesSetting {
	return v.value
}

func (v *NullableClientPrivilegesSetting) Set(val *ClientPrivilegesSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPrivilegesSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPrivilegesSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPrivilegesSetting(val *ClientPrivilegesSetting) *NullableClientPrivilegesSetting {
	return &NullableClientPrivilegesSetting{value: val, isSet: true}
}

func (v NullableClientPrivilegesSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPrivilegesSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


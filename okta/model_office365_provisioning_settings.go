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

// Office365ProvisioningSettings Settings required for the Microsoft Office 365 Provisioning Connection
type Office365ProvisioningSettings struct {
	// Microsoft Office 365 global administrator password
	AdminPassword string `json:"adminPassword"`
	// Microsoft Office 365 global administrator username
	AdminUsername string `json:"adminUsername"`
	AdditionalProperties map[string]interface{}
}

type _Office365ProvisioningSettings Office365ProvisioningSettings

// NewOffice365ProvisioningSettings instantiates a new Office365ProvisioningSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365ProvisioningSettings(adminPassword string, adminUsername string) *Office365ProvisioningSettings {
	this := Office365ProvisioningSettings{}
	this.AdminPassword = adminPassword
	this.AdminUsername = adminUsername
	return &this
}

// NewOffice365ProvisioningSettingsWithDefaults instantiates a new Office365ProvisioningSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365ProvisioningSettingsWithDefaults() *Office365ProvisioningSettings {
	this := Office365ProvisioningSettings{}
	return &this
}

// GetAdminPassword returns the AdminPassword field value
func (o *Office365ProvisioningSettings) GetAdminPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminPassword
}

// GetAdminPasswordOk returns a tuple with the AdminPassword field value
// and a boolean to check if the value has been set.
func (o *Office365ProvisioningSettings) GetAdminPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminPassword, true
}

// SetAdminPassword sets field value
func (o *Office365ProvisioningSettings) SetAdminPassword(v string) {
	o.AdminPassword = v
}

// GetAdminUsername returns the AdminUsername field value
func (o *Office365ProvisioningSettings) GetAdminUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdminUsername
}

// GetAdminUsernameOk returns a tuple with the AdminUsername field value
// and a boolean to check if the value has been set.
func (o *Office365ProvisioningSettings) GetAdminUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminUsername, true
}

// SetAdminUsername sets field value
func (o *Office365ProvisioningSettings) SetAdminUsername(v string) {
	o.AdminUsername = v
}

func (o Office365ProvisioningSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminPassword"] = o.AdminPassword
	}
	if true {
		toSerialize["adminUsername"] = o.AdminUsername
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Office365ProvisioningSettings) UnmarshalJSON(bytes []byte) (err error) {
	varOffice365ProvisioningSettings := _Office365ProvisioningSettings{}

	err = json.Unmarshal(bytes, &varOffice365ProvisioningSettings)
	if err == nil {
		*o = Office365ProvisioningSettings(varOffice365ProvisioningSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "adminPassword")
		delete(additionalProperties, "adminUsername")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOffice365ProvisioningSettings struct {
	value *Office365ProvisioningSettings
	isSet bool
}

func (v NullableOffice365ProvisioningSettings) Get() *Office365ProvisioningSettings {
	return v.value
}

func (v *NullableOffice365ProvisioningSettings) Set(val *Office365ProvisioningSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365ProvisioningSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365ProvisioningSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365ProvisioningSettings(val *Office365ProvisioningSettings) *NullableOffice365ProvisioningSettings {
	return &NullableOffice365ProvisioningSettings{value: val, isSet: true}
}

func (v NullableOffice365ProvisioningSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365ProvisioningSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


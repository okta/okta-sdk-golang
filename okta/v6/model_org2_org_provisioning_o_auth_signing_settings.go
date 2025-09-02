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
	"fmt"
)

// Org2OrgProvisioningOAuthSigningSettings Only used for the Okta Org2Org (`okta_org2org`) app.  The signing key rotation setting.
type Org2OrgProvisioningOAuthSigningSettings struct {
	// The signing key rotation setting for the provisioning connection
	RotationMode string `json:"rotationMode"`
	AdditionalProperties map[string]interface{}
}

type _Org2OrgProvisioningOAuthSigningSettings Org2OrgProvisioningOAuthSigningSettings

// NewOrg2OrgProvisioningOAuthSigningSettings instantiates a new Org2OrgProvisioningOAuthSigningSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrg2OrgProvisioningOAuthSigningSettings(rotationMode string) *Org2OrgProvisioningOAuthSigningSettings {
	this := Org2OrgProvisioningOAuthSigningSettings{}
	this.RotationMode = rotationMode
	return &this
}

// NewOrg2OrgProvisioningOAuthSigningSettingsWithDefaults instantiates a new Org2OrgProvisioningOAuthSigningSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrg2OrgProvisioningOAuthSigningSettingsWithDefaults() *Org2OrgProvisioningOAuthSigningSettings {
	this := Org2OrgProvisioningOAuthSigningSettings{}
	return &this
}

// GetRotationMode returns the RotationMode field value
func (o *Org2OrgProvisioningOAuthSigningSettings) GetRotationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RotationMode
}

// GetRotationModeOk returns a tuple with the RotationMode field value
// and a boolean to check if the value has been set.
func (o *Org2OrgProvisioningOAuthSigningSettings) GetRotationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationMode, true
}

// SetRotationMode sets field value
func (o *Org2OrgProvisioningOAuthSigningSettings) SetRotationMode(v string) {
	o.RotationMode = v
}

func (o Org2OrgProvisioningOAuthSigningSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rotationMode"] = o.RotationMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Org2OrgProvisioningOAuthSigningSettings) UnmarshalJSON(bytes []byte) (err error) {
	varOrg2OrgProvisioningOAuthSigningSettings := _Org2OrgProvisioningOAuthSigningSettings{}

	err = json.Unmarshal(bytes, &varOrg2OrgProvisioningOAuthSigningSettings)
	if err == nil {
		*o = Org2OrgProvisioningOAuthSigningSettings(varOrg2OrgProvisioningOAuthSigningSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "rotationMode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrg2OrgProvisioningOAuthSigningSettings struct {
	value *Org2OrgProvisioningOAuthSigningSettings
	isSet bool
}

func (v NullableOrg2OrgProvisioningOAuthSigningSettings) Get() *Org2OrgProvisioningOAuthSigningSettings {
	return v.value
}

func (v *NullableOrg2OrgProvisioningOAuthSigningSettings) Set(val *Org2OrgProvisioningOAuthSigningSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrg2OrgProvisioningOAuthSigningSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrg2OrgProvisioningOAuthSigningSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrg2OrgProvisioningOAuthSigningSettings(val *Org2OrgProvisioningOAuthSigningSettings) *NullableOrg2OrgProvisioningOAuthSigningSettings {
	return &NullableOrg2OrgProvisioningOAuthSigningSettings{value: val, isSet: true}
}

func (v NullableOrg2OrgProvisioningOAuthSigningSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrg2OrgProvisioningOAuthSigningSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


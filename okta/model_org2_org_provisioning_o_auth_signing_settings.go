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
	"fmt"
)

// checks if the Org2OrgProvisioningOAuthSigningSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Org2OrgProvisioningOAuthSigningSettings{}

// Org2OrgProvisioningOAuthSigningSettings Only used for the Okta Org2Org (`okta_org2org`) app.  The signing key rotation setting.
type Org2OrgProvisioningOAuthSigningSettings struct {
	// The signing key rotation setting for the provisioning connection
	RotationMode         string `json:"rotationMode"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Org2OrgProvisioningOAuthSigningSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rotationMode"] = o.RotationMode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Org2OrgProvisioningOAuthSigningSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rotationMode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrg2OrgProvisioningOAuthSigningSettings := _Org2OrgProvisioningOAuthSigningSettings{}

	err = json.Unmarshal(data, &varOrg2OrgProvisioningOAuthSigningSettings)

	if err != nil {
		return err
	}

	*o = Org2OrgProvisioningOAuthSigningSettings(varOrg2OrgProvisioningOAuthSigningSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rotationMode")
		o.AdditionalProperties = additionalProperties
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

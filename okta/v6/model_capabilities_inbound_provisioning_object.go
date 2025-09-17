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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the CapabilitiesInboundProvisioningObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesInboundProvisioningObject{}

// CapabilitiesInboundProvisioningObject Defines the configuration for the INBOUND_PROVISIONING feature
type CapabilitiesInboundProvisioningObject struct {
	ImportRules          CapabilitiesImportRulesObject    `json:"importRules"`
	ImportSettings       CapabilitiesImportSettingsObject `json:"importSettings"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesInboundProvisioningObject CapabilitiesInboundProvisioningObject

// NewCapabilitiesInboundProvisioningObject instantiates a new CapabilitiesInboundProvisioningObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesInboundProvisioningObject(importRules CapabilitiesImportRulesObject, importSettings CapabilitiesImportSettingsObject) *CapabilitiesInboundProvisioningObject {
	this := CapabilitiesInboundProvisioningObject{}
	this.ImportRules = importRules
	this.ImportSettings = importSettings
	return &this
}

// NewCapabilitiesInboundProvisioningObjectWithDefaults instantiates a new CapabilitiesInboundProvisioningObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesInboundProvisioningObjectWithDefaults() *CapabilitiesInboundProvisioningObject {
	this := CapabilitiesInboundProvisioningObject{}
	return &this
}

// GetImportRules returns the ImportRules field value
func (o *CapabilitiesInboundProvisioningObject) GetImportRules() CapabilitiesImportRulesObject {
	if o == nil {
		var ret CapabilitiesImportRulesObject
		return ret
	}

	return o.ImportRules
}

// GetImportRulesOk returns a tuple with the ImportRules field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesInboundProvisioningObject) GetImportRulesOk() (*CapabilitiesImportRulesObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportRules, true
}

// SetImportRules sets field value
func (o *CapabilitiesInboundProvisioningObject) SetImportRules(v CapabilitiesImportRulesObject) {
	o.ImportRules = v
}

// GetImportSettings returns the ImportSettings field value
func (o *CapabilitiesInboundProvisioningObject) GetImportSettings() CapabilitiesImportSettingsObject {
	if o == nil {
		var ret CapabilitiesImportSettingsObject
		return ret
	}

	return o.ImportSettings
}

// GetImportSettingsOk returns a tuple with the ImportSettings field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesInboundProvisioningObject) GetImportSettingsOk() (*CapabilitiesImportSettingsObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportSettings, true
}

// SetImportSettings sets field value
func (o *CapabilitiesInboundProvisioningObject) SetImportSettings(v CapabilitiesImportSettingsObject) {
	o.ImportSettings = v
}

func (o CapabilitiesInboundProvisioningObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesInboundProvisioningObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["importRules"] = o.ImportRules
	toSerialize["importSettings"] = o.ImportSettings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesInboundProvisioningObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"importRules",
		"importSettings",
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

	varCapabilitiesInboundProvisioningObject := _CapabilitiesInboundProvisioningObject{}

	err = json.Unmarshal(data, &varCapabilitiesInboundProvisioningObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesInboundProvisioningObject(varCapabilitiesInboundProvisioningObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "importRules")
		delete(additionalProperties, "importSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitiesInboundProvisioningObject struct {
	value *CapabilitiesInboundProvisioningObject
	isSet bool
}

func (v NullableCapabilitiesInboundProvisioningObject) Get() *CapabilitiesInboundProvisioningObject {
	return v.value
}

func (v *NullableCapabilitiesInboundProvisioningObject) Set(val *CapabilitiesInboundProvisioningObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesInboundProvisioningObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesInboundProvisioningObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesInboundProvisioningObject(val *CapabilitiesInboundProvisioningObject) *NullableCapabilitiesInboundProvisioningObject {
	return &NullableCapabilitiesInboundProvisioningObject{value: val, isSet: true}
}

func (v NullableCapabilitiesInboundProvisioningObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesInboundProvisioningObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

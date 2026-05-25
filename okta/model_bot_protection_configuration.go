/*
Okta Admin Management API

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

// checks if the BotProtectionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BotProtectionConfiguration{}

// BotProtectionConfiguration Bot protection configuration for the org
type BotProtectionConfiguration struct {
	// The type of enforcement to trigger when a bot is detected
	EnforcementType *string `json:"enforcementType,omitempty"`
	// The sensitivity level of bot detection
	Level string `json:"level"`
	// The enforcement mode for bot protection
	Mode string `json:"mode"`
	// An array of authentication flows that have bot protection enabled
	SupportedFlows       []string   `json:"supportedFlows,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BotProtectionConfiguration BotProtectionConfiguration

// NewBotProtectionConfiguration instantiates a new BotProtectionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotProtectionConfiguration(level string, mode string) *BotProtectionConfiguration {
	this := BotProtectionConfiguration{}
	this.Level = level
	this.Mode = mode
	return &this
}

// NewBotProtectionConfigurationWithDefaults instantiates a new BotProtectionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotProtectionConfigurationWithDefaults() *BotProtectionConfiguration {
	this := BotProtectionConfiguration{}
	return &this
}

// GetEnforcementType returns the EnforcementType field value if set, zero value otherwise.
func (o *BotProtectionConfiguration) GetEnforcementType() string {
	if o == nil || IsNil(o.EnforcementType) {
		var ret string
		return ret
	}
	return *o.EnforcementType
}

// GetEnforcementTypeOk returns a tuple with the EnforcementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProtectionConfiguration) GetEnforcementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnforcementType) {
		return nil, false
	}
	return o.EnforcementType, true
}

// HasEnforcementType returns a boolean if a field has been set.
func (o *BotProtectionConfiguration) HasEnforcementType() bool {
	if o != nil && !IsNil(o.EnforcementType) {
		return true
	}

	return false
}

// SetEnforcementType gets a reference to the given string and assigns it to the EnforcementType field.
func (o *BotProtectionConfiguration) SetEnforcementType(v string) {
	o.EnforcementType = &v
}

// GetLevel returns the Level field value
func (o *BotProtectionConfiguration) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *BotProtectionConfiguration) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *BotProtectionConfiguration) SetLevel(v string) {
	o.Level = v
}

// GetMode returns the Mode field value
func (o *BotProtectionConfiguration) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *BotProtectionConfiguration) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *BotProtectionConfiguration) SetMode(v string) {
	o.Mode = v
}

// GetSupportedFlows returns the SupportedFlows field value if set, zero value otherwise.
func (o *BotProtectionConfiguration) GetSupportedFlows() []string {
	if o == nil || IsNil(o.SupportedFlows) {
		var ret []string
		return ret
	}
	return o.SupportedFlows
}

// GetSupportedFlowsOk returns a tuple with the SupportedFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProtectionConfiguration) GetSupportedFlowsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedFlows) {
		return nil, false
	}
	return o.SupportedFlows, true
}

// HasSupportedFlows returns a boolean if a field has been set.
func (o *BotProtectionConfiguration) HasSupportedFlows() bool {
	if o != nil && !IsNil(o.SupportedFlows) {
		return true
	}

	return false
}

// SetSupportedFlows gets a reference to the given []string and assigns it to the SupportedFlows field.
func (o *BotProtectionConfiguration) SetSupportedFlows(v []string) {
	o.SupportedFlows = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BotProtectionConfiguration) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotProtectionConfiguration) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BotProtectionConfiguration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *BotProtectionConfiguration) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o BotProtectionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BotProtectionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnforcementType) {
		toSerialize["enforcementType"] = o.EnforcementType
	}
	toSerialize["level"] = o.Level
	toSerialize["mode"] = o.Mode
	if !IsNil(o.SupportedFlows) {
		toSerialize["supportedFlows"] = o.SupportedFlows
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BotProtectionConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
		"mode",
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

	varBotProtectionConfiguration := _BotProtectionConfiguration{}

	err = json.Unmarshal(data, &varBotProtectionConfiguration)

	if err != nil {
		return err
	}

	*o = BotProtectionConfiguration(varBotProtectionConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enforcementType")
		delete(additionalProperties, "level")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "supportedFlows")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBotProtectionConfiguration struct {
	value *BotProtectionConfiguration
	isSet bool
}

func (v NullableBotProtectionConfiguration) Get() *BotProtectionConfiguration {
	return v.value
}

func (v *NullableBotProtectionConfiguration) Set(val *BotProtectionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableBotProtectionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableBotProtectionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBotProtectionConfiguration(val *BotProtectionConfiguration) *NullableBotProtectionConfiguration {
	return &NullableBotProtectionConfiguration{value: val, isSet: true}
}

func (v NullableBotProtectionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBotProtectionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

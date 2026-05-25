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

// checks if the ProvisioningCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningCapability{}

// ProvisioningCapability Provisioning capability configuration with embedded protocol details.  *Protocol validation rule:** - The `actions` parameter is required when `supportedProtocols` is set to `ACTIONS`.
type ProvisioningCapability struct {
	// Configuration for the Actions protocol. This parameter is required when `supportedProtocols` is set to `ACTIONS`.
	Actions []SubmissionAction `json:"actions,omitempty"`
	// Provisioning capability identifier
	Capability string `json:"capability"`
	// List of supported provisioning protocols
	SupportedProtocols   []string `json:"supportedProtocols"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningCapability ProvisioningCapability

// NewProvisioningCapability instantiates a new ProvisioningCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningCapability(capability string, supportedProtocols []string) *ProvisioningCapability {
	this := ProvisioningCapability{}
	this.Capability = capability
	this.SupportedProtocols = supportedProtocols
	return &this
}

// NewProvisioningCapabilityWithDefaults instantiates a new ProvisioningCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningCapabilityWithDefaults() *ProvisioningCapability {
	this := ProvisioningCapability{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ProvisioningCapability) GetActions() []SubmissionAction {
	if o == nil || IsNil(o.Actions) {
		var ret []SubmissionAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningCapability) GetActionsOk() ([]SubmissionAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ProvisioningCapability) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []SubmissionAction and assigns it to the Actions field.
func (o *ProvisioningCapability) SetActions(v []SubmissionAction) {
	o.Actions = v
}

// GetCapability returns the Capability field value
func (o *ProvisioningCapability) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCapability) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *ProvisioningCapability) SetCapability(v string) {
	o.Capability = v
}

// GetSupportedProtocols returns the SupportedProtocols field value
func (o *ProvisioningCapability) GetSupportedProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedProtocols
}

// GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field value
// and a boolean to check if the value has been set.
func (o *ProvisioningCapability) GetSupportedProtocolsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedProtocols, true
}

// SetSupportedProtocols sets field value
func (o *ProvisioningCapability) SetSupportedProtocols(v []string) {
	o.SupportedProtocols = v
}

func (o ProvisioningCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	toSerialize["capability"] = o.Capability
	toSerialize["supportedProtocols"] = o.SupportedProtocols

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningCapability) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capability",
		"supportedProtocols",
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

	varProvisioningCapability := _ProvisioningCapability{}

	err = json.Unmarshal(data, &varProvisioningCapability)

	if err != nil {
		return err
	}

	*o = ProvisioningCapability(varProvisioningCapability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "capability")
		delete(additionalProperties, "supportedProtocols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningCapability struct {
	value *ProvisioningCapability
	isSet bool
}

func (v NullableProvisioningCapability) Get() *ProvisioningCapability {
	return v.value
}

func (v *NullableProvisioningCapability) Set(val *ProvisioningCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningCapability(val *ProvisioningCapability) *NullableProvisioningCapability {
	return &NullableProvisioningCapability{value: val, isSet: true}
}

func (v NullableProvisioningCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

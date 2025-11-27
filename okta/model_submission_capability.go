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

// checks if the SubmissionCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionCapability{}

// SubmissionCapability struct for SubmissionCapability
type SubmissionCapability struct {
	Capability           string   `json:"capability"`
	SupportedProtocols   []string `json:"supportedProtocols"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionCapability SubmissionCapability

// NewSubmissionCapability instantiates a new SubmissionCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionCapability(capability string, supportedProtocols []string) *SubmissionCapability {
	this := SubmissionCapability{}
	this.Capability = capability
	this.SupportedProtocols = supportedProtocols
	return &this
}

// NewSubmissionCapabilityWithDefaults instantiates a new SubmissionCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionCapabilityWithDefaults() *SubmissionCapability {
	this := SubmissionCapability{}
	return &this
}

// GetCapability returns the Capability field value
func (o *SubmissionCapability) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *SubmissionCapability) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *SubmissionCapability) SetCapability(v string) {
	o.Capability = v
}

// GetSupportedProtocols returns the SupportedProtocols field value
func (o *SubmissionCapability) GetSupportedProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedProtocols
}

// GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field value
// and a boolean to check if the value has been set.
func (o *SubmissionCapability) GetSupportedProtocolsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedProtocols, true
}

// SetSupportedProtocols sets field value
func (o *SubmissionCapability) SetSupportedProtocols(v []string) {
	o.SupportedProtocols = v
}

func (o SubmissionCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capability"] = o.Capability
	toSerialize["supportedProtocols"] = o.SupportedProtocols

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmissionCapability) UnmarshalJSON(data []byte) (err error) {
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

	varSubmissionCapability := _SubmissionCapability{}

	err = json.Unmarshal(data, &varSubmissionCapability)

	if err != nil {
		return err
	}

	*o = SubmissionCapability(varSubmissionCapability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capability")
		delete(additionalProperties, "supportedProtocols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmissionCapability struct {
	value *SubmissionCapability
	isSet bool
}

func (v NullableSubmissionCapability) Get() *SubmissionCapability {
	return v.value
}

func (v *NullableSubmissionCapability) Set(val *SubmissionCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionCapability(val *SubmissionCapability) *NullableSubmissionCapability {
	return &NullableSubmissionCapability{value: val, isSet: true}
}

func (v NullableSubmissionCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

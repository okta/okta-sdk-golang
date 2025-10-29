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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the SubmissionCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionCapabilities{}

// SubmissionCapabilities struct for SubmissionCapabilities
type SubmissionCapabilities struct {
	Capabilities         []SubmissionCapability `json:"capabilities"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionCapabilities SubmissionCapabilities

// NewSubmissionCapabilities instantiates a new SubmissionCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionCapabilities(capabilities []SubmissionCapability) *SubmissionCapabilities {
	this := SubmissionCapabilities{}
	this.Capabilities = capabilities
	return &this
}

// NewSubmissionCapabilitiesWithDefaults instantiates a new SubmissionCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionCapabilitiesWithDefaults() *SubmissionCapabilities {
	this := SubmissionCapabilities{}
	return &this
}

// GetCapabilities returns the Capabilities field value
func (o *SubmissionCapabilities) GetCapabilities() []SubmissionCapability {
	if o == nil {
		var ret []SubmissionCapability
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *SubmissionCapabilities) GetCapabilitiesOk() ([]SubmissionCapability, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *SubmissionCapabilities) SetCapabilities(v []SubmissionCapability) {
	o.Capabilities = v
}

func (o SubmissionCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capabilities"] = o.Capabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmissionCapabilities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capabilities",
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

	varSubmissionCapabilities := _SubmissionCapabilities{}

	err = json.Unmarshal(data, &varSubmissionCapabilities)

	if err != nil {
		return err
	}

	*o = SubmissionCapabilities(varSubmissionCapabilities)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmissionCapabilities struct {
	value *SubmissionCapabilities
	isSet bool
}

func (v NullableSubmissionCapabilities) Get() *SubmissionCapabilities {
	return v.value
}

func (v *NullableSubmissionCapabilities) Set(val *SubmissionCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionCapabilities(val *SubmissionCapabilities) *NullableSubmissionCapabilities {
	return &NullableSubmissionCapabilities{value: val, isSet: true}
}

func (v NullableSubmissionCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

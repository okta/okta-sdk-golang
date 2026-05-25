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

// checks if the ApiServiceCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiServiceCapability{}

// ApiServiceCapability API Service capability configuration with embedded protocol details
type ApiServiceCapability struct {
	// API Service capability identifier
	Capability string `json:"capability"`
	// List of supported protocols
	SupportedProtocols   []string `json:"supportedProtocols"`
	AdditionalProperties map[string]interface{}
}

type _ApiServiceCapability ApiServiceCapability

// NewApiServiceCapability instantiates a new ApiServiceCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiServiceCapability(capability string, supportedProtocols []string) *ApiServiceCapability {
	this := ApiServiceCapability{}
	this.Capability = capability
	this.SupportedProtocols = supportedProtocols
	return &this
}

// NewApiServiceCapabilityWithDefaults instantiates a new ApiServiceCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiServiceCapabilityWithDefaults() *ApiServiceCapability {
	this := ApiServiceCapability{}
	return &this
}

// GetCapability returns the Capability field value
func (o *ApiServiceCapability) GetCapability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *ApiServiceCapability) GetCapabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *ApiServiceCapability) SetCapability(v string) {
	o.Capability = v
}

// GetSupportedProtocols returns the SupportedProtocols field value
func (o *ApiServiceCapability) GetSupportedProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedProtocols
}

// GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field value
// and a boolean to check if the value has been set.
func (o *ApiServiceCapability) GetSupportedProtocolsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedProtocols, true
}

// SetSupportedProtocols sets field value
func (o *ApiServiceCapability) SetSupportedProtocols(v []string) {
	o.SupportedProtocols = v
}

func (o ApiServiceCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiServiceCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capability"] = o.Capability
	toSerialize["supportedProtocols"] = o.SupportedProtocols

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiServiceCapability) UnmarshalJSON(data []byte) (err error) {
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

	varApiServiceCapability := _ApiServiceCapability{}

	err = json.Unmarshal(data, &varApiServiceCapability)

	if err != nil {
		return err
	}

	*o = ApiServiceCapability(varApiServiceCapability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capability")
		delete(additionalProperties, "supportedProtocols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiServiceCapability struct {
	value *ApiServiceCapability
	isSet bool
}

func (v NullableApiServiceCapability) Get() *ApiServiceCapability {
	return v.value
}

func (v *NullableApiServiceCapability) Set(val *ApiServiceCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableApiServiceCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableApiServiceCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiServiceCapability(val *ApiServiceCapability) *NullableApiServiceCapability {
	return &NullableApiServiceCapability{value: val, isSet: true}
}

func (v NullableApiServiceCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiServiceCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

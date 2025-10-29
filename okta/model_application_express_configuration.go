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
)

// checks if the ApplicationExpressConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationExpressConfiguration{}

// ApplicationExpressConfiguration <div class=\"x-lifecycle-container\"><x-lifecycle class=\"oie\"></x-lifecycle></div> Indicates which Express Configuration capabilities the app supports and has enabled
type ApplicationExpressConfiguration struct {
	// Capabilities currently enabled for the app
	EnabledCapabilities []string `json:"enabledCapabilities,omitempty"`
	// Capabilities supported by the app
	SupportedCapabilities []string `json:"supportedCapabilities,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ApplicationExpressConfiguration ApplicationExpressConfiguration

// NewApplicationExpressConfiguration instantiates a new ApplicationExpressConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationExpressConfiguration() *ApplicationExpressConfiguration {
	this := ApplicationExpressConfiguration{}
	return &this
}

// NewApplicationExpressConfigurationWithDefaults instantiates a new ApplicationExpressConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationExpressConfigurationWithDefaults() *ApplicationExpressConfiguration {
	this := ApplicationExpressConfiguration{}
	return &this
}

// GetEnabledCapabilities returns the EnabledCapabilities field value if set, zero value otherwise.
func (o *ApplicationExpressConfiguration) GetEnabledCapabilities() []string {
	if o == nil || IsNil(o.EnabledCapabilities) {
		var ret []string
		return ret
	}
	return o.EnabledCapabilities
}

// GetEnabledCapabilitiesOk returns a tuple with the EnabledCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExpressConfiguration) GetEnabledCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledCapabilities) {
		return nil, false
	}
	return o.EnabledCapabilities, true
}

// HasEnabledCapabilities returns a boolean if a field has been set.
func (o *ApplicationExpressConfiguration) HasEnabledCapabilities() bool {
	if o != nil && !IsNil(o.EnabledCapabilities) {
		return true
	}

	return false
}

// SetEnabledCapabilities gets a reference to the given []string and assigns it to the EnabledCapabilities field.
func (o *ApplicationExpressConfiguration) SetEnabledCapabilities(v []string) {
	o.EnabledCapabilities = v
}

// GetSupportedCapabilities returns the SupportedCapabilities field value if set, zero value otherwise.
func (o *ApplicationExpressConfiguration) GetSupportedCapabilities() []string {
	if o == nil || IsNil(o.SupportedCapabilities) {
		var ret []string
		return ret
	}
	return o.SupportedCapabilities
}

// GetSupportedCapabilitiesOk returns a tuple with the SupportedCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationExpressConfiguration) GetSupportedCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedCapabilities) {
		return nil, false
	}
	return o.SupportedCapabilities, true
}

// HasSupportedCapabilities returns a boolean if a field has been set.
func (o *ApplicationExpressConfiguration) HasSupportedCapabilities() bool {
	if o != nil && !IsNil(o.SupportedCapabilities) {
		return true
	}

	return false
}

// SetSupportedCapabilities gets a reference to the given []string and assigns it to the SupportedCapabilities field.
func (o *ApplicationExpressConfiguration) SetSupportedCapabilities(v []string) {
	o.SupportedCapabilities = v
}

func (o ApplicationExpressConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationExpressConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnabledCapabilities) {
		toSerialize["enabledCapabilities"] = o.EnabledCapabilities
	}
	if !IsNil(o.SupportedCapabilities) {
		toSerialize["supportedCapabilities"] = o.SupportedCapabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationExpressConfiguration) UnmarshalJSON(data []byte) (err error) {
	varApplicationExpressConfiguration := _ApplicationExpressConfiguration{}

	err = json.Unmarshal(data, &varApplicationExpressConfiguration)

	if err != nil {
		return err
	}

	*o = ApplicationExpressConfiguration(varApplicationExpressConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabledCapabilities")
		delete(additionalProperties, "supportedCapabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationExpressConfiguration struct {
	value *ApplicationExpressConfiguration
	isSet bool
}

func (v NullableApplicationExpressConfiguration) Get() *ApplicationExpressConfiguration {
	return v.value
}

func (v *NullableApplicationExpressConfiguration) Set(val *ApplicationExpressConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationExpressConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationExpressConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationExpressConfiguration(val *ApplicationExpressConfiguration) *NullableApplicationExpressConfiguration {
	return &NullableApplicationExpressConfiguration{value: val, isSet: true}
}

func (v NullableApplicationExpressConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationExpressConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

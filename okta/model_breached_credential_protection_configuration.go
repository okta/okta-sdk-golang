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

// checks if the BreachedCredentialProtectionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BreachedCredentialProtectionConfiguration{}

// BreachedCredentialProtectionConfiguration Breached credential protection configuration for the org
type BreachedCredentialProtectionConfiguration struct {
	// The breach detection method used for breached credential protection
	DetectionMethod      string     `json:"detectionMethod"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BreachedCredentialProtectionConfiguration BreachedCredentialProtectionConfiguration

// NewBreachedCredentialProtectionConfiguration instantiates a new BreachedCredentialProtectionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreachedCredentialProtectionConfiguration(detectionMethod string) *BreachedCredentialProtectionConfiguration {
	this := BreachedCredentialProtectionConfiguration{}
	this.DetectionMethod = detectionMethod
	return &this
}

// NewBreachedCredentialProtectionConfigurationWithDefaults instantiates a new BreachedCredentialProtectionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreachedCredentialProtectionConfigurationWithDefaults() *BreachedCredentialProtectionConfiguration {
	this := BreachedCredentialProtectionConfiguration{}
	return &this
}

// GetDetectionMethod returns the DetectionMethod field value
func (o *BreachedCredentialProtectionConfiguration) GetDetectionMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value
// and a boolean to check if the value has been set.
func (o *BreachedCredentialProtectionConfiguration) GetDetectionMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionMethod, true
}

// SetDetectionMethod sets field value
func (o *BreachedCredentialProtectionConfiguration) SetDetectionMethod(v string) {
	o.DetectionMethod = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BreachedCredentialProtectionConfiguration) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BreachedCredentialProtectionConfiguration) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BreachedCredentialProtectionConfiguration) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *BreachedCredentialProtectionConfiguration) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o BreachedCredentialProtectionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BreachedCredentialProtectionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detectionMethod"] = o.DetectionMethod
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BreachedCredentialProtectionConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detectionMethod",
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

	varBreachedCredentialProtectionConfiguration := _BreachedCredentialProtectionConfiguration{}

	err = json.Unmarshal(data, &varBreachedCredentialProtectionConfiguration)

	if err != nil {
		return err
	}

	*o = BreachedCredentialProtectionConfiguration(varBreachedCredentialProtectionConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detectionMethod")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBreachedCredentialProtectionConfiguration struct {
	value *BreachedCredentialProtectionConfiguration
	isSet bool
}

func (v NullableBreachedCredentialProtectionConfiguration) Get() *BreachedCredentialProtectionConfiguration {
	return v.value
}

func (v *NullableBreachedCredentialProtectionConfiguration) Set(val *BreachedCredentialProtectionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableBreachedCredentialProtectionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableBreachedCredentialProtectionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreachedCredentialProtectionConfiguration(val *BreachedCredentialProtectionConfiguration) *NullableBreachedCredentialProtectionConfiguration {
	return &NullableBreachedCredentialProtectionConfiguration{value: val, isSet: true}
}

func (v NullableBreachedCredentialProtectionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreachedCredentialProtectionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

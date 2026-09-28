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

// checks if the BreachedCredentialProtectionConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BreachedCredentialProtectionConfigurationRequest{}

// BreachedCredentialProtectionConfigurationRequest Request body for updating the breached credential protection configuration
type BreachedCredentialProtectionConfigurationRequest struct {
	// The breach detection method used for breached credential protection
	DetectionMethod      string `json:"detectionMethod"`
	AdditionalProperties map[string]interface{}
}

type _BreachedCredentialProtectionConfigurationRequest BreachedCredentialProtectionConfigurationRequest

// NewBreachedCredentialProtectionConfigurationRequest instantiates a new BreachedCredentialProtectionConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreachedCredentialProtectionConfigurationRequest(detectionMethod string) *BreachedCredentialProtectionConfigurationRequest {
	this := BreachedCredentialProtectionConfigurationRequest{}
	this.DetectionMethod = detectionMethod
	return &this
}

// NewBreachedCredentialProtectionConfigurationRequestWithDefaults instantiates a new BreachedCredentialProtectionConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreachedCredentialProtectionConfigurationRequestWithDefaults() *BreachedCredentialProtectionConfigurationRequest {
	this := BreachedCredentialProtectionConfigurationRequest{}
	return &this
}

// GetDetectionMethod returns the DetectionMethod field value
func (o *BreachedCredentialProtectionConfigurationRequest) GetDetectionMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value
// and a boolean to check if the value has been set.
func (o *BreachedCredentialProtectionConfigurationRequest) GetDetectionMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionMethod, true
}

// SetDetectionMethod sets field value
func (o *BreachedCredentialProtectionConfigurationRequest) SetDetectionMethod(v string) {
	o.DetectionMethod = v
}

func (o BreachedCredentialProtectionConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BreachedCredentialProtectionConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detectionMethod"] = o.DetectionMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BreachedCredentialProtectionConfigurationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBreachedCredentialProtectionConfigurationRequest := _BreachedCredentialProtectionConfigurationRequest{}

	err = json.Unmarshal(data, &varBreachedCredentialProtectionConfigurationRequest)

	if err != nil {
		return err
	}

	*o = BreachedCredentialProtectionConfigurationRequest(varBreachedCredentialProtectionConfigurationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detectionMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBreachedCredentialProtectionConfigurationRequest struct {
	value *BreachedCredentialProtectionConfigurationRequest
	isSet bool
}

func (v NullableBreachedCredentialProtectionConfigurationRequest) Get() *BreachedCredentialProtectionConfigurationRequest {
	return v.value
}

func (v *NullableBreachedCredentialProtectionConfigurationRequest) Set(val *BreachedCredentialProtectionConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBreachedCredentialProtectionConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBreachedCredentialProtectionConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreachedCredentialProtectionConfigurationRequest(val *BreachedCredentialProtectionConfigurationRequest) *NullableBreachedCredentialProtectionConfigurationRequest {
	return &NullableBreachedCredentialProtectionConfigurationRequest{value: val, isSet: true}
}

func (v NullableBreachedCredentialProtectionConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreachedCredentialProtectionConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

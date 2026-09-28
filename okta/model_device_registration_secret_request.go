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

// checks if the DeviceRegistrationSecretRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRegistrationSecretRequest{}

// DeviceRegistrationSecretRequest struct for DeviceRegistrationSecretRequest
type DeviceRegistrationSecretRequest struct {
	// Admin-friendly label for this registration secret. Must be unique per org.
	Description string `json:"description"`
	// Maximum number of successful device registrations this secret may be used for before further registrations are rejected. The cap bounds the blast radius of a leaked secret. Admins must size it to the fleet they intend to register with this secret.
	MaxRegistrations     *int32 `json:"maxRegistrations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRegistrationSecretRequest DeviceRegistrationSecretRequest

// NewDeviceRegistrationSecretRequest instantiates a new DeviceRegistrationSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRegistrationSecretRequest(description string) *DeviceRegistrationSecretRequest {
	this := DeviceRegistrationSecretRequest{}
	this.Description = description
	var maxRegistrations int32 = 10000
	this.MaxRegistrations = &maxRegistrations
	return &this
}

// NewDeviceRegistrationSecretRequestWithDefaults instantiates a new DeviceRegistrationSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRegistrationSecretRequestWithDefaults() *DeviceRegistrationSecretRequest {
	this := DeviceRegistrationSecretRequest{}
	var maxRegistrations int32 = 10000
	this.MaxRegistrations = &maxRegistrations
	return &this
}

// GetDescription returns the Description field value
func (o *DeviceRegistrationSecretRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DeviceRegistrationSecretRequest) SetDescription(v string) {
	o.Description = v
}

// GetMaxRegistrations returns the MaxRegistrations field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretRequest) GetMaxRegistrations() int32 {
	if o == nil || IsNil(o.MaxRegistrations) {
		var ret int32
		return ret
	}
	return *o.MaxRegistrations
}

// GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretRequest) GetMaxRegistrationsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRegistrations) {
		return nil, false
	}
	return o.MaxRegistrations, true
}

// HasMaxRegistrations returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretRequest) HasMaxRegistrations() bool {
	if o != nil && !IsNil(o.MaxRegistrations) {
		return true
	}

	return false
}

// SetMaxRegistrations gets a reference to the given int32 and assigns it to the MaxRegistrations field.
func (o *DeviceRegistrationSecretRequest) SetMaxRegistrations(v int32) {
	o.MaxRegistrations = &v
}

func (o DeviceRegistrationSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRegistrationSecretRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.MaxRegistrations) {
		toSerialize["maxRegistrations"] = o.MaxRegistrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRegistrationSecretRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varDeviceRegistrationSecretRequest := _DeviceRegistrationSecretRequest{}

	err = json.Unmarshal(data, &varDeviceRegistrationSecretRequest)

	if err != nil {
		return err
	}

	*o = DeviceRegistrationSecretRequest(varDeviceRegistrationSecretRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "maxRegistrations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRegistrationSecretRequest struct {
	value *DeviceRegistrationSecretRequest
	isSet bool
}

func (v NullableDeviceRegistrationSecretRequest) Get() *DeviceRegistrationSecretRequest {
	return v.value
}

func (v *NullableDeviceRegistrationSecretRequest) Set(val *DeviceRegistrationSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRegistrationSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRegistrationSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRegistrationSecretRequest(val *DeviceRegistrationSecretRequest) *NullableDeviceRegistrationSecretRequest {
	return &NullableDeviceRegistrationSecretRequest{value: val, isSet: true}
}

func (v NullableDeviceRegistrationSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRegistrationSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

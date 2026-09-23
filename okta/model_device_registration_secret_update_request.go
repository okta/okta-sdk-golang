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
)

// checks if the DeviceRegistrationSecretUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRegistrationSecretUpdateRequest{}

// DeviceRegistrationSecretUpdateRequest struct for DeviceRegistrationSecretUpdateRequest
type DeviceRegistrationSecretUpdateRequest struct {
	// Admin-friendly label for this registration secret. Must be unique per org.
	Description *string `json:"description,omitempty"`
	// Maximum number of successful device registrations. May be raised or lowered. Lowering below the current `registrationCount` is allowed and effectively soft-disables further registrations against this secret.
	MaxRegistrations *int32 `json:"maxRegistrations,omitempty"`
	// Status of the registration secret. * `ACTIVE` — The secret can be used for device registration. * `INACTIVE` — The secret can't be used for device registration. Any token request presenting a client assertion signed with an inactive secret is rejected.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRegistrationSecretUpdateRequest DeviceRegistrationSecretUpdateRequest

// NewDeviceRegistrationSecretUpdateRequest instantiates a new DeviceRegistrationSecretUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRegistrationSecretUpdateRequest() *DeviceRegistrationSecretUpdateRequest {
	this := DeviceRegistrationSecretUpdateRequest{}
	return &this
}

// NewDeviceRegistrationSecretUpdateRequestWithDefaults instantiates a new DeviceRegistrationSecretUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRegistrationSecretUpdateRequestWithDefaults() *DeviceRegistrationSecretUpdateRequest {
	this := DeviceRegistrationSecretUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRegistrationSecretUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMaxRegistrations returns the MaxRegistrations field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretUpdateRequest) GetMaxRegistrations() int32 {
	if o == nil || IsNil(o.MaxRegistrations) {
		var ret int32
		return ret
	}
	return *o.MaxRegistrations
}

// GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretUpdateRequest) GetMaxRegistrationsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRegistrations) {
		return nil, false
	}
	return o.MaxRegistrations, true
}

// HasMaxRegistrations returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretUpdateRequest) HasMaxRegistrations() bool {
	if o != nil && !IsNil(o.MaxRegistrations) {
		return true
	}

	return false
}

// SetMaxRegistrations gets a reference to the given int32 and assigns it to the MaxRegistrations field.
func (o *DeviceRegistrationSecretUpdateRequest) SetMaxRegistrations(v int32) {
	o.MaxRegistrations = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretUpdateRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretUpdateRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceRegistrationSecretUpdateRequest) SetStatus(v string) {
	o.Status = &v
}

func (o DeviceRegistrationSecretUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRegistrationSecretUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MaxRegistrations) {
		toSerialize["maxRegistrations"] = o.MaxRegistrations
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRegistrationSecretUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varDeviceRegistrationSecretUpdateRequest := _DeviceRegistrationSecretUpdateRequest{}

	err = json.Unmarshal(data, &varDeviceRegistrationSecretUpdateRequest)

	if err != nil {
		return err
	}

	*o = DeviceRegistrationSecretUpdateRequest(varDeviceRegistrationSecretUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "maxRegistrations")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRegistrationSecretUpdateRequest struct {
	value *DeviceRegistrationSecretUpdateRequest
	isSet bool
}

func (v NullableDeviceRegistrationSecretUpdateRequest) Get() *DeviceRegistrationSecretUpdateRequest {
	return v.value
}

func (v *NullableDeviceRegistrationSecretUpdateRequest) Set(val *DeviceRegistrationSecretUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRegistrationSecretUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRegistrationSecretUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRegistrationSecretUpdateRequest(val *DeviceRegistrationSecretUpdateRequest) *NullableDeviceRegistrationSecretUpdateRequest {
	return &NullableDeviceRegistrationSecretUpdateRequest{value: val, isSet: true}
}

func (v NullableDeviceRegistrationSecretUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRegistrationSecretUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

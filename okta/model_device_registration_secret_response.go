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
	"time"
)

// checks if the DeviceRegistrationSecretResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceRegistrationSecretResponse{}

// DeviceRegistrationSecretResponse struct for DeviceRegistrationSecretResponse
type DeviceRegistrationSecretResponse struct {
	// Timestamp when the registration secret was created
	Created *time.Time `json:"created,omitempty"`
	// Admin-friendly label for this registration secret
	Description *string `json:"description,omitempty"`
	// Unique identifier for the registration secret (prefix `drs`)
	Id *string `json:"id,omitempty"`
	// Timestamp when the registration secret was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Maximum number of successful device registrations allowed for this secret
	MaxRegistrations *int32 `json:"maxRegistrations,omitempty"`
	// Number of successful HMAC token issuances against this secret. Monotonically increasing for the lifetime of the secret.  > **Note:** A device that obtains a token but doesn't complete `/idp/v2/device/register` leaves its grant `PENDING` while still consuming a slot of the cap.
	RegistrationCount *int32 `json:"registrationCount,omitempty"`
	// Status of the registration secret. * `ACTIVE` — The secret can be used for device registration. * `INACTIVE` — The secret can't be used for device registration. Any token request presenting a client assertion signed with an inactive secret is rejected.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceRegistrationSecretResponse DeviceRegistrationSecretResponse

// NewDeviceRegistrationSecretResponse instantiates a new DeviceRegistrationSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRegistrationSecretResponse() *DeviceRegistrationSecretResponse {
	this := DeviceRegistrationSecretResponse{}
	return &this
}

// NewDeviceRegistrationSecretResponseWithDefaults instantiates a new DeviceRegistrationSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRegistrationSecretResponseWithDefaults() *DeviceRegistrationSecretResponse {
	this := DeviceRegistrationSecretResponse{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DeviceRegistrationSecretResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRegistrationSecretResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceRegistrationSecretResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DeviceRegistrationSecretResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetMaxRegistrations returns the MaxRegistrations field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetMaxRegistrations() int32 {
	if o == nil || IsNil(o.MaxRegistrations) {
		var ret int32
		return ret
	}
	return *o.MaxRegistrations
}

// GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetMaxRegistrationsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRegistrations) {
		return nil, false
	}
	return o.MaxRegistrations, true
}

// HasMaxRegistrations returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasMaxRegistrations() bool {
	if o != nil && !IsNil(o.MaxRegistrations) {
		return true
	}

	return false
}

// SetMaxRegistrations gets a reference to the given int32 and assigns it to the MaxRegistrations field.
func (o *DeviceRegistrationSecretResponse) SetMaxRegistrations(v int32) {
	o.MaxRegistrations = &v
}

// GetRegistrationCount returns the RegistrationCount field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetRegistrationCount() int32 {
	if o == nil || IsNil(o.RegistrationCount) {
		var ret int32
		return ret
	}
	return *o.RegistrationCount
}

// GetRegistrationCountOk returns a tuple with the RegistrationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetRegistrationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RegistrationCount) {
		return nil, false
	}
	return o.RegistrationCount, true
}

// HasRegistrationCount returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasRegistrationCount() bool {
	if o != nil && !IsNil(o.RegistrationCount) {
		return true
	}

	return false
}

// SetRegistrationCount gets a reference to the given int32 and assigns it to the RegistrationCount field.
func (o *DeviceRegistrationSecretResponse) SetRegistrationCount(v int32) {
	o.RegistrationCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceRegistrationSecretResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRegistrationSecretResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceRegistrationSecretResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceRegistrationSecretResponse) SetStatus(v string) {
	o.Status = &v
}

func (o DeviceRegistrationSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceRegistrationSecretResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.MaxRegistrations) {
		toSerialize["maxRegistrations"] = o.MaxRegistrations
	}
	if !IsNil(o.RegistrationCount) {
		toSerialize["registrationCount"] = o.RegistrationCount
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceRegistrationSecretResponse) UnmarshalJSON(data []byte) (err error) {
	varDeviceRegistrationSecretResponse := _DeviceRegistrationSecretResponse{}

	err = json.Unmarshal(data, &varDeviceRegistrationSecretResponse)

	if err != nil {
		return err
	}

	*o = DeviceRegistrationSecretResponse(varDeviceRegistrationSecretResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "maxRegistrations")
		delete(additionalProperties, "registrationCount")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceRegistrationSecretResponse struct {
	value *DeviceRegistrationSecretResponse
	isSet bool
}

func (v NullableDeviceRegistrationSecretResponse) Get() *DeviceRegistrationSecretResponse {
	return v.value
}

func (v *NullableDeviceRegistrationSecretResponse) Set(val *DeviceRegistrationSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRegistrationSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRegistrationSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRegistrationSecretResponse(val *DeviceRegistrationSecretResponse) *NullableDeviceRegistrationSecretResponse {
	return &NullableDeviceRegistrationSecretResponse{value: val, isSet: true}
}

func (v NullableDeviceRegistrationSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRegistrationSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

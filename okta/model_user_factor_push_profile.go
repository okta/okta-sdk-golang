/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserFactorPushProfile struct for UserFactorPushProfile
type UserFactorPushProfile struct {
	// ID for the Factor credential
	CredentialId *string `json:"credentialId,omitempty"`
	// Token used to identify the device
	DeviceToken *string `json:"deviceToken,omitempty"`
	// Type of device
	DeviceType *string `json:"deviceType,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// OS version of the associated device
	Platform *string `json:"platform,omitempty"`
	// Installed version of Okta Verify
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushProfile UserFactorPushProfile

// NewUserFactorPushProfile instantiates a new UserFactorPushProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushProfile() *UserFactorPushProfile {
	this := UserFactorPushProfile{}
	return &this
}

// NewUserFactorPushProfileWithDefaults instantiates a new UserFactorPushProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushProfileWithDefaults() *UserFactorPushProfile {
	this := UserFactorPushProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *UserFactorPushProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

// GetDeviceToken returns the DeviceToken field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetDeviceToken() string {
	if o == nil || o.DeviceToken == nil {
		var ret string
		return ret
	}
	return *o.DeviceToken
}

// GetDeviceTokenOk returns a tuple with the DeviceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetDeviceTokenOk() (*string, bool) {
	if o == nil || o.DeviceToken == nil {
		return nil, false
	}
	return o.DeviceToken, true
}

// HasDeviceToken returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasDeviceToken() bool {
	if o != nil && o.DeviceToken != nil {
		return true
	}

	return false
}

// SetDeviceToken gets a reference to the given string and assigns it to the DeviceToken field.
func (o *UserFactorPushProfile) SetDeviceToken(v string) {
	o.DeviceToken = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *UserFactorPushProfile) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserFactorPushProfile) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *UserFactorPushProfile) SetPlatform(v string) {
	o.Platform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UserFactorPushProfile) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushProfile) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UserFactorPushProfile) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UserFactorPushProfile) SetVersion(v string) {
	o.Version = &v
}

func (o UserFactorPushProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}
	if o.DeviceToken != nil {
		toSerialize["deviceToken"] = o.DeviceToken
	}
	if o.DeviceType != nil {
		toSerialize["deviceType"] = o.DeviceType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorPushProfile := _UserFactorPushProfile{}

	err = json.Unmarshal(bytes, &varUserFactorPushProfile)
	if err == nil {
		*o = UserFactorPushProfile(varUserFactorPushProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentialId")
		delete(additionalProperties, "deviceToken")
		delete(additionalProperties, "deviceType")
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorPushProfile struct {
	value *UserFactorPushProfile
	isSet bool
}

func (v NullableUserFactorPushProfile) Get() *UserFactorPushProfile {
	return v.value
}

func (v *NullableUserFactorPushProfile) Set(val *UserFactorPushProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushProfile(val *UserFactorPushProfile) *NullableUserFactorPushProfile {
	return &NullableUserFactorPushProfile{value: val, isSet: true}
}

func (v NullableUserFactorPushProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the UserFactorSignedNonceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorSignedNonceProfile{}

// UserFactorSignedNonceProfile Profile for the Okta FastPass (signed nonce) factor
type UserFactorSignedNonceProfile struct {
	// ID for the factor credential
	CredentialId *string `json:"credentialId,omitempty"`
	// Type of device
	DeviceType *string `json:"deviceType,omitempty"`
	// Cryptographic keys associated with the signed nonce factor
	Keys []UserFactorSignedNonceProfileKey `json:"keys,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// OS platform of the associated device
	Platform *string `json:"platform,omitempty"`
	// OS version of the associated device
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSignedNonceProfile UserFactorSignedNonceProfile

// NewUserFactorSignedNonceProfile instantiates a new UserFactorSignedNonceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSignedNonceProfile() *UserFactorSignedNonceProfile {
	this := UserFactorSignedNonceProfile{}
	return &this
}

// NewUserFactorSignedNonceProfileWithDefaults instantiates a new UserFactorSignedNonceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSignedNonceProfileWithDefaults() *UserFactorSignedNonceProfile {
	this := UserFactorSignedNonceProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetCredentialId() string {
	if o == nil || IsNil(o.CredentialId) {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialId) {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasCredentialId() bool {
	if o != nil && !IsNil(o.CredentialId) {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *UserFactorSignedNonceProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *UserFactorSignedNonceProfile) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetKeys() []UserFactorSignedNonceProfileKey {
	if o == nil || IsNil(o.Keys) {
		var ret []UserFactorSignedNonceProfileKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetKeysOk() ([]UserFactorSignedNonceProfileKey, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []UserFactorSignedNonceProfileKey and assigns it to the Keys field.
func (o *UserFactorSignedNonceProfile) SetKeys(v []UserFactorSignedNonceProfileKey) {
	o.Keys = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserFactorSignedNonceProfile) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *UserFactorSignedNonceProfile) SetPlatform(v string) {
	o.Platform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfile) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfile) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfile) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UserFactorSignedNonceProfile) SetVersion(v string) {
	o.Version = &v
}

func (o UserFactorSignedNonceProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorSignedNonceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialId) {
		toSerialize["credentialId"] = o.CredentialId
	}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorSignedNonceProfile) UnmarshalJSON(data []byte) (err error) {
	varUserFactorSignedNonceProfile := _UserFactorSignedNonceProfile{}

	err = json.Unmarshal(data, &varUserFactorSignedNonceProfile)

	if err != nil {
		return err
	}

	*o = UserFactorSignedNonceProfile(varUserFactorSignedNonceProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentialId")
		delete(additionalProperties, "deviceType")
		delete(additionalProperties, "keys")
		delete(additionalProperties, "name")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorSignedNonceProfile struct {
	value *UserFactorSignedNonceProfile
	isSet bool
}

func (v NullableUserFactorSignedNonceProfile) Get() *UserFactorSignedNonceProfile {
	return v.value
}

func (v *NullableUserFactorSignedNonceProfile) Set(val *UserFactorSignedNonceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSignedNonceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSignedNonceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSignedNonceProfile(val *UserFactorSignedNonceProfile) *NullableUserFactorSignedNonceProfile {
	return &NullableUserFactorSignedNonceProfile{value: val, isSet: true}
}

func (v NullableUserFactorSignedNonceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSignedNonceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

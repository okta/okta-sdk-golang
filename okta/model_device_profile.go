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

// DeviceProfile struct for DeviceProfile
type DeviceProfile struct {
	// Type of encryption used on the device > **Note:** The following values map to Disk Encryption ON: `FULL`, `USER`, `ALL_INTERNAL_VOLUMES`. All other values map to Disk Encryption OFF.
	DiskEncryptionType *string `json:"diskEncryptionType,omitempty"`
	// Display name of the device
	DisplayName string `json:"displayName"`
	// International Mobile Equipment Identity (IMEI) of the device
	Imei *string `json:"imei,omitempty"`
	// Indicates if the device is jailbroken or rooted. Only applicable to `IOS` and `ANDROID` platforms
	IntegrityJailbreak *bool `json:"integrityJailbreak,omitempty"`
	// Name of the manufacturer of the device
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Mobile equipment identifier of the device
	Meid *string `json:"meid,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// Version of the device OS
	OsVersion *string `json:"osVersion,omitempty"`
	// OS platform of the device
	Platform string `json:"platform"`
	// Indicates if the device is registered at Okta
	Registered bool `json:"registered"`
	// Indicates if the device contains a secure hardware functionality
	SecureHardwarePresent *bool `json:"secureHardwarePresent,omitempty"`
	// Serial number of the device
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Windows Security identifier of the device
	Sid *string `json:"sid,omitempty"`
	// Windows Trusted Platform Module hash value
	TpmPublicKeyHash *string `json:"tpmPublicKeyHash,omitempty"`
	// macOS Unique Device identifier of the device
	Udid *string `json:"udid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceProfile DeviceProfile

// NewDeviceProfile instantiates a new DeviceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceProfile(displayName string, platform string, registered bool) *DeviceProfile {
	this := DeviceProfile{}
	this.DisplayName = displayName
	this.Platform = platform
	this.Registered = registered
	return &this
}

// NewDeviceProfileWithDefaults instantiates a new DeviceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceProfileWithDefaults() *DeviceProfile {
	this := DeviceProfile{}
	return &this
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *DeviceProfile) GetDiskEncryptionType() string {
	if o == nil || o.DiskEncryptionType == nil {
		var ret string
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetDiskEncryptionTypeOk() (*string, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *DeviceProfile) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given string and assigns it to the DiskEncryptionType field.
func (o *DeviceProfile) SetDiskEncryptionType(v string) {
	o.DiskEncryptionType = &v
}

// GetDisplayName returns the DisplayName field value
func (o *DeviceProfile) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *DeviceProfile) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *DeviceProfile) GetImei() string {
	if o == nil || o.Imei == nil {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetImeiOk() (*string, bool) {
	if o == nil || o.Imei == nil {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *DeviceProfile) HasImei() bool {
	if o != nil && o.Imei != nil {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *DeviceProfile) SetImei(v string) {
	o.Imei = &v
}

// GetIntegrityJailbreak returns the IntegrityJailbreak field value if set, zero value otherwise.
func (o *DeviceProfile) GetIntegrityJailbreak() bool {
	if o == nil || o.IntegrityJailbreak == nil {
		var ret bool
		return ret
	}
	return *o.IntegrityJailbreak
}

// GetIntegrityJailbreakOk returns a tuple with the IntegrityJailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetIntegrityJailbreakOk() (*bool, bool) {
	if o == nil || o.IntegrityJailbreak == nil {
		return nil, false
	}
	return o.IntegrityJailbreak, true
}

// HasIntegrityJailbreak returns a boolean if a field has been set.
func (o *DeviceProfile) HasIntegrityJailbreak() bool {
	if o != nil && o.IntegrityJailbreak != nil {
		return true
	}

	return false
}

// SetIntegrityJailbreak gets a reference to the given bool and assigns it to the IntegrityJailbreak field.
func (o *DeviceProfile) SetIntegrityJailbreak(v bool) {
	o.IntegrityJailbreak = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *DeviceProfile) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *DeviceProfile) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *DeviceProfile) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetMeid returns the Meid field value if set, zero value otherwise.
func (o *DeviceProfile) GetMeid() string {
	if o == nil || o.Meid == nil {
		var ret string
		return ret
	}
	return *o.Meid
}

// GetMeidOk returns a tuple with the Meid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetMeidOk() (*string, bool) {
	if o == nil || o.Meid == nil {
		return nil, false
	}
	return o.Meid, true
}

// HasMeid returns a boolean if a field has been set.
func (o *DeviceProfile) HasMeid() bool {
	if o != nil && o.Meid != nil {
		return true
	}

	return false
}

// SetMeid gets a reference to the given string and assigns it to the Meid field.
func (o *DeviceProfile) SetMeid(v string) {
	o.Meid = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DeviceProfile) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DeviceProfile) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DeviceProfile) SetModel(v string) {
	o.Model = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceProfile) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceProfile) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *DeviceProfile) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPlatform returns the Platform field value
func (o *DeviceProfile) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *DeviceProfile) SetPlatform(v string) {
	o.Platform = v
}

// GetRegistered returns the Registered field value
func (o *DeviceProfile) GetRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registered, true
}

// SetRegistered sets field value
func (o *DeviceProfile) SetRegistered(v bool) {
	o.Registered = v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *DeviceProfile) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *DeviceProfile) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *DeviceProfile) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceProfile) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceProfile) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceProfile) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *DeviceProfile) GetSid() string {
	if o == nil || o.Sid == nil {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetSidOk() (*string, bool) {
	if o == nil || o.Sid == nil {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *DeviceProfile) HasSid() bool {
	if o != nil && o.Sid != nil {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *DeviceProfile) SetSid(v string) {
	o.Sid = &v
}

// GetTpmPublicKeyHash returns the TpmPublicKeyHash field value if set, zero value otherwise.
func (o *DeviceProfile) GetTpmPublicKeyHash() string {
	if o == nil || o.TpmPublicKeyHash == nil {
		var ret string
		return ret
	}
	return *o.TpmPublicKeyHash
}

// GetTpmPublicKeyHashOk returns a tuple with the TpmPublicKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetTpmPublicKeyHashOk() (*string, bool) {
	if o == nil || o.TpmPublicKeyHash == nil {
		return nil, false
	}
	return o.TpmPublicKeyHash, true
}

// HasTpmPublicKeyHash returns a boolean if a field has been set.
func (o *DeviceProfile) HasTpmPublicKeyHash() bool {
	if o != nil && o.TpmPublicKeyHash != nil {
		return true
	}

	return false
}

// SetTpmPublicKeyHash gets a reference to the given string and assigns it to the TpmPublicKeyHash field.
func (o *DeviceProfile) SetTpmPublicKeyHash(v string) {
	o.TpmPublicKeyHash = &v
}

// GetUdid returns the Udid field value if set, zero value otherwise.
func (o *DeviceProfile) GetUdid() string {
	if o == nil || o.Udid == nil {
		var ret string
		return ret
	}
	return *o.Udid
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetUdidOk() (*string, bool) {
	if o == nil || o.Udid == nil {
		return nil, false
	}
	return o.Udid, true
}

// HasUdid returns a boolean if a field has been set.
func (o *DeviceProfile) HasUdid() bool {
	if o != nil && o.Udid != nil {
		return true
	}

	return false
}

// SetUdid gets a reference to the given string and assigns it to the Udid field.
func (o *DeviceProfile) SetUdid(v string) {
	o.Udid = &v
}

func (o DeviceProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskEncryptionType != nil {
		toSerialize["diskEncryptionType"] = o.DiskEncryptionType
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Imei != nil {
		toSerialize["imei"] = o.Imei
	}
	if o.IntegrityJailbreak != nil {
		toSerialize["integrityJailbreak"] = o.IntegrityJailbreak
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.Meid != nil {
		toSerialize["meid"] = o.Meid
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["registered"] = o.Registered
	}
	if o.SecureHardwarePresent != nil {
		toSerialize["secureHardwarePresent"] = o.SecureHardwarePresent
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.Sid != nil {
		toSerialize["sid"] = o.Sid
	}
	if o.TpmPublicKeyHash != nil {
		toSerialize["tpmPublicKeyHash"] = o.TpmPublicKeyHash
	}
	if o.Udid != nil {
		toSerialize["udid"] = o.Udid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceProfile) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceProfile := _DeviceProfile{}

	err = json.Unmarshal(bytes, &varDeviceProfile)
	if err == nil {
		*o = DeviceProfile(varDeviceProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "diskEncryptionType")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "imei")
		delete(additionalProperties, "integrityJailbreak")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "meid")
		delete(additionalProperties, "model")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "registered")
		delete(additionalProperties, "secureHardwarePresent")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "sid")
		delete(additionalProperties, "tpmPublicKeyHash")
		delete(additionalProperties, "udid")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceProfile struct {
	value *DeviceProfile
	isSet bool
}

func (v NullableDeviceProfile) Get() *DeviceProfile {
	return v.value
}

func (v *NullableDeviceProfile) Set(val *DeviceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceProfile(val *DeviceProfile) *NullableDeviceProfile {
	return &NullableDeviceProfile{value: val, isSet: true}
}

func (v NullableDeviceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


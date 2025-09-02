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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// LogDevice The entity that describes a device enrolled with passwordless authentication using Okta Verify.
type LogDevice struct {
	// The integration platform or software used with the device
	DeviceIntegrator map[string]interface{} `json:"device_integrator,omitempty"`
	DiskEncryptionType *string `json:"disk_encryption_type,omitempty"`
	// ID of the device
	Id *string `json:"id,omitempty"`
	// If the device has removed software restrictions
	Jailbreak *bool `json:"jailbreak,omitempty"`
	// Indicates if the device is configured for device management and is registered with Okta
	Managed *bool `json:"managed,omitempty"`
	Name *string `json:"name,omitempty"`
	OsPlatform *string `json:"os_platform,omitempty"`
	OsVersion *string `json:"os_version,omitempty"`
	// Indicates if the device is registered with an Okta org and is bound to an Okta Verify instance on the device
	Registered *bool `json:"registered,omitempty"`
	ScreenLockType *string `json:"screen_lock_type,omitempty"`
	// The availability of hardware security on the device
	SecureHardwarePresent *bool `json:"secure_hardware_present,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogDevice LogDevice

// NewLogDevice instantiates a new LogDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDevice() *LogDevice {
	this := LogDevice{}
	return &this
}

// NewLogDeviceWithDefaults instantiates a new LogDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDeviceWithDefaults() *LogDevice {
	this := LogDevice{}
	return &this
}

// GetDeviceIntegrator returns the DeviceIntegrator field value if set, zero value otherwise.
func (o *LogDevice) GetDeviceIntegrator() map[string]interface{} {
	if o == nil || o.DeviceIntegrator == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DeviceIntegrator
}

// GetDeviceIntegratorOk returns a tuple with the DeviceIntegrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetDeviceIntegratorOk() (map[string]interface{}, bool) {
	if o == nil || o.DeviceIntegrator == nil {
		return nil, false
	}
	return o.DeviceIntegrator, true
}

// HasDeviceIntegrator returns a boolean if a field has been set.
func (o *LogDevice) HasDeviceIntegrator() bool {
	if o != nil && o.DeviceIntegrator != nil {
		return true
	}

	return false
}

// SetDeviceIntegrator gets a reference to the given map[string]interface{} and assigns it to the DeviceIntegrator field.
func (o *LogDevice) SetDeviceIntegrator(v map[string]interface{}) {
	o.DeviceIntegrator = v
}

// GetDiskEncryptionType returns the DiskEncryptionType field value if set, zero value otherwise.
func (o *LogDevice) GetDiskEncryptionType() string {
	if o == nil || o.DiskEncryptionType == nil {
		var ret string
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetDiskEncryptionTypeOk() (*string, bool) {
	if o == nil || o.DiskEncryptionType == nil {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *LogDevice) HasDiskEncryptionType() bool {
	if o != nil && o.DiskEncryptionType != nil {
		return true
	}

	return false
}

// SetDiskEncryptionType gets a reference to the given string and assigns it to the DiskEncryptionType field.
func (o *LogDevice) SetDiskEncryptionType(v string) {
	o.DiskEncryptionType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogDevice) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogDevice) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogDevice) SetId(v string) {
	o.Id = &v
}

// GetJailbreak returns the Jailbreak field value if set, zero value otherwise.
func (o *LogDevice) GetJailbreak() bool {
	if o == nil || o.Jailbreak == nil {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetJailbreakOk() (*bool, bool) {
	if o == nil || o.Jailbreak == nil {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *LogDevice) HasJailbreak() bool {
	if o != nil && o.Jailbreak != nil {
		return true
	}

	return false
}

// SetJailbreak gets a reference to the given bool and assigns it to the Jailbreak field.
func (o *LogDevice) SetJailbreak(v bool) {
	o.Jailbreak = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *LogDevice) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *LogDevice) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *LogDevice) SetManaged(v bool) {
	o.Managed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogDevice) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogDevice) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogDevice) SetName(v string) {
	o.Name = &v
}

// GetOsPlatform returns the OsPlatform field value if set, zero value otherwise.
func (o *LogDevice) GetOsPlatform() string {
	if o == nil || o.OsPlatform == nil {
		var ret string
		return ret
	}
	return *o.OsPlatform
}

// GetOsPlatformOk returns a tuple with the OsPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetOsPlatformOk() (*string, bool) {
	if o == nil || o.OsPlatform == nil {
		return nil, false
	}
	return o.OsPlatform, true
}

// HasOsPlatform returns a boolean if a field has been set.
func (o *LogDevice) HasOsPlatform() bool {
	if o != nil && o.OsPlatform != nil {
		return true
	}

	return false
}

// SetOsPlatform gets a reference to the given string and assigns it to the OsPlatform field.
func (o *LogDevice) SetOsPlatform(v string) {
	o.OsPlatform = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *LogDevice) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *LogDevice) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *LogDevice) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *LogDevice) GetRegistered() bool {
	if o == nil || o.Registered == nil {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetRegisteredOk() (*bool, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *LogDevice) HasRegistered() bool {
	if o != nil && o.Registered != nil {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given bool and assigns it to the Registered field.
func (o *LogDevice) SetRegistered(v bool) {
	o.Registered = &v
}

// GetScreenLockType returns the ScreenLockType field value if set, zero value otherwise.
func (o *LogDevice) GetScreenLockType() string {
	if o == nil || o.ScreenLockType == nil {
		var ret string
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetScreenLockTypeOk() (*string, bool) {
	if o == nil || o.ScreenLockType == nil {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *LogDevice) HasScreenLockType() bool {
	if o != nil && o.ScreenLockType != nil {
		return true
	}

	return false
}

// SetScreenLockType gets a reference to the given string and assigns it to the ScreenLockType field.
func (o *LogDevice) SetScreenLockType(v string) {
	o.ScreenLockType = &v
}

// GetSecureHardwarePresent returns the SecureHardwarePresent field value if set, zero value otherwise.
func (o *LogDevice) GetSecureHardwarePresent() bool {
	if o == nil || o.SecureHardwarePresent == nil {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || o.SecureHardwarePresent == nil {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *LogDevice) HasSecureHardwarePresent() bool {
	if o != nil && o.SecureHardwarePresent != nil {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *LogDevice) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

func (o LogDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceIntegrator != nil {
		toSerialize["device_integrator"] = o.DeviceIntegrator
	}
	if o.DiskEncryptionType != nil {
		toSerialize["disk_encryption_type"] = o.DiskEncryptionType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Jailbreak != nil {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OsPlatform != nil {
		toSerialize["os_platform"] = o.OsPlatform
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}
	if o.ScreenLockType != nil {
		toSerialize["screen_lock_type"] = o.ScreenLockType
	}
	if o.SecureHardwarePresent != nil {
		toSerialize["secure_hardware_present"] = o.SecureHardwarePresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogDevice) UnmarshalJSON(bytes []byte) (err error) {
	varLogDevice := _LogDevice{}

	err = json.Unmarshal(bytes, &varLogDevice)
	if err == nil {
		*o = LogDevice(varLogDevice)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "device_integrator")
		delete(additionalProperties, "disk_encryption_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "name")
		delete(additionalProperties, "os_platform")
		delete(additionalProperties, "os_version")
		delete(additionalProperties, "registered")
		delete(additionalProperties, "screen_lock_type")
		delete(additionalProperties, "secure_hardware_present")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogDevice struct {
	value *LogDevice
	isSet bool
}

func (v NullableLogDevice) Get() *LogDevice {
	return v.value
}

func (v *NullableLogDevice) Set(val *LogDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDevice(val *LogDevice) *NullableLogDevice {
	return &NullableLogDevice{value: val, isSet: true}
}

func (v NullableLogDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


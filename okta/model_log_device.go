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

// checks if the LogDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogDevice{}

// LogDevice The entity that describes a device enrolled with passwordless authentication using Okta Verify.
type LogDevice struct {
	// The integration platform or software used with the device
	DeviceIntegrator   map[string]interface{} `json:"device_integrator,omitempty"`
	DiskEncryptionType *string                `json:"disk_encryption_type,omitempty"`
	// ID of the device
	Id *string `json:"id,omitempty"`
	// If the device has removed software restrictions
	Jailbreak *bool `json:"jailbreak,omitempty"`
	// Indicates if the device is configured for device management and is registered with Okta
	Managed    *bool   `json:"managed,omitempty"`
	Name       *string `json:"name,omitempty"`
	OsPlatform *string `json:"os_platform,omitempty"`
	OsVersion  *string `json:"os_version,omitempty"`
	// Indicates if the device is registered with an Okta org and is bound to an Okta Verify instance on the device
	Registered     *bool   `json:"registered,omitempty"`
	ScreenLockType *string `json:"screen_lock_type,omitempty"`
	// The availability of hardware security on the device
	SecureHardwarePresent *bool `json:"secure_hardware_present,omitempty"`
	AdditionalProperties  map[string]interface{}
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
	if o == nil || IsNil(o.DeviceIntegrator) {
		var ret map[string]interface{}
		return ret
	}
	return o.DeviceIntegrator
}

// GetDeviceIntegratorOk returns a tuple with the DeviceIntegrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetDeviceIntegratorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DeviceIntegrator) {
		return map[string]interface{}{}, false
	}
	return o.DeviceIntegrator, true
}

// HasDeviceIntegrator returns a boolean if a field has been set.
func (o *LogDevice) HasDeviceIntegrator() bool {
	if o != nil && !IsNil(o.DeviceIntegrator) {
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
	if o == nil || IsNil(o.DiskEncryptionType) {
		var ret string
		return ret
	}
	return *o.DiskEncryptionType
}

// GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetDiskEncryptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskEncryptionType) {
		return nil, false
	}
	return o.DiskEncryptionType, true
}

// HasDiskEncryptionType returns a boolean if a field has been set.
func (o *LogDevice) HasDiskEncryptionType() bool {
	if o != nil && !IsNil(o.DiskEncryptionType) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogDevice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Jailbreak) {
		var ret bool
		return ret
	}
	return *o.Jailbreak
}

// GetJailbreakOk returns a tuple with the Jailbreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetJailbreakOk() (*bool, bool) {
	if o == nil || IsNil(o.Jailbreak) {
		return nil, false
	}
	return o.Jailbreak, true
}

// HasJailbreak returns a boolean if a field has been set.
func (o *LogDevice) HasJailbreak() bool {
	if o != nil && !IsNil(o.Jailbreak) {
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
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *LogDevice) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.OsPlatform) {
		var ret string
		return ret
	}
	return *o.OsPlatform
}

// GetOsPlatformOk returns a tuple with the OsPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetOsPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.OsPlatform) {
		return nil, false
	}
	return o.OsPlatform, true
}

// HasOsPlatform returns a boolean if a field has been set.
func (o *LogDevice) HasOsPlatform() bool {
	if o != nil && !IsNil(o.OsPlatform) {
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
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *LogDevice) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
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
	if o == nil || IsNil(o.Registered) {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.Registered) {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *LogDevice) HasRegistered() bool {
	if o != nil && !IsNil(o.Registered) {
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
	if o == nil || IsNil(o.ScreenLockType) {
		var ret string
		return ret
	}
	return *o.ScreenLockType
}

// GetScreenLockTypeOk returns a tuple with the ScreenLockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetScreenLockTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScreenLockType) {
		return nil, false
	}
	return o.ScreenLockType, true
}

// HasScreenLockType returns a boolean if a field has been set.
func (o *LogDevice) HasScreenLockType() bool {
	if o != nil && !IsNil(o.ScreenLockType) {
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
	if o == nil || IsNil(o.SecureHardwarePresent) {
		var ret bool
		return ret
	}
	return *o.SecureHardwarePresent
}

// GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogDevice) GetSecureHardwarePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureHardwarePresent) {
		return nil, false
	}
	return o.SecureHardwarePresent, true
}

// HasSecureHardwarePresent returns a boolean if a field has been set.
func (o *LogDevice) HasSecureHardwarePresent() bool {
	if o != nil && !IsNil(o.SecureHardwarePresent) {
		return true
	}

	return false
}

// SetSecureHardwarePresent gets a reference to the given bool and assigns it to the SecureHardwarePresent field.
func (o *LogDevice) SetSecureHardwarePresent(v bool) {
	o.SecureHardwarePresent = &v
}

func (o LogDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceIntegrator) {
		toSerialize["device_integrator"] = o.DeviceIntegrator
	}
	if !IsNil(o.DiskEncryptionType) {
		toSerialize["disk_encryption_type"] = o.DiskEncryptionType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Jailbreak) {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OsPlatform) {
		toSerialize["os_platform"] = o.OsPlatform
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.Registered) {
		toSerialize["registered"] = o.Registered
	}
	if !IsNil(o.ScreenLockType) {
		toSerialize["screen_lock_type"] = o.ScreenLockType
	}
	if !IsNil(o.SecureHardwarePresent) {
		toSerialize["secure_hardware_present"] = o.SecureHardwarePresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogDevice) UnmarshalJSON(data []byte) (err error) {
	varLogDevice := _LogDevice{}

	err = json.Unmarshal(data, &varLogDevice)

	if err != nil {
		return err
	}

	*o = LogDevice(varLogDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

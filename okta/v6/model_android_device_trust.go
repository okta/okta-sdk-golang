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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AndroidDeviceTrust type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AndroidDeviceTrust{}

// AndroidDeviceTrust <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Android Device Trust integration provider
type AndroidDeviceTrust struct {
	// Indicates how well a device can enforce app integrity
	DeviceIntegrityLevel *string `json:"deviceIntegrityLevel,omitempty"`
	// Indicates whether a device has a network proxy disabled
	NetworkProxyDisabled *bool `json:"networkProxyDisabled,omitempty"`
	// Indicates if Google Play Protect is enabled on the device and whether it has found known malware
	PlayProtectVerdict *string `json:"playProtectVerdict,omitempty"`
	// Indicates whether the device needs to be on the latest major version available to the device  **Note:** This option requires an `osVersion.dynamicVersionRequirement` value to be supplied with the `osVersion.dynamicVersionRequirement.type` as either `MINIMUM` or `EXACT`.
	RequireMajorVersionUpdate *bool `json:"requireMajorVersionUpdate,omitempty"`
	// Indicates whether a device has a screen lock set, and the type or complexity of the screen lock  **Note:** This option requires a `screenLockType.include` value.
	ScreenLockComplexity *string `json:"screenLockComplexity,omitempty"`
	// Indicates whether Android Debug Bridge (adb) over USB is disabled
	UsbDebuggingDisabled *bool `json:"usbDebuggingDisabled,omitempty"`
	// Indicates whether a device is on a password-protected Wi-Fi network
	WifiSecured          *bool `json:"wifiSecured,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AndroidDeviceTrust AndroidDeviceTrust

// NewAndroidDeviceTrust instantiates a new AndroidDeviceTrust object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidDeviceTrust() *AndroidDeviceTrust {
	this := AndroidDeviceTrust{}
	return &this
}

// NewAndroidDeviceTrustWithDefaults instantiates a new AndroidDeviceTrust object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidDeviceTrustWithDefaults() *AndroidDeviceTrust {
	this := AndroidDeviceTrust{}
	return &this
}

// GetDeviceIntegrityLevel returns the DeviceIntegrityLevel field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetDeviceIntegrityLevel() string {
	if o == nil || IsNil(o.DeviceIntegrityLevel) {
		var ret string
		return ret
	}
	return *o.DeviceIntegrityLevel
}

// GetDeviceIntegrityLevelOk returns a tuple with the DeviceIntegrityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetDeviceIntegrityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceIntegrityLevel) {
		return nil, false
	}
	return o.DeviceIntegrityLevel, true
}

// HasDeviceIntegrityLevel returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasDeviceIntegrityLevel() bool {
	if o != nil && !IsNil(o.DeviceIntegrityLevel) {
		return true
	}

	return false
}

// SetDeviceIntegrityLevel gets a reference to the given string and assigns it to the DeviceIntegrityLevel field.
func (o *AndroidDeviceTrust) SetDeviceIntegrityLevel(v string) {
	o.DeviceIntegrityLevel = &v
}

// GetNetworkProxyDisabled returns the NetworkProxyDisabled field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetNetworkProxyDisabled() bool {
	if o == nil || IsNil(o.NetworkProxyDisabled) {
		var ret bool
		return ret
	}
	return *o.NetworkProxyDisabled
}

// GetNetworkProxyDisabledOk returns a tuple with the NetworkProxyDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetNetworkProxyDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NetworkProxyDisabled) {
		return nil, false
	}
	return o.NetworkProxyDisabled, true
}

// HasNetworkProxyDisabled returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasNetworkProxyDisabled() bool {
	if o != nil && !IsNil(o.NetworkProxyDisabled) {
		return true
	}

	return false
}

// SetNetworkProxyDisabled gets a reference to the given bool and assigns it to the NetworkProxyDisabled field.
func (o *AndroidDeviceTrust) SetNetworkProxyDisabled(v bool) {
	o.NetworkProxyDisabled = &v
}

// GetPlayProtectVerdict returns the PlayProtectVerdict field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetPlayProtectVerdict() string {
	if o == nil || IsNil(o.PlayProtectVerdict) {
		var ret string
		return ret
	}
	return *o.PlayProtectVerdict
}

// GetPlayProtectVerdictOk returns a tuple with the PlayProtectVerdict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetPlayProtectVerdictOk() (*string, bool) {
	if o == nil || IsNil(o.PlayProtectVerdict) {
		return nil, false
	}
	return o.PlayProtectVerdict, true
}

// HasPlayProtectVerdict returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasPlayProtectVerdict() bool {
	if o != nil && !IsNil(o.PlayProtectVerdict) {
		return true
	}

	return false
}

// SetPlayProtectVerdict gets a reference to the given string and assigns it to the PlayProtectVerdict field.
func (o *AndroidDeviceTrust) SetPlayProtectVerdict(v string) {
	o.PlayProtectVerdict = &v
}

// GetRequireMajorVersionUpdate returns the RequireMajorVersionUpdate field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetRequireMajorVersionUpdate() bool {
	if o == nil || IsNil(o.RequireMajorVersionUpdate) {
		var ret bool
		return ret
	}
	return *o.RequireMajorVersionUpdate
}

// GetRequireMajorVersionUpdateOk returns a tuple with the RequireMajorVersionUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetRequireMajorVersionUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireMajorVersionUpdate) {
		return nil, false
	}
	return o.RequireMajorVersionUpdate, true
}

// HasRequireMajorVersionUpdate returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasRequireMajorVersionUpdate() bool {
	if o != nil && !IsNil(o.RequireMajorVersionUpdate) {
		return true
	}

	return false
}

// SetRequireMajorVersionUpdate gets a reference to the given bool and assigns it to the RequireMajorVersionUpdate field.
func (o *AndroidDeviceTrust) SetRequireMajorVersionUpdate(v bool) {
	o.RequireMajorVersionUpdate = &v
}

// GetScreenLockComplexity returns the ScreenLockComplexity field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetScreenLockComplexity() string {
	if o == nil || IsNil(o.ScreenLockComplexity) {
		var ret string
		return ret
	}
	return *o.ScreenLockComplexity
}

// GetScreenLockComplexityOk returns a tuple with the ScreenLockComplexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetScreenLockComplexityOk() (*string, bool) {
	if o == nil || IsNil(o.ScreenLockComplexity) {
		return nil, false
	}
	return o.ScreenLockComplexity, true
}

// HasScreenLockComplexity returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasScreenLockComplexity() bool {
	if o != nil && !IsNil(o.ScreenLockComplexity) {
		return true
	}

	return false
}

// SetScreenLockComplexity gets a reference to the given string and assigns it to the ScreenLockComplexity field.
func (o *AndroidDeviceTrust) SetScreenLockComplexity(v string) {
	o.ScreenLockComplexity = &v
}

// GetUsbDebuggingDisabled returns the UsbDebuggingDisabled field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetUsbDebuggingDisabled() bool {
	if o == nil || IsNil(o.UsbDebuggingDisabled) {
		var ret bool
		return ret
	}
	return *o.UsbDebuggingDisabled
}

// GetUsbDebuggingDisabledOk returns a tuple with the UsbDebuggingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetUsbDebuggingDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UsbDebuggingDisabled) {
		return nil, false
	}
	return o.UsbDebuggingDisabled, true
}

// HasUsbDebuggingDisabled returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasUsbDebuggingDisabled() bool {
	if o != nil && !IsNil(o.UsbDebuggingDisabled) {
		return true
	}

	return false
}

// SetUsbDebuggingDisabled gets a reference to the given bool and assigns it to the UsbDebuggingDisabled field.
func (o *AndroidDeviceTrust) SetUsbDebuggingDisabled(v bool) {
	o.UsbDebuggingDisabled = &v
}

// GetWifiSecured returns the WifiSecured field value if set, zero value otherwise.
func (o *AndroidDeviceTrust) GetWifiSecured() bool {
	if o == nil || IsNil(o.WifiSecured) {
		var ret bool
		return ret
	}
	return *o.WifiSecured
}

// GetWifiSecuredOk returns a tuple with the WifiSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidDeviceTrust) GetWifiSecuredOk() (*bool, bool) {
	if o == nil || IsNil(o.WifiSecured) {
		return nil, false
	}
	return o.WifiSecured, true
}

// HasWifiSecured returns a boolean if a field has been set.
func (o *AndroidDeviceTrust) HasWifiSecured() bool {
	if o != nil && !IsNil(o.WifiSecured) {
		return true
	}

	return false
}

// SetWifiSecured gets a reference to the given bool and assigns it to the WifiSecured field.
func (o *AndroidDeviceTrust) SetWifiSecured(v bool) {
	o.WifiSecured = &v
}

func (o AndroidDeviceTrust) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AndroidDeviceTrust) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceIntegrityLevel) {
		toSerialize["deviceIntegrityLevel"] = o.DeviceIntegrityLevel
	}
	if !IsNil(o.NetworkProxyDisabled) {
		toSerialize["networkProxyDisabled"] = o.NetworkProxyDisabled
	}
	if !IsNil(o.PlayProtectVerdict) {
		toSerialize["playProtectVerdict"] = o.PlayProtectVerdict
	}
	if !IsNil(o.RequireMajorVersionUpdate) {
		toSerialize["requireMajorVersionUpdate"] = o.RequireMajorVersionUpdate
	}
	if !IsNil(o.ScreenLockComplexity) {
		toSerialize["screenLockComplexity"] = o.ScreenLockComplexity
	}
	if !IsNil(o.UsbDebuggingDisabled) {
		toSerialize["usbDebuggingDisabled"] = o.UsbDebuggingDisabled
	}
	if !IsNil(o.WifiSecured) {
		toSerialize["wifiSecured"] = o.WifiSecured
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AndroidDeviceTrust) UnmarshalJSON(data []byte) (err error) {
	varAndroidDeviceTrust := _AndroidDeviceTrust{}

	err = json.Unmarshal(data, &varAndroidDeviceTrust)

	if err != nil {
		return err
	}

	*o = AndroidDeviceTrust(varAndroidDeviceTrust)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceIntegrityLevel")
		delete(additionalProperties, "networkProxyDisabled")
		delete(additionalProperties, "playProtectVerdict")
		delete(additionalProperties, "requireMajorVersionUpdate")
		delete(additionalProperties, "screenLockComplexity")
		delete(additionalProperties, "usbDebuggingDisabled")
		delete(additionalProperties, "wifiSecured")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAndroidDeviceTrust struct {
	value *AndroidDeviceTrust
	isSet bool
}

func (v NullableAndroidDeviceTrust) Get() *AndroidDeviceTrust {
	return v.value
}

func (v *NullableAndroidDeviceTrust) Set(val *AndroidDeviceTrust) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidDeviceTrust) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidDeviceTrust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidDeviceTrust(val *AndroidDeviceTrust) *NullableAndroidDeviceTrust {
	return &NullableAndroidDeviceTrust{value: val, isSet: true}
}

func (v NullableAndroidDeviceTrust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidDeviceTrust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

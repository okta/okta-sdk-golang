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

// DTCChromeOS Google Chrome Device Trust Connector provider
type DTCChromeOS struct {
	// Indicates whether the AllowScreenLock enterprise policy is enabled
	AllowScreenLock *bool `json:"allowScreenLock,omitempty"`
	BrowserVersion *ChromeBrowserVersion `json:"browserVersion,omitempty"`
	// Indicates if a software stack is used to communicate with the DNS server
	BuiltInDnsClientEnabled *bool `json:"builtInDnsClientEnabled,omitempty"`
	// Indicates whether access to the Chrome Remote Desktop application is blocked through a policy
	ChromeRemoteDesktopAppBlocked *bool `json:"chromeRemoteDesktopAppBlocked,omitempty"`
	// Enrollment domain of the customer that is currently managing the device
	DeviceEnrollmentDomain *string `json:"deviceEnrollmentDomain,omitempty"`
	// Indicates whether the main disk is encrypted
	DiskEncrypted *bool `json:"diskEncrypted,omitempty"`
	// Represents the attestation strength used by the Chrome Verified Access API
	KeyTrustLevel *string `json:"keyTrustLevel,omitempty"`
	// Indicates whether the device is enrolled in ChromeOS device management
	ManagedDevice *bool `json:"managedDevice,omitempty"`
	// Indicates whether a firewall is enabled at the OS-level on the device
	OsFirewall *bool `json:"osFirewall,omitempty"`
	OsVersion *OSVersionFourComponents `json:"osVersion,omitempty"`
	// Indicates whether the Password Protection Warning feature is enabled
	PasswordProtectionWarningTrigger *string `json:"passwordProtectionWarningTrigger,omitempty"`
	// Indicates whether enterprise-grade (custom) unsafe URL scanning is enabled
	RealtimeUrlCheckMode *bool `json:"realtimeUrlCheckMode,omitempty"`
	// Represents the current value of the Safe Browsing protection level
	SafeBrowsingProtectionLevel *string `json:"safeBrowsingProtectionLevel,omitempty"`
	// Indicates whether the device is password-protected
	ScreenLockSecured *bool `json:"screenLockSecured,omitempty"`
	// Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled
	SiteIsolationEnabled *bool `json:"siteIsolationEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DTCChromeOS DTCChromeOS

// NewDTCChromeOS instantiates a new DTCChromeOS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTCChromeOS() *DTCChromeOS {
	this := DTCChromeOS{}
	return &this
}

// NewDTCChromeOSWithDefaults instantiates a new DTCChromeOS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTCChromeOSWithDefaults() *DTCChromeOS {
	this := DTCChromeOS{}
	return &this
}

// GetAllowScreenLock returns the AllowScreenLock field value if set, zero value otherwise.
func (o *DTCChromeOS) GetAllowScreenLock() bool {
	if o == nil || o.AllowScreenLock == nil {
		var ret bool
		return ret
	}
	return *o.AllowScreenLock
}

// GetAllowScreenLockOk returns a tuple with the AllowScreenLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetAllowScreenLockOk() (*bool, bool) {
	if o == nil || o.AllowScreenLock == nil {
		return nil, false
	}
	return o.AllowScreenLock, true
}

// HasAllowScreenLock returns a boolean if a field has been set.
func (o *DTCChromeOS) HasAllowScreenLock() bool {
	if o != nil && o.AllowScreenLock != nil {
		return true
	}

	return false
}

// SetAllowScreenLock gets a reference to the given bool and assigns it to the AllowScreenLock field.
func (o *DTCChromeOS) SetAllowScreenLock(v bool) {
	o.AllowScreenLock = &v
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *DTCChromeOS) GetBrowserVersion() ChromeBrowserVersion {
	if o == nil || o.BrowserVersion == nil {
		var ret ChromeBrowserVersion
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetBrowserVersionOk() (*ChromeBrowserVersion, bool) {
	if o == nil || o.BrowserVersion == nil {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *DTCChromeOS) HasBrowserVersion() bool {
	if o != nil && o.BrowserVersion != nil {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given ChromeBrowserVersion and assigns it to the BrowserVersion field.
func (o *DTCChromeOS) SetBrowserVersion(v ChromeBrowserVersion) {
	o.BrowserVersion = &v
}

// GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field value if set, zero value otherwise.
func (o *DTCChromeOS) GetBuiltInDnsClientEnabled() bool {
	if o == nil || o.BuiltInDnsClientEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BuiltInDnsClientEnabled
}

// GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetBuiltInDnsClientEnabledOk() (*bool, bool) {
	if o == nil || o.BuiltInDnsClientEnabled == nil {
		return nil, false
	}
	return o.BuiltInDnsClientEnabled, true
}

// HasBuiltInDnsClientEnabled returns a boolean if a field has been set.
func (o *DTCChromeOS) HasBuiltInDnsClientEnabled() bool {
	if o != nil && o.BuiltInDnsClientEnabled != nil {
		return true
	}

	return false
}

// SetBuiltInDnsClientEnabled gets a reference to the given bool and assigns it to the BuiltInDnsClientEnabled field.
func (o *DTCChromeOS) SetBuiltInDnsClientEnabled(v bool) {
	o.BuiltInDnsClientEnabled = &v
}

// GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field value if set, zero value otherwise.
func (o *DTCChromeOS) GetChromeRemoteDesktopAppBlocked() bool {
	if o == nil || o.ChromeRemoteDesktopAppBlocked == nil {
		var ret bool
		return ret
	}
	return *o.ChromeRemoteDesktopAppBlocked
}

// GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool) {
	if o == nil || o.ChromeRemoteDesktopAppBlocked == nil {
		return nil, false
	}
	return o.ChromeRemoteDesktopAppBlocked, true
}

// HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.
func (o *DTCChromeOS) HasChromeRemoteDesktopAppBlocked() bool {
	if o != nil && o.ChromeRemoteDesktopAppBlocked != nil {
		return true
	}

	return false
}

// SetChromeRemoteDesktopAppBlocked gets a reference to the given bool and assigns it to the ChromeRemoteDesktopAppBlocked field.
func (o *DTCChromeOS) SetChromeRemoteDesktopAppBlocked(v bool) {
	o.ChromeRemoteDesktopAppBlocked = &v
}

// GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field value if set, zero value otherwise.
func (o *DTCChromeOS) GetDeviceEnrollmentDomain() string {
	if o == nil || o.DeviceEnrollmentDomain == nil {
		var ret string
		return ret
	}
	return *o.DeviceEnrollmentDomain
}

// GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetDeviceEnrollmentDomainOk() (*string, bool) {
	if o == nil || o.DeviceEnrollmentDomain == nil {
		return nil, false
	}
	return o.DeviceEnrollmentDomain, true
}

// HasDeviceEnrollmentDomain returns a boolean if a field has been set.
func (o *DTCChromeOS) HasDeviceEnrollmentDomain() bool {
	if o != nil && o.DeviceEnrollmentDomain != nil {
		return true
	}

	return false
}

// SetDeviceEnrollmentDomain gets a reference to the given string and assigns it to the DeviceEnrollmentDomain field.
func (o *DTCChromeOS) SetDeviceEnrollmentDomain(v string) {
	o.DeviceEnrollmentDomain = &v
}

// GetDiskEncrypted returns the DiskEncrypted field value if set, zero value otherwise.
func (o *DTCChromeOS) GetDiskEncrypted() bool {
	if o == nil || o.DiskEncrypted == nil {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || o.DiskEncrypted == nil {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *DTCChromeOS) HasDiskEncrypted() bool {
	if o != nil && o.DiskEncrypted != nil {
		return true
	}

	return false
}

// SetDiskEncrypted gets a reference to the given bool and assigns it to the DiskEncrypted field.
func (o *DTCChromeOS) SetDiskEncrypted(v bool) {
	o.DiskEncrypted = &v
}

// GetKeyTrustLevel returns the KeyTrustLevel field value if set, zero value otherwise.
func (o *DTCChromeOS) GetKeyTrustLevel() string {
	if o == nil || o.KeyTrustLevel == nil {
		var ret string
		return ret
	}
	return *o.KeyTrustLevel
}

// GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetKeyTrustLevelOk() (*string, bool) {
	if o == nil || o.KeyTrustLevel == nil {
		return nil, false
	}
	return o.KeyTrustLevel, true
}

// HasKeyTrustLevel returns a boolean if a field has been set.
func (o *DTCChromeOS) HasKeyTrustLevel() bool {
	if o != nil && o.KeyTrustLevel != nil {
		return true
	}

	return false
}

// SetKeyTrustLevel gets a reference to the given string and assigns it to the KeyTrustLevel field.
func (o *DTCChromeOS) SetKeyTrustLevel(v string) {
	o.KeyTrustLevel = &v
}

// GetManagedDevice returns the ManagedDevice field value if set, zero value otherwise.
func (o *DTCChromeOS) GetManagedDevice() bool {
	if o == nil || o.ManagedDevice == nil {
		var ret bool
		return ret
	}
	return *o.ManagedDevice
}

// GetManagedDeviceOk returns a tuple with the ManagedDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetManagedDeviceOk() (*bool, bool) {
	if o == nil || o.ManagedDevice == nil {
		return nil, false
	}
	return o.ManagedDevice, true
}

// HasManagedDevice returns a boolean if a field has been set.
func (o *DTCChromeOS) HasManagedDevice() bool {
	if o != nil && o.ManagedDevice != nil {
		return true
	}

	return false
}

// SetManagedDevice gets a reference to the given bool and assigns it to the ManagedDevice field.
func (o *DTCChromeOS) SetManagedDevice(v bool) {
	o.ManagedDevice = &v
}

// GetOsFirewall returns the OsFirewall field value if set, zero value otherwise.
func (o *DTCChromeOS) GetOsFirewall() bool {
	if o == nil || o.OsFirewall == nil {
		var ret bool
		return ret
	}
	return *o.OsFirewall
}

// GetOsFirewallOk returns a tuple with the OsFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetOsFirewallOk() (*bool, bool) {
	if o == nil || o.OsFirewall == nil {
		return nil, false
	}
	return o.OsFirewall, true
}

// HasOsFirewall returns a boolean if a field has been set.
func (o *DTCChromeOS) HasOsFirewall() bool {
	if o != nil && o.OsFirewall != nil {
		return true
	}

	return false
}

// SetOsFirewall gets a reference to the given bool and assigns it to the OsFirewall field.
func (o *DTCChromeOS) SetOsFirewall(v bool) {
	o.OsFirewall = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DTCChromeOS) GetOsVersion() OSVersionFourComponents {
	if o == nil || o.OsVersion == nil {
		var ret OSVersionFourComponents
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetOsVersionOk() (*OSVersionFourComponents, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DTCChromeOS) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersionFourComponents and assigns it to the OsVersion field.
func (o *DTCChromeOS) SetOsVersion(v OSVersionFourComponents) {
	o.OsVersion = &v
}

// GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field value if set, zero value otherwise.
func (o *DTCChromeOS) GetPasswordProtectionWarningTrigger() string {
	if o == nil || o.PasswordProtectionWarningTrigger == nil {
		var ret string
		return ret
	}
	return *o.PasswordProtectionWarningTrigger
}

// GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetPasswordProtectionWarningTriggerOk() (*string, bool) {
	if o == nil || o.PasswordProtectionWarningTrigger == nil {
		return nil, false
	}
	return o.PasswordProtectionWarningTrigger, true
}

// HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.
func (o *DTCChromeOS) HasPasswordProtectionWarningTrigger() bool {
	if o != nil && o.PasswordProtectionWarningTrigger != nil {
		return true
	}

	return false
}

// SetPasswordProtectionWarningTrigger gets a reference to the given string and assigns it to the PasswordProtectionWarningTrigger field.
func (o *DTCChromeOS) SetPasswordProtectionWarningTrigger(v string) {
	o.PasswordProtectionWarningTrigger = &v
}

// GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field value if set, zero value otherwise.
func (o *DTCChromeOS) GetRealtimeUrlCheckMode() bool {
	if o == nil || o.RealtimeUrlCheckMode == nil {
		var ret bool
		return ret
	}
	return *o.RealtimeUrlCheckMode
}

// GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetRealtimeUrlCheckModeOk() (*bool, bool) {
	if o == nil || o.RealtimeUrlCheckMode == nil {
		return nil, false
	}
	return o.RealtimeUrlCheckMode, true
}

// HasRealtimeUrlCheckMode returns a boolean if a field has been set.
func (o *DTCChromeOS) HasRealtimeUrlCheckMode() bool {
	if o != nil && o.RealtimeUrlCheckMode != nil {
		return true
	}

	return false
}

// SetRealtimeUrlCheckMode gets a reference to the given bool and assigns it to the RealtimeUrlCheckMode field.
func (o *DTCChromeOS) SetRealtimeUrlCheckMode(v bool) {
	o.RealtimeUrlCheckMode = &v
}

// GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field value if set, zero value otherwise.
func (o *DTCChromeOS) GetSafeBrowsingProtectionLevel() string {
	if o == nil || o.SafeBrowsingProtectionLevel == nil {
		var ret string
		return ret
	}
	return *o.SafeBrowsingProtectionLevel
}

// GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetSafeBrowsingProtectionLevelOk() (*string, bool) {
	if o == nil || o.SafeBrowsingProtectionLevel == nil {
		return nil, false
	}
	return o.SafeBrowsingProtectionLevel, true
}

// HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.
func (o *DTCChromeOS) HasSafeBrowsingProtectionLevel() bool {
	if o != nil && o.SafeBrowsingProtectionLevel != nil {
		return true
	}

	return false
}

// SetSafeBrowsingProtectionLevel gets a reference to the given string and assigns it to the SafeBrowsingProtectionLevel field.
func (o *DTCChromeOS) SetSafeBrowsingProtectionLevel(v string) {
	o.SafeBrowsingProtectionLevel = &v
}

// GetScreenLockSecured returns the ScreenLockSecured field value if set, zero value otherwise.
func (o *DTCChromeOS) GetScreenLockSecured() bool {
	if o == nil || o.ScreenLockSecured == nil {
		var ret bool
		return ret
	}
	return *o.ScreenLockSecured
}

// GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetScreenLockSecuredOk() (*bool, bool) {
	if o == nil || o.ScreenLockSecured == nil {
		return nil, false
	}
	return o.ScreenLockSecured, true
}

// HasScreenLockSecured returns a boolean if a field has been set.
func (o *DTCChromeOS) HasScreenLockSecured() bool {
	if o != nil && o.ScreenLockSecured != nil {
		return true
	}

	return false
}

// SetScreenLockSecured gets a reference to the given bool and assigns it to the ScreenLockSecured field.
func (o *DTCChromeOS) SetScreenLockSecured(v bool) {
	o.ScreenLockSecured = &v
}

// GetSiteIsolationEnabled returns the SiteIsolationEnabled field value if set, zero value otherwise.
func (o *DTCChromeOS) GetSiteIsolationEnabled() bool {
	if o == nil || o.SiteIsolationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SiteIsolationEnabled
}

// GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCChromeOS) GetSiteIsolationEnabledOk() (*bool, bool) {
	if o == nil || o.SiteIsolationEnabled == nil {
		return nil, false
	}
	return o.SiteIsolationEnabled, true
}

// HasSiteIsolationEnabled returns a boolean if a field has been set.
func (o *DTCChromeOS) HasSiteIsolationEnabled() bool {
	if o != nil && o.SiteIsolationEnabled != nil {
		return true
	}

	return false
}

// SetSiteIsolationEnabled gets a reference to the given bool and assigns it to the SiteIsolationEnabled field.
func (o *DTCChromeOS) SetSiteIsolationEnabled(v bool) {
	o.SiteIsolationEnabled = &v
}

func (o DTCChromeOS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowScreenLock != nil {
		toSerialize["allowScreenLock"] = o.AllowScreenLock
	}
	if o.BrowserVersion != nil {
		toSerialize["browserVersion"] = o.BrowserVersion
	}
	if o.BuiltInDnsClientEnabled != nil {
		toSerialize["builtInDnsClientEnabled"] = o.BuiltInDnsClientEnabled
	}
	if o.ChromeRemoteDesktopAppBlocked != nil {
		toSerialize["chromeRemoteDesktopAppBlocked"] = o.ChromeRemoteDesktopAppBlocked
	}
	if o.DeviceEnrollmentDomain != nil {
		toSerialize["deviceEnrollmentDomain"] = o.DeviceEnrollmentDomain
	}
	if o.DiskEncrypted != nil {
		toSerialize["diskEncrypted"] = o.DiskEncrypted
	}
	if o.KeyTrustLevel != nil {
		toSerialize["keyTrustLevel"] = o.KeyTrustLevel
	}
	if o.ManagedDevice != nil {
		toSerialize["managedDevice"] = o.ManagedDevice
	}
	if o.OsFirewall != nil {
		toSerialize["osFirewall"] = o.OsFirewall
	}
	if o.OsVersion != nil {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.PasswordProtectionWarningTrigger != nil {
		toSerialize["passwordProtectionWarningTrigger"] = o.PasswordProtectionWarningTrigger
	}
	if o.RealtimeUrlCheckMode != nil {
		toSerialize["realtimeUrlCheckMode"] = o.RealtimeUrlCheckMode
	}
	if o.SafeBrowsingProtectionLevel != nil {
		toSerialize["safeBrowsingProtectionLevel"] = o.SafeBrowsingProtectionLevel
	}
	if o.ScreenLockSecured != nil {
		toSerialize["screenLockSecured"] = o.ScreenLockSecured
	}
	if o.SiteIsolationEnabled != nil {
		toSerialize["siteIsolationEnabled"] = o.SiteIsolationEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DTCChromeOS) UnmarshalJSON(bytes []byte) (err error) {
	varDTCChromeOS := _DTCChromeOS{}

	err = json.Unmarshal(bytes, &varDTCChromeOS)
	if err == nil {
		*o = DTCChromeOS(varDTCChromeOS)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowScreenLock")
		delete(additionalProperties, "browserVersion")
		delete(additionalProperties, "builtInDnsClientEnabled")
		delete(additionalProperties, "chromeRemoteDesktopAppBlocked")
		delete(additionalProperties, "deviceEnrollmentDomain")
		delete(additionalProperties, "diskEncrypted")
		delete(additionalProperties, "keyTrustLevel")
		delete(additionalProperties, "managedDevice")
		delete(additionalProperties, "osFirewall")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "passwordProtectionWarningTrigger")
		delete(additionalProperties, "realtimeUrlCheckMode")
		delete(additionalProperties, "safeBrowsingProtectionLevel")
		delete(additionalProperties, "screenLockSecured")
		delete(additionalProperties, "siteIsolationEnabled")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDTCChromeOS struct {
	value *DTCChromeOS
	isSet bool
}

func (v NullableDTCChromeOS) Get() *DTCChromeOS {
	return v.value
}

func (v *NullableDTCChromeOS) Set(val *DTCChromeOS) {
	v.value = val
	v.isSet = true
}

func (v NullableDTCChromeOS) IsSet() bool {
	return v.isSet
}

func (v *NullableDTCChromeOS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTCChromeOS(val *DTCChromeOS) *NullableDTCChromeOS {
	return &NullableDTCChromeOS{value: val, isSet: true}
}

func (v NullableDTCChromeOS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTCChromeOS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


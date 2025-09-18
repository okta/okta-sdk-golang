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

// checks if the DTCMacOS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTCMacOS{}

// DTCMacOS Google Chrome Device Trust Connector provider
type DTCMacOS struct {
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
	// Indicates whether a firewall is enabled at the OS-level on the device
	OsFirewall *bool                     `json:"osFirewall,omitempty"`
	OsVersion  *OSVersionThreeComponents `json:"osVersion,omitempty"`
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

type _DTCMacOS DTCMacOS

// NewDTCMacOS instantiates a new DTCMacOS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTCMacOS() *DTCMacOS {
	this := DTCMacOS{}
	return &this
}

// NewDTCMacOSWithDefaults instantiates a new DTCMacOS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTCMacOSWithDefaults() *DTCMacOS {
	this := DTCMacOS{}
	return &this
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *DTCMacOS) GetBrowserVersion() ChromeBrowserVersion {
	if o == nil || IsNil(o.BrowserVersion) {
		var ret ChromeBrowserVersion
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetBrowserVersionOk() (*ChromeBrowserVersion, bool) {
	if o == nil || IsNil(o.BrowserVersion) {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *DTCMacOS) HasBrowserVersion() bool {
	if o != nil && !IsNil(o.BrowserVersion) {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given ChromeBrowserVersion and assigns it to the BrowserVersion field.
func (o *DTCMacOS) SetBrowserVersion(v ChromeBrowserVersion) {
	o.BrowserVersion = &v
}

// GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field value if set, zero value otherwise.
func (o *DTCMacOS) GetBuiltInDnsClientEnabled() bool {
	if o == nil || IsNil(o.BuiltInDnsClientEnabled) {
		var ret bool
		return ret
	}
	return *o.BuiltInDnsClientEnabled
}

// GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetBuiltInDnsClientEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BuiltInDnsClientEnabled) {
		return nil, false
	}
	return o.BuiltInDnsClientEnabled, true
}

// HasBuiltInDnsClientEnabled returns a boolean if a field has been set.
func (o *DTCMacOS) HasBuiltInDnsClientEnabled() bool {
	if o != nil && !IsNil(o.BuiltInDnsClientEnabled) {
		return true
	}

	return false
}

// SetBuiltInDnsClientEnabled gets a reference to the given bool and assigns it to the BuiltInDnsClientEnabled field.
func (o *DTCMacOS) SetBuiltInDnsClientEnabled(v bool) {
	o.BuiltInDnsClientEnabled = &v
}

// GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field value if set, zero value otherwise.
func (o *DTCMacOS) GetChromeRemoteDesktopAppBlocked() bool {
	if o == nil || IsNil(o.ChromeRemoteDesktopAppBlocked) {
		var ret bool
		return ret
	}
	return *o.ChromeRemoteDesktopAppBlocked
}

// GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.ChromeRemoteDesktopAppBlocked) {
		return nil, false
	}
	return o.ChromeRemoteDesktopAppBlocked, true
}

// HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.
func (o *DTCMacOS) HasChromeRemoteDesktopAppBlocked() bool {
	if o != nil && !IsNil(o.ChromeRemoteDesktopAppBlocked) {
		return true
	}

	return false
}

// SetChromeRemoteDesktopAppBlocked gets a reference to the given bool and assigns it to the ChromeRemoteDesktopAppBlocked field.
func (o *DTCMacOS) SetChromeRemoteDesktopAppBlocked(v bool) {
	o.ChromeRemoteDesktopAppBlocked = &v
}

// GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field value if set, zero value otherwise.
func (o *DTCMacOS) GetDeviceEnrollmentDomain() string {
	if o == nil || IsNil(o.DeviceEnrollmentDomain) {
		var ret string
		return ret
	}
	return *o.DeviceEnrollmentDomain
}

// GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetDeviceEnrollmentDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceEnrollmentDomain) {
		return nil, false
	}
	return o.DeviceEnrollmentDomain, true
}

// HasDeviceEnrollmentDomain returns a boolean if a field has been set.
func (o *DTCMacOS) HasDeviceEnrollmentDomain() bool {
	if o != nil && !IsNil(o.DeviceEnrollmentDomain) {
		return true
	}

	return false
}

// SetDeviceEnrollmentDomain gets a reference to the given string and assigns it to the DeviceEnrollmentDomain field.
func (o *DTCMacOS) SetDeviceEnrollmentDomain(v string) {
	o.DeviceEnrollmentDomain = &v
}

// GetDiskEncrypted returns the DiskEncrypted field value if set, zero value otherwise.
func (o *DTCMacOS) GetDiskEncrypted() bool {
	if o == nil || IsNil(o.DiskEncrypted) {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskEncrypted) {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *DTCMacOS) HasDiskEncrypted() bool {
	if o != nil && !IsNil(o.DiskEncrypted) {
		return true
	}

	return false
}

// SetDiskEncrypted gets a reference to the given bool and assigns it to the DiskEncrypted field.
func (o *DTCMacOS) SetDiskEncrypted(v bool) {
	o.DiskEncrypted = &v
}

// GetKeyTrustLevel returns the KeyTrustLevel field value if set, zero value otherwise.
func (o *DTCMacOS) GetKeyTrustLevel() string {
	if o == nil || IsNil(o.KeyTrustLevel) {
		var ret string
		return ret
	}
	return *o.KeyTrustLevel
}

// GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetKeyTrustLevelOk() (*string, bool) {
	if o == nil || IsNil(o.KeyTrustLevel) {
		return nil, false
	}
	return o.KeyTrustLevel, true
}

// HasKeyTrustLevel returns a boolean if a field has been set.
func (o *DTCMacOS) HasKeyTrustLevel() bool {
	if o != nil && !IsNil(o.KeyTrustLevel) {
		return true
	}

	return false
}

// SetKeyTrustLevel gets a reference to the given string and assigns it to the KeyTrustLevel field.
func (o *DTCMacOS) SetKeyTrustLevel(v string) {
	o.KeyTrustLevel = &v
}

// GetOsFirewall returns the OsFirewall field value if set, zero value otherwise.
func (o *DTCMacOS) GetOsFirewall() bool {
	if o == nil || IsNil(o.OsFirewall) {
		var ret bool
		return ret
	}
	return *o.OsFirewall
}

// GetOsFirewallOk returns a tuple with the OsFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetOsFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.OsFirewall) {
		return nil, false
	}
	return o.OsFirewall, true
}

// HasOsFirewall returns a boolean if a field has been set.
func (o *DTCMacOS) HasOsFirewall() bool {
	if o != nil && !IsNil(o.OsFirewall) {
		return true
	}

	return false
}

// SetOsFirewall gets a reference to the given bool and assigns it to the OsFirewall field.
func (o *DTCMacOS) SetOsFirewall(v bool) {
	o.OsFirewall = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DTCMacOS) GetOsVersion() OSVersionThreeComponents {
	if o == nil || IsNil(o.OsVersion) {
		var ret OSVersionThreeComponents
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetOsVersionOk() (*OSVersionThreeComponents, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DTCMacOS) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersionThreeComponents and assigns it to the OsVersion field.
func (o *DTCMacOS) SetOsVersion(v OSVersionThreeComponents) {
	o.OsVersion = &v
}

// GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field value if set, zero value otherwise.
func (o *DTCMacOS) GetPasswordProtectionWarningTrigger() string {
	if o == nil || IsNil(o.PasswordProtectionWarningTrigger) {
		var ret string
		return ret
	}
	return *o.PasswordProtectionWarningTrigger
}

// GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetPasswordProtectionWarningTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordProtectionWarningTrigger) {
		return nil, false
	}
	return o.PasswordProtectionWarningTrigger, true
}

// HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.
func (o *DTCMacOS) HasPasswordProtectionWarningTrigger() bool {
	if o != nil && !IsNil(o.PasswordProtectionWarningTrigger) {
		return true
	}

	return false
}

// SetPasswordProtectionWarningTrigger gets a reference to the given string and assigns it to the PasswordProtectionWarningTrigger field.
func (o *DTCMacOS) SetPasswordProtectionWarningTrigger(v string) {
	o.PasswordProtectionWarningTrigger = &v
}

// GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field value if set, zero value otherwise.
func (o *DTCMacOS) GetRealtimeUrlCheckMode() bool {
	if o == nil || IsNil(o.RealtimeUrlCheckMode) {
		var ret bool
		return ret
	}
	return *o.RealtimeUrlCheckMode
}

// GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetRealtimeUrlCheckModeOk() (*bool, bool) {
	if o == nil || IsNil(o.RealtimeUrlCheckMode) {
		return nil, false
	}
	return o.RealtimeUrlCheckMode, true
}

// HasRealtimeUrlCheckMode returns a boolean if a field has been set.
func (o *DTCMacOS) HasRealtimeUrlCheckMode() bool {
	if o != nil && !IsNil(o.RealtimeUrlCheckMode) {
		return true
	}

	return false
}

// SetRealtimeUrlCheckMode gets a reference to the given bool and assigns it to the RealtimeUrlCheckMode field.
func (o *DTCMacOS) SetRealtimeUrlCheckMode(v bool) {
	o.RealtimeUrlCheckMode = &v
}

// GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field value if set, zero value otherwise.
func (o *DTCMacOS) GetSafeBrowsingProtectionLevel() string {
	if o == nil || IsNil(o.SafeBrowsingProtectionLevel) {
		var ret string
		return ret
	}
	return *o.SafeBrowsingProtectionLevel
}

// GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetSafeBrowsingProtectionLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SafeBrowsingProtectionLevel) {
		return nil, false
	}
	return o.SafeBrowsingProtectionLevel, true
}

// HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.
func (o *DTCMacOS) HasSafeBrowsingProtectionLevel() bool {
	if o != nil && !IsNil(o.SafeBrowsingProtectionLevel) {
		return true
	}

	return false
}

// SetSafeBrowsingProtectionLevel gets a reference to the given string and assigns it to the SafeBrowsingProtectionLevel field.
func (o *DTCMacOS) SetSafeBrowsingProtectionLevel(v string) {
	o.SafeBrowsingProtectionLevel = &v
}

// GetScreenLockSecured returns the ScreenLockSecured field value if set, zero value otherwise.
func (o *DTCMacOS) GetScreenLockSecured() bool {
	if o == nil || IsNil(o.ScreenLockSecured) {
		var ret bool
		return ret
	}
	return *o.ScreenLockSecured
}

// GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetScreenLockSecuredOk() (*bool, bool) {
	if o == nil || IsNil(o.ScreenLockSecured) {
		return nil, false
	}
	return o.ScreenLockSecured, true
}

// HasScreenLockSecured returns a boolean if a field has been set.
func (o *DTCMacOS) HasScreenLockSecured() bool {
	if o != nil && !IsNil(o.ScreenLockSecured) {
		return true
	}

	return false
}

// SetScreenLockSecured gets a reference to the given bool and assigns it to the ScreenLockSecured field.
func (o *DTCMacOS) SetScreenLockSecured(v bool) {
	o.ScreenLockSecured = &v
}

// GetSiteIsolationEnabled returns the SiteIsolationEnabled field value if set, zero value otherwise.
func (o *DTCMacOS) GetSiteIsolationEnabled() bool {
	if o == nil || IsNil(o.SiteIsolationEnabled) {
		var ret bool
		return ret
	}
	return *o.SiteIsolationEnabled
}

// GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCMacOS) GetSiteIsolationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SiteIsolationEnabled) {
		return nil, false
	}
	return o.SiteIsolationEnabled, true
}

// HasSiteIsolationEnabled returns a boolean if a field has been set.
func (o *DTCMacOS) HasSiteIsolationEnabled() bool {
	if o != nil && !IsNil(o.SiteIsolationEnabled) {
		return true
	}

	return false
}

// SetSiteIsolationEnabled gets a reference to the given bool and assigns it to the SiteIsolationEnabled field.
func (o *DTCMacOS) SetSiteIsolationEnabled(v bool) {
	o.SiteIsolationEnabled = &v
}

func (o DTCMacOS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTCMacOS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrowserVersion) {
		toSerialize["browserVersion"] = o.BrowserVersion
	}
	if !IsNil(o.BuiltInDnsClientEnabled) {
		toSerialize["builtInDnsClientEnabled"] = o.BuiltInDnsClientEnabled
	}
	if !IsNil(o.ChromeRemoteDesktopAppBlocked) {
		toSerialize["chromeRemoteDesktopAppBlocked"] = o.ChromeRemoteDesktopAppBlocked
	}
	if !IsNil(o.DeviceEnrollmentDomain) {
		toSerialize["deviceEnrollmentDomain"] = o.DeviceEnrollmentDomain
	}
	if !IsNil(o.DiskEncrypted) {
		toSerialize["diskEncrypted"] = o.DiskEncrypted
	}
	if !IsNil(o.KeyTrustLevel) {
		toSerialize["keyTrustLevel"] = o.KeyTrustLevel
	}
	if !IsNil(o.OsFirewall) {
		toSerialize["osFirewall"] = o.OsFirewall
	}
	if !IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !IsNil(o.PasswordProtectionWarningTrigger) {
		toSerialize["passwordProtectionWarningTrigger"] = o.PasswordProtectionWarningTrigger
	}
	if !IsNil(o.RealtimeUrlCheckMode) {
		toSerialize["realtimeUrlCheckMode"] = o.RealtimeUrlCheckMode
	}
	if !IsNil(o.SafeBrowsingProtectionLevel) {
		toSerialize["safeBrowsingProtectionLevel"] = o.SafeBrowsingProtectionLevel
	}
	if !IsNil(o.ScreenLockSecured) {
		toSerialize["screenLockSecured"] = o.ScreenLockSecured
	}
	if !IsNil(o.SiteIsolationEnabled) {
		toSerialize["siteIsolationEnabled"] = o.SiteIsolationEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DTCMacOS) UnmarshalJSON(data []byte) (err error) {
	varDTCMacOS := _DTCMacOS{}

	err = json.Unmarshal(data, &varDTCMacOS)

	if err != nil {
		return err
	}

	*o = DTCMacOS(varDTCMacOS)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "browserVersion")
		delete(additionalProperties, "builtInDnsClientEnabled")
		delete(additionalProperties, "chromeRemoteDesktopAppBlocked")
		delete(additionalProperties, "deviceEnrollmentDomain")
		delete(additionalProperties, "diskEncrypted")
		delete(additionalProperties, "keyTrustLevel")
		delete(additionalProperties, "osFirewall")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "passwordProtectionWarningTrigger")
		delete(additionalProperties, "realtimeUrlCheckMode")
		delete(additionalProperties, "safeBrowsingProtectionLevel")
		delete(additionalProperties, "screenLockSecured")
		delete(additionalProperties, "siteIsolationEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDTCMacOS struct {
	value *DTCMacOS
	isSet bool
}

func (v NullableDTCMacOS) Get() *DTCMacOS {
	return v.value
}

func (v *NullableDTCMacOS) Set(val *DTCMacOS) {
	v.value = val
	v.isSet = true
}

func (v NullableDTCMacOS) IsSet() bool {
	return v.isSet
}

func (v *NullableDTCMacOS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTCMacOS(val *DTCMacOS) *NullableDTCMacOS {
	return &NullableDTCMacOS{value: val, isSet: true}
}

func (v NullableDTCMacOS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTCMacOS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

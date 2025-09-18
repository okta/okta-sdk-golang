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

// checks if the DTCWindows type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTCWindows{}

// DTCWindows Google Chrome Device Trust Connector provider
type DTCWindows struct {
	BrowserVersion *ChromeBrowserVersion `json:"browserVersion,omitempty"`
	// Indicates if a software stack is used to communicate with the DNS server
	BuiltInDnsClientEnabled *bool `json:"builtInDnsClientEnabled,omitempty"`
	// Indicates whether access to the Chrome Remote Desktop application is blocked through a policy
	ChromeRemoteDesktopAppBlocked *bool `json:"chromeRemoteDesktopAppBlocked,omitempty"`
	// Agent ID of an installed CrowdStrike agent
	CrowdStrikeAgentId *string `json:"crowdStrikeAgentId,omitempty"`
	// Customer ID of an installed CrowdStrike agent
	CrowdStrikeCustomerId *string `json:"crowdStrikeCustomerId,omitempty"`
	// Enrollment domain of the customer that is currently managing the device
	DeviceEnrollmentDomain *string `json:"deviceEnrollmentDomain,omitempty"`
	// Indicates whether the main disk is encrypted
	DiskEncrypted *bool `json:"diskEncrypted,omitempty"`
	// Represents the attestation strength used by the Chrome Verified Access API
	KeyTrustLevel *string `json:"keyTrustLevel,omitempty"`
	// Indicates whether a firewall is enabled at the OS-level on the device
	OsFirewall *bool                    `json:"osFirewall,omitempty"`
	OsVersion  *OSVersionFourComponents `json:"osVersion,omitempty"`
	// Indicates whether the Password Protection Warning feature is enabled
	PasswordProtectionWarningTrigger *string `json:"passwordProtectionWarningTrigger,omitempty"`
	// Indicates whether enterprise-grade (custom) unsafe URL scanning is enabled
	RealtimeUrlCheckMode *bool `json:"realtimeUrlCheckMode,omitempty"`
	// Represents the current value of the Safe Browsing protection level
	SafeBrowsingProtectionLevel *string `json:"safeBrowsingProtectionLevel,omitempty"`
	// Indicates whether the device is password-protected
	ScreenLockSecured *bool `json:"screenLockSecured,omitempty"`
	// Indicates whether the device's startup software has its Secure Boot feature enabled
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
	// Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled
	SiteIsolationEnabled *bool `json:"siteIsolationEnabled,omitempty"`
	// Indicates whether Chrome is blocking third-party software injection
	// Deprecated
	ThirdPartyBlockingEnabled *bool `json:"thirdPartyBlockingEnabled,omitempty"`
	// Windows domain that the current machine has joined
	WindowsMachineDomain *string `json:"windowsMachineDomain,omitempty"`
	// Windows domain for the current OS user
	WindowsUserDomain    *string `json:"windowsUserDomain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DTCWindows DTCWindows

// NewDTCWindows instantiates a new DTCWindows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTCWindows() *DTCWindows {
	this := DTCWindows{}
	return &this
}

// NewDTCWindowsWithDefaults instantiates a new DTCWindows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTCWindowsWithDefaults() *DTCWindows {
	this := DTCWindows{}
	return &this
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *DTCWindows) GetBrowserVersion() ChromeBrowserVersion {
	if o == nil || IsNil(o.BrowserVersion) {
		var ret ChromeBrowserVersion
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetBrowserVersionOk() (*ChromeBrowserVersion, bool) {
	if o == nil || IsNil(o.BrowserVersion) {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *DTCWindows) HasBrowserVersion() bool {
	if o != nil && !IsNil(o.BrowserVersion) {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given ChromeBrowserVersion and assigns it to the BrowserVersion field.
func (o *DTCWindows) SetBrowserVersion(v ChromeBrowserVersion) {
	o.BrowserVersion = &v
}

// GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field value if set, zero value otherwise.
func (o *DTCWindows) GetBuiltInDnsClientEnabled() bool {
	if o == nil || IsNil(o.BuiltInDnsClientEnabled) {
		var ret bool
		return ret
	}
	return *o.BuiltInDnsClientEnabled
}

// GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetBuiltInDnsClientEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BuiltInDnsClientEnabled) {
		return nil, false
	}
	return o.BuiltInDnsClientEnabled, true
}

// HasBuiltInDnsClientEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasBuiltInDnsClientEnabled() bool {
	if o != nil && !IsNil(o.BuiltInDnsClientEnabled) {
		return true
	}

	return false
}

// SetBuiltInDnsClientEnabled gets a reference to the given bool and assigns it to the BuiltInDnsClientEnabled field.
func (o *DTCWindows) SetBuiltInDnsClientEnabled(v bool) {
	o.BuiltInDnsClientEnabled = &v
}

// GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field value if set, zero value otherwise.
func (o *DTCWindows) GetChromeRemoteDesktopAppBlocked() bool {
	if o == nil || IsNil(o.ChromeRemoteDesktopAppBlocked) {
		var ret bool
		return ret
	}
	return *o.ChromeRemoteDesktopAppBlocked
}

// GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.ChromeRemoteDesktopAppBlocked) {
		return nil, false
	}
	return o.ChromeRemoteDesktopAppBlocked, true
}

// HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.
func (o *DTCWindows) HasChromeRemoteDesktopAppBlocked() bool {
	if o != nil && !IsNil(o.ChromeRemoteDesktopAppBlocked) {
		return true
	}

	return false
}

// SetChromeRemoteDesktopAppBlocked gets a reference to the given bool and assigns it to the ChromeRemoteDesktopAppBlocked field.
func (o *DTCWindows) SetChromeRemoteDesktopAppBlocked(v bool) {
	o.ChromeRemoteDesktopAppBlocked = &v
}

// GetCrowdStrikeAgentId returns the CrowdStrikeAgentId field value if set, zero value otherwise.
func (o *DTCWindows) GetCrowdStrikeAgentId() string {
	if o == nil || IsNil(o.CrowdStrikeAgentId) {
		var ret string
		return ret
	}
	return *o.CrowdStrikeAgentId
}

// GetCrowdStrikeAgentIdOk returns a tuple with the CrowdStrikeAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetCrowdStrikeAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.CrowdStrikeAgentId) {
		return nil, false
	}
	return o.CrowdStrikeAgentId, true
}

// HasCrowdStrikeAgentId returns a boolean if a field has been set.
func (o *DTCWindows) HasCrowdStrikeAgentId() bool {
	if o != nil && !IsNil(o.CrowdStrikeAgentId) {
		return true
	}

	return false
}

// SetCrowdStrikeAgentId gets a reference to the given string and assigns it to the CrowdStrikeAgentId field.
func (o *DTCWindows) SetCrowdStrikeAgentId(v string) {
	o.CrowdStrikeAgentId = &v
}

// GetCrowdStrikeCustomerId returns the CrowdStrikeCustomerId field value if set, zero value otherwise.
func (o *DTCWindows) GetCrowdStrikeCustomerId() string {
	if o == nil || IsNil(o.CrowdStrikeCustomerId) {
		var ret string
		return ret
	}
	return *o.CrowdStrikeCustomerId
}

// GetCrowdStrikeCustomerIdOk returns a tuple with the CrowdStrikeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetCrowdStrikeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CrowdStrikeCustomerId) {
		return nil, false
	}
	return o.CrowdStrikeCustomerId, true
}

// HasCrowdStrikeCustomerId returns a boolean if a field has been set.
func (o *DTCWindows) HasCrowdStrikeCustomerId() bool {
	if o != nil && !IsNil(o.CrowdStrikeCustomerId) {
		return true
	}

	return false
}

// SetCrowdStrikeCustomerId gets a reference to the given string and assigns it to the CrowdStrikeCustomerId field.
func (o *DTCWindows) SetCrowdStrikeCustomerId(v string) {
	o.CrowdStrikeCustomerId = &v
}

// GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field value if set, zero value otherwise.
func (o *DTCWindows) GetDeviceEnrollmentDomain() string {
	if o == nil || IsNil(o.DeviceEnrollmentDomain) {
		var ret string
		return ret
	}
	return *o.DeviceEnrollmentDomain
}

// GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetDeviceEnrollmentDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceEnrollmentDomain) {
		return nil, false
	}
	return o.DeviceEnrollmentDomain, true
}

// HasDeviceEnrollmentDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasDeviceEnrollmentDomain() bool {
	if o != nil && !IsNil(o.DeviceEnrollmentDomain) {
		return true
	}

	return false
}

// SetDeviceEnrollmentDomain gets a reference to the given string and assigns it to the DeviceEnrollmentDomain field.
func (o *DTCWindows) SetDeviceEnrollmentDomain(v string) {
	o.DeviceEnrollmentDomain = &v
}

// GetDiskEncrypted returns the DiskEncrypted field value if set, zero value otherwise.
func (o *DTCWindows) GetDiskEncrypted() bool {
	if o == nil || IsNil(o.DiskEncrypted) {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.DiskEncrypted) {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *DTCWindows) HasDiskEncrypted() bool {
	if o != nil && !IsNil(o.DiskEncrypted) {
		return true
	}

	return false
}

// SetDiskEncrypted gets a reference to the given bool and assigns it to the DiskEncrypted field.
func (o *DTCWindows) SetDiskEncrypted(v bool) {
	o.DiskEncrypted = &v
}

// GetKeyTrustLevel returns the KeyTrustLevel field value if set, zero value otherwise.
func (o *DTCWindows) GetKeyTrustLevel() string {
	if o == nil || IsNil(o.KeyTrustLevel) {
		var ret string
		return ret
	}
	return *o.KeyTrustLevel
}

// GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetKeyTrustLevelOk() (*string, bool) {
	if o == nil || IsNil(o.KeyTrustLevel) {
		return nil, false
	}
	return o.KeyTrustLevel, true
}

// HasKeyTrustLevel returns a boolean if a field has been set.
func (o *DTCWindows) HasKeyTrustLevel() bool {
	if o != nil && !IsNil(o.KeyTrustLevel) {
		return true
	}

	return false
}

// SetKeyTrustLevel gets a reference to the given string and assigns it to the KeyTrustLevel field.
func (o *DTCWindows) SetKeyTrustLevel(v string) {
	o.KeyTrustLevel = &v
}

// GetOsFirewall returns the OsFirewall field value if set, zero value otherwise.
func (o *DTCWindows) GetOsFirewall() bool {
	if o == nil || IsNil(o.OsFirewall) {
		var ret bool
		return ret
	}
	return *o.OsFirewall
}

// GetOsFirewallOk returns a tuple with the OsFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetOsFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.OsFirewall) {
		return nil, false
	}
	return o.OsFirewall, true
}

// HasOsFirewall returns a boolean if a field has been set.
func (o *DTCWindows) HasOsFirewall() bool {
	if o != nil && !IsNil(o.OsFirewall) {
		return true
	}

	return false
}

// SetOsFirewall gets a reference to the given bool and assigns it to the OsFirewall field.
func (o *DTCWindows) SetOsFirewall(v bool) {
	o.OsFirewall = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DTCWindows) GetOsVersion() OSVersionFourComponents {
	if o == nil || IsNil(o.OsVersion) {
		var ret OSVersionFourComponents
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetOsVersionOk() (*OSVersionFourComponents, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DTCWindows) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given OSVersionFourComponents and assigns it to the OsVersion field.
func (o *DTCWindows) SetOsVersion(v OSVersionFourComponents) {
	o.OsVersion = &v
}

// GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field value if set, zero value otherwise.
func (o *DTCWindows) GetPasswordProtectionWarningTrigger() string {
	if o == nil || IsNil(o.PasswordProtectionWarningTrigger) {
		var ret string
		return ret
	}
	return *o.PasswordProtectionWarningTrigger
}

// GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetPasswordProtectionWarningTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordProtectionWarningTrigger) {
		return nil, false
	}
	return o.PasswordProtectionWarningTrigger, true
}

// HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.
func (o *DTCWindows) HasPasswordProtectionWarningTrigger() bool {
	if o != nil && !IsNil(o.PasswordProtectionWarningTrigger) {
		return true
	}

	return false
}

// SetPasswordProtectionWarningTrigger gets a reference to the given string and assigns it to the PasswordProtectionWarningTrigger field.
func (o *DTCWindows) SetPasswordProtectionWarningTrigger(v string) {
	o.PasswordProtectionWarningTrigger = &v
}

// GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field value if set, zero value otherwise.
func (o *DTCWindows) GetRealtimeUrlCheckMode() bool {
	if o == nil || IsNil(o.RealtimeUrlCheckMode) {
		var ret bool
		return ret
	}
	return *o.RealtimeUrlCheckMode
}

// GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetRealtimeUrlCheckModeOk() (*bool, bool) {
	if o == nil || IsNil(o.RealtimeUrlCheckMode) {
		return nil, false
	}
	return o.RealtimeUrlCheckMode, true
}

// HasRealtimeUrlCheckMode returns a boolean if a field has been set.
func (o *DTCWindows) HasRealtimeUrlCheckMode() bool {
	if o != nil && !IsNil(o.RealtimeUrlCheckMode) {
		return true
	}

	return false
}

// SetRealtimeUrlCheckMode gets a reference to the given bool and assigns it to the RealtimeUrlCheckMode field.
func (o *DTCWindows) SetRealtimeUrlCheckMode(v bool) {
	o.RealtimeUrlCheckMode = &v
}

// GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field value if set, zero value otherwise.
func (o *DTCWindows) GetSafeBrowsingProtectionLevel() string {
	if o == nil || IsNil(o.SafeBrowsingProtectionLevel) {
		var ret string
		return ret
	}
	return *o.SafeBrowsingProtectionLevel
}

// GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSafeBrowsingProtectionLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SafeBrowsingProtectionLevel) {
		return nil, false
	}
	return o.SafeBrowsingProtectionLevel, true
}

// HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.
func (o *DTCWindows) HasSafeBrowsingProtectionLevel() bool {
	if o != nil && !IsNil(o.SafeBrowsingProtectionLevel) {
		return true
	}

	return false
}

// SetSafeBrowsingProtectionLevel gets a reference to the given string and assigns it to the SafeBrowsingProtectionLevel field.
func (o *DTCWindows) SetSafeBrowsingProtectionLevel(v string) {
	o.SafeBrowsingProtectionLevel = &v
}

// GetScreenLockSecured returns the ScreenLockSecured field value if set, zero value otherwise.
func (o *DTCWindows) GetScreenLockSecured() bool {
	if o == nil || IsNil(o.ScreenLockSecured) {
		var ret bool
		return ret
	}
	return *o.ScreenLockSecured
}

// GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetScreenLockSecuredOk() (*bool, bool) {
	if o == nil || IsNil(o.ScreenLockSecured) {
		return nil, false
	}
	return o.ScreenLockSecured, true
}

// HasScreenLockSecured returns a boolean if a field has been set.
func (o *DTCWindows) HasScreenLockSecured() bool {
	if o != nil && !IsNil(o.ScreenLockSecured) {
		return true
	}

	return false
}

// SetScreenLockSecured gets a reference to the given bool and assigns it to the ScreenLockSecured field.
func (o *DTCWindows) SetScreenLockSecured(v bool) {
	o.ScreenLockSecured = &v
}

// GetSecureBootEnabled returns the SecureBootEnabled field value if set, zero value otherwise.
func (o *DTCWindows) GetSecureBootEnabled() bool {
	if o == nil || IsNil(o.SecureBootEnabled) {
		var ret bool
		return ret
	}
	return *o.SecureBootEnabled
}

// GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSecureBootEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureBootEnabled) {
		return nil, false
	}
	return o.SecureBootEnabled, true
}

// HasSecureBootEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasSecureBootEnabled() bool {
	if o != nil && !IsNil(o.SecureBootEnabled) {
		return true
	}

	return false
}

// SetSecureBootEnabled gets a reference to the given bool and assigns it to the SecureBootEnabled field.
func (o *DTCWindows) SetSecureBootEnabled(v bool) {
	o.SecureBootEnabled = &v
}

// GetSiteIsolationEnabled returns the SiteIsolationEnabled field value if set, zero value otherwise.
func (o *DTCWindows) GetSiteIsolationEnabled() bool {
	if o == nil || IsNil(o.SiteIsolationEnabled) {
		var ret bool
		return ret
	}
	return *o.SiteIsolationEnabled
}

// GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSiteIsolationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SiteIsolationEnabled) {
		return nil, false
	}
	return o.SiteIsolationEnabled, true
}

// HasSiteIsolationEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasSiteIsolationEnabled() bool {
	if o != nil && !IsNil(o.SiteIsolationEnabled) {
		return true
	}

	return false
}

// SetSiteIsolationEnabled gets a reference to the given bool and assigns it to the SiteIsolationEnabled field.
func (o *DTCWindows) SetSiteIsolationEnabled(v bool) {
	o.SiteIsolationEnabled = &v
}

// GetThirdPartyBlockingEnabled returns the ThirdPartyBlockingEnabled field value if set, zero value otherwise.
// Deprecated
func (o *DTCWindows) GetThirdPartyBlockingEnabled() bool {
	if o == nil || IsNil(o.ThirdPartyBlockingEnabled) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyBlockingEnabled
}

// GetThirdPartyBlockingEnabledOk returns a tuple with the ThirdPartyBlockingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DTCWindows) GetThirdPartyBlockingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdPartyBlockingEnabled) {
		return nil, false
	}
	return o.ThirdPartyBlockingEnabled, true
}

// HasThirdPartyBlockingEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasThirdPartyBlockingEnabled() bool {
	if o != nil && !IsNil(o.ThirdPartyBlockingEnabled) {
		return true
	}

	return false
}

// SetThirdPartyBlockingEnabled gets a reference to the given bool and assigns it to the ThirdPartyBlockingEnabled field.
// Deprecated
func (o *DTCWindows) SetThirdPartyBlockingEnabled(v bool) {
	o.ThirdPartyBlockingEnabled = &v
}

// GetWindowsMachineDomain returns the WindowsMachineDomain field value if set, zero value otherwise.
func (o *DTCWindows) GetWindowsMachineDomain() string {
	if o == nil || IsNil(o.WindowsMachineDomain) {
		var ret string
		return ret
	}
	return *o.WindowsMachineDomain
}

// GetWindowsMachineDomainOk returns a tuple with the WindowsMachineDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetWindowsMachineDomainOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsMachineDomain) {
		return nil, false
	}
	return o.WindowsMachineDomain, true
}

// HasWindowsMachineDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasWindowsMachineDomain() bool {
	if o != nil && !IsNil(o.WindowsMachineDomain) {
		return true
	}

	return false
}

// SetWindowsMachineDomain gets a reference to the given string and assigns it to the WindowsMachineDomain field.
func (o *DTCWindows) SetWindowsMachineDomain(v string) {
	o.WindowsMachineDomain = &v
}

// GetWindowsUserDomain returns the WindowsUserDomain field value if set, zero value otherwise.
func (o *DTCWindows) GetWindowsUserDomain() string {
	if o == nil || IsNil(o.WindowsUserDomain) {
		var ret string
		return ret
	}
	return *o.WindowsUserDomain
}

// GetWindowsUserDomainOk returns a tuple with the WindowsUserDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetWindowsUserDomainOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsUserDomain) {
		return nil, false
	}
	return o.WindowsUserDomain, true
}

// HasWindowsUserDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasWindowsUserDomain() bool {
	if o != nil && !IsNil(o.WindowsUserDomain) {
		return true
	}

	return false
}

// SetWindowsUserDomain gets a reference to the given string and assigns it to the WindowsUserDomain field.
func (o *DTCWindows) SetWindowsUserDomain(v string) {
	o.WindowsUserDomain = &v
}

func (o DTCWindows) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTCWindows) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CrowdStrikeAgentId) {
		toSerialize["crowdStrikeAgentId"] = o.CrowdStrikeAgentId
	}
	if !IsNil(o.CrowdStrikeCustomerId) {
		toSerialize["crowdStrikeCustomerId"] = o.CrowdStrikeCustomerId
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
	if !IsNil(o.SecureBootEnabled) {
		toSerialize["secureBootEnabled"] = o.SecureBootEnabled
	}
	if !IsNil(o.SiteIsolationEnabled) {
		toSerialize["siteIsolationEnabled"] = o.SiteIsolationEnabled
	}
	if !IsNil(o.ThirdPartyBlockingEnabled) {
		toSerialize["thirdPartyBlockingEnabled"] = o.ThirdPartyBlockingEnabled
	}
	if !IsNil(o.WindowsMachineDomain) {
		toSerialize["windowsMachineDomain"] = o.WindowsMachineDomain
	}
	if !IsNil(o.WindowsUserDomain) {
		toSerialize["windowsUserDomain"] = o.WindowsUserDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DTCWindows) UnmarshalJSON(data []byte) (err error) {
	varDTCWindows := _DTCWindows{}

	err = json.Unmarshal(data, &varDTCWindows)

	if err != nil {
		return err
	}

	*o = DTCWindows(varDTCWindows)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "browserVersion")
		delete(additionalProperties, "builtInDnsClientEnabled")
		delete(additionalProperties, "chromeRemoteDesktopAppBlocked")
		delete(additionalProperties, "crowdStrikeAgentId")
		delete(additionalProperties, "crowdStrikeCustomerId")
		delete(additionalProperties, "deviceEnrollmentDomain")
		delete(additionalProperties, "diskEncrypted")
		delete(additionalProperties, "keyTrustLevel")
		delete(additionalProperties, "osFirewall")
		delete(additionalProperties, "osVersion")
		delete(additionalProperties, "passwordProtectionWarningTrigger")
		delete(additionalProperties, "realtimeUrlCheckMode")
		delete(additionalProperties, "safeBrowsingProtectionLevel")
		delete(additionalProperties, "screenLockSecured")
		delete(additionalProperties, "secureBootEnabled")
		delete(additionalProperties, "siteIsolationEnabled")
		delete(additionalProperties, "thirdPartyBlockingEnabled")
		delete(additionalProperties, "windowsMachineDomain")
		delete(additionalProperties, "windowsUserDomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDTCWindows struct {
	value *DTCWindows
	isSet bool
}

func (v NullableDTCWindows) Get() *DTCWindows {
	return v.value
}

func (v *NullableDTCWindows) Set(val *DTCWindows) {
	v.value = val
	v.isSet = true
}

func (v NullableDTCWindows) IsSet() bool {
	return v.isSet
}

func (v *NullableDTCWindows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTCWindows(val *DTCWindows) *NullableDTCWindows {
	return &NullableDTCWindows{value: val, isSet: true}
}

func (v NullableDTCWindows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTCWindows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

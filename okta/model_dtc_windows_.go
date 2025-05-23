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
	// Indicates whether the device's startup software has its Secure Boot feature enabled
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
	// Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled
	SiteIsolationEnabled *bool `json:"siteIsolationEnabled,omitempty"`
	// Indicates whether Chrome is blocking third-party software injection
	ThirdPartyBlockingEnabled *bool `json:"thirdPartyBlockingEnabled,omitempty"`
	// Windows domain that the current machine has joined
	WindowsMachineDomain *string `json:"windowsMachineDomain,omitempty"`
	// Windows domain for the current OS user
	WindowsUserDomain *string `json:"windowsUserDomain,omitempty"`
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
	if o == nil || o.BrowserVersion == nil {
		var ret ChromeBrowserVersion
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetBrowserVersionOk() (*ChromeBrowserVersion, bool) {
	if o == nil || o.BrowserVersion == nil {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *DTCWindows) HasBrowserVersion() bool {
	if o != nil && o.BrowserVersion != nil {
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
	if o == nil || o.BuiltInDnsClientEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BuiltInDnsClientEnabled
}

// GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetBuiltInDnsClientEnabledOk() (*bool, bool) {
	if o == nil || o.BuiltInDnsClientEnabled == nil {
		return nil, false
	}
	return o.BuiltInDnsClientEnabled, true
}

// HasBuiltInDnsClientEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasBuiltInDnsClientEnabled() bool {
	if o != nil && o.BuiltInDnsClientEnabled != nil {
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
	if o == nil || o.ChromeRemoteDesktopAppBlocked == nil {
		var ret bool
		return ret
	}
	return *o.ChromeRemoteDesktopAppBlocked
}

// GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool) {
	if o == nil || o.ChromeRemoteDesktopAppBlocked == nil {
		return nil, false
	}
	return o.ChromeRemoteDesktopAppBlocked, true
}

// HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.
func (o *DTCWindows) HasChromeRemoteDesktopAppBlocked() bool {
	if o != nil && o.ChromeRemoteDesktopAppBlocked != nil {
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
	if o == nil || o.CrowdStrikeAgentId == nil {
		var ret string
		return ret
	}
	return *o.CrowdStrikeAgentId
}

// GetCrowdStrikeAgentIdOk returns a tuple with the CrowdStrikeAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetCrowdStrikeAgentIdOk() (*string, bool) {
	if o == nil || o.CrowdStrikeAgentId == nil {
		return nil, false
	}
	return o.CrowdStrikeAgentId, true
}

// HasCrowdStrikeAgentId returns a boolean if a field has been set.
func (o *DTCWindows) HasCrowdStrikeAgentId() bool {
	if o != nil && o.CrowdStrikeAgentId != nil {
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
	if o == nil || o.CrowdStrikeCustomerId == nil {
		var ret string
		return ret
	}
	return *o.CrowdStrikeCustomerId
}

// GetCrowdStrikeCustomerIdOk returns a tuple with the CrowdStrikeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetCrowdStrikeCustomerIdOk() (*string, bool) {
	if o == nil || o.CrowdStrikeCustomerId == nil {
		return nil, false
	}
	return o.CrowdStrikeCustomerId, true
}

// HasCrowdStrikeCustomerId returns a boolean if a field has been set.
func (o *DTCWindows) HasCrowdStrikeCustomerId() bool {
	if o != nil && o.CrowdStrikeCustomerId != nil {
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
	if o == nil || o.DeviceEnrollmentDomain == nil {
		var ret string
		return ret
	}
	return *o.DeviceEnrollmentDomain
}

// GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetDeviceEnrollmentDomainOk() (*string, bool) {
	if o == nil || o.DeviceEnrollmentDomain == nil {
		return nil, false
	}
	return o.DeviceEnrollmentDomain, true
}

// HasDeviceEnrollmentDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasDeviceEnrollmentDomain() bool {
	if o != nil && o.DeviceEnrollmentDomain != nil {
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
	if o == nil || o.DiskEncrypted == nil {
		var ret bool
		return ret
	}
	return *o.DiskEncrypted
}

// GetDiskEncryptedOk returns a tuple with the DiskEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetDiskEncryptedOk() (*bool, bool) {
	if o == nil || o.DiskEncrypted == nil {
		return nil, false
	}
	return o.DiskEncrypted, true
}

// HasDiskEncrypted returns a boolean if a field has been set.
func (o *DTCWindows) HasDiskEncrypted() bool {
	if o != nil && o.DiskEncrypted != nil {
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
	if o == nil || o.KeyTrustLevel == nil {
		var ret string
		return ret
	}
	return *o.KeyTrustLevel
}

// GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetKeyTrustLevelOk() (*string, bool) {
	if o == nil || o.KeyTrustLevel == nil {
		return nil, false
	}
	return o.KeyTrustLevel, true
}

// HasKeyTrustLevel returns a boolean if a field has been set.
func (o *DTCWindows) HasKeyTrustLevel() bool {
	if o != nil && o.KeyTrustLevel != nil {
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
	if o == nil || o.OsFirewall == nil {
		var ret bool
		return ret
	}
	return *o.OsFirewall
}

// GetOsFirewallOk returns a tuple with the OsFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetOsFirewallOk() (*bool, bool) {
	if o == nil || o.OsFirewall == nil {
		return nil, false
	}
	return o.OsFirewall, true
}

// HasOsFirewall returns a boolean if a field has been set.
func (o *DTCWindows) HasOsFirewall() bool {
	if o != nil && o.OsFirewall != nil {
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
	if o == nil || o.OsVersion == nil {
		var ret OSVersionFourComponents
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetOsVersionOk() (*OSVersionFourComponents, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DTCWindows) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
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
	if o == nil || o.PasswordProtectionWarningTrigger == nil {
		var ret string
		return ret
	}
	return *o.PasswordProtectionWarningTrigger
}

// GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetPasswordProtectionWarningTriggerOk() (*string, bool) {
	if o == nil || o.PasswordProtectionWarningTrigger == nil {
		return nil, false
	}
	return o.PasswordProtectionWarningTrigger, true
}

// HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.
func (o *DTCWindows) HasPasswordProtectionWarningTrigger() bool {
	if o != nil && o.PasswordProtectionWarningTrigger != nil {
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
	if o == nil || o.RealtimeUrlCheckMode == nil {
		var ret bool
		return ret
	}
	return *o.RealtimeUrlCheckMode
}

// GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetRealtimeUrlCheckModeOk() (*bool, bool) {
	if o == nil || o.RealtimeUrlCheckMode == nil {
		return nil, false
	}
	return o.RealtimeUrlCheckMode, true
}

// HasRealtimeUrlCheckMode returns a boolean if a field has been set.
func (o *DTCWindows) HasRealtimeUrlCheckMode() bool {
	if o != nil && o.RealtimeUrlCheckMode != nil {
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
	if o == nil || o.SafeBrowsingProtectionLevel == nil {
		var ret string
		return ret
	}
	return *o.SafeBrowsingProtectionLevel
}

// GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSafeBrowsingProtectionLevelOk() (*string, bool) {
	if o == nil || o.SafeBrowsingProtectionLevel == nil {
		return nil, false
	}
	return o.SafeBrowsingProtectionLevel, true
}

// HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.
func (o *DTCWindows) HasSafeBrowsingProtectionLevel() bool {
	if o != nil && o.SafeBrowsingProtectionLevel != nil {
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
	if o == nil || o.ScreenLockSecured == nil {
		var ret bool
		return ret
	}
	return *o.ScreenLockSecured
}

// GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetScreenLockSecuredOk() (*bool, bool) {
	if o == nil || o.ScreenLockSecured == nil {
		return nil, false
	}
	return o.ScreenLockSecured, true
}

// HasScreenLockSecured returns a boolean if a field has been set.
func (o *DTCWindows) HasScreenLockSecured() bool {
	if o != nil && o.ScreenLockSecured != nil {
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
	if o == nil || o.SecureBootEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SecureBootEnabled
}

// GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSecureBootEnabledOk() (*bool, bool) {
	if o == nil || o.SecureBootEnabled == nil {
		return nil, false
	}
	return o.SecureBootEnabled, true
}

// HasSecureBootEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasSecureBootEnabled() bool {
	if o != nil && o.SecureBootEnabled != nil {
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
	if o == nil || o.SiteIsolationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SiteIsolationEnabled
}

// GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetSiteIsolationEnabledOk() (*bool, bool) {
	if o == nil || o.SiteIsolationEnabled == nil {
		return nil, false
	}
	return o.SiteIsolationEnabled, true
}

// HasSiteIsolationEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasSiteIsolationEnabled() bool {
	if o != nil && o.SiteIsolationEnabled != nil {
		return true
	}

	return false
}

// SetSiteIsolationEnabled gets a reference to the given bool and assigns it to the SiteIsolationEnabled field.
func (o *DTCWindows) SetSiteIsolationEnabled(v bool) {
	o.SiteIsolationEnabled = &v
}

// GetThirdPartyBlockingEnabled returns the ThirdPartyBlockingEnabled field value if set, zero value otherwise.
func (o *DTCWindows) GetThirdPartyBlockingEnabled() bool {
	if o == nil || o.ThirdPartyBlockingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ThirdPartyBlockingEnabled
}

// GetThirdPartyBlockingEnabledOk returns a tuple with the ThirdPartyBlockingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetThirdPartyBlockingEnabledOk() (*bool, bool) {
	if o == nil || o.ThirdPartyBlockingEnabled == nil {
		return nil, false
	}
	return o.ThirdPartyBlockingEnabled, true
}

// HasThirdPartyBlockingEnabled returns a boolean if a field has been set.
func (o *DTCWindows) HasThirdPartyBlockingEnabled() bool {
	if o != nil && o.ThirdPartyBlockingEnabled != nil {
		return true
	}

	return false
}

// SetThirdPartyBlockingEnabled gets a reference to the given bool and assigns it to the ThirdPartyBlockingEnabled field.
func (o *DTCWindows) SetThirdPartyBlockingEnabled(v bool) {
	o.ThirdPartyBlockingEnabled = &v
}

// GetWindowsMachineDomain returns the WindowsMachineDomain field value if set, zero value otherwise.
func (o *DTCWindows) GetWindowsMachineDomain() string {
	if o == nil || o.WindowsMachineDomain == nil {
		var ret string
		return ret
	}
	return *o.WindowsMachineDomain
}

// GetWindowsMachineDomainOk returns a tuple with the WindowsMachineDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetWindowsMachineDomainOk() (*string, bool) {
	if o == nil || o.WindowsMachineDomain == nil {
		return nil, false
	}
	return o.WindowsMachineDomain, true
}

// HasWindowsMachineDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasWindowsMachineDomain() bool {
	if o != nil && o.WindowsMachineDomain != nil {
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
	if o == nil || o.WindowsUserDomain == nil {
		var ret string
		return ret
	}
	return *o.WindowsUserDomain
}

// GetWindowsUserDomainOk returns a tuple with the WindowsUserDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCWindows) GetWindowsUserDomainOk() (*string, bool) {
	if o == nil || o.WindowsUserDomain == nil {
		return nil, false
	}
	return o.WindowsUserDomain, true
}

// HasWindowsUserDomain returns a boolean if a field has been set.
func (o *DTCWindows) HasWindowsUserDomain() bool {
	if o != nil && o.WindowsUserDomain != nil {
		return true
	}

	return false
}

// SetWindowsUserDomain gets a reference to the given string and assigns it to the WindowsUserDomain field.
func (o *DTCWindows) SetWindowsUserDomain(v string) {
	o.WindowsUserDomain = &v
}

func (o DTCWindows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BrowserVersion != nil {
		toSerialize["browserVersion"] = o.BrowserVersion
	}
	if o.BuiltInDnsClientEnabled != nil {
		toSerialize["builtInDnsClientEnabled"] = o.BuiltInDnsClientEnabled
	}
	if o.ChromeRemoteDesktopAppBlocked != nil {
		toSerialize["chromeRemoteDesktopAppBlocked"] = o.ChromeRemoteDesktopAppBlocked
	}
	if o.CrowdStrikeAgentId != nil {
		toSerialize["crowdStrikeAgentId"] = o.CrowdStrikeAgentId
	}
	if o.CrowdStrikeCustomerId != nil {
		toSerialize["crowdStrikeCustomerId"] = o.CrowdStrikeCustomerId
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
	if o.SecureBootEnabled != nil {
		toSerialize["secureBootEnabled"] = o.SecureBootEnabled
	}
	if o.SiteIsolationEnabled != nil {
		toSerialize["siteIsolationEnabled"] = o.SiteIsolationEnabled
	}
	if o.ThirdPartyBlockingEnabled != nil {
		toSerialize["thirdPartyBlockingEnabled"] = o.ThirdPartyBlockingEnabled
	}
	if o.WindowsMachineDomain != nil {
		toSerialize["windowsMachineDomain"] = o.WindowsMachineDomain
	}
	if o.WindowsUserDomain != nil {
		toSerialize["windowsUserDomain"] = o.WindowsUserDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DTCWindows) UnmarshalJSON(bytes []byte) (err error) {
	varDTCWindows := _DTCWindows{}

	err = json.Unmarshal(bytes, &varDTCWindows)
	if err == nil {
		*o = DTCWindows(varDTCWindows)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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
	} else {
		return err
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


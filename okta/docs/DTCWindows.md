# DTCWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserVersion** | Pointer to [**ChromeBrowserVersion**](ChromeBrowserVersion.md) |  | [optional] 
**BuiltInDnsClientEnabled** | Pointer to **bool** | Indicates if a software stack is used to communicate with the DNS server | [optional] 
**ChromeRemoteDesktopAppBlocked** | Pointer to **bool** | Indicates whether access to the Chrome Remote Desktop application is blocked through a policy | [optional] 
**CrowdStrikeAgentId** | Pointer to **string** | Agent ID of an installed CrowdStrike agent | [optional] 
**CrowdStrikeCustomerId** | Pointer to **string** | Customer ID of an installed CrowdStrike agent | [optional] 
**DeviceEnrollmentDomain** | Pointer to **string** | Enrollment domain of the customer that is currently managing the device | [optional] 
**DiskEncrypted** | Pointer to **bool** | Indicates whether the main disk is encrypted | [optional] 
**KeyTrustLevel** | Pointer to **string** | Represents the attestation strength used by the Chrome Verified Access API | [optional] 
**OsFirewall** | Pointer to **bool** | Indicates whether a firewall is enabled at the OS-level on the device | [optional] 
**OsVersion** | Pointer to [**OSVersionFourComponents**](OSVersionFourComponents.md) |  | [optional] 
**PasswordProtectionWarningTrigger** | Pointer to **string** | Indicates whether the Password Protection Warning feature is enabled | [optional] 
**RealtimeUrlCheckMode** | Pointer to **bool** | Indicates whether enterprise-grade (custom) unsafe URL scanning is enabled | [optional] 
**SafeBrowsingProtectionLevel** | Pointer to **string** | Represents the current value of the Safe Browsing protection level | [optional] 
**ScreenLockSecured** | Pointer to **bool** | Indicates whether the device is password-protected | [optional] 
**SecureBootEnabled** | Pointer to **bool** | Indicates whether the device&#39;s startup software has its Secure Boot feature enabled | [optional] 
**SiteIsolationEnabled** | Pointer to **bool** | Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled | [optional] 
**ThirdPartyBlockingEnabled** | Pointer to **bool** | Indicates whether Chrome is blocking third-party software injection | [optional] 
**WindowsMachineDomain** | Pointer to **string** | Windows domain that the current machine has joined | [optional] 
**WindowsUserDomain** | Pointer to **string** | Windows domain for the current OS user | [optional] 

## Methods

### NewDTCWindows

`func NewDTCWindows() *DTCWindows`

NewDTCWindows instantiates a new DTCWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTCWindowsWithDefaults

`func NewDTCWindowsWithDefaults() *DTCWindows`

NewDTCWindowsWithDefaults instantiates a new DTCWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserVersion

`func (o *DTCWindows) GetBrowserVersion() ChromeBrowserVersion`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *DTCWindows) GetBrowserVersionOk() (*ChromeBrowserVersion, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *DTCWindows) SetBrowserVersion(v ChromeBrowserVersion)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *DTCWindows) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetBuiltInDnsClientEnabled

`func (o *DTCWindows) GetBuiltInDnsClientEnabled() bool`

GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field if non-nil, zero value otherwise.

### GetBuiltInDnsClientEnabledOk

`func (o *DTCWindows) GetBuiltInDnsClientEnabledOk() (*bool, bool)`

GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInDnsClientEnabled

`func (o *DTCWindows) SetBuiltInDnsClientEnabled(v bool)`

SetBuiltInDnsClientEnabled sets BuiltInDnsClientEnabled field to given value.

### HasBuiltInDnsClientEnabled

`func (o *DTCWindows) HasBuiltInDnsClientEnabled() bool`

HasBuiltInDnsClientEnabled returns a boolean if a field has been set.

### GetChromeRemoteDesktopAppBlocked

`func (o *DTCWindows) GetChromeRemoteDesktopAppBlocked() bool`

GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field if non-nil, zero value otherwise.

### GetChromeRemoteDesktopAppBlockedOk

`func (o *DTCWindows) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool)`

GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeRemoteDesktopAppBlocked

`func (o *DTCWindows) SetChromeRemoteDesktopAppBlocked(v bool)`

SetChromeRemoteDesktopAppBlocked sets ChromeRemoteDesktopAppBlocked field to given value.

### HasChromeRemoteDesktopAppBlocked

`func (o *DTCWindows) HasChromeRemoteDesktopAppBlocked() bool`

HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.

### GetCrowdStrikeAgentId

`func (o *DTCWindows) GetCrowdStrikeAgentId() string`

GetCrowdStrikeAgentId returns the CrowdStrikeAgentId field if non-nil, zero value otherwise.

### GetCrowdStrikeAgentIdOk

`func (o *DTCWindows) GetCrowdStrikeAgentIdOk() (*string, bool)`

GetCrowdStrikeAgentIdOk returns a tuple with the CrowdStrikeAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrowdStrikeAgentId

`func (o *DTCWindows) SetCrowdStrikeAgentId(v string)`

SetCrowdStrikeAgentId sets CrowdStrikeAgentId field to given value.

### HasCrowdStrikeAgentId

`func (o *DTCWindows) HasCrowdStrikeAgentId() bool`

HasCrowdStrikeAgentId returns a boolean if a field has been set.

### GetCrowdStrikeCustomerId

`func (o *DTCWindows) GetCrowdStrikeCustomerId() string`

GetCrowdStrikeCustomerId returns the CrowdStrikeCustomerId field if non-nil, zero value otherwise.

### GetCrowdStrikeCustomerIdOk

`func (o *DTCWindows) GetCrowdStrikeCustomerIdOk() (*string, bool)`

GetCrowdStrikeCustomerIdOk returns a tuple with the CrowdStrikeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrowdStrikeCustomerId

`func (o *DTCWindows) SetCrowdStrikeCustomerId(v string)`

SetCrowdStrikeCustomerId sets CrowdStrikeCustomerId field to given value.

### HasCrowdStrikeCustomerId

`func (o *DTCWindows) HasCrowdStrikeCustomerId() bool`

HasCrowdStrikeCustomerId returns a boolean if a field has been set.

### GetDeviceEnrollmentDomain

`func (o *DTCWindows) GetDeviceEnrollmentDomain() string`

GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field if non-nil, zero value otherwise.

### GetDeviceEnrollmentDomainOk

`func (o *DTCWindows) GetDeviceEnrollmentDomainOk() (*string, bool)`

GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentDomain

`func (o *DTCWindows) SetDeviceEnrollmentDomain(v string)`

SetDeviceEnrollmentDomain sets DeviceEnrollmentDomain field to given value.

### HasDeviceEnrollmentDomain

`func (o *DTCWindows) HasDeviceEnrollmentDomain() bool`

HasDeviceEnrollmentDomain returns a boolean if a field has been set.

### GetDiskEncrypted

`func (o *DTCWindows) GetDiskEncrypted() bool`

GetDiskEncrypted returns the DiskEncrypted field if non-nil, zero value otherwise.

### GetDiskEncryptedOk

`func (o *DTCWindows) GetDiskEncryptedOk() (*bool, bool)`

GetDiskEncryptedOk returns a tuple with the DiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncrypted

`func (o *DTCWindows) SetDiskEncrypted(v bool)`

SetDiskEncrypted sets DiskEncrypted field to given value.

### HasDiskEncrypted

`func (o *DTCWindows) HasDiskEncrypted() bool`

HasDiskEncrypted returns a boolean if a field has been set.

### GetKeyTrustLevel

`func (o *DTCWindows) GetKeyTrustLevel() string`

GetKeyTrustLevel returns the KeyTrustLevel field if non-nil, zero value otherwise.

### GetKeyTrustLevelOk

`func (o *DTCWindows) GetKeyTrustLevelOk() (*string, bool)`

GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTrustLevel

`func (o *DTCWindows) SetKeyTrustLevel(v string)`

SetKeyTrustLevel sets KeyTrustLevel field to given value.

### HasKeyTrustLevel

`func (o *DTCWindows) HasKeyTrustLevel() bool`

HasKeyTrustLevel returns a boolean if a field has been set.

### GetOsFirewall

`func (o *DTCWindows) GetOsFirewall() bool`

GetOsFirewall returns the OsFirewall field if non-nil, zero value otherwise.

### GetOsFirewallOk

`func (o *DTCWindows) GetOsFirewallOk() (*bool, bool)`

GetOsFirewallOk returns a tuple with the OsFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFirewall

`func (o *DTCWindows) SetOsFirewall(v bool)`

SetOsFirewall sets OsFirewall field to given value.

### HasOsFirewall

`func (o *DTCWindows) HasOsFirewall() bool`

HasOsFirewall returns a boolean if a field has been set.

### GetOsVersion

`func (o *DTCWindows) GetOsVersion() OSVersionFourComponents`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DTCWindows) GetOsVersionOk() (*OSVersionFourComponents, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DTCWindows) SetOsVersion(v OSVersionFourComponents)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DTCWindows) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPasswordProtectionWarningTrigger

`func (o *DTCWindows) GetPasswordProtectionWarningTrigger() string`

GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field if non-nil, zero value otherwise.

### GetPasswordProtectionWarningTriggerOk

`func (o *DTCWindows) GetPasswordProtectionWarningTriggerOk() (*string, bool)`

GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtectionWarningTrigger

`func (o *DTCWindows) SetPasswordProtectionWarningTrigger(v string)`

SetPasswordProtectionWarningTrigger sets PasswordProtectionWarningTrigger field to given value.

### HasPasswordProtectionWarningTrigger

`func (o *DTCWindows) HasPasswordProtectionWarningTrigger() bool`

HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.

### GetRealtimeUrlCheckMode

`func (o *DTCWindows) GetRealtimeUrlCheckMode() bool`

GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field if non-nil, zero value otherwise.

### GetRealtimeUrlCheckModeOk

`func (o *DTCWindows) GetRealtimeUrlCheckModeOk() (*bool, bool)`

GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeUrlCheckMode

`func (o *DTCWindows) SetRealtimeUrlCheckMode(v bool)`

SetRealtimeUrlCheckMode sets RealtimeUrlCheckMode field to given value.

### HasRealtimeUrlCheckMode

`func (o *DTCWindows) HasRealtimeUrlCheckMode() bool`

HasRealtimeUrlCheckMode returns a boolean if a field has been set.

### GetSafeBrowsingProtectionLevel

`func (o *DTCWindows) GetSafeBrowsingProtectionLevel() string`

GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field if non-nil, zero value otherwise.

### GetSafeBrowsingProtectionLevelOk

`func (o *DTCWindows) GetSafeBrowsingProtectionLevelOk() (*string, bool)`

GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeBrowsingProtectionLevel

`func (o *DTCWindows) SetSafeBrowsingProtectionLevel(v string)`

SetSafeBrowsingProtectionLevel sets SafeBrowsingProtectionLevel field to given value.

### HasSafeBrowsingProtectionLevel

`func (o *DTCWindows) HasSafeBrowsingProtectionLevel() bool`

HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.

### GetScreenLockSecured

`func (o *DTCWindows) GetScreenLockSecured() bool`

GetScreenLockSecured returns the ScreenLockSecured field if non-nil, zero value otherwise.

### GetScreenLockSecuredOk

`func (o *DTCWindows) GetScreenLockSecuredOk() (*bool, bool)`

GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockSecured

`func (o *DTCWindows) SetScreenLockSecured(v bool)`

SetScreenLockSecured sets ScreenLockSecured field to given value.

### HasScreenLockSecured

`func (o *DTCWindows) HasScreenLockSecured() bool`

HasScreenLockSecured returns a boolean if a field has been set.

### GetSecureBootEnabled

`func (o *DTCWindows) GetSecureBootEnabled() bool`

GetSecureBootEnabled returns the SecureBootEnabled field if non-nil, zero value otherwise.

### GetSecureBootEnabledOk

`func (o *DTCWindows) GetSecureBootEnabledOk() (*bool, bool)`

GetSecureBootEnabledOk returns a tuple with the SecureBootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBootEnabled

`func (o *DTCWindows) SetSecureBootEnabled(v bool)`

SetSecureBootEnabled sets SecureBootEnabled field to given value.

### HasSecureBootEnabled

`func (o *DTCWindows) HasSecureBootEnabled() bool`

HasSecureBootEnabled returns a boolean if a field has been set.

### GetSiteIsolationEnabled

`func (o *DTCWindows) GetSiteIsolationEnabled() bool`

GetSiteIsolationEnabled returns the SiteIsolationEnabled field if non-nil, zero value otherwise.

### GetSiteIsolationEnabledOk

`func (o *DTCWindows) GetSiteIsolationEnabledOk() (*bool, bool)`

GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIsolationEnabled

`func (o *DTCWindows) SetSiteIsolationEnabled(v bool)`

SetSiteIsolationEnabled sets SiteIsolationEnabled field to given value.

### HasSiteIsolationEnabled

`func (o *DTCWindows) HasSiteIsolationEnabled() bool`

HasSiteIsolationEnabled returns a boolean if a field has been set.

### GetThirdPartyBlockingEnabled

`func (o *DTCWindows) GetThirdPartyBlockingEnabled() bool`

GetThirdPartyBlockingEnabled returns the ThirdPartyBlockingEnabled field if non-nil, zero value otherwise.

### GetThirdPartyBlockingEnabledOk

`func (o *DTCWindows) GetThirdPartyBlockingEnabledOk() (*bool, bool)`

GetThirdPartyBlockingEnabledOk returns a tuple with the ThirdPartyBlockingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyBlockingEnabled

`func (o *DTCWindows) SetThirdPartyBlockingEnabled(v bool)`

SetThirdPartyBlockingEnabled sets ThirdPartyBlockingEnabled field to given value.

### HasThirdPartyBlockingEnabled

`func (o *DTCWindows) HasThirdPartyBlockingEnabled() bool`

HasThirdPartyBlockingEnabled returns a boolean if a field has been set.

### GetWindowsMachineDomain

`func (o *DTCWindows) GetWindowsMachineDomain() string`

GetWindowsMachineDomain returns the WindowsMachineDomain field if non-nil, zero value otherwise.

### GetWindowsMachineDomainOk

`func (o *DTCWindows) GetWindowsMachineDomainOk() (*string, bool)`

GetWindowsMachineDomainOk returns a tuple with the WindowsMachineDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsMachineDomain

`func (o *DTCWindows) SetWindowsMachineDomain(v string)`

SetWindowsMachineDomain sets WindowsMachineDomain field to given value.

### HasWindowsMachineDomain

`func (o *DTCWindows) HasWindowsMachineDomain() bool`

HasWindowsMachineDomain returns a boolean if a field has been set.

### GetWindowsUserDomain

`func (o *DTCWindows) GetWindowsUserDomain() string`

GetWindowsUserDomain returns the WindowsUserDomain field if non-nil, zero value otherwise.

### GetWindowsUserDomainOk

`func (o *DTCWindows) GetWindowsUserDomainOk() (*string, bool)`

GetWindowsUserDomainOk returns a tuple with the WindowsUserDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUserDomain

`func (o *DTCWindows) SetWindowsUserDomain(v string)`

SetWindowsUserDomain sets WindowsUserDomain field to given value.

### HasWindowsUserDomain

`func (o *DTCWindows) HasWindowsUserDomain() bool`

HasWindowsUserDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DTCMacOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserVersion** | Pointer to [**ChromeBrowserVersion**](ChromeBrowserVersion.md) |  | [optional] 
**BuiltInDnsClientEnabled** | Pointer to **bool** | Indicates if a software stack is used to communicate with the DNS server | [optional] 
**ChromeRemoteDesktopAppBlocked** | Pointer to **bool** | Indicates whether access to the Chrome Remote Desktop application is blocked through a policy | [optional] 
**DeviceEnrollmentDomain** | Pointer to **string** | Enrollment domain of the customer that is currently managing the device | [optional] 
**DiskEncrypted** | Pointer to **bool** | Indicates whether the main disk is encrypted | [optional] 
**KeyTrustLevel** | Pointer to **string** | Represents the attestation strength used by the Chrome Verified Access API | [optional] 
**OsFirewall** | Pointer to **bool** | Indicates whether a firewall is enabled at the OS-level on the device | [optional] 
**OsVersion** | Pointer to [**OSVersionThreeComponents**](OSVersionThreeComponents.md) |  | [optional] 
**PasswordProtectionWarningTrigger** | Pointer to **string** | Indicates whether the Password Protection Warning feature is enabled | [optional] 
**RealtimeUrlCheckMode** | Pointer to **bool** | Indicates whether enterprise-grade (custom) unsafe URL scanning is enabled | [optional] 
**SafeBrowsingProtectionLevel** | Pointer to **string** | Represents the current value of the Safe Browsing protection level | [optional] 
**ScreenLockSecured** | Pointer to **bool** | Indicates whether the device is password-protected | [optional] 
**SiteIsolationEnabled** | Pointer to **bool** | Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled | [optional] 

## Methods

### NewDTCMacOS

`func NewDTCMacOS() *DTCMacOS`

NewDTCMacOS instantiates a new DTCMacOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTCMacOSWithDefaults

`func NewDTCMacOSWithDefaults() *DTCMacOS`

NewDTCMacOSWithDefaults instantiates a new DTCMacOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowserVersion

`func (o *DTCMacOS) GetBrowserVersion() ChromeBrowserVersion`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *DTCMacOS) GetBrowserVersionOk() (*ChromeBrowserVersion, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *DTCMacOS) SetBrowserVersion(v ChromeBrowserVersion)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *DTCMacOS) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetBuiltInDnsClientEnabled

`func (o *DTCMacOS) GetBuiltInDnsClientEnabled() bool`

GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field if non-nil, zero value otherwise.

### GetBuiltInDnsClientEnabledOk

`func (o *DTCMacOS) GetBuiltInDnsClientEnabledOk() (*bool, bool)`

GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInDnsClientEnabled

`func (o *DTCMacOS) SetBuiltInDnsClientEnabled(v bool)`

SetBuiltInDnsClientEnabled sets BuiltInDnsClientEnabled field to given value.

### HasBuiltInDnsClientEnabled

`func (o *DTCMacOS) HasBuiltInDnsClientEnabled() bool`

HasBuiltInDnsClientEnabled returns a boolean if a field has been set.

### GetChromeRemoteDesktopAppBlocked

`func (o *DTCMacOS) GetChromeRemoteDesktopAppBlocked() bool`

GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field if non-nil, zero value otherwise.

### GetChromeRemoteDesktopAppBlockedOk

`func (o *DTCMacOS) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool)`

GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeRemoteDesktopAppBlocked

`func (o *DTCMacOS) SetChromeRemoteDesktopAppBlocked(v bool)`

SetChromeRemoteDesktopAppBlocked sets ChromeRemoteDesktopAppBlocked field to given value.

### HasChromeRemoteDesktopAppBlocked

`func (o *DTCMacOS) HasChromeRemoteDesktopAppBlocked() bool`

HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.

### GetDeviceEnrollmentDomain

`func (o *DTCMacOS) GetDeviceEnrollmentDomain() string`

GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field if non-nil, zero value otherwise.

### GetDeviceEnrollmentDomainOk

`func (o *DTCMacOS) GetDeviceEnrollmentDomainOk() (*string, bool)`

GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentDomain

`func (o *DTCMacOS) SetDeviceEnrollmentDomain(v string)`

SetDeviceEnrollmentDomain sets DeviceEnrollmentDomain field to given value.

### HasDeviceEnrollmentDomain

`func (o *DTCMacOS) HasDeviceEnrollmentDomain() bool`

HasDeviceEnrollmentDomain returns a boolean if a field has been set.

### GetDiskEncrypted

`func (o *DTCMacOS) GetDiskEncrypted() bool`

GetDiskEncrypted returns the DiskEncrypted field if non-nil, zero value otherwise.

### GetDiskEncryptedOk

`func (o *DTCMacOS) GetDiskEncryptedOk() (*bool, bool)`

GetDiskEncryptedOk returns a tuple with the DiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncrypted

`func (o *DTCMacOS) SetDiskEncrypted(v bool)`

SetDiskEncrypted sets DiskEncrypted field to given value.

### HasDiskEncrypted

`func (o *DTCMacOS) HasDiskEncrypted() bool`

HasDiskEncrypted returns a boolean if a field has been set.

### GetKeyTrustLevel

`func (o *DTCMacOS) GetKeyTrustLevel() string`

GetKeyTrustLevel returns the KeyTrustLevel field if non-nil, zero value otherwise.

### GetKeyTrustLevelOk

`func (o *DTCMacOS) GetKeyTrustLevelOk() (*string, bool)`

GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTrustLevel

`func (o *DTCMacOS) SetKeyTrustLevel(v string)`

SetKeyTrustLevel sets KeyTrustLevel field to given value.

### HasKeyTrustLevel

`func (o *DTCMacOS) HasKeyTrustLevel() bool`

HasKeyTrustLevel returns a boolean if a field has been set.

### GetOsFirewall

`func (o *DTCMacOS) GetOsFirewall() bool`

GetOsFirewall returns the OsFirewall field if non-nil, zero value otherwise.

### GetOsFirewallOk

`func (o *DTCMacOS) GetOsFirewallOk() (*bool, bool)`

GetOsFirewallOk returns a tuple with the OsFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFirewall

`func (o *DTCMacOS) SetOsFirewall(v bool)`

SetOsFirewall sets OsFirewall field to given value.

### HasOsFirewall

`func (o *DTCMacOS) HasOsFirewall() bool`

HasOsFirewall returns a boolean if a field has been set.

### GetOsVersion

`func (o *DTCMacOS) GetOsVersion() OSVersionThreeComponents`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DTCMacOS) GetOsVersionOk() (*OSVersionThreeComponents, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DTCMacOS) SetOsVersion(v OSVersionThreeComponents)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DTCMacOS) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPasswordProtectionWarningTrigger

`func (o *DTCMacOS) GetPasswordProtectionWarningTrigger() string`

GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field if non-nil, zero value otherwise.

### GetPasswordProtectionWarningTriggerOk

`func (o *DTCMacOS) GetPasswordProtectionWarningTriggerOk() (*string, bool)`

GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtectionWarningTrigger

`func (o *DTCMacOS) SetPasswordProtectionWarningTrigger(v string)`

SetPasswordProtectionWarningTrigger sets PasswordProtectionWarningTrigger field to given value.

### HasPasswordProtectionWarningTrigger

`func (o *DTCMacOS) HasPasswordProtectionWarningTrigger() bool`

HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.

### GetRealtimeUrlCheckMode

`func (o *DTCMacOS) GetRealtimeUrlCheckMode() bool`

GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field if non-nil, zero value otherwise.

### GetRealtimeUrlCheckModeOk

`func (o *DTCMacOS) GetRealtimeUrlCheckModeOk() (*bool, bool)`

GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeUrlCheckMode

`func (o *DTCMacOS) SetRealtimeUrlCheckMode(v bool)`

SetRealtimeUrlCheckMode sets RealtimeUrlCheckMode field to given value.

### HasRealtimeUrlCheckMode

`func (o *DTCMacOS) HasRealtimeUrlCheckMode() bool`

HasRealtimeUrlCheckMode returns a boolean if a field has been set.

### GetSafeBrowsingProtectionLevel

`func (o *DTCMacOS) GetSafeBrowsingProtectionLevel() string`

GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field if non-nil, zero value otherwise.

### GetSafeBrowsingProtectionLevelOk

`func (o *DTCMacOS) GetSafeBrowsingProtectionLevelOk() (*string, bool)`

GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeBrowsingProtectionLevel

`func (o *DTCMacOS) SetSafeBrowsingProtectionLevel(v string)`

SetSafeBrowsingProtectionLevel sets SafeBrowsingProtectionLevel field to given value.

### HasSafeBrowsingProtectionLevel

`func (o *DTCMacOS) HasSafeBrowsingProtectionLevel() bool`

HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.

### GetScreenLockSecured

`func (o *DTCMacOS) GetScreenLockSecured() bool`

GetScreenLockSecured returns the ScreenLockSecured field if non-nil, zero value otherwise.

### GetScreenLockSecuredOk

`func (o *DTCMacOS) GetScreenLockSecuredOk() (*bool, bool)`

GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockSecured

`func (o *DTCMacOS) SetScreenLockSecured(v bool)`

SetScreenLockSecured sets ScreenLockSecured field to given value.

### HasScreenLockSecured

`func (o *DTCMacOS) HasScreenLockSecured() bool`

HasScreenLockSecured returns a boolean if a field has been set.

### GetSiteIsolationEnabled

`func (o *DTCMacOS) GetSiteIsolationEnabled() bool`

GetSiteIsolationEnabled returns the SiteIsolationEnabled field if non-nil, zero value otherwise.

### GetSiteIsolationEnabledOk

`func (o *DTCMacOS) GetSiteIsolationEnabledOk() (*bool, bool)`

GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIsolationEnabled

`func (o *DTCMacOS) SetSiteIsolationEnabled(v bool)`

SetSiteIsolationEnabled sets SiteIsolationEnabled field to given value.

### HasSiteIsolationEnabled

`func (o *DTCMacOS) HasSiteIsolationEnabled() bool`

HasSiteIsolationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



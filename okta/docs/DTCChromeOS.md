# DTCChromeOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowScreenLock** | Pointer to **bool** | Indicates whether the AllowScreenLock enterprise policy is enabled | [optional] 
**BrowserVersion** | Pointer to [**ChromeBrowserVersion**](ChromeBrowserVersion.md) |  | [optional] 
**BuiltInDnsClientEnabled** | Pointer to **bool** | Indicates if a software stack is used to communicate with the DNS server | [optional] 
**ChromeRemoteDesktopAppBlocked** | Pointer to **bool** | Indicates whether access to the Chrome Remote Desktop application is blocked through a policy | [optional] 
**DeviceEnrollmentDomain** | Pointer to **string** | Enrollment domain of the customer that is currently managing the device | [optional] 
**DiskEncrypted** | Pointer to **bool** | Indicates whether the main disk is encrypted | [optional] 
**KeyTrustLevel** | Pointer to **string** | Represents the attestation strength used by the Chrome Verified Access API | [optional] 
**ManagedDevice** | Pointer to **bool** | Indicates whether the device is enrolled in ChromeOS device management | [optional] 
**OsFirewall** | Pointer to **bool** | Indicates whether a firewall is enabled at the OS-level on the device | [optional] 
**OsVersion** | Pointer to [**OSVersionFourComponents**](OSVersionFourComponents.md) |  | [optional] 
**PasswordProtectionWarningTrigger** | Pointer to **string** | Indicates whether the Password Protection Warning feature is enabled | [optional] 
**RealtimeUrlCheckMode** | Pointer to **bool** | Indicates whether enterprise-grade (custom) unsafe URL scanning is enabled | [optional] 
**SafeBrowsingProtectionLevel** | Pointer to **string** | Represents the current value of the Safe Browsing protection level | [optional] 
**ScreenLockSecured** | Pointer to **bool** | Indicates whether the device is password-protected | [optional] 
**SiteIsolationEnabled** | Pointer to **bool** | Indicates whether the Site Isolation (also known as **Site Per Process**) setting is enabled | [optional] 

## Methods

### NewDTCChromeOS

`func NewDTCChromeOS() *DTCChromeOS`

NewDTCChromeOS instantiates a new DTCChromeOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTCChromeOSWithDefaults

`func NewDTCChromeOSWithDefaults() *DTCChromeOS`

NewDTCChromeOSWithDefaults instantiates a new DTCChromeOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowScreenLock

`func (o *DTCChromeOS) GetAllowScreenLock() bool`

GetAllowScreenLock returns the AllowScreenLock field if non-nil, zero value otherwise.

### GetAllowScreenLockOk

`func (o *DTCChromeOS) GetAllowScreenLockOk() (*bool, bool)`

GetAllowScreenLockOk returns a tuple with the AllowScreenLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowScreenLock

`func (o *DTCChromeOS) SetAllowScreenLock(v bool)`

SetAllowScreenLock sets AllowScreenLock field to given value.

### HasAllowScreenLock

`func (o *DTCChromeOS) HasAllowScreenLock() bool`

HasAllowScreenLock returns a boolean if a field has been set.

### GetBrowserVersion

`func (o *DTCChromeOS) GetBrowserVersion() ChromeBrowserVersion`

GetBrowserVersion returns the BrowserVersion field if non-nil, zero value otherwise.

### GetBrowserVersionOk

`func (o *DTCChromeOS) GetBrowserVersionOk() (*ChromeBrowserVersion, bool)`

GetBrowserVersionOk returns a tuple with the BrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserVersion

`func (o *DTCChromeOS) SetBrowserVersion(v ChromeBrowserVersion)`

SetBrowserVersion sets BrowserVersion field to given value.

### HasBrowserVersion

`func (o *DTCChromeOS) HasBrowserVersion() bool`

HasBrowserVersion returns a boolean if a field has been set.

### GetBuiltInDnsClientEnabled

`func (o *DTCChromeOS) GetBuiltInDnsClientEnabled() bool`

GetBuiltInDnsClientEnabled returns the BuiltInDnsClientEnabled field if non-nil, zero value otherwise.

### GetBuiltInDnsClientEnabledOk

`func (o *DTCChromeOS) GetBuiltInDnsClientEnabledOk() (*bool, bool)`

GetBuiltInDnsClientEnabledOk returns a tuple with the BuiltInDnsClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInDnsClientEnabled

`func (o *DTCChromeOS) SetBuiltInDnsClientEnabled(v bool)`

SetBuiltInDnsClientEnabled sets BuiltInDnsClientEnabled field to given value.

### HasBuiltInDnsClientEnabled

`func (o *DTCChromeOS) HasBuiltInDnsClientEnabled() bool`

HasBuiltInDnsClientEnabled returns a boolean if a field has been set.

### GetChromeRemoteDesktopAppBlocked

`func (o *DTCChromeOS) GetChromeRemoteDesktopAppBlocked() bool`

GetChromeRemoteDesktopAppBlocked returns the ChromeRemoteDesktopAppBlocked field if non-nil, zero value otherwise.

### GetChromeRemoteDesktopAppBlockedOk

`func (o *DTCChromeOS) GetChromeRemoteDesktopAppBlockedOk() (*bool, bool)`

GetChromeRemoteDesktopAppBlockedOk returns a tuple with the ChromeRemoteDesktopAppBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeRemoteDesktopAppBlocked

`func (o *DTCChromeOS) SetChromeRemoteDesktopAppBlocked(v bool)`

SetChromeRemoteDesktopAppBlocked sets ChromeRemoteDesktopAppBlocked field to given value.

### HasChromeRemoteDesktopAppBlocked

`func (o *DTCChromeOS) HasChromeRemoteDesktopAppBlocked() bool`

HasChromeRemoteDesktopAppBlocked returns a boolean if a field has been set.

### GetDeviceEnrollmentDomain

`func (o *DTCChromeOS) GetDeviceEnrollmentDomain() string`

GetDeviceEnrollmentDomain returns the DeviceEnrollmentDomain field if non-nil, zero value otherwise.

### GetDeviceEnrollmentDomainOk

`func (o *DTCChromeOS) GetDeviceEnrollmentDomainOk() (*string, bool)`

GetDeviceEnrollmentDomainOk returns a tuple with the DeviceEnrollmentDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentDomain

`func (o *DTCChromeOS) SetDeviceEnrollmentDomain(v string)`

SetDeviceEnrollmentDomain sets DeviceEnrollmentDomain field to given value.

### HasDeviceEnrollmentDomain

`func (o *DTCChromeOS) HasDeviceEnrollmentDomain() bool`

HasDeviceEnrollmentDomain returns a boolean if a field has been set.

### GetDiskEncrypted

`func (o *DTCChromeOS) GetDiskEncrypted() bool`

GetDiskEncrypted returns the DiskEncrypted field if non-nil, zero value otherwise.

### GetDiskEncryptedOk

`func (o *DTCChromeOS) GetDiskEncryptedOk() (*bool, bool)`

GetDiskEncryptedOk returns a tuple with the DiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncrypted

`func (o *DTCChromeOS) SetDiskEncrypted(v bool)`

SetDiskEncrypted sets DiskEncrypted field to given value.

### HasDiskEncrypted

`func (o *DTCChromeOS) HasDiskEncrypted() bool`

HasDiskEncrypted returns a boolean if a field has been set.

### GetKeyTrustLevel

`func (o *DTCChromeOS) GetKeyTrustLevel() string`

GetKeyTrustLevel returns the KeyTrustLevel field if non-nil, zero value otherwise.

### GetKeyTrustLevelOk

`func (o *DTCChromeOS) GetKeyTrustLevelOk() (*string, bool)`

GetKeyTrustLevelOk returns a tuple with the KeyTrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTrustLevel

`func (o *DTCChromeOS) SetKeyTrustLevel(v string)`

SetKeyTrustLevel sets KeyTrustLevel field to given value.

### HasKeyTrustLevel

`func (o *DTCChromeOS) HasKeyTrustLevel() bool`

HasKeyTrustLevel returns a boolean if a field has been set.

### GetManagedDevice

`func (o *DTCChromeOS) GetManagedDevice() bool`

GetManagedDevice returns the ManagedDevice field if non-nil, zero value otherwise.

### GetManagedDeviceOk

`func (o *DTCChromeOS) GetManagedDeviceOk() (*bool, bool)`

GetManagedDeviceOk returns a tuple with the ManagedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevice

`func (o *DTCChromeOS) SetManagedDevice(v bool)`

SetManagedDevice sets ManagedDevice field to given value.

### HasManagedDevice

`func (o *DTCChromeOS) HasManagedDevice() bool`

HasManagedDevice returns a boolean if a field has been set.

### GetOsFirewall

`func (o *DTCChromeOS) GetOsFirewall() bool`

GetOsFirewall returns the OsFirewall field if non-nil, zero value otherwise.

### GetOsFirewallOk

`func (o *DTCChromeOS) GetOsFirewallOk() (*bool, bool)`

GetOsFirewallOk returns a tuple with the OsFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFirewall

`func (o *DTCChromeOS) SetOsFirewall(v bool)`

SetOsFirewall sets OsFirewall field to given value.

### HasOsFirewall

`func (o *DTCChromeOS) HasOsFirewall() bool`

HasOsFirewall returns a boolean if a field has been set.

### GetOsVersion

`func (o *DTCChromeOS) GetOsVersion() OSVersionFourComponents`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DTCChromeOS) GetOsVersionOk() (*OSVersionFourComponents, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DTCChromeOS) SetOsVersion(v OSVersionFourComponents)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DTCChromeOS) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPasswordProtectionWarningTrigger

`func (o *DTCChromeOS) GetPasswordProtectionWarningTrigger() string`

GetPasswordProtectionWarningTrigger returns the PasswordProtectionWarningTrigger field if non-nil, zero value otherwise.

### GetPasswordProtectionWarningTriggerOk

`func (o *DTCChromeOS) GetPasswordProtectionWarningTriggerOk() (*string, bool)`

GetPasswordProtectionWarningTriggerOk returns a tuple with the PasswordProtectionWarningTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordProtectionWarningTrigger

`func (o *DTCChromeOS) SetPasswordProtectionWarningTrigger(v string)`

SetPasswordProtectionWarningTrigger sets PasswordProtectionWarningTrigger field to given value.

### HasPasswordProtectionWarningTrigger

`func (o *DTCChromeOS) HasPasswordProtectionWarningTrigger() bool`

HasPasswordProtectionWarningTrigger returns a boolean if a field has been set.

### GetRealtimeUrlCheckMode

`func (o *DTCChromeOS) GetRealtimeUrlCheckMode() bool`

GetRealtimeUrlCheckMode returns the RealtimeUrlCheckMode field if non-nil, zero value otherwise.

### GetRealtimeUrlCheckModeOk

`func (o *DTCChromeOS) GetRealtimeUrlCheckModeOk() (*bool, bool)`

GetRealtimeUrlCheckModeOk returns a tuple with the RealtimeUrlCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtimeUrlCheckMode

`func (o *DTCChromeOS) SetRealtimeUrlCheckMode(v bool)`

SetRealtimeUrlCheckMode sets RealtimeUrlCheckMode field to given value.

### HasRealtimeUrlCheckMode

`func (o *DTCChromeOS) HasRealtimeUrlCheckMode() bool`

HasRealtimeUrlCheckMode returns a boolean if a field has been set.

### GetSafeBrowsingProtectionLevel

`func (o *DTCChromeOS) GetSafeBrowsingProtectionLevel() string`

GetSafeBrowsingProtectionLevel returns the SafeBrowsingProtectionLevel field if non-nil, zero value otherwise.

### GetSafeBrowsingProtectionLevelOk

`func (o *DTCChromeOS) GetSafeBrowsingProtectionLevelOk() (*string, bool)`

GetSafeBrowsingProtectionLevelOk returns a tuple with the SafeBrowsingProtectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeBrowsingProtectionLevel

`func (o *DTCChromeOS) SetSafeBrowsingProtectionLevel(v string)`

SetSafeBrowsingProtectionLevel sets SafeBrowsingProtectionLevel field to given value.

### HasSafeBrowsingProtectionLevel

`func (o *DTCChromeOS) HasSafeBrowsingProtectionLevel() bool`

HasSafeBrowsingProtectionLevel returns a boolean if a field has been set.

### GetScreenLockSecured

`func (o *DTCChromeOS) GetScreenLockSecured() bool`

GetScreenLockSecured returns the ScreenLockSecured field if non-nil, zero value otherwise.

### GetScreenLockSecuredOk

`func (o *DTCChromeOS) GetScreenLockSecuredOk() (*bool, bool)`

GetScreenLockSecuredOk returns a tuple with the ScreenLockSecured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockSecured

`func (o *DTCChromeOS) SetScreenLockSecured(v bool)`

SetScreenLockSecured sets ScreenLockSecured field to given value.

### HasScreenLockSecured

`func (o *DTCChromeOS) HasScreenLockSecured() bool`

HasScreenLockSecured returns a boolean if a field has been set.

### GetSiteIsolationEnabled

`func (o *DTCChromeOS) GetSiteIsolationEnabled() bool`

GetSiteIsolationEnabled returns the SiteIsolationEnabled field if non-nil, zero value otherwise.

### GetSiteIsolationEnabledOk

`func (o *DTCChromeOS) GetSiteIsolationEnabledOk() (*bool, bool)`

GetSiteIsolationEnabledOk returns a tuple with the SiteIsolationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIsolationEnabled

`func (o *DTCChromeOS) SetSiteIsolationEnabled(v bool)`

SetSiteIsolationEnabled sets SiteIsolationEnabled field to given value.

### HasSiteIsolationEnabled

`func (o *DTCChromeOS) HasSiteIsolationEnabled() bool`

HasSiteIsolationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



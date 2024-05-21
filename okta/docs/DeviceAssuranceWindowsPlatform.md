# DeviceAssuranceWindowsPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskEncryptionType** | Pointer to [**DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType**](DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType.md) |  | [optional] 
**OsVersion** | Pointer to [**OSVersionFourComponents**](OSVersionFourComponents.md) |  | [optional] 
**OsVersionConstraints** | Pointer to [**[]OSVersionConstraint**](OSVersionConstraint.md) | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;Specifies the Windows version requirements for the assurance policy. Each requirement must correspond to a different major version (Windows 11 or Windows 10). If a requirement isn&#39;t specified for a major version, then devices on that major version satisfy the condition.  There are two types of OS requirements: * **Static**: A specific Windows version requirement that doesn&#39;t change until you update the policy. A static OS Windows requirement is specified with &#x60;majorVersionConstraint&#x60; and &#x60;minimum&#x60;. * **Dynamic**: A Windows version requirement that is relative to the latest major release and security patch. A dynamic OS Windows requirement is specified with &#x60;majorVersionConstraint&#x60; and &#x60;dynamicVersionRequirement&#x60;.  &gt; **Note:** Dynamic OS requirements are available only if the **Dynamic OS version compliance** [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled. The &#x60;osVersionConstraints&#x60; property is only supported for the Windows platform. You can&#39;t specify both &#x60;osVersion.minimum&#x60; and &#x60;osVersionConstraints&#x60; properties at the same time.  | [optional] 
**ScreenLockType** | Pointer to [**DeviceAssuranceAndroidPlatformAllOfScreenLockType**](DeviceAssuranceAndroidPlatformAllOfScreenLockType.md) |  | [optional] 
**SecureHardwarePresent** | Pointer to **bool** |  | [optional] 
**ThirdPartySignalProviders** | Pointer to [**DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders**](DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders.md) |  | [optional] 

## Methods

### NewDeviceAssuranceWindowsPlatform

`func NewDeviceAssuranceWindowsPlatform() *DeviceAssuranceWindowsPlatform`

NewDeviceAssuranceWindowsPlatform instantiates a new DeviceAssuranceWindowsPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAssuranceWindowsPlatformWithDefaults

`func NewDeviceAssuranceWindowsPlatformWithDefaults() *DeviceAssuranceWindowsPlatform`

NewDeviceAssuranceWindowsPlatformWithDefaults instantiates a new DeviceAssuranceWindowsPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionType() DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatform) SetDiskEncryptionType(v DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatform) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersion() OSVersionFourComponents`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersionOk() (*OSVersionFourComponents, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceAssuranceWindowsPlatform) SetOsVersion(v OSVersionFourComponents)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceAssuranceWindowsPlatform) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsVersionConstraints

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersionConstraints() []OSVersionConstraint`

GetOsVersionConstraints returns the OsVersionConstraints field if non-nil, zero value otherwise.

### GetOsVersionConstraintsOk

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersionConstraintsOk() (*[]OSVersionConstraint, bool)`

GetOsVersionConstraintsOk returns a tuple with the OsVersionConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersionConstraints

`func (o *DeviceAssuranceWindowsPlatform) SetOsVersionConstraints(v []OSVersionConstraint)`

SetOsVersionConstraints sets OsVersionConstraints field to given value.

### HasOsVersionConstraints

`func (o *DeviceAssuranceWindowsPlatform) HasOsVersionConstraints() bool`

HasOsVersionConstraints returns a boolean if a field has been set.

### GetScreenLockType

`func (o *DeviceAssuranceWindowsPlatform) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *DeviceAssuranceWindowsPlatform) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *DeviceAssuranceWindowsPlatform) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *DeviceAssuranceWindowsPlatform) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatform) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *DeviceAssuranceWindowsPlatform) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatform) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatform) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatform) GetThirdPartySignalProviders() DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders`

GetThirdPartySignalProviders returns the ThirdPartySignalProviders field if non-nil, zero value otherwise.

### GetThirdPartySignalProvidersOk

`func (o *DeviceAssuranceWindowsPlatform) GetThirdPartySignalProvidersOk() (*DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders, bool)`

GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatform) SetThirdPartySignalProviders(v DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders)`

SetThirdPartySignalProviders sets ThirdPartySignalProviders field to given value.

### HasThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatform) HasThirdPartySignalProviders() bool`

HasThirdPartySignalProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



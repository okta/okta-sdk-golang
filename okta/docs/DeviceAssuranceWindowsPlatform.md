# DeviceAssuranceWindowsPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskEncryptionType** | Pointer to [**DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType**](DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType.md) |  | [optional] 
**Jailbreak** | Pointer to **bool** |  | [optional] 
**OsVersion** | Pointer to [**OSVersion**](OSVersion.md) |  | [optional] 
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

`func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionType() DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *DeviceAssuranceWindowsPlatform) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatform) SetDiskEncryptionType(v DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatform) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetJailbreak

`func (o *DeviceAssuranceWindowsPlatform) GetJailbreak() bool`

GetJailbreak returns the Jailbreak field if non-nil, zero value otherwise.

### GetJailbreakOk

`func (o *DeviceAssuranceWindowsPlatform) GetJailbreakOk() (*bool, bool)`

GetJailbreakOk returns a tuple with the Jailbreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbreak

`func (o *DeviceAssuranceWindowsPlatform) SetJailbreak(v bool)`

SetJailbreak sets Jailbreak field to given value.

### HasJailbreak

`func (o *DeviceAssuranceWindowsPlatform) HasJailbreak() bool`

HasJailbreak returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersion() OSVersion`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceAssuranceWindowsPlatform) GetOsVersionOk() (*OSVersion, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceAssuranceWindowsPlatform) SetOsVersion(v OSVersion)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceAssuranceWindowsPlatform) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

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



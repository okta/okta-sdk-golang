# DeviceAssuranceWindowsPlatformAllOf

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

### NewDeviceAssuranceWindowsPlatformAllOf

`func NewDeviceAssuranceWindowsPlatformAllOf() *DeviceAssuranceWindowsPlatformAllOf`

NewDeviceAssuranceWindowsPlatformAllOf instantiates a new DeviceAssuranceWindowsPlatformAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAssuranceWindowsPlatformAllOfWithDefaults

`func NewDeviceAssuranceWindowsPlatformAllOfWithDefaults() *DeviceAssuranceWindowsPlatformAllOf`

NewDeviceAssuranceWindowsPlatformAllOfWithDefaults instantiates a new DeviceAssuranceWindowsPlatformAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetDiskEncryptionType() DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetDiskEncryptionType(v DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetJailbreak

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetJailbreak() bool`

GetJailbreak returns the Jailbreak field if non-nil, zero value otherwise.

### GetJailbreakOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetJailbreakOk() (*bool, bool)`

GetJailbreakOk returns a tuple with the Jailbreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbreak

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetJailbreak(v bool)`

SetJailbreak sets Jailbreak field to given value.

### HasJailbreak

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasJailbreak() bool`

HasJailbreak returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetOsVersion() OSVersion`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetOsVersionOk() (*OSVersion, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetOsVersion(v OSVersion)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetScreenLockType

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetThirdPartySignalProviders() DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders`

GetThirdPartySignalProviders returns the ThirdPartySignalProviders field if non-nil, zero value otherwise.

### GetThirdPartySignalProvidersOk

`func (o *DeviceAssuranceWindowsPlatformAllOf) GetThirdPartySignalProvidersOk() (*DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders, bool)`

GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatformAllOf) SetThirdPartySignalProviders(v DeviceAssuranceWindowsPlatformAllOfThirdPartySignalProviders)`

SetThirdPartySignalProviders sets ThirdPartySignalProviders field to given value.

### HasThirdPartySignalProviders

`func (o *DeviceAssuranceWindowsPlatformAllOf) HasThirdPartySignalProviders() bool`

HasThirdPartySignalProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



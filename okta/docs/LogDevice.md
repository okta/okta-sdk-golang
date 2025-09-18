# LogDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIntegrator** | Pointer to **map[string]interface{}** | The integration platform or software used with the device | [optional] [readonly] 
**DiskEncryptionType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID of the device | [optional] [readonly] 
**Jailbreak** | Pointer to **bool** | If the device has removed software restrictions | [optional] [readonly] 
**Managed** | Pointer to **bool** | Indicates if the device is configured for device management and is registered with Okta | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**OsPlatform** | Pointer to **string** |  | [optional] [readonly] 
**OsVersion** | Pointer to **string** |  | [optional] [readonly] 
**Registered** | Pointer to **bool** | Indicates if the device is registered with an Okta org and is bound to an Okta Verify instance on the device | [optional] [readonly] 
**ScreenLockType** | Pointer to **string** |  | [optional] 
**SecureHardwarePresent** | Pointer to **bool** | The availability of hardware security on the device | [optional] [readonly] 

## Methods

### NewLogDevice

`func NewLogDevice() *LogDevice`

NewLogDevice instantiates a new LogDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDeviceWithDefaults

`func NewLogDeviceWithDefaults() *LogDevice`

NewLogDeviceWithDefaults instantiates a new LogDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIntegrator

`func (o *LogDevice) GetDeviceIntegrator() map[string]interface{}`

GetDeviceIntegrator returns the DeviceIntegrator field if non-nil, zero value otherwise.

### GetDeviceIntegratorOk

`func (o *LogDevice) GetDeviceIntegratorOk() (*map[string]interface{}, bool)`

GetDeviceIntegratorOk returns a tuple with the DeviceIntegrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIntegrator

`func (o *LogDevice) SetDeviceIntegrator(v map[string]interface{})`

SetDeviceIntegrator sets DeviceIntegrator field to given value.

### HasDeviceIntegrator

`func (o *LogDevice) HasDeviceIntegrator() bool`

HasDeviceIntegrator returns a boolean if a field has been set.

### GetDiskEncryptionType

`func (o *LogDevice) GetDiskEncryptionType() string`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *LogDevice) GetDiskEncryptionTypeOk() (*string, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *LogDevice) SetDiskEncryptionType(v string)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *LogDevice) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetId

`func (o *LogDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJailbreak

`func (o *LogDevice) GetJailbreak() bool`

GetJailbreak returns the Jailbreak field if non-nil, zero value otherwise.

### GetJailbreakOk

`func (o *LogDevice) GetJailbreakOk() (*bool, bool)`

GetJailbreakOk returns a tuple with the Jailbreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbreak

`func (o *LogDevice) SetJailbreak(v bool)`

SetJailbreak sets Jailbreak field to given value.

### HasJailbreak

`func (o *LogDevice) HasJailbreak() bool`

HasJailbreak returns a boolean if a field has been set.

### GetManaged

`func (o *LogDevice) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LogDevice) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LogDevice) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *LogDevice) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetName

`func (o *LogDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsPlatform

`func (o *LogDevice) GetOsPlatform() string`

GetOsPlatform returns the OsPlatform field if non-nil, zero value otherwise.

### GetOsPlatformOk

`func (o *LogDevice) GetOsPlatformOk() (*string, bool)`

GetOsPlatformOk returns a tuple with the OsPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPlatform

`func (o *LogDevice) SetOsPlatform(v string)`

SetOsPlatform sets OsPlatform field to given value.

### HasOsPlatform

`func (o *LogDevice) HasOsPlatform() bool`

HasOsPlatform returns a boolean if a field has been set.

### GetOsVersion

`func (o *LogDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *LogDevice) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *LogDevice) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *LogDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetRegistered

`func (o *LogDevice) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *LogDevice) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *LogDevice) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *LogDevice) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetScreenLockType

`func (o *LogDevice) GetScreenLockType() string`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *LogDevice) GetScreenLockTypeOk() (*string, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *LogDevice) SetScreenLockType(v string)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *LogDevice) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *LogDevice) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *LogDevice) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *LogDevice) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *LogDevice) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



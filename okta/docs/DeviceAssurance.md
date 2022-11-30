# DeviceAssurance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedDate** | Pointer to **string** |  | [optional] [readonly] 
**DiskEncryptionType** | Pointer to [**DeviceAssuranceDiskEncryptionType**](DeviceAssuranceDiskEncryptionType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Jailbreak** | Pointer to **bool** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the Device Assurance Policy | [optional] 
**OsVersion** | Pointer to [**VersionObject**](VersionObject.md) |  | [optional] 
**Platform** | Pointer to [**Platform**](Platform.md) |  | [optional] 
**ScreenLockType** | Pointer to [**DeviceAssuranceScreenLockType**](DeviceAssuranceScreenLockType.md) |  | [optional] 
**SecureHardwarePresent** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 

## Methods

### NewDeviceAssurance

`func NewDeviceAssurance() *DeviceAssurance`

NewDeviceAssurance instantiates a new DeviceAssurance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAssuranceWithDefaults

`func NewDeviceAssuranceWithDefaults() *DeviceAssurance`

NewDeviceAssuranceWithDefaults instantiates a new DeviceAssurance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DeviceAssurance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceAssurance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceAssurance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceAssurance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DeviceAssurance) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeviceAssurance) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeviceAssurance) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeviceAssurance) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDiskEncryptionType

`func (o *DeviceAssurance) GetDiskEncryptionType() DeviceAssuranceDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *DeviceAssurance) GetDiskEncryptionTypeOk() (*DeviceAssuranceDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *DeviceAssurance) SetDiskEncryptionType(v DeviceAssuranceDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *DeviceAssurance) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetId

`func (o *DeviceAssurance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAssurance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAssurance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAssurance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJailbreak

`func (o *DeviceAssurance) GetJailbreak() bool`

GetJailbreak returns the Jailbreak field if non-nil, zero value otherwise.

### GetJailbreakOk

`func (o *DeviceAssurance) GetJailbreakOk() (*bool, bool)`

GetJailbreakOk returns a tuple with the Jailbreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbreak

`func (o *DeviceAssurance) SetJailbreak(v bool)`

SetJailbreak sets Jailbreak field to given value.

### HasJailbreak

`func (o *DeviceAssurance) HasJailbreak() bool`

HasJailbreak returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *DeviceAssurance) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *DeviceAssurance) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *DeviceAssurance) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *DeviceAssurance) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *DeviceAssurance) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *DeviceAssurance) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *DeviceAssurance) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *DeviceAssurance) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetName

`func (o *DeviceAssurance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAssurance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAssurance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAssurance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceAssurance) GetOsVersion() VersionObject`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceAssurance) GetOsVersionOk() (*VersionObject, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceAssurance) SetOsVersion(v VersionObject)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceAssurance) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceAssurance) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceAssurance) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceAssurance) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceAssurance) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetScreenLockType

`func (o *DeviceAssurance) GetScreenLockType() DeviceAssuranceScreenLockType`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *DeviceAssurance) GetScreenLockTypeOk() (*DeviceAssuranceScreenLockType, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *DeviceAssurance) SetScreenLockType(v DeviceAssuranceScreenLockType)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *DeviceAssurance) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *DeviceAssurance) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *DeviceAssurance) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *DeviceAssurance) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *DeviceAssurance) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetLinks

`func (o *DeviceAssurance) GetLinks() ApiTokenLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAssurance) GetLinksOk() (*ApiTokenLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAssurance) SetLinks(v ApiTokenLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAssurance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



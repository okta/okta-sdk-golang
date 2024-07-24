# ListDeviceAssurancePolicies200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedDate** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdate** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the Device Assurance Policy | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 
**DiskEncryptionType** | Pointer to [**DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType**](DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType.md) |  | [optional] 
**OsVersion** | Pointer to [**OSVersion**](OSVersion.md) |  | [optional] 
**ScreenLockType** | Pointer to [**DeviceAssuranceAndroidPlatformAllOfScreenLockType**](DeviceAssuranceAndroidPlatformAllOfScreenLockType.md) |  | [optional] 
**SecureHardwarePresent** | Pointer to **bool** |  | [optional] 
**ThirdPartySignalProviders** | Pointer to [**DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders**](DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders.md) |  | [optional] 
**Jailbreak** | Pointer to **bool** |  | [optional] 

## Methods

### NewListDeviceAssurancePolicies200ResponseInner

`func NewListDeviceAssurancePolicies200ResponseInner() *ListDeviceAssurancePolicies200ResponseInner`

NewListDeviceAssurancePolicies200ResponseInner instantiates a new ListDeviceAssurancePolicies200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeviceAssurancePolicies200ResponseInnerWithDefaults

`func NewListDeviceAssurancePolicies200ResponseInnerWithDefaults() *ListDeviceAssurancePolicies200ResponseInner`

NewListDeviceAssurancePolicies200ResponseInnerWithDefaults instantiates a new ListDeviceAssurancePolicies200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetId

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetName

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetLinks

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDiskEncryptionType

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetDiskEncryptionType() DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetDiskEncryptionTypeOk() (*DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetDiskEncryptionType(v DeviceAssuranceAndroidPlatformAllOfDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetOsVersion

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetOsVersion() OSVersion`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetOsVersionOk() (*OSVersion, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetOsVersion(v OSVersion)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetScreenLockType

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetThirdPartySignalProviders

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetThirdPartySignalProviders() DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders`

GetThirdPartySignalProviders returns the ThirdPartySignalProviders field if non-nil, zero value otherwise.

### GetThirdPartySignalProvidersOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetThirdPartySignalProvidersOk() (*DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders, bool)`

GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartySignalProviders

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetThirdPartySignalProviders(v DeviceAssuranceChromeOSPlatformAllOfThirdPartySignalProviders)`

SetThirdPartySignalProviders sets ThirdPartySignalProviders field to given value.

### HasThirdPartySignalProviders

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasThirdPartySignalProviders() bool`

HasThirdPartySignalProviders returns a boolean if a field has been set.

### GetJailbreak

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetJailbreak() bool`

GetJailbreak returns the Jailbreak field if non-nil, zero value otherwise.

### GetJailbreakOk

`func (o *ListDeviceAssurancePolicies200ResponseInner) GetJailbreakOk() (*bool, bool)`

GetJailbreakOk returns a tuple with the Jailbreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbreak

`func (o *ListDeviceAssurancePolicies200ResponseInner) SetJailbreak(v bool)`

SetJailbreak sets Jailbreak field to given value.

### HasJailbreak

`func (o *ListDeviceAssurancePolicies200ResponseInner) HasJailbreak() bool`

HasJailbreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



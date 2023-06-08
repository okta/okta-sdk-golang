# DeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the device | 
**Imei** | Pointer to **string** | International Mobile Equipment Identity of the device | [optional] 
**Manufacturer** | Pointer to **string** | Name of the manufacturer of the device | [optional] 
**Meid** | Pointer to **string** | Mobile equipment identifier of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**OsVersion** | Pointer to **string** | Version of the device OS | [optional] 
**Platform** | [**DevicePlatform**](DevicePlatform.md) |  | 
**Registered** | **bool** | Indicates if the device is registered at Okta | 
**SecureHardwarePresent** | Pointer to **bool** | Indicates if the device constains a secure hardware functionality | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the device | [optional] 
**Sid** | Pointer to **string** | Windows Security identifier of the device | [optional] 
**TpmPublicKeyHash** | Pointer to **string** | Windows Trsted Platform Module hash value | [optional] 
**Udid** | Pointer to **string** | macOS Unique Device identifier of the device | [optional] 

## Methods

### NewDeviceProfile

`func NewDeviceProfile(displayName string, platform DevicePlatform, registered bool, ) *DeviceProfile`

NewDeviceProfile instantiates a new DeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceProfileWithDefaults

`func NewDeviceProfileWithDefaults() *DeviceProfile`

NewDeviceProfileWithDefaults instantiates a new DeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetImei

`func (o *DeviceProfile) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *DeviceProfile) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *DeviceProfile) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *DeviceProfile) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetManufacturer

`func (o *DeviceProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DeviceProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetMeid

`func (o *DeviceProfile) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *DeviceProfile) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *DeviceProfile) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *DeviceProfile) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetModel

`func (o *DeviceProfile) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceProfile) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceProfile) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceProfile) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceProfile) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceProfile) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceProfile) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceProfile) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceProfile) GetPlatform() DevicePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceProfile) GetPlatformOk() (*DevicePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceProfile) SetPlatform(v DevicePlatform)`

SetPlatform sets Platform field to given value.


### GetRegistered

`func (o *DeviceProfile) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *DeviceProfile) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *DeviceProfile) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.


### GetSecureHardwarePresent

`func (o *DeviceProfile) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *DeviceProfile) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *DeviceProfile) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *DeviceProfile) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DeviceProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSid

`func (o *DeviceProfile) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *DeviceProfile) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *DeviceProfile) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *DeviceProfile) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetTpmPublicKeyHash

`func (o *DeviceProfile) GetTpmPublicKeyHash() string`

GetTpmPublicKeyHash returns the TpmPublicKeyHash field if non-nil, zero value otherwise.

### GetTpmPublicKeyHashOk

`func (o *DeviceProfile) GetTpmPublicKeyHashOk() (*string, bool)`

GetTpmPublicKeyHashOk returns a tuple with the TpmPublicKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmPublicKeyHash

`func (o *DeviceProfile) SetTpmPublicKeyHash(v string)`

SetTpmPublicKeyHash sets TpmPublicKeyHash field to given value.

### HasTpmPublicKeyHash

`func (o *DeviceProfile) HasTpmPublicKeyHash() bool`

HasTpmPublicKeyHash returns a boolean if a field has been set.

### GetUdid

`func (o *DeviceProfile) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *DeviceProfile) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *DeviceProfile) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *DeviceProfile) HasUdid() bool`

HasUdid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DevicePreRegistrationEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | Device platform for pre-registration | 
**RegistrationGrants** | [**[]RegistrationGrantRequest**](RegistrationGrantRequest.md) | Registration grants to associate with this device | 
**SerialNumber** | **string** | The serial number of the device | 

## Methods

### NewDevicePreRegistrationEntry

`func NewDevicePreRegistrationEntry(platform string, registrationGrants []RegistrationGrantRequest, serialNumber string, ) *DevicePreRegistrationEntry`

NewDevicePreRegistrationEntry instantiates a new DevicePreRegistrationEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePreRegistrationEntryWithDefaults

`func NewDevicePreRegistrationEntryWithDefaults() *DevicePreRegistrationEntry`

NewDevicePreRegistrationEntryWithDefaults instantiates a new DevicePreRegistrationEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *DevicePreRegistrationEntry) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DevicePreRegistrationEntry) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DevicePreRegistrationEntry) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetRegistrationGrants

`func (o *DevicePreRegistrationEntry) GetRegistrationGrants() []RegistrationGrantRequest`

GetRegistrationGrants returns the RegistrationGrants field if non-nil, zero value otherwise.

### GetRegistrationGrantsOk

`func (o *DevicePreRegistrationEntry) GetRegistrationGrantsOk() (*[]RegistrationGrantRequest, bool)`

GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationGrants

`func (o *DevicePreRegistrationEntry) SetRegistrationGrants(v []RegistrationGrantRequest)`

SetRegistrationGrants sets RegistrationGrants field to given value.


### GetSerialNumber

`func (o *DevicePreRegistrationEntry) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DevicePreRegistrationEntry) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DevicePreRegistrationEntry) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



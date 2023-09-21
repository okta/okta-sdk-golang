# UserLockoutSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreventBruteForceLockoutFromUnknownDevices** | Pointer to **bool** | Prevents brute-force lockout from unknown devices for the password authenticator. | [optional] 

## Methods

### NewUserLockoutSettings

`func NewUserLockoutSettings() *UserLockoutSettings`

NewUserLockoutSettings instantiates a new UserLockoutSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLockoutSettingsWithDefaults

`func NewUserLockoutSettingsWithDefaults() *UserLockoutSettings`

NewUserLockoutSettingsWithDefaults instantiates a new UserLockoutSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreventBruteForceLockoutFromUnknownDevices

`func (o *UserLockoutSettings) GetPreventBruteForceLockoutFromUnknownDevices() bool`

GetPreventBruteForceLockoutFromUnknownDevices returns the PreventBruteForceLockoutFromUnknownDevices field if non-nil, zero value otherwise.

### GetPreventBruteForceLockoutFromUnknownDevicesOk

`func (o *UserLockoutSettings) GetPreventBruteForceLockoutFromUnknownDevicesOk() (*bool, bool)`

GetPreventBruteForceLockoutFromUnknownDevicesOk returns a tuple with the PreventBruteForceLockoutFromUnknownDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventBruteForceLockoutFromUnknownDevices

`func (o *UserLockoutSettings) SetPreventBruteForceLockoutFromUnknownDevices(v bool)`

SetPreventBruteForceLockoutFromUnknownDevices sets PreventBruteForceLockoutFromUnknownDevices field to given value.

### HasPreventBruteForceLockoutFromUnknownDevices

`func (o *UserLockoutSettings) HasPreventBruteForceLockoutFromUnknownDevices() bool`

HasPreventBruteForceLockoutFromUnknownDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



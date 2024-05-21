# DeviceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp when device was created | [optional] 
**ManagementStatus** | Pointer to **string** | The management status of the device | [optional] 
**ScreenLockType** | Pointer to **string** | Screen lock type of the device | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewDeviceUser

`func NewDeviceUser() *DeviceUser`

NewDeviceUser instantiates a new DeviceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUserWithDefaults

`func NewDeviceUserWithDefaults() *DeviceUser`

NewDeviceUserWithDefaults instantiates a new DeviceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeviceUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetManagementStatus

`func (o *DeviceUser) GetManagementStatus() string`

GetManagementStatus returns the ManagementStatus field if non-nil, zero value otherwise.

### GetManagementStatusOk

`func (o *DeviceUser) GetManagementStatusOk() (*string, bool)`

GetManagementStatusOk returns a tuple with the ManagementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementStatus

`func (o *DeviceUser) SetManagementStatus(v string)`

SetManagementStatus sets ManagementStatus field to given value.

### HasManagementStatus

`func (o *DeviceUser) HasManagementStatus() bool`

HasManagementStatus returns a boolean if a field has been set.

### GetScreenLockType

`func (o *DeviceUser) GetScreenLockType() string`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *DeviceUser) GetScreenLockTypeOk() (*string, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *DeviceUser) SetScreenLockType(v string)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *DeviceUser) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetUser

`func (o *DeviceUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DeviceUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DeviceUser) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *DeviceUser) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



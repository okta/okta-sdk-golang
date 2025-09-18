# UserDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the device was created | [optional] [readonly] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**DeviceUserId** | Pointer to **string** | Unique key for the user device link | [optional] 

## Methods

### NewUserDevice

`func NewUserDevice() *UserDevice`

NewUserDevice instantiates a new UserDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDeviceWithDefaults

`func NewUserDeviceWithDefaults() *UserDevice`

NewUserDeviceWithDefaults instantiates a new UserDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserDevice) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserDevice) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserDevice) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserDevice) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDevice

`func (o *UserDevice) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserDevice) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserDevice) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserDevice) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceUserId

`func (o *UserDevice) GetDeviceUserId() string`

GetDeviceUserId returns the DeviceUserId field if non-nil, zero value otherwise.

### GetDeviceUserIdOk

`func (o *UserDevice) GetDeviceUserIdOk() (*string, bool)`

GetDeviceUserIdOk returns a tuple with the DeviceUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUserId

`func (o *UserDevice) SetDeviceUserId(v string)`

SetDeviceUserId sets DeviceUserId field to given value.

### HasDeviceUserId

`func (o *UserDevice) HasDeviceUserId() bool`

HasDeviceUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



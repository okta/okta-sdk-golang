# DeviceContextProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the device context provider | [optional] 
**Key** | **string** | Identifies the type of device context provider | 
**UserIdentification** | Pointer to **string** | Whether or not the device context provider is used to identify the user. &#x60;IGNORE&#x60; prevents the device context provider from being used to authenticate the user. Identification of the device and device context collection happens regardless of this setting. | [optional] 

## Methods

### NewDeviceContextProvider

`func NewDeviceContextProvider(key string, ) *DeviceContextProvider`

NewDeviceContextProvider instantiates a new DeviceContextProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceContextProviderWithDefaults

`func NewDeviceContextProviderWithDefaults() *DeviceContextProvider`

NewDeviceContextProviderWithDefaults instantiates a new DeviceContextProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceContextProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceContextProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceContextProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceContextProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *DeviceContextProvider) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeviceContextProvider) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeviceContextProvider) SetKey(v string)`

SetKey sets Key field to given value.


### GetUserIdentification

`func (o *DeviceContextProvider) GetUserIdentification() string`

GetUserIdentification returns the UserIdentification field if non-nil, zero value otherwise.

### GetUserIdentificationOk

`func (o *DeviceContextProvider) GetUserIdentificationOk() (*string, bool)`

GetUserIdentificationOk returns a tuple with the UserIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentification

`func (o *DeviceContextProvider) SetUserIdentification(v string)`

SetUserIdentification sets UserIdentification field to given value.

### HasUserIdentification

`func (o *DeviceContextProvider) HasUserIdentification() bool`

HasUserIdentification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeviceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**OktaVerifyPayload**](OktaVerifyPayload.md) |  | [optional] 
**PayloadVersion** | Pointer to **string** | The version of the payload schema | [optional] 
**Provider** | Pointer to **string** | The name of the provider | [optional] 

## Methods

### NewDeviceProvider

`func NewDeviceProvider() *DeviceProvider`

NewDeviceProvider instantiates a new DeviceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceProviderWithDefaults

`func NewDeviceProviderWithDefaults() *DeviceProvider`

NewDeviceProviderWithDefaults instantiates a new DeviceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *DeviceProvider) GetPayload() OktaVerifyPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DeviceProvider) GetPayloadOk() (*OktaVerifyPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DeviceProvider) SetPayload(v OktaVerifyPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DeviceProvider) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPayloadVersion

`func (o *DeviceProvider) GetPayloadVersion() string`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *DeviceProvider) GetPayloadVersionOk() (*string, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *DeviceProvider) SetPayloadVersion(v string)`

SetPayloadVersion sets PayloadVersion field to given value.

### HasPayloadVersion

`func (o *DeviceProvider) HasPayloadVersion() bool`

HasPayloadVersion returns a boolean if a field has been set.

### GetProvider

`func (o *DeviceProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DeviceProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DeviceProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DeviceProvider) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



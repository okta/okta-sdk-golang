# EventHookChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**EventHookChannelConfig**](EventHookChannelConfig.md) |  | [optional] 
**Type** | Pointer to [**EventHookChannelType**](EventHookChannelType.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewEventHookChannel

`func NewEventHookChannel() *EventHookChannel`

NewEventHookChannel instantiates a new EventHookChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookChannelWithDefaults

`func NewEventHookChannelWithDefaults() *EventHookChannel`

NewEventHookChannelWithDefaults instantiates a new EventHookChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *EventHookChannel) GetConfig() EventHookChannelConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EventHookChannel) GetConfigOk() (*EventHookChannelConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EventHookChannel) SetConfig(v EventHookChannelConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EventHookChannel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetType

`func (o *EventHookChannel) GetType() EventHookChannelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventHookChannel) GetTypeOk() (*EventHookChannelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventHookChannel) SetType(v EventHookChannelType)`

SetType sets Type field to given value.

### HasType

`func (o *EventHookChannel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *EventHookChannel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventHookChannel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventHookChannel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EventHookChannel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



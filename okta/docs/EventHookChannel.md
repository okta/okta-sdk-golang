# EventHookChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**EventHookChannelConfig**](EventHookChannelConfig.md) |  | 
**Type** | **string** | The channel type. Currently supports &#x60;HTTP&#x60;. | 
**Version** | **string** | Version of the channel. Currently the only supported version is &#x60;1.0.0&#x60;&#x60;. | 

## Methods

### NewEventHookChannel

`func NewEventHookChannel(config EventHookChannelConfig, type_ string, version string, ) *EventHookChannel`

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


### GetType

`func (o *EventHookChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventHookChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventHookChannel) SetType(v string)`

SetType sets Type field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



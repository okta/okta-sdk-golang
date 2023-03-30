# EventHookChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**EventHookChannelConfigAuthScheme**](EventHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]EventHookChannelConfigHeader**](EventHookChannelConfigHeader.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewEventHookChannelConfig

`func NewEventHookChannelConfig() *EventHookChannelConfig`

NewEventHookChannelConfig instantiates a new EventHookChannelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookChannelConfigWithDefaults

`func NewEventHookChannelConfigWithDefaults() *EventHookChannelConfig`

NewEventHookChannelConfigWithDefaults instantiates a new EventHookChannelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *EventHookChannelConfig) GetAuthScheme() EventHookChannelConfigAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *EventHookChannelConfig) GetAuthSchemeOk() (*EventHookChannelConfigAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *EventHookChannelConfig) SetAuthScheme(v EventHookChannelConfigAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *EventHookChannelConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetHeaders

`func (o *EventHookChannelConfig) GetHeaders() []EventHookChannelConfigHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EventHookChannelConfig) GetHeadersOk() (*[]EventHookChannelConfigHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EventHookChannelConfig) SetHeaders(v []EventHookChannelConfigHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EventHookChannelConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUri

`func (o *EventHookChannelConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *EventHookChannelConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *EventHookChannelConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *EventHookChannelConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# EventHookChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**EventHookChannelConfigAuthScheme**](EventHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]EventHookChannelConfigHeader**](EventHookChannelConfigHeader.md) | Optional list of key/value pairs for headers that can be sent with the request to the external service. For example, &#x60;X-Other-Header&#x60; is an example of an optional header, with a value of &#x60;my-header-value&#x60;, that you want Okta to pass to your external service. | [optional] 
**Method** | Pointer to **string** | The method of the Okta event hook request | [optional] [readonly] 
**Uri** | **string** | The external service endpoint called to execute the event hook handler | 

## Methods

### NewEventHookChannelConfig

`func NewEventHookChannelConfig(uri string, ) *EventHookChannelConfig`

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

### GetMethod

`func (o *EventHookChannelConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EventHookChannelConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EventHookChannelConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *EventHookChannelConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InlineHookChannelConfigCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) | An optional list of key/value pairs for headers that you can send with the request to the external service. | [optional] 
**Method** | Pointer to **string** | The method of the Okta inline hook request | [optional] 
**Uri** | Pointer to **string** | The external service endpoint that executes the inline hook handler. It must begin with &#x60;https://&#x60; and be reachable by Okta. No white space is allowed in the URI. | [optional] 

## Methods

### NewInlineHookChannelConfigCreate

`func NewInlineHookChannelConfigCreate() *InlineHookChannelConfigCreate`

NewInlineHookChannelConfigCreate instantiates a new InlineHookChannelConfigCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookChannelConfigCreateWithDefaults

`func NewInlineHookChannelConfigCreateWithDefaults() *InlineHookChannelConfigCreate`

NewInlineHookChannelConfigCreateWithDefaults instantiates a new InlineHookChannelConfigCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *InlineHookChannelConfigCreate) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookChannelConfigCreate) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookChannelConfigCreate) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookChannelConfigCreate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookChannelConfigCreate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookChannelConfigCreate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookChannelConfigCreate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookChannelConfigCreate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookChannelConfigCreate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookChannelConfigCreate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookChannelConfigCreate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookChannelConfigCreate) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



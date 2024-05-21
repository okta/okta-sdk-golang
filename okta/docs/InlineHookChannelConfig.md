# InlineHookChannelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**InlineHookChannelConfigAuthScheme**](InlineHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineHookChannelConfig

`func NewInlineHookChannelConfig() *InlineHookChannelConfig`

NewInlineHookChannelConfig instantiates a new InlineHookChannelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookChannelConfigWithDefaults

`func NewInlineHookChannelConfigWithDefaults() *InlineHookChannelConfig`

NewInlineHookChannelConfigWithDefaults instantiates a new InlineHookChannelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *InlineHookChannelConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookChannelConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookChannelConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookChannelConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineHookChannelConfig) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookChannelConfig) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookChannelConfig) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookChannelConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookChannelConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookChannelConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookChannelConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookChannelConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookChannelConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookChannelConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookChannelConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookChannelConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



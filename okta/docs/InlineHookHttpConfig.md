# InlineHookHttpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**NullableInlineHookChannelConfigAuthSchemeResponse**](InlineHookChannelConfigAuthSchemeResponse.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) | An optional list of key/value pairs for headers that you can send with the request to the external service | [optional] 
**Method** | Pointer to **string** | The method of the Okta inline hook request | [optional] 
**Uri** | Pointer to **string** | The external service endpoint that executes the inline hook handler. It must begin with &#x60;https://&#x60; and be reachable by Okta. No white space is allowed in the URI. | [optional] 

## Methods

### NewInlineHookHttpConfig

`func NewInlineHookHttpConfig() *InlineHookHttpConfig`

NewInlineHookHttpConfig instantiates a new InlineHookHttpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookHttpConfigWithDefaults

`func NewInlineHookHttpConfigWithDefaults() *InlineHookHttpConfig`

NewInlineHookHttpConfigWithDefaults instantiates a new InlineHookHttpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *InlineHookHttpConfig) GetAuthScheme() InlineHookChannelConfigAuthSchemeResponse`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookHttpConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthSchemeResponse, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookHttpConfig) SetAuthScheme(v InlineHookChannelConfigAuthSchemeResponse)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookHttpConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### SetAuthSchemeNil

`func (o *InlineHookHttpConfig) SetAuthSchemeNil(b bool)`

 SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil

### UnsetAuthScheme
`func (o *InlineHookHttpConfig) UnsetAuthScheme()`

UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
### GetHeaders

`func (o *InlineHookHttpConfig) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookHttpConfig) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookHttpConfig) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookHttpConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookHttpConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookHttpConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookHttpConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookHttpConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookHttpConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookHttpConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookHttpConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookHttpConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



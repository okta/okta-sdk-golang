# InlineHookHttpConfigCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to [**NullableInlineHookChannelConfigAuthSchemeBody**](InlineHookChannelConfigAuthSchemeBody.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) | An optional list of key/value pairs for headers that you can send with the request to the external service. | [optional] 
**Method** | Pointer to **string** | The method of the Okta inline hook request | [optional] 
**Uri** | Pointer to **string** | The external service endpoint that executes the inline hook handler. It must begin with &#x60;https://&#x60; and be reachable by Okta. No white space is allowed in the URI. | [optional] 

## Methods

### NewInlineHookHttpConfigCreate

`func NewInlineHookHttpConfigCreate() *InlineHookHttpConfigCreate`

NewInlineHookHttpConfigCreate instantiates a new InlineHookHttpConfigCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookHttpConfigCreateWithDefaults

`func NewInlineHookHttpConfigCreateWithDefaults() *InlineHookHttpConfigCreate`

NewInlineHookHttpConfigCreateWithDefaults instantiates a new InlineHookHttpConfigCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthScheme

`func (o *InlineHookHttpConfigCreate) GetAuthScheme() InlineHookChannelConfigAuthSchemeBody`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookHttpConfigCreate) GetAuthSchemeOk() (*InlineHookChannelConfigAuthSchemeBody, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookHttpConfigCreate) SetAuthScheme(v InlineHookChannelConfigAuthSchemeBody)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookHttpConfigCreate) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### SetAuthSchemeNil

`func (o *InlineHookHttpConfigCreate) SetAuthSchemeNil(b bool)`

 SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil

### UnsetAuthScheme
`func (o *InlineHookHttpConfigCreate) UnsetAuthScheme()`

UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
### GetHeaders

`func (o *InlineHookHttpConfigCreate) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookHttpConfigCreate) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookHttpConfigCreate) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookHttpConfigCreate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookHttpConfigCreate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookHttpConfigCreate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookHttpConfigCreate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookHttpConfigCreate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookHttpConfigCreate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookHttpConfigCreate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookHttpConfigCreate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookHttpConfigCreate) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



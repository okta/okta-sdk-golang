# InlineHookChannelConfigAuthSchemeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The header name for the authorization server | [optional] 
**Type** | Pointer to **string** | The authentication scheme type. Supported type&amp;mdash;&#x60;HEADER&#x60;. | [optional] 
**Value** | Pointer to **string** | The header value. This secret value is passed to your external service endpoint. Your external service can check it as a security measure. | [optional] 

## Methods

### NewInlineHookChannelConfigAuthSchemeBody

`func NewInlineHookChannelConfigAuthSchemeBody() *InlineHookChannelConfigAuthSchemeBody`

NewInlineHookChannelConfigAuthSchemeBody instantiates a new InlineHookChannelConfigAuthSchemeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookChannelConfigAuthSchemeBodyWithDefaults

`func NewInlineHookChannelConfigAuthSchemeBodyWithDefaults() *InlineHookChannelConfigAuthSchemeBody`

NewInlineHookChannelConfigAuthSchemeBodyWithDefaults instantiates a new InlineHookChannelConfigAuthSchemeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *InlineHookChannelConfigAuthSchemeBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InlineHookChannelConfigAuthSchemeBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InlineHookChannelConfigAuthSchemeBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InlineHookChannelConfigAuthSchemeBody) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *InlineHookChannelConfigAuthSchemeBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineHookChannelConfigAuthSchemeBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineHookChannelConfigAuthSchemeBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineHookChannelConfigAuthSchemeBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InlineHookChannelConfigAuthSchemeBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineHookChannelConfigAuthSchemeBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineHookChannelConfigAuthSchemeBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineHookChannelConfigAuthSchemeBody) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



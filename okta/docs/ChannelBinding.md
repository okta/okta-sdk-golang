# ChannelBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to [**RequiredEnum**](RequiredEnum.md) |  | [optional] 
**Style** | Pointer to **string** |  | [optional] 

## Methods

### NewChannelBinding

`func NewChannelBinding() *ChannelBinding`

NewChannelBinding instantiates a new ChannelBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelBindingWithDefaults

`func NewChannelBindingWithDefaults() *ChannelBinding`

NewChannelBindingWithDefaults instantiates a new ChannelBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *ChannelBinding) GetRequired() RequiredEnum`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ChannelBinding) GetRequiredOk() (*RequiredEnum, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ChannelBinding) SetRequired(v RequiredEnum)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ChannelBinding) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetStyle

`func (o *ChannelBinding) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ChannelBinding) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ChannelBinding) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ChannelBinding) HasStyle() bool`

HasStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



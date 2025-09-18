# InlineHookReplace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**InlineHookChannelCreate**](InlineHookChannelCreate.md) |  | [optional] 
**Name** | Pointer to **string** | The display name of the inline hook | [optional] 
**Version** | Pointer to **string** | Version of the inline hook type. The currently supported version is &#x60;1.0.0&#x60;. | [optional] 

## Methods

### NewInlineHookReplace

`func NewInlineHookReplace() *InlineHookReplace`

NewInlineHookReplace instantiates a new InlineHookReplace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookReplaceWithDefaults

`func NewInlineHookReplaceWithDefaults() *InlineHookReplace`

NewInlineHookReplaceWithDefaults instantiates a new InlineHookReplace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InlineHookReplace) GetChannel() InlineHookChannelCreate`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineHookReplace) GetChannelOk() (*InlineHookChannelCreate, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineHookReplace) SetChannel(v InlineHookChannelCreate)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineHookReplace) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetName

`func (o *InlineHookReplace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineHookReplace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineHookReplace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineHookReplace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *InlineHookReplace) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineHookReplace) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineHookReplace) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineHookReplace) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



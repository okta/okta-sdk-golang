# InlineHookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**InlineHookChannelCreate**](InlineHookChannelCreate.md) |  | [optional] 
**Name** | Pointer to **string** | The display name of the inline hook | [optional] 
**Type** | Pointer to **string** | One of the inline hook types | [optional] 
**Version** | Pointer to **string** | Version of the inline hook type. The currently supported version is &#x60;1.0.0&#x60;. | [optional] 

## Methods

### NewInlineHookCreate

`func NewInlineHookCreate() *InlineHookCreate`

NewInlineHookCreate instantiates a new InlineHookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookCreateWithDefaults

`func NewInlineHookCreateWithDefaults() *InlineHookCreate`

NewInlineHookCreateWithDefaults instantiates a new InlineHookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InlineHookCreate) GetChannel() InlineHookChannelCreate`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineHookCreate) GetChannelOk() (*InlineHookChannelCreate, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineHookCreate) SetChannel(v InlineHookChannelCreate)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineHookCreate) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetName

`func (o *InlineHookCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineHookCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineHookCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineHookCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *InlineHookCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineHookCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineHookCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineHookCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *InlineHookCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineHookCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineHookCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineHookCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



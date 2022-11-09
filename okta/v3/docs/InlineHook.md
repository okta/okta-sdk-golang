# InlineHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**InlineHookChannel**](InlineHookChannel.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**InlineHookStatus**](InlineHookStatus.md) |  | [optional] 
**Type** | Pointer to [**InlineHookType**](InlineHookType.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewInlineHook

`func NewInlineHook() *InlineHook`

NewInlineHook instantiates a new InlineHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookWithDefaults

`func NewInlineHookWithDefaults() *InlineHook`

NewInlineHookWithDefaults instantiates a new InlineHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InlineHook) GetChannel() InlineHookChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineHook) GetChannelOk() (*InlineHookChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineHook) SetChannel(v InlineHookChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineHook) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreated

`func (o *InlineHook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InlineHook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InlineHook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InlineHook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *InlineHook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineHook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineHook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineHook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineHook) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineHook) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineHook) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineHook) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *InlineHook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineHook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineHook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineHook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *InlineHook) GetStatus() InlineHookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineHook) GetStatusOk() (*InlineHookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineHook) SetStatus(v InlineHookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineHook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *InlineHook) GetType() InlineHookType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineHook) GetTypeOk() (*InlineHookType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineHook) SetType(v InlineHookType)`

SetType sets Type field to given value.

### HasType

`func (o *InlineHook) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *InlineHook) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineHook) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineHook) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineHook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *InlineHook) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineHook) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineHook) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineHook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



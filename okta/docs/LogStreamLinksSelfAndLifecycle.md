# LogStreamLinksSelfAndLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activate** | Pointer to [**LogStreamActivateLink**](LogStreamActivateLink.md) |  | [optional] 
**Deactivate** | Pointer to [**LogStreamDeactivateLink**](LogStreamDeactivateLink.md) |  | [optional] 
**Self** | [**LogStreamSelfLink**](LogStreamSelfLink.md) |  | 

## Methods

### NewLogStreamLinksSelfAndLifecycle

`func NewLogStreamLinksSelfAndLifecycle(self LogStreamSelfLink, ) *LogStreamLinksSelfAndLifecycle`

NewLogStreamLinksSelfAndLifecycle instantiates a new LogStreamLinksSelfAndLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamLinksSelfAndLifecycleWithDefaults

`func NewLogStreamLinksSelfAndLifecycleWithDefaults() *LogStreamLinksSelfAndLifecycle`

NewLogStreamLinksSelfAndLifecycleWithDefaults instantiates a new LogStreamLinksSelfAndLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivate

`func (o *LogStreamLinksSelfAndLifecycle) GetActivate() LogStreamActivateLink`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *LogStreamLinksSelfAndLifecycle) GetActivateOk() (*LogStreamActivateLink, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *LogStreamLinksSelfAndLifecycle) SetActivate(v LogStreamActivateLink)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *LogStreamLinksSelfAndLifecycle) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *LogStreamLinksSelfAndLifecycle) GetDeactivate() LogStreamDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *LogStreamLinksSelfAndLifecycle) GetDeactivateOk() (*LogStreamDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *LogStreamLinksSelfAndLifecycle) SetDeactivate(v LogStreamDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *LogStreamLinksSelfAndLifecycle) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetSelf

`func (o *LogStreamLinksSelfAndLifecycle) GetSelf() LogStreamSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LogStreamLinksSelfAndLifecycle) GetSelfOk() (*LogStreamSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LogStreamLinksSelfAndLifecycle) SetSelf(v LogStreamSelfLink)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



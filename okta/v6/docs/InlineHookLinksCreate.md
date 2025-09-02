# InlineHookLinksCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObject**](HrefObject.md) | URL to deactivate the inline hook | [optional] 
**Execute** | Pointer to [**HrefObject**](HrefObject.md) | URL to test the inline hook | [optional] 

## Methods

### NewInlineHookLinksCreate

`func NewInlineHookLinksCreate() *InlineHookLinksCreate`

NewInlineHookLinksCreate instantiates a new InlineHookLinksCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookLinksCreateWithDefaults

`func NewInlineHookLinksCreateWithDefaults() *InlineHookLinksCreate`

NewInlineHookLinksCreateWithDefaults instantiates a new InlineHookLinksCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *InlineHookLinksCreate) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *InlineHookLinksCreate) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *InlineHookLinksCreate) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *InlineHookLinksCreate) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetDeactivate

`func (o *InlineHookLinksCreate) GetDeactivate() HrefObject`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *InlineHookLinksCreate) GetDeactivateOk() (*HrefObject, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *InlineHookLinksCreate) SetDeactivate(v HrefObject)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *InlineHookLinksCreate) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetExecute

`func (o *InlineHookLinksCreate) GetExecute() HrefObject`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *InlineHookLinksCreate) GetExecuteOk() (*HrefObject, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *InlineHookLinksCreate) SetExecute(v HrefObject)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *InlineHookLinksCreate) HasExecute() bool`

HasExecute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


